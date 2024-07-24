package types

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// d is a helper function for creating sdkmath.LegacyDec values in tests
func d(str string) sdkmath.LegacyDec { return sdk.MustNewDecFromStr(str) }

// c is a helper function for created sdk.Coin types in tests
func c(denom string, amount int64) sdk.Coin { return sdk.NewInt64Coin(denom, amount) }

// c is a helper function for created sdk.Coins types in tests
func cs(coins ...sdk.Coin) sdk.Coins { return sdk.NewCoins(coins...) }

func TestRewardIndexes(t *testing.T) {
	t.Run("With", func(t *testing.T) {
		arbitraryDec := sdk.MustNewDecFromStr("0.1")

		type args struct {
			denom  string
			factor sdkmath.LegacyDec
		}
		testcases := []struct {
			name          string
			rewardIndexes RewardIndexes
			args          args
			expected      RewardIndexes
		}{
			{
				name: "when index is not present, it's added and original isn't overwritten",
				rewardIndexes: RewardIndexes{
					NewRewardIndex("denom", arbitraryDec),
				},
				args: args{
					denom:  "otherdenom",
					factor: arbitraryDec,
				},
				expected: RewardIndexes{
					NewRewardIndex("denom", arbitraryDec),
					NewRewardIndex("otherdenom", arbitraryDec),
				},
			},
			{
				name: "when index is present, it's updated and original isn't overwritten",
				rewardIndexes: RewardIndexes{
					NewRewardIndex("denom", arbitraryDec),
				},
				args: args{
					denom:  "denom",
					factor: arbitraryDec.MulInt64(2),
				},
				expected: RewardIndexes{
					NewRewardIndex("denom", arbitraryDec.MulInt64(2)),
				},
			},
		}

		for _, tc := range testcases {
			t.Run(tc.name, func(t *testing.T) {
				newIndexes := tc.rewardIndexes.With(tc.args.denom, tc.args.factor)

				require.Equal(t, tc.expected, newIndexes)
				require.NotEqual(t, tc.rewardIndexes, newIndexes) // check original slice not modified
			})
		}
	})
	t.Run("Get", func(t *testing.T) {
		arbitraryDec := sdk.MustNewDecFromStr("0.1")

		type expected struct {
			factor sdkmath.LegacyDec
			found  bool
		}
		testcases := []struct {
			name          string
			rewardIndexes RewardIndexes
			argDenom      string
			expected      expected
		}{
			{
				name: "when index is present, it is found and returned",
				rewardIndexes: RewardIndexes{
					NewRewardIndex("denom", arbitraryDec),
				},
				argDenom: "denom",
				expected: expected{
					factor: arbitraryDec,
					found:  true,
				},
			},
			{
				name: "when index is not present, it is not found",
				rewardIndexes: RewardIndexes{
					NewRewardIndex("denom", arbitraryDec),
				},
				argDenom: "notpresent",
				expected: expected{
					found: false,
				},
			},
		}

		for _, tc := range testcases {
			t.Run(tc.name, func(t *testing.T) {
				factor, found := tc.rewardIndexes.Get(tc.argDenom)

				require.Equal(t, tc.expected.found, found)
				require.Equal(t, tc.expected.factor, factor)
			})
		}
	})
	t.Run("Mul", func(t *testing.T) {
		testcases := []struct {
			name          string
			rewardIndexes RewardIndexes
			multiplier    sdkmath.LegacyDec
			expected      RewardIndexes
		}{
			{
				name: "non zero values are all multiplied",
				rewardIndexes: RewardIndexes{
					NewRewardIndex("denom", d("0.1")),
					NewRewardIndex("denom2", d("0.2")),
				},
				multiplier: d("2.0"),
				expected: RewardIndexes{
					NewRewardIndex("denom", d("0.2")),
					NewRewardIndex("denom2", d("0.4")),
				},
			},
			{
				name: "multiplying by zero, zeros all values",
				rewardIndexes: RewardIndexes{
					NewRewardIndex("denom", d("0.1")),
					NewRewardIndex("denom2", d("0.0")),
				},
				multiplier: d("0.0"),
				expected: RewardIndexes{
					NewRewardIndex("denom", d("0.0")),
					NewRewardIndex("denom2", d("0.0")),
				},
			},
			{
				name:          "empty indexes are unchanged",
				rewardIndexes: RewardIndexes{},
				multiplier:    d("2.0"),
				expected:      RewardIndexes{},
			},
			{
				name:          "nil indexes are unchanged",
				rewardIndexes: nil,
				multiplier:    d("2.0"),
				expected:      nil,
			},
		}

		for _, tc := range testcases {
			t.Run(tc.name, func(t *testing.T) {
				calculated := tc.rewardIndexes.Mul(tc.multiplier)
				require.Equal(t, len(tc.expected), len(calculated))
				for i, el := range tc.expected {
					require.Equal(t, el.RewardFactor.String(), calculated[i].RewardFactor.String())
				}
			})
		}
	})
	t.Run("Quo", func(t *testing.T) {
		type expected struct {
			indexes RewardIndexes
			panics  bool
		}
		testcases := []struct {
			name          string
			rewardIndexes RewardIndexes
			divisor       sdkmath.LegacyDec
			expected      expected
		}{
			{
				name: "non zero values are all divided",
				rewardIndexes: RewardIndexes{
					NewRewardIndex("denom", d("0.6")),
					NewRewardIndex("denom2", d("0.2")),
				},
				divisor: d("3.0"),
				expected: expected{
					indexes: RewardIndexes{
						NewRewardIndex("denom", d("0.2")),
						NewRewardIndex("denom2", d("0.066666666666666667")),
					},
				},
			},
			{
				name: "diving by zero panics when values are present",
				rewardIndexes: RewardIndexes{
					NewRewardIndex("denom", d("0.1")),
					NewRewardIndex("denom2", d("0.0")),
				},
				divisor: d("0.0"),
				expected: expected{
					panics: true,
				},
			},
			{
				name:          "empty indexes are unchanged",
				rewardIndexes: RewardIndexes{},
				divisor:       d("2.0"),
				expected: expected{
					indexes: RewardIndexes{},
				},
			},
			{
				name:          "nil indexes are unchanged",
				rewardIndexes: nil,
				divisor:       d("2.0"),
				expected: expected{
					indexes: nil,
				},
			},
		}

		for _, tc := range testcases {
			t.Run(tc.name, func(t *testing.T) {
				var actual RewardIndexes
				quoFunc := func() { actual = tc.rewardIndexes.Quo(tc.divisor) }
				if tc.expected.panics {
					require.Panics(t, quoFunc)
					return
				} else {
					require.NotPanics(t, quoFunc)
				}
				require.Equal(t, tc.expected.indexes, actual)
			})
		}
	})
	t.Run("Add", func(t *testing.T) {
		testcases := []struct {
			name          string
			rewardIndexes RewardIndexes
			addend        RewardIndexes
			expected      RewardIndexes
		}{
			{
				name: "same denoms are added",
				rewardIndexes: RewardIndexes{
					NewRewardIndex("denom", d("0.1")),
					NewRewardIndex("denom2", d("0.2")),
				},
				addend: RewardIndexes{
					NewRewardIndex("denom", d("0.1")),
					NewRewardIndex("denom2", d("0.2")),
				},
				expected: RewardIndexes{
					NewRewardIndex("denom", d("0.2")),
					NewRewardIndex("denom2", d("0.4")),
				},
			},
			{
				name: "new denoms are appended",
				rewardIndexes: RewardIndexes{
					NewRewardIndex("denom", d("0.1")),
				},
				addend: RewardIndexes{
					NewRewardIndex("denom", d("0.3")),
					NewRewardIndex("denom2", d("0.2")),
				},
				expected: RewardIndexes{
					NewRewardIndex("denom", d("0.4")),
					NewRewardIndex("denom2", d("0.2")),
				},
			},
			{
				name: "missing denoms are unchanged",
				rewardIndexes: RewardIndexes{
					NewRewardIndex("denom", d("0.1")),
					NewRewardIndex("denom2", d("0.2")),
				},
				addend: RewardIndexes{
					NewRewardIndex("denom2", d("0.2")),
				},
				expected: RewardIndexes{
					NewRewardIndex("denom", d("0.1")),
					NewRewardIndex("denom2", d("0.4")),
				},
			},
			{
				name: "adding empty indexes does nothing",
				rewardIndexes: RewardIndexes{
					NewRewardIndex("denom", d("0.1")),
				},
				addend: RewardIndexes{},
				expected: RewardIndexes{
					NewRewardIndex("denom", d("0.1")),
				},
			},
			{
				name: "adding nil indexes does nothing",
				rewardIndexes: RewardIndexes{
					NewRewardIndex("denom", d("0.1")),
				},
				addend: nil,
				expected: RewardIndexes{
					NewRewardIndex("denom", d("0.1")),
				},
			},
			{
				name:          "denom can be added to empty indexes",
				rewardIndexes: RewardIndexes{},
				addend: RewardIndexes{
					NewRewardIndex("denom", d("0.1")),
				},
				expected: RewardIndexes{
					NewRewardIndex("denom", d("0.1")),
				},
			},
			{
				name:          "denom can be added to nil indexes",
				rewardIndexes: nil,
				addend: RewardIndexes{
					NewRewardIndex("denom", d("0.1")),
				},
				expected: RewardIndexes{
					NewRewardIndex("denom", d("0.1")),
				},
			},
			{
				name:          "adding empty indexes to nil does nothing",
				rewardIndexes: nil,
				addend:        RewardIndexes{},
				expected:      nil,
			},
			{
				name:          "adding nil to empty indexes does nothing",
				rewardIndexes: RewardIndexes{},
				addend:        nil,
				expected:      RewardIndexes{},
			},
			{
				name:          "adding nil to nil indexes does nothing",
				rewardIndexes: nil,
				addend:        nil,
				expected:      nil,
			},
			{
				name:          "adding empty indexes to empty indexes does nothing",
				rewardIndexes: RewardIndexes{},
				addend:        RewardIndexes{},
				expected:      RewardIndexes{},
			},
		}
		for _, tc := range testcases {
			t.Run(tc.name, func(t *testing.T) {
				sum := tc.rewardIndexes.Add(tc.addend)
				require.Equal(t, tc.expected, sum)
			})
		}
	})
}

