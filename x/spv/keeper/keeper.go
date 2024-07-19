package keeper

import (
	"fmt"
	"time"

	"cosmossdk.io/math"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
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
		incentivekeeper types.IncentiveKeeper
		hooks           types.SPVHooks
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
	incentiveKeeper types.IncentiveKeeper,
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
		incentivekeeper: incentiveKeeper,
		hooks:           nil,
	}
}

func (k *Keeper) SetIncentiveKeeper(incentiveKeeper types.IncentiveKeeper) *Keeper {
	k.incentivekeeper = incentiveKeeper
	return k
}

// SetHooks adds hooks to the keeper.
func (k *Keeper) SetHooks(sh types.SPVHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set SPV hooks twice")
	}
	k.hooks = sh
	return k
}

func (k *Keeper) IsHookSet() bool {
	return k.hooks != nil
}

func (k Keeper) Logger(ctx context.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetPool sets the pool
func (k Keeper) SetPool(ctx context.Context, poolInfo types.PoolInfo) {
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
func (k Keeper) SetReserve(ctx context.Context, reserved sdk.Coin) {
	storeKey := fmt.Sprintf("%v%v", types.ProjectsKeyPrefix, "reserve")
	reserveStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(storeKey))
	bz := k.cdc.MustMarshal(&reserved)
	key := fmt.Sprintf("reserve-%v", reserved.Denom)
	reserveStore.Set(types.KeyPrefix(key), bz)
}

func (k Keeper) DelPool(ctx context.Context, index string) {
	poolsStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Pool))
	poolsStore.Delete(types.KeyPrefix(index))
}

// SetHistoryPool sets the pool
func (k Keeper) SetHistoryPool(ctx context.Context, poolInfo types.PoolInfo) {
	poolsStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HistoryPool))
	bz := k.cdc.MustMarshal(&poolInfo)
	poolsStore.Set(types.KeyPrefix(poolInfo.Index), bz)
}

// AddInvestorToPool add investors to the give pool
func (k Keeper) AddInvestorToPool(ctx context.Context, poolWithInvestors *types.PoolWithInvestors) {
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

func (k Keeper) GetInvestorToPool(ctx context.Context, poolIndex string) (currentInvestorPool types.PoolWithInvestors, found bool) {
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
func (k Keeper) IterateInvestorPools(ctx context.Context, cb func(poolWithInvestors types.PoolWithInvestors) (stop bool)) {
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
func (k Keeper) GetPools(ctx context.Context, index string) (poolInfo types.PoolInfo, ok bool) {
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

func (k Keeper) GetPoolBorrowed(ctx context.Context, poolIndex string) (borrowed math.Int, ok bool) {
	pool, ok := k.GetPools(ctx, poolIndex)
	if !ok {
		return borrowed, false
	}
	return pool.BorrowedAmount.Amount, true
}

func (k Keeper) GetDepositorTotalBorrowedAmount(ctx context.Context, depositor sdk.AccAddress, poolID string) (borrowed math.Int, found bool) {
	depositorInfo, found := k.GetDepositor(ctx, poolID, depositor)
	if !found {
		return borrowed, false
	}
	return depositorInfo.TotalPaidLiquidationAmount, true
}

func (k Keeper) IterSPVReserve(ctx context.Context, cb func(coin sdk.Coin) (stop bool)) {
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

// GetReserve gets the poolInfo with given pool index
func (k Keeper) GetReserve(ctx context.Context, denom string) (amount sdk.Coin, ok bool) {
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
func (k Keeper) IterateReserve(ctx context.Context, cb func(coin sdk.Coin) (stop bool)) {
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
func (k Keeper) IteratePool(ctx context.Context, cb func(poolInfo types.PoolInfo) (stop bool)) {
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
func (k Keeper) SetDepositorHistory(ctx context.Context, depositor types.DepositorInfo) {
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

// GetHistoryPools gets the poolInfo with given pool index
func (k Keeper) GetHistoryPools(ctx context.Context, index string) (poolInfo types.PoolInfo, ok bool) {
	gasBefore := ctx.GasMeter().GasConsumed()
	defer func() {
		gasAfter := ctx.GasMeter().GasConsumed()
		ctx.GasMeter().RefundGas(gasAfter-gasBefore, "SetPool")
	}()
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HistoryPool))
	bz := store.Get(types.KeyPrefix(index))
	if bz == nil {
		return poolInfo, false
	}
	k.cdc.MustUnmarshal(bz, &poolInfo)
	return poolInfo, true
}

// GetDepositorHistory sets the depositor to history store
func (k Keeper) GetDepositorHistory(ctx context.Context, timeStamp time.Time, poolIndex string, addr sdk.AccAddress) (types.DepositorInfo, bool) {
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

// IteratorAllDepositorHistory gets all the depositor to history store
func (k Keeper) IteratorAllDepositorHistory(ctx context.Context, poolIndex string, cb func(depositor types.DepositorInfo) (stop bool)) {
	gasBefore := ctx.GasMeter().GasConsumed()
	defer func() {
		gasAfter := ctx.GasMeter().GasConsumed()
		ctx.GasMeter().RefundGas(gasAfter-gasBefore, "SetPool")
	}()
	depositorPoolStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolDepositorHistory+poolIndex))

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

// SetDepositor sets the depositor
func (k Keeper) SetDepositor(ctx context.Context, depositor types.DepositorInfo) {
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
func (k Keeper) DelDepositor(ctx context.Context, depositor types.DepositorInfo) {
	gasBefore := ctx.GasMeter().GasConsumed()
	defer func() {
		gasAfter := ctx.GasMeter().GasConsumed()
		ctx.GasMeter().RefundGas(gasAfter-gasBefore, "SetPool")
	}()
	depositorPoolStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolDepositor+depositor.PoolIndex))
	depositorPoolStore.Delete(depositor.GetDepositorAddress().Bytes())
}

func (k Keeper) GetDepositor(ctx context.Context, poolIndex string, walletAddress sdk.AccAddress) (depositor types.DepositorInfo, found bool) {
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
func (k Keeper) IterateDepositors(ctx context.Context, poolIndex string, cb func(depositor types.DepositorInfo) (stop bool)) {
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
