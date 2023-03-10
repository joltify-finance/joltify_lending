package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/suite"
)

// Test suite used for all keeper tests
type withDrawPrincipalSuite struct {
	suite.Suite
	keeper       *spvkeeper.Keeper
	nftKeeper    types.NFTKeeper
	app          types.MsgServer
	ctx          sdk.Context
	investors    []string
	investorPool string
}

func setupPool(suite *withDrawPrincipalSuite) {
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
func (suite *withDrawPrincipalSuite) SetupTest() {

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
func TestWithdrawPrincipalInterest(t *testing.T) {
	suite.Run(t, new(withDrawPrincipalSuite))
}

func (suite *withDrawPrincipalSuite) TestMsgWithdrawPrincipalTest() {

	setupPool(suite)
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
	totalBorrowed := borrow.BorrowAmount

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.Require().EqualValues(borrow.BorrowAmount.Amount, poolInfo.BorrowedAmount.Amount)

	borrowable := sdk.NewIntFromUint64(6e5).Sub(sdk.NewIntFromUint64(1.34e5))
	suite.Require().EqualValues(borrowable, poolInfo.BorrowableAmount.Amount)
	suite.Require().EqualValues(poolInfo.BorrowedAmount.Amount, sdk.NewIntFromUint64(1.34e5))

	depositor1, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	depositor2, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)

	totalWithdrawbleFromInvestor := depositor1.WithdrawalAmount.Add(depositor2.WithdrawalAmount)
	suite.Require().EqualValues(totalWithdrawbleFromInvestor.Amount, poolInfo.BorrowableAmount.Amount)

	suite.Require().EqualValues(sdk.NewInt(4e5), depositor1.LockedAmount.Add(depositor1.WithdrawalAmount).Amount)
	suite.Require().EqualValues(sdk.NewInt(2e5), depositor2.LockedAmount.Add(depositor2.WithdrawalAmount).Amount)

	getRatio := sdk.NewDecFromInt(depositor1.LockedAmount.Amount).Quo(sdk.NewDecFromInt(depositor2.LockedAmount.Amount))
	// the ratio should close to 2
	suite.Require().True(getRatio.Sub(sdk.MustNewDecFromStr("2")).LTE(sdk.NewDecWithPrec(1, 4)))

	// now we borrow more money
	borrow.BorrowAmount = sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1.1e5))
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)
	totalBorrowed = totalBorrowed.Add(borrow.BorrowAmount)

	borrowable = borrowable.Sub(borrow.BorrowAmount.Amount)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.Require().EqualValues(borrowable, poolInfo.BorrowableAmount.Amount)

	depositor1, found = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	depositor2, found = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)
	totalWithdrawbleFromInvestor = depositor1.WithdrawalAmount.Add(depositor2.WithdrawalAmount)
	suite.Require().EqualValues(totalWithdrawbleFromInvestor.Amount, poolInfo.BorrowableAmount.Amount)
	suite.Require().EqualValues(sdk.NewInt(4e5), depositor1.LockedAmount.Add(depositor1.WithdrawalAmount).Amount)
	suite.Require().EqualValues(sdk.NewInt(2e5), depositor2.LockedAmount.Add(depositor2.WithdrawalAmount).Amount)

	getRatio = sdk.NewDecFromInt(depositor1.LockedAmount.Amount).Quo(sdk.NewDecFromInt(depositor2.LockedAmount.Amount))
	// the ratio should close to 2
	suite.Require().True(getRatio.Sub(sdk.MustNewDecFromStr("2")).LTE(sdk.NewDecWithPrec(1, 4)))

	// now we run the withdrawal
	withdrawReq := types.MsgWithdrawPrincipal{Creator: "invalid"}
	_, err = suite.app.WithdrawPrincipal(suite.ctx, &withdrawReq)
	suite.Require().ErrorContains(err, "invalid address")

	withdrawReq = types.MsgWithdrawPrincipal{Creator: suite.investors[0], PoolIndex: "invalid"}
	_, err = suite.app.WithdrawPrincipal(suite.ctx, &withdrawReq)
	suite.Require().ErrorContains(err, "depositor not found for pool")

	withdrawReq.PoolIndex = suite.investorPool
	withdrawReq.Token = sdk.NewCoin("invalid", sdk.NewIntFromUint64(22))
	_, err = suite.app.WithdrawPrincipal(suite.ctx, &withdrawReq)
	suite.Require().ErrorContains(err, "you can only withdraw ausdc")

	fmt.Printf(">>>>>>>%v\n", depositor2.WithdrawalAmount)
	fmt.Printf(">>>>>>>%v\n", depositor2.LockedAmount)

	getRatio = sdk.NewDecFromInt(depositor1.LockedAmount.Amount).Quo(sdk.NewDecFromInt(depositor2.LockedAmount.Amount))
	// the ratio should close to 2
	suite.Require().True(getRatio.Sub(sdk.MustNewDecFromStr("2")).LTE(sdk.NewDecWithPrec(1, 4)))

	//withdrawReq.Token = depositor2.WithdrawalAmount.AddAmount(sdk.NewIntFromU)
	//_, err = suite.app.WithdrawPrincipal(suite.ctx, &withdrawReq)
	//suite.Require().ErrorContains(err, "you can only withdraw ausdc")

}

