package blocktime_test

import (
	"testing"

	testapp "github.com/dydxprotocol/v4-chain/protocol/testutil/app"
	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/blocktime"
	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/blocktime/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()
	got := blocktime.ExportGenesis(ctx, tApp.App.BlockTimeKeeper)
	require.NotNil(t, got)
	require.Equal(t, types.DefaultGenesis(), got)
}
