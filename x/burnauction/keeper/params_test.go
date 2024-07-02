package keeper_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/x/burnauction/types"

	testkeeper "github.com/joltify-finance/joltify_lending/testutil/keeper"

	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, _, _, ctx := testkeeper.BurnauctionKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.BurnThreshold, k.Burnthreshold(ctx))
}
