package vault

import (
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/vault/keeper"
)

func EndBlock(ctx sdk.Context, keeper keeper.Keeper) []abci.ValidatorUpdate {
	// we burn the token after the first churn of the network
	keeper.ProcessAccountLeft(ctx)
	return keeper.NewUpdate(ctx)
}
