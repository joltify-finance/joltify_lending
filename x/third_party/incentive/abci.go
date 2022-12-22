package incentive

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
)

// BeginBlocker runs at the start of every block
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	params := k.GetParams(ctx)

	// not enabled
	//for _, rp := range params.USDXMintingRewardPeriods {
	//	k.AccumulateUSDXMintingRewards(ctx, rp)
	//}
	for _, rp := range params.JoltSupplyRewardPeriods {
		k.AccumulateJoltSupplyRewards(ctx, rp)
	}
	for _, rp := range params.JoltBorrowRewardPeriods {
		k.AccumulateJoltBorrowRewards(ctx, rp)
	}
	//for _, rp := range params.DelegatorRewardPeriods {
	//	k.AccumulateDelegatorRewards(ctx, rp)
	//}
	// swap is not enabled now
	//for _, rp := range params.SwapRewardPeriods {
	//	k.AccumulateSwapRewards(ctx, rp)
	//}
	//for _, rp := range params.SavingsRewardPeriods {
	//	k.AccumulateSavingsRewards(ctx, rp)
	//}
}
