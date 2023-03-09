package keeper_test

import (
	"fmt"
	"testing"
	"time"

	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/require"
)

const oneMonth = 24 * 30 * 3600

func TestMsgRepayInterest(t *testing.T) {

	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)
	app, k, _, wctx := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(wctx)

	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 1, PoolName: "hello", Apy: "7.8", TargetTokenAmount: sdk.NewCoin("ausdc", sdk.NewInt(322))}
	resp, err := app.CreatePool(ctx, &req)
	require.NoError(t, err)

	poolIndex := resp.PoolIndex
	reqRePayInterest := types.MsgRepayInterest{
		Creator:   "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0",
		PoolIndex: poolIndex[0],
		Token:     sdk.NewCoin("invalid", sdk.NewInt(100)),
	}
	_, err = app.RepayInterest(ctx, &reqRePayInterest)
	require.ErrorContains(t, err, "pool denom ausdc and repay is invalid: inconsistency tokens")

	reqRePayInterest.Creator = "invalid address"
	_, err = app.RepayInterest(ctx, &reqRePayInterest)
	require.ErrorContains(t, err, "invalid address")

	reqRePayInterest = types.MsgRepayInterest{
		Creator:   "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0",
		PoolIndex: poolIndex[0],
		Token:     sdk.NewCoin("ausdc", sdk.NewInt(100)),
	}
	_, err = app.RepayInterest(ctx, &reqRePayInterest)
	require.NoError(t, err)

	poolInfo, found := k.GetPools(ctx, poolIndex[0])
	require.True(t, found)
	poolInfo.PoolStatus = types.PoolInfo_CLOSED
	_, err = app.RepayInterest(ctx, &reqRePayInterest)
	require.NoError(t, err, "pool is not active")

	poolInfo.PoolStatus = types.PoolInfo_INACTIVE
	_, err = app.RepayInterest(ctx, &reqRePayInterest)
	require.NoError(t, err, "pool is not active")

}

func mockBorrow(ctx sdk.Context, nftKeeper types.NFTKeeper, poolInfo *types.PoolInfo, borrowAmount sdk.Coin) {

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

	rate := spvkeeper.CalculateInterestRate(poolInfo.Apy, int(poolInfo.PayFreq))
	paymentTime := ctx.BlockTime()
	firstPayment := types.PaymentItem{PaymentTime: paymentTime, PaymentAmount: sdk.NewCoin(borrowAmount.Denom, sdk.NewInt(0))}
	borrowDetails := make([]types.BorrowDetail, 1, 10)
	borrowDetails[0] = types.BorrowDetail{BorrowedAmount: borrowAmount, TimeStamp: ctx.BlockTime()}
	bi := types.BorrowInterest{
		PoolIndex:     poolInfo.Index,
		Apy:           poolInfo.Apy,
		PayFreq:       poolInfo.PayFreq,
		IssueTime:     ctx.BlockTime(),
		BorrowDetails: borrowDetails,
		MonthlyRatio:  i,
		InterestSPY:   rate,
		Payments:      []*types.PaymentItem{&firstPayment},
	}

	data, err := types2.NewAnyWithValue(&bi)
	if err != nil {
		panic(err)
	}
	currentBorrowClass.Data = data
	nftKeeper.SaveClass(ctx, currentBorrowClass)
	poolInfo.PoolNFTIds = append(poolInfo.PoolNFTIds, currentBorrowClass.Id)
}

