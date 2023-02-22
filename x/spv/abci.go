package spv

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"time"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) {

	currentTime := ctx.BlockTime()
	// we firstly handle the interest

	k.IteratePool(ctx, func(poolInfo types.PoolInfo) (stop bool) {
		dueTime := poolInfo.LastPaymentTime.Add(time.Second * time.Duration(poolInfo.PayFreq))
		var delta time.Duration
		if currentTime.After(dueTime) {
			delta = currentTime.Sub(dueTime)
		} else {
			delta = dueTime.Sub(currentTime)
		}

		if delta < time.Minute {
			k.HandleInterest(ctx, &poolInfo)
		
		}
		return false
	})

	return
}
