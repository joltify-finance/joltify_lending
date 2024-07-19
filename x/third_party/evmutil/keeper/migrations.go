package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v2 "github.com/joltify-finance/joltify_lending/x/third_party/evmutil/migrations/v2"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator {
	return Migrator{
		keeper: keeper,
	}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx context.Context) error {
	return v2.MigrateStore(ctx, m.keeper.paramSubspace)
}
