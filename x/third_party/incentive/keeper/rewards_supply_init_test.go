package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"

	"github.com/stretchr/testify/suite"
)

// InitializeJoltSupplyRewardTests runs unit tests for the keeper.InitializeJoltSupplyReward method
type InitializeJoltSupplyRewardTests struct {
	unitTester
}

func TestInitializeJoltSupplyReward(t *testing.T) {
	suite.Run(t, new(InitializeJoltSupplyRewardTests))
}

func (suite *InitializeJoltSupplyRewardTests) TestClaimIndexesAreSetWhenClaimExists() {
	claim := types2.JoltLiquidityProviderClaim{
		BaseMultiClaim: types2.BaseMultiClaim{
			Owner: arbitraryAddress(),
		},
		// Indexes should always be empty when initialize is called.
		// If initialize is called then the user must have repaid their deposit positions,
		// which means UpdateJoltSupplyIndexDenoms was called and should have remove indexes.
		SupplyRewardIndexes: types2.MultiRewardIndexes{},
	}
	suite.storeJoltClaim(claim)

	globalIndexes := nonEmptyMultiRewardIndexes
	suite.storeGlobalSupplyIndexes(globalIndexes)

	deposit := NewJoltDepositBuilder(claim.Owner).
		WithArbitrarySourceShares(extractCollateralTypes(globalIndexes)...).
		Build()

	suite.keeper.InitializeJoltSupplyReward(sdk.UnwrapSDKContext(suite.ctx), deposit)

	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(sdk.UnwrapSDKContext(suite.ctx), claim.Owner)
	suite.Equal(globalIndexes, syncedClaim.SupplyRewardIndexes)
}

func (suite *InitializeJoltSupplyRewardTests) TestClaimIndexesAreSetWhenClaimDoesNotExist() {
	globalIndexes := nonEmptyMultiRewardIndexes
	suite.storeGlobalSupplyIndexes(globalIndexes)

	owner := arbitraryAddress()
	deposit := NewJoltDepositBuilder(owner).
		WithArbitrarySourceShares(extractCollateralTypes(globalIndexes)...).
		Build()

	suite.keeper.InitializeJoltSupplyReward(sdk.UnwrapSDKContext(suite.ctx), deposit)

	syncedClaim, found := suite.keeper.GetJoltLiquidityProviderClaim(sdk.UnwrapSDKContext(suite.ctx), owner)
	suite.True(found)
	suite.Equal(globalIndexes, syncedClaim.SupplyRewardIndexes)
}

func (suite *InitializeJoltSupplyRewardTests) TestClaimIndexesAreSetEmptyForMissingIndexes() {
	globalIndexes := nonEmptyMultiRewardIndexes
	suite.storeGlobalSupplyIndexes(globalIndexes)

	owner := arbitraryAddress()
	// Supply a denom that is not in the global indexes.
	// This happens when a deposit denom has no rewards associated with it.
	expectedIndexes := appendUniqueEmptyMultiRewardIndex(globalIndexes)
	depositedDenoms := extractCollateralTypes(expectedIndexes)
	deposit := NewJoltDepositBuilder(owner).
		WithArbitrarySourceShares(depositedDenoms...).
		Build()

	suite.keeper.InitializeJoltSupplyReward(sdk.UnwrapSDKContext(suite.ctx), deposit)

	syncedClaim, _ := suite.keeper.GetJoltLiquidityProviderClaim(sdk.UnwrapSDKContext(suite.ctx), owner)
	suite.Equal(expectedIndexes, syncedClaim.SupplyRewardIndexes)
}
