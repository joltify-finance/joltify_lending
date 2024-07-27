package keeper_test

import (
	"errors"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
	jolttypes "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// SynchronizeJoltBorrowRewardTests runs unit tests for the keeper.SynchronizeJoltBorrowReward method
type SynchronizeJoltBorrowRewardTests struct {
	unitTester
}

func TestSynchronizeJoltBorrowReward(t *testing.T) {
	suite.Run(t, new(SynchronizeJoltBorrowRewardTests))
}

func (suite *SynchronizeJoltBorrowRewardTests) TestClaimIndexesAreUpdatedWhenGlobalIndexesHaveIncreased() {
	// This is the normal case

	claim := types2.JoltLiquidityProviderClaim{
		BaseMultiClaim: types2.BaseMultiClaim{
			Owner: arbitraryAddress(),
		},
		BorrowRewardIndexes: nonEmptyMultiRewardIndexes,
	}
	suite.storeJoltClaim(claim)

	globalIndexes := increaseAllRewardFactors(nonEmptyMultiRewardIndexes)
	suite.storeGlobalBorrowIndexes(globalIndexes)

	borrow := NewBorrowBuilder(claim.Owner).
		WithArbitrarySourceShares(extractCollateralTypes(claim.BorrowRewardIndexes)...).
		Build()

	suite.keeper.SynchronizeJoltBorrowReward(suite.ctx, borrow)

	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Equal(globalIndexes, syncedClaim.BorrowRewardIndexes)
}

func (suite *SynchronizeJoltBorrowRewardTests) TestClaimIndexesAreUnchangedWhenGlobalIndexesUnchanged() {
	// It should be safe to call SynchronizeJoltBorrowReward multiple times

	unchangingIndexes := nonEmptyMultiRewardIndexes

	claim := types2.JoltLiquidityProviderClaim{
		BaseMultiClaim: types2.BaseMultiClaim{
			Owner: arbitraryAddress(),
		},
		BorrowRewardIndexes: unchangingIndexes,
	}
	suite.storeJoltClaim(claim)

	suite.storeGlobalBorrowIndexes(unchangingIndexes)

	borrow := NewBorrowBuilder(claim.Owner).
		WithArbitrarySourceShares(extractCollateralTypes(unchangingIndexes)...).
		Build()

	suite.keeper.SynchronizeJoltBorrowReward(suite.ctx, borrow)

	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Equal(unchangingIndexes, syncedClaim.BorrowRewardIndexes)
}

func (suite *SynchronizeJoltBorrowRewardTests) TestClaimIndexesAreUpdatedWhenNewRewardAdded() {
	// When a new reward is added (via gov) for a jolt borrow denom the user has already borrowed, and the claim is synced;
	// Then the new reward's index should be added to the claim.

	claim := types2.JoltLiquidityProviderClaim{
		BaseMultiClaim: types2.BaseMultiClaim{
			Owner: arbitraryAddress(),
		},
		BorrowRewardIndexes: nonEmptyMultiRewardIndexes,
	}
	suite.storeJoltClaim(claim)

	globalIndexes := appendUniqueMultiRewardIndex(nonEmptyMultiRewardIndexes)
	suite.storeGlobalBorrowIndexes(globalIndexes)

	borrow := NewBorrowBuilder(claim.Owner).
		WithArbitrarySourceShares(extractCollateralTypes(globalIndexes)...).
		Build()

	suite.keeper.SynchronizeJoltBorrowReward(suite.ctx, borrow)

	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Equal(globalIndexes, syncedClaim.BorrowRewardIndexes)
}

func (suite *SynchronizeJoltBorrowRewardTests) TestClaimIndexesAreUpdatedWhenNewRewardDenomAdded() {
	// When a new reward coin is added (via gov) to an already rewarded borrow denom (that the user has already borrowed), and the claim is synced;
	// Then the new reward coin's index should be added to the claim.

	claim := types2.JoltLiquidityProviderClaim{
		BaseMultiClaim: types2.BaseMultiClaim{
			Owner: arbitraryAddress(),
		},
		BorrowRewardIndexes: nonEmptyMultiRewardIndexes,
	}
	suite.storeJoltClaim(claim)

	globalIndexes := appendUniqueRewardIndexToFirstItem(nonEmptyMultiRewardIndexes)
	suite.storeGlobalBorrowIndexes(globalIndexes)

	borrow := NewBorrowBuilder(claim.Owner).
		WithArbitrarySourceShares(extractCollateralTypes(globalIndexes)...).
		Build()

	suite.keeper.SynchronizeJoltBorrowReward(suite.ctx, borrow)

	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Equal(globalIndexes, syncedClaim.BorrowRewardIndexes)
}

func (suite *SynchronizeJoltBorrowRewardTests) TestRewardIsIncrementedWhenGlobalIndexesHaveIncreased() {
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
		BorrowRewardIndexes: types2.MultiRewardIndexes{
			{
				CollateralType: "borrowdenom",
				RewardIndexes: types2.RewardIndexes{
					{
						CollateralType: "rewarddenom",
						RewardFactor:   d("1000.001").MulInt64(1e12),
					},
				},
			},
		},
	}
	suite.storeJoltClaim(claim)

	suite.storeGlobalBorrowIndexes(types2.MultiRewardIndexes{
		{
			CollateralType: "borrowdenom",
			RewardIndexes: types2.RewardIndexes{
				{
					CollateralType: "rewarddenom",
					RewardFactor:   d("2000.002").MulInt64(1e12),
				},
			},
		},
	})

	borrow := NewBorrowBuilder(claim.Owner).
		WithSourceShares("borrowdenom", 1e9).
		Build()

	suite.keeper.SynchronizeJoltBorrowReward(suite.ctx, borrow)

	// new reward is (new index - old index) * borrow amount
	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Equal(
		cs(c("rewarddenom", 1_000_001_000_000)).Add(originalReward...),
		syncedClaim.Reward,
	)
}

