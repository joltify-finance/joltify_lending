package mint

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/mint/keeper"
	"github.com/joltify-finance/joltify_lending/x/mint/types"
)

// InitGenesis initializes the mint module's state from a provided genesis
// state.
func InitGenesis(ctx context.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)

	err := k.FirstDist(ctx)
	if err != nil {
		panic(err)
	}

	h := types.HistoricalDistInfo{PayoutTime: ctx.BlockTime(), TotalMintCoins: sdk.NewCoins()}
	k.SetDistInfo(ctx, h)
}

// ExportGenesis returns the mint module's exported genesis.
func ExportGenesis(ctx context.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// this line is used by starport scaffolding # genesis/module/export
	h := k.GetDistInfo(ctx)
	genesis.HistoricalDistInfo = &h
	return genesis
}
