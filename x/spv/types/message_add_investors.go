package types

import (
	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddInvestors = "add_investors"

var _ sdk.Msg = &MsgAddInvestors{}

func NewMsgAddInvestors(creator, poolIndex string, investorID []string) *MsgAddInvestors {
	return &MsgAddInvestors{
		Creator:    creator,
		PoolIndex:  poolIndex,
		InvestorID: investorID,
	}
}

func (msg *MsgAddInvestors) Route() string {
	return RouterKey
}

func (msg *MsgAddInvestors) Type() string {
	return TypeMsgAddInvestors
}

func (msg *MsgAddInvestors) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddInvestors) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddInvestors) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.InvestorID) == 0 {
		return coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "empty investors")
	}
	return nil
}
