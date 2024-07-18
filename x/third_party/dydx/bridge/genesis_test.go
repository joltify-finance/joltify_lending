package bridge_test

import (
	"testing"

	testapp "github.com/dydxprotocol/v4-chain/protocol/testutil/app"
	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/bridge"
	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/bridge/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()
	got := bridge.ExportGenesis(ctx, tApp.App.BridgeKeeper)
	require.NotNil(t, got)
	require.Equal(t, types.DefaultGenesis(), got)
}
