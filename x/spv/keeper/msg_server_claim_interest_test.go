package keeper_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/gogo/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/suite"
)

const month = 3600 * 24 * 30

// Test suite used for all keeper tests
type claimInterestSuite struct {
	suite.Suite
	keeper       *spvkeeper.Keeper
	nftKeeper    types.NFTKeeper
	app          types.MsgServer
	ctx          sdk.Context
	investors    []string
	investorPool string
}

func SetupPool(suite *claimInterestSuite) {
	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 3, PoolName: "hello", Apy: "0.15", TargetTokenAmount: sdk.NewCoin("ausdc", sdk.NewInt(3*1e9))}
	resp, err := suite.app.CreatePool(suite.ctx, &req)
	suite.Require().NoError(err)

	depositorPool := resp.PoolIndex[0]

	suite.investorPool = depositorPool

	_, err = suite.app.ActivePool(suite.ctx, types.NewMsgActivePool("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", resp.PoolIndex[0]))
	suite.Require().NoError(err)

	_, err = suite.app.ActivePool(suite.ctx, types.NewMsgActivePool("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", resp.PoolIndex[1]))
	suite.Require().NoError(err)

	req2 := types.MsgAddInvestors{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: resp.PoolIndex[0],
		InvestorID: []string{"2"}}
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
	app, k, nftKeeper, wctx := setupMsgServer(suite.T())
	ctx := sdk.UnwrapSDKContext(wctx)
	// create the first pool apy 7.8%

	suite.ctx = ctx
	suite.keeper = k
	suite.app = app
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
	interestOneYearWithReserve := sdk.NewDecFromInt(lastBorrow.Amount).Mul(sdk.MustNewDecFromStr("0.15")).QuoInt64(12).TruncateInt()
	interestOneYear := interestOneYearWithReserve.Sub(sdk.NewDecFromInt(interestOneYearWithReserve).Mul(sdk.MustNewDecFromStr("0.15")).TruncateInt())

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

	ratio1 := sdk.NewDecFromInt(nftInfo1.Borrowed.Amount).Quo(sdk.NewDecFromInt(lastBorrow.Amount))
	ratio2 := sdk.NewDecFromInt(nftInfo2.Borrowed.Amount).Quo(sdk.NewDecFromInt(lastBorrow.Amount))

	user1Interest := sdk.NewDecFromInt(interestOneYear).Mul(ratio1).TruncateInt()
	user2Interest := sdk.NewDecFromInt(interestOneYear).Mul(ratio2).TruncateInt()

	fmt.Printf(">>>>%v\n", borrowClassInfo.InterestPaid)

	c1, err := sdk.ParseCoinsNormalized(expectedAmount1 + "ausdc")
	c2, err := sdk.ParseCoinsNormalized(expectedAmount2 + "ausdc")

	totalPaid := sdk.ZeroInt()
	for _, el := range borrowClassInfo.Payments {
		totalPaid = totalPaid.Add(el.PaymentAmount.Amount)
	}
	//totalPaid := borrowClassInfo.Payments[len(borrowClassInfo.Payments)-1].PaymentAmount.Amount

	delta := totalPaid.Sub(borrowClassInfo.InterestPaid.Amount)

	fmt.Printf(">>=----->>%v\n", delta)

	fmt.Printf(">>>>>%v---%v\n", expectedAmount1, user1Interest.String())
	fmt.Printf(">>>>>%v---%v\n", expectedAmount2, user2Interest.String())
	fmt.Printf(">>>>>333>>>%v\n", c1[0].Add(c2[0]).Amount.Sub(user1Interest.Add(user2Interest)).Abs())

	suite.Require().True(c1[0].Add(c2[0]).Amount.Sub(user1Interest.Add(user2Interest)).Abs().LTE(delta))

	suite.Require().True(c1[0].Amount.Sub(user1Interest).Abs().LTE(delta))
	suite.Require().True(c2[0].Amount.Sub(user2Interest).Abs().LTE(delta))

}

