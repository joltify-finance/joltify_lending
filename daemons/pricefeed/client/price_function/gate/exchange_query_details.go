package gate

import (
	"github.com/joltify-finance/joltify_lending/daemons/pricefeed/client/constants/exchange_common"
	"github.com/joltify-finance/joltify_lending/daemons/pricefeed/client/types"
)

var (
	GateDetails = types.ExchangeQueryDetails{
		Exchange:      exchange_common.EXCHANGE_ID_GATE,
		Url:           "https://api.gateio.ws/api/v4/spot/tickers",
		PriceFunction: GatePriceFunction,
		IsMultiMarket: true,
	}
)
