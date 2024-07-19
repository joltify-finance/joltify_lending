package jolt

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/jolt/keeper"
)

// BeginBlocker updates interest rates
func BeginBlocker(ctx context.Context, k keeper.Keeper) {
	k.ApplyInterestRateUpdates(ctx)
	err := k.RunSurplusAuctions(ctx)
	if err != nil {
		ctx.Logger().Error("jolt", "surplusAuction", err)
	}
}
