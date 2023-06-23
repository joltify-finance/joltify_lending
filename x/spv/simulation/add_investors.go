package simulation

import (
	"math/rand"

	"github.com/joltify-finance/joltify_lending/x/spv/keeper"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func SimulateMsgAddInvestors(
	_ types.AccountKeeper,
	_ types.BankKeeper,
	_ keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAddInvestors{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the AddInvestors simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "AddInvestors simulation not implemented"), nil, nil
	}
}
