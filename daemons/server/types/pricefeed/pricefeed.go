package pricefeed

import (
	"time"

	pricefeedmetrics "github.com/joltify-finance/joltify_lending/daemons/pricefeed/metrics"

	"github.com/cosmos/cosmos-sdk/telemetry"
	gometrics "github.com/hashicorp/go-metrics"
	"github.com/joltify-finance/joltify_lending/daemons/pricefeed/api"
	"github.com/joltify-finance/joltify_lending/lib"
	"github.com/joltify-finance/joltify_lending/lib/metrics"
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

// NewPriceTimestamp creates a new PriceTimestamp.
func NewPriceTimestamp() *PriceTimestamp {
	return &PriceTimestamp{}
}

// UpdatePrice updates the price if the given update has a greater timestamp. Returns true if
// updating succeeds. Otherwise, returns false.
func (pt *PriceTimestamp) UpdatePrice(price uint64, newUpdateTime *time.Time) bool {
	if newUpdateTime.After(pt.LastUpdateTime) {
		pt.LastUpdateTime = *newUpdateTime
		pt.Price = price

		return true
	}

	return false
}

// GetValidPrice returns (price, true) if the last update time is greater than or
// equal to the given cutoff time. Otherwise returns (0, false).
func (pt *PriceTimestamp) GetValidPrice(cutoffTime time.Time) (uint64, bool) {
	if pt.LastUpdateTime.Before(cutoffTime) {
		return 0, false
	}
	return pt.Price, true
}

// GetValidPrices returns a list of "valid" prices. Prices are considered
// "valid" iff the last update time is greater than or equal to the given cutoff time.
func (etp *ExchangeToPrice) GetValidPrices(
	cutoffTime time.Time,
) []uint64 {
	validExchangePricesForMarket := make([]uint64, 0, len(etp.exchangeToPriceTimestamp))
	for exchangeId, priceTimestamp := range etp.exchangeToPriceTimestamp {
		validity := metrics.Valid

		// PriceTimestamp returns price if the last update time is valid.
		if price, ok := priceTimestamp.GetValidPrice(cutoffTime); ok {
			validExchangePricesForMarket = append(validExchangePricesForMarket, price)
		} else {
			// Price is invalid.
			validity = metrics.PriceIsInvalid
		}

		// Measure count of valid and invalid prices fetched from the in-memory map.
		telemetry.IncrCounterWithLabels(
			[]string{
				metrics.PricefeedServer,
				metrics.GetValidPrices,
				validity,
				metrics.Count,
			},
			1,
			[]gometrics.Label{
				pricefeedmetrics.GetLabelForExchangeId(exchangeId),
				pricefeedmetrics.GetLabelForMarketId(etp.marketId),
			},
		)
	}
	return validExchangePricesForMarket
}

// UpdatePrices updates prices given a list of prices from different exchanges.
// Prices are only updated if the timestamp on the updates are greater than
// the timestamp on existing prices.
func (etp *ExchangeToPrice) UpdatePrices(updates []*api.ExchangePrice) {
	for _, exchangePrice := range updates {
		exchangeId := exchangePrice.ExchangeId
		priceTimestamp, exists := etp.exchangeToPriceTimestamp[exchangeId]
		if !exists {
			priceTimestamp = NewPriceTimestamp()
			etp.exchangeToPriceTimestamp[exchangeId] = priceTimestamp
		}

		isUpdated := priceTimestamp.UpdatePrice(exchangePrice.Price, exchangePrice.LastUpdateTime)

		validity := metrics.Valid
		if exists && !isUpdated {
			validity = metrics.Invalid
		}

		// Measure count of valid and invalid prices inserted into the in-memory map.
		telemetry.IncrCounterWithLabels(
			[]string{metrics.PricefeedServer, metrics.UpdatePrice, validity, metrics.Count},
			1,
			[]gometrics.Label{
				pricefeedmetrics.GetLabelForMarketId(etp.marketId),
				pricefeedmetrics.GetLabelForExchangeId(exchangeId),
			},
		)
	}
}

// NewExchangeToPrice creates a new ExchangeToPrice. It takes a market ID, which is used in logging and metrics to
// identify the market these exchange prices are for. The market ID does not otherwise affect the behavior
// of the ExchangeToPrice.
func NewExchangeToPrice(marketId uint32) *ExchangeToPrice {
	return &ExchangeToPrice{
		marketId:                 marketId,
		exchangeToPriceTimestamp: make(map[string]*PriceTimestamp),
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
