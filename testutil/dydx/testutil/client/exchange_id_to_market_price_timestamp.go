package client

import (
	"github.com/joltify-finance/joltify_lending/daemons/pricefeed/client/types"
)

// ExchangeIdMarketPriceTimestamp contains an `ExchangeId` and an associated
// `types.MarketPriceTimestamp`. This type exists for convenience and clarity in testing the
// pricefeed client.
type ExchangeIdMarketPriceTimestamp struct {
	ExchangeId           types.ExchangeId
	MarketPriceTimestamp *types.MarketPriceTimestamp
}
