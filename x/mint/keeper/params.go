package keeper

import (
	"context"

	sdkmath "cosmossdk.io/math"
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
func (k Keeper) SetParams(rctx context.Context, params types.Params) {
	ctx := sdk.UnwrapSDKContext(rctx)
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) GetNodeAPY(rctx context.Context) (res sdkmath.LegacyDec) {
	ctx := sdk.UnwrapSDKContext(rctx)
	k.paramstore.Get(ctx, types.NodeSPY, &res)
	return
}

// FirstProvision returns the CurrentProvision param
func (k Keeper) FirstProvision(rctx context.Context) (res sdkmath.LegacyDec) {
	ctx := sdk.UnwrapSDKContext(rctx)
	k.paramstore.Get(ctx, types.KeyFirstProvision, &res)
	return
}

// Unit returns the CurrentProvision param
func (k Keeper) Unit(rctx context.Context) (res string) {
	ctx := sdk.UnwrapSDKContext(rctx)
	k.paramstore.Get(ctx, types.KeyUnit, &res)
	return
}
