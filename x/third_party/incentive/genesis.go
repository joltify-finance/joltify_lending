package incentive

import (
	"fmt"
	"strings"
	"time"

	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"

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
	accountKeeper types.AccountKeeper,
	gs types.GenesisState,
) {
	// check if the module account exists
	moduleAcc := accountKeeper.GetModuleAccount(ctx, types.IncentiveMacc)
	if moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.IncentiveMacc))
	}

	if err := gs.Validate(); err != nil {
		panic(fmt.Sprintf("failed to validate %s genesis state: %s", types.ModuleName, err))
	}

	// TODO more param validation?

	k.SetParams(ctx, gs.Params)

	// Jolt Supply / Borrow
	for _, claim := range gs.JoltLiquidityProviderClaims {
		k.SetJoltLiquidityProviderClaim(ctx, claim)
	}
	for _, gat := range gs.JoltSupplyRewardState.AccumulationTimes {
		if err := ValidateAccumulationTime(gat.PreviousAccumulationTime); err != nil {
			panic(err.Error())
		}
		k.SetPreviousJoltSupplyRewardAccrualTime(ctx, gat.CollateralType, gat.PreviousAccumulationTime)
	}
	for _, mri := range gs.JoltSupplyRewardState.MultiRewardIndexes {
		k.SetJoltSupplyRewardIndexes(ctx, mri.CollateralType, mri.RewardIndexes)
	}
	for _, gat := range gs.JoltBorrowRewardState.AccumulationTimes {
		if err := ValidateAccumulationTime(gat.PreviousAccumulationTime); err != nil {
			panic(err.Error())
		}
		k.SetPreviousJoltBorrowRewardAccrualTime(ctx, gat.CollateralType, gat.PreviousAccumulationTime)
	}
	for _, mri := range gs.JoltBorrowRewardState.MultiRewardIndexes {
		k.SetJoltBorrowRewardIndexes(ctx, mri.CollateralType, mri.RewardIndexes)
	}

	// Swap
	for _, claim := range gs.SwapClaims {
		k.SetSwapClaim(ctx, claim)
	}
	for _, gat := range gs.SwapRewardState.AccumulationTimes {
		if err := ValidateAccumulationTime(gat.PreviousAccumulationTime); err != nil {
			panic(err.Error())
		}
		k.SetSwapRewardAccrualTime(ctx, gat.CollateralType, gat.PreviousAccumulationTime)
	}
	for _, mri := range gs.SwapRewardState.MultiRewardIndexes {
		k.SetSwapRewardIndexes(ctx, mri.CollateralType, mri.RewardIndexes)
	}

	// SPV
	for _, gat := range gs.SpvRewardState.AccumulationTimes {
		if err := ValidateAccumulationTime(gat.PreviousAccumulationTime); err != nil {
			panic(err.Error())
		}
		k.SetSPVRewardAccrualTime(ctx, gat.CollateralType, gat.PreviousAccumulationTime)
	}
	for _, mri := range gs.SpvRewardState.AccRewardIndexs {
		k.SetSPVReward(ctx, mri.CollateralType, mri.AccReward)
	}

	for _, eachInvestor := range gs.SpvRewardState.SpvInvestors {
		k.SetSPVInvestorReward(ctx, eachInvestor.Pool, eachInvestor.Wallet, eachInvestor.Reward)
	}
}

func getSPVGenesisRewardState(ctx sdk.Context, keeper keeper.Keeper) types.SPVGenesisRewardState {
	var ats types.AccumulationTimes
	keeper.IterateSPVRewardAccrualTimes(ctx, func(ctype string, accTime time.Time) bool {
		ctype = strings.TrimPrefix(ctype, types.Incentiveprefix)
		ats = append(ats, types.NewAccumulationTime(ctype, accTime))
		return false
	})

	var mris []types.SPVRewardAccIndex
	keeper.IterateSPVRewardIndexes(ctx, func(ctype string, indexes types.SPVRewardAccTokens) bool {
		// we need to remove the preifx
		ctype = strings.TrimPrefix(ctype, types.Incentiveprefix)
		mris = append(mris, types.SPVRewardAccIndex{ctype, indexes})
		return false
	})

	var spvInvestors []*types.SPVGenRewardInvestorState
	keeper.IterateSPVInvestorReward(ctx, func(key string, reward types.SPVRewardAccTokens) bool {
		out := strings.TrimPrefix(key, types.Incentiveclassprefix)
		data := strings.Split(out, "-")
		poolID := data[0]
		walletAddr := data[1]
		spvInvestors = append(spvInvestors, &types.SPVGenRewardInvestorState{Wallet: walletAddr, Pool: poolID, Reward: reward.PaymentAmount})
		return false
	})

	return types.NewSPVGenesisRewardState(ats, mris, spvInvestors)
}