func TestMultiRewardIndexes(t *testing.T) {
	arbitraryRewardIndexes := RewardIndexes{
		{
			CollateralType: "reward",
			RewardFactor:   sdk.MustNewDecFromStr("0.1"),
		},
	}

	t.Run("Get", func(t *testing.T) {
		type expected struct {
			rewardIndexes RewardIndexes
			found         bool
		}
		testcases := []struct {
			name               string
			multiRewardIndexes MultiRewardIndexes
			argDenom           string
			expected           expected
		}{
			{
				name: "when indexes are present, they are found and returned",
				multiRewardIndexes: MultiRewardIndexes{
					{
						CollateralType: "denom",
						RewardIndexes:  arbitraryRewardIndexes,
					},
				},
				argDenom: "denom",
				expected: expected{
					found:         true,
					rewardIndexes: arbitraryRewardIndexes,
				},
			},
			{
				name: "when indexes are not present, they are not found",
				multiRewardIndexes: MultiRewardIndexes{
					{
						CollateralType: "denom",
						RewardIndexes:  arbitraryRewardIndexes,
					},
				},
				argDenom: "notpresent",
				expected: expected{
					found: false,
				},
			},
		}
		for _, tc := range testcases {
			t.Run(tc.name, func(t *testing.T) {
				rewardIndexes, found := tc.multiRewardIndexes.Get(tc.argDenom)

				require.Equal(t, tc.expected.found, found)
				require.Equal(t, tc.expected.rewardIndexes, rewardIndexes)
			})
		}
	})
	t.Run("With", func(t *testing.T) {
		type args struct {
			denom         string
			rewardIndexes RewardIndexes
		}
		testcases := []struct {
			name               string
			multiRewardIndexes MultiRewardIndexes
			args               args
			expected           MultiRewardIndexes
		}{
			{
				name: "when indexes are not present, add them and do not update original",
				multiRewardIndexes: MultiRewardIndexes{
					{
						CollateralType: "denom",
						RewardIndexes:  arbitraryRewardIndexes,
					},
				},
				args: args{
					denom:         "otherdenom",
					rewardIndexes: arbitraryRewardIndexes,
				},
				expected: MultiRewardIndexes{
					{
						CollateralType: "denom",
						RewardIndexes:  arbitraryRewardIndexes,
					},
					{
						CollateralType: "otherdenom",
						RewardIndexes:  arbitraryRewardIndexes,
					},
				},
			},
			{
				name: "when indexes are present, update them and do not update original",
				multiRewardIndexes: MultiRewardIndexes{
					{
						CollateralType: "denom",
						RewardIndexes:  arbitraryRewardIndexes,
					},
				},
				args: args{
					denom:         "denom",
					rewardIndexes: appendUniqueRewardIndex(arbitraryRewardIndexes),
				},
				expected: MultiRewardIndexes{
					{
						CollateralType: "denom",
						RewardIndexes:  appendUniqueRewardIndex(arbitraryRewardIndexes),
					},
				},
			},
		}
		for _, tc := range testcases {
			t.Run(tc.name, func(t *testing.T) {
				oldIndexes := tc.multiRewardIndexes.copy()

				newIndexes := tc.multiRewardIndexes.With(tc.args.denom, tc.args.rewardIndexes)

				require.Equal(t, tc.expected, newIndexes)
				require.Equal(t, oldIndexes, tc.multiRewardIndexes)
			})
		}
	})
	t.Run("RemoveRewardIndex", func(t *testing.T) {
		testcases := []struct {
			name               string
			multiRewardIndexes MultiRewardIndexes
			argDenom           string
			expected           MultiRewardIndexes
		}{
			{
				name: "when indexes are not present, do nothing",
				multiRewardIndexes: MultiRewardIndexes{
					{
						CollateralType: "denom",
						RewardIndexes:  arbitraryRewardIndexes,
					},
				},
				argDenom: "notpresent",
				expected: MultiRewardIndexes{
					{
						CollateralType: "denom",
						RewardIndexes:  arbitraryRewardIndexes,
					},
				},
			},
			{
				name: "when indexes are present, remove them and do not update original",
				multiRewardIndexes: MultiRewardIndexes{
					{
						CollateralType: "denom",
						RewardIndexes:  arbitraryRewardIndexes,
					},
					{
						CollateralType: "otherdenom",
						RewardIndexes:  arbitraryRewardIndexes,
					},
				},
				argDenom: "denom",
				expected: MultiRewardIndexes{
					{
						CollateralType: "otherdenom",
						RewardIndexes:  arbitraryRewardIndexes,
					},
				},
			},
		}
		for _, tc := range testcases {
			t.Run(tc.name, func(t *testing.T) {
				oldIndexes := tc.multiRewardIndexes.copy()

				newIndexes := tc.multiRewardIndexes.RemoveRewardIndex(tc.argDenom)

				require.Equal(t, tc.expected, newIndexes)
				require.Equal(t, oldIndexes, tc.multiRewardIndexes)
			})
		}
	})
	t.Run("Validate", func(t *testing.T) {
		testcases := []struct {
			name               string
			multiRewardIndexes MultiRewardIndexes
			wantErr            bool
		}{
			{
				name: "normal case",
				multiRewardIndexes: MultiRewardIndexes{
					{CollateralType: "btcb", RewardIndexes: normalRewardIndexes},
					{CollateralType: "bnb", RewardIndexes: normalRewardIndexes},
				},
				wantErr: false,
			},
			{
				name:               "empty",
				multiRewardIndexes: nil,
				wantErr:            false,
			},
			{
				name: "empty collateral type",
				multiRewardIndexes: MultiRewardIndexes{
					{RewardIndexes: normalRewardIndexes},
				},
				wantErr: true,
			},
			{
				name: "invalid reward index",
				multiRewardIndexes: MultiRewardIndexes{
					{CollateralType: "btcb", RewardIndexes: invalidRewardIndexes},
				},
				wantErr: true,
			},
		}
		for _, tc := range testcases {
			t.Run(tc.name, func(t *testing.T) {
				err := tc.multiRewardIndexes.Validate()
				if tc.wantErr {
					require.NotNil(t, err)
				} else {
					require.Nil(t, err)
				}
			})
		}
	})
}

var normalRewardIndexes = RewardIndexes{
	NewRewardIndex("jolt", sdk.MustNewDecFromStr("0.000001")),
	NewRewardIndex("ujolt", sdk.MustNewDecFromStr("0.1")),
}

var invalidRewardIndexes = RewardIndexes{
	RewardIndex{"jolt", sdk.MustNewDecFromStr("-0.01")},
}

func appendUniqueRewardIndex(indexes RewardIndexes) RewardIndexes {
	const uniqueDenom = "uniquereward"

	for _, mri := range indexes {
		if mri.CollateralType == uniqueDenom {
			panic(fmt.Sprintf("tried to add unique reward index with denom '%s', but denom already existed", uniqueDenom))
		}
	}

	return append(
		indexes,
		NewRewardIndex(uniqueDenom, sdk.MustNewDecFromStr("0.02")),
	)
}
