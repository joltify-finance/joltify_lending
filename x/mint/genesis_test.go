package mint_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	tmlog "cosmossdk.io/log"
	tmprototypes "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtime "github.com/cometbft/cometbft/types/time"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/stretchr/testify/assert"

	"github.com/joltify-finance/joltify_lending/x/mint"
	"github.com/joltify-finance/joltify_lending/x/mint/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	lg := tmlog.TestingLogger()
	tApp := app.NewTestApp(lg, t.TempDir())
	ctx := tApp.NewContext(true, tmprototypes.Header{Height: 1, Time: tmtime.Now()})
	k := tApp.GetMintKeeper()

	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		HistoricalDistInfo: &types.HistoricalDistInfo{
			PayoutTime:     ctx.BlockTime(),
			TotalMintCoins: sdk.NewCoins(),
		},
	}

	k.SetParams(ctx, genesisState.Params)
	k.SetDistInfo(ctx, *genesisState.HistoricalDistInfo)
	_ = genesisState
	got := mint.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	history := got.Params

	historyTime := got.HistoricalDistInfo.PayoutTime

	defaultparam := types.DefaultParams()
	assert.Equal(t, history.FirstProvisions, defaultparam.FirstProvisions)
	assert.Equal(t, history.NodeSPY, defaultparam.NodeSPY)
	assert.True(t, historyTime.Equal(ctx.BlockTime()))
}
