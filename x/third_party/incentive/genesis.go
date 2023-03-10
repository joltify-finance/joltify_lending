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
var EarliestValidAccumulationTime time.Duration = year

// InitGenesis initializes the store state from a genesis state.
func InitGenesis(
	ctx sdk.Context,
	k keeper.Keeper,
	accountKeeper types2.AccountKeeper,
	bankKeeper types2.BankKeeper,
	cdpKeeper types2.CdpKeeper,
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

	for _, rp := range gs.Params.USDXMintingRewardPeriods {
		if _, found := cdpKeeper.GetCollateral(ctx, rp.CollateralType); !found {
			panic(fmt.Sprintf("incentive params contain collateral not found in cdp params: %s", rp.CollateralType))
		}
	}
	// TODO more param validation?

	k.SetParams(ctx, gs.Params)

	// USDX Minting
	for _, claim := range gs.USDXMintingClaims {
		k.SetUSDXMintingClaim(ctx, claim)
	}
	for _, gat := range gs.USDXRewardState.AccumulationTimes {
		if err := ValidateAccumulationTime(gat.PreviousAccumulationTime, ctx.BlockTime()); err != nil {
			panic(err.Error())
		}
		k.SetPreviousUSDXMintingAccrualTime(ctx, gat.CollateralType, gat.PreviousAccumulationTime)
	}
	for _, mri := range gs.USDXRewardState.MultiRewardIndexes {
		factor, found := mri.RewardIndexes.Get(types2.USDXMintingRewardDenom)
		if !found || len(mri.RewardIndexes) != 1 {
			panic(fmt.Sprintf("USDX Minting reward factors must only have denom %s", types2.USDXMintingRewardDenom))
		}
		k.SetUSDXMintingRewardFactor(ctx, mri.CollateralType, factor)
	}

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

	// Delegator
	for _, claim := range gs.DelegatorClaims {
		k.SetDelegatorClaim(ctx, claim)
	}
	for _, gat := range gs.DelegatorRewardState.AccumulationTimes {
		if err := ValidateAccumulationTime(gat.PreviousAccumulationTime, ctx.BlockTime()); err != nil {
			panic(err.Error())
		}
		k.SetPreviousDelegatorRewardAccrualTime(ctx, gat.CollateralType, gat.PreviousAccumulationTime)
	}
	for _, mri := range gs.DelegatorRewardState.MultiRewardIndexes {
		k.SetDelegatorRewardIndexes(ctx, mri.CollateralType, mri.RewardIndexes)
	}

	// Swap
	for _, claim := range gs.SwapClaims {
		k.SetSwapClaim(ctx, claim)
	}
	for _, gat := range gs.SwapRewardState.AccumulationTimes {
		if err := ValidateAccumulationTime(gat.PreviousAccumulationTime, ctx.BlockTime()); err != nil {
			panic(err.Error())
		}
		k.SetSwapRewardAccrualTime(ctx, gat.CollateralType, gat.PreviousAccumulationTime)
	}
	for _, mri := range gs.SwapRewardState.MultiRewardIndexes {
		k.SetSwapRewardIndexes(ctx, mri.CollateralType, mri.RewardIndexes)
	}

	// Savings
	for _, claim := range gs.SavingsClaims {
		k.SetSavingsClaim(ctx, claim)
	}
	for _, gat := range gs.SavingsRewardState.AccumulationTimes {
		if err := ValidateAccumulationTime(gat.PreviousAccumulationTime, ctx.BlockTime()); err != nil {
			panic(err.Error())
		}
		k.SetSavingsRewardAccrualTime(ctx, gat.CollateralType, gat.PreviousAccumulationTime)
	}
	for _, mri := range gs.SavingsRewardState.MultiRewardIndexes {
		k.SetSavingsRewardIndexes(ctx, mri.CollateralType, mri.RewardIndexes)
	}
}

