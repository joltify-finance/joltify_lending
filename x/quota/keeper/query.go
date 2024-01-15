package keeper

import (
	"github.com/joltify-finance/joltify_lending/x/quota/types"
)

var _ types.QueryServer = Keeper{}
