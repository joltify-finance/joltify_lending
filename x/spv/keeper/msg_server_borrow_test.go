package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/suite"
	"testing"
)

// Test suite used for all keeper tests
type addBorrowSuite struct {
	suite.Suite
	keeper     *spvkeeper.Keeper
	app        types.MsgServer
	ctx        sdk.Context
	poolIndexs []string
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

	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 1, PoolName: "hello", Apy: "7.8", TargetTokenAmount: sdk.NewCoin("ausdc", sdk.NewInt(322))}
	resp, err := app.CreatePool(ctx, &req)
	suite.Require().NoError(err)
	suite.poolIndexs = resp.PoolIndex
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

	testCases := []test{
		{
			name: "invalid address",
			args: args{msgBorrow: &types.MsgBorrow{Creator: "invalid address"}, expectedErr: "invalid address invalid address: invalid address"},
		},

		{
			name: "pool cannot be found",
			args: args{msgBorrow: &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0"}, expectedErr: "pool not found with"},
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
