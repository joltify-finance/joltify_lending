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
		dueTime := poolInfo.LastPaymentTime.Add(time.Second * time.Duration(poolInfo.PayFreq))
		if dueTime.Before(currentTime) {
			if poolInfo.PoolStatus == types.PoolInfo_ACTIVE || poolInfo.PoolStatus == types.PoolInfo_FREEZING {
				err := k.HandleInterest(ctx, &poolInfo)
				if err != nil {
					ctx.Logger().Error(err.Error())
					return false
				}
			}
			if poolInfo.PoolStatus == types.PoolInfo_ACTIVE {
				k.HandleTransfer(ctx, &poolInfo)
				if poolInfo.ProjectDueTime.Before(currentTime) {
					// we pay the partial of the interest
					k.HandlePartialPrincipalPayment(ctx, &poolInfo, poolInfo.GetWithdrawAccounts())
					// we update the project due time to the next cycle
					poolInfo.ProjectDueTime = poolInfo.ProjectDueTime.Add(time.Second * time.Duration(poolInfo.PayFreq))
					k.SetPool(ctx, poolInfo)
				}
			}
			if poolInfo.PoolStatus == types.PoolInfo_FREEZING {
				k.HandlePrincipalPayment(ctx, &poolInfo)
			}
			k.SetPool(ctx, poolInfo)
		}

		return false
	})
}
