package keeper_test

import (
	"errors"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/suite"
)

// Test suite used for all keeper tests
type DepositTestSuite struct {
	suite.Suite
	keeper *spvkeeper.Keeper
	app    types.MsgServer
	ctx    sdk.Context
}

func TestDepositTestSuite(t *testing.T) {
	suite.Run(t, new(DepositTestSuite))
}

// The default state used by each test
func (suite *DepositTestSuite) SetupTest() {

	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)

	app, k, _, wctx := setupMsgServer(suite.T())
	ctx := sdk.UnwrapSDKContext(wctx)

	// create the first pool apy 7.8%

	suite.ctx = ctx
	suite.keeper = k
	suite.app = app
}

func (suite *DepositTestSuite) TestDeposit() {

	type args struct {
		msgDeposit  *types.MsgDeposit
		expectedErr error
	}

	type test struct {
		name string
		args args
	}

	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 2, PoolName: "hello", Apy: "7.8", TargetTokenAmount: sdk.NewCoin("ausdc", sdk.NewInt(3*1e9))}
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
			args: args{msgDeposit: &types.MsgDeposit{Creator: "invalid address"}, expectedErr: errors.New("invalid address invalid address: invalid address")},
		},

		{
			name: "pool cannot be found",
			args: args{msgDeposit: &types.MsgDeposit{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0"}, expectedErr: errors.New("pool cannot be found : not found")},
		},

		{
			name: "pool is full",
			args: args{msgDeposit: &types.MsgDeposit{Creator: "jolt1m28h5mu57ugcpfw2sp5t9chdp69akzc6ze5r0j",
				PoolIndex: resp.PoolIndex[0],
				Token:     sdk.NewCoin("ausdc", sdk.NewInt(3e9))},
				expectedErr: errors.New("pool is full")},
		},

		{
			name: "not on white list",
			args: args{msgDeposit: &types.MsgDeposit{Creator: "jolt1m28h5mu57ugcpfw2sp5t9chdp69akzc6ze5r0j",
				PoolIndex: resp.PoolIndex[1],
				Token:     sdk.NewCoin("ausdc", sdk.NewInt(100))},
				expectedErr: errors.New("the given investor is not allowed to invest jolt1m28h5mu57ugcpfw2sp5t9chdp69akzc6ze5r0j: unauthorized operation")},
		},

		{
			name: "can deposit as expected",
			args: args{msgDeposit: &types.MsgDeposit{Creator: "jolt1kkujrm0lqeu0e5va5f6mmwk87wva0k8cmam8jq",
				PoolIndex: resp.PoolIndex[1],
				Token:     sdk.NewCoin("ausdc", sdk.NewInt(100))},
				expectedErr: nil,
			},
		},
		{
			name: "can deposit as expected with the second investor",
			args: args{msgDeposit: &types.MsgDeposit{Creator: "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl",
				PoolIndex: resp.PoolIndex[1],
				Token:     sdk.NewCoin("ausdc", sdk.NewInt(100))},
				expectedErr: nil,
			},
		},

		{
			name: "incorrect denom",
			args: args{msgDeposit: &types.MsgDeposit{Creator: "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl",
				PoolIndex: resp.PoolIndex[1],
				Token:     sdk.NewCoin("usdd", sdk.NewInt(100))},
				expectedErr: errors.New("we only accept ausdc: invalid coins"),
			},
		},

		{
			name: "can deposit the second time",
			args: args{msgDeposit: &types.MsgDeposit{Creator: "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl",
				PoolIndex: resp.PoolIndex[1],
				Token:     sdk.NewCoin("ausdc", sdk.NewInt(100))},
				expectedErr: nil,
			},
		},

		{
			name: "can deposit in both pools",
			args: args{msgDeposit: &types.MsgDeposit{Creator: "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl",
				PoolIndex: resp.PoolIndex[0],
				Token:     sdk.NewCoin("ausdc", sdk.NewInt(100))},
				expectedErr: nil,
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			_, err := suite.app.Deposit(suite.ctx, tc.args.msgDeposit)
			if tc.args.expectedErr != nil {
				suite.Require().Equal(tc.args.expectedErr.Error(), err.Error())
			} else {
				suite.Require().NoError(err)

			}
		})
	}
}

