package keeper

import (
	"context"
	"errors"

	sdkmath "cosmossdk.io/math"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
)

// Borrow funds
func (k Keeper) Borrow(ctx context.Context, borrower sdk.AccAddress, coins sdk.Coins) error {
	// Set any new denoms' global borrow index to 1.0
	for _, coin := range coins {
		_, foundInterestFactor := k.GetBorrowInterestFactor(ctx, coin.Denom)
		if !foundInterestFactor {
			_, foundMm := k.GetMoneyMarket(ctx, coin.Denom)
			if foundMm {
				k.SetBorrowInterestFactor(ctx, coin.Denom, sdkmath.LegacyOneDec())
			}
		}
	}

	// Call incentive hooks
	existingDeposit, hasExistingDeposit := k.GetDeposit(ctx, borrower)
	if hasExistingDeposit {
		k.BeforeDepositModified(ctx, existingDeposit)
	}
	existingBorrow, hasExistingBorrow := k.GetBorrow(ctx, borrower)
	if hasExistingBorrow {
		k.BeforeBorrowModified(ctx, existingBorrow)
	}

	k.SyncSupplyInterest(ctx, borrower)
	k.SyncBorrowInterest(ctx, borrower)

	// Validate borrow amount within user and protocol limits
	err := k.ValidateBorrow(ctx, borrower, coins)
	if err != nil {
		return err
	}

	// Sends coins from Jolt module account to user
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types2.ModuleAccountName, borrower, coins)
	if err != nil {
		if errors.Is(err, errorsmod.ErrInsufficientFunds) {
			macc := k.accountKeeper.GetModuleAccount(ctx, types2.ModuleAccountName)
			modAccCoins := k.bankKeeper.GetAllBalances(ctx, macc.GetAddress())
			for _, coin := range coins {
				_, isNegative := modAccCoins.SafeSub(coin)
				if isNegative {
					return errorsmod.Wrapf(types2.ErrBorrowExceedsAvailableBalance,
						"the requested borrow amount of %s exceeds the total amount of %s%s available to borrow",
						coin, modAccCoins.AmountOf(coin.Denom), coin.Denom,
					)
				}
			}
		}
		return err
	}

	interestFactors := types2.BorrowInterestFactors{}
	currBorrow, foundBorrow := k.GetBorrow(ctx, borrower)
	if foundBorrow {
		interestFactors = currBorrow.Index
	}
	for _, coin := range coins {
		interestFactorValue, foundValue := k.GetBorrowInterestFactor(ctx, coin.Denom)
		if foundValue {
			interestFactors = interestFactors.SetInterestFactor(coin.Denom, interestFactorValue)
		}
	}

	// Calculate new borrow amount
	var amount sdk.Coins
	if foundBorrow {
		amount = currBorrow.Amount.Add(coins...)
	} else {
		amount = coins
	}

	// Construct the user's new/updated borrow with amount and interest factors
	borrow := types2.NewBorrow(borrower, amount, interestFactors)
	if borrow.Amount.Empty() {
		k.DeleteBorrow(ctx, borrow)
	} else {
		k.SetBorrow(ctx, borrow)
	}

	// Update total borrowed amount by newly borrowed coins. Don't add user's pending interest as
	// it has already been included in the total borrowed coins by the BeginBlocker.
	k.IncrementBorrowedCoins(ctx, coins)

	if !hasExistingBorrow {
		k.AfterBorrowCreated(ctx, borrow)
	} else {
		k.AfterBorrowModified(ctx, borrow)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types2.EventTypeJoltBorrow,
			sdk.NewAttribute(types2.AttributeKeyBorrower, borrower.String()),
			sdk.NewAttribute(types2.AttributeKeyBorrowCoins, coins.String()),
		),
	)

	return nil
}

