package types_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/subaccounts/types"
	"github.com/stretchr/testify/require"
)

func TestModuleAddress(t *testing.T) {
	require.Equal(t, "dydx1v88c3xv9xyv3eetdx0tvcmq7ung3dywp5upwc6", types.ModuleAddress.String())
}