func (suite *withDrawPrincipalSuite) TestWithdrawWithSPVBorrowAndRepay() {

	// skip the test now
	suite.T().Skip("we skip the test now")
	setupPool(suite)
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

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.Require().EqualValues(borrow.BorrowAmount.Amount, poolInfo.BorrowedAmount.Amount)

	borrowable := sdk.NewIntFromUint64(6e5).Sub(sdk.NewIntFromUint64(1.34e5))
	suite.Require().EqualValues(borrowable, poolInfo.BorrowableAmount.Amount)
	suite.Require().EqualValues(poolInfo.BorrowedAmount.Amount, sdk.NewIntFromUint64(1.34e5))

	depositor1, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	depositor2, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)

	totalWithdrawbleFromInvestor := depositor1.WithdrawalAmount.Add(depositor2.WithdrawalAmount)
	suite.Require().EqualValues(totalWithdrawbleFromInvestor.Amount, poolInfo.BorrowableAmount.Amount)

	suite.Require().EqualValues(sdk.NewInt(4e5), depositor1.LockedAmount.Add(depositor1.WithdrawalAmount).Amount)
	suite.Require().EqualValues(sdk.NewInt(2e5), depositor2.LockedAmount.Add(depositor2.WithdrawalAmount).Amount)

	firstWithdrawable := depositor2.WithdrawalAmount
	firstWithdrawable1 := depositor1.WithdrawalAmount

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	// now we pay the principal
	req := types.MsgPayPrincipal{
		Creator:   "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0",
		PoolIndex: suite.investorPool,
		Token:     sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1e5)),
	}
	_, err = suite.app.PayPrincipal(suite.ctx, &req)
	suite.Require().NoError(err)
	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.keeper.HandlePrincipalPayment(suite.ctx, &poolInfo)
	poolInfoNew, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	withdraw := types.MsgWithdrawPrincipal{
		Creator:   suite.investors[1],
		PoolIndex: suite.investorPool,
		Token:     sdk.NewCoin("ausdc", sdk.ZeroInt()),
	}
	_, err = suite.app.WithdrawPrincipal(suite.ctx, &withdraw)
	suite.Require().NoError(err)
	depositor2, found = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)
	delta := depositor2.WithdrawalAmount.Sub(firstWithdrawable)

	totalMoney2 := depositor2.WithdrawalAmount.Add(depositor2.LockedAmount)
	suite.Require().EqualValues(sdk.NewIntFromUint64(2e5), totalMoney2.Amount)

	withdraw.Creator = suite.investors[0]
	_, err = suite.app.WithdrawPrincipal(suite.ctx, &withdraw)
	suite.Require().NoError(err)

	depositor1, found = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)
	delta2 := depositor1.WithdrawalAmount.Sub(firstWithdrawable1)

	totalMoney1 := depositor1.WithdrawalAmount.Add(depositor1.LockedAmount)
	suite.Require().EqualValues(sdk.NewIntFromUint64(4e5), totalMoney1.Amount)

	allWithdraw := delta.Add(delta2)
	suite.Require().True(allWithdraw.SubAmount(sdk.NewIntFromUint64(1e5)).Amount.LT(sdk.NewIntFromUint64(2)))

	// we have another two borrows

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	borrow.BorrowAmount = sdk.NewCoin("ausdc", sdk.NewIntFromUint64(4e3))
	//now we borrow 4e3
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	borrow.BorrowAmount = sdk.NewCoin("ausdc", sdk.NewIntFromUint64(2e3))
	//now we borrow 2e3
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	// now we pay the principal
	req = types.MsgPayPrincipal{
		Creator:   "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0",
		PoolIndex: suite.investorPool,
		Token:     sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1e2)),
	}
	_, err = suite.app.PayPrincipal(suite.ctx, &req)
	suite.Require().NoError(err)

	poolInfoNew, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.keeper.HandlePrincipalPayment(suite.ctx, &poolInfoNew)
	poolInfoNew, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	nfts := poolInfoNew.PoolNFTIds[1]
	class, found := suite.nftKeeper.GetClass(suite.ctx, nfts)
	if !found {
		panic(found)
	}
	var borrowInterest types.BorrowInterest
	err = proto.Unmarshal(class.Data.Value, &borrowInterest)
	if err != nil {
		panic(err)
	}

	// now we verify the depositor 1 and 2
	depositorBeforeWithdraw2, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)
	depositorBeforeWithdraw1, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	withdraw.Creator = suite.investors[1]
	_, err = suite.app.WithdrawPrincipal(suite.ctx, &withdraw)
	suite.Require().NoError(err)

	depositor2, found = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)

	delta = depositor2.WithdrawalAmount.Sub(depositorBeforeWithdraw2.WithdrawalAmount)

	totalMoney2 = depositor2.WithdrawalAmount.Add(depositor2.LockedAmount)
	suite.Require().EqualValues(sdk.NewIntFromUint64(2e5), totalMoney2.Amount)

	withdraw.Creator = suite.investors[0]
	_, err = suite.app.WithdrawPrincipal(suite.ctx, &withdraw)
	suite.Require().NoError(err)

	depositor1, found = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)
	delta2 = depositor1.WithdrawalAmount.Sub(depositorBeforeWithdraw1.WithdrawalAmount)

	totalMoney1 = depositor1.WithdrawalAmount.Add(depositor1.LockedAmount)
	suite.Require().EqualValues(sdk.NewIntFromUint64(4e5), totalMoney1.Amount)

	allWithdrawAbs := delta.Add(delta2).Amount.Abs()
	// as we have two investors, so max of each extra pay is 1, in total is 2 at most
	suite.Require().True(allWithdrawAbs.Sub(sdk.NewIntFromUint64(1e2)).LT(sdk.NewIntFromUint64(3)))

	poolInfoBeforePayAll, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	//
	//
	//
	// now we pay all the money
	req = types.MsgPayPrincipal{
		Creator:   "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0",
		PoolIndex: suite.investorPool,
		Token:     sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1e10)),
	}
	_, err = suite.app.PayPrincipal(suite.ctx, &req)
	suite.Require().NoError(err)

	poolInfoNew, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.keeper.HandlePrincipalPayment(suite.ctx, &poolInfoNew)
	poolInfoNew, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.Require().True(poolInfoNew.EscrowPrincipalAmount.Amount.Equal(sdk.NewIntFromUint64(1e10).Sub(poolInfoBeforePayAll.BorrowedAmount.Amount)))
	paid := poolInfoBeforePayAll.BorrowedAmount

	// now we verify the depositor 1 and 2

	depositorBeforeWithdraw2, found = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)

	depositorBeforeWithdraw1, found = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	fmt.Printf(">>>>WW2222WWWW>>>%v\n", depositorBeforeWithdraw2.WithdrawalAmount)
	fmt.Printf(">>>>WW2222WWWW>>>%v\n", depositorBeforeWithdraw1.WithdrawalAmount)

	withdraw.Creator = suite.investors[1]
	_, err = suite.app.WithdrawPrincipal(suite.ctx, &withdraw)
	suite.Require().NoError(err)

	depositor2, found = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)

	delta = depositor2.WithdrawalAmount.Sub(depositorBeforeWithdraw2.WithdrawalAmount)
	fmt.Printf(">>>>delta1>>>%v\n", delta)

	fmt.Printf(">>>>>>>>>>%v-----%v\n", depositor2.LockedAmount, depositor2.WithdrawalAmount)
	suite.Require().True(depositor2.LockedAmount.Amount.Equal(sdk.ZeroInt()))
	// the original v2 deposit is 2e5
	suite.Require().True(depositor2.WithdrawalAmount.Amount.Equal(sdk.NewIntFromUint64(2e5)))

	totalMoney2 = depositor2.WithdrawalAmount.Add(depositor2.LockedAmount)
	suite.Require().EqualValues(sdk.NewIntFromUint64(2e5), totalMoney2.Amount)

	withdraw.Creator = suite.investors[0]
	_, err = suite.app.WithdrawPrincipal(suite.ctx, &withdraw)
	suite.Require().NoError(err)

	depositor1, found = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)
	delta2 = depositor1.WithdrawalAmount.Sub(depositorBeforeWithdraw1.WithdrawalAmount)

	suite.Require().True(depositor1.LockedAmount.Amount.Equal(sdk.ZeroInt()))
	// the original v2 deposit is 2e5
	suite.Require().True(depositor1.WithdrawalAmount.Amount.Equal(sdk.NewIntFromUint64(4e5)))

	totalMoney1 = depositor1.WithdrawalAmount.Add(depositor1.LockedAmount)
	suite.Require().EqualValues(sdk.NewIntFromUint64(4e5), totalMoney1.Amount)

	allWithdraw = delta.Add(delta2)
	suite.Require().True(allWithdraw.Amount.Sub(paid.Amount).Abs().LT(sdk.NewIntFromUint64(4)))

}
