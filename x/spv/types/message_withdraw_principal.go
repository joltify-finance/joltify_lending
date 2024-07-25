package types

import (
	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWithdrawPrincipal = "withdraw_principal"

var _ sdk.Msg = &MsgWithdrawPrincipal{}

func NewMsgWithdrawPrincipal(creator string, poolIndex string, token sdk.Coin) *MsgWithdrawPrincipal {
	return &MsgWithdrawPrincipal{
		Creator:   creator,
		PoolIndex: poolIndex,
		Token:     token,
	}
}

func (msg *MsgWithdrawPrincipal) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawPrincipal) Type() string {
	return TypeMsgWithdrawPrincipal
}

func (msg *MsgWithdrawPrincipal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdrawPrincipal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawPrincipal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
