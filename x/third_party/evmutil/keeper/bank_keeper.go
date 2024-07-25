package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	"github.com/joltify-finance/joltify_lending/x/third_party/evmutil/types"
)

const (
	// EvmDenom is the gas denom used by the evm
	EvmDenom = "ajolt"

	// CosmosDenom is the gas denom used by the jolt app
	CosmosDenom = "ujolt"
)

// ConversionMultiplier is the conversion multiplier between ajolt and ujolt
var ConversionMultiplier = sdkmath.NewInt(1_000_000_000_000)

var _ evmtypes.BankKeeper = EvmBankKeeper{}

// EvmBankKeeper is a BankKeeper wrapper for the x/evm module to allow the use
// of the 18 decimal ajolt coin on the evm.
// x/evm consumes gas and send coins by minting and burning ajolt coins in its module
// account and then sending the funds to the target account.
// This keeper uses both the ujolt coin and a separate ajolt balance to manage the
// extra percision needed by the evm.
type EvmBankKeeper struct {
	ajoltKeeper Keeper
	bk          types.BankKeeper
	ak          types.AccountKeeper
}

func (k EvmBankKeeper) IsSendEnabledCoins(ctx context.Context, coins ...sdk.Coin) error {
	// TODO implement me
	panic("implement me")
}

func (k EvmBankKeeper) SendCoinsFromModuleToAccountVirtual(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	// TODO implement me
	panic("implement me")
}

func (k EvmBankKeeper) SendCoinsFromAccountToModuleVirtual(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	// TODO implement me
	panic("implement me")
}

func (k EvmBankKeeper) BlockedAddr(addr sdk.AccAddress) bool {
	return k.bk.BlockedAddr(addr)
}

func NewEvmBankKeeper(uJoltKeeper Keeper, bk types.BankKeeper, ak types.AccountKeeper) EvmBankKeeper {
	return EvmBankKeeper{
		ajoltKeeper: uJoltKeeper,
		bk:          bk,
		ak:          ak,
	}
}

// GetBalance returns the total **spendable** balance of ajolt for a given account by address.
func (k EvmBankKeeper) GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	if denom != EvmDenom {
		panic(fmt.Errorf("only evm denom %s is supported by EvmBankKeeper", EvmDenom))
	}

	spendableCoins := k.bk.SpendableCoins(ctx, addr)
	ujolt := spendableCoins.AmountOf(CosmosDenom)
	ajolt := k.ajoltKeeper.GetBalance(ctx, addr)
	total := ujolt.Mul(ConversionMultiplier).Add(ajolt)
	return sdk.NewCoin(EvmDenom, total)
}

// SendCoins transfers ajolt coins from a AccAddress to an AccAddress.
func (k EvmBankKeeper) SendCoins(ctx context.Context, senderAddr sdk.AccAddress, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	// SendCoins method is not used by the evm module, but is required by the
	// evmtypes.BankKeeper interface. This must be updated if the evm module
	// is updated to use SendCoins.
	panic("not implemented")
}

// SendCoinsFromModuleToAccount transfers ajolt coins from a ModuleAccount to an AccAddress.
// It will panic if the module account does not exist. An error is returned if the recipient
// address is black-listed or if sending the tokens fails.
func (k EvmBankKeeper) SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	ujolt, ajolt, err := SplitAJoltCoins(amt)
	if err != nil {
		return err
	}

	if ujolt.Amount.IsPositive() {
		if err := k.bk.SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, sdk.NewCoins(ujolt)); err != nil {
			return err
		}
	}

	senderAddr := k.GetModuleAddress(senderModule)
	if err := k.ConvertOneUjoltToAjoltIfNeeded(ctx, senderAddr, ajolt); err != nil {
		return err
	}

	if err := k.ajoltKeeper.SendBalance(ctx, senderAddr, recipientAddr, ajolt); err != nil {
		return err
	}

	return k.ConvertAJoltToUJolt(ctx, recipientAddr)
}

// SendCoinsFromAccountToModule transfers ajolt coins from an AccAddress to a ModuleAccount.
// It will panic if the module account does not exist.
func (k EvmBankKeeper) SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	ujolt, uJoltNeeded, err := SplitAJoltCoins(amt)
	if err != nil {
		return err
	}

	if ujolt.IsPositive() {
		if err := k.bk.SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, sdk.NewCoins(ujolt)); err != nil {
			return err
		}
	}

	if err := k.ConvertOneUjoltToAjoltIfNeeded(ctx, senderAddr, uJoltNeeded); err != nil {
		return err
	}

	recipientAddr := k.GetModuleAddress(recipientModule)
	if err := k.ajoltKeeper.SendBalance(ctx, senderAddr, recipientAddr, uJoltNeeded); err != nil {
		return err
	}

	return k.ConvertAJoltToUJolt(ctx, recipientAddr)
}

// MintCoins mints ajolt coins by minting the equivalent ujolt coins and any remaining ajolt coins.
// It will panic if the module account does not exist or is unauthorized.
func (k EvmBankKeeper) MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error {
	ujolt, ajolt, err := SplitAJoltCoins(amt)
	if err != nil {
		return err
	}

	if ujolt.IsPositive() {
		if err := k.bk.MintCoins(ctx, moduleName, sdk.NewCoins(ujolt)); err != nil {
			return err
		}
	}

	recipientAddr := k.GetModuleAddress(moduleName)
	if err := k.ajoltKeeper.AddBalance(ctx, recipientAddr, ajolt); err != nil {
		return err
	}

	return k.ConvertAJoltToUJolt(ctx, recipientAddr)
}

