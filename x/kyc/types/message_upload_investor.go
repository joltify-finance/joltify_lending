package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUploadInvestor = "upload_investor"

var _ sdk.Msg = &MsgUploadInvestor{}

func NewMsgUploadInvestor(creator string, investorId string, walletAddress []string) *MsgUploadInvestor {
	return &MsgUploadInvestor{
		Creator:       creator,
		InvestorId:    investorId,
		WalletAddress: walletAddress,
	}
}

func (msg *MsgUploadInvestor) Route() string {
	return RouterKey
}

func (msg *MsgUploadInvestor) Type() string {
	return TypeMsgUploadInvestor
}

func (msg *MsgUploadInvestor) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUploadInvestor) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUploadInvestor) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(errorsmod.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
