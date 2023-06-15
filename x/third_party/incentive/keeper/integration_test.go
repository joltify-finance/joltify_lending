package keeper_test

import (
	"time"

	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/testutil"
	"github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/joltify-finance/joltify_lending/app"
)

// Avoid cluttering test cases with long function names
func i(in int64) sdk.Int                    { return sdk.NewInt(in) }
func d(str string) sdk.Dec                  { return sdk.MustNewDecFromStr(str) }
func c(denom string, amount int64) sdk.Coin { return sdk.NewInt64Coin(denom, amount) }
func cs(coins ...sdk.Coin) sdk.Coins        { return sdk.NewCoins(coins...) }

func NewPricefeedGenStateMultiFromTime(cdc codec.JSONCodec, t time.Time) app.GenesisState {
	expiry := 100 * 365 * 24 * time.Hour // 100 years

	pfGenesis := types.GenesisState{
		Params: types.Params{
			Markets: []types.Market{
				{MarketID: "jolt:usd", BaseAsset: "jolt", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "btc:usd", BaseAsset: "btc", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "xrp:usd", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "bnb:usd", BaseAsset: "bnb", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "sbnb:usd", BaseAsset: "sbnb", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "busd:usd", BaseAsset: "busd", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "zzz:usd", BaseAsset: "zzz", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "pjolt:usd", BaseAsset: "pjolt", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
			},
		},
		PostedPrices: []types.PostedPrice{
			{
				MarketID:      "jolt:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("2.00"),
				Expiry:        t.Add(expiry),
			},
			{
				MarketID:      "btc:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("8000.00"),
				Expiry:        t.Add(expiry),
			},
			{
				MarketID:      "xrp:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("0.25"),
				Expiry:        t.Add(expiry),
			},
			{
				MarketID:      "bnb:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("17.25"),
				Expiry:        t.Add(expiry),
			},
			{
				MarketID:      "sbnb:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.OneDec(),
				Expiry:        t.Add(expiry),
			},
			{
				MarketID:      "busd:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.OneDec(),
				Expiry:        t.Add(expiry),
			},
			{
				MarketID:      "zzz:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("2.00"),
				Expiry:        t.Add(expiry),
			},
			{
				MarketID:      "pjolt:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("2.00"),
				Expiry:        t.Add(expiry),
			},
		},
	}
	return app.GenesisState{types.ModuleName: cdc.MustMarshalJSON(&pfGenesis)}
}

func NewJoltGenStateMulti(genTime time.Time) testutil.JoltGenesisBuilder {
	joltMM := testutil.NewStandardMoneyMarket("ujolt")
	joltMM.SpotMarketID = "jolt:usd"
	btcMM := testutil.NewStandardMoneyMarket("btcb")
	btcMM.SpotMarketID = "btc:usd"

	pjoltMM := testutil.NewStandardMoneyMarket("pjolt")
	pjoltMM.SpotMarketID = "pjolt:usd"

	builder := testutil.NewJoltGenesisBuilder().WithGenesisTime(genTime).
		WithInitializedMoneyMarket(testutil.NewStandardMoneyMarket("usdx")).
		WithInitializedMoneyMarket(joltMM).
		WithInitializedMoneyMarket(testutil.NewStandardMoneyMarket("bnb")).
		WithInitializedMoneyMarket(testutil.NewStandardMoneyMarket("sbnb")).
		WithInitializedMoneyMarket(btcMM).
		WithInitializedMoneyMarket(testutil.NewStandardMoneyMarket("xrp")).
		WithInitializedMoneyMarket(testutil.NewStandardMoneyMarket("zzz")).
		WithInitializedMoneyMarket(pjoltMM)
	return builder
}

func NewStakingGenesisState(cdc codec.JSONCodec) app.GenesisState {
	genState := stakingtypes.DefaultGenesisState()
	genState.Params.BondDenom = "ujolt"
	return app.GenesisState{
		stakingtypes.ModuleName: cdc.MustMarshalJSON(genState),
	}
}
