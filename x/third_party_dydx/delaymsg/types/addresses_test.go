package types_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/delaymsg/types"
	"github.com/stretchr/testify/require"
)

func TestModuleAddress(t *testing.T) {
	require.Equal(t, "jolt1mkkvp26dngu6n8rmalaxyp3gwkjuzztqhm4zca", types.ModuleAddress.String())
}
