package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k Keeper) HandleInterest(ctx sdk.Context, poolInfo *types.PoolInfo) {
	totalAmountDue := k.getAllInterestToBePaid(ctx, poolInfo)

	if poolInfo.EscrowInterestAmount.Amount.LT(totalAmountDue) {
		ctx.Logger().Error("insufficient fund to pay the interest %v<%v", poolInfo.EscrowInterestAmount.String(), totalAmountDue.String())
		return
	}

	// finally, we update the poolinfo
	poolInfo.LastPaymentTime = ctx.BlockTime()
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

}
