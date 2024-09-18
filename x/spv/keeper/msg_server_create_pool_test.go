package keeper_test

import (
	"fmt"
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/utils"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/require"
)

func TestMsgSERvCreatePool(t *testing.T) {
	config := sdk.NewConfig()
	utils.SetBech32AddressPrefixes(config)

	acc, err := sdk.AccAddressFromBech32("jolt1p3jl6udk43vw0cvc5hjqrpnncsqmsz56wd32z8")
	require.NoError(t, err)
	lapp, k, _, _, _, wctx := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(wctx)
	_ = k

	req := types.MsgCreatePool{Creator: acc.String(), ProjectIndex: 4, PoolName: "hello", Apy: []string{"7.8", "7.2"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdkmath.NewInt(322)), sdk.NewCoin("ausdc", sdkmath.NewInt(322))}}
	_, err = lapp.CreatePool(ctx, &req)
	require.Error(t, err)

	req = types.MsgCreatePool{Creator: "invalid address", ProjectIndex: 1, PoolName: "hello", Apy: []string{"7.8", "7.2"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdkmath.NewInt(322)), sdk.NewCoin("ausdc", sdkmath.NewInt(322))}}
	_, err = lapp.CreatePool(ctx, &req)
	require.Error(t, err)

	req = types.MsgCreatePool{Creator: acc.String(), ProjectIndex: 1, PoolName: "hello", Apy: []string{"7.8", "7.2"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdkmath.NewInt(322)), sdk.NewCoin("ausdc", sdkmath.NewInt(322))}}
	_, err = lapp.CreatePool(ctx, &req)
	require.Error(t, err)

	// invalid pay freq
	req = types.MsgCreatePool{Creator: acc.String(), ProjectIndex: 1, PoolName: "hello", Apy: []string{"7.8", "7.2"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdkmath.NewInt(322)), sdk.NewCoin("ausdc", sdkmath.NewInt(322))}}
	_, err = lapp.CreatePool(ctx, &req)
	require.Error(t, err)

	// invalid demon from market
	req = types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 1, PoolName: "hello", Apy: []string{"7.8", "7.2"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("invalid", sdkmath.NewInt(322)), sdk.NewCoin("ausdc", sdkmath.NewInt(322))}}
	_, err = lapp.CreatePool(ctx, &req)
	require.ErrorContains(t, err, "invalid parameter from market: conversion factor")

	pa := k.GetParams(ctx)
	pa.Markets = append(pa.Markets, types.Moneymarket{Denom: "invalid", ConversionFactor: 6})
	k.SetParams(ctx, pa)
	_, err = lapp.CreatePool(ctx, &req)
	require.ErrorContains(t, err, "unsupported token")

	// create the first pool apy 7.8%
	req = types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 1, PoolName: "hello", Apy: []string{"7.8", "7.2"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdkmath.NewInt(322)), sdk.NewCoin("ausdc", sdkmath.NewInt(322))}}
	_, err = lapp.CreatePool(ctx, &req)
	require.NoError(t, err)

	// duplicate pool
	req = types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 1, PoolName: "hello2", Apy: []string{"7.8", "7.2"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdkmath.NewInt(4322)), sdk.NewCoin("ausdc", sdkmath.NewInt(322))}}
	_, err = lapp.CreatePool(ctx, &req)
	fmt.Printf(">>>%v\n", err)
	require.Error(t, err)
}

func TestMsgSERvCreatePoolApyCheck(t *testing.T) {
	config := sdk.NewConfig()
	utils.SetBech32AddressPrefixes(config)
	lapp, k, _, _, _, wctx := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(wctx)

	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 1, PoolName: "hello", Apy: []string{"7.8", "7.2"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdkmath.NewInt(322)), sdk.NewCoin("ausdc", sdkmath.NewInt(322))}}
	resp, err := lapp.CreatePool(ctx, &req)
	require.NoError(t, err)

	p1, ok := k.GetPools(ctx, resp.PoolIndex[0])
	require.True(t, ok)

	p2, ok := k.GetPools(ctx, resp.PoolIndex[1])
	require.True(t, ok)

	require.Equal(t, p1.Apy, sdkmath.LegacyMustNewDecFromStr("7.8"))
	require.Equal(t, p2.Apy, sdkmath.LegacyMustNewDecFromStr("7.2"))
}
