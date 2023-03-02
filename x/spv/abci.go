package spv

import (
	"fmt"
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
		fmt.Printf(">>>>>>%v\n", dueTime.String())
		fmt.Printf(">>>>>>>>>>delta: %v\n", dueTime.Sub(currentTime).Seconds())
		if dueTime.Before(currentTime) {
			k.HandleInterest(ctx, &poolInfo)
			k.HandlePrincipal(ctx, &poolInfo)
		}
		return false
	})
}
