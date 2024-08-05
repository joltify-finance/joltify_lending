package keeper_test

import (
	"time"

	tmlog "cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/x/bank/testutil"

	"github.com/joltify-finance/joltify_lending/x/third_party/jolt"
	types3 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"

	"github.com/cometbft/cometbft/crypto"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/app"
)

func (suite *KeeperTestSuite) TestRepay() {
	type args struct {
		borrower             sdk.AccAddress
		repayer              sdk.AccAddress
		initialBorrowerCoins sdk.Coins
		initialRepayerCoins  sdk.Coins
		initialModuleCoins   sdk.Coins
		depositCoins         []sdk.Coin
		borrowCoins          sdk.Coins
		repayCoins           sdk.Coins
	}

	type errArgs struct {
		expectPass   bool
		expectDelete bool
		contains     string
	}

	type borrowTest struct {
		name    string
		args    args
		errArgs errArgs
	}

	model := types3.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyMustNewDecFromStr("2"), sdkmath.LegacyMustNewDecFromStr("0.8"), sdkmath.LegacyMustNewDecFromStr("10"))

	testCases := []borrowTest{
		{
			"valid: partial repay",
			args{
				borrower:             sdk.AccAddress(crypto.AddressHash([]byte("borrower"))),
				repayer:              sdk.AccAddress(crypto.AddressHash([]byte("borrower"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialRepayerCoins:  sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf)), sdk.NewCoin("usdx", sdkmath.NewInt(1000*UsdxCf))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(50*JoltCf))),
				repayCoins:           sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(10*JoltCf))),
			},
			errArgs{
				expectPass:   true,
				expectDelete: false,
				contains:     "",
			},
		},
		{
			"valid: partial repay by non borrower",
			args{
				borrower:             sdk.AccAddress(crypto.AddressHash([]byte("borrower"))),
				repayer:              sdk.AccAddress(crypto.AddressHash([]byte("repayer"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialRepayerCoins:  sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf)), sdk.NewCoin("usdx", sdkmath.NewInt(1000*UsdxCf))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(50*JoltCf))),
				repayCoins:           sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(10*JoltCf))),
			},
			errArgs{
				expectPass:   true,
				expectDelete: false,
				contains:     "",
			},
		},
		{
			"valid: repay in full",
			args{
				borrower:             sdk.AccAddress(crypto.AddressHash([]byte("borrower"))),
				repayer:              sdk.AccAddress(crypto.AddressHash([]byte("borrower"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialRepayerCoins:  sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf)), sdk.NewCoin("usdx", sdkmath.NewInt(1000*UsdxCf))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(50*JoltCf))),
				repayCoins:           sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(50*JoltCf))),
			},
			errArgs{
				expectPass:   true,
				expectDelete: true,
				contains:     "",
			},
		},
		{
			"valid: overpayment is adjusted",
			args{
				borrower:             sdk.AccAddress(crypto.AddressHash([]byte("borrower"))),
				repayer:              sdk.AccAddress(crypto.AddressHash([]byte("borrower"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialRepayerCoins:  sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf)), sdk.NewCoin("usdx", sdkmath.NewInt(1000*UsdxCf))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(80*JoltCf))), // Deposit less so user still has some KAVA
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(50*JoltCf))),
				repayCoins:           sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(60*JoltCf))), // Exceeds borrowed coins but not user's balance
			},
			errArgs{
				expectPass:   true,
				expectDelete: true,
				contains:     "",
			},
		},
		{
			"invalid: attempt to repay non-supplied coin",
			args{
				borrower:             sdk.AccAddress(crypto.AddressHash([]byte("borrower"))),
				repayer:              sdk.AccAddress(crypto.AddressHash([]byte("borrower"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialRepayerCoins:  sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf)), sdk.NewCoin("usdx", sdkmath.NewInt(1000*UsdxCf))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(50*JoltCf))),
				repayCoins:           sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(10*JoltCf)), sdk.NewCoin("bnb", sdkmath.NewInt(10*JoltCf))),
			},
			errArgs{
				expectPass:   false,
				expectDelete: false,
				contains:     "no coins of this type borrowed",
			},
		},
		{
			"invalid: insufficient balance for repay",
			args{
				borrower:             sdk.AccAddress(crypto.AddressHash([]byte("borrower"))),
				repayer:              sdk.AccAddress(crypto.AddressHash([]byte("repayer"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialRepayerCoins:  sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(49*JoltCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf)), sdk.NewCoin("usdx", sdkmath.NewInt(1000*UsdxCf))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(50*JoltCf))),
				repayCoins:           sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(50*JoltCf))), // Exceeds repayer's balance, but not borrow amount
			},
			errArgs{
				expectPass:   false,
				expectDelete: false,
				contains:     "account can only repay up to 49000000ujolt",
			},
		},
		{
			"invalid: repaying a single coin type results in borrow position below the minimum USD value",
			args{
				borrower:             sdk.AccAddress(crypto.AddressHash([]byte("borrower"))),
				repayer:              sdk.AccAddress(crypto.AddressHash([]byte("borrower"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(100*UsdxCf))),
				initialRepayerCoins:  sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(100*UsdxCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf)), sdk.NewCoin("usdx", sdkmath.NewInt(1000*UsdxCf))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(100*UsdxCf))),
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(50*UsdxCf))),
				repayCoins:           sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(45*UsdxCf))),
			},
			errArgs{
				expectPass:   false,
				expectDelete: false,
				contains:     "proposed borrow's USD value $5.000000000000000000 is below the minimum borrow limit",
			},
		},
		{
			"invalid: repaying multiple coin types results in borrow position below the minimum USD value",
			args{
				borrower:             sdk.AccAddress(crypto.AddressHash([]byte("borrower"))),
				repayer:              sdk.AccAddress(crypto.AddressHash([]byte("borrower"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(100*UsdxCf))),
				initialRepayerCoins:  sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(100*UsdxCf)), sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf)), sdk.NewCoin("usdx", sdkmath.NewInt(1000*UsdxCf))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(100*UsdxCf))),
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(50*UsdxCf)), sdk.NewCoin("ujolt", sdkmath.NewInt(10*JoltCf))), // (50*$1)+(10*$2) = $70
				repayCoins:           sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(45*UsdxCf)), sdk.NewCoin("ujolt", sdkmath.NewInt(8*JoltCf))),  // (45*$1)+(8*$2) = $61
			},
			errArgs{
				expectPass:   false,
				expectDelete: false,
				contains:     "proposed borrow's USD value $9.000000000000000000 is below the minimum borrow limit",
			},
		},
		{
			"invalid: overpaying multiple coin types results in borrow position below the minimum USD value",
			args{
				borrower:             sdk.AccAddress(crypto.AddressHash([]byte("borrower"))),
				repayer:              sdk.AccAddress(crypto.AddressHash([]byte("borrower"))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(100*UsdxCf))),
				initialRepayerCoins:  sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(100*UsdxCf)), sdk.NewCoin("ujolt", sdkmath.NewInt(100*JoltCf))),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1000*JoltCf)), sdk.NewCoin("usdx", sdkmath.NewInt(1000*UsdxCf))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(100*UsdxCf))),
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(50*UsdxCf)), sdk.NewCoin("ujolt", sdkmath.NewInt(10*JoltCf))), // (50*$1)+(10*$2) = $70
				repayCoins:           sdk.NewCoins(sdk.NewCoin("usdx", sdkmath.NewInt(500*UsdxCf)), sdk.NewCoin("ujolt", sdkmath.NewInt(8*JoltCf))), // (500*$1)+(8*$2) = $516, or capping to borrowed amount, (50*$1)+(8*$2) = $66
			},
			errArgs{
				expectPass:   false,
				expectDelete: false,
				contains:     "proposed borrow's USD value $4.000000000000000000 is below the minimum borrow limit",
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			// Initialize test app and set context
			tApp := app.NewTestApp(tmlog.TestingLogger(), suite.T().TempDir())
			ctx := tApp.NewContext(true)

			// Auth module genesis state
			addrs, coinses := uniqueAddressCoins(
				[]sdk.AccAddress{tc.args.borrower, tc.args.repayer},
				[]sdk.Coins{tc.args.initialBorrowerCoins, tc.args.initialRepayerCoins},
			)
			authGS := app.NewFundedGenStateWithCoins(
				tApp.AppCodec(),
				coinses,
				addrs,
			)

			// Hard module genesis state
			hardGS := types3.NewGenesisState(types3.NewParams(
				types3.MoneyMarkets{
					types3.NewMoneyMarket("usdx",
						types3.NewBorrowLimit(false, sdkmath.LegacyNewDec(100000000*UsdxCf), sdkmath.LegacyMustNewDecFromStr("1")), // Borrow Limit
						"usdx:usd",                               // Market ID
						sdkmath.NewInt(UsdxCf),                   // Conversion Factor
						model,                                    // Interest Rate Model
						sdkmath.LegacyMustNewDecFromStr("0.05"),  // Reserve Factor
						sdkmath.LegacyMustNewDecFromStr("0.05")), // Keeper Reward Percent
					types3.NewMoneyMarket("ujolt",
						types3.NewBorrowLimit(false, sdkmath.LegacyNewDec(100000000*JoltCf), sdkmath.LegacyMustNewDecFromStr("0.8")), // Borrow Limit
						"joltify:usd",                            // Market ID
						sdkmath.NewInt(JoltCf),                   // Conversion Factor
						model,                                    // Interest Rate Model
						sdkmath.LegacyMustNewDecFromStr("0.05"),  // Reserve Factor
						sdkmath.LegacyMustNewDecFromStr("0.05")), // Keeper Reward Percent
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
			tApp.InitializeFromGenesisStates(nil, nil, authGS,
				app.GenesisState{types2.ModuleName: tApp.AppCodec().MustMarshalJSON(&pricefeedGS)},
				app.GenesisState{types3.ModuleName: tApp.AppCodec().MustMarshalJSON(&hardGS)},
			)

			// Mint coins to Hard module account
			bankKeeper := tApp.GetBankKeeper()
			err := bankKeeper.MintCoins(ctx, types3.ModuleAccountName, tc.args.initialModuleCoins)
			suite.Require().NoError(err)

			suite.app = tApp
			suite.ctx = ctx
			suite.keeper = tApp.GetJoltKeeper()

			// Run BeginBlocker once to transition MoneyMarkets
			jolt.BeginBlocker(suite.ctx, suite.keeper)

			err = testutil.FundAccount(suite.app.GetBankKeeper(), suite.ctx, tc.args.borrower, tc.args.initialBorrowerCoins)
			suite.Require().NoError(err)

			err = testutil.FundAccount(suite.app.GetBankKeeper(), suite.ctx, tc.args.repayer, tc.args.initialRepayerCoins)
			suite.Require().NoError(err)

			// Deposit coins to jolt
			err = suite.keeper.Deposit(suite.ctx, tc.args.borrower, tc.args.depositCoins)
			suite.Require().NoError(err)

			// Borrow coins from jolt
			err = suite.keeper.Borrow(suite.ctx, tc.args.borrower, tc.args.borrowCoins)
			suite.Require().NoError(err)

			repayerAcc := suite.getAccount(tc.args.repayer)
			previousRepayerCoins := bankKeeper.GetAllBalances(suite.ctx, repayerAcc.GetAddress())

			err = suite.keeper.Repay(suite.ctx, tc.args.repayer, tc.args.borrower, tc.args.repayCoins)
			if tc.errArgs.expectPass {
				suite.Require().NoError(err)
				// If we overpaid expect an adjustment
				repaymentCoins, err := suite.keeper.CalculatePaymentAmount(tc.args.borrowCoins, tc.args.repayCoins)
				suite.Require().NoError(err)

				// Check repayer balance
				expectedRepayerCoins := previousRepayerCoins.Sub(repaymentCoins...)
				acc := suite.getAccount(tc.args.repayer)
				// use IsEqual for sdk.Coins{nil} vs sdk.Coins{}
				suite.Require().True(expectedRepayerCoins.IsEqual(bankKeeper.GetAllBalances(suite.ctx, acc.GetAddress())))

				// Check module account balance
				expectedModuleCoins := tc.args.initialModuleCoins.Add(tc.args.depositCoins...).Sub(tc.args.borrowCoins...).Add(repaymentCoins...)
				mAcc := suite.getModuleAccount(types3.ModuleAccountName)
				suite.Require().Equal(expectedModuleCoins, bankKeeper.GetAllBalances(suite.ctx, mAcc.GetAddress()))

				// Check user's borrow object
				borrow, foundBorrow := suite.keeper.GetBorrow(suite.ctx, tc.args.borrower)
				expectedBorrowCoins := tc.args.borrowCoins.Sub(repaymentCoins...)

				if tc.errArgs.expectDelete {
					suite.Require().False(foundBorrow)
				} else {
					suite.Require().True(foundBorrow)
					suite.Require().Equal(expectedBorrowCoins, borrow.Amount)
				}
			} else {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.errArgs.contains)

				// Check repayer balance (no repay coins)
				acc := suite.getAccount(tc.args.repayer)
				suite.Require().Equal(previousRepayerCoins, bankKeeper.GetAllBalances(suite.ctx, acc.GetAddress()))

				// Check module account balance (no repay coins)
				expectedModuleCoins := tc.args.initialModuleCoins.Add(tc.args.depositCoins...).Sub(tc.args.borrowCoins...)
				mAcc := suite.getModuleAccount(types3.ModuleAccountName)
				suite.Require().Equal(expectedModuleCoins, bankKeeper.GetAllBalances(suite.ctx, mAcc.GetAddress()))

				// Check user's borrow object (no repay coins)
				borrow, foundBorrow := suite.keeper.GetBorrow(suite.ctx, tc.args.borrower)
				suite.Require().True(foundBorrow)
				suite.Require().Equal(tc.args.borrowCoins, borrow.Amount)
			}
		})
	}
}

// uniqueAddressCoins removes duplicate addresses, and the corresponding elements in a list of coins.
func uniqueAddressCoins(addresses []sdk.AccAddress, coinses []sdk.Coins) ([]sdk.AccAddress, []sdk.Coins) {
	var uniqueAddresses []sdk.AccAddress
	var filteredCoins []sdk.Coins

	addrMap := map[string]bool{}
	for i, a := range addresses {
		if !addrMap[a.String()] {
			uniqueAddresses = append(uniqueAddresses, a)
			filteredCoins = append(filteredCoins, coinses[i])
		}
		addrMap[a.String()] = true
	}
	return uniqueAddresses, filteredCoins
}
