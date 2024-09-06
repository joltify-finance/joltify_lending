package keeper_test

import (
	"fmt"
	"strings"
	"time"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/davecgh/go-spew/spew"
	"github.com/gogo/protobuf/proto"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func checkValueEqualWithExchange(localAmount, usdAmount sdkmath.Int) bool {
	aUSD := convertBorrowToUsd(localAmount)
	fmt.Printf("exchange >>>>>ausd %v    b:%v  abs:%v\n", aUSD, usdAmount, aUSD.Sub(usdAmount).Abs())
	return usdAmount.Sub(aUSD).Abs().LTE(sdkmath.OneInt().MulRaw(3))
}

func checkValueWithRangeTwo(a, b sdkmath.Int) bool {
	fmt.Printf("check value>>>>>a:%v, b:%v  abs:%v\n", a, b, a.Sub(b).Abs())
	return a.Sub(b).Abs().LTE(sdkmath.OneInt().MulRaw(2))
}

func (suite *withDrawPrincipalSuite) TestTransferOwnershipOneInvestor() {
	setupPool(suite)
	// now we deposit some token and it should be enough to borrow
	creator1 := suite.investors[0]
	creatorAddr1, err := sdk.AccAddressFromBech32(creator1)
	suite.Require().NoError(err)
	// creatorAddr2, err := sdk.AccAddressFromBech32(creator2)
	// suite.Require().NoError(err)
	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(4e5))

	// suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{
		Creator:   creator1,
		PoolIndex: suite.investorPool,
		Token:     depositAmount,
	}
	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	//_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	//suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	// now we borrow 1.34e5
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Second * 20))
	_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{Creator: suite.investors[1], PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1e9))})
	suite.Require().NoError(err)

	reqOwner := types.MsgTransferOwnership{Creator: suite.investors[0], PoolIndex: suite.investorPool}
	_, err = suite.app.TransferOwnership(suite.ctx, &reqOwner)
	suite.Require().NoError(err)

	poolInfoBefore, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Duration(poolInfo.PayFreq) * time.Second))

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	// err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)

	depositorBefore, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	suite.keeper.HandleTransfer(suite.ctx, &poolInfo)

	depositorAfterTransfer, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.Require().True(checkValueWithRangeTwo(poolInfoBefore.BorrowedAmount.Amount, poolInfo.BorrowedAmount.Amount))
	suite.Require().True(poolInfoBefore.UsableAmount.Equal(poolInfo.UsableAmount))

	// fixme need to check the interest
	borrowed := sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))
	suite.Require().True(checkValueEqualWithExchange(depositorAfterTransfer.LockedAmount.Amount, borrowed.Amount))
	suite.Require().Equal(depositorAfterTransfer.DepositType, types.DepositorInfo_processed)
	spew.Dump(depositorAfterTransfer)
	suite.Require().True(depositorAfterTransfer.WithdrawalAmount.IsZero())

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
	suite.Require().True(checkValueEqualWithExchange(nftInfo.Borrowed.Amount, borrowed.Amount))

	// now we deposit more
	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Hour * 24))
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
	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(4e5))
	depositAmount2 := sdk.NewCoin("ausdc", sdkmath.NewInt(2e5))

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
		Token:     depositAmount2,
	}

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	//_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	//suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	// now we borrow 1.34e5
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Second * 20))
	_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{Creator: suite.investors[1], PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1e9))})
	suite.Require().NoError(err)

	addr, err := sdk.AccAddressFromBech32(suite.investors[0])
	suite.Require().NoError(err)
	d, ok := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, addr)
	suite.Require().True(ok)
	dUser0Returned := d.WithdrawalAmount.Amount
	reqOwner := types.MsgTransferOwnership{Creator: suite.investors[0], PoolIndex: suite.investorPool}
	_, err = suite.app.TransferOwnership(suite.ctx, &reqOwner)

	suite.Require().NoError(err)

	poolInfoBefore, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Duration(poolInfo.PayFreq) * time.Second))

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	// err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)

	depositorBefore, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	suite.keeper.HandleTransfer(suite.ctx, &poolInfo)

	depositorAfterTransfer, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.Require().True(checkValueWithRangeTwo(poolInfoBefore.BorrowedAmount.Amount, poolInfo.BorrowedAmount.Amount))
	suite.Require().True(poolInfoBefore.UsableAmount.Equal(poolInfo.UsableAmount))

	// fixme need to check the interest
	borrowed := sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))
	all1 := depositAmount
	suite.Require().True(checkValueEqualWithExchange(depositorAfterTransfer.LockedAmount.Amount, borrowed.Amount))
	suite.Require().Equal(depositorAfterTransfer.DepositType, types.DepositorInfo_processed)
	suite.Require().True(depositorAfterTransfer.WithdrawalAmount.IsEqual(all1.Sub(borrowed).SubAmount(dUser0Returned)))

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
	suite.Require().True(checkValueEqualWithExchange(nftInfo.Borrowed.Amount, borrowed.Amount))

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Hour * 24))
	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	// 4 +2=6 withdrawable
	poolInfoBefore, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	_, err = suite.app.TransferOwnership(suite.ctx, &reqOwner)
	suite.Require().NoError(err)

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Duration(poolInfo.PayFreq) * time.Second))

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	// err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)

	depositor2Before, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)

	suite.Require().Len(depositor2Before.LinkedNFT, 0)
	suite.Require().True(depositor2Before.LockedAmount.Amount.Equal(sdkmath.ZeroInt()))
	suite.Require().True(depositor2Before.WithdrawalAmount.Amount.Equal(sdkmath.NewIntFromUint64(2e5)))

	d, ok = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(ok)
	dUser1LockedUSD := convertBorrowToUsd(d.LockedAmount.Amount)

	suite.keeper.HandleTransfer(suite.ctx, &poolInfo)

	depositor1AfterTransfer, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	suite.keeper.HandleTransfer(suite.ctx, &poolInfo)

	depositor2AfterTransfer, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.Require().Len(depositor1AfterTransfer.LinkedNFT, 0)
	// all its locked is withdrawable now
	suite.Require().True(checkValueWithRangeTwo(depositor1AfterTransfer.WithdrawalAmount.Amount, dUser1LockedUSD))
	suite.Require().True(depositor1AfterTransfer.LockedAmount.Amount.Equal(sdkmath.NewIntFromUint64(0)))
	suite.Require().True(depositor1AfterTransfer.DepositType == types.DepositorInfo_deposit_close)

	suite.Require().Len(depositor2AfterTransfer.LinkedNFT, 1)
	suite.Require().True(checkValueWithRangeTwo(depositor2AfterTransfer.WithdrawalAmount.Amount, sdkmath.NewIntFromUint64(66000)))
	suite.Require().True(checkValueEqualWithExchange(depositor2AfterTransfer.LockedAmount.Amount, sdkmath.NewIntFromUint64(1.34e5)))
	suite.Require().True(depositor2AfterTransfer.DepositType == types.DepositorInfo_unset)

	ids = strings.Split(depositor2AfterTransfer.LinkedNFT[0], ":")
	nft2, found := suite.nftKeeper.GetNFT(suite.ctx, ids[0], ids[1])
	suite.Require().True(found)

	suite.Require().True(checkValueEqualWithExchange(nftInfo.Borrowed.Amount, sdkmath.NewIntFromUint64(1.34e5)))
	err = proto.Unmarshal(nft2.Data.Value, &nftInfo)
	if err != nil {
		panic(err)
	}

	spew.Dump(depositor1AfterTransfer)
	spew.Dump(depositor2AfterTransfer)

	spew.Dump(nftInfo)

	suite.Require().True(checkValueWithRangeTwo(poolInfoBefore.BorrowedAmount.Amount, poolInfo.BorrowedAmount.Amount))
	// now only depositor2 offers the money for the whole pool
	suite.Require().True(checkValueWithRangeTwo(depositor2AfterTransfer.WithdrawalAmount.Amount, poolInfo.UsableAmount.Amount))
}

