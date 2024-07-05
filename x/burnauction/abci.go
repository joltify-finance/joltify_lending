package burnauction

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/burnauction/keeper"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) {
	k.RunSurplusAuctions(ctx)
}
