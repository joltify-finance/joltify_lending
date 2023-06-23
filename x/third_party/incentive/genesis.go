package incentive

import (
	"fmt"
	"time"

	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const year = 365 * 24 * time.Hour

// EarliestValidAccumulationTime is how far behind the genesis time an accumulation time can be for it to be valid.
// It's a safety check to ensure rewards aren't accidentally accumulated for many years on the first block (eg since Jan 1970).
var EarliestValidAccumulationTime = year

// InitGenesis initializes the store state from a genesis state.
func InitGenesis(
	ctx sdk.Context,
	k keeper.Keeper,
	accountKeeper types2.AccountKeeper,
	gs types2.GenesisState,
) {
	// check if the module account exists
	moduleAcc := accountKeeper.GetModuleAccount(ctx, types2.IncentiveMacc)
	if moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", types2.IncentiveMacc))
	}

	if err := gs.Validate(); err != nil {
		panic(fmt.Sprintf("failed to validate %s genesis state: %s", types2.ModuleName, err))
	}

	// TODO more param validation?

	k.SetParams(ctx, gs.Params)

	// Jolt Supply / Borrow
	for _, claim := range gs.JoltLiquidityProviderClaims {
		k.SetJoltLiquidityProviderClaim(ctx, claim)
	}
	for _, gat := range gs.JoltSupplyRewardState.AccumulationTimes {
		if err := ValidateAccumulationTime(gat.PreviousAccumulationTime, ctx.BlockTime()); err != nil {
			panic(err.Error())
		}
		k.SetPreviousJoltSupplyRewardAccrualTime(ctx, gat.CollateralType, gat.PreviousAccumulationTime)
	}
	for _, mri := range gs.JoltSupplyRewardState.MultiRewardIndexes {
		k.SetJoltSupplyRewardIndexes(ctx, mri.CollateralType, mri.RewardIndexes)
	}
	for _, gat := range gs.JoltBorrowRewardState.AccumulationTimes {
		if err := ValidateAccumulationTime(gat.PreviousAccumulationTime, ctx.BlockTime()); err != nil {
			panic(err.Error())
		}
		k.SetPreviousJoltBorrowRewardAccrualTime(ctx, gat.CollateralType, gat.PreviousAccumulationTime)
	}
	for _, mri := range gs.JoltBorrowRewardState.MultiRewardIndexes {
		k.SetJoltBorrowRewardIndexes(ctx, mri.CollateralType, mri.RewardIndexes)
	}
}

// ExportGenesis export genesis state for incentive module
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types2.GenesisState {
	params := k.GetParams(ctx)

	joltClaims := k.GetAllJoltLiquidityProviderClaims(ctx)
	joltSupplyRewardState := getJoltSupplyGenesisRewardState(ctx, k)
	joltBorrowRewardState := getJoltBorrowGenesisRewardState(ctx, k)

	return types2.NewGenesisState(
		params, joltSupplyRewardState, joltBorrowRewardState, joltClaims,
	)
}

func getJoltSupplyGenesisRewardState(ctx sdk.Context, keeper keeper.Keeper) types2.GenesisRewardState {
	var ats types2.AccumulationTimes
	keeper.IterateJoltSupplyRewardAccrualTimes(ctx, func(ctype string, accTime time.Time) bool {
		ats = append(ats, types2.NewAccumulationTime(ctype, accTime))
		return false
	})

	var mris types2.MultiRewardIndexes
	keeper.IterateJoltSupplyRewardIndexes(ctx, func(ctype string, indexes types2.RewardIndexes) bool {
		mris = append(mris, types2.NewMultiRewardIndex(ctype, indexes))
		return false
	})

	return types2.NewGenesisRewardState(ats, mris)
}

func getJoltBorrowGenesisRewardState(ctx sdk.Context, keeper keeper.Keeper) types2.GenesisRewardState {
	var ats types2.AccumulationTimes
	keeper.IterateJoltBorrowRewardAccrualTimes(ctx, func(ctype string, accTime time.Time) bool {
		ats = append(ats, types2.NewAccumulationTime(ctype, accTime))
		return false
	})

	var mris types2.MultiRewardIndexes
	keeper.IterateJoltBorrowRewardIndexes(ctx, func(ctype string, indexes types2.RewardIndexes) bool {
		mris = append(mris, types2.NewMultiRewardIndex(ctype, indexes))
		return false
	})

	return types2.NewGenesisRewardState(ats, mris)
}

func ValidateAccumulationTime(previousAccumulationTime, genesisTime time.Time) error {
	if previousAccumulationTime.Before(genesisTime.Add(-1 * EarliestValidAccumulationTime)) {
		return fmt.Errorf(
			"found accumulation time '%s' more than '%s' behind genesis time '%s'",
			previousAccumulationTime,
			EarliestValidAccumulationTime,
			genesisTime,
		)
	}
	return nil
}
