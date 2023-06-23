package keeper_test

import (
	"time"

	types3 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/app"
)

func NewJoltGenState(cdc codec.JSONCodec) app.GenesisState {
	hardGenesis := types3.GenesisState{
		Params: types3.NewParams(
			types3.MoneyMarkets{
				types3.MoneyMarket{
					Denom: "usdx",
					BorrowLimit: types3.BorrowLimit{
						HasMaxLimit:  true,
						MaximumLimit: sdk.MustNewDecFromStr("100000000000"),
						LoanToValue:  sdk.MustNewDecFromStr("1"),
					},
					SpotMarketID:     "usdx:usd",
					ConversionFactor: sdk.NewInt(UsdxCf),
					InterestRateModel: types3.InterestRateModel{
						BaseRateAPY:    sdk.MustNewDecFromStr("0.05"),
						BaseMultiplier: sdk.MustNewDecFromStr("2"),
						Kink:           sdk.MustNewDecFromStr("0.8"),
						JumpMultiplier: sdk.MustNewDecFromStr("10"),
					},
					ReserveFactor:          sdk.MustNewDecFromStr("0.05"),
					KeeperRewardPercentage: sdk.ZeroDec(),
				},
				types3.MoneyMarket{
					Denom: "bnb",
					BorrowLimit: types3.BorrowLimit{
						HasMaxLimit:  true,
						MaximumLimit: sdk.MustNewDecFromStr("3000000000000"),
						LoanToValue:  sdk.MustNewDecFromStr("0.5"),
					},
					SpotMarketID:     "bnb:usd",
					ConversionFactor: sdk.NewInt(UsdxCf),
					InterestRateModel: types3.InterestRateModel{
						BaseRateAPY:    sdk.MustNewDecFromStr("0"),
						BaseMultiplier: sdk.MustNewDecFromStr("0.05"),
						Kink:           sdk.MustNewDecFromStr("0.8"),
						JumpMultiplier: sdk.MustNewDecFromStr("5.0"),
					},
					ReserveFactor:          sdk.MustNewDecFromStr("0.025"),
					KeeperRewardPercentage: sdk.MustNewDecFromStr("0.02"),
				},
				types3.MoneyMarket{
					Denom: "busd",
					BorrowLimit: types3.BorrowLimit{
						HasMaxLimit:  true,
						MaximumLimit: sdk.MustNewDecFromStr("1000000000000000"),
						LoanToValue:  sdk.MustNewDecFromStr("0.5"),
					},
					SpotMarketID:     "busd:usd",
					ConversionFactor: sdk.MustNewDecFromStr("100000000").RoundInt(),
					InterestRateModel: types3.InterestRateModel{
						BaseRateAPY:    sdk.MustNewDecFromStr("0"),
						BaseMultiplier: sdk.MustNewDecFromStr("0.5"),
						Kink:           sdk.MustNewDecFromStr("0.8"),
						JumpMultiplier: sdk.MustNewDecFromStr("5"),
					},
					ReserveFactor:          sdk.MustNewDecFromStr("0.025"),
					KeeperRewardPercentage: sdk.MustNewDecFromStr("0.02"),
				},
			},
			sdk.MustNewDecFromStr("10"),
		),
		PreviousAccumulationTimes: types3.GenesisAccumulationTimes{
			types3.NewGenesisAccumulationTime(
				"usdx",
				time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC),
				sdk.OneDec(),
				sdk.OneDec(),
			),
		},
		Deposits:      types3.DefaultDeposits,
		Borrows:       types3.DefaultBorrows,
		TotalSupplied: sdk.NewCoins(),
		TotalBorrowed: sdk.NewCoins(),
		TotalReserves: sdk.NewCoins(),
	}
	return app.GenesisState{types3.ModuleName: cdc.MustMarshalJSON(&hardGenesis)}
}

func NewPricefeedGenStateMulti(cdc codec.JSONCodec) app.GenesisState {
	pfGenesis := types2.GenesisState{
		Params: types2.Params{
			Markets: []types2.Market{
				{MarketID: "usdx:usd", BaseAsset: "usdx", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "xrp:usd", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "bnb:usd", BaseAsset: "bnb", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "busd:usd", BaseAsset: "busd", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
			},
		},
		PostedPrices: []types2.PostedPrice{
			{
				MarketID:      "usdx:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.OneDec(),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketID:      "bnb:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("618.13"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketID:      "busd:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.OneDec(),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
		},
	}
	return app.GenesisState{types2.ModuleName: cdc.MustMarshalJSON(&pfGenesis)}
}
