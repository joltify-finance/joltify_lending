package v5

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
)

// these codes are only for migration and may out of date

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	storeHandler := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.OutboundTxKeyPrefix))

	oldOutBoundIter := storeHandler.Iterator(nil, nil)
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

	newTxs := convertOutboundTxs(oldOutTxs)

	for _, el := range newTxs {
		each := el
		b := cdc.MustMarshal(&each)
		storeHandler.Set(types.OutboundTxKey(
			el.Index,
		), b)
	}

	// verify everything is ok
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

func convertOutboundTxs(old []types.OutboundTxV04) []types.OutboundTx {

	newOutboundTx := make([]types.OutboundTx, len(old))
	for i, el := range old {
		//newOutboundTx[i]

		setItemProposals(el.Items)

		types.OutboundTx{
			Index:           el.Index,
			Processed:       el.Processed,
			OutboundTxs:     el.OutboundTxs,
			ChainType:       el.ChainType,
			InTxHash:        el.InTxHash,
			ReceiverAddress: el.ReceiverAddress,
			NeedMint:        el.NeedMint,
			Feecoin:         el.Feecoin,
		}

	}

}

func setItemProposals(proposals map[string]types.ProposalsV04) {

}

func convertOutboundTx(old types.OutboundTxV04) types.OutboundTx {
	index := old.Index
	for txID, info := range old.Items {
		entities := make([]*types.Entity, len(info.Entry))
		for i, el := range info.Entry {
			e := types.Entity{
				Address: el.Address,
				Feecoin: []sdk.Coin{},
			}
			entities[i] = &e
		}
		proposals := types.Proposals{Entry: entities}
		_ = proposals
		_ = txID
		// items[txID] = proposals
	}
	return types.OutboundTx{
		Index:     index,
		Processed: true,
		ChainType: "",
		InTxHash:  "",
		NeedMint:  false,
	}
}
