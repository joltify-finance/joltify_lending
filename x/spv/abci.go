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
	c := true
	k.IteratePool(ctx, func(poolInfo types.PoolInfo) (stop bool) {
		for c {
			dueTime := poolInfo.LastPaymentTime.Add(time.Second * time.Duration(poolInfo.PayFreq))
			poolReady := poolInfo.PoolStatus == types.PoolInfo_ACTIVE || poolInfo.PoolStatus == types.PoolInfo_FREEZING || poolInfo.PoolStatus == types.PoolInfo_PooLPayPartially
			if dueTime.Before(currentTime) && poolReady {
				err := k.HandleInterest(ctx, &poolInfo)
				if err != nil {
					if err.Error() == "pay interest too early" {
						c = false
					}
					panic(err)
				}
				ctx.Logger().Info("process interest", "pool Index:", poolInfo.Index, "latest payment", poolInfo.LastPaymentTime.Local().String())
			} else {
				c = false
			}
		}

		if poolInfo.PoolStatus == types.PoolInfo_ACTIVE || poolInfo.PoolStatus == types.PoolInfo_PooLPayPartially {
			processed := k.HandleTransfer(ctx, &poolInfo)
			if processed {
				k.SetPool(ctx, poolInfo)
				return false
			}
			ctx.Logger().Info("pool due time update", "index", poolInfo.Index, "due time", poolInfo.ProjectDueTime.Local().String())
			if poolInfo.ProjectDueTime.Before(currentTime) {
				// we pay the partial of the interest
				if poolInfo.PoolStatus == types.PoolInfo_PooLPayPartially {
					k.HandlePartialPrincipalPayment(ctx, &poolInfo, poolInfo.GetWithdrawAccounts())
				}
				// we update the project due time to the next cycle
				poolInfo.ProjectDueTime = poolInfo.ProjectDueTime.Add(time.Second * time.Duration(poolInfo.ProjectLength))
			}
			k.SetPool(ctx, poolInfo)
			return false
		}
		if poolInfo.PoolStatus == types.PoolInfo_FREEZING {
			k.HandlePrincipalPayment(ctx, &poolInfo)
		}
		k.SetPool(ctx, poolInfo)

		return false
	})
}
