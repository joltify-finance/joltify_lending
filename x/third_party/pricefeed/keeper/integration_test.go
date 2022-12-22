package keeper_test

import (
	"time"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/app"
)

func NewPricefeedGenStateMulti() app.GenesisState {
	pfGenesis := types2.GenesisState{
		Params: types2.Params{
			Markets: []types2.Market{
				{MarketID: "btc:usd", BaseAsset: "btc", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "xrp:usd", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
			},
		},
		PostedPrices: []types2.PostedPrice{
			{
				MarketID:      "btc:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("8000.00"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketID:      "xrp:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("0.25"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
		},
	}
	return app.GenesisState{types2.ModuleName: types2.ModuleCdc.LegacyAmino.MustMarshalJSON(pfGenesis)}
}
