package types

import (
	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgLiquidate = "liquidate"

var _ sdk.Msg = &MsgLiquidate{}

func NewMsgLiquidate(creator string, poolIndex string) *MsgLiquidate {
	return &MsgLiquidate{
		Creator:   creator,
		PoolIndex: poolIndex,
	}
}

func (msg *MsgLiquidate) Route() string {
	return RouterKey
}

func (msg *MsgLiquidate) Type() string {
	return TypeMsgLiquidate
}

func (msg *MsgLiquidate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLiquidate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLiquidate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
