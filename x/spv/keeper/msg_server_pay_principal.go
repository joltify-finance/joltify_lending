package keeper

import (
	"context"
	"time"

	sdkmath "cosmossdk.io/math"

	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k msgServer) calulatetotalDueInterest(ctx sdk.Context, poolInfo types.PoolInfo) (sdkmath.Int, error) {

	totalAmount := sdk.ZeroInt()
	for _, el := range poolInfo.PoolNFTIds {
		class, found := k.nftKeeper.GetClass(ctx, el)
		if !found {
			panic(found)
		}
		var borrowInterest types.BorrowInterest
		err := proto.Unmarshal(class.Data.Value, &borrowInterest)
		if err != nil {
			return sdkmath.ZeroInt(), coserrors.Wrapf(err, "invalid unmarshal of the borrow interest")
		}
		paymentAmount := borrowInterest.MonthlyRatio.Mul(sdk.NewDecFromInt(poolInfo.BorrowedAmount.Amount)).TruncateInt()
		totalAmount = totalAmount.Add(paymentAmount)
	}
	return totalAmount, nil
}

func (k msgServer) PayPrincipal(goCtx context.Context, msg *types.MsgPayPrincipal) (*types.MsgPayPrincipalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	spv, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	poolInfo, found := k.GetPools(ctx, msg.GetPoolIndex())
	if !found {
		return nil, coserrors.Wrapf(sdkerrors.ErrNotFound, "pool cannot be found %v", msg.GetPoolIndex())
	}

	if poolInfo.PoolStatus != types.PoolInfo_ACTIVE {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "pool is not active")
	}

	// we do not allow the spv to pay principal at the time fram [due-3*withdrawWindow, due]
	dueDate := poolInfo.ProjectDueTime
	secondTimeStampBeforeProjectDueDate := dueDate.Add(-time.Second * time.Duration(poolInfo.WithdrawRequestWindowSeconds*3+10))

	currentTime := ctx.BlockTime()
	if currentTime.After(secondTimeStampBeforeProjectDueDate) && currentTime.Before(dueDate) {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "%v: principal can not be paid between %v <-> %v", currentTime, secondTimeStampBeforeProjectDueDate, dueDate)
	}

	if msg.Token.Denom != poolInfo.BorrowedAmount.Denom {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid token demo, want %v", poolInfo.BorrowedAmount.Denom)
	}

	if !spv.Equals(poolInfo.OwnerAddress) {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "only pool owner can pay the principal")
	}

	interestPaymentAmount, err := k.calulatetotalDueInterest(ctx, poolInfo)
	if err != nil {
		return nil, err
	}
	// to assure the interest is enough, we ask the spv to pay at least the 1.3 interest ahead, will be refund if it has leftover
	a, _ := denomConvertToLocalAndUsd(poolInfo.BorrowedAmount.Denom)
	paymentAmountUsd, _, err := k.outboundConvertToUSDWithMarketID(ctx, a, interestPaymentAmount)
	if err != nil {
		return nil, coserrors.Wrapf(err, "fail to convert to USD")
	}

	paymentAmountUsd = sdk.NewDecFromInt(paymentAmountUsd).Mul(sdk.NewDecWithPrec(13, 1)).TruncateInt()
	if poolInfo.EscrowInterestAmount.LT(paymentAmountUsd) {
		return nil, coserrors.Wrapf(types.ErrInsufficientFund, "not enough interest to be paid to close the pool, at least %v is needed", interestPaymentAmount)
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, spv, types.ModuleAccount, sdk.NewCoins(msg.Token))
	if err != nil {
		return nil, coserrors.Wrapf(err, "fail to pay all the principal")
	}

	poolInfo.EscrowPrincipalAmount = poolInfo.EscrowPrincipalAmount.Add(msg.Token)

	a, _ = denomConvertToLocalAndUsd(poolInfo.BorrowedAmount.Denom)
	principalEscrowAmountLocal, ratio, err := k.inboundConvertFromUSDWithMarketID(ctx, a, poolInfo.EscrowPrincipalAmount.Amount)
	if err != nil {
		return nil, coserrors.Wrapf(err, "fail to convert to USD")
	}

	// we only close the pool when the escrow principal is later than the total borrowed and the project pass the project length
	if principalEscrowAmountLocal.GTE(poolInfo.BorrowedAmount.Amount) && ctx.BlockTime().After(poolInfo.ProjectDueTime) {
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
	b := outboundConvertToUSD(poolInfo.BorrowedAmount.Amount, ratio)
	return &types.MsgPayPrincipalResponse{}, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "principal is not fully paid. you have paid %v and borrowed %v", poolInfo.EscrowPrincipalAmount, b)

}

