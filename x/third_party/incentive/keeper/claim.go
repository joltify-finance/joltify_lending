package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
)

// ClaimJoltReward pays out funds from a claim to a receiver account.
// Rewards are removed from a claim and paid out according to the multiplier, which reduces the reward amount in exchange for shorter vesting times.
func (k Keeper) ClaimJoltReward(ctx sdk.Context, owner, receiver sdk.AccAddress, denom string, multiplierName string) error {
	multiplier, found := k.GetMultiplierByDenom(ctx, denom, multiplierName)
	if !found {
		return errorsmod.Wrapf(types.ErrInvalidMultiplier, "denom '%s' has no multiplier '%s'", denom, multiplierName)
	}

	claimEnd := k.GetClaimEnd(ctx)

	if ctx.BlockTime().After(claimEnd) {
		return errorsmod.Wrapf(types.ErrClaimExpired, "block time %s > claim end time %s", ctx.BlockTime(), claimEnd)
	}

	k.SynchronizeJoltLiquidityProviderClaim(ctx, owner)

	syncedClaim, found := k.GetJoltLiquidityProviderClaim(ctx, owner)
	if !found {
		return sdkerrors.Wrapf(types.ErrClaimNotFound, "address: %s", owner)
	}

	amt := syncedClaim.Reward.AmountOf(denom)

	claimingCoins := sdk.NewCoins(sdk.NewCoin(denom, amt))
	rewardCoins := sdk.NewCoins(sdk.NewCoin(denom, sdk.NewDecFromInt(amt).Mul(multiplier.Factor).RoundInt()))
	if rewardCoins.IsZero() {
		return types.ErrZeroClaim
	}
	length := k.GetPeriodLength(ctx.BlockTime(), multiplier.MonthsLockup)

	err := k.SendTimeLockedCoinsToAccount(ctx, types.IncentiveMacc, receiver, rewardCoins, length)
	if err != nil {
		return err
	}

	// remove claimed coins (NOT reward coins)
	syncedClaim.Reward = syncedClaim.Reward.Sub(claimingCoins...)
	k.SetJoltLiquidityProviderClaim(ctx, syncedClaim)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeClaim,
			sdk.NewAttribute(types.AttributeKeyClaimedBy, owner.String()),
			sdk.NewAttribute(types.AttributeKeyClaimAmount, claimingCoins.String()),
			sdk.NewAttribute(types.AttributeKeyClaimType, syncedClaim.GetType()),
		),
	)
	return nil
}

// ClaimSwapReward pays out funds from a claim to a receiver account.
// Rewards are removed from a claim and paid out according to the multiplier, which reduces the reward amount in exchange for shorter vesting times.
func (k Keeper) ClaimSwapReward(ctx sdk.Context, owner, receiver sdk.AccAddress, denom string, multiplierName string) error {
	multiplier, found := k.GetMultiplierByDenom(ctx, denom, multiplierName)
	if !found {
		return errorsmod.Wrapf(types.ErrInvalidMultiplier, "denom '%s' has no multiplier '%s'", denom, multiplierName)
	}

	claimEnd := k.GetClaimEnd(ctx)

	if ctx.BlockTime().After(claimEnd) {
		return errorsmod.Wrapf(types.ErrClaimExpired, "block time %s > claim end time %s", ctx.BlockTime(), claimEnd)
	}

	syncedClaim, found := k.GetSynchronizedSwapClaim(ctx, owner)
	if !found {
		return errorsmod.Wrapf(types.ErrClaimNotFound, "address: %s", owner)
	}

	amt := syncedClaim.Reward.AmountOf(denom)

	claimingCoins := sdk.NewCoins(sdk.NewCoin(denom, amt))
	rewardCoins := sdk.NewCoins(sdk.NewCoin(denom, sdk.NewDecFromInt(amt).Mul(multiplier.Factor).RoundInt()))
	if rewardCoins.IsZero() {
		return types.ErrZeroClaim
	}
	length := k.GetPeriodLength(ctx.BlockTime(), multiplier.MonthsLockup)

	err := k.SendTimeLockedCoinsToAccount(ctx, types.IncentiveMacc, receiver, rewardCoins, length)
	if err != nil {
		return err
	}

	// remove claimed coins (NOT reward coins)
	syncedClaim.Reward = syncedClaim.Reward.Sub(claimingCoins...)
	k.SetSwapClaim(ctx, syncedClaim)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeClaim,
			sdk.NewAttribute(types.AttributeKeyClaimedBy, owner.String()),
			sdk.NewAttribute(types.AttributeKeyClaimAmount, claimingCoins.String()),
			sdk.NewAttribute(types.AttributeKeyClaimType, syncedClaim.GetType()),
		),
	)
	return nil
}
