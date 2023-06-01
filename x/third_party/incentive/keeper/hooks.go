package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	cdptypes "github.com/joltify-finance/joltify_lending/x/third_party/cdp/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"
)

// Hooks wrapper struct for hooks
type Hooks struct {
	k Keeper
}

// _ cdptypes.CDPHooks   = Hooks{}
var _ types.JOLTHooks = Hooks{}

// Hooks create new incentive hooks
func (k Keeper) Hooks() Hooks { return Hooks{k} }

// ------------------- Cdp Module Hooks -------------------

// AfterCDPCreated function that runs after a cdp is created
func (h Hooks) AfterCDPCreated(ctx sdk.Context, cdp cdptypes.CDP) {
	// todo we need to something here once we enable CDP
}

// BeforeCDPModified function that runs before a cdp is modified
// note that this is called immediately after interest is synchronized, and so could potentially
// be called AfterCDPInterestUpdated or something like that, if we we're to expand the scope of cdp hooks
func (h Hooks) BeforeCDPModified(ctx sdk.Context, cdp cdptypes.CDP) {
	// todo we need to something here once we enable CDP
}

// ------------------- Jolt Module Hooks -------------------

// AfterDepositCreated function that runs after a deposit is created
func (h Hooks) AfterDepositCreated(ctx sdk.Context, deposit types.Deposit) {
	h.k.InitializeJoltSupplyReward(ctx, deposit)
}

// BeforeDepositModified function that runs before a deposit is modified
func (h Hooks) BeforeDepositModified(ctx sdk.Context, deposit types.Deposit) {
	h.k.SynchronizeJoltSupplyReward(ctx, deposit)
}

// AfterDepositModified function that runs after a deposit is modified
func (h Hooks) AfterDepositModified(ctx sdk.Context, deposit types.Deposit) {
	h.k.UpdateJoltSupplyIndexDenoms(ctx, deposit)
}

// AfterBorrowCreated function that runs after a borrow is created
func (h Hooks) AfterBorrowCreated(ctx sdk.Context, borrow types.Borrow) {
	h.k.InitializeJoltBorrowReward(ctx, borrow)
}

// BeforeBorrowModified function that runs before a borrow is modified
func (h Hooks) BeforeBorrowModified(ctx sdk.Context, borrow types.Borrow) {
	h.k.SynchronizeJoltBorrowReward(ctx, borrow)
}

// AfterBorrowModified function that runs after a borrow is modified
func (h Hooks) AfterBorrowModified(ctx sdk.Context, borrow types.Borrow) {
	h.k.UpdateJoltBorrowIndexDenoms(ctx, borrow)
}
