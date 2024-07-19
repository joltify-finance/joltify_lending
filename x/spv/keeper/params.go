package keeper

import (
	"context"
	"strings"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(rctx context.Context) types.Params {
	var param types.Params
	ctx := sdk.UnwrapSDKContext(rctx)
	k.paramstore.GetParamSet(ctx, &param)
	return param
}

// GetParams get all parameters as types.Params
func (k Keeper) GetParamsV21(rctx context.Context) sdkmath.Int {
	ctx := sdk.UnwrapSDKContext(rctx)
	ret := k.paramstore.GetRaw(ctx, types.KeyBurnThreshold)
	out := strings.Split(string(ret), "\"")

	amt, ok := sdkmath.NewIntFromString(out[len(out)-2])
	if !ok {
		panic("fail to convert")
	}
	return amt
}

// SetParams set the params
func (k Keeper) SetParams(rctx context.Context, params types.Params) {
	ctx := sdk.UnwrapSDKContext(rctx)
	k.paramstore.SetParamSet(ctx, &params)
}
