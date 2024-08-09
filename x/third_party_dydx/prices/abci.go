package prices

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/types"
)

// PreBlocker executes all ABCI PreBlock logic respective to the prices module.
func PreBlocker(
	ctx sdk.Context,
	keeper types.PricesKeeper,
) {
	keeper.InitializeCurrencyPairIdCache(ctx)
}
