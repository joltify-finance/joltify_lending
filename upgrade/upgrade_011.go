package upgrade

import (
	"fmt"
	"strings"
	"time"

	incentivemodulekeeper "github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
	incentivetypes "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"

	quotamodulekeeper "github.com/joltify-finance/joltify_lending/x/quota/keeper"
	spvmoduletypes "github.com/joltify-finance/joltify_lending/x/spv/types"

	kycmoduletypes "github.com/joltify-finance/joltify_lending/x/kyc/types"

	spvmodulekeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"

	kycmodulekeeper "github.com/joltify-finance/joltify_lending/x/kyc/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

const (
	V011UpgradeName = "v011_upgrade"
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

		burncoin := sdk.NewCoins(sdk.NewCoin("ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3", sdk.NewInt(10000000)))

		m := spvmoduletypes.Moneymarket{Denom: "ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3", ConversionFactor: 6}

		// fixme double check the poolid
		myIncentive := spvmoduletypes.Incentive{
			Poolid: "0x3a0e72aefc820a7ec5a04cd3b987df8794d5adc48df082a5f8c2aba80a5f6e20",
			Spy:    "1.000000005262847188",
		}
		pa := spvmoduletypes.Params{
			BurnThreshold: burncoin,
			Markets:       []spvmoduletypes.Moneymarket{m},
			Incentives:    []spvmoduletypes.Incentive{myIncentive},
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

		// update the incentive module parameter
		currentTime := ctx.BlockTime()
		incentiveParams := incentiveKeeper.GetParams(ctx)
		addedIncentive := incentivetypes.NewMultiRewardPeriod(true, "0x3a0e72aefc820a7ec5a04cd3b987df8794d5adc48df082a5f8c2aba80a5f6e20", currentTime.Add(-1*24*time.Hour), time.Now().Add(oneyear), sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(0))))

		incentiveParams.SPVRewardPeriods = append(incentiveParams.SPVRewardPeriods, addedIncentive)
		incentiveKeeper.SetParams(ctx, incentiveParams)

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
