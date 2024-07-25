package types

import (
	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRepayInterest = "repay_interest"

var _ sdk.Msg = &MsgRepayInterest{}

func NewMsgRepayInterest(creator string, poolIndex string, amount sdk.Coin) *MsgRepayInterest {
	return &MsgRepayInterest{
		Creator:   creator,
		PoolIndex: poolIndex,
		Token:     amount,
	}
}

func (msg *MsgRepayInterest) Route() string {
	return RouterKey
}

func (msg *MsgRepayInterest) Type() string {
	return TypeMsgRepayInterest
}

func (msg *MsgRepayInterest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRepayInterest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRepayInterest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
