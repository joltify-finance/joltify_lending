package prices

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	pricestypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/types"
)

// PriceUpdateGenerator is an interface to abstract the logic of retrieving a
// `MsgUpdateMarketPrices` for any block.
type PriceUpdateGenerator interface {
	GetValidMarketPriceUpdates(ctx sdk.Context, extCommitBz []byte) (*pricestypes.MsgUpdateMarketPrices, error)
}
