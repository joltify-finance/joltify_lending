package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/vault module sentinel errors
var (
	ErrFormat = errorsmod.Register(ModuleName, 1, "format convert error")
	ErrUpdate = errorsmod.Register(ModuleName, 2, "error in update the validtor")
	ErrPool   = errorsmod.Register(ModuleName, 3, "fail to get two pools")
	// this line is used by starport scaffolding # ibc/errors
)
