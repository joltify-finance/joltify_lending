package v1

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	spvmodulekeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	spvmoduletypes "github.com/joltify-finance/joltify_lending/x/spv/types"
)

const V018UpgradeName = "v018_upgrade"

func CreateUpgradeHandlerForV018Upgrade(
	mm *module.Manager,
	configurator module.Configurator,
	spvKeeper spvmodulekeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for i := 0; i < 5; i++ {
			ctx.Logger().Info("we upgrade to v018")
		}
		spvKeeper.IteratePool(ctx, func(poolInfo spvmoduletypes.PoolInfo) bool {
			for {
				if poolInfo.ProjectDueTime.After(ctx.BlockTime()) {
					break
				}
				poolInfo.ProjectDueTime = poolInfo.ProjectDueTime.Add(time.Second * time.Duration(poolInfo.ProjectLength))
			}
			spvKeeper.SetPool(ctx, poolInfo)
			return false
		})
		return mm.RunMigrations(ctx, configurator, vm)
	}
}
