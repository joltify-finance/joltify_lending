package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
	hardtypes "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"
)

// AccumulateJoltSupplyRewards calculates new rewards to distribute this block and updates the global indexes to reflect this.
// The provided rewardPeriod must be valid to avoid panics in calculating time durations.
func (k Keeper) AccumulateJoltSupplyRewards(ctx sdk.Context, rewardPeriod types2.MultiRewardPeriod) {
	previousAccrualTime, found := k.GetPreviousJoltSupplyRewardAccrualTime(ctx, rewardPeriod.CollateralType)
	if !found {
		previousAccrualTime = ctx.BlockTime()
	}

	indexes, found := k.GetJoltSupplyRewardIndexes(ctx, rewardPeriod.CollateralType)
	if !found {
		indexes = types2.RewardIndexes{}
	}

	acc := types2.NewAccumulator(previousAccrualTime, indexes)

	totalSource := k.getJoltSupplyTotalSourceShares(ctx, rewardPeriod.CollateralType)

	acc.Accumulate(rewardPeriod, totalSource, ctx.BlockTime())

	k.SetPreviousJoltSupplyRewardAccrualTime(ctx, rewardPeriod.CollateralType, acc.PreviousAccumulationTime)
	if len(acc.Indexes) > 0 {
		// the store panics when setting empty or nil indexes
		k.SetJoltSupplyRewardIndexes(ctx, rewardPeriod.CollateralType, acc.Indexes)
	}
}

// getJoltSupplyTotalSourceShares fetches the sum of all source shares for a supply reward.
// In the case of jolt supply, this is the total supplied divided by the supply interest factor.
// This gives the "pre interest" value of the total supplied.
func (k Keeper) getJoltSupplyTotalSourceShares(ctx sdk.Context, denom string) sdk.Dec {
	totalSuppliedCoins, found := k.joltKeeper.GetSuppliedCoins(ctx)
	if !found {
		// assume no coins have been supplied
		totalSuppliedCoins = sdk.NewCoins()
	}
	totalSupplied := totalSuppliedCoins.AmountOf(denom)

	interestFactor, found := k.joltKeeper.GetSupplyInterestFactor(ctx, denom)
	if !found {
		// assume nothing has been borrowed so the factor starts at it's default value
		interestFactor = sdk.OneDec()
	}

	// return supplied/factor to get the "pre interest" value of the current total supplied
	return sdk.NewDecFromInt(totalSupplied).Quo(interestFactor)
}

// InitializeJoltSupplyReward initializes the supply-side of a jolt liquidity provider claim
// by creating the claim and setting the supply reward factor index
func (k Keeper) InitializeJoltSupplyReward(ctx sdk.Context, deposit hardtypes.Deposit) {
	claim, found := k.GetJoltLiquidityProviderClaim(ctx, deposit.Depositor)
	if !found {
		claim = types2.NewJoltLiquidityProviderClaim(deposit.Depositor, sdk.Coins{}, nil, nil)
	}

	var supplyRewardIndexes types2.MultiRewardIndexes
	for _, coin := range deposit.Amount {
		globalRewardIndexes, found := k.GetJoltSupplyRewardIndexes(ctx, coin.Denom)
		if !found {
			globalRewardIndexes = types2.RewardIndexes{}
		}
		supplyRewardIndexes = supplyRewardIndexes.With(coin.Denom, globalRewardIndexes)
	}

	claim.SupplyRewardIndexes = supplyRewardIndexes
	k.SetJoltLiquidityProviderClaim(ctx, claim)
}

// SynchronizeJoltSupplyReward updates the claim object by adding any accumulated rewards
// and updating the reward index value
func (k Keeper) SynchronizeJoltSupplyReward(ctx sdk.Context, deposit hardtypes.Deposit) {
	claim, found := k.GetJoltLiquidityProviderClaim(ctx, deposit.Depositor)
	if !found {
		return
	}

	// Source shares for jolt deposits is their normalized deposit amount
	normalizedDeposit, err := deposit.NormalizedDeposit()
	if err != nil {
		panic(fmt.Sprintf("during deposit reward sync, could not get normalized deposit for %s: %s", deposit.Depositor, err.Error()))
	}

	for _, normedDeposit := range normalizedDeposit {
		claim = k.synchronizeSingleJoltSupplyReward(ctx, claim, normedDeposit.Denom, normedDeposit.Amount)
	}
	k.SetJoltLiquidityProviderClaim(ctx, claim)
}

