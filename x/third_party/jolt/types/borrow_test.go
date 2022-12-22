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
		expect    sdk.DecCoins
		expectErr string
	}{
		{
			name: "multiple denoms are calculated correctly",
			borrow: types2.Borrow{
				Amount: sdk.NewCoins(
					sdk.NewInt64Coin("bnb", 100e8),
					sdk.NewInt64Coin("xrpb", 1e8),
				),
				Index: types2.BorrowInterestFactors{
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
			name: "empty borrow amount returns empty dec coins",
			borrow: types2.Borrow{
				Amount: sdk.Coins{},
				Index:  types2.BorrowInterestFactors{},
			},
			expect: sdk.DecCoins{},
		},
		{
			name: "nil borrow amount returns empty dec coins",
			borrow: types2.Borrow{
				Amount: nil,
				Index:  types2.BorrowInterestFactors{},
			},
			expect: sdk.DecCoins{},
		},
		{
			name: "missing indexes return error",
			borrow: types2.Borrow{
				Amount: sdk.NewCoins(
					sdk.NewInt64Coin("bnb", 100e8),
				),
				Index: types2.BorrowInterestFactors{
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
			borrow: types2.Borrow{
				Amount: sdk.NewCoins(
					sdk.NewInt64Coin("bnb", 100e8),
				),
				Index: types2.BorrowInterestFactors{
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
			borrow: types2.Borrow{
				Amount: sdk.NewCoins(
					sdk.NewInt64Coin("bnb", 100e8),
				),
				Index: types2.BorrowInterestFactors{
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
			nb, err := tc.borrow.NormalizedBorrow()

			require.Equal(t, tc.expect, nb)

			if len(tc.expectErr) > 0 {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectErr)
			}
		})
	}
}
