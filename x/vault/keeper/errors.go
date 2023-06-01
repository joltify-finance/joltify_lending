package keeper

import (
	errorsmod "cosmossdk.io/errors"
)

var ErrSuspend = errorsmod.Register("oppyChain", 1, "bridge transfer suspended")
