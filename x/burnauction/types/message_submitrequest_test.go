package types

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/testutil/sample"

	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgSubmitrequest_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSubmitrequest
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSubmitrequest{
				Creator: "invalid_address",
			},
			err: errorsmod.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSubmitrequest{
				Creator: sample.AccAddress(),
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
