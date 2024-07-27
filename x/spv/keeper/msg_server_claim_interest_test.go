package keeper_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/stretchr/testify/require"

	sdkmath "cosmossdk.io/math"

	"github.com/gogo/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/suite"
)

// Test suite used for all keeper tests
type claimInterestSuite struct {
	suite.Suite
	keeper       *spvkeeper.Keeper
	nftKeeper    types.NFTKeeper
	app          types.MsgServer
	ctx          context.Context
	investors    []string
	investorPool string
}

func SetupPool(suite *claimInterestSuite) {
	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 3, PoolName: "hello", Apy: []string{"0.15", "0.0875"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdkmath.NewInt(3*1e9)), sdk.NewCoin("ausdc", sdkmath.NewInt(3*1e9))}}
	resp, err := suite.app.CreatePool(suite.ctx, &req)
	suite.Require().NoError(err)

	depositorPool := resp.PoolIndex[0]

	suite.investorPool = depositorPool

	_, err = suite.app.ActivePool(suite.ctx, types.NewMsgActivePool("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", resp.PoolIndex[0]))
	suite.Require().NoError(err)

	_, err = suite.app.ActivePool(suite.ctx, types.NewMsgActivePool("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", resp.PoolIndex[1]))
	suite.Require().NoError(err)

	req2 := types.MsgAddInvestors{
		Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: resp.PoolIndex[0],
		InvestorID: []string{"2"},
	}
	_, err = suite.app.AddInvestors(suite.ctx, &req2)
	suite.Require().NoError(err)

	creator1 := "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl"
	creator2 := "jolt1kkujrm0lqeu0e5va5f6mmwk87wva0k8cmam8jq"

	suite.investors = []string{creator1, creator2}
}

// The default state used by each test
func (suite *claimInterestSuite) SetupTest() {
	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)
	lapp, k, nftKeeper, _, _, wctx := setupMsgServer(suite.T())
	ctx := sdk.UnwrapSDKContext(wctx)
	// create the first pool apy 7.8%

	suite.ctx = ctx
	suite.keeper = k
	suite.app = lapp
	suite.nftKeeper = nftKeeper
}

func TestClaimInterestTestSuite(t *testing.T) {
	suite.Run(t, new(claimInterestSuite))
}

func checkInterestCorrectness(suite *claimInterestSuite, creatorAddr1, creatorAddr2 sdk.AccAddress, nftIndex int, expectedAmount1, expectedAmount2 string) {
	depositor1, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	depositor2, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)

	dat := strings.Split(depositor1.LinkedNFT[nftIndex], ":")
	nftUser1, found := suite.nftKeeper.GetNFT(suite.ctx, dat[0], dat[1])
	suite.Require().True(found)

	Borrowclass, _ := suite.nftKeeper.GetClass(suite.ctx, dat[0])

	var borrowClassInfo types.BorrowInterest
	err := proto.Unmarshal(Borrowclass.Data.Value, &borrowClassInfo)
	if err != nil {
		panic(err)
	}

	lastBorrow := borrowClassInfo.BorrowDetails[len(borrowClassInfo.BorrowDetails)-1].BorrowedAmount
	period := spvkeeper.OneYear / spvkeeper.OneMonth
	interestOneYearWithReserve := sdkmath.LegacyNewDecFromInt(lastBorrow.Amount).Mul(sdkmath.LegacyMustNewDecFromStr("0.15")).QuoInt64(int64(period)).TruncateInt()

	interestOneMonth := interestOneYearWithReserve.Sub(sdkmath.LegacyNewDecFromInt(interestOneYearWithReserve).Mul(sdkmath.LegacyMustNewDecFromStr("0.15")).TruncateInt())

	dat = strings.Split(depositor2.LinkedNFT[nftIndex], ":")
	nftUser2, found := suite.nftKeeper.GetNFT(suite.ctx, dat[0], dat[1])
	suite.Require().True(found)

	var nftInfo1 types.NftInfo
	err = proto.Unmarshal(nftUser1.Data.Value, &nftInfo1)
	if err != nil {
		panic(err)
	}

	var nftInfo2 types.NftInfo
	err = proto.Unmarshal(nftUser2.Data.Value, &nftInfo2)
	if err != nil {
		panic(err)
	}

	totalBorrowed := nftInfo1.Borrowed.Add(nftInfo2.Borrowed)

	suite.Require().True(totalBorrowed.Amount.Equal(lastBorrow.Amount))

	ratio1 := sdkmath.LegacyNewDecFromInt(nftInfo1.Borrowed.Amount).Quo(sdkmath.LegacyNewDecFromInt(lastBorrow.Amount))
	ratio2 := sdkmath.LegacyNewDecFromInt(nftInfo2.Borrowed.Amount).Quo(sdkmath.LegacyNewDecFromInt(lastBorrow.Amount))

	user1Interest := sdkmath.LegacyNewDecFromInt(interestOneMonth).Mul(ratio1).TruncateInt()
	user2Interest := sdkmath.LegacyNewDecFromInt(interestOneMonth).Mul(ratio2).TruncateInt()

	c1, err := sdk.ParseCoinsNormalized(expectedAmount1 + "ausdc")
	suite.Require().NoError(err)
	c2, err := sdk.ParseCoinsNormalized(expectedAmount2 + "ausdc")
	suite.Require().NoError(err)

	totalPaid := sdkmath.ZeroInt()
	for _, el := range borrowClassInfo.Payments {
		totalPaid = totalPaid.Add(el.PaymentAmount.Amount)
	}
	// totalPaid := borrowClassInfo.Payments[len(borrowClassInfo.Payments)-1].PaymentAmount.Amount
	delta2 := borrowClassInfo.AccInterest.Amount.Sub(borrowClassInfo.InterestPaid.Amount)
	delta := totalPaid.Sub(borrowClassInfo.InterestPaid.Amount)

	suite.Require().Equal(delta, delta2)
	suite.Require().True(c1[0].Add(c2[0]).Amount.Sub(user1Interest.Add(user2Interest)).Abs().LTE(delta))

	suite.Require().True(c1[0].Amount.Sub(user1Interest).Abs().LTE(delta))
	suite.Require().True(c2[0].Amount.Sub(user2Interest).Abs().LTE(delta))
}

