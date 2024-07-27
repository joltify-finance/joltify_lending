package keeper

import (
	"context"
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
func (k Keeper) GetJoltLiquidityProviderClaim(ctx context.Context, addr sdk.AccAddress) (types.JoltLiquidityProviderClaim, bool) {
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
func (k Keeper) SetJoltLiquidityProviderClaim(ctx context.Context, c types.JoltLiquidityProviderClaim) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.JoltLiquidityClaimKeyPrefix)
	bz := k.cdc.MustMarshal(&c)
	store.Set(c.Owner, bz)
}

// DeleteJoltLiquidityProviderClaim deletes the claim in the store corresponding to the input address, collateral type, and id
func (k Keeper) DeleteJoltLiquidityProviderClaim(rctx context.Context, owner sdk.AccAddress) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types.JoltLiquidityClaimKeyPrefix)
	store.Delete(owner)
}

// IterateJoltLiquidityProviderClaims iterates over all claim  objects in the store and preforms a callback function
func (k Keeper) IterateJoltLiquidityProviderClaims(rctx context.Context, cb func(c types.JoltLiquidityProviderClaim) (stop bool)) {
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
func (k Keeper) GetAllJoltLiquidityProviderClaims(ctx context.Context) types.JoltLiquidityProviderClaims {
	cs := types.JoltLiquidityProviderClaims{}
	k.IterateJoltLiquidityProviderClaims(ctx, func(c types.JoltLiquidityProviderClaim) (stop bool) {
		cs = append(cs, c)
		return false
	})
	return cs
}

// SetJoltSupplyRewardIndexes sets the current reward indexes for an individual denom
func (k Keeper) SetJoltSupplyRewardIndexes(ctx context.Context, denom string, indexes types.RewardIndexes) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.JoltSupplyRewardIndexesKeyPrefix)
	bz := k.cdc.MustMarshal(&types.RewardIndexesProto{
		RewardIndexes: indexes,
	})
	store.Set([]byte(denom), bz)
}

// GetJoltSupplyRewardIndexes gets the current reward indexes for an individual denom
func (k Keeper) GetJoltSupplyRewardIndexes(ctx context.Context, denom string) (types.RewardIndexes, bool) {
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
func (k Keeper) IterateJoltSupplyRewardIndexes(ctx context.Context, cb func(denom string, indexes types.RewardIndexes) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.JoltSupplyRewardIndexesKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var proto types.RewardIndexesProto
		k.cdc.MustUnmarshal(iterator.Value(), &proto)
		if cb(string(iterator.Key()), proto.RewardIndexes) {
			break
		}
	}
}

func (k Keeper) IterateJoltSupplyRewardAccrualTimes(ctx context.Context, cb func(string, time.Time) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousJoltSupplyRewardAccrualTimeKeyPrefix)
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
func (k Keeper) SetJoltBorrowRewardIndexes(ctx context.Context, denom string, indexes types.RewardIndexes) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.JoltBorrowRewardIndexesKeyPrefix)
	bz := k.cdc.MustMarshal(&types.RewardIndexesProto{
		RewardIndexes: indexes,
	})
	store.Set([]byte(denom), bz)
}

// GetJoltBorrowRewardIndexes gets the current reward indexes for an individual denom
func (k Keeper) GetJoltBorrowRewardIndexes(ctx context.Context, denom string) (types.RewardIndexes, bool) {
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
func (k Keeper) IterateJoltBorrowRewardIndexes(ctx context.Context, cb func(denom string, indexes types.RewardIndexes) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.JoltBorrowRewardIndexesKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var proto types.RewardIndexesProto
		k.cdc.MustUnmarshal(iterator.Value(), &proto)
		if cb(string(iterator.Key()), proto.RewardIndexes) {
			break
		}
	}
}

func (k Keeper) IterateJoltBorrowRewardAccrualTimes(ctx context.Context, cb func(string, time.Time) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousJoltBorrowRewardAccrualTimeKeyPrefix)
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
func (k Keeper) GetPreviousJoltSupplyRewardAccrualTime(ctx context.Context, denom string) (blockTime time.Time, found bool) {
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
func (k Keeper) SetPreviousJoltSupplyRewardAccrualTime(ctx context.Context, denom string, blockTime time.Time) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousJoltSupplyRewardAccrualTimeKeyPrefix)
	bz, err := blockTime.MarshalBinary()
	if err != nil {
		panic(err)
	}
	store.Set([]byte(denom), bz)
}

// GetPreviousJoltBorrowRewardAccrualTime returns the last time a denom accrued Hard protocol borrow-side rewards
func (k Keeper) GetPreviousJoltBorrowRewardAccrualTime(ctx context.Context, denom string) (blockTime time.Time, found bool) {
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
func (k Keeper) SetPreviousJoltBorrowRewardAccrualTime(ctx context.Context, denom string, blockTime time.Time) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousJoltBorrowRewardAccrualTimeKeyPrefix)
	bz, err := blockTime.MarshalBinary()
	if err != nil {
		panic(err)
	}
	store.Set([]byte(denom), bz)
}

// GetPreviousDelegatorRewardAccrualTime returns the last time a denom accrued protocol delegator rewards
func (k Keeper) GetPreviousDelegatorRewardAccrualTime(ctx context.Context, denom string) (blockTime time.Time, found bool) {
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
func (k Keeper) SetPreviousDelegatorRewardAccrualTime(ctx context.Context, denom string, blockTime time.Time) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousDelegatorRewardAccrualTimeKeyPrefix)
	bz, err := blockTime.MarshalBinary()
	if err != nil {
		panic(err)
	}
	store.Set([]byte(denom), bz)
}

