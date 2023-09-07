package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	jolttypes "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"
)

// ParamSubspace defines the expected Subspace interfacace
type ParamSubspace interface {
	GetParamSet(sdk.Context, paramtypes.ParamSet)
	SetParamSet(sdk.Context, paramtypes.ParamSet)
	WithKeyTable(paramtypes.KeyTable) paramtypes.Subspace
	HasKeyTable() bool
}

// BankKeeper defines the expected interface needed to send coins
type BankKeeper interface {
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}

// JoltKeeper defines the expected jolt keeper for interacting with Jolt protocol
type JoltKeeper interface {
	GetDeposit(ctx sdk.Context, depositor sdk.AccAddress) (jolttypes.Deposit, bool)
	GetBorrow(ctx sdk.Context, borrower sdk.AccAddress) (jolttypes.Borrow, bool)

	GetSupplyInterestFactor(ctx sdk.Context, denom string) (sdk.Dec, bool)
	GetBorrowInterestFactor(ctx sdk.Context, denom string) (sdk.Dec, bool)
	GetBorrowedCoins(ctx sdk.Context) (coins sdk.Coins, found bool)
	GetSuppliedCoins(ctx sdk.Context) (coins sdk.Coins, found bool)
}

// AccountKeeper expected interface for the account keeper (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	SetAccount(ctx sdk.Context, acc authtypes.AccountI)
	GetModuleAccount(ctx sdk.Context, name string) authtypes.ModuleAccountI
}

// JOLTHooks event hooks for other keepers to run code in response to HARD modifications
type JOLTHooks interface {
	AfterDepositCreated(ctx sdk.Context, deposit jolttypes.Deposit)
	BeforeDepositModified(ctx sdk.Context, deposit jolttypes.Deposit)
	AfterDepositModified(ctx sdk.Context, deposit jolttypes.Deposit)
	AfterBorrowCreated(ctx sdk.Context, borrow jolttypes.Borrow)
	BeforeBorrowModified(ctx sdk.Context, borrow jolttypes.Borrow)
	AfterBorrowModified(ctx sdk.Context, deposit jolttypes.Deposit)
}
