package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitWitdrawProposal = "submit_witdraw_proposal"

var _ sdk.Msg = &MsgSubmitWitdrawProposal{}

func NewMsgSubmitWitdrawProposal(creator string, poolIndex string) *MsgSubmitWitdrawProposal {
  return &MsgSubmitWitdrawProposal{
		Creator: creator,
    PoolIndex: poolIndex,
	}
}

func (msg *MsgSubmitWitdrawProposal) Route() string {
  return RouterKey
}

func (msg *MsgSubmitWitdrawProposal) Type() string {
  return TypeMsgSubmitWitdrawProposal
}

func (msg *MsgSubmitWitdrawProposal) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitWitdrawProposal) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitWitdrawProposal) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

