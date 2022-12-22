package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"
)

// Implements StakingHooks interface
var _ types2.JOLTHooks = Keeper{}

// AfterDepositCreated - call hook if registered
func (k Keeper) AfterDepositCreated(ctx sdk.Context, deposit types2.Deposit) {
	if k.hooks != nil {
		k.hooks.AfterDepositCreated(ctx, deposit)
	}
}

// BeforeDepositModified - call hook if registered
func (k Keeper) BeforeDepositModified(ctx sdk.Context, deposit types2.Deposit) {
	if k.hooks != nil {
		k.hooks.BeforeDepositModified(ctx, deposit)
	}
}

// AfterDepositModified - call hook if registered
func (k Keeper) AfterDepositModified(ctx sdk.Context, deposit types2.Deposit) {
	if k.hooks != nil {
		k.hooks.AfterDepositModified(ctx, deposit)
	}
}

// AfterBorrowCreated - call hook if registered
func (k Keeper) AfterBorrowCreated(ctx sdk.Context, borrow types2.Borrow) {
	if k.hooks != nil {
		k.hooks.AfterBorrowCreated(ctx, borrow)
	}
}

// BeforeBorrowModified - call hook if registered
func (k Keeper) BeforeBorrowModified(ctx sdk.Context, borrow types2.Borrow) {
	if k.hooks != nil {
		k.hooks.BeforeBorrowModified(ctx, borrow)
	}
}

// AfterBorrowModified - call hook if registered
func (k Keeper) AfterBorrowModified(ctx sdk.Context, borrow types2.Borrow) {
	if k.hooks != nil {
		k.hooks.AfterBorrowModified(ctx, borrow)
	}
}
