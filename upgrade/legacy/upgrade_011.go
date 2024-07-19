package legacy

import (
	"strings"

	kycmoduletypes "github.com/joltify-finance/joltify_lending/x/kyc/types"

	spvmodulekeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"

	kycmodulekeeper "github.com/joltify-finance/joltify_lending/x/kyc/keeper"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

const V011UpgradeName = "v011_upgrade"

func CreateUpgradeHandlerForV011Upgrade(
	mm *module.Manager,
	configurator module.Configurator,
	kycKeeper kycmodulekeeper.Keeper,
	spvKeeper spvmodulekeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for i := 0; i < 5; i++ {
			ctx.Logger().Info("we upgrade to v011")
		}

		kycKeeper.IterateProject(ctx, func(projectInfo kycmoduletypes.ProjectInfo) bool {
			if strings.Contains(projectInfo.SPVName, "test projects") {
				if projectInfo.Index == 0 {
					return false
				}
			}
			kycKeeper.DeleteProject(ctx, projectInfo.Index)
			return false
		})
		return mm.RunMigrations(ctx, configurator, vm)
	}
}
