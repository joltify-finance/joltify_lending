package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/joltify-finance/joltify_lending/x/spv/types"
    "github.com/joltify-finance/joltify_lending/x/spv/keeper"
    keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SpvKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
