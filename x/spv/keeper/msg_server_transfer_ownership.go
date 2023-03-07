package keeper

import (
	"context"

	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

// TransferOwnership will allow the investor to submit the request to transfer/update their ratio in the pool
func (k msgServer) TransferOwnership(goCtx context.Context, msg *types.MsgTransferOwnership) (*types.MsgTransferOwnershipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	caller, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	d, found := k.GetDepositor(ctx, msg.PoolIndex, caller)
	if !found {
		return &types.MsgTransferOwnershipResponse{}, types.ErrDepositorNotFound
	}

	poolInfo, found := k.GetPools(ctx, msg.GetPoolIndex())
	if !found {
		return &types.MsgTransferOwnershipResponse{}, types.ErrPoolNotFound
	}

	if poolInfo.PoolStatus != types.PoolInfo_ACTIVE {
		return &types.MsgTransferOwnershipResponse{}, types.ErrUNEXPECTEDSTATUS
	}

	poolInfo.TransferAccounts = append(poolInfo.TransferAccounts, caller)
	d.TransferRequest = true

	k.SetDepositor(ctx, d)
	k.SetPool(ctx, poolInfo)

	return &types.MsgTransferOwnershipResponse{OperationResult: true}, nil
}
