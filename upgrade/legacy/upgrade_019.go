package legacy

import (
	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	kycmodulekeeper "github.com/joltify-finance/joltify_lending/x/kyc/keeper"
	kycmoduletypes "github.com/joltify-finance/joltify_lending/x/kyc/types"
	spvmodulekeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	spvmoduletypes "github.com/joltify-finance/joltify_lending/x/spv/types"
)

const V019UpgradeName = "v019_upgrade"

func CreateUpgradeHandlerForV019Upgrade(
	mm *module.Manager,
	configurator module.Configurator,
	spvKeeper spvmodulekeeper.Keeper,
	kycKeeper kycmodulekeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for i := 0; i < 5; i++ {
			ctx.Logger().Info("we upgrade to v019")
		}
		defaultAmount, _ := sdk.NewIntFromString("200000000000000000000")
		spvKeeper.IteratePool(ctx, func(poolInfo spvmoduletypes.PoolInfo) bool {
			poolInfo.MinDepositAmount = defaultAmount
			spvKeeper.SetPool(ctx, poolInfo)
			return false
		})

		kycKeeper.IterateProject(ctx, func(projectInfo kycmoduletypes.ProjectInfo) bool {
			projectInfo.MinDepositAmount = defaultAmount
			kycKeeper.UpdateProject(ctx, &projectInfo)
			return false
		})

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
