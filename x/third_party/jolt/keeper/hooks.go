package keeper

import (
	"context"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"
)

// Implements StakingHooks interface
var _ types2.JOLTHooks = Keeper{}

// AfterDepositCreated - call hook if registered
func (k Keeper) AfterDepositCreated(ctx context.Context, deposit types2.Deposit) {
	if k.hooks != nil {
		k.hooks.AfterDepositCreated(ctx, deposit)
	}
}

// BeforeDepositModified - call hook if registered
func (k Keeper) BeforeDepositModified(ctx context.Context, deposit types2.Deposit) {
	if k.hooks != nil {
		k.hooks.BeforeDepositModified(ctx, deposit)
	}
}

// AfterDepositModified - call hook if registered
func (k Keeper) AfterDepositModified(ctx context.Context, deposit types2.Deposit) {
	if k.hooks != nil {
		k.hooks.AfterDepositModified(ctx, deposit)
	}
}

// AfterBorrowCreated - call hook if registered
func (k Keeper) AfterBorrowCreated(ctx context.Context, borrow types2.Borrow) {
	if k.hooks != nil {
		k.hooks.AfterBorrowCreated(ctx, borrow)
	}
}

// BeforeBorrowModified - call hook if registered
func (k Keeper) BeforeBorrowModified(ctx context.Context, borrow types2.Borrow) {
	if k.hooks != nil {
		k.hooks.BeforeBorrowModified(ctx, borrow)
	}
}

// AfterBorrowModified - call hook if registered
func (k Keeper) AfterBorrowModified(ctx context.Context, borrow types2.Borrow) {
	if k.hooks != nil {
		k.hooks.AfterBorrowModified(ctx, borrow)
	}
}
