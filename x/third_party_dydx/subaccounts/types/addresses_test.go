package types_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/subaccounts/types"
	"github.com/stretchr/testify/require"
)

func TestModuleAddress(t *testing.T) {
	require.Equal(t, "jolt14jux2kfgelja5394dxquqn3wh974psqzv4hzgg", types.ModuleAddress.String())
}
