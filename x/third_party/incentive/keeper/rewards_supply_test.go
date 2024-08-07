package keeper_test

import (
	"context"
	"testing"
	"time"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/joltify-finance/joltify_lending/x/third_party/jolt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testutil2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/testutil"
	"github.com/stretchr/testify/suite"

	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/testutil"
	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
	joltkeeper "github.com/joltify-finance/joltify_lending/x/third_party/jolt/keeper"
)

type SupplyIntegrationTests struct {
	testutil2.IntegrationTester

	genesisTime time.Time
	addrs       []sdk.AccAddress
}

func TestSupplyIntegration(t *testing.T) {
	suite.Run(t, new(SupplyIntegrationTests))
}

// SetupTest is run automatically before each suite test
func (suite *SupplyIntegrationTests) SetupTest() {
	_, suite.addrs = app.GeneratePrivKeyAddressPairs(5)

	suite.genesisTime = time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC)
}

func (suite *SupplyIntegrationTests) TestSingleUserAccumulatesRewardsAfterSyncing() {
	userA := suite.addrs[0]

	authBulder := app.NewAuthBankGenesisBuilder().
		WithSimpleModuleAccount(types.ModuleName, cs(c("jjolt", 1e18))). // Fill mintt  with enough coins to pay out any reward
		WithSimpleAccount(userA, cs(c("bnb", 1e12))).                    // give the user some coins
		WithSimpleAccount(userA, cs(c("sbnb", 1e12)))                    // give the user some coins

	incentBuilder := testutil.NewIncentiveGenesisBuilder().
		WithGenesisTime(suite.genesisTime).
		WithMultipliers(types.MultipliersPerDenoms{{
			Denom:       "jjolt",
			Multipliers: types.Multipliers{types.NewMultiplier("large", 12, d("1.0"))}, // keep payout at 1.0 to make maths easier
		}}).
		WithSimpleSupplyRewardPeriod("bnb", cs(c("jjolt", 1e6))). // only borrow rewards
		WithSimpleSupplyRewardPeriod("sbnb", cs(c("jjolt", 1e6))) // only borrow rewards

	var genAcc []authtypes.GenesisAccount
	for _, el := range suite.addrs {
		b := authtypes.NewBaseAccount(el, nil, 0, 0)
		genAcc = append(genAcc, b)
	}

	suite.SetApp()
	suite.StartChain(genAcc, cs(c("bnb", 1e12), c("sbnb", 1e12)),
		suite.genesisTime,
		NewPricefeedGenStateMultiFromTime(suite.App.AppCodec(), suite.genesisTime),
		NewJoltGenStateMulti(suite.genesisTime).BuildMarshalled(suite.App.AppCodec()),
		authBulder.BuildMarshalled(suite.App.AppCodec()),
		incentBuilder.BuildMarshalled(suite.App.AppCodec()),
	)

	err := fundModuleAccount(suite.App.GetBankKeeper(), suite.Ctx, types.ModuleName, cs(c("jjolt", 1e18)))
	suite.Require().NoError(err)

	// Create a deposit
	suite.NoError(suite.DeliverJoltMsgDeposit(userA, cs(c("bnb", 1e11))))
	// Also create a borrow so interest accumulates on the deposit
	suite.NoError(suite.DeliverJoltMsgBorrow(userA, cs(c("bnb", 1e10))))

	// Let time pass to accumulate interest on the deposit
	// Use one long block instead of many to reduce any rounding errors, and speed up tests.
	suite.NextBlockAfter(1e6 * time.Second) // about 12 days

	// User withdraw and redeposits just to sync their deposit.
	suite.NoError(suite.DeliverJoltMsgWithdraw(userA, cs(c("bnb", 1))))
	suite.NoError(suite.DeliverJoltMsgDeposit(userA, cs(c("bnb", 1))))

	// Accumulate more rewards.
	// The user still has the same percentage of all deposits (100%) so their rewards should be the same as in the previous block.
	suite.NextBlockAfter(1e6 * time.Second) // about 12 days

	msg := types.NewMsgClaimJoltReward(
		userA.String(),
		types.Selections{
			types.NewSelection("jjolt", "large"),
		})

	// User claims all their rewards
	suite.NoError(suite.DeliverIncentiveMsg(&msg))

	// The users has always had 100% of deposits, so they should receive all rewards for the previous two blocks.
	// Total rewards for each block is block duration * rewards per second
	accuracy := 1e-10 // using a very high accuracy to flag future small calculation changes
	suite.BalanceInEpsilon(userA, cs(c("bnb", 1e12-1e11+1e10), c("jjolt", 2*1e6*1e6), c("stake", 100000000000000), c("sbnb", 1e12)), accuracy)
}

// Test suite used for all keeper tests
type SupplyRewardsTestSuite struct {
	suite.Suite

	keeper     keeper.Keeper
	joltKeeper joltkeeper.Keeper

	app app.TestApp
	ctx context.Context

	genesisTime time.Time
	addrs       []sdk.AccAddress
}

// SetupTest is run automatically before each suite test
func (suite *SupplyRewardsTestSuite) SetupTest() {
	config := sdk.GetConfig()
	app.SetBech32AddressPrefixes(config)

	_, suite.addrs = app.GeneratePrivKeyAddressPairs(5)

	suite.genesisTime = time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC)
}

func (suite *SupplyRewardsTestSuite) SetupApp() {
	suite.app = app.NewTestApp(log.NewTestLogger(suite.T()), suite.T().TempDir())

	suite.keeper = suite.app.GetIncentiveKeeper()
	suite.joltKeeper = suite.app.GetJoltKeeper()

	suite.ctx = suite.app.Ctx
}

func (suite *SupplyRewardsTestSuite) SetupWithGenState(genAcc []authtypes.GenesisAccount, coins sdk.Coins, authBuilder *app.AuthBankGenesisBuilder, incentBuilder testutil.IncentiveGenesisBuilder, joltBuilder testutil.JoltGenesisBuilder, genTime time.Time) {
	suite.SetupApp()

	suite.app.InitializeFromGenesisStatesWithTime(suite.T(),
		suite.genesisTime, genAcc, coins,
		authBuilder.BuildMarshalled(suite.app.AppCodec()),
		NewPricefeedGenStateMultiFromTime(suite.app.AppCodec(), suite.genesisTime),
		joltBuilder.BuildMarshalled(suite.app.AppCodec()),
		incentBuilder.BuildMarshalled(suite.app.AppCodec()),
	)
	suite.ctx = suite.app.NewContextLegacy(false, tmproto.Header{Height: 1, Time: suite.genesisTime})
}

