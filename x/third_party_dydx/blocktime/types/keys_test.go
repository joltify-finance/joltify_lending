package types_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/blocktime/types"
	"github.com/stretchr/testify/require"
)

func TestModuleKeys(t *testing.T) {
	require.Equal(t, "blocktime", types.ModuleName)
	require.Equal(t, "blocktime", types.StoreKey)
}

func TestStateKeys(t *testing.T) {
	require.Equal(t, "DowntimeParams", types.DowntimeParamsKey)
	require.Equal(t, "AllDowntimeInfo", types.AllDowntimeInfoKey)
	require.Equal(t, "PreviousBlockInfo", types.PreviousBlockInfoKey)
}
