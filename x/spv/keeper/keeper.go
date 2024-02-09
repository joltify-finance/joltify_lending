package keeper

import (
	"fmt"
	"time"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

type (
	Keeper struct {
		cdc             codec.BinaryCodec
		storeKey        storetypes.StoreKey
		memKey          storetypes.StoreKey
		paramstore      paramtypes.Subspace
		kycKeeper       types.KycKeeper
		bankKeeper      types.BankKeeper
		accKeeper       types.AccountKeeper
		NftKeeper       types.NFTKeeper
		priceFeedKeeper types.PriceFeedKeeper
		auctionKeeper   types.AuctionKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	kycKeeper types.KycKeeper,
	bankKeeper types.BankKeeper,
	accKeeper types.AccountKeeper,
	nftKeeper types.NFTKeeper,
	pricefeedkeeper types.PriceFeedKeeper,
	auctionKeeper types.AuctionKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:             cdc,
		storeKey:        storeKey,
		memKey:          memKey,
		paramstore:      ps,
		kycKeeper:       kycKeeper,
		bankKeeper:      bankKeeper,
		accKeeper:       accKeeper,
		NftKeeper:       nftKeeper,
		priceFeedKeeper: pricefeedkeeper,
		auctionKeeper:   auctionKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetPool sets the pool
func (k Keeper) SetPool(ctx sdk.Context, poolInfo types.PoolInfo) {
	gasBefore := ctx.GasMeter().GasConsumed()
	defer func() {
		gasAfter := ctx.GasMeter().GasConsumed()
		ctx.GasMeter().RefundGas(gasAfter-gasBefore, "SetPool")
	}()

	poolsStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Pool))
	bz := k.cdc.MustMarshal(&poolInfo)
	poolsStore.Set(types.KeyPrefix(poolInfo.Index), bz)
}

// SetReserve sets the pool
func (k Keeper) SetReserve(ctx sdk.Context, reserved sdk.Coin) {
	storeKey := fmt.Sprintf("%v%v", types.ProjectsKeyPrefix, "reserve")
	reserveStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(storeKey))
	bz := k.cdc.MustMarshal(&reserved)
	key := fmt.Sprintf("reserve-%v", reserved.Denom)
	reserveStore.Set(types.KeyPrefix(key), bz)
}

func (k Keeper) DelPool(ctx sdk.Context, index string) {
	poolsStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Pool))
	poolsStore.Delete(types.KeyPrefix(index))
}

// SetHistoryPool sets the pool
func (k Keeper) SetHistoryPool(ctx sdk.Context, poolInfo types.PoolInfo) {
	poolsStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HistoryProjectsKeyPrefix))
	bz := k.cdc.MustMarshal(&poolInfo)
	poolsStore.Set(types.KeyPrefix(poolInfo.Index), bz)
}

// AddInvestorToPool add investors to the give pool
func (k Keeper) AddInvestorToPool(ctx sdk.Context, poolWithInvestors *types.PoolWithInvestors) {
	gasBefore := ctx.GasMeter().GasConsumed()
	defer func() {
		gasAfter := ctx.GasMeter().GasConsumed()
		ctx.GasMeter().RefundGas(gasAfter-gasBefore, "SetPool")
	}()
	poolsStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolInvestor))
	key := types.KeyPrefix(poolWithInvestors.PoolIndex)
	bz := k.cdc.MustMarshal(poolWithInvestors)
	poolsStore.Set(key, bz)
}

func (k Keeper) GetInvestorToPool(ctx sdk.Context, poolIndex string) (currentInvestorPool types.PoolWithInvestors, found bool) {
	gasBefore := ctx.GasMeter().GasConsumed()
	defer func() {
		gasAfter := ctx.GasMeter().GasConsumed()
		ctx.GasMeter().RefundGas(gasAfter-gasBefore, "SetPool")
	}()
	poolsStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolInvestor))
	key := types.KeyPrefix(poolIndex)
	bz := poolsStore.Get(key)
	if bz == nil {
		return currentInvestorPool, false
	}
	k.cdc.MustUnmarshal(bz, &currentInvestorPool)
	return currentInvestorPool, true
}

