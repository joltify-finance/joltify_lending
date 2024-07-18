package stats_test

import (
	"testing"

	testapp "github.com/dydxprotocol/v4-chain/protocol/testutil/app"
	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/stats"
	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/stats/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()
	got := stats.ExportGenesis(ctx, tApp.App.StatsKeeper)
	require.NotNil(t, got)
	require.Equal(t, types.DefaultGenesis(), got)
}
