package incentive_test

import (
	"testing"
	"time"

	tmlog "github.com/tendermint/tendermint/libs/log"

	"github.com/joltify-finance/joltify_lending/x/third_party/incentive"
	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
	types3 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/joltify-finance/joltify_lending/app"
)

const (
	oneYear = 365 * 24 * time.Hour
)

type GenesisTestSuite struct {
	suite.Suite

	ctx    sdk.Context
	app    app.TestApp
	keeper keeper.Keeper
	addrs  []sdk.AccAddress

	genesisTime time.Time
}

func (suite *GenesisTestSuite) SetupTest() {
	tApp := app.NewTestApp(tmlog.TestingLogger(), suite.T().TempDir())
	suite.app = tApp
	keeper := tApp.GetIncentiveKeeper()
	suite.genesisTime = time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC)

	_, addrs := app.GeneratePrivKeyAddressPairs(5)

	authBuilder := app.NewAuthBankGenesisBuilder().
		WithSimpleAccount(addrs[0], cs(c("bnb", 1e10), c("ujolt", 1e10))).
		WithSimpleModuleAccount(types2.IncentiveMacc, cs(c("hard", 1e15), c("ujolt", 1e15)))

	loanToValue, _ := sdk.NewDecFromStr("0.6")
	borrowLimit := sdk.NewDec(1000000000000000)
	joltGS := types3.NewGenesisState(
		types3.NewParams(
			types3.MoneyMarkets{
				types3.NewMoneyMarket("ujolt", types3.NewBorrowLimit(false, borrowLimit, loanToValue), "jolt:usd", sdk.NewInt(1000000), types3.NewInterestRateModel(sdk.MustNewDecFromStr("0.05"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10")), sdk.MustNewDecFromStr("0.05"), sdk.ZeroDec()),
				types3.NewMoneyMarket("bnb", types3.NewBorrowLimit(false, borrowLimit, loanToValue), "bnb:usd", sdk.NewInt(1000000), types3.NewInterestRateModel(sdk.MustNewDecFromStr("0.05"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10")), sdk.MustNewDecFromStr("0.05"), sdk.ZeroDec()),
			},
			sdk.NewDec(10),
		),
		types3.DefaultAccumulationTimes,
		types3.DefaultDeposits,
		types3.DefaultBorrows,
		types3.DefaultTotalSupplied,
		types3.DefaultTotalBorrowed,
		types3.DefaultTotalReserves,
	)

	pa := types2.NewParams(
		types2.MultiRewardPeriods{types2.NewMultiRewardPeriod(true, "btcb/usdx", suite.genesisTime.Add(-1*oneYear), suite.genesisTime.Add(oneYear), cs(c("swp", 122354)))},
		types2.MultiRewardPeriods{types2.NewMultiRewardPeriod(true, "ujolt", suite.genesisTime.Add(-1*oneYear), suite.genesisTime.Add(oneYear), cs(c("hard", 122354)))},
		types2.MultipliersPerDenoms{
			{
				Denom: "ujolt",
				Multipliers: types2.Multipliers{
					types2.NewMultiplier("large", 12, d("1.0")),
				},
			},
			{
				Denom: "hard",
				Multipliers: types2.Multipliers{
					types2.NewMultiplier("small", 1, d("0.25")),
					types2.NewMultiplier("large", 12, d("1.0")),
				},
			},
			{
				Denom: "swp",
				Multipliers: types2.Multipliers{
					types2.NewMultiplier("small", 1, d("0.25")),
					types2.NewMultiplier("medium", 6, d("0.8")),
				},
			},
		},
		suite.genesisTime.Add(5*oneYear),
	)

	incentiveGS := types2.NewGenesisState(
		pa,
		types2.DefaultGenesisRewardState,
		types2.DefaultGenesisRewardState,
		types2.DefaultJoltClaims,
	)

	cdc := suite.app.AppCodec()

	tApp.InitializeFromGenesisStatesWithTime(
		suite.genesisTime, nil, nil,
		app.GenesisState{types2.ModuleName: cdc.MustMarshalJSON(&incentiveGS)},
		app.GenesisState{types3.ModuleName: cdc.MustMarshalJSON(&joltGS)},
		NewPricefeedGenStateMultiFromTime(cdc, suite.genesisTime),
		authBuilder.BuildMarshalled(cdc),
	)

	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: suite.genesisTime})

	suite.addrs = addrs
	suite.keeper = keeper
	suite.ctx = ctx
}

func (suite *GenesisTestSuite) TestExportedGenesisMatchesImported() {
	genesisTime := time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC)
	pa := types2.NewParams(
		types2.MultiRewardPeriods{types2.NewMultiRewardPeriod(true, "btcb/usdx", genesisTime.Add(-1*oneYear), genesisTime.Add(oneYear), cs(c("swp", 122354)))},
		types2.MultiRewardPeriods{types2.NewMultiRewardPeriod(true, "ujolt", genesisTime.Add(-1*oneYear), genesisTime.Add(oneYear), cs(c("hard", 122354)))},
		types2.MultipliersPerDenoms{
			{
				Denom: "ujolt",
				Multipliers: types2.Multipliers{
					types2.NewMultiplier("large", 12, d("1.0")),
				},
			},
			{
				Denom: "hard",
				Multipliers: types2.Multipliers{
					types2.NewMultiplier("small", 1, d("0.25")),
					types2.NewMultiplier("large", 12, d("1.0")),
				},
			},
			{
				Denom: "swp",
				Multipliers: types2.Multipliers{
					types2.NewMultiplier("small", 1, d("0.25")),
					types2.NewMultiplier("medium", 6, d("0.8")),
				},
			},
		},
		genesisTime.Add(5*oneYear),
	)
	genesisState := types2.NewGenesisState(
		pa,
		types2.NewGenesisRewardState(
			types2.AccumulationTimes{
				types2.NewAccumulationTime("bnb-a", genesisTime),
			},
			types2.MultiRewardIndexes{
				types2.NewMultiRewardIndex("bnb-a", types2.RewardIndexes{{CollateralType: "ujolt", RewardFactor: d("0.3")}}),
			},
		),
		types2.NewGenesisRewardState(
			types2.AccumulationTimes{
				types2.NewAccumulationTime("bnb", genesisTime.Add(-1*time.Hour)),
			},
			types2.MultiRewardIndexes{
				types2.NewMultiRewardIndex("bnb", types2.RewardIndexes{{CollateralType: "hard", RewardFactor: d("0.1")}}),
			},
		),

		types2.JoltLiquidityProviderClaims{
			types2.NewJoltLiquidityProviderClaim(
				suite.addrs[0],
				cs(c("ujolt", 1e9), c("hard", 1e9)),
				types2.MultiRewardIndexes{{CollateralType: "bnb", RewardIndexes: types2.RewardIndexes{{CollateralType: "hard", RewardFactor: d("0.01")}}}},
				types2.MultiRewardIndexes{{CollateralType: "bnb", RewardIndexes: types2.RewardIndexes{{CollateralType: "hard", RewardFactor: d("0.0")}}}},
			),
			types2.NewJoltLiquidityProviderClaim(
				suite.addrs[1],
				cs(c("hard", 1)),
				types2.MultiRewardIndexes{{CollateralType: "bnb", RewardIndexes: types2.RewardIndexes{{CollateralType: "hard", RewardFactor: d("0.1")}}}},
				types2.MultiRewardIndexes{{CollateralType: "bnb", RewardIndexes: types2.RewardIndexes{{CollateralType: "hard", RewardFactor: d("0.0")}}}},
			),
		},
	)

	tApp := app.NewTestApp(tmlog.TestingLogger(), suite.T().TempDir())
	ctx := tApp.NewContext(true, tmproto.Header{Height: 0, Time: genesisTime})

	// Incentive init genesis reads from the cdp keeper to check params are ok. So it needs to be initialized first.
	// Then the cdp keeper reads from pricefeed keeper to check its params are ok. So it also need initialization.
	tApp.InitializeFromGenesisStates(nil, nil,
		NewPricefeedGenStateMultiFromTime(tApp.AppCodec(), genesisTime),
	)

	incentive.InitGenesis(
		ctx,
		tApp.GetIncentiveKeeper(),
		tApp.GetAccountKeeper(),
		genesisState,
	)

	exportedGenesisState := incentive.ExportGenesis(ctx, tApp.GetIncentiveKeeper())

	suite.Equal(genesisState, exportedGenesisState)
}

func (suite *GenesisTestSuite) TestInitGenesisPanicsWhenAccumulationTimesToLongAgo() {
	genesisTime := time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC)
	invalidRewardState := types2.NewGenesisRewardState(
		types2.AccumulationTimes{
			types2.NewAccumulationTime(
				"bnb",
				genesisTime.Add(-23*incentive.EarliestValidAccumulationTime).Add(-time.Nanosecond),
			),
		},
		types2.MultiRewardIndexes{},
	)
	minimalParams := types2.Params{
		ClaimEnd: genesisTime.Add(5 * oneYear),
	}

	testCases := []struct {
		genesisState types2.GenesisState
	}{
		{
			types2.GenesisState{
				Params:                minimalParams,
				JoltSupplyRewardState: invalidRewardState,
			},
		},
		{
			types2.GenesisState{
				Params:                minimalParams,
				JoltBorrowRewardState: invalidRewardState,
			},
		},
	}

	for _, tc := range testCases {
		tApp := app.NewTestApp(tmlog.TestingLogger(), suite.T().TempDir())
		ctx := tApp.NewContext(true, tmproto.Header{Height: 0, Time: genesisTime})

		// Incentive init genesis reads from the cdp keeper to check params are ok. So it needs to be initialized first.
		// Then the cdp keeper reads from pricefeed keeper to check its params are ok. So it also need initialization.
		tApp.InitializeFromGenesisStates(nil, nil,
			NewPricefeedGenStateMultiFromTime(tApp.AppCodec(), genesisTime),
		)

		suite.PanicsWithValue(
			"found accumulation time '1975-01-06 23:59:59.999999999 +0000 UTC' more than '8760h0m0s' behind genesis time '1998-01-01 00:00:00 +0000 UTC'",
			func() {
				incentive.InitGenesis(
					ctx, tApp.GetIncentiveKeeper(),
					tApp.GetAccountKeeper(),
					tc.genesisState,
				)
			},
		)
	}
}

func (suite *GenesisTestSuite) TestValidateAccumulationTime() {
	genTime := time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC)

	err := incentive.ValidateAccumulationTime(
		genTime.Add(-incentive.EarliestValidAccumulationTime).Add(-time.Nanosecond),
		genTime,
	)
	suite.Error(err)
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}