func (suite *claimInterestSuite) TestClaimInterestMultipleMonth() {
	investor1TotalClaimed := sdk.OneInt()
	SetupPool(suite)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	poolInfo.PoolTotalBorrowLimit = 100
	poolInfo.TargetAmount = sdk.NewCoin("ausdc", sdkmath.NewInt(600000))
	suite.keeper.SetPool(suite.ctx, poolInfo)

	// now we deposit some token and it should be enough to borrow
	creator1 := suite.investors[0]
	creator2 := suite.investors[1]
	creatorAddr1, err := sdk.AccAddressFromBech32(creator1)
	suite.Require().NoError(err)
	creatorAddr2, err := sdk.AccAddressFromBech32(creator2)
	suite.Require().NoError(err)
	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(4e5))
	// suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{
		Creator:   creator1,
		PoolIndex: suite.investorPool,
		Token:     depositAmount,
	}

	// user two deposit half of the amount of the user 1
	msgDepositUser2 := &types.MsgDeposit{
		Creator:   creator2,
		PoolIndex: suite.investorPool,
		Token:     depositAmount.SubAmount(sdkmath.NewInt(2e5)),
	}

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	// now we borrow 1.34e5
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	depositor1, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	investor1Borrowed := depositor1.LockedAmount.Amount

	reqInterest := types.MsgRepayInterest{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(8e9))}
	_, err = suite.app.RepayInterest(suite.ctx, &reqInterest)
	suite.Require().NoError(err)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(spvkeeper.OneMonth)))
	err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)
	suite.Require().NoError(err)

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(spvkeeper.OneMonth)))
	err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)
	suite.Require().NoError(err)

	req := types.MsgClaimInterest{
		Creator:   creator1,
		PoolIndex: suite.investorPool,
	}

	result1, err := suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)

	req.Creator = suite.investors[1]
	result2, err := suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)

	a1, _ := sdk.ParseCoinsNormalized(result1.Amount)

	investor1TotalClaimed = investor1TotalClaimed.Add(a1.AmountOf("ausdc"))
	a2, _ := sdk.ParseCoinsNormalized(result2.Amount)

	amount1 := a1.AmountOf("ausdc").Quo(sdkmath.NewInt(2))
	amount2 := a2.AmountOf("ausdc").Quo(sdkmath.NewInt(2))

	checkInterestCorrectness(suite, creatorAddr1, creatorAddr2, 0, convertBorrowToLocal(amount1).String(), convertBorrowToLocal(amount2).String())

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(spvkeeper.OneMonth)))
	err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)
	suite.Require().NoError(err)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	req.Creator = suite.investors[0]
	result1, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)

	req.Creator = suite.investors[1]
	result2, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)

	a1, _ = sdk.ParseCoinsNormalized(result1.Amount)
	investor1TotalClaimed = investor1TotalClaimed.Add(a1.AmountOf("ausdc"))
	a2, _ = sdk.ParseCoinsNormalized(result2.Amount)

	amount1 = a1.AmountOf("ausdc").Quo(sdkmath.NewInt(1))
	amount2 = a2.AmountOf("ausdc").Quo(sdkmath.NewInt(1))

	fmt.Printf("amoutn 1 %v amount 2 %v\n", amount1, amount2)
	checkInterestCorrectness(suite, creatorAddr1, creatorAddr2, 0, convertBorrowToLocal(amount1).String(), convertBorrowToLocal(amount2).String())

	// for the rest of 10 month
	for i := 0; i < 10; i++ {
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(spvkeeper.OneMonth)))
		err := suite.keeper.HandleInterest(suite.ctx, &poolInfo)
		suite.Require().NoError(err)
	}

	req.Creator = suite.investors[0]
	result1, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)

	req.Creator = suite.investors[1]
	result2, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)

	a1, _ = sdk.ParseCoinsNormalized(result1.Amount)
	investor1TotalClaimed = investor1TotalClaimed.Add(a1.AmountOf("ausdc"))
	a2, _ = sdk.ParseCoinsNormalized(result2.Amount)

	amount1 = a1.AmountOf("ausdc").Quo(sdkmath.NewInt(10))
	amount2 = a2.AmountOf("ausdc").Quo(sdkmath.NewInt(10))
	checkInterestCorrectness(suite, creatorAddr1, creatorAddr2, 0, convertBorrowToLocal(amount1).String(), convertBorrowToLocal(amount2).String())

	expectedInterest := sdkmath.LegacyNewDecFromInt(convertBorrowToUsd(investor1Borrowed)).Mul(sdkmath.LegacyMustNewDecFromStr("0.15")).Mul(sdk.NewDecWithPrec(85, 2)).TruncateInt()
	fmt.Printf("expected interest %v, investor1TotalClaimed %v", expectedInterest, investor1TotalClaimed)

	dat := strings.Split(depositor1.LinkedNFT[0], ":")
	Borrowclass, _ := suite.nftKeeper.GetClass(suite.ctx, dat[0])

	var borrowClassInfo types.BorrowInterest
	err = proto.Unmarshal(Borrowclass.Data.Value, &borrowClassInfo)
	suite.Require().NoError(err)

	delta := borrowClassInfo.AccInterest.Sub(borrowClassInfo.InterestPaid)

	suite.Require().True(expectedInterest.Equal(investor1TotalClaimed.Add(delta.Amount)))
}