func (suite *DepositTestSuite) TestDepositWithAmountCorrect() {

	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 1, PoolName: "hello", Apy: "7.8", TargetTokenAmount: sdk.NewCoin("usdc", sdk.NewInt(0))}
	_, err := suite.app.CreatePool(suite.ctx, &req)
	suite.Require().NoError(err)

	// create the first pool apy 8.8%
	req = types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 1, PoolName: "hello", Apy: "8.8", TargetTokenAmount: sdk.NewCoin("usdc", sdk.NewInt(322))}
	_, err = suite.app.CreatePool(suite.ctx, &req)
	suite.Require().NoError(err)

	_, err = suite.app.ActivePool(suite.ctx, types.NewMsgActivePool("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", "0x86a7506e61dfab773c243762f636ea428a5f497ba69205729a12dc428ce4abf3"))
	suite.Require().NoError(err)

	_, err = suite.app.ActivePool(suite.ctx, types.NewMsgActivePool("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", "0xcac0266ccc0bedb38f3ccc4da1edf884c8d3960497e9a038802fb73c5c0e18bc"))
	suite.Require().NoError(err)

	req2 := types.MsgAddInvestors{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: "0x86a7506e61dfab773c243762f636ea428a5f497ba69205729a12dc428ce4abf3",
		InvestorID: []string{"2"}}
	_, err = suite.app.AddInvestors(suite.ctx, &req2)
	suite.Require().NoError(err)

	pool, found := suite.keeper.GetPools(suite.ctx, "0x86a7506e61dfab773c243762f636ea428a5f497ba69205729a12dc428ce4abf3")
	suite.Require().True(found)
	suite.Require().True(pool.TargetAmount.Equal(sdk.NewCoin("usdc", sdk.NewInt(322))))
	suite.Require().True(pool.TotalAmount.Equal(sdk.NewCoin("usdc", sdk.NewInt(0))))

	suite.Require().True(pool.BorrowedAmount.Equal(sdk.NewCoin("usdc", sdk.NewInt(0))))
	suite.Require().True(pool.BorrowableAmount.Equal(sdk.NewCoin("usdc", sdk.NewInt(0))))

	depositAmount := sdk.NewCoin("usdc", sdk.NewInt(100))
	msgDepositor := types.MsgDeposit{Creator: "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl",
		PoolIndex: "0x86a7506e61dfab773c243762f636ea428a5f497ba69205729a12dc428ce4abf3",
		Token:     depositAmount}

	_, err = suite.app.Deposit(suite.ctx, &msgDepositor)
	suite.Require().NoError(err)

	pool, found = suite.keeper.GetPools(suite.ctx, "0x86a7506e61dfab773c243762f636ea428a5f497ba69205729a12dc428ce4abf3")
	suite.Require().True(found)
	suite.Require().True(pool.TargetAmount.Equal(sdk.NewCoin("usdc", sdk.NewInt(322))))
	suite.Require().True(pool.TotalAmount.Equal(depositAmount))

	suite.Require().True(pool.BorrowedAmount.Equal(sdk.NewCoin("usdc", sdk.NewInt(0))))
	suite.Require().True(pool.BorrowableAmount.Equal(depositAmount))

	depositerAddr, err := sdk.AccAddressFromBech32("jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl")
	suite.Require().NoError(err)

	depositorData, found := suite.keeper.GetDepositor(suite.ctx, "0x86a7506e61dfab773c243762f636ea428a5f497ba69205729a12dc428ce4abf3", depositerAddr)
	suite.Require().True(found)

	suite.Require().True(depositorData.LockedAmount.Equal(sdk.NewCoin("usdc", sdk.NewInt(0))))

	suite.Require().True(depositorData.WithdrawalAmount.Equal(depositAmount))

	// we deposit more money

	_, err = suite.app.Deposit(suite.ctx, &msgDepositor)
	suite.Require().NoError(err)

	pool, found = suite.keeper.GetPools(suite.ctx, "0x86a7506e61dfab773c243762f636ea428a5f497ba69205729a12dc428ce4abf3")
	suite.Require().True(found)
	suite.Require().True(pool.TargetAmount.Equal(sdk.NewCoin("usdc", sdk.NewInt(322))))
	suite.Require().True(pool.TotalAmount.Equal(depositAmount.Add(depositAmount)))

	suite.Require().True(pool.BorrowedAmount.Equal(sdk.NewCoin("usdc", sdk.NewInt(0))))
	suite.Require().True(pool.BorrowableAmount.Equal(depositAmount.Add(depositAmount)))

	depositorData, found = suite.keeper.GetDepositor(suite.ctx, "0x86a7506e61dfab773c243762f636ea428a5f497ba69205729a12dc428ce4abf3", depositerAddr)
	suite.Require().True(found)

	suite.Require().True(depositorData.LockedAmount.Equal(sdk.NewCoin("ausdc", sdk.NewInt(0))))

	suite.Require().True(depositorData.WithdrawalAmount.Equal(depositAmount.Add(depositAmount)))

}
