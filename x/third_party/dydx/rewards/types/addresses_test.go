package types_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/rewards/types"
	"github.com/stretchr/testify/require"
)

func TestTreasuryModuleAddress(t *testing.T) {
	require.Equal(t, "dydx16wrau2x4tsg033xfrrdpae6kxfn9kyuerr5jjp", types.TreasuryModuleAddress.String())
}
