package keeper

import (
	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
)

// SetQuotaData set a specific createPool in the store from its index
func (k Keeper) SetQuotaData(ctx context.Context, coinsQuota types.CoinsQuota) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.QuotaStoreKey))
	b := k.cdc.MustMarshal(&coinsQuota)

	store.Set(types.KeyPrefix("info"), b)
}

// GetQuotaData returns a createPool from its index
func (k Keeper) GetQuotaData(ctx context.Context) (val types.CoinsQuota, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.QuotaStoreKey))

	b := store.Get(types.KeyPrefix("info"))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
