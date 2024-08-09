package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	testkeeper "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/x/quota/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.QuotaKeeper(t)
	params := testParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

func TestSetNonSortedParas(t *testing.T) {
	k, ctx := testkeeper.QuotaKeeper(t)
	params := types.DefaultParams()
	// unsorted, err := sdk.ParseCoinsNormalized("10000000ujolt,10000000uatom")

	c1 := sdk.NewCoin("ujolt", sdkmath.NewInt(10))
	c2 := sdk.NewCoin("uatom", sdkmath.NewInt(32))
	unsorted := []sdk.Coin{c1, c2}
	params.Targets[0].CoinsSum = unsorted

	require.Panics(t, func() { k.SetParams(ctx, params) })
}
