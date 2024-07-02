package keeper

import (
	"github.com/joltify-finance/joltify_lending/x/burnauction/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.Burnthreshold(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// Burnthreshold returns the Burnthreshold param
func (k Keeper) Burnthreshold(ctx sdk.Context) (res sdk.Coins) {
	k.paramstore.Get(ctx, types.KeyBurnThreshold, &res)
	return
}
