package grpc

import pricetypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/types"

type QueryServer interface {
	pricetypes.QueryServer
}
