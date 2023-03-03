package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k Keeper) HandleInterest(ctx sdk.Context, poolInfo *types.PoolInfo) error {
	totalAmountDue, err := k.getAllInterestToBePaid(ctx, poolInfo)
	if err != nil {
		ctx.Logger().Info(err.Error())
		return err
	}

	if poolInfo.EscrowInterestAmount.Amount.LT(totalAmountDue) {
		ctx.Logger().Error("insufficient fund to pay the interest %v<%v", poolInfo.EscrowInterestAmount.String(), totalAmountDue.String())
		return types.ErrInsufficientFund
	}

	// finally, we update the poolinfo
	currentTimeTruncated := ctx.BlockTime().Truncate(time.Duration(poolInfo.PayFreq) * time.Second)
	poolInfo.LastPaymentTime = currentTimeTruncated
	if poolInfo.BorrowedAmount.Equal(sdk.ZeroInt()) {
		poolInfo.PoolStatus = types.PoolInfo_CLOSED
	}

	poolInfo.EscrowInterestAmount = poolInfo.EscrowInterestAmount.SubAmount(totalAmountDue)

	k.SetPool(ctx, *poolInfo)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRepayInterest,
			sdk.NewAttribute("amount", totalAmountDue.String()),
		),
	)
	return nil
}

func (k Keeper) HandlePrincipal(ctx sdk.Context, poolInfo *types.PoolInfo) {
	token := poolInfo.EscrowPrincipalAmount
	if token.Amount.Equal(sdk.ZeroInt()) {
		return
	}
	aboutToPay := sdk.NewCoin(token.Denom, token.Amount)
	if token.IsGTE(poolInfo.BorrowedAmount) {
		aboutToPay = poolInfo.BorrowedAmount
	}
	k.travelThoughPrincipalToBePaid(ctx, poolInfo, aboutToPay)
	// now we query all the borrows

	poolInfo.BorrowedAmount = poolInfo.BorrowedAmount.Sub(aboutToPay)
	poolInfo.BorrowableAmount = poolInfo.BorrowableAmount.Add(aboutToPay)

	// once the pool borrowed is 0, we will deactive the pool
	if poolInfo.BorrowedAmount.Amount.Equal(sdk.ZeroInt()) {
		poolInfo.PoolStatus = types.PoolInfo_CLOSING
	}
	poolInfo.EscrowPrincipalAmount = poolInfo.EscrowPrincipalAmount.Sub(aboutToPay)
	k.SetPool(ctx, *poolInfo)

}
