package types

import "errors"

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AllCoinsQuota: []CoinsQuota{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in quota
	coinsQuota := gs.AllCoinsQuota
	for _, el := range coinsQuota {
		if el.ModuleName == "" {
			return errors.New("invalid module name")
		}
		if el.CoinsSum.IsZero() {
			return errors.New("invalid quota sum")
		}

		for _, ee := range el.History {
			if ee.Amount.IsZero() {
				return errors.New("invalid quota amount")
			}
			if ee.BlockHeight < 1 {
				return errors.New("invalid block height")
			}
		}

	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
