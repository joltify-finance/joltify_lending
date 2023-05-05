package keeper

import (
	"context"

	"github.com/ethereum/go-ethereum/crypto"

	coserrors "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k msgServer) ActivePool(goCtx context.Context, msg *types.MsgActivePool) (*types.MsgActivePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message

	caller, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	poolInfo1, found := k.GetPools(ctx, msg.GetPoolIndex())
	if !found {
		return nil, coserrors.Wrapf(sdkerrors.ErrNotFound, "pool cannot be found %v", msg.GetPoolIndex())
	}

	if poolInfo1.PoolStatus != types.PoolInfo_PREPARE {
		return nil, types.ErrUNEXPECTEDSTATUS
	}

	if !poolInfo1.OwnerAddress.Equals(caller) {
		return nil, coserrors.Wrapf(types.ErrUnauthorized, "%v is not authorized to active the pool", msg.Creator)
	}

	poolInfo1.PoolStatus = types.PoolInfo_ACTIVE
	poolInfo1.PoolCreatedTime = ctx.BlockTime()
	poolInfo1.LastPaymentTime = ctx.BlockTime()

	allProjects := k.kycKeeper.GetProjects(ctx)

	targetProject := allProjects[poolInfo1.LinkedProject-1]

	var t string
	if poolInfo1.PoolType == types.PoolInfo_JUNIOR {
		t = "senior"
	} else {
		t = "junior"
	}

	secondPoolIndex := crypto.Keccak256Hash([]byte(targetProject.BasicInfo.ProjectName), poolInfo1.OwnerAddress.Bytes(), []byte(t))

	poolInfo2, found := k.GetPools(ctx, secondPoolIndex.Hex())
	if !found {
		return nil, coserrors.Wrapf(sdkerrors.ErrNotFound, "pool cannot be found %v", msg.GetPoolIndex())
	}

	if poolInfo2.PoolStatus != types.PoolInfo_PREPARE {
		return nil, types.ErrUNEXPECTEDSTATUS
	}

	if !poolInfo2.OwnerAddress.Equals(caller) {
		return nil, coserrors.Wrapf(types.ErrUnauthorized, "%v is not authorized to active the pool", msg.Creator)
	}

	poolInfo2.PoolStatus = types.PoolInfo_ACTIVE
	poolInfo2.PoolCreatedTime = ctx.BlockTime()
	poolInfo2.LastPaymentTime = ctx.BlockTime()

	totalTarget := poolInfo1.TargetAmount.Add(poolInfo2.TargetAmount)

	var poolAmount sdk.Dec
	if poolInfo1.PoolType == types.PoolInfo_JUNIOR {
		poolAmount = sdk.NewDecFromInt(poolInfo1.TargetAmount.Amount)
	} else {
		poolAmount = sdk.NewDecFromInt(poolInfo2.TargetAmount.Amount)
	}

	ratio := poolAmount.QuoInt(totalTarget.Amount)

	if ratio.LT(targetProject.JuniorMinRatio) {
		return nil, coserrors.Wrapf(types.ErrInvalidParameter, "junior ratio is too low, current is %v and target is %v", ratio, targetProject.JuniorMinRatio)
	}

	k.SetPool(ctx, poolInfo1)
	k.SetPool(ctx, poolInfo2)

	return &types.MsgActivePoolResponse{}, nil
}
