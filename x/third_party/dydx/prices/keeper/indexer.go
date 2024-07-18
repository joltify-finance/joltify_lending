package keeper

import (
	indexerevents "github.com/dydxprotocol/v4-chain/protocol/indexer/events"
	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/prices/types"
)

// GenerateMarketPriceUpdateIndexerEvents takes in a slice of market prices
// and returns a slice of price updates.
func GenerateMarketPriceUpdateIndexerEvents(
	markets []types.MarketPrice,
) []*indexerevents.MarketEventV1 {
	events := make([]*indexerevents.MarketEventV1, 0, len(markets))
	for _, market := range markets {
		events = append(
			events,
			indexerevents.NewMarketPriceUpdateEvent(
				market.Id,
				market.Price,
			),
		)
	}
	return events
}
