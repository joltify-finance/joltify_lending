package keeper

import (
	"context"

	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k msgServer) Borrow(goCtx context.Context, msg *types.MsgBorrow) (*types.MsgBorrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	caller, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	poolInfo, found := k.GetPools(ctx, msg.GetPoolIndex())
	if !found {
		return nil, coserrors.Wrapf(sdkerrors.ErrNotFound, "pool cannot be found %v", msg.GetPoolIndex())
	}

	if poolInfo.PoolStatus != types.PoolInfo_ACTIVE {
		return nil, coserrors.Wrap(types.ErrPoolNotActive, "pool has been closed")
	}

	if !poolInfo.OwnerAddress.Equals(caller) {
		return nil, coserrors.Wrapf(types.ErrUnauthorized, "%v is not authorized to borrow money", msg.Creator)
	}

	if msg.BorrowAmount.Denom != poolInfo.TotalAmount.Denom {
		return nil, coserrors.Wrap(types.ErrInconsistencyToken, "token to be borrowed is inconsistency")
	}

	if poolInfo.UsableAmount.IsLT(msg.BorrowAmount) {
		return nil, types.ErrInsufficientFund
	}

	k.doBorrow(ctx, &poolInfo, msg.BorrowAmount, true, nil, sdk.ZeroInt())

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeBorrow,
			sdk.NewAttribute(types.AttributeCreator, msg.Creator),
			sdk.NewAttribute("amount", msg.BorrowAmount.Amount.String()),
		),
	)

	return &types.MsgBorrowResponse{BorrowAmount: msg.BorrowAmount.String()}, nil
}
