package keeper_test

import (
	"testing"
	"time"

	"github.com/cosmos/gogoproto/proto"
	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"

	types3 "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"

	"cosmossdk.io/x/nft"

	sdkmath "cosmossdk.io/math"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
	"github.com/stretchr/testify/suite"
)

type SPVRewardsTests struct {
	unitTester
}

func (suite *SPVRewardsTests) storedTimeEquals(denom string, expected time.Time) {
	storedTime, found := suite.keeper.GetSPVRewardAccrualTime(suite.ctx, denom)
	suite.True(found)
	suite.Equal(expected, storedTime)
}

func (suite *SPVRewardsTests) storedIndexesEqual(denom string, expected types2.RewardIndexes) {
	storedIndexes, found := suite.keeper.GetSPVReward(suite.ctx, denom)
	suite.Equal(found, expected != nil)

	if found {
		suite.Equal(expected, storedIndexes)
	} else {
		suite.Empty(storedIndexes)
	}
}

func TestSPVReward(t *testing.T) {
	suite.Run(t, new(SPVRewardsTests))
}

func (suite *SPVRewardsTests) TestStateUpdatedWhenBlockTimeHasIncreased() {
	denom := "bnb"

	spvKeeper := newFakeSPVKeeper()
	suite.keeper = suite.NewKeeper(&fakeParamSubspace{}, nil, nil, nil, nil, spvKeeper, nil)

	previousAccrualTime := time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC)
	suite.keeper.SetSPVRewardAccrualTime(suite.ctx, denom, previousAccrualTime)

	newAccrualTime := previousAccrualTime.Add(1 * time.Hour)
	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(newAccrualTime)

	endTime := previousAccrualTime.Add(time.Hour * 2)

	b1amount := int64(2000)
	b2amount := int64(1000)
	period := types2.NewMultiRewardPeriod(
		true,
		denom,
		time.Unix(0, 0), // ensure the test is within start and end times
		endTime,
		cs(c("jolt", b1amount), c("ujolt", b2amount)), // same denoms as in global indexes
	)

	suite.keeper.AccumulateSPVRewards(suite.ctx, period)
	reward, ok := suite.keeper.GetSPVReward(suite.ctx, denom)
	suite.Require().True(ok)

	// after one hour,the token amout should be
	expectedB1 := 2000 * 3600
	expectedB2 := 1000 * 3600

	expectedAmount := cs(c("jolt", int64(expectedB1)), c("ujolt", int64(expectedB2)))
	suite.Require().Equal(expectedAmount, reward.PaymentAmount)

	lasttime, ok := suite.keeper.GetSPVRewardAccrualTime(suite.ctx, denom)
	suite.Require().True(ok)

	suite.Require().True(lasttime.Equal(newAccrualTime))

	// 4 hours later
	newAccrualTime = newAccrualTime.Add(4 * time.Hour)
	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(newAccrualTime)

	suite.keeper.AccumulateSPVRewards(suite.ctx, period)
	reward, ok = suite.keeper.GetSPVReward(suite.ctx, denom)
	suite.Require().True(ok)

	expectedB1 = 2000 * 3600 * 2
	expectedB2 = 1000 * 3600 * 2
	expectedAmount = cs(c("jolt", int64(expectedB1)), c("ujolt", int64(expectedB2)))
	suite.Require().Equal(expectedAmount, reward.PaymentAmount)

	lasttime, ok = suite.keeper.GetSPVRewardAccrualTime(suite.ctx, denom)
	suite.Require().True(ok)

	suite.Require().True(lasttime.Equal(previousAccrualTime.Add(time.Hour * 2)))
}

