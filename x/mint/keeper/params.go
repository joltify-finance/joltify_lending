package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/mint/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx context.Context) types.Params {
	return types.NewParams(
		k.FirstProvision(ctx),
		k.Unit(ctx),
		k.GetNodeAPY(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx context.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) GetNodeAPY(ctx context.Context) (res sdkmath.LegacyDec) {
	k.paramstore.Get(ctx, types.NodeSPY, &res)
	return
}

// FirstProvision returns the CurrentProvision param
func (k Keeper) FirstProvision(ctx context.Context) (res sdkmath.LegacyDec) {
	k.paramstore.Get(ctx, types.KeyFirstProvision, &res)
	return
}

// Unit returns the CurrentProvision param
func (k Keeper) Unit(ctx context.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyUnit, &res)
	return
}
