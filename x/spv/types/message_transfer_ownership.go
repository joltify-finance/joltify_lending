package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTransferOwnership = "transfer_ownership"

var _ sdk.Msg = &MsgTransferOwnership{}

func NewMsgTransferOwnership(creator string, poolIndex string) *MsgTransferOwnership {
  return &MsgTransferOwnership{
		Creator: creator,
    PoolIndex: poolIndex,
	}
}

func (msg *MsgTransferOwnership) Route() string {
  return RouterKey
}

func (msg *MsgTransferOwnership) Type() string {
  return TypeMsgTransferOwnership
}

func (msg *MsgTransferOwnership) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgTransferOwnership) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgTransferOwnership) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

