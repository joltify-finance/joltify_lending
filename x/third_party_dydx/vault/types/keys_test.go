package types_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/vault/types"
	"github.com/stretchr/testify/require"
)

func TestModuleKeys(t *testing.T) {
	require.Equal(t, "vault", types.ModuleName)
	require.Equal(t, "vault", types.StoreKey)
}

func TestStateKeys(t *testing.T) {
	require.Equal(t, "TotalShares:", types.TotalSharesKeyPrefix)
}
