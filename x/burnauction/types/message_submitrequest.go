package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitrequest = "submitrequest"

var _ sdk.Msg = &MsgSubmitrequest{}

func NewMsgSubmitrequest(creator string, tokens string) *MsgSubmitrequest {
	return &MsgSubmitrequest{
		Creator: creator,
		Tokens:  tokens,
	}
}

func (msg *MsgSubmitrequest) Route() string {
	return RouterKey
}

func (msg *MsgSubmitrequest) Type() string {
	return TypeMsgSubmitrequest
}

func (msg *MsgSubmitrequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitrequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitrequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(errorsmod.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = sdk.ParseCoinsNormalized(msg.Tokens)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidCoin, "invalid tokens (%s)", err)
	}
	return nil
}
