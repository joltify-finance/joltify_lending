package keeper

import (
	"context"

	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PoolInvestors(goCtx context.Context, req *types.QueryPoolInvestorsRequest) (*types.QueryPoolInvestorsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	investorsResp, found := k.GetInvestorToPool(ctx, req.PoolIndex)
	if !found {
		return nil, coserrors.Wrapf(types.ErrPoolNotFound, "the pool %v does not exist", req.PoolIndex)
	}

	return &types.QueryPoolInvestorsResponse{Investors: investorsResp.Investors}, nil
}
