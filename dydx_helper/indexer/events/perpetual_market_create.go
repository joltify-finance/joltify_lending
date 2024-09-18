package events

import (
	v1 "github.com/joltify-finance/joltify_lending/dydx_helper/indexer/protocol/v1"
	clobtypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/types"
	perptypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/perpetuals/types"
)

// NewPerpetualMarketCreateEvent creates a PerpetualMarketCreateEvent
// representing creation of a perpetual market.
func NewPerpetualMarketCreateEvent(
	id uint32,
	clobPairId uint32,
	ticker string,
	marketId uint32,
	status clobtypes.ClobPair_Status,
	quantumConversionExponent int32,
	atomicResolution int32,
	subticksPerTick uint32,
	stepBaseQuantums uint64,
	liquidityTier uint32,
	marketType perptypes.PerpetualMarketType,
) *PerpetualMarketCreateEventV2 {
	return &PerpetualMarketCreateEventV2{
		Id:                        id,
		ClobPairId:                clobPairId,
		Ticker:                    ticker,
		MarketId:                  marketId,
		Status:                    v1.ConvertToClobPairStatus(status),
		QuantumConversionExponent: quantumConversionExponent,
		AtomicResolution:          atomicResolution,
		SubticksPerTick:           subticksPerTick,
		StepBaseQuantums:          stepBaseQuantums,
		LiquidityTier:             liquidityTier,
		MarketType:                v1.ConvertToPerpetualMarketType(marketType),
	}
}
