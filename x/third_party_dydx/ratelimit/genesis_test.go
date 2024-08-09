package ratelimit_test

import (
	"testing"

	testapp "github.com/joltify-finance/joltify_lending/dydx_helper/testutil/app"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/ratelimit"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/ratelimit/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()
	got := ratelimit.ExportGenesis(ctx, tApp.App.RatelimitKeeper)
	require.NotNil(t, got)
	require.Equal(t, types.DefaultGenesis(), got)
}
