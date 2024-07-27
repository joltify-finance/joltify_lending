package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/suite"
)

// Test suite used for all keeper tests
type withdrawProposalSuite struct {
	suite.Suite
	keeper       *spvkeeper.Keeper
	nftKeeper    types.NFTKeeper
	app          types.MsgServer
	ctx          context.Context
	investorPool string
	investors    []string
}

func TestWithdrawProposalSuite(t *testing.T) {
	suite.Run(t, new(withdrawProposalSuite))
}

// The default state used by each test
func (suite *withdrawProposalSuite) SetupTest() {
	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)

	lapp, k, nftKeeper, _, _, wctx := setupMsgServer(suite.T())
	ctx := sdk.UnwrapSDKContext(wctx)

	suite.ctx = ctx
	suite.keeper = k
	suite.nftKeeper = nftKeeper
	suite.app = lapp
}

func setupWithdrawProposal(suite *withdrawProposalSuite) {
	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 3, PoolName: "hello", Apy: []string{"0.15", "0.15"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdkmath.NewInt(3*1e9)), sdk.NewCoin("ausdc", sdkmath.NewInt(3*1e9))}}
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

	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(4e5))
	// suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{
		Creator:   creator1,
		PoolIndex: suite.investorPool,
		Token:     depositAmount,
	}

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	poolInfo.PoolTotalBorrowLimit = 100
	poolInfo.TargetAmount = sdk.NewCoin("ausdc", sdkmath.NewInt(4e5))
	suite.keeper.SetPool(suite.ctx, poolInfo)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: depositorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	suite.investors = []string{creator1, creator2}
}

func (suite *withdrawProposalSuite) TestWithdrawProposal() {
	setupWithdrawProposal(suite)

	type args struct {
		msgWithdraw *types.MsgSubmitWithdrawProposal
		expectedErr string
	}

	type test struct {
		name string
		args args
	}

	testCases := []test{
		{
			name: "invalid address",
			args: args{msgWithdraw: &types.MsgSubmitWithdrawProposal{Creator: "invalid address"}, expectedErr: "invalid address invalid address: invalid address"},
		},

		{
			name: "pool cannot be found",
			args: args{msgWithdraw: &types.MsgSubmitWithdrawProposal{Creator: suite.investors[0], PoolIndex: "invalid"}, expectedErr: "not found for pool index"},
		},
		{
			name: "pool cannot be found",
			args: args{msgWithdraw: &types.MsgSubmitWithdrawProposal{Creator: suite.investors[1], PoolIndex: suite.investorPool}, expectedErr: "not found for pool index"},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			_, err := suite.app.SubmitWithdrawProposal(suite.ctx, tc.args.msgWithdraw)
			if tc.args.expectedErr != "" {
				suite.Require().ErrorContains(err, tc.args.expectedErr)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *withdrawProposalSuite) TestWithdrawProposalTooEarlyOrLate() {
	setupWithdrawProposal(suite)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	currentBlockTime := suite.ctx.BlockTime()

	ctx1 := suite.ctx.WithBlockTime(currentBlockTime.Add(time.Second * time.Duration(poolInfo.ProjectLength-uint64(poolInfo.WithdrawRequestWindowSeconds)*3-1)))
	req := types.MsgSubmitWithdrawProposal{Creator: suite.investors[0], PoolIndex: suite.investorPool}
	_, err := suite.app.SubmitWithdrawProposal(ctx1, &req)
	suite.Require().ErrorContains(err, "submit the proposal too early")

	ctx2 := suite.ctx.WithBlockTime(currentBlockTime.Add(time.Second * time.Duration(poolInfo.ProjectLength-uint64(poolInfo.WithdrawRequestWindowSeconds)*2+1)))
	_, err = suite.app.SubmitWithdrawProposal(ctx2, &req)
	suite.Require().ErrorContains(err, "submit the proposal too late")

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.Require().Len(poolInfo.WithdrawAccounts, 0)
	suite.Require().True(poolInfo.WithdrawProposalAmount.Amount.IsZero())
	withdrawable := poolInfo.UsableAmount
	borrowed := poolInfo.BorrowedAmount
	ctx3 := suite.ctx.WithBlockTime(currentBlockTime.Add(time.Second * time.Duration(poolInfo.ProjectLength-uint64(poolInfo.WithdrawRequestWindowSeconds)*3)))
	_, err = suite.app.SubmitWithdrawProposal(ctx3, &req)
	suite.Require().NoError(err)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	acc, err := sdk.AccAddressFromBech32(suite.investors[0])
	suite.Require().NoError(err)

	depositor, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, acc)
	suite.Require().True(found)

	suite.Require().EqualValues([]sdk.AccAddress{acc}, poolInfo.WithdrawAccounts)
	suite.Require().True(poolInfo.WithdrawProposalAmount.IsEqual(depositor.LockedAmount))

	suite.Require().Equal(poolInfo.UsableAmount.String(), withdrawable.Sub(depositor.WithdrawalAmount).String())
	suite.Require().Equal(poolInfo.BorrowedAmount.String(), borrowed.String())

	ctx4 := suite.ctx.WithBlockTime(currentBlockTime.Add(time.Second * time.Duration(poolInfo.ProjectLength-spvkeeper.OneMonth-1)))
	_, err = suite.app.SubmitWithdrawProposal(ctx4, &req)
	suite.Require().ErrorContains(err, "is not in unset status (current status withdraw_proposal")
}
