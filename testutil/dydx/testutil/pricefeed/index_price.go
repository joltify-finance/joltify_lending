package pricefeed

import (
	pricefeedapi "github.com/joltify-finance/joltify_lending/daemons/pricefeed/api"
	"github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/constants"
)

func GetTestMarketPriceUpdates(n int) (indexPrices []*pricefeedapi.MarketPriceUpdate) {
	for i := 0; i < n; i++ {
		indexPrices = append(
			indexPrices,
			&pricefeedapi.MarketPriceUpdate{
				MarketId: uint32(i),
				ExchangePrices: []*pricefeedapi.ExchangePrice{
					constants.Exchange1_Price1_TimeT,
					constants.Exchange2_Price2_TimeT,
				},
			},
		)
	}
	return indexPrices
}
