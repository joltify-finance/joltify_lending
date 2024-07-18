package keeper

import (
	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/listing/types"
)

var _ types.QueryServer = Keeper{}
