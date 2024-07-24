package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"
)

// GetParams returns the params from the store
func (k Keeper) GetParams(ctx context.Context) types.Params {
	var p types.Params
	k.paramSubspace.GetParamSet(ctx, &p)
	return p
}

// SetParams sets params on the store
func (k Keeper) SetParams(ctx context.Context, params types.Params) {
	k.paramSubspace.SetParamSet(ctx, &params)
}

// GetMinimumBorrowUSDValue returns the minimum borrow USD value
func (k Keeper) GetMinimumBorrowUSDValue(ctx context.Context) sdkmath.LegacyDec {
	params := k.GetParams(ctx)
	return params.MinimumBorrowUSDValue
}
