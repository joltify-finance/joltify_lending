package types_test

import (
	"testing"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

type ParamTestSuite struct {
	suite.Suite
}

func (suite *ParamTestSuite) TestParamValidation() {
	type args struct {
		minBorrowVal sdk.Dec
		mms          types2.MoneyMarkets
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
				minBorrowVal: types2.DefaultMinimumBorrowUSDValue,
				mms:          types2.DefaultMoneyMarkets,
			},
			expectPass:  true,
			expectedErr: "",
		},
		{
			name: "invalid: conversion factor < one",
			args: args{
				minBorrowVal: types2.DefaultMinimumBorrowUSDValue,
				mms: types2.MoneyMarkets{
					{
						Denom: "btcb",
						BorrowLimit: types2.NewBorrowLimit(
							false,
							sdk.MustNewDecFromStr("100000000000"),
							sdk.MustNewDecFromStr("0.5"),
						),
						SpotMarketID:           "btc:usd",
						ConversionFactor:       sdk.NewInt(0),
						InterestRateModel:      types2.InterestRateModel{},
						ReserveFactor:          sdk.MustNewDecFromStr("0.05"),
						KeeperRewardPercentage: sdk.MustNewDecFromStr("0.05"),
					},
				},
			},
			expectPass:  false,
			expectedErr: "conversion '0' factor must be ≥ one",
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			params := types2.NewParams(tc.args.mms, tc.args.minBorrowVal)
			err := params.Validate()
			if tc.expectPass {
				suite.NoError(err)
			} else {
				suite.Error(err)
				suite.Require().Contains(err.Error(), tc.expectedErr)
			}
		})
	}
}

func TestParamTestSuite(t *testing.T) {
	suite.Run(t, new(ParamTestSuite))
}
