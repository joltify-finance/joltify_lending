package types_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/prices/types"
	"github.com/stretchr/testify/require"
)

func TestModuleKeys(t *testing.T) {
	require.Equal(t, "prices", types.ModuleName)
	require.Equal(t, "prices", types.StoreKey)
}

func TestStateKeys(t *testing.T) {
	require.Equal(t, "Param:", types.MarketParamKeyPrefix)
	require.Equal(t, "Price:", types.MarketPriceKeyPrefix)
}
