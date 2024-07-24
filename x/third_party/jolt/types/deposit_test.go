package types_test

import (
	"testing"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestDeposit_NormalizedDeposit(t *testing.T) {
	testCases := []struct {
		name      string
		deposit   types2.Deposit
		expect    sdkmath.LegacyDecCoins
		expectErr string
	}{
		{
			name: "multiple denoms are calculated correctly",
			deposit: types2.Deposit{
				Amount: sdk.NewCoins(
					sdk.NewInt64Coin("bnb", 100e8),
					sdk.NewInt64Coin("xrpb", 1e8),
				),
				Index: types2.SupplyInterestFactors{
					{
						Denom: "xrpb",
						Value: sdk.MustNewDecFromStr("1.25"),
					},
					{
						Denom: "bnb",
						Value: sdk.MustNewDecFromStr("2.0"),
					},
				},
			},
			expect: sdk.NewDecCoins(
				sdk.NewInt64DecCoin("bnb", 50e8),
				sdk.NewInt64DecCoin("xrpb", 8e7),
			),
		},
		{
			name: "empty deposit amount returns empty dec coins",
			deposit: types2.Deposit{
				Amount: sdk.Coins{},
				Index:  types2.SupplyInterestFactors{},
			},
			expect: sdkmath.LegacyDecCoins{},
		},
		{
			name: "nil deposit amount returns empty dec coins",
			deposit: types2.Deposit{
				Amount: nil,
				Index:  types2.SupplyInterestFactors{},
			},
			expect: sdkmath.LegacyDecCoins{},
		},
		{
			name: "missing indexes return error",
			deposit: types2.Deposit{
				Amount: sdk.NewCoins(
					sdk.NewInt64Coin("bnb", 100e8),
				),
				Index: types2.SupplyInterestFactors{
					{
						Denom: "xrpb",
						Value: sdk.MustNewDecFromStr("1.25"),
					},
				},
			},
			expectErr: "missing interest factor",
		},
		{
			name: "invalid indexes return error",
			deposit: types2.Deposit{
				Amount: sdk.NewCoins(
					sdk.NewInt64Coin("bnb", 100e8),
				),
				Index: types2.SupplyInterestFactors{
					{
						Denom: "bnb",
						Value: sdk.MustNewDecFromStr("0.999999999999999999"),
					},
				},
			},
			expectErr: "< 1",
		},
		{
			name: "zero indexes return error rather than panicking",
			deposit: types2.Deposit{
				Amount: sdk.NewCoins(
					sdk.NewInt64Coin("bnb", 100e8),
				),
				Index: types2.SupplyInterestFactors{
					{
						Denom: "bnb",
						Value: sdk.MustNewDecFromStr("0"),
					},
				},
			},
			expectErr: "< 1",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nb, err := tc.deposit.NormalizedDeposit()

			require.Equal(t, tc.expect, nb)

			if len(tc.expectErr) > 0 {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectErr)
			}
		})
	}
}
