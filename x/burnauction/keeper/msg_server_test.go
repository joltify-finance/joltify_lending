package keeper_test

import (
	"context"
	"testing"

	"github.com/joltify-finance/joltify_lending/x/burnauction/keeper"
	"github.com/joltify-finance/joltify_lending/x/burnauction/types"

	keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BurnauctionKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}