func (suite *claimInterestSuite) TestClaimInterestMultipleMonth() {

	investor1TotalClaimed := sdk.OneInt()
	SetupPool(suite)
	// now we deposit some token and it should be enough to borrow
	creator1 := suite.investors[0]
	creator2 := suite.investors[1]
	creatorAddr1, err := sdk.AccAddressFromBech32(creator1)
	suite.Require().NoError(err)
	creatorAddr2, err := sdk.AccAddressFromBech32(creator2)
	suite.Require().NoError(err)
	depositAmount := sdk.NewCoin("ausdc", sdk.NewInt(4e5))
	//suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{Creator: creator1,
		PoolIndex: suite.investorPool,
		Token:     depositAmount}

	// user two deposit half of the amount of the user 1
	msgDepositUser2 := &types.MsgDeposit{Creator: creator2,
		PoolIndex: suite.investorPool,
		Token:     depositAmount.SubAmount(sdk.NewInt(2e5))}

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1.34e5))}

	//now we borrow 1.34e5
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	depositor1, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	investor1Borrowd := depositor1.LockedAmount.Amount

	reqInterest := types.MsgRepayInterest{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdk.NewIntFromUint64(8e9))}
	_, err = suite.app.RepayInterest(suite.ctx, &reqInterest)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(month)))
	suite.keeper.HandleInterest(suite.ctx, &poolInfo)

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(month)))
	suite.keeper.HandleInterest(suite.ctx, &poolInfo)

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

	amount1 := a1.AmountOf("ausdc").Quo(sdk.NewInt(2))
	amount2 := a2.AmountOf("ausdc").Quo(sdk.NewInt(2))

	checkInterestCorrectness(suite, creatorAddr1, creatorAddr2, 0, amount1.String(), amount2.String())

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(month)))
	suite.keeper.HandleInterest(suite.ctx, &poolInfo)

	req.Creator = suite.investors[0]
	result1, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)

	req.Creator = suite.investors[1]
	result2, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)

	a1, _ = sdk.ParseCoinsNormalized(result1.Amount)
	investor1TotalClaimed = investor1TotalClaimed.Add(a1.AmountOf("ausdc"))
	a2, _ = sdk.ParseCoinsNormalized(result2.Amount)

	amount1 = a1.AmountOf("ausdc").Quo(sdk.NewInt(1))
	amount2 = a2.AmountOf("ausdc").Quo(sdk.NewInt(1))
	checkInterestCorrectness(suite, creatorAddr1, creatorAddr2, 0, amount1.String(), amount2.String())

	// for the rest of 9 month

	for i := 0; i < 9; i++ {
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(month)))
		suite.keeper.HandleInterest(suite.ctx, &poolInfo)
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

	amount1 = a1.AmountOf("ausdc").Quo(sdk.NewInt(9))
	amount2 = a2.AmountOf("ausdc").Quo(sdk.NewInt(9))
	checkInterestCorrectness(suite, creatorAddr1, creatorAddr2, 0, amount1.String(), amount2.String())

	expectedInterest := sdk.NewDecFromInt(investor1Borrowd).Mul(sdk.MustNewDecFromStr("0.15")).Mul(sdk.MustNewDecFromStr("0.85")).TruncateInt()
	fmt.Printf(">>>>>%v===%v\n", investor1TotalClaimed.String(), expectedInterest.String())
	suite.Require().True(expectedInterest.Equal(investor1TotalClaimed))
}