func (suite *withDrawPrincipalSuite) TestTransferOwnershipTwoInvestorBoth() {
	setupPool(suite)
	// now we deposit some token and it should be enough to borrow
	creator1 := suite.investors[0]
	creator2 := suite.investors[1]
	creatorAddr1, err := sdk.AccAddressFromBech32(creator1)
	suite.Require().NoError(err)
	creatorAddr2, err := sdk.AccAddressFromBech32(creator2)
	suite.Require().NoError(err)
	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(4e5))
	depositAmount2 := sdk.NewCoin("ausdc", sdkmath.NewInt(1e5))

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
		Token:     depositAmount2,
	}

	_ = msgDepositUser2
	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	//_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	//suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	// now we borrow 1.34e5
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Second * 20))
	_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{Creator: suite.investors[1], PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1e9))})
	suite.Require().NoError(err)

	reqOwner := types.MsgTransferOwnership{Creator: suite.investors[0], PoolIndex: suite.investorPool}
	_, err = suite.app.TransferOwnership(suite.ctx, &reqOwner)
	suite.Require().NoError(err)

	poolInfoBefore, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Duration(poolInfo.PayFreq) * time.Second))
	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	// err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)

	depositorBefore, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	suite.keeper.HandleTransfer(suite.ctx, &poolInfo)

	depositorAfterTransfer, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.Require().True(checkValueWithRangeTwo(poolInfoBefore.BorrowedAmount.Amount, poolInfo.BorrowedAmount.Amount))
	suite.Require().True(checkValueWithRangeTwo(poolInfoBefore.UsableAmount.Amount, poolInfo.UsableAmount.Amount))

	// fixme need to check the interest
	borrowed := sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))
	suite.Require().True(checkValueEqualWithExchange(depositorAfterTransfer.LockedAmount.Amount, borrowed.Amount))
	suite.Require().Equal(depositorAfterTransfer.DepositType, types.DepositorInfo_processed)
	suite.Require().True(checkValueWithRangeTwo(depositorAfterTransfer.WithdrawalAmount.Amount, sdkmath.ZeroInt()))

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
	suite.Require().True(checkValueEqualWithExchange(nftInfo.Borrowed.Amount, borrowed.Amount))

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Hour * 24))
	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	// usable amount is 4e5 + 1e5 = 5e5
	poolInfoBefore, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	depositor1BeforeTransfer, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	_ = depositor1BeforeTransfer

	_, err = suite.app.TransferOwnership(suite.ctx, &reqOwner)
	suite.Require().NoError(err)

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Duration(poolInfo.PayFreq) * time.Second))

	//  after transfer the ownership, only user2's withdrawable money in the pool now which is 1e5
	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	depositor2Before, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)

	suite.Require().Len(depositor2Before.LinkedNFT, 0)
	suite.Require().True(depositor2Before.LockedAmount.Amount.Equal(sdkmath.ZeroInt()))
	suite.Require().True(depositor2Before.WithdrawalAmount.Amount.Equal(sdkmath.NewIntFromUint64(1e5)))

	suite.keeper.HandleTransfer(suite.ctx, &poolInfo)

	depositor1AfterTransfer, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	suite.keeper.HandleTransfer(suite.ctx, &poolInfo)

	depositor2AfterTransfer, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.Require().Len(depositor1AfterTransfer.LinkedNFT, 1)

	// user1 will get the locked 1e5 back
	suite.Require().True(checkValueWithRangeTwo(depositor1AfterTransfer.WithdrawalAmount.Amount, sdkmath.NewIntFromUint64(1e5)))
	suite.Require().True(checkValueEqualWithExchange(depositor1AfterTransfer.LockedAmount.Amount, sdkmath.NewIntFromUint64(0.34e5)))
	suite.Require().True(depositor1AfterTransfer.DepositType == types.DepositorInfo_processed)

	suite.Require().Len(depositor2AfterTransfer.LinkedNFT, 1)
	suite.Require().True(depositor2AfterTransfer.WithdrawalAmount.Amount.Equal(sdkmath.NewIntFromUint64(0)))
	suite.Require().True(checkValueEqualWithExchange(depositor2AfterTransfer.LockedAmount.Amount, sdkmath.NewIntFromUint64(1e5)))
	suite.Require().True(depositor2AfterTransfer.DepositType == types.DepositorInfo_unset)

	ids = strings.Split(depositor2AfterTransfer.LinkedNFT[0], ":")
	nft2, found := suite.nftKeeper.GetNFT(suite.ctx, ids[0], ids[1])
	suite.Require().True(found)

	err = proto.Unmarshal(nft2.Data.Value, &nftInfo)
	if err != nil {
		panic(err)
	}

	suite.Require().True(checkValueEqualWithExchange(nftInfo.Borrowed.Amount, sdkmath.NewIntFromUint64(1e5)))
	spew.Dump(depositor1AfterTransfer)
	spew.Dump(depositor2AfterTransfer)

	suite.Require().True(checkValueWithRangeTwo(poolInfoBefore.BorrowedAmount.Amount, poolInfo.BorrowedAmount.Amount))
	// the 8e5 is not released unless deposit more,so the amount is
	suite.Require().True(poolInfo.UsableAmount.Amount.IsZero())
}

