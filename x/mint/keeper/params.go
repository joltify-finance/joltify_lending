package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/mint/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.FirstProvision(ctx),
		k.Unit(ctx),
		k.GetNodeAPY(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) GetNodeAPY(ctx sdk.Context) (res sdk.Dec) {
	k.paramstore.Get(ctx, types.NodeSPY, &res)
	return
}

// FirstProvision returns the CurrentProvision param
func (k Keeper) FirstProvision(ctx sdk.Context) (res sdk.Dec) {
	k.paramstore.Get(ctx, types.KeyFirstProvision, &res)
	return
}

// Unit returns the CurrentProvision param
func (k Keeper) Unit(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyUnit, &res)
	return
}
