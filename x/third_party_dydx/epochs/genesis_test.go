package epochs_test

import (
	"testing"
	"time"

	keepertest "github.com/joltify-finance/joltify_lending/dydx_helper/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/epochs"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/epochs/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		EpochInfoList: []types.EpochInfo{
			{
				Name:                "0",
				Duration:            60,
				FastForwardNextTick: true,
			},
			{
				Name:                "1",
				Duration:            60,
				FastForwardNextTick: true,
			},
		},
	}

	expectedExportState := types.GenesisState{
		EpochInfoList: []types.EpochInfo{
			{
				Name:                "0",
				Duration:            60,
				FastForwardNextTick: true,
			},
			{
				Name:                "1",
				Duration:            60,
				FastForwardNextTick: true,
			},
		},
	}

	ctx, k, _ := keepertest.EpochsKeeper(t)
	initCtx := ctx.WithBlockTime(time.Unix(1800000000, 0))
	epochs.InitGenesis(initCtx, *k, genesisState)
	got := epochs.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.ElementsMatch(t, expectedExportState.EpochInfoList, got.EpochInfoList)
	// this line is used by starport scaffolding # genesis/test/assert
}
