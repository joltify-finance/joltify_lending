package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, *keeper.Keeper, types.NFTKeeper, context.Context) {
	k, nftKeeper, ctx := keepertest.SpvKeeper(t)
	return keeper.NewMsgServerImpl(*k), k, nftKeeper, sdk.WrapSDKContext(ctx)
}
