package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

const V013UpgradeName = "v013_upgrade"

func CreateUpgradeHandlerForV013Upgrade(
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for i := 0; i < 5; i++ {
			ctx.Logger().Info("we upgrade to v013")
		}
		return mm.RunMigrations(ctx, configurator, vm)
	}
}
