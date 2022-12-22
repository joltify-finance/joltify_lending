package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
)

// SetStoreFeeAmount set a specific outboundTx in the store from its index
func (k Keeper) SetStoreFeeAmount(ctx sdk.Context, fees sdk.Coins) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeStoreKey))
	for _, el := range fees {
		each := el
		b := k.cdc.MustMarshal(&each)
		store.Set(types.OutboundTxKey(
			el.Denom,
		), b)
	}
}

// GetFeeAmount returns a outboundTx from its index
func (k Keeper) GetFeeAmount(
	ctx sdk.Context,
	denom string,
) (fee sdk.Coin, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeStoreKey))
	b := store.Get(types.OutboundTxKey(
		denom,
	))
	if b == nil {
		return fee, false
	}

	k.cdc.MustUnmarshal(b, &fee)
	return fee, true
}

// GetAllFeeAmount returns all outboundTx
func (k Keeper) GetAllFeeAmount(ctx sdk.Context) (fees sdk.Coins) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeStoreKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var val sdk.Coin
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		fees = append(fees, val)
	}
	return
}
