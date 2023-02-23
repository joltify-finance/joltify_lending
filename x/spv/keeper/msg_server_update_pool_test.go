package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/require"
)

func verifyInterest(t *testing.T, bApy, apy1, apy2 sdk.Dec, totalAmount, apy1Amount, apy2Amount sdkmath.Int) {
	interest1 := bApy.MulInt(totalAmount)
	interest2 := apy1.MulInt(apy1Amount)
	interest3 := apy2.MulInt(apy2Amount)
	total := interest2.Add(interest3)
	require.True(t, total.Sub(interest1).Abs().LT(sdk.NewDecWithPrec(1, 8)))
}

func TestMsgSERvUpdatePool(t *testing.T) {

	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)
	app, k, wctx := setupMsgServer(t)
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

	verifyInterest(t, bapy, p1.Apy, p2.Apy, coin.Amount, p1.TargetAmount.Amount, p2.TargetAmount.Amount)

	reqUpdate := types.MsgUpdatePool{
		Creator:           p1.OwnerAddress.String(),
		PoolIndex:         poolsIndex[0],
		PoolName:          "updatedpool-1",
		PoolApy:           "0.3",
		TargetTokenAmount: sdk.NewCoin("ausdc", sdk.NewInt(550)),
	}
	_, err = app.UpdatePool(ctx, &reqUpdate)
	require.NoError(t, err)
	p1, ok = k.GetPools(ctx, poolsIndex[0])
	require.True(t, ok)

	p2, ok = k.GetPools(ctx, poolsIndex[1])
	require.True(t, ok)

	require.EqualValues(t, "updatedpool-1", p1.PoolName)
	require.EqualValues(t, "updatedpool-1", p2.PoolName)
	verifyInterest(t, bapy, p1.Apy, p2.Apy, coin.Amount, p1.TargetAmount.Amount, p2.TargetAmount.Amount)

	reqUpdate = types.MsgUpdatePool{
		Creator:           p1.OwnerAddress.String(),
		PoolIndex:         poolsIndex[1],
		PoolName:          "updatedpool-2",
		PoolApy:           "0.5",
		TargetTokenAmount: sdk.NewCoin("ausdc", sdk.NewInt(750)),
	}
	_, err = app.UpdatePool(ctx, &reqUpdate)
	require.NoError(t, err)
	p1, ok = k.GetPools(ctx, poolsIndex[0])
	require.True(t, ok)
	p2, ok = k.GetPools(ctx, poolsIndex[1])
	require.True(t, ok)
	require.EqualValues(t, "updatedpool-2", p1.PoolName)
	require.EqualValues(t, "updatedpool-2", p2.PoolName)
	verifyInterest(t, bapy, p1.Apy, p2.Apy, coin.Amount, p1.TargetAmount.Amount, p2.TargetAmount.Amount)
}

func TestMsgSERvUpdatePoolWithError(t *testing.T) {

	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)
	app, _, wctx := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(wctx)

	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 1, PoolName: "hello", Apy: "7.8", TargetTokenAmount: sdk.NewCoin("ausdc", sdk.NewInt(322))}
	_, err := app.CreatePool(ctx, &req)
	require.NoError(t, err)

	poolsIndex := []string{"0x907ef9f3822d51da01347e81d1d859b95caf42b18ccc8a266044f5900b8c1371", "0x69da3ffa942f73a8bc9751541c8675cf28a4ffb68ef89559502391e06d37804a"}

	reqUpdate := types.MsgUpdatePool{
		Creator:           "invalid",
		PoolIndex:         poolsIndex[0],
		PoolName:          "updatedpool-1",
		PoolApy:           "0.3",
		TargetTokenAmount: sdk.NewCoin("ausdc", sdk.NewInt(550)),
	}

	_, err = app.UpdatePool(ctx, &reqUpdate)
	//fmt.Println(err)
	//require.Error(t, err)
	require.Error(t, err, "invalid address invalid: invalid address")

	reqUpdate = types.MsgUpdatePool{
		Creator:           "jolt10nsg95f7geuhf9dm8v2r4d7jxvnjk23aaufq3p",
		PoolIndex:         "3322323",
		PoolName:          "updatedpool-1",
		PoolApy:           "0.3",
		TargetTokenAmount: sdk.NewCoin("ausdc", sdk.NewInt(550)),
	}

	_, err = app.UpdatePool(ctx, &reqUpdate)
	require.ErrorContains(t, err, "pool cannot be found")

	reqUpdate = types.MsgUpdatePool{
		Creator:           "jolt10nsg95f7geuhf9dm8v2r4d7jxvnjk23aaufq3p",
		PoolIndex:         poolsIndex[0],
		PoolName:          "updatedpool-1",
		PoolApy:           "0.3",
		TargetTokenAmount: sdk.NewCoin("ausdc", sdk.NewInt(550)),
	}

	_, err = app.UpdatePool(ctx, &reqUpdate)
	require.ErrorContains(t, err, "is not authorized to update the pool")

}
