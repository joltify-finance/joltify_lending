package burnauction_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/testutil/nullify"
	"github.com/joltify-finance/joltify_lending/x/burnauction"
	"github.com/joltify-finance/joltify_lending/x/burnauction/types"

	keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, _, _, ctx := keepertest.BurnauctionKeeper(t)
	burnauction.InitGenesis(ctx, *k, genesisState)
	got := burnauction.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
