package grpc

import (
	bridgetypes "github.com/joltify-finance/joltify_lending/daemons/bridge/api"
	liquidationtypes "github.com/joltify-finance/joltify_lending/daemons/liquidation/api"
	blocktimetypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/blocktime/types"
	clobtypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/types"
	perptypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/perpetuals/types"
	pricetypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/types"
	satypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/subaccounts/types"
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
}
