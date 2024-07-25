package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k Keeper) processEachReserve(ctx context.Context, c sdk.Coin) (bool, error) {
	pa := k.GetParams(ctx)
	// now we set the auction

	thresholdsTokens := pa.GetBurnThreshold()
	amt := thresholdsTokens.AmountOf(c.Denom)

	if c.Amount.GTE(amt) {
		_, err := k.auctionKeeper.StartSurplusAuction(ctx, types.ModuleAccount, c, "ujolt")
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

// RunSurplusAuctions nets the surplus and debt balances and then creates surplus or debt auctions if the remaining balance is above the auction threshold parameter
func (k Keeper) RunSurplusAuctions(ctx context.Context) {
	k.IterSPVReserve(ctx, func(totalReserve sdk.Coin) (stop bool) {
		if totalReserve.IsZero() {
			return false
		}

		acc := k.accKeeper.GetModuleAccount(ctx, types.ModuleAccount)
		currentBalance := k.bankKeeper.GetBalance(ctx, acc.GetAddress(), totalReserve.Denom)
		if currentBalance.IsLT(totalReserve) {
			err := errorsmod.Wrapf(errorsmod.ErrInsufficientFunds, "we need to burn %v and we only have %v in account", totalReserve.String(), currentBalance.String())
			ctx.Logger().Error("run surplus auction error", "error msg", err.Error())
			return false
		}

		processed, err := k.processEachReserve(ctx, totalReserve)
		if err != nil {
			ctx.Logger().Error("spv", "surplusAuction", err)
			return false
		}
		if processed {
			k.SetReserve(ctx, sdk.NewCoin(totalReserve.Denom, sdkmath.ZeroInt()))
		}
		return false
	})
}
