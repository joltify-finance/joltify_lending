package keeper_test

import (
	"testing"
	"time"

	tmlog "github.com/cometbft/cometbft/libs/log"

	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
	"github.com/stretchr/testify/suite"

	tmprototypes "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/app"
)

// Test suite used for all keeper tests
type KeeperTestSuite struct {
	suite.Suite

	keeper keeper.Keeper

	app app.TestApp
	ctx sdk.Context

	genesisTime time.Time
	addrs       []sdk.AccAddress
}

// SetupTest is run automatically before each suite test
func (suite *KeeperTestSuite) SetupTest() {
	config := sdk.GetConfig()
	app.SetBech32AddressPrefixes(config)

	_, suite.addrs = app.GeneratePrivKeyAddressPairs(5)

	suite.genesisTime = time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC)
}

func (suite *KeeperTestSuite) SetupApp() {
	suite.app = app.NewTestApp(tmlog.TestingLogger(), suite.T().TempDir())

	suite.keeper = suite.app.GetIncentiveKeeper()

	suite.ctx = suite.app.NewContext(true, tmprototypes.Header{Time: suite.genesisTime})
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
		suite.keeper.SetPreviousJoltSupplyRewardAccrualTime(suite.ctx, at.denom, at.time)
	}

	var actualAccrualTimes []accrualtime
	suite.keeper.IterateJoltSupplyRewardAccrualTimes(suite.ctx, func(denom string, accrualTime time.Time) bool {
		actualAccrualTimes = append(actualAccrualTimes, accrualtime{denom: denom, time: accrualTime})
		return false
	})

	suite.Equal(expectedAccrualTimes, actualAccrualTimes)
}

func (suite *KeeperTestSuite) TestIterateHardBorrowrRewardAccrualTimes() {
	suite.SetupApp()

	expectedAccrualTimes := nonEmptyAccrualTimes

	for _, at := range expectedAccrualTimes {
		suite.keeper.SetPreviousJoltBorrowRewardAccrualTime(suite.ctx, at.denom, at.time)
	}

	var actualAccrualTimes []accrualtime
	suite.keeper.IterateJoltBorrowRewardAccrualTimes(suite.ctx, func(denom string, accrualTime time.Time) bool {
		actualAccrualTimes = append(actualAccrualTimes, accrualtime{denom: denom, time: accrualTime})
		return false
	})

	suite.Equal(expectedAccrualTimes, actualAccrualTimes)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