// ValidateBorrow validates a borrow request against borrower and protocol requirements
func (k Keeper) ValidateBorrow(ctx context.Context, borrower sdk.AccAddress, amount sdk.Coins) error {
	if amount.IsZero() {
		return types2.ErrBorrowEmptyCoins
	}

	// The reserve coins aren't available for users to borrow
	macc := k.accountKeeper.GetModuleAccount(ctx, types2.ModuleName)
	joltMaccCoins := k.bankKeeper.GetAllBalances(ctx, macc.GetAddress())
	reserveCoins, foundReserveCoins := k.GetTotalReserves(ctx)
	if !foundReserveCoins {
		reserveCoins = sdk.NewCoins()
	}
	fundsAvailableToBorrow, isNegative := joltMaccCoins.SafeSub(reserveCoins...)
	if isNegative {
		return errorsmod.Wrapf(types2.ErrReservesExceedCash, "reserves %s > cash %s", reserveCoins, joltMaccCoins)
	}
	if amount.IsAnyGT(fundsAvailableToBorrow) {
		return errorsmod.Wrapf(types2.ErrExceedsProtocolBorrowableBalance, "requested borrow %s > available to borrow %s", amount, fundsAvailableToBorrow)
	}

	// Get the proposed borrow USD value
	proprosedBorrowUSDValue := sdk.ZeroDec()
	for _, coin := range amount {
		moneyMarket, found := k.GetMoneyMarket(ctx, coin.Denom)
		if !found {
			return errorsmod.Wrapf(types2.ErrMarketNotFound, "no money market found for denom %s", coin.Denom)
		}

		// Calculate this coin's USD value and add it borrow's total USD value
		assetPriceInfo, err := k.pricefeedKeeper.GetCurrentPrice(ctx, moneyMarket.SpotMarketID)
		if err != nil {
			return errorsmod.Wrapf(types2.ErrPriceNotFound, "no price found for market %s", moneyMarket.SpotMarketID)
		}
		coinUSDValue := sdk.NewDecFromInt(coin.Amount).Quo(sdk.NewDecFromInt(moneyMarket.ConversionFactor)).Mul(assetPriceInfo.Price)

		// Validate the requested borrow value for the asset against the money market's global borrow limit
		if moneyMarket.BorrowLimit.HasMaxLimit {
			var assetTotalBorrowedAmount sdk.Int
			totalBorrowedCoins, found := k.GetBorrowedCoins(ctx)
			if !found {
				assetTotalBorrowedAmount = sdk.ZeroInt()
			} else {
				assetTotalBorrowedAmount = totalBorrowedCoins.AmountOf(coin.Denom)
			}
			newProposedAssetTotalBorrowedAmount := sdk.NewDecFromInt(assetTotalBorrowedAmount.Add(coin.Amount))
			if newProposedAssetTotalBorrowedAmount.GT(moneyMarket.BorrowLimit.MaximumLimit) {
				return errorsmod.Wrapf(types2.ErrGreaterThanAssetBorrowLimit,
					"proposed borrow would result in %s borrowed, but the maximum global asset borrow limit is %s",
					newProposedAssetTotalBorrowedAmount, moneyMarket.BorrowLimit.MaximumLimit)
			}
		}
		proprosedBorrowUSDValue = proprosedBorrowUSDValue.Add(coinUSDValue)

	}

	// Get the total borrowable USD amount at user's existing deposits
	deposit, found := k.GetDeposit(ctx, borrower)
	if !found {
		return errorsmod.Wrapf(types2.ErrDepositsNotFound, "no deposits found for %s", borrower)
	}
	totalUsableAmount := sdk.ZeroDec()
	for _, coin := range deposit.Amount {
		moneyMarket, found := k.GetMoneyMarket(ctx, coin.Denom)
		if !found {
			return errorsmod.Wrapf(types2.ErrMarketNotFound, "no money market found for denom %s", coin.Denom)
		}

		// Calculate the borrowable amount and add it to the user's total borrowable amount
		assetPriceInfo, err := k.pricefeedKeeper.GetCurrentPrice(ctx, moneyMarket.SpotMarketID)
		if err != nil {
			return errorsmod.Wrapf(types2.ErrPriceNotFound, "no price found for market %s", moneyMarket.SpotMarketID)
		}
		depositUSDValue := sdk.NewDecFromInt(coin.Amount).Quo(sdk.NewDecFromInt(moneyMarket.ConversionFactor)).Mul(assetPriceInfo.Price)
		UsableAmountForDeposit := depositUSDValue.Mul(moneyMarket.BorrowLimit.LoanToValue)
		totalUsableAmount = totalUsableAmount.Add(UsableAmountForDeposit)
	}

	// Get the total USD value of user's existing borrows
	existingBorrowUSDValue := sdk.ZeroDec()
	existingBorrow, found := k.GetBorrow(ctx, borrower)
	if found {
		for _, coin := range existingBorrow.Amount {
			moneyMarket, found := k.GetMoneyMarket(ctx, coin.Denom)
			if !found {
				return errorsmod.Wrapf(types2.ErrMarketNotFound, "no money market found for denom %s", coin.Denom)
			}

			// Calculate this borrow coin's USD value and add it to the total previous borrowed USD value
			assetPriceInfo, err := k.pricefeedKeeper.GetCurrentPrice(ctx, moneyMarket.SpotMarketID)
			if err != nil {
				return errorsmod.Wrapf(types2.ErrPriceNotFound, "no price found for market %s", moneyMarket.SpotMarketID)
			}
			coinUSDValue := sdk.NewDecFromInt(coin.Amount).Quo(sdk.NewDecFromInt(moneyMarket.ConversionFactor)).Mul(assetPriceInfo.Price)
			existingBorrowUSDValue = existingBorrowUSDValue.Add(coinUSDValue)
		}
	}

	// Borrow's updated total USD value must be greater than the minimum global USD borrow limit
	totalBorrowUSDValue := proprosedBorrowUSDValue.Add(existingBorrowUSDValue)
	if totalBorrowUSDValue.LT(k.GetMinimumBorrowUSDValue(ctx)) {
		return errorsmod.Wrapf(types2.ErrBelowMinimumBorrowValue, "the proposed borrow's USD value $%s is below the minimum borrow limit $%s", totalBorrowUSDValue, k.GetMinimumBorrowUSDValue(ctx))
	}

	// Validate that the proposed borrow's USD value is within user's borrowable limit
	if proprosedBorrowUSDValue.GT(totalUsableAmount.Sub(existingBorrowUSDValue)) {
		return errorsmod.Wrapf(types2.ErrInsufficientLoanToValue, "requested borrow %s exceeds the allowable amount as determined by the collateralization ratio", amount)
	}
	return nil
}

