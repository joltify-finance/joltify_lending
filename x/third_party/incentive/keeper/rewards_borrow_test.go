package keeper_test

import (
	"context"
	"testing"
	"time"

	tmlog "github.com/cometbft/cometbft/libs/log"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
	testutil2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/testutil"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/jolt"
	joltkeeper "github.com/joltify-finance/joltify_lending/x/third_party/jolt/keeper"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/joltify-finance/joltify_lending/app"
)

type BorrowIntegrationTests struct {
	testutil2.IntegrationTester

	genesisTime time.Time
	addrs       []sdk.AccAddress
}

func TestBorrowIntegration(t *testing.T) {
	suite.Run(t, new(BorrowIntegrationTests))
}

// SetupTest is run automatically before each suite test
func (suite *BorrowIntegrationTests) SetupTest() {
	_, suite.addrs = app.GeneratePrivKeyAddressPairs(5)

	suite.genesisTime = time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC)
}

func (suite *BorrowIntegrationTests) TestSingleUserAccumulatesRewardsAfterSyncing() {
	userA := suite.addrs[0]

	authBulder := app.NewAuthBankGenesisBuilder().
		WithSimpleModuleAccount(types2.ModuleName, cs(c("uexam", 1e18))). // Fill mint with enough coins to pay out any reward
		WithSimpleAccount(userA, cs(c("bnb", 1e12)))                      // give the user some coins

	incentBuilder := testutil2.NewIncentiveGenesisBuilder().
		WithGenesisTime(suite.genesisTime).
		WithMultipliers(types2.MultipliersPerDenoms{{
			Denom:       "uexam",
			Multipliers: types2.Multipliers{types2.NewMultiplier("large", 12, d("1.0"))}, // keep payout at 1.0 to make maths easier
		}}).
		WithSimpleBorrowRewardPeriod("bnb", cs(c("uexam", 1e6))) // only borrow rewards

	suite.SetApp()

	var genAcc []authtypes.GenesisAccount
	for _, el := range suite.addrs {
		b := authtypes.NewBaseAccount(el, nil, 0, 0)
		genAcc = append(genAcc, b)
	}

	suite.StartChain(genAcc, cs(c("bnb", 1e12)), suite.genesisTime,
		NewPricefeedGenStateMultiFromTime(suite.App.AppCodec(), suite.genesisTime),
		NewJoltGenStateMulti(suite.genesisTime).BuildMarshalled(suite.App.AppCodec()),
		authBulder.BuildMarshalled(suite.App.AppCodec()),
		incentBuilder.BuildMarshalled(suite.App.AppCodec()),
	)

	err := fundModuleAccount(suite.App.GetBankKeeper(), suite.Ctx, types2.ModuleName, cs(c("uexam", 1e18)))
	suite.Require().NoError(err)

	// Create a borrow (need to first deposit to allow it)
	suite.NoError(suite.DeliverJoltMsgDeposit(userA, cs(c("bnb", 1e11))))
	suite.NoError(suite.DeliverJoltMsgBorrow(userA, cs(c("bnb", 1e10))))

	// Let time pass to accumulate interest on the borrow
	// Use one long block instead of many to reduce any rounding errors, and speed up tests.
	suite.NextBlockAfter(1e6 * time.Second) // about 12 days

	// User borrows and repays just to sync their borrow.
	suite.NoError(suite.DeliverHardMsgRepay(userA, cs(c("bnb", 1))))
	suite.NoError(suite.DeliverJoltMsgBorrow(userA, cs(c("bnb", 1))))

	// Accumulate more rewards.
	// The user still has the same percentage of all borrows (100%) so their rewards should be the same as in the previous block.
	suite.NextBlockAfter(1e6 * time.Second) // about 12 days

	msg := types2.NewMsgClaimJoltReward(userA.String(), types2.Selections{
		types2.NewSelection("uexam", "large"),
	})

	// User claims all their rewards
	suite.NoError(suite.DeliverIncentiveMsg(&msg))

	// The users has always had 100% of borrows, so they should receive all rewards for the previous two blocks.
	// Total rewards for each block is block duration * rewards per second
	// we need to add 100000000000000stake token as it is initlised in genesis
	accuracy := 1e-10 // using a very high accuracy to flag future small calculation changes
	suite.BalanceInEpsilon(userA, cs(c("bnb", 1e12-1e11+1e10), c("stake", 100000000000000), c("uexam", 2*1e6*1e6)), accuracy)
}

// Test suite used for all keeper tests
type BorrowRewardsTestSuite struct {
	suite.Suite

	keeper     keeper.Keeper
	joltKeeper joltkeeper.Keeper

	app app.TestApp
	ctx context.Context

	genesisTime time.Time
	addrs       []sdk.AccAddress
}

// SetupTest is run automatically before each suite test
func (suite *BorrowRewardsTestSuite) SetupTest() {
	config := sdk.GetConfig()
	app.SetBech32AddressPrefixes(config)

	_, suite.addrs = app.GeneratePrivKeyAddressPairs(5)

	suite.genesisTime = time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC)
}

func (suite *BorrowRewardsTestSuite) SetupApp() {
	suite.app = app.NewTestApp(tmlog.TestingLogger(), suite.T().TempDir())

	suite.keeper = suite.app.GetIncentiveKeeper()
	suite.joltKeeper = suite.app.GetJoltKeeper()

	suite.ctx = suite.app.NewContext(true, tmproto.Header{Height: 1, Time: suite.genesisTime})
}

func (suite *BorrowRewardsTestSuite) SetupWithGenState(authBuilder *app.AuthBankGenesisBuilder, incentBuilder testutil2.IncentiveGenesisBuilder, hardBuilder testutil2.JoltGenesisBuilder) {
	suite.SetupApp()

	suite.app.InitializeFromGenesisStatesWithTime(
		suite.genesisTime, nil, nil,
		authBuilder.BuildMarshalled(suite.app.AppCodec()),
		NewPricefeedGenStateMultiFromTime(suite.app.AppCodec(), suite.genesisTime),
		hardBuilder.BuildMarshalled(suite.app.AppCodec()),
		incentBuilder.BuildMarshalled(suite.app.AppCodec()),
	)
}

