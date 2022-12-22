package types_test

import (
	"fmt"
	"testing"
	"time"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/suite"
)

type ParamTestSuite struct {
	suite.Suite
}

func (suite *ParamTestSuite) SetupTest() {}

var rewardPeriodWithInvalidRewardsPerSecond = types2.NewRewardPeriod(
	true,
	"bnb",
	time.Date(2020, 10, 15, 14, 0, 0, 0, time.UTC),
	time.Date(2024, 10, 15, 14, 0, 0, 0, time.UTC),
	sdk.Coin{Denom: "INVALID!@#😫", Amount: sdk.ZeroInt()},
)

var rewardMultiPeriodWithInvalidRewardsPerSecond = types2.NewMultiRewardPeriod(
	true,
	"bnb",
	time.Date(2020, 10, 15, 14, 0, 0, 0, time.UTC),
	time.Date(2024, 10, 15, 14, 0, 0, 0, time.UTC),
	sdk.Coins{sdk.Coin{Denom: "INVALID!@#😫", Amount: sdk.ZeroInt()}},
)

var validMultiRewardPeriod = types2.NewMultiRewardPeriod(
	true,
	"bnb",
	time.Date(2020, 10, 15, 14, 0, 0, 0, time.UTC),
	time.Date(2024, 10, 15, 14, 0, 0, 0, time.UTC),
	sdk.NewCoins(sdk.NewInt64Coin("swap", 1e9)),
)

var validRewardPeriod = types2.NewRewardPeriod(
	true,
	"bnb-a",
	time.Date(2020, 10, 15, 14, 0, 0, 0, time.UTC),
	time.Date(2024, 10, 15, 14, 0, 0, 0, time.UTC),
	sdk.NewInt64Coin(types2.USDXMintingRewardDenom, 1e9),
)

