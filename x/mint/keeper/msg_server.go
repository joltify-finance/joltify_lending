package keeper

import (
	"github.com/joltify-finance/joltify_lending/x/mint/types"
)

type msgServer struct {
	Keeper
}

var _ types.MsgServer = msgServer{}
