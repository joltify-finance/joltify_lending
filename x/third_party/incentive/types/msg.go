package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

const MaxDenomsToClaim = 1000

// ensure Msg interface compliance at compile time

var (
	_ sdk.Msg = &MsgClaimJoltReward{}
	_ sdk.Msg = &MsgClaimSwapReward{}
	_ sdk.Msg = &MsgClaimSPVReward{}

	_ legacytx.LegacyMsg = &MsgClaimSwapReward{}
	_ legacytx.LegacyMsg = &MsgClaimJoltReward{}
	_ legacytx.LegacyMsg = &MsgClaimSPVReward{}
)

const (
	TypeMsgClaimHardReward = "claim_jolt_reward"
	TypeMsgClaimSwapReward = "claim_swap_reward"
	TypeMsgClaimSPVReward  = "claim_spv_reward"
)

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
	return TypeMsgClaimHardReward
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgClaimJoltReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty or invalid")
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

// NewMsgClaimSwapReward returns a new MsgClaimSwapReward.
func NewMsgClaimSwapReward(sender string, denomsToClaim Selections) MsgClaimSwapReward {
	return MsgClaimSwapReward{
		Sender:        sender,
		DenomsToClaim: denomsToClaim,
	}
}

// Route return the message type used for routing the message.
func (msg MsgClaimSwapReward) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgClaimSwapReward) Type() string {
	return TypeMsgClaimSwapReward
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgClaimSwapReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty or invalid")
	}
	if err := msg.DenomsToClaim.Validate(); err != nil {
		return err
	}
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgClaimSwapReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgClaimSwapReward) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// NewMsgClaimSPVReward returns a new MsgClaimSwapReward.
func NewMsgClaimSPVReward(sender string, poolIndex string) MsgClaimSPVReward {
	return MsgClaimSPVReward{
		Sender:    sender,
		PoolIndex: poolIndex,
	}
}

// Route return the message type used for routing the message.
func (msg MsgClaimSPVReward) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgClaimSPVReward) Type() string {
	return TypeMsgClaimSwapReward
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgClaimSPVReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty or invalid")
	}
	if msg.PoolIndex == "" {
		return errorsmod.Wrap(errorsmod.ErrInvalidRequest, "pool index cannot be empty")
	}
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgClaimSPVReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgClaimSPVReward) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}