func (suite *ParamTestSuite) TestParamValidation() {
	type errArgs struct {
		expectPass bool
		contains   string
	}
	type test struct {
		name    string
		params  types2.Params
		errArgs errArgs
	}

	testCases := []test{
		{
			"default is valid",
			types2.DefaultParams(),
			errArgs{
				expectPass: true,
			},
		},
		{
			"valid",
			types2.Params{
				USDXMintingRewardPeriods: types2.RewardPeriods{
					types2.NewRewardPeriod(
						true,
						"bnb-a",
						time.Date(2020, 10, 15, 14, 0, 0, 0, time.UTC),
						time.Date(2024, 10, 15, 14, 0, 0, 0, time.UTC),
						sdk.NewCoin(types2.USDXMintingRewardDenom, sdk.NewInt(122354)),
					),
				},
				JoltSupplyRewardPeriods: types2.DefaultMultiRewardPeriods,
				JoltBorrowRewardPeriods: types2.DefaultMultiRewardPeriods,
				DelegatorRewardPeriods:  types2.DefaultMultiRewardPeriods,
				SwapRewardPeriods:       types2.DefaultMultiRewardPeriods,
				SavingsRewardPeriods:    types2.DefaultMultiRewardPeriods,
				ClaimMultipliers: types2.MultipliersPerDenoms{
					{
						Denom: "jolt",
						Multipliers: types2.Multipliers{
							types2.NewMultiplier("small", 1, sdk.MustNewDecFromStr("0.25")),
							types2.NewMultiplier("large", 12, sdk.MustNewDecFromStr("1.0")),
						},
					},
					{
						Denom: "ujolt",
						Multipliers: types2.Multipliers{
							types2.NewMultiplier("small", 1, sdk.MustNewDecFromStr("0.2")),
							types2.NewMultiplier("large", 12, sdk.MustNewDecFromStr("1.0")),
						},
					},
				},
				ClaimEnd: time.Date(2025, 10, 15, 14, 0, 0, 0, time.UTC),
			},
			errArgs{
				expectPass: true,
			},
		},
		{
			"invalid usdx minting period makes params invalid",
			types2.Params{
				USDXMintingRewardPeriods: types2.RewardPeriods{rewardPeriodWithInvalidRewardsPerSecond},
				JoltSupplyRewardPeriods:  types2.DefaultMultiRewardPeriods,
				JoltBorrowRewardPeriods:  types2.DefaultMultiRewardPeriods,
				DelegatorRewardPeriods:   types2.DefaultMultiRewardPeriods,
				SwapRewardPeriods:        types2.DefaultMultiRewardPeriods,
				SavingsRewardPeriods:     types2.DefaultMultiRewardPeriods,
				ClaimMultipliers:         types2.DefaultMultipliers,
				ClaimEnd:                 time.Date(2025, 10, 15, 14, 0, 0, 0, time.UTC),
			},
			errArgs{
				expectPass: false,
				contains:   fmt.Sprintf("reward denom must be %s", types2.USDXMintingRewardDenom),
			},
		},
		{
			"invalid jolt supply periods makes params invalid",
			types2.Params{
				USDXMintingRewardPeriods: types2.DefaultRewardPeriods,
				JoltSupplyRewardPeriods:  types2.MultiRewardPeriods{rewardMultiPeriodWithInvalidRewardsPerSecond},
				JoltBorrowRewardPeriods:  types2.DefaultMultiRewardPeriods,
				DelegatorRewardPeriods:   types2.DefaultMultiRewardPeriods,
				SwapRewardPeriods:        types2.DefaultMultiRewardPeriods,
				SavingsRewardPeriods:     types2.DefaultMultiRewardPeriods,
				ClaimMultipliers:         types2.DefaultMultipliers,
				ClaimEnd:                 time.Date(2025, 10, 15, 14, 0, 0, 0, time.UTC),
			},
			errArgs{
				expectPass: false,
				contains:   "invalid reward amount",
			},
		},
		{
			"invalid jolt borrow periods makes params invalid",
			types2.Params{
				USDXMintingRewardPeriods: types2.DefaultRewardPeriods,
				JoltSupplyRewardPeriods:  types2.DefaultMultiRewardPeriods,
				JoltBorrowRewardPeriods:  types2.MultiRewardPeriods{rewardMultiPeriodWithInvalidRewardsPerSecond},
				DelegatorRewardPeriods:   types2.DefaultMultiRewardPeriods,
				SwapRewardPeriods:        types2.DefaultMultiRewardPeriods,
				SavingsRewardPeriods:     types2.DefaultMultiRewardPeriods,
				ClaimMultipliers:         types2.DefaultMultipliers,
				ClaimEnd:                 time.Date(2025, 10, 15, 14, 0, 0, 0, time.UTC),
			},
			errArgs{
				expectPass: false,
				contains:   "invalid reward amount",
			},
		},
		{
			"invalid delegator periods makes params invalid",
			types2.Params{
				USDXMintingRewardPeriods: types2.DefaultRewardPeriods,
				JoltSupplyRewardPeriods:  types2.DefaultMultiRewardPeriods,
				JoltBorrowRewardPeriods:  types2.DefaultMultiRewardPeriods,
				DelegatorRewardPeriods:   types2.MultiRewardPeriods{rewardMultiPeriodWithInvalidRewardsPerSecond},
				SwapRewardPeriods:        types2.DefaultMultiRewardPeriods,
				SavingsRewardPeriods:     types2.DefaultMultiRewardPeriods,
				ClaimMultipliers:         types2.DefaultMultipliers,
				ClaimEnd:                 time.Date(2025, 10, 15, 14, 0, 0, 0, time.UTC),
			},
			errArgs{
				expectPass: false,
				contains:   "invalid reward amount",
			},
		},
		{
			"invalid swap periods makes params invalid",
			types2.Params{
				USDXMintingRewardPeriods: types2.DefaultRewardPeriods,
				JoltSupplyRewardPeriods:  types2.DefaultMultiRewardPeriods,
				JoltBorrowRewardPeriods:  types2.DefaultMultiRewardPeriods,
				DelegatorRewardPeriods:   types2.DefaultMultiRewardPeriods,
				SwapRewardPeriods:        types2.MultiRewardPeriods{rewardMultiPeriodWithInvalidRewardsPerSecond},
				SavingsRewardPeriods:     types2.DefaultMultiRewardPeriods,
				ClaimMultipliers:         types2.DefaultMultipliers,
				ClaimEnd:                 time.Date(2025, 10, 15, 14, 0, 0, 0, time.UTC),
			},
			errArgs{
				expectPass: false,
				contains:   "invalid reward amount",
			},
		},
		{
			"invalid multipliers makes params invalid",
			types2.Params{
				USDXMintingRewardPeriods: types2.DefaultRewardPeriods,
				JoltSupplyRewardPeriods:  types2.DefaultMultiRewardPeriods,
				JoltBorrowRewardPeriods:  types2.DefaultMultiRewardPeriods,
				DelegatorRewardPeriods:   types2.DefaultMultiRewardPeriods,
				SwapRewardPeriods:        types2.DefaultMultiRewardPeriods,
				SavingsRewardPeriods:     types2.DefaultMultiRewardPeriods,
				ClaimMultipliers: types2.MultipliersPerDenoms{
					{
						Denom: "jolt",
						Multipliers: types2.Multipliers{
							types2.NewMultiplier("small", -9999, sdk.MustNewDecFromStr("0.25")),
						},
					},
				},
				ClaimEnd: time.Date(2025, 10, 15, 14, 0, 0, 0, time.UTC),
			},
			errArgs{
				expectPass: false,
				contains:   "expected non-negative lockup",
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			err := tc.params.Validate()

			if tc.errArgs.expectPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.errArgs.contains)
			}
		})
	}
}

