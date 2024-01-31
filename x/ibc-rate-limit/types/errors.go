package types

import (
	errorsmod "cosmossdk.io/errors"
)

var ErrQuota = errorsmod.Register(ModuleName, 1, "quota reached")
