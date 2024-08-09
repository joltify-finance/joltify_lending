package upgrade

import (
	"context"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

const (
	V014UpgradeName       = "v014_testnet_upgrade"
	V014HotFixUpgradeName = "v014_testnet_hotfix_upgrade"
)

func CreateUpgradeHandlerForV014Upgrade(
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for i := 0; i < 5; i++ {
			sdk.UnwrapSDKContext(ctx).Logger().Info("we upgrade to v014")
		}

		return mm.RunMigrations(ctx, configurator, vm)
	}
}

func CreateUpgradeHandlerForV014hotFixUpgrade(
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for i := 0; i < 5; i++ {
			sdk.UnwrapSDKContext(ctx).Logger().Info("we upgrade to v014 hotfix")
		}

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
