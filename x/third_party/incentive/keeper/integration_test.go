package keeper_test

import (
	"time"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/cdp/types"
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

func NewCDPGenStateMulti(cdc codec.JSONCodec) app.GenesisState {
	cdpGenesis := types2.GenesisState{
		Params: types2.Params{
			GlobalDebtLimit:         sdk.NewInt64Coin("usdx", 2000000000000),
			SurplusAuctionThreshold: types2.DefaultSurplusThreshold,
			SurplusAuctionLot:       types2.DefaultSurplusLot,
			DebtAuctionThreshold:    types2.DefaultDebtThreshold,
			DebtAuctionLot:          types2.DefaultDebtLot,
			CollateralParams: types2.CollateralParams{
				{
					Denom:               "xrp",
					Type:                "xrp-a",
					LiquidationRatio:    sdk.MustNewDecFromStr("2.0"),
					DebtLimit:           sdk.NewInt64Coin("usdx", 500000000000),
					StabilityFee:        sdk.MustNewDecFromStr("1.000000001547125958"), // %5 apr
					LiquidationPenalty:  d("0.05"),
					AuctionSize:         i(7000000000),
					SpotMarketID:        "xrp:usd",
					LiquidationMarketID: "xrp:usd",
					ConversionFactor:    i(6),
				},
				{
					Denom:               "btc",
					Type:                "btc-a",
					LiquidationRatio:    sdk.MustNewDecFromStr("1.5"),
					DebtLimit:           sdk.NewInt64Coin("usdx", 500000000000),
					StabilityFee:        sdk.MustNewDecFromStr("1.000000000782997609"), // %2.5 apr
					LiquidationPenalty:  d("0.025"),
					AuctionSize:         i(10000000),
					SpotMarketID:        "btc:usd",
					LiquidationMarketID: "btc:usd",
					ConversionFactor:    i(8),
				},
				{
					Denom:               "bnb",
					Type:                "bnb-a",
					LiquidationRatio:    sdk.MustNewDecFromStr("1.5"),
					DebtLimit:           sdk.NewInt64Coin("usdx", 500000000000),
					StabilityFee:        sdk.MustNewDecFromStr("1.000000001547125958"), // %5 apr
					LiquidationPenalty:  d("0.05"),
					AuctionSize:         i(50000000000),
					SpotMarketID:        "bnb:usd",
					LiquidationMarketID: "bnb:usd",
					ConversionFactor:    i(8),
				},
				{
					Denom:               "busd",
					Type:                "busd-a",
					LiquidationRatio:    d("1.01"),
					DebtLimit:           sdk.NewInt64Coin("usdx", 500000000000),
					StabilityFee:        sdk.OneDec(), // %0 apr
					LiquidationPenalty:  d("0.05"),
					AuctionSize:         i(10000000000),
					SpotMarketID:        "busd:usd",
					LiquidationMarketID: "busd:usd",
					ConversionFactor:    i(8),
				},
			},
			DebtParam: types2.DebtParam{
				Denom:            "usdx",
				ReferenceAsset:   "usd",
				ConversionFactor: i(6),
				DebtFloor:        i(10000000),
			},
		},
		StartingCdpID: types2.DefaultCdpStartingID,
		DebtDenom:     types2.DefaultDebtDenom,
		GovDenom:      types2.DefaultGovDenom,
		CDPs:          types2.CDPs{},
		PreviousAccumulationTimes: types2.GenesisAccumulationTimes{
			types2.NewGenesisAccumulationTime("btc-a", time.Time{}, sdk.OneDec()),
			types2.NewGenesisAccumulationTime("xrp-a", time.Time{}, sdk.OneDec()),
			types2.NewGenesisAccumulationTime("busd-a", time.Time{}, sdk.OneDec()),
			types2.NewGenesisAccumulationTime("bnb-a", time.Time{}, sdk.OneDec()),
		},
		TotalPrincipals: types2.GenesisTotalPrincipals{
			types2.NewGenesisTotalPrincipal("btc-a", sdk.ZeroInt()),
			types2.NewGenesisTotalPrincipal("xrp-a", sdk.ZeroInt()),
			types2.NewGenesisTotalPrincipal("busd-a", sdk.ZeroInt()),
			types2.NewGenesisTotalPrincipal("bnb-a", sdk.ZeroInt()),
		},
	}
	return app.GenesisState{types2.ModuleName: cdc.MustMarshalJSON(&cdpGenesis)}
}

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