func (suite *SupplyRewardsTestSuite) TestAccumulateJoltSupplyRewards() {
	type args struct {
		deposit               sdk.Coin
		rewardsPerSecond      sdk.Coins
		timeElapsed           int
		expectedRewardIndexes types.RewardIndexes
	}
	type test struct {
		name string
		args args
	}
	testCases := []test{
		{
			"single reward denom: 7 seconds",
			args{
				deposit:               c("bnb", 1000000000000),
				rewardsPerSecond:      cs(c("jjolt", 122354)),
				timeElapsed:           7,
				expectedRewardIndexes: types.RewardIndexes{types.NewRewardIndex("jjolt", d("0.000000856478000000"))},
			},
		},
		{
			"single reward denom: 1 day",
			args{
				deposit:               c("bnb", 1000000000000),
				rewardsPerSecond:      cs(c("jjolt", 122354)),
				timeElapsed:           86400,
				expectedRewardIndexes: types.RewardIndexes{types.NewRewardIndex("jjolt", d("0.010571385600000000"))},
			},
		},
		{
			"single reward denom: 0 seconds",
			args{
				deposit:               c("bnb", 1000000000000),
				rewardsPerSecond:      cs(c("jjolt", 122354)),
				timeElapsed:           0,
				expectedRewardIndexes: types.RewardIndexes{types.NewRewardIndex("jjolt", d("0.0"))},
			},
		},
		{
			"multiple reward denoms: 7 seconds",
			args{
				deposit:          c("bnb", 1000000000000),
				rewardsPerSecond: cs(c("jjolt", 122354), c("ujolt", 122354)),
				timeElapsed:      7,
				expectedRewardIndexes: types.RewardIndexes{
					types.NewRewardIndex("jjolt", d("0.000000856478000000")),
					types.NewRewardIndex("ujolt", d("0.000000856478000000")),
				},
			},
		},
		{
			"multiple reward denoms: 1 day",
			args{
				deposit:          c("bnb", 1000000000000),
				rewardsPerSecond: cs(c("jjolt", 122354), c("ujolt", 122354)),
				timeElapsed:      86400,
				expectedRewardIndexes: types.RewardIndexes{
					types.NewRewardIndex("jjolt", d("0.010571385600000000")),
					types.NewRewardIndex("ujolt", d("0.010571385600000000")),
				},
			},
		},
		{
			"multiple reward denoms: 0 seconds",
			args{
				deposit:          c("bnb", 1000000000000),
				rewardsPerSecond: cs(c("jjolt", 122354), c("ujolt", 122354)),
				timeElapsed:      0,
				expectedRewardIndexes: types.RewardIndexes{
					types.NewRewardIndex("jjolt", d("0.0")),
					types.NewRewardIndex("ujolt", d("0.0")),
				},
			},
		},
		{
			"multiple reward denoms with different rewards per second: 1 day",
			args{
				deposit:          c("bnb", 1000000000000),
				rewardsPerSecond: cs(c("jjolt", 122354), c("ujolt", 555555)),
				timeElapsed:      86400,
				expectedRewardIndexes: types.RewardIndexes{
					types.NewRewardIndex("jjolt", d("0.010571385600000000")),
					types.NewRewardIndex("ujolt", d("0.047999952000000000")),
				},
			},
		},
	}
	for _, tc := range testCases {

		coins := cs(c("bnb", 1e15), c("ujolt", 1e15), c("btcb", 1e15), c("xrp", 1e15), c("zzz", 1e15))
		suite.Run(tc.name, func() {
			userAddr := suite.addrs[3]
			authBuilder := app.NewAuthBankGenesisBuilder().WithSimpleAccount(
				userAddr,
				coins,
			)
			// suite.SetupWithGenState(authBuilder)
			incentBuilder := testutil.NewIncentiveGenesisBuilder().
				WithGenesisTime(suite.genesisTime)
			if tc.args.rewardsPerSecond != nil {
				incentBuilder = incentBuilder.WithSimpleSupplyRewardPeriod(tc.args.deposit.Denom, tc.args.rewardsPerSecond)
			}

			var genAcc []authtypes.GenesisAccount
			b := authtypes.NewBaseAccount(userAddr, nil, 0, 0)
			genAcc = append(genAcc, b)

			suite.SetupWithGenState(genAcc, coins, authBuilder, incentBuilder, NewJoltGenStateMulti(suite.genesisTime), suite.genesisTime)

			// User deposits to increase total supplied amount
			err := suite.joltKeeper.Deposit(suite.ctx, userAddr, sdk.NewCoins(tc.args.deposit))
			suite.Require().NoError(err)

			// Set up chain context at future time
			runAtTime := sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Duration(int(time.Second) * tc.args.timeElapsed))
			runCtx := sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(runAtTime)

			// Run Jolt begin blocker in order to update the denom's index factor
			jolt.BeginBlocker(runCtx, suite.joltKeeper)

			// Accumulate jolt supply rewards for the deposit denom
			multiRewardPeriod, found := suite.keeper.GetJoltSupplyRewardPeriods(runCtx, tc.args.deposit.Denom)
			suite.Require().True(found)
			suite.keeper.AccumulateJoltSupplyRewards(runCtx, multiRewardPeriod)

			// Check that each expected reward index matches the current stored reward index for the denom
			globalRewardIndexes, found := suite.keeper.GetJoltSupplyRewardIndexes(runCtx, tc.args.deposit.Denom)
			if len(tc.args.rewardsPerSecond) > 0 {
				suite.Require().True(found)
				for _, expectedRewardIndex := range tc.args.expectedRewardIndexes {
					expectedRewardIndex.RewardFactor = expectedRewardIndex.RewardFactor.MulInt64(1000000000000)
					globalRewardIndex, found := globalRewardIndexes.GetRewardIndex(expectedRewardIndex.CollateralType)
					suite.Require().True(found)
					suite.Require().Equal(expectedRewardIndex, globalRewardIndex)
				}
			} else {
				suite.Require().False(found)
			}
		})
	}
}

