package cdp_test

import (
	"fmt"
	tmlog "github.com/tendermint/tendermint/libs/log"
	"sort"
	"strings"
	"testing"
	"time"

	cdp2 "github.com/joltify-finance/joltify_lending/x/third_party/cdp"
	"github.com/joltify-finance/joltify_lending/x/third_party/cdp/keeper"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/cdp/types"

	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/app"
)

type GenesisTestSuite struct {
	suite.Suite

	app     app.TestApp
	ctx     sdk.Context
	genTime time.Time
	keeper  keeper.Keeper
	addrs   []sdk.AccAddress
}

func (suite *GenesisTestSuite) SetupTest() {
	lg := tmlog.TestingLogger()
	tApp := app.NewTestApp(lg, suite.T().TempDir())
	suite.genTime = tmtime.Canonical(time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC))
	suite.ctx = tApp.NewContext(true, tmproto.Header{Height: 1, Time: suite.genTime})
	suite.keeper = tApp.GetCDPKeeper()
	suite.app = tApp

	_, addrs := app.GeneratePrivKeyAddressPairs(3)
	suite.addrs = addrs
}

func (suite *GenesisTestSuite) TestInvalidGenState() {
	type args struct {
		params             types2.Params
		cdps               types2.CDPs
		deposits           types2.Deposits
		startingID         uint64
		debtDenom          string
		govDenom           string
		genAccumTimes      types2.GenesisAccumulationTimes
		genTotalPrincipals types2.GenesisTotalPrincipals
	}
	type errArgs struct {
		expectPass bool
		contains   string
	}

	testCases := []struct {
		name    string
		args    args
		errArgs errArgs
	}{
		{
			name: "empty debt denom",
			args: args{
				params:             types2.DefaultParams(),
				cdps:               types2.CDPs{},
				deposits:           types2.Deposits{},
				debtDenom:          "",
				govDenom:           types2.DefaultGovDenom,
				genAccumTimes:      types2.DefaultGenesisState().PreviousAccumulationTimes,
				genTotalPrincipals: types2.DefaultGenesisState().TotalPrincipals,
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "debt denom invalid",
			},
		},
		{
			name: "empty gov denom",
			args: args{
				params:             types2.DefaultParams(),
				cdps:               types2.CDPs{},
				deposits:           types2.Deposits{},
				debtDenom:          types2.DefaultDebtDenom,
				govDenom:           "",
				genAccumTimes:      types2.DefaultGenesisState().PreviousAccumulationTimes,
				genTotalPrincipals: types2.DefaultGenesisState().TotalPrincipals,
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "gov denom invalid",
			},
		},
		{
			name: "interest factor below one",
			args: args{
				params:             types2.DefaultParams(),
				cdps:               types2.CDPs{},
				deposits:           types2.Deposits{},
				debtDenom:          types2.DefaultDebtDenom,
				govDenom:           types2.DefaultGovDenom,
				genAccumTimes:      types2.GenesisAccumulationTimes{types2.NewGenesisAccumulationTime("bnb-a", time.Time{}, sdk.OneDec().Sub(sdk.SmallestDec()))},
				genTotalPrincipals: types2.DefaultGenesisState().TotalPrincipals,
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "interest factor should be ≥ 1.0",
			},
		},
		{
			name: "negative total principal",
			args: args{
				params:             types2.DefaultParams(),
				cdps:               types2.CDPs{},
				deposits:           types2.Deposits{},
				debtDenom:          types2.DefaultDebtDenom,
				govDenom:           types2.DefaultGovDenom,
				genAccumTimes:      types2.DefaultGenesisState().PreviousAccumulationTimes,
				genTotalPrincipals: types2.GenesisTotalPrincipals{types2.NewGenesisTotalPrincipal("bnb-a", sdk.NewInt(-1))},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "total principal should be positive",
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			gs := types2.NewGenesisState(tc.args.params, tc.args.cdps, tc.args.deposits, tc.args.startingID,
				tc.args.debtDenom, tc.args.govDenom, tc.args.genAccumTimes, tc.args.genTotalPrincipals)
			err := gs.Validate()
			if tc.errArgs.expectPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().True(strings.Contains(err.Error(), tc.errArgs.contains))
			}
		})
	}
}

