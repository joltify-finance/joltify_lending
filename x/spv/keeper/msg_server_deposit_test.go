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

	app, k, _, _, wctx := setupMsgServer(suite.T())
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

	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 2, PoolName: "hello", Apy: []string{"7.8", "7.2"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdk.NewInt(3*1e9)), sdk.NewCoin("ausdc", sdk.NewInt(3*1e9))}}
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
			args: args{msgDeposit: &types.MsgDeposit{Creator: "invalid address"}, expectedErr: errors.New("invalid address invalid address: invalid address")},
		},

		{
			name: "pool cannot be found",
			args: args{msgDeposit: &types.MsgDeposit{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0"}, expectedErr: errors.New("pool cannot be found : not found")},
		},

		{
			name: "not on white list",
			args: args{
				msgDeposit: &types.MsgDeposit{
					Creator:   "jolt1m28h5mu57ugcpfw2sp5t9chdp69akzc6ze5r0j",
					PoolIndex: resp.PoolIndex[1],
					Token:     sdk.NewCoin("ausdc", sdk.NewInt(100)),
				},
				expectedErr: errors.New("the given investor is not allowed to invest jolt1m28h5mu57ugcpfw2sp5t9chdp69akzc6ze5r0j: unauthorized operation"),
			},
		},

		{
			name: "can deposit as expected",
			args: args{
				msgDeposit: &types.MsgDeposit{
					Creator:   "jolt1kkujrm0lqeu0e5va5f6mmwk87wva0k8cmam8jq",
					PoolIndex: resp.PoolIndex[1],
					Token:     sdk.NewCoin("ausdc", sdk.NewInt(100)),
				},
				expectedErr: nil,
			},
		},
		{
			name: "can deposit as expected with the second investor",
			args: args{
				msgDeposit: &types.MsgDeposit{
					Creator:   "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl",
					PoolIndex: resp.PoolIndex[1],
					Token:     sdk.NewCoin("ausdc", sdk.NewInt(100)),
				},
				expectedErr: nil,
			},
		},

		{
			name: "incorrect denom",
			args: args{
				msgDeposit: &types.MsgDeposit{
					Creator:   "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl",
					PoolIndex: resp.PoolIndex[1],
					Token:     sdk.NewCoin("usdd", sdk.NewInt(100)),
				},
				expectedErr: errors.New("we only accept ausdc: invalid coins"),
			},
		},

		{
			name: "can deposit the second time",
			args: args{
				msgDeposit: &types.MsgDeposit{
					Creator:   "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl",
					PoolIndex: resp.PoolIndex[1],
					Token:     sdk.NewCoin("ausdc", sdk.NewInt(100)),
				},
				expectedErr: nil,
			},
		},

		{
			name: "can deposit in both pools",
			args: args{
				msgDeposit: &types.MsgDeposit{
					Creator:   "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl",
					PoolIndex: resp.PoolIndex[0],
					Token:     sdk.NewCoin("ausdc", sdk.NewInt(100)),
				},
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
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 1, PoolName: "hello", Apy: []string{"7.8", "7.2"}, TargetTokenAmount: sdk.NewCoins(sdk.NewCoin("usdc", sdk.NewInt(0)), sdk.NewCoin("usdc", sdk.NewInt(0)))}
	_, err := suite.app.CreatePool(suite.ctx, &req)
	suite.Require().ErrorContains(err, "the amount cannot be 0")

	// create the first pool apy 8.8%
	req = types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 1, PoolName: "hello", Apy: []string{"8.8", "7.2"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdk.NewInt(322)), sdk.NewCoin("ausdc", sdk.NewInt(322))}}
	resp, err := suite.app.CreatePool(suite.ctx, &req)
	suite.Require().NoError(err)

	_, err = suite.app.ActivePool(suite.ctx, types.NewMsgActivePool("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", resp.PoolIndex[0]))
	suite.Require().NoError(err)

	_, err = suite.app.ActivePool(suite.ctx, types.NewMsgActivePool("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", resp.PoolIndex[1]))
	suite.Require().NoError(err)

	req2 := types.MsgAddInvestors{
		Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: resp.PoolIndex[1],
		InvestorID: []string{"2"},
	}
	_, err = suite.app.AddInvestors(suite.ctx, &req2)
	suite.Require().NoError(err)

	pool, found := suite.keeper.GetPools(suite.ctx, resp.PoolIndex[0])
	suite.Require().True(found)
	// suite.Require().True(pool.TargetAmount.Equal(sdk.NewCoin("ausdc", sdk.NewInt(322))))
	suite.Require().True(pool.BorrowedAmount.Equal(sdk.NewCoin("aud-ausdc", sdk.NewInt(0))))
	suite.Require().True(pool.UsableAmount.Equal(sdk.NewCoin("ausdc", sdk.NewInt(0))))

	depositAmount := sdk.NewCoin("ausdc", sdk.NewInt(100))
	msgDepositor := types.MsgDeposit{
		Creator:   "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl",
		PoolIndex: resp.PoolIndex[0],
		Token:     depositAmount,
	}

	_, err = suite.app.Deposit(suite.ctx, &msgDepositor)
	suite.Require().NoError(err)

	pool, found = suite.keeper.GetPools(suite.ctx, resp.PoolIndex[0])
	suite.Require().True(found)
	suite.Require().True(checkValueWithRangeTwo(pool.TargetAmount.Amount, sdk.NewCoin("ausdc", sdk.NewInt(322)).Amount))

	suite.Require().True(pool.BorrowedAmount.Equal(sdk.NewCoin("aud-ausdc", sdk.NewInt(0))))
	suite.Require().True(pool.UsableAmount.Equal(depositAmount))

	depositerAddr, err := sdk.AccAddressFromBech32("jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl")
	suite.Require().NoError(err)

	depositorData, found := suite.keeper.GetDepositor(suite.ctx, resp.PoolIndex[0], depositerAddr)
	suite.Require().True(found)

	suite.Require().True(depositorData.LockedAmount.Equal(sdk.NewCoin("aud-ausdc", sdk.NewInt(0))))

	suite.Require().True(depositorData.WithdrawalAmount.Equal(depositAmount))

	// we deposit more money

	_, err = suite.app.Deposit(suite.ctx, &msgDepositor)
	suite.Require().NoError(err)

	pool, found = suite.keeper.GetPools(suite.ctx, resp.PoolIndex[0])
	suite.Require().True(found)
	suite.Require().True(pool.TargetAmount.Equal(sdk.NewCoin("ausdc", sdk.NewInt(322))))

	suite.Require().True(pool.BorrowedAmount.Equal(sdk.NewCoin("aud-ausdc", sdk.NewInt(0))))
	suite.Require().True(pool.UsableAmount.Equal(depositAmount.Add(depositAmount)))

	depositorData, found = suite.keeper.GetDepositor(suite.ctx, resp.PoolIndex[0], depositerAddr)
	suite.Require().True(found)

	suite.Require().True(depositorData.LockedAmount.Equal(sdk.NewCoin("aud-ausdc", sdk.NewInt(0))))

	suite.Require().True(depositorData.WithdrawalAmount.Equal(depositAmount.Add(depositAmount)))
}
