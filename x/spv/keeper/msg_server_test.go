package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, *keeper.Keeper, types.NFTKeeper, types.BankKeeper, keepertest.MockAuctionKeeper, context.Context) {
	k, nftType, bankKeeper, auctionKeeper, _, ctx := keepertest.SpvKeeper(t)
	k.SetParams(ctx, types.NewTestParams())
	return keeper.NewMsgServerImpl(*k), k, nftType, bankKeeper, auctionKeeper, ctx
}

func setupMsgServerWithIncentiveKeeper(t testing.TB) (types.MsgServer, *keeper.Keeper, types.NFTKeeper, types.BankKeeper, keepertest.MockAuctionKeeper, keepertest.FakeIncentiveKeeper, context.Context) {
	k, nftType, bankKeeper, auctionKeeper, incentiveKeeper, ctx := keepertest.SpvKeeper(t)
	k.SetParams(ctx, types.NewTestParams())
	return keeper.NewMsgServerImpl(*k), k, nftType, bankKeeper, auctionKeeper, incentiveKeeper, ctx
}
