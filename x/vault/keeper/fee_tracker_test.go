package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/stretchr/testify/assert"
)

func TestSetStoreFeeAmount(t *testing.T) {
	app, ctx := keepertest.SetupVaultApp(t)
	a := sdk.NewCoin("mock", sdkmath.NewInt(12))
	b := sdk.NewCoin("mock2", sdkmath.NewInt(22))

	fees := sdk.NewCoins(a, b)
	app.VaultKeeper.SetStoreFeeAmount(ctx, fees)

	feeGet, ok := app.VaultKeeper.GetFeeAmount(ctx, "mock")
	assert.Equal(t, true, ok)
	assert.True(t, feeGet.IsEqual(sdk.NewCoin("mock", sdkmath.NewInt(12))))

	feesGet := app.VaultKeeper.GetAllFeeAmount(ctx)
	assert.Equal(t, true, feesGet.IsEqual(fees))

	feesGet[0].Amount = sdkmath.NewInt(2222)
	app.VaultKeeper.SetStoreFeeAmount(ctx, feesGet)
	feesGet = app.VaultKeeper.GetAllFeeAmount(ctx)
	assert.Equal(t, true, feesGet[0].Amount.Equal(sdk.NewInt(2222)))
}
