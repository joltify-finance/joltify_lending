package module_test

import (
	"encoding/json"
	"os"
	"testing"

	appconfig "github.com/joltify-finance/joltify_lending/app/config"
	"github.com/joltify-finance/joltify_lending/dydx_helper/module"
	"github.com/stretchr/testify/require"
)

func TestDefaultGenesis(t *testing.T) {
	encodingConfig := appconfig.MakeEncodingConfig()

	defaultGenesis := module.SlashingModuleBasic{}.DefaultGenesis(encodingConfig.Codec)
	humanReadableDefaultGenesisState, unmarshalErr := json.Marshal(&defaultGenesis)

	expectedDefaultGenesisState, fileReadErr := os.ReadFile("testdata/slashing_default_genesis_state.json")

	require.NoError(t, unmarshalErr)
	require.NoError(t, fileReadErr)
	require.JSONEq(t,
		string(expectedDefaultGenesisState), string(humanReadableDefaultGenesisState))
}
