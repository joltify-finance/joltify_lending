package accountplus

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/accountplus/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/accountplus/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	err := k.SetGenesisState(ctx, data)
	if err != nil {
		panic(err)
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return types.GenesisState{
		Accounts: k.GetAllAccountStates(ctx),
	}
}
