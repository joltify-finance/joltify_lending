package keeper_test

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	app2 "github.com/joltify-finance/joltify_lending/app"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/testutil/nullify"
	"github.com/joltify-finance/joltify_lending/x/vault/keeper"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNOutboundTx(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.OutboundTx {
	items := make([]types.OutboundTx, n)

	r := rand.New(rand.NewSource(time.Now().Unix())) //nolint:gosec
	accs := simulation.RandomAccounts(r, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].Feecoin = []sdk.Coin{{Denom: "mock", Amount: sdk.NewInt(1)}}
		items[i].ReceiverAddress = accs[i].Address
		keeper.SetOutboundTx(ctx, items[i])
	}
	return items
}

func TestOutboundTxGet(t *testing.T) {
	app2.SetSDKConfig()

	app, ctx := keepertest.SetupVaultApp(t)
	items := createNOutboundTx(&app.VaultKeeper, ctx, 10)
	for _, item := range items {
		rst, found := app.VaultKeeper.GetOutboundTx(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestOutboundTxGetAll(t *testing.T) {
	app, ctx := keepertest.SetupVaultApp(t)
	items := createNOutboundTx(&app.VaultKeeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(app.VaultKeeper.GetAllOutboundTx(ctx)),
	)
}
