package keeper_test

import (
	"testing"
	"time"

	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmprototypes "github.com/tendermint/tendermint/proto/tendermint/types"

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
	suite.app = app.NewTestApp()

	suite.keeper = suite.app.GetIncentiveKeeper()

	suite.ctx = suite.app.NewContext(true, tmprototypes.Header{Time: suite.genesisTime})
}

func (suite *KeeperTestSuite) TestGetSetDeleteUSDXMintingClaim() {
	suite.SetupApp()
	c := types2.NewUSDXMintingClaim(suite.addrs[0], c("ujolt", 1000000), types2.RewardIndexes{types2.NewRewardIndex("bnb-a", sdk.ZeroDec())})
	_, found := suite.keeper.GetUSDXMintingClaim(suite.ctx, suite.addrs[0])
	suite.Require().False(found)
	suite.Require().NotPanics(func() {
		suite.keeper.SetUSDXMintingClaim(suite.ctx, c)
	})
	testC, found := suite.keeper.GetUSDXMintingClaim(suite.ctx, suite.addrs[0])
	suite.Require().True(found)
	suite.Require().Equal(c, testC)
	suite.Require().NotPanics(func() {
		suite.keeper.DeleteUSDXMintingClaim(suite.ctx, suite.addrs[0])
	})
	_, found = suite.keeper.GetUSDXMintingClaim(suite.ctx, suite.addrs[0])
	suite.Require().False(found)
}

func (suite *KeeperTestSuite) TestIterateUSDXMintingClaims() {
	suite.SetupApp()
	for i := 0; i < len(suite.addrs); i++ {
		c := types2.NewUSDXMintingClaim(suite.addrs[i], c("ujolt", 100000), types2.RewardIndexes{types2.NewRewardIndex("bnb-a", sdk.ZeroDec())})
		suite.Require().NotPanics(func() {
			suite.keeper.SetUSDXMintingClaim(suite.ctx, c)
		})
	}
	claims := types2.USDXMintingClaims{}
	suite.keeper.IterateUSDXMintingClaims(suite.ctx, func(c types2.USDXMintingClaim) bool {
		claims = append(claims, c)
		return false
	})
	suite.Require().Equal(len(suite.addrs), len(claims))

	claims = suite.keeper.GetAllUSDXMintingClaims(suite.ctx)
	suite.Require().Equal(len(suite.addrs), len(claims))
}

func (suite *KeeperTestSuite) TestGetSetDeleteSwapClaims() {
	suite.SetupApp()
	c := types2.NewSwapClaim(suite.addrs[0], arbitraryCoins(), nonEmptyMultiRewardIndexes)

	_, found := suite.keeper.GetSwapClaim(suite.ctx, suite.addrs[0])
	suite.Require().False(found)

	suite.Require().NotPanics(func() {
		suite.keeper.SetSwapClaim(suite.ctx, c)
	})
	testC, found := suite.keeper.GetSwapClaim(suite.ctx, suite.addrs[0])
	suite.Require().True(found)
	suite.Require().Equal(c, testC)

	suite.Require().NotPanics(func() {
		suite.keeper.DeleteSwapClaim(suite.ctx, suite.addrs[0])
	})
	_, found = suite.keeper.GetSwapClaim(suite.ctx, suite.addrs[0])
	suite.Require().False(found)
}

func (suite *KeeperTestSuite) TestIterateSwapClaims() {
	suite.SetupApp()
	claims := types2.SwapClaims{
		types2.NewSwapClaim(suite.addrs[0], arbitraryCoins(), nonEmptyMultiRewardIndexes),
		types2.NewSwapClaim(suite.addrs[1], nil, nil), // different claim to the first
	}
	for _, claim := range claims {
		suite.keeper.SetSwapClaim(suite.ctx, claim)
	}

	var actualClaims types2.SwapClaims
	suite.keeper.IterateSwapClaims(suite.ctx, func(c types2.SwapClaim) bool {
		actualClaims = append(actualClaims, c)
		return false
	})

	suite.Require().Equal(claims, actualClaims)
}

func (suite *KeeperTestSuite) TestGetSetSwapRewardIndexes() {
	testCases := []struct {
		name      string
		poolName  string
		indexes   types2.RewardIndexes
		wantIndex types2.RewardIndexes
		panics    bool
	}{
		{
			name:     "two factors can be written and read",
			poolName: "btc/usdx",
			indexes: types2.RewardIndexes{
				{
					CollateralType: "jolt",
					RewardFactor:   d("0.02"),
				},
				{
					CollateralType: "ujolt",
					RewardFactor:   d("0.04"),
				},
			},
			wantIndex: types2.RewardIndexes{
				{
					CollateralType: "jolt",
					RewardFactor:   d("0.02"),
				},
				{
					CollateralType: "ujolt",
					RewardFactor:   d("0.04"),
				},
			},
		},
		{
			name:     "indexes with empty pool name panics",
			poolName: "",
			indexes: types2.RewardIndexes{
				{
					CollateralType: "jolt",
					RewardFactor:   d("0.02"),
				},
				{
					CollateralType: "ujolt",
					RewardFactor:   d("0.04"),
				},
			},
			panics: true,
		},
		{
			// this test is to detect any changes in behavior
			name:     "setting empty indexes does not panic",
			poolName: "btc/usdx",
			// Marshalling empty slice results in [] bytes, unmarshalling the []
			// empty bytes results in a nil slice instead of an empty slice
			indexes:   types2.RewardIndexes{},
			wantIndex: nil,
			panics:    false,
		},
		{
			// this test is to detect any changes in behavior
			name:      "setting nil indexes does not panic",
			poolName:  "btc/usdx",
			indexes:   nil,
			wantIndex: nil,
			panics:    false,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupApp()

			_, found := suite.keeper.GetSwapRewardIndexes(suite.ctx, tc.poolName)
			suite.False(found)

			setFunc := func() { suite.keeper.SetSwapRewardIndexes(suite.ctx, tc.poolName, tc.indexes) }
			if tc.panics {
				suite.Panics(setFunc)
				return
			} else {
				suite.NotPanics(setFunc)
			}

			storedIndexes, found := suite.keeper.GetSwapRewardIndexes(suite.ctx, tc.poolName)
			suite.True(found)
			suite.Equal(tc.wantIndex, storedIndexes)
		})
	}
}

