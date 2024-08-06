package types

import (
	"strings"
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	type errArgs struct {
		expectPass bool
		contains   string
	}

	pa := NewParams(
		DefaultMultiRewardPeriods,
		DefaultMultiRewardPeriods,
		DefaultMultiRewardPeriods,
		DefaultMultiRewardPeriods,
		MultipliersPerDenoms{
			{
				Denom: "ujolt",
				Multipliers: Multipliers{
					NewMultiplier("small", 1, sdkmath.LegacyMustNewDecFromStr("0.33")),
					NewMultiplier("large", 12, sdkmath.LegacyMustNewDecFromStr("1.00")),
				},
			},
		},
		time.Date(2025, 10, 15, 14, 0, 0, 0, time.UTC),
	)
	//
	//pa2 := NewParams(
	//	DefaultMultiRewardPeriods,
	//	DefaultMultiRewardPeriods,
	//	nil,
	//	time.Date(2025, 10, 15, 14, 0, 0, 0, time.UTC),
	//)

	state := GenesisRewardState{
		AccumulationTimes: AccumulationTimes{
			{CollateralType: "", PreviousAccumulationTime: normalAccumulationtime},
		},
	}

	state2 := GenesisRewardState{
		AccumulationTimes: AccumulationTimes{
			{CollateralType: "bnb", PreviousAccumulationTime: normalAccumulationtime},
		},
	}

	claimState := JoltLiquidityProviderClaim{}

	testCases := []struct {
		name    string
		genesis GenesisState
		errArgs errArgs
	}{
		{
			name:    "default",
			genesis: DefaultGenesisState(),
			errArgs: errArgs{
				expectPass: true,
			},
		},
		{
			name: "valid",
			genesis: GenesisState{
				Params: pa,
			},
			errArgs: errArgs{
				expectPass: true,
			},
		},
		{
			name: "invalid genesis accumulation time",
			genesis: GenesisState{
				Params:                DefaultParams(),
				JoltSupplyRewardState: state,
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "collateral type must be defined",
			},
		},
		{
			name: "invalid claim",
			genesis: GenesisState{
				Params:                      DefaultParams(),
				JoltSupplyRewardState:       state2,
				JoltLiquidityProviderClaims: JoltLiquidityProviderClaims{claimState},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "claim owner cannot be empty",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.genesis.Validate()
			if tc.errArgs.expectPass {
				require.NoError(t, err, tc.name)
			} else {
				require.Error(t, err, tc.name)
				require.True(t, strings.Contains(err.Error(), tc.errArgs.contains))
			}
		})
	}
}

func TestGenesisAccumulationTimes_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		gats    AccumulationTimes
		wantErr bool
	}{
		{
			name: "normal",
			gats: AccumulationTimes{
				{CollateralType: "btcb", PreviousAccumulationTime: normalAccumulationtime},
				{CollateralType: "bnb", PreviousAccumulationTime: normalAccumulationtime},
			},
			wantErr: false,
		},
		{
			name:    "empty",
			gats:    nil,
			wantErr: false,
		},
		{
			name: "empty collateral type",
			gats: AccumulationTimes{
				{PreviousAccumulationTime: normalAccumulationtime},
			},
			wantErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.gats.Validate()
			if tc.wantErr {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

var normalAccumulationtime = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