// IncrementBorrowedCoins increments the total amount of borrowed coins by the newCoins parameter
func (k Keeper) IncrementBorrowedCoins(ctx context.Context, newCoins sdk.Coins) {
	borrowedCoins, found := k.GetBorrowedCoins(ctx)
	if !found {
		if !newCoins.Empty() {
			k.SetBorrowedCoins(ctx, newCoins)
		}
	} else {
		k.SetBorrowedCoins(ctx, borrowedCoins.Add(newCoins...))
	}
}

// DecrementBorrowedCoins decrements the total amount of borrowed coins by the coins parameter
func (k Keeper) DecrementBorrowedCoins(ctx context.Context, coins sdk.Coins) error {
	borrowedCoins, found := k.GetBorrowedCoins(ctx)
	if !found {
		return errorsmod.Wrapf(types2.ErrBorrowedCoinsNotFound, "cannot repay coins if no coins are currently borrowed")
	}

	updatedBorrowedCoins, isNegative := borrowedCoins.SafeSub(coins...)
	if isNegative {
		coinsToSubtract := sdk.NewCoins()
		for _, coin := range coins {
			if borrowedCoins.AmountOf(coin.Denom).LT(coin.Amount) {
				if borrowedCoins.AmountOf(coin.Denom).GT(sdk.ZeroInt()) {
					coinsToSubtract = coinsToSubtract.Add(sdk.NewCoin(coin.Denom, borrowedCoins.AmountOf(coin.Denom)))
				}
			} else {
				coinsToSubtract = coinsToSubtract.Add(coin)
			}
		}
		updatedBorrowedCoins = borrowedCoins.Sub(coinsToSubtract...)
	}

	k.SetBorrowedCoins(ctx, updatedBorrowedCoins)
	return nil
}

// GetSyncedBorrow returns a borrow object containing current balances and indexes
func (k Keeper) GetSyncedBorrow(ctx context.Context, borrower sdk.AccAddress) (types2.Borrow, bool) {
	borrow, found := k.GetBorrow(ctx, borrower)
	if !found {
		return types2.Borrow{}, false
	}

	return k.loadSyncedBorrow(ctx, borrow), true
}

// loadSyncedBorrow calculates a user's synced borrow, but does not update state
func (k Keeper) loadSyncedBorrow(ctx context.Context, borrow types2.Borrow) types2.Borrow {
	totalNewInterest := sdk.Coins{}
	newBorrowIndexes := types2.BorrowInterestFactors{}
	for _, coin := range borrow.Amount {
		interestFactorValue, foundInterestFactorValue := k.GetBorrowInterestFactor(ctx, coin.Denom)
		if foundInterestFactorValue {
			// Locate the interest factor by coin denom in the user's list of interest factors
			foundAtIndex := -1
			for i := range borrow.Index {
				if borrow.Index[i].Denom == coin.Denom {
					foundAtIndex = i
					break
				}
			}

			// Calculate interest owed by user for this asset
			if foundAtIndex != -1 {
				storedAmount := sdk.NewDecFromInt(borrow.Amount.AmountOf(coin.Denom))
				userLastInterestFactor := borrow.Index[foundAtIndex].Value
				coinInterest := (storedAmount.Quo(userLastInterestFactor).Mul(interestFactorValue)).Sub(storedAmount)
				totalNewInterest = totalNewInterest.Add(sdk.NewCoin(coin.Denom, coinInterest.TruncateInt()))
			}
		}

		borrowIndex := types2.NewBorrowInterestFactor(coin.Denom, interestFactorValue)
		newBorrowIndexes = append(newBorrowIndexes, borrowIndex)
	}

	return types2.NewBorrow(borrow.Borrower, borrow.Amount.Add(totalNewInterest...), newBorrowIndexes)
}
