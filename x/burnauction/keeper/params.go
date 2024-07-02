package keeper

import (
	"github.com/joltify-finance/joltify_lending/x/burnauction/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.SupportCoins(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// SupportCoins returns the SupportCoins param
func (k Keeper) SupportCoins(ctx sdk.Context) (res int32) {
	k.paramstore.Get(ctx, types.KeySupportCoins, &res)
	return
}
