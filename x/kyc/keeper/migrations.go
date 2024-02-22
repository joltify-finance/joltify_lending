package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v16 "github.com/joltify-finance/joltify_lending/x/kyc/migrate"
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
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v16.MigrateStore(ctx, m.keeper.paramstore, m.keeper.storeKey, m.keeper.cdc)
}
