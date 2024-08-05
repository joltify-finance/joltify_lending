package incentive_test

import (
	"testing"
	"time"

	tmlog "cosmossdk.io/log"

	"github.com/joltify-finance/joltify_lending/x/third_party/incentive"
	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/joltify-finance/joltify_lending/app"
)

const (
	oneYear = 365 * 24 * time.Hour
)

type GenesisTestSuite struct {
	suite.Suite

	ctx    context.Context
	app    app.TestApp
	keeper keeper.Keeper
	addrs  []sdk.AccAddress

	genesisTime time.Time
}

func (suite *GenesisTestSuite) SetupTest() {
	tApp := app.NewTestApp(tmlog.TestingLogger(), suite.T().TempDir())
	suite.app = tApp
	k := tApp.GetIncentiveKeeper()
	suite.genesisTime = time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC)

	_, addrs := app.GeneratePrivKeyAddressPairs(5)

	authBuilder := app.NewAuthBankGenesisBuilder().
		WithSimpleAccount(addrs[0], cs(c("bnb", 1e10), c("ujolt", 1e10))).
		WithSimpleModuleAccount(types.IncentiveMacc, cs(c("hard", 1e15), c("ujolt", 1e15)))

	loanToValue, _ := sdkmath.LegacyNewDecFromStr("0.6")
	borrowLimit := sdkmath.LegacyNewDec(1000000000000000)
	joltGS := types2.NewGenesisState(
		types2.NewParams(
			types2.MoneyMarkets{
				types2.NewMoneyMarket("ujolt", types2.NewBorrowLimit(false, borrowLimit, loanToValue), "jolt:usd", sdkmath.NewInt(1000000), types2.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyMustNewDecFromStr("2"), sdkmath.LegacyMustNewDecFromStr("0.8"), sdkmath.LegacyMustNewDecFromStr("10")), sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyZeroDec()),
				types2.NewMoneyMarket("bnb", types2.NewBorrowLimit(false, borrowLimit, loanToValue), "bnb:usd", sdkmath.NewInt(1000000), types2.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyMustNewDecFromStr("2"), sdkmath.LegacyMustNewDecFromStr("0.8"), sdkmath.LegacyMustNewDecFromStr("10")), sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyZeroDec()),
			},
			sdkmath.LegacyNewDec(10),
		),
		types2.DefaultAccumulationTimes,
		types2.DefaultDeposits,
		types2.DefaultBorrows,
		types2.DefaultTotalSupplied,
		types2.DefaultTotalBorrowed,
		types2.DefaultTotalReserves,
	)

	pa := types.NewParams(
		types.MultiRewardPeriods{types.NewMultiRewardPeriod(true, "btcb/usdx", suite.genesisTime.Add(-1*oneYear), suite.genesisTime.Add(oneYear), cs(c("swp", 122354)))},
		types.MultiRewardPeriods{types.NewMultiRewardPeriod(true, "ujolt", suite.genesisTime.Add(-1*oneYear), suite.genesisTime.Add(oneYear), cs(c("hard", 122354)))},
		types.MultiRewardPeriods{types.NewMultiRewardPeriod(true, "ujolt", suite.genesisTime.Add(-1*oneYear), suite.genesisTime.Add(oneYear), cs(c("hard", 122354)))},
		types.MultiRewardPeriods{types.NewMultiRewardPeriod(true, "spv:1", suite.genesisTime.Add(-1*oneYear), suite.genesisTime.Add(oneYear), cs(c("jolt", 122354)))},
		types.MultipliersPerDenoms{
			{
				Denom: "ujolt",
				Multipliers: types.Multipliers{
					types.NewMultiplier("large", 12, d("1.0")),
				},
			},
			{
				Denom: "hard",
				Multipliers: types.Multipliers{
					types.NewMultiplier("small", 1, d("0.25")),
					types.NewMultiplier("large", 12, d("1.0")),
				},
			},
			{
				Denom: "swp",
				Multipliers: types.Multipliers{
					types.NewMultiplier("small", 1, d("0.25")),
					types.NewMultiplier("medium", 6, d("0.8")),
				},
			},
		},
		suite.genesisTime.Add(5*oneYear),
	)

	incentiveGS := types.NewGenesisState(
		pa,
		types.DefaultGenesisRewardState,
		types.DefaultGenesisRewardState,
		types.DefaultGenesisRewardState,
		types.DefaultSPVGenesisRewardState,
		types.DefaultJoltClaims,
		types.DefaultSwapClaims,
	)

	cdc := suite.app.AppCodec()

	tApp.InitializeFromGenesisStatesWithTime(
		suite.genesisTime, nil, nil,
		app.GenesisState{types.ModuleName: cdc.MustMarshalJSON(&incentiveGS)},
		app.GenesisState{types2.ModuleName: cdc.MustMarshalJSON(&joltGS)},
		NewPricefeedGenStateMultiFromTime(cdc, suite.genesisTime),
		authBuilder.BuildMarshalled(cdc),
	)

	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: suite.genesisTime})

	suite.addrs = addrs
	suite.keeper = k
	suite.ctx = ctx
}

