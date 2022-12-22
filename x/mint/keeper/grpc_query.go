package keeper

import (
	"github.com/joltify-finance/joltify_lending/x/mint/types"
)

var _ types.QueryServer = Keeper{}
