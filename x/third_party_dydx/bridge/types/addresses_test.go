package types_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/bridge/types"
	"github.com/stretchr/testify/require"
)

func TestModuleAddress(t *testing.T) {
	require.Equal(t, "jolt12p5np79t6w6vfzc7q37uncxce2f4qtjqfny0nz", types.ModuleAddress.String())
}
