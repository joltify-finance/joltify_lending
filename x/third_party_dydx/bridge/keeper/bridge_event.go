package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/bridge/types"
)

// `GetBridgeEventFromServer` returns the bridge event with the given id from the server. `found` is false
// if the event is not found.
func (k Keeper) GetBridgeEventFromServer(ctx sdk.Context, id uint32) (event types.BridgeEvent, found bool) {
	event, _, found = k.bridgeEventManager.GetBridgeEventById(id)
	return event, found
}