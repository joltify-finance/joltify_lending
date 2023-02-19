package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgClaimInterest = "claim_interest"

var _ sdk.Msg = &MsgClaimInterest{}

func NewMsgClaimInterest(creator string, borrowAmount string) *MsgClaimInterest {
	return &MsgClaimInterest{
		Creator: creator,
		//BorrowAmount: borrowAmount,
	}
}

func (msg *MsgClaimInterest) Route() string {
	return RouterKey
}

func (msg *MsgClaimInterest) Type() string {
	return TypeMsgClaimInterest
}

func (msg *MsgClaimInterest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgClaimInterest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimInterest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