// IterateInvestorPools iterates over all pools objects in the store and performs a callback function
func (k Keeper) IterateInvestorPools(ctx sdk.Context, cb func(poolWithInvestors types.PoolWithInvestors) (stop bool)) {
	gasBefore := ctx.GasMeter().GasConsumed()
	defer func() {
		gasAfter := ctx.GasMeter().GasConsumed()
		ctx.GasMeter().RefundGas(gasAfter-gasBefore, "SetPool")
	}()
	poolsStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolInvestor))
	iterator := sdk.KVStorePrefixIterator(poolsStore, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var depositor types.PoolWithInvestors
		k.cdc.MustUnmarshal(iterator.Value(), &depositor)
		if cb(depositor) {
			break
		}
	}
}

// GetPools gets the poolInfo with given pool index
func (k Keeper) GetPools(ctx sdk.Context, index string) (poolInfo types.PoolInfo, ok bool) {
	gasBefore := ctx.GasMeter().GasConsumed()
	defer func() {
		gasAfter := ctx.GasMeter().GasConsumed()
		ctx.GasMeter().RefundGas(gasAfter-gasBefore, "SetPool")
	}()
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Pool))
	bz := store.Get(types.KeyPrefix(index))
	if bz == nil {
		return poolInfo, false
	}
	k.cdc.MustUnmarshal(bz, &poolInfo)
	return poolInfo, true
}

// GetReserve gets the poolInfo with given pool index
func (k Keeper) GetReserve(ctx sdk.Context, denom string) (amount sdk.Coin, ok bool) {
	storeKey := fmt.Sprintf("%v%v", types.ProjectsKeyPrefix, "reserve")
	reserveStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(storeKey))
	key := fmt.Sprintf("reserve-%v", denom)
	bz := reserveStore.Get(types.KeyPrefix(key))
	if bz == nil {
		return amount, false
	}
	k.cdc.MustUnmarshal(bz, &amount)
	return amount, true
}

// IterateReserve get the spv reserve token
func (k Keeper) IterateReserve(ctx sdk.Context, cb func(coin sdk.Coin) (stop bool)) {
	gasBefore := ctx.GasMeter().GasConsumed()
	defer func() {
		gasAfter := ctx.GasMeter().GasConsumed()
		ctx.GasMeter().RefundGas(gasAfter-gasBefore, "SetPool")
	}()
	storeKey := fmt.Sprintf("%v%v", types.ProjectsKeyPrefix, "reserve")
	reserveStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(storeKey))
	iterator := sdk.KVStorePrefixIterator(reserveStore, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var coin sdk.Coin
		k.cdc.MustUnmarshal(iterator.Value(), &coin)
		if cb(coin) {
			break
		}
	}
}

// IteratePool iterates over all deposit objects in the store and performs a callback function
func (k Keeper) IteratePool(ctx sdk.Context, cb func(poolInfo types.PoolInfo) (stop bool)) {
	gasBefore := ctx.GasMeter().GasConsumed()
	defer func() {
		gasAfter := ctx.GasMeter().GasConsumed()
		ctx.GasMeter().RefundGas(gasAfter-gasBefore, "SetPool")
	}()
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Pool))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var poolInfo types.PoolInfo
		k.cdc.MustUnmarshal(iterator.Value(), &poolInfo)
		if cb(poolInfo) {
			break
		}
	}
}

