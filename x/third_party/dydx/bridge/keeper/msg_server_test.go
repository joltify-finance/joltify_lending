package keeper_test

import (
	"context"
	"testing"

	testapp "github.com/dydxprotocol/v4-chain/protocol/testutil/app"
	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/bridge/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/bridge/types"
)

func setupMsgServer(t *testing.T) (keeper.Keeper, types.MsgServer, context.Context) {
	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()
	k := tApp.App.BridgeKeeper

	return k, keeper.NewMsgServerImpl(k), ctx
}
