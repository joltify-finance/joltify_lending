package issuance_test

import (
	"testing"
	"time"

	tmlog "github.com/tendermint/tendermint/libs/log"

	"github.com/joltify-finance/joltify_lending/x/third_party/issuance"
	"github.com/joltify-finance/joltify_lending/x/third_party/issuance/keeper"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/issuance/types"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/joltify-finance/joltify_lending/app"
)

// Test suite used for all keeper tests
type ABCITestSuite struct {
	suite.Suite

	keeper     keeper.Keeper
	app        app.TestApp
	ctx        sdk.Context
	addrs      []sdk.AccAddress
	modAccount sdk.AccAddress
	blockTime  time.Time
}

// The default state used by each test
func (suite *ABCITestSuite) SetupTest() {
	tApp := app.NewTestApp(tmlog.TestingLogger(), suite.T().TempDir())
	blockTime := tmtime.Now()
	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: blockTime})
	tApp.InitializeFromGenesisStates(nil, nil)
	_, addrs := app.GeneratePrivKeyAddressPairs(5)
	keeper := tApp.GetIssuanceKeeper()
	modAccount, err := sdk.AccAddressFromBech32("jolt10wlnqzyss4accfqmyxwx5jy5x9nfkwh6kwhjs9")
	suite.Require().NoError(err)
	suite.app = tApp
	suite.ctx = ctx
	suite.keeper = keeper
	suite.addrs = addrs
	suite.modAccount = modAccount
	suite.blockTime = blockTime
}

func (suite *ABCITestSuite) TestRateLimitingTimePassage() {
	type args struct {
		assets         []types2.Asset
		supplies       []types2.AssetSupply
		blockTimes     []time.Duration
		expectedSupply types2.AssetSupply
	}
	testCases := []struct {
		name string
		args args
	}{
		{
			"time passage same period",
			args{
				assets: []types2.Asset{
					types2.NewAsset(suite.addrs[0].String(), "usdtoken", []string{suite.addrs[1].String()}, false, true, types2.NewRateLimit(true, sdk.NewInt(10000000000), time.Hour*24)),
				},
				supplies: []types2.AssetSupply{
					types2.NewAssetSupply(sdk.NewCoin("usdtoken", sdk.ZeroInt()), time.Hour),
				},
				blockTimes:     []time.Duration{time.Hour},
				expectedSupply: types2.NewAssetSupply(sdk.NewCoin("usdtoken", sdk.ZeroInt()), time.Hour*2),
			},
		},
		{
			"time passage new period",
			args{
				assets: []types2.Asset{
					types2.NewAsset(suite.addrs[0].String(), "usdtoken", []string{suite.addrs[1].String()}, false, true, types2.NewRateLimit(true, sdk.NewInt(10000000000), time.Hour*24)),
				},
				supplies: []types2.AssetSupply{
					types2.NewAssetSupply(sdk.NewCoin("usdtoken", sdk.ZeroInt()), time.Hour),
				},
				blockTimes:     []time.Duration{time.Hour * 12, time.Hour * 12},
				expectedSupply: types2.NewAssetSupply(sdk.NewCoin("usdtoken", sdk.ZeroInt()), time.Duration(0)),
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			params := types2.NewParams(tc.args.assets)
			suite.keeper.SetParams(suite.ctx, params)
			for _, supply := range tc.args.supplies {
				suite.keeper.SetAssetSupply(suite.ctx, supply, supply.GetDenom())
			}
			suite.keeper.SetPreviousBlockTime(suite.ctx, suite.blockTime)
			for _, bt := range tc.args.blockTimes {
				nextBlockTime := suite.ctx.BlockTime().Add(bt)
				suite.ctx = suite.ctx.WithBlockTime(nextBlockTime)
				suite.Require().NotPanics(func() {
					issuance.BeginBlocker(suite.ctx, suite.keeper)
				})
			}
			actualSupply, found := suite.keeper.GetAssetSupply(suite.ctx, tc.args.expectedSupply.GetDenom())
			suite.Require().True(found)
			suite.Require().Equal(tc.args.expectedSupply, actualSupply)
		})
	}
}

func TestABCITestSuite(t *testing.T) {
	suite.Run(t, new(ABCITestSuite))
}
