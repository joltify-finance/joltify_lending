package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/joltify-finance/joltify_lending/testutil/sample"
)

func TestMsgTransferOwnership_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgTransferOwnership
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgTransferOwnership{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgTransferOwnership{
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
