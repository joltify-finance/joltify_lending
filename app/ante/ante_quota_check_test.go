package ante_test

import (
	"encoding/hex"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/cometbft/cometbft/libs/log"

	tmdb "github.com/cometbft/cometbft-db"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	types2 "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/app/ante"
	vaultkeeper "github.com/joltify-finance/joltify_lending/x/vault/keeper"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockStakingKeeper struct{}

func (m mockStakingKeeper) IterateLastValidators(_ sdk.Context, _ func(index int64, validator types2.ValidatorI) (stop bool)) {
	// TODO implement me
	panic("implement me")
}

func (m mockStakingKeeper) GetParams(_ sdk.Context) types2.Params {
	// TODO implement me
	panic("implement me")
}

func (m mockStakingKeeper) LastValidatorsIterator(_ sdk.Context) (iterator sdk.Iterator) {
	// TODO implement me
	panic("implement me")
}

func (m mockStakingKeeper) GetValidator(_ sdk.Context, _ sdk.ValAddress) (validator types2.Validator, found bool) {
	// TODO implement me
	panic("implement me")
}

func (m mockStakingKeeper) GetHistoricalInfo(_ sdk.Context, _ int64) (types2.HistoricalInfo, bool) {
	// TODO implement me
	panic("implement me")
}

func (m mockStakingKeeper) GetBondedValidatorsByPower(_ sdk.Context) (validators []types2.Validator) {
	v := types2.Validator{}
	return []types2.Validator{v, v}
}

func setupVaultApp(t *testing.T) (sdk.Context, *vaultkeeper.Keeper) {
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
		"VaultParams",
	)

	stakingKeeper := mockStakingKeeper{}
	k := vaultkeeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		stakingKeeper,
		nil,
		paramsSubspace,
		nil,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())
	return ctx, k
}

func generateNValidators(t *testing.T, n int) (types2.Validators, []sdk.AccAddress) {
	testValidators := make(types2.Validators, n)
	creators := make([]sdk.AccAddress, n)
	for i := 0; i < n; i++ {
		sk := ed25519.GenPrivKey()
		desc := types2.NewDescription("tester", "testId", "www.test.com", "aaa", "aaa")

		skCreator := secp256k1.GenPrivKey()

		skCreator.PubKey().Address()
		creator := skCreator.PubKey().Address().Bytes()

		valAddr, err := sdk.ValAddressFromHex(hex.EncodeToString(creator))
		require.NoError(t, err)
		testValidator, err := types2.NewValidator(valAddr, sk.PubKey(), desc)
		require.NoError(t, err)
		testValidators[i] = testValidator
		creators[i] = creator
	}
	return testValidators, creators
}

