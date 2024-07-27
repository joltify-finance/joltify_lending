package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func c(denom string, amount int64) sdk.Coin { return sdkmath.NewInt64Coin(denom, amount) }
func cs(coins ...sdk.Coin) sdk.Coins        { return sdk.NewCoins(coins...) }
func is(ns ...int64) (is []sdkmath.Int) {
	for _, n := range ns {
		is = append(is, sdkmath.NewInt(n))
	}
	return
}
