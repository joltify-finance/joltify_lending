package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k Keeper) processEachReserve(ctx sdk.Context, c sdk.Coin) (bool, error) {

	pa := k.GetParams(ctx)
	// now we set the auction
	if c.IsGTE(pa.BurnThreshold) {
		_, err := k.auctionKeeper.StartSurplusAuction(ctx, types.ModuleAccount, c, "ujolt")
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

// RunSurplusAuctions nets the surplus and debt balances and then creates surplus or debt auctions if the remaining balance is above the auction threshold parameter
func (k Keeper) RunSurplusAuctions(ctx sdk.Context) error {
	totalReserve, found := k.GetReserve(ctx, types.SupportedToken)
	if !found {
		return nil
	}
	processed, err := k.processEachReserve(ctx, totalReserve)
	if err != nil {
		ctx.Logger().Error("spv", "surplusAuction", err)
		return err
	}
	if processed {
		k.SetReserve(ctx, sdk.NewCoin(types.SupportedToken, sdk.ZeroInt()))
	}
	return nil
}
