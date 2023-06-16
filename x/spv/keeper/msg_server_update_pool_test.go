package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/require"
)

func TestMsgSERvUpdatePool(t *testing.T) {
	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)
	app, k, _, _, wctx := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(wctx)

	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 1, PoolName: "hello", Apy: []string{"7.8", "7.2"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdk.NewInt(322)), sdk.NewCoin("ausdc", sdk.NewInt(322))}}

	_, err := app.CreatePool(ctx, &req)
	require.NoError(t, err)

	poolsIndex := []string{"0x907ef9f3822d51da01347e81d1d859b95caf42b18ccc8a266044f5900b8c1371", "0x69da3ffa942f73a8bc9751541c8675cf28a4ffb68ef89559502391e06d37804a"}

	p1, ok := k.GetPools(ctx, poolsIndex[0])
	require.True(t, ok)

	p2, ok := k.GetPools(ctx, poolsIndex[1])
	require.True(t, ok)

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

	require.EqualValues(t, "updatedpool-1-junior", p1.PoolName)

	reqUpdate = types.MsgUpdatePool{
		Creator:           p1.OwnerAddress.String(),
		PoolIndex:         poolsIndex[1],
		PoolName:          "updatedpool-2",
		PoolApy:           "0.5",
		TargetTokenAmount: sdk.NewCoin("ausdc", sdk.NewInt(750)),
	}
	_, err = app.UpdatePool(ctx, &reqUpdate)
	require.NoError(t, err)
	p2, ok = k.GetPools(ctx, poolsIndex[1])
	require.True(t, ok)
	require.EqualValues(t, "updatedpool-2-senior", p2.PoolName)
}

func TestMsgSERvUpdatePoolWithError(t *testing.T) {
	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)
	app, _, _, _, wctx := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(wctx)

	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 1, PoolName: "hello", Apy: []string{"7.8", "7.2"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdk.NewInt(322)), sdk.NewCoin("ausdc", sdk.NewInt(322))}}
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
	// fmt.Println(err)
	// require.Error(t, err)
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
		TargetTokenAmount: sdk.NewCoin("invalid", sdk.NewInt(550)),
	}

	_, err = app.UpdatePool(ctx, &reqUpdate)
	require.ErrorContains(t, err, "target amount denom is not matched")

	reqUpdate = types.MsgUpdatePool{
		Creator:           "jolt10nsg95f7geuhf9dm8v2r4d7jxvnjk23aaufq3p",
		PoolIndex:         poolsIndex[0],
		PoolName:          "updatedpool-1",
		PoolApy:           "0.3",
		TargetTokenAmount: sdk.NewCoin("invalid", sdk.NewInt(550)),
	}

	_, err = app.UpdatePool(ctx, &reqUpdate)
	require.ErrorContains(t, err, "target amount denom is not matched")

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
