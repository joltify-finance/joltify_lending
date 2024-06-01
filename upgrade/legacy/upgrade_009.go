package v1

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ethermint "github.com/evmos/ethermint/types"
)

const V009UpgradeName = "v009_upgrade"

func CreateUpgradeHandlerForV009Upgrade(
	mm *module.Manager,
	configurator module.Configurator,
	ak authkeeper.AccountKeeper,
	bk bankkeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for i := 0; i < 5; i++ {
			ctx.Logger().Info("we update the parameter for v009")
		}

		allAccounts := ak.GetAllAccounts(ctx)
		for _, ac := range allAccounts {

			ethAcc, ok := ac.(ethermint.EthAccountI)
			if ok {
				// we check whether it has the balance
				coins := bk.GetAllBalances(ctx, ethAcc.GetAddress())
				if coins.IsZero() {
					continue
				}
				fmt.Printf(">>>>%v coins it has %v\n", ethAcc.GetAddress().String(), coins)
				bAcc := authtypes.NewBaseAccount(ethAcc.GetAddress(), ethAcc.GetPubKey(), ethAcc.GetAccountNumber(), ethAcc.GetSequence())
				ak.SetAccount(ctx, bAcc)
			}
			continue
		}
		return mm.RunMigrations(ctx, configurator, vm)
	}
}
