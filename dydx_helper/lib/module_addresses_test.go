package lib_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/dydx_helper/lib"
	"github.com/stretchr/testify/require"
)

func TestGovModuleAddress(t *testing.T) {
	require.Equal(t, "dydx10d07y265gmmuvt4z0w9aw880jnsr700jnmapky", lib.GovModuleAddress.String())
}
