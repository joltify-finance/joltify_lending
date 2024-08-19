package keeper

import (
	"testing"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/joltify-finance/joltify_lending/lib"
	"github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/ante"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/joltify-finance/joltify_lending/mocks"
	bridgekeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/bridge/keeper"
	bridgetypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/bridge/types"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/delaymsg/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/delaymsg/types"
)

func DelayMsgKeepers(
	t testing.TB,
) (
	ctx sdk.Context,
	delayMsgKeeper *keeper.Keeper,
	storeKey storetypes.StoreKey,
	bridgeKeeper *bridgekeeper.Keeper,
	bankKeeper bankkeeper.Keeper,
	authorities []string,
) {
	ctx = initKeepers(t, func(
		db *dbm.MemDB,
		_ codectypes.InterfaceRegistry,
		_ *codec.ProtoCodec,
		stateStore storetypes.CommitMultiStore,
		transientStoreKey storetypes.StoreKey,
	) []GenesisInitializer {
		encCfg := ante.MakeTestEncodingConfig()
		cdc := encCfg.Codec.(*codec.ProtoCodec)
		registry := encCfg.InterfaceRegistry

		router := baseapp.NewMsgServiceRouter()
		router.SetInterfaceRegistry(registry)

		// Register bridge messages for encoding / decoding.
		bridgetypes.RegisterInterfaces(registry)

		accountKeeper, _ := createAccountKeeper(stateStore, db, cdc, registry)
		bankKeeper, _ = createBankKeeper(stateStore, db, cdc, accountKeeper)
		bridgeKeeper, _, _, _, _ = createBridgeKeeper(stateStore, db, cdc, transientStoreKey, bankKeeper)

		// Register bridge keeper msg server for msg routing.
		bridgetypes.RegisterMsgServer(router, bridgekeeper.NewMsgServerImpl(bridgeKeeper))

		authorities = []string{
			bridgetypes.ModuleAddress.String(),
			lib.GovModuleAddress.String(),
		}
		delayMsgKeeper, storeKey = createDelayMsgKeeper(
			stateStore,
			db,
			cdc,
			router,
			authorities,
		)

		return []GenesisInitializer{
			delayMsgKeeper,
		}
	})
	return ctx, delayMsgKeeper, storeKey, bridgeKeeper, bankKeeper, authorities
}

func DelayMsgKeeperWithMockBridgeKeeper(
	t testing.TB,
) (
	ctx sdk.Context,
	delayMsgKeeper *keeper.Keeper,
	storeKey storetypes.StoreKey,
	bridgeKeeper *mocks.BridgeKeeper,
	authorities []string,
) {
	ctx = initKeepers(t, func(
		db *dbm.MemDB,
		_ codectypes.InterfaceRegistry,
		_ *codec.ProtoCodec,
		stateStore storetypes.CommitMultiStore,
		transientStoreKey storetypes.StoreKey,
	) []GenesisInitializer {
		encCfg := ante.MakeTestEncodingConfig()
		cdc := encCfg.Codec.(*codec.ProtoCodec)
		registry := encCfg.InterfaceRegistry

		router := baseapp.NewMsgServiceRouter()
		router.SetInterfaceRegistry(registry)

		// Register bridge messages for encoding / decoding.
		bridgetypes.RegisterInterfaces(registry)

		bridgeKeeper = &mocks.BridgeKeeper{}

		// Register bridge keeper msg server for msg routing.
		bridgetypes.RegisterMsgServer(router, bridgekeeper.NewMsgServerImpl(bridgeKeeper))

		authorities = []string{
			bridgetypes.ModuleAddress.String(),
			lib.GovModuleAddress.String(),
		}
		delayMsgKeeper, storeKey = createDelayMsgKeeper(
			stateStore,
			db,
			cdc,
			router,
			authorities,
		)

		return []GenesisInitializer{delayMsgKeeper}
	})
	return ctx, delayMsgKeeper, storeKey, bridgeKeeper, authorities
}

func createDelayMsgKeeper(
	stateStore storetypes.CommitMultiStore,
	db *dbm.MemDB,
	cdc *codec.ProtoCodec,
	router *baseapp.MsgServiceRouter,
	authorities []string,
) (*keeper.Keeper, storetypes.StoreKey) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		router,
		authorities,
	)
	return k, storeKey
}