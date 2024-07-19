package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/auction/types"
)

func (k Keeper) SetParams(ctx context.Context, params types.Params) {
	k.paramSubspace.SetParamSet(ctx, &params)
}

func (k Keeper) GetParams(ctx context.Context) (params types.Params) {
	k.paramSubspace.GetParamSet(ctx, &params)
	return
}
