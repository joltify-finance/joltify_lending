package keeper_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"
	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/joltify-finance/joltify_lending/app"
	appconfig "github.com/joltify-finance/joltify_lending/app/config"
	testapp "github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/app"
	"github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/constants"
	minttypes "github.com/joltify-finance/joltify_lending/x/mint/types"
	testutil2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/testutil"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
	jolttypes "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"
	pricefeedtypes "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"
	prices "github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/types"
	"github.com/stretchr/testify/suite"
)

// Test suite used for all keeper tests
type HandlerTestSuite struct {
	testutil2.IntegrationTester

	genesisTime time.Time
	addrs       []sdk.AccAddress
}

func TestHandlerTestSuite(t *testing.T) {
	t.SkipNow()
	suite.Run(t, new(HandlerTestSuite))
}

// SetupTest is run automatically before each suite test
func (suite *HandlerTestSuite) SetupTest() {
	appconfig.SetupConfig()

	_, suite.addrs = app.GeneratePrivKeyAddressPairs(5)

	suite.genesisTime = time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC)
}

func (suite *HandlerTestSuite) SetupApp() {
	suite.App = app.NewTestApp(log.NewTestLogger(suite.T()), suite.T().TempDir())
	suite.Ctx = suite.App.NewContext(true)
}

type genesisBuilder interface {
	BuildMarshalled(cdc codec.JSONCodec) app.GenesisState
}

func (suite *HandlerTestSuite) SetupWithGenState(genAcc []authtypes.GenesisAccount, coins sdk.Coins, builders ...genesisBuilder) {
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

	suite.App.InitializeFromGenesisStatesWithTime(suite.T(),
		suite.genesisTime, genAcc, coins,
		builtGenStates...,
	)
}

// authBuilder returns a new auth genesis builder with a full mint  module account.
func (suite *HandlerTestSuite) authBuilder(authgen *authtypes.GenesisState, bankgen *banktypes.GenesisState) *app.AuthBankGenesisBuilder {
	return app.NewAuthBankGenesisBuilder(authgen, bankgen).
		WithSimpleModuleAccount(minttypes.ModuleName, cs(c(types2.RewardDenom, 1e18), c("hard", 1e18), c("swap", 1e18)))
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

	authBuilder := suite.authBuilder(nil, nil).
		WithSimpleAccount(userAddr, cs(c("bnb", 1e12))).
		WithSimpleAccount(receiverAddr, nil).
		WithSimpleModuleAccount(minttypes.ModuleName, cs(), "minter").
		WithSimpleModuleAccount(types2.ModuleName, cs(), "minter")

	incentBuilder := suite.incentiveBuilder().
		WithSimpleSupplyRewardPeriod("bnb", cs(c("hard", 1e6), c("swap", 1e6))).
		WithSimpleBorrowRewardPeriod("bnb", cs(c("hard", 1e6), c("swap", 1e6)))

	var genAcc []authtypes.GenesisAccount
	b := authtypes.NewBaseAccount(userAddr, nil, 0, 0)
	genAcc = append(genAcc, b)
	coin := cs(c("bnb", 1e12))

	suite.SetupApp()

	mapp := suite.App.InitializeFromGenesisStatesWithTime(suite.T(),
		suite.genesisTime, genAcc, coin,
		authBuilder.BuildMarshalled(suite.App.AppCodec()),
		NewPricefeedGenStateMultiFromTime(suite.App.AppCodec(), suite.genesisTime),
		NewJoltGenStateMulti(suite.genesisTime).BuildMarshalled(suite.App.AppCodec()),
		incentBuilder.BuildMarshalled(suite.App.AppCodec()),
	)
	suite.App = mapp
	suite.App.App = mapp.App
	suite.Ctx = mapp.Ctx
	suite.App.Ctx = mapp.Ctx

	err := suite.App.GetBankKeeper().MintCoins(suite.Ctx, types2.ModuleName, cs(c("hard", 1e12), c("swap", 1e12)))
	suite.Require().NoError(err)
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
	err = suite.DeliverIncentiveMsg(&msg)
	suite.NoError(err)

	// Check rewards were paid out
	expectedRewardsJolt := c("hard", int64(0.2*float64(2*7*1e6)))
	expectedRewardsSwap := c("swap", int64(0.2*float64(2*7*1e6)))
	suite.BalanceEquals(userAddr, preClaimBal.Add(expectedRewardsJolt, expectedRewardsSwap))
	suite.JoltRewardEquals(userAddr, nil)
}

