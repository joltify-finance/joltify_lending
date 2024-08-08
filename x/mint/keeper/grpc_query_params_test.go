package keeper_test

import (
	"testing"

	"cosmossdk.io/log"

	"github.com/stretchr/testify/require"

	"github.com/joltify-finance/joltify_lending/app"

	"github.com/joltify-finance/joltify_lending/x/mint/types"
)

func TestParamsQuery(t *testing.T) {
	lg := log.NewNopLogger()
	tApp := app.NewTestApp(lg, t.TempDir())
	ctx := tApp.NewContext(true)

	keeper := tApp.GetMintKeeper()

	wctx := ctx
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
