package vault_test

import (
	"testing"

	keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/testutil/nullify"
	"github.com/joltify-finance/joltify_lending/x/vault"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		OutboundTxList: []types.OutboundTx{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	tapp, ctx := keepertest.SetupVaultApp(t)
	vault.InitGenesis(ctx, tapp.App.VaultKeeper, genesisState)
	got := vault.ExportGenesis(ctx, tapp.App.VaultKeeper)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.OutboundTxList, got.OutboundTxList)
	// this line is used by starport scaffolding # genesis/test/assert
}
