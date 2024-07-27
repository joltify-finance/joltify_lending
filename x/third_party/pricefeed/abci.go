package pricefeed

import (
	"context"
	"errors"

	"github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"
)

// EndBlocker updates the current pricefeed
func EndBlocker(ctx context.Context, k keeper.Keeper) {
	// Update the current price of each asset.
	for _, market := range k.GetMarkets(ctx) {
		if !market.Active {
			continue
		}

		err := k.SetCurrentPrices(ctx, market.MarketID)
		if err != nil && !errors.Is(err, types.ErrNoValidPrice) {
			panic(err)
		}
	}
}
