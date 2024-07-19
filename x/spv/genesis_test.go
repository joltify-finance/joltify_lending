package spv_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/testutil/nullify"
	"github.com/joltify-finance/joltify_lending/x/spv"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, _, _, _, _, ctx := keepertest.SpvKeeper(t)
	spv.InitGenesis(sdk.UnwrapSDKContext(ctx), *k, genesisState)
	got := spv.ExportGenesis(sdk.UnwrapSDKContext(ctx), *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
