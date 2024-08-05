package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

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

func testGenesis() *types.GenesisState {
	return &types.GenesisState{
		AllCoinsQuota: []types.CoinsQuota{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: testParams(),
	}
}

func TestGenesisState_Validate(t *testing.T) {
	ht := types.HistoricalAmount{
		100,
		sdk.NewCoins(sdk.NewCoin("test", sdkmath.NewInt(100))),
		1,
	}

	cq := types.CoinsQuota{
		ModuleName: "testmodule",
		History:    []*types.HistoricalAmount{&ht},
		CoinsSum:   sdk.NewCoins(sdk.NewCoin("test", sdkmath.NewInt(100))),
	}

	cqNoModuleName := types.CoinsQuota{
		ModuleName: "",
		History:    []*types.HistoricalAmount{&ht},
		CoinsSum:   sdk.NewCoins(sdk.NewCoin("test", sdkmath.NewInt(100))),
	}

	cqNoCoinsNoCoins := types.CoinsQuota{
		ModuleName: "",
		History:    []*types.HistoricalAmount{&ht},
		CoinsSum:   sdk.Coins{},
	}

	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: testGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				Params:        testParams(),
				AllCoinsQuota: []types.CoinsQuota{cq},
			},
			valid: true,
		},

		{
			desc: "valid genesis state with no quota",
			genState: &types.GenesisState{
				Params:        testParams(),
				AllCoinsQuota: []types.CoinsQuota{},
			},
			valid: true,
		},
		{
			desc: "invalid genesis state with no module name",
			genState: &types.GenesisState{
				Params:        types.DefaultParams(),
				AllCoinsQuota: []types.CoinsQuota{cqNoModuleName},
			},
			valid: false,
		},
		{
			desc: "invalid genesis state with no coins",
			genState: &types.GenesisState{
				Params:        types.DefaultParams(),
				AllCoinsQuota: []types.CoinsQuota{cqNoCoinsNoCoins},
			},
			valid: false,
		},

		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
