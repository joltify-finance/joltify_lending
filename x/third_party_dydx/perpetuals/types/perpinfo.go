package types

import (
	pricestypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/types"
)

// PerpInfo contains all information needed to calculate margin requirements for a perpetual.
type PerpInfo struct {
	Perpetual     Perpetual
	Price         pricestypes.MarketPrice
	LiquidityTier LiquidityTier
}
