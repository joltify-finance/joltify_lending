package simulation

import (
	"math/rand"

	"github.com/joltify-finance/joltify_lending/x/kyc/keeper"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
)

func SimulateMsgUploadInvestor(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgUploadInvestor{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the UploadInvestor simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "UploadInvestor simulation not implemented"), nil, nil
	}
}
