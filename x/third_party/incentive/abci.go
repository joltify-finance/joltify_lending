package incentive

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
)

// BeginBlocker runs at the start of every block
func BeginBlocker(ctx context.Context, k keeper.Keeper) {
	params := k.GetParams(sdk.UnwrapSDKContext(ctx))

	for _, rp := range params.JoltSupplyRewardPeriods {
		k.AccumulateJoltSupplyRewards(sdk.UnwrapSDKContext(ctx), rp)
	}
	for _, rp := range params.JoltBorrowRewardPeriods {
		k.AccumulateJoltBorrowRewards(ctx, rp)
	}
	for _, rp := range params.SwapRewardPeriods {
		k.AccumulateSwapRewards(sdk.UnwrapSDKContext(ctx), rp)
	}
	for _, rp := range params.SPVRewardPeriods {
		k.AccumulateSPVRewards(sdk.UnwrapSDKContext(ctx), rp)
	}
}
