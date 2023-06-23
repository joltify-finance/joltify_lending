package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Querier routes for the incentive module
const (
	QueryGetJoltRewards   = "jolt-rewards"
	QueryGetRewardFactors = "reward-factors"
	QueryGetParams        = "parameters"
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
	JoltSupplyRewardFactors MultiRewardIndexes `json:"jolt_supply_reward_factors" yaml:"jolt_supply_reward_factors"`
	JoltBorrowRewardFactors MultiRewardIndexes `json:"jolt_borrow_reward_factors" yaml:"jolt_borrow_reward_factors"`
}

// NewQueryGetRewardFactorsResponse returns a new instance of QueryAllRewardFactorsResponse
func NewQueryGetRewardFactorsResponse(supplyFactors,
	joltBorrowFactors MultiRewardIndexes,
) QueryGetRewardFactorsResponse {
	return QueryGetRewardFactorsResponse{
		JoltSupplyRewardFactors: supplyFactors,
		JoltBorrowRewardFactors: joltBorrowFactors,
	}
}