func (suite *SynchronizeJoltBorrowRewardTests) TestRewardIsIncrementedWhenNewRewardAdded() {
	// When a new reward is added (via gov) for a jolt borrow denom the user has already borrowed, and the claim is synced
	// Then the user earns rewards for the time since the reward was added

	originalReward := arbitraryCoins()
	claim := types2.JoltLiquidityProviderClaim{
		BaseMultiClaim: types2.BaseMultiClaim{
			Owner:  arbitraryAddress(),
			Reward: originalReward,
		},
		BorrowRewardIndexes: types2.MultiRewardIndexes{
			{
				CollateralType: "rewarded",
				RewardIndexes: types2.RewardIndexes{
					{
						CollateralType: "reward",
						RewardFactor:   d("1000.001").MulInt64(1e12),
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
					RewardFactor:   d("2000.002").MulInt64(1e12),
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
					RewardFactor: d("1000.001").MulInt64(1e12),
				},
			},
		},
	}
	suite.storeGlobalBorrowIndexes(globalIndexes)

	borrow := NewBorrowBuilder(claim.Owner).
		WithSourceShares("rewarded", 1e9).
		WithSourceShares("newlyrewarded", 1e9).
		Build()

	suite.keeper.SynchronizeJoltBorrowReward(suite.ctx, borrow)

	// new reward is (new index - old index) * borrow amount for each borrowed denom
	// The old index for `newlyrewarded` isn't in the claim, so it's added starting at 0 for calculating the reward.
	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Equal(
		cs(c("otherreward", 1_000_001_000_000), c("reward", 1_000_001_000_000)).Add(originalReward...),
		syncedClaim.Reward,
	)
}

func (suite *SynchronizeJoltBorrowRewardTests) TestRewardIsIncrementedWhenNewRewardDenomAdded() {
	// When a new reward coin is added (via gov) to an already rewarded borrow denom (that the user has already borrowed), and the claim is synced;
	// Then the user earns rewards for the time since the reward was added

	originalReward := arbitraryCoins()
	claim := types2.JoltLiquidityProviderClaim{
		BaseMultiClaim: types2.BaseMultiClaim{
			Owner:  arbitraryAddress(),
			Reward: originalReward,
		},
		BorrowRewardIndexes: types2.MultiRewardIndexes{
			{
				CollateralType: "borrowed",
				RewardIndexes: types2.RewardIndexes{
					{
						CollateralType: "reward",
						RewardFactor:   d("1000.001").MulInt64(1e12),
					},
				},
			},
		},
	}
	suite.storeJoltClaim(claim)

	globalIndexes := types2.MultiRewardIndexes{
		{
			CollateralType: "borrowed",
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
	suite.storeGlobalBorrowIndexes(globalIndexes)

	borrow := NewBorrowBuilder(claim.Owner).
		WithSourceShares("borrowed", 1e9).
		Build()

	suite.keeper.SynchronizeJoltBorrowReward(suite.ctx, borrow)

	// new reward is (new index - old index) * borrow amount for each borrowed denom
	// The old index for `otherreward` isn't in the claim, so it's added starting at 0 for calculating the reward.
	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Equal(
		cs(c("reward", 1_000_001_000_000), c("otherreward", 1_000_001_000_000)).Add(originalReward...),
		syncedClaim.Reward,
	)
}

// BorrowBuilder is a tool for creating a jolt borrows.
// The builder inherits from jolt.Borrow, so fields can be accessed directly if a helper method doesn't exist.
type BorrowBuilder struct {
	jolttypes.Borrow
}

// NewBorrowBuilder creates a BorrowBuilder containing an empty borrow.
func NewBorrowBuilder(borrower sdk.AccAddress) BorrowBuilder {
	return BorrowBuilder{
		Borrow: jolttypes.Borrow{
			Borrower: borrower,
		},
	}
}

// Build assembles and returns the final borrow.
func (builder BorrowBuilder) Build() jolttypes.Borrow { return builder.Borrow }

// WithSourceShares adds a borrow amount and factor such that the source shares for this borrow is equal to specified.
// With a factor of 1, the borrow amount is the source shares. This picks an arbitrary factor to ensure factors are accounted for in production code.
func (builder BorrowBuilder) WithSourceShares(denom string, shares int64) BorrowBuilder {
	if !builder.Amount.AmountOf(denom).Equal(sdkmath.ZeroInt()) {
		panic("adding to amount with existing denom not implemented")
	}
	if _, f := builder.Index.GetInterestFactor(denom); f {
		panic("adding to indexes with existing denom not implemented")
	}

	// pick arbitrary factor
	factor := sdkmath.LegacyMustNewDecFromStr("2")

	// Calculate borrow amount that would equal the requested source shares given the above factor.
	amt := sdkmath.NewInt(shares).Mul(factor.RoundInt())

	builder.Amount = builder.Amount.Add(sdk.NewCoin(denom, amt))
	builder.Index = builder.Index.SetInterestFactor(denom, factor)
	return builder
}

// WithArbitrarySourceShares adds arbitrary borrow amounts and indexes for each specified denom.
func (builder BorrowBuilder) WithArbitrarySourceShares(denoms ...string) BorrowBuilder {
	const arbitraryShares = 1e9
	var builderHandler BorrowBuilder
	builderHandler = builder
	for _, denom := range denoms {
		builderHandler = builderHandler.WithSourceShares(denom, arbitraryShares)
	}
	return builderHandler
}

func TestCalculateRewards(t *testing.T) {
	type expected struct {
		err   error
		coins sdk.Coins
	}
	type args struct {
		oldIndexes, newIndexes types2.RewardIndexes
		sourceAmount           sdkmath.LegacyDec
	}
	testcases := []struct {
		name     string
		args     args
		expected expected
	}{
		{
			name: "when old and new indexes have same denoms, rewards are calculated correctly",
			args: args{
				oldIndexes: types2.RewardIndexes{
					{
						CollateralType: "jolt",
						RewardFactor:   d("0.000000001").MulInt64(1e12),
					},
					{
						CollateralType: "ujolt",
						RewardFactor:   d("0.1").MulInt64(1e12),
					},
				},
				newIndexes: types2.RewardIndexes{
					{
						CollateralType: "jolt",
						RewardFactor:   d("1000.0").MulInt64(1e12),
					},
					{
						CollateralType: "ujolt",
						RewardFactor:   d("0.100000001").MulInt64(1e12),
					},
				},
				sourceAmount: d("1000000000"),
			},
			expected: expected{
				// for each denom: (new - old) * sourceAmount
				coins: cs(c("jolt", 999999999999), c("ujolt", 1)),
			},
		},
		{
			name: "when new indexes have an extra denom, rewards are calculated as if it was 0 in old indexes",
			args: args{
				oldIndexes: types2.RewardIndexes{
					{
						CollateralType: "jolt",
						RewardFactor:   d("0.000000001").MulInt64(1e12),
					},
				},
				newIndexes: types2.RewardIndexes{
					{
						CollateralType: "jolt",
						RewardFactor:   d("1000.0").MulInt64(1e12),
					},
					{
						CollateralType: "ujolt",
						RewardFactor:   d("0.100000001").MulInt64(1e12),
					},
				},
				sourceAmount: d("1000000000"),
			},
			expected: expected{
				// for each denom: (new - old) * sourceAmount
				coins: cs(c("jolt", 999999999999), c("ujolt", 100000001)),
			},
		},
		{
			name: "when new indexes are smaller than old, an error is returned",
			args: args{
				oldIndexes: types2.RewardIndexes{
					{
						CollateralType: "jolt",
						RewardFactor:   d("0.2").MulInt64(1e12),
					},
				},
				newIndexes: types2.RewardIndexes{
					{
						CollateralType: "jolt",
						RewardFactor:   d("0.1").MulInt64(1e12),
					},
				},
				sourceAmount: d("1000000000"),
			},
			expected: expected{
				err: types2.ErrDecreasingRewardFactor,
			},
		},
		{
			name: "when old indexes have an extra denom, an error is returned",
			args: args{
				oldIndexes: types2.RewardIndexes{
					{
						CollateralType: "jolt",
						RewardFactor:   d("0.1").MulInt64(1e12),
					},
					{
						CollateralType: "ujolt",
						RewardFactor:   d("0.1").MulInt64(1e12),
					},
				},
				newIndexes: types2.RewardIndexes{
					{
						CollateralType: "jolt",
						RewardFactor:   d("0.2"),
					},
				},
				sourceAmount: d("1000000000"),
			},
			expected: expected{
				err: types2.ErrDecreasingRewardFactor,
			},
		},
		{
			name: "when old and new indexes are 0, rewards are 0",
			args: args{
				oldIndexes: types2.RewardIndexes{
					{
						CollateralType: "jolt",
						RewardFactor:   d("0.0"),
					},
				},
				newIndexes: types2.RewardIndexes{
					{
						CollateralType: "jolt",
						RewardFactor:   d("0.0"),
					},
				},
				sourceAmount: d("1000000000"),
			},
			expected: expected{
				coins: sdk.Coins{},
			},
		},
		{
			name: "when old and new indexes are empty, rewards are 0",
			args: args{
				oldIndexes:   types2.RewardIndexes{},
				newIndexes:   nil,
				sourceAmount: d("1000000000"),
			},
			expected: expected{
				coins: nil,
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			coins, err := keeper.Keeper{}.CalculateRewards(tc.args.oldIndexes, tc.args.newIndexes, tc.args.sourceAmount)
			if tc.expected.err != nil {
				require.True(t, errors.Is(err, tc.expected.err))
			} else {
				require.Equal(t, tc.expected.coins, coins)
			}
		})
	}
}

func TestCalculateSingleReward(t *testing.T) {
	type expected struct {
		err    error
		reward sdkmath.Int
	}
	type args struct {
		oldIndex, newIndex sdkmath.LegacyDec
		sourceAmount       sdkmath.LegacyDec
	}
	testcases := []struct {
		name     string
		args     args
		expected expected
	}{
		{
			name: "when new index is > old, rewards are calculated correctly",
			args: args{
				oldIndex:     d("0.000000001"),
				newIndex:     d("1000.0"),
				sourceAmount: d("1000000000"),
			},
			expected: expected{
				// (new - old) * sourceAmount
				reward: i(999999999999),
			},
		},
		{
			name: "when new index is < old, an error is returned",
			args: args{
				oldIndex:     d("0.000000001"),
				newIndex:     d("0.0"),
				sourceAmount: d("1000000000"),
			},
			expected: expected{
				err: types2.ErrDecreasingRewardFactor,
			},
		},
		{
			name: "when old and new indexes are 0, rewards are 0",
			args: args{
				oldIndex:     d("0.0"),
				newIndex:     d("0.0"),
				sourceAmount: d("1000000000"),
			},
			expected: expected{
				reward: sdkmath.ZeroInt(),
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			reward, err := keeper.Keeper{}.CalculateSingleReward(tc.args.oldIndex, tc.args.newIndex, tc.args.sourceAmount)
			if tc.expected.err != nil {
				require.True(t, errors.Is(err, tc.expected.err))
			} else {
				require.Equal(t, tc.expected.reward, reward)
			}
		})
	}
}
