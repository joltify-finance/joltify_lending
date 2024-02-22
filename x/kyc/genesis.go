package kyc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/client"
	"github.com/joltify-finance/joltify_lending/x/kyc/keeper"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	if client.MAINNETFLAG == "false" {
		projects := types.GenerateTestProjects()
		for _, p := range projects {
			_, err := k.SetProject(ctx, p)
			if err != nil {
				panic("fail to set the test projects")
			}
		}
	}

	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	// this line is used by starport scaffolding # genesis/module/export
	return genesis
}
