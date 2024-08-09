package types_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/govplus/types"
	"github.com/stretchr/testify/require"
)

func TestModuleKeys(t *testing.T) {
	require.Equal(t, "govplus", types.ModuleName)
	require.Equal(t, "dydxgovplus", types.StoreKey)
}
