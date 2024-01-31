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

// NewParams creates a new Params instance
func testParams() types.Params {
	// the coin list is the amount of USD for the given token, 100jolt means 100 USD value of jolt
	quota, err := sdk.ParseCoinsNormalized("100000ujolt,1000000usdt")
	if err != nil {
		panic(err)
	}

	quotaAcc, err := sdk.ParseCoinsNormalized("10000000ujolt,100000000usdt")
	if err != nil {
		panic(err)
	}

	targets := types.Target{
		ModuleName:    "ibc",
		CoinsSum:      quota,
		HistoryLength: 512,
	}

	targets2 := types.Target{
		ModuleName:    "bridge",
		CoinsSum:      quota,
		HistoryLength: 512,
	}

	targetsAcc := types.Target{
		ModuleName:    "ibc",
		CoinsSum:      quotaAcc,
		HistoryLength: 512,
	}

	targets2Acc := types.Target{
		ModuleName:    "bridge",
		CoinsSum:      quotaAcc,
		HistoryLength: 512,
	}

	return types.Params{Targets: []*types.Target{&targets, &targets2}, PerAccounttargets: []*types.Target{&targetsAcc, &targets2Acc}}
}

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
		Params:        testParams(),
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
