package keeper_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"

	"github.com/stretchr/testify/suite"
)

// UpdateJoltBorrowIndexDenomsTests runs unit tests for the keeper.UpdateJoltBorrowIndexDenoms method
type UpdateJoltBorrowIndexDenomsTests struct {
	unitTester
}

func TestUpdateJoltBorrowIndexDenoms(t *testing.T) {
	suite.Run(t, new(UpdateJoltBorrowIndexDenomsTests))
}

func (suite *UpdateJoltBorrowIndexDenomsTests) TestClaimIndexesAreRemovedForDenomsNoLongerBorrowed() {
	claim := types.JoltLiquidityProviderClaim{
		BaseMultiClaim: types.BaseMultiClaim{
			Owner: arbitraryAddress(),
		},
		BorrowRewardIndexes: nonEmptyMultiRewardIndexes,
	}
	suite.storeJoltClaim(claim)
	suite.storeGlobalBorrowIndexes(claim.BorrowRewardIndexes)

	// remove one denom from the indexes already in the borrow
	expectedIndexes := claim.BorrowRewardIndexes[1:]
	borrow := NewBorrowBuilder(claim.Owner).
		WithArbitrarySourceShares(extractCollateralTypes(expectedIndexes)...).
		Build()

	suite.keeper.UpdateJoltBorrowIndexDenoms(suite.ctx, borrow)

	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Equal(expectedIndexes, syncedClaim.BorrowRewardIndexes)
}

func (suite *UpdateJoltBorrowIndexDenomsTests) TestClaimIndexesAreAddedForNewlyBorrowedDenoms() {
	claim := types.JoltLiquidityProviderClaim{
		BaseMultiClaim: types.BaseMultiClaim{
			Owner: arbitraryAddress(),
		},
		BorrowRewardIndexes: nonEmptyMultiRewardIndexes,
	}
	suite.storeJoltClaim(claim)
	globalIndexes := appendUniqueMultiRewardIndex(claim.BorrowRewardIndexes)
	suite.storeGlobalBorrowIndexes(globalIndexes)

	borrow := NewBorrowBuilder(claim.Owner).
		WithArbitrarySourceShares(extractCollateralTypes(globalIndexes)...).
		Build()

	suite.keeper.UpdateJoltBorrowIndexDenoms(suite.ctx, borrow)

	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Equal(globalIndexes, syncedClaim.BorrowRewardIndexes)
}

func (suite *UpdateJoltBorrowIndexDenomsTests) TestClaimIndexesAreUnchangedWhenBorrowedDenomsUnchanged() {
	claim := types.JoltLiquidityProviderClaim{
		BaseMultiClaim: types.BaseMultiClaim{
			Owner: arbitraryAddress(),
		},
		BorrowRewardIndexes: nonEmptyMultiRewardIndexes,
	}
	suite.storeJoltClaim(claim)
	// Set global indexes with same denoms but different values.
	// UpdateJoltBorrowIndexDenoms should ignore the new values.
	suite.storeGlobalBorrowIndexes(increaseAllRewardFactors(claim.BorrowRewardIndexes))

	borrow := NewBorrowBuilder(claim.Owner).
		WithArbitrarySourceShares(extractCollateralTypes(claim.BorrowRewardIndexes)...).
		Build()

	suite.keeper.UpdateJoltBorrowIndexDenoms(suite.ctx, borrow)

	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Equal(claim.BorrowRewardIndexes, syncedClaim.BorrowRewardIndexes)
}

func (suite *UpdateJoltBorrowIndexDenomsTests) TestEmptyClaimIndexesAreAddedForNewlyBorrowedButNotRewardedDenoms() {
	claim := types.JoltLiquidityProviderClaim{
		BaseMultiClaim: types.BaseMultiClaim{
			Owner: arbitraryAddress(),
		},
		BorrowRewardIndexes: nonEmptyMultiRewardIndexes,
	}
	suite.storeJoltClaim(claim)
	suite.storeGlobalBorrowIndexes(claim.BorrowRewardIndexes)

	// add a denom to the borrowed amount that is not in the global or claim's indexes
	expectedIndexes := appendUniqueEmptyMultiRewardIndex(claim.BorrowRewardIndexes)
	borrowedDenoms := extractCollateralTypes(expectedIndexes)
	borrow := NewBorrowBuilder(claim.Owner).
		WithArbitrarySourceShares(borrowedDenoms...).
		Build()

	suite.keeper.UpdateJoltBorrowIndexDenoms(suite.ctx, borrow)

	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Equal(expectedIndexes, syncedClaim.BorrowRewardIndexes)
}
