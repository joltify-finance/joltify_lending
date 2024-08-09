package keeper_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/x/bank/testutil"

	sdkmath "cosmossdk.io/math"
	"github.com/joltify-finance/joltify_lending/x/third_party/jolt"
	"github.com/joltify-finance/joltify_lending/x/third_party/jolt/keeper"
	types3 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"

	"github.com/cometbft/cometbft/crypto"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/joltify-finance/joltify_lending/app"
)

type InterestTestSuite struct {
	suite.Suite
}

func (suite *InterestTestSuite) TestCalculateUtilizationRatio() {
	type args struct {
		cash          sdkmath.LegacyDec
		borrows       sdkmath.LegacyDec
		reserves      sdkmath.LegacyDec
		expectedValue sdkmath.LegacyDec
	}

	type test struct {
		name string
		args args
	}

	testCases := []test{
		{
			"normal",
			args{
				cash:          sdkmath.LegacyMustNewDecFromStr("1000"),
				borrows:       sdkmath.LegacyMustNewDecFromStr("5000"),
				reserves:      sdkmath.LegacyMustNewDecFromStr("100"),
				expectedValue: sdkmath.LegacyMustNewDecFromStr("0.847457627118644068"),
			},
		},
		{
			"high util ratio",
			args{
				cash:          sdkmath.LegacyMustNewDecFromStr("1000"),
				borrows:       sdkmath.LegacyMustNewDecFromStr("250000"),
				reserves:      sdkmath.LegacyMustNewDecFromStr("100"),
				expectedValue: sdkmath.LegacyMustNewDecFromStr("0.996412913511359107"),
			},
		},
		{
			"very high util ratio",
			args{
				cash:          sdkmath.LegacyMustNewDecFromStr("1000"),
				borrows:       sdkmath.LegacyMustNewDecFromStr("250000000000"),
				reserves:      sdkmath.LegacyMustNewDecFromStr("100"),
				expectedValue: sdkmath.LegacyMustNewDecFromStr("0.999999996400000013"),
			},
		},
		{
			"low util ratio",
			args{
				cash:          sdkmath.LegacyMustNewDecFromStr("1000"),
				borrows:       sdkmath.LegacyMustNewDecFromStr("50"),
				reserves:      sdkmath.LegacyMustNewDecFromStr("100"),
				expectedValue: sdkmath.LegacyMustNewDecFromStr("0.052631578947368421"),
			},
		},
		{
			"very low util ratio",
			args{
				cash:          sdkmath.LegacyMustNewDecFromStr("10000000"),
				borrows:       sdkmath.LegacyMustNewDecFromStr("50"),
				reserves:      sdkmath.LegacyMustNewDecFromStr("100"),
				expectedValue: sdkmath.LegacyMustNewDecFromStr("0.000005000025000125"),
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			utilRatio := keeper.CalculateUtilizationRatio(tc.args.cash, tc.args.borrows, tc.args.reserves)
			suite.Require().Equal(tc.args.expectedValue, utilRatio)
		})
	}
}

func (suite *InterestTestSuite) TestCalculateBorrowRate() {
	type args struct {
		cash          sdkmath.LegacyDec
		borrows       sdkmath.LegacyDec
		reserves      sdkmath.LegacyDec
		model         types3.InterestRateModel
		expectedValue sdkmath.LegacyDec
	}

	type test struct {
		name string
		args args
	}

	// Normal model has:
	// 	- BaseRateAPY:      0.0
	// 	- BaseMultiplier:   0.1
	// 	- Kink:             0.8
	// 	- JumpMultiplier:   0.5
	normalModel := types3.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("0"), sdkmath.LegacyMustNewDecFromStr("0.1"), sdkmath.LegacyMustNewDecFromStr("0.8"), sdkmath.LegacyMustNewDecFromStr("0.5"))

	testCases := []test{
		{
			"normal no jump",
			args{
				cash:          sdkmath.LegacyMustNewDecFromStr("5000"),
				borrows:       sdkmath.LegacyMustNewDecFromStr("1000"),
				reserves:      sdkmath.LegacyMustNewDecFromStr("1000"),
				model:         normalModel,
				expectedValue: sdkmath.LegacyMustNewDecFromStr("0.020000000000000000"),
			},
		},
		{
			"normal with jump",
			args{
				cash:          sdkmath.LegacyMustNewDecFromStr("1000"),
				borrows:       sdkmath.LegacyMustNewDecFromStr("5000"),
				reserves:      sdkmath.LegacyMustNewDecFromStr("100"),
				model:         normalModel,
				expectedValue: sdkmath.LegacyMustNewDecFromStr("0.103728813559322034"),
			},
		},
		{
			"high cash",
			args{
				cash:          sdkmath.LegacyMustNewDecFromStr("10000000"),
				borrows:       sdkmath.LegacyMustNewDecFromStr("5000"),
				reserves:      sdkmath.LegacyMustNewDecFromStr("100"),
				model:         normalModel,
				expectedValue: sdkmath.LegacyMustNewDecFromStr("0.000049975511999120"),
			},
		},
		{
			"high borrows",
			args{
				cash:          sdkmath.LegacyMustNewDecFromStr("1000"),
				borrows:       sdkmath.LegacyMustNewDecFromStr("5000000000000"),
				reserves:      sdkmath.LegacyMustNewDecFromStr("100"),
				model:         normalModel,
				expectedValue: sdkmath.LegacyMustNewDecFromStr("0.179999999910000000"),
			},
		},
		{
			"high reserves",
			args{
				cash:          sdkmath.LegacyMustNewDecFromStr("1000"),
				borrows:       sdkmath.LegacyMustNewDecFromStr("5000"),
				reserves:      sdkmath.LegacyMustNewDecFromStr("1000000000000"),
				model:         normalModel,
				expectedValue: sdkmath.LegacyMustNewDecFromStr("0.180000000000000000"),
			},
		},
		{
			"random numbers",
			args{
				cash:          sdkmath.LegacyMustNewDecFromStr("125"),
				borrows:       sdkmath.LegacyMustNewDecFromStr("11"),
				reserves:      sdkmath.LegacyMustNewDecFromStr("82"),
				model:         normalModel,
				expectedValue: sdkmath.LegacyMustNewDecFromStr("0.020370370370370370"),
			},
		},
		{
			"increased base multiplier",
			args{
				cash:          sdkmath.LegacyMustNewDecFromStr("1000"),
				borrows:       sdkmath.LegacyMustNewDecFromStr("5000"),
				reserves:      sdkmath.LegacyMustNewDecFromStr("100"),
				model:         types3.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("0"), sdkmath.LegacyMustNewDecFromStr("0.5"), sdkmath.LegacyMustNewDecFromStr("0.8"), sdkmath.LegacyMustNewDecFromStr("1.0")),
				expectedValue: sdkmath.LegacyMustNewDecFromStr("0.447457627118644068"),
			},
		},
		{
			"decreased kink",
			args{
				cash:          sdkmath.LegacyMustNewDecFromStr("1000"),
				borrows:       sdkmath.LegacyMustNewDecFromStr("5000"),
				reserves:      sdkmath.LegacyMustNewDecFromStr("100"),
				model:         types3.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("0"), sdkmath.LegacyMustNewDecFromStr("0.5"), sdkmath.LegacyMustNewDecFromStr("0.1"), sdkmath.LegacyMustNewDecFromStr("1.0")),
				expectedValue: sdkmath.LegacyMustNewDecFromStr("0.797457627118644068"),
			},
		},
		{
			"zero model returns zero",
			args{
				cash:     sdkmath.LegacyMustNewDecFromStr("1000"),
				borrows:  sdkmath.LegacyMustNewDecFromStr("5000"),
				reserves: sdkmath.LegacyMustNewDecFromStr("100"),
				model: types3.NewInterestRateModel(
					sdkmath.LegacyMustNewDecFromStr("0.0"),
					sdkmath.LegacyMustNewDecFromStr("0.0"),
					sdkmath.LegacyMustNewDecFromStr("0.8"),
					sdkmath.LegacyMustNewDecFromStr("0.0"),
				),
				expectedValue: sdkmath.LegacyMustNewDecFromStr("0.0"),
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			borrowRate, err := keeper.CalculateBorrowRate(tc.args.model, tc.args.cash, tc.args.borrows, tc.args.reserves)
			suite.Require().NoError(err)
			suite.Require().Equal(tc.args.expectedValue, borrowRate)
		})
	}
}

