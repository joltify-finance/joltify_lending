package legacy

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

const V008UpgradeName = "v008_upgrade"

func CreateUpgradeHandlerForV008Upgrade(
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for i := 0; i < 5; i++ {
			sdk.UnwrapSDKContext(ctx).Logger().Info("we update the parameter for v008")
		}

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
