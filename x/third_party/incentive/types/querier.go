package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Querier routes for the incentive module
const (
	QueryGetJoltRewards        = "jolt-rewards"
	QueryGetUSDXMintingRewards = "usdx-minting-rewards"
	QueryGetDelegatorRewards   = "delegator-rewards"
	QueryGetSwapRewards        = "swap-rewards"
	QueryGetSavingsRewards     = "savings-rewards"
	QueryGetRewardFactors      = "reward-factors"
	QueryGetParams             = "parameters"

	RestClaimCollateralType = "collateral_type"
	RestClaimOwner          = "owner"
	RestClaimType           = "type"
	RestUnsynced            = "unsynced"
)

// QueryRewardsParams params for query /incentive/rewards/<claim type>
type QueryRewardsParams struct {
	Page           int            `json:"page" yaml:"page"`
	Limit          int            `json:"limit" yaml:"limit"`
	Owner          sdk.AccAddress `json:"owner" yaml:"owner"`
	Unsynchronized bool           `json:"unsynchronized" yaml:"unsynchronized"`
}

// NewQueryRewardsParams returns QueryRewardsParams
func NewQueryRewardsParams(page, limit int, owner sdk.AccAddress, unsynchronized bool) QueryRewardsParams {
	return QueryRewardsParams{
		Page:           page,
		Limit:          limit,
		Owner:          owner,
		Unsynchronized: unsynchronized,
	}
}

// QueryGetRewardFactorsResponse holds the response to a reward factor query
type QueryGetRewardFactorsResponse struct {
	USDXMintingRewardFactors RewardIndexes      `json:"usdx_minting_reward_factors" yaml:"usdx_minting_reward_factors"`
	JoltSupplyRewardFactors  MultiRewardIndexes `json:"jolt_supply_reward_factors" yaml:"jolt_supply_reward_factors"`
	JoltBorrowRewardFactors  MultiRewardIndexes `json:"jolt_borrow_reward_factors" yaml:"jolt_borrow_reward_factors"`
	DelegatorRewardFactors   MultiRewardIndexes `json:"delegator_reward_factors" yaml:"delegator_reward_factors"`
	SwapRewardFactors        MultiRewardIndexes `json:"swap_reward_factors" yaml:"swap_reward_factors"`
	SavingsRewardFactors     MultiRewardIndexes `json:"savings_reward_factors" yaml:"savings_reward_factors"`
}

// NewQueryGetRewardFactorsResponse returns a new instance of QueryAllRewardFactorsResponse
func NewQueryGetRewardFactorsResponse(usdxMintingFactors RewardIndexes, supplyFactors,
	joltBorrowFactors, delegatorFactors, swapFactors, savingsFactors MultiRewardIndexes,
) QueryGetRewardFactorsResponse {
	return QueryGetRewardFactorsResponse{
		USDXMintingRewardFactors: usdxMintingFactors,
		JoltSupplyRewardFactors:  supplyFactors,
		JoltBorrowRewardFactors:  joltBorrowFactors,
		DelegatorRewardFactors:   delegatorFactors,
		SwapRewardFactors:        swapFactors,
		SavingsRewardFactors:     savingsFactors,
	}
}
