package keeper_test

import (
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogo/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/suite"
)

// Test suite used for all keeper tests
type addBorrowSuite struct {
	suite.Suite
	keeper     *spvkeeper.Keeper
	nftKeeper  types.NFTKeeper
	bankKeeper types.BankKeeper
	app        types.MsgServer
	ctx        context.Context
}

func TestBorrowTestSuite(t *testing.T) {
	suite.Run(t, new(addBorrowSuite))
}

// The default state used by each test
func (suite *addBorrowSuite) SetupTest() {
	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)

	lapp, k, nftKeeper, bankKeeper, _, wctx := setupMsgServer(suite.T())
	ctx := sdk.UnwrapSDKContext(wctx)
	suite.ctx = ctx
	suite.keeper = k
	suite.nftKeeper = nftKeeper
	suite.bankKeeper = bankKeeper
	suite.app = lapp
}

func (suite *addBorrowSuite) TestAddBorrow() {
	type args struct {
		msgBorrow   *types.MsgBorrow
		expectedErr string
	}

	type test struct {
		name string
		args args
	}

	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 2, PoolName: "hello", Apy: []string{"7.8", "7.2"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdkmath.NewInt(1*1e3)), sdk.NewCoin("ausdc", sdkmath.NewInt(1*1e3))}}
	resp, err := suite.app.CreatePool(suite.ctx, &req)
	suite.Require().NoError(err)

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

	testCases := []test{
		{
			name: "invalid address",
			args: args{msgBorrow: &types.MsgBorrow{Creator: "invalid address"}, expectedErr: "invalid address invalid address: invalid address"},
		},

		{
			name: "pool cannot be found",
			args: args{msgBorrow: &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0"}, expectedErr: "pool cannot be found"},
		},
		{
			name: "is not authorised to borrow",
			args: args{msgBorrow: &types.MsgBorrow{Creator: "jolt1m28h5mu57ugcpfw2sp5t9chdp69akzc6ze5r0j", PoolIndex: resp.PoolIndex[0], BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(100))}, expectedErr: "not authorized to borrow money"},
		},

		{
			name: "inconsistency toekn denom",
			args: args{msgBorrow: &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: resp.PoolIndex[0], BorrowAmount: sdk.NewCoin("aaa", sdkmath.NewIntFromUint64(2))}, expectedErr: "token to be borrowed is inconsistency"},
		},
		{
			name: "reach borrow limit",
			args: args{msgBorrow: &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: resp.PoolIndex[0], BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(2233))}, expectedErr: "pool reached borrow limit"},
		},
		{
			name: "not enough to borrow",
			args: args{msgBorrow: &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: resp.PoolIndex[0], BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(2233))}, expectedErr: "insufficient tokens"},
		},

		{
			name: "not reach the minimum borrow amount",
			args: args{msgBorrow: &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: resp.PoolIndex[0], BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(2))}, expectedErr: "pool minimal borrow is 100ausdc and you try to borrow 2ausdc: invalid parameter"},
		},

		{
			name: "expire borrow time",
			args: args{msgBorrow: &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: resp.PoolIndex[0], BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(2233))}, expectedErr: "pool borrow time window expired"},
		},

		{
			name: "not enough to borrow",
			args: args{msgBorrow: &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: resp.PoolIndex[0], BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(2233))}, expectedErr: "pool borrow time window expired"},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			poolInfo, found := suite.keeper.GetPools(suite.ctx, resp.PoolIndex[0])
			suite.Require().True(found)
			poolInfo.CurrentPoolTotalBorrowCounter = 0
			poolInfo.PoolTotalBorrowLimit = 1
			poolInfo.UsableAmount = poolInfo.TargetAmount
			poolInfo.MinBorrowAmount = sdk.NewCoin(poolInfo.TargetAmount.Denom, sdkmath.NewIntFromUint64(100))

			if tc.name == "reach borrow limit" {
				poolInfo.PoolTotalBorrowLimit = 0
			}
			if tc.name == "not enough to borrow" {
				poolInfo.UsableAmount = sdk.NewCoin("ausdc", sdkmath.NewInt(100))
			}

			if tc.name == "expire borrow time" {
				expiredTime := poolInfo.PoolCreatedTime.Add(time.Second*time.Duration(poolInfo.PoolLockedSeconds) + poolInfo.GraceTime + time.Second)
				suite.ctx = suite.ctx.WithBlockTime(expiredTime)
			}

			suite.keeper.SetPool(suite.ctx, poolInfo)
			_, err := suite.app.Borrow(suite.ctx, tc.args.msgBorrow)
			if tc.args.expectedErr != "" {
				suite.Require().ErrorContains(err, tc.args.expectedErr)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *addBorrowSuite) compareDepositor(expected, actual types.DepositorInfo) {
	suite.Require().Equal(expected.InvestorId, actual.InvestorId)
	suite.Require().True(expected.DepositorAddress.Equals(actual.DepositorAddress))
	suite.Require().True(expected.LockedAmount.IsEqual(actual.LockedAmount))
	suite.Require().True(expected.WithdrawalAmount.IsEqual(actual.WithdrawalAmount))
	suite.Require().Equal(len(expected.LinkedNFT), len(actual.LinkedNFT))
	for i, el := range expected.LinkedNFT {
		suite.Require().Equal(el, actual.LinkedNFT[i])
	}
}

