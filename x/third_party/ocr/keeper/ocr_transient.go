package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/x/third_party/ocr/types"
)

func (k *Keeper) SetTransientLatestEpochAndRound(
	ctx sdk.Context,
	feedId string,
	epochAndRound *types.EpochAndRound,
) {
	key := types.GetLatestEpochAndRoundKey(feedId)
	bz := k.cdc.MustMarshal(epochAndRound)
	k.getTransientStore(ctx).Set(key, bz)
}

func (k *Keeper) GetTransientLatestEpochAndRound(
	ctx sdk.Context,
	feedId string,
) *types.EpochAndRound {
	bz := k.getTransientStore(ctx).Get(types.GetLatestEpochAndRoundKey(feedId))
	if bz == nil {
		return nil
	}

	var epochAndRound types.EpochAndRound
	k.cdc.MustUnmarshal(bz, &epochAndRound)
	return &epochAndRound
}
