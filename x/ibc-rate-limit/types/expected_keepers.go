package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// QuotaKeeper defines the expected bank keeper
type QuotaKeeper interface {
	WhetherOnwhitelist(ctx sdk.Context, moduleName, sender string) bool
	UpdateQuota(ctx sdk.Context, coins sdk.Coins, ibcSeq uint64, moduleName string) error
	RevokeHistory(ctx sdk.Context, moduleName string, seq uint64)
}
