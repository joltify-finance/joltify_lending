package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, *keeper.Keeper, types.NFTKeeper, types.BankKeeper, context.Context) {
	k, nftType, bankKeeper, ctx := keepertest.SpvKeeper(t)
	return keeper.NewMsgServerImpl(*k), k, nftType, bankKeeper, sdk.WrapSDKContext(ctx)
}
