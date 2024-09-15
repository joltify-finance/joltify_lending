package keeper_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/joltify-finance/joltify_lending/utils"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/require"
)

func TestMsgRepayInterest(t *testing.T) {
	config := sdk.NewConfig()
	utils.SetBech32AddressPrefixes(config)
	lapp, k, _, _, _, wctx := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(wctx)

	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 1, PoolName: "hello", Apy: []string{"7.8", "7.2"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdkmath.NewInt(322)), sdk.NewCoin("ausdc", sdkmath.NewInt(322))}}
	resp, err := lapp.CreatePool(ctx, &req)
	require.NoError(t, err)

	poolIndex := resp.PoolIndex
	reqRePayInterest := types.MsgRepayInterest{
		Creator:   "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0",
		PoolIndex: poolIndex[0],
		Token:     sdk.NewCoin("invalid", sdkmath.NewInt(100)),
	}
	_, err = lapp.RepayInterest(ctx, &reqRePayInterest)
	require.ErrorContains(t, err, "pool denom ausdc and repay is invalid: inconsistency tokens")

	reqRePayInterest.Creator = "invalid address"
	_, err = lapp.RepayInterest(ctx, &reqRePayInterest)
	require.ErrorContains(t, err, "invalid address")

	reqRePayInterest = types.MsgRepayInterest{
		Creator:   "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0",
		PoolIndex: poolIndex[0],
		Token:     sdk.NewCoin("ausdc", sdkmath.NewInt(100)),
	}
	_, err = lapp.RepayInterest(ctx, &reqRePayInterest)
	require.ErrorContains(t, err, "borrow amount is zero, no interest to be paid or interest paid is zero")

	poolInfo, found := k.GetPools(ctx, poolIndex[0])
	require.True(t, found)
	poolInfo.PoolStatus = types.PoolInfo_FROZEN
	poolInfo.BorrowedAmount = sdk.NewCoin("ausdc", sdkmath.NewInt(100))
	k.SetPool(ctx, poolInfo)
	_, err = lapp.RepayInterest(ctx, &reqRePayInterest)
	require.ErrorContains(t, err, "pool is not active")

	poolInfo.PoolStatus = types.PoolInfo_INACTIVE
	poolInfo.BorrowedAmount = sdk.NewCoin("ausdc", sdkmath.NewInt(100))
	k.SetPool(ctx, poolInfo)
	_, err = lapp.RepayInterest(ctx, &reqRePayInterest)
	require.ErrorContains(t, err, "pool is not active")
}

func mockBorrow(rctx context.Context, nftKeeper types.NFTKeeper, poolInfo *types.PoolInfo, borrowAmount sdk.Coin) {
	ctx := sdk.UnwrapSDKContext(rctx)
	classID := fmt.Sprintf("class-%v", poolInfo.Index[2:])
	poolClass, found := nftKeeper.GetClass(ctx, classID)
	if !found {
		panic("pool class must have already been set")
	}

	latestSeries := len(poolInfo.PoolNFTIds)

	currentBorrowClass := poolClass
	currentBorrowClass.Id = fmt.Sprintf("%v-%v", currentBorrowClass.Id, latestSeries)

	i, err := spvkeeper.CalculateInterestAmount(poolInfo.Apy, int(poolInfo.PayFreq))
	if err != nil {
		panic(err)
	}
	localAmount := convertBorrowToLocal(borrowAmount.Amount)

	rate := spvkeeper.CalculateInterestRate(poolInfo.Apy, int(poolInfo.PayFreq))
	paymentTime := ctx.BlockTime()
	firstPayment := types.PaymentItem{PaymentTime: paymentTime, PaymentAmount: sdk.NewCoin(borrowAmount.Denom, sdkmath.NewInt(0))}
	borrowDetails := make([]types.BorrowDetail, 1, 10)
	borrowDetails[0] = types.BorrowDetail{BorrowedAmount: sdk.NewCoin("aud-ausdc", localAmount), TimeStamp: ctx.BlockTime()}
	bi := types.BorrowInterest{
		PoolIndex:     poolInfo.Index,
		Apy:           poolInfo.Apy,
		PayFreq:       poolInfo.PayFreq,
		IssueTime:     ctx.BlockTime(),
		BorrowDetails: borrowDetails,
		MonthlyRatio:  i,
		InterestSPY:   rate,
		Payments:      []*types.PaymentItem{&firstPayment},
		AccInterest:   sdk.NewCoin("ausdc", sdkmath.ZeroInt()),
	}

	data, err := types2.NewAnyWithValue(&bi)
	if err != nil {
		panic(err)
	}
	currentBorrowClass.Data = data
	err = nftKeeper.SaveClass(ctx, currentBorrowClass)
	if err != nil {
		panic(err)
	}
	poolInfo.PoolNFTIds = append(poolInfo.PoolNFTIds, currentBorrowClass.Id)
	poolInfo.BorrowedAmount = poolInfo.BorrowedAmount.AddAmount(localAmount)
}

