package incentive

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
)

// BeginBlocker runs at the start of every block
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	params := k.GetParams(ctx)

	for _, rp := range params.JoltSupplyRewardPeriods {
		k.AccumulateJoltSupplyRewards(ctx, rp)
	}
	for _, rp := range params.JoltBorrowRewardPeriods {
		k.AccumulateJoltBorrowRewards(ctx, rp)
	}
	for _, rp := range params.SwapRewardPeriods {
		k.AccumulateSwapRewards(ctx, rp)
	}
	for _, rp := range params.SPVRewardPeriods {
		k.AccumulateSPVRewards(ctx, rp)
	}
}