// GetSwapClaim returns the claim in the store corresponding the the input address.
func (k Keeper) GetSwapClaim(ctx context.Context, addr sdk.AccAddress) (types.SwapClaim, bool) {
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
func (k Keeper) SetSwapClaim(ctx context.Context, c types.SwapClaim) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SwapClaimKeyPrefix)
	bz := k.cdc.MustMarshal(&c)
	store.Set(c.Owner, bz)
}

// DeleteSwapClaim deletes the claim in the store corresponding to the input address.
func (k Keeper) DeleteSwapClaim(ctx context.Context, owner sdk.AccAddress) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SwapClaimKeyPrefix)
	store.Delete(owner)
}

// IterateSwapClaims iterates over all claim  objects in the store and preforms a callback function
func (k Keeper) IterateSwapClaims(ctx context.Context, cb func(c types.SwapClaim) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SwapClaimKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
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
func (k Keeper) GetAllSwapClaims(ctx context.Context) types.SwapClaims {
	cs := types.SwapClaims{}
	k.IterateSwapClaims(ctx, func(c types.SwapClaim) (stop bool) {
		cs = append(cs, c)
		return false
	})
	return cs
}

// SetSwapRewardIndexes stores the global reward indexes that track total rewards to a swap pool.
func (k Keeper) SetSwapRewardIndexes(ctx context.Context, poolID string, indexes types.RewardIndexes) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SwapRewardIndexesKeyPrefix)
	bz := k.cdc.MustMarshal(&types.RewardIndexesProto{
		RewardIndexes: indexes,
	})
	store.Set([]byte(poolID), bz)
}

// GetSwapRewardIndexes fetches the global reward indexes that track total rewards to a swap pool.
func (k Keeper) GetSwapRewardIndexes(ctx context.Context, poolID string) (types.RewardIndexes, bool) {
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
func (k Keeper) IterateSwapRewardIndexes(ctx context.Context, cb func(poolID string, indexes types.RewardIndexes) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SwapRewardIndexesKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
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
func (k Keeper) GetSwapRewardAccrualTime(ctx context.Context, poolID string) (blockTime time.Time, found bool) {
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
func (k Keeper) SetSwapRewardAccrualTime(ctx context.Context, poolID string, blockTime time.Time) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousSwapRewardAccrualTimeKeyPrefix)
	bz, err := blockTime.MarshalBinary()
	if err != nil {
		panic(err)
	}
	store.Set([]byte(poolID), bz)
}

func (k Keeper) IterateSwapRewardAccrualTimes(ctx context.Context, cb func(string, time.Time) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousSwapRewardAccrualTimeKeyPrefix)
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

// SetSPVReward stores the global reward indexes that track total rewards to a SPV pool.
func (k Keeper) SetSPVReward(ctx context.Context, poolID string, accRewardTokens types.SPVRewardAccTokens) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SPVRewardIndexesKeyPrefix)
	bz := k.cdc.MustMarshal(&types.SPVRewardAccTokens{
		PaymentAmount: accRewardTokens.PaymentAmount,
	})
	incentivePool := types.Incentiveprefix + poolID
	store.Set([]byte(incentivePool), bz)
}

// GetSPVReward fetches the global reward indexes that track total rewards to a SPV pool.
func (k Keeper) GetSPVReward(ctx context.Context, poolID string) (types.SPVRewardAccTokens, bool) {
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
func (k Keeper) SetSPVInvestorReward(ctx context.Context, poolID, walletAddr string, incentiveTokens sdk.Coins) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SPVRewardInvestorKeyPrefix)
	bz := k.cdc.MustMarshal(&types.SPVRewardAccTokens{
		PaymentAmount: incentiveTokens,
	})
	incentivePool := types.Incentiveclassprefix + fmt.Sprintf("%s-%s", poolID, walletAddr)
	store.Set([]byte(incentivePool), bz)
}

// GetSPVInvestorReward fetches the investor reward indexes that track total rewards to a SPV pool.
func (k Keeper) GetSPVInvestorReward(ctx context.Context, poolID, walletAddr string) (sdk.Coins, bool) {
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
func (k Keeper) IterateSPVInvestorReward(ctx context.Context, cb func(key string, accTokens types.SPVRewardAccTokens) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SPVRewardInvestorKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
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
func (k Keeper) LegacyIterateSPVInvestorReward(ctx context.Context, cb func(key string, accTokens types.SPVRewardAccTokens) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SPVRewardIndexesKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
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

func (k Keeper) DeleteSPVInvestorReward(ctx context.Context, poolID, walletAddr string) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SPVRewardInvestorKeyPrefix)
	incentivePool := types.Incentiveclassprefix + fmt.Sprintf("%s-%s", poolID, walletAddr)
	store.Delete([]byte(incentivePool))
}

// IterateSPVRewardIndexes iterates over all SPV global reward index objects in the store and preforms a callback function
func (k Keeper) IterateSPVRewardIndexes(ctx context.Context, cb func(poolID string, accTokens types.SPVRewardAccTokens) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.SPVRewardIndexesKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
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
func (k Keeper) GetSPVRewardAccrualTime(ctx context.Context, poolID string) (blockTime time.Time, found bool) {
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
func (k Keeper) SetSPVRewardAccrualTime(ctx context.Context, poolID string, blockTime time.Time) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousSPVRewardAccrualTimeKeyPrefix)
	bz, err := blockTime.MarshalBinary()
	if err != nil {
		panic(err)
	}
	incentivePool := types.Incentiveprefix + poolID
	store.Set([]byte(incentivePool), bz)
}

func (k Keeper) IterateSPVRewardAccrualTimes(ctx context.Context, cb func(string, time.Time) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousSPVRewardAccrualTimeKeyPrefix)
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