func (suite *claimInterestSuite) TestClaimInterestMultipleBorrow() {

	investor1TotalClaimed := sdk.OneInt()
	SetupPool(suite)
	// now we deposit some token and it should be enough to borrow
	creator1 := suite.investors[0]
	creator2 := suite.investors[1]
	creatorAddr1, err := sdk.AccAddressFromBech32(creator1)
	suite.Require().NoError(err)
	creatorAddr2, err := sdk.AccAddressFromBech32(creator2)
	suite.Require().NoError(err)
	depositAmount := sdk.NewCoin("ausdc", sdk.NewInt(4e5))
	//suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{Creator: creator1,
		PoolIndex: suite.investorPool,
		Token:     depositAmount}

	// user two deposit half of the amount of the user 1
	msgDepositUser2 := &types.MsgDeposit{Creator: creator2,
		PoolIndex: suite.investorPool,
		Token:     depositAmount.SubAmount(sdk.NewInt(2e5))}

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1.34e5))}

	//now we borrow 1.34e5
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	depositor1, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	reqInterest := types.MsgRepayInterest{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdk.NewIntFromUint64(8e9))}
	_, err = suite.app.RepayInterest(suite.ctx, &reqInterest)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(month)))
	suite.keeper.HandleInterest(suite.ctx, &poolInfo)

	req := types.MsgClaimInterest{
		Creator:   creator1,
		PoolIndex: suite.investorPool,
	}

	result1, err := suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)

	req.Creator = suite.investors[1]
	result2, err := suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)

	//interestOneYearWithReserve := sdk.NewDecFromInt(sdk.NewIntFromUint64(1.34e5)).Mul(sdk.MustNewDecFromStr("0.15")).QuoInt64(12).TruncateInt()
	//interestOneYear := interestOneYearWithReserve.Sub(sdk.NewDecFromInt(interestOneYearWithReserve).Mul(sdk.MustNewDecFromStr("0.15")).TruncateInt())

	a1, _ := sdk.ParseCoinsNormalized(result1.Amount)

	investor1TotalClaimed = investor1TotalClaimed.Add(a1.AmountOf("ausdc"))
	a2, _ := sdk.ParseCoinsNormalized(result2.Amount)

	amount1 := a1.AmountOf("ausdc").Quo(sdk.NewInt(1))
	amount2 := a2.AmountOf("ausdc").Quo(sdk.NewInt(1))

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

	checkInterestCorrectness(suite, creatorAddr1, creatorAddr2, 0, amount1.String(), amount2.String())

	// we borrow again after 15 days
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(3600*24*15)))
	borroweTime := suite.ctx.BlockTime()
	fmt.Printf(">.>>.biorrow time %v\n", borroweTime.String())

	//######now we borrow 2.1e5########
	borrow.BorrowAmount = sdk.NewCoin("ausdc", sdk.NewIntFromUint64(2.1e5))
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)
	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	// after another 16 days
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(3600*24*16)))
	suite.keeper.HandleInterest(suite.ctx, &poolInfo)

	// this payfreq is from the testutil spv.go
	payFreq := 2592000
	r := spvkeeper.CalculateInterestRate(sdk.MustNewDecFromStr("0.15"), payFreq)
	delta := suite.ctx.BlockTime().Truncate(time.Second * time.Duration(payFreq)).Sub(borroweTime).Seconds()

	interest := r.Power(uint64(delta)).Sub(sdk.OneDec())
	paymentAmount := interest.MulInt(borrow.BorrowAmount.Amount).TruncateInt()

	reservedAmount := sdk.NewDecFromInt(paymentAmount).Mul(sdk.MustNewDecFromStr("0.15")).TruncateInt()
	toInvestors := paymentAmount.Sub(reservedAmount)

	depositor1, found = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	depositor2, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)

	dat = strings.Split(depositor1.LinkedNFT[0], ":")
	nftUser1, found = suite.nftKeeper.GetNFT(suite.ctx, dat[0], dat[1])
	suite.Require().True(found)

	dat = strings.Split(depositor2.LinkedNFT[0], ":")
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

	Borrowclass, found = suite.nftKeeper.GetClass(suite.ctx, dat[0])
	suite.Require().True(found)

	var borrowClassInfo types.BorrowInterest
	err = proto.Unmarshal(Borrowclass.Data.Value, &borrowClassInfo)
	if err != nil {
		panic(err)
	}

	lastBorrow := borrowClassInfo.BorrowDetails[len(borrowClassInfo.BorrowDetails)-1].BorrowedAmount
	ratio1 := sdk.NewDecFromInt(nftInfo1.Borrowed.Amount).Quo(sdk.NewDecFromInt(lastBorrow.Amount))
	ratio2 := sdk.NewDecFromInt(nftInfo2.Borrowed.Amount).Quo(sdk.NewDecFromInt(lastBorrow.Amount))

	expectedToUser2 := sdk.NewDecFromInt(toInvestors).Mul(ratio2).TruncateInt()
	expectedToUser2 = expectedToUser2.Add(amount2)
	expectedToUser1 := sdk.NewDecFromInt(toInvestors).Mul(ratio1).TruncateInt()
	expectedToUser1 = expectedToUser1.Add(amount1)

	req.Creator = suite.investors[0]
	result1, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)

	req.Creator = suite.investors[1]
	result2, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)

	suite.Require().Equal(result1.Amount, expectedToUser1.String()+"ausdc")
	suite.Require().Equal(result2.Amount, expectedToUser2.String()+"ausdc")

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(month)))
	suite.keeper.HandleInterest(suite.ctx, &poolInfo)

	req.Creator = suite.investors[0]
	result1, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)
	req.Creator = suite.investors[1]
	result2, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)

	// for the new borrow, each month, it pays

	//interestOneYearWithReserve2 := sdk.NewDecFromInt(sdk.NewIntFromUint64(1.34e5)).Mul(sdk.MustNewDecFromStr("0.15")).QuoInt64(12).TruncateInt()
	//interestOneYear2 := interestOneYearWithReserve.Sub(sdk.NewDecFromInt(interestOneYearWithReserve).Mul(sdk.MustNewDecFromStr("0.15")).TruncateInt())

	amount1New, _ := sdk.ParseCoinsNormalized(result1.Amount)
	amount2New, _ := sdk.ParseCoinsNormalized(result2.Amount)

	newIncome1 := amount1New[0].SubAmount(amount1)
	newIncome2 := amount2New[0].SubAmount(amount2)

	checkInterestCorrectness(suite, creatorAddr1, creatorAddr2, 1, newIncome1.Amount.String(), newIncome2.Amount.String())

	for i := 0; i < 9; i++ {
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(month)))
		suite.keeper.HandleInterest(suite.ctx, &poolInfo)
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
	suite.Require().True(a1[0].Amount.Equal(amount1New[0].Amount.MulRaw(9)))
	suite.Require().True(a2[0].Amount.Equal(amount2New[0].Amount.MulRaw(9)))

}

