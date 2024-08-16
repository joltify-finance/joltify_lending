package keeper

import (
	"testing"

	storetypes "cosmossdk.io/store/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/lib"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/blocktime/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/blocktime/types"
	delaymsgtypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/delaymsg/types"
)

func createBlockTimeKeeper(
	stateStore storetypes.CommitMultiStore,
	db *dbm.MemDB,
	cdc *codec.ProtoCodec,
) (*keeper.Keeper, storetypes.StoreKey) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)

	authorities := []string{
		delaymsgtypes.ModuleAddress.String(),
		lib.GovModuleAddress.String(),
	}
	k := keeper.NewKeeper(
		cdc,
		storeKey,
		authorities,
	)

	return k, storeKey
}

func BlcokTimeKeepers(
	t testing.TB,
) (
	ctx sdk.Context,
	keeper *keeper.Keeper,
	storeKey storetypes.StoreKey,
) {
	ctx = initKeepers(t, func(
		db *dbm.MemDB,
		registry codectypes.InterfaceRegistry,
		cdc *codec.ProtoCodec,
		stateStore storetypes.CommitMultiStore,
		transientStoreKey storetypes.StoreKey,
	) []GenesisInitializer {
		// Define necessary keepers here for unit tests
		keeper, storeKey = createBlockTimeKeeper(stateStore, db, cdc)

		return []GenesisInitializer{keeper}
	})
	return ctx, keeper, storeKey
}
