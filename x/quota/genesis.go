package quota

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/quota/keeper"
	"github.com/joltify-finance/joltify_lending/x/quota/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the quota
	for _, elem := range genState.AllCoinsQuota {
		k.SetQuotaData(ctx, elem)
	}

	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.AllCoinsQuota = k.GetAllQuota(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