// ExportGenesis export genesis state for incentive module
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types2.GenesisState {
	params := k.GetParams(ctx)

	usdxClaims := k.GetAllUSDXMintingClaims(ctx)
	usdxRewardState := getUSDXMintingGenesisRewardState(ctx, k)

	joltClaims := k.GetAllJoltLiquidityProviderClaims(ctx)
	joltSupplyRewardState := getJoltSupplyGenesisRewardState(ctx, k)
	joltBorrowRewardState := getJoltBorrowGenesisRewardState(ctx, k)

	delegatorClaims := k.GetAllDelegatorClaims(ctx)
	delegatorRewardState := getDelegatorGenesisRewardState(ctx, k)

	swapClaims := k.GetAllSwapClaims(ctx)
	swapRewardState := getSwapGenesisRewardState(ctx, k)

	savingsClaims := k.GetAllSavingsClaims(ctx)
	savingsRewardState := getSavingsGenesisRewardState(ctx, k)

	return types2.NewGenesisState(
		params,
		usdxRewardState, joltSupplyRewardState, joltBorrowRewardState, delegatorRewardState, swapRewardState,
		savingsRewardState, usdxClaims, joltClaims, delegatorClaims, swapClaims, savingsClaims,
	)
}

func getUSDXMintingGenesisRewardState(ctx sdk.Context, keeper keeper.Keeper) types2.GenesisRewardState {
	var ats types2.AccumulationTimes
	keeper.IterateUSDXMintingAccrualTimes(ctx, func(ctype string, accTime time.Time) bool {
		ats = append(ats, types2.NewAccumulationTime(ctype, accTime))
		return false
	})

	var mris types2.MultiRewardIndexes
	keeper.IterateUSDXMintingRewardFactors(ctx, func(ctype string, factor sdk.Dec) bool {
		mris = append(
			mris,
			types2.NewMultiRewardIndex(
				ctype,
				types2.RewardIndexes{types2.NewRewardIndex(types2.USDXMintingRewardDenom, factor)},
			),
		)
		return false
	})

	return types2.NewGenesisRewardState(ats, mris)
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

func getDelegatorGenesisRewardState(ctx sdk.Context, keeper keeper.Keeper) types2.GenesisRewardState {
	var ats types2.AccumulationTimes
	keeper.IterateDelegatorRewardAccrualTimes(ctx, func(ctype string, accTime time.Time) bool {
		ats = append(ats, types2.NewAccumulationTime(ctype, accTime))
		return false
	})

	var mris types2.MultiRewardIndexes
	keeper.IterateDelegatorRewardIndexes(ctx, func(ctype string, indexes types2.RewardIndexes) bool {
		mris = append(mris, types2.NewMultiRewardIndex(ctype, indexes))
		return false
	})

	return types2.NewGenesisRewardState(ats, mris)
}

func getSwapGenesisRewardState(ctx sdk.Context, keeper keeper.Keeper) types2.GenesisRewardState {
	var ats types2.AccumulationTimes
	keeper.IterateSwapRewardAccrualTimes(ctx, func(ctype string, accTime time.Time) bool {
		ats = append(ats, types2.NewAccumulationTime(ctype, accTime))
		return false
	})

	var mris types2.MultiRewardIndexes
	keeper.IterateSwapRewardIndexes(ctx, func(ctype string, indexes types2.RewardIndexes) bool {
		mris = append(mris, types2.NewMultiRewardIndex(ctype, indexes))
		return false
	})

	return types2.NewGenesisRewardState(ats, mris)
}

func getSavingsGenesisRewardState(ctx sdk.Context, keeper keeper.Keeper) types2.GenesisRewardState {
	var ats types2.AccumulationTimes
	keeper.IterateSavingsRewardAccrualTimes(ctx, func(ctype string, accTime time.Time) bool {
		ats = append(ats, types2.NewAccumulationTime(ctype, accTime))
		return false
	})

	var mris types2.MultiRewardIndexes
	keeper.IterateSavingsRewardIndexes(ctx, func(ctype string, indexes types2.RewardIndexes) bool {
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
