package upgrade

import (
	"fmt"
	"strings"
	"time"

	quotamodulekeeper "github.com/joltify-finance/joltify_lending/x/quota/keeper"
	spvmoduletypes "github.com/joltify-finance/joltify_lending/x/spv/types"
	incentivemodulekeeper "github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"

	kycmoduletypes "github.com/joltify-finance/joltify_lending/x/kyc/types"

	spvmodulekeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"

	kycmodulekeeper "github.com/joltify-finance/joltify_lending/x/kyc/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

const (
	V011UpgradeName = "v011_upgrade_testnet_2"
	oneyear         = time.Hour * 24 * 365
)

func CreateUpgradeHandlerForV011Upgrade(
	mm *module.Manager,
	configurator module.Configurator,
	kycKeeper kycmodulekeeper.Keeper,
	spvKeeper spvmodulekeeper.Keeper,
	quotaKeeper quotamodulekeeper.Keeper,
	incentiveKeeper incentivemodulekeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for i := 0; i < 5; i++ {
			ctx.Logger().Info("we upgrade to v011")
		}

		// fixme double check the poolid
		myIncentive := spvmoduletypes.Incentive{
			Poolid: "0x70606714efcc24afe4736427c8a3df8168865daf01413008d7d98efcf03466b9",
			Spy:    "1.000000005262847188",
		}

		newparams := spvmoduletypes.Params{
			BurnThreshold: sdk.NewCoins(),
			Markets:       []spvmoduletypes.Moneymarket{},
			Incentives:    []spvmoduletypes.Incentive{myIncentive},
		}
		b, m := spvKeeper.GetParamsFromV22(ctx)
		spvKeeper.SetParams(ctx, newparams)
		spvKeeper.SetParamsFromV22(ctx, b, m)
		paget := spvKeeper.GetParams(ctx)

		fmt.Printf(">>>>>%v\n", paget.String())

		kycKeeper.IterateProject(ctx, func(projectInfo kycmoduletypes.ProjectInfo) bool {
			if strings.Contains(projectInfo.SPVName, "test projects") {
				fmt.Printf("now we delete project %s\n", projectInfo.SPVName)
				kycKeeper.DeleteProject(ctx, projectInfo.Index)
			}
			return false
		})

		// update the quota demon !!!!
		quotaParams := quotaKeeper.GetParams(ctx)

		fmt.Printf(">>>beofre <<<%v\n", quotaParams.String())
		globalTargets := quotaParams.Targets
		perAccountTargets := quotaParams.PerAccounttargets

		for i, el := range globalTargets {
			if el.ModuleName == "ibc" {
				var newcoin sdk.Coins
				for _, coin := range el.CoinsSum {
					demo := strings.ToLower(coin.Denom)
					n := sdk.NewCoin(demo, coin.Amount)
					newcoin = append(newcoin, n)
				}
				globalTargets[i].CoinsSum = newcoin
			}
		}
		quotaParams.Targets = globalTargets

		for i, el := range perAccountTargets {
			if el.ModuleName == "ibc" {
				var newcoin sdk.Coins
				for _, coin := range el.CoinsSum {
					demo := strings.ToLower(coin.Denom)
					n := sdk.NewCoin(demo, coin.Amount)
					newcoin = append(newcoin, n)
				}
				perAccountTargets[i].CoinsSum = newcoin
			}
		}

		quotaParams.PerAccounttargets = perAccountTargets
		quotaKeeper.SetParams(ctx, quotaParams)

		panew := quotaKeeper.GetParams(ctx)
		fmt.Printf(">>>>%v\n", panew.String())

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