// synchronizeSingleJoltSupplyReward synchronizes a single rewarded supply denom in a jolt claim.
// It returns the claim without setting in the store.
// The public methods for accessing and modifying claims are preferred over this one. Direct modification of claims is easy to get wrong.
func (k Keeper) synchronizeSingleJoltSupplyReward(ctx sdk.Context, claim types2.JoltLiquidityProviderClaim, denom string, sourceShares sdk.Dec) types2.JoltLiquidityProviderClaim {
	globalRewardIndexes, found := k.GetJoltSupplyRewardIndexes(ctx, denom)
	if !found {
		// The global factor is only not found if
		// - the supply denom has not started accumulating rewards yet (either there is no reward specified in params, or the reward start time hasn't been hit)
		// - OR it was wrongly deleted from state (factors should never be removed while unsynced claims exist)
		// If not found we could either skip this sync, or assume the global factor is zero.
		// Skipping will avoid storing unnecessary factors in the claim for non rewarded denoms.
		// And in the event a global factor is wrongly deleted, it will avoid this function panicking when calculating rewards.
		return claim
	}

	userRewardIndexes, found := claim.SupplyRewardIndexes.Get(denom)
	if !found {
		// Normally the reward indexes should always be found.
		// But if a denom was not rewarded then becomes rewarded (ie a reward period is added to params), then the indexes will be missing from claims for that supplied denom.
		// So given the reward period was just added, assume the starting value for any global reward indexes, which is an empty slice.
		userRewardIndexes = types2.RewardIndexes{}
	}

	newRewards, err := k.CalculateRewards(userRewardIndexes, globalRewardIndexes, sourceShares)
	if err != nil {
		// Global reward factors should never decrease, as it would lead to a negative update to claim.Rewards.
		// This panics if a global reward factor decreases or disappears between the old and new indexes.
		panic(fmt.Sprintf("corrupted global reward indexes found: %v", err))
	}

	claim.Reward = claim.Reward.Add(newRewards...)
	claim.SupplyRewardIndexes = claim.SupplyRewardIndexes.With(denom, globalRewardIndexes)

	return claim
}

// UpdateJoltSupplyIndexDenoms adds any new deposit denoms to the claim's supply reward index
func (k Keeper) UpdateJoltSupplyIndexDenoms(ctx sdk.Context, deposit hardtypes.Deposit) {
	claim, found := k.GetJoltLiquidityProviderClaim(ctx, deposit.Depositor)
	if !found {
		claim = types2.NewJoltLiquidityProviderClaim(deposit.Depositor, sdk.Coins{}, nil, nil)
	}

	depositDenoms := getDenoms(deposit.Amount)
	supplyRewardIndexDenoms := claim.SupplyRewardIndexes.GetCollateralTypes()

	supplyRewardIndexes := claim.SupplyRewardIndexes

	// Create a new multi-reward index in the claim for every new deposit denom
	uniqueDepositDenoms := setDifference(depositDenoms, supplyRewardIndexDenoms)

	for _, denom := range uniqueDepositDenoms {
		globalSupplyRewardIndexes, found := k.GetJoltSupplyRewardIndexes(ctx, denom)
		if !found {
			globalSupplyRewardIndexes = types2.RewardIndexes{}
		}
		supplyRewardIndexes = supplyRewardIndexes.With(denom, globalSupplyRewardIndexes)
	}

	// Delete multi-reward index from claim if the collateral type is no longer deposited
	uniqueSupplyRewardDenoms := setDifference(supplyRewardIndexDenoms, depositDenoms)

	for _, denom := range uniqueSupplyRewardDenoms {
		supplyRewardIndexes = supplyRewardIndexes.RemoveRewardIndex(denom)
	}

	claim.SupplyRewardIndexes = supplyRewardIndexes
	k.SetJoltLiquidityProviderClaim(ctx, claim)
}

// SynchronizeJoltLiquidityProviderClaim adds any accumulated rewards
func (k Keeper) SynchronizeJoltLiquidityProviderClaim(ctx sdk.Context, owner sdk.AccAddress) {
	// Synchronize any jolt liquidity supply-side rewards
	deposit, foundDeposit := k.joltKeeper.GetDeposit(ctx, owner)
	if foundDeposit {
		k.SynchronizeJoltSupplyReward(ctx, deposit)
	}

	// Synchronize any jolt liquidity borrow-side rewards
	borrow, foundBorrow := k.joltKeeper.GetBorrow(ctx, owner)
	if foundBorrow {
		k.SynchronizeJoltBorrowReward(ctx, borrow)
	}
}

