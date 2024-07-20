package types

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	// Methods imported from account should be defined here
	GetModuleAccount(ctx context.Context, name string) sdk.ModuleAccountI
	GetModuleAddress(name string) sdk.AccAddress
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	// Methods imported from bank should be defined here
	GetAllBalances(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}

// AuctionKeeper expected interface for the auction keeper (noalias)
type AuctionKeeper interface {
	StartSurplusAuction(ctx context.Context, seller string, lot sdk.Coin, bidDenom string) (uint64, error)
}
