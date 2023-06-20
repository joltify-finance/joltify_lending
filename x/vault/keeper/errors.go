package keeper

import (
	errorsmod "cosmossdk.io/errors"
)

var ErrSuspend = errorsmod.Register("JoltifyChain", 1, "bridge transfer suspended as quota reached")
