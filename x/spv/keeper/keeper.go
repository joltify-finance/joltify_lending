package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
		kycKeeper  types.KycKeeper
		bankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	kycKeeper types.KycKeeper,
	bankKeeper types.BankKeeper,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		kycKeeper:  kycKeeper,
		bankKeeper: bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetPool sets the pool
func (k Keeper) SetPool(ctx sdk.Context, poolInfo types.PoolInfo) {
	poolsStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectsKeyPrefix))
	bz := k.cdc.MustMarshal(&poolInfo)
	poolsStore.Set(types.KeyPrefix(poolInfo.Index), bz)
}

// AddInvestorToPool add investors to the give poolindex
func (k Keeper) AddInvestorToPool(ctx sdk.Context, poolWithInvestors *types.PoolWithInvestors) {
	poolsStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolInvestor))
	key := types.KeyPrefix(poolWithInvestors.PoolIndex)
	bz := k.cdc.MustMarshal(poolWithInvestors)
	poolsStore.Set(key, bz)
}

func (k Keeper) GetInvestorToPool(ctx sdk.Context, poolIndex string) (currentInvestorPool types.PoolWithInvestors, found bool) {
	poolsStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolInvestor))
	key := types.KeyPrefix(poolIndex)
	bz := poolsStore.Get(key)
	if bz == nil {
		return currentInvestorPool, false
	}
	k.cdc.MustUnmarshal(bz, &currentInvestorPool)
	return currentInvestorPool, true
}

// GetPools gets the poolInfo with given pool index
func (k Keeper) GetPools(ctx sdk.Context, index string) (poolInfo types.PoolInfo, ok bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectsKeyPrefix))
	bz := store.Get(types.KeyPrefix(index))
	if bz == nil {
		return poolInfo, false
	}
	k.cdc.MustUnmarshal(bz, &poolInfo)
	return poolInfo, true
}

// SetDepositor sets the depositor
func (k Keeper) SetDepositor(ctx sdk.Context, depositor types.DepositorInfo) {
	depositorPoolStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolDepositor+depositor.PoolIndex))

	bz := k.cdc.MustMarshal(&depositor)
	depositorPoolStore.Set(depositor.GetDepositorAddress().Bytes(), bz)
}

func (k Keeper) GetDepositor(ctx sdk.Context, poolIndex string, walletAddress sdk.AccAddress) (depositor types.DepositorInfo, found bool) {
	depositorPoolStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(poolIndex))

	bz := depositorPoolStore.Get(walletAddress.Bytes())
	if bz == nil {
		return depositor, found
	}

	k.cdc.MustUnmarshal(bz, &depositor)
	return depositor, true
}