func (suite *KeeperTestSuite) TestIterateSwapRewardIndexes() {
	suite.SetupApp()
	multiIndexes := types2.MultiRewardIndexes{
		{
			CollateralType: "bnb/usdx",
			RewardIndexes: types2.RewardIndexes{
				{
					CollateralType: "swap",
					RewardFactor:   d("0.0000002"),
				},
				{
					CollateralType: "ujolt",
					RewardFactor:   d("0.04"),
				},
			},
		},
		{
			CollateralType: "btcb/usdx",
			RewardIndexes: types2.RewardIndexes{
				{
					CollateralType: "jolt",
					RewardFactor:   d("0.02"),
				},
			},
		},
	}
	for _, mi := range multiIndexes {
		suite.keeper.SetSwapRewardIndexes(suite.ctx, mi.CollateralType, mi.RewardIndexes)
	}

	var actualMultiIndexes types2.MultiRewardIndexes
	suite.keeper.IterateSwapRewardIndexes(suite.ctx, func(poolID string, i types2.RewardIndexes) bool {
		actualMultiIndexes = actualMultiIndexes.With(poolID, i)
		return false
	})

	suite.Require().Equal(multiIndexes, actualMultiIndexes)
}

func (suite *KeeperTestSuite) TestGetSetSwapRewardAccrualTimes() {
	testCases := []struct {
		name        string
		poolName    string
		accrualTime time.Time
		panics      bool
	}{
		{
			name:        "normal time can be written and read",
			poolName:    "btc/usdx",
			accrualTime: time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:        "zero time can be written and read",
			poolName:    "btc/usdx",
			accrualTime: time.Time{},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupApp()

			_, found := suite.keeper.GetSwapRewardAccrualTime(suite.ctx, tc.poolName)
			suite.False(found)

			setFunc := func() { suite.keeper.SetSwapRewardAccrualTime(suite.ctx, tc.poolName, tc.accrualTime) }
			if tc.panics {
				suite.Panics(setFunc)
				return
			} else {
				suite.NotPanics(setFunc)
			}

			storedTime, found := suite.keeper.GetSwapRewardAccrualTime(suite.ctx, tc.poolName)
			suite.True(found)
			suite.Equal(tc.accrualTime, storedTime)
		})
	}
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

func (suite *KeeperTestSuite) TestIterateUSDXMintingAccrualTimes() {
	suite.SetupApp()

	expectedAccrualTimes := nonEmptyAccrualTimes

	for _, at := range expectedAccrualTimes {
		suite.keeper.SetPreviousUSDXMintingAccrualTime(suite.ctx, at.denom, at.time)
	}

	var actualAccrualTimes []accrualtime
	suite.keeper.IterateUSDXMintingAccrualTimes(suite.ctx, func(denom string, accrualTime time.Time) bool {
		actualAccrualTimes = append(actualAccrualTimes, accrualtime{denom: denom, time: accrualTime})
		return false
	})

	suite.Equal(expectedAccrualTimes, actualAccrualTimes)
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

func (suite *KeeperTestSuite) TestIterateDelegatorRewardAccrualTimes() {
	suite.SetupApp()

	expectedAccrualTimes := nonEmptyAccrualTimes

	for _, at := range expectedAccrualTimes {
		suite.keeper.SetPreviousDelegatorRewardAccrualTime(suite.ctx, at.denom, at.time)
	}

	var actualAccrualTimes []accrualtime
	suite.keeper.IterateDelegatorRewardAccrualTimes(suite.ctx, func(denom string, accrualTime time.Time) bool {
		actualAccrualTimes = append(actualAccrualTimes, accrualtime{denom: denom, time: accrualTime})
		return false
	})

	suite.Equal(expectedAccrualTimes, actualAccrualTimes)
}

func (suite *KeeperTestSuite) TestIterateSwapRewardAccrualTimes() {
	suite.SetupApp()

	expectedAccrualTimes := nonEmptyAccrualTimes

	for _, at := range expectedAccrualTimes {
		suite.keeper.SetSwapRewardAccrualTime(suite.ctx, at.denom, at.time)
	}

	var actualAccrualTimes []accrualtime
	suite.keeper.IterateSwapRewardAccrualTimes(suite.ctx, func(denom string, accrualTime time.Time) bool {
		actualAccrualTimes = append(actualAccrualTimes, accrualtime{denom: denom, time: accrualTime})
		return false
	})

	suite.Equal(expectedAccrualTimes, actualAccrualTimes)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
