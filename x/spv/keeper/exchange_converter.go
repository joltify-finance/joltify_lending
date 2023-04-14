package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) inboundConvertFromUSDWithMarketID(ctx sdk.Context, marketID string, amount sdkmath.Int) (sdkmath.Int, sdk.Dec, error) {
	currencyPrice, err := k.priceFeedKeeper.GetCurrentPrice(ctx, marketID)
	if err != nil {
		return sdk.ZeroInt(), sdk.ZeroDec(), err
	}
	outAmount := currencyPrice.Price.Quo(sdk.NewDecFromInt(amount)).TruncateInt()
	return outAmount, currencyPrice.Price, nil
}

func (k Keeper) outboundConvertToUSDWithMarketID(ctx sdk.Context, marketID string, amount sdkmath.Int) (sdkmath.Int, sdk.Dec, error) {
	currencyPrice, err := k.priceFeedKeeper.GetCurrentPrice(ctx, marketID)
	if err != nil {
		return sdk.ZeroInt(), sdk.ZeroDec(), err
	}
	outAmount := currencyPrice.Price.Mul(sdk.NewDecFromInt(amount)).TruncateInt()
	return outAmount, currencyPrice.Price, nil
}

func (k Keeper) inboundConvertFromUSD(inAmount sdkmath.Int, ratio sdk.Dec) sdkmath.Int {
	outAmount := ratio.Quo(sdk.NewDecFromInt(inAmount)).TruncateInt()
	return outAmount
}

func (k Keeper) outboundConvertToUSD(inAmount sdkmath.Int, ratio sdk.Dec) sdkmath.Int {
	outAmount := ratio.Mul(sdk.NewDecFromInt(inAmount)).TruncateInt()
	return outAmount
}