func TestQuotaCheck(t *testing.T) {
	encod := app.MakeEncodingConfig()
	txConfig := encod.TxConfig
	ctx, vaultKeeper := setupVaultApp(t)

	vd := ante.NewVaultQuotaDecorate(*vaultKeeper)

	params := vaultKeeper.GetParams(ctx)

	testValidators, creators := generateNValidators(t, 4)
	testPrivKeys, testAddresses := app.GeneratePrivKeyAddressPairs(5)

	p1 := types.PoolProposal{PoolAddr: creators[0], Nodes: []sdk.AccAddress{creators[0], creators[1], creators[2]}}
	p2 := types.PoolProposal{PoolAddr: creators[1], Nodes: []sdk.AccAddress{creators[0], creators[1], creators[2]}}

	position1 := ctx.BlockHeight() - params.BlockChurnInterval + 1
	position2 := ctx.BlockHeight() - params.BlockChurnInterval*2 + 1

	createPool := types.CreatePool{
		BlockHeight: strconv.FormatInt(position1, 10),
		Validators:  testValidators,
		Proposal:    []*types.PoolProposal{&p1, &p1, &p1},
	}
	vaultKeeper.SetCreatePool(ctx, createPool)

	createPool2 := types.CreatePool{
		BlockHeight: strconv.FormatInt(position2, 10),
		Validators:  testValidators,
		Proposal:    []*types.PoolProposal{&p2, &p2, &p2},
	}
	vaultKeeper.SetCreatePool(ctx, createPool)
	q := types.CoinsQuota{
		History:  []*types.HistoricalAmount{},
		CoinsSum: sdk.NewCoins(),
	}

	vaultKeeper.GenSetLastTwoPool(ctx, []*types.CreatePool{&createPool, &createPool2})
	vaultKeeper.SetQuotaData(ctx, q)

	decorator := ante.NewVaultQuotaDecorate(*vaultKeeper)

	coins := params.GetTargetQuota()
	ret := vd.QuotaCheck(ctx, coins)
	assert.True(t, ret)

	q.CoinsSum = coins
	vaultKeeper.SetQuotaData(ctx, q)

	ret = vd.QuotaCheck(ctx, coins)
	assert.False(t, ret)

	// tEth := sdk.NewCoins(sdk.NewCoin("aeth", sdk.NewInt(1)))
	tx, err := simtestutil.GenSignedMockTx(
		rand.New(rand.NewSource(time.Now().UnixNano())),
		txConfig,
		[]sdk.Msg{
			&banktypes.MsgSend{
				FromAddress: testAddresses[0].String(),
				ToAddress:   creators[0].String(),
				Amount:      coins,
			},
		},
		sdk.NewCoins(),
		simtestutil.DefaultGenTxGas,
		"testing-chain-id",
		[]uint64{0},
		[]uint64{0},
		testPrivKeys[0],
	)
	require.NoError(t, err)
	ctx = ctx.WithIsCheckTx(true)
	_, err = decorator.AnteHandle(ctx, tx, false, mockAnteHandler)
	require.Error(t, err)

	t1 := sdk.NewCoins(sdk.NewCoin("abnb", sdk.NewInt(100)))
	q.CoinsSum = q.CoinsSum.Sub(t1...)
	vaultKeeper.SetQuotaData(ctx, q)
	ret = vd.QuotaCheck(ctx, t1)
	assert.True(t, ret)

	tx, err = simtestutil.GenSignedMockTx(
		rand.New(rand.NewSource(time.Now().UnixNano())),
		txConfig,
		[]sdk.Msg{
			&banktypes.MsgSend{
				FromAddress: testAddresses[0].String(),
				ToAddress:   creators[0].String(),
				Amount:      t1,
			},
		},
		sdk.NewCoins(),
		simtestutil.DefaultGenTxGas,
		"testing-chain-id",
		[]uint64{0},
		[]uint64{0},
		testPrivKeys[0],
	)
	require.NoError(t, err)
	ctx = ctx.WithIsCheckTx(true)
	_, err = decorator.AnteHandle(ctx, tx, false, mockAnteHandler)
	require.NoError(t, err)

	tEth := sdk.NewCoins(sdk.NewCoin("aeth", sdk.NewInt(1)))

	tx, err = simtestutil.GenSignedMockTx(
		rand.New(rand.NewSource(time.Now().UnixNano())),
		txConfig,
		[]sdk.Msg{
			&banktypes.MsgSend{
				FromAddress: testAddresses[0].String(),
				ToAddress:   creators[0].String(),
				Amount:      tEth,
			},
		},
		sdk.NewCoins(),
		simtestutil.DefaultGenTxGas,
		"testing-chain-id",
		[]uint64{0},
		[]uint64{0},
		testPrivKeys[0],
	)
	require.NoError(t, err)
	ctx = ctx.WithIsCheckTx(true)
	_, err = decorator.AnteHandle(ctx, tx, false, mockAnteHandler)
	require.ErrorContainsf(t, err, "has reached the quota target", "quota check")

	//
	tEthandBnb := sdk.NewCoins(sdk.NewCoin("aeth", sdk.NewInt(1)), t1[0])
	ret = vd.QuotaCheck(ctx, tEthandBnb)
	assert.False(t, ret)

	tx, err = simtestutil.GenSignedMockTx(
		rand.New(rand.NewSource(time.Now().UnixNano())),
		txConfig,
		[]sdk.Msg{
			&banktypes.MsgSend{
				FromAddress: testAddresses[0].String(),
				ToAddress:   creators[0].String(),
				Amount:      tEthandBnb,
			},
		},
		sdk.NewCoins(),
		simtestutil.DefaultGenTxGas,
		"testing-chain-id",
		[]uint64{0},
		[]uint64{0},
		testPrivKeys[0],
	)
	require.NoError(t, err)
	ctx = ctx.WithIsCheckTx(true)
	_, err = decorator.AnteHandle(ctx, tx, false, mockAnteHandler)
	require.ErrorContainsf(t, err, "has reached the quota target", "quota check")

	t2 := t1.Add(sdk.NewCoin("abnb", sdk.NewInt(1)))
	tx, err = simtestutil.GenSignedMockTx(
		rand.New(rand.NewSource(time.Now().UnixNano())),
		txConfig,
		[]sdk.Msg{
			&banktypes.MsgSend{
				FromAddress: testAddresses[0].String(),
				ToAddress:   creators[0].String(),
				Amount:      t2,
			},
		},
		sdk.NewCoins(),
		simtestutil.DefaultGenTxGas,
		"testing-chain-id",
		[]uint64{0},
		[]uint64{0},
		testPrivKeys[0],
	)
	require.NoError(t, err)
	ctx = ctx.WithIsCheckTx(true)
	_, err = decorator.AnteHandle(ctx, tx, false, mockAnteHandler)
	require.ErrorContainsf(t, err, "has reached the quota target", "quota check")

	// not to the pool, should pass
	tx, err = simtestutil.GenSignedMockTx(
		rand.New(rand.NewSource(time.Now().UnixNano())),
		txConfig,
		[]sdk.Msg{
			&banktypes.MsgSend{
				FromAddress: testAddresses[0].String(),
				ToAddress:   creators[3].String(),
				Amount:      t2,
			},
		},
		sdk.NewCoins(),
		simtestutil.DefaultGenTxGas,
		"testing-chain-id",
		[]uint64{0},
		[]uint64{0},
		testPrivKeys[0],
	)
	require.NoError(t, err)
	ctx = ctx.WithIsCheckTx(true)
	_, err = decorator.AnteHandle(ctx, tx, false, mockAnteHandler)
	require.NoError(t, err)

	tx, err = simtestutil.GenSignedMockTx(
		rand.New(rand.NewSource(time.Now().UnixNano())),
		txConfig,
		[]sdk.Msg{
			&banktypes.MsgSend{
				FromAddress: testAddresses[0].String(),
				ToAddress:   creators[0].String(),
				Amount:      t2,
			},
			&banktypes.MsgSend{
				FromAddress: testAddresses[0].String(),
				ToAddress:   creators[3].String(),
				Amount:      t2,
			},
		},
		sdk.NewCoins(),
		simtestutil.DefaultGenTxGas,
		"testing-chain-id",
		[]uint64{0},
		[]uint64{0},
		testPrivKeys[0],
	)
	require.NoError(t, err)
	ctx = ctx.WithIsCheckTx(true)
	_, err = decorator.AnteHandle(ctx, tx, false, mockAnteHandler)
	require.ErrorContainsf(t, err, "has reached the quota target", "quota check")
}
