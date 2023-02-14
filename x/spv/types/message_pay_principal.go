package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPayPrincipal = "pay_principal"

var _ sdk.Msg = &MsgPayPrincipal{}

func NewMsgPayPrincipal(creator string, poolIndex string, token sdk.Coin) *MsgPayPrincipal {
	return &MsgPayPrincipal{
		Creator:   creator,
		PoolIndex: poolIndex,
		Token:     token,
	}
}

func (msg *MsgPayPrincipal) Route() string {
	return RouterKey
}

func (msg *MsgPayPrincipal) Type() string {
	return TypeMsgPayPrincipal
}

func (msg *MsgPayPrincipal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPayPrincipal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPayPrincipal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