func (k msgServer) PayPrincipalForWithdrawalRequests(goCtx context.Context, msg *types.MsgPayPrincipalPartial) (*types.MsgPayPrincipalPartialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	spv, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	poolInfo, found := k.GetPools(ctx, msg.GetPoolIndex())
	if !found {
		return nil, coserrors.Wrapf(sdkerrors.ErrNotFound, "pool cannot be found %v", msg.GetPoolIndex())
	}

	if poolInfo.WithdrawProposalAmount.IsZero() {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "no withdraw proposal to be paid")
	}

	if poolInfo.PoolStatus != types.PoolInfo_ACTIVE {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "pool is not active")
	}

	if !spv.Equals(poolInfo.OwnerAddress) {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "only pool owner can pay the principal")
	}

	paymentAmount, err := k.calulatetotalDueInterest(ctx, poolInfo)
	if err != nil {
		return nil, err
	}

	// to assure the interest is enough, we ask the spv to pay at least the 1.3 interest ahead, will be refund if it has leftover
	a, _ := denomConvertToLocalAndUsd(poolInfo.BorrowedAmount.Denom)
	paymentAmountUsd, _, err := k.outboundConvertToUSDWithMarketID(ctx, a, paymentAmount)
	if err != nil {
		return nil, coserrors.Wrapf(err, "fail to convert to USD")
	}

	paymentAmountUsd = sdk.NewDecFromInt(paymentAmountUsd).Mul(sdk.NewDecWithPrec(13, 1)).TruncateInt()
	if poolInfo.EscrowInterestAmount.LT(paymentAmountUsd) {
		return nil, coserrors.Wrapf(types.ErrInsufficientFund, "not enough interest to be paid to close the pool, at least %v is needed", paymentAmount)
	}

	if poolInfo.EscrowInterestAmount.LT(paymentAmountUsd) {
		return nil, coserrors.Wrapf(types.ErrInsufficientFund, "not enough interest to be paid to close the pool, at least %v is needed", paymentAmount)
	}

	currentTime := ctx.BlockTime()
	dueDate := poolInfo.ProjectDueTime
	secondTimeStampBeforeProjectDueDate := dueDate.Add(-time.Second * time.Duration(poolInfo.WithdrawRequestWindowSeconds*2-10))

	condition := currentTime.After(secondTimeStampBeforeProjectDueDate) && currentTime.Before(dueDate)
	if !condition {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "principal can only  be paid between %v <-> %v", secondTimeStampBeforeProjectDueDate, dueDate)
	}

	if msg.Token.Denom != poolInfo.BorrowedAmount.Denom {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid token demo, want %v", poolInfo.BorrowedAmount.Denom)
	}

	poolInfo.EscrowPrincipalAmount = poolInfo.EscrowPrincipalAmount.Add(msg.Token)

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, spv, types.ModuleAccount, sdk.NewCoins(msg.Token))
	if err != nil {
		return nil, coserrors.Wrapf(err, "fail to pay all the principal")
	}

	a, _ = denomConvertToLocalAndUsd(poolInfo.BorrowedAmount.Denom)
	principalEscrowAmountLocal, ratio, err := k.inboundConvertFromUSDWithMarketID(ctx, a, poolInfo.EscrowPrincipalAmount.Amount)
	if err != nil {
		return nil, coserrors.Wrapf(err, "fail to convert to USD")
	}

	if principalEscrowAmountLocal.LT(poolInfo.WithdrawProposalAmount.Amount) {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "you must pay at least %v( current amount %v)", poolInfo.WithdrawProposalAmount, principalEscrowAmountLocal)
	}

	poolInfo.PrincipalWithdrawalRequestPaymentRatio = ratio
	k.SetPool(ctx, poolInfo)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypePayPrincipal,
			sdk.NewAttribute(types.AttributeCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeAmount, msg.Token.String()),
		),
	)

	return &types.MsgPayPrincipalPartialResponse{}, nil
}
