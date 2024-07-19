package types

import (
	"context"

	sdkmath "cosmossdk.io/math"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"

	nfttypes "cosmossdk.io/x/nft"
	sdk "github.com/cosmos/cosmos-sdk/types"
	kyctypes "github.com/joltify-finance/joltify_lending/x/kyc/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	GetModuleAccount(ctx context.Context, name string) sdk.ModuleAccountI
	GetModuleAddress(name string) sdk.AccAddress
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected bank keeper
type BankKeeper interface {
	SendCoinsFromModuleToModule(ctx context.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error

	GetSupply(ctx context.Context, denom string) sdk.Coin
	GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin
	GetAllBalances(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	BurnCoins(ctx context.Context, name string, amt sdk.Coins) error
}

type KycKeeper interface {
	GetProject(ctx context.Context, index int32) (val kyctypes.ProjectInfo, found bool)
	GetInvestorWallets(ctx context.Context, investorID string) (kyctypes.Investor, error)
	GetByWallet(ctx context.Context, wallet string) (kyctypes.Investor, error)
}

// AuctionKeeper expected interface for the auction keeper (noalias)
type AuctionKeeper interface {
	StartSurplusAuction(ctx context.Context, seller string, lot sdk.Coin, bidDenom string) (uint64, error)
}

type NFTKeeper interface {
	Mint(ctx context.Context, nft nfttypes.NFT, receiver sdk.AccAddress) error
	SaveClass(ctx context.Context, class nfttypes.Class) error
	UpdateClass(ctx context.Context, class nfttypes.Class) error
	GetClass(ctx context.Context, classID string) (nfttypes.Class, bool)
	GetNFT(ctx context.Context, classID, nftID string) (nfttypes.NFT, bool)
	GetOwner(ctx context.Context, classID string, nftID string) sdk.AccAddress
	Update(ctx context.Context, token nfttypes.NFT) error
	Burn(ctx context.Context, classID string, nftID string) error
	GetTotalSupply(ctx context.Context, classID string) uint64
}

type PriceFeedKeeper interface {
	GetCurrentPrice(ctx context.Context, marketID string) (types2.CurrentPrice, error)
}

// SPVHooks are event hooks called when the interest is paid to the SPV pool
type SPVHooks interface {
	AfterSPVInterestPaid(ctx context.Context, poolID string, interestPaid sdkmath.Int)
	BeforeNFTBurned(ctx context.Context, poolIndex, investorID string, linkednfts []string) error
}

type IncentiveKeeper interface {
	SetSPVRewardTokens(ctx context.Context, poolId string, rewardTokens sdk.Coins)
}
