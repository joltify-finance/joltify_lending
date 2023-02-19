package types

// DONTCOVER

import (
	coserrors "cosmossdk.io/errors"
)

// x/spv module sentinel errors
var (
	ErrInvalidParameter   = coserrors.Register(ModuleName, 1, "invalid parameter")
	ErrPoolNotFound       = coserrors.Register(ModuleName, 2, "pool not found")
	ErrUnauthorized       = coserrors.Register(ModuleName, 3, "unauthorized operation")
	ErrPoolClosed         = coserrors.Register(ModuleName, 4, "pool closed")
	ErrPoolExisted        = coserrors.Register(ModuleName, 5, "pool existed")
	ErrInconsistencyToken = coserrors.Register(ModuleName, 6, "inconsistency tokens")
	ErrInsufficientFund   = coserrors.Register(ModuleName, 7, "insufficient tokens")
	ErrPoolFull           = coserrors.Register(ModuleName, 8, "pool is full")
	ErrMINTNFT            = coserrors.Register(ModuleName, 9, "fail to mint nft")
	ErrINCONSISTENCYDENOM = coserrors.Register(ModuleName, 10, "denom inconsistency")
	ErrDepositorNotFound  = coserrors.Register(ModuleName, 11, "depositor not found")
)