func (suite *claimInterestSuite) TestClaimInterestMultipleBorrow() {
	investor1TotalClaimed := sdk.OneInt()
	SetupPool(suite)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	poolInfo.PoolTotalBorrowLimit = 100
	poolInfo.TargetAmount = sdk.NewCoin("ausdc", sdkmath.NewInt(600000))
	suite.keeper.SetPool(suite.ctx, poolInfo)

	// now we deposit some token and it should be enough to borrow1
	creator1 := suite.investors[0]
	creator2 := suite.investors[1]
	creatorAddr1, err := sdk.AccAddressFromBech32(creator1)
	suite.Require().NoError(err)
	creatorAddr2, err := sdk.AccAddressFromBech32(creator2)
	suite.Require().NoError(err)
	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(4e5))
	// suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{
		Creator:   creator1,
		PoolIndex: suite.investorPool,
		Token:     depositAmount,
	}

	// user two deposit half of the amount of the user 1
	msgDepositUser2 := &types.MsgDeposit{
		Creator:   creator2,
		PoolIndex: suite.investorPool,
		Token:     depositAmount.SubAmount(sdkmath.NewInt(2e5)),
	}

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	borrow1 := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	// now we borrow1 1.34e5
	_, err = suite.app.Borrow(suite.ctx, borrow1)
	suite.Require().NoError(err)

	depositor1, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	reqInterest := types.MsgRepayInterest{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1e4))}
	_, err = suite.app.RepayInterest(suite.ctx, &reqInterest)
	suite.Require().NoError(err)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(spvkeeper.OneMonth)))
	err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)
	suite.Require().NoError(err)

	req := types.MsgClaimInterest{
		Creator:   creator1,
		PoolIndex: suite.investorPool,
	}

	result1, err := suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)

	req.Creator = suite.investors[1]
	result2, err := suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)

	a1, _ := sdk.ParseCoinsNormalized(result1.Amount)

	investor1TotalClaimed = investor1TotalClaimed.Add(a1.AmountOf("ausdc"))
	a2, _ := sdk.ParseCoinsNormalized(result2.Amount)

	amount1 := a1.AmountOf("ausdc").Quo(sdkmath.NewInt(1))
	amount2 := a2.AmountOf("ausdc").Quo(sdkmath.NewInt(1))

	dat := strings.Split(depositor1.LinkedNFT[0], ":")
	nftUser1, found := suite.nftKeeper.GetNFT(suite.ctx, dat[0], dat[1])
	suite.Require().True(found)

	Borrowclass, _ := suite.nftKeeper.GetClass(suite.ctx, dat[0])

	var borrowClassInfo1 types.BorrowInterest
	err = proto.Unmarshal(Borrowclass.Data.Value, &borrowClassInfo1)
	if err != nil {
		panic(err)
	}

	suite.Require().True(a1.Add(a2[0])[0].IsEqual(borrowClassInfo1.InterestPaid))

	checkInterestCorrectness(suite, creatorAddr1, creatorAddr2, 0, convertBorrowToLocal(amount1).String(), convertBorrowToLocal(amount2).String())

	// now we pay out the interest for all the prepaid interest

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	payCounter := poolInfo.InterestPrepayment.Counter
	for i := 1; i < int(payCounter+1); i++ {
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(poolInfo.PayFreq)))
		err := suite.keeper.HandleInterest(suite.ctx, &poolInfo)
		suite.Require().NoError(err)
	}

	poolCheck, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(poolCheck.InterestPrepayment == nil)

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(3600*24*15)))

	// clear the interest
	req.Creator = suite.investors[0]
	_, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)

	req.Creator = suite.investors[1]
	_, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)

	// ########we will run the payment again

	//######now we borrow1 2.1e5########
	borrow2 := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(2.1e5))}
	_, err = suite.app.Borrow(suite.ctx, borrow2)
	suite.Require().NoError(err)
	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	depositor1, found = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	depositor2, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)

	dat = strings.Split(depositor1.LinkedNFT[1], ":")
	nftUser1, found = suite.nftKeeper.GetNFT(suite.ctx, dat[0], dat[1])
	suite.Require().True(found)

	dat = strings.Split(depositor2.LinkedNFT[1], ":")
	nftUser2, found := suite.nftKeeper.GetNFT(suite.ctx, dat[0], dat[1])
	suite.Require().True(found)

	var nftInfo12 types.NftInfo
	err = proto.Unmarshal(nftUser1.Data.Value, &nftInfo12)
	if err != nil {
		panic(err)
	}

	var nftInfo22 types.NftInfo
	err = proto.Unmarshal(nftUser2.Data.Value, &nftInfo22)
	if err != nil {
		panic(err)
	}

	Borrowclass, found = suite.nftKeeper.GetClass(suite.ctx, dat[0])
	suite.Require().True(found)

	var borrowClassInfo types.BorrowInterest
	err = proto.Unmarshal(Borrowclass.Data.Value, &borrowClassInfo)
	if err != nil {
		panic(err)
	}

	// pay the interest again
	reqInterest.Token = sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1e5))
	_, err = suite.app.RepayInterest(suite.ctx, &reqInterest)
	suite.Require().NoError(err)
	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	totalCounter := poolInfo.InterestPrepayment.Counter
	for i := 0; i < int(totalCounter); i++ {
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(poolInfo.PayFreq)))
		err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)
	}

	suite.Require().Nil(poolInfo.InterestPrepayment)

	req.Creator = suite.investors[0]
	result1, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)

	req.Creator = suite.investors[1]
	result2, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)

	expectedUser1TotalfromTheFirstBorrow := amount1.MulRaw(int64(totalCounter))
	expectedUser2TotalfromTheFirstBorrow := amount2.MulRaw(int64(totalCounter))

	totalPayment := sdkmath.ZeroInt()

	Borrowclasss, _ := suite.nftKeeper.GetClass(suite.ctx, "class-e0d49c3eed41e408b493a14042a8aa31375d64e3e357f911afbb085e02bde083-1")
	err = proto.Unmarshal(Borrowclasss.Data.Value, &borrowClassInfo)
	if err != nil {
		panic(err)
	}

	for _, el := range borrowClassInfo.Payments {
		totalPayment = totalPayment.Add(el.PaymentAmount.Amount)
	}

	// this payfreq is from the testutil spv.go
	paymentAmount := totalPayment
	// reservedAmount := sdkmath.LegacyNewDecFromInt(paymentAmount).Mul(sdkmath.LegacyMustNewDecFromStr("0.15")).TruncateInt()
	toInvestors := paymentAmount

	lastBorrow := borrowClassInfo.BorrowDetails[len(borrowClassInfo.BorrowDetails)-1].BorrowedAmount

	ratio1 := sdkmath.LegacyNewDecFromInt(nftInfo12.Borrowed.Amount).Quo(sdkmath.LegacyNewDecFromInt(lastBorrow.Amount))
	ratio2 := sdkmath.LegacyNewDecFromInt(nftInfo22.Borrowed.Amount).Quo(sdkmath.LegacyNewDecFromInt(lastBorrow.Amount))

	expectedToUser2 := sdkmath.LegacyNewDecFromInt(toInvestors).Mul(ratio2).TruncateInt()
	expectedToUser1 := sdkmath.LegacyNewDecFromInt(toInvestors).Mul(ratio1).TruncateInt()
	// we add the interest from the first borrow
	expectedToUser1 = expectedToUser1.Add(expectedUser1TotalfromTheFirstBorrow)
	expectedToUser2 = expectedToUser2.Add(expectedUser2TotalfromTheFirstBorrow)

	resultCoinUser1, err := sdk.ParseCoinNormalized(result1.Amount)
	if err != nil {
		panic(err)
	}
	resultCoinUser2, err := sdk.ParseCoinNormalized(result2.Amount)
	if err != nil {
		panic(err)
	}
	// as we have the division, so the max error is totalCounter *1 (1: the max error of the division for each payment)

	suite.Require().True(expectedToUser1.Sub(resultCoinUser1.Amount).LT(sdkmath.NewInt(int64(totalCounter))))
	suite.Require().True(expectedToUser2.Sub(resultCoinUser2.Amount).LT(sdkmath.NewInt(int64(totalCounter))))

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(spvkeeper.OneMonth)))
	req.Creator = suite.investors[0]
	result1, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)
	req.Creator = suite.investors[1]
	result2, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)
}

