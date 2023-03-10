package keeper_test

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
	minttypes "github.com/joltify-finance/joltify_lending/x/mint/types"
	testutil2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/testutil"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

const secondsPerDay = 24 * 60 * 60

// Test suite used for all keeper tests
type HandlerTestSuite struct {
	testutil2.IntegrationTester

	genesisTime time.Time
	addrs       []sdk.AccAddress
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

// SetupTest is run automatically before each suite test
func (suite *HandlerTestSuite) SetupTest() {
	config := sdk.GetConfig()
	app.SetBech32AddressPrefixes(config)

	_, suite.addrs = app.GeneratePrivKeyAddressPairs(5)

	suite.genesisTime = time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC)
}

func (suite *HandlerTestSuite) SetupApp() {
	suite.App = app.NewTestApp()
	suite.Ctx = suite.App.NewContext(true, tmproto.Header{Height: 1, Time: suite.genesisTime})
}

type genesisBuilder interface {
	BuildMarshalled(cdc codec.JSONCodec) app.GenesisState
}

func (suite *HandlerTestSuite) SetupWithGenState(builders ...genesisBuilder) {
	suite.SetupApp()

	builtGenStates := []app.GenesisState{
		NewStakingGenesisState(suite.App.AppCodec()),
		NewPricefeedGenStateMultiFromTime(suite.App.AppCodec(), suite.genesisTime),
		// NewCDPGenStateMulti(suite.App.AppCodec()),
		NewJoltGenStateMulti(suite.genesisTime).BuildMarshalled(suite.App.AppCodec()),
	}
	for _, builder := range builders {
		builtGenStates = append(builtGenStates, builder.BuildMarshalled(suite.App.AppCodec()))
	}

	suite.App.InitializeFromGenesisStatesWithTime(
		suite.genesisTime,
		builtGenStates...,
	)
}

// authBuilder returns a new auth genesis builder with a full mint  module account.
func (suite *HandlerTestSuite) authBuilder() *app.AuthBankGenesisBuilder {
	return app.NewAuthBankGenesisBuilder().
		WithSimpleModuleAccount(minttypes.ModuleName, cs(c(types2.USDXMintingRewardDenom, 1e18), c("hard", 1e18), c("swap", 1e18)))
}

// incentiveBuilder returns a new incentive genesis builder with a genesis time and multipliers set
func (suite *HandlerTestSuite) incentiveBuilder() testutil2.IncentiveGenesisBuilder {
	return testutil2.NewIncentiveGenesisBuilder().
		WithGenesisTime(suite.genesisTime).
		WithMultipliers(types2.MultipliersPerDenoms{
			{
				Denom: "hard",
				Multipliers: types2.Multipliers{
					types2.NewMultiplier("small", 0, d("0.2")),
					types2.NewMultiplier("large", 12, d("1.0")),
				},
			},
			{
				Denom: "swap",
				Multipliers: types2.Multipliers{
					types2.NewMultiplier("small", 0, d("0.2")),
					types2.NewMultiplier("large", 12, d("1.0")),
				},
			},
			{
				Denom: "ujolt",
				Multipliers: types2.Multipliers{
					types2.NewMultiplier("small", 0, d("0.2")),
					types2.NewMultiplier("large", 12, d("1.0")),
				},
			},
		})
}

func (suite *HandlerTestSuite) TestPayoutJoltClaimMultiDenom() {
	userAddr, receiverAddr := suite.addrs[0], suite.addrs[1]

	authBulder := suite.authBuilder().
		WithSimpleAccount(userAddr, cs(c("bnb", 1e12))).
		WithSimpleAccount(receiverAddr, nil).
		WithSimpleModuleAccount(minttypes.ModuleName, cs(), "minter").
		WithSimpleModuleAccount(types2.ModuleName, cs(), "minter")

	incentBuilder := suite.incentiveBuilder().
		WithSimpleSupplyRewardPeriod("bnb", cs(c("hard", 1e6), c("swap", 1e6))).
		WithSimpleBorrowRewardPeriod("bnb", cs(c("hard", 1e6), c("swap", 1e6)))

	suite.SetupWithGenState(authBulder, incentBuilder)

	suite.App.GetBankKeeper().MintCoins(suite.Ctx, types2.ModuleName, cs(c("hard", 1e12), c("swap", 1e12)))
	// create a deposit and borrow
	suite.NoError(suite.DeliverJoltMsgDeposit(userAddr, cs(c("bnb", 1e11))))
	suite.NoError(suite.DeliverJoltMsgBorrow(userAddr, cs(c("bnb", 1e10))))

	// accumulate some rewards
	suite.NextBlockAfter(7 * time.Second)

	preClaimBal := suite.GetBalance(userAddr)

	msg := types2.NewMsgClaimJoltReward(
		userAddr.String(),
		types2.Selections{
			types2.NewSelection("hard", "small"),
			types2.NewSelection("swap", "small"),
		},
	)

	// Claim denoms
	err := suite.DeliverIncentiveMsg(&msg)
	suite.NoError(err)

	// Check rewards were paid out
	expectedRewardsJolt := c("hard", int64(0.2*float64(2*7*1e6)))
	expectedRewardsSwap := c("swap", int64(0.2*float64(2*7*1e6)))
	suite.BalanceEquals(userAddr, preClaimBal.Add(expectedRewardsJolt, expectedRewardsSwap))
	suite.JoltRewardEquals(userAddr, nil)
}

func (suite *HandlerTestSuite) TestPayoutHardClaimSingleDenom() {
	userAddr := suite.addrs[0]

	authBulder := suite.authBuilder().
		WithSimpleAccount(userAddr, cs(c("bnb", 1e12))).
		WithSimpleModuleAccount(minttypes.ModuleName, cs(), "minter").
		WithSimpleModuleAccount(types2.ModuleName, cs(), "minter")

	incentBuilder := suite.incentiveBuilder().
		WithSimpleSupplyRewardPeriod("bnb", cs(c("jolt", 1e6), c("swap", 1e6))).
		WithSimpleBorrowRewardPeriod("bnb", cs(c("jolt", 1e6), c("swap", 1e6)))

	suite.SetupWithGenState(authBulder, incentBuilder)

	suite.App.GetBankKeeper().MintCoins(suite.Ctx, types2.ModuleName, cs(c("hard", 1e12), c("swap", 1e12)))

	// create a deposit and borrow
	suite.NoError(suite.DeliverJoltMsgDeposit(userAddr, cs(c("bnb", 1e11))))
	suite.NoError(suite.DeliverJoltMsgBorrow(userAddr, cs(c("bnb", 1e10))))

	// accumulate some rewards
	suite.NextBlockAfter(7 * time.Second)

	preClaimBal := suite.GetBalance(userAddr)

	msg := types2.NewMsgClaimJoltReward(
		userAddr.String(),
		types2.Selections{
			types2.NewSelection("swap", "large"),
		},
	)

	// Claim rewards
	err := suite.DeliverIncentiveMsg(&msg)
	suite.NoError(err)

	// Check rewards were paid out
	expectedRewards := c("swap", 2*7*1e6)
	suite.BalanceEquals(userAddr, preClaimBal.Add(expectedRewards))

	// Check that claimed coins have been removed from a claim's reward
	suite.JoltRewardEquals(userAddr, cs(c("jolt", 2*7*1e6)))
}
