package types

import (
	"encoding/base64"
	"errors"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/gogoproto/proto"
)

const TypeMsgCreateProject = "create_project"

var _ sdk.Msg = &MsgCreateProject{}

func NewMsgCreateProject(creator string, encodedProject string) *MsgCreateProject {
	return &MsgCreateProject{
		Creator:        creator,
		EncodedProject: encodedProject,
	}
}

func (msg *MsgCreateProject) Route() string {
	return RouterKey
}

func (msg *MsgCreateProject) Type() string {
	return TypeMsgCreateProject
}

func (msg *MsgCreateProject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateProject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func ValidateProject(p ProjectInfo) error {
	if p.PoolLockedSeconds < 0 {
		return errors.New("project time related setting cannot be negative")
	}

	if p.WithdrawRequestWindowSeconds < 0 {
		return errors.New("project time related setting cannot be negative")
	}
	if p.GraceTime.Seconds() < 0 {
		return errors.New("project time related setting cannot be negative")
	}

	if !p.MinBorrowAmount.IsPositive() {
		return errors.New("min borrow amount cannot be negative")
	}

	if !p.MinDepositAmount.IsPositive() {
		return errors.New("min deposit amount cannot be negative")
	}

	freq, err := strconv.ParseInt(p.PayFreq, 10, 64)
	if err != nil {
		return err
	}
	if freq < 0 {
		return errors.New("pay freq cannot be negative")
	}

	if p.ProjectLength%uint64(freq) != 0 {
		return errors.New("project length should be multiple of pay freq")
	}

	return nil
}

func (msg *MsgCreateProject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(errorsmod.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	out, err := base64.StdEncoding.DecodeString(msg.EncodedProject)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidProject, "fail to decode the project base64 string: %v", err)
	}

	var project ProjectInfo
	err = proto.Unmarshal(out, &project)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidProject, "fail to unmarshal the project: %v", err)
	}

	err = ValidateProject(project)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidProject, "invalid project: %v", err)
	}

	return nil
}
