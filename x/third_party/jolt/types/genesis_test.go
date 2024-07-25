package types_test

import (
	"strings"
	"testing"
	"time"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	UsdxCf = 1000000
)

type GenesisTestSuite struct {
	suite.Suite
}

func (suite *GenesisTestSuite) TestGenesisValidation() {
	type args struct {
		params types2.Params
		gats   types2.GenesisAccumulationTimes
		deps   types2.Deposits
		brws   types2.Borrows
		ts     sdk.Coins
		tb     sdk.Coins
		tr     sdk.Coins
	}
	testCases := []struct {
		name        string
		args        args
		expectPass  bool
		expectedErr string
	}{
		{
			name: "default",
			args: args{
				params: types2.DefaultParams(),
				gats:   types2.DefaultAccumulationTimes,
				deps:   types2.DefaultDeposits,
				brws:   types2.DefaultBorrows,
				ts:     types2.DefaultTotalSupplied,
				tb:     types2.DefaultTotalBorrowed,
				tr:     types2.DefaultTotalReserves,
			},
			expectPass:  true,
			expectedErr: "",
		},
		{
			name: "valid",
			args: args{
				params: types2.NewParams(
					types2.MoneyMarkets{
						types2.NewMoneyMarket("usdx", types2.NewBorrowLimit(true, sdk.MustNewDecFromStr("100000000000"), sdk.MustNewDecFromStr("1")), "usdx:usd", sdk.NewInt(UsdxCf), types2.NewInterestRateModel(sdk.MustNewDecFromStr("0.05"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10")), sdk.MustNewDecFromStr("0.05"), sdk.ZeroDec()),
					},
					sdk.MustNewDecFromStr("10"),
				),
				gats: types2.GenesisAccumulationTimes{
					types2.NewGenesisAccumulationTime("usdx", time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC), sdkmath.LegacyOneDec(), sdkmath.LegacyOneDec()),
				},
				deps: types2.DefaultDeposits,
				brws: types2.DefaultBorrows,
				ts:   sdk.Coins{},
				tb:   sdk.Coins{},
				tr:   sdk.Coins{},
			},
			expectPass:  true,
			expectedErr: "",
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			gs := types2.NewGenesisState(tc.args.params, tc.args.gats, tc.args.deps, tc.args.brws, tc.args.ts, tc.args.tb, tc.args.tr)
			err := gs.Validate()
			if tc.expectPass {
				suite.NoError(err)
			} else {
				suite.Error(err)
				suite.Require().True(strings.Contains(err.Error(), tc.expectedErr))
			}
		})
	}
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}
