package keeper

import (
	dbm "github.com/cosmos/cosmos-db"
	"github.com/joltify-finance/joltify_lending/lib"
	"github.com/joltify-finance/joltify_lending/mocks"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	delaymsgtypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/delaymsg/types"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/feetiers/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/feetiers/types"
	statskeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/stats/keeper"
	vaultkeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/vault/keeper"
)

func createFeeTiersKeeper(
	stateStore storetypes.CommitMultiStore,
	statsKeeper *statskeeper.Keeper,
	vaultKeeper *vaultkeeper.Keeper,
	db *dbm.MemDB,
	cdc *codec.ProtoCodec,
) (*keeper.Keeper, storetypes.StoreKey) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)

	mockMsgSender := &mocks.IndexerMessageSender{}
	mockMsgSender.On("Enabled").Return(true)

	authorities := []string{
		delaymsgtypes.ModuleAddress.String(),
		lib.GovModuleAddress.String(),
	}
	k := keeper.NewKeeper(
		cdc,
		statsKeeper,
		storeKey,
		authorities,
	)
	k.SetVaultKeeper(vaultKeeper)

	return k, storeKey
}
