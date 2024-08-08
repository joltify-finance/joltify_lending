package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// QuotaKeeper defines the expected bank keeper
type QuotaKeeper interface {
	WhetherOnwhitelist(ctx context.Context, moduleName, sender string) bool
	UpdateQuota(ctx context.Context, coins sdk.Coins, sender string, ibcSeq uint64, moduleName string) error
	RevokeHistory(ctx context.Context, moduleName string, seq uint64)
	WhetherOnBanlist(ctx context.Context, moduleName, sender string) bool
}
