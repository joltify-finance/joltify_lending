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
type addInvestorSuite struct {
	suite.Suite
	keeper     *spvkeeper.Keeper
	app        types.MsgServer
	ctx        sdk.Context
	poolIndexs []string
}

func TestAddInvestorTestSuite(t *testing.T) {
	suite.Run(t, new(addInvestorSuite))
}

// The default state used by each test
func (suite *addInvestorSuite) SetupTest() {
	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)

	app, k, _, _, wctx := setupMsgServer(suite.T())
	ctx := sdk.UnwrapSDKContext(wctx)

	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 1, PoolName: "hello", Apy: []string{"7.8", "7.2"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdk.NewInt(322)), sdk.NewCoin("ausdc", sdk.NewInt(322))}}
	resp, err := app.CreatePool(ctx, &req)
	suite.Require().NoError(err)
	suite.poolIndexs = resp.PoolIndex
	suite.ctx = ctx
	suite.keeper = k
	suite.app = app
}

func (suite *addInvestorSuite) TestAddInvestor() {
	type args struct {
		msgAddInvestor *types.MsgAddInvestors
		expectedErr    string
	}

	type test struct {
		name string
		args args
	}

	testCases := []test{
		{
			name: "invalid address",
			args: args{msgAddInvestor: &types.MsgAddInvestors{Creator: "invalid address"}, expectedErr: "invalid address invalid address: invalid address"},
		},

		{
			name: "pool cannot be found",
			args: args{msgAddInvestor: &types.MsgAddInvestors{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0"}, expectedErr: "pool not found with"},
		},

		{
			name: "not the pool owner",
			args: args{msgAddInvestor: &types.MsgAddInvestors{Creator: "jolt1kkujrm0lqeu0e5va5f6mmwk87wva0k8cmam8jq", PoolIndex: suite.poolIndexs[0]}, expectedErr: "unauthorized operations: unauthorized operation"},
		},

		{
			name: "add valid address",
			args: args{msgAddInvestor: &types.MsgAddInvestors{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.poolIndexs[0], InvestorID: []string{"123", "324"}}, expectedErr: ""},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			_, err := suite.app.AddInvestors(suite.ctx, tc.args.msgAddInvestor)
			if tc.args.expectedErr != "" {
				suite.Require().ErrorContains(err, tc.args.expectedErr)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *addInvestorSuite) TestAddDuplicateInvestors() {
	req := &types.MsgAddInvestors{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.poolIndexs[0], InvestorID: []string{"123", "324"}}

	_, err := suite.app.AddInvestors(suite.ctx, req)
	suite.Require().NoError(err)

	r, found := suite.keeper.GetInvestorToPool(suite.ctx, suite.poolIndexs[0])
	suite.Require().True(found)
	suite.Require().EqualValues(r.Investors, []string{"123", "324"})

	req.InvestorID = []string{"324", "444"}
	_, err = suite.app.AddInvestors(suite.ctx, req)
	suite.Require().NoError(err)

	r, found = suite.keeper.GetInvestorToPool(suite.ctx, suite.poolIndexs[0])
	suite.Require().True(found)

	suite.Require().EqualValues(r.Investors, []string{"123", "324", "444"})

	req.InvestorID = []string{"324", "444"}
	_, err = suite.app.AddInvestors(suite.ctx, req)
	suite.Require().NoError(err)

	r, found = suite.keeper.GetInvestorToPool(suite.ctx, suite.poolIndexs[0])
	suite.Require().True(found)

	suite.Require().EqualValues(r.Investors, []string{"123", "324", "444"})

	req.InvestorID = []string{"555"}
	_, err = suite.app.AddInvestors(suite.ctx, req)
	suite.Require().NoError(err)

	r, found = suite.keeper.GetInvestorToPool(suite.ctx, suite.poolIndexs[0])
	suite.Require().True(found)

	suite.Require().EqualValues(r.Investors, []string{"123", "324", "444", "555"})
}
