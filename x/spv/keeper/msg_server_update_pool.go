package keeper

import (
	"context"
	coserrors "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k msgServer) UpdatePool(goCtx context.Context, msg *types.MsgUpdatePool) (*types.MsgUpdatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	caller, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	poolInfo, found := k.GetPools(ctx, msg.GetPoolIndex())
	if !found {
		return nil, coserrors.Wrapf(sdkerrors.ErrNotFound, "pool cannot be found %v", msg.GetPoolIndex())
	}

	if poolInfo.PoolStatus != types.PoolInfo_PREPARE {
		return nil, types.ErrUNEXPECTEDSTATUS
	}

	if !poolInfo.OwnerAddress.Equals(caller) {
		return nil, coserrors.Wrapf(types.ErrUnauthorized, "%v is not authorized to update the pool", msg.Creator)
	}

	apy, payfreq, err := parameterSanitize(msg.PayFreq, msg.PoolApy)
	if err != nil {
		return nil, coserrors.Wrapf(types.ErrInvalidParameter, "invalid parameter: %v", err.Error())
	}

	poolInfo.PoolName = msg.PoolName
	poolInfo.Apy = apy
	poolInfo.PayFreq = payfreq
	poolInfo.TotalAmount = msg.TargetTokenAmount

	k.SetPool(ctx, poolInfo)
	return &types.MsgUpdatePoolResponse{}, nil
}
