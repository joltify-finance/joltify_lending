package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"
)

// Withdraw returns some or all of a deposit back to original depositor
func (k Keeper) Withdraw(rctx context.Context, depositor sdk.AccAddress, coins sdk.Coins) error {
	ctx := sdk.UnwrapSDKContext(rctx)
	// Call incentive hooks
	existingDeposit, found := k.GetDeposit(ctx, depositor)
	if !found {
		return errorsmod.Wrapf(types2.ErrDepositNotFound, "no deposit found for %s", depositor)
	}
	k.BeforeDepositModified(ctx, existingDeposit)

	existingBorrow, hasExistingBorrow := k.GetBorrow(ctx, depositor)
	if hasExistingBorrow {
		k.BeforeBorrowModified(ctx, existingBorrow)
	}

	// Sync interest
	k.SyncBorrowInterest(ctx, depositor)
	k.SyncSupplyInterest(ctx, depositor)

	// Refresh Deposit after syncing interest
	deposit, _ := k.GetDeposit(ctx, depositor)

	amount, err := k.CalculateWithdrawAmount(deposit.Amount, coins)
	if err != nil {
		return err
	}

	borrow, found := k.GetBorrow(ctx, depositor)
	if !found {
		borrow = types2.Borrow{}
	}

	proposedDeposit := types2.NewDeposit(deposit.Depositor, deposit.Amount.Sub(amount...), types2.SupplyInterestFactors{})
	valid, _, err := k.IsWithinValidLtvRange(ctx, proposedDeposit, borrow)
	if err != nil {
		return err
	}
	if !valid {
		return errorsmod.Wrapf(types2.ErrInvalidWithdrawAmount, "proposed withdraw outside loan-to-value range")
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types2.ModuleAccountName, depositor, amount)
	if err != nil {
		return err
	}

	// If any coin denoms have been completely withdrawn reset the denom's supply index factor
	for _, coin := range deposit.Amount {
		if !sdk.NewCoins(coin).DenomsSubsetOf(proposedDeposit.Amount) {
			depositIndex, removed := deposit.Index.RemoveInterestFactor(coin.Denom)
			if !removed {
				return errorsmod.Wrapf(types2.ErrInvalidIndexFactorDenom, "%s", coin.Denom)
			}
			deposit.Index = depositIndex
		}
	}

	deposit.Amount = deposit.Amount.Sub(amount...)
	if deposit.Amount.Empty() {
		k.DeleteDeposit(ctx, deposit)
	} else {
		k.SetDeposit(ctx, deposit)
	}

	// Update total supplied amount
	err = k.DecrementSuppliedCoins(ctx, amount)
	if err != nil {
		return err
	}

	// Call incentive hook
	k.AfterDepositModified(ctx, deposit)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types2.EventTypeHardWithdrawal,
			sdk.NewAttribute(sdk.AttributeKeyAmount, amount.String()),
			sdk.NewAttribute(types2.AttributeKeyDepositor, depositor.String()),
		),
	)
	return nil
}

// CalculateWithdrawAmount enables full withdraw of deposited coins by adjusting withdraw amount
// to equal total deposit amount if the requested withdraw amount > current deposit amount
func (k Keeper) CalculateWithdrawAmount(available sdk.Coins, request sdk.Coins) (sdk.Coins, error) {
	result := sdk.Coins{}

	if !request.DenomsSubsetOf(available) {
		return result, types2.ErrInvalidWithdrawDenom
	}

	for _, coin := range request {
		if coin.Amount.GT(available.AmountOf(coin.Denom)) {
			result = append(result, sdk.NewCoin(coin.Denom, available.AmountOf(coin.Denom)))
		} else {
			result = append(result, coin)
		}
	}
	return result, nil
}
