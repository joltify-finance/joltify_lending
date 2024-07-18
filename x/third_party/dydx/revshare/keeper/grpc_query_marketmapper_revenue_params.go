package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/revshare/types"
)

func (k Keeper) MarketMapperRevenueShareParams(
	ctx context.Context,
	req *types.QueryMarketMapperRevenueShareParams,
) (*types.QueryMarketMapperRevenueShareParamsResponse, error) {
	params := k.GetMarketMapperRevenueShareParams(sdk.UnwrapSDKContext(ctx))
	return &types.QueryMarketMapperRevenueShareParamsResponse{
		Params: params,
	}, nil
}
