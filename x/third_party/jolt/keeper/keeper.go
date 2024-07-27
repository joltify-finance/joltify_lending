package keeper

import (
	"context"
	"time"

	sdkmath "cosmossdk.io/math"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Keeper keeper for the jolt module
type Keeper struct {
	key             storetypes.StoreKey
	cdc             codec.Codec
	paramSubspace   paramtypes.Subspace
	accountKeeper   types2.AccountKeeper
	bankKeeper      types2.BankKeeper
	pricefeedKeeper types2.PricefeedKeeper
	auctionKeeper   types2.AuctionKeeper
	hooks           types2.JOLTHooks
}

// NewKeeper creates a new keeper
func NewKeeper(cdc codec.Codec, key storetypes.StoreKey, paramstore paramtypes.Subspace,
	ak types2.AccountKeeper, bk types2.BankKeeper,
	pfk types2.PricefeedKeeper, auk types2.AuctionKeeper,
) Keeper {
	if !paramstore.HasKeyTable() {
		paramstore = paramstore.WithKeyTable(types2.ParamKeyTable())
	}

	return Keeper{
		key:             key,
		cdc:             cdc,
		paramSubspace:   paramstore,
		accountKeeper:   ak,
		bankKeeper:      bk,
		pricefeedKeeper: pfk,
		auctionKeeper:   auk,
		hooks:           nil,
	}
}

// SetHooks adds hooks to the keeper.
func (k *Keeper) SetHooks(hooks types2.JOLTHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set jolt hooks twice")
	}
	k.hooks = hooks
	return k
}

// GetDeposit returns a deposit from the store for a particular depositor address, deposit denom
func (k Keeper) GetDeposit(ctx context.Context, depositor sdk.AccAddress) (types2.Deposit, bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.DepositsKeyPrefix)
	bz := store.Get(depositor.Bytes())
	if len(bz) == 0 {
		return types2.Deposit{}, false
	}
	var deposit types2.Deposit
	k.cdc.MustUnmarshal(bz, &deposit)
	return deposit, true
}

// SetDeposit sets the input deposit in the store, prefixed by the deposit type, deposit denom, and depositor address, in that order
func (k Keeper) SetDeposit(rctx context.Context, deposit types2.Deposit) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types2.DepositsKeyPrefix)
	bz := k.cdc.MustMarshal(&deposit)
	store.Set(deposit.Depositor.Bytes(), bz)
}

// DeleteDeposit deletes a deposit from the store
func (k Keeper) DeleteDeposit(rctx context.Context, deposit types2.Deposit) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types2.DepositsKeyPrefix)
	store.Delete(deposit.Depositor.Bytes())
}

// IterateDeposits iterates over all deposit objects in the store and performs a callback function
func (k Keeper) IterateDeposits(rctx context.Context, cb func(deposit types2.Deposit) (stop bool)) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types2.DepositsKeyPrefix)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var deposit types2.Deposit
		k.cdc.MustUnmarshal(iterator.Value(), &deposit)
		if cb(deposit) {
			break
		}
	}
}

// GetDepositsByUser gets all deposits for an individual user
func (k Keeper) GetDepositsByUser(ctx context.Context, user sdk.AccAddress) []types2.Deposit {
	var deposits []types2.Deposit
	k.IterateDeposits(ctx, func(deposit types2.Deposit) (stop bool) {
		if deposit.Depositor.Equals(user) {
			deposits = append(deposits, deposit)
		}
		return false
	})
	return deposits
}

// GetBorrow returns a Borrow from the store for a particular borrower address and borrow denom
func (k Keeper) GetBorrow(rctx context.Context, borrower sdk.AccAddress) (types2.Borrow, bool) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types2.BorrowsKeyPrefix)
	bz := store.Get(borrower)
	if len(bz) == 0 {
		return types2.Borrow{}, false
	}
	var borrow types2.Borrow
	k.cdc.MustUnmarshal(bz, &borrow)
	return borrow, true
}

// SetBorrow sets the input borrow in the store, prefixed by the borrower address and borrow denom
func (k Keeper) SetBorrow(rctx context.Context, borrow types2.Borrow) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types2.BorrowsKeyPrefix)
	bz := k.cdc.MustMarshal(&borrow)
	store.Set(borrow.Borrower, bz)
}

// DeleteBorrow deletes a borrow from the store
func (k Keeper) DeleteBorrow(rctx context.Context, borrow types2.Borrow) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types2.BorrowsKeyPrefix)
	store.Delete(borrow.Borrower)
}