func testWithDifferentPrePayAmount(t *testing.T, ctx context.Context, app types.MsgServer, poolIndex string, k *spvkeeper.Keeper, expectedInterest sdkmath.Int) {
	poolnfoBeforeTest, found := k.GetPools(ctx, poolIndex)
	require.True(t, found)

	repayMsg := types.MsgRepayInterest{
		Creator:   "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0",
		PoolIndex: poolIndex,
		Token:     sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(2e8)),
	}
	_, err := app.RepayInterest(ctx, &repayMsg)
	require.NoError(t, err)
	poolnfo, found := k.GetPools(ctx, poolIndex)
	require.True(t, found)

	counter := sdkmath.NewIntFromUint64(2e8).Quo(expectedInterest)
	require.EqualValues(t, counter.Int64(), int64(poolnfo.InterestPrepayment.Counter))

	k.SetPool(ctx, poolnfoBeforeTest)
	// now we pay very small amount
	repayMsg.Token = sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1))

	_, err = app.RepayInterest(ctx, &repayMsg)
	require.ErrorContains(t, err, "you must pay at least one interest cycle")
	poolnfo, found = k.GetPools(ctx, poolIndex)
	require.True(t, found)

	require.Nil(t, poolnfo.InterestPrepayment)
	k.SetPool(ctx, poolnfoBeforeTest)
}

