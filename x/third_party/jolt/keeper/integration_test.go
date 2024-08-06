package keeper_test

import (
	"time"

	sdkmath "cosmossdk.io/math"
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
						MaximumLimit: sdkmath.LegacyMustNewDecFromStr("100000000000"),
						LoanToValue:  sdkmath.LegacyMustNewDecFromStr("1"),
					},
					SpotMarketID:     "usdx:usd",
					ConversionFactor: sdkmath.NewInt(UsdxCf),
					InterestRateModel: types3.InterestRateModel{
						BaseRateAPY:    sdkmath.LegacyMustNewDecFromStr("0.05"),
						BaseMultiplier: sdkmath.LegacyMustNewDecFromStr("2"),
						Kink:           sdkmath.LegacyMustNewDecFromStr("0.8"),
						JumpMultiplier: sdkmath.LegacyMustNewDecFromStr("10"),
					},
					ReserveFactor:          sdkmath.LegacyMustNewDecFromStr("0.05"),
					KeeperRewardPercentage: sdkmath.LegacyZeroDec(),
				},
				types3.MoneyMarket{
					Denom: "bnb",
					BorrowLimit: types3.BorrowLimit{
						HasMaxLimit:  true,
						MaximumLimit: sdkmath.LegacyMustNewDecFromStr("3000000000000"),
						LoanToValue:  sdkmath.LegacyMustNewDecFromStr("0.5"),
					},
					SpotMarketID:     "bnb:usd",
					ConversionFactor: sdkmath.NewInt(UsdxCf),
					InterestRateModel: types3.InterestRateModel{
						BaseRateAPY:    sdkmath.LegacyMustNewDecFromStr("0"),
						BaseMultiplier: sdkmath.LegacyMustNewDecFromStr("0.05"),
						Kink:           sdkmath.LegacyMustNewDecFromStr("0.8"),
						JumpMultiplier: sdkmath.LegacyMustNewDecFromStr("5.0"),
					},
					ReserveFactor:          sdkmath.LegacyMustNewDecFromStr("0.025"),
					KeeperRewardPercentage: sdkmath.LegacyMustNewDecFromStr("0.02"),
				},
				types3.MoneyMarket{
					Denom: "busd",
					BorrowLimit: types3.BorrowLimit{
						HasMaxLimit:  true,
						MaximumLimit: sdkmath.LegacyMustNewDecFromStr("1000000000000000"),
						LoanToValue:  sdkmath.LegacyMustNewDecFromStr("0.5"),
					},
					SpotMarketID:     "busd:usd",
					ConversionFactor: sdkmath.LegacyMustNewDecFromStr("100000000").RoundInt(),
					InterestRateModel: types3.InterestRateModel{
						BaseRateAPY:    sdkmath.LegacyMustNewDecFromStr("0"),
						BaseMultiplier: sdkmath.LegacyMustNewDecFromStr("0.5"),
						Kink:           sdkmath.LegacyMustNewDecFromStr("0.8"),
						JumpMultiplier: sdkmath.LegacyMustNewDecFromStr("5"),
					},
					ReserveFactor:          sdkmath.LegacyMustNewDecFromStr("0.025"),
					KeeperRewardPercentage: sdkmath.LegacyMustNewDecFromStr("0.02"),
				},
			},
			sdkmath.LegacyMustNewDecFromStr("10"),
		),
		PreviousAccumulationTimes: types3.GenesisAccumulationTimes{
			types3.NewGenesisAccumulationTime(
				"usdx",
				time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC),
				sdkmath.LegacyOneDec(),
				sdkmath.LegacyOneDec(),
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
				Price:         sdkmath.LegacyOneDec(),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketID:      "bnb:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdkmath.LegacyMustNewDecFromStr("618.13"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketID:      "busd:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdkmath.LegacyOneDec(),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
		},
	}
	return app.GenesisState{types2.ModuleName: cdc.MustMarshalJSON(&pfGenesis)}
}
