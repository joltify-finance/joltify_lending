package legacy

import (
	"context"
	"time"

	sdkmath "cosmossdk.io/math"

	nftmodulekeeper "cosmossdk.io/x/nft/keeper"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
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
	return func(ctx context.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for i := 0; i < 5; i++ {
			sdk.UnwrapSDKContext(ctx).Logger().Info("we upgrade to v020")
		}
		defaultAmount, _ := sdkmath.NewIntFromString("200000000000000000000")
		spvKeeper.IteratePool(ctx, func(poolInfo spvmoduletypes.PoolInfo) bool {
			poolInfo.MinDepositAmount = defaultAmount
			spvKeeper.SetPool(ctx, poolInfo)
			updatenftmodule(sdk.UnwrapSDKContext(ctx), nftKeeper, poolInfo.PoolNFTIds)
			return false
		})

		paOld := incentiveKeeper.GetParamsV19(sdk.UnwrapSDKContext(ctx))

		// we give 3000 jolt per day to the pool with 3466b9
		rewards := incentivetypes.MultiRewardPeriods{
			incentivetypes.NewMultiRewardPeriod(true, "0x70606714efcc24afe4736427c8a3df8168865daf01413008d7d98efcf03466b9", sdk.UnwrapSDKContext(ctx).BlockTime().Add(-time.Hour), sdk.UnwrapSDKContext(ctx).BlockTime().Add(time.Hour*24*365), sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(34722222222222222)))),
		}

		newParamns := incentivetypes.Params{
			ClaimEnd:                paOld.ClaimEnd,
			ClaimMultipliers:        paOld.ClaimMultipliers,
			JoltSupplyRewardPeriods: paOld.JoltSupplyRewardPeriods,
			JoltBorrowRewardPeriods: paOld.JoltBorrowRewardPeriods,
			SwapRewardPeriods:       paOld.SwapRewardPeriods,
			SPVRewardPeriods:        rewards,
		}

		incentiveKeeper.SetParams(sdk.UnwrapSDKContext(ctx), newParamns)

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
