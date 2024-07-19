package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
)

// GetParams returns the params from the store
func (k Keeper) GetParams(ctx context.Context) types.Params {
	var p types.Params
	k.paramSubspace.GetParamSet(ctx, &p)
	return p
}

// legacy function GetParamsV19 returns the params from the store
func (k Keeper) GetParamsV19(ctx context.Context) types.ParamsV19 {
	var p types.ParamsV19
	k.paramSubspace.GetParamSet(ctx, &p)
	return p
}

// SetParams sets params on the store
func (k Keeper) SetParams(ctx context.Context, params types.Params) {
	k.paramSubspace.SetParamSet(ctx, &params)
}

// update spv reward tokens
func (k Keeper) SetSPVRewardTokens(ctx context.Context, poolId string, rewardTokens sdk.Coins) {
	params := k.GetParams(ctx)
	for i, el := range params.SPVRewardPeriods {
		if !el.Active {
			continue
		}
		if el.CollateralType == poolId {
			el.RewardsPerSecond = rewardTokens
			params.SPVRewardPeriods[i] = el
			break
		}
		continue
	}
	k.SetParams(ctx, params)
}

// GetJoltSupplyRewardPeriods returns the reward period with the specified collateral type if it's found in the params
func (k Keeper) GetJoltSupplyRewardPeriods(ctx context.Context, denom string) (types.MultiRewardPeriod, bool) {
	params := k.GetParams(ctx)
	for _, rp := range params.JoltSupplyRewardPeriods {
		if rp.CollateralType == denom {
			return rp, true
		}
	}
	return types.MultiRewardPeriod{}, false
}

// GetJoltBorrowRewardPeriods returns the reward period with the specified collateral type if it's found in the params
func (k Keeper) GetJoltBorrowRewardPeriods(ctx context.Context, denom string) (types.MultiRewardPeriod, bool) {
	params := k.GetParams(ctx)
	for _, rp := range params.JoltBorrowRewardPeriods {
		if rp.CollateralType == denom {
			return rp, true
		}
	}
	return types.MultiRewardPeriod{}, false
}

// GetMultiplierByDenom fetches a multiplier from the params matching the denom and name.
func (k Keeper) GetMultiplierByDenom(ctx context.Context, denom string, name string) (types.Multiplier, bool) {
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
func (k Keeper) GetClaimEnd(ctx context.Context) time.Time {
	params := k.GetParams(ctx)
	return params.ClaimEnd
}
