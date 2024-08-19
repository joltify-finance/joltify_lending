package keeper

import (
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/sending/types"
)

type msgServer struct {
	Keeper types.SendingKeeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper types.SendingKeeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}