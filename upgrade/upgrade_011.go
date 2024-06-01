package upgrade

import (
	"fmt"
	"strings"

	spvmoduletypes "github.com/joltify-finance/joltify_lending/x/spv/types"

	kycmoduletypes "github.com/joltify-finance/joltify_lending/x/kyc/types"

	spvmodulekeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"

	kycmodulekeeper "github.com/joltify-finance/joltify_lending/x/kyc/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
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

		burncoin := sdk.NewCoins(sdk.NewCoin("ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3", sdk.NewInt(10000000)))

		m := spvmoduletypes.Moneymarket{Denom: "ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3", ConversionFactor: 6}

		pa := spvmoduletypes.Params{
			BurnThreshold: burncoin,
			Markets:       []spvmoduletypes.Moneymarket{m},
		}
		spvKeeper.SetParams(ctx, pa)

		paget := spvKeeper.GetParams(ctx)

		fmt.Printf(">>>>>%v\n", paget.String())

		kycKeeper.IterateProject(ctx, func(projectInfo kycmoduletypes.ProjectInfo) bool {
			if strings.Contains(projectInfo.SPVName, "test projects") {
				fmt.Printf("now we delete project %s\n", projectInfo.SPVName)
				kycKeeper.DeleteProject(ctx, projectInfo.Index)
			}
			return false
		})
		return mm.RunMigrations(ctx, configurator, vm)
	}
}
