package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
)

// SetValidators set a specific validator in the store from its index
func (k Keeper) SetValidators(rctx context.Context, index string, validators types.Validators) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorsStoreKey))
	b := k.cdc.MustMarshal(&validators)
	store.Set(types.KeyPrefix(index), b)
}

// GetValidatorsByHeight returns a validators group from its index
func (k Keeper) GetValidatorsByHeight(rctx context.Context, index string) (val types.Validators, found bool) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorsStoreKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// DoGetAllValidators returns all issueToken
func (k Keeper) DoGetAllValidators(rctx context.Context) (list []types.Validators) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorsStoreKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Validators
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// SetStandbyPower set a specific validator in the store from its index
func (k Keeper) SetStandbyPower(rctx context.Context, addr string, powerItem types.StandbyPower) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StandbyPwoerStoreKey))
	b := k.cdc.MustMarshal(&powerItem)
	store.Set(types.KeyPrefix(addr), b)
}

// DelStandbyPower set a specific validator in the store from its index
func (k Keeper) DelStandbyPower(rctx context.Context, addr string) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StandbyPwoerStoreKey))
	store.Delete(types.KeyPrefix(addr))
}

// GetStandbyPower returns a validators group from its index
func (k Keeper) GetStandbyPower(rctx context.Context, addr string) (val types.StandbyPower, found bool) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StandbyPwoerStoreKey))

	b := store.Get(types.KeyPrefix(addr))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// DoGetAllStandbyPower returns all issueToken
func (k Keeper) DoGetAllStandbyPower(rctx context.Context) (list []types.StandbyPower) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StandbyPwoerStoreKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StandbyPower
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
