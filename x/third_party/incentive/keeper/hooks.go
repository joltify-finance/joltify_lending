package keeper

import (
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
func (h Hooks) AfterDepositCreated(ctx context.Context, deposit types.Deposit) {
	h.k.InitializeJoltSupplyReward(ctx, deposit)
}

// BeforeDepositModified function that runs before a deposit is modified
func (h Hooks) BeforeDepositModified(ctx context.Context, deposit types.Deposit) {
	h.k.SynchronizeJoltSupplyReward(ctx, deposit)
}

// AfterDepositModified function that runs after a deposit is modified
func (h Hooks) AfterDepositModified(ctx context.Context, deposit types.Deposit) {
	h.k.UpdateJoltSupplyIndexDenoms(ctx, deposit)
}

// AfterBorrowCreated function that runs after a borrow is created
func (h Hooks) AfterBorrowCreated(ctx context.Context, borrow types.Borrow) {
	h.k.InitializeJoltBorrowReward(ctx, borrow)
}

// BeforeBorrowModified function that runs before a borrow is modified
func (h Hooks) BeforeBorrowModified(ctx context.Context, borrow types.Borrow) {
	h.k.SynchronizeJoltBorrowReward(ctx, borrow)
}

// AfterBorrowModified function that runs after a borrow is modified
func (h Hooks) AfterBorrowModified(ctx context.Context, borrow types.Borrow) {
	h.k.UpdateJoltBorrowIndexDenoms(ctx, borrow)
}

// ------------------- Swap Module Hooks -------------------

func (h Hooks) AfterPoolDepositCreated(ctx context.Context, poolID string, depositor sdk.AccAddress, _ sdkmath.Int) {
	h.k.InitializeSwapReward(ctx, poolID, depositor)
}

func (h Hooks) BeforePoolDepositModified(ctx context.Context, poolID string, depositor sdk.AccAddress, sharesOwned sdkmath.Int) {
	h.k.SynchronizeSwapReward(ctx, poolID, depositor, sharesOwned)
}

// ------------------- SPV Module Hooks -------------------

func (h Hooks) AfterSPVInterestPaid(ctx context.Context, poolID string, interestPaid sdkmath.Int) {
	h.k.AfterSPVInterestPaid(ctx, poolID, interestPaid)
}

func (h Hooks) BeforeNFTBurned(ctx context.Context, poolIndex, investorId string, linkednfts []string) error {
	return h.k.BeforeNFTBurn(ctx, poolIndex, investorId, linkednfts)
}