// SimulateJoltSynchronization calculates a user's outstanding jolt rewards by simulating reward synchronization
func (k Keeper) SimulateJoltSynchronization(ctx sdk.Context, claim types2.JoltLiquidityProviderClaim) types2.JoltLiquidityProviderClaim {
	// 1. Simulate Hard supply-side rewards
	for _, ri := range claim.SupplyRewardIndexes {
		globalRewardIndexes, foundGlobalRewardIndexes := k.GetJoltSupplyRewardIndexes(ctx, ri.CollateralType)
		if !foundGlobalRewardIndexes {
			continue
		}

		userRewardIndexes, foundUserRewardIndexes := claim.SupplyRewardIndexes.GetRewardIndex(ri.CollateralType)
		if !foundUserRewardIndexes {
			continue
		}

		userRewardIndexIndex, foundUserRewardIndexIndex := claim.SupplyRewardIndexes.GetRewardIndexIndex(ri.CollateralType)
		if !foundUserRewardIndexIndex {
			continue
		}

		for _, globalRewardIndex := range globalRewardIndexes {
			userRewardIndex, foundUserRewardIndex := userRewardIndexes.RewardIndexes.GetRewardIndex(globalRewardIndex.CollateralType)
			if !foundUserRewardIndex {
				userRewardIndex = types2.NewRewardIndex(globalRewardIndex.CollateralType, sdk.ZeroDec())
				userRewardIndexes.RewardIndexes = append(userRewardIndexes.RewardIndexes, userRewardIndex)
				claim.SupplyRewardIndexes[userRewardIndexIndex].RewardIndexes = append(claim.SupplyRewardIndexes[userRewardIndexIndex].RewardIndexes, userRewardIndex)
			}

			globalRewardFactor := globalRewardIndex.RewardFactor
			userRewardFactor := userRewardIndex.RewardFactor
			rewardsAccumulatedFactor := globalRewardFactor.Sub(userRewardFactor)
			if rewardsAccumulatedFactor.IsZero() {
				continue
			}
			deposit, found := k.joltKeeper.GetDeposit(ctx, claim.GetOwner())
			if !found {
				continue
			}
			newRewardsAmount := rewardsAccumulatedFactor.Mul(sdk.NewDecFromInt(deposit.Amount.AmountOf(ri.CollateralType))).QuoInt(sdk.NewInt(1e12)).RoundInt()
			if newRewardsAmount.IsZero() || newRewardsAmount.IsNegative() {
				continue
			}

			factorIndex, foundFactorIndex := userRewardIndexes.RewardIndexes.GetFactorIndex(globalRewardIndex.CollateralType)
			if !foundFactorIndex {
				continue
			}
			claim.SupplyRewardIndexes[userRewardIndexIndex].RewardIndexes[factorIndex].RewardFactor = globalRewardIndex.RewardFactor
			newRewardsCoin := sdk.NewCoin(userRewardIndex.CollateralType, newRewardsAmount)
			claim.Reward = claim.Reward.Add(newRewardsCoin)
		}
	}

	// 2. Simulate Hard borrow-side rewards
	for _, ri := range claim.BorrowRewardIndexes {
		globalRewardIndexes, foundGlobalRewardIndexes := k.GetJoltBorrowRewardIndexes(ctx, ri.CollateralType)
		if !foundGlobalRewardIndexes {
			continue
		}

		userRewardIndexes, foundUserRewardIndexes := claim.BorrowRewardIndexes.GetRewardIndex(ri.CollateralType)
		if !foundUserRewardIndexes {
			continue
		}

		userRewardIndexIndex, foundUserRewardIndexIndex := claim.BorrowRewardIndexes.GetRewardIndexIndex(ri.CollateralType)
		if !foundUserRewardIndexIndex {
			continue
		}

		for _, globalRewardIndex := range globalRewardIndexes {
			userRewardIndex, foundUserRewardIndex := userRewardIndexes.RewardIndexes.GetRewardIndex(globalRewardIndex.CollateralType)
			if !foundUserRewardIndex {
				userRewardIndex = types2.NewRewardIndex(globalRewardIndex.CollateralType, sdk.ZeroDec())
				userRewardIndexes.RewardIndexes = append(userRewardIndexes.RewardIndexes, userRewardIndex)
				claim.BorrowRewardIndexes[userRewardIndexIndex].RewardIndexes = append(claim.BorrowRewardIndexes[userRewardIndexIndex].RewardIndexes, userRewardIndex)
			}

			globalRewardFactor := globalRewardIndex.RewardFactor
			userRewardFactor := userRewardIndex.RewardFactor
			rewardsAccumulatedFactor := globalRewardFactor.Sub(userRewardFactor)
			if rewardsAccumulatedFactor.IsZero() {
				continue
			}
			borrow, found := k.joltKeeper.GetBorrow(ctx, claim.GetOwner())
			if !found {
				continue
			}
			newRewardsAmount := rewardsAccumulatedFactor.Mul(sdk.NewDecFromInt(borrow.Amount.AmountOf(ri.CollateralType))).QuoInt(sdk.NewInt(1e12)).TruncateInt()
			if newRewardsAmount.IsZero() || newRewardsAmount.IsNegative() {
				continue
			}

			factorIndex, foundFactorIndex := userRewardIndexes.RewardIndexes.GetFactorIndex(globalRewardIndex.CollateralType)
			if !foundFactorIndex {
				continue
			}
			claim.BorrowRewardIndexes[userRewardIndexIndex].RewardIndexes[factorIndex].RewardFactor = globalRewardIndex.RewardFactor
			newRewardsCoin := sdk.NewCoin(userRewardIndex.CollateralType, newRewardsAmount)
			claim.Reward = claim.Reward.Add(newRewardsCoin)
		}
	}

	return claim
}

// Set setDifference: A - B
func setDifference(a, b []string) (diff []string) {
	m := make(map[string]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}

func getDenoms(coins sdk.Coins) []string {
	var denoms []string
	for _, coin := range coins {
		denoms = append(denoms, coin.Denom)
	}
	return denoms
}
