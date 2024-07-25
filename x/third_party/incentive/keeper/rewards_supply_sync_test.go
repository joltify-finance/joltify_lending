package keeper_test

import (
	"testing"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
	hardtypes "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

// SynchronizeHardSupplyRewardTests runs unit tests for the keeper.SynchronizeJoltSupplyReward method
type SynchronizeHardSupplyRewardTests struct {
	unitTester
}

func TestSynchronizeHardSupplyReward(t *testing.T) {
	suite.Run(t, new(SynchronizeHardSupplyRewardTests))
}

func (suite *SynchronizeHardSupplyRewardTests) TestClaimIndexesAreUpdatedWhenGlobalIndexesHaveIncreased() {
	// This is the normal case

	claim := types2.JoltLiquidityProviderClaim{
		BaseMultiClaim: types2.BaseMultiClaim{
			Owner: arbitraryAddress(),
		},
		SupplyRewardIndexes: nonEmptyMultiRewardIndexes,
	}
	suite.storeJoltClaim(claim)

	globalIndexes := increaseAllRewardFactors(nonEmptyMultiRewardIndexes)
	suite.storeGlobalSupplyIndexes(globalIndexes)
	deposit := NewJoltDepositBuilder(claim.Owner).
		WithArbitrarySourceShares(extractCollateralTypes(claim.SupplyRewardIndexes)...).
		Build()

	suite.keeper.SynchronizeJoltSupplyReward(suite.ctx, deposit)

	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Equal(globalIndexes, syncedClaim.SupplyRewardIndexes)
}

func (suite *SynchronizeHardSupplyRewardTests) TestClaimIndexesAreUnchangedWhenGlobalIndexesUnchanged() {
	// It should be safe to call SynchronizeJoltSupplyReward multiple times

	unchangingIndexes := nonEmptyMultiRewardIndexes

	claim := types2.JoltLiquidityProviderClaim{
		BaseMultiClaim: types2.BaseMultiClaim{
			Owner: arbitraryAddress(),
		},
		SupplyRewardIndexes: unchangingIndexes,
	}
	suite.storeJoltClaim(claim)

	suite.storeGlobalSupplyIndexes(unchangingIndexes)

	deposit := NewJoltDepositBuilder(claim.Owner).
		WithArbitrarySourceShares(extractCollateralTypes(unchangingIndexes)...).
		Build()

	suite.keeper.SynchronizeJoltSupplyReward(suite.ctx, deposit)

	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Equal(unchangingIndexes, syncedClaim.SupplyRewardIndexes)
}

func (suite *SynchronizeHardSupplyRewardTests) TestClaimIndexesAreUpdatedWhenNewRewardAdded() {
	// When a new reward is added (via gov) for a jolt deposit denom the user has already deposited, and the claim is synced;
	// Then the new reward's index should be added to the claim.

	claim := types2.JoltLiquidityProviderClaim{
		BaseMultiClaim: types2.BaseMultiClaim{
			Owner: arbitraryAddress(),
		},
		SupplyRewardIndexes: nonEmptyMultiRewardIndexes,
	}
	suite.storeJoltClaim(claim)

	globalIndexes := appendUniqueMultiRewardIndex(nonEmptyMultiRewardIndexes)
	suite.storeGlobalSupplyIndexes(globalIndexes)

	deposit := NewJoltDepositBuilder(claim.Owner).
		WithArbitrarySourceShares(extractCollateralTypes(globalIndexes)...).
		Build()

	suite.keeper.SynchronizeJoltSupplyReward(suite.ctx, deposit)

	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Equal(globalIndexes, syncedClaim.SupplyRewardIndexes)
}

func (suite *SynchronizeHardSupplyRewardTests) TestClaimIndexesAreUpdatedWhenNewRewardDenomAdded() {
	// When a new reward coin is added (via gov) to an already rewarded deposit denom (that the user has already deposited), and the claim is synced;
	// Then the new reward coin's index should be added to the claim.

	claim := types2.JoltLiquidityProviderClaim{
		BaseMultiClaim: types2.BaseMultiClaim{
			Owner: arbitraryAddress(),
		},
		SupplyRewardIndexes: nonEmptyMultiRewardIndexes,
	}
	suite.storeJoltClaim(claim)

	globalIndexes := appendUniqueRewardIndexToFirstItem(nonEmptyMultiRewardIndexes)
	suite.storeGlobalSupplyIndexes(globalIndexes)

	deposit := NewJoltDepositBuilder(claim.Owner).
		WithArbitrarySourceShares(extractCollateralTypes(globalIndexes)...).
		Build()

	suite.keeper.SynchronizeJoltSupplyReward(suite.ctx, deposit)

	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Equal(globalIndexes, syncedClaim.SupplyRewardIndexes)
}

func (suite *SynchronizeHardSupplyRewardTests) TestRewardIsIncrementedWhenGlobalIndexesHaveIncreased() {
	// This is the normal case
	// Given some time has passed (meaning the global indexes have increased)
	// When the claim is synced
	// The user earns rewards for the time passed

	originalReward := arbitraryCoins()

	claim := types2.JoltLiquidityProviderClaim{
		BaseMultiClaim: types2.BaseMultiClaim{
			Owner:  arbitraryAddress(),
			Reward: originalReward,
		},
		SupplyRewardIndexes: types2.MultiRewardIndexes{
			{
				CollateralType: "depositdenom",
				RewardIndexes: types2.RewardIndexes{
					{
						CollateralType: "rewarddenom",
						RewardFactor:   d("1000001000000000"),
					},
				},
			},
		},
	}
	suite.storeJoltClaim(claim)

	suite.storeGlobalSupplyIndexes(types2.MultiRewardIndexes{
		{
			CollateralType: "depositdenom",
			RewardIndexes: types2.RewardIndexes{
				{
					CollateralType: "rewarddenom",
					RewardFactor:   d("2000002000000000"),
				},
			},
		},
	})

	deposit := NewJoltDepositBuilder(claim.Owner).
		WithSourceShares("depositdenom", 1e9).
		Build()

	suite.keeper.SynchronizeJoltSupplyReward(suite.ctx, deposit)

	// new reward is (new index - old index) * deposit amount
	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Equal(
		cs(c("rewarddenom", 1_000_001_000_000)).Add(originalReward...),
		syncedClaim.Reward,
	)
}

func (suite *SynchronizeHardSupplyRewardTests) TestRewardIsIncrementedWhenNewRewardAdded() {
	// When a new reward is added (via gov) for a jolt deposit denom the user has already deposited, and the claim is synced
	// Then the user earns rewards for the time since the reward was added

	originalReward := arbitraryCoins()
	claim := types2.JoltLiquidityProviderClaim{
		BaseMultiClaim: types2.BaseMultiClaim{
			Owner:  arbitraryAddress(),
			Reward: originalReward,
		},
		SupplyRewardIndexes: types2.MultiRewardIndexes{
			{
				CollateralType: "rewarded",
				RewardIndexes: types2.RewardIndexes{
					{
						CollateralType: "reward",
						RewardFactor:   d("1000001000000000"),
					},
				},
			},
		},
	}
	suite.storeJoltClaim(claim)

	globalIndexes := types2.MultiRewardIndexes{
		{
			CollateralType: "rewarded",
			RewardIndexes: types2.RewardIndexes{
				{
					CollateralType: "reward",
					RewardFactor:   d("2000002000000000"),
				},
			},
		},
		{
			CollateralType: "newlyrewarded",
			RewardIndexes: types2.RewardIndexes{
				{
					CollateralType: "otherreward",
					// Indexes start at 0 when the reward is added by gov,
					// so this represents the syncing happening some time later.
					RewardFactor: d("1000001000000000"),
				},
			},
		},
	}
	suite.storeGlobalSupplyIndexes(globalIndexes)

	deposit := NewJoltDepositBuilder(claim.Owner).
		WithSourceShares("rewarded", 1e9).
		WithSourceShares("newlyrewarded", 1e9).
		Build()

	suite.keeper.SynchronizeJoltSupplyReward(suite.ctx, deposit)

	// new reward is (new index - old index) * deposit amount for each deposited denom
	// The old index for `newlyrewarded` isn't in the claim, so it's added starting at 0 for calculating the reward.
	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Equal(
		cs(c("otherreward", 1_000_001_000_000), c("reward", 1_000_001_000_000)).Add(originalReward...),
		syncedClaim.Reward,
	)
}

func (suite *SynchronizeHardSupplyRewardTests) TestRewardIsIncrementedWhenNewRewardDenomAdded() {
	// When a new reward coin is added (via gov) to an already rewarded deposit denom (that the user has already deposited), and the claim is synced;
	// Then the user earns rewards for the time since the reward was added

	originalReward := arbitraryCoins()
	claim := types2.JoltLiquidityProviderClaim{
		BaseMultiClaim: types2.BaseMultiClaim{
			Owner:  arbitraryAddress(),
			Reward: originalReward,
		},
		SupplyRewardIndexes: types2.MultiRewardIndexes{
			{
				CollateralType: "deposited",
				RewardIndexes: types2.RewardIndexes{
					{
						CollateralType: "reward",
						RewardFactor:   d("1000001000000000"),
					},
				},
			},
		},
	}
	suite.storeJoltClaim(claim)

	globalIndexes := types2.MultiRewardIndexes{
		{
			CollateralType: "deposited",
			RewardIndexes: types2.RewardIndexes{
				{
					CollateralType: "reward",
					RewardFactor:   d("2000.002").MulInt64(1e12),
				},
				{
					CollateralType: "otherreward",
					// Indexes start at 0 when the reward is added by gov,
					// so this represents the syncing happening some time later.
					RewardFactor: d("1000.001").MulInt64(1e12),
				},
			},
		},
	}
	suite.storeGlobalSupplyIndexes(globalIndexes)

	deposit := NewJoltDepositBuilder(claim.Owner).
		WithSourceShares("deposited", 1e9).
		Build()

	suite.keeper.SynchronizeJoltSupplyReward(suite.ctx, deposit)

	// new reward is (new index - old index) * deposit amount for each deposited denom
	// The old index for `otherreward` isn't in the claim, so it's added starting at 0 for calculating the reward.
	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Equal(
		cs(c("reward", 1_000_001_000_000), c("otherreward", 1_000_001_000_000)).Add(originalReward...),
		syncedClaim.Reward,
	)
}

// JoltDepositBuilder is a tool for creating a jolt deposit in tests.
// The builder inherits from jolt.Deposit, so fields can be accessed directly if a helper method doesn't exist.
type JoltDepositBuilder struct {
	hardtypes.Deposit
}

// NewJoltDepositBuilder creates a JoltDepositBuilder containing an empty deposit.
func NewJoltDepositBuilder(depositor sdk.AccAddress) JoltDepositBuilder {
	return JoltDepositBuilder{
		Deposit: hardtypes.Deposit{
			Depositor: depositor,
		},
	}
}

// Build assembles and returns the final deposit.
func (builder JoltDepositBuilder) Build() hardtypes.Deposit { return builder.Deposit }

// WithSourceShares adds a deposit amount and factor such that the source shares for this deposit is equal to specified.
// With a factor of 1, the deposit amount is the source shares. This picks an arbitrary factor to ensure factors are accounted for in production code.
func (builder JoltDepositBuilder) WithSourceShares(denom string, shares int64) JoltDepositBuilder {
	if !builder.Amount.AmountOf(denom).Equal(sdkmath.ZeroInt()) {
		panic("adding to amount with existing denom not implemented")
	}
	if _, f := builder.Index.GetInterestFactor(denom); f {
		panic("adding to indexes with existing denom not implemented")
	}

	// pick arbitrary factor
	factor := sdk.MustNewDecFromStr("2")

	// Calculate deposit amount that would equal the requested source shares given the above factor.
	amt := sdk.NewInt(shares).Mul(factor.RoundInt())

	builder.Amount = builder.Amount.Add(sdk.NewCoin(denom, amt))
	builder.Index = builder.Index.SetInterestFactor(denom, factor)
	return builder
}

// WithArbitrarySourceShares adds arbitrary deposit amounts and indexes for each specified denom.
func (builder JoltDepositBuilder) WithArbitrarySourceShares(denoms ...string) JoltDepositBuilder {
	const arbitraryShares = 1e9
	var builderHandler JoltDepositBuilder
	builderHandler = builder
	for _, denom := range denoms {
		builderHandler = builderHandler.WithSourceShares(denom, arbitraryShares)
	}
	return builderHandler
}
