package incentive_test

import (
	"time"

	"github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
)

// Avoid cluttering test cases with long function names
func i(in int64) sdk.Int                    { return sdk.NewInt(in) }
func d(str string) sdk.Dec                  { return sdk.MustNewDecFromStr(str) }
func cs(coins ...sdk.Coin) sdk.Coins        { return sdk.NewCoins(coins...) }
func c(denom string, amount int64) sdk.Coin { return sdk.NewInt64Coin(denom, amount) }

func NewPricefeedGenStateMultiFromTime(cdc codec.JSONCodec, t time.Time) app.GenesisState {
	pfGenesis := types.GenesisState{
		Params: types.Params{
			Markets: []types.Market{
				{MarketID: "jolt:usd", BaseAsset: "jolt", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "btc:usd", BaseAsset: "btc", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "xrp:usd", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "bnb:usd", BaseAsset: "bnb", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "busd:usd", BaseAsset: "busd", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "zzz:usd", BaseAsset: "zzz", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
			},
		},
		PostedPrices: []types.PostedPrice{
			{
				MarketID:      "jolt:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("2.00"),
				Expiry:        t.Add(1 * time.Hour),
			},
			{
				MarketID:      "btc:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("8000.00"),
				Expiry:        t.Add(1 * time.Hour),
			},
			{
				MarketID:      "xrp:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("0.25"),
				Expiry:        t.Add(1 * time.Hour),
			},
			{
				MarketID:      "bnb:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("17.25"),
				Expiry:        t.Add(1 * time.Hour),
			},
			{
				MarketID:      "busd:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.OneDec(),
				Expiry:        t.Add(1 * time.Hour),
			},
			{
				MarketID:      "zzz:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("2.00"),
				Expiry:        t.Add(1 * time.Hour),
			},
		},
	}
	return app.GenesisState{types.ModuleName: cdc.MustMarshalJSON(&pfGenesis)}
}

// func NewJoltGenStateMulti(genTime time.Time) testutil.JoltGenesisBuilder {
//	joltMM := testutil.NewStandardMoneyMarket("ujolt")
//	joltMM.SpotMarketID = "jolt:usd"
//	btcMM := testutil.NewStandardMoneyMarket("btcb")
//	btcMM.SpotMarketID = "btc:usd"
//
//	builder := testutil.NewJoltGenesisBuilder().WithGenesisTime(genTime).
//		WithInitializedMoneyMarket(testutil.NewStandardMoneyMarket("usdx")).
//		WithInitializedMoneyMarket(joltMM).
//		WithInitializedMoneyMarket(testutil.NewStandardMoneyMarket("bnb")).
//		WithInitializedMoneyMarket(btcMM).
//		WithInitializedMoneyMarket(testutil.NewStandardMoneyMarket("xrp")).
//		WithInitializedMoneyMarket(testutil.NewStandardMoneyMarket("zzz"))
//	return builder
// }

// func NewStakingGenesisState(cdc codec.JSONCodec) app.GenesisState {
//	genState := stakingtypes.DefaultGenesisState()
//	genState.Params.BondDenom = "ujolt"
//	return app.GenesisState{
//		stakingtypes.ModuleName: cdc.MustMarshalJSON(genState),
//	}
// }