func getSwapGenesisRewardState(ctx sdk.Context, keeper keeper.Keeper) types.GenesisRewardState {
	var ats types.AccumulationTimes
	keeper.IterateSwapRewardAccrualTimes(ctx, func(ctype string, accTime time.Time) bool {
		ats = append(ats, types.NewAccumulationTime(ctype, accTime))
		return false
	})

	var mris types.MultiRewardIndexes
	keeper.IterateSwapRewardIndexes(ctx, func(ctype string, indexes types.RewardIndexes) bool {
		mris = append(mris, types.NewMultiRewardIndex(ctype, indexes))
		return false
	})

	return types.NewGenesisRewardState(ats, mris)
}

// ExportGenesis export genesis state for incentive module
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	params := k.GetParams(ctx)

	joltClaims := k.GetAllJoltLiquidityProviderClaims(ctx)
	joltSupplyRewardState := getJoltSupplyGenesisRewardState(ctx, k)
	joltBorrowRewardState := getJoltBorrowGenesisRewardState(ctx, k)

	swapClaims := k.GetAllSwapClaims(ctx)
	swapRewardState := getSwapGenesisRewardState(ctx, k)

	// fixme we need to fix the spv claims later
	// spvClaims := k.GetAllSPVClaims(ctx)
	spvRewardState := getSPVGenesisRewardState(ctx, k)

	return types.NewGenesisState(
		params, joltSupplyRewardState, joltBorrowRewardState, swapRewardState, spvRewardState, joltClaims, swapClaims,
	)
}

func getJoltSupplyGenesisRewardState(ctx sdk.Context, keeper keeper.Keeper) types.GenesisRewardState {
	var ats types.AccumulationTimes
	keeper.IterateJoltSupplyRewardAccrualTimes(ctx, func(ctype string, accTime time.Time) bool {
		ats = append(ats, types.NewAccumulationTime(ctype, accTime))
		return false
	})

	var mris types.MultiRewardIndexes
	keeper.IterateJoltSupplyRewardIndexes(ctx, func(ctype string, indexes types.RewardIndexes) bool {
		mris = append(mris, types.NewMultiRewardIndex(ctype, indexes))
		return false
	})

	return types.NewGenesisRewardState(ats, mris)
}

func getJoltBorrowGenesisRewardState(ctx sdk.Context, keeper keeper.Keeper) types.GenesisRewardState {
	var ats types.AccumulationTimes
	keeper.IterateJoltBorrowRewardAccrualTimes(ctx, func(ctype string, accTime time.Time) bool {
		ats = append(ats, types.NewAccumulationTime(ctype, accTime))
		return false
	})

	var mris types.MultiRewardIndexes
	keeper.IterateJoltBorrowRewardIndexes(ctx, func(ctype string, indexes types.RewardIndexes) bool {
		mris = append(mris, types.NewMultiRewardIndex(ctype, indexes))
		return false
	})

	return types.NewGenesisRewardState(ats, mris)
}

func ValidateAccumulationTime(previousAccumulationTime time.Time) error {
	//if previousAccumulationTime.Before(genesisTime.Add(-1 * EarliestValidAccumulationTime)) {
	//	return fmt.Errorf(
	//		"found accumulation time '%s' more than '%s' behind genesis time '%s'",
	//		previousAccumulationTime,
	//		EarliestValidAccumulationTime,
	//		genesisTime,
	//	)
	//}
	if previousAccumulationTime.Equal(time.Time{}) {
		return fmt.Errorf("accumulation time is not set")
	}
	return nil
}
