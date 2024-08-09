package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"

	"github.com/stretchr/testify/suite"
)

// InitializeJoltBorrowRewardTests runs unit tests for the keeper.InitializeJoltBorrowReward method
type InitializeJoltBorrowRewardTests struct {
	unitTester
}

func TestInitializeHardBorrowReward(t *testing.T) {
	suite.Run(t, new(InitializeJoltBorrowRewardTests))
}

func (suite *InitializeJoltBorrowRewardTests) TestClaimIndexesAreSetWhenClaimExists() {
	claim := types2.JoltLiquidityProviderClaim{
		BaseMultiClaim: types2.BaseMultiClaim{
			Owner: arbitraryAddress(),
		},
		// Indexes should always be empty when initialize is called.
		// If initialize is called then the user must have repaid their borrow positions,
		// which means UpdateJoltBorrowIndexDenoms was called and should have remove indexes.
		BorrowRewardIndexes: types2.MultiRewardIndexes{},
	}
	suite.storeJoltClaim(claim)

	globalIndexes := nonEmptyMultiRewardIndexes
	suite.storeGlobalBorrowIndexes(globalIndexes)

	borrow := NewBorrowBuilder(claim.Owner).
		WithArbitrarySourceShares(extractCollateralTypes(globalIndexes)...).
		Build()

	suite.keeper.InitializeJoltBorrowReward(sdk.UnwrapSDKContext(suite.ctx), borrow)

	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(sdk.UnwrapSDKContext(suite.ctx), claim.Owner)
	suite.Equal(globalIndexes, syncedClaim.BorrowRewardIndexes)
}

func (suite *InitializeJoltBorrowRewardTests) TestClaimIndexesAreSetWhenClaimDoesNotExist() {
	globalIndexes := nonEmptyMultiRewardIndexes
	suite.storeGlobalBorrowIndexes(globalIndexes)

	owner := arbitraryAddress()
	borrow := NewBorrowBuilder(owner).
		WithArbitrarySourceShares(extractCollateralTypes(globalIndexes)...).
		Build()

	suite.keeper.InitializeJoltBorrowReward(sdk.UnwrapSDKContext(suite.ctx), borrow)

	syncedClaim, found := suite.keeper.GetJoltLiquidityProviderClaim(sdk.UnwrapSDKContext(suite.ctx), owner)
	suite.True(found)
	suite.Equal(globalIndexes, syncedClaim.BorrowRewardIndexes)
}

func (suite *InitializeJoltBorrowRewardTests) TestClaimIndexesAreSetEmptyForMissingIndexes() {
	globalIndexes := nonEmptyMultiRewardIndexes
	suite.storeGlobalBorrowIndexes(globalIndexes)

	owner := arbitraryAddress()
	// Borrow a denom that is not in the global indexes.
	// This happens when a borrow denom has no rewards associated with it.
	expectedIndexes := appendUniqueEmptyMultiRewardIndex(globalIndexes)
	borrowedDenoms := extractCollateralTypes(expectedIndexes)
	borrow := NewBorrowBuilder(owner).
		WithArbitrarySourceShares(borrowedDenoms...).
		Build()

	suite.keeper.InitializeJoltBorrowReward(sdk.UnwrapSDKContext(suite.ctx), borrow)

	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(sdk.UnwrapSDKContext(suite.ctx), owner)
	suite.Equal(expectedIndexes, syncedClaim.BorrowRewardIndexes)
}
