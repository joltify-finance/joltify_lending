package grpc

import pricetypes "github.com/joltify-finance/joltify_lending/dydx_helper/x/prices/types"

type QueryServer interface {
	pricetypes.QueryServer
}
