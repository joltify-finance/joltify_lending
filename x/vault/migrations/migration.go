package migrations

import (
	vaultmodulekeeper "github.com/joltify-finance/joltify_lending/x/vault/keeper"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper vaultmodulekeeper.Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper vaultmodulekeeper.Keeper) Migrator {
	return Migrator{keeper: keeper}
}
