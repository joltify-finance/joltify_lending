package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/types"
)

func SimulateMsgProposedOperations(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := simtypes.NoOpMsg(
			types.ModuleName,
			types.TypeMsgProposedOperations,
			"ProposedOperations simulation not implemented",
		)
		return msg, nil, nil
	}
}