func (suite *claimInterestSuite) TestClaimInterest() {
	SetupPool(suite)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	poolInfo.PoolTotalBorrowLimit = 100
	poolInfo.TargetAmount = sdk.NewCoin("ausdc", sdkmath.NewInt(600000))
	suite.keeper.SetPool(suite.ctx, poolInfo)

	// now we deposit some token and it should be enough to borrow
	creator1 := suite.investors[0]
	creator2 := suite.investors[1]
	creatorAddr1, err := sdk.AccAddressFromBech32(creator1)
	suite.Require().NoError(err)
	creatorAddr2, err := sdk.AccAddressFromBech32(creator2)
	suite.Require().NoError(err)
	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(4e5))
	// suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{
		Creator:   creator1,
		PoolIndex: suite.investorPool,
		Token:     depositAmount,
	}

	// user two deposit half of the amount of the user 1
	msgDepositUser2 := &types.MsgDeposit{
		Creator:   creator2,
		PoolIndex: suite.investorPool,
		Token:     depositAmount.SubAmount(sdkmath.NewInt(2e5)),
	}

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	// now we borrow 1.34e5
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	reqInterest := types.MsgRepayInterest{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(8e9))}
	_, err = suite.app.RepayInterest(suite.ctx, &reqInterest)
	suite.Require().NoError(err)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(spvkeeper.OneMonth)))
	err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)
	suite.Require().NoError(err)

	req := types.MsgClaimInterest{
		Creator:   creator1,
		PoolIndex: suite.investorPool,
	}

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(spvkeeper.OneMonth)))
	result1, err := suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)

	req.Creator = suite.investors[1]
	result2, err := suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)
	// total borrow is 1.34e5 and the user1 offer 4e5 while the second user offer 2e5, reserve is 0.15, the apy is 0.15
	// after the first month, user 1 get 1.34e5*0.85*0.15/12*4/6
	// interestOneYearWithReserve := sdkmath.LegacyNewDecFromInt(sdkmath.NewIntFromUint64(1.34e5)).Mul(sdkmath.LegacyMustNewDecFromStr("0.15")).QuoInt64(12).TruncateInt()
	// interestOneYear := interestOneYearWithReserve.Sub(sdkmath.LegacyNewDecFromInt(interestOneYearWithReserve).Mul(sdkmath.LegacyMustNewDecFromStr("0.15")).TruncateInt())

	a1, _ := sdk.ParseCoinsNormalized(result1.Amount)
	a2, _ := sdk.ParseCoinsNormalized(result2.Amount)

	amount1 := a1.AmountOf("ausdc").Quo(sdkmath.NewInt(1))
	amount2 := a2.AmountOf("ausdc").Quo(sdkmath.NewInt(1))

	checkInterestCorrectness(suite, creatorAddr1, creatorAddr2, 0, convertBorrowToLocal(amount1).String(), convertBorrowToLocal(amount2).String())

	result1, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)
	suite.Require().True(result1.Amount == "0ausdc")

	// we add one second after the withdraw the interest

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(100)))
	result1, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)
	suite.Require().True(result1.Amount == "0ausdc")

	// we add less than a month, so the amount should still be zero
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(spvkeeper.OneMonth-101)))
	result1, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)
	suite.Require().True(result1.Amount == "0ausdc")

	// since the spv not paid the interest, we cannont claim the interest
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(spvkeeper.OneMonth+150)))
	result1, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)
	suite.Require().True(result1.Amount == "0ausdc")
}

