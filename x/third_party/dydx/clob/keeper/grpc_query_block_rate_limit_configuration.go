package keeper

import (
	"context"

	"github.com/dydxprotocol/v4-chain/protocol/lib"
	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/clob/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) BlockRateLimitConfiguration(
	c context.Context,
	req *types.QueryBlockRateLimitConfigurationRequest,
) (*types.QueryBlockRateLimitConfigurationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := lib.UnwrapSDKContext(c, types.ModuleName)
	blockRateLimitConfig := k.GetBlockRateLimitConfiguration(ctx)

	return &types.QueryBlockRateLimitConfigurationResponse{
		BlockRateLimitConfig: blockRateLimitConfig,
	}, nil
}
