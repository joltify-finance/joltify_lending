package keeper

import (
	"context"

	coserrors "cosmossdk.io/errors"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k msgServer) ActivePool(goCtx context.Context, msg *types.MsgActivePool) (*types.MsgActivePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message

	caller, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	poolInfo1, found := k.GetPools(ctx, msg.GetPoolIndex())
	if !found {
		return nil, coserrors.Wrapf(errorsmod.ErrNotFound, "pool cannot be found %v", msg.GetPoolIndex())
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

	targetProject, ok := k.kycKeeper.GetProject(ctx, poolInfo1.LinkedProject)
	if !ok {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidRequest, "the given project %v cannot be found", poolInfo1.LinkedProject)
	}

	if poolInfo1.SeparatePool || poolInfo1.PoolType == types.PoolInfo_JUNIOR {
		k.SetPool(ctx, poolInfo1)
		return &types.MsgActivePoolResponse{}, nil
	}

	juniorPoolIndex := crypto.Keccak256Hash([]byte(targetProject.BasicInfo.ProjectName), poolInfo1.OwnerAddress.Bytes(), []byte("junior"))

	juniorPoolInfo, found := k.GetPools(ctx, juniorPoolIndex.Hex())
	if !found {
		return nil, coserrors.Wrapf(errorsmod.ErrNotFound, "pool cannot be found %v", msg.GetPoolIndex())
	}

	if juniorPoolInfo.PoolStatus != types.PoolInfo_ACTIVE {
		return nil, coserrors.Wrapf(types.ErrUNEXPECTEDSTATUS, "junior pool must be activated firstly")
	}

	totalTarget := poolInfo1.TargetAmount.Add(juniorPoolInfo.TargetAmount)
	poolAmount := sdk.NewDecFromInt(juniorPoolInfo.TargetAmount.Amount)

	ratio := poolAmount.QuoInt(totalTarget.Amount)

	if ratio.LT(targetProject.JuniorMinRatio) {
		return nil, coserrors.Wrapf(types.ErrInvalidParameter, "junior ratio is too low, current is %v and target is %v", ratio, targetProject.JuniorMinRatio)
	}

	k.SetPool(ctx, poolInfo1)
	return &types.MsgActivePoolResponse{}, nil
}
