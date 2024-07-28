package keeper

import (
	"fmt"
	"strings"
	"time"

	storetypes "cosmossdk.io/store/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper keeper for the incentive module
type Keeper struct {
	cdc           codec.Codec
	key           storetypes.StoreKey
	paramSubspace types.ParamSubspace
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	joltKeeper    types.JoltKeeper
	swapKeeper    types.SwapKeeper
	spvKeeper     types.SPVKeeper
	NftKeeper     types.NFTKeeper
}

// NewKeeper creates a new keeper
func NewKeeper(
	cdc codec.Codec, key storetypes.StoreKey, paramstore types.ParamSubspace, bk types.BankKeeper,
	joltKeeper types.JoltKeeper, ak types.AccountKeeper, swapKeeper types.SwapKeeper, spvKeeper types.SPVKeeper, nftKeeper types.NFTKeeper,
) Keeper {
	if !paramstore.HasKeyTable() {
		paramstore = paramstore.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		accountKeeper: ak,
		cdc:           cdc,
		key:           key,
		paramSubspace: paramstore,
		bankKeeper:    bk,
		joltKeeper:    joltKeeper,
		swapKeeper:    swapKeeper,
		spvKeeper:     spvKeeper,
		NftKeeper:     nftKeeper,
	}
}

// GetJoltLiquidityProviderClaim returns the claim in the store corresponding the the input address collateral type and id and a boolean for if the claim was found
func (k Keeper) GetJoltLiquidityProviderClaim(ctx sdk.Context, addr sdk.AccAddress) (types.JoltLiquidityProviderClaim, bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.JoltLiquidityClaimKeyPrefix)
	bz := store.Get(addr)
	if bz == nil {
		return types.JoltLiquidityProviderClaim{}, false
	}
	var c types.JoltLiquidityProviderClaim
	k.cdc.MustUnmarshal(bz, &c)
	return c, true
}

// SetJoltLiquidityProviderClaim sets the claim in the store corresponding to the input address, collateral type, and id
func (k Keeper) SetJoltLiquidityProviderClaim(ctx sdk.Context, c types.JoltLiquidityProviderClaim) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.JoltLiquidityClaimKeyPrefix)
	bz := k.cdc.MustMarshal(&c)
	store.Set(c.Owner, bz)
}

// DeleteJoltLiquidityProviderClaim deletes the claim in the store corresponding to the input address, collateral type, and id
func (k Keeper) DeleteJoltLiquidityProviderClaim(rctx sdk.Context, owner sdk.AccAddress) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types.JoltLiquidityClaimKeyPrefix)
	store.Delete(owner)
}

// IterateJoltLiquidityProviderClaims iterates over all claim  objects in the store and preforms a callback function
func (k Keeper) IterateJoltLiquidityProviderClaims(rctx sdk.Context, cb func(c types.JoltLiquidityProviderClaim) (stop bool)) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types.JoltLiquidityClaimKeyPrefix)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var c types.JoltLiquidityProviderClaim
		k.cdc.MustUnmarshal(iterator.Value(), &c)
		if cb(c) {
			break
		}
	}
}

// GetAllJoltLiquidityProviderClaims returns all Claim objects in the store
func (k Keeper) GetAllJoltLiquidityProviderClaims(ctx sdk.Context) types.JoltLiquidityProviderClaims {
	cs := types.JoltLiquidityProviderClaims{}
	k.IterateJoltLiquidityProviderClaims(ctx, func(c types.JoltLiquidityProviderClaim) (stop bool) {
		cs = append(cs, c)
		return false
	})
	return cs
}

// SetJoltSupplyRewardIndexes sets the current reward indexes for an individual denom
func (k Keeper) SetJoltSupplyRewardIndexes(ctx sdk.Context, denom string, indexes types.RewardIndexes) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.JoltSupplyRewardIndexesKeyPrefix)
	bz := k.cdc.MustMarshal(&types.RewardIndexesProto{
		RewardIndexes: indexes,
	})
	store.Set([]byte(denom), bz)
}

// GetJoltSupplyRewardIndexes gets the current reward indexes for an individual denom
func (k Keeper) GetJoltSupplyRewardIndexes(ctx sdk.Context, denom string) (types.RewardIndexes, bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.JoltSupplyRewardIndexesKeyPrefix)
	bz := store.Get([]byte(denom))
	if bz == nil {
		return types.RewardIndexes{}, false
	}
	var proto types.RewardIndexesProto
	k.cdc.MustUnmarshal(bz, &proto)

	return proto.RewardIndexes, true
}