func (suite *GenesisTestSuite) TestExportedGenesisMatchesImported() {
	a := types.SPVRewardAccTokens{PaymentAmount: cs(c("jolt", 1e9))}
	genesisTime := time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC)
	pa := types.NewParams(
		types.MultiRewardPeriods{types.NewMultiRewardPeriod(true, "ujolt", genesisTime.Add(-1*oneYear), genesisTime.Add(oneYear), cs(c("hard", 122354)))},
		types.MultiRewardPeriods{types.NewMultiRewardPeriod(true, "ujolt", genesisTime.Add(-1*oneYear), genesisTime.Add(oneYear), cs(c("hard", 122354)))},
		types.MultiRewardPeriods{types.NewMultiRewardPeriod(true, "btcb/usdx", genesisTime.Add(-1*oneYear), genesisTime.Add(oneYear), cs(c("swp", 122354)))},
		types.MultiRewardPeriods{types.NewMultiRewardPeriod(true, "spv:1", genesisTime.Add(-1*oneYear), genesisTime.Add(oneYear), cs(c("jolt", 122354)))},
		types.MultipliersPerDenoms{
			{
				Denom: "ujolt",
				Multipliers: types.Multipliers{
					types.NewMultiplier("large", 12, d("1.0")),
				},
			},
			{
				Denom: "hard",
				Multipliers: types.Multipliers{
					types.NewMultiplier("small", 1, d("0.25")),
					types.NewMultiplier("large", 12, d("1.0")),
				},
			},
			{
				Denom: "swp",
				Multipliers: types.Multipliers{
					types.NewMultiplier("small", 1, d("0.25")),
					types.NewMultiplier("medium", 6, d("0.8")),
				},
			},
		},
		genesisTime.Add(5*oneYear),
	)
	genesisState := types.NewGenesisState(
		pa,
		types.NewGenesisRewardState(
			types.AccumulationTimes{
				types.NewAccumulationTime("bnb-a", genesisTime),
			},
			types.MultiRewardIndexes{
				types.NewMultiRewardIndex("bnb-a", types.RewardIndexes{{CollateralType: "ujolt", RewardFactor: d("0.3")}}),
			},
		),
		types.NewGenesisRewardState(
			types.AccumulationTimes{
				types.NewAccumulationTime("bnb", genesisTime.Add(-1*time.Hour)),
			},
			types.MultiRewardIndexes{
				types.NewMultiRewardIndex("bnb", types.RewardIndexes{{CollateralType: "hard", RewardFactor: d("0.1")}}),
			},
		),
		types.NewGenesisRewardState(
			types.AccumulationTimes{
				types.NewAccumulationTime("bnb", genesisTime.Add(-1*time.Hour)),
			},
			types.MultiRewardIndexes{
				types.NewMultiRewardIndex("bnb", types.RewardIndexes{{CollateralType: "hard", RewardFactor: d("0.1")}}),
			},
		),

		types.NewSPVGenesisRewardState(
			types.AccumulationTimes{
				types.NewAccumulationTime("bnb", genesisTime.Add(-1*time.Hour)),
			},
			[]types.SPVRewardAccIndex{{
				CollateralType: "bnb",
				AccReward:      a,
			}},
			[]*types.SPVGenRewardInvestorState{{
				Pool:   "spv:1",
				Wallet: suite.addrs[0].String(),
				Reward: sdk.NewCoins(sdk.NewCoin("jolt", sdkmath.NewInt(1e9))),
			}},
		),

		types.JoltLiquidityProviderClaims{
			types.NewJoltLiquidityProviderClaim(
				suite.addrs[0],
				cs(c("ujolt", 1e9), c("hard", 1e9)),
				types.MultiRewardIndexes{{CollateralType: "bnb", RewardIndexes: types.RewardIndexes{{CollateralType: "hard", RewardFactor: d("0.01")}}}},
				types.MultiRewardIndexes{{CollateralType: "bnb", RewardIndexes: types.RewardIndexes{{CollateralType: "hard", RewardFactor: d("0.0")}}}},
			),
			types.NewJoltLiquidityProviderClaim(
				suite.addrs[1],
				cs(c("ujolt", 1)),
				types.MultiRewardIndexes{{CollateralType: "bnb", RewardIndexes: types.RewardIndexes{{CollateralType: "hard", RewardFactor: d("0.1")}}}},
				types.MultiRewardIndexes{{CollateralType: "bnb", RewardIndexes: types.RewardIndexes{{CollateralType: "hard", RewardFactor: d("0.0")}}}},
			),
		},

		types.SwapClaims{
			types.NewSwapClaim(
				suite.addrs[3],
				nil,
				types.MultiRewardIndexes{{CollateralType: "btcb/usdx", RewardIndexes: types.RewardIndexes{{CollateralType: "swap", RewardFactor: d("0.0")}}}},
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
	suite.True(exportedGenesisState.String() == genesisState.String())
	// suite.Equal(genesisState.String(), exportedGenesisState.String())
}

func (suite *GenesisTestSuite) TestValidateAccumulationTime() {
	// valid when set
	accTime := time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC)
	suite.NoError(incentive.ValidateAccumulationTime(accTime))

	// invalid when nil value
	suite.Error(incentive.ValidateAccumulationTime(time.Time{}))
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}
