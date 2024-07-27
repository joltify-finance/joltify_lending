package v5

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
)

// these codes are only for migration and may out of date

func MigrateStore(rctx context.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	oldPrefix := "OutboundTx/value/"
	ctx := sdk.UnwrapSDKContext(rctx)
	storeHandler := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.OutboundTxKeyPrefix))
	storeHandlerOld := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(oldPrefix))

	oldOutBoundIter := storeHandlerOld.Iterator(nil, nil)
	var oldOutTxs []types.OutboundTxV04
	for ; oldOutBoundIter.Valid(); oldOutBoundIter.Next() {
		var oldOutTx types.OutboundTxV04
		if err := cdc.Unmarshal(oldOutBoundIter.Value(), &oldOutTx); err != nil {
			return err
		}
		oldOutTxs = append(oldOutTxs, oldOutTx)
	}
	err := oldOutBoundIter.Close()
	if err != nil {
		panic(err)
	}

	if len(oldOutTxs) == 0 {
		panic("fail to load the old outbound txs")
	}

	newTxs := convertOutboundTxsAndSetProposal(ctx, storeKey, cdc, oldOutTxs)

	for _, el := range newTxs {
		each := el
		b := cdc.MustMarshal(&each)
		storeHandler.Set(types.OutboundTxKey(
			el.Index,
		), b)
	}

	// verify everything is ok
	newOutBoundIter2 := storeHandler.Iterator(nil, nil)
	for ; newOutBoundIter2.Valid(); newOutBoundIter2.Next() {
		var oldOutTx types.OutboundTx
		if err := cdc.Unmarshal(newOutBoundIter2.Value(), &oldOutTx); err != nil {
			return err
		}
	}
	err = newOutBoundIter2.Close()
	if err != nil {
		panic(err)
	}

	// now we delete the old stored data
	for _, el := range oldOutTxs {
		da := storeHandlerOld.Get(types.OutboundTxKey(el.Index))
		if da == nil {
			panic("fail to find the old outbound txs")
		}
		storeHandlerOld.Delete(types.OutboundTxKey(
			el.Index,
		))
	}

	return nil
}

func convertOutboundTxsAndSetProposal(ctx context.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, old []types.OutboundTxV04) []types.OutboundTx {
	newOutboundTxs := make([]types.OutboundTx, len(old))
	for i, el := range old {
		// newOutboundTx[i]
		outboundTxs := setItemProposals(ctx, storeKey, cdc, el.Index, el.Items)
		newOutboundTx := types.OutboundTx{
			Index:           el.Index,
			Processed:       el.Processed,
			OutboundTxs:     outboundTxs,
			ChainType:       el.ChainType,
			InTxHash:        el.InTxHash,
			ReceiverAddress: el.ReceiverAddress,
			NeedMint:        el.NeedMint,
			Feecoin:         el.Feecoin,
		}
		newOutboundTxs[i] = newOutboundTx
	}
	return newOutboundTxs
}

func setItemProposals(rctx context.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, reqIndex string, proposals map[string]types.ProposalsV04) []string {
	ctx := sdk.UnwrapSDKContext(rctx)
	items := make([]string, 0, len(proposals))
	for txID := range proposals {
		items = append(items, txID)
	}
	sort.Strings(items)
	for _, el := range items {
		proposal, found := proposals[el]
		if !found {
			panic("item must be found")
		}
		storeHandler := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.OutboundTxProposalKeyPrefix))
		b := cdc.MustMarshal(&proposal)
		key := fmt.Sprintf("%v:%v", reqIndex, el)
		storeHandler.Set(types.OutboundTxKey(
			strings.ToLower(key),
		), b)
	}
	return items
}
