package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPayPrincipalPartial = "pay_principal_partial"

var _ sdk.Msg = &MsgPayPrincipalPartial{}

func NewMsgPayPrincipalPartial(creator string, poolIndex string, token sdk.Coin) *MsgPayPrincipalPartial {
	return &MsgPayPrincipalPartial{
		Creator:   creator,
		PoolIndex: poolIndex,
		Token:     token,
	}
}

func (msg *MsgPayPrincipalPartial) Route() string {
	return RouterKey
}

func (msg *MsgPayPrincipalPartial) Type() string {
	return TypeMsgPayPrincipal
}

func (msg *MsgPayPrincipalPartial) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPayPrincipalPartial) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPayPrincipalPartial) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