func (suite *InterestTestSuite) TestCalculateBorrowInterestFactor() {
	type args struct {
		perSecondInterestRate sdkmath.LegacyDec
		timeElapsed           sdkmath.Int
		expectedValue         sdkmath.LegacyDec
	}

	type test struct {
		name string
		args args
	}

	oneYearInSeconds := int64(31536000)

	testCases := []test{
		{
			"1 year",
			args{
				perSecondInterestRate: sdkmath.LegacyMustNewDecFromStr("1.000000005555"),
				timeElapsed:           sdkmath.NewInt(oneYearInSeconds),
				expectedValue:         sdkmath.LegacyMustNewDecFromStr("1.191463614477847370"),
			},
		},
		{
			"10 year",
			args{
				perSecondInterestRate: sdkmath.LegacyMustNewDecFromStr("1.000000005555"),
				timeElapsed:           sdkmath.NewInt(oneYearInSeconds * 10),
				expectedValue:         sdkmath.LegacyMustNewDecFromStr("5.765113233897391189"),
			},
		},
		{
			"1 month",
			args{
				perSecondInterestRate: sdkmath.LegacyMustNewDecFromStr("1.000000005555"),
				timeElapsed:           sdkmath.NewInt(oneYearInSeconds / 12),
				expectedValue:         sdkmath.LegacyMustNewDecFromStr("1.014705619075717373"),
			},
		},
		{
			"1 day",
			args{
				perSecondInterestRate: sdkmath.LegacyMustNewDecFromStr("1.000000005555"),
				timeElapsed:           sdkmath.NewInt(oneYearInSeconds / 365),
				expectedValue:         sdkmath.LegacyMustNewDecFromStr("1.000480067194057924"),
			},
		},
		{
			"1 year: low interest rate",
			args{
				perSecondInterestRate: sdkmath.LegacyMustNewDecFromStr("1.000000000555"),
				timeElapsed:           sdkmath.NewInt(oneYearInSeconds),
				expectedValue:         sdkmath.LegacyMustNewDecFromStr("1.017656545925063632"),
			},
		},
		{
			"1 year, lower interest rate",
			args{
				perSecondInterestRate: sdkmath.LegacyMustNewDecFromStr("1.000000000055"),
				timeElapsed:           sdkmath.NewInt(oneYearInSeconds),
				expectedValue:         sdkmath.LegacyMustNewDecFromStr("1.001735985079841390"),
			},
		},
		{
			"1 year, lowest interest rate",
			args{
				perSecondInterestRate: sdkmath.LegacyMustNewDecFromStr("1.000000000005"),
				timeElapsed:           sdkmath.NewInt(oneYearInSeconds),
				expectedValue:         sdkmath.LegacyMustNewDecFromStr("1.000157692432076670"),
			},
		},
		{
			"1 year: high interest rate",
			args{
				perSecondInterestRate: sdkmath.LegacyMustNewDecFromStr("1.000000055555"),
				timeElapsed:           sdkmath.NewInt(oneYearInSeconds),
				expectedValue:         sdkmath.LegacyMustNewDecFromStr("5.766022095987868825"),
			},
		},
		{
			"1 year: higher interest rate",
			args{
				perSecondInterestRate: sdkmath.LegacyMustNewDecFromStr("1.000000555555"),
				timeElapsed:           sdkmath.NewInt(oneYearInSeconds),
				expectedValue:         sdkmath.LegacyMustNewDecFromStr("40628388.864535408465693310"),
			},
		},
		{
			"1 year: highest interest rate",
			args{
				perSecondInterestRate: sdkmath.LegacyMustNewDecFromStr("1.000001555555"),
				timeElapsed:           sdkmath.NewInt(oneYearInSeconds),
				expectedValue:         sdkmath.LegacyMustNewDecFromStr("2017093013158200407564.613502861572552603"),
			},
		},
		{
			"largest per second interest rate with practical elapsed time",
			args{
				perSecondInterestRate: sdkmath.LegacyMustNewDecFromStr("18.445"), // Begins to panic at ~18.45 (1845%/second interest rate)
				timeElapsed:           sdkmath.NewInt(30),                        // Assume a 30 second period, longer than any expected individual block
				expectedValue:         sdkmath.LegacyMustNewDecFromStr("94702138679846565921082258202543002089.215969366091911769"),
			},
		},
		{
			"supports calculated values greater than 1.84x10^19",
			args{
				perSecondInterestRate: sdkmath.LegacyMustNewDecFromStr("18.5"), // Old uint64 conversion would panic at ~18.45 (1845%/second interest rate)
				timeElapsed:           sdkmath.NewInt(30),                      // Assume a 30 second period, longer than any expected individual block
				expectedValue:         sdkmath.LegacyMustNewDecFromStr("103550416986452240450480615551792302106.072205164469778538"),
			},
		},
		{
			"largest per second interest rate before sdk.Uint overflows 256 bytes",
			args{
				perSecondInterestRate: sdkmath.LegacyMustNewDecFromStr("23.3"), // 23.4 overflows bit length 256 by 1 byte
				timeElapsed:           sdkmath.NewInt(30),                      // Assume a 30 second period, longer than any expected individual block
				expectedValue:         sdkmath.LegacyMustNewDecFromStr("104876366068119517411103023062013348034546.437155815200037999"),
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			interestFactor := keeper.CalculateBorrowInterestFactor(tc.args.perSecondInterestRate, tc.args.timeElapsed)
			suite.Require().Equal(tc.args.expectedValue, interestFactor)
		})
	}
}