func (suite *claimInterestSuite) TestClaimInterestNoAuthorized() {
	SetupPool(suite)
	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	poolInfo.PoolTotalBorrowLimit = 100
	poolInfo.TargetAmount = sdk.NewCoin("ausdc", sdkmath.NewInt(600000))
	suite.keeper.SetPool(suite.ctx, poolInfo)

	// now we deposit some token and it should be enough to borrow
	creator1 := suite.investors[0]
	creator2 := suite.investors[1]
	// creatorAddr1, err := sdk.AccAddressFromBech32(creator1)
	// suite.Require().NoError(err)
	// creatorAddr2, err := sdk.AccAddressFromBech32(creator2)
	// suite.Require().NoError(err)
	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(4e5))
	// suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{
		Creator:   creator1,
		PoolIndex: suite.investorPool,
		Token:     depositAmount,
	}

	// user two deposit half of the amount of the user 1
	msgDepositUser2 := &types.MsgDeposit{
		Creator:   creator2,
		PoolIndex: suite.investorPool,
		Token:     depositAmount.SubAmount(sdkmath.NewInt(2e5)),
	}

	_, err := suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	// now we borrow 1.34e5
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	reqInterest := types.MsgRepayInterest{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(8e9))}
	_, err = suite.app.RepayInterest(suite.ctx, &reqInterest)
	suite.Require().NoError(err)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	month := 3600 * 24 * 30

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(month)))
	err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)
	suite.Require().NoError(err)

	req := types.MsgClaimInterest{
		Creator:   creator1,
		PoolIndex: suite.investorPool,
	}

	req.Creator = "jolt15anuxmcus4tyh2rttydj0cyfa8ldfg9akdek0f"

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(month)))
	_, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().ErrorContains(err, "not found for pool index")

	req.Creator = "not the depositor"
	_, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().ErrorContains(err, "invalid address")

	req.Creator = suite.investors[0]
	req.PoolIndex = "invalid"
	_, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().ErrorContains(err, "not found for pool index")
}

