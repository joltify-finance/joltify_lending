package mexc

import (
	"github.com/joltify-finance/joltify_lending/daemons/pricefeed/client/constants/exchange_common"
	"github.com/joltify-finance/joltify_lending/daemons/pricefeed/client/types"
)

var MexcDetails = types.ExchangeQueryDetails{
	Exchange:      exchange_common.EXCHANGE_ID_MEXC,
	Url:           "https://www.mexc.com/open/api/v2/market/ticker",
	PriceFunction: MexcPriceFunction,
	IsMultiMarket: true,
}
