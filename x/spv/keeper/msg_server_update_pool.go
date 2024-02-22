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

	targetProject, ok := k.kycKeeper.GetProject(ctx, poolInfo.LinkedProject)
	if !ok {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "the given project %v cannot be found", poolInfo.LinkedProject)
	}

	// we use the second one as the mock apy
	apy, _, err := parameterSanitize(targetProject.PayFreq, []string{msg.PoolApy, "0"})
	if err != nil {
		return nil, coserrors.Wrapf(types.ErrInvalidParameter, "invalid parameter: %v", err.Error())
	}

	if msg.TargetTokenAmount.Denom != poolInfo.TargetAmount.Denom {
		return nil, coserrors.Wrapf(types.ErrInvalidParameter, "invalid parameter: %v", "target amount denom is not matched")
	}

	if poolInfo.PoolStatus != types.PoolInfo_PREPARE {
		return nil, types.ErrUNEXPECTEDSTATUS
	}

	if !poolInfo.OwnerAddress.Equals(caller) {
		return nil, coserrors.Wrapf(types.ErrUnauthorized, "%v is not authorized to update the pool", msg.Creator)
	}

	pType := "-senior"
	if poolInfo.PoolType == types.PoolInfo_JUNIOR {
		pType = "-junior"
	}

	poolInfo.PoolName = msg.PoolName + pType
	poolInfo.Apy = apy[0]
	poolInfo.TargetAmount = msg.TargetTokenAmount
	k.SetPool(ctx, poolInfo)

	return &types.MsgUpdatePoolResponse{}, nil
}
