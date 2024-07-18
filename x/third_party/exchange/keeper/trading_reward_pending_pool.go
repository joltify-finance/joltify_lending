package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/x/third_party/exchange/types"
)

// GetCampaignRewardPendingPool fetches the trading reward pool corresponding to a given start timestamp.
func (k *Keeper) GetCampaignRewardPendingPool(ctx sdk.Context, startTimestamp int64) *types.CampaignRewardPool {
	store := k.getStore(ctx)
	bz := store.Get(types.GetCampaignRewardPendingPoolKey(startTimestamp))
	if bz == nil {
		return nil
	}

	var rewardPool types.CampaignRewardPool
	k.cdc.MustUnmarshal(bz, &rewardPool)
	return &rewardPool
}

// DeleteCampaignRewardPendingPool deletes the trading reward pool corresponding to a given start timestamp.
func (k *Keeper) DeleteCampaignRewardPendingPool(ctx sdk.Context, startTimestamp int64) {
	store := k.getStore(ctx)
	store.Delete(types.GetCampaignRewardPendingPoolKey(startTimestamp))
}

// SetCampaignRewardPendingPool sets the trading reward pool corresponding to a given start timestamp.
func (k *Keeper) SetCampaignRewardPendingPool(ctx sdk.Context, rewardPool *types.CampaignRewardPool) {
	store := k.getStore(ctx)
	bz := k.cdc.MustMarshal(rewardPool)
	store.Set(types.GetCampaignRewardPendingPoolKey(rewardPool.StartTimestamp), bz)
}

// GetAllCampaignRewardPendingPools gets all campaign reward pools
func (k *Keeper) GetAllCampaignRewardPendingPools(ctx sdk.Context) []*types.CampaignRewardPool {
	rewardPools := make([]*types.CampaignRewardPool, 0)

	appendPool := func(pool *types.CampaignRewardPool) (stop bool) {
		rewardPools = append(rewardPools, pool)
		return false
	}

	k.IterateCampaignRewardPendingPools(ctx, false, appendPool)
	return rewardPools
}

// GetFirstCampaignRewardPendingPool gets the first campaign reward pool.
func (k *Keeper) GetFirstCampaignRewardPendingPool(ctx sdk.Context) (rewardPool *types.CampaignRewardPool) {
	appendPool := func(pool *types.CampaignRewardPool) (stop bool) {
		rewardPool = pool
		return true
	}

	k.IterateCampaignRewardPendingPools(ctx, false, appendPool)
	return rewardPool
}

// IterateCampaignRewardPendingPools iterates over the trading reward pools
func (k *Keeper) IterateCampaignRewardPendingPools(
	ctx sdk.Context,
	shouldReverseIterate bool,
	process func(*types.CampaignRewardPool) (stop bool),
) {
	store := k.getStore(ctx)

	rewardPoolStore := prefix.NewStore(store, types.TradingRewardCampaignRewardPendingPoolPrefix)

	var iterator storetypes.Iterator
	if shouldReverseIterate {
		iterator = rewardPoolStore.ReverseIterator(nil, nil)
	} else {
		iterator = rewardPoolStore.Iterator(nil, nil)
	}
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var pool types.CampaignRewardPool
		bz := iterator.Value()
		k.cdc.MustUnmarshal(bz, &pool)
		if process(&pool) {
			return
		}
	}
}
