package keeper

import (
	"context"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	spv "github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"
	swap "github.com/joltify-finance/joltify_lending/x/third_party/swap/types"
)

// Hooks wrapper struct for hooks
type Hooks struct {
	k Keeper
}

// _ cdptypes.CDPHooks   = Hooks{}
var _ types.JOLTHooks = Hooks{}

var _ spv.SPVHooks = Hooks{}

var _ swap.SwapHooks = Hooks{}

// Hooks create new incentive hooks
func (k Keeper) Hooks() Hooks { return Hooks{k} }

// ------------------- Jolt Module Hooks -------------------

// AfterDepositCreated function that runs after a deposit is created
func (h Hooks) AfterDepositCreated(rctx context.Context, deposit types.Deposit) {
	ctx := sdk.UnwrapSDKContext(rctx)
	h.k.InitializeJoltSupplyReward(ctx, deposit)
}

// BeforeDepositModified function that runs before a deposit is modified
func (h Hooks) BeforeDepositModified(rctx context.Context, deposit types.Deposit) {
	ctx := sdk.UnwrapSDKContext(rctx)
	h.k.SynchronizeJoltSupplyReward(ctx, deposit)
}

// AfterDepositModified function that runs after a deposit is modified
func (h Hooks) AfterDepositModified(rctx context.Context, deposit types.Deposit) {
	ctx := sdk.UnwrapSDKContext(rctx)
	h.k.UpdateJoltSupplyIndexDenoms(ctx, deposit)
}

// AfterBorrowCreated function that runs after a borrow is created
func (h Hooks) AfterBorrowCreated(rctx context.Context, borrow types.Borrow) {
	ctx := sdk.UnwrapSDKContext(rctx)
	h.k.InitializeJoltBorrowReward(ctx, borrow)
}

// BeforeBorrowModified function that runs before a borrow is modified
func (h Hooks) BeforeBorrowModified(rctx context.Context, borrow types.Borrow) {
	ctx := sdk.UnwrapSDKContext(rctx)
	h.k.SynchronizeJoltBorrowReward(ctx, borrow)
}

// AfterBorrowModified function that runs after a borrow is modified
func (h Hooks) AfterBorrowModified(rctx context.Context, borrow types.Borrow) {
	ctx := sdk.UnwrapSDKContext(rctx)
	h.k.UpdateJoltBorrowIndexDenoms(ctx, borrow)
}

// ------------------- Swap Module Hooks -------------------

func (h Hooks) AfterPoolDepositCreated(rctx context.Context, poolID string, depositor sdk.AccAddress, _ sdkmath.Int) {
	ctx := sdk.UnwrapSDKContext(rctx)
	h.k.InitializeSwapReward(ctx, poolID, depositor)
}

func (h Hooks) BeforePoolDepositModified(rctx context.Context, poolID string, depositor sdk.AccAddress, sharesOwned sdkmath.Int) {
	ctx := sdk.UnwrapSDKContext(rctx)
	h.k.SynchronizeSwapReward(ctx, poolID, depositor, sharesOwned)
}

// ------------------- SPV Module Hooks -------------------

func (h Hooks) AfterSPVInterestPaid(rctx context.Context, poolID string, interestPaid sdkmath.Int) {
	ctx := sdk.UnwrapSDKContext(rctx)
	h.k.AfterSPVInterestPaid(ctx, poolID, interestPaid)
}

func (h Hooks) BeforeNFTBurned(rctx context.Context, poolIndex, investorId string, linkednfts []string) error {
	ctx := sdk.UnwrapSDKContext(rctx)
	return h.k.BeforeNFTBurn(ctx, poolIndex, investorId, linkednfts)
}
