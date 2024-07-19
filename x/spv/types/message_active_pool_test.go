package types

import (
	"testing"

	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgActivePool_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgActivePool
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgActivePool{
				Creator: "invalid_address",
			},
			err: errorsmod.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgActivePool{
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
