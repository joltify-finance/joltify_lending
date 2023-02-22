package spv

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/keeper"
	abci "github.com/tendermint/tendermint/abci/types"
)

func EndBlock(ctx sdk.Context, keeper keeper.Keeper) []abci.ValidatorUpdate {

	return keeper.NewUpdate(ctx)
}
