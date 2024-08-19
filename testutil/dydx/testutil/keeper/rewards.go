package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/proto"

	storetypes "cosmossdk.io/store/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	indexerevents "github.com/joltify-finance/joltify_lending/dydx_helper/indexer/events"
	"github.com/joltify-finance/joltify_lending/dydx_helper/indexer/indexer_manager"
	"github.com/joltify-finance/joltify_lending/lib"
	"github.com/joltify-finance/joltify_lending/mocks"
	assetskeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/assets/keeper"
	delaymsgtypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/delaymsg/types"
	feetierskeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/feetiers/keeper"
	priceskeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/keeper"
	rewardskeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/rewards/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/rewards/types"
)

func RewardsKeepers(
	t testing.TB,
) (
	ctx sdk.Context,
	rewardsKeeper *rewardskeeper.Keeper,
	feetiersKeeper *feetierskeeper.Keeper,
	bankKeeper bankkeeper.Keeper,
	assetsKeeper *assetskeeper.Keeper,
	pricesKeeper *priceskeeper.Keeper,
	indexerEventManager indexer_manager.IndexerEventManager,
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
		pricesKeeper, _, _, _ = createPricesKeeper(stateStore, db, cdc, transientStoreKey)
		// Mock time provider response for market creation.
		epochsKeeper, _ := createEpochsKeeper(stateStore, db, cdc)
		assetsKeeper, _ = createAssetsKeeper(
			stateStore,
			db,
			cdc,
			pricesKeeper,
			transientStoreKey,
			true,
		)
		statsKeeper, _ := createStatsKeeper(
			stateStore,
			epochsKeeper,
			db,
			cdc,
		)
		vaultKeeper, _ := createVaultKeeper(
			stateStore,
			db,
			cdc,
			transientStoreKey,
		)
		feetiersKeeper, _ = createFeeTiersKeeper(
			stateStore,
			statsKeeper,
			vaultKeeper,
			db,
			cdc,
		)
		rewardsKeeper, storeKey = createRewardsKeeper(
			stateStore,
			assetsKeeper,
			bankKeeper,
			feetiersKeeper,
			pricesKeeper,
			indexerEventManager,
			db,
			cdc,
		)

		return []GenesisInitializer{
			pricesKeeper,
			assetsKeeper,
			feetiersKeeper,
			statsKeeper,
		}
	})
	return ctx, rewardsKeeper, feetiersKeeper, bankKeeper, assetsKeeper, pricesKeeper, indexerEventManager, storeKey
}

func createRewardsKeeper(
	stateStore storetypes.CommitMultiStore,
	assetsKeeper *assetskeeper.Keeper,
	bankKeeper bankkeeper.Keeper,
	feeTiersKeeper *feetierskeeper.Keeper,
	pricesKeeper *priceskeeper.Keeper,
	indexerEventManager indexer_manager.IndexerEventManager,
	db *dbm.MemDB,
	cdc *codec.ProtoCodec,
) (*rewardskeeper.Keeper, storetypes.StoreKey) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	transientStoreKey := storetypes.NewTransientStoreKey(types.TransientStoreKey)

	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(transientStoreKey, storetypes.StoreTypeTransient, db)

	mockMsgSender := &mocks.IndexerMessageSender{}
	mockMsgSender.On("Enabled").Return(true)

	authorities := []string{
		delaymsgtypes.ModuleAddress.String(),
		lib.GovModuleAddress.String(),
	}
	k := rewardskeeper.NewKeeper(
		cdc,
		storeKey,
		transientStoreKey,
		assetsKeeper,
		bankKeeper,
		feeTiersKeeper,
		pricesKeeper,
		indexerEventManager,
		authorities,
	)

	return k, storeKey
}

func GetTradingRewardEventsFromIndexerTendermintBlock(
	block indexer_manager.IndexerTendermintBlock,
) []*indexerevents.TradingRewardsEventV1 {
	var rewardEvents []*indexerevents.TradingRewardsEventV1
	for _, event := range block.Events {
		if event.Subtype != indexerevents.SubtypeTradingReward {
			continue
		}
		var rewardEvent indexerevents.TradingRewardsEventV1
		err := proto.Unmarshal(event.DataBytes, &rewardEvent)
		if err != nil {
			panic(err)
		}
		rewardEvents = append(rewardEvents, &rewardEvent)
	}
	return rewardEvents
}