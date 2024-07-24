package keeper

import (
	"sort"

	sdkmath "cosmossdk.io/math"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
)

// LiqData holds liquidation-related data
type LiqData struct {
	price            sdkmath.LegacyDec
	ltv              sdkmath.LegacyDec
	conversionFactor sdk.Int
}

// AttemptKeeperLiquidation enables a keeper to liquidate an individual borrower's position
func (k Keeper) AttemptKeeperLiquidation(ctx context.Context, keeper sdk.AccAddress, borrower sdk.AccAddress) error {
	deposit, found := k.GetDeposit(ctx, borrower)
	if !found {
		return types2.ErrDepositNotFound
	}

	borrow, found := k.GetBorrow(ctx, borrower)
	if !found {
		return types2.ErrBorrowNotFound
	}

	// Call incentive hooks
	k.BeforeDepositModified(ctx, deposit)
	k.BeforeBorrowModified(ctx, borrow)

	k.SyncBorrowInterest(ctx, borrower)
	k.SyncSupplyInterest(ctx, borrower)

	deposit, found = k.GetDeposit(ctx, borrower)
	if !found {
		return types2.ErrDepositNotFound
	}

	borrow, found = k.GetBorrow(ctx, borrower)
	if !found {
		return types2.ErrBorrowNotFound
	}

	isWithinRange, _, err := k.IsWithinValidLtvRange(ctx, deposit, borrow)
	if err != nil {
		return err
	}
	if isWithinRange {
		return errorsmod.Wrapf(types2.ErrBorrowNotLiquidatable, "position is within valid LTV range")
	}

	// Sending coins to auction module with keeper address getting % of the profits
	borrowDenoms := getDenoms(borrow.Amount)
	depositDenoms := getDenoms(deposit.Amount)
	err = k.SeizeDeposits(ctx, keeper, deposit, borrow, depositDenoms, borrowDenoms)
	if err != nil {
		return err
	}

	deposit.Amount = sdk.NewCoins()
	k.DeleteDeposit(ctx, deposit)
	k.AfterDepositModified(ctx, deposit)

	borrow.Amount = sdk.NewCoins()
	k.DeleteBorrow(ctx, borrow)
	k.AfterBorrowModified(ctx, borrow)
	return nil
}

// SeizeDeposits seizes a list of deposits and sends them to auction
func (k Keeper) SeizeDeposits(ctx context.Context, keeper sdk.AccAddress, deposit types2.Deposit,
	borrow types2.Borrow, dDenoms, bDenoms []string,
) error {
	liqMap, err := k.LoadLiquidationData(ctx, deposit, borrow)
	if err != nil {
		return err
	}

	// Seize % of every deposit and send to the keeper
	keeperRewardCoins := sdk.Coins{}
	for _, depCoin := range deposit.Amount {
		mm, _ := k.GetMoneyMarket(ctx, depCoin.Denom)
		keeperReward := mm.KeeperRewardPercentage.MulInt(depCoin.Amount).TruncateInt()
		if keeperReward.GT(sdk.ZeroInt()) {
			// Send keeper their reward
			keeperCoin := sdk.NewCoin(depCoin.Denom, keeperReward)
			keeperRewardCoins = append(keeperRewardCoins, keeperCoin)
		}
	}
	if !keeperRewardCoins.Empty() {
		if err := k.DecrementSuppliedCoins(ctx, keeperRewardCoins); err != nil {
			return err
		}
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types2.ModuleAccountName, keeper, keeperRewardCoins); err != nil {
			return err
		}
	}

	// All deposit amounts not given to keeper as rewards are eligible to be auctioned off
	aucDeposits := deposit.Amount.Sub(keeperRewardCoins...)

	// Build valuation map to hold deposit coin USD valuations
	depositCoinValues := types2.NewValuationMap()
	for _, deposit := range aucDeposits {
		dData := liqMap[deposit.Denom]
		dCoinUsdValue := sdk.NewDecFromInt(deposit.Amount).Quo(sdk.NewDecFromInt(dData.conversionFactor)).Mul(dData.price)
		depositCoinValues.Increment(deposit.Denom, dCoinUsdValue)
	}

	// Build valuation map to hold borrow coin USD valuations
	borrowCoinValues := types2.NewValuationMap()
	for _, bCoin := range borrow.Amount {
		bData := liqMap[bCoin.Denom]
		bCoinUsdValue := sdk.NewDecFromInt(bCoin.Amount).Quo(sdk.NewDecFromInt(bData.conversionFactor)).Mul(bData.price)
		borrowCoinValues.Increment(bCoin.Denom, bCoinUsdValue)
	}

	// Loan-to-Value ratio after sending keeper their reward
	depositUsdValue := depositCoinValues.Sum()
	if depositUsdValue.IsZero() {
		// Deposit value can be zero if params.KeeperRewardPercent is 1.0, or all deposit asset prices are zero.
		// In this case the full deposit will be sent to the keeper and no auctions started.
		return nil
	}
	ltv := borrowCoinValues.Sum().Quo(depositUsdValue)

	liquidatedCoins, err := k.StartAuctions(ctx, deposit.Depositor, borrow.Amount, aucDeposits, depositCoinValues, borrowCoinValues, ltv, liqMap)
	// If some coins were liquidated and sent to auction prior to error, still need to emit liquidation event
	if !liquidatedCoins.Empty() {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types2.EventTypeHardLiquidation,
				sdk.NewAttribute(types2.AttributeKeyLiquidatedOwner, deposit.Depositor.String()),
				sdk.NewAttribute(types2.AttributeKeyLiquidatedCoins, liquidatedCoins.String()),
				sdk.NewAttribute(types2.AttributeKeyKeeper, keeper.String()),
				sdk.NewAttribute(types2.AttributeKeyKeeperRewardCoins, keeperRewardCoins.String()),
			),
		)
	}
	// Returns nil if there's no error
	return err
}

