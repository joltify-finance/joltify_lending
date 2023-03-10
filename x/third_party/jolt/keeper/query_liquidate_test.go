package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
	types4 "github.com/joltify-finance/joltify_lending/x/third_party/auction/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/jolt"
	joltKeeper "github.com/joltify-finance/joltify_lending/x/third_party/jolt/keeper"
	types3 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"
	"github.com/tendermint/tendermint/crypto"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func (suite *KeeperTestSuite) TestKeeperQueryLiquidation() {
	type args struct {
		borrower                   []sdk.AccAddress
		keeper                     sdk.AccAddress
		keeperRewardPercent        sdk.Dec
		initialModuleCoins         sdk.Coins
		initialBorrowerCoins       sdk.Coins
		initialKeeperCoins         sdk.Coins
		depositCoins               []sdk.Coin
		borrowCoins                sdk.Coins
		liquidateAfter             time.Duration
		expectedTotalSuppliedCoins sdk.Coins
		expectedTotalBorrowedCoins sdk.Coins
		expectedKeeperCoins        sdk.Coins        // coins keeper address should have after successfully liquidating position
		expectedBorrowerCoins      sdk.Coins        // additional coins (if any) the borrower address should have after successfully liquidating position
		expectedAuctions           []types4.Auction // the auctions we should expect to find have been started
	}

	type errArgs struct {
		expectPass bool
		contains   string
	}

	type liqTest struct {
		name    string
		args    args
		errArgs errArgs
	}

	// Set up test constants
	model := types3.NewInterestRateModel(sdk.MustNewDecFromStr("0"), sdk.MustNewDecFromStr("0.1"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("0.5"))
	reserveFactor := sdk.MustNewDecFromStr("0.05")
	oneMonthDur := time.Second * 30 * 24 * 3600
	borrower := sdk.AccAddress(crypto.AddressHash([]byte("testborrower")))
	keeper := sdk.AccAddress(crypto.AddressHash([]byte("testkeeper")))

	// Set up auction constants
	layout := "2006-01-02T15:04:05.000Z"
	endTimeStr := "9000-01-01T00:00:00.000Z"
	endTime, _ := time.Parse(layout, endTimeStr)

	lotReturns, _ := types4.NewWeightedAddresses([]sdk.AccAddress{borrower}, []sdk.Int{sdk.NewInt(100)})

	testCases := []liqTest{
		{
			"valid: keeper liquidates borrow",
			args{
				borrower:                   []sdk.AccAddress{borrower},
				keeper:                     keeper,
				keeperRewardPercent:        sdk.MustNewDecFromStr("0.05"),
				initialModuleCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF))),
				initialBorrowerCoins:       sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF))),
				initialKeeperCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF))),
				depositCoins:               sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(10*JOLT_CF))),
				borrowCoins:                sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(8*JOLT_CF))),
				liquidateAfter:             oneMonthDur,
				expectedTotalSuppliedCoins: sdk.NewCoins(sdk.NewInt64Coin("ujolt", 100004118)),
				expectedTotalBorrowedCoins: nil,
				expectedKeeperCoins:        sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100500020))),
				expectedBorrowerCoins:      sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(98000001))), // initial - deposit + borrow + liquidation leftovers
				expectedAuctions: []types4.Auction{
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              1,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("ujolt", 9500390),
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("ujolt", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("ujolt", 8004766),
						LotReturns:        lotReturns,
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"valid: 0% keeper rewards",
			args{
				borrower:                   []sdk.AccAddress{borrower},
				keeper:                     keeper,
				keeperRewardPercent:        sdk.MustNewDecFromStr("0.0"),
				initialModuleCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF))),
				initialBorrowerCoins:       sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF))),
				initialKeeperCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF))),
				depositCoins:               sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(10*JOLT_CF))),
				borrowCoins:                sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(8*JOLT_CF))),
				liquidateAfter:             oneMonthDur,
				expectedTotalSuppliedCoins: sdk.NewCoins(sdk.NewInt64Coin("ujolt", 100_004_117)),
				expectedTotalBorrowedCoins: sdk.NewCoins(sdk.NewInt64Coin("ujolt", 1)),
				expectedKeeperCoins:        sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF))),
				expectedBorrowerCoins:      sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(98*JOLT_CF))), // initial - deposit + borrow + liquidation leftovers
				expectedAuctions: []types4.Auction{
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              1,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("ujolt", 10000411),
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("ujolt", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("ujolt", 8004765),
						LotReturns:        lotReturns,
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"valid: 100% keeper reward",
			args{
				borrower:                   []sdk.AccAddress{borrower},
				keeper:                     keeper,
				keeperRewardPercent:        sdk.MustNewDecFromStr("1.0"),
				initialModuleCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF))),
				initialBorrowerCoins:       sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF))),
				initialKeeperCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF))),
				depositCoins:               sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(10*JOLT_CF))),
				borrowCoins:                sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(8*JOLT_CF))),
				liquidateAfter:             oneMonthDur,
				expectedTotalSuppliedCoins: sdk.NewCoins(sdk.NewInt64Coin("ujolt", 100_004_117)),
				expectedTotalBorrowedCoins: sdk.NewCoins(sdk.NewInt64Coin("ujolt", 8_004_766)),
				expectedKeeperCoins:        sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(110_000_411))),
				expectedBorrowerCoins:      sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(98*JOLT_CF))), // initial - deposit + borrow + liquidation leftovers
				expectedAuctions:           nil,
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"valid: single deposit, multiple borrows",
			args{
				borrower:             []sdk.AccAddress{borrower},
				keeper:               keeper,
				keeperRewardPercent:  sdk.MustNewDecFromStr("0.05"),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("usdc", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("bnb", sdk.NewInt(1000*BNB_CF)), sdk.NewCoin("btc", sdk.NewInt(1000*BTCB_CF))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF))),
				initialKeeperCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(50*JOLT_CF))),                                                                                                                                     // $100 * 0.8 = $80 borrowable
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("usdc", sdk.NewInt(20*JOLT_CF)), sdk.NewCoin("ujolt", sdk.NewInt(10*JOLT_CF)), sdk.NewCoin("bnb", sdk.NewInt(2*BNB_CF)), sdk.NewCoin("btc", sdk.NewInt(0.2*BTCB_CF))), // $20+$20+$20 = $80 borrowed
				liquidateAfter:       oneMonthDur,
				expectedTotalSuppliedCoins: sdk.NewCoins(
					sdk.NewInt64Coin("ujolt", 1000000710),
					sdk.NewInt64Coin("usdc", 1000003120),
					sdk.NewInt64Coin("bnb", 100000003123),
					sdk.NewInt64Coin("btc", 100000000031),
				),
				expectedTotalBorrowedCoins: nil,
				expectedKeeperCoins:        sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(102500001))),
				expectedBorrowerCoins:      sdk.NewCoins(sdk.NewCoin("usdc", sdk.NewInt(20*JOLT_CF)), sdk.NewCoin("ujolt", sdk.NewInt(60000002)), sdk.NewCoin("bnb", sdk.NewInt(2*BNB_CF)), sdk.NewCoin("btc", sdk.NewInt(0.2*BTCB_CF))), // initial - deposit + borrow + liquidation leftovers
				expectedAuctions: []types4.Auction{
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              1,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("ujolt", 11874430),
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("bnb", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("bnb", 200003287),
						LotReturns:        lotReturns,
					},
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              2,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("ujolt", 11874254),
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("btc", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("btc", 20000032),
						LotReturns:        lotReturns,
					},
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              3,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("ujolt", 11875163),
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("ujolt", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("ujolt", 10000782),
						LotReturns:        lotReturns,
					},
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              4,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("ujolt", 11876185),
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("usdc", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("usdc", 20003284),
						LotReturns:        lotReturns,
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"valid: multiple deposits, single borrow",
			args{
				borrower:             []sdk.AccAddress{borrower},
				keeper:               keeper,
				keeperRewardPercent:  sdk.MustNewDecFromStr("0.05"),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(1000*JOLT_CF))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF)), sdk.NewCoin("bnb", sdk.NewInt(100*BNB_CF)), sdk.NewCoin("btc", sdk.NewInt(100*BTCB_CF))),
				initialKeeperCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(50*JOLT_CF)), sdk.NewCoin("bnb", sdk.NewInt(10*BNB_CF)), sdk.NewCoin("btc", sdk.NewInt(1*BTCB_CF))), // $100 + $100 + $100 = $300 * 0.8 = $240 borrowable                                                                                                                                       // $100 * 0.8 = $80 borrowable
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(120*JOLT_CF))),                                                                                      // $240 borrowed
				liquidateAfter:       oneMonthDur,
				expectedTotalSuppliedCoins: sdk.NewCoins(
					sdk.NewInt64Coin("ujolt", 1000101456),
				),
				expectedTotalBorrowedCoins: nil,
				expectedKeeperCoins:        sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(102500253)), sdk.NewCoin("bnb", sdk.NewInt(0.5*BNB_CF)), sdk.NewCoin("btc", sdk.NewInt(0.05*BTCB_CF))), // 5% of each seized coin + initial balances
				expectedBorrowerCoins:      sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(170.000001*JOLT_CF)), sdk.NewCoin("bnb", sdk.NewInt(90*BNB_CF)), sdk.NewCoin("btc", sdk.NewInt(99*BTCB_CF))),
				expectedAuctions: []types4.Auction{
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              1,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("bnb", 950000000),
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("ujolt", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("ujolt", 40036023),
						LotReturns:        lotReturns,
					},
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              2,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("btc", 95000000),
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("ujolt", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("ujolt", 40036023),
						LotReturns:        lotReturns,
					},
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              3,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("ujolt", 47504818),
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("ujolt", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("ujolt", 40040087),
						LotReturns:        lotReturns,
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"valid: multiple stablecoin deposits, multiple variable coin borrows",
			// Auctions: total lot value = $285 ($300 of deposits - $15 keeper reward), total max bid value = $270
			args{
				borrower:             []sdk.AccAddress{borrower},
				keeper:               keeper,
				keeperRewardPercent:  sdk.MustNewDecFromStr("0.05"),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("bnb", sdk.NewInt(1000*BNB_CF)), sdk.NewCoin("btc", sdk.NewInt(1000*BTCB_CF))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF)), sdk.NewCoin("usdc", sdk.NewInt(100*JOLT_CF)), sdk.NewCoin("usdt", sdk.NewInt(100*JOLT_CF)), sdk.NewCoin("usdx", sdk.NewInt(100*JOLT_CF))),
				initialKeeperCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("usdc", sdk.NewInt(100*JOLT_CF)), sdk.NewCoin("usdt", sdk.NewInt(100*JOLT_CF)), sdk.NewCoin("usdx", sdk.NewInt(100*JOLT_CF))), // $100 + $100 + $100 = $300 * 0.9 = $270 borrowable
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(35*JOLT_CF)), sdk.NewCoin("bnb", sdk.NewInt(10*BNB_CF)), sdk.NewCoin("btc", sdk.NewInt(1*BTCB_CF))),       // $270 borrowed
				liquidateAfter:       oneMonthDur,
				expectedTotalSuppliedCoins: sdk.NewCoins(
					sdk.NewInt64Coin("bnb", 100000078047),
					sdk.NewInt64Coin("btc", 100000000780),
					sdk.NewInt64Coin("ujolt", 1000009550),
					sdk.NewInt64Coin("usdx", 1),
				),
				expectedTotalBorrowedCoins: nil,
				expectedKeeperCoins:        sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF)), sdk.NewCoin("usdc", sdk.NewInt(5*JOLT_CF)), sdk.NewCoin("usdt", sdk.NewInt(5*JOLT_CF)), sdk.NewCoin("usdx", sdk.NewInt(5*JOLT_CF))), // 5% of each seized coin + initial balances
				expectedBorrowerCoins:      sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(135*JOLT_CF)), sdk.NewCoin("bnb", sdk.NewInt(10*BNB_CF)), sdk.NewCoin("btc", sdk.NewInt(1*BTCB_CF)), sdk.NewCoin("usdx", sdk.NewInt(0.000001*JOLT_CF))),
				expectedAuctions: []types4.Auction{
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              1,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("usdc", 95000000), // $95.00
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("bnb", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("bnb", 900097134), // $90.00
						LotReturns:        lotReturns,
					},
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              2,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("usdt", 10552835), // $10.55
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("bnb", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("bnb", 99985020), // $10.00
						LotReturns:        lotReturns,
					},
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              3,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("usdt", 84447165), // $84.45
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("btc", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("btc", 80011211), // $80.01
						LotReturns:        lotReturns,
					},
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              4,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("usdx", 21097866), // $21.10
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("btc", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("btc", 19989610), // $19.99
						LotReturns:        lotReturns,
					},
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              5,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("usdx", 73902133), //$73.90
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("ujolt", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("ujolt", 35010052), // $70.02
						LotReturns:        lotReturns,
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"valid: multiple stablecoin deposits, multiple stablecoin borrows",
			args{
				borrower:             []sdk.AccAddress{borrower},
				keeper:               keeper,
				keeperRewardPercent:  sdk.MustNewDecFromStr("0.05"),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("usdx", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("usdt", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("dai", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("usdc", sdk.NewInt(1000*JOLT_CF))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("usdx", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("usdt", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("dai", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("usdc", sdk.NewInt(1000*JOLT_CF))),
				initialKeeperCoins:   sdk.NewCoins(sdk.NewCoin("usdx", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("usdt", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("dai", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("usdc", sdk.NewInt(1000*JOLT_CF))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("dai", sdk.NewInt(350*JOLT_CF)), sdk.NewCoin("usdc", sdk.NewInt(200*JOLT_CF))),
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("usdt", sdk.NewInt(250*JOLT_CF)), sdk.NewCoin("usdx", sdk.NewInt(245*JOLT_CF))),
				liquidateAfter:       oneMonthDur,
				expectedTotalSuppliedCoins: sdk.NewCoins(
					sdk.NewInt64Coin("dai", 1000000000),
					sdk.NewInt64Coin("usdc", 1000000001),
					sdk.NewInt64Coin("usdt", 1000482503),
					sdk.NewInt64Coin("usdx", 1000463500),
				),
				expectedTotalBorrowedCoins: nil,
				expectedKeeperCoins:        sdk.NewCoins(sdk.NewCoin("dai", sdk.NewInt(1017.50*JOLT_CF)), sdk.NewCoin("usdt", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("usdc", sdk.NewInt(1010*JOLT_CF)), sdk.NewCoin("usdx", sdk.NewInt(1000*JOLT_CF))),
				expectedBorrowerCoins:      sdk.NewCoins(sdk.NewCoin("dai", sdk.NewInt(650*JOLT_CF)), sdk.NewCoin("usdc", sdk.NewInt(800000001)), sdk.NewCoin("usdt", sdk.NewInt(1250*JOLT_CF)), sdk.NewCoin("usdx", sdk.NewInt(1245*JOLT_CF))),
				expectedAuctions: []types4.Auction{
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              1,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("dai", 263894126),
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("usdt", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("usdt", 250507897),
						LotReturns:        lotReturns,
					},
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              2,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("dai", 68605874),
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("usdx", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("usdx", 65125788),
						LotReturns:        lotReturns,
					},
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              3,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("usdc", 189999999),
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("usdx", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("usdx", 180362106),
						LotReturns:        lotReturns,
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			// Initialize test app and set context
			tApp := app.NewTestApp()
			ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC)})

			// account which will deposit "initial module account coins"
			depositor := sdk.AccAddress(crypto.AddressHash([]byte("testdepositor")))

			// Auth module genesis state
			authGS := app.NewFundedGenStateWithCoins(
				tApp.AppCodec(),
				[]sdk.Coins{
					tc.args.initialBorrowerCoins,
					tc.args.initialKeeperCoins,
					tc.args.initialModuleCoins,
				},
				[]sdk.AccAddress{
					tc.args.borrower[0],
					tc.args.keeper,
					depositor,
				},
			)

			// Hard module genesis state
			joltGS := types3.NewGenesisState(types3.NewParams(
				types3.MoneyMarkets{
					types3.NewMoneyMarket("usdx",
						types3.NewBorrowLimit(false, sdk.NewDec(100000000*JOLT_CF), sdk.MustNewDecFromStr("0.9")), // Borrow Limit
						"usdx:usd",                   // Market ID
						sdk.NewInt(JOLT_CF),          // Conversion Factor
						model,                        // Interest Rate Model
						reserveFactor,                // Reserve Factor
						tc.args.keeperRewardPercent), // Keeper Reward Percent
					types3.NewMoneyMarket("usdt",
						types3.NewBorrowLimit(false, sdk.NewDec(100000000*JOLT_CF), sdk.MustNewDecFromStr("0.9")), // Borrow Limit
						"usdt:usd",                   // Market ID
						sdk.NewInt(JOLT_CF),          // Conversion Factor
						model,                        // Interest Rate Model
						reserveFactor,                // Reserve Factor
						tc.args.keeperRewardPercent), // Keeper Reward Percent
					types3.NewMoneyMarket("usdc",
						types3.NewBorrowLimit(false, sdk.NewDec(100000000*JOLT_CF), sdk.MustNewDecFromStr("0.9")), // Borrow Limit
						"usdc:usd",                   // Market ID
						sdk.NewInt(JOLT_CF),          // Conversion Factor
						model,                        // Interest Rate Model
						reserveFactor,                // Reserve Factor
						tc.args.keeperRewardPercent), // Keeper Reward Percent
					types3.NewMoneyMarket("dai",
						types3.NewBorrowLimit(false, sdk.NewDec(100000000*JOLT_CF), sdk.MustNewDecFromStr("0.9")), // Borrow Limit
						"dai:usd",                    // Market ID
						sdk.NewInt(JOLT_CF),          // Conversion Factor
						model,                        // Interest Rate Model
						reserveFactor,                // Reserve Factor
						tc.args.keeperRewardPercent), // Keeper Reward Percent
					types3.NewMoneyMarket("ujolt",
						types3.NewBorrowLimit(false, sdk.NewDec(100000000*JOLT_CF), sdk.MustNewDecFromStr("0.8")), // Borrow Limit
						"joltify:usd",                // Market ID
						sdk.NewInt(JOLT_CF),          // Conversion Factor
						model,                        // Interest Rate Model
						reserveFactor,                // Reserve Factor
						tc.args.keeperRewardPercent), // Keeper Reward Percent
					types3.NewMoneyMarket("bnb",
						types3.NewBorrowLimit(false, sdk.NewDec(100000000*BNB_CF), sdk.MustNewDecFromStr("0.8")), // Borrow Limit
						"bnb:usd",                    // Market ID
						sdk.NewInt(BNB_CF),           // Conversion Factor
						model,                        // Interest Rate Model
						reserveFactor,                // Reserve Factor
						tc.args.keeperRewardPercent), // Keeper Reward Percent
					types3.NewMoneyMarket("btc",
						types3.NewBorrowLimit(false, sdk.NewDec(100000000*BTCB_CF), sdk.MustNewDecFromStr("0.8")), // Borrow Limit
						"btc:usd",                    // Market ID
						sdk.NewInt(BTCB_CF),          // Conversion Factor
						model,                        // Interest Rate Model
						reserveFactor,                // Reserve Factor
						tc.args.keeperRewardPercent), // Keeper Reward Percent
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
						{MarketID: "usdt:usd", BaseAsset: "usdt", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "usdc:usd", BaseAsset: "usdc", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "dai:usd", BaseAsset: "dai", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "joltify:usd", BaseAsset: "joltify", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "bnb:usd", BaseAsset: "bnb", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "btc:usd", BaseAsset: "btc", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
					},
				},
				PostedPrices: []types2.PostedPrice{
					{
						MarketID:      "usdx:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdk.MustNewDecFromStr("1.00"),
						Expiry:        time.Now().Add(100 * time.Hour),
					},
					{
						MarketID:      "usdt:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdk.MustNewDecFromStr("1.00"),
						Expiry:        time.Now().Add(100 * time.Hour),
					},
					{
						MarketID:      "usdc:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdk.MustNewDecFromStr("1.00"),
						Expiry:        time.Now().Add(100 * time.Hour),
					},
					{
						MarketID:      "dai:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdk.MustNewDecFromStr("1.00"),
						Expiry:        time.Now().Add(100 * time.Hour),
					},
					{
						MarketID:      "joltify:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdk.MustNewDecFromStr("2.00"),
						Expiry:        time.Now().Add(100 * time.Hour),
					},
					{
						MarketID:      "bnb:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdk.MustNewDecFromStr("10.00"),
						Expiry:        time.Now().Add(100 * time.Hour),
					},
					{
						MarketID:      "btc:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdk.MustNewDecFromStr("100.00"),
						Expiry:        time.Now().Add(100 * time.Hour),
					},
				},
			}

			// Initialize test application
			tApp.InitializeFromGenesisStates(authGS,
				app.GenesisState{types2.ModuleName: tApp.AppCodec().MustMarshalJSON(&pricefeedGS)},
				app.GenesisState{types3.ModuleName: tApp.AppCodec().MustMarshalJSON(&joltGS)})

			keeper := tApp.GetJoltKeeper()
			suite.app = tApp
			suite.ctx = ctx
			suite.keeper = keeper
			suite.auctionKeeper = tApp.GetAuctionKeeper()

			var err error

			// Run begin blocker to set up state
			jolt.BeginBlocker(suite.ctx, suite.keeper)

			// Deposit initial module account coins
			err = suite.keeper.Deposit(suite.ctx, depositor, tc.args.initialModuleCoins)
			suite.Require().NoError(err)

			// Deposit coins
			err = suite.keeper.Deposit(suite.ctx, tc.args.borrower[0], tc.args.depositCoins)
			suite.Require().NoError(err)

			// Borrow coins
			err = suite.keeper.Borrow(suite.ctx, tc.args.borrower[0], tc.args.borrowCoins)
			suite.Require().NoError(err)

			// Set up liquidation chain context and run begin blocker
			runAtTime := suite.ctx.BlockTime().Add(tc.args.liquidateAfter)
			liqCtx := suite.ctx.WithBlockTime(runAtTime)
			jolt.BeginBlocker(liqCtx, suite.keeper)

			// Check borrow exists before liquidation
			_, foundBorrowBefore := suite.keeper.GetBorrow(liqCtx, tc.args.borrower[0])
			suite.Require().True(foundBorrowBefore)
			// Check that the user's deposit exists before liquidation
			_, foundDepositBefore := suite.keeper.GetDeposit(liqCtx, tc.args.borrower[0])
			suite.Require().True(foundDepositBefore)

			queryServer := joltKeeper.NewQueryServerImpl(suite.keeper, suite.app.GetAccountKeeper(), suite.app.GetBankKeeper())

			req := types3.QueryLiquidateRequest{}
			ret, err := queryServer.Liquidate(sdk.WrapSDKContext(suite.ctx), &req)
			suite.Require().NoError(err)

			result := sdk.MustNewDecFromStr(ret.LiquidateItems[0].Ltv)
			suite.Require().True(result.GTE(sdk.MustNewDecFromStr("1.0")))
		})
	}
}

// we test with multiple liquidate items

func (suite *KeeperTestSuite) TestKeeperMultiQueryLiquidation() {
	type args struct {
		borrower                   []sdk.AccAddress
		keeper                     sdk.AccAddress
		keeperRewardPercent        sdk.Dec
		initialModuleCoins         sdk.Coins
		initialBorrowerCoins       sdk.Coins
		initialKeeperCoins         sdk.Coins
		depositCoins               []sdk.Coin
		borrowCoins                sdk.Coins
		liquidateAfter             time.Duration
		expectedTotalSuppliedCoins sdk.Coins
		expectedTotalBorrowedCoins sdk.Coins
		expectedKeeperCoins        sdk.Coins        // coins keeper address should have after successfully liquidating position
		expectedBorrowerCoins      sdk.Coins        // additional coins (if any) the borrower address should have after successfully liquidating position
		expectedAuctions           []types4.Auction // the auctions we should expect to find have been started
	}

	type errArgs struct {
		expectPass bool
		contains   string
	}

	type liqTest struct {
		name    string
		args    args
		errArgs errArgs
	}

	// Set up test constants
	model := types3.NewInterestRateModel(sdk.MustNewDecFromStr("0"), sdk.MustNewDecFromStr("0.1"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("0.5"))
	reserveFactor := sdk.MustNewDecFromStr("0.05")
	oneMonthDur := time.Second * 30 * 24 * 3600
	borrower1 := sdk.AccAddress(crypto.AddressHash([]byte("testborrower")))
	borrower2 := sdk.AccAddress(crypto.AddressHash([]byte("testborrower2")))
	keeper := sdk.AccAddress(crypto.AddressHash([]byte("testkeeper")))

	// Set up auction constants
	layout := "2006-01-02T15:04:05.000Z"
	endTimeStr := "9000-01-01T00:00:00.000Z"
	endTime, _ := time.Parse(layout, endTimeStr)

	lotReturns, _ := types4.NewWeightedAddresses([]sdk.AccAddress{borrower1}, []sdk.Int{sdk.NewInt(100)})

	testCases := []liqTest{
		{
			"valid: keeper liquidates borrow",
			args{
				borrower:                   []sdk.AccAddress{borrower1, borrower2},
				keeper:                     keeper,
				keeperRewardPercent:        sdk.MustNewDecFromStr("0.05"),
				initialModuleCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF))),
				initialBorrowerCoins:       sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF))),
				initialKeeperCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF))),
				depositCoins:               sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(10*JOLT_CF))),
				borrowCoins:                sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(8*JOLT_CF))),
				liquidateAfter:             oneMonthDur,
				expectedTotalSuppliedCoins: sdk.NewCoins(sdk.NewInt64Coin("ujolt", 100004118)),
				expectedTotalBorrowedCoins: nil,
				expectedKeeperCoins:        sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100500020))),
				expectedBorrowerCoins:      sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(98000001))), // initial - deposit + borrow + liquidation leftovers
				expectedAuctions: []types4.Auction{
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              1,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("ujolt", 9500390),
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("ujolt", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("ujolt", 8004766),
						LotReturns:        lotReturns,
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},

		{
			"valid: single deposit, multiple borrows",
			args{
				borrower:             []sdk.AccAddress{borrower1, borrower2},
				keeper:               keeper,
				keeperRewardPercent:  sdk.MustNewDecFromStr("0.05"),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("usdc", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("bnb", sdk.NewInt(1000*BNB_CF)), sdk.NewCoin("btc", sdk.NewInt(1000*BTCB_CF))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF))),
				initialKeeperCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(50*JOLT_CF))),                                                                                                                                     // $100 * 0.8 = $80 borrowable
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("usdc", sdk.NewInt(20*JOLT_CF)), sdk.NewCoin("ujolt", sdk.NewInt(10*JOLT_CF)), sdk.NewCoin("bnb", sdk.NewInt(2*BNB_CF)), sdk.NewCoin("btc", sdk.NewInt(0.2*BTCB_CF))), // $20+$20+$20 = $80 borrowed
				liquidateAfter:       oneMonthDur,
				expectedTotalSuppliedCoins: sdk.NewCoins(
					sdk.NewInt64Coin("ujolt", 1000000710),
					sdk.NewInt64Coin("usdc", 1000003120),
					sdk.NewInt64Coin("bnb", 100000003123),
					sdk.NewInt64Coin("btc", 100000000031),
				),
				expectedTotalBorrowedCoins: nil,
				expectedKeeperCoins:        sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(102500001))),
				expectedBorrowerCoins:      sdk.NewCoins(sdk.NewCoin("usdc", sdk.NewInt(20*JOLT_CF)), sdk.NewCoin("ujolt", sdk.NewInt(60000002)), sdk.NewCoin("bnb", sdk.NewInt(2*BNB_CF)), sdk.NewCoin("btc", sdk.NewInt(0.2*BTCB_CF))), // initial - deposit + borrow + liquidation leftovers
				expectedAuctions: []types4.Auction{
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              1,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("ujolt", 11874430),
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("bnb", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("bnb", 200003287),
						LotReturns:        lotReturns,
					},
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              2,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("ujolt", 11874254),
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("btc", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("btc", 20000032),
						LotReturns:        lotReturns,
					},
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              3,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("ujolt", 11875163),
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("ujolt", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("ujolt", 10000782),
						LotReturns:        lotReturns,
					},
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              4,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("ujolt", 11876185),
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("usdc", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("usdc", 20003284),
						LotReturns:        lotReturns,
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"valid: multiple deposits, single borrow",
			args{
				borrower:             []sdk.AccAddress{borrower1, borrower2},
				keeper:               keeper,
				keeperRewardPercent:  sdk.MustNewDecFromStr("0.05"),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(1000*JOLT_CF))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF)), sdk.NewCoin("bnb", sdk.NewInt(100*BNB_CF)), sdk.NewCoin("btc", sdk.NewInt(100*BTCB_CF))),
				initialKeeperCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(50*JOLT_CF)), sdk.NewCoin("bnb", sdk.NewInt(10*BNB_CF)), sdk.NewCoin("btc", sdk.NewInt(1*BTCB_CF))), // $100 + $100 + $100 = $300 * 0.8 = $240 borrowable                                                                                                                                       // $100 * 0.8 = $80 borrowable
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(120*JOLT_CF))),                                                                                      // $240 borrowed
				liquidateAfter:       oneMonthDur,
				expectedTotalSuppliedCoins: sdk.NewCoins(
					sdk.NewInt64Coin("ujolt", 1000101456),
				),
				expectedTotalBorrowedCoins: nil,
				expectedKeeperCoins:        sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(102500253)), sdk.NewCoin("bnb", sdk.NewInt(0.5*BNB_CF)), sdk.NewCoin("btc", sdk.NewInt(0.05*BTCB_CF))), // 5% of each seized coin + initial balances
				expectedBorrowerCoins:      sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(170.000001*JOLT_CF)), sdk.NewCoin("bnb", sdk.NewInt(90*BNB_CF)), sdk.NewCoin("btc", sdk.NewInt(99*BTCB_CF))),
				expectedAuctions: []types4.Auction{
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              1,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("bnb", 950000000),
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("ujolt", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("ujolt", 40036023),
						LotReturns:        lotReturns,
					},
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              2,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("btc", 95000000),
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("ujolt", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("ujolt", 40036023),
						LotReturns:        lotReturns,
					},
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              3,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("ujolt", 47504818),
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("ujolt", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("ujolt", 40040087),
						LotReturns:        lotReturns,
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"valid: multiple stablecoin deposits, multiple variable coin borrows",
			// Auctions: total lot value = $285 ($300 of deposits - $15 keeper reward), total max bid value = $270
			args{
				borrower:             []sdk.AccAddress{borrower1, borrower2},
				keeper:               keeper,
				keeperRewardPercent:  sdk.MustNewDecFromStr("0.05"),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("bnb", sdk.NewInt(1000*BNB_CF)), sdk.NewCoin("btc", sdk.NewInt(1000*BTCB_CF))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF)), sdk.NewCoin("usdc", sdk.NewInt(100*JOLT_CF)), sdk.NewCoin("usdt", sdk.NewInt(100*JOLT_CF)), sdk.NewCoin("usdx", sdk.NewInt(100*JOLT_CF))),
				initialKeeperCoins:   sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("usdc", sdk.NewInt(100*JOLT_CF)), sdk.NewCoin("usdt", sdk.NewInt(100*JOLT_CF)), sdk.NewCoin("usdx", sdk.NewInt(100*JOLT_CF))), // $100 + $100 + $100 = $300 * 0.9 = $270 borrowable
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(35*JOLT_CF)), sdk.NewCoin("bnb", sdk.NewInt(10*BNB_CF)), sdk.NewCoin("btc", sdk.NewInt(1*BTCB_CF))),       // $270 borrowed
				liquidateAfter:       oneMonthDur,
				expectedTotalSuppliedCoins: sdk.NewCoins(
					sdk.NewInt64Coin("bnb", 100000078047),
					sdk.NewInt64Coin("btc", 100000000780),
					sdk.NewInt64Coin("ujolt", 1000009550),
					sdk.NewInt64Coin("usdx", 1),
				),
				expectedTotalBorrowedCoins: nil,
				expectedKeeperCoins:        sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(100*JOLT_CF)), sdk.NewCoin("usdc", sdk.NewInt(5*JOLT_CF)), sdk.NewCoin("usdt", sdk.NewInt(5*JOLT_CF)), sdk.NewCoin("usdx", sdk.NewInt(5*JOLT_CF))), // 5% of each seized coin + initial balances
				expectedBorrowerCoins:      sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(135*JOLT_CF)), sdk.NewCoin("bnb", sdk.NewInt(10*BNB_CF)), sdk.NewCoin("btc", sdk.NewInt(1*BTCB_CF)), sdk.NewCoin("usdx", sdk.NewInt(0.000001*JOLT_CF))),
				expectedAuctions: []types4.Auction{
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              1,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("usdc", 95000000), // $95.00
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("bnb", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("bnb", 900097134), // $90.00
						LotReturns:        lotReturns,
					},
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              2,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("usdt", 10552835), // $10.55
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("bnb", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("bnb", 99985020), // $10.00
						LotReturns:        lotReturns,
					},
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              3,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("usdt", 84447165), // $84.45
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("btc", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("btc", 80011211), // $80.01
						LotReturns:        lotReturns,
					},
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              4,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("usdx", 21097866), // $21.10
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("btc", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("btc", 19989610), // $19.99
						LotReturns:        lotReturns,
					},
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              5,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("usdx", 73902133), //$73.90
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("ujolt", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("ujolt", 35010052), // $70.02
						LotReturns:        lotReturns,
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"valid: multiple stablecoin deposits, multiple stablecoin borrows",
			args{
				borrower:             []sdk.AccAddress{borrower1, borrower2},
				keeper:               keeper,
				keeperRewardPercent:  sdk.MustNewDecFromStr("0.05"),
				initialModuleCoins:   sdk.NewCoins(sdk.NewCoin("usdx", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("usdt", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("dai", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("usdc", sdk.NewInt(1000*JOLT_CF))),
				initialBorrowerCoins: sdk.NewCoins(sdk.NewCoin("usdx", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("usdt", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("dai", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("usdc", sdk.NewInt(1000*JOLT_CF))),
				initialKeeperCoins:   sdk.NewCoins(sdk.NewCoin("usdx", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("usdt", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("dai", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("usdc", sdk.NewInt(1000*JOLT_CF))),
				depositCoins:         sdk.NewCoins(sdk.NewCoin("dai", sdk.NewInt(350*JOLT_CF)), sdk.NewCoin("usdc", sdk.NewInt(200*JOLT_CF))),
				borrowCoins:          sdk.NewCoins(sdk.NewCoin("usdt", sdk.NewInt(250*JOLT_CF)), sdk.NewCoin("usdx", sdk.NewInt(245*JOLT_CF))),
				liquidateAfter:       oneMonthDur,
				expectedTotalSuppliedCoins: sdk.NewCoins(
					sdk.NewInt64Coin("dai", 1000000000),
					sdk.NewInt64Coin("usdc", 1000000001),
					sdk.NewInt64Coin("usdt", 1000482503),
					sdk.NewInt64Coin("usdx", 1000463500),
				),
				expectedTotalBorrowedCoins: nil,
				expectedKeeperCoins:        sdk.NewCoins(sdk.NewCoin("dai", sdk.NewInt(1017.50*JOLT_CF)), sdk.NewCoin("usdt", sdk.NewInt(1000*JOLT_CF)), sdk.NewCoin("usdc", sdk.NewInt(1010*JOLT_CF)), sdk.NewCoin("usdx", sdk.NewInt(1000*JOLT_CF))),
				expectedBorrowerCoins:      sdk.NewCoins(sdk.NewCoin("dai", sdk.NewInt(650*JOLT_CF)), sdk.NewCoin("usdc", sdk.NewInt(800000001)), sdk.NewCoin("usdt", sdk.NewInt(1250*JOLT_CF)), sdk.NewCoin("usdx", sdk.NewInt(1245*JOLT_CF))),
				expectedAuctions: []types4.Auction{
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              1,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("dai", 263894126),
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("usdt", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("usdt", 250507897),
						LotReturns:        lotReturns,
					},
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              2,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("dai", 68605874),
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("usdx", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("usdx", 65125788),
						LotReturns:        lotReturns,
					},
					&types4.CollateralAuction{
						BaseAuction: types4.BaseAuction{
							ID:              3,
							Initiator:       "jolt",
							Lot:             sdk.NewInt64Coin("usdc", 189999999),
							Bidder:          sdk.AccAddress(nil),
							Bid:             sdk.NewInt64Coin("usdx", 0),
							HasReceivedBids: false,
							EndTime:         endTime,
							MaxEndTime:      endTime,
						},
						CorrespondingDebt: sdk.NewInt64Coin("debt", 0),
						MaxBid:            sdk.NewInt64Coin("usdx", 180362106),
						LotReturns:        lotReturns,
					},
				},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
	}

	for i, tc := range testCases {
		suite.Run(tc.name, func() {
			// Initialize test app and set context
			tApp := app.NewTestApp()
			ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC)})

			// account which will deposit "initial module account coins"
			depositor1 := sdk.AccAddress(crypto.AddressHash([]byte("testdepositor1")))

			// Auth module genesis state
			authGS := app.NewFundedGenStateWithCoins(
				tApp.AppCodec(),
				[]sdk.Coins{
					tc.args.initialBorrowerCoins,
					tc.args.initialBorrowerCoins,
					tc.args.initialKeeperCoins,
					tc.args.initialModuleCoins,
				},
				[]sdk.AccAddress{
					tc.args.borrower[0],
					tc.args.borrower[1],
					tc.args.keeper,
					depositor1,
				},
			)

			// Hard module genesis state
			joltGS := types3.NewGenesisState(types3.NewParams(
				types3.MoneyMarkets{
					types3.NewMoneyMarket("usdx",
						types3.NewBorrowLimit(false, sdk.NewDec(100000000*JOLT_CF), sdk.MustNewDecFromStr("0.9")), // Borrow Limit
						"usdx:usd",                   // Market ID
						sdk.NewInt(JOLT_CF),          // Conversion Factor
						model,                        // Interest Rate Model
						reserveFactor,                // Reserve Factor
						tc.args.keeperRewardPercent), // Keeper Reward Percent
					types3.NewMoneyMarket("usdt",
						types3.NewBorrowLimit(false, sdk.NewDec(100000000*JOLT_CF), sdk.MustNewDecFromStr("0.9")), // Borrow Limit
						"usdt:usd",                   // Market ID
						sdk.NewInt(JOLT_CF),          // Conversion Factor
						model,                        // Interest Rate Model
						reserveFactor,                // Reserve Factor
						tc.args.keeperRewardPercent), // Keeper Reward Percent
					types3.NewMoneyMarket("usdc",
						types3.NewBorrowLimit(false, sdk.NewDec(100000000*JOLT_CF), sdk.MustNewDecFromStr("0.9")), // Borrow Limit
						"usdc:usd",                   // Market ID
						sdk.NewInt(JOLT_CF),          // Conversion Factor
						model,                        // Interest Rate Model
						reserveFactor,                // Reserve Factor
						tc.args.keeperRewardPercent), // Keeper Reward Percent
					types3.NewMoneyMarket("dai",
						types3.NewBorrowLimit(false, sdk.NewDec(100000000*JOLT_CF), sdk.MustNewDecFromStr("0.9")), // Borrow Limit
						"dai:usd",                    // Market ID
						sdk.NewInt(JOLT_CF),          // Conversion Factor
						model,                        // Interest Rate Model
						reserveFactor,                // Reserve Factor
						tc.args.keeperRewardPercent), // Keeper Reward Percent
					types3.NewMoneyMarket("ujolt",
						types3.NewBorrowLimit(false, sdk.NewDec(100000000*JOLT_CF), sdk.MustNewDecFromStr("0.8")), // Borrow Limit
						"joltify:usd",                // Market ID
						sdk.NewInt(JOLT_CF),          // Conversion Factor
						model,                        // Interest Rate Model
						reserveFactor,                // Reserve Factor
						tc.args.keeperRewardPercent), // Keeper Reward Percent
					types3.NewMoneyMarket("bnb",
						types3.NewBorrowLimit(false, sdk.NewDec(100000000*BNB_CF), sdk.MustNewDecFromStr("0.8")), // Borrow Limit
						"bnb:usd",                    // Market ID
						sdk.NewInt(BNB_CF),           // Conversion Factor
						model,                        // Interest Rate Model
						reserveFactor,                // Reserve Factor
						tc.args.keeperRewardPercent), // Keeper Reward Percent
					types3.NewMoneyMarket("btc",
						types3.NewBorrowLimit(false, sdk.NewDec(100000000*BTCB_CF), sdk.MustNewDecFromStr("0.8")), // Borrow Limit
						"btc:usd",                    // Market ID
						sdk.NewInt(BTCB_CF),          // Conversion Factor
						model,                        // Interest Rate Model
						reserveFactor,                // Reserve Factor
						tc.args.keeperRewardPercent), // Keeper Reward Percent
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
						{MarketID: "usdt:usd", BaseAsset: "usdt", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "usdc:usd", BaseAsset: "usdc", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "dai:usd", BaseAsset: "dai", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "joltify:usd", BaseAsset: "joltify", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "bnb:usd", BaseAsset: "bnb", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "btc:usd", BaseAsset: "btc", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
					},
				},
				PostedPrices: []types2.PostedPrice{
					{
						MarketID:      "usdx:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdk.MustNewDecFromStr("1.00"),
						Expiry:        time.Now().Add(100 * time.Hour),
					},
					{
						MarketID:      "usdt:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdk.MustNewDecFromStr("1.00"),
						Expiry:        time.Now().Add(100 * time.Hour),
					},
					{
						MarketID:      "usdc:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdk.MustNewDecFromStr("1.00"),
						Expiry:        time.Now().Add(100 * time.Hour),
					},
					{
						MarketID:      "dai:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdk.MustNewDecFromStr("1.00"),
						Expiry:        time.Now().Add(100 * time.Hour),
					},
					{
						MarketID:      "joltify:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdk.MustNewDecFromStr("2.00"),
						Expiry:        time.Now().Add(100 * time.Hour),
					},
					{
						MarketID:      "bnb:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdk.MustNewDecFromStr("10.00"),
						Expiry:        time.Now().Add(100 * time.Hour),
					},
					{
						MarketID:      "btc:usd",
						OracleAddress: sdk.AccAddress{},
						Price:         sdk.MustNewDecFromStr("100.00"),
						Expiry:        time.Now().Add(100 * time.Hour),
					},
				},
			}

			// Initialize test application
			tApp.InitializeFromGenesisStates(authGS,
				app.GenesisState{types2.ModuleName: tApp.AppCodec().MustMarshalJSON(&pricefeedGS)},
				app.GenesisState{types3.ModuleName: tApp.AppCodec().MustMarshalJSON(&joltGS)})

			keeper := tApp.GetJoltKeeper()
			suite.app = tApp
			suite.ctx = ctx
			suite.keeper = keeper
			suite.auctionKeeper = tApp.GetAuctionKeeper()

			var err error

			// Run begin blocker to set up state
			jolt.BeginBlocker(suite.ctx, suite.keeper)

			// Deposit initial module account coins
			err = suite.keeper.Deposit(suite.ctx, depositor1, tc.args.initialModuleCoins)
			suite.Require().NoError(err)

			oneLiquidate := false
			for j, el := range tc.args.borrower {
				// Deposit coins
				err = suite.keeper.Deposit(suite.ctx, el, tc.args.depositCoins)
				suite.Require().NoError(err)

				if j == 1 && i == 0 {
					// Borrow coins
					oneLiquidate = true
					borrowCoinsSmall := sdk.NewCoins(sdk.NewCoin("ujolt", sdk.NewInt(5*JOLT_CF)))
					err = suite.keeper.Borrow(suite.ctx, el, borrowCoinsSmall)
					suite.Require().NoError(err)
				} else {
					// Borrow coins
					err = suite.keeper.Borrow(suite.ctx, el, tc.args.borrowCoins)
					suite.Require().NoError(err)
				}
			}
			// Set up liquidation chain context and run begin blocker
			runAtTime := suite.ctx.BlockTime().Add(tc.args.liquidateAfter)
			liqCtx := suite.ctx.WithBlockTime(runAtTime)
			jolt.BeginBlocker(liqCtx, suite.keeper)

			for _, el := range tc.args.borrower {
				// Check borrow exists before liquidation
				_, foundBorrowBefore := suite.keeper.GetBorrow(liqCtx, el)
				suite.Require().True(foundBorrowBefore)
				// Check that the user's deposit exists before liquidation
				_, foundDepositBefore := suite.keeper.GetDeposit(liqCtx, el)
				suite.Require().True(foundDepositBefore)
			}
			queryServer := joltKeeper.NewQueryServerImpl(suite.keeper, suite.app.GetAccountKeeper(), suite.app.GetBankKeeper())
			req := types3.QueryLiquidateRequest{}
			ret, err := queryServer.Liquidate(sdk.WrapSDKContext(suite.ctx), &req)
			suite.Require().NoError(err)
			if oneLiquidate {
				suite.Require().Equal(1, len(ret.LiquidateItems))
			} else {
				suite.Require().Equal(2, len(ret.LiquidateItems))
			}
			result := sdk.MustNewDecFromStr(ret.LiquidateItems[0].Ltv)
			suite.Require().True(result.GTE(sdk.MustNewDecFromStr("1.0")))
		})
	}
}