// IterateBorrows iterates over all borrow objects in the store and performs a callback function
func (k Keeper) IterateBorrows(rctx context.Context, cb func(borrow types2.Borrow) (stop bool)) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types2.BorrowsKeyPrefix)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var borrow types2.Borrow
		k.cdc.MustUnmarshal(iterator.Value(), &borrow)
		if cb(borrow) {
			break
		}
	}
}

// SetBorrowedCoins sets the total amount of coins currently borrowed in the store
func (k Keeper) SetBorrowedCoins(rctx context.Context, borrowedCoins sdk.Coins) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types2.BorrowedCoinsPrefix)
	if borrowedCoins.Empty() {
		store.Set(types2.BorrowedCoinsPrefix, []byte{})
	} else {
		bz := k.cdc.MustMarshal(&types2.CoinsProto{
			Coins: borrowedCoins,
		})
		store.Set(types2.BorrowedCoinsPrefix, bz)
	}
}

// GetBorrowedCoins returns an sdk.Coins object from the store representing all currently borrowed coins
func (k Keeper) GetBorrowedCoins(rctx context.Context) (sdk.Coins, bool) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types2.BorrowedCoinsPrefix)
	bz := store.Get(types2.BorrowedCoinsPrefix)
	if len(bz) == 0 {
		return sdk.Coins{}, false
	}
	var borrowed types2.CoinsProto
	k.cdc.MustUnmarshal(bz, &borrowed)
	return borrowed.Coins, true
}

// SetSuppliedCoins sets the total amount of coins currently supplied in the store
func (k Keeper) SetSuppliedCoins(rctx context.Context, suppliedCoins sdk.Coins) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types2.SuppliedCoinsPrefix)
	if suppliedCoins.Empty() {
		store.Set(types2.SuppliedCoinsPrefix, []byte{})
	} else {
		bz := k.cdc.MustMarshal(&types2.CoinsProto{
			Coins: suppliedCoins,
		})
		store.Set(types2.SuppliedCoinsPrefix, bz)
	}
}

// GetSuppliedCoins returns an sdk.Coins object from the store representing all currently supplied coins
func (k Keeper) GetSuppliedCoins(rctx context.Context) (sdk.Coins, bool) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types2.SuppliedCoinsPrefix)
	bz := store.Get(types2.SuppliedCoinsPrefix)
	if len(bz) == 0 {
		return sdk.Coins{}, false
	}
	var supplied types2.CoinsProto
	k.cdc.MustUnmarshal(bz, &supplied)
	return supplied.Coins, true
}

// GetMoneyMarket returns a money market from the store for a denom
func (k Keeper) GetMoneyMarket(rctx context.Context, denom string) (types2.MoneyMarket, bool) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types2.MoneyMarketsPrefix)
	bz := store.Get([]byte(denom))
	if len(bz) == 0 {
		return types2.MoneyMarket{}, false
	}
	var moneyMarket types2.MoneyMarket
	k.cdc.MustUnmarshal(bz, &moneyMarket)
	return moneyMarket, true
}

// SetMoneyMarket sets a money market in the store for a denom
func (k Keeper) SetMoneyMarket(rctx context.Context, denom string, moneyMarket types2.MoneyMarket) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types2.MoneyMarketsPrefix)
	bz := k.cdc.MustMarshal(&moneyMarket)
	store.Set([]byte(denom), bz)
}

// DeleteMoneyMarket deletes a money market from the store
func (k Keeper) DeleteMoneyMarket(rctx context.Context, denom string) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types2.MoneyMarketsPrefix)
	store.Delete([]byte(denom))
}

// IterateMoneyMarkets iterates over all money markets objects in the store and performs a callback function
//
//	that returns both the money market and the key (denom) it's stored under
func (k Keeper) IterateMoneyMarkets(rctx context.Context, cb func(denom string, moneyMarket types2.MoneyMarket) (stop bool)) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types2.MoneyMarketsPrefix)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var moneyMarket types2.MoneyMarket
		k.cdc.MustUnmarshal(iterator.Value(), &moneyMarket)
		if cb(string(iterator.Key()), moneyMarket) {
			break
		}
	}
}

// GetAllMoneyMarkets returns all money markets from the store
func (k Keeper) GetAllMoneyMarkets(ctx context.Context) (moneyMarkets types2.MoneyMarkets) {
	k.IterateMoneyMarkets(ctx, func(denom string, moneyMarket types2.MoneyMarket) bool {
		moneyMarkets = append(moneyMarkets, moneyMarket)
		return false
	})
	return
}

// GetPreviousAccrualTime returns the last time an individual market accrued interest
func (k Keeper) GetPreviousAccrualTime(rctx context.Context, denom string) (time.Time, bool) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types2.PreviousAccrualTimePrefix)
	bz := store.Get([]byte(denom))
	if len(bz) == 0 {
		return time.Time{}, false
	}

	var previousAccrualTime time.Time
	if err := previousAccrualTime.UnmarshalBinary(bz); err != nil {
		panic(err)
	}
	return previousAccrualTime, true
}

