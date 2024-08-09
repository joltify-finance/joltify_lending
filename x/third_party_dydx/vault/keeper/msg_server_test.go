package keeper_test

import (
	"context"
	"testing"

	testapp "github.com/joltify-finance/joltify_lending/dydx_helper/testutil/app"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/vault/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/vault/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t *testing.T) (keeper.Keeper, types.MsgServer, context.Context) {
	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()
	k := tApp.App.VaultKeeper

	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, k)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}
