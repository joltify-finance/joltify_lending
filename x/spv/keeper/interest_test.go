package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"testing"
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
				apy:           sdk.MustNewDecFromStr("0.005"),
				payfrq:        4,
				expectedValue: sdk.MustNewDecFromStr("0.999999831991472557"),
			},
			false,
		},

		{
			"lowest apy",
			args{
				apy:           sdk.MustNewDecFromStr("0.051271109622422061"),
				payfrq:        4,
				expectedValue: sdk.MustNewDecFromStr("0.999999831991472557"),
			},
			false,
		},
		{
			"lower apy",
			args{
				apy:           sdk.MustNewDecFromStr("0.05"),
				payfrq:        4,
				expectedValue: sdk.MustNewDecFromStr("0.999999905005957279"),
			},
			false,
		},
		{
			"medium-low apy",
			args{
				apy:           sdk.MustNewDecFromStr("0.5"),
				payfrq:        4,
				expectedValue: sdk.MustNewDecFromStr("0.999999978020447332"),
			},
			false,
		},
		{
			"medium-high apy",
			args{
				apy:           sdk.MustNewDecFromStr("5"),
				payfrq:        4,
				expectedValue: sdk.MustNewDecFromStr("1.000000051034942717"),
			},
			false,
		},
		{
			"high apy",
			args{
				apy:           sdk.MustNewDecFromStr("50"),
				payfrq:        4,
				expectedValue: sdk.MustNewDecFromStr("1.000000124049443433"),
			},
			false,
		},
		{
			"highest apy",
			args{
				apy:    sdk.MustNewDecFromStr("177"),
				payfrq: 4,
				// fixme previous was 1.000002441641340532
				expectedValue: sdk.MustNewDecFromStr("1.000000164134644767"),
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
			i, err := CalculateInterest(tc.args.apy, tc.args.payfrq)
			if tc.expectError {
				suite.Require().NoError(err)
			} else {
				suite.Require().NoError(err)

				accTime := tc.args.payfrq * OneWeek
				accumulate := i.Power(uint64(accTime))

				total := accumulate.Mul(sdk.NewDec(OneYear / int64(accTime)))
				gap := total.Sub(tc.args.apy)
				fmt.Printf("gap>>>>>%v\n", gap)
				suite.Require().True(gap.LT(sdk.NewDecFromIntWithPrec(sdk.NewInt(1), 8)))

			}
		})
	}
}

func (suite *InterestTestSuite) TestMe() {

	d := sdk.MustNewDecFromStr("1.05127")

	rate, err := d.ApproxRoot(OneYear)
	suite.Require().NoError(err)
	fmt.Printf(">>>>%v\n", rate.String())

	i := sdk.MustNewDecFromStr("0.0000000015854900")

	_ = i

	rate = rate.Sub(sdk.NewDec(1))
	p := sdk.NewDec(1)
	for j := 0; j < OneYear; j++ {
		in := p.Mul(rate)
		p = p.Add(in)
	}
	fmt.Printf(">>>>%v\n", p.String())

}

func (suite *InterestTestSuite) TestMe2() {

	for qq := 1; qq < 13; qq++ {
		d := sdk.MustNewDecFromStr("1.05127")

		d2 := d.Quo(sdk.NewDec(OneYear / OneWeek / int64(qq)))

		root, err := d2.ApproxRoot(OneWeek * uint64(qq))
		suite.Require().NoError(err)
		fmt.Printf(">>>>%v\n", root.String())

		//i := sdk.MustNewDecFromStr("0.0000000015854900")

		i := root.Sub(sdk.NewDec(1))

		p := sdk.NewDec(1)
		for j := 0; j < OneWeek*qq; j++ {
			in := p.Mul(i)
			p = p.Add(in)
		}

		total := p.Mul(sdk.NewDec(OneYear / OneWeek / int64(qq)))

		fmt.Printf(">>>>month %v :%v\n", qq, total.String())
	}
}

type InterestTestSuite struct {
	suite.Suite
}

func TestInterestTestSuite(t *testing.T) {
	suite.Run(t, new(InterestTestSuite))
}
