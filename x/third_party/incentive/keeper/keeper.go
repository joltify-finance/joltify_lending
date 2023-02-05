package keeper

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"time"

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
	cdpKeeper     types2.CdpKeeper
	joltKeeper    types2.JoltKeeper
}

// NewKeeper creates a new keeper
func NewKeeper(
	cdc codec.Codec, key storetypes.StoreKey, paramstore types2.ParamSubspace, bk types2.BankKeeper,
	cdpk types2.CdpKeeper, hk types2.JoltKeeper, ak types2.AccountKeeper,
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
		cdpKeeper:     cdpk,
		joltKeeper:    hk,
	}
}

// GetUSDXMintingClaim returns the claim in the store corresponding the the input address collateral type and id and a boolean for if the claim was found
func (k Keeper) GetUSDXMintingClaim(ctx sdk.Context, addr sdk.AccAddress) (types2.USDXMintingClaim, bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.USDXMintingClaimKeyPrefix)
	bz := store.Get(addr)
	if bz == nil {
		return types2.USDXMintingClaim{}, false
	}
	var c types2.USDXMintingClaim
	k.cdc.MustUnmarshal(bz, &c)
	return c, true
}

// SetUSDXMintingClaim sets the claim in the store corresponding to the input address, collateral type, and id
func (k Keeper) SetUSDXMintingClaim(ctx sdk.Context, c types2.USDXMintingClaim) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.USDXMintingClaimKeyPrefix)
	bz := k.cdc.MustMarshal(&c)
	store.Set(c.Owner, bz)
}

// DeleteUSDXMintingClaim deletes the claim in the store corresponding to the input address, collateral type, and id
func (k Keeper) DeleteUSDXMintingClaim(ctx sdk.Context, owner sdk.AccAddress) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.USDXMintingClaimKeyPrefix)
	store.Delete(owner)
}

// IterateUSDXMintingClaims iterates over all claim  objects in the store and preforms a callback function
func (k Keeper) IterateUSDXMintingClaims(ctx sdk.Context, cb func(c types2.USDXMintingClaim) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.USDXMintingClaimKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var c types2.USDXMintingClaim
		k.cdc.MustUnmarshal(iterator.Value(), &c)
		if cb(c) {
			break
		}
	}
}

// GetAllUSDXMintingClaims returns all Claim objects in the store
func (k Keeper) GetAllUSDXMintingClaims(ctx sdk.Context) types2.USDXMintingClaims {
	cs := types2.USDXMintingClaims{}
	k.IterateUSDXMintingClaims(ctx, func(c types2.USDXMintingClaim) (stop bool) {
		cs = append(cs, c)
		return false
	})
	return cs
}

// GetPreviousUSDXMintingAccrualTime returns the last time a collateral type accrued USDX minting rewards
func (k Keeper) GetPreviousUSDXMintingAccrualTime(ctx sdk.Context, ctype string) (blockTime time.Time, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.PreviousUSDXMintingRewardAccrualTimeKeyPrefix)
	b := store.Get([]byte(ctype))
	if b == nil {
		return time.Time{}, false
	}
	if err := blockTime.UnmarshalBinary(b); err != nil {
		panic(err)
	}
	return blockTime, true
}

// SetPreviousUSDXMintingAccrualTime sets the last time a collateral type accrued USDX minting rewards
func (k Keeper) SetPreviousUSDXMintingAccrualTime(ctx sdk.Context, ctype string, blockTime time.Time) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.PreviousUSDXMintingRewardAccrualTimeKeyPrefix)
	bz, err := blockTime.MarshalBinary()
	if err != nil {
		panic(err)
	}
	store.Set([]byte(ctype), bz)
}

// IterateUSDXMintingAccrualTimes iterates over all previous USDX minting accrual times and preforms a callback function
func (k Keeper) IterateUSDXMintingAccrualTimes(ctx sdk.Context, cb func(string, time.Time) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.PreviousUSDXMintingRewardAccrualTimeKeyPrefix)
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

// GetUSDXMintingRewardFactor returns the current reward factor for an individual collateral type
func (k Keeper) GetUSDXMintingRewardFactor(ctx sdk.Context, ctype string) (factor sdk.Dec, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.USDXMintingRewardFactorKeyPrefix)
	bz := store.Get([]byte(ctype))
	if bz == nil {
		return sdk.ZeroDec(), false
	}
	if err := factor.Unmarshal(bz); err != nil {
		panic(err)
	}
	return factor, true
}