// SetPreviousAccrualTime sets the most recent accrual time for a particular market
func (k Keeper) SetPreviousAccrualTime(rctx context.Context, denom string, previousAccrualTime time.Time) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types2.PreviousAccrualTimePrefix)
	bz, err := previousAccrualTime.MarshalBinary()
	if err != nil {
		panic(err)
	}
	store.Set([]byte(denom), bz)
}

// SetTotalReserves sets the total reserves for an individual market
func (k Keeper) SetTotalReserves(rctx context.Context, coins sdk.Coins) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types2.TotalReservesPrefix)
	if coins.Empty() {
		store.Set(types2.TotalReservesPrefix, []byte{})
		return
	}

	bz := k.cdc.MustMarshal(&types2.CoinsProto{
		Coins: coins,
	})
	store.Set(types2.TotalReservesPrefix, bz)
}

// GetTotalReserves returns the total reserves for an individual market
func (k Keeper) GetTotalReserves(rctx context.Context) (sdk.Coins, bool) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types2.TotalReservesPrefix)
	bz := store.Get(types2.TotalReservesPrefix)
	if len(bz) == 0 {
		return sdk.Coins{}, false
	}

	var totalReserves types2.CoinsProto
	k.cdc.MustUnmarshal(bz, &totalReserves)
	return totalReserves.Coins, true
}

// GetBorrowInterestFactor returns the current borrow interest factor for an individual market
func (k Keeper) GetBorrowInterestFactor(rctx context.Context, denom string) (sdkmath.LegacyDec, bool) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types2.BorrowInterestFactorPrefix)
	bz := store.Get([]byte(denom))
	if len(bz) == 0 {
		return sdkmath.LegacyZeroDec(), false
	}
	var borrowInterestFactor sdkmath.LegacyDec
	k.cdc.MustUnmarshal(bz, &borrowInterestFactor)
	return borrowInterestFactor.Dec, true
}

// SetBorrowInterestFactor sets the current borrow interest factor for an individual market
func (k Keeper) SetBorrowInterestFactor(rctx context.Context, denom string, borrowInterestFactor sdkmath.LegacyDec) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types2.BorrowInterestFactorPrefix)
	bz := k.cdc.MustMarshal(&sdkmath.LegacyDecProto{Dec: borrowInterestFactor})
	store.Set([]byte(denom), bz)
}

// IterateBorrowInterestFactors iterates over all borrow interest factors in the store and returns
// both the borrow interest factor and the key (denom) it's stored under
func (k Keeper) IterateBorrowInterestFactors(rctx context.Context, cb func(denom string, factor sdkmath.LegacyDec) (stop bool)) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types2.BorrowInterestFactorPrefix)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var factor sdkmath.LegacyDecProto
		k.cdc.MustUnmarshal(iterator.Value(), &factor)
		if cb(string(iterator.Key()), factor.Dec) {
			break
		}
	}
}

// GetSupplyInterestFactor returns the current supply interest factor for an individual market
func (k Keeper) GetSupplyInterestFactor(ctx context.Context, denom string) (sdkmath.LegacyDec, bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.SupplyInterestFactorPrefix)
	bz := store.Get([]byte(denom))
	if len(bz) == 0 {
		return sdkmath.LegacyZeroDec(), false
	}
	var supplyInterestFactor sdkmath.LegacyDecProto
	k.cdc.MustUnmarshal(bz, &supplyInterestFactor)
	return supplyInterestFactor.Dec, true
}

// SetSupplyInterestFactor sets the current supply interest factor for an individual market
func (k Keeper) SetSupplyInterestFactor(ctx context.Context, denom string, supplyInterestFactor sdkmath.LegacyDec) {
	store := prefix.NewStore(ctx.KVStore(k.key), types2.SupplyInterestFactorPrefix)
	bz := k.cdc.MustMarshal(&sdkmath.LegacyDecProto{Dec: supplyInterestFactor})
	store.Set([]byte(denom), bz)
}

// IterateSupplyInterestFactors iterates over all supply interest factors in the store and returns
// both the supply interest factor and the key (denom) it's stored under
func (k Keeper) IterateSupplyInterestFactors(rctx context.Context, cb func(denom string, factor sdkmath.LegacyDec) (stop bool)) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.key), types2.SupplyInterestFactorPrefix)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var factor sdkmath.LegacyDecProto

		k.cdc.MustUnmarshal(iterator.Value(), &factor)
		if cb(string(iterator.Key()), factor.Dec) {
			break
		}
	}
}
