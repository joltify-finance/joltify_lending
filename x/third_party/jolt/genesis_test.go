package jolt_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"

	"github.com/joltify-finance/joltify_lending/x/third_party/jolt"
	"github.com/joltify-finance/joltify_lending/x/third_party/jolt/keeper"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	"github.com/stretchr/testify/suite"

	tmtime "github.com/cometbft/cometbft/types/time"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/app"
)

type GenesisTestSuite struct {
	suite.Suite

	app     app.TestApp
	genTime time.Time
	ctx     context.Context
	keeper  keeper.Keeper
	addrs   []sdk.AccAddress
}

func (suite *GenesisTestSuite) SetupTest() {
	tApp := app.NewTestApp(log.NewTestLogger(suite.T()), suite.T().TempDir())
	suite.genTime = tmtime.Canonical(time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC))
	suite.ctx = tApp.Ctx
	suite.keeper = tApp.GetJoltKeeper()
	suite.app = tApp

	_, addrs := app.GeneratePrivKeyAddressPairs(3)
	suite.addrs = addrs
}

func (suite *GenesisTestSuite) Test_InitExportGenesis() {
	loanToValue, _ := sdkmath.LegacyNewDecFromStr("0.6")
	params := types2.NewParams(
		types2.MoneyMarkets{
			types2.NewMoneyMarket(
				"ujolt",
				types2.NewBorrowLimit(
					false,
					sdkmath.LegacyNewDec(1e15),
					loanToValue,
				),
				"joltify:usd",
				sdkmath.NewInt(1e6),
				types2.NewInterestRateModel(
					sdkmath.LegacyMustNewDecFromStr("0.05"),
					sdkmath.LegacyMustNewDecFromStr("2"),
					sdkmath.LegacyMustNewDecFromStr("0.8"),
					sdkmath.LegacyMustNewDecFromStr("10"),
				),
				sdkmath.LegacyMustNewDecFromStr("0.05"),
				sdkmath.LegacyZeroDec(),
			),
		},
		sdkmath.LegacyNewDec(10),
	)

	deposits := types2.Deposits{
		types2.NewDeposit(
			suite.addrs[0],
			sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1e8))), // 100 ujolt
			types2.SupplyInterestFactors{
				{
					Denom: "ujolt",
					Value: sdkmath.LegacyNewDec(1),
				},
			},
		),
	}

	var totalSupplied sdk.Coins
	for _, deposit := range deposits {
		totalSupplied = totalSupplied.Add(deposit.Amount...)
	}

	borrows := types2.Borrows{
		types2.NewBorrow(
			suite.addrs[1],
			sdk.NewCoins(sdk.NewCoin("ujolt", sdkmath.NewInt(1e7))), // 10 ujolt
			types2.BorrowInterestFactors{
				{
					Denom: "ujolt",
					Value: sdkmath.LegacyNewDec(1),
				},
			},
		),
	}

	var totalBorrowed sdk.Coins
	for _, borrow := range borrows {
		totalBorrowed = totalBorrowed.Add(borrow.Amount...)
	}

	supplyInterestFactor := sdkmath.LegacyMustNewDecFromStr("1.0001")
	borrowInterestFactor := sdkmath.LegacyMustNewDecFromStr("1.1234")
	accuralTimes := types2.GenesisAccumulationTimes{
		types2.NewGenesisAccumulationTime("ujolt", suite.genTime, supplyInterestFactor, borrowInterestFactor),
	}

	joltGenesis := types2.NewGenesisState(
		params,
		accuralTimes,
		deposits,
		borrows,
		totalSupplied,
		totalBorrowed,
		sdk.Coins{},
	)

	suite.NotPanics(
		func() {
			suite.app.InitializeFromGenesisStatesWithTime(
				suite.genTime, nil, nil,
				app.GenesisState{types2.ModuleName: suite.app.AppCodec().MustMarshalJSON(&joltGenesis)},
			)
		},
	)

	var expectedDeposits types2.Deposits
	for _, deposit := range deposits {
		// Deposit coin amounts
		var depositAmount sdk.Coins
		for _, coin := range deposit.Amount {
			accrualTime, found := getGenesisAccumulationTime(coin.Denom, accuralTimes)
			if !found {
				panic(fmt.Sprintf("accrual time not found %s", coin.Denom))
			}
			expectedAmt := accrualTime.SupplyInterestFactor.MulInt(coin.Amount).RoundInt()
			depositAmount = depositAmount.Add(sdk.NewCoin(coin.Denom, expectedAmt))
		}
		deposit.Amount = depositAmount
		// Deposit interest factor indexes
		var indexes types2.SupplyInterestFactors
		for _, index := range deposit.Index {
			accrualTime, found := getGenesisAccumulationTime(index.Denom, accuralTimes)
			if !found {
				panic(fmt.Sprintf("accrual time not found %s", index.Denom))
			}
			index.Value = accrualTime.SupplyInterestFactor
			indexes = append(indexes, index)
		}
		deposit.Index = indexes
		expectedDeposits = append(expectedDeposits, deposit)
	}

	var expectedBorrows types2.Borrows
	for _, borrow := range borrows {
		// Borrow coin amounts
		var borrowAmount sdk.Coins
		for _, coin := range borrow.Amount {
			accrualTime, found := getGenesisAccumulationTime(coin.Denom, accuralTimes)
			if !found {
				panic(fmt.Sprintf("accrual time not found %s", coin.Denom))
			}
			expectedAmt := accrualTime.BorrowInterestFactor.MulInt(coin.Amount).RoundInt()
			borrowAmount = borrowAmount.Add(sdk.NewCoin(coin.Denom, expectedAmt))

		}
		borrow.Amount = borrowAmount
		// Borrow interest factor indexes
		var indexes types2.BorrowInterestFactors
		for _, index := range borrow.Index {
			accrualTime, found := getGenesisAccumulationTime(index.Denom, accuralTimes)
			if !found {
				panic(fmt.Sprintf("accrual time not found %s", index.Denom))
			}
			index.Value = accrualTime.BorrowInterestFactor
			indexes = append(indexes, index)
		}
		borrow.Index = indexes
		expectedBorrows = append(expectedBorrows, borrow)
	}

	expectedGenesis := joltGenesis
	expectedGenesis.Deposits = expectedDeposits
	expectedGenesis.Borrows = expectedBorrows
	exportedGenesis := jolt.ExportGenesis(sdk.UnwrapSDKContext(suite.ctx), suite.keeper)
	suite.Equal(expectedGenesis, exportedGenesis)
}

func getGenesisAccumulationTime(denom string, ts types2.GenesisAccumulationTimes) (types2.GenesisAccumulationTime, bool) {
	for _, t := range ts {
		if t.CollateralType == denom {
			return t, true
		}
	}
	return types2.GenesisAccumulationTime{}, false
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}