// BurnCoins burns ajolt coins by burning the equivalent ujolt coins and any remaining ajolt coins.
// It will panic if the module account does not exist or is unauthorized.
func (k EvmBankKeeper) BurnCoins(ctx context.Context, moduleName string, amt sdk.Coins) error {
	ujolt, ajolt, err := SplitAJoltCoins(amt)
	if err != nil {
		return err
	}

	if ujolt.IsPositive() {
		if err := k.bk.BurnCoins(ctx, moduleName, sdk.NewCoins(ujolt)); err != nil {
			return err
		}
	}

	moduleAddr := k.GetModuleAddress(moduleName)
	if err := k.ConvertOneUjoltToAjoltIfNeeded(ctx, moduleAddr, ajolt); err != nil {
		return err
	}

	return k.ajoltKeeper.RemoveBalance(ctx, moduleAddr, ajolt)
}

// ConvertOneuJoltTouJoltIfNeeded converts 1 ujolt to ajolt for an address if
// its ajolt balance is smaller than the uJoltNeeded amount.
func (k EvmBankKeeper) ConvertOneUjoltToAjoltIfNeeded(ctx context.Context, addr sdk.AccAddress, uJoltNeeded sdkmath.Int) error {
	ajoltBal := k.ajoltKeeper.GetBalance(ctx, addr)
	if ajoltBal.GTE(uJoltNeeded) {
		return nil
	}

	uJoltToStore := sdk.NewCoins(sdk.NewCoin(CosmosDenom, sdkmath.OneInt()))
	if err := k.bk.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, uJoltToStore); err != nil {
		return err
	}

	// add 1uJolt equivalent of ajolt to addr
	uJoltToReceive := ConversionMultiplier
	if err := k.ajoltKeeper.AddBalance(ctx, addr, uJoltToReceive); err != nil {
		return err
	}

	return nil
}

// ConvertAJoltToUJolt converts all available ajolt to ujolt for a given AccAddress.
func (k EvmBankKeeper) ConvertAJoltToUJolt(ctx context.Context, addr sdk.AccAddress) error {
	totaluJolt := k.ajoltKeeper.GetBalance(ctx, addr)
	ujolt, _, err := SplitAJoltCoins(sdk.NewCoins(sdk.NewCoin(EvmDenom, totaluJolt)))
	if err != nil {
		return err
	}

	// do nothing if account does not have enough ajolt for a single ujolt
	uJoltToReceive := ujolt.Amount
	if !uJoltToReceive.IsPositive() {
		return nil
	}

	// remove ajolt used for converting to ujolt
	uJoltToBurn := uJoltToReceive.Mul(ConversionMultiplier)
	finalBal := totaluJolt.Sub(uJoltToBurn)
	if err := k.ajoltKeeper.SetBalance(ctx, addr, finalBal); err != nil {
		return err
	}

	fromAddr := k.GetModuleAddress(types.ModuleName)
	if err := k.bk.SendCoins(ctx, fromAddr, addr, sdk.NewCoins(ujolt)); err != nil {
		return err
	}

	return nil
}

func (k EvmBankKeeper) GetModuleAddress(moduleName string) sdk.AccAddress {
	addr := k.ak.GetModuleAddress(moduleName)
	if addr == nil {
		panic(errorsmod.Wrapf(errorsmod.ErrUnknownAddress, "module account %s does not exist", moduleName))
	}
	return addr
}

// SplitAJoltCoins splits ajolt coins to the equivalent ujolt coins and any remaining ajolt balance.
// An error will be returned if the coins are not valid or if the coins are not the ajolt denom.
func SplitAJoltCoins(coins sdk.Coins) (sdk.Coin, sdkmath.Int, error) {
	ajolt := sdkmath.ZeroInt()
	ujolt := sdk.NewCoin(CosmosDenom, sdkmath.ZeroInt())

	if len(coins) == 0 {
		return ujolt, ajolt, nil
	}

	if err := ValidateEvmCoins(coins); err != nil {
		return ujolt, ajolt, err
	}

	// note: we should always have len(coins) == 1 here since coins cannot have dup denoms after we validate.
	coin := coins[0]
	remainingBalance := coin.Amount.Mod(ConversionMultiplier)
	if remainingBalance.IsPositive() {
		ajolt = remainingBalance
	}
	uJoltAmount := coin.Amount.Quo(ConversionMultiplier)
	if uJoltAmount.IsPositive() {
		ujolt = sdk.NewCoin(CosmosDenom, uJoltAmount)
	}

	return ujolt, ajolt, nil
}

// ValidateEvmCoins validates the coins from evm is valid and is the EvmDenom (ajolt).
func ValidateEvmCoins(coins sdk.Coins) error {
	if len(coins) == 0 {
		return nil
	}

	// validate that coins are non-negative, sorted, and no dup denoms
	if err := coins.Validate(); err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, coins.String())
	}

	// validate that coin denom is ajolt
	if len(coins) != 1 || coins[0].Denom != EvmDenom {
		errMsg := fmt.Sprintf("invalid evm coin denom, only %s is supported", EvmDenom)
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, errMsg)
	}

	return nil
}
