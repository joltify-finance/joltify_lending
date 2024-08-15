package memclob

import (
	"testing"

	sdktest "github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/sdk"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/types"
	satypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/subaccounts/types"
	"github.com/stretchr/testify/require"
)

func TestGetOrder_Success(t *testing.T) {
	ctx, _, _ := sdktest.NewSdkContextWithMultistore()

	memclob := NewMemClobPriceTimePriority(false)

	orderId := types.OrderId{
		SubaccountId: satypes.SubaccountId{
			Owner: "testGetOrder",
		},
	}
	order := types.Order{OrderId: orderId}

	memclob.openOrders.orderIdToLevelOrder[orderId] = &types.LevelOrder{
		Value: types.ClobOrder{
			Order: order,
		},
	}

	foundOrder, found := memclob.GetOrder(ctx, orderId)
	require.True(t, found)
	require.Equal(t, order, foundOrder)
}

func TestGetOrder_ErrDoesNotExist(t *testing.T) {
	ctx, _, _ := sdktest.NewSdkContextWithMultistore()

	memclob := NewMemClobPriceTimePriority(false)

	_, found := memclob.GetOrder(ctx, types.OrderId{})
	require.False(t, found)
}
