package keeper

import (
	"testing"

	"cosmossdk.io/store/metrics"

	dbm "github.com/cosmos/cosmos-db"

	"github.com/joltify-finance/joltify_lending/x/burnauction/keeper"
	"github.com/joltify-finance/joltify_lending/x/burnauction/types"

	"cosmossdk.io/store"
	storetypes "cosmossdk.io/store/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"
)

func BurnauctionKeeper(t testing.TB) (*keeper.Keeper, types.BankKeeper, types.AccountKeeper, sdk.Context) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"BurnauctionParams",
	)

	bankKeeper := mockbankKeeper{make(map[string]sdk.Coins)}

	auctionKeeper := MockAuctionKeeper{
		AuctionAmount: make([]sdk.Coin, 1),
		SellerBid:     make([]string, 2),
		mockbank:      &bankKeeper,
	}

	accKeeper := mockAccKeeper{}
	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		accKeeper,
		bankKeeper,
		auctionKeeper,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, bankKeeper, accKeeper, ctx
}
