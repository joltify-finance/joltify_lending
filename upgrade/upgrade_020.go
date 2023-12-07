package v1

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	nftmodulekeeper "github.com/cosmos/cosmos-sdk/x/nft/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/cosmos/gogoproto/proto"
	spvmodulekeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	spvmoduletypes "github.com/joltify-finance/joltify_lending/x/spv/types"
	incentivemodulekeeper "github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
	incentivetypes "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
)

const V020UpgradeName = "v020_upgrade"

func updatenftmodule(ctx sdk.Context, nftKeeper nftmodulekeeper.Keeper, classes []string) {
	for _, classID := range classes {
		classItem, found := nftKeeper.GetClass(ctx, classID)
		if !found {
			panic("should never fail to find the nft")
		}

		var borrowInterest spvmoduletypes.BorrowInterest
		var err error
		err = proto.Unmarshal(classItem.Data.Value, &borrowInterest)
		if err != nil {
			panic("we failed at unmarshal" + err.Error())
		}

		var allincentivePayment []*spvmoduletypes.IncentivePaymentItem
		for _, payment := range borrowInterest.Payments {
			item := spvmoduletypes.IncentivePaymentItem{
				PaymentAmount:  sdk.NewCoins(),
				PaymentTime:    payment.PaymentTime,
				BorrowedAmount: payment.BorrowedAmount,
			}
			allincentivePayment = append(allincentivePayment, &item)
		}
		borrowInterest.IncentivePayments = allincentivePayment
		bz, err := proto.Marshal(&borrowInterest)
		if err != nil {
			panic("we failed here" + err.Error())
		}
		classItem.Data.Value = bz
		err = nftKeeper.UpdateClass(ctx, classItem)
		if err != nil {
			panic(err)
		}

	}
}

func CreateUpgradeHandlerForV020Upgrade(
	mm *module.Manager,
	configurator module.Configurator,
	spvKeeper spvmodulekeeper.Keeper,
	nftKeeper nftmodulekeeper.Keeper,
	incentiveKeeper incentivemodulekeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for i := 0; i < 5; i++ {
			ctx.Logger().Info("we upgrade to v020")
		}
		defaultAmount, _ := sdk.NewIntFromString("200000000000000000000")
		spvKeeper.IteratePool(ctx, func(poolInfo spvmoduletypes.PoolInfo) bool {
			poolInfo.MinDepositAmount = defaultAmount
			spvKeeper.SetPool(ctx, poolInfo)
			updatenftmodule(ctx, nftKeeper, poolInfo.PoolNFTIds)
			return false
		})

		paOld := incentiveKeeper.GetParamsV19(ctx)

		// we give 3000 jolt per day to the pool with 3466b9
		rewards := incentivetypes.MultiRewardPeriods{
			incentivetypes.NewMultiRewardPeriod(true, "0x70606714efcc24afe4736427c8a3df8168865daf01413008d7d98efcf03466b9", ctx.BlockTime().Add(-time.Hour), ctx.BlockTime().Add(time.Hour*24*365), sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(34722222222222222)))),
		}

		newParamns := incentivetypes.Params{
			ClaimEnd:                paOld.ClaimEnd,
			ClaimMultipliers:        paOld.ClaimMultipliers,
			JoltSupplyRewardPeriods: paOld.JoltSupplyRewardPeriods,
			JoltBorrowRewardPeriods: paOld.JoltBorrowRewardPeriods,
			SwapRewardPeriods:       paOld.SwapRewardPeriods,
			SPVRewardPeriods:        rewards,
		}

		incentiveKeeper.SetParams(ctx, newParamns)

		return mm.RunMigrations(ctx, configurator, vm)
	}
}