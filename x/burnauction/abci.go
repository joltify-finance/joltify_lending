package burnauction

import (
	"context"

	"github.com/joltify-finance/joltify_lending/x/burnauction/keeper"
)

func EndBlock(ctx context.Context, k keeper.Keeper) {
	k.RunSurplusAuctions(ctx)
}