// StartAuctions attempts to start auctions for seized assets
func (k Keeper) StartAuctions(ctx context.Context, borrower sdk.AccAddress, borrows, deposits sdk.Coins,
	depositCoinValues, borrowCoinValues types2.ValuationMap, ltv sdkmath.LegacyDec, liqMap map[string]LiqData,
) (sdk.Coins, error) {
	// Sort keys to ensure deterministic behavior
	bKeys := borrowCoinValues.GetSortedKeys()
	dKeys := depositCoinValues.GetSortedKeys()

	// Set up auction constants
	returnAddrs := []sdk.AccAddress{borrower}
	weights := []sdkmath.Int{sdk.NewInt(100)}
	debt := sdk.NewCoin("debt", sdk.ZeroInt())

	macc := k.accountKeeper.GetModuleAccount(ctx, types2.ModuleAccountName)
	maccCoins := k.bankKeeper.SpendableCoins(ctx, macc.GetAddress())

	var liquidatedCoins sdk.Coins
	for _, bKey := range bKeys {
		bValue := borrowCoinValues.Get(bKey)
		maxLotSize := bValue.Quo(ltv)

		for _, dKey := range dKeys {
			dValue := depositCoinValues.Get(dKey)
			if maxLotSize.Equal(sdk.ZeroDec()) {
				break // exit out of the loop if we have cleared the full amount
			}

			if dValue.GTE(maxLotSize) { // We can start an auction for the whole borrow amount]
				bid := sdk.NewCoin(bKey, borrows.AmountOf(bKey))

				lotSize := maxLotSize.MulInt(liqMap[dKey].conversionFactor).Quo(liqMap[dKey].price)
				if lotSize.TruncateInt().Equal(sdk.ZeroInt()) {
					continue
				}
				lot := sdk.NewCoin(dKey, lotSize.TruncateInt())

				insufficientLotFunds := false
				if lot.Amount.GT(maccCoins.AmountOf(dKey)) {
					insufficientLotFunds = true
					lot = sdk.NewCoin(lot.Denom, maccCoins.AmountOf(dKey))
				}

				// Sanity check that we can deliver coins to the liquidator account
				if deposits.AmountOf(dKey).LT(lot.Amount) {
					return liquidatedCoins, types2.ErrInsufficientCoins
				}

				// Start auction: bid = full borrow amount, lot = maxLotSize
				_, err := k.auctionKeeper.StartCollateralAuction(ctx, types2.ModuleAccountName, lot, bid, returnAddrs, weights, debt)
				if err != nil {
					return liquidatedCoins, err
				}
				// Decrement supplied coins and decrement borrowed coins optimistically
				err = k.DecrementSuppliedCoins(ctx, sdk.Coins{lot})
				if err != nil {
					return liquidatedCoins, err
				}
				err = k.DecrementBorrowedCoins(ctx, sdk.Coins{bid})
				if err != nil {
					return liquidatedCoins, err
				}

				// Add lot to liquidated coins
				liquidatedCoins = liquidatedCoins.Add(lot)

				// Update USD valuation maps
				borrowCoinValues.SetZero(bKey)
				depositCoinValues.Decrement(dKey, maxLotSize)
				// Update deposits, borrows
				borrows = borrows.Sub(bid)
				if insufficientLotFunds {
					deposits = deposits.Sub(sdk.NewCoin(dKey, deposits.AmountOf(dKey)))
				} else {
					deposits = deposits.Sub(lot)
				}
				// Update max lot size
				maxLotSize = sdk.ZeroDec()
			} else { // We can only start an auction for the partial borrow amount
				maxBid := dValue.Mul(ltv)
				bidSize := maxBid.MulInt(liqMap[bKey].conversionFactor).Quo(liqMap[bKey].price)
				bid := sdk.NewCoin(bKey, bidSize.TruncateInt())
				lot := sdk.NewCoin(dKey, deposits.AmountOf(dKey))

				if bid.Amount.Equal(sdk.ZeroInt()) || lot.Amount.Equal(sdk.ZeroInt()) {
					continue
				}

				insufficientLotFunds := false
				if lot.Amount.GT(maccCoins.AmountOf(dKey)) {
					insufficientLotFunds = true
					lot = sdk.NewCoin(lot.Denom, maccCoins.AmountOf(dKey))
				}

				// Sanity check that we can deliver coins to the liquidator account
				if deposits.AmountOf(dKey).LT(lot.Amount) {
					return liquidatedCoins, types2.ErrInsufficientCoins
				}

				// Start auction: bid = maxBid, lot = whole deposit amount
				_, err := k.auctionKeeper.StartCollateralAuction(ctx, types2.ModuleAccountName, lot, bid, returnAddrs, weights, debt)
				if err != nil {
					return liquidatedCoins, err
				}
				// Decrement supplied coins and decrement borrowed coins optimistically
				err = k.DecrementSuppliedCoins(ctx, sdk.Coins{lot})
				if err != nil {
					return liquidatedCoins, err
				}
				err = k.DecrementBorrowedCoins(ctx, sdk.Coins{bid})
				if err != nil {
					return liquidatedCoins, err
				}

				// Add lot to liquidated coins
				liquidatedCoins = liquidatedCoins.Add(lot)

				// Update variables to account for partial auction
				borrowCoinValues.Decrement(bKey, maxBid)
				depositCoinValues.SetZero(dKey)

				borrows = borrows.Sub(bid)
				if insufficientLotFunds {
					deposits = deposits.Sub(sdk.NewCoin(dKey, deposits.AmountOf(dKey)))
				} else {
					deposits = deposits.Sub(lot)
				}

				// Update max lot size
				maxLotSize = borrowCoinValues.Get(bKey).Quo(ltv)
			}
		}
	}

	// Send any remaining deposit back to the original borrower
	for _, dKey := range dKeys {
		remaining := deposits.AmountOf(dKey)
		if remaining.GT(sdk.ZeroInt()) {
			returnCoin := sdk.NewCoins(sdk.NewCoin(dKey, remaining))
			err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types2.ModuleAccountName, borrower, returnCoin)
			if err != nil {
				return liquidatedCoins, err
			}
		}
	}

	return liquidatedCoins, nil
}