// SetUSDXMintingRewardFactor sets the current reward factor for an individual collateral type
func (k Keeper) SetUSDXMintingRewardFactor(ctx sdk.Context, ctype string, factor sdk.Dec) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.USDXMintingRewardFactorKeyPrefix)
	bz, err := factor.Marshal()
	if err != nil {
		panic(err)
	}
	store.Set([]byte(ctype), bz)
}

// IterateUSDXMintingRewardFactors iterates over all USDX Minting reward factor objects in the store and preforms a callback function
func (k Keeper) IterateUSDXMintingRewardFactors(ctx sdk.Context, cb func(denom string, factor sdk.Dec) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.USDXMintingRewardFactorKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var factor sdk.Dec
		if err := factor.Unmarshal(iterator.Value()); err != nil {
			panic(err)
		}
		if cb(string(iterator.Key()), factor) {
			break
		}
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

// GetDelegatorClaim returns the claim in the store corresponding the the input address collateral type and id and a boolean for if the claim was found
func (k Keeper) GetDelegatorClaim(ctx sdk.Context, addr sdk.AccAddress) (types2.DelegatorClaim, bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.DelegatorClaimKeyPrefix)
	bz := store.Get(addr)
	if bz == nil {
		return types2.DelegatorClaim{}, false
	}
	var c types2.DelegatorClaim
	k.cdc.MustUnmarshal(bz, &c)
	return c, true
}

// SetDelegatorClaim sets the claim in the store corresponding to the input address, collateral type, and id
func (k Keeper) SetDelegatorClaim(ctx sdk.Context, c types2.DelegatorClaim) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.DelegatorClaimKeyPrefix)
	bz := k.cdc.MustMarshal(&c)
	store.Set(c.Owner, bz)
}

// DeleteDelegatorClaim deletes the claim in the store corresponding to the input address, collateral type, and id
func (k Keeper) DeleteDelegatorClaim(ctx sdk.Context, owner sdk.AccAddress) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.DelegatorClaimKeyPrefix)
	store.Delete(owner)
}

// IterateDelegatorClaims iterates over all claim  objects in the store and preforms a callback function
func (k Keeper) IterateDelegatorClaims(ctx sdk.Context, cb func(c types2.DelegatorClaim) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.DelegatorClaimKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var c types2.DelegatorClaim
		k.cdc.MustUnmarshal(iterator.Value(), &c)
		if cb(c) {
			break
		}
	}
}

// GetAllDelegatorClaims returns all DelegatorClaim objects in the store
func (k Keeper) GetAllDelegatorClaims(ctx sdk.Context) types2.DelegatorClaims {
	cs := types2.DelegatorClaims{}
	k.IterateDelegatorClaims(ctx, func(c types2.DelegatorClaim) (stop bool) {
		cs = append(cs, c)
		return false
	})
	return cs
}

// GetSwapClaim returns the claim in the store corresponding the the input address.
func (k Keeper) GetSwapClaim(ctx sdk.Context, addr sdk.AccAddress) (types2.SwapClaim, bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.SwapClaimKeyPrefix)
	bz := store.Get(addr)
	if bz == nil {
		return types2.SwapClaim{}, false
	}
	var c types2.SwapClaim
	k.cdc.MustUnmarshal(bz, &c)
	return c, true
}

// SetSwapClaim sets the claim in the store corresponding to the input address.
func (k Keeper) SetSwapClaim(ctx sdk.Context, c types2.SwapClaim) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.SwapClaimKeyPrefix)
	bz := k.cdc.MustMarshal(&c)
	store.Set(c.Owner, bz)
}

// DeleteSwapClaim deletes the claim in the store corresponding to the input address.
func (k Keeper) DeleteSwapClaim(ctx sdk.Context, owner sdk.AccAddress) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.SwapClaimKeyPrefix)
	store.Delete(owner)
}

// IterateSwapClaims iterates over all claim  objects in the store and preforms a callback function
func (k Keeper) IterateSwapClaims(ctx sdk.Context, cb func(c types2.SwapClaim) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.SwapClaimKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var c types2.SwapClaim
		k.cdc.MustUnmarshal(iterator.Value(), &c)
		if cb(c) {
			break
		}
	}
}

