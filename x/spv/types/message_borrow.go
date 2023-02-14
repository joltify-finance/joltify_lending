package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBorrow = "borrow"

var _ sdk.Msg = &MsgBorrow{}

func NewMsgBorrow(creator, poolIndex string, borrowAmount sdk.Coin) *MsgBorrow {
	return &MsgBorrow{
		Creator:      creator,
		PoolIndex:    poolIndex,
		BorrowAmount: borrowAmount,
	}
}

func (msg *MsgBorrow) Route() string {
	return RouterKey
}

func (msg *MsgBorrow) Type() string {
	return TypeMsgBorrow
}

func (msg *MsgBorrow) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBorrow) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBorrow) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
