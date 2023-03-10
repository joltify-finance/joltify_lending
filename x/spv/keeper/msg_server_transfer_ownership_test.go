package keeper_test

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (suite *withDrawPrincipalSuite) TestTransferOwnershipOneInvestor() {

	setupPool(suite)
	// now we deposit some token and it should be enough to borrow
	creator1 := suite.investors[0]
	creator2 := suite.investors[1]
	creatorAddr1, err := sdk.AccAddressFromBech32(creator1)
	suite.Require().NoError(err)
	//creatorAddr2, err := sdk.AccAddressFromBech32(creator2)
	//suite.Require().NoError(err)
	depositAmount := sdk.NewCoin("ausdc", sdk.NewInt(4e5))
	depositAmount2 := sdk.NewCoin("ausdc", sdk.NewInt(2e5))

	//suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{Creator: creator1,
		PoolIndex: suite.investorPool,
		Token:     depositAmount}

	// user two deposit half of the amount of the user 1
	msgDepositUser2 := &types.MsgDeposit{Creator: creator2,
		PoolIndex: suite.investorPool,
		Token:     depositAmount2}

	_ = msgDepositUser2
	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	//_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	//suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1.34e5))}

	//now we borrow 1.34e5
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * 20))
	_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{suite.investors[1], suite.investorPool, sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1e9))})
	suite.Require().NoError(err)

	reqOwner := types.MsgTransferOwnership{Creator: suite.investors[0], PoolIndex: suite.investorPool}
	_, err = suite.app.TransferOwnership(suite.ctx, &reqOwner)
	suite.Require().NoError(err)

	poolInfoBefore, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Duration(poolInfo.PayFreq) * time.Second))

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	//err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)
	suite.Require().NoError(err)
	suite.keeper.HandleTransfer(suite.ctx, &poolInfo)

	fmt.Printf(">>>>>>>>>%v\n", poolInfoBefore.PoolNFTIds)

	depositor, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)
	_ = depositor

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.Require().True(poolInfoBefore.BorrowedAmount.Equal(poolInfoBefore.BorrowedAmount))
	suite.Require().True(poolInfoBefore.BorrowableAmount.Equal(poolInfoBefore.BorrowableAmount))
	fmt.Printf(">>>>>>>>>%v\n", poolInfo.BorrowableAmount)
	fmt.Printf(">>>>>>>>>%v\n", poolInfo.BorrowedAmount)
	fmt.Printf(">>>>>>>>>%v\n", poolInfo.PoolNFTIds)

	borrowed := sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1.34e5))
	all1 := depositAmount
	suite.Require().True(depositor.LockedAmount.IsEqual(borrowed))
	fmt.Printf(">>>>%v\n", depositor.WithdrawalAmount)
	fmt.Printf(">>>>%v\n", all1.Sub(borrowed))
	suite.Require().True(depositor.WithdrawalAmount.IsEqual(all1.Sub(borrowed)))

	suite.nftKeeper.GetNFT()

}
