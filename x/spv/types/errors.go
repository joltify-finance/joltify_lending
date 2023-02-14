package types

// DONTCOVER

import (
	coserrors "cosmossdk.io/errors"
)

// x/spv module sentinel errors
var (
	InvalidParameter = coserrors.Register(ModuleName, 1, "invalid parameter")
	PoolNotFound     = coserrors.Register(ModuleName, 2, "pool not found")
	Unauthorized     = coserrors.Register(ModuleName, 3, "unauthorized operation")
)
