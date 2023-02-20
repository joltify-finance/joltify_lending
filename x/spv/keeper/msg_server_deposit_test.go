package keeper_test

import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/suite"
	"testing"
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

	app, k, wctx := setupMsgServer(suite.T())
	ctx := sdk.UnwrapSDKContext(wctx)

	suite.ctx = ctx
	suite.keeper = k
	suite.app = app
}

func (suite *DepositTestSuite) TestRepay() {

	type args struct {
		msgDeposit  *types.MsgDeposit
		expectedErr error
	}

	type test struct {
		name string
		args args
	}

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
			name: "not whitelist investor",
			args: args{msgDeposit: &types.MsgDeposit{Creator: "jolt1m28h5mu57ugcpfw2sp5t9chdp69akzc6ze5r0j", PoolIndex: "0x230365e8dc1e9b0e080e8986a3168b60c3d26ebe5f6746a00f5abd1fd62e8461"}, expectedErr: errors.New("pool cannot be found : not found")},
		},
	}

	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 1, PoolName: "hello", Apy: "7.8", PayFreq: "6", TargetTokenAmount: sdk.NewCoin("demo", sdk.NewInt(322))}
	_, err := suite.app.CreatePool(suite.ctx, &req)
	suite.Require().NoError(err)

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
