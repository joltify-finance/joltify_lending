package keeper

import (
	"time"

	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams returns the params from the store
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	var p types.Params
	k.paramSubspace.GetParamSet(ctx, &p)
	return p
}

// SetParams sets params on the store
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSubspace.SetParamSet(ctx, &params)
}

// GetUSDXMintingRewardPeriod returns the reward period with the specified collateral type if it's found in the params
func (k Keeper) GetUSDXMintingRewardPeriod(ctx sdk.Context, collateralType string) (types.RewardPeriod, bool) {
	params := k.GetParams(ctx)
	for _, rp := range params.USDXMintingRewardPeriods {
		if rp.CollateralType == collateralType {
			return rp, true
		}
	}
	return types.RewardPeriod{}, false
}

// GetJoltSupplyRewardPeriods returns the reward period with the specified collateral type if it's found in the params
func (k Keeper) GetJoltSupplyRewardPeriods(ctx sdk.Context, denom string) (types.MultiRewardPeriod, bool) {
	params := k.GetParams(ctx)
	for _, rp := range params.JoltSupplyRewardPeriods {
		if rp.CollateralType == denom {
			return rp, true
		}
	}
	return types.MultiRewardPeriod{}, false
}

// GetJoltBorrowRewardPeriods returns the reward period with the specified collateral type if it's found in the params
func (k Keeper) GetJoltBorrowRewardPeriods(ctx sdk.Context, denom string) (types.MultiRewardPeriod, bool) {
	params := k.GetParams(ctx)
	for _, rp := range params.JoltBorrowRewardPeriods {
		if rp.CollateralType == denom {
			return rp, true
		}
	}
	return types.MultiRewardPeriod{}, false
}

// GetDelegatorRewardPeriods returns the reward period with the specified collateral type if it's found in the params
func (k Keeper) GetDelegatorRewardPeriods(ctx sdk.Context, denom string) (types.MultiRewardPeriod, bool) {
	params := k.GetParams(ctx)
	for _, rp := range params.DelegatorRewardPeriods {
		if rp.CollateralType == denom {
			return rp, true
		}
	}
	return types.MultiRewardPeriod{}, false
}

// GetSavingsRewardPeriods returns the reward period with the specified collateral type if it's found in the params
func (k Keeper) GetSavingsRewardPeriods(ctx sdk.Context, denom string) (types.MultiRewardPeriod, bool) {
	params := k.GetParams(ctx)
	for _, rp := range params.SavingsRewardPeriods {
		if rp.CollateralType == denom {
			return rp, true
		}
	}
	return types.MultiRewardPeriod{}, false
}

// GetMultiplierByDenom fetches a multiplier from the params matching the denom and name.
func (k Keeper) GetMultiplierByDenom(ctx sdk.Context, denom string, name string) (types.Multiplier, bool) {
	params := k.GetParams(ctx)

	for _, dm := range params.ClaimMultipliers {
		if dm.Denom == denom {
			m, found := dm.Multipliers.Get(name)
			return m, found
		}
	}
	return types.Multiplier{}, false
}

// GetClaimEnd returns the claim end time for the params
func (k Keeper) GetClaimEnd(ctx sdk.Context) time.Time {
	params := k.GetParams(ctx)
	return params.ClaimEnd
}
