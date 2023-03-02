package keeper_test

import (
	"fmt"
	"strings"
	"testing"

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
	keeper    *spvkeeper.Keeper
	nftKeeper types.NFTKeeper
	app       types.MsgServer
	ctx       sdk.Context
}

func TestBorrowTestSuite(t *testing.T) {
	suite.Run(t, new(addBorrowSuite))
}

// The default state used by each test
func (suite *addBorrowSuite) SetupTest() {

	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)

	app, k, nftKeeper, wctx := setupMsgServer(suite.T())
	ctx := sdk.UnwrapSDKContext(wctx)

	suite.ctx = ctx
	suite.keeper = k
	suite.nftKeeper = nftKeeper
	suite.app = app
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
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 2, PoolName: "hello", Apy: "7.8", TargetTokenAmount: sdk.NewCoin("ausdc", sdk.NewInt(1*1e3))}
	resp, err := suite.app.CreatePool(suite.ctx, &req)
	suite.Require().NoError(err)

	_, err = suite.app.ActivePool(suite.ctx, types.NewMsgActivePool("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", resp.PoolIndex[0]))
	suite.Require().NoError(err)

	_, err = suite.app.ActivePool(suite.ctx, types.NewMsgActivePool("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", resp.PoolIndex[1]))
	suite.Require().NoError(err)

	req2 := types.MsgAddInvestors{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: resp.PoolIndex[0],
		InvestorID: []string{"2"}}
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
			args: args{msgBorrow: &types.MsgBorrow{Creator: "jolt1m28h5mu57ugcpfw2sp5t9chdp69akzc6ze5r0j", PoolIndex: resp.PoolIndex[0]}, expectedErr: "not authorized to borrow money"},
		},

		{
			name: "inconsistency toekn denom",
			args: args{msgBorrow: &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: resp.PoolIndex[0], BorrowAmount: sdk.NewCoin("aaa", sdk.NewIntFromUint64(2233))}, expectedErr: "token to be borrowed is inconsistency"},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			_, err := suite.app.Borrow(suite.ctx, tc.args.msgBorrow)
			if tc.args.expectedErr != "" {
				suite.Require().ErrorContains(err, tc.args.expectedErr)
			} else {
				suite.Require().NoError(err)

			}
		})
	}
}

