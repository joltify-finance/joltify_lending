package keeper

import (
	"context"
	"github.com/joltify-finance/joltify_lending/x/burnauction/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx context.Context) types.Params {
	return types.NewParams(
		k.Burnthreshold(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx context.Context, params types.Params) {
	k.paramstore.SetParamSet(sdk.UnwrapSDKContext(ctx), &params)
}

// Burnthreshold returns the Burnthreshold param
func (k Keeper) Burnthreshold(ctx context.Context) (res sdk.Coins) {
	k.paramstore.Get(sdk.UnwrapSDKContext(ctx), types.KeyBurnThreshold, &res)
	return
}
