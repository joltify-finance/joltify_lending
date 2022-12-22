package auction

import (
	"fmt"

	"github.com/joltify-finance/joltify_lending/x/third_party/auction/keeper"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/auction/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the store state from a genesis state.
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, bankKeeper types2.BankKeeper, accountKeeper types2.AccountKeeper, gs *types2.GenesisState) {
	if err := gs.Validate(); err != nil {
		panic(fmt.Sprintf("failed to validate %s genesis state: %s", types2.ModuleName, err))
	}

	keeper.SetNextAuctionID(ctx, gs.NextAuctionId)

	keeper.SetParams(ctx, gs.Params)

	totalAuctionCoins := sdk.NewCoins()

	auctions, err := types2.UnpackGenesisAuctions(gs.Auctions)
	if err != nil {
		panic(fmt.Sprintf("failed to unpack genesis auctions: %s", err))
	}
	for _, a := range auctions {
		keeper.SetAuction(ctx, a)
		// find the total coins that should be present in the module account
		totalAuctionCoins = totalAuctionCoins.Add(a.GetModuleAccountCoins()...)
	}

	// check if the module account exists
	moduleAcc := accountKeeper.GetModuleAccount(ctx, types2.ModuleName)
	if moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", types2.ModuleName))
	}

	maccCoins := bankKeeper.GetAllBalances(ctx, moduleAcc.GetAddress())

	// check module coins match auction coins
	// Note: Other sdk modules do not check this, instead just using the existing module account coins, or if zero, setting them.
	if !maccCoins.IsEqual(totalAuctionCoins) {
		panic(fmt.Sprintf("total auction coins (%s) do not equal (%s) module account (%s) ", maccCoins, types2.ModuleName, totalAuctionCoins))
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types2.GenesisState {
	nextAuctionID, err := keeper.GetNextAuctionID(ctx)
	if err != nil {
		panic(err)
	}

	params := keeper.GetParams(ctx)

	genAuctions := []types2.GenesisAuction{} // return empty list instead of nil if no auctions
	keeper.IterateAuctions(ctx, func(a types2.Auction) bool {
		ga, ok := a.(types2.GenesisAuction)
		if !ok {
			panic("could not convert stored auction to GenesisAuction type")
		}
		genAuctions = append(genAuctions, ga)
		return false
	})

	gs, err := types2.NewGenesisState(nextAuctionID, params, genAuctions)
	if err != nil {
		panic(err)
	}

	return gs
}
