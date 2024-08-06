package keeper_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"cosmossdk.io/log"

	"github.com/cosmos/cosmos-sdk/x/bank/testutil"

	keeper2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/keeper"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/stretchr/testify/suite"
)

type grpcQueryTestSuite struct {
	suite.Suite

	tApp        app.TestApp
	ctx         context.Context
	keeper      keeper2.Keeper
	queryServer types2.QueryServer
	addrs       []sdk.AccAddress
}

func (suite *grpcQueryTestSuite) SetupTest() {
	suite.tApp = app.NewTestApp(log.NewTestLogger(suite.T()), suite.T().TempDir())
	_, addrs := app.GeneratePrivKeyAddressPairs(2)

	suite.addrs = addrs

	suite.ctx = suite.tApp.NewContext(true).WithBlockTime(time.Now().UTC())
	suite.keeper = suite.tApp.GetJoltKeeper()
	suite.queryServer = keeper2.NewQueryServerImpl(suite.keeper, suite.tApp.GetAccountKeeper(), suite.tApp.GetBankKeeper())

	err := suite.tApp.FundModuleAccount(
		sdk.UnwrapSDKContext(suite.ctx),
		types2.ModuleAccountName,
		cs(
			c("usdx", 10000000000),
			c("busd", 10000000000),
		),
	)
	suite.Require().NoError(err)

	suite.tApp.InitializeFromGenesisStates(nil, nil,
		NewPricefeedGenStateMulti(suite.tApp.AppCodec()),
		NewJoltGenState(suite.tApp.AppCodec()),
		app.NewFundedGenStateWithSameCoins(
			suite.tApp.AppCodec(),
			cs(
				c("bnb", 10000000000),
				c("busd", 20000000000),
			),
			addrs,
		),
	)
	coins := cs(
		c("bnb", 10000000000),
		c("busd", 20000000000),
	)

	for _, addr := range addrs {
		err = testutil.FundAccount(suite.ctx, suite.tApp.GetBankKeeper(), addr, coins)
		suite.Require().NoError(err)
	}
}

func (suite *grpcQueryTestSuite) TestGrpcQueryParams() {
	res, err := suite.queryServer.Params(suite.ctx, &types2.QueryParamsRequest{})
	suite.Require().NoError(err)

	var expected types2.GenesisState
	defaultHARDState := NewJoltGenState(suite.tApp.AppCodec())
	suite.tApp.AppCodec().MustUnmarshalJSON(defaultHARDState[types2.ModuleName], &expected)

	suite.Equal(expected.Params, res.Params, "params should equal test genesis state")
}

func (suite *grpcQueryTestSuite) TestGrpcQueryAccounts() {
	res, err := suite.queryServer.Accounts(suite.ctx, &types2.QueryAccountsRequest{})
	suite.Require().NoError(err)

	ak := suite.tApp.GetAccountKeeper()
	acc := ak.GetModuleAccount(suite.ctx, types2.ModuleName)

	suite.Len(res.Accounts, 1)
	suite.Equal(acc, &res.Accounts[0], "accounts should include module account")
}

func (suite *grpcQueryTestSuite) TestGrpcQueryDeposits_EmptyResponse() {
	res, err := suite.queryServer.Deposits(suite.ctx, &types2.QueryDepositsRequest{})
	suite.Require().NoError(err)

	fmt.Printf(">>>>>%v\n", res.Pagination)
}

func (suite *grpcQueryTestSuite) addDeposits() {
	deposits := []struct {
		Address sdk.AccAddress
		Coins   sdk.Coins
	}{
		{
			suite.addrs[0],
			cs(c("bnb", 100000000)),
		},
		{
			suite.addrs[1],
			cs(c("bnb", 20000000)),
		},
		{
			suite.addrs[0],
			cs(c("busd", 20000000)),
		},
		{
			suite.addrs[0],
			cs(c("busd", 8000000)),
		},
	}

	for _, dep := range deposits {
		suite.NotPanics(func() {
			err := suite.keeper.Deposit(suite.ctx, dep.Address, dep.Coins)
			suite.Require().NoError(err)
		})
	}
}

func (suite *grpcQueryTestSuite) addBorrows() {
	borrows := []struct {
		Address sdk.AccAddress
		Coins   sdk.Coins
	}{
		{
			suite.addrs[0],
			cs(c("usdx", 10000000)),
		},
		{
			suite.addrs[1],
			cs(c("usdx", 20000000)),
		},
		{
			suite.addrs[0],
			cs(c("usdx", 40000000)),
		},
		{
			suite.addrs[0],
			cs(c("busd", 80000000)),
		},
	}

	for _, dep := range borrows {
		suite.NotPanics(func() {
			err := suite.keeper.Borrow(suite.ctx, dep.Address, dep.Coins)
			suite.Require().NoErrorf(err, "borrow %s should not error", dep.Coins)
		})
	}
}

