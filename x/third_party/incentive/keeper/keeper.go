package keeper

import (
	"time"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper keeper for the incentive module
type Keeper struct {
	cdc           codec.Codec
	key           storetypes.StoreKey
	paramSubspace types2.ParamSubspace
	accountKeeper types2.AccountKeeper
	bankKeeper    types2.BankKeeper
	joltKeeper    types2.JoltKeeper
}

// NewKeeper creates a new keeper
func NewKeeper(
	cdc codec.Codec, key storetypes.StoreKey, paramstore types2.ParamSubspace, bk types2.BankKeeper,
	joltKeeper types2.JoltKeeper, ak types2.AccountKeeper,
) Keeper {
	if !paramstore.HasKeyTable() {
		paramstore = paramstore.WithKeyTable(types2.ParamKeyTable())
	}

	return Keeper{
		accountKeeper: ak,
		cdc:           cdc,
		key:           key,
		paramSubspace: paramstore,
		bankKeeper:    bk,
		joltKeeper:    joltKeeper,
	}
}

// GetJoltLiquidityProviderClaim returns the claim in the store corresponding the the input address collateral type and id and a boolean for if the claim was found
func (k Keeper) GetJoltLiquidityProviderClaim(ctx sdk.Context, addr sdk.AccAddress) (types2.JoltLiquidityProviderClaim, bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.JoltLiquidityClaimKeyPrefix)
	bz := store.Get(addr)
	if bz == nil {
		return types2.JoltLiquidityProviderClaim{}, false
	}
	var c types2.JoltLiquidityProviderClaim
	k.cdc.MustUnmarshal(bz, &c)
	return c, true
}

// SetJoltLiquidityProviderClaim sets the claim in the store corresponding to the input address, collateral type, and id
func (k Keeper) SetJoltLiquidityProviderClaim(ctx sdk.Context, c types2.JoltLiquidityProviderClaim) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.JoltLiquidityClaimKeyPrefix)
	bz := k.cdc.MustMarshal(&c)
	store.Set(c.Owner, bz)
}

// DeleteJoltLiquidityProviderClaim deletes the claim in the store corresponding to the input address, collateral type, and id
func (k Keeper) DeleteJoltLiquidityProviderClaim(ctx sdk.Context, owner sdk.AccAddress) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.JoltLiquidityClaimKeyPrefix)
	store.Delete(owner)
}

// IterateJoltLiquidityProviderClaims iterates over all claim  objects in the store and preforms a callback function
func (k Keeper) IterateJoltLiquidityProviderClaims(ctx sdk.Context, cb func(c types2.JoltLiquidityProviderClaim) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.JoltLiquidityClaimKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var c types2.JoltLiquidityProviderClaim
		k.cdc.MustUnmarshal(iterator.Value(), &c)
		if cb(c) {
			break
		}
	}
}

// GetAllJoltLiquidityProviderClaims returns all Claim objects in the store
func (k Keeper) GetAllJoltLiquidityProviderClaims(ctx sdk.Context) types2.JoltLiquidityProviderClaims {
	cs := types2.JoltLiquidityProviderClaims{}
	k.IterateJoltLiquidityProviderClaims(ctx, func(c types2.JoltLiquidityProviderClaim) (stop bool) {
		cs = append(cs, c)
		return false
	})
	return cs
}

// SetJoltSupplyRewardIndexes sets the current reward indexes for an individual denom
func (k Keeper) SetJoltSupplyRewardIndexes(ctx sdk.Context, denom string, indexes types2.RewardIndexes) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.JoltSupplyRewardIndexesKeyPrefix)
	bz := k.cdc.MustMarshal(&types2.RewardIndexesProto{
		RewardIndexes: indexes,
	})
	store.Set([]byte(denom), bz)
}

// GetJoltSupplyRewardIndexes gets the current reward indexes for an individual denom
func (k Keeper) GetJoltSupplyRewardIndexes(ctx sdk.Context, denom string) (types2.RewardIndexes, bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.JoltSupplyRewardIndexesKeyPrefix)
	bz := store.Get([]byte(denom))
	if bz == nil {
		return types2.RewardIndexes{}, false
	}
	var proto types2.RewardIndexesProto
	k.cdc.MustUnmarshal(bz, &proto)

	return proto.RewardIndexes, true
}

// IterateJoltSupplyRewardIndexes iterates over all Hard supply reward index objects in the store and preforms a callback function
func (k Keeper) IterateJoltSupplyRewardIndexes(ctx sdk.Context, cb func(denom string, indexes types2.RewardIndexes) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.JoltSupplyRewardIndexesKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var proto types2.RewardIndexesProto
		k.cdc.MustUnmarshal(iterator.Value(), &proto)
		if cb(string(iterator.Key()), proto.RewardIndexes) {
			break
		}
	}
}

func (k Keeper) IterateJoltSupplyRewardAccrualTimes(ctx sdk.Context, cb func(string, time.Time) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.PreviousJoltSupplyRewardAccrualTimeKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var accrualTime time.Time
		if err := accrualTime.UnmarshalBinary(iterator.Value()); err != nil {
			panic(err)
		}
		denom := string(iterator.Key())
		if cb(denom, accrualTime) {
			break
		}
	}
}