func (suite *SupplyRewardsTestSuite) TestInitializeJoltSupplyRewards() {
	type args struct {
		moneyMarketRewardDenoms          map[string]sdk.Coins
		deposit                          sdk.Coins
		expectedClaimSupplyRewardIndexes types.MultiRewardIndexes
	}
	type test struct {
		name string
		args args
	}

	standardMoneyMarketRewardDenoms := map[string]sdk.Coins{
		"bnb":  cs(c("jjolt", 1)),
		"btcb": cs(c("jjolt", 1), c("ujolt", 1)),
	}

	testCases := []test{
		{
			"single deposit denom, single reward denom",
			args{
				moneyMarketRewardDenoms: standardMoneyMarketRewardDenoms,
				deposit:                 cs(c("bnb", 1000000000000)),
				expectedClaimSupplyRewardIndexes: types.MultiRewardIndexes{
					types.NewMultiRewardIndex(
						"bnb",
						types.RewardIndexes{
							types.NewRewardIndex("jjolt", d("0.0")),
						},
					),
				},
			},
		},
		{
			"single deposit denom, multiple reward denoms",
			args{
				moneyMarketRewardDenoms: standardMoneyMarketRewardDenoms,
				deposit:                 cs(c("btcb", 1000000000000)),
				expectedClaimSupplyRewardIndexes: types.MultiRewardIndexes{
					types.NewMultiRewardIndex(
						"btcb",
						types.RewardIndexes{
							types.NewRewardIndex("jjolt", d("0.0")),
							types.NewRewardIndex("ujolt", d("0.0")),
						},
					),
				},
			},
		},
		{
			"single deposit denom, no reward denoms",
			args{
				moneyMarketRewardDenoms: standardMoneyMarketRewardDenoms,
				deposit:                 cs(c("xrp", 1000000000000)),
				expectedClaimSupplyRewardIndexes: types.MultiRewardIndexes{
					types.NewMultiRewardIndex(
						"xrp",
						nil,
					),
				},
			},
		},
		{
			"multiple deposit denoms, multiple overlapping reward denoms",
			args{
				moneyMarketRewardDenoms: standardMoneyMarketRewardDenoms,
				deposit:                 cs(c("bnb", 1000000000000), c("btcb", 1000000000000)),
				expectedClaimSupplyRewardIndexes: types.MultiRewardIndexes{
					types.NewMultiRewardIndex(
						"bnb",
						types.RewardIndexes{
							types.NewRewardIndex("jjolt", d("0.0")),
						},
					),
					types.NewMultiRewardIndex(
						"btcb",
						types.RewardIndexes{
							types.NewRewardIndex("jjolt", d("0.0")),
							types.NewRewardIndex("ujolt", d("0.0")),
						},
					),
				},
			},
		},
		{
			"multiple deposit denoms, correct discrete reward denoms",
			args{
				moneyMarketRewardDenoms: standardMoneyMarketRewardDenoms,
				deposit:                 cs(c("bnb", 1000000000000), c("xrp", 1000000000000)),
				expectedClaimSupplyRewardIndexes: types.MultiRewardIndexes{
					types.NewMultiRewardIndex(
						"bnb",
						types.RewardIndexes{
							types.NewRewardIndex("jjolt", d("0.0")),
						},
					),
					types.NewMultiRewardIndex(
						"xrp",
						nil,
					),
				},
			},
		},
	}
	for _, tc := range testCases {
		coins := cs(c("bnb", 1e15), c("ujolt", 1e15), c("btcb", 1e15), c("sbnb", 1e15), c("xrp", 1e15), c("zzz", 1e15))
		suite.Run(tc.name, func() {
			userAddr := suite.addrs[3]
			authBuilder := app.NewAuthBankGenesisBuilder().WithSimpleAccount(
				userAddr,
				coins,
			)

			incentBuilder := testutil.NewIncentiveGenesisBuilder().WithGenesisTime(suite.genesisTime)
			for moneyMarketDenom, rewardsPerSecond := range tc.args.moneyMarketRewardDenoms {
				incentBuilder = incentBuilder.WithSimpleSupplyRewardPeriod(moneyMarketDenom, rewardsPerSecond)
			}

			var genAcc []authtypes.GenesisAccount
			b := authtypes.NewBaseAccount(userAddr, nil, 0, 0)
			genAcc = append(genAcc, b)

			suite.SetupWithGenState(genAcc, coins, authBuilder, incentBuilder, NewJoltGenStateMulti(suite.genesisTime), suite.genesisTime)

			// User deposits
			err := suite.joltKeeper.Deposit(suite.ctx, userAddr, tc.args.deposit)
			suite.Require().NoError(err)

			claim, foundClaim := suite.keeper.GetJoltLiquidityProviderClaim(sdk.UnwrapSDKContext(suite.ctx), userAddr)
			suite.Require().True(foundClaim)
			suite.Require().Equal(tc.args.expectedClaimSupplyRewardIndexes, claim.SupplyRewardIndexes)
		})
	}
}