func (suite *addBorrowSuite) TestBorrowValueCheck() {
	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 2, PoolName: "hello", Apy: []string{"0.15", "0.15"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdkmath.NewInt(1*1e6)), sdk.NewCoin("ausdc", sdkmath.NewInt(1e6))}}
	resp, err := suite.app.CreatePool(suite.ctx, &req)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, resp.PoolIndex[0])
	suite.Require().True(found)
	poolInfo.CurrentPoolTotalBorrowCounter = 0
	poolInfo.PoolTotalBorrowLimit = 10
	poolInfo.TargetAmount = sdk.NewCoin("ausdc", sdkmath.NewInt(1*1e6))
	suite.keeper.SetPool(suite.ctx, poolInfo)

	depositorPool := resp.PoolIndex[0]

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

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: depositorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().ErrorContains(err, "insufficient tokens")

	// now we deposit some token and it should be enough to borrow
	creator1 := "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl"
	creator2 := "jolt1kkujrm0lqeu0e5va5f6mmwk87wva0k8cmam8jq"
	creatorAddr1, err := sdk.AccAddressFromBech32(creator1)
	suite.Require().NoError(err)
	creatorAddr2, err := sdk.AccAddressFromBech32(creator2)
	suite.Require().NoError(err)
	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(4e5))
	suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{
		Creator:   creator1,
		PoolIndex: depositorPool,
		Token:     depositAmount,
	}

	// user two deposit half of the amount of the user 1
	msgDepositUser2 := &types.MsgDeposit{
		Creator:   creator2,
		PoolIndex: depositorPool,
		Token:     depositAmount.SubAmount(sdk.NewInt(2e5)),
	}

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	// we check the depositor info
	depositor, found := suite.keeper.GetDepositor(suite.ctx, depositorPool, creatorAddr1)
	suite.Require().True(found)
	targetDepositor := types.DepositorInfo{
		InvestorId:       "2",
		DepositorAddress: creatorAddr1,
		PoolIndex:        depositorPool,
		LockedAmount:     sdk.NewCoin("aud-ausdc", sdkmath.ZeroInt()),
		WithdrawalAmount: depositAmount,
		LinkedNFT:        []string{},
	}

	suite.compareDepositor(targetDepositor, depositor)
	// we deposit again,so withdrawal is doubled

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	targetDepositor.WithdrawalAmount = targetDepositor.WithdrawalAmount.Add(depositAmount)

	depositor, found = suite.keeper.GetDepositor(suite.ctx, depositorPool, creatorAddr1)
	suite.Require().True(found)
	suite.compareDepositor(targetDepositor, depositor)

	// we mock the second user deposits the token, now we have 3*4e5 tokens
	//_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	//suite.Require().NoError(err)

	pool, found := suite.keeper.GetPools(suite.ctx, depositorPool)
	suite.Require().True(found)

	totalBorrowable := msgDepositUser1.Token.Add(msgDepositUser1.Token).Add(msgDepositUser2.Token)
	suite.Require().True(totalBorrowable.IsEqual(pool.UsableAmount))

	user1Ratio := sdkmath.LegacyNewDecFromInt(msgDepositUser1.Token.Amount.Mul(sdk.NewInt(2))).Quo(sdkmath.LegacyNewDecFromInt(totalBorrowable.Amount))

	user2Ratio := sdkmath.LegacyNewDecFromInt(msgDepositUser2.Token.Amount).Quo(sdkmath.LegacyNewDecFromInt(totalBorrowable.Amount))

	// now we borrow 2e5
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	failedBorrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: depositorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(100.34e5))}
	// pool reached the limit test
	_, err = suite.app.Borrow(suite.ctx, failedBorrow)
	suite.Require().ErrorContains(err, "insufficient tokens")

	p1, found := suite.keeper.GetDepositor(suite.ctx, depositorPool, creatorAddr1)
	suite.Require().True(found)

	p2, found := suite.keeper.GetDepositor(suite.ctx, depositorPool, creatorAddr2)
	suite.Require().True(found)

	poolNow, found := suite.keeper.GetPools(suite.ctx, depositorPool)
	suite.Require().True(found)

	suite.Require().True(checkValueEqualWithExchange(poolNow.BorrowedAmount.Amount, borrow.BorrowAmount.Amount))
	suite.Require().True(checkValueWithRangeTwo(totalBorrowable.Amount, poolNow.UsableAmount.Amount.Add(borrow.BorrowAmount.Amount)))

	borrowedFromUser1 := sdkmath.LegacyNewDecFromInt(borrow.BorrowAmount.Amount).Mul(user1Ratio).TruncateInt()
	borrowedFromUser2 := borrow.BorrowAmount.Amount.Sub(borrowedFromUser1)

	borrowedFromUser2Ratio := sdkmath.LegacyNewDecFromInt(borrow.BorrowAmount.Amount).Mul(user2Ratio).TruncateInt()

	suite.Require().True(checkValueEqualWithExchange(p1.LockedAmount.Amount, borrowedFromUser1))
	suite.Require().True(checkValueEqualWithExchange(p2.LockedAmount.Amount, borrowedFromUser2))
	suite.Require().True(checkValueWithRangeTwo(borrowedFromUser2Ratio, borrowedFromUser2))

	// total amount shoube be locked+withdrawable
	suite.Require().True(p1.WithdrawalAmount.AddAmount(borrowedFromUser1).IsEqual(msgDepositUser1.Token.Add(msgDepositUser1.Token)))

	nftUser1 := p1.LinkedNFT[0]
	nftUser2 := p2.LinkedNFT[0]

	nftClassID := fmt.Sprintf("class-%v-0", depositorPool[2:])
	nft, found := suite.nftKeeper.GetClass(suite.ctx, nftClassID)
	suite.Require().True(found)

	var borrowClassInfo types.BorrowInterest
	err = proto.Unmarshal(nft.Data.Value, &borrowClassInfo)
	if err != nil {
		panic(err)
	}

	lastBorrow := borrowClassInfo.BorrowDetails[len(borrowClassInfo.BorrowDetails)-1].BorrowedAmount
	suite.True(checkValueEqualWithExchange(lastBorrow.Amount, borrow.BorrowAmount.Amount))
	suite.Require().True(borrowClassInfo.Apy.Equal(sdk.NewDecWithPrec(15, 2)))

	// nft ID is the hash(nft class ID, investorWallet)
	indexHash := crypto.Keccak256Hash([]byte(nftClassID), p1.DepositorAddress)
	expectedID1 := fmt.Sprintf("%v:invoice-%v", nftClassID, indexHash.String()[2:])
	suite.Require().Equal(nftUser1, expectedID1)

	indexHash = crypto.Keccak256Hash([]byte(nftClassID), p2.DepositorAddress)
	expectedID2 := fmt.Sprintf("%v:invoice-%v", nftClassID, indexHash.String()[2:])
	suite.Require().Equal(nftUser2, expectedID2)

	dat := strings.Split(nftUser1, ":")
	nft1, found := suite.nftKeeper.GetNFT(suite.ctx, dat[0], dat[1])
	suite.Require().True(found)

	var nftInfo types.NftInfo
	err = proto.Unmarshal(nft1.Data.Value, &nftInfo)
	if err != nil {
		panic(err)
	}

	lastBorrow = borrowClassInfo.BorrowDetails[len(borrowClassInfo.BorrowDetails)-1].BorrowedAmount
	ration1 := sdkmath.LegacyNewDecFromInt(nftInfo.Borrowed.Amount).Quo(sdkmath.LegacyNewDecFromInt(lastBorrow.Amount))

	fmt.Printf(">>>>%v===%v\n", ration1, user1Ratio)
	suite.Require().True(ration1.Sub(user1Ratio).Abs().LTE(sdkmath.LegacyMustNewDecFromStr("0.0001")))

	// now, user 2 deposits more money and then, the spv borrow more. the ratio should  be changed.
	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	user1Cached := p1.WithdrawalAmount
	user2Cached := p2.WithdrawalAmount

	p2, found = suite.keeper.GetDepositor(suite.ctx, depositorPool, creatorAddr2)

	suite.Require().True(found)
	suite.Require().True(p2.WithdrawalAmount.IsEqual(user2Cached.AddAmount(msgDepositUser2.Token.Amount)))
	suite.Require().True(poolNow.GetUsableAmount().IsEqual(user1Cached.Add(user2Cached)))
	user2Cached = p2.WithdrawalAmount
	newuser1Ratio1 := sdkmath.LegacyNewDecFromInt(user1Cached.Amount).Quo(sdkmath.LegacyNewDecFromInt(user1Cached.Add(user2Cached).Amount))

	poolAfterUser2SecondDeposit, found := suite.keeper.GetPools(suite.ctx, depositorPool)
	suite.Require().True(found)
	suite.Require().True(poolAfterUser2SecondDeposit.UsableAmount.IsEqual(user2Cached.Add(user1Cached)))

	// NOW we borrow
	borrow.BorrowAmount = sdk.NewCoin(borrow.BorrowAmount.Denom, sdkmath.NewInt(1.2e5))
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	previousAmountBorrowed := poolNow.BorrowedAmount
	previousBorrowAble := poolNow.UsableAmount
	poolNow, found = suite.keeper.GetPools(suite.ctx, depositorPool)
	suite.Require().True(found)

	suite.Require().True(checkValueEqualWithExchange(poolNow.BorrowedAmount.Amount, borrow.BorrowAmount.AddAmount(convertBorrowToUsd(previousAmountBorrowed.Amount)).Amount))

	suite.Require().True(poolNow.UsableAmount.Equal(previousBorrowAble.Add(msgDepositUser2.Token).Sub(borrow.BorrowAmount)))

	beforeLockedAmount := p1.LockedAmount
	// now we check the nfts
	p1, found = suite.keeper.GetDepositor(suite.ctx, depositorPool, creatorAddr1)
	suite.Require().True(found)

	p2, found = suite.keeper.GetDepositor(suite.ctx, depositorPool, creatorAddr2)
	suite.Require().True(found)

	lockedThistime := p1.LockedAmount.Sub(beforeLockedAmount)
	shouldLocked := newuser1Ratio1.Mul(sdkmath.LegacyNewDecFromInt(borrow.BorrowAmount.Amount)).TruncateInt()
	suite.Require().True(checkValueEqualWithExchange(lockedThistime.Amount, shouldLocked))

	// we check the total deposit of the user1 is correct
	suite.Require().True(checkValueWithRangeTwo(convertBorrowToUsd(p1.LockedAmount.Amount).Add(p1.WithdrawalAmount.Amount), msgDepositUser1.Token.Add(msgDepositUser1.Token).Amount))

	nft2User1 := p1.LinkedNFT[1]
	nft2User2 := p2.LinkedNFT[1]

	nftClassID = fmt.Sprintf("class-%v-1", depositorPool[2:])
	nft, found = suite.nftKeeper.GetClass(suite.ctx, nftClassID)
	suite.Require().True(found)

	err = proto.Unmarshal(nft.Data.Value, &borrowClassInfo)
	if err != nil {
		panic(err)
	}

	lastBorrow = borrowClassInfo.BorrowDetails[len(borrowClassInfo.BorrowDetails)-1].BorrowedAmount
	suite.True(checkValueEqualWithExchange(lastBorrow.Amount, borrow.BorrowAmount.Amount))

	// nft ID is the hash(nft class ID, investorWallet)
	indexHash = crypto.Keccak256Hash([]byte(nftClassID), p1.DepositorAddress)
	expectedID1 = fmt.Sprintf("%v:invoice-%v", nftClassID, indexHash.String()[2:])
	suite.Require().Equal(nft2User1, expectedID1)

	indexHash = crypto.Keccak256Hash([]byte(nftClassID), p2.DepositorAddress)
	expectedID2 = fmt.Sprintf("%v:invoice-%v", nftClassID, indexHash.String()[2:])
	suite.Require().Equal(nft2User2, expectedID2)

	dat = strings.Split(nft2User1, ":")
	nft1, found = suite.nftKeeper.GetNFT(suite.ctx, dat[0], dat[1])
	suite.Require().True(found)

	err = proto.Unmarshal(nft1.Data.Value, &nftInfo)
	if err != nil {
		panic(err)
	}

	// this calculates the ratio that user1 contribute to this borrow
	suite.Require().True(checkValueEqualWithExchange(nftInfo.Borrowed.Amount, shouldLocked))
}