func (suite *claimInterestSuite) TestQueryOutStandingInterest() {
	SetupPool(suite)
	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	poolInfo.PoolTotalBorrowLimit = 100
	poolInfo.TargetAmount = sdk.NewCoin("ausdc", sdkmath.NewInt(600000))
	suite.keeper.SetPool(suite.ctx, poolInfo)

	// now we deposit some token and it should be enough to borrow
	creator1 := suite.investors[0]
	creator2 := suite.investors[1]
	// creatorAddr1, err := sdk.AccAddressFromBech32(creator1)
	// suite.Require().NoError(err)
	// creatorAddr2, err := sdk.AccAddressFromBech32(creator2)
	// suite.Require().NoError(err)
	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(4e5))
	// suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{
		Creator:   creator1,
		PoolIndex: suite.investorPool,
		Token:     depositAmount,
	}

	// user two deposit half of the amount of the user 1
	msgDepositUser2 := &types.MsgDeposit{
		Creator:   creator2,
		PoolIndex: suite.investorPool,
		Token:     depositAmount.SubAmount(sdkmath.NewInt(2e5)),
	}

	_, err := suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	// now we borrow 1.34e5
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	// now we test the claimable interest

	resp, err := suite.keeper.OutstandingInterest(suite.ctx, &types.QueryOutstandingInterestRequest{
		Wallet:    creator1,
		PoolIndex: suite.investorPool,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(resp.Amount, "0")
	fmt.Printf(">>>>>%v\n", poolInfo.PayFreq)

	reqInterest := types.MsgRepayInterest{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(8e9))}
	_, err = suite.app.RepayInterest(suite.ctx, &reqInterest)
	suite.Require().NoError(err)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(spvkeeper.OneMonth)))
	err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)
	suite.Require().NoError(err)

	req := types.MsgClaimInterest{
		Creator:   creator1,
		PoolIndex: suite.investorPool,
	}

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(spvkeeper.OneMonth)))
	result1, err := suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)

	req.Creator = suite.investors[1]
	result2, err := suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)
	// total borrow is 1.34e5 and the user1 offer 4e5 while the second user offer 2e5, reserve is 0.15, the apy is 0.15
	// after the first month, user 1 get 1.34e5*0.85*0.15/12*4/6
	// interestOneYearWithReserve := sdkmath.LegacyNewDecFromInt(sdkmath.NewIntFromUint64(1.34e5)).Mul(sdkmath.LegacyMustNewDecFromStr("0.15")).QuoInt64(12).TruncateInt()
	// interestOneYear := interestOneYearWithReserve.Sub(sdkmath.LegacyNewDecFromInt(interestOneYearWithReserve).Mul(sdkmath.LegacyMustNewDecFromStr("0.15")).TruncateInt())

	a1, _ := sdk.ParseCoinsNormalized(result1.Amount)
	a2, _ := sdk.ParseCoinsNormalized(result2.Amount)

	oneMonthInvestor1 := a1[0]
	oneMonthInvestor2 := a2[0]

	creatorAddr1, err := sdk.AccAddressFromBech32(creator1)
	suite.Require().NoError(err)
	creatorAddr2, err := sdk.AccAddressFromBech32(creator2)
	suite.Require().NoError(err)

	checkInterestCorrectness(suite, creatorAddr1, creatorAddr2, 0, convertBorrowToLocal(oneMonthInvestor1.Amount).String(), convertBorrowToLocal(oneMonthInvestor2.Amount).String())

	result1, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)
	suite.Require().True(result1.Amount == "0ausdc")

	resp, err = suite.keeper.OutstandingInterest(suite.ctx, &types.QueryOutstandingInterestRequest{
		Wallet:    creator1,
		PoolIndex: suite.investorPool,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(resp.Amount, oneMonthInvestor1.Amount.String())

	resp, err = suite.keeper.OutstandingInterest(suite.ctx, &types.QueryOutstandingInterestRequest{
		Wallet:    creator2,
		PoolIndex: suite.investorPool,
	})
	suite.Require().NoError(err)
	val, ok := sdkmath.NewIntFromString(resp.Amount)
	suite.Require().True(ok)
	suite.Require().True(checkValueWithRangeTwo(val, oneMonthInvestor2.Amount))

	dueTime := suite.ctx.BlockTime().Add(time.Second * time.Duration(poolInfo.PayFreq))
	// the correctness of the calculation is verified in  interest_test.go
	for {
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(200)))
		resp, err = suite.keeper.OutstandingInterest(suite.ctx, &types.QueryOutstandingInterestRequest{
			Wallet:    creator1,
			PoolIndex: suite.investorPool,
		})
		suite.Require().NoError(err)
		if suite.ctx.BlockTime().After(dueTime) {
			break
		}
	}
}

