package keeper_test

import (
	"encoding/hex"
	"testing"

	app2 "github.com/joltify-finance/joltify_lending/app"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	types2 "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreatePoolMsgServerCreate(t *testing.T) {
	app2.SetSDKConfig()
	app, srv, wctx := setupMsgServer(t)
	k := &app.VaultKeeper

	sk := ed25519.GenPrivKey()
	desc := types2.NewDescription("tester", "testId", "www.test.com", "aaa", "aaa")
	creatorStr := "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0"
	creator, err := sdk.AccAddressFromBech32(creatorStr)
	assert.Nil(t, err)
	valAddr, err := sdk.ValAddressFromHex(hex.EncodeToString(creator.Bytes()))
	assert.Nil(t, err)
	testValidator, err := types2.NewValidator(valAddr, sk.PubKey(), desc)
	require.NoError(t, err)

	pubkey := "joltpub1addwnpepqdq8w6407qzrqd39pan2zrq3a5n3zj5tcuf6ggc0px5jjylq5wsh2zk42cg"

	ctx := sdk.UnwrapSDKContext(wctx)
	historyInfo := types2.HistoricalInfo{
		Valset: types2.Validators{testValidator},
	}
	app.GetStakingKeeper().SetHistoricalInfo(ctx, int64(1), &historyInfo)

	expected := &types.MsgCreateCreatePool{Creator: creator, BlockHeight: "1", PoolPubKey: pubkey}
	_, err = srv.CreateCreatePool(wctx, expected)
	require.NoError(t, err)

	rst, found := k.GetCreatePool(ctx, expected.BlockHeight)
	require.True(t, found)
	assert.Equal(t, expected.PoolPubKey, rst.Proposal[0].PoolPubKey)
}

func TestCreatePoolMsgServerCreateNotValidator(t *testing.T) {
	app2.SetSDKConfig()
	app, srv, wctx := setupMsgServer(t)
	k := &app.VaultKeeper
	ctx := sdk.UnwrapSDKContext(wctx)

	sk := ed25519.GenPrivKey()
	desc := types2.NewDescription("tester", "testId", "www.test.com", "aaa", "aaa")

	valAddr := sk.PubKey().Address().Bytes()
	testValidator, err := types2.NewValidator(valAddr, sk.PubKey(), desc)
	require.NoError(t, err)
	historyInfo := types2.HistoricalInfo{
		Valset: types2.Validators{testValidator},
	}
	app.GetStakingKeeper().SetHistoricalInfo(ctx, int64(1), &historyInfo)

	creatorStr := "jolt18mdnq8x9m07dryymlyf8jknagp87yga0hpe7n6"
	creator, err := sdk.AccAddressFromBech32(creatorStr)
	assert.Nil(t, err)
	expected := &types.MsgCreateCreatePool{Creator: creator, BlockHeight: "1", PoolPubKey: creatorStr}
	_, err = srv.CreateCreatePool(wctx, expected)
	require.NoError(t, err)
	_, found := k.GetCreatePool(ctx, expected.BlockHeight)
	require.False(t, found)
}
