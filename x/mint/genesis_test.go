package mint_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"

	"github.com/joltify-finance/joltify_lending/app"

	"github.com/joltify-finance/joltify_lending/x/mint"
	"github.com/joltify-finance/joltify_lending/x/mint/types"
)

func TestGenesis(t *testing.T) {
	lg := log.NewTestLogger(t)
	tApp := app.NewTestApp(lg, t.TempDir())
	ctx := tApp.NewContext(true)
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
