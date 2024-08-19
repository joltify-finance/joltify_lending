package bridge_test

import (
	"fmt"
	"testing"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	testapp "github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/app"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/bridge"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/bridge/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {

	ma := authtypes.NewModuleAddress(types.ModuleName)
	fmt.Printf(">>>>>%v\n", ma.String())
	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()
	got := bridge.ExportGenesis(ctx, tApp.App.BridgeKeeper)
	require.NotNil(t, got)
	require.Equal(t, types.DefaultGenesis(), got)
}
