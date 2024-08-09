package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/stats/types"
)

// StatsKeeper defines the expected stats keeper
type StatsKeeper interface {
	GetUserStats(ctx sdk.Context, address string) *types.UserStats
	GetGlobalStats(ctx sdk.Context) *types.GlobalStats
}

// VaultKeeper defines the expected vault keeper.
type VaultKeeper interface {
	IsVault(ctx sdk.Context, address string) bool
}
