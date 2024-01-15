package quota_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/testutil/nullify"
	"github.com/joltify-finance/joltify_lending/x/quota"
	"github.com/joltify-finance/joltify_lending/x/quota/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	ht := types.HistoricalAmount{
		100,
		sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
		1,
	}

	cq := types.CoinsQuota{
		ModuleName: "testmodule",
		History:    []*types.HistoricalAmount{&ht},
		CoinsSum:   sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
	}

	genesisState := types.GenesisState{
		Params:        types.DefaultParams(),
		AllCoinsQuota: []types.CoinsQuota{cq},
	}

	k, ctx := keepertest.QuotaKeeper(t)
	quota.InitGenesis(ctx, *k, genesisState)
	got := quota.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.AllCoinsQuota, got.AllCoinsQuota)
	require.Equal(t, genesisState.Params, got.Params)
	// this line is used by starport scaffolding # genesis/test/assert
}
