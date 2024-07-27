package keeper

import (
	"context"

	v17 "github.com/joltify-finance/joltify_lending/x/vault/migrations/v17"
	v5 "github.com/joltify-finance/joltify_lending/x/vault/migrations/v5"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

// Migrate2to3 from version 2 to 3.
func (m Migrator) Migrate2to3(ctx context.Context) error {
	return v17.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}

// Migrate3to4 from version3 to 4.
func (m Migrator) Migrate3to4(ctx context.Context) error {
	return v5.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}
