package keeper

import (
	"context"

	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QueryPool(goCtx context.Context, req *types.QueryQueryPoolRequest) (*types.QueryQueryPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	pool, ok := k.GetPools(ctx, req.GetPoolIndex())
	if !ok {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidRequest, "index cannot be found %v", req.GetPoolIndex())
	}
	return &types.QueryQueryPoolResponse{PoolInfo: &pool}, nil
}

func (k Keeper) QueryHistoryPool(goCtx context.Context, req *types.QueryQueryPoolRequest) (*types.QueryQueryPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	pool, ok := k.GetHistoryPools(ctx, req.GetPoolIndex())
	if !ok {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidRequest, "index cannot be found %v", req.GetPoolIndex())
	}
	return &types.QueryQueryPoolResponse{PoolInfo: &pool}, nil
}
