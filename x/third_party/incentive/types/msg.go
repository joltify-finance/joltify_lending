package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const MaxDenomsToClaim = 1000

// ensure Msg interface compliance at compile time

var _ sdk.Msg = &MsgClaimJoltReward{}

// NewMsgClaimJoltReward returns a new MsgClaimJoltReward.
func NewMsgClaimJoltReward(sender string, denomsToClaim Selections) MsgClaimJoltReward {
	return MsgClaimJoltReward{
		Sender:        sender,
		DenomsToClaim: denomsToClaim,
	}
}

// Route return the message type used for routing the message.
func (msg MsgClaimJoltReward) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgClaimJoltReward) Type() string {
	return "claim_jolt_reward"
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgClaimJoltReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty or invalid")
	}
	if err := msg.DenomsToClaim.Validate(); err != nil {
		return err
	}
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgClaimJoltReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgClaimJoltReward) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}
