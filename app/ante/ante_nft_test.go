package ante_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/tendermint/tendermint/libs/log"
	tmdb "github.com/tendermint/tm-db"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/app/ante"
)

func setupApp(t *testing.T) (sdk.Context, *spvkeeper.Keeper) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := paramtypes.NewSubspace(cdc,
		codec.NewLegacyAmino(),
		storeKey,
		memStoreKey,
		"SpvParams",
	)

	k := spvkeeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())
	return ctx, k
}

func mockAnteHandler(ctx sdk.Context, tx sdk.Tx, simulate bool) (newCtx sdk.Context, err error) {
	return ctx, nil
}

func TestTransferSPVNFT(t *testing.T) {
	encod := app.MakeEncodingConfig()
	txConfig := encod.TxConfig

	ctx, k := setupApp(t)

	testPrivKeys, testAddresses := app.GeneratePrivKeyAddressPairs(5)

	mockPool := types.PoolInfo{
		Index:      "demoIndex",
		PoolNFTIds: []string{"demonft"},
	}
	k.SetPool(ctx, mockPool)

	decorator := ante.NewSPVNFTDecorator(*k)

	tx, err := helpers.GenSignedMockTx(
		rand.New(rand.NewSource(time.Now().UnixNano())),
		txConfig,
		[]sdk.Msg{
			&nfttypes.MsgSend{
				ClassId:  "demonft",
				Id:       "whatever",
				Sender:   testAddresses[0].String(),
				Receiver: testAddresses[1].String(),
			},
		},
		sdk.NewCoins(),
		helpers.DefaultGenTxGas,
		"testing-chain-id",
		[]uint64{0},
		[]uint64{0},
		testPrivKeys[0],
	)
	require.NoError(t, err)
	ctx = ctx.WithIsCheckTx(true)
	_, err = decorator.AnteHandle(ctx, tx, false, mockAnteHandler)
	require.Error(t, err)
	require.Contains(t, err.Error(), "found disabled spv nft")

	tx2, err := helpers.GenSignedMockTx(
		rand.New(rand.NewSource(time.Now().UnixNano())),
		txConfig,
		[]sdk.Msg{
			&nfttypes.MsgSend{
				ClassId:  "demonft_not_spv",
				Id:       "whatever",
				Sender:   testAddresses[0].String(),
				Receiver: testAddresses[1].String(),
			},
		},
		sdk.NewCoins(),
		helpers.DefaultGenTxGas,
		"testing-chain-id",
		[]uint64{0},
		[]uint64{0},
		testPrivKeys[0],
	)
	require.NoError(t, err)
	_, err = decorator.AnteHandle(ctx, tx2, false, mockAnteHandler)
	require.NoError(t, err)

	tx3, err := helpers.GenSignedMockTx(
		rand.New(rand.NewSource(time.Now().UnixNano())),
		txConfig,
		[]sdk.Msg{
			&nfttypes.MsgSend{
				ClassId:  "demonft_not_spv",
				Id:       "whatever",
				Sender:   testAddresses[0].String(),
				Receiver: testAddresses[1].String(),
			},
			&nfttypes.MsgSend{
				ClassId:  "demonft",
				Id:       "whatever",
				Sender:   testAddresses[0].String(),
				Receiver: testAddresses[1].String(),
			},
		},
		sdk.NewCoins(),
		helpers.DefaultGenTxGas,
		"testing-chain-id",
		[]uint64{0},
		[]uint64{0},
		testPrivKeys[0],
	)
	require.NoError(t, err)
	_, err = decorator.AnteHandle(ctx, tx3, false, mockAnteHandler)
	require.Error(t, err)
	require.Contains(t, err.Error(), "found disabled spv nft")
}