package keeper

import (
	"context"
	"time"

	coserrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k msgServer) calculateTotalDueInterest(ctx context.Context, poolInfo types.PoolInfo) (sdkmath.Int, error) {
	totalAmount := sdk.ZeroInt()
	for _, el := range poolInfo.PoolNFTIds {
		class, found := k.NftKeeper.GetClass(ctx, el)
		if !found {
			panic(found)
		}
		var borrowInterest types.BorrowInterest
		err := proto.Unmarshal(class.Data.Value, &borrowInterest)
		if err != nil {
			return sdkmath.ZeroInt(), coserrors.Wrapf(err, "invalid unmarshal of the borrow interest")
		}
		borrow := borrowInterest.BorrowDetails[len(borrowInterest.BorrowDetails)-1].BorrowedAmount
		paymentAmount := borrowInterest.MonthlyRatio.Mul(sdk.NewDecFromInt(borrow.Amount)).TruncateInt()
		totalAmount = totalAmount.Add(paymentAmount)
	}
	return totalAmount, nil
}

func (k msgServer) PayPrincipal(goCtx context.Context, msg *types.MsgPayPrincipal) (*types.MsgPayPrincipalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ctx = ctx.WithGasMeter(sdk.NewInfiniteGasMeter())

	spv, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	poolInfo, found := k.GetPools(ctx, msg.GetPoolIndex())
	if !found {
		return nil, coserrors.Wrapf(errorsmod.ErrNotFound, "pool cannot be found %v", msg.GetPoolIndex())
	}

	if poolInfo.PrincipalPaid {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidRequest, "principal has been paid")
	}

	if poolInfo.PoolStatus != types.PoolInfo_ACTIVE {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidRequest, "pool is not active current status is %v", poolInfo.PoolStatus)
	}

	// we do not allow the spv to pay principal at the time fram [due-3*withdrawWindow, due]
	dueDate := poolInfo.ProjectDueTime
	secondTimeStampBeforeProjectDueDate := dueDate.Add(-time.Second * time.Duration(poolInfo.WithdrawRequestWindowSeconds*3+10))

	currentTime := ctx.BlockTime()
	if currentTime.After(secondTimeStampBeforeProjectDueDate) && currentTime.Before(dueDate) {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidRequest, "%v: principal can not be paid between %v <-> %v", currentTime, secondTimeStampBeforeProjectDueDate, dueDate)
	}

	if msg.Token.Denom != poolInfo.TargetAmount.Denom {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidRequest, "invalid token demo, want %v", poolInfo.BorrowedAmount.Denom)
	}

	if !spv.Equals(poolInfo.OwnerAddress) {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidRequest, "only pool owner can pay the principal")
	}

	if poolInfo.InterestPrepayment == nil {
		return nil, coserrors.Wrapf(types.ErrInvalidParameter, "you need to pay interest firstly")
	}

	if poolInfo.EscrowInterestAmount.IsNegative() {
		panic("if the interest prepayment is not nil, the escrow interest amount should not be negative")
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, spv, types.ModuleAccount, sdk.NewCoins(msg.Token))
	if err != nil {
		return nil, coserrors.Wrapf(err, "fail to pay all the principal")
	}

	poolInfo.EscrowPrincipalAmount = poolInfo.EscrowPrincipalAmount.Add(msg.Token)

	a, _ := denomConvertToLocalAndUsd(poolInfo.BorrowedAmount.Denom)
	principalEscrowAmountLocal, ratio, err := k.inboundConvertFromUSDWithMarketID(ctx, denomConvertToMarketID(a), poolInfo.EscrowPrincipalAmount.Amount)
	if err != nil {
		return nil, coserrors.Wrapf(err, "fail to convert to USD with market id %v", denomConvertToMarketID(a))
	}

	// we only close the pool when the escrow principal is later than the total borrowed and the project pass the project length
	if principalEscrowAmountLocal.GTE(poolInfo.BorrowedAmount.Amount) {
		// once we are in the freezing state, the usable amount will not be accurate any longer
		poolInfo.PoolStatus = types.PoolInfo_FREEZING
		poolInfo.PrincipalPaymentExchangeRatio = ratio

		k.SetPool(ctx, poolInfo)
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypePayPrincipal,
				sdk.NewAttribute(types.AttributeCreator, msg.Creator),
				sdk.NewAttribute(types.AttributeAmount, msg.Token.String()),
			),
		)

		return &types.MsgPayPrincipalResponse{}, nil
	}
	return &types.MsgPayPrincipalResponse{}, coserrors.Wrapf(errorsmod.ErrInvalidRequest, "principal is not fully paid. you have paid %v and borrowed %v", principalEscrowAmountLocal, outboundConvertToUSD(poolInfo.BorrowedAmount.Amount, ratio))
}

