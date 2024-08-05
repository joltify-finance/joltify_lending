package keeper_test

import (
	"testing"

	tmlog "cosmossdk.io/log"
	tmprototypes "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtime "github.com/cometbft/cometbft/types/time"

	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/x/mint/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	lg := tmlog.TestingLogger()
	tApp := app.NewTestApp(lg, t.TempDir())
	ctx := tApp.NewContext(true, tmprototypes.Header{Height: 1, Time: tmtime.Now()})
	k := tApp.GetMintKeeper()

	params := types.DefaultParams()
	k.SetParams(ctx, params)
	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.FirstProvisions, k.FirstProvision(ctx))
	require.EqualValues(t, params.Unit, k.Unit(ctx))
	require.EqualValues(t, params.NodeSPY, k.GetNodeAPY(ctx))
}
