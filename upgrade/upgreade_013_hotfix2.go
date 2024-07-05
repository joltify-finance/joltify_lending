package upgrade

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

const (
	V013UpgradeNameHotfix2 = "v013_upgrade_testnet-hotfix2"
)

func CreateUpgradeHandlerForV013Hotfix2Upgrade(
	mm *module.Manager,
	configurator module.Configurator,
	acckeeper authkeeper.AccountKeeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for i := 0; i < 5; i++ {
			ctx.Logger().Info("we upgrade to v013")
		}

		acc := acckeeper.GetModuleAccount(ctx, "burn_auction")
		base := authtypes.NewBaseAccount(acc.GetAddress(), acc.GetPubKey(), acc.GetAccountNumber(), acc.GetSequence())
		newacc := authtypes.NewModuleAccount(base, "auction_burner", authtypes.Burner)
		acckeeper.SetModuleAccount(ctx, newacc)

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