// SetJoltBorrowRewardIndexes sets the current reward indexes for an individual denom
func (k Keeper) SetJoltBorrowRewardIndexes(ctx sdk.Context, denom string, indexes types2.RewardIndexes) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.JoltBorrowRewardIndexesKeyPrefix)
	bz := k.cdc.MustMarshal(&types2.RewardIndexesProto{
		RewardIndexes: indexes,
	})
	store.Set([]byte(denom), bz)
}

// GetJoltBorrowRewardIndexes gets the current reward indexes for an individual denom
func (k Keeper) GetJoltBorrowRewardIndexes(ctx sdk.Context, denom string) (types2.RewardIndexes, bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.JoltBorrowRewardIndexesKeyPrefix)
	bz := store.Get([]byte(denom))
	if bz == nil {
		return types2.RewardIndexes{}, false
	}
	var proto types2.RewardIndexesProto
	k.cdc.MustUnmarshal(bz, &proto)

	return proto.RewardIndexes, true
}

// IterateJoltBorrowRewardIndexes iterates over all Hard borrow reward index objects in the store and preforms a callback function
func (k Keeper) IterateJoltBorrowRewardIndexes(ctx sdk.Context, cb func(denom string, indexes types2.RewardIndexes) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.JoltBorrowRewardIndexesKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var proto types2.RewardIndexesProto
		k.cdc.MustUnmarshal(iterator.Value(), &proto)
		if cb(string(iterator.Key()), proto.RewardIndexes) {
			break
		}
	}
}

func (k Keeper) IterateJoltBorrowRewardAccrualTimes(ctx sdk.Context, cb func(string, time.Time) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.PreviousJoltBorrowRewardAccrualTimeKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		denom := string(iterator.Key())
		var accrualTime time.Time
		if err := accrualTime.UnmarshalBinary(iterator.Value()); err != nil {
			panic(err)
		}
		if cb(denom, accrualTime) {
			break
		}
	}
}

// GetPreviousJoltSupplyRewardAccrualTime returns the last time a denom accrued Hard protocol supply-side rewards
func (k Keeper) GetPreviousJoltSupplyRewardAccrualTime(ctx sdk.Context, denom string) (blockTime time.Time, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.PreviousJoltSupplyRewardAccrualTimeKeyPrefix)
	bz := store.Get([]byte(denom))
	if bz == nil {
		return time.Time{}, false
	}
	if err := blockTime.UnmarshalBinary(bz); err != nil {
		panic(err)
	}
	return blockTime, true
}

// SetPreviousJoltSupplyRewardAccrualTime sets the last time a denom accrued Hard protocol supply-side rewards
func (k Keeper) SetPreviousJoltSupplyRewardAccrualTime(ctx sdk.Context, denom string, blockTime time.Time) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.PreviousJoltSupplyRewardAccrualTimeKeyPrefix)
	bz, err := blockTime.MarshalBinary()
	if err != nil {
		panic(err)
	}
	store.Set([]byte(denom), bz)
}

// GetPreviousJoltBorrowRewardAccrualTime returns the last time a denom accrued Hard protocol borrow-side rewards
func (k Keeper) GetPreviousJoltBorrowRewardAccrualTime(ctx sdk.Context, denom string) (blockTime time.Time, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.PreviousJoltBorrowRewardAccrualTimeKeyPrefix)
	b := store.Get([]byte(denom))
	if b == nil {
		return time.Time{}, false
	}
	if err := blockTime.UnmarshalBinary(b); err != nil {
		panic(err)
	}
	return blockTime, true
}

// SetPreviousJoltBorrowRewardAccrualTime sets the last time a denom accrued Hard protocol borrow-side rewards
func (k Keeper) SetPreviousJoltBorrowRewardAccrualTime(ctx sdk.Context, denom string, blockTime time.Time) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.PreviousJoltBorrowRewardAccrualTimeKeyPrefix)
	bz, err := blockTime.MarshalBinary()
	if err != nil {
		panic(err)
	}
	store.Set([]byte(denom), bz)
}

// GetPreviousDelegatorRewardAccrualTime returns the last time a denom accrued protocol delegator rewards
func (k Keeper) GetPreviousDelegatorRewardAccrualTime(ctx sdk.Context, denom string) (blockTime time.Time, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.PreviousDelegatorRewardAccrualTimeKeyPrefix)
	bz := store.Get([]byte(denom))
	if bz == nil {
		return time.Time{}, false
	}
	if err := blockTime.UnmarshalBinary(bz); err != nil {
		panic(err)
	}
	return blockTime, true
}

// SetPreviousDelegatorRewardAccrualTime sets the last time a denom accrued protocol delegator rewards
func (k Keeper) SetPreviousDelegatorRewardAccrualTime(ctx sdk.Context, denom string, blockTime time.Time) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.PreviousDelegatorRewardAccrualTimeKeyPrefix)
	bz, err := blockTime.MarshalBinary()
	if err != nil {
		panic(err)
	}
	store.Set([]byte(denom), bz)
}
