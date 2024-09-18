package events

import (
	"testing"

	v1types "github.com/joltify-finance/joltify_lending/dydx_helper/indexer/protocol/v1/types"
	clobtypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/types"
	perptypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/perpetuals/types"

	"github.com/stretchr/testify/require"
)

func TestNewPerpetualMarketCreateEvent_Success(t *testing.T) {
	perpetualMarketCreateEvent := NewPerpetualMarketCreateEvent(
		0,
		0,
		"BTC",
		0,
		clobtypes.ClobPair_STATUS_ACTIVE,
		-8,
		8,
		5,
		5,
		0,
		perptypes.PerpetualMarketType_PERPETUAL_MARKET_TYPE_CROSS,
	)
	expectedPerpetualMarketCreateEventProto := &PerpetualMarketCreateEventV2{
		Id:                        0,
		ClobPairId:                0,
		Ticker:                    "BTC",
		MarketId:                  0,
		Status:                    v1types.ClobPairStatus_CLOB_PAIR_STATUS_ACTIVE,
		QuantumConversionExponent: -8,
		AtomicResolution:          8,
		SubticksPerTick:           5,
		StepBaseQuantums:          5,
		LiquidityTier:             0,
		MarketType:                v1types.PerpetualMarketType_PERPETUAL_MARKET_TYPE_CROSS,
	}
	require.Equal(t, expectedPerpetualMarketCreateEventProto, perpetualMarketCreateEvent)
}
