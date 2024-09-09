package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	appconfig "github.com/joltify-finance/joltify_lending/app/config"
	"github.com/stretchr/testify/require"

	sdkmath "cosmossdk.io/math"

	"github.com/joltify-finance/joltify_lending/x/burnauction/types"

	testkeeper "github.com/joltify-finance/joltify_lending/testutil/keeper"
)

func TestAuction(t *testing.T) {
	appconfig.SetupConfig()
	k, bk, ac, ctx := testkeeper.BurnauctionKeeper(t)
	k.RunSurplusAuctions(ctx)
	acc := ac.GetModuleAccount(ctx, types.ModuleAccount)
	balance := bk.GetAllBalances(ctx, acc.GetAddress())
	require.True(t, balance.Empty())

	burntokens := sdk.NewCoins(
		sdk.NewCoin("afake", sdkmath.NewInt(100)),
		sdk.NewCoin("bfake", sdkmath.NewInt(200)),
		sdk.NewCoin("cfake", sdkmath.NewInt(300)),
	)

	k.SetParams(ctx, types.Params{
		BurnThreshold: burntokens,
	})

	tb := sdk.NewCoins(sdk.NewCoin("afake", sdkmath.NewInt(1)))
	sender := sdk.AccAddress([]byte("sender"))
	err := bk.SendCoinsFromAccountToModule(ctx, sender, types.ModuleAccount, tb)
	require.NoError(t, err)

	k.RunSurplusAuctions(ctx)
	balance = bk.GetAllBalances(ctx, acc.GetAddress())
	require.True(t, balance.Equal(tb))

	// we put more token
	tb = tb.Add(sdk.NewCoin("bfake", sdkmath.NewInt(1)))
	err = bk.SendCoinsFromAccountToModule(ctx, sender, types.ModuleAccount, sdk.NewCoins(sdk.NewCoin("bfake", sdkmath.NewInt(1))))
	require.NoError(t, err)
	k.RunSurplusAuctions(ctx)
	balance = bk.GetAllBalances(ctx, acc.GetAddress())
	require.True(t, balance.Equal(tb))

	// we put more token
	tb = tb.Add(sdk.NewCoin("afake", sdkmath.NewInt(99)))
	err = bk.SendCoinsFromAccountToModule(ctx, sender, types.ModuleAccount, sdk.NewCoins(sdk.NewCoin("afake", sdkmath.NewInt(99))))
	require.NoError(t, err)
	k.RunSurplusAuctions(ctx)
	balance = bk.GetAllBalances(ctx, acc.GetAddress())
	require.True(t, balance.Equal(sdk.NewCoins(sdk.NewCoin("bfake", sdkmath.NewInt(1)))))

	// we add c coin
	tb = tb.Add(sdk.NewCoin("cfake", sdkmath.NewInt(1)))
	err = bk.SendCoinsFromAccountToModule(ctx, sender, types.ModuleAccount, sdk.NewCoins(sdk.NewCoin("cfake", sdkmath.NewInt(1))))
	require.NoError(t, err)
	k.RunSurplusAuctions(ctx)
	balance = bk.GetAllBalances(ctx, acc.GetAddress())
	require.True(t, balance.Equal(sdk.NewCoins(sdk.NewCoin("bfake", sdkmath.NewInt(1)), sdk.NewCoin("cfake", sdkmath.NewInt(1)))))
	prebalance := balance

	// coin not in threshold
	err = bk.SendCoinsFromAccountToModule(ctx, sender, types.ModuleAccount, sdk.NewCoins(sdk.NewCoin("ffake", sdkmath.NewInt(1))))
	require.NoError(t, err)
	k.RunSurplusAuctions(ctx)
	balance = bk.GetAllBalances(ctx, acc.GetAddress())
	require.True(t, balance.Equal(prebalance.Add(sdk.NewCoin("ffake", sdkmath.NewInt(1)))))

	// empty balance
	err = bk.SendCoinsFromAccountToModule(ctx, sender, types.ModuleAccount, burntokens)
	require.NoError(t, err)
	k.RunSurplusAuctions(ctx)
	balance = bk.GetAllBalances(ctx, acc.GetAddress())
	require.True(t, balance.Equal(sdk.NewCoins(sdk.NewCoin("ffake", sdkmath.NewInt(1)))))
}
