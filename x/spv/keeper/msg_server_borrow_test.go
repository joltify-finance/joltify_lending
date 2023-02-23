package keeper_test

import (
	"testing"

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
	keeper *spvkeeper.Keeper
	app    types.MsgServer
	ctx    sdk.Context
}

func TestBorrowTestSuite(t *testing.T) {
	suite.Run(t, new(addBorrowSuite))
}

// The default state used by each test
func (suite *addBorrowSuite) SetupTest() {

	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)

	app, k, wctx := setupMsgServer(suite.T())
	ctx := sdk.UnwrapSDKContext(wctx)

	suite.ctx = ctx
	suite.keeper = k
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

func (suite *addBorrowSuite) TestBorrowValueCheck() {

	// create the first pool apy 7.8%
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

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: resp.PoolIndex[0], BorrowAmount: sdk.NewCoin("ausdc", sdk.NewIntFromUint64(2233))}

	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().ErrorContains(err, "insufficient tokens")

	// now we deposit some token and it should be enough to borrow
	msgDeposit := &types.MsgDeposit{Creator: "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl",
		PoolIndex: resp.PoolIndex[1],
		Token:     sdk.NewCoin("ausdc", sdk.NewInt(4e5))}

	_, err = suite.app.Deposit(suite.ctx, msgDeposit)
	suite.Require().NoError(err)

	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

}
