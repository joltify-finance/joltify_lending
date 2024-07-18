package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/x/third_party/exchange/types"
)

// GetCampaignRewardPool fetches the trading reward pool corresponding to a given start timestamp.
func (k *Keeper) GetCampaignRewardPool(ctx sdk.Context, startTimestamp int64) *types.CampaignRewardPool {
	store := k.getStore(ctx)
	bz := store.Get(types.GetCampaignRewardPoolKey(startTimestamp))
	if bz == nil {
		return nil
	}

	var rewardPool types.CampaignRewardPool
	k.cdc.MustUnmarshal(bz, &rewardPool)
	return &rewardPool
}

// DeleteCampaignRewardPool deletes the trading reward pool corresponding to a given start timestamp.
func (k *Keeper) DeleteCampaignRewardPool(ctx sdk.Context, startTimestamp int64) {
	store := k.getStore(ctx)
	store.Delete(types.GetCampaignRewardPoolKey(startTimestamp))
}

// SetCampaignRewardPool sets the trading reward pool corresponding to a given start timestamp.
func (k *Keeper) SetCampaignRewardPool(ctx sdk.Context, rewardPool *types.CampaignRewardPool) {
	store := k.getStore(ctx)
	bz := k.cdc.MustMarshal(rewardPool)
	store.Set(types.GetCampaignRewardPoolKey(rewardPool.StartTimestamp), bz)
}

// GetAllCampaignRewardPools gets all campaign reward pools
func (k *Keeper) GetAllCampaignRewardPools(ctx sdk.Context) []*types.CampaignRewardPool {
	rewardPools := make([]*types.CampaignRewardPool, 0)

	appendPool := func(pool *types.CampaignRewardPool) (stop bool) {
		rewardPools = append(rewardPools, pool)
		return false
	}

	k.IterateCampaignRewardPools(ctx, false, appendPool)
	return rewardPools
}

// GetFirstCampaignRewardPool gets the first campaign reward pool.
func (k *Keeper) GetFirstCampaignRewardPool(ctx sdk.Context) (rewardPool *types.CampaignRewardPool) {
	appendPool := func(pool *types.CampaignRewardPool) (stop bool) {
		rewardPool = pool
		return true
	}

	k.IterateCampaignRewardPools(ctx, false, appendPool)
	return rewardPool
}

// IterateCampaignRewardPools iterates over the trading reward pools
func (k *Keeper) IterateCampaignRewardPools(
	ctx sdk.Context,
	shouldReverseIterate bool,
	process func(*types.CampaignRewardPool) (stop bool),
) {
	store := k.getStore(ctx)

	rewardPoolStore := prefix.NewStore(store, types.TradingRewardCampaignRewardPoolPrefix)

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
