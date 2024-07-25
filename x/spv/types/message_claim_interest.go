package types

import (
	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgClaimInterest = "claim_interest"

var _ sdk.Msg = &MsgClaimInterest{}

func NewMsgClaimInterest(creator string, poolIndex string) *MsgClaimInterest {
	return &MsgClaimInterest{
		Creator:   creator,
		PoolIndex: poolIndex,
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
		return coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