func (suite *claimInterestSuite) TestClaimInterest() {

	SetupPool(suite)
	// now we deposit some token and it should be enough to borrow
	creator1 := suite.investors[0]
	creator2 := suite.investors[1]
	creatorAddr1, err := sdk.AccAddressFromBech32(creator1)
	suite.Require().NoError(err)
	creatorAddr2, err := sdk.AccAddressFromBech32(creator2)
	suite.Require().NoError(err)
	depositAmount := sdk.NewCoin("ausdc", sdk.NewInt(4e5))
	//suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{Creator: creator1,
		PoolIndex: suite.investorPool,
		Token:     depositAmount}

	// user two deposit half of the amount of the user 1
	msgDepositUser2 := &types.MsgDeposit{Creator: creator2,
		PoolIndex: suite.investorPool,
		Token:     depositAmount.SubAmount(sdk.NewInt(2e5))}

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1.34e5))}

	//now we borrow 1.34e5
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	reqInterest := types.MsgRepayInterest{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdk.NewIntFromUint64(8e9))}
	_, err = suite.app.RepayInterest(suite.ctx, &reqInterest)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(month)))
	suite.keeper.HandleInterest(suite.ctx, &poolInfo)

	req := types.MsgClaimInterest{
		Creator:   creator1,
		PoolIndex: suite.investorPool,
	}

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(month)))
	result1, err := suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)

	req.Creator = suite.investors[1]
	result2, err := suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)
	// total borrow is 1.34e5 and the user1 offer 4e5 while the second user offer 2e5, reserve is 0.15, the apy is 0.15
	// after the first month, user 1 get 1.34e5*0.85*0.15/12*4/6
	//interestOneYearWithReserve := sdk.NewDecFromInt(sdk.NewIntFromUint64(1.34e5)).Mul(sdk.MustNewDecFromStr("0.15")).QuoInt64(12).TruncateInt()
	//interestOneYear := interestOneYearWithReserve.Sub(sdk.NewDecFromInt(interestOneYearWithReserve).Mul(sdk.MustNewDecFromStr("0.15")).TruncateInt())

	a1, _ := sdk.ParseCoinsNormalized(result1.Amount)
	a2, _ := sdk.ParseCoinsNormalized(result2.Amount)

	amount1 := a1.AmountOf("ausdc").Quo(sdk.NewInt(1))
	amount2 := a2.AmountOf("ausdc").Quo(sdk.NewInt(1))

	checkInterestCorrectness(suite, creatorAddr1, creatorAddr2, 0, amount1.String(), amount2.String())

	result1, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)
	suite.Require().True(result1.Amount == "0ausdc")

	// we add one second after the withdraw the interest

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(100)))
	result1, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)
	suite.Require().True(result1.Amount == "0ausdc")

	// we add less than a month, so the amount should still be zero
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(month-101)))
	result1, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)
	suite.Require().True(result1.Amount == "0ausdc")

	// since the spv not paid the interest, we cannont claim the interest
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(month+150)))
	result1, err = suite.app.ClaimInterest(suite.ctx, &req)
	suite.Require().NoError(err)
	suite.Require().True(result1.Amount == "0ausdc")

	// we check the depositor info
	//depositor, found := suite.keeper.GetDepositor(suite.ctx, depositorPool, creatorAddr1)
	//suite.Require().True(found)
	//targetDepositor := types.DepositorInfo{
	//	InvestorId:       "2",
	//	DepositorAddress: creatorAddr1,
	//	PoolIndex:        depositorPool,
	//	LockedAmount:     sdk.NewCoin("ausdc", sdk.ZeroInt()),
	//	WithdrawalAmount: depositAmount,
	//	LinkedNFT:        []string{},
	//}
	//
	//compareDepositor(suite.Suite, targetDepositor, depositor)
	//// we deposit again,so withdrawal is doubled
	//
	//_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	//suite.Require().NoError(err)
	//
	//targetDepositor.WithdrawalAmount = targetDepositor.WithdrawalAmount.Add(depositAmount)
	//
	//depositor, found = suite.keeper.GetDepositor(suite.ctx, depositorPool, creatorAddr1)
	//suite.Require().True(found)
	//compareDepositor(suite.Suite, targetDepositor, depositor)
	//
	//// we mock the second user deposits the token, now we have 3*4e5 tokens
	////_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	////suite.Require().NoError(err)
	//
	//pool, found := suite.keeper.GetPools(suite.ctx, depositorPool)
	//suite.Require().True(found)
	//
	//totalBorrowable := msgDepositUser1.Token.Add(msgDepositUser1.Token).Add(msgDepositUser2.Token)
	//suite.Require().True(totalBorrowable.IsEqual(pool.UsableAmount))
	//
	//user1Ratio := sdk.NewDecFromInt(msgDepositUser1.Token.Amount.Mul(sdk.NewInt(2))).Quo(sdk.NewDecFromInt(totalBorrowable.Amount))
	//
	//user2Ratio := sdk.NewDecFromInt(msgDepositUser2.Token.Amount).Quo(sdk.NewDecFromInt(totalBorrowable.Amount))
	//
	////now we borrow 2e5
	//_, err = suite.app.Borrow(suite.ctx, borrow)
	//suite.Require().NoError(err)
	//
	//p1, found := suite.keeper.GetDepositor(suite.ctx, depositorPool, creatorAddr1)
	//suite.Require().True(found)
	//
	//p2, found := suite.keeper.GetDepositor(suite.ctx, depositorPool, creatorAddr2)
	//suite.Require().True(found)
	//
	//poolNow, found := suite.keeper.GetPools(suite.ctx, depositorPool)
	//suite.Require().True(found)
	//suite.Require().True(poolNow.BorrowedAmount.IsEqual(borrow.BorrowAmount))
	//suite.Require().True(totalBorrowable.Sub(borrow.BorrowAmount).IsEqual(poolNow.UsableAmount))
	//
	//borrowedFromUser1 := sdk.NewDecFromInt(borrow.BorrowAmount.Amount).Mul(user1Ratio).TruncateInt()
	//borrowedFromUser2 := borrow.BorrowAmount.Amount.Sub(borrowedFromUser1)
	//
	//borrowedFromUser2Ratio := sdk.NewDecFromInt(borrow.BorrowAmount.Amount).Mul(user2Ratio).TruncateInt()
	//
	//suite.Require().True(p1.LockedAmount.Amount.Equal(borrowedFromUser1))
	//suite.Require().True(p2.LockedAmount.Amount.Equal(borrowedFromUser2))
	//suite.Require().True(borrowedFromUser2Ratio.Equal(borrowedFromUser2))
	//
	//// total amount shoube be locked+withdrawable
	//suite.Require().True(p1.LockedAmount.Add(p1.WithdrawalAmount).IsEqual(msgDepositUser1.Token.Add(msgDepositUser1.Token)))
	//
	//nftUser1 := p1.LinkedNFT[0]
	//nftUser2 := p2.LinkedNFT[0]
	//
	//nftClassID := fmt.Sprintf("nft-%v-0", depositorPool[2:])
	//nft, found := suite.nftKeeper.GetClass(suite.ctx, nftClassID)
	//suite.Require().True(found)
	//
	//var borrowClassInfo types.BorrowInterest
	//err = proto.Unmarshal(nft.Data.Value, &borrowClassInfo)
	//if err != nil {
	//	panic(err)
	//}
	//
	//suite.True(borrowClassInfo.Borrowed.IsEqual(borrow.BorrowAmount))
	//suite.True(borrowClassInfo.BorrowedLast.IsEqual(borrow.BorrowAmount))
	//fmt.Printf(">>>>%v\n", borrowClassInfo.Apy)
	//suite.True(borrowClassInfo.Apy.Equal(sdk.NewDecWithPrec(15, 2)))
	//
	//// nft ID is the hash(nft class ID, investorWallet)
	//indexHash := crypto.Keccak256Hash([]byte(nftClassID), p1.DepositorAddress)
	//expectedID1 := fmt.Sprintf("%v:invoice-%v", nftClassID, indexHash.String()[2:])
	//suite.Require().Equal(nftUser1, expectedID1)
	//
	//indexHash = crypto.Keccak256Hash([]byte(nftClassID), p2.DepositorAddress)
	//expectedID2 := fmt.Sprintf("%v:invoice-%v", nftClassID, indexHash.String()[2:])
	//suite.Require().Equal(nftUser2, expectedID2)
	//
	//dat := strings.Split(nftUser1, ":")
	//nft1, found := suite.nftKeeper.GetNFT(suite.ctx, dat[0], dat[1])
	//suite.Require().True(found)
	//
	//var nftInfo types.NftInfo
	//err = proto.Unmarshal(nft1.Data.Value, &nftInfo)
	//if err != nil {
	//	panic(err)
	//}
	//suite.Require().True(nftInfo.Ratio.Equal(user1Ratio))
	//
	//// now, user 2 deposits more money and then, the spv borrow more. the ratio should  be changed.
	//_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	//suite.Require().NoError(err)
	//
	//user1Cached := p1.WithdrawalAmount
	//user2Cached := p2.WithdrawalAmount
	//
	//p2, found = suite.keeper.GetDepositor(suite.ctx, depositorPool, creatorAddr2)
	//
	//suite.Require().True(found)
	//suite.Require().True(p2.WithdrawalAmount.IsEqual(user2Cached.AddAmount(msgDepositUser2.Token.Amount)))
	//suite.Require().True(poolNow.GetUsableAmount().IsEqual(user1Cached.Add(user2Cached)))
	//user2Cached = p2.WithdrawalAmount
	//newuser1Ratio1 := sdk.NewDecFromInt(user1Cached.Amount).Quo(sdk.NewDecFromInt(user1Cached.Add(user2Cached).Amount))
	//
	//poolAfterUser2SecondDeposit, found := suite.keeper.GetPools(suite.ctx, depositorPool)
	//suite.Require().True(found)
	//suite.Require().True(poolAfterUser2SecondDeposit.UsableAmount.IsEqual(user2Cached.Add(user1Cached)))
	//
	//// NOW we borrow
	//borrow.BorrowAmount = sdk.NewCoin(borrow.BorrowAmount.Denom, sdk.NewInt(1.2e5))
	//_, err = suite.app.Borrow(suite.ctx, borrow)
	//suite.Require().NoError(err)
	//
	//previousAmountBorrowed := poolNow.BorrowedAmount
	//previousBorrowAble := poolNow.UsableAmount
	//poolNow, found = suite.keeper.GetPools(suite.ctx, depositorPool)
	//suite.Require().True(found)
	//
	//suite.Require().True(poolNow.BorrowedAmount.Equal(borrow.BorrowAmount.AddAmount(previousAmountBorrowed.Amount)))
	//
	//suite.Require().True(poolNow.UsableAmount.Equal(previousBorrowAble.Add(msgDepositUser2.Token).Sub(borrow.BorrowAmount)))
	//
	//beforeLockedAmount := p1.LockedAmount
	//// now we check the nfts
	//p1, found = suite.keeper.GetDepositor(suite.ctx, depositorPool, creatorAddr1)
	//suite.Require().True(found)
	//
	//p2, found = suite.keeper.GetDepositor(suite.ctx, depositorPool, creatorAddr2)
	//suite.Require().True(found)
	//
	//lockedThistime := p1.LockedAmount.Sub(beforeLockedAmount)
	//shouldLocked := newuser1Ratio1.Mul(sdk.NewDecFromInt(borrow.BorrowAmount.Amount)).TruncateInt()
	//suite.Require().True(lockedThistime.Amount.Equal(shouldLocked))
	//
	//// we check the total deposit of the user1 is correct
	//suite.Require().True(p1.LockedAmount.Add(p1.WithdrawalAmount).IsEqual(msgDepositUser1.Token.Add(msgDepositUser1.Token)))
	//
	//nft2User1 := p1.LinkedNFT[1]
	//nft2User2 := p2.LinkedNFT[1]
	//
	//nftClassID = fmt.Sprintf("nft-%v-1", depositorPool[2:])
	//nft, found = suite.nftKeeper.GetClass(suite.ctx, nftClassID)
	//suite.Require().True(found)
	//
	//err = proto.Unmarshal(nft.Data.Value, &borrowClassInfo)
	//if err != nil {
	//	panic(err)
	//}
	//
	//suite.True(borrowClassInfo.Borrowed.IsEqual(borrow.BorrowAmount))
	//suite.True(borrowClassInfo.BorrowedLast.IsEqual(borrow.BorrowAmount))
	//
	////nft ID is the hash(nft class ID, investorWallet)
	//indexHash = crypto.Keccak256Hash([]byte(nftClassID), p1.DepositorAddress)
	//expectedID1 = fmt.Sprintf("%v:invoice-%v", nftClassID, indexHash.String()[2:])
	//suite.Require().Equal(nft2User1, expectedID1)
	//
	//indexHash = crypto.Keccak256Hash([]byte(nftClassID), p2.DepositorAddress)
	//expectedID2 = fmt.Sprintf("%v:invoice-%v", nftClassID, indexHash.String()[2:])
	//suite.Require().Equal(nft2User2, expectedID2)
	//
	//dat = strings.Split(nft2User1, ":")
	//nft1, found = suite.nftKeeper.GetNFT(suite.ctx, dat[0], dat[1])
	//suite.Require().True(found)
	//
	//err = proto.Unmarshal(nft1.Data.Value, &nftInfo)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// this calculates the ratio that user1 contribute to this borrow
	//shoudRatio := sdk.NewDecFromInt(shouldLocked).QuoInt(borrow.BorrowAmount.Amount)
	//suite.Require().True(nftInfo.Ratio.Equal(shoudRatio))

}

