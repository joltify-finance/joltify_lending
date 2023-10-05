package keeper_test

import (
	"testing"

	app2 "github.com/joltify-finance/joltify_lending/app"
	keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"
)

func TestUpateStakingInfo(t *testing.T) {
	app2.SetSDKConfig()

	app, ctx := keepertest.SetupVaultApp(t)
	k := app.VaultKeeper
	k.UpdateStakingInfo(ctx)
}