func (suite *HandlerTestSuite) TestPayoutHardClaimSingleDenom() {
	cfg := appconfig.MakeEncodingConfig()
	app.ModuleBasics.RegisterInterfaces(cfg.InterfaceRegistry)
	userAddr := suite.addrs[0]

	incentBuilder := suite.incentiveBuilder().
		WithSimpleSupplyRewardPeriod("bnb", cs(c("jolt", 1e6), c("swap", 1e6))).
		WithSimpleBorrowRewardPeriod("bnb", cs(c("jolt", 1e6), c("swap", 1e6)))

	tApp := testapp.NewTestAppBuilder(suite.T()).WithGenesisDocFn(func() (genesis tmtypes.GenesisDoc) {
		genesis = testapp.DefaultGenesis()
		genesis.GenesisTime = time.Now()
		testapp.UpdateGenesisDocWithAppStateForModule(
			&genesis,
			func(genesisState *prices.GenesisState) {
				*genesisState = constants.TestPricesGenesisState
			},
		)

		var appState map[string]json.RawMessage
		err := json.Unmarshal(genesis.AppState, &appState)
		suite.NoError(err)

		var authgen authtypes.GenesisState
		constants.TestEncodingCfg.Codec.MustUnmarshalJSON(appState[authtypes.ModuleName], &authgen)
		suite.NoError(err)

		var bankgen banktypes.GenesisState
		constants.TestEncodingCfg.Codec.MustUnmarshalJSON(appState[banktypes.ModuleName], &bankgen)
		suite.NoError(err)

		authBuilder := suite.authBuilder(&authgen, &bankgen).
			WithSimpleAccount(userAddr, cs(c("bnb", 1e12))).
			WithSimpleModuleAccount(minttypes.ModuleName, cs(), "minter").
			WithSimpleModuleAccount(types2.ModuleName, cs(), "minter")

		testapp.UpdateGenesisDocWithAppStateForModule(
			&genesis,
			func(genesisState *authtypes.GenesisState) {
				*genesisState = authBuilder.AuthGenesis
			},
		)
		testapp.UpdateGenesisDocWithAppStateForModule(
			&genesis,
			func(genesisState *banktypes.GenesisState) {
				// authgenesis:=authBuilder.BuildMarshalled(suite.App.AppCodec())
				*genesisState = authBuilder.BankGenesis
			},
		)
		testapp.UpdateGenesisDocWithAppStateForModule(
			&genesis,
			func(genesisState *types2.GenesisState) {
				// authgenesis:=authBuilder.BuildMarshalled(suite.App.AppCodec())
				*genesisState = incentBuilder.GenesisState
			},
		)
		testapp.UpdateGenesisDocWithAppStateForModule(
			&genesis,
			func(genesisState *jolttypes.GenesisState) {
				// authgenesis:=authBuilder.BuildMarshalled(suite.App.AppCodec())
				*genesisState = jolttypes.DefaultGenesisState()
			},
		)

		testapp.UpdateGenesisDocWithAppStateForModule(
			&genesis,
			func(genesisState *pricefeedtypes.GenesisState) {
				_, addrs := app.GeneratePrivKeyAddressPairs(10)

				pfGenesis := pricefeedtypes.GenesisState{
					Params: pricefeedtypes.Params{
						Markets: []pricefeedtypes.Market{
							{MarketID: "btc:usd", BaseAsset: "btc", QuoteAsset: "usd", Oracles: addrs, Active: true},
							{MarketID: "xrp:usd", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: addrs, Active: true},
							{MarketID: "bnb:usd", BaseAsset: "bnb", QuoteAsset: "usd", Oracles: addrs, Active: true},
						},
					},
					PostedPrices: []pricefeedtypes.PostedPrice{
						{
							MarketID:      "btc:usd",
							OracleAddress: addrs[0],
							Price:         sdkmath.LegacyMustNewDecFromStr("8000.00"),
							Expiry:        time.Now().Add(1 * time.Hour),
						},
						{
							MarketID:      "xrp:usd",
							OracleAddress: addrs[0],
							Price:         sdkmath.LegacyMustNewDecFromStr("0.25"),
							Expiry:        time.Now().Add(1 * time.Hour),
						},
						{
							MarketID:      "bnb:usd",
							OracleAddress: addrs[0],
							Price:         sdkmath.LegacyMustNewDecFromStr("100000000000"),
							Expiry:        time.Now().Add(1 * time.Hour),
						},
					},
				}
				*genesisState = pfGenesis
			},
		)

		return genesis
	}).Build()

	ctx := tApp.InitChain()

	// var genAcc []authtypes.GenesisAccount
	// b := authtypes.NewBaseAccount(userAddr, nil, 0, 0)
	// genAcc = append(genAcc, b)
	// coin := cs(c("bnb", 1e12))

	// suite.SetupApp()

	//mapp := suite.App.InitializeFromGenesisStatesWithTime(suite.T(),
	//	suite.genesisTime, genAcc, coin,
	//	authBuilder.BuildMarshalled(suite.App.AppCodec()),
	//	NewPricefeedGenStateMultiFromTime(suite.App.AppCodec(), suite.genesisTime),
	//	NewJoltGenStateMulti(suite.genesisTime).BuildMarshalled(suite.App.AppCodec()),
	//	incentBuilder.BuildMarshalled(suite.App.AppCodec()),
	//)
	//suite.App.App = tApp
	//suite.App.App = mapp.App
	//suite.Ctx = mapp.Ctx
	//suite.App.Ctx = mapp.Ctx

	suite.Ctx = ctx
	suite.App.Ctx = ctx
	suite.App.App = *tApp.App
	//
	//loanToValue, _ := sdkmath.LegacyNewDecFromStr("0.8")
	//m2 := jolttypes.NewMoneyMarket(
	//	"abnb",
	//	jolttypes.NewBorrowLimit(
	//		false,
	//		sdkmath.LegacyNewDec(1e15),
	//		loanToValue,
	//	),
	//	"bnb:usd",
	//	sdkmath.NewInt(1e18),
	//	jolttypes.NewInterestRateModel(
	//		sdkmath.LegacyMustNewDecFromStr("0.0"),
	//		sdkmath.LegacyMustNewDecFromStr("0.02"),
	//		sdkmath.LegacyMustNewDecFromStr("0.8"),
	//		sdkmath.LegacyMustNewDecFromStr("5"),
	//	),
	//	sdkmath.LegacyMustNewDecFromStr("0.02"),
	//	sdkmath.LegacyMustNewDecFromStr("0.02"),
	//)
	//
	//suite.App.GetJoltKeeper().SetMoneyMarket(suite.Ctx, "bnb", m2)
	err := suite.App.GetBankKeeper().MintCoins(suite.Ctx, types2.ModuleName, cs(c("hard", 1e12), c("swap", 1e12)))
	// err := tApp.App.BankKeeper.MintCoins(suite.Ctx, types2.ModuleName, cs(c("hard", 1e12), c("swap", 1e12)))
	suite.Require().NoError(err)

	// err := fundModuleAccount(suite.App.GetBankKeeper(), suite.Ctx, types2.ModuleName, cs(c("jjolt", 1e18)))
	// suite.Require().NoError(err)

	// create a deposit and borrow
	suite.NoError(suite.DeliverJoltMsgDeposit(userAddr, cs(c("bnb", 1e11))))
	suite.NoError(suite.DeliverJoltMsgBorrow(userAddr, cs(c("bnb", 1e10))))

	// de, found := suite.App.GetJoltKeeper().GetDeposit(suite.Ctx, userAddr)

	fmt.Printf("current height %v\n", suite.App.Ctx.BlockHeight())
	// accumulate some rewards
	newctx := tApp.AdvanceToBlock(2, testapp.AdvanceToBlockOptions{BlockTime: suite.App.Ctx.BlockTime().Add(time.Second * 7)})
	suite.Ctx = newctx
	suite.App.Ctx = newctx
	suite.App.App = *tApp.App

	preClaimBal := suite.GetBalance(userAddr)

	msg := types2.NewMsgClaimJoltReward(
		userAddr.String(),
		types2.Selections{
			types2.NewSelection("swap", "large"),
		},
	)

	// Claim rewards
	err = suite.DeliverIncentiveMsg(&msg)
	suite.NoError(err)

	// Check rewards were paid out
	expectedRewards := c("swap", 2*7*1e6)
	suite.BalanceEquals(userAddr, preClaimBal.Add(expectedRewards))

	// Check that claimed coins have been removed from a claim's reward
	suite.JoltRewardEquals(userAddr, cs(c("jolt", 2*7*1e6)))
}
