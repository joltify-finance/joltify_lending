package upgrade

import (
	"fmt"
	"strings"

	mintkeeper "github.com/joltify-finance/joltify_lending/x/mint/keeper"
	minttypes "github.com/joltify-finance/joltify_lending/x/mint/types"
	incentivetypes "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"

	spvmoduletypes "github.com/joltify-finance/joltify_lending/x/spv/types"

	spvmodulekeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	incentivekeeper "github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
)

const V011UpgradeNameTestnet = "v011_upgrade_testnet"

func CreateUpgradeHandlerForV011TestnetUpgrade(
	mm *module.Manager,
	configurator module.Configurator,
	spvKeeper spvmodulekeeper.Keeper,
	mintKeeper mintkeeper.Keeper,
	incentiveKeeper incentivekeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for i := 0; i < 5; i++ {
			ctx.Logger().Info("we upgrade to v011 testnet")
		}

		// set the mint new prameter
		mintKeeper.SetParams(ctx, minttypes.Params{
			FirstProvisions: minttypes.DefaultFirstProvision,
			Unit:            "minute",
			NodeSPY:         minttypes.DefaultNodeSPY,
		})

		mintKeeper.SetDistInfo(ctx, minttypes.HistoricalDistInfo{
			PayoutTime:     ctx.BlockTime(),
			TotalMintCoins: sdk.NewCoins(),
		})

		// fix the incorrect spv incentive rewards reported by commmunity

		incentiveKeeper.LegacyIterateSPVInvestorReward(ctx, func(key string, reward incentivetypes.SPVRewardAccTokens) bool {
			out := strings.TrimPrefix(key, incentivetypes.Incentiveclassprefix)
			data := strings.Split(out, "-")
			poolID := data[0]

			if reward.PaymentAmount == nil {
				return false
			}

			myreward := reward.PaymentAmount[0]
			amtAdj := myreward.Amount.Quo(sdk.NewInt(1e12))

			if amtAdj.IsZero() {
				incentiveKeeper.SetSPVInvestorReward(ctx, poolID, data[1], reward.PaymentAmount)
				return false
			}
			adjReward := sdk.NewCoins(sdk.NewCoin(myreward.Denom, amtAdj))
			fmt.Printf(">>>>>>%v----%v\n", data, adjReward.String())
			incentiveKeeper.SetSPVInvestorReward(ctx, poolID, data[1], adjReward)

			return false
		})

		amt := spvKeeper.GetParamsV21(ctx)
		burncoin := sdk.NewCoins(sdk.NewCoin("ausdc", amt))

		m := spvmoduletypes.Moneymarket{Denom: "ausdc", ConversionFactor: 18}
		pa := spvmoduletypes.Params{
			BurnThreshold: burncoin,
			Markets:       []spvmoduletypes.Moneymarket{m},
		}
		spvKeeper.SetParams(ctx, pa)

		paget := spvKeeper.GetParams(ctx)

		fmt.Printf(">>>>>%v\n", paget.String())
		return mm.RunMigrations(ctx, configurator, vm)
	}
}
