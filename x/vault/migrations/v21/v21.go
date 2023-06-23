// Package v21 these codes are only for migration and may out of date
package v21

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	types2 "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/types"
	types3 "github.com/joltify-finance/joltify_lending/x/vault/types"
)

func MigrateStore(ctx types.Context, storeKey types2.StoreKey, cdc codec.BinaryCodec) error {
	storeHandler := prefix.NewStore(ctx.KVStore(storeKey), types3.KeyPrefix(types3.OutboundTxKeyPrefix))

	oldOutBoundIter := storeHandler.Iterator(nil, nil)
	var newTxs []types3.OutboundTx
	for ; oldOutBoundIter.Valid(); oldOutBoundIter.Next() {
		var oldOutTx types3.OutboundTxV120
		if err := cdc.Unmarshal(oldOutBoundIter.Value(), &oldOutTx); err != nil {
			return err
		}
		newTx := convertOutboundTx(oldOutTx)
		newTxs = append(newTxs, newTx)
	}
	oldOutBoundIter.Close()

	for _, el := range newTxs {
		each := el
		b := cdc.MustMarshal(&each)
		storeHandler.Set(types3.OutboundTxKey(
			el.Index,
		), b)
	}

	oldOutBoundIter2 := storeHandler.Iterator(nil, nil)
	for ; oldOutBoundIter2.Valid(); oldOutBoundIter2.Next() {
		var oldOutTx types3.OutboundTx
		if err := cdc.Unmarshal(oldOutBoundIter2.Value(), &oldOutTx); err != nil {
			return err
		}
	}
	oldOutBoundIter2.Close()

	return nil
}

func convertOutboundTx(old types3.OutboundTxV120) types3.OutboundTx {
	index := old.Index
	for txID, info := range old.Items {
		entities := make([]*types3.Entity, len(info.Entry))
		for i, el := range info.Entry {
			e := types3.Entity{
				Address: el.Address,
				Feecoin: []types.Coin{},
			}
			entities[i] = &e
		}
		proposals := types3.Proposals{Entry: entities}
		_ = proposals
		_ = txID
		// items[txID] = proposals
	}
	return types3.OutboundTx{
		Index:     index,
		Processed: true,
		ChainType: "",
		InTxHash:  "",
		NeedMint:  false,
	}
}