func (suite *claimInterestSuite) TestClaimInterestMultipleMonthWithSomePaymentMissing() {
	SetupPool(suite)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	poolInfo.PoolTotalBorrowLimit = 100
	poolInfo.TargetAmount = sdk.NewCoin("ausdc", sdkmath.NewInt(600000))
	suite.keeper.SetPool(suite.ctx, poolInfo)

	// now we deposit some token and it should be enough to borrow
	creator1 := suite.investors[0]
	creator2 := suite.investors[1]

	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(4e5))
	// suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{
		Creator:   creator1,
		PoolIndex: suite.investorPool,
		Token:     depositAmount,
	}

	// user two deposit half of the amount of the user 1
	msgDepositUser2 := &types.MsgDeposit{
		Creator:   creator2,
		PoolIndex: suite.investorPool,
		Token:     depositAmount.SubAmount(sdkmath.NewInt(2e5)),
	}

	_, err := suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	// now we borrow 1.34e5
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	reqInterest := types.MsgRepayInterest{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(8e9))}
	_, err = suite.app.RepayInterest(suite.ctx, &reqInterest)
	suite.Require().NoError(err)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(spvkeeper.OneMonth)))
	err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)
	suite.Require().NoError(err)

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(spvkeeper.OneMonth)))
	err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)
	suite.Require().NoError(err)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(spvkeeper.OneMonth)))
	err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)
	suite.Require().NoError(err)

	// now we borrow 1.34e5
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	// for the rest of 10 month
	for i := 0; i < 10; i++ {
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(spvkeeper.OneMonth)))
		err := suite.keeper.HandleInterest(suite.ctx, &poolInfo)
		suite.Require().NoError(err)
	}

	// now we delete some payment info to simulate the missing payment
	// classid := fmt.Sprintf("%v-0", poolInfo.PoolNFTIds)
	classInfo1, ok := suite.keeper.NftKeeper.GetClass(suite.ctx, poolInfo.PoolNFTIds[0])
	suite.Require().True(ok)

	var borrowInterest1 types.BorrowInterest
	err = proto.Unmarshal(classInfo1.Data.Value, &borrowInterest1)
	if err != nil {
		panic(err)
	}

	classInfo2, ok := suite.keeper.NftKeeper.GetClass(suite.ctx, poolInfo.PoolNFTIds[1])
	suite.Require().True(ok)

	var borrowInterest2 types.BorrowInterest
	err = proto.Unmarshal(classInfo2.Data.Value, &borrowInterest2)
	if err != nil {
		panic(err)
	}

	previousLength := len(borrowInterest1.Payments)
	previousLength2 := len(borrowInterest2.Payments)

	// now we delete 3 last payment
	borrowInterest1.Payments = borrowInterest1.Payments[:len(borrowInterest1.Payments)-3]
	require.Equal(suite.T(), previousLength-3, len(borrowInterest1.Payments))

	classInfo1.Data, err = types2.NewAnyWithValue(&borrowInterest1)
	suite.Require().NoError(err)
	err = suite.keeper.NftKeeper.UpdateClass(suite.ctx, classInfo1)
	suite.Require().NoError(err)

	escrowInterestAmount1 := poolInfo.EscrowInterestAmount

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(spvkeeper.OneMonth)))
	err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)
	suite.Require().NoError(err)

	classInfo1after, ok := suite.keeper.NftKeeper.GetClass(suite.ctx, poolInfo.PoolNFTIds[0])
	suite.Require().True(ok)

	var borrowInterest1after types.BorrowInterest
	err = proto.Unmarshal(classInfo1after.Data.Value, &borrowInterest1after)
	if err != nil {
		panic(err)
	}

	classInfo2after, ok := suite.keeper.NftKeeper.GetClass(suite.ctx, poolInfo.PoolNFTIds[1])
	suite.Require().True(ok)

	var borrowInterest2after types.BorrowInterest
	err = proto.Unmarshal(classInfo2after.Data.Value, &borrowInterest2after)
	if err != nil {
		panic(err)
	}

	suite.Require().Equal(len(borrowInterest1after.Payments), previousLength+1)
	suite.Require().Equal(len(borrowInterest2after.Payments), previousLength2+1)

	// and the amount should be the same as before deleted
	first := borrowInterest1after.Payments[1].PaymentAmount.Amount
	for _, el := range borrowInterest1after.Payments[1:] {
		suite.Require().True(first.Equal(el.PaymentAmount.Amount))
	}

	// as each payment is 1545, and we have 3 missing payment and 2 normal payment, so the escrow should be 1545*5
	suite.Require().True((escrowInterestAmount1.Sub(poolInfo.EscrowInterestAmount)).Equal(sdkmath.NewInt(1545 * 5)))
}
