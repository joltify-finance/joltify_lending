package rewards_test

import (
	"testing"

	testapp "github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/app"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/rewards"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/rewards/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()
	k := tApp.App.RewardsKeeper

	rewards.InitGenesis(ctx, k, genesisState)
	got := rewards.ExportGenesis(ctx, k)
	require.NotNil(t, got)
}
