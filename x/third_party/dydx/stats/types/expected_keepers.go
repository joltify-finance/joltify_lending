package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/epochs/types"
)

// EpochsKeeper defines the expected epochs keeper to get epoch info.
type EpochsKeeper interface {
	MustGetStatsEpochInfo(ctx sdk.Context) types.EpochInfo
}
