package types

import (
	"testing"

	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgUpdatePool_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdatePool
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdatePool{
				Creator: "invalid_address",
			},
			err: errorsmod.ErrInvalidAddress,
		},
		{
			name: "valid address",
			msg: MsgUpdatePool{
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
