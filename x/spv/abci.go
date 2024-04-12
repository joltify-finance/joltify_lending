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
	err := k.RunSurplusAuctions(ctx)
	if err != nil {
		ctx.Logger().Error("run surplus auction error", "error msg", err.Error())
	}
	k.IteratePool(ctx, func(poolInfo types.PoolInfo) (stop bool) {
		// it means we need to catchup, we give extra 30 seconds to allow the delay caused by block process time
		// currently, the block process time is 5 seconds
		for int32(currentTime.Sub(poolInfo.LastPaymentTime).Seconds()) > poolInfo.PayFreq*2 {
			if poolInfo.PoolStatus != types.PoolInfo_ACTIVE {
				break
			}

			// the pool has been stop too long
			if currentTime.Sub(poolInfo.LastPaymentTime).Hours() > 28*24*time.Hour.Hours() {
				break
			}

			err := k.HandleInterest(ctx, &poolInfo)
			if err != nil {
				if err.Error() == "pay interest too early" {
					break
				}
				if err.Error() == "no interest to be paid" {
					break
				}
				panic(err)
			}
			ctx.Logger().Info("####process interest", "pool Index:", poolInfo.Index, "latest payment", poolInfo.LastPaymentTime.Local().String())
		}
		dueTime := poolInfo.LastPaymentTime.Add(time.Second * time.Duration(poolInfo.PayFreq))
		poolReady := poolInfo.PoolStatus == types.PoolInfo_ACTIVE || poolInfo.PoolStatus == types.PoolInfo_FREEZING || poolInfo.PoolStatus == types.PoolInfo_PooLPayPartially
		if dueTime.Before(currentTime) && poolReady {
			// dueTime is the time to pay the interest for the whole cycle
			err := k.HandleInterest(ctx, &poolInfo)
			if err != nil && (err.Error() != "no interest to be paid" && err.Error() != "pay interest too early") {
				panic(err)
			}

			ctx.Logger().Info("process interest", "pool Index:", poolInfo.Index, "latest payment", poolInfo.LastPaymentTime.Local().String())

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
				processed := k.HandleTransfer(ctx, &poolInfo)
				if processed {
					ctx.Logger().Info("handler transfer", "pool status", poolInfo.PoolStatus)
				}
				k.HandlePrincipalPayment(ctx, &poolInfo)
			}
			k.SetPool(ctx, poolInfo)
		}
		return false
	})
}
