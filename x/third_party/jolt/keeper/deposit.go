package keeper

import (
	"errors"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
)

// Deposit deposit
func (k Keeper) Deposit(ctx context.Context, depositor sdk.AccAddress, coins sdk.Coins) error {
	// Set any new denoms' global supply index to 1.0
	for _, coin := range coins {
		_, foundInterestFactor := k.GetSupplyInterestFactor(ctx, coin.Denom)
		if !foundInterestFactor {
			_, foundMm := k.GetMoneyMarket(ctx, coin.Denom)
			if foundMm {
				k.SetSupplyInterestFactor(ctx, coin.Denom, sdk.OneDec())
			}
		}
	}

	// Call incentive hook
	existingDeposit, hasExistingDeposit := k.GetDeposit(ctx, depositor)
	if hasExistingDeposit {
		k.BeforeDepositModified(ctx, existingDeposit)
	}

	// Sync any outstanding interest
	k.SyncSupplyInterest(ctx, depositor)

	err := k.ValidateDeposit(ctx, coins)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, depositor, types2.ModuleAccountName, coins)
	if err != nil {
		if errors.Is(err, errorsmod.ErrInsufficientFunds) {
			acc := k.accountKeeper.GetAccount(ctx, depositor)
			accCoins := k.bankKeeper.SpendableCoins(ctx, acc.GetAddress())
			for _, coin := range coins {
				_, isNegative := accCoins.SafeSub(coin)
				if isNegative {
					return errorsmod.Wrapf(types2.ErrBorrowExceedsAvailableBalance,
						"insufficient funds: the requested deposit amount of %s exceeds the total available account funds of %s%s",
						coin, accCoins.AmountOf(coin.Denom), coin.Denom,
					)
				}
			}
		}
	}
	if err != nil {
		return err
	}

	interestFactors := types2.SupplyInterestFactors{}
	currDeposit, foundDeposit := k.GetDeposit(ctx, depositor)
	if foundDeposit {
		interestFactors = currDeposit.Index
	}
	for _, coin := range coins {
		interestFactorValue, foundValue := k.GetSupplyInterestFactor(ctx, coin.Denom)
		if foundValue {
			interestFactors = interestFactors.SetInterestFactor(coin.Denom, interestFactorValue)
		}
	}

	// Calculate new deposit amount
	var amount sdk.Coins
	if foundDeposit {
		amount = currDeposit.Amount.Add(coins...)
	} else {
		amount = coins
	}
	// Update the depositer's amount and supply interest factors in the store
	deposit := types2.NewDeposit(depositor, amount, interestFactors)

	if deposit.Amount.Empty() {
		k.DeleteDeposit(ctx, deposit)
	} else {
		k.SetDeposit(ctx, deposit)
	}

	k.IncrementSuppliedCoins(ctx, coins)
	if !foundDeposit { // User's first deposit
		k.AfterDepositCreated(ctx, deposit)
	} else {
		k.AfterDepositModified(ctx, deposit)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types2.EventTypeHardDeposit,
			sdk.NewAttribute(sdk.AttributeKeyAmount, coins.String()),
			sdk.NewAttribute(types2.AttributeKeyDepositor, deposit.Depositor.String()),
		),
	)

	return nil
}

// ValidateDeposit validates a deposit
func (k Keeper) ValidateDeposit(ctx context.Context, coins sdk.Coins) error {
	for _, depCoin := range coins {
		_, foundMm := k.GetMoneyMarket(ctx, depCoin.Denom)
		if !foundMm {
			return errorsmod.Wrapf(types2.ErrInvalidDepositDenom, "money market denom %s not found", depCoin.Denom)
		}
	}

	return nil
}

// GetTotalDeposited returns the total amount deposited for the input deposit type and deposit denom
func (k Keeper) GetTotalDeposited(ctx context.Context, depositDenom string) (total sdk.Int) {
	macc := k.accountKeeper.GetModuleAccount(ctx, types2.ModuleAccountName)
	return k.bankKeeper.GetBalance(ctx, macc.GetAddress(), depositDenom).Amount
}

