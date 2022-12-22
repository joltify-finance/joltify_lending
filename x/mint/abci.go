package mint

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/mint/keeper"
)

func BeginBlock(ctx sdk.Context, keeper keeper.Keeper) {
	// we mint the tokens
	keeper.DoDistribute(ctx)
}
