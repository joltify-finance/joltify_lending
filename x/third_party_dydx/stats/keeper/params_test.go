package keeper_test

import (
	"testing"
	"time"

	testapp "github.com/joltify-finance/joltify_lending/dydx_helper/testutil/app"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/stats/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()
	k := tApp.App.StatsKeeper

	require.Equal(t, types.DefaultGenesis().Params, k.GetParams(ctx))
}

func TestSetParams_Success(t *testing.T) {
	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()
	k := tApp.App.StatsKeeper

	params := types.Params{
		WindowDuration: time.Duration(30 * 24 * time.Hour),
	}
	require.NoError(t, params.Validate())

	require.NoError(t, k.SetParams(ctx, params))
	require.Equal(t, params, k.GetParams(ctx))
}
