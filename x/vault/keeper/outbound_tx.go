package keeper

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
)

// SetOutboundTx set a specific outboundTx in the store from its index
func (k Keeper) SetOutboundTx(ctx sdk.Context, outboundTx types.OutboundTx) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OutboundTxKeyPrefix))
	b := k.cdc.MustMarshal(&outboundTx)
	store.Set(types.OutboundTxKey(
		outboundTx.Index,
	), b)
}

// GetOutboundTx returns a outboundTx from its index
func (k Keeper) GetOutboundTx(
	ctx sdk.Context,
	requestID string,
) (val types.OutboundTx, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OutboundTxKeyPrefix))

	b := store.Get(types.OutboundTxKey(
		requestID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// SetOutboundTxProposal set proposals based on its requestID:outboundTxID
func (k Keeper) SetOutboundTxProposal(ctx sdk.Context, reqID, outboundTxID string, proposals types.Proposals) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OutboundTxProposalKeyPrefix))
	b := k.cdc.MustMarshal(&proposals)
	key := fmt.Sprintf("%v:%v", reqID, outboundTxID)
	store.Set(types.OutboundTxKey(
		strings.ToLower(key),
	), b)
}

// GetOutboundTxProposal returns proposals from its requestID:outboundTxID
func (k Keeper) GetOutboundTxProposal(
	ctx sdk.Context,
	reqID, outboundTxID string,
) (val types.Proposals, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OutboundTxProposalKeyPrefix))

	key := fmt.Sprintf("%v:%v", reqID, outboundTxID)
	b := store.Get(types.OutboundTxKey(
		strings.ToLower(key),
	))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllOutboundTx returns all outboundTx
func (k Keeper) GetAllOutboundTx(ctx sdk.Context) (list []types.OutboundTx) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OutboundTxKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.OutboundTx
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
