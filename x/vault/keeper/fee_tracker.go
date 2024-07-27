package keeper

import (
	"context"
	"encoding/json"
	"fmt"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
)

// SetStoreFeeAmount set a specific outboundTx in the store from its index
func (k Keeper) SetStoreFeeAmount(rctx context.Context, fees sdk.Coins) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeStoreKey))
	feeBytes, err := fees.MarshalJSON()
	if err != nil {
		panic("marshal coins failed")
	}
	store.Set(types.OutboundTxKey("-fee"), feeBytes)
}

// GetFeeAmount returns a outboundTx from its index
func (k Keeper) GetFeeAmount(
	ctx context.Context,
	denom string,
) (fee sdk.Coin, found bool) {
	fees := k.GetAllFeeAmount(ctx)
	amount := fees.AmountOf(denom)
	return sdk.NewCoin(denom, amount), true
}

func (k Keeper) LegacyGetAllFeeAMountAndDelete(rctx context.Context) sdk.Coins {
	ctx := sdk.UnwrapSDKContext(rctx)
	var fees sdk.Coins
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeStoreKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	var deleteDenom []string
	for ; iterator.Valid(); iterator.Next() {
		var val sdk.Coin
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		fees = append(fees, val)
	}
	for _, el := range deleteDenom {
		store.Delete(types.OutboundTxKey(el))
	}
	err := iterator.Close()
	if err != nil {
		msg := fmt.Errorf("fail to close the iterator %w", err)
		panic(msg)
	}
	return fees
}

// GetAllFeeAmount returns all outboundTx
func (k Keeper) GetAllFeeAmount(rctx context.Context) sdk.Coins {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeStoreKey))
	result := store.Get(types.OutboundTxKey("-fee"))
	var fees sdk.Coins
	err := json.Unmarshal(result, &fees)
	if err != nil {
		panic("fail to get the fee")
	}
	return fees
}
