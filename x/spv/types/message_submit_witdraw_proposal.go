package types

import (
	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitWitdrawProposal = "submit_withdraw_proposal"

var _ sdk.Msg = &MsgSubmitWithdrawProposal{}

func NewMsgSubmitWitdrawProposal(creator string, poolIndex string) *MsgSubmitWithdrawProposal {
	return &MsgSubmitWithdrawProposal{
		Creator:   creator,
		PoolIndex: poolIndex,
	}
}

func (msg *MsgSubmitWithdrawProposal) Route() string {
	return RouterKey
}

func (msg *MsgSubmitWithdrawProposal) Type() string {
	return TypeMsgSubmitWitdrawProposal
}

func (msg *MsgSubmitWithdrawProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitWithdrawProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitWithdrawProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