// SetDepositorHistory sets the depositor to history store
func (k Keeper) SetDepositorHistory(ctx sdk.Context, depositor types.DepositorInfo) {
	gasBefore := ctx.GasMeter().GasConsumed()
	defer func() {
		gasAfter := ctx.GasMeter().GasConsumed()
		ctx.GasMeter().RefundGas(gasAfter-gasBefore, "SetPool")
	}()
	depositorPoolStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolDepositorHistory+depositor.PoolIndex))
	bz := k.cdc.MustMarshal(&depositor)
	timeBytes, err := ctx.BlockTime().MarshalBinary()
	if err != nil {
		panic(err)
	}
	key := append(depositor.DepositorAddress.Bytes(), timeBytes...)
	depositorPoolStore.Set(key, bz)
}

// GetDepositorHistory sets the depositor to history store
func (k Keeper) GetDepositorHistory(ctx sdk.Context, timeStamp time.Time, poolIndex string, addr sdk.AccAddress) (types.DepositorInfo, bool) {
	gasBefore := ctx.GasMeter().GasConsumed()
	defer func() {
		gasAfter := ctx.GasMeter().GasConsumed()
		ctx.GasMeter().RefundGas(gasAfter-gasBefore, "SetPool")
	}()
	depositorPoolStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolDepositorHistory+poolIndex))
	var depositor types.DepositorInfo
	timeBytes, err := timeStamp.MarshalBinary()
	if err != nil {
		panic(err)
	}
	key := append(addr.Bytes(), timeBytes...)
	bz := depositorPoolStore.Get(key)
	if bz == nil {
		return depositor, false
	}
	k.cdc.MustUnmarshal(bz, &depositor)
	return depositor, true
}

// SetDepositor sets the depositor
func (k Keeper) SetDepositor(ctx sdk.Context, depositor types.DepositorInfo) {
	gasBefore := ctx.GasMeter().GasConsumed()
	defer func() {
		gasAfter := ctx.GasMeter().GasConsumed()
		ctx.GasMeter().RefundGas(gasAfter-gasBefore, "SetPool")
	}()
	depositorPoolStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolDepositor+depositor.PoolIndex))
	bz := k.cdc.MustMarshal(&depositor)

	depositorPoolStore.Set(depositor.GetDepositorAddress().Bytes(), bz)
}

// DelDepositor sets the depositor
func (k Keeper) DelDepositor(ctx sdk.Context, depositor types.DepositorInfo) {
	gasBefore := ctx.GasMeter().GasConsumed()
	defer func() {
		gasAfter := ctx.GasMeter().GasConsumed()
		ctx.GasMeter().RefundGas(gasAfter-gasBefore, "SetPool")
	}()
	depositorPoolStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolDepositor+depositor.PoolIndex))
	depositorPoolStore.Delete(depositor.GetDepositorAddress().Bytes())
}

func (k Keeper) GetDepositor(ctx sdk.Context, poolIndex string, walletAddress sdk.AccAddress) (depositor types.DepositorInfo, found bool) {
	gasBefore := ctx.GasMeter().GasConsumed()
	defer func() {
		gasAfter := ctx.GasMeter().GasConsumed()
		ctx.GasMeter().RefundGas(gasAfter-gasBefore, "SetPool")
	}()
	depositorPoolStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolDepositor+poolIndex))

	bz := depositorPoolStore.Get(walletAddress.Bytes())
	if bz == nil {
		return depositor, found
	}

	k.cdc.MustUnmarshal(bz, &depositor)
	return depositor, true
}

// IterateDepositors iterates over all deposit objects in the store and performs a callback function
func (k Keeper) IterateDepositors(ctx sdk.Context, poolIndex string, cb func(depositor types.DepositorInfo) (stop bool)) {
	gasBefore := ctx.GasMeter().GasConsumed()
	defer func() {
		gasAfter := ctx.GasMeter().GasConsumed()
		ctx.GasMeter().RefundGas(gasAfter-gasBefore, "SetPool")
	}()
	depositorPoolStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolDepositor+poolIndex))
	iterator := sdk.KVStorePrefixIterator(depositorPoolStore, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var depositor types.DepositorInfo
		k.cdc.MustUnmarshal(iterator.Value(), &depositor)
		if cb(depositor) {
			break
		}
	}
}
