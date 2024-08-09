package types_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/perpetuals/types"
	"github.com/stretchr/testify/require"
)

func TestModuleKeys(t *testing.T) {
	require.Equal(t, "perpetuals", types.ModuleName)
	require.Equal(t, "perpetuals", types.StoreKey)
}

func TestStateKeys(t *testing.T) {
	require.Equal(t, "Perp:", types.PerpetualKeyPrefix)
	require.Equal(t, "PremVotes", types.PremiumVotesKey)
	require.Equal(t, "PremSamples", types.PremiumSamplesKey)
	require.Equal(t, "LiqTier:", types.LiquidityTierKeyPrefix)
	require.Equal(t, "Params", types.ParamsKey)
}

func TestModuleAccountKeys(t *testing.T) {
	require.Equal(t, "insurance_fund", types.InsuranceFundName)
}
