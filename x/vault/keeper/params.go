package keeper

import (
	"context"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	vaultmoduletypes "github.com/joltify-finance/joltify_lending/x/vault/types"
)

func (k Keeper) BlockChurnInterval(rctx context.Context) (res int64) {
	ctx := sdk.UnwrapSDKContext(rctx)
	k.paramstore.Get(ctx, vaultmoduletypes.KeyBlockChurnInterval, &res)
	return
}

func (k Keeper) Power(rctx context.Context) (res int64) {
	ctx := sdk.UnwrapSDKContext(rctx)
	k.paramstore.Get(ctx, vaultmoduletypes.KeyPower, &res)
	return
}

func (k Keeper) Step(rctx context.Context) (res int64) {
	ctx := sdk.UnwrapSDKContext(rctx)
	k.paramstore.Get(ctx, vaultmoduletypes.KeyStep, &res)
	return
}

func (k Keeper) CandidateRatio(rctx context.Context) (res sdkmath.LegacyDec) {
	ctx := sdk.UnwrapSDKContext(rctx)
	k.paramstore.Get(ctx, vaultmoduletypes.KeyRatio, &res)
	return
}

func (k Keeper) TargetQuota(rctx context.Context) (res sdk.Coins) {
	ctx := sdk.UnwrapSDKContext(rctx)
	k.paramstore.Get(ctx, vaultmoduletypes.KeyTargetQuota, &res)
	return
}

func (k Keeper) QuotaHistoryLength(rctx context.Context) (res int32) {
	ctx := sdk.UnwrapSDKContext(rctx)
	k.paramstore.Get(ctx, vaultmoduletypes.KeyHistoryLength, &res)
	return
}

// GetParams Get all parameteras as types.Params
func (k Keeper) GetParams(ctx context.Context) vaultmoduletypes.Params {
	return vaultmoduletypes.NewParams(
		k.BlockChurnInterval(ctx),
		k.Power(ctx),
		k.Step(ctx),
		k.CandidateRatio(ctx),
		k.TargetQuota(ctx),
		k.QuotaHistoryLength(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(rctx context.Context, params vaultmoduletypes.Params) {
	ctx := sdk.UnwrapSDKContext(rctx)
	k.paramstore.SetParamSet(ctx, &params)
}
