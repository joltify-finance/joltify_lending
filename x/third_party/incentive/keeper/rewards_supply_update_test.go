package keeper_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"

	"github.com/stretchr/testify/suite"
)

// UpdateJoltSupplyIndexDenomsTests runs unit tests for the keeper.UpdateJoltSupplyIndexDenoms method
type UpdateJoltSupplyIndexDenomsTests struct {
	unitTester
}

func TestUpdateHardSupplyIndexDenoms(t *testing.T) {
	suite.Run(t, new(UpdateJoltSupplyIndexDenomsTests))
}

func (suite *UpdateJoltSupplyIndexDenomsTests) TestClaimIndexesAreRemovedForDenomsNoLongerSupplied() {
	claim := types.JoltLiquidityProviderClaim{
		BaseMultiClaim: types.BaseMultiClaim{
			Owner: arbitraryAddress(),
		},
		SupplyRewardIndexes: nonEmptyMultiRewardIndexes,
	}
	suite.storeJoltClaim(claim)
	suite.storeGlobalSupplyIndexes(claim.SupplyRewardIndexes)

	// remove one denom from the indexes already in the deposit
	expectedIndexes := claim.SupplyRewardIndexes[1:]
	deposit := NewJoltDepositBuilder(claim.Owner).
		WithArbitrarySourceShares(extractCollateralTypes(expectedIndexes)...).
		Build()

	suite.keeper.UpdateJoltSupplyIndexDenoms(suite.ctx, deposit)

	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Equal(expectedIndexes, syncedClaim.SupplyRewardIndexes)
}

func (suite *UpdateJoltSupplyIndexDenomsTests) TestClaimIndexesAreAddedForNewlySuppliedDenoms() {
	claim := types.JoltLiquidityProviderClaim{
		BaseMultiClaim: types.BaseMultiClaim{
			Owner: arbitraryAddress(),
		},
		SupplyRewardIndexes: nonEmptyMultiRewardIndexes,
	}
	suite.storeJoltClaim(claim)
	globalIndexes := appendUniqueMultiRewardIndex(claim.SupplyRewardIndexes)
	suite.storeGlobalSupplyIndexes(globalIndexes)

	deposit := NewJoltDepositBuilder(claim.Owner).
		WithArbitrarySourceShares(extractCollateralTypes(globalIndexes)...).
		Build()

	suite.keeper.UpdateJoltSupplyIndexDenoms(suite.ctx, deposit)

	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Equal(globalIndexes, syncedClaim.SupplyRewardIndexes)
}

func (suite *UpdateJoltSupplyIndexDenomsTests) TestClaimIndexesAreUnchangedWhenSuppliedDenomsUnchanged() {
	claim := types.JoltLiquidityProviderClaim{
		BaseMultiClaim: types.BaseMultiClaim{
			Owner: arbitraryAddress(),
		},
		SupplyRewardIndexes: nonEmptyMultiRewardIndexes,
	}
	suite.storeJoltClaim(claim)
	// Set global indexes with same denoms but different values.
	// UpdateJoltSupplyIndexDenoms should ignore the new values.
	suite.storeGlobalSupplyIndexes(increaseAllRewardFactors(claim.SupplyRewardIndexes))

	deposit := NewJoltDepositBuilder(claim.Owner).
		WithArbitrarySourceShares(extractCollateralTypes(claim.SupplyRewardIndexes)...).
		Build()

	suite.keeper.UpdateJoltSupplyIndexDenoms(suite.ctx, deposit)

	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Equal(claim.SupplyRewardIndexes, syncedClaim.SupplyRewardIndexes)
}

func (suite *UpdateJoltSupplyIndexDenomsTests) TestEmptyClaimIndexesAreAddedForNewlySuppliedButNotRewardedDenoms() {
	claim := types.JoltLiquidityProviderClaim{
		BaseMultiClaim: types.BaseMultiClaim{
			Owner: arbitraryAddress(),
		},
		SupplyRewardIndexes: nonEmptyMultiRewardIndexes,
	}
	suite.storeJoltClaim(claim)
	suite.storeGlobalSupplyIndexes(claim.SupplyRewardIndexes)

	// add a denom to the deposited amount that is not in the global or claim's indexes
	expectedIndexes := appendUniqueEmptyMultiRewardIndex(claim.SupplyRewardIndexes)
	depositedDenoms := extractCollateralTypes(expectedIndexes)
	deposit := NewJoltDepositBuilder(claim.Owner).
		WithArbitrarySourceShares(depositedDenoms...).
		Build()

	suite.keeper.UpdateJoltSupplyIndexDenoms(suite.ctx, deposit)

	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Equal(expectedIndexes, syncedClaim.SupplyRewardIndexes)
}
