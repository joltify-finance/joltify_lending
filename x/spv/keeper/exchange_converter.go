package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) inboundConvertFromUSDWithMarketID(ctx context.Context, marketID string, amount sdkmath.Int) (sdkmath.Int, sdk.Dec, error) {
	currencyPrice, err := k.priceFeedKeeper.GetCurrentPrice(ctx, marketID)
	if err != nil {
		return sdk.ZeroInt(), sdk.ZeroDec(), err
	}
	outAmount := sdk.NewDecFromInt(amount).Quo(currencyPrice.Price).TruncateInt()
	return outAmount, currencyPrice.Price, nil
}

func (k Keeper) outboundConvertToUSDWithMarketID(ctx context.Context, marketID string, amount sdkmath.Int) (sdkmath.Int, sdk.Dec, error) {
	currencyPrice, err := k.priceFeedKeeper.GetCurrentPrice(ctx, marketID)
	if err != nil {
		return sdk.ZeroInt(), sdk.ZeroDec(), err
	}
	outAmount := currencyPrice.Price.Mul(sdk.NewDecFromInt(amount)).TruncateInt()
	return outAmount, currencyPrice.Price, nil
}

func inboundConvertFromUSD(inAmount sdkmath.Int, ratio sdk.Dec) sdkmath.Int {
	outAmount := sdk.NewDecFromInt(inAmount).Quo(ratio).TruncateInt()
	return outAmount
}

func outboundConvertToUSD(inAmount sdkmath.Int, ratio sdk.Dec) sdkmath.Int {
	outAmount := ratio.Mul(sdk.NewDecFromInt(inAmount)).TruncateInt()
	return outAmount
}