func compareDepositor(suite suite.Suite, expected, actual types.DepositorInfo) {

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
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 2, PoolName: "hello", Apy: "0.15", TargetTokenAmount: sdk.NewCoin("ausdc", sdk.NewInt(3*1e9))}
	resp, err := suite.app.CreatePool(suite.ctx, &req)
	suite.Require().NoError(err)

	depositorPool := resp.PoolIndex[0]

	_, err = suite.app.ActivePool(suite.ctx, types.NewMsgActivePool("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", resp.PoolIndex[0]))
	suite.Require().NoError(err)

	_, err = suite.app.ActivePool(suite.ctx, types.NewMsgActivePool("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", resp.PoolIndex[1]))
	suite.Require().NoError(err)

	req2 := types.MsgAddInvestors{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: resp.PoolIndex[0],
		InvestorID: []string{"2"}}
	_, err = suite.app.AddInvestors(suite.ctx, &req2)
	suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: depositorPool, BorrowAmount: sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1.34e5))}

	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().ErrorContains(err, "insufficient tokens")

	// now we deposit some token and it should be enough to borrow
	creator1 := "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl"
	creator2 := "jolt1kkujrm0lqeu0e5va5f6mmwk87wva0k8cmam8jq"
	creatorAddr1, err := sdk.AccAddressFromBech32(creator1)
	suite.Require().NoError(err)
	creatorAddr2, err := sdk.AccAddressFromBech32(creator2)
	suite.Require().NoError(err)
	depositAmount := sdk.NewCoin("ausdc", sdk.NewInt(4e5))
	suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{Creator: creator1,
		PoolIndex: depositorPool,
		Token:     depositAmount}

	// user two deposit half of the amount of the user 1
	msgDepositUser2 := &types.MsgDeposit{Creator: creator2,
		PoolIndex: depositorPool,
		Token:     depositAmount.SubAmount(sdk.NewInt(2e5))}

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
		LockedAmount:     sdk.NewCoin("ausdc", sdk.ZeroInt()),
		WithdrawalAmount: depositAmount,
		LinkedNFT:        []string{},
	}

	compareDepositor(suite.Suite, targetDepositor, depositor)
	// we deposit again,so withdrawal is doubled

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	targetDepositor.WithdrawalAmount = targetDepositor.WithdrawalAmount.Add(depositAmount)

	depositor, found = suite.keeper.GetDepositor(suite.ctx, depositorPool, creatorAddr1)
	suite.Require().True(found)
	compareDepositor(suite.Suite, targetDepositor, depositor)

	// we mock the second user deposits the token, now we have 3*4e5 tokens
	//_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	//suite.Require().NoError(err)

	pool, found := suite.keeper.GetPools(suite.ctx, depositorPool)
	suite.Require().True(found)

	totalBorrowable := msgDepositUser1.Token.Add(msgDepositUser1.Token).Add(msgDepositUser2.Token)
	suite.Require().True(totalBorrowable.IsEqual(pool.BorrowableAmount))

	user1Ratio := sdk.NewDecFromInt(msgDepositUser1.Token.Amount.Mul(sdk.NewInt(2))).Quo(sdk.NewDecFromInt(totalBorrowable.Amount))

	user2Ratio := sdk.NewDecFromInt(msgDepositUser2.Token.Amount).Quo(sdk.NewDecFromInt(totalBorrowable.Amount))

	//now we borrow 2e5
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	p1, found := suite.keeper.GetDepositor(suite.ctx, depositorPool, creatorAddr1)
	suite.Require().True(found)

	p2, found := suite.keeper.GetDepositor(suite.ctx, depositorPool, creatorAddr2)
	suite.Require().True(found)

	poolNow, found := suite.keeper.GetPools(suite.ctx, depositorPool)
	suite.Require().True(found)
	suite.Require().True(poolNow.BorrowedAmount.IsEqual(borrow.BorrowAmount))
	suite.Require().True(totalBorrowable.Sub(borrow.BorrowAmount).IsEqual(poolNow.BorrowableAmount))

	borrowedFromUser1 := sdk.NewDecFromInt(borrow.BorrowAmount.Amount).Mul(user1Ratio).TruncateInt()
	borrowedFromUser2 := borrow.BorrowAmount.Amount.Sub(borrowedFromUser1)

	borrowedFromUser2Ratio := sdk.NewDecFromInt(borrow.BorrowAmount.Amount).Mul(user2Ratio).TruncateInt()

	suite.Require().True(p1.LockedAmount.Amount.Equal(borrowedFromUser1))
	suite.Require().True(p2.LockedAmount.Amount.Equal(borrowedFromUser2))
	suite.Require().True(borrowedFromUser2Ratio.Equal(borrowedFromUser2))

	// total amount shoube be locked+withdrawable
	suite.Require().True(p1.LockedAmount.Add(p1.WithdrawalAmount).IsEqual(msgDepositUser1.Token.Add(msgDepositUser1.Token)))

	nftUser1 := p1.LinkedNFT[0]
	nftUser2 := p2.LinkedNFT[0]

	nftClassID := fmt.Sprintf("nft-%v-0", depositorPool[2:])
	nft, found := suite.nftKeeper.GetClass(suite.ctx, nftClassID)
	suite.Require().True(found)

	var borrowClassInfo types.BorrowInterest
	err = proto.Unmarshal(nft.Data.Value, &borrowClassInfo)
	if err != nil {
		panic(err)
	}

	suite.True(borrowClassInfo.Borrowed.IsEqual(borrow.BorrowAmount))
	suite.True(borrowClassInfo.BorrowedLast.IsEqual(borrow.BorrowAmount))
	fmt.Printf(">>>>%v\n", borrowClassInfo.Apy)
	suite.True(borrowClassInfo.Apy.Equal(sdk.NewDecWithPrec(15, 2)))

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
	suite.Require().True(nftInfo.Ratio.Equal(user1Ratio))

	// now, user 2 deposits more money and then, the spv borrow more. the ratio should  be changed.
	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	user1Cached := p1.WithdrawalAmount
	user2Cached := p2.WithdrawalAmount

	p2, found = suite.keeper.GetDepositor(suite.ctx, depositorPool, creatorAddr2)

	suite.Require().True(found)
	suite.Require().True(p2.WithdrawalAmount.IsEqual(user2Cached.AddAmount(msgDepositUser2.Token.Amount)))
	suite.Require().True(poolNow.GetBorrowableAmount().IsEqual(user1Cached.Add(user2Cached)))
	user2Cached = p2.WithdrawalAmount
	newuser1Ratio1 := sdk.NewDecFromInt(user1Cached.Amount).Quo(sdk.NewDecFromInt(user1Cached.Add(user2Cached).Amount))

	poolAfterUser2SecondDeposit, found := suite.keeper.GetPools(suite.ctx, depositorPool)
	suite.Require().True(found)
	suite.Require().True(poolAfterUser2SecondDeposit.BorrowableAmount.IsEqual(user2Cached.Add(user1Cached)))

	// NOW we borrow
	borrow.BorrowAmount = sdk.NewCoin(borrow.BorrowAmount.Denom, sdk.NewInt(1.2e5))
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	previousAmountBorrowed := poolNow.BorrowedAmount
	previousBorrowAble := poolNow.BorrowableAmount
	poolNow, found = suite.keeper.GetPools(suite.ctx, depositorPool)
	suite.Require().True(found)

	suite.Require().True(poolNow.BorrowedAmount.Equal(borrow.BorrowAmount.AddAmount(previousAmountBorrowed.Amount)))

	suite.Require().True(poolNow.BorrowableAmount.Equal(previousBorrowAble.Add(msgDepositUser2.Token).Sub(borrow.BorrowAmount)))

	beforeLockedAmount := p1.LockedAmount
	// now we check the nfts
	p1, found = suite.keeper.GetDepositor(suite.ctx, depositorPool, creatorAddr1)
	suite.Require().True(found)

	p2, found = suite.keeper.GetDepositor(suite.ctx, depositorPool, creatorAddr2)
	suite.Require().True(found)

	lockedThistime := p1.LockedAmount.Sub(beforeLockedAmount)
	shouldLocked := newuser1Ratio1.Mul(sdk.NewDecFromInt(borrow.BorrowAmount.Amount)).TruncateInt()
	suite.Require().True(lockedThistime.Amount.Equal(shouldLocked))

	// we check the total deposit of the user1 is correct
	suite.Require().True(p1.LockedAmount.Add(p1.WithdrawalAmount).IsEqual(msgDepositUser1.Token.Add(msgDepositUser1.Token)))

	nft2User1 := p1.LinkedNFT[1]
	nft2User2 := p2.LinkedNFT[1]

	nftClassID = fmt.Sprintf("nft-%v-1", depositorPool[2:])
	nft, found = suite.nftKeeper.GetClass(suite.ctx, nftClassID)
	suite.Require().True(found)

	err = proto.Unmarshal(nft.Data.Value, &borrowClassInfo)
	if err != nil {
		panic(err)
	}

	suite.True(borrowClassInfo.Borrowed.IsEqual(borrow.BorrowAmount))
	suite.True(borrowClassInfo.BorrowedLast.IsEqual(borrow.BorrowAmount))

	//nft ID is the hash(nft class ID, investorWallet)
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
	shoudRatio := sdk.NewDecFromInt(shouldLocked).QuoInt(borrow.BorrowAmount.Amount)
	suite.Require().True(nftInfo.Ratio.Equal(shoudRatio))

}
