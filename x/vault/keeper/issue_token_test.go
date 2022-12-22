package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/x/vault/keeper"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
	"github.com/stretchr/testify/assert"
)

func createNIssueToken(keeper *keeper.Keeper, ctx sdk.Context, n int) ([]types.IssueToken, error) {
	items := make([]types.IssueToken, n)
	creator, err := sdk.AccAddressFromBech32("jolt1rfmwldwrm3652shx3a7say0v4vvtglast0l05d")
	if err != nil {
		return nil, err
	}
	for i := range items {
		items[i].Creator = creator
		items[i].Index = fmt.Sprintf("%d", i)
		keeper.SetIssueToken(ctx, items[i])
	}
	return items, nil
}

func TestIssueTokenGet(t *testing.T) {
	app, ctx := keepertest.SetupVaultApp(t)
	items, err := createNIssueToken(&app.VaultKeeper, ctx, 10)
	assert.Nil(t, err)
	for _, item := range items {
		rst, found := app.VaultKeeper.GetIssueToken(ctx, item.Index)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}

func TestIssueTokenGetAll(t *testing.T) {
	app, ctx := keepertest.SetupVaultApp(t)
	items, err := createNIssueToken(&app.VaultKeeper, ctx, 10)
	assert.Nil(t, err)
	assert.Equal(t, items, app.VaultKeeper.GetAllIssueToken(ctx))
}
