package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
)

// SetValidators set a specific validator in the store from its index
func (k Keeper) SetValidators(ctx sdk.Context, index string, validators types.Validators) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorsStoreKey))
	b := k.cdc.MustMarshal(&validators)
	store.Set(types.KeyPrefix(index), b)
}

// GetValidatorsByHeight returns a validators group from its index
func (k Keeper) GetValidatorsByHeight(ctx sdk.Context, index string) (val types.Validators, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorsStoreKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// DoGetAllValidators returns all issueToken
func (k Keeper) DoGetAllValidators(ctx sdk.Context) (list []types.Validators) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorsStoreKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Validators
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// SetStandbyPower set a specific validator in the store from its index
func (k Keeper) SetStandbyPower(ctx sdk.Context, addr string, powerItem types.StandbyPower) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StandbyPwoerStoreKey))
	b := k.cdc.MustMarshal(&powerItem)
	store.Set(types.KeyPrefix(addr), b)
}

// DelStandbyPower set a specific validator in the store from its index
func (k Keeper) DelStandbyPower(ctx sdk.Context, addr string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StandbyPwoerStoreKey))
	store.Delete(types.KeyPrefix(addr))
}

// GetStandbyPower returns a validators group from its index
func (k Keeper) GetStandbyPower(ctx sdk.Context, addr string) (val types.StandbyPower, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StandbyPwoerStoreKey))

	b := store.Get(types.KeyPrefix(addr))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// DoGetAllStandbyPower returns all issueToken
func (k Keeper) DoGetAllStandbyPower(ctx sdk.Context) (list []types.StandbyPower) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StandbyPwoerStoreKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StandbyPower
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
