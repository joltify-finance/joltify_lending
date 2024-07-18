package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/joltify-finance/joltify_lending/x/third_party/exchange/types"
)

func (k *Keeper) HasSubaccountAlreadyPlacedMarketOrder(ctx sdk.Context, marketID, subaccountID common.Hash) bool {
	// use transient store key
	store := k.getTransientStore(ctx)

	key := types.GetSubaccountMarketOrderIndicatorKey(marketID, subaccountID)

	return store.Has(key)
}

func (k *Keeper) HasSubaccountAlreadyPlacedLimitOrder(ctx sdk.Context, marketID, subaccountID common.Hash) bool {
	// use transient store key
	store := k.getTransientStore(ctx)

	key := types.GetSubaccountLimitOrderIndicatorKey(marketID, subaccountID)

	return store.Has(key)
}

func (k *Keeper) SetTransientSubaccountMarketOrderIndicator(ctx sdk.Context, marketID, subaccountID common.Hash) {
	// use transient store key
	store := k.getTransientStore(ctx)

	key := types.GetSubaccountMarketOrderIndicatorKey(marketID, subaccountID)
	store.Set(key, []byte{})
}

func (k *Keeper) SetTransientSubaccountLimitOrderIndicator(ctx sdk.Context, marketID, subaccountID common.Hash) {
	// use transient store key
	store := k.getTransientStore(ctx)

	key := types.GetSubaccountLimitOrderIndicatorKey(marketID, subaccountID)
	store.Set(key, []byte{})
}