func (suite *SPVRewardsTests) TestStateUpdatedWhenBlockTimeHasnotIncreased() {
	denom := "bnb"

	spvKeeper := newFakeSPVKeeper()
	suite.keeper = suite.NewKeeper(&fakeParamSubspace{}, nil, nil, nil, nil, spvKeeper, nil)

	previousAccrualTime := time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC)
	suite.keeper.SetSPVRewardAccrualTime(suite.ctx, denom, previousAccrualTime)

	newAccrualTime := previousAccrualTime.Add(10 * time.Microsecond)
	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(newAccrualTime)

	endTime := previousAccrualTime.Add(time.Hour * 2)

	b1amount := int64(2000)
	b2amount := int64(1000)
	period := types2.NewMultiRewardPeriod(
		true,
		denom,
		time.Unix(0, 0), // ensure the test is within start and end times
		endTime,
		cs(c("jolt", b1amount), c("ujolt", b2amount)), // same denoms as in global indexes
	)

	suite.keeper.AccumulateSPVRewards(suite.ctx, period)
	reward, ok := suite.keeper.GetSPVReward(suite.ctx, denom)
	suite.Require().True(ok)

	// after less then one second, the token amout should be 0
	suite.Require().Nil(reward.PaymentAmount)

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(previousAccrualTime.Add(1 * time.Hour))

	expectedB1 := 2000 * 3600
	expectedB2 := 1000 * 3600

	suite.keeper.AccumulateSPVRewards(suite.ctx, period)
	reward, ok = suite.keeper.GetSPVReward(suite.ctx, denom)
	suite.Require().True(ok)

	expectedAmount := cs(c("jolt", int64(expectedB1)), c("ujolt", int64(expectedB2)))
	suite.Require().Equal(expectedAmount, reward.PaymentAmount)

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(10 * time.Microsecond))

	suite.keeper.AccumulateSPVRewards(suite.ctx, period)
	reward, ok = suite.keeper.GetSPVReward(suite.ctx, denom)
	suite.Require().True(ok)
	suite.Require().Equal(expectedAmount, reward.PaymentAmount)
}

