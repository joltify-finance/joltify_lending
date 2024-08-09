package types

import "context"

// MultiHARDHooks combine multiple HARD hooks, all hook functions are run in array sequence
type MultiHARDHooks []JOLTHooks

// NewMultiJoltHooks returns a new MultiHARDHooks
func NewMultiJoltHooks(hooks ...JOLTHooks) MultiHARDHooks {
	return hooks
}

// AfterDepositCreated runs after a deposit is created
func (h MultiHARDHooks) AfterDepositCreated(ctx context.Context, deposit Deposit) {
	for i := range h {
		h[i].AfterDepositCreated(ctx, deposit)
	}
}

// BeforeDepositModified runs before a deposit is modified
func (h MultiHARDHooks) BeforeDepositModified(ctx context.Context, deposit Deposit) {
	for i := range h {
		h[i].BeforeDepositModified(ctx, deposit)
	}
}

// AfterDepositModified runs after a deposit is modified
func (h MultiHARDHooks) AfterDepositModified(ctx context.Context, deposit Deposit) {
	for i := range h {
		h[i].AfterDepositModified(ctx, deposit)
	}
}

// AfterBorrowCreated runs after a borrow is created
func (h MultiHARDHooks) AfterBorrowCreated(ctx context.Context, borrow Borrow) {
	for i := range h {
		h[i].AfterBorrowCreated(ctx, borrow)
	}
}

// BeforeBorrowModified runs before a borrow is modified
func (h MultiHARDHooks) BeforeBorrowModified(ctx context.Context, borrow Borrow) {
	for i := range h {
		h[i].BeforeBorrowModified(ctx, borrow)
	}
}

// AfterBorrowModified runs after a borrow is modified
func (h MultiHARDHooks) AfterBorrowModified(ctx context.Context, borrow Borrow) {
	for i := range h {
		h[i].AfterBorrowModified(ctx, borrow)
	}
}