// IncrementSuppliedCoins increments the total amount of supplied coins by the newCoins parameter
func (k Keeper) IncrementSuppliedCoins(ctx context.Context, newCoins sdk.Coins) {
	suppliedCoins, found := k.GetSuppliedCoins(ctx)
	if !found {
		if !newCoins.Empty() {
			k.SetSuppliedCoins(ctx, newCoins)
		}
	} else {
		k.SetSuppliedCoins(ctx, suppliedCoins.Add(newCoins...))
	}
}

// DecrementSuppliedCoins decrements the total amount of supplied coins by the coins parameter
func (k Keeper) DecrementSuppliedCoins(ctx context.Context, coins sdk.Coins) error {
	suppliedCoins, found := k.GetSuppliedCoins(ctx)
	if !found {
		return errorsmod.Wrapf(types2.ErrSuppliedCoinsNotFound, "cannot withdraw if no coins are deposited")
	}
	updatedSuppliedCoins, isNegative := suppliedCoins.SafeSub(coins...)
	if isNegative {
		coinsToSubtract := sdk.NewCoins()
		for _, coin := range coins {
			if suppliedCoins.AmountOf(coin.Denom).LT(coin.Amount) {
				if suppliedCoins.AmountOf(coin.Denom).GT(sdk.ZeroInt()) {
					coinsToSubtract = coinsToSubtract.Add(sdk.NewCoin(coin.Denom, suppliedCoins.AmountOf(coin.Denom)))
				}
			} else {
				coinsToSubtract = coinsToSubtract.Add(coin)
			}
		}
		updatedSuppliedCoins = suppliedCoins.Sub(coinsToSubtract...)
	}

	k.SetSuppliedCoins(ctx, updatedSuppliedCoins)
	return nil
}

// GetSyncedDeposit returns a deposit object containing current balances and indexes
func (k Keeper) GetSyncedDeposit(ctx context.Context, depositor sdk.AccAddress) (types2.Deposit, bool) {
	deposit, found := k.GetDeposit(ctx, depositor)
	if !found {
		return types2.Deposit{}, false
	}

	return k.loadSyncedDeposit(ctx, deposit), true
}

// loadSyncedDeposit calculates a user's synced deposit, but does not update state
func (k Keeper) loadSyncedDeposit(ctx context.Context, deposit types2.Deposit) types2.Deposit {
	totalNewInterest := sdk.Coins{}
	newSupplyIndexes := types2.SupplyInterestFactors{}
	for _, coin := range deposit.Amount {
		interestFactorValue, foundInterestFactorValue := k.GetSupplyInterestFactor(ctx, coin.Denom)
		if foundInterestFactorValue {
			// Locate the interest factor by coin denom in the user's list of interest factors
			foundAtIndex := -1
			for i := range deposit.Index {
				if deposit.Index[i].Denom == coin.Denom {
					foundAtIndex = i
					break
				}
			}

			// Calculate interest that will be paid to user for this asset
			if foundAtIndex != -1 {
				storedAmount := sdk.NewDecFromInt(deposit.Amount.AmountOf(coin.Denom))
				userLastInterestFactor := deposit.Index[foundAtIndex].Value
				coinInterest := (storedAmount.Quo(userLastInterestFactor).Mul(interestFactorValue)).Sub(storedAmount)
				totalNewInterest = totalNewInterest.Add(sdk.NewCoin(coin.Denom, coinInterest.TruncateInt()))
			}
		}

		supplyIndex := types2.NewSupplyInterestFactor(coin.Denom, interestFactorValue)
		newSupplyIndexes = append(newSupplyIndexes, supplyIndex)
	}

	return types2.NewDeposit(deposit.Depositor, deposit.Amount.Add(totalNewInterest...), newSupplyIndexes)
}
