package keeper_test

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/davecgh/go-spew/spew"
	"github.com/gogo/protobuf/proto"
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

	depositorBefore, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	suite.keeper.HandleTransfer(suite.ctx, &poolInfo)

	depositorAfterTransfer, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.Require().True(poolInfoBefore.BorrowedAmount.Equal(poolInfo.BorrowedAmount))
	suite.Require().True(poolInfoBefore.BorrowableAmount.Equal(poolInfo.BorrowableAmount))
	fmt.Printf(">>>>>>>>>%v\n", poolInfo.BorrowableAmount)
	fmt.Printf(">>>>>>>>>%v\n", poolInfo.BorrowedAmount)
	fmt.Printf(">>>>>>>>>%v\n", poolInfo.PoolNFTIds)

	//fixme need to check the interest
	borrowed := sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1.34e5))
	all1 := depositAmount
	suite.Require().True(depositorAfterTransfer.LockedAmount.IsEqual(borrowed))
	suite.Require().Equal(depositorAfterTransfer.DepositType, types.DepositorInfo_processed)
	suite.Require().True(depositorAfterTransfer.WithdrawalAmount.IsEqual(all1.Sub(borrowed)))

	ids := strings.Split(depositorBefore.LinkedNFT[0], ":")
	_, found = suite.nftKeeper.GetNFT(suite.ctx, ids[0], ids[1])
	suite.Require().False(found)

	ids = strings.Split(depositorAfterTransfer.LinkedNFT[0], ":")
	nft1, found := suite.nftKeeper.GetNFT(suite.ctx, ids[0], ids[1])
	suite.Require().True(found)

	var nftInfo types.NftInfo
	err = proto.Unmarshal(nft1.Data.Value, &nftInfo)
	if err != nil {
		panic(err)
	}
	suite.Require().True(nftInfo.Borrowed.IsEqual(borrowed))

	// now we deposit more
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Hour * 24))
	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	depositorAfterDepositAgain, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)
	spew.Dump(depositorAfterDepositAgain)
	spew.Dump(depositorBefore)
	suite.Require().True(depositorAfterDepositAgain.GetWithdrawalAmount().Sub(depositorBefore.WithdrawalAmount).IsEqual(msgDepositUser1.Token))
}

func (suite *withDrawPrincipalSuite) TestTransferOwnershipTwoInvestor() {

	setupPool(suite)
	// now we deposit some token and it should be enough to borrow
	creator1 := suite.investors[0]
	creator2 := suite.investors[1]
	creatorAddr1, err := sdk.AccAddressFromBech32(creator1)
	suite.Require().NoError(err)
	creatorAddr2, err := sdk.AccAddressFromBech32(creator2)
	suite.Require().NoError(err)
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

	depositorBefore, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	suite.keeper.HandleTransfer(suite.ctx, &poolInfo)

	depositorAfterTransfer, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.Require().True(poolInfoBefore.BorrowedAmount.Equal(poolInfo.BorrowedAmount))
	suite.Require().True(poolInfoBefore.BorrowableAmount.Equal(poolInfo.BorrowableAmount))

	//fixme need to check the interest
	borrowed := sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1.34e5))
	all1 := depositAmount
	suite.Require().True(depositorAfterTransfer.LockedAmount.IsEqual(borrowed))
	suite.Require().Equal(depositorAfterTransfer.DepositType, types.DepositorInfo_processed)
	suite.Require().True(depositorAfterTransfer.WithdrawalAmount.IsEqual(all1.Sub(borrowed)))

	ids := strings.Split(depositorBefore.LinkedNFT[0], ":")
	_, found = suite.nftKeeper.GetNFT(suite.ctx, ids[0], ids[1])
	suite.Require().False(found)

	ids = strings.Split(depositorAfterTransfer.LinkedNFT[0], ":")
	nft1, found := suite.nftKeeper.GetNFT(suite.ctx, ids[0], ids[1])
	suite.Require().True(found)

	var nftInfo types.NftInfo
	err = proto.Unmarshal(nft1.Data.Value, &nftInfo)
	if err != nil {
		panic(err)
	}
	suite.Require().True(nftInfo.Borrowed.IsEqual(borrowed))

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Hour * 24))
	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	// 8+2-1.34=8.66 withdrawable
	poolInfoBefore, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	_, err = suite.app.TransferOwnership(suite.ctx, &reqOwner)
	suite.Require().NoError(err)

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Duration(poolInfo.PayFreq) * time.Second))

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	//err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)

	depositor1Before, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	fmt.Printf(">>>>>>>>>>>>>>>>>\n")
	spew.Dump(depositor1Before)
	fmt.Printf(">>>>>>>>>>>>>>>>>\n")

	depositor2Before, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)

	fmt.Printf(">>>>>>>>>>>>>>>>>\n")
	spew.Dump(depositor2Before)
	fmt.Printf(">>>>>>>>>>>>>>>>>\n")

	suite.keeper.HandleTransfer(suite.ctx, &poolInfo)

	depositor1AfterTransfer, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	fmt.Printf(">>>>>>>>>>>>>>>>>\n")
	spew.Dump(depositor1AfterTransfer)
	fmt.Printf(">>>>>>>>>>>>>>>>>\n")

	suite.keeper.HandleTransfer(suite.ctx, &poolInfo)

	depositor2AfterTransfer, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)

	fmt.Printf(">>>>>>>>>>>>>>>>>\n")
	spew.Dump(depositor2AfterTransfer)
	fmt.Printf(">>>>>>>>>>>>>>>>>\n")

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.Require().Len(depositor1AfterTransfer.LinkedNFT, 0)
	suite.Require().True(depositor1AfterTransfer.WithdrawalAmount.Amount.Equal(sdk.NewIntFromUint64(8e5)))
	suite.Require().True(depositor1AfterTransfer.LockedAmount.Amount.Equal(sdk.NewIntFromUint64(0)))
	suite.Require().True(depositor1AfterTransfer.DepositType == types.DepositorInfo_deposit_close)

	suite.Require().Len(depositor2AfterTransfer.LinkedNFT, 1)
	suite.Require().True(depositor2AfterTransfer.WithdrawalAmount.Amount.Equal(sdk.NewIntFromUint64(66000)))
	suite.Require().True(depositor2AfterTransfer.LockedAmount.Amount.Equal(sdk.NewIntFromUint64(1.34e5)))
	suite.Require().True(depositor2AfterTransfer.DepositType == types.DepositorInfo_unset)

	ids = strings.Split(depositor2AfterTransfer.LinkedNFT[0], ":")
	nft2, found := suite.nftKeeper.GetNFT(suite.ctx, ids[0], ids[1])
	suite.Require().True(found)

	suite.Require().True(nftInfo.Borrowed.Amount.Equal(sdk.NewIntFromUint64(1.34e5)))
	err = proto.Unmarshal(nft2.Data.Value, &nftInfo)
	if err != nil {
		panic(err)
	}

	spew.Dump(nftInfo)

	suite.Require().True(poolInfoBefore.BorrowedAmount.Equal(poolInfo.BorrowedAmount))
	// the 8e5 is not released unless deposit more,so the amount is
	suite.Require().True(poolInfoBefore.BorrowableAmount.SubAmount(sdk.NewIntFromUint64(8e5)).Amount.Equal(poolInfo.BorrowableAmount.Amount))

}

// need the case some each of the investor contribute some parts of the new transfer
