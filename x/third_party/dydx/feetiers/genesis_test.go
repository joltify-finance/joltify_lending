package feetiers_test

import (
	"testing"

	testapp "github.com/dydxprotocol/v4-chain/protocol/testutil/app"
	feetiers "github.com/joltify-finance/joltify_lending/x/third_party/dydx/feetiers"
	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/feetiers/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()
	got := feetiers.ExportGenesis(ctx, *tApp.App.FeeTiersKeeper)
	require.NotNil(t, got)
	require.Equal(t, types.DefaultGenesis(), got)
}
