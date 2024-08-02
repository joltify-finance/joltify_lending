package upgrade

import (
	"context"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

const (
	V014UpgradeName = "v014_upgrade"
)

func CreateUpgradeHandlerForV014Upgrade(
	mm *module.Manager,
	configurator module.Configurator,
	keeper *ibckeeper.Keeper,
	subspace paramtypes.Subspace,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for i := 0; i < 5; i++ {
			sdk.UnwrapSDKContext(ctx).Logger().Info("we upgrade to v014")
		}

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