func (suite *SupplyRewardsTestSuite) TestSynchronizeJoltSupplyReward() {
	type args struct {
		incentiveSupplyRewardDenom   string
		deposit                      sdk.Coin
		rewardsPerSecond             sdk.Coins
		blockTimes                   []int
		expectedRewardIndexes        types.RewardIndexes
		expectedRewards              sdk.Coins
		updateRewardsViaCommmittee   bool
		updatedBaseDenom             string
		updatedRewardsPerSecond      sdk.Coins
		updatedExpectedRewardIndexes types.RewardIndexes
		updatedExpectedRewards       sdk.Coins
		updatedTimeDuration          int
	}
	type test struct {
		name string
		args args
	}

	testCases := []test{
		{
			"single reward denom: 10 blocks",
			args{
				incentiveSupplyRewardDenom: "bnb",
				deposit:                    c("bnb", 10000000000),
				rewardsPerSecond:           cs(c("jjolt", 122354)),
				blockTimes:                 []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
				expectedRewardIndexes:      types.RewardIndexes{types.NewRewardIndex("jjolt", d("0.001223540000000000"))},
				expectedRewards:            cs(c("jjolt", 12235400)),
				updateRewardsViaCommmittee: false,
			},
		},
		{
			"single reward denom: 10 blocks - long block time",
			args{
				incentiveSupplyRewardDenom: "bnb",
				deposit:                    c("bnb", 10000000000),
				rewardsPerSecond:           cs(c("jjolt", 122354)),
				blockTimes:                 []int{86400, 86400, 86400, 86400, 86400, 86400, 86400, 86400, 86400, 86400},
				expectedRewardIndexes:      types.RewardIndexes{types.NewRewardIndex("jjolt", d("10.571385600000000000"))},
				expectedRewards:            cs(c("jjolt", 105713856000)),
				updateRewardsViaCommmittee: false,
			},
		},
		{
			"single reward denom: user reward index updated when reward is zero",
			args{
				incentiveSupplyRewardDenom: "sbnb",
				deposit:                    c("sbnb", 1),
				rewardsPerSecond:           cs(c("jjolt", 122354)),
				blockTimes:                 []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
				expectedRewardIndexes:      types.RewardIndexes{types.NewRewardIndex("jjolt", d("0.122353998776460012"))},
				expectedRewards:            cs(),
				updateRewardsViaCommmittee: false,
			},
		},
		{
			"multiple reward denoms: 10 blocks",
			args{
				incentiveSupplyRewardDenom: "bnb",
				deposit:                    c("bnb", 10000000000),
				rewardsPerSecond:           cs(c("jjolt", 122354), c("ujolt", 122354)),
				blockTimes:                 []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
				expectedRewardIndexes: types.RewardIndexes{
					types.NewRewardIndex("jjolt", d("0.001223540000000000")),
					types.NewRewardIndex("ujolt", d("0.001223540000000000")),
				},
				expectedRewards:            cs(c("jjolt", 12235400), c("ujolt", 12235400)),
				updateRewardsViaCommmittee: false,
			},
		},
		{
			"multiple reward denoms: 10 blocks - long block time",
			args{
				incentiveSupplyRewardDenom: "bnb",
				deposit:                    c("bnb", 10000000000),
				rewardsPerSecond:           cs(c("jjolt", 122354), c("ujolt", 122354)),
				blockTimes:                 []int{86400, 86400, 86400, 86400, 86400, 86400, 86400, 86400, 86400, 86400},
				expectedRewardIndexes: types.RewardIndexes{
					types.NewRewardIndex("jjolt", d("10.571385600000000000")),
					types.NewRewardIndex("ujolt", d("10.571385600000000000")),
				},
				expectedRewards:            cs(c("jjolt", 105713856000), c("ujolt", 105713856000)),
				updateRewardsViaCommmittee: false,
			},
		},
		{
			"multiple reward denoms with different rewards per second: 10 blocks",
			args{
				incentiveSupplyRewardDenom: "bnb",
				deposit:                    c("bnb", 10000000000),
				rewardsPerSecond:           cs(c("jjolt", 122354), c("ujolt", 555555)),
				blockTimes:                 []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
				expectedRewardIndexes: types.RewardIndexes{
					types.NewRewardIndex("jjolt", d("0.001223540000000000")),
					types.NewRewardIndex("ujolt", d("0.005555550000000000")),
				},
				expectedRewards:            cs(c("jjolt", 12235400), c("ujolt", 55555500)),
				updateRewardsViaCommmittee: false,
			},
		},
		{
			"denom is in incentive's jolt supply reward params and has rewards; add new reward type",
			args{
				incentiveSupplyRewardDenom: "bnb",
				deposit:                    c("bnb", 10000000000),
				rewardsPerSecond:           cs(c("jjolt", 122354)),
				blockTimes:                 []int{86400},
				expectedRewardIndexes: types.RewardIndexes{
					types.NewRewardIndex("jjolt", d("1.057138560000000000")),
				},
				expectedRewards:            cs(c("jjolt", 10571385600)),
				updateRewardsViaCommmittee: true,
				updatedBaseDenom:           "bnb",
				updatedRewardsPerSecond:    cs(c("jjolt", 122354), c("ujolt", 100000)),
				updatedExpectedRewards:     cs(c("jjolt", 21142771200), c("ujolt", 8640000000)),
				updatedExpectedRewardIndexes: types.RewardIndexes{
					types.NewRewardIndex("jjolt", d("2.114277120000000000")),
					types.NewRewardIndex("ujolt", d("0.864000000000000000")),
				},
				updatedTimeDuration: 86400,
			},
		},
		{
			"denom is in jolt's money market params but not in incentive's jolt supply reward params; add reward",
			args{
				incentiveSupplyRewardDenom: "bnb",
				deposit:                    c("zzz", 10000000000),
				rewardsPerSecond:           nil,
				blockTimes:                 []int{100},
				expectedRewardIndexes:      types.RewardIndexes{},
				expectedRewards:            sdk.Coins{},
				updateRewardsViaCommmittee: true,
				updatedBaseDenom:           "zzz",
				updatedRewardsPerSecond:    cs(c("jjolt", 100000)),
				updatedExpectedRewards:     cs(c("jjolt", 8640000000)),
				updatedExpectedRewardIndexes: types.RewardIndexes{
					types.NewRewardIndex("jjolt", d("0.864")),
				},
				updatedTimeDuration: 86400,
			},
		},
		{
			"denom is in jolt's money market params but not in incentive's jolt supply reward params; add multiple reward types",
			args{
				incentiveSupplyRewardDenom: "bnb",
				deposit:                    c("zzz", 10000000000),
				rewardsPerSecond:           nil,
				blockTimes:                 []int{100},
				expectedRewardIndexes:      types.RewardIndexes{},
				expectedRewards:            sdk.Coins{},
				updateRewardsViaCommmittee: true,
				updatedBaseDenom:           "zzz",
				updatedRewardsPerSecond:    cs(c("jjolt", 100000), c("ujolt", 100500), c("swap", 500)),
				updatedExpectedRewards:     cs(c("jjolt", 8640000000), c("ujolt", 8683200000), c("swap", 43200000)),
				updatedExpectedRewardIndexes: types.RewardIndexes{
					types.NewRewardIndex("jjolt", d("0.864")),
					types.NewRewardIndex("ujolt", d("0.86832")),
					types.NewRewardIndex("swap", d("0.00432")),
				},
				updatedTimeDuration: 86400,
			},
		},
	}
	for _, tc := range testCases {
		coins := cs(c("bnb", 1e15), c("ujolt", 1e15), c("btcb", 1e15), c("sbnb", 1e15), c("xrp", 1e15), c("zzz", 1e15))
		suite.Run(tc.name, func() {
			userAddr := suite.addrs[3]
			authBuilder := app.NewAuthBankGenesisBuilder().
				WithSimpleAccount(suite.addrs[2], cs(c("ujolt", 1e9), c("sbnb", 1e9))).
				WithSimpleAccount(userAddr, coins)

			incentBuilder := testutil.NewIncentiveGenesisBuilder().
				WithGenesisTime(suite.genesisTime)
			if tc.args.rewardsPerSecond != nil {
				incentBuilder = incentBuilder.WithSimpleSupplyRewardPeriod(tc.args.incentiveSupplyRewardDenom, tc.args.rewardsPerSecond)
			}

			var genAcc []authtypes.GenesisAccount
			b := authtypes.NewBaseAccount(userAddr, nil, 0, 0)
			genAcc = append(genAcc, b)

			suite.SetupWithGenState(genAcc, coins, authBuilder, incentBuilder, NewJoltGenStateMulti(suite.genesisTime), suite.genesisTime)

			err := fundAccount(suite.app.GetBankKeeper(), suite.ctx, suite.addrs[2], cs(c("ujolt", 1e9), c("sbnb", 1e9)))
			suite.Require().NoError(err)

			// Deposit a fixed amount from another user to dilute primary user's rewards per second.
			suite.Require().NoError(
				suite.joltKeeper.Deposit(suite.ctx, suite.addrs[2], cs(c("sbnb", 100_000_000))),
			)

			// User deposits and borrows to increase total borrowed amount
			err = suite.joltKeeper.Deposit(suite.ctx, userAddr, sdk.NewCoins(tc.args.deposit))
			suite.Require().NoError(err)

			// Check that Jolt hooks initialized a JoltLiquidityProviderClaim with 0 reward indexes
			claim, found := suite.keeper.GetJoltLiquidityProviderClaim(sdk.UnwrapSDKContext(suite.ctx), userAddr)
			suite.Require().True(found)
			multiRewardIndex, _ := claim.SupplyRewardIndexes.GetRewardIndex(tc.args.deposit.Denom)
			for _, expectedRewardIndex := range tc.args.expectedRewardIndexes {
				expectedRewardIndex.RewardFactor = expectedRewardIndex.RewardFactor.MulInt64(1000000000000)
				currRewardIndex, found := multiRewardIndex.RewardIndexes.GetRewardIndex(expectedRewardIndex.CollateralType)
				suite.Require().True(found)
				suite.Require().Equal(sdkmath.LegacyZeroDec(), currRewardIndex.RewardFactor)
			}

			// Run accumulator at several intervals
			var timeElapsed int
			previousBlockTime := sdk.UnwrapSDKContext(suite.ctx).BlockTime()
			for _, t := range tc.args.blockTimes {
				timeElapsed += t
				updatedBlockTime := previousBlockTime.Add(time.Duration(int(time.Second) * t))
				previousBlockTime = updatedBlockTime
				blockCtx := sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(updatedBlockTime)

				// Run Jolt begin blocker for each block ctx to update denom's interest factor
				jolt.BeginBlocker(blockCtx, suite.joltKeeper)

				// Accumulate jolt supply-side rewards
				multiRewardPeriod, found := suite.keeper.GetJoltSupplyRewardPeriods(blockCtx, tc.args.deposit.Denom)
				if found {
					suite.keeper.AccumulateJoltSupplyRewards(blockCtx, multiRewardPeriod)
				}
			}
			updatedBlockTime := sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Duration(int(time.Second) * timeElapsed))
			suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(updatedBlockTime)

			// After we've accumulated, run synchronize
			deposit, found := suite.joltKeeper.GetDeposit(suite.ctx, userAddr)
			suite.Require().True(found)
			suite.Require().NotPanics(func() {
				suite.keeper.SynchronizeJoltSupplyReward(sdk.UnwrapSDKContext(suite.ctx), deposit)
			})

			// Check that the global reward index's reward factor and user's claim have been updated as expected
			claim, found = suite.keeper.GetJoltLiquidityProviderClaim(sdk.UnwrapSDKContext(suite.ctx), userAddr)
			suite.Require().True(found)
			globalRewardIndexes, foundGlobalRewardIndexes := suite.keeper.GetJoltSupplyRewardIndexes(sdk.UnwrapSDKContext(suite.ctx), tc.args.deposit.Denom)
			if len(tc.args.rewardsPerSecond) > 0 {
				suite.Require().True(foundGlobalRewardIndexes)
				for _, expectedRewardIndex := range tc.args.expectedRewardIndexes {
					globalRewardIndex, found := globalRewardIndexes.GetRewardIndex(expectedRewardIndex.CollateralType)
					// in joltify we need to quo by 10^12
					globalRewardIndex.RewardFactor = globalRewardIndex.RewardFactor.QuoInt64(1000000000000)
					suite.Require().True(found)
					suite.Require().Equal(expectedRewardIndex, globalRewardIndex)

					// Check that the user's claim's reward index matches the corresponding global reward index
					multiRewardIndex, found := claim.SupplyRewardIndexes.GetRewardIndex(tc.args.deposit.Denom)
					suite.Require().True(found)
					rewardIndex, found := multiRewardIndex.RewardIndexes.GetRewardIndex(expectedRewardIndex.CollateralType)

					rewardIndex.RewardFactor = rewardIndex.RewardFactor.QuoInt64(1000000000000)
					suite.Require().True(found)
					suite.Require().Equal(expectedRewardIndex, rewardIndex)

					// Check that the user's claim holds the expected amount of reward coins
					suite.Require().Equal(
						tc.args.expectedRewards.AmountOf(expectedRewardIndex.CollateralType),
						claim.Reward.AmountOf(expectedRewardIndex.CollateralType),
					)
				}
			}

			// Only test cases with reward param updates continue past this point
			if !tc.args.updateRewardsViaCommmittee {
				return
			}

			// If are no initial rewards per second, add new rewards through a committee param change
			// 1. Construct incentive's new JoltSupplyRewardPeriods param
			currIncentiveJoltSupplyRewardPeriods := suite.keeper.GetParams(sdk.UnwrapSDKContext(suite.ctx)).JoltSupplyRewardPeriods
			multiRewardPeriod, found := currIncentiveJoltSupplyRewardPeriods.GetMultiRewardPeriod(tc.args.deposit.Denom)
			if found {
				// Deposit denom's reward period exists, but it doesn't have any rewards per second
				index, found := currIncentiveJoltSupplyRewardPeriods.GetMultiRewardPeriodIndex(tc.args.deposit.Denom)
				suite.Require().True(found)
				multiRewardPeriod.RewardsPerSecond = tc.args.updatedRewardsPerSecond
				currIncentiveJoltSupplyRewardPeriods[index] = multiRewardPeriod
			} else {
				// Deposit denom's reward period does not exist
				_, found := currIncentiveJoltSupplyRewardPeriods.GetMultiRewardPeriodIndex(tc.args.deposit.Denom)
				suite.Require().False(found)
				newMultiRewardPeriod := types.NewMultiRewardPeriod(true, tc.args.deposit.Denom, suite.genesisTime, suite.genesisTime.Add(time.Hour*24*365*4), tc.args.updatedRewardsPerSecond)
				currIncentiveJoltSupplyRewardPeriods = append(currIncentiveJoltSupplyRewardPeriods, newMultiRewardPeriod)
			}

			// 2. Construct the parameter change proposal to update JoltSupplyRewardPeriods param
			//pubProposal := proposaltypes.NewParameterChangeProposal(
			//	"Update jolt supply rewards", "Adds a new reward coin to the incentive module's jolt supply rewards.",
			//	[]proposaltypes.ParamChange{
			//		{
			//			Subspace: types.ModuleName,                         // target incentive module
			//			Key:      string(types.KeyJoltSupplyRewardPeriods), // target jolt supply rewards key
			//			Value:    string(suite.app.LegacyAmino().MustMarshalJSON(&currIncentiveJoltSupplyRewardPeriods)),
			//		},
			//	},
			//)

			// types.KeyJoltBorrowRewardPeriods

			// Value:    string(suite.app.LegacyAmino().MustMarshalJSON(&currIncentiveJoltSupplyRewardPeriods)),

			params := suite.app.GetIncentiveKeeper().GetParams(sdk.UnwrapSDKContext(suite.ctx))
			params.JoltSupplyRewardPeriods = currIncentiveJoltSupplyRewardPeriods
			suite.app.GetIncentiveKeeper().SetParams(sdk.UnwrapSDKContext(suite.ctx), params)

			// We need to accumulate jolt supply-side rewards again
			multiRewardPeriod, found = suite.keeper.GetJoltSupplyRewardPeriods(suite.ctx, tc.args.deposit.Denom)
			suite.Require().True(found)

			// But new deposit denoms don't have their PreviousJoltSupplyRewardAccrualTime set yet,
			// so we need to call the accumulation method once to set the initial reward accrual time
			if tc.args.deposit.Denom != tc.args.incentiveSupplyRewardDenom {
				suite.keeper.AccumulateJoltSupplyRewards(sdk.UnwrapSDKContext(suite.ctx), multiRewardPeriod)
			}

			// Now we can jump forward in time and accumulate rewards
			updatedBlockTime = previousBlockTime.Add(time.Duration(int(time.Second) * tc.args.updatedTimeDuration))
			suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(updatedBlockTime)
			suite.keeper.AccumulateJoltSupplyRewards(sdk.UnwrapSDKContext(suite.ctx), multiRewardPeriod)

			// After we've accumulated, run synchronize
			deposit, found = suite.joltKeeper.GetDeposit(suite.ctx, userAddr)
			suite.Require().True(found)
			suite.Require().NotPanics(func() {
				suite.keeper.SynchronizeJoltSupplyReward(sdk.UnwrapSDKContext(suite.ctx), deposit)
			})

			// Check that the global reward index's reward factor and user's claim have been updated as expected
			globalRewardIndexes, found = suite.keeper.GetJoltSupplyRewardIndexes(sdk.UnwrapSDKContext(suite.ctx), tc.args.deposit.Denom)
			suite.Require().True(found)
			claim, found = suite.keeper.GetJoltLiquidityProviderClaim(sdk.UnwrapSDKContext(suite.ctx), userAddr)
			suite.Require().True(found)
			for _, expectedRewardIndex := range tc.args.updatedExpectedRewardIndexes {
				expectedRewardIndex.RewardFactor = expectedRewardIndex.RewardFactor.MulInt64(1000000000000)
				// Check that global reward index has been updated as expected
				globalRewardIndex, found := globalRewardIndexes.GetRewardIndex(expectedRewardIndex.CollateralType)
				suite.Require().True(found)
				suite.Require().Equal(expectedRewardIndex, globalRewardIndex)

				// Check that the user's claim's reward index matches the corresponding global reward index
				multiRewardIndex, found := claim.SupplyRewardIndexes.GetRewardIndex(tc.args.deposit.Denom)
				suite.Require().True(found)
				rewardIndex, found := multiRewardIndex.RewardIndexes.GetRewardIndex(expectedRewardIndex.CollateralType)
				suite.Require().True(found)
				suite.Require().Equal(expectedRewardIndex, rewardIndex)

				// Check that the user's claim holds the expected amount of reward coins
				suite.Require().Equal(
					tc.args.updatedExpectedRewards.AmountOf(expectedRewardIndex.CollateralType),
					claim.Reward.AmountOf(expectedRewardIndex.CollateralType),
				)
			}
		})
	}
}