func TestGetAllInterestWithInterestPaid(t *testing.T) {
	config := sdk.NewConfig()
	utils.SetBech32AddressPrefixes(config)
	lapp, k, nftKeeper, _, _, wctx := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(wctx)

	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 3, PoolName: "hello", Apy: []string{"0.078", "0.072"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdkmath.NewInt(322)), sdk.NewCoin("ausdc", sdkmath.NewInt(322))}}
	resp, err := lapp.CreatePool(ctx, &req)
	require.NoError(t, err)

	poolIndex := resp.PoolIndex[0]
	samplePool, found := k.GetPools(ctx, poolIndex)
	require.True(t, found)

	samplePool.EscrowInterestAmount = sdkmath.NewIntFromUint64(10e12)
	k.SetPool(ctx, samplePool)

	samplePool.UsableAmount = sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(8e12))
	samplePool.PoolStatus = types.PoolInfo_ACTIVE
	samplePool.PoolTotalBorrowLimit = 100
	samplePool.GraceTime = time.Hour * 1000
	samplePool.TargetAmount = sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(100e12))
	firstBorrowTime := ctx.BlockTime()
	mockBorrow(ctx, nftKeeper, &samplePool, sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(2e8)))
	samplePool.CurrentPoolTotalBorrowCounter = 1
	k.SetPool(ctx, samplePool)
	err = k.HandleInterest(ctx, &samplePool)
	require.ErrorContains(t, err, "pay interest too early")
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Second * time.Duration(spvkeeper.OneMonth-1)))
	err = k.HandleInterest(ctx, &samplePool)
	require.ErrorContains(t, err, "pay interest too early")

	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Second * 2))
	err = k.HandleInterest(ctx, &samplePool)
	require.NoError(t, err)
	poolInfo, _ := k.GetPools(ctx, poolIndex)

	b1 := poolInfo.PoolNFTIds[0]

	nclass, _ := nftKeeper.GetClass(ctx, b1)
	var borrowInterest types.BorrowInterest
	err = proto.Unmarshal(nclass.Data.Value, &borrowInterest)
	if err != nil {
		panic(err)
	}

	period := spvkeeper.OneYear / spvkeeper.OneMonth
	interestOneMonthWithReserve := sdkmath.LegacyNewDecFromInt(sdkmath.NewIntFromUint64(2e8)).Mul(samplePool.Apy).QuoInt64(int64(period)).TruncateInt()

	interestOneMonth := interestOneMonthWithReserve.Sub(sdkmath.LegacyNewDecFromInt(interestOneMonthWithReserve).Mul(sdkmath.LegacyMustNewDecFromStr("0.15")).TruncateInt())

	testWithDifferentPrePayAmount(t, ctx, lapp, poolIndex, k, interestOneMonthWithReserve)

	repayMsg := types.MsgRepayInterest{
		Creator:   "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0",
		PoolIndex: poolIndex,
		Token:     sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(2e8)),
	}
	_, err = lapp.RepayInterest(ctx, &repayMsg)

	paymentTime := borrowInterest.Payments[1].PaymentTime
	require.EqualValues(t, firstBorrowTime.Add(time.Second*spvkeeper.OneMonth), paymentTime)
	require.True(t, checkValueWithRangeTwo(interestOneMonth, borrowInterest.Payments[1].PaymentAmount.Amount))

	// at the middle of the month, we borrow
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Second * 15 * 24 * 3600))
	secondBorrowTime := ctx.BlockTime()

	mockBorrow(ctx, nftKeeper, &samplePool, sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(2e8)))

	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Second * 15 * 24 * 3600))
	err = k.HandleInterest(ctx, &samplePool)
	require.NoError(t, err)

	poolInfo, _ = k.GetPools(ctx, poolIndex)

	b1 = poolInfo.PoolNFTIds[0]
	nclass, _ = nftKeeper.GetClass(ctx, b1)
	err = proto.Unmarshal(nclass.Data.Value, &borrowInterest)
	if err != nil {
		panic(err)
	}
	paymentTime = borrowInterest.Payments[2].PaymentTime
	require.EqualValues(t, firstBorrowTime.Add(time.Second*spvkeeper.OneMonth*2), paymentTime)
	require.True(t, checkValueWithRangeTwo(interestOneMonth, borrowInterest.Payments[1].PaymentAmount.Amount))

	b2 := poolInfo.PoolNFTIds[1]
	nclass, _ = nftKeeper.GetClass(ctx, b2)
	err = proto.Unmarshal(nclass.Data.Value, &borrowInterest)
	if err != nil {
		panic(err)
	}

	// we check the interest of the second borrow from mid of the month
	delta := firstBorrowTime.Add(time.Second * spvkeeper.OneMonth * 2).Sub(secondBorrowTime)

	r := spvkeeper.CalculateInterestRate(poolInfo.Apy, int(poolInfo.PayFreq))
	interest := r.Power(uint64(delta.Seconds())).Sub(sdkmath.LegacyOneDec())

	paymentAmount := interest.MulInt(sdkmath.NewIntFromUint64(2e8)).TruncateInt()
	reservedAmount := sdkmath.LegacyNewDecFromInt(paymentAmount).Mul(sdkmath.LegacyMustNewDecFromStr("0.15")).TruncateInt()
	toInvestors := paymentAmount.Sub(reservedAmount)

	err = proto.Unmarshal(nclass.Data.Value, &borrowInterest)
	if err != nil {
		panic(err)
	}

	paymentTime = borrowInterest.Payments[1].PaymentTime
	require.EqualValues(t, firstBorrowTime.Add(time.Second*spvkeeper.OneMonth*2), paymentTime)
	require.True(t, checkValueWithRangeTwo(toInvestors, borrowInterest.Payments[1].PaymentAmount.Amount))

	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Second * 30 * 24 * 3600))
	err = k.HandleInterest(ctx, &samplePool)
	require.NoError(t, err)

	nclass1, _ := nftKeeper.GetClass(ctx, b1)
	err = proto.Unmarshal(nclass1.Data.Value, &borrowInterest)
	require.NoError(t, err)

	paymentTime = borrowInterest.Payments[3].PaymentTime
	require.EqualValues(t, firstBorrowTime.Add(time.Second*spvkeeper.OneMonth*3), paymentTime)
	require.True(t, checkValueWithRangeTwo(interestOneMonth, borrowInterest.Payments[3].PaymentAmount.Amount))

	nclass2, _ := nftKeeper.GetClass(ctx, b2)
	err = proto.Unmarshal(nclass2.Data.Value, &borrowInterest)
	require.NoError(t, err)

	paymentTime = borrowInterest.Payments[2].PaymentTime
	require.EqualValues(t, firstBorrowTime.Add(time.Second*spvkeeper.OneMonth*3), paymentTime)
	require.True(t, checkValueWithRangeTwo(interestOneMonth, borrowInterest.Payments[2].PaymentAmount.Amount))

	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Second * time.Duration(24*3600*13)))
	thirdBorrowTime := ctx.BlockTime()
	mockBorrow(ctx, nftKeeper, &samplePool, sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(2e8)))

	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Second * time.Duration(24*3600*19)))
	err = k.HandleInterest(ctx, &samplePool)
	require.NoError(t, err)

	delta = firstBorrowTime.Add(time.Second * spvkeeper.OneMonth * 4).Sub(thirdBorrowTime)

	interest = r.Power(uint64(delta.Seconds())).Sub(sdkmath.LegacyOneDec())
	paymentAmount = interest.MulInt(sdkmath.NewIntFromUint64(2e8)).TruncateInt()
	reservedAmount = sdkmath.LegacyNewDecFromInt(paymentAmount).Mul(sdkmath.LegacyMustNewDecFromStr("0.15")).TruncateInt()
	toInvestors = paymentAmount.Sub(reservedAmount)

	poolInfo, _ = k.GetPools(ctx, poolIndex)
	b3 := poolInfo.PoolNFTIds[2]

	nclass3, _ := nftKeeper.GetClass(ctx, b3)
	err = proto.Unmarshal(nclass3.Data.Value, &borrowInterest)
	if err != nil {
		panic(err)
	}

	paymentTime = borrowInterest.Payments[1].PaymentTime
	require.EqualValues(t, firstBorrowTime.Add(time.Second*spvkeeper.OneMonth*4), paymentTime)
	require.True(t, checkValueWithRangeTwo(toInvestors, borrowInterest.Payments[1].PaymentAmount.Amount))

	nclass2, _ = nftKeeper.GetClass(ctx, b2)
	err = proto.Unmarshal(nclass2.Data.Value, &borrowInterest)
	if err != nil {
		panic(err)
	}
	paymentTime = borrowInterest.Payments[3].PaymentTime
	require.EqualValues(t, firstBorrowTime.Add(time.Second*spvkeeper.OneMonth*4), paymentTime)
	require.True(t, checkValueWithRangeTwo(interestOneMonth, borrowInterest.Payments[3].PaymentAmount.Amount))

	nclass1, _ = nftKeeper.GetClass(ctx, b1)
	err = proto.Unmarshal(nclass1.Data.Value, &borrowInterest)
	if err != nil {
		panic(err)
	}
	paymentTime = borrowInterest.Payments[4].PaymentTime
	require.EqualValues(t, firstBorrowTime.Add(time.Second*spvkeeper.OneMonth*4), paymentTime)
	require.True(t, checkValueWithRangeTwo(interestOneMonth, borrowInterest.Payments[4].PaymentAmount.Amount))
}
