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

// Migrate1to2 from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v16.MigrateStore(sdk.UnwrapSDKContext(ctx), m.keeper.paramstore, m.keeper.storeKey, m.keeper.cdc)
}
