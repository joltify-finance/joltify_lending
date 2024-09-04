package testexchange

import (
	"fmt"

	"github.com/joltify-finance/joltify_lending/daemons/pricefeed/client/constants/exchange_common"
	"github.com/joltify-finance/joltify_lending/daemons/pricefeed/client/price_function/coinbase_pro"
	"github.com/joltify-finance/joltify_lending/daemons/pricefeed/client/types"
)

// Exchange used for testing purposes. We'll reuse the CoinbasePro price function.
var (
	TestExchangeHost    = "test.exchange"
	TestExchangePort    = "9888"
	TestExchangeDetails = types.ExchangeQueryDetails{
		Exchange:      exchange_common.EXCHANGE_ID_TEST_EXCHANGE,
		Url:           fmt.Sprintf("http://%s:%s/ticker?symbol=$", TestExchangeHost, TestExchangePort),
		PriceFunction: coinbase_pro.CoinbaseProPriceFunction,
	}
)
