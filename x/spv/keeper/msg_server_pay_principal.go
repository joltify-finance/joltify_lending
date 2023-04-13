package keeper

import (
	"context"
	sdkmath "cosmossdk.io/math"
	"time"

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

	exchangeRatio, err := sdk.NewDecFromStr(msg.ExchangeRatio)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid exchange ratio %v", msg.ExchangeRatio)
	}

	poolInfo, found := k.GetPools(ctx, msg.GetPoolIndex())
	if !found {
		return nil, coserrors.Wrapf(sdkerrors.ErrNotFound, "pool cannot be found %v", msg.GetPoolIndex())
	}

	if poolInfo.PrincipalPaid {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "principal already paid")
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

	adjAmount := sdk.NewCoin(msg.Token.Denom, exchangeRatio.MulInt(msg.Token.Amount).TruncateInt())
	if !adjAmount.Equal(poolInfo.BorrowedAmount) {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "you must pay exact full principal %v(adjust amount %v)", poolInfo.BorrowedAmount, adjAmount)
	}

	if !spv.Equals(poolInfo.OwnerAddress) {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "only pool owner can pay the principal")
	}

	paymentAmount, err := k.calulatetotalDueInterest(ctx, poolInfo)
	if err != nil {
		return nil, err
	}

	if poolInfo.EscrowInterestAmount.LT(paymentAmount) {
		return nil, coserrors.Wrapf(types.ErrInsufficientFund, "not enough interest to be paid to close the pool, at least %v is needed", paymentAmount)
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, spv, types.ModuleAccount, sdk.NewCoins(msg.Token))
	if err != nil {
		return nil, coserrors.Wrapf(err, "fail to pay all the principal")
	}

	poolInfo.EscrowPrincipalAmount = poolInfo.EscrowPrincipalAmount.Add(adjAmount)

	exchange, found := k.GetExchangeInfo(ctx, poolInfo.Index)
	item := types.ExchangeItem{
		ExchangeRatio:          exchangeRatio,
		ActualPrincipalPayment: msg.Token.Amount,
		Proof:                  msg.ProofUrl,
	}
	if !found {
		items := make([]*types.ExchangeItem, 0, 10)
		exchange = types.ExchangeInfo{
			PoolIndex:                      poolInfo.Index,
			ExchangeItemsForPartialPayment: items,
			ExchangeItemForFullPayment:     &item,
		}
		k.SetExchangeInfo(ctx, exchange)
	} else {
		exchange.ExchangeItemForFullPayment = &item
		k.SetExchangeInfo(ctx, exchange)
	}

	// we only close the pool when the escrow principal is later than the total borrowed and the project pass the project length
	if poolInfo.EscrowPrincipalAmount.IsGTE(poolInfo.BorrowedAmount) && ctx.BlockTime().After(poolInfo.ProjectDueTime) {
		// once we are in the freezing state, the usable amount will not be accurate any longer
		poolInfo.PoolStatus = types.PoolInfo_FREEZING
	}
	poolInfo.PrincipalPaid = true
	poolInfo.ExchangeRatio = exchangeRatio

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

func (k msgServer) PayPrincipalForWithdrawalRequests(goCtx context.Context, msg *types.MsgPayPrincipalPartial) (*types.MsgPayPrincipalPartialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	spv, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	exchangeRatio, err := sdk.NewDecFromStr(msg.ExchangeRatio)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid exchange ratio %v", msg.ExchangeRatio)
	}

	poolInfo, found := k.GetPools(ctx, msg.GetPoolIndex())
	if !found {
		return nil, coserrors.Wrapf(sdkerrors.ErrNotFound, "pool cannot be found %v", msg.GetPoolIndex())
	}

	if poolInfo.WithdrawProposalAmount.IsZero() {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "no withdraw proposal to be paid")
	}

	if poolInfo.PrincipalPaid {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "principal already paid")
	}

	if !spv.Equals(poolInfo.OwnerAddress) {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "only pool owner can pay the principal")
	}

	paymentAmount, err := k.calulatetotalDueInterest(ctx, poolInfo)
	if err != nil {
		return nil, err
	}

	if poolInfo.EscrowInterestAmount.LT(paymentAmount) {
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

	adjAmount := sdk.NewCoin(msg.Token.Denom, exchangeRatio.MulInt(msg.Token.Amount).TruncateInt())
	if !adjAmount.Equal(poolInfo.GetWithdrawProposalAmount()) {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "you must pay exact full principal %v(adjust amount %v)", poolInfo.BorrowedAmount, adjAmount)
	}

	exchange, found := k.GetExchangeInfo(ctx, poolInfo.Index)
	item := types.ExchangeItem{
		ExchangeRatio:          exchangeRatio,
		ActualPrincipalPayment: msg.Token.Amount,
		Proof:                  msg.ProofUrl,
	}
	if !found {
		items := make([]*types.ExchangeItem, 1, 10)
		items[0] = &item
		exchange = types.ExchangeInfo{
			PoolIndex:                      poolInfo.Index,
			ExchangeItemsForPartialPayment: items,
			ExchangeItemForFullPayment:     nil,
		}
		k.SetExchangeInfo(ctx, exchange)
	} else {
		exchange.ExchangeItemsForPartialPayment = append(exchange.ExchangeItemsForPartialPayment, &item)
		k.SetExchangeInfo(ctx, exchange)
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, spv, types.ModuleAccount, sdk.NewCoins(msg.Token))
	if err != nil {
		return nil, coserrors.Wrapf(err, "fail to pay all the principal")
	}

	poolInfo.EscrowPrincipalAmount = poolInfo.EscrowPrincipalAmount.Add(adjAmount)
	poolInfo.PrincipalPaid = true
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
