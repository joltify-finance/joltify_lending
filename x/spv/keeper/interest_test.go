package keeper

import (
	"fmt"
	"testing"

	sdkmath "cosmossdk.io/math"

	"github.com/stretchr/testify/suite"
)

func (suite *InterestTestSuite) TestAPYToSPY() {
	type args struct {
		apy           sdkmath.LegacyDec
		payfrq        int
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
				apy:    sdkmath.LegacyMustNewDecFromStr("0.005"),
				payfrq: 3600 * 24 * 7,
			},
			false,
		},

		{
			"lowest apy",
			args{
				apy:    sdkmath.LegacyMustNewDecFromStr("0.051271109622422061"),
				payfrq: 4,
			},
			false,
		},
		{
			"lower apy",
			args{
				apy:    sdkmath.LegacyMustNewDecFromStr("0.05"),
				payfrq: 4,
			},
			false,
		},
		{
			"medium-low apy",
			args{
				apy:    sdkmath.LegacyMustNewDecFromStr("0.5"),
				payfrq: 4,
			},
			false,
		},
		{
			"medium-high apy",
			args{
				apy:    sdkmath.LegacyMustNewDecFromStr("5"),
				payfrq: 4,
			},
			false,
		},
		{
			"high apy",
			args{
				apy:    sdkmath.LegacyMustNewDecFromStr("50"),
				payfrq: 4,
			},
			false,
		},
		{
			"highest apy",
			args{
				apy:    sdkmath.LegacyMustNewDecFromStr("177"),
				payfrq: 4,
			},
			false,
		},
		{
			"out of bounds error after 178",
			args{
				apy:           sdkmath.LegacyMustNewDecFromStr("179"),
				payfrq:        4,
				expectedValue: sdkmath.LegacyZeroDec(),
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			i := CalculateInterestRate(tc.args.apy, tc.args.payfrq)

			accTime := tc.args.payfrq
			accumulate := i.Power(uint64(accTime))

			total := (accumulate.Sub(sdkmath.LegacyOneDec())).Mul(sdkmath.LegacyNewDec(OneYear / int64(accTime)))
			gap := total.Sub(tc.args.apy)
			suite.Require().True(gap.LT(sdkmath.LegacyNewDecFromIntWithPrec(sdkmath.NewInt(1), 8)))
		})
	}
}

func checkPayFreqApy(oneYearApy sdkmath.LegacyDec, freqApy sdkmath.LegacyDec, circle uint64) bool {
	return oneYearApy.Sub(freqApy.MulInt(sdkmath.NewIntFromUint64(circle))).Abs().LTE(sdkmath.LegacyNewDecWithPrec(1, 8))
}

func (suite *InterestTestSuite) TestCalculateInterestAmount() {
	testapy := sdkmath.LegacyMustNewDecFromStr("0.15")
	_, err := CalculateInterestAmount(testapy, 0)
	suite.Require().ErrorContains(err, "payFreq cannot be zero")
	for i := 1; i < 52; i++ {
		payfreq := OneWeek * i
		result, err := CalculateInterestAmount(testapy, payfreq)
		suite.Require().NoError(err)
		circle := OneYear / payfreq
		checkPayFreqApy(testapy, result, uint64(circle))
	}
}

func (suite *InterestTestSuite) TestCalculateInterestFactor() {
	testapy := sdkmath.LegacyMustNewDecFromStr("0.25")
	payfreq := OneWeek * 8

	apyEachPayment, err := CalculateInterestAmount(testapy, payfreq)
	suite.Require().NoError(err)
	spy, err := apyTospy(apyEachPayment, uint64(payfreq))
	suite.Require().NoError(err)

	result := CalculateInterestFactor(spy, sdkmath.NewIntFromUint64(uint64(payfreq)))
	suite.Require().True(apyEachPayment.Sub(result).Abs().LTE(sdkmath.LegacyNewDecWithPrec(1, 8)))
}

func (suite *InterestTestSuite) TestCalculateInterestPerSecond() {
	testapy := sdkmath.LegacyMustNewDecFromStr("0.18")

	adjMonthAPY := sdkmath.LegacyOneDec().Add(testapy)

	val, err := apyTospy(adjMonthAPY, uint64(OneYear))

	fmt.Printf(">>>>val is %v\n", val.String())

	result := (val.Sub(sdkmath.LegacyOneDec())).Mul(sdkmath.LegacyNewDec(OneYear))

	gap := testapy.Sub(result)

	fmt.Printf(">>>>%v===%v(%v)\n", result, gap, gap.Quo(testapy))

	suite.Require().NoError(err)
	fmt.Printf(">>>>%v\n", val.String())
}

type InterestTestSuite struct {
	suite.Suite
}

func TestInterestTestSuite(t *testing.T) {
	suite.Run(t, new(InterestTestSuite))
}