func (suite *SupplyRewardsTestSuite) TestUpdateJoltSupplyIndexDenoms() {
	type depositModification struct {
		coins    sdk.Coins
		withdraw bool
	}

	type args struct {
		firstDeposit              sdk.Coins
		modification              depositModification
		rewardsPerSecond          sdk.Coins
		expectedSupplyIndexDenoms []string
	}
	type test struct {
		name string
		args args
	}

	testCases := []test{
		{
			"single reward denom: update adds one supply reward index",
			args{
				firstDeposit:              cs(c("bnb", 10000000000)),
				modification:              depositModification{coins: cs(c("ujolt", 10000000000))},
				rewardsPerSecond:          cs(c("jjolt", 122354)),
				expectedSupplyIndexDenoms: []string{"bnb", "ujolt"},
			},
		},
		{
			"single reward denom: update adds multiple supply reward indexes",
			args{
				firstDeposit:              cs(c("bnb", 10000000000)),
				modification:              depositModification{coins: cs(c("ujolt", 10000000000), c("btcb", 10000000000), c("xrp", 10000000000))},
				rewardsPerSecond:          cs(c("jjolt", 122354)),
				expectedSupplyIndexDenoms: []string{"bnb", "ujolt", "btcb", "xrp"},
			},
		},
		{
			"single reward denom: update doesn't add duplicate supply reward index for same denom",
			args{
				firstDeposit:              cs(c("bnb", 10000000000)),
				modification:              depositModification{coins: cs(c("bnb", 5000000000))},
				rewardsPerSecond:          cs(c("jjolt", 122354)),
				expectedSupplyIndexDenoms: []string{"bnb"},
			},
		},
		{
			"multiple reward denoms: update adds one supply reward index",
			args{
				firstDeposit:              cs(c("bnb", 10000000000)),
				modification:              depositModification{coins: cs(c("ujolt", 10000000000))},
				rewardsPerSecond:          cs(c("jjolt", 122354), c("ujolt", 122354)),
				expectedSupplyIndexDenoms: []string{"bnb", "ujolt"},
			},
		},
		{
			"multiple reward denoms: update adds multiple supply reward indexes",
			args{
				firstDeposit:              cs(c("bnb", 10000000000)),
				modification:              depositModification{coins: cs(c("ujolt", 10000000000), c("btcb", 10000000000), c("xrp", 10000000000))},
				rewardsPerSecond:          cs(c("jjolt", 122354), c("ujolt", 122354)),
				expectedSupplyIndexDenoms: []string{"bnb", "ujolt", "btcb", "xrp"},
			},
		},
		{
			"multiple reward denoms: update doesn't add duplicate supply reward index for same denom",
			args{
				firstDeposit:              cs(c("bnb", 10000000000)),
				modification:              depositModification{coins: cs(c("bnb", 5000000000))},
				rewardsPerSecond:          cs(c("jjolt", 122354), c("ujolt", 122354)),
				expectedSupplyIndexDenoms: []string{"bnb"},
			},
		},
		{
			"single reward denom: fully withdrawing a denom deletes the denom's supply reward index",
			args{
				firstDeposit:              cs(c("bnb", 1000000000)),
				modification:              depositModification{coins: cs(c("bnb", 1100000000)), withdraw: true},
				rewardsPerSecond:          cs(c("jjolt", 122354)),
				expectedSupplyIndexDenoms: []string{},
			},
		},
		{
			"single reward denom: fully withdrawing a denom deletes only the denom's supply reward index",
			args{
				firstDeposit:              cs(c("bnb", 1000000000), c("ujolt", 100000000)),
				modification:              depositModification{coins: cs(c("bnb", 1100000000)), withdraw: true},
				rewardsPerSecond:          cs(c("jjolt", 122354)),
				expectedSupplyIndexDenoms: []string{"ujolt"},
			},
		},
		{
			"multiple reward denoms: fully repaying a denom deletes the denom's supply reward index",
			args{
				firstDeposit:              cs(c("bnb", 1000000000)),
				modification:              depositModification{coins: cs(c("bnb", 1100000000)), withdraw: true},
				rewardsPerSecond:          cs(c("jjolt", 122354), c("ujolt", 122354)),
				expectedSupplyIndexDenoms: []string{},
			},
		},
	}
	for _, tc := range testCases {
		coins := cs(c("bnb", 1e15), c("ujolt", 1e15), c("btcb", 1e15), c("xrp", 1e15), c("zzz", 1e15))
		suite.Run(tc.name, func() {
			userAddr := suite.addrs[3]
			authBuilder := app.NewAuthBankGenesisBuilder().WithSimpleAccount(
				userAddr,
				coins,
			)
			incentBuilder := testutil.NewIncentiveGenesisBuilder().
				WithGenesisTime(suite.genesisTime).
				WithSimpleSupplyRewardPeriod("bnb", tc.args.rewardsPerSecond).
				WithSimpleSupplyRewardPeriod("ujolt", tc.args.rewardsPerSecond).
				WithSimpleSupplyRewardPeriod("btcb", tc.args.rewardsPerSecond).
				WithSimpleSupplyRewardPeriod("xrp", tc.args.rewardsPerSecond)

			var genAcc []authtypes.GenesisAccount
			b := authtypes.NewBaseAccount(userAddr, nil, 0, 0)
			genAcc = append(genAcc, b)
			suite.SetupWithGenState(genAcc, coins, authBuilder, incentBuilder, NewJoltGenStateMulti(suite.genesisTime), suite.genesisTime)

			// User deposits (first time)
			err := suite.joltKeeper.Deposit(suite.ctx, userAddr, tc.args.firstDeposit)
			suite.Require().NoError(err)

			// Confirm that a claim was created and populated with the correct supply indexes
			claimAfterFirstDeposit, found := suite.keeper.GetJoltLiquidityProviderClaim(sdk.UnwrapSDKContext(suite.ctx), userAddr)
			suite.Require().True(found)
			for _, coin := range tc.args.firstDeposit {
				_, hasIndex := claimAfterFirstDeposit.HasSupplyRewardIndex(coin.Denom)
				suite.Require().True(hasIndex)
			}
			suite.Require().True(len(claimAfterFirstDeposit.SupplyRewardIndexes) == len(tc.args.firstDeposit))

			// User modifies their Deposit by withdrawing or depositing more
			if tc.args.modification.withdraw {
				err = suite.joltKeeper.Withdraw(suite.ctx, userAddr, tc.args.modification.coins)
			} else {
				err = suite.joltKeeper.Deposit(suite.ctx, userAddr, tc.args.modification.coins)
			}
			suite.Require().NoError(err)

			// Confirm that the claim contains all expected supply indexes
			claimAfterModification, found := suite.keeper.GetJoltLiquidityProviderClaim(sdk.UnwrapSDKContext(suite.ctx), userAddr)
			suite.Require().True(found)
			for _, denom := range tc.args.expectedSupplyIndexDenoms {
				_, hasIndex := claimAfterModification.HasSupplyRewardIndex(denom)
				suite.Require().True(hasIndex)
			}
			suite.Require().True(len(claimAfterModification.SupplyRewardIndexes) == len(tc.args.expectedSupplyIndexDenoms))
		})
	}
}

