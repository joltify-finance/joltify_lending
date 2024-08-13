package keeper

import (
	storetypes "cosmossdk.io/store/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/joltify-finance/joltify_lending/dydx_helper/lib"
	"github.com/joltify-finance/joltify_lending/dydx_helper/x/blocktime/keeper"
	"github.com/joltify-finance/joltify_lending/dydx_helper/x/blocktime/types"
	delaymsgtypes "github.com/joltify-finance/joltify_lending/dydx_helper/x/delaymsg/types"
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
