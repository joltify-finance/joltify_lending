package keeper

import (
	"strings"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx context.Context) types.Params {
	var param types.Params
	k.paramstore.GetParamSet(ctx, &param)
	return param
}

// GetParams get all parameters as types.Params
func (k Keeper) GetParamsV21(ctx context.Context) sdkmath.Int {
	ret := k.paramstore.GetRaw(ctx, types.KeyBurnThreshold)
	out := strings.Split(string(ret), "\"")

	amt, ok := sdkmath.NewIntFromString(out[len(out)-2])
	if !ok {
		panic("fail to convert")
	}
	return amt
}

// SetParams set the params
func (k Keeper) SetParams(ctx context.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
