package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	vaultmoduletypes "github.com/joltify-finance/joltify_lending/x/vault/types"
)

func (k Keeper) BlockChurnInterval(ctx context.Context) (res int64) {
	k.paramstore.Get(ctx, vaultmoduletypes.KeyBlockChurnInterval, &res)
	return
}

func (k Keeper) Power(ctx context.Context) (res int64) {
	k.paramstore.Get(ctx, vaultmoduletypes.KeyPower, &res)
	return
}

func (k Keeper) Step(ctx context.Context) (res int64) {
	k.paramstore.Get(ctx, vaultmoduletypes.KeyStep, &res)
	return
}

func (k Keeper) CandidateRatio(ctx context.Context) (res sdk.Dec) {
	k.paramstore.Get(ctx, vaultmoduletypes.KeyRatio, &res)
	return
}

func (k Keeper) TargetQuota(ctx context.Context) (res sdk.Coins) {
	k.paramstore.Get(ctx, vaultmoduletypes.KeyTargetQuota, &res)
	return
}

func (k Keeper) QuotaHistoryLength(ctx context.Context) (res int32) {
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
func (k Keeper) SetParams(ctx context.Context, params vaultmoduletypes.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
