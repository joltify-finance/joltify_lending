package blocktime_test

import (
	"testing"

	keepertest "github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/blocktime"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/blocktime/types"
	"github.com/stretchr/testify/require"
)

//func TestGenesis(t *testing.T) {
//	tApp := testapp.NewTestAppBuilder(t).Build()
//	ctx := tApp.InitChain()
//	got := blocktime.ExportGenesis(ctx, tApp.App.BlockTimeKeeper)
//	require.NotNil(t, got)
//	require.Equal(t, types.DefaultGenesis(), got)
//}

func TestGenesis(t *testing.T) {
	expected := types.DefaultGenesis()
	ctx, k, _ := keepertest.BlcokTimeKeepers(t)
	blocktime.InitGenesis(ctx, *k, *expected)
	actual := blocktime.ExportGenesis(ctx, *k)
	require.NotNil(t, actual)
	require.Equal(t, types.DefaultGenesis(), actual)
}
