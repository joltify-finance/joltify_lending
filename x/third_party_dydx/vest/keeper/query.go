package keeper

import (
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/vest/types"
)

var _ types.QueryServer = Keeper{}
