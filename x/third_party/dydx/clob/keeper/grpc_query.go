package keeper

import (
	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/clob/types"
)

var _ types.QueryServer = Keeper{}