func (suite *InterestTestSuite) TestCalculateSupplyInterestFactor() {
	type args struct {
		newInterest   sdkmath.LegacyDec
		cash          sdkmath.LegacyDec
		borrows       sdkmath.LegacyDec
		reserves      sdkmath.LegacyDec
		reserveFactor sdkmath.LegacyDec
		expectedValue sdkmath.LegacyDec
	}

	type test struct {
		name string
		args args
	}

	testCases := []test{
		{
			"low new interest",
			args{
				newInterest:   sdkmath.LegacyMustNewDecFromStr("1"),
				cash:          sdkmath.LegacyMustNewDecFromStr("100.0"),
				borrows:       sdkmath.LegacyMustNewDecFromStr("1000.0"),
				reserves:      sdkmath.LegacyMustNewDecFromStr("10.0"),
				reserveFactor: sdkmath.LegacyMustNewDecFromStr("0.05"),
				expectedValue: sdkmath.LegacyMustNewDecFromStr("1.000917431192660550"),
			},
		},
		{
			"medium new interest",
			args{
				newInterest:   sdkmath.LegacyMustNewDecFromStr("5"),
				cash:          sdkmath.LegacyMustNewDecFromStr("100.0"),
				borrows:       sdkmath.LegacyMustNewDecFromStr("1000.0"),
				reserves:      sdkmath.LegacyMustNewDecFromStr("10.0"),
				reserveFactor: sdkmath.LegacyMustNewDecFromStr("0.05"),
				expectedValue: sdkmath.LegacyMustNewDecFromStr("1.004587155963302752"),
			},
		},
		{
			"high new interest",
			args{
				newInterest:   sdkmath.LegacyMustNewDecFromStr("10"),
				cash:          sdkmath.LegacyMustNewDecFromStr("100.0"),
				borrows:       sdkmath.LegacyMustNewDecFromStr("1000.0"),
				reserves:      sdkmath.LegacyMustNewDecFromStr("10.0"),
				reserveFactor: sdkmath.LegacyMustNewDecFromStr("0.05"),
				expectedValue: sdkmath.LegacyMustNewDecFromStr("1.009174311926605505"),
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			interestFactor := keeper.CalculateSupplyInterestFactor(tc.args.newInterest,
				tc.args.cash, tc.args.borrows, tc.args.reserves)
			suite.Require().Equal(tc.args.expectedValue, interestFactor)
		})
	}
}

func (suite *InterestTestSuite) TestAPYToSPY() {
	type args struct {
		apy           sdkmath.LegacyDec
		expectedValue sdkmath.LegacyDec
	}

	type test struct {
		name        string
		args        args
		expectError bool
	}

	testCases := []test{
		{
			"lowest apy",
			args{
				apy:           sdkmath.LegacyMustNewDecFromStr("0.005"),
				expectedValue: sdkmath.LegacyMustNewDecFromStr("0.999999831991472557"),
			},
			false,
		},
		{
			"lower apy",
			args{
				apy:           sdkmath.LegacyMustNewDecFromStr("0.05"),
				expectedValue: sdkmath.LegacyMustNewDecFromStr("0.999999905005957279"),
			},
			false,
		},
		{
			"medium-low apy",
			args{
				apy:           sdkmath.LegacyMustNewDecFromStr("0.5"),
				expectedValue: sdkmath.LegacyMustNewDecFromStr("0.999999978020447332"),
			},
			false,
		},
		{
			"medium-high apy",
			args{
				apy:           sdkmath.LegacyMustNewDecFromStr("5"),
				expectedValue: sdkmath.LegacyMustNewDecFromStr("1.000000051034942717"),
			},
			false,
		},
		{
			"high apy",
			args{
				apy:           sdkmath.LegacyMustNewDecFromStr("50"),
				expectedValue: sdkmath.LegacyMustNewDecFromStr("1.000000124049443433"),
			},
			false,
		},
		{
			"highest apy",
			args{
				apy: sdkmath.LegacyMustNewDecFromStr("177"),
				// fixme previous was 1.000002441641340532
				expectedValue: sdkmath.LegacyMustNewDecFromStr("1.000000164134644767"),
			},
			false,
		},
		{
			"out of bounds error after 178",
			args{
				apy:           sdkmath.LegacyMustNewDecFromStr("179"),
				expectedValue: sdkmath.LegacyZeroDec(),
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			spy, err := keeper.APYToSPY(tc.args.apy)
			if tc.expectError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.args.expectedValue, spy)
			}
		})
	}
}

func (suite *InterestTestSuite) TestSPYToEstimatedAPY() {
	type args struct {
		spy             sdkmath.LegacyDec
		expectedAPY     float64
		acceptableRange float64
	}

	type test struct {
		name string
		args args
	}

	testCases := []test{
		{
			"lowest apy",
			args{
				spy:             sdkmath.LegacyMustNewDecFromStr("0.999999831991472557"),
				expectedAPY:     0.005,   // Returned value: 0.004999999888241291
				acceptableRange: 0.00001, // +/- 1/10000th of a precent
			},
		},
		{
			"lower apy",
			args{
				spy:             sdkmath.LegacyMustNewDecFromStr("0.999999905005957279"),
				expectedAPY:     0.05,    // Returned value: 0.05000000074505806
				acceptableRange: 0.00001, // +/- 1/10000th of a precent
			},
		},
		{
			"medium-low apy",
			args{
				spy:             sdkmath.LegacyMustNewDecFromStr("0.999999978020447332"),
				expectedAPY:     0.5,     // Returned value: 0.5
				acceptableRange: 0.00001, // +/- 1/10000th of a precent
			},
		},
		{
			"medium-high apy",
			args{
				spy:             sdkmath.LegacyMustNewDecFromStr("1.000000051034942717"),
				expectedAPY:     5,       // Returned value: 5
				acceptableRange: 0.00001, // +/- 1/10000th of a precent
			},
		},
		{
			"high apy",
			args{
				spy:             sdkmath.LegacyMustNewDecFromStr("1.000000124049443433"),
				expectedAPY:     50,      // Returned value: 50
				acceptableRange: 0.00001, // +/- 1/10000th of a precent
			},
		},
		{
			"highest apy",
			args{
				spy:             sdkmath.LegacyMustNewDecFromStr("1.000000146028999310"),
				expectedAPY:     100,     // 100
				acceptableRange: 0.00001, // +/- 1/10000th of a precent
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			// From SPY calculate APY and parse result from sdkmath.LegacyDec to float64
			calculatedAPY := keeper.SPYToEstimatedAPY(tc.args.spy)
			calculatedAPYFloat, err := strconv.ParseFloat(calculatedAPY.String(), 32)
			suite.Require().NoError(err)

			// Check that the calculated value is within an acceptable percentage range
			suite.Require().InEpsilon(tc.args.expectedAPY, calculatedAPYFloat, tc.args.acceptableRange)
		})
	}
}

