package pricefeed

import (
	"sync"
	"time"
)

// MarketToExchangePrices maintains price info for multiple markets. Each
// market can support prices from multiple exchange sources. Specifically,
// MarketToExchangePrices supports methods to update prices and to retrieve
// median prices. Methods are goroutine safe.
type MarketToExchangePrices struct {
	sync.Mutex                                         // lock
	marketToExchangePrices map[uint32]*ExchangeToPrice // {k: market id, v: exchange prices}
	// maxPriceAge is the maximum age of a price before it is considered too stale to be used.
	// Prices older than this age will not be used to calculate the median price.
	maxPriceAge time.Duration
}

// NewMarketToExchangePrices creates a new MarketToExchangePrices.
func NewMarketToExchangePrices(maxPriceAge time.Duration) *MarketToExchangePrices {
	return &MarketToExchangePrices{
		marketToExchangePrices: make(map[uint32]*ExchangeToPrice),
		maxPriceAge:            maxPriceAge,
	}
}