// now we have 2 investors to "buy" the nft from the first user
func (suite *withDrawPrincipalSuite) TestTransferOwnershipSharedByTwoInvestors() {
	setupPool(suite)
	// now we deposit some token and it should be enough to borrow
	creator1 := suite.investors[0]
	creator2 := suite.investors[1]
	creator3 := suite.investors[2]
	creatorAddr1, err := sdk.AccAddressFromBech32(creator1)
	suite.Require().NoError(err)
	creatorAddr2, err := sdk.AccAddressFromBech32(creator2)
	suite.Require().NoError(err)
	creatorAddr3, err := sdk.AccAddressFromBech32(creator3)
	suite.Require().NoError(err)

	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(4e5))
	depositAmount2 := sdk.NewCoin("ausdc", sdkmath.NewInt(1e5))
	depositAmount3 := sdk.NewCoin("ausdc", sdkmath.NewInt(5e4))

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
		Token:     depositAmount2,
	}

	// user two deposit half of the amount of the user 1
	msgDepositUser3 := &types.MsgDeposit{
		Creator:   creator3,
		PoolIndex: suite.investorPool,
		Token:     depositAmount3,
	}

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	//_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	//suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	// now we borrow 1.34e5
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Second * 20))
	_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{Creator: suite.investors[1], PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1e9))})
	suite.Require().NoError(err)

	addr, err := sdk.AccAddressFromBech32(suite.investors[0])
	suite.Require().NoError(err)
	d, ok := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, addr)
	suite.Require().True(ok)
	user0Returned := d.WithdrawalAmount.Amount
	_ = user0Returned

	reqOwner := types.MsgTransferOwnership{Creator: suite.investors[0], PoolIndex: suite.investorPool}
	_, err = suite.app.TransferOwnership(suite.ctx, &reqOwner)

	suite.Require().NoError(err)

	poolInfoBefore, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Duration(poolInfo.PayFreq) * time.Second))

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	// err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)

	depositorBefore, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	suite.keeper.HandleTransfer(suite.ctx, &poolInfo)

	depositorAfterTransfer, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.Require().True(checkValueWithRangeTwo(poolInfoBefore.BorrowedAmount.Amount, poolInfoBefore.BorrowedAmount.Amount))
	suite.Require().True(poolInfoBefore.UsableAmount.Equal(poolInfo.UsableAmount))

	// fixme need to check the interest
	borrowed := sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))
	all1 := depositAmount
	suite.Require().True(checkValueEqualWithExchange(depositorAfterTransfer.LockedAmount.Amount, borrowed.Amount))
	suite.Require().Equal(depositorAfterTransfer.DepositType, types.DepositorInfo_processed)
	suite.Require().True(depositorAfterTransfer.WithdrawalAmount.IsEqual(all1.Sub(borrowed).SubAmount(user0Returned)))

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
	suite.Require().True(checkValueEqualWithExchange(nftInfo.Borrowed.Amount, borrowed.Amount))

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Hour * 24))
	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	// 8+1+0.5-1.34=7.66 withdrawable
	poolInfoBefore, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	_, err = suite.app.TransferOwnership(suite.ctx, &reqOwner)
	suite.Require().NoError(err)

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser3)
	suite.Require().NoError(err)

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime((sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Duration(poolInfo.PayFreq) * time.Second)))

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	// err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)

	depositor2Before, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)

	suite.Require().Len(depositor2Before.LinkedNFT, 0)
	suite.Require().True(depositor2Before.LockedAmount.Amount.Equal(sdkmath.ZeroInt()))
	suite.Require().True(depositor2Before.WithdrawalAmount.Amount.Equal(sdkmath.NewIntFromUint64(1e5)))

	suite.keeper.HandleTransfer(suite.ctx, &poolInfo)

	depositor1AfterTransfer, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	suite.keeper.HandleTransfer(suite.ctx, &poolInfo)

	depositor2AfterTransfer, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)

	depositor3AfterTransfer, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr3)
	suite.Require().True(found)

	suite.Require().Len(depositor1AfterTransfer.LinkedNFT, 0)
	suite.Require().Equal(depositor1AfterTransfer.DepositType, types.DepositorInfo_deposit_close)
	// the second investor need to locked all his money
	suite.Require().True(checkValueEqualWithExchange(depositor2AfterTransfer.LockedAmount.Amount, sdkmath.NewIntFromUint64(89333)))
	borrowable := sdkmath.NewIntFromUint64(1e5).Sub(sdkmath.NewIntFromUint64(89333))
	suite.Require().True(checkValueWithRangeTwo(depositor2AfterTransfer.WithdrawalAmount.Amount, borrowable))

	// investor3 get
	locked := sdkmath.NewIntFromUint64(1.34e5).Sub(sdkmath.NewIntFromUint64(89333))
	borrowable = sdkmath.NewIntFromUint64(5e4).Sub(locked)

	suite.Require().True(checkValueEqualWithExchange(depositor3AfterTransfer.LockedAmount.Amount, locked))
	suite.Require().True(checkValueWithRangeTwo(depositor3AfterTransfer.WithdrawalAmount.Amount, borrowable))
	// the first investor locked the reset of the money

	ids = strings.Split(depositor2AfterTransfer.LinkedNFT[0], ":")
	nft2, found := suite.nftKeeper.GetNFT(suite.ctx, ids[0], ids[1])
	suite.Require().True(found)

	err = proto.Unmarshal(nft2.Data.Value, &nftInfo)
	if err != nil {
		panic(err)
	}
	suite.Require().True(nftInfo.Borrowed.IsEqual(depositor2AfterTransfer.LockedAmount))

	ids = strings.Split(depositor2AfterTransfer.LinkedNFT[0], ":")
	nft3, found := suite.nftKeeper.GetNFT(suite.ctx, ids[0], ids[1])
	suite.Require().True(found)

	err = proto.Unmarshal(nft3.Data.Value, &nftInfo)
	if err != nil {
		panic(err)
	}

	suite.Require().True(nftInfo.Borrowed.IsEqual(depositor2AfterTransfer.LockedAmount))

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.Require().True(checkValueWithRangeTwo(poolInfoBefore.BorrowedAmount.Amount, poolInfo.BorrowedAmount.Amount))
	// the 8e5 is not released unless deposit more,so the amount is
	// 1.5e5-1.34e5=0.16e5
	suite.Require().True(checkValueWithRangeTwo(poolInfo.UsableAmount.Amount, sdkmath.NewIntFromUint64(0.16e5)))
}