// IterateJoltSupplyRewardIndexes iterates over all Hard supply reward index objects in the store and preforms a callback function
func (k Keeper) IterateJoltSupplyRewardIndexes(ctx sdk.Context, cb func(denom string, indexes types.RewardIndexes) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.JoltSupplyRewardIndexesKeyPrefix)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var proto types.RewardIndexesProto
		k.cdc.MustUnmarshal(iterator.Value(), &proto)
		if cb(string(iterator.Key()), proto.RewardIndexes) {
			break
		}
	}
}

func (k Keeper) IterateJoltSupplyRewardAccrualTimes(ctx sdk.Context, cb func(string, time.Time) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousJoltSupplyRewardAccrualTimeKeyPrefix)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
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
func (k Keeper) SetJoltBorrowRewardIndexes(ctx sdk.Context, denom string, indexes types.RewardIndexes) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.JoltBorrowRewardIndexesKeyPrefix)
	bz := k.cdc.MustMarshal(&types.RewardIndexesProto{
		RewardIndexes: indexes,
	})
	store.Set([]byte(denom), bz)
}

// GetJoltBorrowRewardIndexes gets the current reward indexes for an individual denom
func (k Keeper) GetJoltBorrowRewardIndexes(ctx sdk.Context, denom string) (types.RewardIndexes, bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.JoltBorrowRewardIndexesKeyPrefix)
	bz := store.Get([]byte(denom))
	if bz == nil {
		return types.RewardIndexes{}, false
	}
	var proto types.RewardIndexesProto
	k.cdc.MustUnmarshal(bz, &proto)

	return proto.RewardIndexes, true
}

// IterateJoltBorrowRewardIndexes iterates over all Hard borrow reward index objects in the store and preforms a callback function
func (k Keeper) IterateJoltBorrowRewardIndexes(ctx sdk.Context, cb func(denom string, indexes types.RewardIndexes) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.JoltBorrowRewardIndexesKeyPrefix)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var proto types.RewardIndexesProto
		k.cdc.MustUnmarshal(iterator.Value(), &proto)
		if cb(string(iterator.Key()), proto.RewardIndexes) {
			break
		}
	}
}

func (k Keeper) IterateJoltBorrowRewardAccrualTimes(ctx sdk.Context, cb func(string, time.Time) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousJoltBorrowRewardAccrualTimeKeyPrefix)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
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
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousJoltSupplyRewardAccrualTimeKeyPrefix)
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
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousJoltSupplyRewardAccrualTimeKeyPrefix)
	bz, err := blockTime.MarshalBinary()
	if err != nil {
		panic(err)
	}
	store.Set([]byte(denom), bz)
}

// GetPreviousJoltBorrowRewardAccrualTime returns the last time a denom accrued Hard protocol borrow-side rewards
func (k Keeper) GetPreviousJoltBorrowRewardAccrualTime(ctx sdk.Context, denom string) (blockTime time.Time, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousJoltBorrowRewardAccrualTimeKeyPrefix)
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
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousJoltBorrowRewardAccrualTimeKeyPrefix)
	bz, err := blockTime.MarshalBinary()
	if err != nil {
		panic(err)
	}
	store.Set([]byte(denom), bz)
}

// GetPreviousDelegatorRewardAccrualTime returns the last time a denom accrued protocol delegator rewards
func (k Keeper) GetPreviousDelegatorRewardAccrualTime(ctx sdk.Context, denom string) (blockTime time.Time, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousDelegatorRewardAccrualTimeKeyPrefix)
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
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousDelegatorRewardAccrualTimeKeyPrefix)
	bz, err := blockTime.MarshalBinary()
	if err != nil {
		panic(err)
	}
	store.Set([]byte(denom), bz)
}

// GetSwapClaim returns the claim in the store corresponding the the input address.
func (k Keeper) GetSwapClaim(ctx sdk.Context, addr sdk.AccAddress) (types.SwapClaim, bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SwapClaimKeyPrefix)
	bz := store.Get(addr)
	if bz == nil {
		return types.SwapClaim{}, false
	}
	var c types.SwapClaim
	k.cdc.MustUnmarshal(bz, &c)
	return c, true
}

