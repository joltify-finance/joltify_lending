package types

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	kyctypes "github.com/joltify-finance/joltify_lending/x/kyc/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	GetModuleAccount(ctx sdk.Context, name string) types.ModuleAccountI
	GetModuleAddress(name string) sdk.AccAddress
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected bank keeper
type BankKeeper interface {
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error

	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	BurnCoins(ctx sdk.Context, name string, amt sdk.Coins) error
}

type KycKeeper interface {
	GetProjects(ctx sdk.Context) (projectsInfo []*kyctypes.ProjectInfo)
	QueryInvestorWallets(goCtx context.Context, req *kyctypes.QueryInvestorWalletsRequest) (*kyctypes.QueryInvestorWalletsResponse, error)
}

type NFTKeeper interface {
	Mint(ctx sdk.Context, nft nfttypes.NFT, receiver sdk.AccAddress) error
	SaveClass(ctx sdk.Context, class nfttypes.Class) error
	GetClass(ctx sdk.Context, classID string) (nfttypes.Class, bool)
}
