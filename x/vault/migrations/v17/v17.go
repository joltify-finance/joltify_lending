// these codes are only for migration and may out of date

package keeper

import (
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
)

func convertOutboundTx(old types.OutboundTxV16) types.OutboundTx {
	index := old.Index
	items := make(map[string]types.Proposals)
	for txID, info := range old.Items {
		entities := make([]*types.Entity, len(info.Address))
		for i, el := range info.Address {
			e := types.Entity{
				Address: el,
				Feecoin: []sdk.Coin{},
			}
			entities[i] = &e
		}
		proposals := types.Proposals{Entry: entities}
		items[txID] = proposals
	}
	return types.OutboundTx{
		Index:     index,
		Processed: true,
		// Items:     items,
	}
}

func MigrateStore(ctx context.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	storeHandler := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.OutboundTxKeyPrefix))

	oldOutBoundIter := storeHandler.Iterator(nil, nil)
	var newTxs []types.OutboundTx
	for ; oldOutBoundIter.Valid(); oldOutBoundIter.Next() {
		var oldOutTx types.OutboundTxV16
		if err := cdc.Unmarshal(oldOutBoundIter.Value(), &oldOutTx); err != nil {
			return err
		}
		newTx := convertOutboundTx(oldOutTx)
		newTxs = append(newTxs, newTx)
	}
	err := oldOutBoundIter.Close()
	if err != nil {
		return err
	}

	for _, el := range newTxs {
		each := el
		b := cdc.MustMarshal(&each)
		storeHandler.Set(types.OutboundTxKey(
			el.Index,
		), b)
	}

	oldOutBoundIter2 := storeHandler.Iterator(nil, nil)
	for ; oldOutBoundIter2.Valid(); oldOutBoundIter2.Next() {
		var oldOutTx types.OutboundTx
		if err := cdc.Unmarshal(oldOutBoundIter2.Value(), &oldOutTx); err != nil {
			return err
		}
	}
	oldOutBoundIter2.Close()

	return nil
}
