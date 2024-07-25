package types

import (
	"testing"

	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgAddInvestors_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddInvestors
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddInvestors{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgAddInvestors{
				Creator:    sample.AccAddress(),
				InvestorID: []string{"1", "2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
