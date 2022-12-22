package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
)

// ClaimJoltReward pays out funds from a claim to a receiver account.
// Rewards are removed from a claim and paid out according to the multiplier, which reduces the reward amount in exchange for shorter vesting times.
func (k Keeper) ClaimJoltReward(ctx sdk.Context, owner, receiver sdk.AccAddress, denom string, multiplierName string) error {
	multiplier, found := k.GetMultiplierByDenom(ctx, denom, multiplierName)
	if !found {
		return sdkerrors.Wrapf(types2.ErrInvalidMultiplier, "denom '%s' has no multiplier '%s'", denom, multiplierName)
	}

	claimEnd := k.GetClaimEnd(ctx)

	if ctx.BlockTime().After(claimEnd) {
		return sdkerrors.Wrapf(types2.ErrClaimExpired, "block time %s > claim end time %s", ctx.BlockTime(), claimEnd)
	}

	k.SynchronizeJoltLiquidityProviderClaim(ctx, owner)

	syncedClaim, found := k.GetJoltLiquidityProviderClaim(ctx, owner)
	if !found {
		return sdkerrors.Wrapf(types2.ErrClaimNotFound, "address: %s", owner)
	}

	amt := syncedClaim.Reward.AmountOf(denom)

	claimingCoins := sdk.NewCoins(sdk.NewCoin(denom, amt))
	rewardCoins := sdk.NewCoins(sdk.NewCoin(denom, amt.ToDec().Mul(multiplier.Factor).RoundInt()))
	if rewardCoins.IsZero() {
		return types2.ErrZeroClaim
	}
	length := k.GetPeriodLength(ctx.BlockTime(), multiplier.MonthsLockup)

	err := k.SendTimeLockedCoinsToAccount(ctx, types2.IncentiveMacc, receiver, rewardCoins, length)
	if err != nil {
		return err
	}

	// remove claimed coins (NOT reward coins)
	syncedClaim.Reward = syncedClaim.Reward.Sub(claimingCoins)
	k.SetJoltLiquidityProviderClaim(ctx, syncedClaim)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types2.EventTypeClaim,
			sdk.NewAttribute(types2.AttributeKeyClaimedBy, owner.String()),
			sdk.NewAttribute(types2.AttributeKeyClaimAmount, claimingCoins.String()),
			sdk.NewAttribute(types2.AttributeKeyClaimType, syncedClaim.GetType()),
		),
	)
	return nil
}
