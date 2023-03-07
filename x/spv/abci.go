package spv

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) {

	currentTime := ctx.BlockTime()
	// we firstly handle the interest

	k.IteratePool(ctx, func(poolInfo types.PoolInfo) (stop bool) {
		if poolInfo.PoolStatus != types.PoolInfo_ACTIVE {
			return false
		}

		dueTime := poolInfo.LastPaymentTime.Add(time.Second * time.Duration(poolInfo.PayFreq))
		if dueTime.Before(currentTime) {
			k.HandleInterest(ctx, &poolInfo)
			k.HandleTransfer(ctx, &poolInfo)

			if poolInfo.ProjectDueTime.Before(currentTime) {
				// we pay the partial of the interest
				k.HandlePartialPrincipalPayment(ctx, &poolInfo)
			}

			k.HandlePrincipalPayment(ctx, &poolInfo)
		}
		return false
	})
}
