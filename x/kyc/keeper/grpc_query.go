package keeper

import (
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
)

var _ types.QueryServer = Keeper{}