func (suite *SupplyRewardsTestSuite) TestSimulateJoltSupplyRewardSynchronization() {
	type args struct {
		deposit               sdk.Coin
		rewardsPerSecond      sdk.Coins
		blockTimes            []int
		expectedRewardIndexes types.RewardIndexes
		expectedRewards       sdk.Coins
	}
	type test struct {
		name string
		args args
	}

	testCases := []test{
		{
			"10 blocks",
			args{
				deposit:               c("bnb", 10000000000),
				rewardsPerSecond:      cs(c("jjolt", 122354)),
				blockTimes:            []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
				expectedRewardIndexes: types.RewardIndexes{types.NewRewardIndex("jjolt", d("0.001223540000000000"))},
				expectedRewards:       cs(c("jjolt", 12235400)),
			},
		},
		{
			"10 blocks - long block time",
			args{
				deposit:               c("bnb", 10000000000),
				rewardsPerSecond:      cs(c("jjolt", 122354)),
				blockTimes:            []int{86400, 86400, 86400, 86400, 86400, 86400, 86400, 86400, 86400, 86400},
				expectedRewardIndexes: types.RewardIndexes{types.NewRewardIndex("jjolt", d("10.571385600000000000"))},
				expectedRewards:       cs(c("jjolt", 105713856000)),
			},
		},
	}
	for _, tc := range testCases {
		coins := cs(c("bnb", 1e15), c("ujolt", 1e15), c("btcb", 1e15), c("xrp", 1e15), c("zzz", 1e15))
		suite.Run(tc.name, func() {
			userAddr := suite.addrs[3]
			authBuilder := app.NewAuthBankGenesisBuilder().WithSimpleAccount(
				userAddr,
				coins,
			)
			incentBuilder := testutil.NewIncentiveGenesisBuilder().
				WithGenesisTime(suite.genesisTime).
				WithSimpleSupplyRewardPeriod(tc.args.deposit.Denom, tc.args.rewardsPerSecond)

			var genAcc []authtypes.GenesisAccount
			b := authtypes.NewBaseAccount(userAddr, nil, 0, 0)
			genAcc = append(genAcc, b)
			suite.SetupWithGenState(genAcc, coins, authBuilder, incentBuilder, NewJoltGenStateMulti(suite.genesisTime), suite.genesisTime)

			// User deposits and borrows to increase total borrowed amount
			err := suite.joltKeeper.Deposit(suite.ctx, userAddr, sdk.NewCoins(tc.args.deposit))
			suite.Require().NoError(err)

			// Run accumulator at several intervals
			var timeElapsed int
			previousBlockTime := sdk.UnwrapSDKContext(suite.ctx).BlockTime()
			for _, t := range tc.args.blockTimes {
				timeElapsed += t
				updatedBlockTime := previousBlockTime.Add(time.Duration(int(time.Second) * t))
				previousBlockTime = updatedBlockTime
				blockCtx := sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(updatedBlockTime)

				// Run Jolt begin blocker for each block ctx to update denom's interest factor
				jolt.BeginBlocker(blockCtx, suite.joltKeeper)

				// Accumulate jolt supply-side rewards
				multiRewardPeriod, found := suite.keeper.GetJoltSupplyRewardPeriods(blockCtx, tc.args.deposit.Denom)
				suite.Require().True(found)
				suite.keeper.AccumulateJoltSupplyRewards(blockCtx, multiRewardPeriod)
			}
			updatedBlockTime := sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Duration(int(time.Second) * timeElapsed))
			suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(updatedBlockTime)

			// Confirm that the user's claim hasn't been synced
			claimPre, foundPre := suite.keeper.GetJoltLiquidityProviderClaim(sdk.UnwrapSDKContext(suite.ctx), userAddr)
			suite.Require().True(foundPre)
			multiRewardIndexPre, _ := claimPre.SupplyRewardIndexes.GetRewardIndex(tc.args.deposit.Denom)
			for _, expectedRewardIndex := range tc.args.expectedRewardIndexes {
				currRewardIndex, found := multiRewardIndexPre.RewardIndexes.GetRewardIndex(expectedRewardIndex.CollateralType)
				suite.Require().True(found)
				suite.Require().Equal(sdkmath.LegacyZeroDec(), currRewardIndex.RewardFactor)
			}

			// Check that the synced claim held in memory has properly simulated syncing
			syncedClaim := suite.keeper.SimulateJoltSynchronization(sdk.UnwrapSDKContext(suite.ctx), claimPre)
			for _, expectedRewardIndex := range tc.args.expectedRewardIndexes {
				expectedRewardIndex.RewardFactor = expectedRewardIndex.RewardFactor.MulInt64(1000000000000)
				// Check that the user's claim's reward index matches the expected reward index
				multiRewardIndex, found := syncedClaim.SupplyRewardIndexes.GetRewardIndex(tc.args.deposit.Denom)
				suite.Require().True(found)
				rewardIndex, found := multiRewardIndex.RewardIndexes.GetRewardIndex(expectedRewardIndex.CollateralType)
				suite.Require().True(found)
				suite.Require().Equal(expectedRewardIndex, rewardIndex)

				// Check that the user's claim holds the expected amount of reward coins
				suite.Require().Equal(
					tc.args.expectedRewards.AmountOf(expectedRewardIndex.CollateralType),
					syncedClaim.Reward.AmountOf(expectedRewardIndex.CollateralType),
				)
			}
		})
	}
}

func TestSupplyRewardsTestSuite(t *testing.T) {
	suite.Run(t, new(SupplyRewardsTestSuite))
}
