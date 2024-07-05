package keeper

import (
	"github.com/joltify-finance/joltify_lending/x/burnauction/types"
)

var _ types.QueryServer = Keeper{}