func (suite *ParamTestSuite) TestRewardPeriods() {
	suite.Run("Validate", func() {
		type err struct {
			pass     bool
			contains string
		}
		testCases := []struct {
			name    string
			periods types2.RewardPeriods
			expect  err
		}{
			{
				name: "single period is valid",
				periods: types2.RewardPeriods{
					validRewardPeriod,
				},
				expect: err{
					pass: true,
				},
			},
			{
				name: "duplicated reward period is invalid",
				periods: types2.RewardPeriods{
					validRewardPeriod,
					validRewardPeriod,
				},
				expect: err{
					contains: "duplicated reward period",
				},
			},
			{
				name: "invalid reward denom is invalid",
				periods: types2.RewardPeriods{
					types2.NewRewardPeriod(
						true,
						"bnb-a",
						time.Date(2020, 10, 15, 14, 0, 0, 0, time.UTC),
						time.Date(2024, 10, 15, 14, 0, 0, 0, time.UTC),
						sdk.NewInt64Coin("jolt", 1e9),
					),
				},
				expect: err{
					contains: fmt.Sprintf("reward denom must be %s", types2.USDXMintingRewardDenom),
				},
			},
		}
		for _, tc := range testCases {

			err := tc.periods.Validate()

			if tc.expect.pass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Contains(err.Error(), tc.expect.contains)
			}
		}
	})
}

func (suite *ParamTestSuite) TestMultiRewardPeriods() {
	suite.Run("Validate", func() {
		type err struct {
			pass     bool
			contains string
		}
		testCases := []struct {
			name    string
			periods types2.MultiRewardPeriods
			expect  err
		}{
			{
				name: "single period is valid",
				periods: types2.MultiRewardPeriods{
					validMultiRewardPeriod,
				},
				expect: err{
					pass: true,
				},
			},
			{
				name: "duplicated reward period is invalid",
				periods: types2.MultiRewardPeriods{
					validMultiRewardPeriod,
					validMultiRewardPeriod,
				},
				expect: err{
					contains: "duplicated reward period",
				},
			},
			{
				name: "invalid reward period is invalid",
				periods: types2.MultiRewardPeriods{
					rewardMultiPeriodWithInvalidRewardsPerSecond,
				},
				expect: err{
					contains: "invalid reward amount",
				},
			},
		}
		for _, tc := range testCases {

			err := tc.periods.Validate()

			if tc.expect.pass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Contains(err.Error(), tc.expect.contains)
			}
		}
	})
}

func TestParamTestSuite(t *testing.T) {
	suite.Run(t, new(ParamTestSuite))
}
