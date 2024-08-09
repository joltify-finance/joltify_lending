package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/x/mint/types"
)

func TestGetParams(t *testing.T) {
	lg := log.NewTestLogger(t)
	tApp := app.NewTestApp(lg, t.TempDir())
	ctx := tApp.NewContext(true)
	k := tApp.GetMintKeeper()

	params := types.DefaultParams()
	k.SetParams(ctx, params)
	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.FirstProvisions, k.FirstProvision(ctx))
	require.EqualValues(t, params.Unit, k.Unit(ctx))
	require.EqualValues(t, params.NodeSPY, k.GetNodeAPY(ctx))
}