func (suite *grpcQueryTestSuite) TestGrpcQueryDeposits() {
	suite.addDeposits()

	tests := []struct {
		giveName          string
		giveRequest       *types2.QueryDepositsRequest
		wantDepositCounts int
		shouldError       bool
		errorSubstr       string
	}{
		{
			"empty query",
			&types2.QueryDepositsRequest{},
			2,
			false,
			"",
		},
		{
			"owner",
			&types2.QueryDepositsRequest{
				Owner: suite.addrs[0].String(),
			},
			// Excludes the second address
			1,
			false,
			"",
		},
		{
			"invalid owner",
			&types2.QueryDepositsRequest{
				Owner: "invalid address",
			},
			// No deposits
			0,
			true,
			"decoding bech32 failed",
		},
		{
			"owner and denom",
			&types2.QueryDepositsRequest{
				Owner: suite.addrs[0].String(),
				Denom: "bnb",
			},
			// Only the first one
			1,
			false,
			"",
		},
		{
			"owner and invalid denom empty response",
			&types2.QueryDepositsRequest{
				Owner: suite.addrs[0].String(),
				Denom: "invalid denom",
			},
			0,
			false,
			"",
		},
		{
			"denom",
			&types2.QueryDepositsRequest{
				Denom: "bnb",
			},
			2,
			false,
			"",
		},
	}

	for _, tt := range tests {
		suite.Run(tt.giveName, func() {
			res, err := suite.queryServer.Deposits(suite.ctx, tt.giveRequest)

			if tt.shouldError {
				suite.Error(err)
				suite.Contains(err.Error(), tt.errorSubstr)
			} else {
				suite.NoError(err)
				suite.Equal(tt.wantDepositCounts, len(res.Deposits))
			}
		})

		// Unsynced deposits should be the same
		suite.Run(tt.giveName+"_unsynced", func() {
			res, err := suite.queryServer.UnsyncedDeposits(suite.ctx, &types2.QueryUnsyncedDepositsRequest{
				Denom:      tt.giveRequest.Denom,
				Owner:      tt.giveRequest.Owner,
				Pagination: tt.giveRequest.Pagination,
			})

			if tt.shouldError {
				suite.Error(err)
				suite.Contains(err.Error(), tt.errorSubstr)
			} else {
				suite.NoError(err)
				suite.Equal(tt.wantDepositCounts, len(res.Deposits))
			}
		})
	}
}

func (suite *grpcQueryTestSuite) TestGrpcQueryBorrows() {
	suite.addDeposits()
	suite.addBorrows()

	tests := []struct {
		giveName          string
		giveRequest       *types2.QueryBorrowsRequest
		wantDepositCounts int
		shouldError       bool
		errorSubstr       string
	}{
		{
			"empty query",
			&types2.QueryBorrowsRequest{},
			2,
			false,
			"",
		},
		{
			"owner",
			&types2.QueryBorrowsRequest{
				Owner: suite.addrs[0].String(),
			},
			// Excludes the second address
			1,
			false,
			"",
		},
		{
			"invalid owner",
			&types2.QueryBorrowsRequest{
				Owner: "invalid address",
			},
			// No deposits
			0,
			true,
			"decoding bech32 failed",
		},
		{
			"owner and denom",
			&types2.QueryBorrowsRequest{
				Owner: suite.addrs[0].String(),
				Denom: "usdx",
			},
			// Only the first one
			1,
			false,
			"",
		},
		{
			"owner and invalid denom empty response",
			&types2.QueryBorrowsRequest{
				Owner: suite.addrs[0].String(),
				Denom: "invalid denom",
			},
			0,
			false,
			"",
		},
		{
			"denom",
			&types2.QueryBorrowsRequest{
				Denom: "usdx",
			},
			2,
			false,
			"",
		},
	}

	for _, tt := range tests {
		suite.Run(tt.giveName, func() {
			res, err := suite.queryServer.Borrows(suite.ctx, tt.giveRequest)

			if tt.shouldError {
				suite.Error(err)
				suite.Contains(err.Error(), tt.errorSubstr)
			} else {
				suite.NoError(err)
				suite.Equal(tt.wantDepositCounts, len(res.Borrows))
			}
		})

		// Unsynced deposits should be the same
		suite.Run(tt.giveName+"_unsynced", func() {
			res, err := suite.queryServer.UnsyncedBorrows(suite.ctx, &types2.QueryUnsyncedBorrowsRequest{
				Denom:      tt.giveRequest.Denom,
				Owner:      tt.giveRequest.Owner,
				Pagination: tt.giveRequest.Pagination,
			})

			if tt.shouldError {
				suite.Error(err)
				suite.Contains(err.Error(), tt.errorSubstr)
			} else {
				suite.NoError(err)
				suite.Equal(tt.wantDepositCounts, len(res.Borrows))
			}
		})
	}
}

