package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
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

// SwapKeeper defines the required methods needed by this modules keeper
type SwapKeeper interface {
	GetPoolShares(ctx sdk.Context, poolID string) (shares sdkmath.Int, found bool)
	GetDepositorSharesAmount(ctx sdk.Context, depositor sdk.AccAddress, poolID string) (shares sdkmath.Int, found bool)
}

// SPVKeeper defines the required methods needed by this modules keeper
type SPVKeeper interface {
	AfterSPVInterestPaid(ctx sdk.Context, poolID string, interestPaid sdkmath.Int)
	GetPools(ctx sdk.Context, index string) (poolInfo types.PoolInfo, ok bool)
	GetDepositor(ctx sdk.Context, poolIndex string, walletAddress sdk.AccAddress) (depositor types.DepositorInfo, found bool)
}

type NFTKeeper interface {
	GetClass(ctx sdk.Context, classID string) (nfttypes.Class, bool)
	UpdateClass(ctx sdk.Context, class nfttypes.Class) error
	GetNFT(ctx sdk.Context, classID, nftID string) (nfttypes.NFT, bool)
	Update(ctx sdk.Context, token nfttypes.NFT) error
}
