package v1

import (
	"fmt"
	"time"

	"github.com/joltify-finance/joltify_lending/x/spv/types"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/cosmos/gogoproto/proto"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	spvmoduletypes "github.com/joltify-finance/joltify_lending/x/spv/types"

	incentivekeeper "github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
)

const V021UpgradeName = "v021_upgrade"

func CreateUpgradeHandlerForV021Upgrade(
	mm *module.Manager,
	configurator module.Configurator,
	incentiveKeeper incentivekeeper.Keeper,
	spvKeeper spvkeeper.Keeper,
	nftKeeper types.NFTKeeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for i := 0; i < 5; i++ {
			ctx.Logger().Info("we upgrade to v021")
		}

		incentiveParams := incentiveKeeper.GetParams(ctx)

		for i, el := range incentiveParams.SPVRewardPeriods {
			if el.CollateralType == "0x70606714efcc24afe4736427c8a3df8168865daf01413008d7d98efcf03466b9" {
				el.RewardsPerSecond = sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(34722)))
				incentiveParams.SPVRewardPeriods[i] = el
			}
		}

		incentiveKeeper.SetParams(ctx, incentiveParams)

		newincentive := incentiveKeeper.GetParams(ctx)
		for _, el := range newincentive.SPVRewardPeriods {
			fmt.Printf(">>>%v\n", el)
		}

		// now we correct the spv reward

		req := &types.QueryQueryPoolRequest{
			PoolIndex: "0x70606714efcc24afe4736427c8a3df8168865daf01413008d7d98efcf03466b9",
		}
		ret, err := spvKeeper.QueryPool(ctx, req)
		if err != nil {
			panic("should never find pool" + err.Error())
		}

		nfts := ret.PoolInfo.PoolNFTIds
		for _, eachnftID := range nfts {
			a, ok := nftKeeper.GetClass(ctx, eachnftID)
			if !ok {
				panic("should never find nft")
			}

			var borrowInterest types.BorrowInterest
			var err error
			err = proto.Unmarshal(a.Data.Value, &borrowInterest)
			if err != nil {
				panic("we failed at unmarshal" + err.Error())
			}

			incentivepaments := borrowInterest.IncentivePayments
			for index, eachincentive := range incentivepaments {

				if eachincentive.PaymentAmount.IsZero() {
					continue
				}

				onlyIncentiveCoin := eachincentive.PaymentAmount[0]

				amt := onlyIncentiveCoin.Amount.Quo(sdk.NewInt(1e12))
				incentivepaments[index].PaymentAmount = sdk.NewCoins(sdk.NewCoin("ujolt", amt))
			}

			bz, err := proto.Marshal(&borrowInterest)
			if err != nil {
				panic("we failed here" + err.Error())
			}
			a.Data.Value = bz
			err = nftKeeper.UpdateClass(ctx, a)
			if err != nil {
				panic("fail to update final payment" + err.Error())
			}
		}

		// fix the incorrect reward store
		rewards, ok := incentiveKeeper.GetSPVReward(ctx, "0x70606714efcc24afe4736427c8a3df8168865daf01413008d7d98efcf03466b9")
		if !ok {
			panic("fail to find the reward")
		}

		var newrewards sdk.Coins
		for _, el := range rewards.PaymentAmount {
			amt := el.Amount.Quo(sdk.NewInt(1e12))
			c := sdk.NewCoin("ujolt", amt)
			newrewards.Add(c)
		}

		rt := types2.SPVRewardAccTokens{
			PaymentAmount: newrewards,
		}
		incentiveKeeper.SetSPVReward(ctx, "0x70606714efcc24afe4736427c8a3df8168865daf01413008d7d98efcf03466b9", rt)

		// we now update the project due time
		spvKeeper.IteratePool(ctx, func(poolInfo spvmoduletypes.PoolInfo) bool {
			if poolInfo.PoolStatus == types.PoolInfo_ACTIVE || poolInfo.PoolStatus == types.PoolInfo_PooLPayPartially {
				previousDueTime := poolInfo.ProjectDueTime
				if poolInfo.ProjectDueTime.Before(poolInfo.LastPaymentTime) {
					fmt.Printf("we have to correct the due time")
					previousDueTime = poolInfo.LastPaymentTime.Add(time.Duration(poolInfo.PayFreq) * time.Second)
				}

				poolInfo.ProjectDueTime = previousDueTime.Truncate(time.Duration(poolInfo.PayFreq) * time.Second)
				spvKeeper.SetPool(ctx, poolInfo)
			}
			return false
		})

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
