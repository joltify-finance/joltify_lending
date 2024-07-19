package types

import (
	"context"

	sdkmath "cosmossdk.io/math"
	nfttypes "cosmossdk.io/x/nft"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	jolttypes "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"
)

// ParamSubspace defines the expected Subspace interfacace
type ParamSubspace interface {
	GetParamSet(context.Context, paramtypes.ParamSet)
	SetParamSet(context.Context, paramtypes.ParamSet)
	WithKeyTable(paramtypes.KeyTable) paramtypes.Subspace
	HasKeyTable() bool
}

// BankKeeper defines the expected interface needed to send coins
type BankKeeper interface {
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	GetAllBalances(ctx context.Context, addr sdk.AccAddress) sdk.Coins
}

// JoltKeeper defines the expected jolt keeper for interacting with Jolt protocol
type JoltKeeper interface {
	GetDeposit(ctx context.Context, depositor sdk.AccAddress) (jolttypes.Deposit, bool)
	GetBorrow(ctx context.Context, borrower sdk.AccAddress) (jolttypes.Borrow, bool)

	GetSupplyInterestFactor(ctx context.Context, denom string) (sdkmath.LegacyDec, bool)
	GetBorrowInterestFactor(ctx context.Context, denom string) (sdkmath.LegacyDec, bool)
	GetBorrowedCoins(ctx context.Context) (coins sdk.Coins, found bool)
	GetSuppliedCoins(ctx context.Context) (coins sdk.Coins, found bool)
}

// AccountKeeper expected interface for the account keeper (noalias)
type AccountKeeper interface {
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	SetAccount(ctx context.Context, acc sdk.AccountI)
	GetModuleAccount(ctx context.Context, name string) sdk.ModuleAccountI
}

// JOLTHooks event hooks for other keepers to run code in response to HARD modifications
type JOLTHooks interface {
	AfterDepositCreated(ctx context.Context, deposit jolttypes.Deposit)
	BeforeDepositModified(ctx context.Context, deposit jolttypes.Deposit)
	AfterDepositModified(ctx context.Context, deposit jolttypes.Deposit)
	AfterBorrowCreated(ctx context.Context, borrow jolttypes.Borrow)
	BeforeBorrowModified(ctx context.Context, borrow jolttypes.Borrow)
	AfterBorrowModified(ctx context.Context, deposit jolttypes.Deposit)
}

// SwapKeeper defines the required methods needed by this modules keeper
type SwapKeeper interface {
	GetPoolShares(ctx context.Context, poolID string) (shares sdkmath.Int, found bool)
	GetDepositorSharesAmount(ctx context.Context, depositor sdk.AccAddress, poolID string) (shares sdkmath.Int, found bool)
}

// SPVKeeper defines the required methods needed by this modules keeper
type SPVKeeper interface {
	AfterSPVInterestPaid(ctx context.Context, poolID string, interestPaid sdkmath.Int)
	GetPools(ctx context.Context, index string) (poolInfo types.PoolInfo, ok bool)
	GetDepositor(ctx context.Context, poolIndex string, walletAddress sdk.AccAddress) (depositor types.DepositorInfo, found bool)
}

type NFTKeeper interface {
	GetClass(ctx context.Context, classID string) (nfttypes.Class, bool)
	UpdateClass(ctx context.Context, class nfttypes.Class) error
	GetNFT(ctx context.Context, classID, nftID string) (nfttypes.NFT, bool)
	Update(ctx context.Context, token nfttypes.NFT) error
}
