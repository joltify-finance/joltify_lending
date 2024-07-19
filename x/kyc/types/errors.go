package types

import (
	errorsmod "cosmossdk.io/errors"
)

// DONTCOVER

// x/kyc module sentinel errors
var (
	ErrUnauthorised        = errorsmod.Register(ModuleName, 1, "unauthorised submitter")
	ErrExceedMaxWalletsNum = errorsmod.Register(ModuleName, 2, "wallets number exceed max allowed")
	ErrInvalidWallets      = errorsmod.Register(ModuleName, 3, "wallets address are invalid")
	ErrInvalidProject      = errorsmod.Register(ModuleName, 4, "project is invalid")
)