func TestGetAllInterestToBePaid(t *testing.T) {

	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)
	app, k, nftKeeper, wctx := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(wctx)

	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 3, PoolName: "hello", Apy: "7.8", TargetTokenAmount: sdk.NewCoin("ausdc", sdk.NewInt(322))}
	resp, err := app.CreatePool(ctx, &req)
	require.NoError(t, err)

	poolIndex := resp.PoolIndex[0]
	samplePool, found := k.GetPools(ctx, poolIndex)
	require.True(t, found)

	samplePool.EscrowInterestAmount = sdk.NewCoin("ausdc", sdk.NewIntFromUint64(10e12))
	k.SetPool(ctx, samplePool)

	samplePool.BorrowableAmount = sdk.NewCoin("ausdc", sdk.NewIntFromUint64(8e12))
	samplePool.PoolStatus = types.PoolInfo_ACTIVE
	firstBorrowTime := ctx.BlockTime()
	mockBorrow(ctx, nftKeeper, &samplePool, sdk.NewCoin("ausdc", sdk.NewIntFromUint64(2e8)))
	k.SetPool(ctx, samplePool)
	err = k.HandleInterest(ctx, &samplePool)
	require.ErrorContains(t, err, "pay interest too early")
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Second * time.Duration(oneMonth-1)))
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

	interestOneYearWithReserve := sdk.NewDecFromInt(sdk.NewIntFromUint64(2e8)).Mul(samplePool.Apy).QuoInt64(12).TruncateInt()
	interestOneYear := interestOneYearWithReserve.Sub(sdk.NewDecFromInt(interestOneYearWithReserve).Mul(sdk.MustNewDecFromStr("0.15")).TruncateInt())

	paymentTime := borrowInterest.Payments[1].PaymentTime
	require.EqualValues(t, firstBorrowTime.Add(time.Second*oneMonth), paymentTime)
	require.EqualValues(t, borrowInterest.Payments[1].PaymentAmount.Amount.String(), interestOneYear.String())

	// at the middle of the month, we borrow
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Second * 15 * 24 * 3600))
	secondBorrowTime := ctx.BlockTime()
	mockBorrow(ctx, nftKeeper, &samplePool, sdk.NewCoin("ausdc", sdk.NewIntFromUint64(2e8)))

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
	require.EqualValues(t, firstBorrowTime.Add(time.Second*oneMonth*2), paymentTime)
	require.EqualValues(t, borrowInterest.Payments[1].PaymentAmount.Amount.String(), interestOneYear.String())

	b2 := poolInfo.PoolNFTIds[1]
	nclass, _ = nftKeeper.GetClass(ctx, b2)
	err = proto.Unmarshal(nclass.Data.Value, &borrowInterest)
	if err != nil {
		panic(err)
	}

	err = proto.Unmarshal(nclass.Data.Value, &borrowInterest)
	if err != nil {
		panic(err)
	}
	paymentTime = borrowInterest.Payments[1].PaymentTime

	delta := firstBorrowTime.Add(time.Second * month * 2).Sub(secondBorrowTime)

	r := spvkeeper.CalculateInterestRate(poolInfo.Apy, int(poolInfo.PayFreq))
	interest := r.Power(uint64(delta.Seconds())).Sub(sdk.OneDec())

	paymentAmount := interest.MulInt(sdk.NewIntFromUint64(2e8)).TruncateInt()
	reservedAmount := sdk.NewDecFromInt(paymentAmount).Mul(sdk.MustNewDecFromStr("0.15")).TruncateInt()
	toInvestors := paymentAmount.Sub(reservedAmount)

	err = proto.Unmarshal(nclass.Data.Value, &borrowInterest)
	if err != nil {
		panic(err)
	}

	paymentTime = borrowInterest.Payments[1].PaymentTime
	require.EqualValues(t, firstBorrowTime.Add(time.Second*oneMonth*2), paymentTime)
	require.EqualValues(t, borrowInterest.Payments[1].PaymentAmount.Amount.String(), toInvestors.String())

	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Second * 30 * 24 * 3600))
	err = k.HandleInterest(ctx, &samplePool)
	require.NoError(t, err)

	nclass1, _ := nftKeeper.GetClass(ctx, b1)
	err = proto.Unmarshal(nclass1.Data.Value, &borrowInterest)
	require.NoError(t, err)

	paymentTime = borrowInterest.Payments[3].PaymentTime
	require.EqualValues(t, firstBorrowTime.Add(time.Second*oneMonth*3), paymentTime)
	require.EqualValues(t, borrowInterest.Payments[3].PaymentAmount.Amount.String(), interestOneYear.String())

	nclass2, _ := nftKeeper.GetClass(ctx, b2)
	err = proto.Unmarshal(nclass2.Data.Value, &borrowInterest)
	require.NoError(t, err)

	paymentTime = borrowInterest.Payments[2].PaymentTime
	require.EqualValues(t, firstBorrowTime.Add(time.Second*oneMonth*3), paymentTime)
	require.EqualValues(t, borrowInterest.Payments[2].PaymentAmount.Amount.String(), interestOneYear.String())

	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Second * time.Duration(24*3600*13)))
	thirdBorrowTime := ctx.BlockTime()
	mockBorrow(ctx, nftKeeper, &samplePool, sdk.NewCoin("ausdc", sdk.NewIntFromUint64(2e8)))

	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Second * time.Duration(24*3600*19)))
	err = k.HandleInterest(ctx, &samplePool)
	require.NoError(t, err)

	delta = firstBorrowTime.Add(time.Second * month * 4).Sub(thirdBorrowTime)

	interest = r.Power(uint64(delta.Seconds())).Sub(sdk.OneDec())
	paymentAmount = interest.MulInt(sdk.NewIntFromUint64(2e8)).TruncateInt()
	reservedAmount = sdk.NewDecFromInt(paymentAmount).Mul(sdk.MustNewDecFromStr("0.15")).TruncateInt()
	toInvestors = paymentAmount.Sub(reservedAmount)

	poolInfo, _ = k.GetPools(ctx, poolIndex)
	b3 := poolInfo.PoolNFTIds[2]

	nclass3, _ := nftKeeper.GetClass(ctx, b3)
	err = proto.Unmarshal(nclass3.Data.Value, &borrowInterest)
	if err != nil {
		panic(err)
	}

	paymentTime = borrowInterest.Payments[1].PaymentTime
	require.EqualValues(t, firstBorrowTime.Add(time.Second*oneMonth*4), paymentTime)
	require.EqualValues(t, borrowInterest.Payments[1].PaymentAmount.Amount.String(), toInvestors.String())

	nclass2, _ = nftKeeper.GetClass(ctx, b2)
	err = proto.Unmarshal(nclass2.Data.Value, &borrowInterest)
	if err != nil {
		panic(err)
	}
	paymentTime = borrowInterest.Payments[3].PaymentTime
	require.EqualValues(t, firstBorrowTime.Add(time.Second*oneMonth*4), paymentTime)
	require.EqualValues(t, borrowInterest.Payments[3].PaymentAmount.Amount.String(), interestOneYear.String())

	nclass1, _ = nftKeeper.GetClass(ctx, b1)
	err = proto.Unmarshal(nclass1.Data.Value, &borrowInterest)
	if err != nil {
		panic(err)
	}
	paymentTime = borrowInterest.Payments[4].PaymentTime
	require.EqualValues(t, firstBorrowTime.Add(time.Second*oneMonth*4), paymentTime)
	require.EqualValues(t, borrowInterest.Payments[4].PaymentAmount.Amount.String(), interestOneYear.String())

}