// SetSwapClaim sets the claim in the store corresponding to the input address.
func (k Keeper) SetSwapClaim(ctx sdk.Context, c types.SwapClaim) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SwapClaimKeyPrefix)
	bz := k.cdc.MustMarshal(&c)
	store.Set(c.Owner, bz)
}

// DeleteSwapClaim deletes the claim in the store corresponding to the input address.
func (k Keeper) DeleteSwapClaim(ctx sdk.Context, owner sdk.AccAddress) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SwapClaimKeyPrefix)
	store.Delete(owner)
}

// IterateSwapClaims iterates over all claim  objects in the store and preforms a callback function
func (k Keeper) IterateSwapClaims(ctx sdk.Context, cb func(c types.SwapClaim) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SwapClaimKeyPrefix)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var c types.SwapClaim
		k.cdc.MustUnmarshal(iterator.Value(), &c)
		if cb(c) {
			break
		}
	}
}

// GetAllSwapClaims returns all Claim objects in the store
func (k Keeper) GetAllSwapClaims(ctx sdk.Context) types.SwapClaims {
	cs := types.SwapClaims{}
	k.IterateSwapClaims(ctx, func(c types.SwapClaim) (stop bool) {
		cs = append(cs, c)
		return false
	})
	return cs
}

// SetSwapRewardIndexes stores the global reward indexes that track total rewards to a swap pool.
func (k Keeper) SetSwapRewardIndexes(ctx sdk.Context, poolID string, indexes types.RewardIndexes) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SwapRewardIndexesKeyPrefix)
	bz := k.cdc.MustMarshal(&types.RewardIndexesProto{
		RewardIndexes: indexes,
	})
	store.Set([]byte(poolID), bz)
}

// GetSwapRewardIndexes fetches the global reward indexes that track total rewards to a swap pool.
func (k Keeper) GetSwapRewardIndexes(ctx sdk.Context, poolID string) (types.RewardIndexes, bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SwapRewardIndexesKeyPrefix)
	bz := store.Get([]byte(poolID))
	if bz == nil {
		return types.RewardIndexes{}, false
	}
	var proto types.RewardIndexesProto
	k.cdc.MustUnmarshal(bz, &proto)
	return proto.RewardIndexes, true
}

// IterateSwapRewardIndexes iterates over all swap reward index objects in the store and preforms a callback function
func (k Keeper) IterateSwapRewardIndexes(ctx sdk.Context, cb func(poolID string, indexes types.RewardIndexes) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SwapRewardIndexesKeyPrefix)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var proto types.RewardIndexesProto
		k.cdc.MustUnmarshal(iterator.Value(), &proto)
		if cb(string(iterator.Key()), proto.RewardIndexes) {
			break
		}
	}
}

// GetSwapRewardAccrualTime fetches the last time rewards were accrued for a swap pool.
func (k Keeper) GetSwapRewardAccrualTime(ctx sdk.Context, poolID string) (blockTime time.Time, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousSwapRewardAccrualTimeKeyPrefix)
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
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousSwapRewardAccrualTimeKeyPrefix)
	bz, err := blockTime.MarshalBinary()
	if err != nil {
		panic(err)
	}
	store.Set([]byte(poolID), bz)
}

func (k Keeper) IterateSwapRewardAccrualTimes(ctx sdk.Context, cb func(string, time.Time) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousSwapRewardAccrualTimeKeyPrefix)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
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

// SetSPVReward stores the global reward indexes that track total rewards to a SPV pool.
func (k Keeper) SetSPVReward(ctx sdk.Context, poolID string, accRewardTokens types.SPVRewardAccTokens) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SPVRewardIndexesKeyPrefix)
	bz := k.cdc.MustMarshal(&types.SPVRewardAccTokens{
		PaymentAmount: accRewardTokens.PaymentAmount,
	})
	incentivePool := types.Incentiveprefix + poolID
	store.Set([]byte(incentivePool), bz)
}

// GetSPVReward fetches the global reward indexes that track total rewards to a SPV pool.
func (k Keeper) GetSPVReward(ctx sdk.Context, poolID string) (types.SPVRewardAccTokens, bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SPVRewardIndexesKeyPrefix)
	incentivePool := types.Incentiveprefix + poolID
	bz := store.Get([]byte(incentivePool))
	if bz == nil {
		return types.SPVRewardAccTokens{}, false
	}
	var accTokens types.SPVRewardAccTokens
	k.cdc.MustUnmarshal(bz, &accTokens)
	return accTokens, true
}

