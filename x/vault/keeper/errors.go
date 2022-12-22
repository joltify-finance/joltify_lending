package keeper

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var ErrSuspend = sdkerrors.Register("oppyChain", 1, "bridge transfer suspended")
