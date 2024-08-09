package keeper_test

import (
	"strings"
	"time"

	sdkmath "cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/x/bank/testutil"

	"github.com/joltify-finance/joltify_lending/x/third_party/jolt"
	types3 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"

	"github.com/cometbft/cometbft/crypto"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/app"
)

func (suite *KeeperTestSuite) TestDeposit() {
	type args struct {
		depositor                 sdk.AccAddress
		amount                    sdk.Coins
		numberDeposits            int
		expectedAccountBalance    sdk.Coins
		expectedModAccountBalance sdk.Coins
		expectedDepositCoins      sdk.Coins
	}
	type errArgs struct {
		expectPass bool
		contains   string
	}
	type depositTest struct {
		name    string
		args    args
		errArgs errArgs
	}
	testCases := []depositTest{
		{
			"valid",
			args{
				depositor:                 sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				amount:                    sdk.NewCoins(sdk.NewCoin("bnb", sdkmath.NewInt(100))),
				numberDeposits:            1,
				expectedAccountBalance:    sdk.NewCoins(sdk.NewCoin("bnb", sdkmath.NewInt(900)), sdk.NewCoin("btcb", sdkmath.NewInt(1000))),
				expectedModAccountBalance: sdk.NewCoins(sdk.NewCoin("bnb", sdkmath.NewInt(100))),
				expectedDepositCoins:      sdk.NewCoins(sdk.NewCoin("bnb", sdkmath.NewInt(100))),
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"valid multi deposit",
			args{
				depositor:                 sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				amount:                    sdk.NewCoins(sdk.NewCoin("bnb", sdkmath.NewInt(100))),
				numberDeposits:            2,
				expectedAccountBalance:    sdk.NewCoins(sdk.NewCoin("bnb", sdkmath.NewInt(800)), sdk.NewCoin("btcb", sdkmath.NewInt(1000))),
				expectedModAccountBalance: sdk.NewCoins(sdk.NewCoin("bnb", sdkmath.NewInt(200))),
				expectedDepositCoins:      sdk.NewCoins(sdk.NewCoin("bnb", sdkmath.NewInt(200))),
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"invalid deposit denom",
			args{
				depositor:                 sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				amount:                    sdk.NewCoins(sdk.NewCoin("fake", sdkmath.NewInt(100))),
				numberDeposits:            1,
				expectedAccountBalance:    sdk.Coins{},
				expectedModAccountBalance: sdk.Coins{},
				expectedDepositCoins:      sdk.Coins{},
			},
			errArgs{
				expectPass: false,
				contains:   "invalid deposit denom",
			},
		},
		{
			"insufficient funds",
			args{
				depositor:                 sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				amount:                    sdk.NewCoins(sdk.NewCoin("bnb", sdkmath.NewInt(10000))),
				numberDeposits:            1,
				expectedAccountBalance:    sdk.Coins{},
				expectedModAccountBalance: sdk.Coins{},
				expectedDepositCoins:      sdk.Coins{},
			},
			errArgs{
				expectPass: false,
				contains:   "insufficient funds: the requested deposit amount",
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			// create new app with one funded account

			coins := sdk.NewCoins(
				sdk.NewCoin("bnb", sdkmath.NewInt(1000)),
				sdk.NewCoin("btcb", sdkmath.NewInt(1000)),
			)

			// Initialize test app and set context
			authGS := app.NewFundedGenStateWithCoins(
				suite.app.AppCodec(),
				[]sdk.Coins{
					coins,
				},
				[]sdk.AccAddress{tc.args.depositor},
			)
			loanToValue, _ := sdkmath.LegacyNewDecFromStr("0.6")
			hardGS := types3.NewGenesisState(types3.NewParams(
				types3.MoneyMarkets{
					types3.NewMoneyMarket("usdx", types3.NewBorrowLimit(false, sdkmath.LegacyNewDec(1000000000000000), loanToValue), "usdx:usd", sdkmath.NewInt(1000000), types3.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyMustNewDecFromStr("2"), sdkmath.LegacyMustNewDecFromStr("0.8"), sdkmath.LegacyMustNewDecFromStr("10")), sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyZeroDec()),
					types3.NewMoneyMarket("ujolt", types3.NewBorrowLimit(false, sdkmath.LegacyNewDec(1000000000000000), loanToValue), "joltify:usd", sdkmath.NewInt(1000000), types3.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyMustNewDecFromStr("2"), sdkmath.LegacyMustNewDecFromStr("0.8"), sdkmath.LegacyMustNewDecFromStr("10")), sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyZeroDec()),
					types3.NewMoneyMarket("bnb", types3.NewBorrowLimit(false, sdkmath.LegacyNewDec(1000000000000000), loanToValue), "bnb:usd", sdkmath.NewInt(1000000), types3.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyMustNewDecFromStr("2"), sdkmath.LegacyMustNewDecFromStr("0.8"), sdkmath.LegacyMustNewDecFromStr("10")), sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyZeroDec()),
					types3.NewMoneyMarket("btcb", types3.NewBorrowLimit(false, sdkmath.LegacyNewDec(1000000000000000), loanToValue), "btcb:usd", sdkmath.NewInt(1000000), types3.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyMustNewDecFromStr("2"), sdkmath.LegacyMustNewDecFromStr("0.8"), sdkmath.LegacyMustNewDecFromStr("10")), sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyZeroDec()),
				},
				sdkmath.LegacyNewDec(10),
			), types3.DefaultAccumulationTimes, types3.DefaultDeposits, types3.DefaultBorrows,
				types3.DefaultTotalSupplied, types3.DefaultTotalBorrowed, types3.DefaultTotalReserves,
			)

			// Pricefeed module genesis state
			pricefeedGS := types2.GenesisState{
				Params: types2.Params{
					Markets: []types2.Market{
						{MarketID: "usdx:usd", BaseAsset: "usdx", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "joltify:usd", BaseAsset: "joltify", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "btcb:usd", BaseAsset: "btcb", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "bnb:usd", BaseAsset: "bnb", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
					},
				},
				PostedPrices: []types2.PostedPrice{
					{
						MarketID:      "usdx:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdkmath.LegacyMustNewDecFromStr("1.00"),
						Expiry:        time.Now().Add(1 * time.Hour),
					},
					{
						MarketID:      "joltify:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdkmath.LegacyMustNewDecFromStr("2.00"),
						Expiry:        time.Now().Add(1 * time.Hour),
					},
					{
						MarketID:      "btcb:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdkmath.LegacyMustNewDecFromStr("100.00"),
						Expiry:        time.Now().Add(1 * time.Hour),
					},
					{
						MarketID:      "bnb:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdkmath.LegacyMustNewDecFromStr("10.00"),
						Expiry:        time.Now().Add(1 * time.Hour),
					},
				},
			}

			mapp := suite.app.InitializeFromGenesisStates(suite.T(), time.Now(), nil, nil, authGS,
				app.GenesisState{types2.ModuleName: suite.app.AppCodec().MustMarshalJSON(&pricefeedGS)},
				app.GenesisState{types3.ModuleName: suite.app.AppCodec().MustMarshalJSON(&hardGS)},
			)

			suite.app = mapp
			suite.app.App = mapp.App
			suite.ctx = mapp.Ctx
			suite.app.Ctx = mapp.Ctx
			suite.keeper = mapp.GetJoltKeeper()

			// Run BeginBlocker once to transition MoneyMarkets
			jolt.BeginBlocker(suite.ctx, suite.keeper)

			var err error
			err = testutil.FundAccount(suite.ctx, suite.app.GetBankKeeper(), tc.args.depositor, coins)
			suite.Require().NoError(err)

			// run the test
			for i := 0; i < tc.args.numberDeposits; i++ {
				err = suite.keeper.Deposit(suite.ctx, tc.args.depositor, tc.args.amount)
			}

			// verify results
			if tc.errArgs.expectPass {
				suite.Require().NoError(err)
				acc := suite.getAccount(tc.args.depositor)
				suite.Require().Equal(tc.args.expectedAccountBalance, suite.getAccountCoins(acc))
				mAcc := suite.getModuleAccount(types3.ModuleAccountName)
				suite.Require().Equal(tc.args.expectedModAccountBalance, suite.getAccountCoins(mAcc))
				dep, f := suite.keeper.GetDeposit(suite.ctx, tc.args.depositor)
				suite.Require().True(f)
				suite.Require().Equal(tc.args.expectedDepositCoins, dep.Amount)
			} else {
				suite.Require().Error(err)
				suite.Require().True(strings.Contains(err.Error(), tc.errArgs.contains))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestDecrementSuppliedCoins() {
	type args struct {
		suppliedInitial       sdk.Coins
		decrementCoins        sdk.Coins
		expectedSuppliedFinal sdk.Coins
	}
	type errArgs struct {
		expectPass bool
		contains   string
	}
	type decrementTest struct {
		name    string
		args    args
		errArgs errArgs
	}
	testCases := []decrementTest{
		{
			"valid",
			args{
				suppliedInitial:       cs(c("bnb", 10000000000000), c("busd", 3000000000000), c("xrpb", 2500000000000)),
				decrementCoins:        cs(c("bnb", 5000000000000)),
				expectedSuppliedFinal: cs(c("bnb", 5000000000000), c("busd", 3000000000000), c("xrpb", 2500000000000)),
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"valid-negative",
			args{
				suppliedInitial:       cs(c("bnb", 10000000000000), c("busd", 3000000000000), c("xrpb", 2500000000000)),
				decrementCoins:        cs(c("bnb", 10000000000001)),
				expectedSuppliedFinal: cs(c("busd", 3000000000000), c("xrpb", 2500000000000)),
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"valid-multiple negative",
			args{
				suppliedInitial:       cs(c("bnb", 10000000000000), c("busd", 3000000000000), c("xrpb", 2500000000000)),
				decrementCoins:        cs(c("bnb", 10000000000001), c("busd", 5000000000000)),
				expectedSuppliedFinal: cs(c("xrpb", 2500000000000)),
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"valid-absent coin denom",
			args{
				suppliedInitial:       cs(c("bnb", 10000000000000), c("xrpb", 2500000000000)),
				decrementCoins:        cs(c("busd", 5)),
				expectedSuppliedFinal: cs(c("bnb", 10000000000000), c("xrpb", 2500000000000)),
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			loanToValue, _ := sdkmath.LegacyNewDecFromStr("0.6")
			depositor := sdk.AccAddress(crypto.AddressHash([]byte("test")))
			authGS := app.NewFundedGenStateWithCoins(
				suite.app.AppCodec(),
				[]sdk.Coins{tc.args.suppliedInitial},
				[]sdk.AccAddress{depositor},
			)
			hardGS := types3.NewGenesisState(types3.NewParams(
				types3.MoneyMarkets{
					types3.NewMoneyMarket("bnb", types3.NewBorrowLimit(false, sdkmath.LegacyNewDec(1000000000000000), loanToValue), "bnb:usd", sdkmath.NewInt(100000000), types3.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyMustNewDecFromStr("2"), sdkmath.LegacyMustNewDecFromStr("0.8"), sdkmath.LegacyMustNewDecFromStr("10")), sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyZeroDec()),
					types3.NewMoneyMarket("busd", types3.NewBorrowLimit(false, sdkmath.LegacyNewDec(1000000000000000), loanToValue), "busd:usd", sdkmath.NewInt(100000000), types3.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyMustNewDecFromStr("2"), sdkmath.LegacyMustNewDecFromStr("0.8"), sdkmath.LegacyMustNewDecFromStr("10")), sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyZeroDec()),
					types3.NewMoneyMarket("xrpb", types3.NewBorrowLimit(false, sdkmath.LegacyNewDec(1000000000000000), loanToValue), "xrpb:usd", sdkmath.NewInt(100000000), types3.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyMustNewDecFromStr("2"), sdkmath.LegacyMustNewDecFromStr("0.8"), sdkmath.LegacyMustNewDecFromStr("10")), sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyZeroDec()),
				},
				sdkmath.LegacyMustNewDecFromStr("10"),
			), types3.DefaultAccumulationTimes, types3.DefaultDeposits, types3.DefaultBorrows,
				types3.DefaultTotalSupplied, types3.DefaultTotalBorrowed, types3.DefaultTotalReserves,
			)
			// Pricefeed module genesis state
			pricefeedGS := types2.GenesisState{
				Params: types2.Params{
					Markets: []types2.Market{
						{MarketID: "xrpb:usd", BaseAsset: "joltify", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "busd:usd", BaseAsset: "btcb", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "bnb:usd", BaseAsset: "bnb", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
					},
				},
				PostedPrices: []types2.PostedPrice{
					{
						MarketID:      "busd:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdkmath.LegacyMustNewDecFromStr("1.00"),
						Expiry:        time.Now().Add(1 * time.Hour),
					},
					{
						MarketID:      "xrpb:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdkmath.LegacyMustNewDecFromStr("2.00"),
						Expiry:        time.Now().Add(1 * time.Hour),
					},
					{
						MarketID:      "bnb:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdkmath.LegacyMustNewDecFromStr("200.00"),
						Expiry:        time.Now().Add(1 * time.Hour),
					},
				},
			}
			mapp := suite.app.InitializeFromGenesisStates(suite.T(), time.Now(), nil, nil, authGS,
				app.GenesisState{types2.ModuleName: suite.app.AppCodec().MustMarshalJSON(&pricefeedGS)},
				app.GenesisState{types3.ModuleName: suite.app.AppCodec().MustMarshalJSON(&hardGS)},
			)

			suite.app = mapp
			suite.app.App = mapp.App
			suite.ctx = mapp.Ctx
			suite.app.Ctx = mapp.Ctx
			suite.keeper = mapp.GetJoltKeeper()

			keeper := mapp.GetJoltKeeper()
			suite.keeper = keeper

			// Run BeginBlocker once to transition MoneyMarkets
			jolt.BeginBlocker(suite.ctx, suite.keeper)

			err := testutil.FundAccount(suite.ctx, suite.app.GetBankKeeper(), depositor, tc.args.suppliedInitial)
			suite.Require().NoError(err)

			err = suite.keeper.Deposit(suite.ctx, depositor, tc.args.suppliedInitial)
			suite.Require().NoError(err)
			err = suite.keeper.DecrementSuppliedCoins(suite.ctx, tc.args.decrementCoins)
			suite.Require().NoError(err)
			totalSuppliedActual, found := suite.keeper.GetSuppliedCoins(suite.ctx)
			suite.Require().True(found)
			suite.Require().Equal(totalSuppliedActual, tc.args.expectedSuppliedFinal)
		})
	}
}

func c(denom string, amount int64) sdk.Coin { return sdk.NewInt64Coin(denom, amount) }
func cs(coins ...sdk.Coin) sdk.Coins        { return sdk.NewCoins(coins...) }