// IsWithinValidLtvRange compares a borrow and deposit to see if it's within a valid LTV range at current prices
func (k Keeper) IsWithinValidLtvRange(ctx context.Context, deposit types2.Deposit, borrow types2.Borrow) (bool, sdkmath.LegacyDec, error) {
	liqMap, err := k.LoadLiquidationData(ctx, deposit, borrow)
	if err != nil {
		return false, sdkmath.LegacyDec{}, err
	}

	totalBorrowableUSDAmount := sdk.ZeroDec()
	for _, depCoin := range deposit.Amount {
		lData := liqMap[depCoin.Denom]
		usdValue := sdk.NewDecFromInt(depCoin.Amount).Quo(sdk.NewDecFromInt(lData.conversionFactor)).Mul(lData.price)
		borrowableUSDAmountForDeposit := usdValue.Mul(lData.ltv)
		totalBorrowableUSDAmount = totalBorrowableUSDAmount.Add(borrowableUSDAmountForDeposit)
	}

	totalBorrowedUSDAmount := sdk.ZeroDec()
	for _, coin := range borrow.Amount {
		lData := liqMap[coin.Denom]
		usdValue := sdk.NewDecFromInt(coin.Amount).Quo(sdk.NewDecFromInt(lData.conversionFactor)).Mul(lData.price)
		totalBorrowedUSDAmount = totalBorrowedUSDAmount.Add(usdValue)
	}

	var ratio sdkmath.LegacyDec
	if totalBorrowableUSDAmount.Equal(sdk.ZeroDec()) {
		ratio = sdk.MustNewDecFromStr("10000")
	} else {
		ratio = totalBorrowedUSDAmount.Quo(totalBorrowableUSDAmount)
	}
	// Check if the user's has borrowed more than they're allowed to
	if totalBorrowedUSDAmount.GT(totalBorrowableUSDAmount) {
		return false, ratio, nil
	}
	return true, ratio, nil
}

