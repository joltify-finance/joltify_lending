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
			if poolInfo.PoolStatus == types.PoolInfo_ACTIVE || poolInfo.PoolStatus == types.PoolInfo_CLOSING {
				err := k.HandleInterest(ctx, &poolInfo)
				if err != nil {
					ctx.Logger().Error(err.Error())
					return false
				}
			}
			if poolInfo.PoolStatus == types.PoolInfo_ACTIVE {
				k.HandleTransfer(ctx, &poolInfo)
				if poolInfo.ProjectDueTime.Truncate(time.Duration(poolInfo.PayFreq) * time.Second).Before(currentTime) {
					// we pay the partial of the interest
					k.HandlePartialPrincipalPayment(ctx, &poolInfo, poolInfo.GetWithdrawAccounts())
				}
			}
			if poolInfo.PoolStatus == types.PoolInfo_CLOSING {
				k.HandlePrincipalPayment(ctx, &poolInfo)
			}
			k.SetPool(ctx, poolInfo)
		}

		return false
	})
}
