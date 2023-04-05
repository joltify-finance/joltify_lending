package keeper

import (
	"context"
	"time"

	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func checkEligibility(blockTime time.Time, poolInfo types.PoolInfo) error {

	if poolInfo.PoolStatus != types.PoolInfo_ACTIVE {
		return coserrors.Wrap(types.ErrPoolNotActive, "pool is not in active status")
	}

	if poolInfo.CurrentPoolTotalBorrowCounter >= poolInfo.PoolTotalBorrowLimit {
		return types.ErrPoolBorrowLimit
	}

	if poolInfo.CurrentPoolTotalBorrowCounter == 0 && poolInfo.PoolCreatedTime.Add(time.Second*time.Duration(poolInfo.PoolLockedSeconds)).Before(blockTime) {
		return types.ErrPoolBorrowExpire
	}

	if poolInfo.CurrentPoolTotalBorrowCounter == 0 && poolInfo.UsableAmount.IsLT(poolInfo.TargetAmount) {
		return coserrors.Wrapf(types.ErrInsufficientFund, "pool target is %v and we have %v usable", poolInfo.TargetAmount, poolInfo.UsableAmount)
	}
	return nil
}

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

	err = checkEligibility(ctx.BlockTime(), poolInfo)
	if err != nil {
		return nil, err
	}

	poolInfo.CurrentPoolTotalBorrowCounter += 1

	if !poolInfo.OwnerAddress.Equals(caller) {
		return nil, coserrors.Wrapf(types.ErrUnauthorized, "%v is not authorized to borrow money", msg.Creator)
	}

	if msg.BorrowAmount.Denom != poolInfo.TotalAmount.Denom {
		return nil, coserrors.Wrap(types.ErrInconsistencyToken, "token to be borrowed is inconsistency")
	}

	if poolInfo.UsableAmount.IsLT(msg.BorrowAmount) {
		return nil, types.ErrInsufficientFund
	}

	if poolInfo.TargetAmount.IsLT(poolInfo.BorrowedAmount.Add(msg.BorrowAmount)) {
		return nil, types.ErrPoolFull
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
