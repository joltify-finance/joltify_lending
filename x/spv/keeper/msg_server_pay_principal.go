package keeper

import (
	"context"
	"errors"

	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

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

	if msg.Token.Denom != poolInfo.BorrowedAmount.Denom {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid token demo, want %v", poolInfo.BorrowedAmount.Denom)
	}

	class, found := k.nftKeeper.GetClass(ctx, poolInfo.PoolNFTIds[0])
	if !found {
		panic(found)
	}
	var borrowInterest types.BorrowInterest
	err = proto.Unmarshal(class.Data.Value, &borrowInterest)
	if err != nil {
		return nil, coserrors.Wrapf(err, "invalid unmarshal of the borrow interest")
	}
	paymentAmount := borrowInterest.MonthlyRatio.Mul(sdk.NewDecFromInt(poolInfo.BorrowedAmount.Amount)).TruncateInt()

	if poolInfo.EscrowInterestAmount.LT(paymentAmount) {
		return nil, errors.New("not enough interest to be paid to close the pool")
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, spv, types.ModuleAccount, sdk.NewCoins(msg.Token))
	if err != nil {
		return nil, coserrors.Wrapf(err, "fail to pay all the principal")
	}

	poolInfo.EscrowPrincipalAmount = poolInfo.EscrowPrincipalAmount.Add(msg.Token)

	// we only close the pool when the escrow principal is later than the total borrowed and the project pass the project length
	if poolInfo.EscrowPrincipalAmount.IsGTE(poolInfo.BorrowedAmount) && ctx.BlockTime().After(poolInfo.ProjectDueTime) {
		// once we are in the freezing state, the usable amount will not be accurate any longer
		poolInfo.PoolStatus = types.PoolInfo_FREEZING
	}

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

	poolInfo, found := k.GetPools(ctx, msg.GetPoolIndex())
	if !found {
		return nil, coserrors.Wrapf(sdkerrors.ErrNotFound, "pool cannot be found %v", msg.GetPoolIndex())
	}

	if msg.Token.Denom != poolInfo.BorrowedAmount.Denom {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid token demo, want %v", poolInfo.BorrowedAmount.Denom)
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, spv, types.ModuleAccount, sdk.NewCoins(msg.Token))
	if err != nil {
		return nil, coserrors.Wrapf(err, "fail to pay all the principal")
	}

	poolInfo.EscrowPrincipalAmount = poolInfo.EscrowPrincipalAmount.Add(msg.Token)
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