func (suite *SPVRewardsTests) TestStateUpdatedWhenBlockTimeHasDecreased() {
	rewardIndex := "test-pool"

	interestPaid := sdkmath.NewInt(7.058823529e18)

	spvKeeper := newFakeSPVKeeper()

	var borrowInterest types.BorrowInterest

	p1 := types.PaymentItem{
		PaymentAmount:  c("usdc", 1e18),
		BorrowedAmount: c("usdt", 1e18),
	}

	cc := c("usdc", 1e18)
	aa := types.BorrowDetail{BorrowedAmount: cc}
	datap1, err := types3.NewAnyWithValue(&types.BorrowInterest{
		Payments:      []*types.PaymentItem{&p1},
		BorrowDetails: []types.BorrowDetail{aa},
	})
	if err != nil {
		panic("pack class any data failed")
	}

	p2 := types.PaymentItem{
		PaymentAmount:  c("usdc", 0),
		BorrowedAmount: c("usdt", 0),
	}

	cc = c("usdc", 0)
	aa = types.BorrowDetail{BorrowedAmount: cc}

	datap2, err := types3.NewAnyWithValue(&types.BorrowInterest{
		Payments:      []*types.PaymentItem{&p2},
		BorrowDetails: []types.BorrowDetail{aa},
	})
	if err != nil {
		panic("pack class any data failed")
	}

	p3 := types.PaymentItem{
		PaymentAmount:  c("usdc", 2e18),
		BorrowedAmount: c("usdt", 2e18),
	}

	cc = c("usdc", 2e18)
	aa = types.BorrowDetail{BorrowedAmount: cc}

	datap3, err := types3.NewAnyWithValue(&types.BorrowInterest{
		Payments:      []*types.PaymentItem{&p3},
		BorrowDetails: []types.BorrowDetail{aa},
	})
	if err != nil {
		panic("pack class any data failed")
	}

	p4 := types.PaymentItem{
		PaymentAmount:  c("usdc", 25e17),
		BorrowedAmount: c("usdt", 25e17),
	}

	cc = c("usdc", 25e17)
	aa = types.BorrowDetail{BorrowedAmount: cc}

	datap4, err := types3.NewAnyWithValue(&types.BorrowInterest{Payments: []*types.PaymentItem{&p4}, BorrowDetails: []types.BorrowDetail{aa}})
	if err != nil {
		panic("pack class any data failed")
	}

	p5 := types.PaymentItem{
		PaymentAmount:  c("usdc", 5e17),
		BorrowedAmount: c("usdt", 5e17),
	}

	cc = c("usdc", 5e17)
	aa = types.BorrowDetail{BorrowedAmount: cc}
	datap5, err := types3.NewAnyWithValue(&types.BorrowInterest{Payments: []*types.PaymentItem{&p5}, BorrowDetails: []types.BorrowDetail{aa}})
	if err != nil {
		panic("pack class any data failed")
	}

	borrowInterest = types.BorrowInterest{
		Payments: nil,
	}

	datap6, err := types3.NewAnyWithValue(&borrowInterest)
	if err != nil {
		panic("pack class any data failed")
	}

	mockClass1 := nft.Class{
		Id:   "c1",
		Data: datap1,
	}

	mockClass2 := nft.Class{
		Id:   "c2",
		Data: datap2,
	}

	mockClass3 := nft.Class{
		Id:   "c3",
		Data: datap3,
	}

	mockClass4 := nft.Class{
		Id:   "c4",
		Data: datap4,
	}

	mockClass5 := nft.Class{
		Id:   "c5",
		Data: datap5,
	}

	mockClass6 := nft.Class{
		Id:   "c6",
		Data: datap6,
	}

	mockclasses := []*nft.Class{&mockClass1, &mockClass2, &mockClass3, &mockClass4, &mockClass5, &mockClass6}

	nftKeeper := newFakeNFTKeeper(mockclasses, nil)

	suite.keeper = suite.NewKeeper(&fakeParamSubspace{}, nil, nil, nil, nil, spvKeeper, nftKeeper)

	previousAccrualTime := time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC)
	suite.keeper.SetSPVRewardAccrualTime(suite.ctx, rewardIndex, previousAccrualTime)

	newAccrualTime := previousAccrualTime.Add(1 * time.Hour)
	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(newAccrualTime)

	endTime := previousAccrualTime.Add(time.Hour * 2)

	b1amount := int64(2000)
	b2amount := int64(1000)
	coins := cs(c("jolt", b1amount), c("ujolt", b2amount))
	period := types2.NewMultiRewardPeriod(
		true,
		rewardIndex,
		time.Unix(0, 0), // ensure the test is within start and end times
		endTime,
		coins, // same denoms as in global indexes
	)

	suite.keeper.AccumulateSPVRewards(suite.ctx, period)
	reward, ok := suite.keeper.GetSPVReward(suite.ctx, rewardIndex)
	suite.Require().True(ok)

	// after one hour,the token amout should be
	expectedB1 := 2000 * 3600
	expectedB2 := 1000 * 3600

	expectedAmount := cs(c("jolt", int64(expectedB1)), c("ujolt", int64(expectedB2)))
	suite.Require().Equal(expectedAmount, reward.PaymentAmount)

	suite.keeper.AfterSPVInterestPaid(suite.ctx, rewardIndex, interestPaid)

	reserveAmt := sdkmath.LegacyNewDecFromInt(interestPaid).Mul(sdkmath.LegacyMustNewDecFromStr("0.15")).TruncateInt()
	paymentToInvestor := interestPaid.Sub(reserveAmt)

	var allincentives []sdk.Coins
	allincentives = []sdk.Coins{}

	allps := []types.PaymentItem{p1, p2, p3, p4, p5}
	for _, el := range allps {
		ratio := sdkmath.LegacyNewDecFromInt(el.PaymentAmount.Amount).Quo(sdkmath.LegacyNewDecFromInt(paymentToInvestor))
		var incentiveCoins sdk.Coins
		for _, eachCoin := range reward.PaymentAmount {
			amt := sdkmath.LegacyNewDecFromInt(eachCoin.Amount).Mul(ratio).TruncateInt()
			incentiveCoins = incentiveCoins.Add(sdk.NewCoin(eachCoin.Denom, amt))
		}
		allincentives = append(allincentives, incentiveCoins)
	}

	ids := []string{"c1", "c2", "c3", "c4", "c5", "c6"}
	for index, el := range ids {
		nft, found := suite.keeper.NftKeeper.GetClass(suite.ctx, el)
		suite.Assert().True(found)
		var borrowInterest types.BorrowInterest
		var err error
		err = proto.Unmarshal(nft.Data.Value, &borrowInterest)
		if err != nil {
			panic(err)
		}
		// if we do not have the payment info, we skip
		if len(borrowInterest.IncentivePayments) == 0 {
			suite.Assert().Equal("c6", nft.Id)
			continue
		}
		lastpayment := borrowInterest.IncentivePayments[len(borrowInterest.Payments)-1]
		suite.Require().True(lastpayment.PaymentAmount.IsEqual(allincentives[index]))
	}

	reward, ok = suite.keeper.GetSPVReward(suite.ctx, rewardIndex)
	suite.Require().True(ok)

	suite.Require().True(reward.PaymentAmount.IsZero())

	// now we do the allication again

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Hour))
	suite.keeper.AccumulateSPVRewards(suite.ctx, period)

	suite.keeper.AfterSPVInterestPaid(suite.ctx, rewardIndex, interestPaid)

	reserveAmt = sdkmath.LegacyNewDecFromInt(interestPaid).Mul(sdkmath.LegacyMustNewDecFromStr("0.15")).TruncateInt()
	paymentToInvestor = interestPaid.Sub(reserveAmt)

	for index, el := range ids {
		nft, found := suite.keeper.NftKeeper.GetClass(suite.ctx, el)
		suite.Assert().True(found)
		var borrowInterest types.BorrowInterest
		var err error
		err = proto.Unmarshal(nft.Data.Value, &borrowInterest)
		if err != nil {
			panic(err)
		}
		// if we do not have the payment info, we skip
		if len(borrowInterest.IncentivePayments) == 0 {
			suite.Assert().Equal("c6", nft.Id)
			continue
		}
		lastpayment := borrowInterest.IncentivePayments[len(borrowInterest.Payments)-1]
		suite.Require().True(lastpayment.PaymentAmount.IsEqual(allincentives[index]))
		suite.Require().Len(borrowInterest.IncentivePayments, 2)
	}

	reward, ok = suite.keeper.GetSPVReward(suite.ctx, rewardIndex)
	suite.Require().True(ok)

	suite.Require().True(reward.PaymentAmount.IsZero())
}

