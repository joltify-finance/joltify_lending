package keeper

import (
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

var _ types.QueryServer = Keeper{}
