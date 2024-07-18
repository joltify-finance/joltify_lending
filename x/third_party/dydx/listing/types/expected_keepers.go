package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	clobtypes "github.com/joltify-finance/joltify_lending/x/third_party/dydx/clob/types"
	perpetualtypes "github.com/joltify-finance/joltify_lending/x/third_party/dydx/perpetuals/types"
	pricestypes "github.com/joltify-finance/joltify_lending/x/third_party/dydx/prices/types"
	satypes "github.com/joltify-finance/joltify_lending/x/third_party/dydx/subaccounts/types"
	marketmaptypes "github.com/skip-mev/slinky/x/marketmap/types"
)

type PricesKeeper interface {
	CreateMarket(
		ctx sdk.Context,
		marketParam pricestypes.MarketParam,
		marketPrice pricestypes.MarketPrice,
	) (pricestypes.MarketParam, error)
	AcquireNextMarketID(ctx sdk.Context) uint32
}

type ClobKeeper interface {
	CreatePerpetualClobPair(
		ctx sdk.Context,
		clobPairId uint32,
		perpetualId uint32,
		stepSizeBaseQuantums satypes.BaseQuantums,
		quantumConversionExponent int32,
		subticksPerTick uint32,
		status clobtypes.ClobPair_Status,
	) (clobtypes.ClobPair, error)
	AcquireNextClobPairID(ctx sdk.Context) uint32
}

type MarketMapKeeper interface {
	GetMarket(
		ctx sdk.Context,
		ticker string,
	) (marketmaptypes.Market, error)
}

type PerpetualsKeeper interface {
	CreatePerpetual(
		ctx sdk.Context,
		id uint32,
		ticker string,
		marketId uint32,
		atomicResolution int32,
		defaultFundingPpm int32,
		liquidityTier uint32,
		marketType perpetualtypes.PerpetualMarketType,
	) (perpetualtypes.Perpetual, error)
	AcquireNextPerpetualID(ctx sdk.Context) uint32
}
