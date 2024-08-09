package upgrade

import (
	"context"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

const (
	V012UpgradeName = "v012_upgrade"
)

func CreateUpgradeHandlerForV012Upgrade(
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for i := 0; i < 5; i++ {
			sdk.UnwrapSDKContext(ctx).Logger().Info("we upgrade to v012")
		}

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
