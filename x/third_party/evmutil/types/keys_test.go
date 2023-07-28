package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/joltify-finance/joltify_lending/x/third_party/evmutil/types"
)

func TestDeployedCosmosCoinContractKey(t *testing.T) {
	denom := "magic"
	key := types.DeployedCosmosCoinContractKey(denom)
	require.Equal(t, key, append([]byte{0x01}, []byte(denom)...))
	require.Equal(t, denom, types.DenomFromDeployedCosmosCoinContractKey(key))
}