func (suite *withDrawPrincipalSuite) TestTransferOwnershipSharedByMultipleEnoughMoney() {
	setupPool(suite)
	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(1e5))
	depositAmount2 := sdk.NewCoin("ausdc", sdkmath.NewInt(1.2e5))
	depositAmount3 := sdk.NewCoin("ausdc", sdkmath.NewInt(1.4e5))
	depositAmount4 := sdk.NewCoin("ausdc", sdkmath.NewInt(1.6e5))
	depositAmount5 := sdk.NewCoin("ausdc", sdkmath.NewInt(1e5))
	depositAmount6 := sdk.NewCoin("ausdc", sdkmath.NewInt(1.3e5))
	depositAmount7 := sdk.NewCoin("ausdc", sdkmath.NewInt(1.4e5))
	depositAmount8 := sdk.NewCoin("ausdc", sdkmath.NewInt(1.5e5))

	depositorAmounts := []sdk.Coin{depositAmount, depositAmount2, depositAmount3, depositAmount4, depositAmount5, depositAmount6, depositAmount7, depositAmount8}

	// now first 3 investor deposits
	for i := 0; i < 3; i++ {
		msgDeposit := &types.MsgDeposit{Creator: suite.investors[i], PoolIndex: suite.investorPool, Token: depositorAmounts[i]}
		_, err := suite.app.Deposit(suite.ctx, msgDeposit)
		suite.Require().NoError(err)
	}

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	// now we borrow 1.34e5
	_, err := suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime((sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Second * 20)))
	_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{Creator: suite.investors[1], PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1e9))})
	suite.Require().NoError(err)

	for i := 0; i < 3; i++ {
		reqOwner := types.MsgTransferOwnership{Creator: suite.investors[i], PoolIndex: suite.investorPool}
		_, err = suite.app.TransferOwnership(suite.ctx, &reqOwner)
	}

	suite.Require().NoError(err)

	poolInfoBefore, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime((sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Duration(poolInfo.PayFreq) * time.Second)))

	suite.Require().True(checkValueEqualWithExchange(poolInfoBefore.BorrowedAmount.Amount, sdkmath.NewIntFromUint64(1.34e5)))
	suite.Require().True(poolInfoBefore.UsableAmount.Amount.IsZero())
	// err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)

	totalBorrowed := sdkmath.ZeroInt()
	for i := 0; i < 3; i++ {
		creatorAddr, err := sdk.AccAddressFromBech32(suite.investors[i])
		suite.Require().NoError(err)
		depositorBefore, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr)
		suite.Require().True(found)
		totalBorrowed = totalBorrowed.Add(depositorBefore.LockedAmount.Amount)
	}

	suite.Require().True(totalBorrowed.Equal(poolInfoBefore.BorrowedAmount.Amount))

	for i := 3; i < 8; i++ {
		msgDeposit := &types.MsgDeposit{Creator: suite.investors[i], PoolIndex: suite.investorPool, Token: depositorAmounts[i]}
		_, err := suite.app.Deposit(suite.ctx, msgDeposit)
		suite.Require().NoError(err)
	}
	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdkmath.NewIntFromUint64(6.8e5)))

	suite.keeper.HandleTransfer(suite.ctx, &poolInfo)

	totalBorrowed2 := sdkmath.ZeroInt()

	for i := 0; i < 8; i++ {
		creatorAddr, err := sdk.AccAddressFromBech32(suite.investors[i])
		suite.Require().NoError(err)
		depositor, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr)
		suite.Require().True(found)
		spew.Dump(depositor)
		totalBorrowed2 = totalBorrowed2.Add(depositor.LockedAmount.Amount)
		if i < 3 {
			suite.Require().Len(depositor.LinkedNFT, 0)
			continue
		}
		ids := strings.Split(depositor.LinkedNFT[0], ":")
		nft1, found := suite.nftKeeper.GetNFT(suite.ctx, ids[0], ids[1])
		suite.Require().True(found)
		var nftInfo types.NftInfo
		err = proto.Unmarshal(nft1.Data.Value, &nftInfo)
		if err != nil {
			panic(err)
		}
		suite.Require().True(nftInfo.Borrowed.IsEqual(depositor.LockedAmount))
	}
	suite.Require().True(totalBorrowed.Equal(poolInfoBefore.BorrowedAmount.Amount))
	// 6.8-1.34
	suite.Require().True(checkValueWithRangeTwo(poolInfo.UsableAmount.Amount, sdkmath.NewIntFromUint64(5.46e5)))
}

