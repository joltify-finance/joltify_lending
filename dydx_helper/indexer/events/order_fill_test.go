package events_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/dydx_helper/indexer/events"
	v1 "github.com/joltify-finance/joltify_lending/dydx_helper/indexer/protocol/v1"
	"github.com/joltify-finance/joltify_lending/dydx_helper/testutil/constants"
	clobtypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/types"
	satypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/subaccounts/types"
	"github.com/stretchr/testify/require"
)

var (
	makerOrder            = constants.Order_Alice_Num0_Id0_Clob0_Buy5_Price10_GTB15
	indexerMakerOrder     = v1.OrderToIndexerOrder(makerOrder)
	takerOrder            = constants.Order_Alice_Num0_Id2_Clob1_Sell5_Price10_GTB15
	indexerTakerOrder     = v1.OrderToIndexerOrder(takerOrder)
	liquidationTakerOrder = constants.LiquidationOrder_Carl_Num0_Clob0_Buy3_Price50_BTC
	fillAmount            = satypes.BaseQuantums(5)
	makerFee              = int64(-2)
	takerFee              = int64(5)
)

func TestNewOrderFillEvent_Success(t *testing.T) {
	orderFillEvent := events.NewOrderFillEvent(
		makerOrder,
		takerOrder,
		fillAmount,
		makerFee,
		takerFee,
		fillAmount,
		fillAmount,
	)

	expectedOrderFillEventProto := &events.OrderFillEventV1{
		MakerOrder: indexerMakerOrder,
		TakerOrder: &events.OrderFillEventV1_Order{
			Order: &indexerTakerOrder,
		},
		FillAmount:       fillAmount.ToUint64(),
		MakerFee:         makerFee,
		TakerFee:         takerFee,
		TotalFilledMaker: fillAmount.ToUint64(),
		TotalFilledTaker: fillAmount.ToUint64(),
	}
	require.Equal(t, expectedOrderFillEventProto, orderFillEvent)
}

func TestNewLiquidationOrderFillEvent_Success(t *testing.T) {
	var matchableTakerOrder clobtypes.MatchableOrder = &liquidationTakerOrder
	liquidationOrderFillEvent := events.NewLiquidationOrderFillEvent(
		makerOrder,
		matchableTakerOrder,
		fillAmount,
		makerFee,
		takerFee,
		fillAmount,
	)

	expectedLiquidationOrder := events.LiquidationOrderV1{
		Liquidated:  v1.SubaccountIdToIndexerSubaccountId(liquidationTakerOrder.GetSubaccountId()),
		ClobPairId:  liquidationTakerOrder.GetClobPairId().ToUint32(),
		PerpetualId: liquidationTakerOrder.MustGetLiquidatedPerpetualId(),
		TotalSize:   uint64(liquidationTakerOrder.GetBaseQuantums()),
		IsBuy:       liquidationTakerOrder.IsBuy(),
		Subticks:    uint64(liquidationTakerOrder.GetOrderSubticks()),
	}
	expectedOrderFillEventProto := &events.OrderFillEventV1{
		MakerOrder: indexerMakerOrder,
		TakerOrder: &events.OrderFillEventV1_LiquidationOrder{
			LiquidationOrder: &expectedLiquidationOrder,
		},
		FillAmount:       fillAmount.ToUint64(),
		MakerFee:         makerFee,
		TakerFee:         takerFee,
		TotalFilledMaker: fillAmount.ToUint64(),
		TotalFilledTaker: fillAmount.ToUint64(),
	}
	require.Equal(t, expectedOrderFillEventProto, liquidationOrderFillEvent)
}
