package keeper_test

import (
	"context"
	"testing"
	"time"

	"cosmossdk.io/log"
	appconfig "github.com/joltify-finance/joltify_lending/app/config"

	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/app"
)

// Test suite used for all keeper tests
type KeeperTestSuite struct {
	suite.Suite

	keeper keeper.Keeper

	app app.TestApp
	ctx context.Context

	genesisTime time.Time
	addrs       []sdk.AccAddress
}

// SetupTest is run automatically before each suite test
func (suite *KeeperTestSuite) SetupTest() {
	appconfig.SetupConfig()

	_, suite.addrs = app.GeneratePrivKeyAddressPairs(5)

	suite.genesisTime = time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC)
}

func (suite *KeeperTestSuite) SetupApp() {
	suite.app = app.NewTestApp(log.NewTestLogger(suite.T()), suite.T().TempDir())

	suite.keeper = suite.app.GetIncentiveKeeper()

	suite.ctx = suite.app.NewContext(true)
}

type accrualtime struct {
	denom string
	time  time.Time
}

var nonEmptyAccrualTimes = []accrualtime{
	{
		denom: "btcb",
		time:  time.Date(1998, 1, 1, 0, 0, 0, 1, time.UTC),
	},
	{
		denom: "ujolt",
		time:  time.Time{},
	},
}

func (suite *KeeperTestSuite) TestIterateHardSupplyRewardAccrualTimes() {
	suite.SetupApp()

	expectedAccrualTimes := nonEmptyAccrualTimes

	for _, at := range expectedAccrualTimes {
		suite.keeper.SetPreviousJoltSupplyRewardAccrualTime(sdk.UnwrapSDKContext(suite.ctx), at.denom, at.time)
	}

	var actualAccrualTimes []accrualtime
	suite.keeper.IterateJoltSupplyRewardAccrualTimes(sdk.UnwrapSDKContext(suite.ctx), func(denom string, accrualTime time.Time) bool {
		actualAccrualTimes = append(actualAccrualTimes, accrualtime{denom: denom, time: accrualTime})
		return false
	})

	suite.Equal(expectedAccrualTimes, actualAccrualTimes)
}

func (suite *KeeperTestSuite) TestIterateHardBorrowrRewardAccrualTimes() {
	suite.SetupApp()

	expectedAccrualTimes := nonEmptyAccrualTimes

	for _, at := range expectedAccrualTimes {
		suite.keeper.SetPreviousJoltBorrowRewardAccrualTime(sdk.UnwrapSDKContext(suite.ctx), at.denom, at.time)
	}

	var actualAccrualTimes []accrualtime
	suite.keeper.IterateJoltBorrowRewardAccrualTimes(sdk.UnwrapSDKContext(suite.ctx), func(denom string, accrualTime time.Time) bool {
		actualAccrualTimes = append(actualAccrualTimes, accrualtime{denom: denom, time: accrualTime})
		return false
	})

	suite.Equal(expectedAccrualTimes, actualAccrualTimes)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
