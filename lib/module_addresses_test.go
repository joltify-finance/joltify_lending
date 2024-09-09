package lib_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/lib"
	"github.com/stretchr/testify/require"
)

func TestGovModuleAddress(t *testing.T) {
	require.Equal(t, "jolt10d07y265gmmuvt4z0w9aw880jnsr700jszwe96", lib.GovModuleAddress.String())
}
