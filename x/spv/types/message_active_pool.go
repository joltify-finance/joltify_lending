package types

import (
	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgActivePool = "active_pool"

var _ sdk.Msg = &MsgActivePool{}

func NewMsgActivePool(creator string, poolIndex string) *MsgActivePool {
	return &MsgActivePool{
		Creator:   creator,
		PoolIndex: poolIndex,
	}
}

func (msg *MsgActivePool) Route() string {
	return RouterKey
}

func (msg *MsgActivePool) Type() string {
	return TypeMsgActivePool
}

func (msg *MsgActivePool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgActivePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgActivePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
