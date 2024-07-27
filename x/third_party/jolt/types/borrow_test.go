package types_test

import (
	"testing"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestBorrow_NormalizedBorrow(t *testing.T) {
	testCases := []struct {
		name      string
		borrow    types2.Borrow
		expect    sdkmath.LegacyDecCoins
		expectErr string
	}{
		{
			name: "multiple denoms are calculated correctly",
			borrow: types2.Borrow{
				Amount: sdk.NewCoins(
					sdkmath.NewInt64Coin("bnb", 100e8),
					sdkmath.NewInt64Coin("xrpb", 1e8),
				),
				Index: types2.BorrowInterestFactors{
					{
						Denom: "xrpb",
						Value: sdkmath.LegacyMustNewDecFromStr("1.25"),
					},
					{
						Denom: "bnb",
						Value: sdkmath.LegacyMustNewDecFromStr("2.0"),
					},
				},
			},
			expect: sdk.NewDecCoins(
				sdkmath.NewInt64DecCoin("bnb", 50e8),
				sdkmath.NewInt64DecCoin("xrpb", 8e7),
			),
		},
		{
			name: "empty borrow amount returns empty dec coins",
			borrow: types2.Borrow{
				Amount: sdk.Coins{},
				Index:  types2.BorrowInterestFactors{},
			},
			expect: sdkmath.LegacyDecCoins{},
		},
		{
			name: "nil borrow amount returns empty dec coins",
			borrow: types2.Borrow{
				Amount: nil,
				Index:  types2.BorrowInterestFactors{},
			},
			expect: sdkmath.LegacyDecCoins{},
		},
		{
			name: "missing indexes return error",
			borrow: types2.Borrow{
				Amount: sdk.NewCoins(
					sdkmath.NewInt64Coin("bnb", 100e8),
				),
				Index: types2.BorrowInterestFactors{
					{
						Denom: "xrpb",
						Value: sdkmath.LegacyMustNewDecFromStr("1.25"),
					},
				},
			},
			expectErr: "missing interest factor",
		},
		{
			name: "invalid indexes return error",
			borrow: types2.Borrow{
				Amount: sdk.NewCoins(
					sdkmath.NewInt64Coin("bnb", 100e8),
				),
				Index: types2.BorrowInterestFactors{
					{
						Denom: "bnb",
						Value: sdkmath.LegacyMustNewDecFromStr("0.999999999999999999"),
					},
				},
			},
			expectErr: "< 1",
		},
		{
			name: "zero indexes return error rather than panicking",
			borrow: types2.Borrow{
				Amount: sdk.NewCoins(
					sdkmath.NewInt64Coin("bnb", 100e8),
				),
				Index: types2.BorrowInterestFactors{
					{
						Denom: "bnb",
						Value: sdkmath.LegacyMustNewDecFromStr("0"),
					},
				},
			},
			expectErr: "< 1",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nb, err := tc.borrow.NormalizedBorrow()

			require.Equal(t, tc.expect, nb)

			if len(tc.expectErr) > 0 {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectErr)
			}
		})
	}
}
