package keeper_test

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	app "github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func generateNAddr(n int) []string {
	addresses := make([]string, n)
	for i := 0; i < n; i++ {
		pk := ed25519.GenPrivKey().PubKey()
		addr := pk.Address().Bytes()
		a := sdk.AccAddress(addr)
		addresses[i] = a.String()
	}
	return addresses
}

func TestMsgSERvCreatePool(t *testing.T) {
	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)

	acc, err := sdk.AccAddressFromBech32("jolt1p3jl6udk43vw0cvc5hjqrpnncsqmsz56wd32z8")
	require.NoError(t, err)
	app, k, wctx := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(wctx)
	_ = k
	//pa := types.Params{Submitter: []sdk.AccAddress{acc}}

	req := types.MsgCreatePool{Creator: acc.String(), ProjectIndex: 4, PoolName: "hello", Apy: "7.8", PayFreq: "3", TargetTokenAmount: sdk.NewCoin("demo", sdk.NewInt(322))}
	_, err = app.CreatePool(ctx, &req)
	require.Error(t, err)

	req = types.MsgCreatePool{Creator: "invalid address", ProjectIndex: 1, PoolName: "hello", Apy: "7.8", PayFreq: "3", TargetTokenAmount: sdk.NewCoin("demo", sdk.NewInt(322))}
	_, err = app.CreatePool(ctx, &req)
	require.Error(t, err)

	req = types.MsgCreatePool{Creator: acc.String(), ProjectIndex: 1, PoolName: "hello", Apy: "7.8", PayFreq: "3", TargetTokenAmount: sdk.NewCoin("demo", sdk.NewInt(322))}
	_, err = app.CreatePool(ctx, &req)
	require.Error(t, err)

	// invalid pay freq
	req = types.MsgCreatePool{Creator: acc.String(), ProjectIndex: 1, PoolName: "hello", Apy: "7.8", PayFreq: "300", TargetTokenAmount: sdk.NewCoin("demo", sdk.NewInt(322))}
	_, err = app.CreatePool(ctx, &req)
	require.Error(t, err)

	// create the first pool apy 7.8%
	req = types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", IsSenior: true, ProjectIndex: 1, PoolName: "hello", Apy: "7.8", PayFreq: "6", TargetTokenAmount: sdk.NewCoin("demo", sdk.NewInt(322))}
	_, err = app.CreatePool(ctx, &req)
	require.NoError(t, err)

	// duplicate pool
	req = types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", IsSenior: true, ProjectIndex: 1, PoolName: "hello2", Apy: "7.8", PayFreq: "7", TargetTokenAmount: sdk.NewCoin("demo", sdk.NewInt(4322))}
	_, err = app.CreatePool(ctx, &req)
	fmt.Printf(">>>%v\n", err)
	require.Error(t, err)

	// create the second pool apy 17.8%
	req = types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", IsSenior: false, ProjectIndex: 1, PoolName: "hello", Apy: "17.8", PayFreq: "7", TargetTokenAmount: sdk.NewCoin("demo", sdk.NewInt(322))}
	_, err = app.CreatePool(ctx, &req)
	require.NoError(t, err)

}
