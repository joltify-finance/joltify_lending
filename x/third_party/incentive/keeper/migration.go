package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v5 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/migrations/v5"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

// Migrate1to2 from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v5.MigrateStore(ctx, m.keeper.key, m.keeper.cdc)
}
