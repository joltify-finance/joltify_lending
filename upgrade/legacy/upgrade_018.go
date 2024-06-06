package legacy

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
			if poolInfo.Index == "0xaac7b8bd2bf82a8cc4d7f3647f3ec067ca9cdd9a854d493cc983fdc1cf91ab21" || poolInfo.Index == "0xf5de65c0804ddfd4988996d6c80e228dab89d86ada184830178f94020f80247d" {
				poolInfo.ProjectLength = 86400 * 3
			}
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