func (suite *addBorrowSuite) TestMultipleBorrowWithInterestPaid() {
	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 2, PoolName: "hello", Apy: []string{"0.15", "0.15"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdkmath.NewInt(1*1e6)), sdk.NewCoin("ausdc", sdkmath.NewInt(1e6))}}
	resp, err := suite.app.CreatePool(suite.ctx, &req)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, resp.PoolIndex[0])
	suite.Require().True(found)
	poolInfo.CurrentPoolTotalBorrowCounter = 0
	poolInfo.PoolTotalBorrowLimit = 10
	poolInfo.TargetAmount = sdk.NewCoin("ausdc", sdkmath.NewInt(1*1e6))
	suite.keeper.SetPool(suite.ctx, poolInfo)

	depositorPool := resp.PoolIndex[0]

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

	// now we deposit some token and it should be enough to borrow
	creator1 := "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl"
	creator2 := "jolt1kkujrm0lqeu0e5va5f6mmwk87wva0k8cmam8jq"
	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(4e5))
	suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{
		Creator:   creator1,
		PoolIndex: depositorPool,
		Token:     depositAmount,
	}

	// user two deposit half of the amount of the user 1
	msgDepositUser2 := &types.MsgDeposit{
		Creator:   creator2,
		PoolIndex: depositorPool,
		Token:     depositAmount.SubAmount(sdk.NewInt(2e5)),
	}

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)
	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: depositorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	_, err = suite.app.Borrow(suite.ctx, borrow)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, resp.PoolIndex[0])
	suite.Require().True(found)

	period := spvkeeper.OneYear / poolInfo.PayFreq
	interestWithReserve := sdkmath.LegacyNewDecFromInt(sdk.NewIntFromUint64(1.34e5)).Mul(sdkmath.LegacyMustNewDecFromStr("0.15")).QuoInt64(int64(period)).TruncateInt()

	for i := 0; i < 10; i++ {
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(poolInfo.PayFreq)))
		err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)
		suite.Require().NoError(err)
		if i == 0 {
			suite.Require().True(checkValueWithRangeTwo(interestWithReserve, poolInfo.EscrowInterestAmount.Mul(sdk.NewIntFromBigInt(big.NewInt(-1)))))
		}
	}

	totalInterestOwn := interestWithReserve.MulRaw(10)
	_, err = suite.app.RepayInterest(suite.ctx, types.NewMsgRepayInterest("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", depositorPool, sdk.NewCoin("ausdc", totalInterestOwn.SubRaw(100))))
	suite.Require().ErrorContains(err, "you must pay all the outstanding interest which is")

	addr := authtypes.NewModuleAddress("spv")
	coinBefore := suite.bankKeeper.GetBalance(suite.ctx, addr, "ausdc")

	planToPay := interestWithReserve.MulRaw(11).AddRaw(101)
	_, err = suite.app.RepayInterest(suite.ctx, types.NewMsgRepayInterest("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", depositorPool, sdk.NewCoin("ausdc", planToPay)))
	suite.Require().NoError(err)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, resp.PoolIndex[0])
	suite.Require().True(found)

	coinAfter := suite.bankKeeper.GetBalance(suite.ctx, addr, "ausdc")

	delta := coinAfter.Sub(coinBefore)
	oneInterest := poolInfo.EscrowInterestAmount
	suite.Require().True(delta.Amount.Equal(oneInterest.MulRaw(11)))

	half := poolInfo.PayFreq / 2

	borrower, err := sdk.AccAddressFromBech32(borrow.Creator)
	suite.Require().NoError(err)
	userCoinBefore := suite.bankKeeper.GetBalance(suite.ctx, borrower, "ausdc")
	ctx := suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(half)))
	// now we borrow again
	_, err = suite.app.Borrow(ctx, borrow)
	suite.Require().NoError(err)
	userCoinAfter := suite.bankKeeper.GetBalance(suite.ctx, borrower, "ausdc")

	p, found := suite.keeper.GetPools(suite.ctx, resp.PoolIndex[0])
	suite.Require().True(found)
	// now we return the escrow interest as it is not enough to pay the interest
	suite.Require().True(p.EscrowInterestAmount.IsZero())
	suite.Require().Nil(p.InterestPrepayment)

	deltaUser := userCoinAfter.Sub(userCoinBefore).Amount.Sub(sdk.NewIntFromUint64(1.34e5))
	suite.Require().True(checkValueWithRangeTwo(deltaUser, oneInterest))
	fmt.Printf(">>1111>>>%v\n", deltaUser)
}

