package types_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/constants"
	types "github.com/joltify-finance/joltify_lending/x/third_party_dydx/perpetuals/types"
	"github.com/stretchr/testify/require"
)

// validAuthority is a valid bech32 address string.
var validAuthority = constants.AliceAccAddress.String()

func TestMsgCreatePerpetual_ValidateBasic(t *testing.T) {
	tests := []struct {
		desc        string
		msg         types.MsgCreatePerpetual
		expectedErr string
	}{
		{
			desc: "Invalid authority",
			msg: types.MsgCreatePerpetual{
				Authority: "",
			},
			expectedErr: "Authority is invalid",
		},
		{
			desc: "Empty ticker",
			msg: types.MsgCreatePerpetual{
				Authority: validAuthority,
				Params: types.PerpetualParams{
					Ticker:     "",
					MarketType: types.PerpetualMarketType_PERPETUAL_MARKET_TYPE_CROSS,
				},
			},
			expectedErr: "Ticker must be non-empty string",
		},
		{
			desc: "DefaultFundingPpm >= MaxDefaultFundingPpmAbs",
			msg: types.MsgCreatePerpetual{
				Authority: validAuthority,
				Params: types.PerpetualParams{
					Ticker:            "test",
					DefaultFundingPpm: 100_000_000,
					MarketType:        types.PerpetualMarketType_PERPETUAL_MARKET_TYPE_CROSS,
				},
			},
			expectedErr: "DefaultFundingPpm magnitude exceeds maximum value",
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			require.ErrorContains(t, err, tc.expectedErr)
		})
	}
}