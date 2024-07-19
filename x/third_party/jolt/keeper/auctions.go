package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) processEachReserve(ctx context.Context, c sdk.Coin, threshold sdk.Int) (bool, error) {
	liqMap := make(map[string]LiqData)

	// Load required liquidation data for every deposit/borrow denom
	mm, found := k.GetMoneyMarket(ctx, c.Denom)
	if !found {
		return false, errorsmod.Wrapf(types.ErrMarketNotFound, "no market found for denom %s", c.Denom)
	}

	priceData, err := k.pricefeedKeeper.GetCurrentPrice(ctx, mm.SpotMarketID)
	if err != nil {
		return false, err
	}

	liqMap[c.Denom] = LiqData{priceData.Price, mm.BorrowLimit.LoanToValue, mm.ConversionFactor}
	lData := liqMap[c.Denom]
	usdValue := sdk.NewDecFromInt(c.Amount).Quo(sdk.NewDecFromInt(lData.conversionFactor)).Mul(lData.price)

	// now we set the auction
	if usdValue.TruncateInt().GTE(threshold) {
		lot := c
		// if it is the jolt, we burn it directly.
		if lot.Denom == "ujolt" {
			err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(lot))
			if err != nil {
				return false, err
			}
			return true, nil
		}
		_, err := k.auctionKeeper.StartSurplusAuction(ctx, types.ModuleAccountName, lot, "ujolt")
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

// RunSurplusAuctions nets the surplus and debt balances and then creates surplus or debt auctions if the remaining balance is above the auction threshold parameter
func (k Keeper) RunSurplusAuctions(ctx context.Context) error {
	totalReserves, found := k.GetTotalReserves(ctx)
	if !found {
		return nil
	}
	pamas := k.GetParams(ctx)
	processedCoins := sdk.NewCoins()
	for _, el := range totalReserves {
		processed, err := k.processEachReserve(ctx, el, pamas.SurplusAuctionThreshold)
		if err != nil {
			ctx.Logger().Error("jolt", "surplusAuction", err)
			continue
		}
		if processed {
			processedCoins = processedCoins.Add(el)
		}
	}
	if !processedCoins.Empty() {
		newReserved := totalReserves.Sub(processedCoins...)
		k.SetTotalReserves(ctx, newReserved)
	}
	return nil
}
