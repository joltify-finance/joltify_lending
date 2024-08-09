package keeper_test

import (
	"testing"

	testapp "github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/app"
	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()

	logger := tApp.App.BridgeKeeper.Logger(ctx)
	require.NotNil(t, logger)
}
