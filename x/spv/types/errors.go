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
	ErrPoolNotActive      = coserrors.Register(ModuleName, 4, "pool is not active")
	ErrPoolExisted        = coserrors.Register(ModuleName, 5, "pool existed")
	ErrInconsistencyToken = coserrors.Register(ModuleName, 6, "inconsistency tokens")
	ErrInsufficientFund   = coserrors.Register(ModuleName, 7, "insufficient tokens")
	ErrPoolFull           = coserrors.Register(ModuleName, 8, "pool is full")
	ErrMINTNFT            = coserrors.Register(ModuleName, 9, "fail to mint nft")
	ErrINCONSISTENCYDENOM = coserrors.Register(ModuleName, 10, "denom inconsistency")
	ErrDepositorNotFound  = coserrors.Register(ModuleName, 11, "depositor not found")
	ErrUNEXPECTEDSTATUS   = coserrors.Register(ModuleName, 12, "unexpected pool status")
	ErrClassNotFound      = coserrors.Register(ModuleName, 13, "class cannot be found")
	ErrTransferNFT        = coserrors.Register(ModuleName, 14, "fail to transfer nft ownership")
	ErrBurnNFT            = coserrors.Register(ModuleName, 15, "fail to burn the nft")
	ErrNFTNotFound        = coserrors.Register(ModuleName, 16, "nft not found")
	ErrTime               = coserrors.Register(ModuleName, 17, "invalid time stamp")
	ErrDeposit            = coserrors.Register(ModuleName, 18, "fail to deposit")
	ErrClaimInterest      = coserrors.Register(ModuleName, 19, "fail to claim interest")
)
