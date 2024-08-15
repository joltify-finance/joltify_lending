package keeper

import (
	"context"

	"google.golang.org/grpc/status"

	"github.com/joltify-finance/joltify_lending/lib"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/types"
	"google.golang.org/grpc/codes"
)

func (k Keeper) EquityTierLimitConfiguration(
	c context.Context,
	req *types.QueryEquityTierLimitConfigurationRequest,
) (*types.QueryEquityTierLimitConfigurationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := lib.UnwrapSDKContext(c, types.ModuleName)

	return &types.QueryEquityTierLimitConfigurationResponse{
		EquityTierLimitConfig: k.GetEquityTierLimitConfiguration(ctx),
	}, nil
}
