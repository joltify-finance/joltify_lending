package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/oracle/exported"
	v2 "github.com/joltify-finance/joltify_lending/x/third_party/oracle/migrations/v2"
)

type Migrator struct {
	keeper   Keeper
	subspace exported.Subspace
}

func NewMigrator(k Keeper, ss exported.Subspace) Migrator {
	return Migrator{
		keeper:   k,
		subspace: ss,
	}
}

// Migrate1to2 migrates oracle's consensus version from 1 to 2. Specifically, it migrates
// Params kept in x/params directly to oracle's module state
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v2.Migrate(
		ctx,
		ctx.KVStore(m.keeper.storeKey),
		m.subspace,
		m.keeper.cdc,
	)
}