type ExpectedBorrowInterest struct {
	elapsedTime  int64
	shouldBorrow bool
	borrowCoin   sdk.Coin
}

func (suite *KeeperTestSuite) TestBorrowInterest() {
	type args struct {
		user                     sdk.AccAddress
		initialBorrowerCoins     sdk.Coins
		initialModuleCoins       sdk.Coins
		borrowCoinDenom          string
		borrowCoins              sdk.Coins
		interestRateModel        types3.InterestRateModel
		reserveFactor            sdkmath.LegacyDec
		expectedInterestSnaphots []ExpectedBorrowInterest
	}

	type errArgs struct {
		expectPass bool
		contains   string
	}

	type interestTest struct {
		name    string
		args    args
		errArgs errArgs
	}

	normalModel := types3.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("0"), sdkmath.LegacyMustNewDecFromStr("0.1"), sdkmath.LegacyMustNewDecFromStr("0.8"), sdkmath.LegacyMustNewDecFromStr("0.5"))

	oneDayInSeconds := int64(86400)
	oneWeekInSeconds := int64(604800)
	oneMonthInSeconds := int64(2592000)
	oneYearInSeconds := int64(31536000)

	testCases := []interestTest{
		{
			"one day",
			args{
				user:                 sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf))),
				borrowCoinDenom:      "ujolt",
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(20*JoltCf))),
				interestRateModel:    normalModel,
				reserveFactor:        sdkmath.LegacyMustNewDecFromStr("0.05"),
				expectedInterestSnaphots: []ExpectedBorrowInterest{
					{
						elapsedTime:  oneDayInSeconds,
						shouldBorrow: false,
						borrowCoin:   sdk.Coin{},
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"one week",
			args{
				user:                 sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf))),
				borrowCoinDenom:      "ujolt",
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(20*JoltCf))),
				interestRateModel:    normalModel,
				reserveFactor:        sdkmath.LegacyMustNewDecFromStr("0.05"),
				expectedInterestSnaphots: []ExpectedBorrowInterest{
					{
						elapsedTime:  oneWeekInSeconds,
						shouldBorrow: false,
						borrowCoin:   sdk.Coin{},
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"one month",
			args{
				user:                 sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf))),
				borrowCoinDenom:      "ujolt",
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(20*JoltCf))),
				interestRateModel:    normalModel,
				reserveFactor:        sdkmath.LegacyMustNewDecFromStr("0.05"),
				expectedInterestSnaphots: []ExpectedBorrowInterest{
					{
						elapsedTime:  oneMonthInSeconds,
						shouldBorrow: false,
						borrowCoin:   sdk.Coin{},
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"one year",
			args{
				user:                 sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf))),
				borrowCoinDenom:      "ujolt",
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(20*JoltCf))),
				interestRateModel:    normalModel,
				reserveFactor:        sdkmath.LegacyMustNewDecFromStr("0.05"),
				expectedInterestSnaphots: []ExpectedBorrowInterest{
					{
						elapsedTime:  oneYearInSeconds,
						shouldBorrow: false,
						borrowCoin:   sdk.Coin{},
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"0 reserve factor",
			args{
				user:                 sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf))),
				borrowCoinDenom:      "ujolt",
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(20*JoltCf))),
				interestRateModel:    normalModel,
				reserveFactor:        sdkmath.LegacyMustNewDecFromStr("0"),
				expectedInterestSnaphots: []ExpectedBorrowInterest{
					{
						elapsedTime:  oneYearInSeconds,
						shouldBorrow: false,
						borrowCoin:   sdk.Coin{},
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"borrow during snapshot",
			args{
				user:                 sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf))),
				borrowCoinDenom:      "ujolt",
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(20*JoltCf))),
				interestRateModel:    normalModel,
				reserveFactor:        sdkmath.LegacyMustNewDecFromStr("0.05"),
				expectedInterestSnaphots: []ExpectedBorrowInterest{
					{
						elapsedTime:  oneYearInSeconds,
						shouldBorrow: true,
						borrowCoin:   sdk.NewCoin("ujolt", sdkmath.NewInt(1*JoltCf)),
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"multiple snapshots",
			args{
				user:                 sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf))),
				borrowCoinDenom:      "ujolt",
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(20*JoltCf))),
				interestRateModel:    normalModel,
				reserveFactor:        sdkmath.LegacyMustNewDecFromStr("0.05"),
				expectedInterestSnaphots: []ExpectedBorrowInterest{
					{
						elapsedTime:  oneMonthInSeconds,
						shouldBorrow: false,
						borrowCoin:   sdk.Coin{},
					},
					{
						elapsedTime:  oneMonthInSeconds,
						shouldBorrow: false,
						borrowCoin:   sdk.Coin{},
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"varied snapshots",
			args{
				user:                 sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf))),
				borrowCoinDenom:      "ujolt",
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(20*JoltCf))),
				interestRateModel:    normalModel,
				reserveFactor:        sdkmath.LegacyMustNewDecFromStr("0.05"),
				expectedInterestSnaphots: []ExpectedBorrowInterest{
					{
						elapsedTime:  oneDayInSeconds,
						shouldBorrow: false,
						borrowCoin:   sdk.Coin{},
					},
					{
						elapsedTime:  oneWeekInSeconds,
						shouldBorrow: false,
						borrowCoin:   sdk.Coin{},
					},
					{
						elapsedTime:  oneMonthInSeconds,
						shouldBorrow: false,
						borrowCoin:   sdk.Coin{},
					},
					{
						elapsedTime:  oneYearInSeconds,
						shouldBorrow: false,
						borrowCoin:   sdk.Coin{},
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			// Auth module genesis state
			authGS := app.NewFundedGenStateWithCoins(
				suite.app.AppCodec(),
				[]sdk.Coins{tc.args.initialBorrowerCoins},
				[]sdk.AccAddress{tc.args.user},
			)

			// Hard module genesis state
			hardGS := types3.NewGenesisState(types3.NewParams(
				types3.MoneyMarkets{
					types3.NewMoneyMarket("ujolt",
						types3.NewBorrowLimit(false, sdkmath.LegacyNewDec(100000000*JoltCf), sdkmath.LegacyMustNewDecFromStr("0.8")), // Borrow Limit
						"joltify:usd",             // Market ID
						sdkmath.NewInt(JoltCf),    // Conversion Factor
						tc.args.interestRateModel, // Interest Rate Model
						tc.args.reserveFactor,     // Reserve Factor
						sdkmath.LegacyZeroDec()),  // Keeper Reward Percentage
				},
				sdkmath.LegacyNewDec(10),
			), types3.DefaultAccumulationTimes, types3.DefaultDeposits, types3.DefaultBorrows,
				types3.DefaultTotalSupplied, types3.DefaultTotalBorrowed, types3.DefaultTotalReserves,
			)

			// Pricefeed module genesis state
			pricefeedGS := types2.GenesisState{
				Params: types2.Params{
					Markets: []types2.Market{
						{MarketID: "joltify:usd", BaseAsset: "joltify", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
					},
				},
				PostedPrices: []types2.PostedPrice{
					{
						MarketID:      "joltify:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdkmath.LegacyMustNewDecFromStr("2.00"),
						Expiry:        time.Now().Add(100 * time.Hour),
					},
				},
			}

			// Initialize test application
			mapp := suite.app.InitializeFromGenesisStates(suite.T(), time.Now(), nil, nil, authGS,
				app.GenesisState{types2.ModuleName: suite.app.AppCodec().MustMarshalJSON(&pricefeedGS)},
				app.GenesisState{types3.ModuleName: suite.app.AppCodec().MustMarshalJSON(&hardGS)})

			suite.app = mapp
			suite.app.App = mapp.App
			suite.ctx = mapp.Ctx
			suite.app.Ctx = mapp.Ctx
			ctx := mapp.Ctx
			suite.keeper = mapp.GetJoltKeeper()

			// Mint coins to Hard module account
			bankKeeper := suite.app.GetBankKeeper()
			err := bankKeeper.MintCoins(ctx, types3.ModuleAccountName, tc.args.initialModuleCoins)
			suite.Require().NoError(err)

			// Run begin blocker and store initial block time
			jolt.BeginBlocker(suite.ctx, suite.keeper)

			// Deposit 2x as many coins for each coin we intend to borrow
			depositCoins := sdk.NewCoins()
			for _, borrowCoin := range tc.args.borrowCoins {
				depositCoins = depositCoins.Add(sdk.NewCoin(borrowCoin.Denom, borrowCoin.Amount.Mul(sdkmath.NewInt(2))))
			}

			err = testutil.FundAccount(suite.ctx, suite.app.GetBankKeeper(), tc.args.user, tc.args.initialBorrowerCoins)
			suite.Require().NoError(err)

			err = suite.keeper.Deposit(suite.ctx, tc.args.user, depositCoins)
			suite.Require().NoError(err)

			// Borrow coins
			err = suite.keeper.Borrow(suite.ctx, tc.args.user, tc.args.borrowCoins)
			suite.Require().NoError(err)

			// Check that the initial module-level borrow balance is correct and store it
			initialBorrowedCoins, _ := suite.keeper.GetBorrowedCoins(suite.ctx)
			suite.Require().Equal(tc.args.borrowCoins, initialBorrowedCoins)

			// Check interest levels for each snapshot
			prevCtx := suite.ctx
			for _, snapshot := range tc.args.expectedInterestSnaphots {
				// ---------------------------- Calculate expected interest ----------------------------
				// 1. Get cash, borrows, reserves, and borrow index
				cashPrior := suite.getAccountCoins(suite.getModuleAccountAtCtx(types3.ModuleName, prevCtx)).AmountOf(tc.args.borrowCoinDenom)

				borrowCoinsPrior, borrowCoinsPriorFound := suite.keeper.GetBorrowedCoins(prevCtx)
				suite.Require().True(borrowCoinsPriorFound)
				borrowCoinPriorAmount := borrowCoinsPrior.AmountOf(tc.args.borrowCoinDenom)

				reservesPrior, foundReservesPrior := suite.keeper.GetTotalReserves(prevCtx)
				if !foundReservesPrior {
					reservesPrior = sdk.NewCoins(sdk.NewCoin(tc.args.borrowCoinDenom, sdkmath.ZeroInt()))
				}

				interestFactorPrior, foundInterestFactorPrior := suite.keeper.GetBorrowInterestFactor(prevCtx, tc.args.borrowCoinDenom)
				suite.Require().True(foundInterestFactorPrior)

				// 2. Calculate expected interest owed
				borrowRateApy, err := keeper.CalculateBorrowRate(tc.args.interestRateModel, sdkmath.LegacyNewDecFromInt(cashPrior), sdkmath.LegacyNewDecFromInt(borrowCoinPriorAmount), sdkmath.LegacyNewDecFromInt(reservesPrior.AmountOf(tc.args.borrowCoinDenom)))
				suite.Require().NoError(err)

				// Convert from APY to SPY, expressed as (1 + borrow rate)
				borrowRateSpy, err := keeper.APYToSPY(sdkmath.LegacyOneDec().Add(borrowRateApy))
				suite.Require().NoError(err)

				interestFactor := keeper.CalculateBorrowInterestFactor(borrowRateSpy, sdkmath.NewInt(snapshot.elapsedTime))
				expectedInterest := (interestFactor.Mul(sdkmath.LegacyNewDecFromInt(borrowCoinPriorAmount)).TruncateInt()).Sub(borrowCoinPriorAmount)
				expectedReserves := reservesPrior.Add(sdk.NewCoin(tc.args.borrowCoinDenom, sdkmath.LegacyNewDecFromInt(expectedInterest).Mul(tc.args.reserveFactor).TruncateInt()))
				expectedInterestFactor := interestFactorPrior.Mul(interestFactor)
				// -------------------------------------------------------------------------------------

				// Set up snapshot chain context and run begin blocker
				runAtTime := sdk.UnwrapSDKContext(prevCtx).BlockTime().Add(time.Duration(int64(time.Second) * snapshot.elapsedTime))
				snapshotCtx := sdk.UnwrapSDKContext(prevCtx).WithBlockTime(runAtTime)
				jolt.BeginBlocker(snapshotCtx, suite.keeper)

				// Check that the total amount of borrowed coins has increased by expected interest amount
				expectedBorrowedCoins := borrowCoinsPrior.AmountOf(tc.args.borrowCoinDenom).Add(expectedInterest)
				currBorrowedCoins, _ := suite.keeper.GetBorrowedCoins(snapshotCtx)
				suite.Require().Equal(expectedBorrowedCoins, currBorrowedCoins.AmountOf(tc.args.borrowCoinDenom))

				// Check that the total reserves have changed as expected
				currTotalReserves, _ := suite.keeper.GetTotalReserves(snapshotCtx)
				suite.Require().True(expectedReserves.Equal(currTotalReserves))

				// Check that the borrow index has increased as expected
				currIndexPrior, _ := suite.keeper.GetBorrowInterestFactor(snapshotCtx, tc.args.borrowCoinDenom)
				suite.Require().Equal(expectedInterestFactor, currIndexPrior)

				// After borrowing again user's borrow balance should have any outstanding interest applied
				if snapshot.shouldBorrow {
					borrowCoinsBefore, _ := suite.keeper.GetBorrow(snapshotCtx, tc.args.user)
					expectedInterestCoins := sdk.NewCoin(tc.args.borrowCoinDenom, expectedInterest)
					expectedBorrowCoinsAfter := borrowCoinsBefore.Amount.Add(snapshot.borrowCoin).Add(expectedInterestCoins)

					err = suite.keeper.Borrow(snapshotCtx, tc.args.user, sdk.NewCoins(snapshot.borrowCoin))
					suite.Require().NoError(err)

					borrowCoinsAfter, _ := suite.keeper.GetBorrow(snapshotCtx, tc.args.user)
					suite.Require().Equal(expectedBorrowCoinsAfter, borrowCoinsAfter.Amount)
				}
				// Update previous context to this snapshot's context, segmenting time periods between snapshots
				prevCtx = snapshotCtx
			}
		})
	}
}

