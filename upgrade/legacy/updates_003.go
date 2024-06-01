package legacy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	vaultmodulekeeper "github.com/joltify-finance/joltify_lending/x/vault/keeper"
)

const V003UpgradeName = "v003_upgrade"

func CreateUpgradeHandlerForV003Upgrade(
	mm *module.Manager,
	k *vaultmodulekeeper.Keeper,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for i := 0; i < 2; i++ {
			ctx.Logger().Info("we update the parameter for v003")
		}
		fees := k.LegacyGetAllFeeAMountAndDelete(ctx)
		ctx.Logger().Info("we relocate fees:", "fee:", fees.String())
		k.SetStoreFeeAmount(ctx, sdk.NewCoins(fees...))
		afterStorage := k.GetAllFeeAmount(ctx)
		ctx.Logger().Info("we retrieve fees from new storage:", "fee:", afterStorage.String())
		return mm.RunMigrations(ctx, configurator, vm)
	}
}
