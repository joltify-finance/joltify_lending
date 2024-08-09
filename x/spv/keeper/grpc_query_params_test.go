package keeper_test

import (
	"testing"

	testkeeper "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	keeper, _, _, _, _, ctx := testkeeper.SpvKeeper(t)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