// SetSPVInvestorReward stores the investor reward indexes that track total rewards to a SPV pool.
func (k Keeper) SetSPVInvestorReward(ctx sdk.Context, poolID, walletAddr string, incentiveTokens sdk.Coins) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SPVRewardInvestorKeyPrefix)
	bz := k.cdc.MustMarshal(&types.SPVRewardAccTokens{
		PaymentAmount: incentiveTokens,
	})
	incentivePool := types.Incentiveclassprefix + fmt.Sprintf("%s-%s", poolID, walletAddr)
	store.Set([]byte(incentivePool), bz)
}

// GetSPVInvestorReward fetches the investor reward indexes that track total rewards to a SPV pool.
func (k Keeper) GetSPVInvestorReward(ctx sdk.Context, poolID, walletAddr string) (sdk.Coins, bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SPVRewardInvestorKeyPrefix)
	incentivePool := types.Incentiveclassprefix + fmt.Sprintf("%s-%s", poolID, walletAddr)
	bz := store.Get([]byte(incentivePool))
	if bz == nil {
		return sdk.NewCoins(), false
	}
	var accTokens types.SPVRewardAccTokens
	k.cdc.MustUnmarshal(bz, &accTokens)
	return accTokens.PaymentAmount, true
}

// IterateSPVInvestorRewards iterates over all SPV reward index objects in the store and preforms a callback function
func (k Keeper) IterateSPVInvestorReward(ctx sdk.Context, cb func(key string, accTokens types.SPVRewardAccTokens) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SPVRewardInvestorKeyPrefix)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var accTokens types.SPVRewardAccTokens
		k.cdc.MustUnmarshal(iterator.Value(), &accTokens)
		if cb(string(iterator.Key()), accTokens) {
			break
		}
	}
}

// IterateSPVInvestorRewards iterates over all SPV reward index objects in the store and preforms a callback function
func (k Keeper) LegacyIterateSPVInvestorReward(ctx sdk.Context, cb func(key string, accTokens types.SPVRewardAccTokens) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SPVRewardIndexesKeyPrefix)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		if !strings.Contains(string(iterator.Key()), types.Incentiveclassprefix) {
			continue
		}
		var accTokens types.SPVRewardAccTokens
		k.cdc.MustUnmarshal(iterator.Value(), &accTokens)
		if cb(string(iterator.Key()), accTokens) {
			break
		}
	}
}

func (k Keeper) DeleteSPVInvestorReward(ctx sdk.Context, poolID, walletAddr string) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SPVRewardInvestorKeyPrefix)
	incentivePool := types.Incentiveclassprefix + fmt.Sprintf("%s-%s", poolID, walletAddr)
	store.Delete([]byte(incentivePool))
}

// IterateSPVRewardIndexes iterates over all SPV global reward index objects in the store and preforms a callback function
func (k Keeper) IterateSPVRewardIndexes(ctx sdk.Context, cb func(poolID string, accTokens types.SPVRewardAccTokens) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SPVRewardIndexesKeyPrefix)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var proto types.SPVRewardAccTokens
		k.cdc.MustUnmarshal(iterator.Value(), &proto)
		if cb(string(iterator.Key()), proto) {
			break
		}
	}
}

// GetSPVRewardAccrualTime fetches the last time rewards were accrued for a SPV pool.
func (k Keeper) GetSPVRewardAccrualTime(ctx sdk.Context, poolID string) (blockTime time.Time, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousSPVRewardAccrualTimeKeyPrefix)
	incentivePool := types.Incentiveprefix + poolID
	b := store.Get([]byte(incentivePool))
	if b == nil {
		return time.Time{}, false
	}
	if err := blockTime.UnmarshalBinary(b); err != nil {
		panic(err)
	}
	return blockTime, true
}

// SetSPVRewardAccrualTime stores the last time rewards were accrued for a SPV pool.
func (k Keeper) SetSPVRewardAccrualTime(ctx sdk.Context, poolID string, blockTime time.Time) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousSPVRewardAccrualTimeKeyPrefix)
	bz, err := blockTime.MarshalBinary()
	if err != nil {
		panic(err)
	}
	incentivePool := types.Incentiveprefix + poolID
	store.Set([]byte(incentivePool), bz)
}

func (k Keeper) IterateSPVRewardAccrualTimes(ctx sdk.Context, cb func(string, time.Time) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousSPVRewardAccrualTimeKeyPrefix)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
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