// GetAllSwapClaims returns all Claim objects in the store
func (k Keeper) GetAllSwapClaims(ctx sdk.Context) types2.SwapClaims {
	cs := types2.SwapClaims{}
	k.IterateSwapClaims(ctx, func(c types2.SwapClaim) (stop bool) {
		cs = append(cs, c)
		return false
	})
	return cs
}

// GetSavingsClaim returns the claim in the store corresponding the the input address.
func (k Keeper) GetSavingsClaim(ctx sdk.Context, addr sdk.AccAddress) (types2.SavingsClaim, bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.SavingsClaimKeyPrefix)
	bz := store.Get(addr)
	if bz == nil {
		return types2.SavingsClaim{}, false
	}
	var c types2.SavingsClaim
	k.cdc.MustUnmarshal(bz, &c)
	return c, true
}

// SetSavingsClaim sets the claim in the store corresponding to the input address.
func (k Keeper) SetSavingsClaim(ctx sdk.Context, c types2.SavingsClaim) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.SavingsClaimKeyPrefix)
	bz := k.cdc.MustMarshal(&c)
	store.Set(c.Owner, bz)
}

// DeleteSavingsClaim deletes the claim in the store corresponding to the input address.
func (k Keeper) DeleteSavingsClaim(ctx sdk.Context, owner sdk.AccAddress) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.SavingsClaimKeyPrefix)
	store.Delete(owner)
}

// IterateSavingsClaims iterates over all savings claim objects in the store and preforms a callback function
func (k Keeper) IterateSavingsClaims(ctx sdk.Context, cb func(c types2.SavingsClaim) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.SavingsClaimKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var c types2.SavingsClaim
		k.cdc.MustUnmarshal(iterator.Value(), &c)
		if cb(c) {
			break
		}
	}
}

// GetAllSavingsClaims returns all savings claim objects in the store
func (k Keeper) GetAllSavingsClaims(ctx sdk.Context) types2.SavingsClaims {
	cs := types2.SavingsClaims{}
	k.IterateSavingsClaims(ctx, func(c types2.SavingsClaim) (stop bool) {
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

// GetDelegatorRewardIndexes gets the current reward indexes for an individual denom
func (k Keeper) GetDelegatorRewardIndexes(ctx sdk.Context, denom string) (types2.RewardIndexes, bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.DelegatorRewardIndexesKeyPrefix)
	bz := store.Get([]byte(denom))
	if bz == nil {
		return types2.RewardIndexes{}, false
	}
	var proto types2.RewardIndexesProto
	k.cdc.MustUnmarshal(bz, &proto)

	return proto.RewardIndexes, true
}

// SetDelegatorRewardIndexes sets the current reward indexes for an individual denom
func (k Keeper) SetDelegatorRewardIndexes(ctx sdk.Context, denom string, indexes types2.RewardIndexes) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.DelegatorRewardIndexesKeyPrefix)
	bz := k.cdc.MustMarshal(&types2.RewardIndexesProto{
		RewardIndexes: indexes,
	})
	store.Set([]byte(denom), bz)
}

// IterateDelegatorRewardIndexes iterates over all delegator reward index objects in the store and preforms a callback function
func (k Keeper) IterateDelegatorRewardIndexes(ctx sdk.Context, cb func(denom string, indexes types2.RewardIndexes) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.DelegatorRewardIndexesKeyPrefix)
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

func (k Keeper) IterateDelegatorRewardAccrualTimes(ctx sdk.Context, cb func(string, time.Time) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.PreviousDelegatorRewardAccrualTimeKeyPrefix)
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

// SetSwapRewardIndexes stores the global reward indexes that track total rewards to a swap pool.
func (k Keeper) SetSwapRewardIndexes(ctx sdk.Context, poolID string, indexes types2.RewardIndexes) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.SwapRewardIndexesKeyPrefix)
	bz := k.cdc.MustMarshal(&types2.RewardIndexesProto{
		RewardIndexes: indexes,
	})
	store.Set([]byte(poolID), bz)
}

// GetSwapRewardIndexes fetches the global reward indexes that track total rewards to a swap pool.
func (k Keeper) GetSwapRewardIndexes(ctx sdk.Context, poolID string) (types2.RewardIndexes, bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.SwapRewardIndexesKeyPrefix)
	bz := store.Get([]byte(poolID))
	if bz == nil {
		return types2.RewardIndexes{}, false
	}
	var proto types2.RewardIndexesProto
	k.cdc.MustUnmarshal(bz, &proto)
	return proto.RewardIndexes, true
}

