package pricefeed

import (
	"sync"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	gometrics "github.com/hashicorp/go-metrics"
	"github.com/joltify-finance/joltify_lending/daemons/pricefeed/api"
	"github.com/joltify-finance/joltify_lending/dydx_helper/lib"
	"github.com/joltify-finance/joltify_lending/dydx_helper/lib/metrics"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/types"
)

const (
	// MaxPriceAge defines the duration in which a price update is valid for.
	MaxPriceAge = time.Duration(30_000_000_000) // 30 sec, duration uses nanoseconds.
)

// PriceTimestamp maintains a price and its last update timestamp.
type PriceTimestamp struct {
	LastUpdateTime time.Time
	Price          uint64
}

// ExchangeToPrice maintains multiple prices from different exchanges for
// the same market, along with the last time the each exchange price was updated.
type ExchangeToPrice struct {
	marketId                 uint32
	exchangeToPriceTimestamp map[string]*PriceTimestamp
}

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

// NewExchangeToPrice creates a new ExchangeToPrice. It takes a market ID, which is used in logging and metrics to
// identify the market these exchange prices are for. The market ID does not otherwise affect the behavior
// of the ExchangeToPrice.
func NewExchangeToPrice(marketId uint32) *ExchangeToPrice {
	return &ExchangeToPrice{
		marketId:                 marketId,
		exchangeToPriceTimestamp: make(map[string]*types.PriceTimestamp),
	}
}

// UpdatePrices updates market prices given a list of price updates. Prices are
// only updated if the timestamp on the updates are greater than the timestamp
// on existing prices.
func (mte *MarketToExchangePrices) UpdatePrices(
	updates []*api.MarketPriceUpdate,
) {
	mte.Lock()
	defer mte.Unlock()
	for _, marketPriceUpdate := range updates {
		marketId := marketPriceUpdate.MarketId
		exchangeToPrices, ok := mte.marketToExchangePrices[marketId]
		if !ok {
			exchangeToPrices = NewExchangeToPrice(marketId)
			mte.marketToExchangePrices[marketId] = exchangeToPrices
		}
		exchangeToPrices.UpdatePrices(marketPriceUpdate.ExchangePrices)
	}
}

// GetValidMedianPrices returns median prices for multiple markets.
// Specifically, it returns a map where the key is the market ID and the value
// is the median price for the market. It only returns "valid" prices where
// a price is valid iff
// 1) the last update time is within a predefined threshold away from the given
// read time.
func (mte *MarketToExchangePrices) GetValidMedianPrices(
	marketParams []types.MarketParam,
	readTime time.Time,
) map[uint32]uint64 {
	cutoffTime := readTime.Add(-mte.maxPriceAge)
	marketIdToMedianPrice := make(map[uint32]uint64)

	mte.Lock()
	defer mte.Unlock()
	for _, marketParam := range marketParams {
		marketId := marketParam.Id
		exchangeToPrice, ok := mte.marketToExchangePrices[marketId]
		if !ok {
			// No market price info yet, skip this market.
			telemetry.IncrCounterWithLabels(
				[]string{
					metrics.PricefeedServer,
					metrics.NoMarketPrice,
					metrics.Count,
				},
				1,
				[]gometrics.Label{
					pricefeedmetrics.GetLabelForMarketId(marketId),
				},
			)
			continue
		}

		// GetValidPriceForMarket filters prices based on cutoff time.
		validPrices := exchangeToPrice.GetValidPrices(cutoffTime)
		telemetry.SetGaugeWithLabels(
			[]string{
				metrics.PricefeedServer,
				metrics.ValidPrices,
				metrics.Count,
			},
			float32(len(validPrices)),
			[]gometrics.Label{
				pricefeedmetrics.GetLabelForMarketId(marketId),
			},
		)

		// Calculate the median. Returns an error if the input is empty.
		median, err := lib.Median(validPrices)
		if err != nil {
			telemetry.IncrCounterWithLabels(
				[]string{
					metrics.PricefeedServer,
					metrics.NoValidMedianPrice,
					metrics.Count,
				},
				1,
				[]gometrics.Label{
					pricefeedmetrics.GetLabelForMarketId(marketId),
				},
			)
			continue
		}
		marketIdToMedianPrice[marketId] = median
	}

	return marketIdToMedianPrice
}