func (k msgServer) PayPrincipalForWithdrawalRequests(goCtx context.Context, msg *types.MsgPayPrincipalPartial) (*types.MsgPayPrincipalPartialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	spv, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	poolInfo, found := k.GetPools(ctx, msg.GetPoolIndex())
	if !found {
		return nil, coserrors.Wrapf(errorsmod.ErrNotFound, "pool cannot be found %v", msg.GetPoolIndex())
	}

	if poolInfo.PrincipalPaid {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidRequest, "principal has been paid")
	}

	if poolInfo.WithdrawProposalAmount.IsZero() {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidRequest, "no withdraw proposal to be paid")
	}

	if poolInfo.PoolStatus != types.PoolInfo_PooLPayPartially {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidRequest, "pool is not in request to pay partiallly status current status is %v", poolInfo.PoolStatus)
	}

	if !spv.Equals(poolInfo.OwnerAddress) {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidRequest, "only pool owner can pay the principal")
	}

	if poolInfo.InterestPrepayment == nil {
		return nil, coserrors.Wrapf(types.ErrInsufficientFund, "not enough interest to be paid to close the pool")
	}

	if poolInfo.EscrowInterestAmount.IsNegative() {
		panic("if the interest prepayment is not nil, the escrow interest amount should not be negative")
	}

	currentTime := ctx.BlockTime()
	dueDate := poolInfo.ProjectDueTime
	secondTimeStampBeforeProjectDueDate := dueDate.Add(-time.Second * time.Duration(poolInfo.WithdrawRequestWindowSeconds*2-10))

	condition := currentTime.After(secondTimeStampBeforeProjectDueDate) && currentTime.Before(dueDate)
	if !condition {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidRequest, "principal can only  be paid between %v <-> %v (current time %v)", secondTimeStampBeforeProjectDueDate, dueDate, currentTime)
	}

	if msg.Token.Denom != poolInfo.TargetAmount.Denom {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidRequest, "invalid token demo, want %v", poolInfo.BorrowedAmount.Denom)
	}

	poolInfo.EscrowPrincipalAmount = poolInfo.EscrowPrincipalAmount.Add(msg.Token)

	a, _ := denomConvertToLocalAndUsd(poolInfo.WithdrawProposalAmount.Denom)
	withdrawProposalAmountUsd, ratio, err := k.outboundConvertToUSDWithMarketID(ctx, denomConvertToMarketID(a), poolInfo.WithdrawProposalAmount.Amount)
	if err != nil {
		return nil, coserrors.Wrapf(err, "fail to convert to USD with market id %v", denomConvertToMarketID(a))
	}

	if poolInfo.EscrowPrincipalAmount.Amount.LT(withdrawProposalAmountUsd) {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidRequest, "you must pay at least %v( current amount in escrow %v)", withdrawProposalAmountUsd, poolInfo.EscrowPrincipalAmount)
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, spv, types.ModuleAccount, sdk.NewCoins(msg.Token))
	if err != nil {
		return nil, coserrors.Wrapf(err, "fail to pay all the principal")
	}

	poolInfo.PrincipalWithdrawalRequestPaymentRatio = ratio
	k.SetPool(ctx, poolInfo)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypePayPrincipalPartial,
			sdk.NewAttribute(types.AttributeCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeAmount, msg.Token.String()),
		),
	)

	return &types.MsgPayPrincipalPartialResponse{}, nil
}
