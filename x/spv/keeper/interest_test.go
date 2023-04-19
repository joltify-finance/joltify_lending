package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

func (suite *InterestTestSuite) TestAPYToSPY() {
	type args struct {
		apy           sdk.Dec
		payfrq        int
		expectedValue sdk.Dec
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
				apy:    sdk.MustNewDecFromStr("0.005"),
				payfrq: 3600 * 24 * 7,
			},
			false,
		},

		{
			"lowest apy",
			args{
				apy:    sdk.MustNewDecFromStr("0.051271109622422061"),
				payfrq: 4,
			},
			false,
		},
		{
			"lower apy",
			args{
				apy:    sdk.MustNewDecFromStr("0.05"),
				payfrq: 4,
			},
			false,
		},
		{
			"medium-low apy",
			args{
				apy:    sdk.MustNewDecFromStr("0.5"),
				payfrq: 4,
			},
			false,
		},
		{
			"medium-high apy",
			args{
				apy:    sdk.MustNewDecFromStr("5"),
				payfrq: 4,
			},
			false,
		},
		{
			"high apy",
			args{
				apy:    sdk.MustNewDecFromStr("50"),
				payfrq: 4,
			},
			false,
		},
		{
			"highest apy",
			args{
				apy:    sdk.MustNewDecFromStr("177"),
				payfrq: 4,
			},
			false,
		},
		{
			"out of bounds error after 178",
			args{
				apy:           sdk.MustNewDecFromStr("179"),
				payfrq:        4,
				expectedValue: sdk.ZeroDec(),
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			i := CalculateInterestRate(tc.args.apy, tc.args.payfrq)

			accTime := tc.args.payfrq
			accumulate := i.Power(uint64(accTime))

			total := (accumulate.Sub(sdk.OneDec())).Mul(sdk.NewDec(OneYear / int64(accTime)))
			gap := total.Sub(tc.args.apy)
			suite.Require().True(gap.LT(sdk.NewDecFromIntWithPrec(sdk.NewInt(1), 8)))

		})
	}
}

func checkPayFreqApy(oneYearApy sdk.Dec, freqApy sdk.Dec, circle uint64) bool {
	return oneYearApy.Sub(freqApy.MulInt(sdk.NewIntFromUint64(circle))).Abs().LTE(sdk.NewDecWithPrec(1, 8))
}

func (suite *InterestTestSuite) TestCalculateInterestAmount() {

	testapy := sdk.MustNewDecFromStr("0.15")
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

	testapy := sdk.MustNewDecFromStr("0.25")
	payfreq := OneWeek * 8

	apyEachPayment, err := CalculateInterestAmount(testapy, payfreq)
	suite.Require().NoError(err)
	spy, err := apyTospy(apyEachPayment, uint64(payfreq))
	suite.Require().NoError(err)

	result := CalculateInterestFactor(spy, sdk.NewIntFromUint64(uint64(payfreq)))
	suite.Require().True(apyEachPayment.Sub(result).Abs().LTE(sdk.NewDecWithPrec(1, 8)))
}

type InterestTestSuite struct {
	suite.Suite
}

func TestInterestTestSuite(t *testing.T) {
	suite.Run(t, new(InterestTestSuite))
}