func (suite *claimInterestSuite) TestCalimInterestNoAuthorized() {

	SetupPool(suite)
	// now we deposit some token and it should be enough to borrow
	creator1 := suite.investors[0]
	creator2 := suite.investors[1]
	//creatorAddr1, err := sdk.AccAddressFromBech32(creator1)
	//suite.Require().NoError(err)
	//creatorAddr2, err := sdk.AccAddressFromBech32(creator2)
	//suite.Require().NoError(err)
	depositAmount := sdk.NewCoin("ausdc", sdk.NewInt(4e5))
	//suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{Creator: creator1,
		PoolIndex: suite.investorPool,
		Token:     depositAmount}

	// user two deposit half of the amount of the user 1
	msgDepositUser2 := &types.MsgDeposit{Creator: creator2,
		PoolIndex: suite.investorPool,
		Token:     depositAmount.SubAmount(sdk.NewInt(2e5))}

	_, err := suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1.34e5))}

	//now we borrow 1.34e5
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	reqInterest := types.MsgRepayInterest{"jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", suite.investorPool, sdk.NewCoin("ausdc", sdk.NewIntFromUint64(8e9))}
	_, err = suite.app.RepayInterest(suite.ctx, &reqInterest)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	month := 3600 * 24 * 30

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(month)))
	suite.keeper.HandleInterest(suite.ctx, &poolInfo)

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
