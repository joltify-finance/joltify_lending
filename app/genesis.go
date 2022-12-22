package app

import (
	"encoding/json"

	"github.com/tendermint/spm/cosmoscmd"
)

// GenesisState represents the genesis state of the blockchain. It is a map from module names to module genesis states.
type GenesisState map[string]json.RawMessage

// NewDefaultGenesisState generates the default state for the application.
func NewDefaultGenesisState() GenesisState {
	encoding := cosmoscmd.MakeEncodingConfig(ModuleBasics)
	return ModuleBasics.DefaultGenesis(encoding.Marshaler)
}
