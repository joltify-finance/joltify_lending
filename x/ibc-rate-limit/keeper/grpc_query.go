package keeper

import (
	"github.com/joltify-finance/joltify_lending/x/ibc-rate-limit"
	"github.com/joltify-finance/joltify_lending/x/ibc-rate-limit/types"
)

var _ types.QueryServer = ibc_rate_limit.ICS4Wrapper{}
