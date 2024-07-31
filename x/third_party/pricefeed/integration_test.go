package pricefeed_test

import (
	"time"

	sdkmath "cosmossdk.io/math"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/app"
)

func NewPricefeedGen() types2.GenesisState {
	return types2.GenesisState{
		Params: types2.Params{
			Markets: []types2.Market{
				{MarketID: "btc:usd", BaseAsset: "btc", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "xrp:usd", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
			},
		},
		PostedPrices: []types2.PostedPrice{
			{
				MarketID:      "btc:usd",
				OracleAddress: sdk.AccAddress("oracle1"),
				Price:         sdkmath.LegacyMustNewDecFromStr("8000.00"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketID:      "xrp:usd",
				OracleAddress: sdk.AccAddress("oracle2"),
				Price:         sdkmath.LegacyMustNewDecFromStr("0.25"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
		},
	}
}

func NewPricefeedGenStateMulti() app.GenesisState {
	pfGenesis := NewPricefeedGen()
	return app.GenesisState{types2.ModuleName: types2.ModuleCdc.LegacyAmino.MustMarshalJSON(pfGenesis)}
}

func NewPricefeedGenStateWithOracles(addrs []sdk.AccAddress) app.GenesisState {
	pfGenesis := types2.GenesisState{
		Params: types2.Params{
			Markets: []types2.Market{
				{MarketID: "btc:usd", BaseAsset: "btc", QuoteAsset: "usd", Oracles: addrs, Active: true},
				{MarketID: "xrp:usd", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: addrs, Active: true},
			},
		},
		PostedPrices: []types2.PostedPrice{
			{
				MarketID:      "btc:usd",
				OracleAddress: addrs[0],
				Price:         sdkmath.LegacyMustNewDecFromStr("8000.00"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketID:      "xrp:usd",
				OracleAddress: addrs[0],
				Price:         sdkmath.LegacyMustNewDecFromStr("0.25"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
		},
	}
	return app.GenesisState{types2.ModuleName: types2.ModuleCdc.LegacyAmino.MustMarshalJSON(pfGenesis)}
}
