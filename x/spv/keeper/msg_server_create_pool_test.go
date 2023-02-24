package keeper_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	app "github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/require"
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
	app, k, _, wctx := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(wctx)
	_ = k
	//pa := types.Params{Submitter: []sdk.AccAddress{acc}}

	req := types.MsgCreatePool{Creator: acc.String(), ProjectIndex: 4, PoolName: "hello", Apy: "7.8", TargetTokenAmount: sdk.NewCoin("demo", sdk.NewInt(322))}
	_, err = app.CreatePool(ctx, &req)
	require.Error(t, err)

	req = types.MsgCreatePool{Creator: "invalid address", ProjectIndex: 1, PoolName: "hello", Apy: "7.8", TargetTokenAmount: sdk.NewCoin("demo", sdk.NewInt(322))}
	_, err = app.CreatePool(ctx, &req)
	require.Error(t, err)

	req = types.MsgCreatePool{Creator: acc.String(), ProjectIndex: 1, PoolName: "hello", Apy: "7.8", TargetTokenAmount: sdk.NewCoin("demo", sdk.NewInt(322))}
	_, err = app.CreatePool(ctx, &req)
	require.Error(t, err)

	// invalid pay freq
	req = types.MsgCreatePool{Creator: acc.String(), ProjectIndex: 1, PoolName: "hello", Apy: "7.8", TargetTokenAmount: sdk.NewCoin("demo", sdk.NewInt(322))}
	_, err = app.CreatePool(ctx, &req)
	require.Error(t, err)

	// create the first pool apy 7.8%
	req = types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 1, PoolName: "hello", Apy: "7.8", TargetTokenAmount: sdk.NewCoin("ausdc", sdk.NewInt(322))}
	_, err = app.CreatePool(ctx, &req)
	require.NoError(t, err)

	// duplicate pool
	req = types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 1, PoolName: "hello2", Apy: "7.8", TargetTokenAmount: sdk.NewCoin("ausdc", sdk.NewInt(4322))}
	_, err = app.CreatePool(ctx, &req)
	fmt.Printf(">>>%v\n", err)
	require.Error(t, err)
}

func TestMsgSERvCreatePoolApyCheck(t *testing.T) {

	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)
	app, k, _, wctx := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(wctx)

	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 1, PoolName: "hello", Apy: "7.8", TargetTokenAmount: sdk.NewCoin("ausdc", sdk.NewInt(322))}
	_, err := app.CreatePool(ctx, &req)
	require.NoError(t, err)

	poolsIndex := []string{"0x907ef9f3822d51da01347e81d1d859b95caf42b18ccc8a266044f5900b8c1371", "0x69da3ffa942f73a8bc9751541c8675cf28a4ffb68ef89559502391e06d37804a"}

	p1, ok := k.GetPools(ctx, poolsIndex[0])
	require.True(t, ok)

	p2, ok := k.GetPools(ctx, poolsIndex[1])
	require.True(t, ok)

	bapy := sdk.NewDecWithPrec(12, 2)
	coin := sdk.NewCoin("ausdc", sdk.NewIntFromUint64(100000000))

	interest1 := bapy.MulInt(coin.Amount)

	interest2 := p1.Apy.MulInt(p1.TargetAmount.Amount)
	interest3 := p2.Apy.MulInt(p2.TargetAmount.Amount)

	total := interest2.Add(interest3)
	require.True(t, total.Sub(interest1).Abs().LT(sdk.NewDecWithPrec(1, 8)))
}