// IterateSwapRewardIndexes iterates over all swap reward index objects in the store and preforms a callback function
func (k Keeper) IterateSwapRewardIndexes(ctx sdk.Context, cb func(poolID string, indexes types2.RewardIndexes) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.SwapRewardIndexesKeyPrefix)
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

// GetSwapRewardAccrualTime fetches the last time rewards were accrued for a swap pool.
func (k Keeper) GetSwapRewardAccrualTime(ctx sdk.Context, poolID string) (blockTime time.Time, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.PreviousSwapRewardAccrualTimeKeyPrefix)
	b := store.Get([]byte(poolID))
	if b == nil {
		return time.Time{}, false
	}
	if err := blockTime.UnmarshalBinary(b); err != nil {
		panic(err)
	}
	return blockTime, true
}

// SetSwapRewardAccrualTime stores the last time rewards were accrued for a swap pool.
func (k Keeper) SetSwapRewardAccrualTime(ctx sdk.Context, poolID string, blockTime time.Time) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.PreviousSwapRewardAccrualTimeKeyPrefix)
	bz, err := blockTime.MarshalBinary()
	if err != nil {
		panic(err)
	}
	store.Set([]byte(poolID), bz)
}

func (k Keeper) IterateSwapRewardAccrualTimes(ctx sdk.Context, cb func(string, time.Time) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.PreviousSwapRewardAccrualTimeKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		poolID := string(iterator.Key())
		var accrualTime time.Time
		if err := accrualTime.UnmarshalBinary(iterator.Value()); err != nil {
			panic(err)
		}
		if cb(poolID, accrualTime) {
			break
		}
	}
}

// SetSavingsRewardIndexes stores the global reward indexes that rewards for an individual denom type
func (k Keeper) SetSavingsRewardIndexes(ctx sdk.Context, denom string, indexes types2.RewardIndexes) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.SavingsRewardIndexesKeyPrefix)
	bz := k.cdc.MustMarshal(&types2.RewardIndexesProto{
		RewardIndexes: indexes,
	})
	store.Set([]byte(denom), bz)
}

// GetSavingsRewardIndexes fetches the global reward indexes that track rewards for an individual denom type
func (k Keeper) GetSavingsRewardIndexes(ctx sdk.Context, denom string) (types2.RewardIndexes, bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.SavingsRewardIndexesKeyPrefix)
	bz := store.Get([]byte(denom))
	if bz == nil {
		return types2.RewardIndexes{}, false
	}
	var proto types2.RewardIndexesProto
	k.cdc.MustUnmarshal(bz, &proto)
	return proto.RewardIndexes, true
}

// IterateSavingsRewardIndexes iterates over all savings reward index objects in the store and preforms a callback function
func (k Keeper) IterateSavingsRewardIndexes(ctx sdk.Context, cb func(poolID string, indexes types2.RewardIndexes) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.SavingsRewardIndexesKeyPrefix)
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

// GetSavingsRewardAccrualTime fetches the last time rewards were accrued for an individual denom type
func (k Keeper) GetSavingsRewardAccrualTime(ctx sdk.Context, poolID string) (blockTime time.Time, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.PreviousSavingsRewardAccrualTimeKeyPrefix)
	b := store.Get([]byte(poolID))
	if b == nil {
		return time.Time{}, false
	}
	if err := blockTime.UnmarshalBinary(b); err != nil {
		panic(err)
	}
	return blockTime, true
}

// SetSavingsRewardAccrualTime stores the last time rewards were accrued for a savings deposit denom type
func (k Keeper) SetSavingsRewardAccrualTime(ctx sdk.Context, poolID string, blockTime time.Time) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.PreviousSavingsRewardAccrualTimeKeyPrefix)
	bz, err := blockTime.MarshalBinary()
	if err != nil {
		panic(err)
	}
	store.Set([]byte(poolID), bz)
}

// IterateSavingsRewardAccrualTimesiterates over all the previous savings reward accrual times in the store
func (k Keeper) IterateSavingsRewardAccrualTimes(ctx sdk.Context, cb func(string, time.Time) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.PreviousSavingsRewardAccrualTimeKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		poolID := string(iterator.Key())
		var accrualTime time.Time
		if err := accrualTime.UnmarshalBinary(iterator.Value()); err != nil {
			panic(err)
		}
		if cb(poolID, accrualTime) {
			break
		}
	}
}