func (suite *addBorrowSuite) TestMultipleBorrowWithInterestPaidUpdatePrePaid() {
	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 2, PoolName: "hello", Apy: []string{"0.15", "0.15"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdkmath.NewInt(1*1e6)), sdk.NewCoin("ausdc", sdkmath.NewInt(1e6))}}
	resp, err := suite.app.CreatePool(suite.ctx, &req)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, resp.PoolIndex[0])
	suite.Require().True(found)
	poolInfo.CurrentPoolTotalBorrowCounter = 0
	poolInfo.PoolTotalBorrowLimit = 10
	poolInfo.TargetAmount = sdk.NewCoin("ausdc", sdkmath.NewInt(1*1e6))
	suite.keeper.SetPool(suite.ctx, poolInfo)

	depositorPool := resp.PoolIndex[0]

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

	// now we deposit some token and it should be enough to borrow
	creator1 := "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl"
	creator2 := "jolt1kkujrm0lqeu0e5va5f6mmwk87wva0k8cmam8jq"
	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(4e5))
	suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{
		Creator:   creator1,
		PoolIndex: depositorPool,
		Token:     depositAmount,
	}

	// user two deposit half of the amount of the user 1
	msgDepositUser2 := &types.MsgDeposit{
		Creator:   creator2,
		PoolIndex: depositorPool,
		Token:     depositAmount.SubAmount(sdk.NewInt(2e5)),
	}

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)
	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: depositorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	_, err = suite.app.Borrow(suite.ctx, borrow)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, resp.PoolIndex[0])
	suite.Require().True(found)

	period := spvkeeper.OneYear / poolInfo.PayFreq
	interestWithReserve := sdkmath.LegacyNewDecFromInt(sdk.NewIntFromUint64(1.34e5)).Mul(sdkmath.LegacyMustNewDecFromStr("0.15")).QuoInt64(int64(period)).TruncateInt()

	var oneInterest sdkmath.Int
	for i := 0; i < 10; i++ {
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(poolInfo.PayFreq)))
		err = suite.keeper.HandleInterest(suite.ctx, &poolInfo)
		suite.Require().NoError(err)
		if i == 0 {
			suite.Require().True(checkValueWithRangeTwo(interestWithReserve, poolInfo.EscrowInterestAmount.Mul(sdk.NewIntFromBigInt(big.NewInt(-1)))))
			oneInterest = poolInfo.EscrowInterestAmount.MulRaw(-1)
		}
	}

	totalInterestOwn := interestWithReserve.MulRaw(10)
	_, err = suite.app.RepayInterest(suite.ctx, types.NewMsgRepayInterest("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", depositorPool, sdk.NewCoin("ausdc", totalInterestOwn.SubRaw(100))))
	suite.Require().ErrorContains(err, "you must pay all the outstanding interest which is")

	addr := authtypes.NewModuleAddress("spv")
	coinBefore := suite.bankKeeper.GetBalance(suite.ctx, addr, "ausdc")

	planToPay := interestWithReserve.MulRaw(15).AddRaw(101)
	_, err = suite.app.RepayInterest(suite.ctx, types.NewMsgRepayInterest("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", depositorPool, sdk.NewCoin("ausdc", planToPay)))
	suite.Require().NoError(err)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, resp.PoolIndex[0])
	suite.Require().True(found)

	coinAfter := suite.bankKeeper.GetBalance(suite.ctx, addr, "ausdc")

	delta := coinAfter.Sub(coinBefore)
	suite.Require().Equal(poolInfo.InterestPrepayment.Counter, int32(5))

	suite.Require().True(delta.Amount.Equal(oneInterest.MulRaw(15)))
	suite.Require().True(poolInfo.EscrowInterestAmount.Equal(oneInterest.MulRaw(5)))

	escrowInterestBefore := poolInfo.EscrowInterestAmount

	half := poolInfo.PayFreq / 2
	ctx := suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(half)))
	// now we borrow again

	borrower, err := sdk.AccAddressFromBech32(borrow.Creator)
	suite.Require().NoError(err)
	userCoinBefore := suite.bankKeeper.GetBalance(suite.ctx, borrower, "ausdc")
	_, err = suite.app.Borrow(ctx, borrow)
	suite.Require().NoError(err)
	userCoinAfter := suite.bankKeeper.GetBalance(suite.ctx, borrower, "ausdc")
	currentInterest := oneInterest.MulRaw(2)

	p, found := suite.keeper.GetPools(suite.ctx, resp.PoolIndex[0])
	suite.Require().True(found)
	// now the interest is doubled after second borrow
	counter := p.EscrowInterestAmount.Quo(currentInterest)
	suite.Require().Equal(p.InterestPrepayment.Counter, int32(counter.Int64()))

	returned := escrowInterestBefore.Sub(currentInterest.Mul(counter))

	spvReturned := userCoinAfter.Sub(userCoinBefore).Amount.Sub(sdk.NewIntFromUint64(1.34e5))

	suite.Require().True(checkValueWithRangeTwo(returned, spvReturned))
}
