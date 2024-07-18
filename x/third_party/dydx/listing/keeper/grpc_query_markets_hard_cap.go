package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/listing/types"
)

func (k Keeper) MarketsHardCap(
	ctx context.Context,
	req *types.QueryMarketsHardCap,
) (*types.QueryMarketsHardCapResponse, error) {
	hardCap := k.GetMarketsHardCap(sdk.UnwrapSDKContext(ctx))
	return &types.QueryMarketsHardCapResponse{
		HardCap: hardCap,
	}, nil
}
