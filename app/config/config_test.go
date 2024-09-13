package config_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app/config"
	"github.com/stretchr/testify/require"
)

func TestSetupConfig_SealsConfig(t *testing.T) {
	sdkConfig := sdk.GetConfig()

	// A successful set confirms the config is not yet sealed
	sdkConfig.SetPurpose(0)
	require.Equal(t, uint32(0), sdkConfig.GetPurpose(), "Expected purpose to match set value")

	// Should set default app values and seal the config
	config.SetupConfig()

	require.Panicsf(t, func() { sdkConfig.SetPurpose(0) }, "Expected config to be sealed after SetupConfig")
}

func TestSetAddressPrefixes(t *testing.T) {
	sdkConfig := sdk.GetConfig()

	require.Equal(t, "jolt", sdkConfig.GetBech32AccountAddrPrefix())
	require.Equal(t, "joltpub", sdkConfig.GetBech32AccountPubPrefix())

	require.Equal(t, "joltvaloper", sdkConfig.GetBech32ValidatorAddrPrefix())
	require.Equal(t, "joltvaloperpub", sdkConfig.GetBech32ValidatorPubPrefix())

	require.Equal(t, "joltvalcons", sdkConfig.GetBech32ConsensusAddrPrefix())
	require.Equal(t, "joltvalconspub", sdkConfig.GetBech32ConsensusPubPrefix())
}