func (suite *withDrawPrincipalSuite) TestTransferOwnershipSharedByMultipleNotEnoughMoneyAllHaveNFT() {
	setupPool(suite)

	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(1e5))
	depositAmount2 := sdk.NewCoin("ausdc", sdkmath.NewInt(1.2e5))
	depositAmount3 := sdk.NewCoin("ausdc", sdkmath.NewInt(1.4e5))
	depositAmount4 := sdk.NewCoin("ausdc", sdkmath.NewInt(1.6e4))
	depositAmount5 := sdk.NewCoin("ausdc", sdkmath.NewInt(1e4))
	depositAmount6 := sdk.NewCoin("ausdc", sdkmath.NewInt(1.3e4))
	depositAmount7 := sdk.NewCoin("ausdc", sdkmath.NewInt(1.4e4))
	depositAmount8 := sdk.NewCoin("ausdc", sdkmath.NewInt(1.5e4))

	depositorAmounts := []sdk.Coin{depositAmount, depositAmount2, depositAmount3, depositAmount4, depositAmount5, depositAmount6, depositAmount7, depositAmount8}

	totalMoney := sdkmath.ZeroInt()
	for _, el := range depositorAmounts {
		totalMoney = totalMoney.Add(el.Amount)
	}
	fmt.Printf(">>>>%v\n", totalMoney)

	// now first 3 investor deposits

	for i := 0; i < 3; i++ {
		msgDeposit := &types.MsgDeposit{Creator: suite.investors[i], PoolIndex: suite.investorPool, Token: depositorAmounts[i]}
		_, err := suite.app.Deposit(suite.ctx, msgDeposit)
		suite.Require().NoError(err)
	}

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	// now we borrow 1.34e5
	_, err := suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime((sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Second * 20)))
	_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{Creator: suite.investors[1], PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1e9))})
	suite.Require().NoError(err)

	totalRetuened := sdkmath.ZeroInt()
	for i := 0; i < 3; i++ {
		addr, err := sdk.AccAddressFromBech32(suite.investors[i])
		suite.Require().NoError(err)
		d, ok := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, addr)
		suite.Require().True(ok)
		totalRetuened = totalRetuened.Add(d.WithdrawalAmount.Amount)
		reqOwner := types.MsgTransferOwnership{Creator: suite.investors[i], PoolIndex: suite.investorPool}
		_, err = suite.app.TransferOwnership(suite.ctx, &reqOwner)
	}

	suite.Require().NoError(err)

	poolInfoBefore, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime((sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Duration(poolInfo.PayFreq) * time.Second)))

	suite.Require().True(checkValueEqualWithExchange(poolInfoBefore.BorrowedAmount.Amount, sdkmath.NewIntFromUint64(1.34e5)))
	suite.Require().True(poolInfoBefore.UsableAmount.Amount.IsZero())
	// err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)

	totalBorrowed := sdkmath.ZeroInt()
	for i := 0; i < 3; i++ {
		creatorAddr, err := sdk.AccAddressFromBech32(suite.investors[i])
		suite.Require().NoError(err)
		depositorBefore, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr)
		suite.Require().True(found)
		totalBorrowed = totalBorrowed.Add(depositorBefore.LockedAmount.Amount)
	}

	suite.Require().True(totalBorrowed.Equal(poolInfoBefore.BorrowedAmount.Amount))

	totalDeposit := sdkmath.ZeroInt()
	for i := 3; i < 8; i++ {
		msgDeposit := &types.MsgDeposit{Creator: suite.investors[i], PoolIndex: suite.investorPool, Token: depositorAmounts[i]}
		_, err := suite.app.Deposit(suite.ctx, msgDeposit)
		suite.Require().NoError(err)
		totalDeposit = totalDeposit.Add(msgDeposit.Token.Amount)
	}
	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdkmath.NewIntFromUint64(6.8e4)))

	suite.keeper.HandleTransfer(suite.ctx, &poolInfo)
	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	totalBorrowed2 := sdkmath.ZeroInt()
	totalBorrowable := sdkmath.ZeroInt()
	for i := 0; i < 8; i++ {
		creatorAddr, err := sdk.AccAddressFromBech32(suite.investors[i])
		suite.Require().NoError(err)
		depositor, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr)
		suite.Require().True(found)
		spew.Dump(depositor)
		totalBorrowed2 = totalBorrowed2.Add(depositor.LockedAmount.Amount)
		totalBorrowable = totalBorrowable.Add(depositor.WithdrawalAmount.Amount)
		ids := strings.Split(depositor.LinkedNFT[0], ":")
		nft1, found := suite.nftKeeper.GetNFT(suite.ctx, ids[0], ids[1])
		suite.Require().True(found)
		var nftInfo types.NftInfo
		err = proto.Unmarshal(nft1.Data.Value, &nftInfo)
		if err != nil {
			panic(err)
		}
		suite.Require().True(nftInfo.Borrowed.IsEqual(depositor.LockedAmount))
	}
	suite.Require().True(totalBorrowed2.Equal(poolInfo.BorrowedAmount.Amount))
	suite.Require().True(poolInfo.UsableAmount.Amount.IsZero())
	// 4.28-1.34=2.94
	suite.Require().True(checkValueWithRangeTwo(totalBorrowable, sdkmath.NewIntFromUint64(2.94e5).Sub(totalRetuened)))
}

