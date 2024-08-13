package grpc

import (
	bridgetypes "github.com/joltify-finance/joltify_lending/dydx_helper/daemons/bridge/api"
	liquidationtypes "github.com/joltify-finance/joltify_lending/dydx_helper/daemons/liquidation/api"
	pricefeedtypes "github.com/joltify-finance/joltify_lending/dydx_helper/daemons/pricefeed/api"
	blocktimetypes "github.com/joltify-finance/joltify_lending/dydx_helper/x/blocktime/types"
	perptypes "github.com/joltify-finance/joltify_lending/dydx_helper/x/perpetuals/types"
	pricetypes "github.com/joltify-finance/joltify_lending/dydx_helper/x/prices/types"
	satypes "github.com/joltify-finance/joltify_lending/dydx_helper/x/subaccounts/types"
	clobtypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/types"
)

// QueryClient combines all the query clients used in testing into a single mock interface for testing convenience.
type QueryClient interface {
	blocktimetypes.QueryClient
	satypes.QueryClient
	clobtypes.QueryClient
	perptypes.QueryClient
	pricetypes.QueryClient
	bridgetypes.BridgeServiceClient
	liquidationtypes.LiquidationServiceClient
	pricefeedtypes.PriceFeedServiceClient
}
