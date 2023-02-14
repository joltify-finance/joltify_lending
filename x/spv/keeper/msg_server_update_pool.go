package keeper

import (
	"context"

	coserrors "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/crypto"

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

	allProjects := k.kycKeeper.GetProjects(ctx)
	if allProjects == nil || int32(len(allProjects)) < poolInfo.LinkedProject {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "the given project %v cannot be found", poolInfo.LinkedProject)
	}

	targetProject := allProjects[poolInfo.LinkedProject-1]

	apy, _, err := parameterSanitize(targetProject.PayFreq, msg.PoolApy)
	if err != nil {
		return nil, coserrors.Wrapf(types.ErrInvalidParameter, "invalid parameter: %v", err.Error())
	}

	var poolJunior, poolSenior *types.PoolInfo
	queryType := "junior"
	if poolInfo.PoolType == types.PoolInfo_JUNIOR {
		queryType = "senior"
	}

	indexHash2 := crypto.Keccak256Hash([]byte(targetProject.BasicInfo.ProjectName), poolInfo.OwnerAddress.Bytes(), []byte(queryType))

	poolInfo2, found := k.GetPools(ctx, indexHash2.Hex())
	if !found {
		return nil, coserrors.Wrapf(sdkerrors.ErrNotFound, "pool cannot be found %v", msg.GetPoolIndex())
	}

	if poolInfo.PoolType == types.PoolInfo_JUNIOR {
		poolJunior = &poolInfo
		poolSenior = &poolInfo2
	} else {
		poolJunior = &poolInfo2
		poolSenior = &poolInfo
	}

	if poolInfo.PoolStatus != types.PoolInfo_PREPARE || poolInfo2.PoolStatus != types.PoolInfo_PREPARE {
		return nil, types.ErrUNEXPECTEDSTATUS
	}

	if !poolInfo.OwnerAddress.Equals(caller) {
		return nil, coserrors.Wrapf(types.ErrUnauthorized, "%v is not authorized to update the pool", msg.Creator)
	}

	if !poolInfo2.OwnerAddress.Equals(caller) {
		return nil, coserrors.Wrapf(types.ErrUnauthorized, "%v is not authorized to update the pool", msg.Creator)
	}

	isJunior := poolInfo.PoolType == types.PoolInfo_JUNIOR

	poolsInfoAPY, poolsInfoAmount, err := calculateApys(targetProject.ProjectTargetAmount, msg.TargetTokenAmount, targetProject.BaseApy, apy, isJunior)

	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "junior pool amount larger than target")
	}

	poolInfo.PoolName = msg.PoolName
	poolInfo.Apy = apy
	poolInfo.TargetAmount = msg.TargetTokenAmount
	k.SetPool(ctx, poolInfo)

	if isJunior {
		poolSenior.Apy = poolsInfoAPY["senior"]
		poolSenior.TargetAmount = poolsInfoAmount["senior"]
		poolSenior.PoolName = msg.PoolName
		k.SetPool(ctx, *poolSenior)
	} else {
		poolJunior.Apy = poolsInfoAPY["junior"]
		poolJunior.TargetAmount = poolsInfoAmount["junior"]
		poolJunior.PoolName = msg.PoolName
		k.SetPool(ctx, *poolJunior)
	}

	return &types.MsgUpdatePoolResponse{}, nil
}