func (suite *withDrawPrincipalSuite) TestTransferOwnershipSharedMultipleBorrowByMultipleNotEnoughMoneyAllHaveNFT() {
	setupPool(suite)

	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(1e5))
	depositAmount2 := sdk.NewCoin("ausdc", sdkmath.NewInt(1.2e5))
	depositAmount3 := sdk.NewCoin("ausdc", sdkmath.NewInt(1.4e5))
	depositAmount4 := sdk.NewCoin("ausdc", sdkmath.NewInt(1.6e4))
	depositAmount5 := sdk.NewCoin("ausdc", sdkmath.NewInt(1e4))
	depositAmount6 := sdk.NewCoin("ausdc", sdkmath.NewInt(1.3e4))
	depositAmount7 := sdk.NewCoin("ausdc", sdkmath.NewInt(1.4e4))
	depositAmount8 := sdk.NewCoin("ausdc", sdkmath.NewInt(1.5e4))

	depositorAmounts := []sdk.Coin{depositAmount, depositAmount2, depositAmount3, depositAmount4, depositAmount5, depositAmount6, depositAmount7, depositAmount8}

	totalMoney := sdkmath.ZeroInt()
	lastFour := sdkmath.ZeroInt()
	for i, el := range depositorAmounts {
		totalMoney = totalMoney.Add(el.Amount)
		if i >= 3 {
			lastFour = lastFour.Add(el.Amount)
		}
	}
	fmt.Printf(">>>>totalMoney %v\n", totalMoney)
	fmt.Printf(">>>>last4 %v\n", lastFour)

	// now first 3 investor deposits

	for i := 0; i < 3; i++ {
		msgDeposit := &types.MsgDeposit{Creator: suite.investors[i], PoolIndex: suite.investorPool, Token: depositorAmounts[i]}
		_, err := suite.app.Deposit(suite.ctx, msgDeposit)
		suite.Require().NoError(err)
	}

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	// now we borrow 1.34e5
	_, err := suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime((sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Second * 180)))
	borrow.BorrowAmount = sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(2e4))
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime((sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Second * 180)))
	borrow.BorrowAmount = sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.1e4))
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime((sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Second * 20)))
	_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{Creator: suite.investors[1], PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1e9))})
	suite.Require().NoError(err)

	totalReturned := sdkmath.ZeroInt()
	nodeRetuened := make([]sdkmath.Int, 3)
	for i := 0; i < 3; i++ {

		addr, err := sdk.AccAddressFromBech32(suite.investors[i])
		suite.Require().NoError(err)
		depositor, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, addr)
		suite.Require().True(found)
		totalReturned = totalReturned.Add(depositor.WithdrawalAmount.Amount)
		nodeRetuened[i] = depositor.WithdrawalAmount.Amount

		reqOwner := types.MsgTransferOwnership{Creator: suite.investors[i], PoolIndex: suite.investorPool}
		_, err = suite.app.TransferOwnership(suite.ctx, &reqOwner)
	}

	suite.Require().NoError(err)

	poolInfoBefore, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime((sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Duration(poolInfo.PayFreq) * time.Second)))

	suite.Require().True(checkValueEqualWithExchange(poolInfoBefore.BorrowedAmount.Amount, sdkmath.NewIntFromUint64(1.65e5)))
	suite.Require().True(poolInfoBefore.UsableAmount.Amount.IsZero())
	// err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)

	totalBorrowed := sdkmath.ZeroInt()
	for i := 0; i < 3; i++ {
		creatorAddr, err := sdk.AccAddressFromBech32(suite.investors[i])
		suite.Require().NoError(err)
		depositorBefore, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr)
		suite.Require().True(found)
		totalBorrowed = totalBorrowed.Add(depositorBefore.LockedAmount.Amount)
	}

	suite.Require().True(totalBorrowed.Equal(poolInfoBefore.BorrowedAmount.Amount))

	totalDeposit := sdkmath.ZeroInt()
	for i := 3; i < 8; i++ {
		msgDeposit := &types.MsgDeposit{Creator: suite.investors[i], PoolIndex: suite.investorPool, Token: depositorAmounts[i]}
		_, err := suite.app.Deposit(suite.ctx, msgDeposit)
		suite.Require().NoError(err)
		totalDeposit = totalDeposit.Add(msgDeposit.Token.Amount)
	}
	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdkmath.NewIntFromUint64(6.8e4)))
	creatorAddr, err := sdk.AccAddressFromBech32(suite.investors[0])
	suite.Require().NoError(err)
	depositor, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr)
	suite.Require().Len(depositor.LinkedNFT, 3)

	suite.keeper.HandleTransfer(suite.ctx, &poolInfo)
	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	totalBorrowed2 := sdkmath.ZeroInt()
	totalBorrowable := sdkmath.ZeroInt()
	for i := 0; i < 8; i++ {
		creatorAddr, err := sdk.AccAddressFromBech32(suite.investors[i])
		suite.Require().NoError(err)
		depositor, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr)
		suite.Require().True(found)
		totalBorrowed2 = totalBorrowed2.Add(depositor.LockedAmount.Amount)
		totalBorrowable = totalBorrowable.Add(depositor.WithdrawalAmount.Amount)
		ids := strings.Split(depositor.LinkedNFT[0], ":")
		nft1, found := suite.nftKeeper.GetNFT(suite.ctx, ids[0], ids[1])
		suite.Require().True(found)
		var nftInfo types.NftInfo
		err = proto.Unmarshal(nft1.Data.Value, &nftInfo)
		if err != nil {
			panic(err)
		}
		suite.Require().True(nftInfo.Borrowed.IsEqual(depositor.LockedAmount))
		if i >= 3 {
			borrowable := depositorAmounts[i].Amount
			ratio1 := sdkmath.LegacyNewDecFromInt(borrowable).QuoInt(sdkmath.NewIntFromUint64(0.68e5))
			ratio1Expected := sdkmath.LegacyNewDecFromInt(convertBorrowToUsd(nftInfo.Borrowed.Amount)).QuoInt(sdkmath.NewIntFromUint64(0.68e5))
			suite.Require().True(ratio1.Sub(ratio1Expected).Abs().LTE(sdkmath.LegacyNewDecWithPrec(1, 2)))
		} else {
			borrowable := depositorAmounts[i].Amount
			ratio1 := sdkmath.LegacyNewDecFromInt(borrowable).QuoInt(sdkmath.NewIntFromUint64(3.6e5))
			// 1.65-0.68
			ratio1Expected := sdkmath.LegacyNewDecFromInt(convertBorrowToUsd(nftInfo.Borrowed.Amount)).QuoInt(sdkmath.NewIntFromUint64(0.97e5))
			suite.Require().True(ratio1.Sub(ratio1Expected).Abs().LTE(sdkmath.LegacyNewDecWithPrec(1, 2)))
		}
	}
	suite.Require().True(totalBorrowed2.Equal(poolInfo.BorrowedAmount.Amount))
	suite.Require().True(poolInfo.UsableAmount.Amount.IsZero())
	// 4.28-1.65=2.94
	suite.Require().True(checkValueWithRangeTwo(totalBorrowable, sdkmath.NewIntFromUint64(2.63e5).Sub(totalReturned)))

	poolInfo, _ = suite.keeper.GetPools(suite.ctx, suite.investorPool)

	for i := 0; i < 3; i++ {
		req := types.NewMsgWithdrawPrincipal(suite.investors[i], suite.investorPool, sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(2e3)))
		resp, err := suite.app.WithdrawPrincipal(suite.ctx, req)
		suite.Require().NoError(err)
		parsed, err := sdk.ParseCoinsNormalized(resp.Amount)
		suite.Require().NoError(err)
		suite.Require().True(parsed[0].Amount.Equal(sdkmath.NewIntFromUint64(2e3)))
	}
	poolInfo, _ = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	// the first 3 investor money = 4.28(total)-0.68(last 4 investor money)
	// the first 3 leftover 3.6-(1.65-0.68 locked in pool)-0.06

	suite.Require().True(checkValueWithRangeTwo(poolInfo.UsableAmount.Amount, sdkmath.NewIntFromUint64(2.63e5).Sub(sdkmath.NewIntFromUint64(6e3)).Sub(totalReturned)))

	// now the first investor deposit and then withdraw all
	for i := 0; i < 1; i++ {
		msgDeposit := &types.MsgDeposit{Creator: suite.investors[i], PoolIndex: suite.investorPool, Token: depositorAmounts[i]}
		_, err := suite.app.Deposit(suite.ctx, msgDeposit)
		suite.Require().NoError(err)
		totalDeposit = totalDeposit.Add(msgDeposit.Token.Amount)
	}

	poolInfoBefore = poolInfo
	poolInfo, _ = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	//(4.28-0.68)-2e3*3
	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(poolInfoBefore.UsableAmount.Amount.Add(depositorAmounts[0].Amount)))

	// now we withdraw, it will send all the amount
	// first one deposit 26945
	// second one deposit 32333
	// the third deposit 37722
	lockeds := []sdkmath.Int{sdkmath.NewIntFromUint64(26945), sdkmath.NewIntFromUint64(32333), sdkmath.NewIntFromUint64(37722)}
	for i := 0; i < 3; i++ {
		req := types.NewMsgWithdrawPrincipal(suite.investors[i], suite.investorPool, sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(2e10)))
		resp, err := suite.app.WithdrawPrincipal(suite.ctx, req)
		suite.Require().NoError(err)
		parsed, err := sdk.ParseCoinsNormalized(resp.Amount)
		suite.Require().NoError(err)

		creatorAddr, err := sdk.AccAddressFromBech32(suite.investors[i])
		suite.Require().NoError(err)
		d, _ := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr)
		spew.Dump(d)
		if i == 0 {
			shouldGet := depositorAmounts[i].Amount.MulRaw(2).Sub(lockeds[i]).Sub(sdkmath.NewIntFromUint64(2e3)).Sub(nodeRetuened[i])
			suite.Require().True(checkValueWithRangeTwo(parsed[0].Amount, shouldGet))
			continue
		}

		shouldGet := depositorAmounts[i].Amount.Sub(lockeds[i]).Sub(sdkmath.NewIntFromUint64(2e3)).Sub(nodeRetuened[i])
		suite.Require().True(checkValueWithRangeTwo(parsed[0].Amount, shouldGet))
	}

	poolInfo, _ = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	// 88.74-2e3*3
	// suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdkmath.NewIntFromUint64(88.74e5).Sub(sdkmath.NewIntFromUint64(6e3).Add(depositorAmounts[0].Amount))))
	suite.Require().True(poolInfo.UsableAmount.Amount.IsZero())
	suite.Require().True(checkValueEqualWithExchange(poolInfo.BorrowedAmount.Amount, sdkmath.NewIntFromUint64(1.65e5)))
}

