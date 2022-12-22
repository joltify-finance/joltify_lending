package issuance

import (
	"fmt"

	"github.com/joltify-finance/joltify_lending/x/third_party/issuance/keeper"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/issuance/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the store state from a genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, accountKeeper types2.AccountKeeper, gs types2.GenesisState) {
	if err := gs.Validate(); err != nil {
		panic(fmt.Sprintf("failed to validate %s genesis state: %s", types2.ModuleName, err))
	}

	// check if the module account exists
	moduleAcc := accountKeeper.GetModuleAccount(ctx, types2.ModuleAccountName)
	if moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", types2.ModuleAccountName))
	}

	k.SetParams(ctx, gs.Params)

	for _, supply := range gs.Supplies {
		k.SetAssetSupply(ctx, supply, supply.GetDenom())
	}

	for _, asset := range gs.Params.Assets {
		if asset.RateLimit.Active {
			_, found := k.GetAssetSupply(ctx, asset.Denom)
			if !found {
				k.CreateNewAssetSupply(ctx, asset.Denom)
			}
		}
	}
}

// ExportGenesis export genesis state for issuance module
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types2.GenesisState {
	params := k.GetParams(ctx)
	supplies := k.GetAllAssetSupplies(ctx)
	return types2.NewGenesisState(params, supplies)
}
