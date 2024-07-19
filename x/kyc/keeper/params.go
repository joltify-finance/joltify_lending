package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx context.Context) types.Params {
	var param types.Params
	k.paramstore.GetParamSet(ctx, &param)
	return param
}

// SetParams set the params
func (k Keeper) SetParams(ctx context.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) GetSubmitter(ctx context.Context) (submitters []sdk.AccAddress) {
	k.paramstore.Get(ctx, types.KeyKycSubmitter, &submitters)
	return submitters
}