func (suite *BorrowRewardsTestSuite) TestAccumulateHardBorrowRewards() {
	type args struct {
		borrow                sdk.Coin
		rewardsPerSecond      sdk.Coins
		timeElapsed           int
		expectedRewardIndexes types2.RewardIndexes
	}
	type test struct {
		name string
		args args
	}
	testCases := []test{
		{
			"single reward denom: 7 seconds",
			args{
				borrow:                c("bnb", 1000000000000),
				rewardsPerSecond:      cs(c("hard", 122354)),
				timeElapsed:           7,
				expectedRewardIndexes: types2.RewardIndexes{types2.NewRewardIndex("hard", d("856478.000000741989898080"))},
			},
		},
		{
			"single reward denom: 1 day",
			args{
				borrow:                c("bnb", 1000000000000),
				rewardsPerSecond:      cs(c("hard", 122354)),
				timeElapsed:           86400,
				expectedRewardIndexes: types2.RewardIndexes{types2.NewRewardIndex("hard", d("10571385600.010177134215704527"))},
			},
		},
		{
			"single reward denom: 0 seconds",
			args{
				borrow:                c("bnb", 1000000000000),
				rewardsPerSecond:      cs(c("hard", 122354)),
				timeElapsed:           0,
				expectedRewardIndexes: types2.RewardIndexes{types2.NewRewardIndex("hard", d("0.0"))},
			},
		},
		{
			"multiple reward denoms: 7 seconds",
			args{
				borrow:           c("bnb", 1000000000000),
				rewardsPerSecond: cs(c("hard", 122354), c("ujolt", 122354)),
				timeElapsed:      7,
				expectedRewardIndexes: types2.RewardIndexes{
					types2.NewRewardIndex("hard", d("856478.000000741989898080")),
					types2.NewRewardIndex("ujolt", d("856478.000000741989898080")),
				},
			},
		},
		{
			"multiple reward denoms: 1 day",
			args{
				borrow:           c("bnb", 1000000000000),
				rewardsPerSecond: cs(c("hard", 122354), c("ujolt", 122354)),
				timeElapsed:      86400,
				expectedRewardIndexes: types2.RewardIndexes{
					types2.NewRewardIndex("hard", d("10571385600.010177134215704527")),
					types2.NewRewardIndex("ujolt", d("10571385600.010177134215704527")),
				},
			},
		},
		{
			"multiple reward denoms: 0 seconds",
			args{
				borrow:           c("bnb", 1000000000000),
				rewardsPerSecond: cs(c("hard", 122354), c("ujolt", 122354)),
				timeElapsed:      0,
				expectedRewardIndexes: types2.RewardIndexes{
					types2.NewRewardIndex("hard", d("0.0")),
					types2.NewRewardIndex("ujolt", d("0.0")),
				},
			},
		},
		{
			"multiple reward denoms with different rewards per second: 1 day",
			args{
				borrow:           c("bnb", 1000000000000),
				rewardsPerSecond: cs(c("hard", 122354), c("ujolt", 555555)),
				timeElapsed:      86400,
				expectedRewardIndexes: types2.RewardIndexes{
					types2.NewRewardIndex("hard", d("10571385600.010177134215704527")),
					types2.NewRewardIndex("ujolt", d("47999952000.046209832119961168")),
				},
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			userAddr := suite.addrs[3]
			authBuilder := app.NewAuthBankGenesisBuilder().WithSimpleAccount(
				userAddr,
				cs(c("bnb", 1e15), c("ujolt", 1e15), c("btcb", 1e15), c("xrp", 1e15), c("zzz", 1e15)),
			)

			incentBuilder := testutil2.NewIncentiveGenesisBuilder().
				WithGenesisTime(suite.genesisTime).
				WithSimpleBorrowRewardPeriod(tc.args.borrow.Denom, tc.args.rewardsPerSecond)

			suite.SetupWithGenState(authBuilder, incentBuilder, NewJoltGenStateMulti(suite.genesisTime))

			err := fundAccount(suite.app.GetBankKeeper(), suite.ctx, userAddr, cs(c("bnb", 1e15), c("ujolt", 1e15), c("btcb", 1e15), c("xrp", 1e15), c("zzz", 1e15)))
			suite.Require().NoError(err)

			// User deposits and borrows to increase total borrowed amount
			err = suite.joltKeeper.Deposit(suite.ctx, userAddr, sdk.NewCoins(sdk.NewCoin(tc.args.borrow.Denom, tc.args.borrow.Amount.Mul(sdkmath.NewInt(2)))))
			suite.Require().NoError(err)
			err = suite.joltKeeper.Borrow(suite.ctx, userAddr, sdk.NewCoins(tc.args.borrow))
			suite.Require().NoError(err)

			// Set up chain context at future time
			runAtTime := suite.ctx.BlockTime().Add(time.Duration(int(time.Second) * tc.args.timeElapsed))
			runCtx := suite.ctx.WithBlockTime(runAtTime)

			// Run Hard begin blocker in order to update the denom's index factor
			jolt.BeginBlocker(runCtx, suite.joltKeeper)

			// Accumulate hard borrow rewards for the deposit denom
			multiRewardPeriod, found := suite.keeper.GetJoltBorrowRewardPeriods(runCtx, tc.args.borrow.Denom)
			suite.Require().True(found)
			suite.keeper.AccumulateJoltBorrowRewards(runCtx, multiRewardPeriod)

			// Check that each expected reward index matches the current stored reward index for the denom
			globalRewardIndexes, found := suite.keeper.GetJoltBorrowRewardIndexes(runCtx, tc.args.borrow.Denom)
			suite.Require().True(found)
			for _, expectedRewardIndex := range tc.args.expectedRewardIndexes {
				globalRewardIndex, found := globalRewardIndexes.GetRewardIndex(expectedRewardIndex.CollateralType)
				suite.Require().True(found)
				suite.Require().Equal(expectedRewardIndex, globalRewardIndex)
			}
		})
	}
}

