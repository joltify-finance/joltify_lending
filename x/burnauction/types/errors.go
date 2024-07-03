package types

// DONTCOVER

import (
	coserrors "cosmossdk.io/errors"
)

// x/burnauction module sentinel errors
var (
	ErrInvalidCoin = coserrors.Register(ModuleName, 1, "invalid coins")
	ErrTransfer    = coserrors.Register(ModuleName, 2, "invalid transfer")
)
