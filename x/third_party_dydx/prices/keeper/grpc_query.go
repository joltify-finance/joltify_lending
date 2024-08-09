package keeper

import (
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/types"
)

var _ types.QueryServer = Keeper{}