func (suite *SPVRewardsTests) TestClaimIncentives() {
	spvKeeper := newFakeSPVKeeper()

	var allIncentivepayments []*types.IncentivePaymentItem

	for i := 0; i < 50; i++ {
		amt := int64(i * 100e5)
		amt2 := int64(2 * i * 700e5)
		p1 := types.IncentivePaymentItem{
			PaymentAmount:  cs(c("usdc", amt), c("usdt", amt2)),
			BorrowedAmount: c("usdt", 1e18),
		}
		allIncentivepayments = append(allIncentivepayments, &p1)
	}

	cc := c("usdc", 1e18)
	aa := types.BorrowDetail{BorrowedAmount: cc}

	datap2, err := types3.NewAnyWithValue(&types.BorrowInterest{
		IncentivePayments: allIncentivepayments,
		BorrowDetails:     []types.BorrowDetail{aa},
	})
	suite.Require().NoError(err)

	mockClass4 := nft.Class{
		Id:   "p1",
		Data: datap2,
	}

	nft1Info := types.NftInfo{
		Borrowed: c("usdt", 2e17),
	}

	nft2info := types.NftInfo{
		Borrowed: c("usdt", 3e17),
	}

	nft1d, err := types3.NewAnyWithValue(&nft1Info)
	suite.Require().NoError(err)
	nft2d, err := types3.NewAnyWithValue(&nft2info)

	nft1 := nft.NFT{
		Id:   "test-nft1",
		Data: nft1d,
	}

	nft2 := nft.NFT{
		Id:   "test-nft2",
		Data: nft2d,
	}

	nftKeeper := newFakeNFTKeeper([]*nft.Class{&mockClass4}, []*nft.NFT{&nft1, &nft2})

	suite.keeper = suite.NewKeeper(&fakeParamSubspace{}, nil, nil, nil, nil, spvKeeper, nftKeeper)

	c, err := keeper.CalculateTotalIncentives(suite.ctx, []string{"p1:test-nft1", "p1:test-nft2"}, nftKeeper, false)
	suite.Require().NoError(err)

	circleNum := 0
	rewards := calcualteInterest([]sdkmath.Int{sdkmath.NewInt(2e17), sdkmath.NewInt(3e17)}, allIncentivepayments, circleNum)
	suite.Require().True(c.IsEqual(rewards))

	c2, err := keeper.CalculateTotalIncentives(suite.ctx, []string{"p1:test-nft1", "p1:test-nft2"}, nftKeeper, false)
	suite.Require().NoError(err)
	suite.Require().True(c2.IsEqual(rewards))

	c, err = keeper.CalculateTotalIncentives(suite.ctx, []string{"p1:test-nft1", "p1:test-nft2"}, nftKeeper, true)
	suite.Require().NoError(err)

	suite.Require().True(c.IsEqual(rewards))

	c2, err = keeper.CalculateTotalIncentives(suite.ctx, []string{"p1:test-nft1", "p1:test-nft2"}, nftKeeper, false)
	suite.Require().NoError(err)
	suite.Require().True(c2.IsZero())

	// now we have new reward
	classinf, ok := suite.keeper.NftKeeper.GetClass(suite.ctx, "p1")
	suite.Require().True(ok)
	var borrowInterest types.BorrowInterest
	err = proto.Unmarshal(classinf.Data.Value, &borrowInterest)
	suite.Require().NoError(err)

	new2payments := allIncentivepayments[:2]

	borrowInterest.IncentivePayments = append(borrowInterest.IncentivePayments, new2payments...)

	datap2, err = types3.NewAnyWithValue(&borrowInterest)
	suite.Require().NoError(err)
	classinf.Data = datap2
	err = suite.keeper.NftKeeper.UpdateClass(suite.ctx, classinf)
	suite.Require().NoError(err)

	c, err = keeper.CalculateTotalIncentives(suite.ctx, []string{"p1:test-nft1", "p1:test-nft2"}, nftKeeper, false)
	suite.Require().NoError(err)

	rewards = calcualteInterest([]sdkmath.Int{sdkmath.NewInt(2e17), sdkmath.NewInt(3e17)}, new2payments, circleNum)
	suite.Require().True(c.IsEqual(rewards))

	c, err = keeper.CalculateTotalIncentives(suite.ctx, []string{"p1:test-nft1", "p1:test-nft2"}, nftKeeper, true)
	suite.Require().NoError(err)
	suite.Require().True(c.IsEqual(rewards))

	c2, err = keeper.CalculateTotalIncentives(suite.ctx, []string{"p1:test-nft1", "p1:test-nft2"}, nftKeeper, true)
	suite.Require().NoError(err)
	suite.Require().True(c2.IsZero())
}