func (suite *grpcQueryTestSuite) TestGrpcQueryTotalDeposited() {
	suite.addDeposits()

	totalDeposited, err := suite.queryServer.TotalDeposited(suite.ctx, &types2.QueryTotalDepositedRequest{})
	suite.Require().NoError(err)

	suite.Equal(&types2.QueryTotalDepositedResponse{
		SuppliedCoins: cs(
			c("bnb", 100000000+20000000),
			c("busd", 20000000+8000000),
		),
	}, totalDeposited)
}

func (suite *grpcQueryTestSuite) TestGrpcQueryTotalDeposited_Denom() {
	suite.addDeposits()

	totalDeposited, err := suite.queryServer.TotalDeposited(suite.ctx, &types2.QueryTotalDepositedRequest{
		Denom: "bnb",
	})
	suite.Require().NoError(err)

	suite.Equal(&types2.QueryTotalDepositedResponse{
		SuppliedCoins: cs(
			c("bnb", 100000000+20000000),
		),
	}, totalDeposited)
}

func (suite *grpcQueryTestSuite) TestGrpcQueryTotalBorrowed() {
	suite.addDeposits()
	suite.addBorrows()

	totalDeposited, err := suite.queryServer.TotalBorrowed(suite.ctx, &types2.QueryTotalBorrowedRequest{})
	suite.Require().NoError(err)

	suite.Equal(&types2.QueryTotalBorrowedResponse{
		BorrowedCoins: cs(
			c("usdx", 10000000+20000000+40000000),
			c("busd", 80000000),
		),
	}, totalDeposited)
}

func (suite *grpcQueryTestSuite) TestGrpcQueryTotalBorrowed_denom() {
	suite.addDeposits()
	suite.addBorrows()

	totalDeposited, err := suite.queryServer.TotalBorrowed(suite.ctx, &types2.QueryTotalBorrowedRequest{
		Denom: "usdx",
	})
	suite.Require().NoError(err)

	suite.Equal(&types2.QueryTotalBorrowedResponse{
		BorrowedCoins: cs(
			c("usdx", 10000000+20000000+40000000),
		),
	}, totalDeposited)
}

func (suite *grpcQueryTestSuite) TestGrpcQueryInterestRate() {
	tests := []struct {
		giveName          string
		giveDenom         string
		wantInterestRates types2.MoneyMarketInterestRates
		shouldError       bool
	}{
		{
			"no denom",
			"",
			types2.MoneyMarketInterestRates{
				{
					Denom:              "usdx",
					SupplyInterestRate: "0.000000000000000000",
					BorrowInterestRate: "0.050000000000000000",
				},
				{
					Denom:              "bnb",
					SupplyInterestRate: "0.000000000000000000",
					BorrowInterestRate: "0.000000000000000000",
				},
				{
					Denom:              "busd",
					SupplyInterestRate: "0.000000000000000000",
					BorrowInterestRate: "0.000000000000000000",
				},
			},
			false,
		},
		{
			"denom",
			"usdx",
			types2.MoneyMarketInterestRates{
				{
					Denom:              "usdx",
					SupplyInterestRate: "0.000000000000000000",
					BorrowInterestRate: "0.050000000000000000",
				},
			},
			false,
		},
		{
			"invalid denom",
			"bun",
			types2.MoneyMarketInterestRates{},
			true,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.giveName, func() {
			res, err := suite.queryServer.InterestRate(suite.ctx, &types2.QueryInterestRateRequest{
				Denom: tt.giveDenom,
			})

			if tt.shouldError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				suite.ElementsMatch(tt.wantInterestRates, res.InterestRates)
			}
		})
	}
}

func (suite *grpcQueryTestSuite) TestGrpcQueryInterestFactors() {
	res, err := suite.queryServer.InterestFactors(suite.ctx, &types2.QueryInterestFactorsRequest{
		Denom: "usdx",
	})
	suite.Require().NoError(err)

	suite.Equal(&types2.QueryInterestFactorsResponse{
		InterestFactors: types2.InterestFactors{
			{
				Denom:                "usdx",
				BorrowInterestFactor: "1.000000000000000000",
				SupplyInterestFactor: "1.000000000000000000",
			},
		},
	}, res)
}

func (suite *grpcQueryTestSuite) TestGrpcQueryReserves() {
	suite.addDeposits()
	suite.addBorrows()

	res, err := suite.queryServer.Reserves(suite.ctx, &types2.QueryReservesRequest{})
	suite.Require().NoError(err)

	suite.Equal(&types2.QueryReservesResponse{
		Amount: sdk.Coins{},
	}, res)
}

func TestGrpcQueryTestSuite(t *testing.T) {
	suite.Run(t, new(grpcQueryTestSuite))
}
