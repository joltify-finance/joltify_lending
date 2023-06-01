package keeper_test

import (
	"testing"

	tmlog "github.com/tendermint/tendermint/libs/log"
	tmprototypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

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
	require.EqualValues(t, params.FirstProvisions, k.CurrentProvision(ctx))
	require.EqualValues(t, params.FirstProvisions, k.FirstProvision(ctx))
	require.EqualValues(t, params.Unit, k.Unit(ctx))
	require.EqualValues(t, params.CommunityProvisions, k.CommunityProvision(ctx))
	require.EqualValues(t, params.HalfCount, k.HalfCount(ctx))
}
