package prices

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	indexerevents "github.com/joltify-finance/joltify_lending/dydx_helper/indexer/events"
	"github.com/joltify-finance/joltify_lending/dydx_helper/indexer/indexer_manager"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/types"
)

// InitGenesis initializes the x/prices module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.InitializeForGenesis(ctx)

	if len(genState.MarketPrices) != len(genState.MarketParams) {
		panic("Expected the same number of market prices and market params")
	}

	// Set all the market params and prices.
	for i, param := range genState.MarketParams {
		if _, err := k.CreateMarket(ctx, param, genState.MarketPrices[i]); err != nil {
			panic(err)
		}
	}

	// Generate indexer events.
	priceUpdateIndexerEvents := keeper.GenerateMarketPriceUpdateIndexerEvents(genState.MarketPrices)
	for _, update := range priceUpdateIndexerEvents {
		k.GetIndexerEventManager().AddTxnEvent(
			ctx,
			indexerevents.SubtypeMarket,
			indexerevents.MarketEventVersion,
			indexer_manager.GetBytes(
				update,
			),
		)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.MarketParams = k.GetAllMarketParams(ctx)
	genesis.MarketPrices = k.GetAllMarketPrices(ctx)

	return genesis
}
