package keeper

import (
	"fmt"
	"strings"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	var param types.Params
	k.paramstore.GetParamSet(ctx, &param)
	return param
}

// GetParams get all parameters as types.Params
func (k Keeper) GetParamsV21(ctx sdk.Context) sdkmath.Int {
	ret := k.paramstore.GetRaw(ctx, types.KeyBurnThreshold)
	out := strings.Split(string(ret), "\"")

	amt, ok := sdk.NewIntFromString(out[len(out)-2])
	if !ok {
		panic("fail to convert")
	}
	return amt
}

func (k Keeper) GetParamsFromV22(ctx sdk.Context) (sdk.Coins, []types.Moneymarket) {
	var burnThreshold sdk.Coins
	var markets []types.Moneymarket
	k.paramstore.Get(ctx, types.KeyBurnThreshold, &burnThreshold)
	k.paramstore.Get(ctx, types.KeyMoneyMarket, &markets)

	fmt.Printf(">>>burn threshold %v\n", burnThreshold)
	fmt.Printf(">>>markets %v\n", markets)
	return burnThreshold, markets
}

func (k Keeper) SetParamsFromV22(ctx sdk.Context, burnThreshold sdk.Coins, markets []types.Moneymarket) {
	k.paramstore.Set(ctx, types.KeyBurnThreshold, burnThreshold)
	k.paramstore.Set(ctx, types.KeyMoneyMarket, markets)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
