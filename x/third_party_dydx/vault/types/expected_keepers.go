package types

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clobtypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/types"
	perptypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/perpetuals/types"
	pricestypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/types"
	sendingtypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/sending/types"
	satypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/subaccounts/types"
)

type ClobKeeper interface {
	// Clob Pair.
	GetClobPair(ctx sdk.Context, id clobtypes.ClobPairId) (val clobtypes.ClobPair, found bool)

	// Order.
	GetLongTermOrderPlacement(
		ctx sdk.Context,
		orderId clobtypes.OrderId,
	) (val clobtypes.LongTermOrderPlacement, found bool)
	HandleMsgCancelOrder(
		ctx sdk.Context,
		msg *clobtypes.MsgCancelOrder,
	) (err error)
	HandleMsgPlaceOrder(
		ctx sdk.Context,
		msg *clobtypes.MsgPlaceOrder,
		isInternalOrder bool,
	) (err error)
}

type PerpetualsKeeper interface {
	GetPerpetual(
		ctx sdk.Context,
		id uint32,
	) (val perptypes.Perpetual, err error)
}

type PricesKeeper interface {
	GetMarketParam(
		ctx sdk.Context,
		id uint32,
	) (market pricestypes.MarketParam, exists bool)
	GetMarketPrice(
		ctx sdk.Context,
		id uint32,
	) (pricestypes.MarketPrice, error)
}

type SendingKeeper interface {
	ProcessTransfer(
		ctx sdk.Context,
		pendingTransfer *sendingtypes.Transfer,
	) (err error)
}

type SubaccountsKeeper interface {
	GetNetCollateralAndMarginRequirements(
		ctx sdk.Context,
		update satypes.Update,
	) (
		bigNetCollateral *big.Int,
		bigInitialMargin *big.Int,
		bigMaintenanceMargin *big.Int,
		err error,
	)
	GetSubaccount(
		ctx sdk.Context,
		id satypes.SubaccountId,
	) satypes.Subaccount
}