func (suite *withDrawPrincipalSuite) TestTransferOwnershipSharedMultipleBorrowByMultipleEnoughMoneyAllHaveNFT() {
	setupPool(suite)

	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(1e5))
	depositAmount2 := sdk.NewCoin("ausdc", sdkmath.NewInt(1.2e5))
	depositAmount3 := sdk.NewCoin("ausdc", sdkmath.NewInt(1.4e5))
	depositAmount4 := sdk.NewCoin("ausdc", sdkmath.NewInt(1.6e4))
	depositAmount5 := sdk.NewCoin("ausdc", sdkmath.NewInt(1e4))
	depositAmount6 := sdk.NewCoin("ausdc", sdkmath.NewInt(1.3e4))
	depositAmount7 := sdk.NewCoin("ausdc", sdkmath.NewInt(4e6))
	depositAmount8 := sdk.NewCoin("ausdc", sdkmath.NewInt(5e6))

	depositorAmounts := []sdk.Coin{depositAmount, depositAmount2, depositAmount3, depositAmount4, depositAmount5, depositAmount6, depositAmount7, depositAmount8}

	totalMoney := sdkmath.ZeroInt()
	for _, el := range depositorAmounts {
		totalMoney = totalMoney.Add(el.Amount)
	}

	// now first 3 investor deposits

	for i := 0; i < 3; i++ {
		msgDeposit := &types.MsgDeposit{Creator: suite.investors[i], PoolIndex: suite.investorPool, Token: depositorAmounts[i]}
		_, err := suite.app.Deposit(suite.ctx, msgDeposit)
		suite.Require().NoError(err)
	}

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	// now we borrow 1.34e5
	_, err := suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime((sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Second * 180)))
	borrow.BorrowAmount = sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(2e4))
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime((sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Second * 180)))
	borrow.BorrowAmount = sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.1e4))
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime((sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Second * 20)))
	_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{Creator: suite.investors[1], PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1e9))})
	suite.Require().NoError(err)

	totalReturned := sdkmath.ZeroInt()
	nodesReturned := make([]sdkmath.Int, 3)

	for i := 0; i < 3; i++ {

		addr, err := sdk.AccAddressFromBech32(suite.investors[i])
		suite.Require().NoError(err)
		d, ok := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, addr)
		suite.Require().True(ok)
		nodesReturned[i] = d.WithdrawalAmount.Amount
		totalReturned = totalReturned.Add(d.WithdrawalAmount.Amount)

		reqOwner := types.MsgTransferOwnership{Creator: suite.investors[i], PoolIndex: suite.investorPool}
		_, err = suite.app.TransferOwnership(suite.ctx, &reqOwner)
	}

	suite.Require().NoError(err)

	poolInfoBefore, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime((sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Duration(poolInfo.PayFreq) * time.Second)))

	suite.Require().True(checkValueEqualWithExchange(poolInfoBefore.BorrowedAmount.Amount, sdkmath.NewIntFromUint64(1.65e5)))
	suite.Require().True(poolInfoBefore.UsableAmount.Amount.IsZero())
	// err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)

	totalBorrowed := sdkmath.ZeroInt()
	for i := 0; i < 3; i++ {
		creatorAddr, err := sdk.AccAddressFromBech32(suite.investors[i])
		suite.Require().NoError(err)
		depositorBefore, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr)
		suite.Require().True(found)
		totalBorrowed = totalBorrowed.Add(depositorBefore.LockedAmount.Amount)
	}

	suite.Require().True(totalBorrowed.Equal(poolInfoBefore.BorrowedAmount.Amount))

	totalDeposit := sdkmath.ZeroInt()
	for i := 3; i < 8; i++ {
		msgDeposit := &types.MsgDeposit{Creator: suite.investors[i], PoolIndex: suite.investorPool, Token: depositorAmounts[i]}
		_, err := suite.app.Deposit(suite.ctx, msgDeposit)
		suite.Require().NoError(err)
		totalDeposit = totalDeposit.Add(msgDeposit.Token.Amount)
	}
	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdkmath.NewIntFromUint64(9039000)))
	creatorAddr, err := sdk.AccAddressFromBech32(suite.investors[0])
	suite.Require().NoError(err)
	depositor, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr)
	suite.Require().Len(depositor.LinkedNFT, 3)

	suite.keeper.HandleTransfer(suite.ctx, &poolInfo)
	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	totalBorrowed2 := sdkmath.ZeroInt()
	totalBorrowable := sdkmath.ZeroInt()
	for i := 0; i < 8; i++ {
		creatorAddr, err := sdk.AccAddressFromBech32(suite.investors[i])
		suite.Require().NoError(err)
		depositor, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr)
		suite.Require().True(found)
		spew.Dump(depositor)

		if i < 3 {
			suite.Require().Len(depositor.LinkedNFT, 0)
			suite.Require().True(depositor.DepositType == types.DepositorInfo_deposit_close)
			continue
		}

		totalBorrowed2 = totalBorrowed2.Add(depositor.LockedAmount.Amount)
		totalBorrowable = totalBorrowable.Add(depositor.WithdrawalAmount.Amount)
		ids := strings.Split(depositor.LinkedNFT[0], ":")
		nft1, found := suite.nftKeeper.GetNFT(suite.ctx, ids[0], ids[1])
		suite.Require().True(found)
		var nftInfo types.NftInfo
		err = proto.Unmarshal(nft1.Data.Value, &nftInfo)
		if err != nil {
			panic(err)
		}
		suite.Require().True(nftInfo.Borrowed.IsEqual(depositor.LockedAmount))
		borrowable := depositorAmounts[i].Amount
		ratio1 := sdkmath.LegacyNewDecFromInt(borrowable).QuoInt(sdkmath.NewIntFromUint64(90.39e5))
		ratio1Expected := sdkmath.LegacyNewDecFromInt(convertBorrowToUsd(nftInfo.Borrowed.Amount)).QuoInt(sdkmath.NewIntFromUint64(1.65e5))
		suite.Require().True(ratio1.Sub(ratio1Expected).Abs().LTE(sdkmath.LegacyNewDecWithPrec(1, 2)))
	}
	suite.Require().True(totalBorrowed2.Equal(poolInfo.BorrowedAmount.Amount))

	// 	93.99-1.65-3.6 =88.74
	suite.Require().True(checkValueWithRangeTwo(poolInfo.UsableAmount.Amount, sdkmath.NewIntFromUint64(88.74e5)))
	suite.Require().True(checkValueWithRangeTwo(totalBorrowable, sdkmath.NewIntFromUint64(88.74e5)))

	// now we withdraw, it will send all the amount
	for i := 0; i < 3; i++ {
		req := types.NewMsgWithdrawPrincipal(suite.investors[i], suite.investorPool, sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(2e3)))
		resp, err := suite.app.WithdrawPrincipal(suite.ctx, req)
		suite.Require().NoError(err)
		parsed, err := sdk.ParseCoinsNormalized(resp.Amount)
		suite.Require().NoError(err)
		suite.Require().True(checkValueWithRangeTwo(parsed[0].Amount, depositorAmounts[i].Amount.Sub(nodesReturned[i])))
		creatorAddr, err := sdk.AccAddressFromBech32(suite.investors[i])
		suite.Require().NoError(err)
		_, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr)
		suite.Require().False(found)
	}
}
