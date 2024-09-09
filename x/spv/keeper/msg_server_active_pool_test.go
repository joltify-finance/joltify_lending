package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/utils"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/require"
)

func TestActivatePool(t *testing.T) {
	config := sdk.NewConfig()
	utils.SetBech32AddressPrefixes(config)

	lapp, _, _, _, _, wctx := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(wctx)

	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 1, PoolName: "hello", Apy: []string{"7.8", "7.2"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdkmath.NewInt(322)), sdk.NewCoin("ausdc", sdkmath.NewInt(222))}}
	resp, err := lapp.CreatePool(ctx, &req)
	require.NoError(t, err)

	// test not the owner
	_, err = lapp.ActivePool(ctx, &types.MsgActivePool{
		PoolIndex: resp.PoolIndex[0],
		Creator:   "jolt10nsg95f7geuhf9dm8v2r4d7jxvnjk23aaufq3p",
	})

	require.ErrorContains(t, err, "is not authorized to active the pool")

	_, err = lapp.ActivePool(ctx, &types.MsgActivePool{
		PoolIndex: "invalid address",
		Creator:   "jolt10nsg95f7geuhf9dm8v2r4d7jxvnjk23aaufq3p",
	})
	require.ErrorContains(t, err, "pool cannot be found")

	_, err = lapp.ActivePool(ctx, &types.MsgActivePool{
		PoolIndex: "invalid address",
		Creator:   "jolt10nsg95f7geuhf9dm8v2k23aaufq3p",
	})
	require.ErrorContains(t, err, "invalid address")

	_, err = lapp.ActivePool(ctx, &types.MsgActivePool{
		PoolIndex: resp.GetPoolIndex()[0],
		Creator:   "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0",
	})
	require.NoError(t, err)

	_, err = lapp.ActivePool(ctx, &types.MsgActivePool{
		PoolIndex: resp.GetPoolIndex()[0],
		Creator:   "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0",
	})
	require.ErrorContains(t, err, "unexpected pool status")
}
