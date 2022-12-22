package keeper_test

import (
	"time"

	types3 "github.com/joltify-finance/joltify_lending/x/third_party/cdp/types"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/joltify-finance/joltify_lending/app"
)

// Avoid cluttering test cases with long function names
func i(in int64) sdk.Int                    { return sdk.NewInt(in) }
func d(str string) sdk.Dec                  { return sdk.MustNewDecFromStr(str) }
func c(denom string, amount int64) sdk.Coin { return sdk.NewInt64Coin(denom, amount) }
func cs(coins ...sdk.Coin) sdk.Coins        { return sdk.NewCoins(coins...) }

func NewPricefeedGenState(cdc codec.JSONCodec, asset string, price sdk.Dec) app.GenesisState {
	pfGenesis := types2.GenesisState{
		Params: types2.Params{
			Markets: []types2.Market{
				{MarketID: asset + ":usd", BaseAsset: asset, QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
			},
		},
		PostedPrices: []types2.PostedPrice{
			{
				MarketID:      asset + ":usd",
				OracleAddress: sdk.AccAddress{},
				Price:         price,
				Expiry:        time.Now().Add(1 * time.Hour),
			},
		},
	}
	return app.GenesisState{types2.ModuleName: cdc.MustMarshalJSON(&pfGenesis)}
}

func NewCDPGenState(cdc codec.JSONCodec, asset string, liquidationRatio sdk.Dec) app.GenesisState {
	cdpGenesis := types3.GenesisState{
		Params: types3.Params{
			GlobalDebtLimit:         sdk.NewInt64Coin("usdx", 1000000000000),
			SurplusAuctionThreshold: types3.DefaultSurplusThreshold,
			SurplusAuctionLot:       types3.DefaultSurplusLot,
			DebtAuctionThreshold:    types3.DefaultDebtThreshold,
			DebtAuctionLot:          types3.DefaultDebtLot,
			CollateralParams: types3.CollateralParams{
				{
					Denom:                            asset,
					Type:                             asset + "-a",
					LiquidationRatio:                 liquidationRatio,
					DebtLimit:                        sdk.NewInt64Coin("usdx", 1000000000000),
					StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"), // %5 apr
					LiquidationPenalty:               d("0.05"),
					AuctionSize:                      i(100),
					SpotMarketID:                     asset + ":usd",
					LiquidationMarketID:              asset + ":usd",
					KeeperRewardPercentage:           d("0.01"),
					CheckCollateralizationIndexCount: i(10),
					ConversionFactor:                 i(6),
				},
			},
			DebtParam: types3.DebtParam{
				Denom:            "usdx",
				ReferenceAsset:   "usd",
				ConversionFactor: i(6),
				DebtFloor:        i(10000000),
			},
		},
		StartingCdpID: types3.DefaultCdpStartingID,
		DebtDenom:     types3.DefaultDebtDenom,
		GovDenom:      types3.DefaultGovDenom,
		CDPs:          types3.CDPs{},
		PreviousAccumulationTimes: types3.GenesisAccumulationTimes{
			types3.NewGenesisAccumulationTime(asset+"-a", time.Time{}, sdk.OneDec()),
		},
		TotalPrincipals: types3.GenesisTotalPrincipals{
			types3.NewGenesisTotalPrincipal(asset+"-a", sdk.ZeroInt()),
		},
	}
	return app.GenesisState{types3.ModuleName: cdc.MustMarshalJSON(&cdpGenesis)}
}

func NewPricefeedGenStateMulti(cdc codec.JSONCodec) app.GenesisState {
	pfGenesis := types2.GenesisState{
		Params: types2.Params{
			Markets: []types2.Market{
				{MarketID: "btc:usd", BaseAsset: "btc", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "btc:usd:30", BaseAsset: "btc", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "xrp:usd", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "xrp:usd:30", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "bnb:usd", BaseAsset: "bnb", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "bnb:usd:30", BaseAsset: "bnb", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "busd:usd", BaseAsset: "busd", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "busd:usd:30", BaseAsset: "busd", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
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
				MarketID:      "btc:usd:30",
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
			{
				MarketID:      "xrp:usd:30",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("0.25"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketID:      "bnb:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("17.25"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketID:      "bnb:usd:30",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("17.25"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketID:      "busd:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.OneDec(),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketID:      "busd:usd:30",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.OneDec(),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
		},
	}
	return app.GenesisState{types2.ModuleName: cdc.MustMarshalJSON(&pfGenesis)}
}

func NewCDPGenStateMulti(cdc codec.JSONCodec) app.GenesisState {
	cdpGenesis := types3.GenesisState{
		Params: types3.Params{
			GlobalDebtLimit:         sdk.NewInt64Coin("usdx", 2000000000000),
			SurplusAuctionThreshold: types3.DefaultSurplusThreshold,
			SurplusAuctionLot:       types3.DefaultSurplusLot,
			DebtAuctionThreshold:    types3.DefaultDebtThreshold,
			DebtAuctionLot:          types3.DefaultDebtLot,
			CollateralParams: types3.CollateralParams{
				{
					Denom:                            "xrp",
					Type:                             "xrp-a",
					LiquidationRatio:                 sdk.MustNewDecFromStr("2.0"),
					DebtLimit:                        sdk.NewInt64Coin("usdx", 500000000000),
					StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"), // %5 apr
					LiquidationPenalty:               d("0.05"),
					AuctionSize:                      i(7000000000),
					SpotMarketID:                     "xrp:usd",
					LiquidationMarketID:              "xrp:usd:30",
					KeeperRewardPercentage:           d("0.01"),
					CheckCollateralizationIndexCount: i(10),
					ConversionFactor:                 i(6),
				},
				{
					Denom:                            "btc",
					Type:                             "btc-a",
					LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
					DebtLimit:                        sdk.NewInt64Coin("usdx", 500000000000),
					StabilityFee:                     sdk.MustNewDecFromStr("1.000000000782997609"), // %2.5 apr
					LiquidationPenalty:               d("0.025"),
					AuctionSize:                      i(10000000),
					SpotMarketID:                     "btc:usd",
					LiquidationMarketID:              "btc:usd:30",
					KeeperRewardPercentage:           d("0.01"),
					CheckCollateralizationIndexCount: i(10),
					ConversionFactor:                 i(8),
				},
				{
					Denom:                            "bnb",
					Type:                             "bnb-a",
					LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
					DebtLimit:                        sdk.NewInt64Coin("usdx", 500000000000),
					StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"), // %5 apr
					LiquidationPenalty:               d("0.05"),
					AuctionSize:                      i(50000000000),
					SpotMarketID:                     "bnb:usd",
					LiquidationMarketID:              "bnb:usd:30",
					KeeperRewardPercentage:           d("0.01"),
					CheckCollateralizationIndexCount: i(10),
					ConversionFactor:                 i(8),
				},
				{
					Denom:                            "busd",
					Type:                             "busd-a",
					LiquidationRatio:                 d("1.01"),
					DebtLimit:                        sdk.NewInt64Coin("usdx", 500000000000),
					StabilityFee:                     sdk.OneDec(), // %0 apr
					LiquidationPenalty:               d("0.05"),
					AuctionSize:                      i(10000000000),
					SpotMarketID:                     "busd:usd",
					LiquidationMarketID:              "busd:usd:30",
					KeeperRewardPercentage:           d("0.01"),
					CheckCollateralizationIndexCount: i(10),
					ConversionFactor:                 i(8),
				},
			},
			DebtParam: types3.DebtParam{
				Denom:            "usdx",
				ReferenceAsset:   "usd",
				ConversionFactor: i(6),
				DebtFloor:        i(10000000),
			},
		},
		StartingCdpID: types3.DefaultCdpStartingID,
		DebtDenom:     types3.DefaultDebtDenom,
		GovDenom:      types3.DefaultGovDenom,
		CDPs:          types3.CDPs{},
		PreviousAccumulationTimes: types3.GenesisAccumulationTimes{
			types3.NewGenesisAccumulationTime("btc-a", time.Time{}, sdk.OneDec()),
			types3.NewGenesisAccumulationTime("xrp-a", time.Time{}, sdk.OneDec()),
			types3.NewGenesisAccumulationTime("busd-a", time.Time{}, sdk.OneDec()),
			types3.NewGenesisAccumulationTime("bnb-a", time.Time{}, sdk.OneDec()),
		},
		TotalPrincipals: types3.GenesisTotalPrincipals{
			types3.NewGenesisTotalPrincipal("btc-a", sdk.ZeroInt()),
			types3.NewGenesisTotalPrincipal("xrp-a", sdk.ZeroInt()),
			types3.NewGenesisTotalPrincipal("busd-a", sdk.ZeroInt()),
			types3.NewGenesisTotalPrincipal("bnb-a", sdk.ZeroInt()),
		},
	}
	return app.GenesisState{types3.ModuleName: cdc.MustMarshalJSON(&cdpGenesis)}
}

func NewCDPGenStateHighDebtLimit(cdc codec.JSONCodec) app.GenesisState {
	cdpGenesis := types3.GenesisState{
		Params: types3.Params{
			GlobalDebtLimit:         sdk.NewInt64Coin("usdx", 100000000000000),
			SurplusAuctionThreshold: types3.DefaultSurplusThreshold,
			SurplusAuctionLot:       types3.DefaultSurplusLot,
			DebtAuctionThreshold:    types3.DefaultDebtThreshold,
			DebtAuctionLot:          types3.DefaultDebtLot,
			CollateralParams: types3.CollateralParams{
				{
					Denom:                            "xrp",
					Type:                             "xrp-a",
					LiquidationRatio:                 sdk.MustNewDecFromStr("2.0"),
					DebtLimit:                        sdk.NewInt64Coin("usdx", 50000000000000),
					StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"), // %5 apr
					LiquidationPenalty:               d("0.05"),
					AuctionSize:                      i(7000000000),
					SpotMarketID:                     "xrp:usd",
					LiquidationMarketID:              "xrp:usd",
					KeeperRewardPercentage:           d("0.01"),
					CheckCollateralizationIndexCount: i(10),
					ConversionFactor:                 i(6),
				},
				{
					Denom:                            "btc",
					Type:                             "btc-a",
					LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
					DebtLimit:                        sdk.NewInt64Coin("usdx", 50000000000000),
					StabilityFee:                     sdk.MustNewDecFromStr("1.000000000782997609"), // %2.5 apr
					LiquidationPenalty:               d("0.025"),
					AuctionSize:                      i(10000000),
					SpotMarketID:                     "btc:usd",
					LiquidationMarketID:              "btc:usd",
					KeeperRewardPercentage:           d("0.01"),
					CheckCollateralizationIndexCount: i(10),
					ConversionFactor:                 i(8),
				},
			},
			DebtParam: types3.DebtParam{
				Denom:            "usdx",
				ReferenceAsset:   "usd",
				ConversionFactor: i(6),
				DebtFloor:        i(10000000),
			},
		},
		StartingCdpID: types3.DefaultCdpStartingID,
		DebtDenom:     types3.DefaultDebtDenom,
		GovDenom:      types3.DefaultGovDenom,
		CDPs:          types3.CDPs{},
		PreviousAccumulationTimes: types3.GenesisAccumulationTimes{
			types3.NewGenesisAccumulationTime("btc-a", time.Time{}, sdk.OneDec()),
			types3.NewGenesisAccumulationTime("xrp-a", time.Time{}, sdk.OneDec()),
		},
		TotalPrincipals: types3.GenesisTotalPrincipals{
			types3.NewGenesisTotalPrincipal("btc-a", sdk.ZeroInt()),
			types3.NewGenesisTotalPrincipal("xrp-a", sdk.ZeroInt()),
		},
	}
	return app.GenesisState{types3.ModuleName: cdc.MustMarshalJSON(&cdpGenesis)}
}

func cdps() (cdps types3.CDPs) {
	_, addrs := app.GeneratePrivKeyAddressPairs(3)
	c1 := types3.NewCDP(uint64(1), addrs[0], sdk.NewCoin("xrp", sdk.NewInt(10000000)), "xrp-a", sdk.NewCoin("usdx", sdk.NewInt(8000000)), tmtime.Canonical(time.Now()), sdk.OneDec())
	c2 := types3.NewCDP(uint64(2), addrs[1], sdk.NewCoin("xrp", sdk.NewInt(100000000)), "xrp-a", sdk.NewCoin("usdx", sdk.NewInt(10000000)), tmtime.Canonical(time.Now()), sdk.OneDec())
	c3 := types3.NewCDP(uint64(3), addrs[1], sdk.NewCoin("btc", sdk.NewInt(1000000000)), "btc-a", sdk.NewCoin("usdx", sdk.NewInt(10000000)), tmtime.Canonical(time.Now()), sdk.OneDec())
	c4 := types3.NewCDP(uint64(4), addrs[2], sdk.NewCoin("xrp", sdk.NewInt(1000000000)), "xrp-a", sdk.NewCoin("usdx", sdk.NewInt(500000000)), tmtime.Canonical(time.Now()), sdk.OneDec())
	cdps = append(cdps, c1, c2, c3, c4)
	return
}
