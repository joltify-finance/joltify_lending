package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/x/quota/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
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

	cqNoModuleName := types.CoinsQuota{
		ModuleName: "",
		History:    []*types.HistoricalAmount{&ht},
		CoinsSum:   sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
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
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				Params:        types.DefaultParams(),
				AllCoinsQuota: []types.CoinsQuota{cq},
			},
			valid: true,
		},

		{
			desc: "valid genesis state with no quota",
			genState: &types.GenesisState{
				Params:        types.DefaultParams(),
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
