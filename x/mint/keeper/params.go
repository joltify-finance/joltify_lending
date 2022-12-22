package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/mint/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.FirstProvision(ctx),
		k.CurrentProvision(ctx),
		k.Unit(ctx),
		k.CommunityProvision(ctx),
		k.HalfCount(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// CurrentProvision returns the CurrentProvision param
func (k Keeper) CurrentProvision(ctx sdk.Context) (res sdk.Dec) {
	k.paramstore.Get(ctx, types.KeyCurrentProvision, &res)
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

// CommunityProvision returns the CurrentProvision param
func (k Keeper) CommunityProvision(ctx sdk.Context) (res sdk.Dec) {
	k.paramstore.Get(ctx, types.KeyCommunityProvision, &res)
	return
}

// HalfCount returns the half count param
func (k Keeper) HalfCount(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyHalfCount, &res)
	return
}
