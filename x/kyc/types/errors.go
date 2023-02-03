package types

import (
	sdkerrors "cosmossdk.io/errors"
)

// DONTCOVER

// x/kyc module sentinel errors
var (
	ErrUnauthorised        = sdkerrors.Register(ModuleName, 1, "unauthorised submitter")
	ErrExceedMaxWalletsNum = sdkerrors.Register(ModuleName, 2, "wallets number exceed max allowed")
	ErrInvalidWallets      = sdkerrors.Register(ModuleName, 3, "wallets address are invalid")
)
