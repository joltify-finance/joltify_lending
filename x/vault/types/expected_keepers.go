package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type VaultStaking interface {
	IterateLastValidators(context.Context, func(index int64, validator stakingtypes.ValidatorI) (stop bool))

	GetBondedValidatorsByPower(ctx context.Context) []stakingtypes.Validator

	GetParams(ctx context.Context) stakingtypes.Params

	LastValidatorsIterator(ctx context.Context) (iterator sdk.Iterator)

	GetValidator(ctx context.Context, addr sdk.ValAddress) (validator stakingtypes.Validator, found bool)

	GetHistoricalInfo(ctx context.Context, height int64) (stakingtypes.HistoricalInfo, bool)
}

// BankKeeper Methods imported from bank should be defined here
type BankKeeper interface {
	SendKeeper
	SendCoinsFromModuleToModule(ctx context.Context, senderModule, recipientModule string, amt sdk.Coins) error
	MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	BurnCoins(ctx context.Context, moduleName string, amt sdk.Coins) error

	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins
}

type SendKeeper interface {
	ViewKeeper
}
type ViewKeeper interface {
	GetAllBalances(ctx context.Context, addr sdk.AccAddress) sdk.Coins
}
