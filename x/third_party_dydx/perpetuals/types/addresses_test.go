package types_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/perpetuals/types"
	"github.com/stretchr/testify/require"
)

func TestInsuranceFundModuleAddress(t *testing.T) {
	require.Equal(t, "jolt1c7ptc87hkd54e3r7zjy92q29xkq7t79wevr8s7", types.InsuranceFundModuleAddress.String())
}
