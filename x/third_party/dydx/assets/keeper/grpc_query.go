package keeper

import (
	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/assets/types"
)

var _ types.QueryServer = Keeper{}