func (suite *GenesisTestSuite) TestValidGenState() {
	cdc := suite.app.AppCodec()

	suite.NotPanics(func() {
		suite.app.InitializeFromGenesisStates(nil, nil,
			NewPricefeedGenStateMulti(cdc),
			NewCDPGenStateMulti(cdc),
		)
	})

	cdpGS := NewCDPGenStateMulti(cdc)
	gs := types2.GenesisState{}
	suite.app.AppCodec().MustUnmarshalJSON(cdpGS["cdp"], &gs)
	gs.CDPs = cdps()
	gs.StartingCdpID = uint64(5)
	appGS := app.GenesisState{"cdp": suite.app.AppCodec().MustMarshalJSON(&gs)}
	suite.NotPanics(func() {
		suite.SetupTest()
		suite.app.InitializeFromGenesisStates(nil, nil,
			NewPricefeedGenStateMulti(cdc),
			appGS,
		)
	})
}

func (suite *GenesisTestSuite) Test_InitExportGenesis() {
	cdps := types2.CDPs{
		{
			ID:              2,
			Owner:           suite.addrs[0],
			Type:            "xrp-a",
			Collateral:      c("xrp", 200000000),
			Principal:       c("usdx", 10000000),
			AccumulatedFees: c("usdx", 0),
			FeesUpdated:     suite.genTime,
			InterestFactor:  sdk.NewDec(1),
		},
	}

	genTotalPrincipals := types2.GenesisTotalPrincipals{
		types2.NewGenesisTotalPrincipal("btc-a", sdk.ZeroInt()),
		types2.NewGenesisTotalPrincipal("xrp-a", sdk.ZeroInt()),
	}

	var deposits types2.Deposits
	for _, c := range cdps {
		deposit := types2.Deposit{
			CdpID:     c.ID,
			Depositor: c.Owner,
			Amount:    c.Collateral,
		}
		deposits = append(deposits, deposit)

		for i, p := range genTotalPrincipals {
			if p.CollateralType == c.Type {
				genTotalPrincipals[i].TotalPrincipal = genTotalPrincipals[i].TotalPrincipal.Add(c.Principal.Amount)
			}
		}
	}

	cdpGenesis := types2.GenesisState{
		Params: types2.Params{
			GlobalDebtLimit:         sdk.NewInt64Coin("usdx", 1000000000000),
			SurplusAuctionThreshold: types2.DefaultSurplusThreshold,
			SurplusAuctionLot:       types2.DefaultSurplusLot,
			DebtAuctionThreshold:    types2.DefaultDebtThreshold,
			DebtAuctionLot:          types2.DefaultDebtLot,
			CollateralParams: types2.CollateralParams{
				{
					Denom:                            "xrp",
					Type:                             "xrp-a",
					LiquidationRatio:                 sdk.MustNewDecFromStr("2.0"),
					DebtLimit:                        sdk.NewInt64Coin("usdx", 500000000000),
					StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"), // 5% apr
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
					DebtLimit:                        sdk.NewInt64Coin("usdx", 500000000000),
					StabilityFee:                     sdk.MustNewDecFromStr("1.000000000782997609"), // 2.5% apr
					LiquidationPenalty:               d("0.025"),
					AuctionSize:                      i(10000000),
					SpotMarketID:                     "btc:usd",
					LiquidationMarketID:              "btc:usd",
					KeeperRewardPercentage:           d("0.01"),
					CheckCollateralizationIndexCount: i(10),
					ConversionFactor:                 i(8),
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
		CDPs:          cdps,
		Deposits:      deposits,
		PreviousAccumulationTimes: types2.GenesisAccumulationTimes{
			types2.NewGenesisAccumulationTime("btc-a", suite.genTime, sdk.OneDec()),
			types2.NewGenesisAccumulationTime("xrp-a", suite.genTime, sdk.OneDec()),
		},
		TotalPrincipals: genTotalPrincipals,
	}

	suite.NotPanics(func() {
		suite.app.InitializeFromGenesisStatesWithTime(
			suite.genTime, nil, nil,
			NewPricefeedGenStateMulti(suite.app.AppCodec()),
			app.GenesisState{types2.ModuleName: suite.app.AppCodec().MustMarshalJSON(&cdpGenesis)},
		)
	})

	// We run the BeginBlocker at time.Now() to accumulate interest
	suite.ctx = suite.ctx.WithBlockTime(time.Now())
	cdp2.BeginBlocker(suite.ctx, abci.RequestBeginBlock{}, suite.keeper)

	expectedGenesis := cdpGenesis

	// Update previous accrual times in expected genesis
	var expectedPrevAccTimes types2.GenesisAccumulationTimes
	for _, prevAccTime := range cdpGenesis.PreviousAccumulationTimes {
		time, found := suite.keeper.GetPreviousAccrualTime(suite.ctx, prevAccTime.CollateralType)
		if !found {
			panic(fmt.Sprintf("couldn't find previous accrual time for %s", prevAccTime.CollateralType))
		}
		prevAccTime.PreviousAccumulationTime = time

		interestFactor, found := suite.keeper.GetInterestFactor(suite.ctx, prevAccTime.CollateralType)
		if !found {
			panic(fmt.Sprintf("couldn't find interest factor for %s", prevAccTime.CollateralType))
		}
		prevAccTime.InterestFactor = interestFactor

		expectedPrevAccTimes = append(expectedPrevAccTimes, prevAccTime)
	}
	expectedGenesis.PreviousAccumulationTimes = expectedPrevAccTimes

	// Update total principals
	var totalPrincipals types2.GenesisTotalPrincipals
	for _, p := range expectedGenesis.TotalPrincipals {
		totalPrincipal := suite.keeper.GetTotalPrincipal(suite.ctx, p.CollateralType, "usdx")
		p.TotalPrincipal = totalPrincipal
		totalPrincipals = append(totalPrincipals, p)
	}
	expectedGenesis.TotalPrincipals = totalPrincipals

	// Update CDPs
	expectedGenesis.CDPs = suite.keeper.GetAllCdps(suite.ctx)

	exportedGenesis := cdp2.ExportGenesis(suite.ctx, suite.keeper)

	// Sort TotalPrincipals in both genesis files so slice order matches
	sort.SliceStable(expectedGenesis.TotalPrincipals, func(i, j int) bool {
		return expectedGenesis.TotalPrincipals[i].CollateralType < expectedGenesis.TotalPrincipals[j].CollateralType
	})
	sort.SliceStable(exportedGenesis.TotalPrincipals, func(i, j int) bool {
		return exportedGenesis.TotalPrincipals[i].CollateralType < exportedGenesis.TotalPrincipals[j].CollateralType
	})

	// Sort PreviousAccumulationTimes in both genesis files so slice order matches
	sort.SliceStable(expectedGenesis.PreviousAccumulationTimes, func(i, j int) bool {
		return expectedGenesis.PreviousAccumulationTimes[i].CollateralType < expectedGenesis.PreviousAccumulationTimes[j].CollateralType
	})
	sort.SliceStable(exportedGenesis.PreviousAccumulationTimes, func(i, j int) bool {
		return exportedGenesis.PreviousAccumulationTimes[i].CollateralType < exportedGenesis.PreviousAccumulationTimes[j].CollateralType
	})

	suite.Equal(expectedGenesis, exportedGenesis)
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}