func calcualteInterest(nftBorrow []sdkmath.Int, incentivePayments []*types.IncentivePaymentItem, offset int) sdk.Coins {
	rewards := sdk.NewCoins()

	for _, eachNFTBorrow := range nftBorrow {

		payments := incentivePayments[offset:]

		for _, eachIncentivePayment := range payments {
			if eachIncentivePayment.PaymentAmount.IsZero() {
				continue
			}
			classBorrowedAmount := eachIncentivePayment.BorrowedAmount
			incentivePaymentAmount := eachIncentivePayment.PaymentAmount
			// todo there may be the case that because of the truncate, the total payment is larger than the interest paid to investors
			// fixme we should NEVER calculate the interest after the pool status is in luquidation as the user ratio is not correct any more

			var incentiveCoins sdk.Coins
			for _, eachCoin := range incentivePaymentAmount {
				incentiveAmt := sdkmath.LegacyNewDecFromInt(eachCoin.Amount).Mul(sdkmath.LegacyNewDecFromInt(eachNFTBorrow)).Quo(sdkmath.LegacyNewDecFromInt(classBorrowedAmount.Amount)).TruncateInt()
				incentive := sdk.NewCoin(eachCoin.Denom, incentiveAmt)
				incentiveCoins = incentiveCoins.Add(incentive)
			}

			incentiveCoins.Sort()
			rewards = rewards.Add(incentiveCoins...)
		}
	}
	return rewards
}
