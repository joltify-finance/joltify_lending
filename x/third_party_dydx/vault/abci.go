package vault

import (
	"runtime/debug"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/dydx_helper/lib/abci"
	"github.com/joltify-finance/joltify_lending/dydx_helper/lib/log"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/vault/keeper"
)

func BeginBlocker(
	ctx sdk.Context,
	keeper *keeper.Keeper,
) {
	// Panic is not expected in BeginBlocker, but we should recover instead of
	// halting the chain.
	if err := abci.RunCached(ctx, func(ctx sdk.Context) error {
		keeper.DecommissionNonPositiveEquityVaults(ctx)
		return nil
	}); err != nil {
		log.ErrorLog(
			ctx,
			"panic in vault BeginBlocker",
			"stack",
			string(debug.Stack()),
		)
	}
}

func EndBlocker(
	ctx sdk.Context,
	keeper *keeper.Keeper,
) {
	// Panic is not expected in EndBlocker, but we should recover instead of
	// halting the chain.
	if err := abci.RunCached(ctx, func(ctx sdk.Context) error {
		keeper.RefreshAllVaultOrders(ctx)
		return nil
	}); err != nil {
		log.ErrorLog(
			ctx,
			"panic in vault EndBlocker",
			"stack",
			string(debug.Stack()),
		)
	}
}
