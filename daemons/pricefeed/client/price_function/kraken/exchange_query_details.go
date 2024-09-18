package kraken

import (
	"github.com/joltify-finance/joltify_lending/daemons/pricefeed/client/constants/exchange_common"
	"github.com/joltify-finance/joltify_lending/daemons/pricefeed/client/types"
)

var KrakenDetails = types.ExchangeQueryDetails{
	Exchange:      exchange_common.EXCHANGE_ID_KRAKEN,
	Url:           "https://api.kraken.com/0/public/Ticker",
	PriceFunction: KrakenPriceFunction,
	IsMultiMarket: true,
}
