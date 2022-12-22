package v0_16

import (
	"github.com/cosmos/cosmos-sdk/types"
	v015pricefeed "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/legacy/v0_15"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"
)

var NewIBCMarkets = []types2.Market{
	{
		MarketID:   "atom:usd",
		BaseAsset:  "atom",
		QuoteAsset: "usd",
		Oracles:    nil,
		Active:     true,
	},
	{
		MarketID:   "atom:usd:30",
		BaseAsset:  "atom",
		QuoteAsset: "usd",
		Oracles:    nil,
		Active:     true,
	},
	{
		MarketID:   "akt:usd",
		BaseAsset:  "akt",
		QuoteAsset: "usd",
		Oracles:    nil,
		Active:     true,
	},
	{
		MarketID:   "akt:usd:30",
		BaseAsset:  "akt",
		QuoteAsset: "usd",
		Oracles:    nil,
		Active:     true,
	},
	{
		MarketID:   "luna:usd",
		BaseAsset:  "luna",
		QuoteAsset: "usd",
		Oracles:    nil,
		Active:     true,
	},
	{
		MarketID:   "luna:usd:30",
		BaseAsset:  "luna",
		QuoteAsset: "usd",
		Oracles:    nil,
		Active:     true,
	},
	{
		MarketID:   "osmo:usd",
		BaseAsset:  "osmo",
		QuoteAsset: "usd",
		Oracles:    nil,
		Active:     true,
	},
	{
		MarketID:   "osmo:usd:30",
		BaseAsset:  "osmo",
		QuoteAsset: "usd",
		Oracles:    nil,
		Active:     true,
	},
	{
		MarketID:   "ust:usd",
		BaseAsset:  "ust",
		QuoteAsset: "usd",
		Oracles:    nil,
		Active:     true,
	},
	{
		MarketID:   "ust:usd:30",
		BaseAsset:  "ust",
		QuoteAsset: "usd",
		Oracles:    nil,
		Active:     true,
	},
}

func migrateParams(params v015pricefeed.Params) types2.Params {
	markets := make(types2.Markets, len(params.Markets))
	for i, market := range params.Markets {
		markets[i] = types2.Market{
			MarketID:   market.MarketID,
			BaseAsset:  market.BaseAsset,
			QuoteAsset: market.QuoteAsset,
			Oracles:    market.Oracles,
			Active:     market.Active,
		}
	}

	markets = addIbcMarkets(markets)

	return types2.Params{Markets: markets}
}

func addIbcMarkets(markets types2.Markets) types2.Markets {
	var oracles []types.AccAddress

	if len(markets) > 0 {
		oracles = markets[0].Oracles
	}

	for _, newMarket := range NewIBCMarkets {
		// newMarket is a copy, should not affect other uses of NewIBCMarkets
		newMarket.Oracles = oracles
		markets = append(markets, newMarket)
	}

	return markets
}

func migratePostedPrices(oldPostedPrices v015pricefeed.PostedPrices) types2.PostedPrices {
	newPrices := make(types2.PostedPrices, len(oldPostedPrices))
	for i, price := range oldPostedPrices {
		newPrices[i] = types2.PostedPrice{
			MarketID:      price.MarketID,
			OracleAddress: price.OracleAddress,
			Price:         price.Price,
			Expiry:        price.Expiry,
		}
	}
	return newPrices
}

// Migrate converts v0.15 pricefeed state and returns it in v0.16 format
func Migrate(oldState v015pricefeed.GenesisState) *types2.GenesisState {
	return &types2.GenesisState{
		Params:       migrateParams(oldState.Params),
		PostedPrices: migratePostedPrices(oldState.PostedPrices),
	}
}
