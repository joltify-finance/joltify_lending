package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
)

// GetQuotaData returns a createPool from its index
func (k Keeper) GetDebug(ctx sdk.Context) (val types.ConsensusMap, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DebugKey))

	b := store.Get(types.KeyPrefix("info"))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SaveDebug(ctx sdk.Context, val types.ConsensusMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DebugKey))
	b := k.cdc.MustMarshal(&val)
	store.Set(types.KeyPrefix("info"), b)
}

func (k Keeper) DebugConsensus(ctx sdk.Context) {

	blockTime := ctx.BlockTime().String()
	target := crypto.Keccak256Hash([]byte(blockTime))

	val, found := k.GetDebug(ctx)
	if !found {
		v := make(map[string]int64)
		v[target.Hex()] = ctx.BlockHeight()
		w := types.ConsensusMap{Items: v}
		fmt.Printf(">>>>>we save %v>>>%v\n", target.Hex(), ctx.BlockHeight())
		k.SaveDebug(ctx, w)
		return
	}

	val.Items[target.Hex()] = ctx.BlockHeight()
	fmt.Printf(">>>>>we save %v>>>%v\n", target.Hex(), ctx.BlockHeight())
	k.SaveDebug(ctx, val)
}
