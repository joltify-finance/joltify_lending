package keeper

import (
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k Keeper) processEachReserve(ctx sdk.Context, c sdk.Coin) (bool, error) {
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
func (k Keeper) RunSurplusAuctions(ctx sdk.Context) error {
	supportedtokens := strings.Split(types.SupportedToken, ",")

	for _, eachSupported := range supportedtokens {

		totalReserve, found := k.GetReserve(ctx, eachSupported)
		if !found {
			return nil
		}

		if totalReserve.IsZero() {
			return nil
		}

		acc := k.accKeeper.GetModuleAccount(ctx, types.ModuleAccount)
		currentBalance := k.bankKeeper.GetBalance(ctx, acc.GetAddress(), eachSupported)
		if currentBalance.IsLT(totalReserve) {
			return errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "we need to burn %v and we only have %v in account", totalReserve.String(), currentBalance.String())
		}

		processed, err := k.processEachReserve(ctx, totalReserve)
		if err != nil {
			ctx.Logger().Error("spv", "surplusAuction", err)
			return err
		}
		if processed {
			k.SetReserve(ctx, sdk.NewCoin(eachSupported, sdk.ZeroInt()))
		}
	}
	return nil
}