func (suite *BorrowRewardsTestSuite) TestInitializeHardBorrowRewards() {
	type args struct {
		moneyMarketRewardDenoms          map[string]sdk.Coins
		deposit                          sdk.Coins
		borrow                           sdk.Coins
		expectedClaimBorrowRewardIndexes types2.MultiRewardIndexes
	}
	type test struct {
		name string
		args args
	}

	standardMoneyMarketRewardDenoms := map[string]sdk.Coins{
		"bnb":  cs(c("hard", 1)),
		"btcb": cs(c("hard", 1), c("ujolt", 1)),
	}

	testCases := []test{
		{
			"single deposit denom, single reward denom",
			args{
				moneyMarketRewardDenoms: standardMoneyMarketRewardDenoms,
				deposit:                 cs(c("bnb", 1000000000000)),
				borrow:                  cs(c("bnb", 100000000000)),
				expectedClaimBorrowRewardIndexes: types2.MultiRewardIndexes{
					types2.NewMultiRewardIndex(
						"bnb",
						types2.RewardIndexes{
							types2.NewRewardIndex("hard", d("0.0")),
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
				borrow:                  cs(c("btcb", 100000000000)),
				expectedClaimBorrowRewardIndexes: types2.MultiRewardIndexes{
					types2.NewMultiRewardIndex(
						"btcb",
						types2.RewardIndexes{
							types2.NewRewardIndex("hard", d("0.0")),
							types2.NewRewardIndex("ujolt", d("0.0")),
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
				borrow:                  cs(c("xrp", 100000000000)),
				expectedClaimBorrowRewardIndexes: types2.MultiRewardIndexes{
					types2.NewMultiRewardIndex(
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
				borrow:                  cs(c("bnb", 100000000000), c("btcb", 100000000000)),
				expectedClaimBorrowRewardIndexes: types2.MultiRewardIndexes{
					types2.NewMultiRewardIndex(
						"bnb",
						types2.RewardIndexes{
							types2.NewRewardIndex("hard", d("0.0")),
						},
					),
					types2.NewMultiRewardIndex(
						"btcb",
						types2.RewardIndexes{
							types2.NewRewardIndex("hard", d("0.0")),
							types2.NewRewardIndex("ujolt", d("0.0")),
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
				borrow:                  cs(c("bnb", 100000000000), c("xrp", 100000000000)),
				expectedClaimBorrowRewardIndexes: types2.MultiRewardIndexes{
					types2.NewMultiRewardIndex(
						"bnb",
						types2.RewardIndexes{
							types2.NewRewardIndex("hard", d("0.0")),
						},
					),
					types2.NewMultiRewardIndex(
						"xrp",
						nil,
					),
				},
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			userAddr := suite.addrs[3]
			authBuilder := app.NewAuthBankGenesisBuilder().WithSimpleAccount(
				userAddr,
				cs(c("bnb", 1e15), c("ujolt", 1e15), c("btcb", 1e15), c("xrp", 1e15), c("zzz", 1e15)),
			)

			incentBuilder := testutil2.NewIncentiveGenesisBuilder().WithGenesisTime(suite.genesisTime)
			for moneyMarketDenom, rewardsPerSecond := range tc.args.moneyMarketRewardDenoms {
				incentBuilder = incentBuilder.WithSimpleBorrowRewardPeriod(moneyMarketDenom, rewardsPerSecond)
			}

			suite.SetupWithGenState(authBuilder, incentBuilder, NewJoltGenStateMulti(suite.genesisTime))

			err := fundAccount(suite.app.GetBankKeeper(), suite.ctx, userAddr, cs(c("bnb", 1e15), c("ujolt", 1e15), c("btcb", 1e15), c("xrp", 1e15), c("zzz", 1e15)))
			suite.Require().NoError(err)

			// User deposits
			err = suite.joltKeeper.Deposit(suite.ctx, userAddr, tc.args.deposit)
			suite.Require().NoError(err)
			// User borrows
			err = suite.joltKeeper.Borrow(suite.ctx, userAddr, tc.args.borrow)
			suite.Require().NoError(err)

			claim, foundClaim := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, userAddr)
			suite.Require().True(foundClaim)
			suite.Require().Equal(tc.args.expectedClaimBorrowRewardIndexes, claim.BorrowRewardIndexes)
		})
	}
}

func (suite *BorrowRewardsTestSuite) TestSynchronizeHardBorrowReward() {
	type args struct {
		incentiveBorrowRewardDenom   string
		borrow                       sdk.Coin
		rewardsPerSecond             sdk.Coins
		blockTimes                   []int
		expectedRewardIndexes        types2.RewardIndexes
		expectedRewards              sdk.Coins
		updateRewardsViaCommmittee   bool
		updatedBaseDenom             string
		updatedRewardsPerSecond      sdk.Coins
		updatedExpectedRewardIndexes types2.RewardIndexes
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
				incentiveBorrowRewardDenom: "bnb",
				borrow:                     c("bnb", 10000000000),
				rewardsPerSecond:           cs(c("hard", 122354)),
				blockTimes:                 []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
				expectedRewardIndexes:      types2.RewardIndexes{types2.NewRewardIndex("hard", d("1223540000.173229514684687054"))},
				expectedRewards:            cs(c("hard", 12235400)),
				updateRewardsViaCommmittee: false,
			},
		},
		{
			"single reward denom: 10 blocks - long block time",
			args{
				incentiveBorrowRewardDenom: "bnb",
				borrow:                     c("bnb", 10000000000),
				rewardsPerSecond:           cs(c("hard", 122354)),
				blockTimes:                 []int{86400, 86400, 86400, 86400, 86400, 86400, 86400, 86400, 86400, 86400},
				expectedRewardIndexes:      types2.RewardIndexes{types2.NewRewardIndex("hard", d("10571385603126.235338908064289886"))},
				expectedRewards:            cs(c("hard", 105713856031)),
			},
		},
		{
			"single reward denom: user reward index updated when reward is zero",
			args{
				incentiveBorrowRewardDenom: "pjolt",
				borrow:                     c("pjolt", 1), // borrow a tiny amount so that rewards round to zero
				rewardsPerSecond:           cs(c("inc", 122354)),
				blockTimes:                 []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
				expectedRewardIndexes:      types2.RewardIndexes{types2.NewRewardIndex("inc", d("122354003908.172327815873921369"))},
				expectedRewards:            cs(),
				updateRewardsViaCommmittee: false,
			},
		},
		{
			"multiple reward denoms: 10 blocks",
			args{
				incentiveBorrowRewardDenom: "bnb",
				borrow:                     c("bnb", 10000000000),
				rewardsPerSecond:           cs(c("hard", 122354), c("ujolt", 122354)),
				blockTimes:                 []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
				expectedRewardIndexes: types2.RewardIndexes{
					types2.NewRewardIndex("hard", d("1223540000.173229514684687054")),
					types2.NewRewardIndex("ujolt", d("1223540000.173229514684687054")),
				},
				expectedRewards: cs(c("hard", 12235400), c("ujolt", 12235400)),
			},
		},
		{
			"multiple reward denoms: 10 blocks - long block time",
			args{
				incentiveBorrowRewardDenom: "bnb",
				borrow:                     c("bnb", 10000000000),
				rewardsPerSecond:           cs(c("hard", 122354), c("ujolt", 122354)),
				blockTimes:                 []int{86400, 86400, 86400, 86400, 86400, 86400, 86400, 86400, 86400, 86400},
				expectedRewardIndexes: types2.RewardIndexes{
					types2.NewRewardIndex("hard", d("10571385603126.235338908064289886")),
					types2.NewRewardIndex("ujolt", d("10571385603126.235338908064289886")),
				},
				expectedRewards: cs(c("hard", 105713856031), c("ujolt", 105713856031)),
			},
		},
		{
			"multiple reward denoms with different rewards per second: 10 blocks",
			args{
				incentiveBorrowRewardDenom: "bnb",
				borrow:                     c("bnb", 10000000000),
				rewardsPerSecond:           cs(c("hard", 122354), c("ujolt", 555555)),
				blockTimes:                 []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
				expectedRewardIndexes: types2.RewardIndexes{
					types2.NewRewardIndex("hard", d("1223540000.173229514684687054")),
					types2.NewRewardIndex("ujolt", d("5555550000.786558044940511272")),
				},
				expectedRewards: cs(c("hard", 12235400), c("ujolt", 55555500)),
			},
		},
		{
			"denom is in incentive's jolt borrow reward params and has rewards; add new reward type",
			args{
				incentiveBorrowRewardDenom: "bnb",
				borrow:                     c("bnb", 10000000000),
				rewardsPerSecond:           cs(c("hard", 122354)),
				blockTimes:                 []int{86400},
				expectedRewardIndexes: types2.RewardIndexes{
					types2.NewRewardIndex("hard", d("1057138560060.101159956651277179")),
				},
				// fixme the old one is 10571385600 me new reward, 21142771202 -> 21142771200
				expectedRewards:            cs(c("hard", 10571385600)),
				updateRewardsViaCommmittee: true,
				updatedBaseDenom:           "bnb",
				updatedRewardsPerSecond:    cs(c("hard", 122354), c("ujolt", 100000)),
				updatedExpectedRewards:     cs(c("hard", 21142771200), c("ujolt", 8640000000)),
				updatedExpectedRewardIndexes: types2.RewardIndexes{
					types2.NewRewardIndex("hard", d("2114277120120.202319913302554358")),
					types2.NewRewardIndex("ujolt", d("864000000049.120715266073260522")),
				},
				updatedTimeDuration: 86400,
			},
		},
		{
			"denom is in hard's money market params but not in incentive's hard supply reward params; add reward",
			args{
				incentiveBorrowRewardDenom: "bnb",
				borrow:                     c("zzz", 10000000000),
				rewardsPerSecond:           nil,
				blockTimes:                 []int{100},
				expectedRewardIndexes:      types2.RewardIndexes{},
				expectedRewards:            sdk.Coins{},
				updateRewardsViaCommmittee: true,
				updatedBaseDenom:           "zzz",
				updatedRewardsPerSecond:    cs(c("hard", 100000)),
				updatedExpectedRewards:     cs(c("hard", 8640000000)),
				updatedExpectedRewardIndexes: types2.RewardIndexes{
					types2.NewRewardIndex("hard", d("864000000049.803065390262558660")),
				},
				updatedTimeDuration: 86400,
			},
		},
		{
			"denom is in hard's money market params but not in incentive's hard supply reward params; add multiple reward types",
			args{
				incentiveBorrowRewardDenom: "bnb",
				borrow:                     c("zzz", 10000000000),
				rewardsPerSecond:           nil,
				blockTimes:                 []int{100},
				expectedRewardIndexes:      types2.RewardIndexes{},
				expectedRewards:            sdk.Coins{},
				updateRewardsViaCommmittee: true,
				updatedBaseDenom:           "zzz",
				updatedRewardsPerSecond:    cs(c("hard", 100000), c("ujolt", 100500), c("swap", 500)),
				updatedExpectedRewards:     cs(c("hard", 8640000000), c("ujolt", 8683200000), c("swap", 43200000)),
				updatedExpectedRewardIndexes: types2.RewardIndexes{
					types2.NewRewardIndex("hard", d("864000000049.803065390262558660")),
					types2.NewRewardIndex("ujolt", d("868320000050.052080717213871453")),
					types2.NewRewardIndex("swap", d("4320000000.249015326951312793")),
				},
				updatedTimeDuration: 86400,
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			userAddr := suite.addrs[3]
			authBuilder := app.NewAuthBankGenesisBuilder().
				WithSimpleAccount(suite.addrs[2], cs(c("pjolt", 1e9))).
				WithSimpleAccount(userAddr, cs(c("pjolt", 1e9))).
				WithSimpleAccount(userAddr, cs(c("bnb", 1e15), c("ujolt", 1e15), c("btcb", 1e15), c("xrp", 1e15), c("zzz", 1e15)))

			incentBuilder := testutil2.NewIncentiveGenesisBuilder().WithGenesisTime(suite.genesisTime)
			if tc.args.rewardsPerSecond != nil {
				incentBuilder = incentBuilder.WithSimpleBorrowRewardPeriod(tc.args.incentiveBorrowRewardDenom, tc.args.rewardsPerSecond)
			}
			// Set the minimum borrow to 0 to allow testing small borrows
			hardBuilder := NewJoltGenStateMulti(suite.genesisTime).WithMinBorrow(sdkmath.LegacyZeroDec())

			suite.SetupWithGenState(authBuilder, incentBuilder, hardBuilder)

			err := fundAccount(suite.app.GetBankKeeper(), suite.ctx, suite.addrs[2], cs(c("pjolt", 1e9)))
			suite.Require().NoError(err)

			err = fundAccount(suite.app.GetBankKeeper(), suite.ctx, userAddr, cs(c("pjolt", 1e9), c("bnb", 1e15), c("ujolt", 1e15), c("btcb", 1e15), c("xrp", 1e15), c("zzz", 1e15)))
			suite.Require().NoError(err)

			// Borrow a fixed amount from another user to dilute primary user's rewards per second.
			suite.Require().NoError(
				suite.joltKeeper.Deposit(suite.ctx, suite.addrs[2], cs(c("pjolt", 200_000_000))),
			)
			suite.Require().NoError(
				suite.joltKeeper.Borrow(suite.ctx, suite.addrs[2], cs(c("pjolt", 100_000_000))),
			)

			// User deposits and borrows to increase total borrowed amount
			err = suite.joltKeeper.Deposit(suite.ctx, userAddr, sdk.NewCoins(sdk.NewCoin(tc.args.borrow.Denom, tc.args.borrow.Amount.Mul(sdkmath.NewInt(2)))))
			suite.Require().NoError(err)
			err = suite.joltKeeper.Borrow(suite.ctx, userAddr, sdk.NewCoins(tc.args.borrow))
			suite.Require().NoError(err)

			// Check that Hard hooks initialized a HardLiquidityProviderClaim
			claim, found := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, userAddr)
			suite.Require().True(found)
			multiRewardIndex, _ := claim.BorrowRewardIndexes.GetRewardIndex(tc.args.borrow.Denom)
			for _, expectedRewardIndex := range tc.args.expectedRewardIndexes {
				currRewardIndex, found := multiRewardIndex.RewardIndexes.GetRewardIndex(expectedRewardIndex.CollateralType)
				suite.Require().True(found)
				suite.Require().Equal(sdkmath.LegacyZeroDec(), currRewardIndex.RewardFactor)
			}

			// Run accumulator at several intervals
			var timeElapsed int
			previousBlockTime := suite.ctx.BlockTime()
			for _, t := range tc.args.blockTimes {
				timeElapsed += t
				updatedBlockTime := previousBlockTime.Add(time.Duration(int(time.Second) * t))
				previousBlockTime = updatedBlockTime
				blockCtx := suite.ctx.WithBlockTime(updatedBlockTime)

				// Run Hard begin blocker for each block ctx to update denom's interest factor
				jolt.BeginBlocker(blockCtx, suite.joltKeeper)

				// Accumulate hard borrow-side rewards
				multiRewardPeriod, found := suite.keeper.GetJoltBorrowRewardPeriods(blockCtx, tc.args.borrow.Denom)
				if found {
					suite.keeper.AccumulateJoltBorrowRewards(blockCtx, multiRewardPeriod)
				}
			}
			updatedBlockTime := suite.ctx.BlockTime().Add(time.Duration(int(time.Second) * timeElapsed))
			suite.ctx = suite.ctx.WithBlockTime(updatedBlockTime)

			// After we've accumulated, run synchronize
			borrow, found := suite.joltKeeper.GetBorrow(suite.ctx, userAddr)
			suite.Require().True(found)
			suite.Require().NotPanics(func() {
				suite.keeper.SynchronizeJoltBorrowReward(suite.ctx, borrow)
			})

			// Check that the global reward index's reward factor and user's claim have been updated as expected
			claim, found = suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, userAddr)
			suite.Require().True(found)
			globalRewardIndexes, foundGlobalRewardIndexes := suite.keeper.GetJoltBorrowRewardIndexes(suite.ctx, tc.args.borrow.Denom)
			if len(tc.args.rewardsPerSecond) > 0 {
				suite.Require().True(foundGlobalRewardIndexes)
				for _, expectedRewardIndex := range tc.args.expectedRewardIndexes {
					// Check that global reward index has been updated as expected
					globalRewardIndex, found := globalRewardIndexes.GetRewardIndex(expectedRewardIndex.CollateralType)
					suite.Require().True(found)
					suite.Require().Equal(expectedRewardIndex, globalRewardIndex)

					// Check that the user's claim's reward index matches the corresponding global reward index
					multiRewardIndex, found := claim.BorrowRewardIndexes.GetRewardIndex(tc.args.borrow.Denom)
					suite.Require().True(found)
					rewardIndex, found := multiRewardIndex.RewardIndexes.GetRewardIndex(expectedRewardIndex.CollateralType)
					suite.Require().True(found)
					suite.Require().Equal(expectedRewardIndex, rewardIndex)

					// Check that the user's claim holds the expected amount of reward coins
					suite.Require().Equal(
						tc.args.expectedRewards.AmountOf(expectedRewardIndex.CollateralType).String(),
						claim.Reward.AmountOf(expectedRewardIndex.CollateralType).String(),
					)
				}
			}

			// Only test cases with reward param updates continue past this point
			if !tc.args.updateRewardsViaCommmittee {
				return
			}

			// If are no initial rewards per second, add new rewards through a committee param change
			// 1. Construct incentive's new JoltBorrowRewardPeriods param
			currIncentiveHardBorrowRewardPeriods := suite.keeper.GetParams(suite.ctx).JoltBorrowRewardPeriods
			multiRewardPeriod, found := currIncentiveHardBorrowRewardPeriods.GetMultiRewardPeriod(tc.args.borrow.Denom)
			if found {
				// Borrow denom's reward period exists, but it doesn't have any rewards per second
				index, found := currIncentiveHardBorrowRewardPeriods.GetMultiRewardPeriodIndex(tc.args.borrow.Denom)
				suite.Require().True(found)
				multiRewardPeriod.RewardsPerSecond = tc.args.updatedRewardsPerSecond
				currIncentiveHardBorrowRewardPeriods[index] = multiRewardPeriod
			} else {
				// Borrow denom's reward period does not exist
				_, found := currIncentiveHardBorrowRewardPeriods.GetMultiRewardPeriodIndex(tc.args.borrow.Denom)
				suite.Require().False(found)
				newMultiRewardPeriod := types2.NewMultiRewardPeriod(true, tc.args.borrow.Denom, suite.genesisTime, suite.genesisTime.Add(time.Hour*24*365*4), tc.args.updatedRewardsPerSecond)
				currIncentiveHardBorrowRewardPeriods = append(currIncentiveHardBorrowRewardPeriods, newMultiRewardPeriod)
			}

			params := suite.keeper.GetParams(suite.ctx)
			params.JoltBorrowRewardPeriods = currIncentiveHardBorrowRewardPeriods
			suite.keeper.SetParams(suite.ctx, params)

			// We need to accumulate hard supply-side rewards again
			multiRewardPeriod, found = suite.keeper.GetJoltBorrowRewardPeriods(suite.ctx, tc.args.borrow.Denom)
			suite.Require().True(found)

			// But new borrow denoms don't have their PreviousHardBorrowRewardAccrualTime set yet,
			// so we need to call the accumulation method once to set the initial reward accrual time
			if tc.args.borrow.Denom != tc.args.incentiveBorrowRewardDenom {
				suite.keeper.AccumulateJoltBorrowRewards(suite.ctx, multiRewardPeriod)
			}

			// Now we can jump forward in time and accumulate rewards
			updatedBlockTime = previousBlockTime.Add(time.Duration(int(time.Second) * tc.args.updatedTimeDuration))
			suite.ctx = suite.ctx.WithBlockTime(updatedBlockTime)
			suite.keeper.AccumulateJoltBorrowRewards(suite.ctx, multiRewardPeriod)

			// After we've accumulated, run synchronize
			borrow, found = suite.joltKeeper.GetBorrow(suite.ctx, userAddr)
			suite.Require().True(found)
			suite.Require().NotPanics(func() {
				suite.keeper.SynchronizeJoltBorrowReward(suite.ctx, borrow)
			})

			// Check that the global reward index's reward factor and user's claim have been updated as expected
			globalRewardIndexes, found = suite.keeper.GetJoltBorrowRewardIndexes(suite.ctx, tc.args.borrow.Denom)
			suite.Require().True(found)
			claim, found = suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, userAddr)
			suite.Require().True(found)

			for _, expectedRewardIndex := range tc.args.updatedExpectedRewardIndexes {
				// Check that global reward index has been updated as expected
				globalRewardIndex, found := globalRewardIndexes.GetRewardIndex(expectedRewardIndex.CollateralType)
				suite.Require().True(found)
				suite.Require().Equal(expectedRewardIndex, globalRewardIndex)
				// Check that the user's claim's reward index matches the corresponding global reward index
				multiRewardIndex, found := claim.BorrowRewardIndexes.GetRewardIndex(tc.args.borrow.Denom)
				suite.Require().True(found)
				rewardIndex, found := multiRewardIndex.RewardIndexes.GetRewardIndex(expectedRewardIndex.CollateralType)
				suite.Require().True(found)
				suite.Require().Equal(expectedRewardIndex, rewardIndex)

				// Check that the user's claim holds the expected amount of reward coins
				suite.Require().Equal(
					tc.args.updatedExpectedRewards.AmountOf(expectedRewardIndex.CollateralType).String(),
					claim.Reward.AmountOf(expectedRewardIndex.CollateralType).String(),
				)
			}
		})
	}
}

func (suite *BorrowRewardsTestSuite) TestUpdateHardBorrowIndexDenoms() {
	type withdrawModification struct {
		coins sdk.Coins
		repay bool
	}

	type args struct {
		initialDeposit            sdk.Coins
		firstBorrow               sdk.Coins
		modification              withdrawModification
		rewardsPerSecond          sdk.Coins
		expectedBorrowIndexDenoms []string
	}
	type test struct {
		name string
		args args
	}

	testCases := []test{
		{
			"single reward denom: update adds one borrow reward index",
			args{
				initialDeposit:            cs(c("bnb", 10000000000)),
				firstBorrow:               cs(c("bnb", 50000000)),
				modification:              withdrawModification{coins: cs(c("ujolt", 500000000))},
				rewardsPerSecond:          cs(c("hard", 122354)),
				expectedBorrowIndexDenoms: []string{"bnb", "ujolt"},
			},
		},
		{
			"single reward denom: update adds multiple borrow supply reward indexes",
			args{
				initialDeposit:            cs(c("btcb", 10000000000)),
				firstBorrow:               cs(c("btcb", 50000000)),
				modification:              withdrawModification{coins: cs(c("ujolt", 500000000), c("bnb", 50000000000), c("xrp", 50000000000))},
				rewardsPerSecond:          cs(c("hard", 122354)),
				expectedBorrowIndexDenoms: []string{"btcb", "ujolt", "bnb", "xrp"},
			},
		},
		{
			"single reward denom: update doesn't add duplicate borrow reward index for same denom",
			args{
				initialDeposit:            cs(c("bnb", 100000000000)),
				firstBorrow:               cs(c("bnb", 50000000)),
				modification:              withdrawModification{coins: cs(c("bnb", 50000000000))},
				rewardsPerSecond:          cs(c("hard", 122354)),
				expectedBorrowIndexDenoms: []string{"bnb"},
			},
		},
		{
			"multiple reward denoms: update adds one borrow reward index",
			args{
				initialDeposit:            cs(c("bnb", 10000000000)),
				firstBorrow:               cs(c("bnb", 50000000)),
				modification:              withdrawModification{coins: cs(c("ujolt", 500000000))},
				rewardsPerSecond:          cs(c("hard", 122354), c("ujolt", 122354)),
				expectedBorrowIndexDenoms: []string{"bnb", "ujolt"},
			},
		},
		{
			"multiple reward denoms: update adds multiple borrow supply reward indexes",
			args{
				initialDeposit:            cs(c("btcb", 10000000000)),
				firstBorrow:               cs(c("btcb", 50000000)),
				modification:              withdrawModification{coins: cs(c("ujolt", 500000000), c("bnb", 50000000000), c("xrp", 50000000000))},
				rewardsPerSecond:          cs(c("hard", 122354), c("ujolt", 122354)),
				expectedBorrowIndexDenoms: []string{"btcb", "ujolt", "bnb", "xrp"},
			},
		},
		{
			"multiple reward denoms: update doesn't add duplicate borrow reward index for same denom",
			args{
				initialDeposit:            cs(c("bnb", 100000000000)),
				firstBorrow:               cs(c("bnb", 50000000)),
				modification:              withdrawModification{coins: cs(c("bnb", 50000000000))},
				rewardsPerSecond:          cs(c("hard", 122354), c("ujolt", 122354)),
				expectedBorrowIndexDenoms: []string{"bnb"},
			},
		},
		{
			"single reward denom: fully repaying a denom deletes the denom's supply reward index",
			args{
				initialDeposit:            cs(c("bnb", 1000000000)),
				firstBorrow:               cs(c("bnb", 100000000)),
				modification:              withdrawModification{coins: cs(c("bnb", 1100000000)), repay: true},
				rewardsPerSecond:          cs(c("hard", 122354)),
				expectedBorrowIndexDenoms: []string{},
			},
		},
		{
			"single reward denom: fully repaying a denom deletes only the denom's supply reward index",
			args{
				initialDeposit:            cs(c("bnb", 1000000000)),
				firstBorrow:               cs(c("bnb", 100000000), c("ujolt", 10000000)),
				modification:              withdrawModification{coins: cs(c("bnb", 1100000000)), repay: true},
				rewardsPerSecond:          cs(c("hard", 122354)),
				expectedBorrowIndexDenoms: []string{"ujolt"},
			},
		},
		{
			"multiple reward denoms: fully repaying a denom deletes the denom's supply reward index",
			args{
				initialDeposit:            cs(c("bnb", 1000000000)),
				firstBorrow:               cs(c("bnb", 100000000), c("ujolt", 10000000)),
				modification:              withdrawModification{coins: cs(c("bnb", 1100000000)), repay: true},
				rewardsPerSecond:          cs(c("hard", 122354), c("ujolt", 122354)),
				expectedBorrowIndexDenoms: []string{"ujolt"},
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			userAddr := suite.addrs[3]
			authBuilder := app.NewAuthBankGenesisBuilder().
				WithSimpleAccount(
					userAddr,
					cs(c("bnb", 1e15), c("ujolt", 1e15), c("btcb", 1e15), c("xrp", 1e15), c("zzz", 1e15)),
				).
				WithSimpleAccount(
					suite.addrs[0],
					cs(c("bnb", 1e15), c("ujolt", 1e15), c("btcb", 1e15), c("xrp", 1e15), c("zzz", 1e15)),
				)

			incentBuilder := testutil2.NewIncentiveGenesisBuilder().
				WithGenesisTime(suite.genesisTime).
				WithSimpleBorrowRewardPeriod("bnb", tc.args.rewardsPerSecond).
				WithSimpleBorrowRewardPeriod("ujolt", tc.args.rewardsPerSecond).
				WithSimpleBorrowRewardPeriod("btcb", tc.args.rewardsPerSecond).
				WithSimpleBorrowRewardPeriod("xrp", tc.args.rewardsPerSecond)

			suite.SetupWithGenState(authBuilder, incentBuilder, NewJoltGenStateMulti(suite.genesisTime))

			err := fundAccount(suite.app.GetBankKeeper(), suite.ctx, suite.addrs[0], cs(c("bnb", 1e15), c("ujolt", 1e15), c("btcb", 1e15), c("xrp", 1e15), c("zzz", 1e15)))
			suite.Require().NoError(err)

			err = fundAccount(suite.app.GetBankKeeper(), suite.ctx, userAddr, cs(c("bnb", 1e15), c("ujolt", 1e15), c("btcb", 1e15), c("xrp", 1e15), c("zzz", 1e15)))
			suite.Require().NoError(err)

			// Fill the hard supply to allow user to borrow
			err = suite.joltKeeper.Deposit(suite.ctx, suite.addrs[0], tc.args.firstBorrow.Add(tc.args.modification.coins...))
			suite.Require().NoError(err)

			// User deposits initial funds (so that user can borrow)
			err = suite.joltKeeper.Deposit(suite.ctx, userAddr, tc.args.initialDeposit)
			suite.Require().NoError(err)

			// Confirm that claim exists but no borrow reward indexes have been added
			claimAfterDeposit, found := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, userAddr)
			suite.Require().True(found)
			suite.Require().Equal(0, len(claimAfterDeposit.BorrowRewardIndexes))

			// User borrows (first time)
			err = suite.joltKeeper.Borrow(suite.ctx, userAddr, tc.args.firstBorrow)
			suite.Require().NoError(err)

			// Confirm that claim's borrow reward indexes have been updated
			claimAfterFirstBorrow, found := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, userAddr)
			suite.Require().True(found)
			for _, coin := range tc.args.firstBorrow {
				_, hasIndex := claimAfterFirstBorrow.HasBorrowRewardIndex(coin.Denom)
				suite.Require().True(hasIndex)
			}
			suite.Require().True(len(claimAfterFirstBorrow.BorrowRewardIndexes) == len(tc.args.firstBorrow))

			// User modifies their Borrow by either repaying or borrowing more
			if tc.args.modification.repay {
				err = suite.joltKeeper.Repay(suite.ctx, userAddr, userAddr, tc.args.modification.coins)
			} else {
				err = suite.joltKeeper.Borrow(suite.ctx, userAddr, tc.args.modification.coins)
			}
			suite.Require().NoError(err)

			// Confirm that claim's borrow reward indexes contain expected values
			claimAfterModification, found := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, userAddr)
			suite.Require().True(found)
			for _, coin := range tc.args.modification.coins {
				_, hasIndex := claimAfterModification.HasBorrowRewardIndex(coin.Denom)
				if tc.args.modification.repay {
					// Only false if denom is repaid in full
					if tc.args.modification.coins.AmountOf(coin.Denom).GTE(tc.args.firstBorrow.AmountOf(coin.Denom)) {
						suite.Require().False(hasIndex)
					}
				} else {
					suite.Require().True(hasIndex)
				}
			}
			suite.Require().True(len(claimAfterModification.BorrowRewardIndexes) == len(tc.args.expectedBorrowIndexDenoms))
		})
	}
}

func (suite *BorrowRewardsTestSuite) TestSimulateHardBorrowRewardSynchronization() {
	type args struct {
		borrow                sdk.Coin
		rewardsPerSecond      sdk.Coins
		blockTimes            []int
		expectedRewardIndexes types2.RewardIndexes
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
				borrow:                c("bnb", 10000000000),
				rewardsPerSecond:      cs(c("hard", 122354)),
				blockTimes:            []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
				expectedRewardIndexes: types2.RewardIndexes{types2.NewRewardIndex("hard", d("1223540000.173229514684687054"))},
				expectedRewards:       cs(c("hard", 12235400)),
			},
		},
		{
			"10 blocks - long block time",
			args{
				borrow:                c("bnb", 10000000000),
				rewardsPerSecond:      cs(c("hard", 122354)),
				blockTimes:            []int{86400, 86400, 86400, 86400, 86400, 86400, 86400, 86400, 86400, 86400},
				expectedRewardIndexes: types2.RewardIndexes{types2.NewRewardIndex("hard", d("10571385603126.235338908064289886"))},
				expectedRewards:       cs(c("hard", 105713856031)),
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			userAddr := suite.addrs[3]
			authBuilder := app.NewAuthBankGenesisBuilder().WithSimpleAccount(userAddr, cs(c("bnb", 1e15), c("ujolt", 1e15), c("btcb", 1e15), c("xrp", 1e15), c("zzz", 1e15)))

			incentBuilder := testutil2.NewIncentiveGenesisBuilder().
				WithGenesisTime(suite.genesisTime).
				WithSimpleBorrowRewardPeriod(tc.args.borrow.Denom, tc.args.rewardsPerSecond)

			suite.SetupWithGenState(authBuilder, incentBuilder, NewJoltGenStateMulti(suite.genesisTime))

			err := fundAccount(suite.app.GetBankKeeper(), suite.ctx, userAddr, cs(c("bnb", 1e15), c("ujolt", 1e15), c("btcb", 1e15), c("xrp", 1e15), c("zzz", 1e15)))
			suite.Require().NoError(err)

			// User deposits and borrows to increase total borrowed amount
			err = suite.joltKeeper.Deposit(suite.ctx, userAddr, sdk.NewCoins(sdk.NewCoin(tc.args.borrow.Denom, tc.args.borrow.Amount.Mul(sdkmath.NewInt(2)))))
			suite.Require().NoError(err)
			err = suite.joltKeeper.Borrow(suite.ctx, userAddr, sdk.NewCoins(tc.args.borrow))
			suite.Require().NoError(err)

			// Run accumulator at several intervals
			var timeElapsed int
			previousBlockTime := suite.ctx.BlockTime()
			for _, t := range tc.args.blockTimes {
				timeElapsed += t
				updatedBlockTime := previousBlockTime.Add(time.Duration(int(time.Second) * t))
				previousBlockTime = updatedBlockTime
				blockCtx := suite.ctx.WithBlockTime(updatedBlockTime)

				// Run Hard begin blocker for each block ctx to update denom's interest factor
				jolt.BeginBlocker(blockCtx, suite.joltKeeper)

				// Accumulate hard borrow-side rewards
				multiRewardPeriod, found := suite.keeper.GetJoltBorrowRewardPeriods(blockCtx, tc.args.borrow.Denom)
				suite.Require().True(found)
				suite.keeper.AccumulateJoltBorrowRewards(blockCtx, multiRewardPeriod)
			}
			updatedBlockTime := suite.ctx.BlockTime().Add(time.Duration(int(time.Second) * timeElapsed))
			suite.ctx = suite.ctx.WithBlockTime(updatedBlockTime)

			// Confirm that the user's claim hasn't been synced
			claimPre, foundPre := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, userAddr)
			suite.Require().True(foundPre)
			multiRewardIndexPre, _ := claimPre.BorrowRewardIndexes.GetRewardIndex(tc.args.borrow.Denom)
			for _, expectedRewardIndex := range tc.args.expectedRewardIndexes {
				currRewardIndex, found := multiRewardIndexPre.RewardIndexes.GetRewardIndex(expectedRewardIndex.CollateralType)
				suite.Require().True(found)
				suite.Require().Equal(sdkmath.LegacyZeroDec(), currRewardIndex.RewardFactor)
			}

			// Check that the synced claim held in memory has properly simulated syncing
			syncedClaim := suite.keeper.SimulateJoltSynchronization(suite.ctx, claimPre)
			for _, expectedRewardIndex := range tc.args.expectedRewardIndexes {
				// Check that the user's claim's reward index matches the expected reward index
				multiRewardIndex, found := syncedClaim.BorrowRewardIndexes.GetRewardIndex(tc.args.borrow.Denom)
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

func TestBorrowRewardsTestSuite(t *testing.T) {
	suite.Run(t, new(BorrowRewardsTestSuite))
}
