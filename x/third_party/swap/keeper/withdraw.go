package keeper

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/x/third_party/swap/types"
)

// Withdraw removes liquidity from an existing pool from an owners deposit, converting the provided shares for
// the returned pool liquidity.
//
// If 100% of the owners shares are removed, then the deposit is deleted.  In addition, if all the pool shares
// are removed then the pool is deleted.
//
// The number of shares must be large enough to result in at least 1 unit of the smallest reserve in the pool.
// If the share input is below the minimum required for positive liquidity to be remove from both reserves, a
// insufficient error is returned.
//
// In addition, if the withdrawn liquidity for each reserve is below the provided minimum, a slippage exceeded
// error is returned.
func (k Keeper) Withdraw(ctx context.Context, owner sdk.AccAddress, shares sdkmath.Int, minCoinA, minCoinB sdk.Coin) error {
	poolID := types.PoolID(minCoinA.Denom, minCoinB.Denom)

	shareRecord, found := k.GetDepositorShares(ctx, owner, poolID)
	if !found {
		return errorsmod.Wrapf(types.ErrDepositNotFound, "no deposit for account %s and pool %s", owner, poolID)
	}

	if shares.GT(shareRecord.SharesOwned) {
		return errorsmod.Wrapf(types.ErrInvalidShares, "withdraw of %s shares greater than %s shares owned", shares, shareRecord.SharesOwned)
	}

	poolRecord, found := k.GetPool(ctx, poolID)
	if !found {
		panic(fmt.Sprintf("pool %s not found", poolID))
	}

	pool, err := types.NewDenominatedPoolWithExistingShares(poolRecord.Reserves(), poolRecord.TotalShares)
	if err != nil {
		panic(fmt.Sprintf("invalid pool %s: %s", poolID, err))
	}

	withdrawnAmount := pool.RemoveLiquidity(shares)
	if withdrawnAmount.AmountOf(minCoinA.Denom).IsZero() || withdrawnAmount.AmountOf(minCoinB.Denom).IsZero() {
		return errorsmod.Wrap(types.ErrInsufficientLiquidity, "shares must be increased")
	}
	if withdrawnAmount.AmountOf(minCoinA.Denom).LT(minCoinA.Amount) || withdrawnAmount.AmountOf(minCoinB.Denom).LT(minCoinB.Amount) {
		return errorsmod.Wrap(types.ErrSlippageExceeded, "minimum withdraw not met")
	}

	k.updatePool(ctx, poolID, pool)
	k.BeforePoolDepositModified(ctx, poolID, owner, shareRecord.SharesOwned)
	k.updateDepositorShares(ctx, owner, poolID, shareRecord.SharesOwned.Sub(shares))

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccountName, owner, withdrawnAmount)
	if err != nil {
		panic(err)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSwapWithdraw,
			sdk.NewAttribute(types.AttributeKeyPoolID, poolID),
			sdk.NewAttribute(types.AttributeKeyOwner, owner.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, withdrawnAmount.String()),
			sdk.NewAttribute(types.AttributeKeyShares, shares.String()),
		),
	)

	return nil
}
