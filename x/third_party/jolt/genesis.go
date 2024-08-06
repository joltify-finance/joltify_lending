package jolt

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkmath "cosmossdk.io/math"
	"github.com/joltify-finance/joltify_lending/x/third_party/jolt/keeper"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"
)

// InitGenesis initializes the store state from a genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, accountKeeper types2.AccountKeeper, gs types2.GenesisState) {
	if err := gs.Validate(); err != nil {
		panic(fmt.Sprintf("failed to validate %s genesis state: %s", types2.ModuleName, err))
	}

	k.SetParams(ctx, gs.Params)

	for _, mm := range gs.Params.MoneyMarkets {
		k.SetMoneyMarket(ctx, mm.Denom, mm)
	}

	for _, gat := range gs.PreviousAccumulationTimes {
		k.SetPreviousAccrualTime(ctx, gat.CollateralType, gat.PreviousAccumulationTime)
		k.SetSupplyInterestFactor(ctx, gat.CollateralType, gat.SupplyInterestFactor)
		k.SetBorrowInterestFactor(ctx, gat.CollateralType, gat.BorrowInterestFactor)
	}

	_, found := k.GetMoneyMarket(ctx, "ujolt")
	fmt.Printf("found is jolt module %v\n", found)

	for _, deposit := range gs.Deposits {
		k.SetDeposit(ctx, deposit)
	}

	for _, borrow := range gs.Borrows {
		k.SetBorrow(ctx, borrow)
	}

	k.SetSuppliedCoins(ctx, gs.TotalSupplied)
	k.SetBorrowedCoins(ctx, gs.TotalBorrowed)
	k.SetTotalReserves(ctx, gs.TotalReserves)

	// check if the module account exists
	DepositModuleAccount := accountKeeper.GetModuleAccount(ctx, types2.ModuleAccountName)
	if DepositModuleAccount == nil {
		panic(fmt.Sprintf("%s module account has not been set", DepositModuleAccount))
	}
}

// ExportGenesis export genesis state for jolt module
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types2.GenesisState {
	params := k.GetParams(ctx)

	gats := types2.GenesisAccumulationTimes{}
	deposits := types2.Deposits{}
	borrows := types2.Borrows{}

	k.IterateDeposits(ctx, func(d types2.Deposit) bool {
		k.BeforeDepositModified(ctx, d)
		syncedDeposit, found := k.GetSyncedDeposit(ctx, d.Depositor)
		if !found {
			panic(fmt.Sprintf("syncable deposit not found for %s", d.Depositor))
		}
		deposits = append(deposits, syncedDeposit)
		return false
	})

	k.IterateBorrows(ctx, func(b types2.Borrow) bool {
		k.BeforeBorrowModified(ctx, b)
		syncedBorrow, found := k.GetSyncedBorrow(ctx, b.Borrower)
		if !found {
			panic(fmt.Sprintf("syncable borrow not found for %s", b.Borrower))
		}
		borrows = append(borrows, syncedBorrow)
		return false
	})

	totalSupplied, found := k.GetSuppliedCoins(ctx)
	if !found {
		totalSupplied = types2.DefaultTotalSupplied
	}
	totalBorrowed, found := k.GetBorrowedCoins(ctx)
	if !found {
		totalBorrowed = types2.DefaultTotalBorrowed
	}
	totalReserves, found := k.GetTotalReserves(ctx)
	if !found {
		totalReserves = types2.DefaultTotalReserves
	}

	for _, mm := range params.MoneyMarkets {
		supplyFactor, f := k.GetSupplyInterestFactor(ctx, mm.Denom)
		if !f {
			supplyFactor = sdkmath.LegacyOneDec()
		}
		borrowFactor, f := k.GetBorrowInterestFactor(ctx, mm.Denom)
		if !f {
			borrowFactor = sdkmath.LegacyOneDec()
		}
		previousAccrualTime, f := k.GetPreviousAccrualTime(ctx, mm.Denom)
		if !f {
			// Goverance adds new params at end of block, but mm's previous accrual time is set in begin blocker.
			// If a new money market is added and chain is exported before begin blocker runs, then the previous
			// accrual time will not be found. We can't set it here because our ctx doesn't contain current block
			// time; if we set it to ctx.BlockTime() then on the next block it could accrue interest from Jan 1st
			// 0001 to now. To avoid setting up a bad state, we panic.
			panic(fmt.Sprintf("expected previous accrual time to be set in state for %s", mm.Denom))
		}
		gat := types2.NewGenesisAccumulationTime(mm.Denom, previousAccrualTime, supplyFactor, borrowFactor)
		gats = append(gats, gat)

	}
	return types2.NewGenesisState(
		params, gats, deposits, borrows,
		totalSupplied, totalBorrowed, totalReserves,
	)
}
