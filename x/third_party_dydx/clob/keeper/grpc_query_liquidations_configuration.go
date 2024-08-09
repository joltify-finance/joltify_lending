package keeper

import (
	"context"

	"github.com/joltify-finance/joltify_lending/dydx_helper/lib"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LiquidationsConfiguration returns the liquidations configuration.
func (k Keeper) LiquidationsConfiguration(
	c context.Context,
	req *types.QueryLiquidationsConfigurationRequest,
) (*types.QueryLiquidationsConfigurationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := lib.UnwrapSDKContext(c, types.ModuleName)
	liquidationsConfig := k.GetLiquidationsConfig(ctx)
	return &types.QueryLiquidationsConfigurationResponse{
		LiquidationsConfig: liquidationsConfig,
	}, nil
}
