package legacy

import (
	"context"
	"strings"

	sdkmath "cosmossdk.io/math"
	minttypes "github.com/joltify-finance/joltify_lending/x/mint/types"

	"github.com/joltify-finance/joltify_lending/x/spv/types"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	mintkeeper "github.com/joltify-finance/joltify_lending/x/mint/keeper"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	incentivekeeper "github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
	incentivetypes "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
)

const V022UpgradeName = "v022_upgrade"

func CreateUpgradeHandlerForV022Upgrade(
	mm *module.Manager,
	configurator module.Configurator,
	spvKeeper spvkeeper.Keeper,
	mintKeeper mintkeeper.Keeper,
	incentiveKeeper incentivekeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for i := 0; i < 5; i++ {
			sdk.UnwrapSDKContext(ctx).Logger().Info("we upgrade to v022")
		}

		amt := spvKeeper.GetParamsV21(ctx)
		// set the new params
		// spvKeeper.SetParams(ctx, types.Params{
		//	BurnThreshold: sdk.NewCoins(threshold),
		// })
		spvKeeper.SetParams(ctx, types.Params{BurnThreshold: sdk.NewCoins(sdk.NewCoin("ausdc", amt))})
		//

		// set the mint new prameter
		mintKeeper.SetParams(ctx, minttypes.Params{
			FirstProvisions: minttypes.DefaultFirstProvision,
			Unit:            "minute",
			NodeSPY:         minttypes.DefaultNodeSPY,
		})

		mintKeeper.SetDistInfo(ctx, minttypes.HistoricalDistInfo{
			PayoutTime:     sdk.UnwrapSDKContext(ctx).BlockTime(),
			TotalMintCoins: sdk.NewCoins(),
		})

		// fix the incorrect spv incentive rewards reported by commmunity

		tpoolID := "0x4f1f7526042987d595fa135ed33a392a98bcc31f7ad79d6a5928e753ff7e8c8c"

		incentiveKeeper.LegacyIterateSPVInvestorReward(sdk.UnwrapSDKContext(ctx), func(key string, reward incentivetypes.SPVRewardAccTokens) bool {
			out := strings.TrimPrefix(key, incentivetypes.Incentiveclassprefix)
			data := strings.Split(out, "-")
			poolID := data[0]
			if poolID != tpoolID {
				return false
			}

			myreward := reward.PaymentAmount[0]
			amtAdj := myreward.Amount.Quo(sdkmath.NewInt(1e12))

			if amtAdj.IsZero() {
				incentiveKeeper.SetSPVInvestorReward(sdk.UnwrapSDKContext(ctx), poolID, data[1], reward.PaymentAmount)
				return false
			}
			adjReward := sdk.NewCoins(sdk.NewCoin(myreward.Denom, amtAdj))
			incentiveKeeper.SetSPVInvestorReward(sdk.UnwrapSDKContext(ctx), poolID, data[1], adjReward)

			return false
		})

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
