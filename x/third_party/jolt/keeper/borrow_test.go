package keeper_test

import (
	"strings"
	"time"

	tmlog "github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/x/bank/testutil"

	"github.com/joltify-finance/joltify_lending/x/third_party/jolt"
	types3 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"

	"github.com/cometbft/cometbft/crypto"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtime "github.com/cometbft/cometbft/types/time"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/app"
)

const (
	UsdxCf = 1000000
	JoltCf = 1000000
	BtcbCf = 100000000
	BnbCf  = 100000000
	BusdCf = 100000000
)

func (suite *KeeperTestSuite) TestBorrow() {
	type args struct {
		usdxBorrowLimit           sdkmath.LegacyDec
		priceJolt                 sdkmath.LegacyDec
		loanToValueJolt           sdkmath.LegacyDec
		priceBTCB                 sdkmath.LegacyDec
		loanToValueBTCB           sdkmath.LegacyDec
		priceBNB                  sdkmath.LegacyDec
		loanToValueBNB            sdkmath.LegacyDec
		borrower                  sdk.AccAddress
		depositCoins              []sdk.Coin
		previousBorrowCoins       sdk.Coins
		borrowCoins               sdk.Coins
		expectedAccountBalance    sdk.Coins
		expectedModAccountBalance sdk.Coins
	}
	type errArgs struct {
		expectPass bool
		contains   string
	}
	type borrowTest struct {
		name    string
		args    args
		errArgs errArgs
	}
	testCases := []borrowTest{
		{
			"valid",
			args{
				usdxBorrowLimit:           sdkmath.LegacyMustNewDecFromStr("100000000000"),
				priceJolt:                 sdkmath.LegacyMustNewDecFromStr("5.00"),
				loanToValueJolt:           sdkmath.LegacyMustNewDecFromStr("0.6"),
				priceBTCB:                 sdkmath.LegacyMustNewDecFromStr("0.00"),
				loanToValueBTCB:           sdkmath.LegacyMustNewDecFromStr("0.01"),
				priceBNB:                  sdkmath.LegacyMustNewDecFromStr("0.00"),
				loanToValueBNB:            sdkmath.LegacyMustNewDecFromStr("0.01"),
				borrower:                  sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				depositCoins:              []sdk.Coin{sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))},
				previousBorrowCoins:       sdk.NewCoins(),
				borrowCoins:               sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(20*JoltCf))),
				expectedAccountBalance:    sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(20*JoltCf)), sdk.NewCoin("btcb", sdkmath.NewInt(100*BtcbCf)), sdk.NewCoin("bnb", sdkmath.NewInt(100*BnbCf)), sdk.NewCoin("xyz", sdkmath.NewInt(1))),
				expectedModAccountBalance: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1080*JoltCf)), sdk.NewCoin("usdx", sdkmath.NewInt(200*UsdxCf)), sdk.NewCoin("busd", sdkmath.NewInt(100*BusdCf))),
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"invalid: loan-to-value limited",
			args{
				usdxBorrowLimit:           sdkmath.LegacyMustNewDecFromStr("100000000000"),
				priceJolt:                 sdkmath.LegacyMustNewDecFromStr("5.00"),
				loanToValueJolt:           sdkmath.LegacyMustNewDecFromStr("0.6"),
				priceBTCB:                 sdkmath.LegacyMustNewDecFromStr("0.00"),
				loanToValueBTCB:           sdkmath.LegacyMustNewDecFromStr("0.01"),
				priceBNB:                  sdkmath.LegacyMustNewDecFromStr("0.00"),
				loanToValueBNB:            sdkmath.LegacyMustNewDecFromStr("0.01"),
				borrower:                  sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				depositCoins:              []sdk.Coin{sdk.NewCoin("ujolt", sdkmath.NewInt(20*JoltCf))},  // 20 JOLTx $5.00 price = $100
				borrowCoins:               sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(61*UsdxCf))), // 61 USDX x $1 price = $61
				expectedAccountBalance:    sdk.NewCoins(),
				expectedModAccountBalance: sdk.NewCoins(),
			},
			errArgs{
				expectPass: false,
				contains:   "exceeds the allowable amount as determined by the collateralization ratio",
			},
		},
		{
			"valid: multiple deposits",
			args{
				usdxBorrowLimit:           sdkmath.LegacyMustNewDecFromStr("100000000000"),
				priceJolt:                 sdkmath.LegacyMustNewDecFromStr("2.00"),
				loanToValueJolt:           sdkmath.LegacyMustNewDecFromStr("0.80"),
				priceBTCB:                 sdkmath.LegacyMustNewDecFromStr("10000.00"),
				loanToValueBTCB:           sdkmath.LegacyMustNewDecFromStr("0.10"),
				priceBNB:                  sdkmath.LegacyMustNewDecFromStr("0.00"),
				loanToValueBNB:            sdkmath.LegacyMustNewDecFromStr("0.01"),
				borrower:                  sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				depositCoins:              sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(50*JoltCf)), sdk.NewCoin("btcb", sdkmath.NewInt(0.1*BtcbCf))),
				borrowCoins:               sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(180*UsdxCf))),
				expectedAccountBalance:    sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(50*JoltCf)), sdk.NewCoin("btcb", sdkmath.NewInt(99.9*BtcbCf)), sdk.NewCoin("usdx", sdkmath.NewInt(180*UsdxCf)), sdk.NewCoin("bnb", sdkmath.NewInt(100*BnbCf)), sdk.NewCoin("xyz", sdkmath.NewInt(1))),
				expectedModAccountBalance: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1050*JoltCf)), sdk.NewCoin("usdx", sdkmath.NewInt(20*UsdxCf)), sdk.NewCoin("btcb", sdkmath.NewInt(0.1*BtcbCf)), sdk.NewCoin("busd", sdkmath.NewInt(100*BusdCf))),
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"invalid: multiple deposits",
			args{
				usdxBorrowLimit:           sdkmath.LegacyMustNewDecFromStr("100000000000"),
				priceJolt:                 sdkmath.LegacyMustNewDecFromStr("2.00"),
				loanToValueJolt:           sdkmath.LegacyMustNewDecFromStr("0.80"),
				priceBTCB:                 sdkmath.LegacyMustNewDecFromStr("10000.00"),
				loanToValueBTCB:           sdkmath.LegacyMustNewDecFromStr("0.10"),
				priceBNB:                  sdkmath.LegacyMustNewDecFromStr("0.00"),
				loanToValueBNB:            sdkmath.LegacyMustNewDecFromStr("0.01"),
				borrower:                  sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				depositCoins:              sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(50*JoltCf)), sdk.NewCoin("btcb", sdkmath.NewInt(0.1*BtcbCf))),
				borrowCoins:               sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(181*UsdxCf))),
				expectedAccountBalance:    sdk.NewCoins(),
				expectedModAccountBalance: sdk.NewCoins(),
			},
			errArgs{
				expectPass: false,
				contains:   "exceeds the allowable amount as determined by the collateralization ratio",
			},
		},
		{
			"valid: multiple previous borrows",
			args{
				usdxBorrowLimit:           sdkmath.LegacyMustNewDecFromStr("100000000000"),
				priceJolt:                 sdkmath.LegacyMustNewDecFromStr("2.00"),
				loanToValueJolt:           sdkmath.LegacyMustNewDecFromStr("0.8"),
				priceBTCB:                 sdkmath.LegacyMustNewDecFromStr("0.00"),
				loanToValueBTCB:           sdkmath.LegacyMustNewDecFromStr("0.01"),
				priceBNB:                  sdkmath.LegacyMustNewDecFromStr("5.00"),
				loanToValueBNB:            sdkmath.LegacyMustNewDecFromStr("0.8"),
				borrower:                  sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				depositCoins:              sdk.NewCoins(sdk.NewCoin("bnb", sdkmath.NewInt(30*BnbCf)), sdk.NewCoin("ujolt", sdkmath.NewInt(50*JoltCf))), // (50 JOLT x $2.00 price = $100) + (30 BNB x $5.00 price = $150) = $250
				previousBorrowCoins:       sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(99*UsdxCf)), sdk.NewCoin("busd", sdkmath.NewInt(100*BusdCf))),
				borrowCoins:               sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(1*UsdxCf))),
				expectedAccountBalance:    sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(50*JoltCf)), sdk.NewCoin("btcb", sdkmath.NewInt(100*BtcbCf)), sdk.NewCoin("usdx", sdkmath.NewInt(100*UsdxCf)), sdk.NewCoin("busd", sdkmath.NewInt(100*BusdCf)), sdk.NewCoin("bnb", sdkmath.NewInt(70*BnbCf)), sdk.NewCoin("xyz", sdkmath.NewInt(1))),
				expectedModAccountBalance: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1050*JoltCf)), sdk.NewCoin("bnb", sdkmath.NewInt(30*BusdCf)), sdk.NewCoin("usdx", sdkmath.NewInt(100*UsdxCf))),
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"invalid: over loan-to-value with multiple previous borrows",
			args{
				usdxBorrowLimit:           sdkmath.LegacyMustNewDecFromStr("100000000000"),
				priceJolt:                 sdkmath.LegacyMustNewDecFromStr("2.00"),
				loanToValueJolt:           sdkmath.LegacyMustNewDecFromStr("0.8"),
				priceBTCB:                 sdkmath.LegacyMustNewDecFromStr("0.00"),
				loanToValueBTCB:           sdkmath.LegacyMustNewDecFromStr("0.01"),
				priceBNB:                  sdkmath.LegacyMustNewDecFromStr("5.00"),
				loanToValueBNB:            sdkmath.LegacyMustNewDecFromStr("0.8"),
				borrower:                  sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				depositCoins:              sdk.NewCoins(sdk.NewCoin("bnb", sdkmath.NewInt(30*BnbCf)), sdk.NewCoin("ujolt", sdkmath.NewInt(50*JoltCf))), // (50 KAVA x $2.00 price = $100) + (30 BNB x $5.00 price = $150) = $250
				previousBorrowCoins:       sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(100*UsdxCf)), sdk.NewCoin("busd", sdkmath.NewInt(100*BusdCf))),
				borrowCoins:               sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(1*UsdxCf))),
				expectedAccountBalance:    sdk.NewCoins(),
				expectedModAccountBalance: sdk.NewCoins(),
			},
			errArgs{
				expectPass: false,
				contains:   "exceeds the allowable amount as determined by the collateralization ratio",
			},
		},
		{
			"invalid: no price for asset",
			args{
				usdxBorrowLimit:           sdkmath.LegacyMustNewDecFromStr("100000000000"),
				priceJolt:                 sdkmath.LegacyMustNewDecFromStr("5.00"),
				loanToValueJolt:           sdkmath.LegacyMustNewDecFromStr("0.6"),
				priceBTCB:                 sdkmath.LegacyMustNewDecFromStr("0.00"),
				loanToValueBTCB:           sdkmath.LegacyMustNewDecFromStr("0.01"),
				priceBNB:                  sdkmath.LegacyMustNewDecFromStr("0.00"),
				loanToValueBNB:            sdkmath.LegacyMustNewDecFromStr("0.01"),
				borrower:                  sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				depositCoins:              sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				previousBorrowCoins:       sdk.NewCoins(),
				borrowCoins:               sdk.NewCoins(sdk.NewCoin("xyz", sdkmath.NewInt(1))),
				expectedAccountBalance:    sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(20*JoltCf)), sdk.NewCoin("btcb", sdkmath.NewInt(100*BtcbCf)), sdk.NewCoin("bnb", sdkmath.NewInt(100*BnbCf)), sdk.NewCoin("xyz", sdkmath.NewInt(1))),
				expectedModAccountBalance: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1080*JoltCf)), sdk.NewCoin("usdx", sdkmath.NewInt(200*UsdxCf)), sdk.NewCoin("busd", sdkmath.NewInt(100*BusdCf))),
			},
			errArgs{
				expectPass: false,
				contains:   "no price found for market",
			},
		},
		{
			"invalid: borrow exceed module account balance",
			args{
				usdxBorrowLimit:           sdkmath.LegacyMustNewDecFromStr("100000000000"),
				priceJolt:                 sdkmath.LegacyMustNewDecFromStr("2.00"),
				loanToValueJolt:           sdkmath.LegacyMustNewDecFromStr("0.8"),
				priceBTCB:                 sdkmath.LegacyMustNewDecFromStr("0.00"),
				loanToValueBTCB:           sdkmath.LegacyMustNewDecFromStr("0.01"),
				priceBNB:                  sdkmath.LegacyMustNewDecFromStr("0.00"),
				loanToValueBNB:            sdkmath.LegacyMustNewDecFromStr("0.01"),
				borrower:                  sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				depositCoins:              sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				previousBorrowCoins:       sdk.NewCoins(),
				borrowCoins:               sdk.NewCoins(sdk.NewCoin("busd", sdkmath.NewInt(101*BusdCf))),
				expectedAccountBalance:    sdk.NewCoins(),
				expectedModAccountBalance: sdk.NewCoins(),
			},
			errArgs{
				expectPass: false,
				contains:   "exceeds borrowable module account balance",
			},
		},
		{
			"invalid: over global asset borrow limit",
			args{
				usdxBorrowLimit:           sdkmath.LegacyMustNewDecFromStr("20000000"),
				priceJolt:                 sdkmath.LegacyMustNewDecFromStr("2.00"),
				loanToValueJolt:           sdkmath.LegacyMustNewDecFromStr("0.8"),
				priceBTCB:                 sdkmath.LegacyMustNewDecFromStr("0.00"),
				loanToValueBTCB:           sdkmath.LegacyMustNewDecFromStr("0.01"),
				priceBNB:                  sdkmath.LegacyMustNewDecFromStr("0.00"),
				loanToValueBNB:            sdkmath.LegacyMustNewDecFromStr("0.01"),
				borrower:                  sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				depositCoins:              sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(50*JoltCf))),
				previousBorrowCoins:       sdk.NewCoins(),
				borrowCoins:               sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(25*UsdxCf))),
				expectedAccountBalance:    sdk.NewCoins(),
				expectedModAccountBalance: sdk.NewCoins(),
			},
			errArgs{
				expectPass: false,
				contains:   "fails global asset borrow limit validation",
			},
		},
		{
			"invalid: borrowing an individual coin type results in a borrow that's under the minimum USD borrow limit",
			args{
				usdxBorrowLimit:           sdkmath.LegacyMustNewDecFromStr("20000000"),
				priceJolt:                 sdkmath.LegacyMustNewDecFromStr("2.00"),
				loanToValueJolt:           sdkmath.LegacyMustNewDecFromStr("0.8"),
				priceBTCB:                 sdkmath.LegacyMustNewDecFromStr("0.00"),
				loanToValueBTCB:           sdkmath.LegacyMustNewDecFromStr("0.01"),
				priceBNB:                  sdkmath.LegacyMustNewDecFromStr("0.00"),
				loanToValueBNB:            sdkmath.LegacyMustNewDecFromStr("0.01"),
				borrower:                  sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				depositCoins:              sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(50*JoltCf))),
				previousBorrowCoins:       sdk.NewCoins(),
				borrowCoins:               sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(5*UsdxCf))),
				expectedAccountBalance:    sdk.NewCoins(),
				expectedModAccountBalance: sdk.NewCoins(),
			},
			errArgs{
				expectPass: false,
				contains:   "below the minimum borrow limit",
			},
		},
		{
			"invalid: borrowing multiple coins results in a borrow that's under the minimum USD borrow limit",
			args{
				usdxBorrowLimit:           sdkmath.LegacyMustNewDecFromStr("20000000"),
				priceJolt:                 sdkmath.LegacyMustNewDecFromStr("2.00"),
				loanToValueJolt:           sdkmath.LegacyMustNewDecFromStr("0.8"),
				priceBTCB:                 sdkmath.LegacyMustNewDecFromStr("0.00"),
				loanToValueBTCB:           sdkmath.LegacyMustNewDecFromStr("0.01"),
				priceBNB:                  sdkmath.LegacyMustNewDecFromStr("0.00"),
				loanToValueBNB:            sdkmath.LegacyMustNewDecFromStr("0.01"),
				borrower:                  sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				depositCoins:              sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(50*JoltCf))),
				previousBorrowCoins:       sdk.NewCoins(),
				borrowCoins:               sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(5*UsdxCf)), sdk.NewCoin("ujolt", sdkmath.NewInt(2*UsdxCf))),
				expectedAccountBalance:    sdk.NewCoins(),
				expectedModAccountBalance: sdk.NewCoins(),
			},
			errArgs{
				expectPass: false,
				contains:   "below the minimum borrow limit",
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			// Initialize test app and set context
			tApp := app.NewTestApp(tmlog.TestingLogger(), suite.T().TempDir())
			ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})

			coins := sdk.NewCoins(
				sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf)),
				sdk.NewCoin("btcb", sdkmath.NewInt(100*BtcbCf)),
				sdk.NewCoin("bnb", sdkmath.NewInt(100*BnbCf)),
				sdk.NewCoin("xyz", sdkmath.NewInt(1)),
			)

			// Auth module genesis state
			authGS := app.NewFundedGenStateWithCoins(
				tApp.AppCodec(),
				[]sdk.Coins{
					coins,
				},
				[]sdk.AccAddress{tc.args.borrower},
			)

			// jolt module genesis state
			hardGS := types3.NewGenesisState(types3.NewParams(
				types3.MoneyMarkets{
					types3.NewMoneyMarket("usdx", types3.NewBorrowLimit(true, tc.args.usdxBorrowLimit, sdkmath.LegacyMustNewDecFromStr("1")), "usdx:usd", sdkmath.NewInt(UsdxCf), types3.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyMustNewDecFromStr("2"), sdkmath.LegacyMustNewDecFromStr("0.8"), sdkmath.LegacyMustNewDecFromStr("10")), sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyZeroDec()),
					types3.NewMoneyMarket("busd", types3.NewBorrowLimit(false, sdk.NewDec(100000000*BusdCf), sdkmath.LegacyMustNewDecFromStr("1")), "busd:usd", sdkmath.NewInt(BusdCf), types3.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyMustNewDecFromStr("2"), sdkmath.LegacyMustNewDecFromStr("0.8"), sdkmath.LegacyMustNewDecFromStr("10")), sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyZeroDec()),
					types3.NewMoneyMarket("ujolt", types3.NewBorrowLimit(false, sdk.NewDec(100000000*JoltCf), tc.args.loanToValueJolt), "joltify:usd", sdkmath.NewInt(JoltCf), types3.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyMustNewDecFromStr("2"), sdkmath.LegacyMustNewDecFromStr("0.8"), sdkmath.LegacyMustNewDecFromStr("10")), sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyZeroDec()),
					types3.NewMoneyMarket("btcb", types3.NewBorrowLimit(false, sdk.NewDec(100000000*BtcbCf), tc.args.loanToValueBTCB), "btcb:usd", sdkmath.NewInt(BtcbCf), types3.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyMustNewDecFromStr("2"), sdkmath.LegacyMustNewDecFromStr("0.8"), sdkmath.LegacyMustNewDecFromStr("10")), sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyZeroDec()),
					types3.NewMoneyMarket("bnb", types3.NewBorrowLimit(false, sdk.NewDec(100000000*BnbCf), tc.args.loanToValueBNB), "bnb:usd", sdkmath.NewInt(BnbCf), types3.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyMustNewDecFromStr("2"), sdkmath.LegacyMustNewDecFromStr("0.8"), sdkmath.LegacyMustNewDecFromStr("10")), sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyZeroDec()),
					types3.NewMoneyMarket("xyz", types3.NewBorrowLimit(false, sdk.NewDec(1), tc.args.loanToValueBNB), "xyz:usd", sdkmath.NewInt(1), types3.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyMustNewDecFromStr("2"), sdkmath.LegacyMustNewDecFromStr("0.8"), sdkmath.LegacyMustNewDecFromStr("10")), sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyZeroDec()),
				},
				sdk.NewDec(10),
			), types3.DefaultAccumulationTimes, types3.DefaultDeposits, types3.DefaultBorrows,
				types3.DefaultTotalSupplied, types3.DefaultTotalBorrowed, types3.DefaultTotalReserves,
			)

			// Pricefeed module genesis state
			pricefeedGS := types2.GenesisState{
				Params: types2.Params{
					Markets: []types2.Market{
						{MarketID: "usdx:usd", BaseAsset: "usdx", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "busd:usd", BaseAsset: "busd", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "joltify:usd", BaseAsset: "joltify", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "btcb:usd", BaseAsset: "btcb", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "bnb:usd", BaseAsset: "bnb", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "xyz:usd", BaseAsset: "xyz", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
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
						MarketID:      "busd:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdkmath.LegacyMustNewDecFromStr("1.00"),
						Expiry:        time.Now().Add(1 * time.Hour),
					},
					{
						MarketID:      "joltify:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         tc.args.priceJolt,
						Expiry:        time.Now().Add(1 * time.Hour),
					},
					{
						MarketID:      "btcb:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         tc.args.priceBTCB,
						Expiry:        time.Now().Add(1 * time.Hour),
					},
					{
						MarketID:      "bnb:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         tc.args.priceBNB,
						Expiry:        time.Now().Add(1 * time.Hour),
					},
				},
			}

			// Initialize test application
			tApp.InitializeFromGenesisStates(nil, nil, authGS,
				app.GenesisState{types2.ModuleName: tApp.AppCodec().MustMarshalJSON(&pricefeedGS)},
				app.GenesisState{types3.ModuleName: tApp.AppCodec().MustMarshalJSON(&hardGS)})

			// Mint coins to jolt module account
			bankKeeper := tApp.GetBankKeeper()
			hardMaccCoins := sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf)),
				sdk.NewCoin("usdx", sdkmath.NewInt(200*UsdxCf)), sdk.NewCoin("busd", sdkmath.NewInt(100*BusdCf)))
			err := bankKeeper.MintCoins(ctx, types3.ModuleAccountName, hardMaccCoins)
			suite.Require().NoError(err)

			keeper := tApp.GetJoltKeeper()
			suite.app = tApp
			suite.ctx = ctx
			suite.keeper = keeper

			// Run BeginBlocker once to transition MoneyMarkets
			jolt.BeginBlocker(suite.ctx, suite.keeper)

			err = testutil.FundAccount(suite.app.GetBankKeeper(), suite.ctx, tc.args.borrower, coins)
			suite.Require().NoError(err)

			err = suite.keeper.Deposit(suite.ctx, tc.args.borrower, tc.args.depositCoins)
			suite.Require().NoError(err)

			// Execute user's previous borrows
			err = suite.keeper.Borrow(suite.ctx, tc.args.borrower, tc.args.previousBorrowCoins)
			if tc.args.previousBorrowCoins.IsZero() {
				suite.Require().True(strings.Contains(err.Error(), "cannot borrow zero coins"))
			} else {
				suite.Require().NoError(err)
			}

			// Now that our state is properly set up, execute the last borrow
			err = suite.keeper.Borrow(suite.ctx, tc.args.borrower, tc.args.borrowCoins)

			if tc.errArgs.expectPass {
				suite.Require().NoError(err)

				// Check borrower balance
				acc := suite.getAccount(tc.args.borrower)
				suite.Require().Equal(tc.args.expectedAccountBalance, suite.getAccountCoins(acc))

				// Check module account balance
				mAcc := suite.getModuleAccount(types3.ModuleAccountName)
				suite.Require().Equal(tc.args.expectedModAccountBalance, suite.getAccountCoins(mAcc))

				// Check that borrow struct is in store
				_, f := suite.keeper.GetBorrow(suite.ctx, tc.args.borrower)
				suite.Require().True(f)
			} else {
				suite.Require().Error(err)
				suite.Require().True(strings.Contains(err.Error(), tc.errArgs.contains))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestValidateBorrow() {
	blockDuration := time.Second * 3600 * 24 // long blocks to accumulate larger interest

	_, addrs := app.GeneratePrivKeyAddressPairs(5)
	borrower := addrs[0]
	initialBorrowerBalance := sdk.NewCoins(
		sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf)),
		sdk.NewCoin("usdx", sdkmath.NewInt(1000*JoltCf)),
	)

	model := types3.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("1.0"), sdkmath.LegacyMustNewDecFromStr("2"), sdkmath.LegacyMustNewDecFromStr("0.8"), sdkmath.LegacyMustNewDecFromStr("10"))

	// Initialize test app and set context
	tApp := app.NewTestApp(tmlog.TestingLogger(), suite.T().TempDir())
	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})

	// Auth module genesis state
	authGS := app.NewFundedGenStateWithSameCoins(
		tApp.AppCodec(),
		initialBorrowerBalance,
		[]sdk.AccAddress{borrower},
	)

	// Hard module genesis state
	hardGS := types3.NewGenesisState(
		types3.NewParams(
			types3.MoneyMarkets{
				types3.NewMoneyMarket("usdx",
					types3.NewBorrowLimit(false, sdk.NewDec(100000000*UsdxCf), sdkmath.LegacyMustNewDecFromStr("1")), // Borrow Limit
					"usdx:usd",                               // Market ID
					sdkmath.NewInt(UsdxCf),                   // Conversion Factor
					model,                                    // Interest Rate Model
					sdkmath.LegacyMustNewDecFromStr("1.0"),   // Reserve Factor (high)
					sdkmath.LegacyMustNewDecFromStr("0.05")), // Keeper Reward Percent
				types3.NewMoneyMarket("ujolt",
					types3.NewBorrowLimit(false, sdk.NewDec(100000000*JoltCf), sdkmath.LegacyMustNewDecFromStr("0.8")), // Borrow Limit
					"joltify:usd",                            // Market ID
					sdkmath.NewInt(JoltCf),                   // Conversion Factor
					model,                                    // Interest Rate Model
					sdkmath.LegacyMustNewDecFromStr("1.0"),   // Reserve Factor (high)
					sdkmath.LegacyMustNewDecFromStr("0.05")), // Keeper Reward Percent
			},
			sdk.NewDec(10),
		),
		types3.DefaultAccumulationTimes,
		types3.DefaultDeposits,
		types3.DefaultBorrows,
		types3.DefaultTotalSupplied,
		types3.DefaultTotalBorrowed,
		types3.DefaultTotalReserves,
	)

	// Pricefeed module genesis state
	pricefeedGS := types2.GenesisState{
		Params: types2.Params{
			Markets: []types2.Market{
				{MarketID: "usdx:usd", BaseAsset: "usdx", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "joltify:usd", BaseAsset: "joltify", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
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
		},
	}

	// Initialize test application
	tApp.InitializeFromGenesisStates(
		nil, nil, authGS,
		app.GenesisState{types2.ModuleName: tApp.AppCodec().MustMarshalJSON(&pricefeedGS)},
		app.GenesisState{types3.ModuleName: tApp.AppCodec().MustMarshalJSON(&hardGS)},
	)

	keeper := tApp.GetJoltKeeper()
	suite.app = tApp
	suite.ctx = ctx
	suite.keeper = keeper

	var err error

	// Run BeginBlocker once to transition MoneyMarkets
	jolt.BeginBlocker(suite.ctx, suite.keeper)

	// Setup borrower with some collateral to borrow against, and some reserve in the protocol.
	depositCoins := sdk.NewCoins(
		sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf)),
		sdk.NewCoin("usdx", sdkmath.NewInt(100*UsdxCf)),
	)

	err = testutil.FundAccount(suite.app.GetBankKeeper(), suite.ctx, borrower, depositCoins)
	suite.Require().NoError(err)

	err = suite.keeper.Deposit(suite.ctx, borrower, depositCoins)
	suite.Require().NoError(err)

	initialBorrowCoins := sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(70*JoltCf)))
	err = suite.keeper.Borrow(suite.ctx, borrower, initialBorrowCoins)
	suite.Require().NoError(err)

	runAtTime := suite.ctx.BlockTime().Add(blockDuration)
	suite.ctx = suite.ctx.WithBlockTime(runAtTime)
	jolt.BeginBlocker(suite.ctx, suite.keeper)

	repayCoins := sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))) // repay everything including accumulated interest

	err = testutil.FundAccount(suite.app.GetBankKeeper(), suite.ctx, borrower, repayCoins)
	suite.Require().NoError(err)

	err = suite.keeper.Repay(suite.ctx, borrower, borrower, repayCoins)
	suite.Require().NoError(err)

	// Get the total borrowable amount from the protocol, taking into account the reserves.
	modAccBalance := suite.getAccountCoins(suite.getModuleAccountAtCtx(types3.ModuleAccountName, suite.ctx))
	reserves, found := suite.keeper.GetTotalReserves(suite.ctx)
	suite.Require().True(found)
	availableToBorrow := modAccBalance.Sub(reserves...)

	// Test borrowing one over the available amount (try to borrow from the reserves)
	err = suite.keeper.Borrow(
		suite.ctx,
		borrower,
		sdk.NewCoins(sdk.NewCoin("ujolt", availableToBorrow.AmountOf("ujolt").Add(sdk.OneInt()))),
	)
	suite.Require().Error(err)

	// Test borrowing exactly the limit
	err = suite.keeper.Borrow(
		suite.ctx,
		borrower,
		sdk.NewCoins(sdk.NewCoin("ujolt", availableToBorrow.AmountOf("ujolt"))),
	)
	suite.Require().NoError(err)
}
