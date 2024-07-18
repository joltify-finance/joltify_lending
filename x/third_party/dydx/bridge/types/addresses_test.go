package types_test

import (
	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/bridge/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestModuleAddress(t *testing.T) {
	require.Equal(t, "dydx1zlefkpe3g0vvm9a4h0jf9000lmqutlh9jwjnsv", types.ModuleAddress.String())
}