type ExpectedSupplyInterest struct {
	elapsedTime  int64
	shouldSupply bool
	supplyCoin   sdk.Coin
}

func (suite *KeeperTestSuite) TestSupplyInterest() {
	type args struct {
		user                     sdk.AccAddress
		initialBorrowerCoins     sdk.Coins
		initialModuleCoins       sdk.Coins
		depositCoins             sdk.Coins
		coinDenoms               []string
		borrowCoins              sdk.Coins
		interestRateModel        types3.InterestRateModel
		reserveFactor            sdkmath.LegacyDec
		expectedInterestSnaphots []ExpectedSupplyInterest
	}

	type errArgs struct {
		expectPass bool
		contains   string
	}

	type interestTest struct {
		name    string
		args    args
		errArgs errArgs
	}

	normalModel := types3.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("0"), sdkmath.LegacyMustNewDecFromStr("0.1"), sdkmath.LegacyMustNewDecFromStr("0.8"), sdkmath.LegacyMustNewDecFromStr("0.5"))

	oneDayInSeconds := int64(86400)
	oneWeekInSeconds := int64(604800)
	oneMonthInSeconds := int64(2592000)
	oneYearInSeconds := int64(31536000)

	testCases := []interestTest{
		{
			"one day",
			args{
				user:                 sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				coinDenoms:           []string{"ujolt"},
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(20*JoltCf))),
				interestRateModel:    normalModel,
				reserveFactor:        sdkmath.LegacyMustNewDecFromStr("0.05"),
				expectedInterestSnaphots: []ExpectedSupplyInterest{
					{
						elapsedTime:  oneDayInSeconds,
						shouldSupply: false,
						supplyCoin:   sdk.Coin{},
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"one week",
			args{
				user:                 sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				coinDenoms:           []string{"ujolt"},
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(20*JoltCf))),
				interestRateModel:    normalModel,
				reserveFactor:        sdkmath.LegacyMustNewDecFromStr("0.05"),
				expectedInterestSnaphots: []ExpectedSupplyInterest{
					{
						elapsedTime:  oneWeekInSeconds,
						shouldSupply: false,
						supplyCoin:   sdk.Coin{},
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"one month",
			args{
				user:                 sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				coinDenoms:           []string{"ujolt"},
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(20*JoltCf))),
				interestRateModel:    normalModel,
				reserveFactor:        sdkmath.LegacyMustNewDecFromStr("0.05"),
				expectedInterestSnaphots: []ExpectedSupplyInterest{
					{
						elapsedTime:  oneMonthInSeconds,
						shouldSupply: false,
						supplyCoin:   sdk.Coin{},
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"one year",
			args{
				user:                 sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				coinDenoms:           []string{"ujolt"},
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(20*JoltCf))),
				interestRateModel:    normalModel,
				reserveFactor:        sdkmath.LegacyMustNewDecFromStr("0.05"),
				expectedInterestSnaphots: []ExpectedSupplyInterest{
					{
						elapsedTime:  oneYearInSeconds,
						shouldSupply: false,
						supplyCoin:   sdk.Coin{},
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"supply/borrow multiple coins",
			args{
				user:                 sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf)), sdk.NewCoin("bnb", sdkmath.NewInt(100*BnbCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf)), sdk.NewCoin("bnb", sdkmath.NewInt(100*BnbCf))),
				coinDenoms:           []string{"ujolt"},
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(20*JoltCf)), sdk.NewCoin("bnb", sdkmath.NewInt(20*BnbCf))),
				interestRateModel:    normalModel,
				reserveFactor:        sdkmath.LegacyMustNewDecFromStr("0.05"),
				expectedInterestSnaphots: []ExpectedSupplyInterest{
					{
						elapsedTime:  oneMonthInSeconds,
						shouldSupply: false,
						supplyCoin:   sdk.Coin{},
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"supply during snapshot",
			args{
				user:                 sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				coinDenoms:           []string{"ujolt"},
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(20*JoltCf))),
				interestRateModel:    normalModel,
				reserveFactor:        sdkmath.LegacyMustNewDecFromStr("0.05"),
				expectedInterestSnaphots: []ExpectedSupplyInterest{
					{
						elapsedTime:  oneMonthInSeconds,
						shouldSupply: true,
						supplyCoin:   sdk.NewCoin("ujolt", sdkmath.NewInt(20*JoltCf)),
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"multiple snapshots",
			args{
				user:                 sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				coinDenoms:           []string{"ujolt"},
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(80*JoltCf))),
				interestRateModel:    normalModel,
				reserveFactor:        sdkmath.LegacyMustNewDecFromStr("0.05"),
				expectedInterestSnaphots: []ExpectedSupplyInterest{
					{
						elapsedTime:  oneMonthInSeconds,
						shouldSupply: false,
						supplyCoin:   sdk.Coin{},
					},
					{
						elapsedTime:  oneMonthInSeconds,
						shouldSupply: false,
						supplyCoin:   sdk.Coin{},
					},
					{
						elapsedTime:  oneMonthInSeconds,
						shouldSupply: false,
						supplyCoin:   sdk.Coin{},
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"varied snapshots",
			args{
				user:                 sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				coinDenoms:           []string{"ujolt"},
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(50*JoltCf))),
				interestRateModel:    normalModel,
				reserveFactor:        sdkmath.LegacyMustNewDecFromStr("0.05"),
				expectedInterestSnaphots: []ExpectedSupplyInterest{
					{
						elapsedTime:  oneMonthInSeconds,
						shouldSupply: false,
						supplyCoin:   sdk.Coin{},
					},
					{
						elapsedTime:  oneDayInSeconds,
						shouldSupply: false,
						supplyCoin:   sdk.Coin{},
					},
					{
						elapsedTime:  oneYearInSeconds,
						shouldSupply: false,
						supplyCoin:   sdk.Coin{},
					},
					{
						elapsedTime:  oneWeekInSeconds,
						shouldSupply: false,
						supplyCoin:   sdk.Coin{},
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			// Initialize test app and set context
			// Auth module genesis state
			authGS := app.NewFundedGenStateWithCoins(
				suite.app.AppCodec(),
				[]sdk.Coins{tc.args.initialBorrowerCoins},
				[]sdk.AccAddress{tc.args.user},
			)

			// Hard module genesis state
			hardGS := types3.NewGenesisState(types3.NewParams(
				types3.MoneyMarkets{
					types3.NewMoneyMarket("ujolt",
						types3.NewBorrowLimit(false, sdkmath.LegacyNewDec(100000000*JoltCf), sdkmath.LegacyMustNewDecFromStr("0.8")), // Borrow Limit
						"joltify:usd",             // Market ID
						sdkmath.NewInt(JoltCf),    // Conversion Factor
						tc.args.interestRateModel, // Interest Rate Model
						tc.args.reserveFactor,     // Reserve Factor
						sdkmath.LegacyZeroDec()),  // Keeper Reward Percentage
					types3.NewMoneyMarket("bnb",
						types3.NewBorrowLimit(false, sdkmath.LegacyNewDec(100000000*BnbCf), sdkmath.LegacyMustNewDecFromStr("0.8")), // Borrow Limit
						"bnb:usd",                 // Market ID
						sdkmath.NewInt(BnbCf),     // Conversion Factor
						tc.args.interestRateModel, // Interest Rate Model
						tc.args.reserveFactor,     // Reserve Factor
						sdkmath.LegacyZeroDec()),  // Keeper Reward Percentage
				},
				sdkmath.LegacyNewDec(10),
			), types3.DefaultAccumulationTimes, types3.DefaultDeposits, types3.DefaultBorrows,
				types3.DefaultTotalSupplied, types3.DefaultTotalBorrowed, types3.DefaultTotalReserves,
			)

			// Pricefeed module genesis state
			pricefeedGS := types2.GenesisState{
				Params: types2.Params{
					Markets: []types2.Market{
						{MarketID: "joltify:usd", BaseAsset: "joltify", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "bnb:usd", BaseAsset: "bnb", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
					},
				},
				PostedPrices: []types2.PostedPrice{
					{
						MarketID:      "joltify:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdkmath.LegacyMustNewDecFromStr("2.00"),
						Expiry:        time.Now().Add(100 * time.Hour),
					},
					{
						MarketID:      "bnb:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdkmath.LegacyMustNewDecFromStr("20.00"),
						Expiry:        time.Now().Add(100 * time.Hour),
					},
				},
			}

			// Initialize test application
			mapp := suite.app.InitializeFromGenesisStates(suite.T(), time.Now(), nil, nil, authGS,
				app.GenesisState{types2.ModuleName: suite.app.AppCodec().MustMarshalJSON(&pricefeedGS)},
				app.GenesisState{types3.ModuleName: suite.app.AppCodec().MustMarshalJSON(&hardGS)})

			suite.app = mapp
			suite.app.App = mapp.App
			suite.ctx = mapp.Ctx
			suite.app.Ctx = mapp.Ctx
			suite.keeper = mapp.GetJoltKeeper()

			// Mint coins to Hard module account
			bankKeeper := suite.app.GetBankKeeper()
			err := bankKeeper.MintCoins(mapp.Ctx, types3.ModuleAccountName, tc.args.initialModuleCoins)
			suite.Require().NoError(err)

			suite.keeper.SetSuppliedCoins(mapp.Ctx, tc.args.initialModuleCoins)

			// Run begin blocker
			jolt.BeginBlocker(suite.ctx, suite.keeper)

			err = testutil.FundAccount(suite.ctx, suite.app.GetBankKeeper(), tc.args.user, tc.args.depositCoins)
			suite.Require().NoError(err)

			// // Deposit coins
			err = suite.keeper.Deposit(suite.ctx, tc.args.user, tc.args.depositCoins)
			suite.Require().NoError(err)

			// Borrow coins
			err = suite.keeper.Borrow(suite.ctx, tc.args.user, tc.args.borrowCoins)
			suite.Require().NoError(err)

			// Check interest levels for each snapshot
			prevCtx := suite.ctx
			for _, snapshot := range tc.args.expectedInterestSnaphots {
				for _, coinDenom := range tc.args.coinDenoms {
					// ---------------------------- Calculate expected supply interest ----------------------------
					// 1. Get cash, borrows, reserves, and borrow index
					cashPrior := suite.getAccountCoins(suite.getModuleAccountAtCtx(types3.ModuleName, prevCtx)).AmountOf(coinDenom)

					var borrowCoinPriorAmount sdkmath.Int
					borrowCoinsPrior, borrowCoinsPriorFound := suite.keeper.GetBorrowedCoins(prevCtx)
					suite.Require().True(borrowCoinsPriorFound)
					borrowCoinPriorAmount = borrowCoinsPrior.AmountOf(coinDenom)

					var supplyCoinPriorAmount sdkmath.Int
					supplyCoinsPrior, supplyCoinsPriorFound := suite.keeper.GetSuppliedCoins(prevCtx)
					suite.Require().True(supplyCoinsPriorFound)
					supplyCoinPriorAmount = supplyCoinsPrior.AmountOf(coinDenom)

					reservesPrior, foundReservesPrior := suite.keeper.GetTotalReserves(prevCtx)
					if !foundReservesPrior {
						reservesPrior = sdk.NewCoins(sdk.NewCoin(coinDenom, sdkmath.ZeroInt()))
					}

					borrowInterestFactorPrior, foundBorrowInterestFactorPrior := suite.keeper.GetBorrowInterestFactor(prevCtx, coinDenom)
					suite.Require().True(foundBorrowInterestFactorPrior)

					supplyInterestFactorPrior, foundSupplyInterestFactorPrior := suite.keeper.GetSupplyInterestFactor(prevCtx, coinDenom)
					suite.Require().True(foundSupplyInterestFactorPrior)

					// 2. Calculate expected borrow interest owed
					borrowRateApy, err := keeper.CalculateBorrowRate(tc.args.interestRateModel, sdkmath.LegacyNewDecFromInt(cashPrior), sdkmath.LegacyNewDecFromInt(borrowCoinPriorAmount), sdkmath.LegacyNewDecFromInt(reservesPrior.AmountOf(coinDenom)))
					suite.Require().NoError(err)

					// Convert from APY to SPY, expressed as (1 + borrow rate)
					borrowRateSpy, err := keeper.APYToSPY(sdkmath.LegacyOneDec().Add(borrowRateApy))
					suite.Require().NoError(err)

					newBorrowInterestFactor := keeper.CalculateBorrowInterestFactor(borrowRateSpy, sdkmath.NewInt(snapshot.elapsedTime))
					expectedBorrowInterest := (newBorrowInterestFactor.Mul(sdkmath.LegacyNewDecFromInt(borrowCoinPriorAmount)).TruncateInt()).Sub(borrowCoinPriorAmount)
					expectedReserves := reservesPrior.Add(sdk.NewCoin(coinDenom, sdkmath.LegacyNewDecFromInt(expectedBorrowInterest).Mul(tc.args.reserveFactor).TruncateInt())).Sub(reservesPrior...)
					expectedTotalReserves := expectedReserves.Add(reservesPrior...)

					expectedBorrowInterestFactor := borrowInterestFactorPrior.Mul(newBorrowInterestFactor)
					expectedSupplyInterest := expectedBorrowInterest.Sub(expectedReserves.AmountOf(coinDenom))

					newSupplyInterestFactor := keeper.CalculateSupplyInterestFactor(sdkmath.LegacyNewDecFromInt(expectedSupplyInterest), sdkmath.LegacyNewDecFromInt(cashPrior), sdkmath.LegacyNewDecFromInt(borrowCoinPriorAmount), sdkmath.LegacyNewDecFromInt(reservesPrior.AmountOf(coinDenom)))
					expectedSupplyInterestFactor := supplyInterestFactorPrior.Mul(newSupplyInterestFactor)
					// -------------------------------------------------------------------------------------

					// Set up snapshot chain context and run begin blocker
					runAtTime := sdk.UnwrapSDKContext(prevCtx).BlockTime().Add(time.Duration(int64(time.Second) * snapshot.elapsedTime))
					snapshotCtx := sdk.UnwrapSDKContext(prevCtx).WithBlockTime(runAtTime)
					jolt.BeginBlocker(snapshotCtx, suite.keeper)

					borrowInterestFactor, _ := suite.keeper.GetBorrowInterestFactor(mapp.Ctx, coinDenom)
					suite.Require().Equal(expectedBorrowInterestFactor, borrowInterestFactor)
					suite.Require().Equal(expectedBorrowInterest, expectedSupplyInterest.Add(expectedReserves.AmountOf(coinDenom)))

					// Check that the total amount of borrowed coins has increased by expected borrow interest amount
					borrowCoinsPost, _ := suite.keeper.GetBorrowedCoins(snapshotCtx)
					borrowCoinPostAmount := borrowCoinsPost.AmountOf(coinDenom)
					suite.Require().Equal(borrowCoinPostAmount, borrowCoinPriorAmount.Add(expectedBorrowInterest))

					// Check that the total amount of supplied coins has increased by expected supply interest amount
					supplyCoinsPost, _ := suite.keeper.GetSuppliedCoins(prevCtx)
					supplyCoinPostAmount := supplyCoinsPost.AmountOf(coinDenom)
					suite.Require().Equal(supplyCoinPostAmount, supplyCoinPriorAmount.Add(expectedSupplyInterest))

					// Check current total reserves
					totalReserves, _ := suite.keeper.GetTotalReserves(snapshotCtx)
					suite.Require().Equal(
						sdk.NewCoin(coinDenom, expectedTotalReserves.AmountOf(coinDenom)),
						sdk.NewCoin(coinDenom, totalReserves.AmountOf(coinDenom)),
					)

					// Check that the supply index has increased as expected
					currSupplyIndexPrior, _ := suite.keeper.GetSupplyInterestFactor(snapshotCtx, coinDenom)
					suite.Require().Equal(expectedSupplyInterestFactor, currSupplyIndexPrior)

					// // Check that the borrow index has increased as expected
					currBorrowIndexPrior, _ := suite.keeper.GetBorrowInterestFactor(snapshotCtx, coinDenom)
					suite.Require().Equal(expectedBorrowInterestFactor, currBorrowIndexPrior)

					// After supplying again user's supplied balance should have owed supply interest applied
					if snapshot.shouldSupply {
						// Calculate percentage of supply interest profits owed to user
						userSupplyBefore, _ := suite.keeper.GetDeposit(snapshotCtx, tc.args.user)
						userSupplyCoinAmount := userSupplyBefore.Amount.AmountOf(coinDenom)
						userPercentOfTotalSupplied := sdkmath.LegacyNewDecFromInt(userSupplyCoinAmount).Quo(sdkmath.LegacyNewDecFromInt(supplyCoinPriorAmount))
						userExpectedSupplyInterestCoin := sdk.NewCoin(coinDenom, userPercentOfTotalSupplied.MulInt(expectedSupplyInterest).TruncateInt())

						// Supplying syncs user's owed supply and borrow interest
						err = suite.keeper.Deposit(snapshotCtx, tc.args.user, sdk.NewCoins(snapshot.supplyCoin))
						suite.Require().NoError(err)

						// Fetch user's new borrow and supply balance post-interaction
						userSupplyAfter, _ := suite.keeper.GetDeposit(snapshotCtx, tc.args.user)

						// Confirm that user's supply index for the denom has increased as expected
						var userSupplyAfterIndexFactor sdkmath.LegacyDec
						for _, indexFactor := range userSupplyAfter.Index {
							if indexFactor.Denom == coinDenom {
								userSupplyAfterIndexFactor = indexFactor.Value
							}
						}
						suite.Require().Equal(userSupplyAfterIndexFactor, currSupplyIndexPrior)

						// Check user's supplied amount increased by supply interest owed + the newly supplied coins
						expectedSupplyCoinsAfter := userSupplyBefore.Amount.Add(snapshot.supplyCoin).Add(userExpectedSupplyInterestCoin)
						suite.Require().Equal(expectedSupplyCoinsAfter, userSupplyAfter.Amount)
					}
					prevCtx = snapshotCtx
				}
			}
		})
	}
}

func TestInterestTestSuite(t *testing.T) {
	suite.Run(t, new(InterestTestSuite))
}
