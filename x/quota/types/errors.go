package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/quota module sentinel errors
var (
	AccErrQuotaExceed = errorsmod.Register(ModuleName, 1100, "Account Quota Exceed")
	ErrQuotaExceed    = errorsmod.Register(ModuleName, 1101, "Module Quota Exceed")
)
