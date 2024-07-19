package keeper

import (
	"testing"

	"cosmossdk.io/store"
	storetypes "cosmossdk.io/store/types"
	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/joltify-finance/joltify_lending/x/quota/keeper"
	"github.com/joltify-finance/joltify_lending/x/quota/types"
	"github.com/stretchr/testify/require"
)

// NewParams creates a new Params instance
func testParams() types.Params {
	// the coin list is the amount of USD for the given token, 100jolt means 100 USD value of jolt
	quota, err := sdk.ParseCoinsNormalized("100000ujolt,1000000usdt")
	if err != nil {
		panic(err)
	}

	quotaAcc, err := sdk.ParseCoinsNormalized("10000000ujolt,100000000usdt")
	if err != nil {
		panic(err)
	}

	targets := types.Target{
		ModuleName:    "ibc",
		CoinsSum:      quota,
		HistoryLength: 512,
	}

	targets2 := types.Target{
		ModuleName:    "bridge",
		CoinsSum:      quota,
		HistoryLength: 512,
	}

	targetsAcc := types.Target{
		ModuleName:    "ibc",
		CoinsSum:      quotaAcc,
		HistoryLength: 512,
	}

	targets2Acc := types.Target{
		ModuleName:    "bridge",
		CoinsSum:      quotaAcc,
		HistoryLength: 512,
	}

	return types.Params{Targets: []*types.Target{&targets, &targets2}, PerAccounttargets: []*types.Target{&targetsAcc, &targets2Acc}}
}

func QuotaKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"QuotaParams",
	)
	k := keeper.NewKeeper(
		cdc,
		storeKey,
		paramsSubspace,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, testParams())

	return k, ctx
}
