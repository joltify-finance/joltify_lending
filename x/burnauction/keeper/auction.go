package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/burnauction/types"
)

// RunSurplusAuctions nets the surplus and debt balances and then creates surplus or debt auctions if the remaining balance is above the auction threshold parameter
func (k Keeper) RunSurplusAuctions(ctx context.Context) {
	acc := k.accKeeper.GetModuleAccount(ctx, types.ModuleAccount)
	currentBalances := k.bankKeeper.GetAllBalances(ctx, acc.GetAddress())

	pa := k.GetParams(ctx)
	burnThreshold := pa.GetBurnThreshold()

	burnTokens := sdk.NewCoins()
	for _, eachThreshold := range burnThreshold {
		if currentBalances.AmountOf(eachThreshold.Denom).GTE(eachThreshold.Amount) {
			ok, c := currentBalances.Find(eachThreshold.Denom)
			if !ok {
				panic("should never fail to find the token")
			}
			burnTokens = burnTokens.Add(c)
		}
	}

	if burnTokens.Empty() {
		return
	}

	for _, el := range burnTokens {
		_, err := k.auctionKeeper.StartSurplusAuction(ctx, types.ModuleAccount, el, "ujolt")
		if err != nil {
			sdk.UnwrapSDKContext(ctx).Logger().Error("failed to start surplus auction in auction module", "error", err)
			return
		}
	}

	sdk.UnwrapSDKContext(ctx).EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(types.EventModuleName, types.ModuleName),
			sdk.NewAttribute("burn_amount", burnTokens.String()),
		),
	)
}
