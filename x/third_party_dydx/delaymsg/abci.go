package delaymsg

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/delaymsg/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/delaymsg/types"
)

// EndBlocker executes all ABCI EndBlock logic respective to the delaymsg module.
func EndBlocker(ctx sdk.Context, k types.DelayMsgKeeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)
	keeper.DispatchMessagesForBlock(k, ctx)
}