// GetStoreLTV calculates the user's current LTV based on their deposits/borrows in the store
// and does not include any outsanding interest.
func (k Keeper) GetStoreLTV(ctx context.Context, addr sdk.AccAddress) (sdkmath.LegacyDec, error) {
	// Fetch deposits and parse coin denoms
	deposit, found := k.GetDeposit(ctx, addr)
	if !found {
		return sdk.ZeroDec(), nil
	}

	// Fetch borrow balances and parse coin denoms
	borrow, found := k.GetBorrow(ctx, addr)
	if !found {
		return sdk.ZeroDec(), nil
	}

	return k.CalculateLtv(ctx, deposit, borrow)
}

// CalculateLtv calculates the potential LTV given a user's deposits and borrows.
// The boolean returned indicates if the LTV should be added to the store's LTV index.
func (k Keeper) CalculateLtv(ctx context.Context, deposit types2.Deposit, borrow types2.Borrow) (sdkmath.LegacyDec, error) {
	// Load required liquidation data for every deposit/borrow denom
	liqMap, err := k.LoadLiquidationData(ctx, deposit, borrow)
	if err != nil {
		return sdk.ZeroDec(), nil
	}

	// Build valuation map to hold deposit coin USD valuations
	depositCoinValues := types2.NewValuationMap()
	for _, depCoin := range deposit.Amount {
		dData := liqMap[depCoin.Denom]
		dCoinUsdValue := sdk.NewDecFromInt(depCoin.Amount).Quo(sdk.NewDecFromInt(dData.conversionFactor)).Mul(dData.price)
		depositCoinValues.Increment(depCoin.Denom, dCoinUsdValue)
	}

	// Build valuation map to hold borrow coin USD valuations
	borrowCoinValues := types2.NewValuationMap()
	for _, bCoin := range borrow.Amount {
		bData := liqMap[bCoin.Denom]
		bCoinUsdValue := sdk.NewDecFromInt(bCoin.Amount).Quo(sdk.NewDecFromInt(bData.conversionFactor)).Mul(bData.price)
		borrowCoinValues.Increment(bCoin.Denom, bCoinUsdValue)
	}

	// User doesn't have any deposits, catch divide by 0 error
	sumDeposits := depositCoinValues.Sum()
	if sumDeposits.Equal(sdk.ZeroDec()) {
		return sdk.ZeroDec(), nil
	}

	// Loan-to-Value ratio
	return borrowCoinValues.Sum().Quo(sumDeposits), nil
}

// LoadLiquidationData returns liquidation data, deposit, borrow
func (k Keeper) LoadLiquidationData(ctx context.Context, deposit types2.Deposit, borrow types2.Borrow) (map[string]LiqData, error) {
	liqMap := make(map[string]LiqData)

	borrowDenoms := getDenoms(borrow.Amount)
	depositDenoms := getDenoms(deposit.Amount)
	denoms := removeDuplicates(borrowDenoms, depositDenoms)

	// Load required liquidation data for every deposit/borrow denom
	for _, denom := range denoms {
		mm, found := k.GetMoneyMarket(ctx, denom)
		if !found {
			return liqMap, errorsmod.Wrapf(types2.ErrMarketNotFound, "no market found for denom %s", denom)
		}

		priceData, err := k.pricefeedKeeper.GetCurrentPrice(ctx, mm.SpotMarketID)
		if err != nil {
			return liqMap, err
		}

		liqMap[denom] = LiqData{priceData.Price, mm.BorrowLimit.LoanToValue, mm.ConversionFactor}
	}

	return liqMap, nil
}

func getDenoms(coins sdk.Coins) []string {
	var denoms []string
	for _, coin := range coins {
		denoms = append(denoms, coin.Denom)
	}
	return denoms
}

func removeDuplicates(one []string, two []string) []string {
	check := make(map[string]int)
	fullList := one
	fullList = append(fullList, two...)

	var res []string
	for _, val := range fullList {
		check[val] = 1
	}

	for key := range check {
		res = append(res, key)
	}
	sort.Strings(res)
	return res
}
