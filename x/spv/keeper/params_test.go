package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SpvKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
