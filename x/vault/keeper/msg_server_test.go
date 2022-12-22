package keeper_test

import (
	"context"
	"testing"

	joltapp "github.com/joltify-finance/joltify_lending/app"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/x/vault/keeper"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
)

func setupMsgServer(t testing.TB) (*joltapp.TestApp, types.MsgServer, context.Context) {
	app, ctx := keepertest.SetupVaultApp(t)
	return app, keeper.NewMsgServerImpl(app.VaultKeeper), sdk.WrapSDKContext(ctx)
}
