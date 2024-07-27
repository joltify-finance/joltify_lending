package keeper

import (
	sdkmath "cosmossdk.io/math"
)

func (k Keeper) inboundConvertFromUSDWithMarketID(ctx context.Context, marketID string, amount sdkmath.Int) (sdkmath.Int, sdkmath.LegacyDec, error) {
	currencyPrice, err := k.priceFeedKeeper.GetCurrentPrice(ctx, marketID)
	if err != nil {
		return sdkmath.ZeroInt(), sdkmath.LegacyZeroDec(), err
	}
	outAmount := sdkmath.LegacyNewDecFromInt(amount).Quo(currencyPrice.Price).TruncateInt()
	return outAmount, currencyPrice.Price, nil
}

func (k Keeper) outboundConvertToUSDWithMarketID(ctx context.Context, marketID string, amount sdkmath.Int) (sdkmath.Int, sdkmath.LegacyDec, error) {
	currencyPrice, err := k.priceFeedKeeper.GetCurrentPrice(ctx, marketID)
	if err != nil {
		return sdkmath.ZeroInt(), sdkmath.LegacyZeroDec(), err
	}
	outAmount := currencyPrice.Price.Mul(sdkmath.LegacyNewDecFromInt(amount)).TruncateInt()
	return outAmount, currencyPrice.Price, nil
}

func inboundConvertFromUSD(inAmount sdkmath.Int, ratio sdkmath.LegacyDec) sdkmath.Int {
	outAmount := sdkmath.LegacyNewDecFromInt(inAmount).Quo(ratio).TruncateInt()
	return outAmount
}

func outboundConvertToUSD(inAmount sdkmath.Int, ratio sdkmath.LegacyDec) sdkmath.Int {
	outAmount := ratio.Mul(sdkmath.LegacyNewDecFromInt(inAmount)).TruncateInt()
	return outAmount
}
