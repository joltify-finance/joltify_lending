package mint_test

import (
	"testing"

	tmlog "github.com/cometbft/cometbft/libs/log"
	tmprototypes "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtime "github.com/cometbft/cometbft/types/time"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/stretchr/testify/assert"

	"github.com/joltify-finance/joltify_lending/x/mint"
	"github.com/joltify-finance/joltify_lending/x/mint/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # genesis/test/state
	}
	lg := tmlog.TestingLogger()
	tApp := app.NewTestApp(lg, t.TempDir())
	ctx := tApp.NewContext(true, tmprototypes.Header{Height: 1, Time: tmtime.Now()})
	k := tApp.GetMintKeeper()
	mint.InitGenesis(ctx, k, genesisState)
	got := mint.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	history := got.Params
	defaultparam := types.DefaultParams()
	assert.Equal(t, history.CommunityProvisions, defaultparam.CommunityProvisions)
	assert.Equal(t, history.CurrentProvisions, defaultparam.CurrentProvisions)
	assert.Equal(t, history.HalfCount, defaultparam.GetHalfCount())
	assert.Equal(t, history.FirstProvisions, defaultparam.FirstProvisions)
}
