package keeper_test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"cosmossdk.io/log"

	"github.com/joltify-finance/joltify_lending/client"

	jolttypes "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/testutil"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
	joltkeeper "github.com/joltify-finance/joltify_lending/x/third_party/jolt/keeper"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"

	"github.com/joltify-finance/joltify_lending/app"
)

// Test suite used for all keeper tests
type PayoutTestSuite struct {
	suite.Suite

	keeper     keeper.Keeper
	joltKeeper joltkeeper.Keeper

	app app.TestApp
	ctx context.Context

	genesisTime time.Time
	addrs       []sdk.AccAddress
}

// SetupTest is run automatically before each suite test
func (suite *PayoutTestSuite) SetupTest() {
	config := sdk.GetConfig()
	app.SetBech32AddressPrefixes(config)

	_, suite.addrs = app.GeneratePrivKeyAddressPairs(5)

	suite.genesisTime = time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC)
}

func (suite *PayoutTestSuite) SetupApp() {
	suite.app = app.NewTestApp(log.NewTestLogger(suite.T()), suite.T().TempDir())

	suite.keeper = suite.app.GetIncentiveKeeper()
	suite.joltKeeper = suite.app.GetJoltKeeper()

	suite.ctx = suite.app.NewContext(true)
}

func (suite *PayoutTestSuite) SetupWithGenState(authBuilder app.AuthBankGenesisBuilder, incentBuilder testutil.IncentiveGenesisBuilder, hardBuilder testutil.JoltGenesisBuilder) {
	suite.SetupApp()
	suite.app.InitializeFromGenesisStatesWithTime(
		suite.genesisTime, nil, nil,
		authBuilder.BuildMarshalled(suite.app.AppCodec()),
		NewPricefeedGenStateMultiFromTime(suite.app.AppCodec(), suite.genesisTime),
		hardBuilder.BuildMarshalled(suite.app.AppCodec()),
		incentBuilder.BuildMarshalled(suite.app.AppCodec()),
	)
}

func (suite *PayoutTestSuite) getAccount(addr sdk.AccAddress) sdk.AccountI {
	ak := suite.app.GetAccountKeeper()
	return ak.GetAccount(suite.ctx, addr)
}

func (suite *PayoutTestSuite) getModuleAccount(name string) authtypes.ModuleAccountI {
	ak := suite.app.GetAccountKeeper()
	return ak.GetModuleAccount(suite.ctx, name)
}

func (suite *PayoutTestSuite) TestSendCoinsToPeriodicVestingAccount() {
	// need to add the offset of 50000000000000 as we transfer that amount in genesis to incentives module
	type accountArgs struct {
		periods          []vestingtypes.Period
		origVestingCoins sdk.Coins
		startTime        int64
		endTime          int64
	}
	type args struct {
		accArgs             accountArgs
		period              vestingtypes.Period
		ctxTime             time.Time
		mintModAccountCoins bool
		expectedPeriods     []vestingtypes.Period
		expectedStartTime   int64
		expectedEndTime     int64
	}
	type errArgs struct {
		expectErr bool
		contains  string
	}
	type testCase struct {
		name    string
		args    args
		errArgs errArgs
	}
	type testCases []testCase

	// we need to set the flag to avoid sending tokens to incentive module
	client.MAINNETFLAG = "unittest"

	tests := testCases{
		{
			name: "insert period at beginning schedule",
			args: args{
				accArgs: accountArgs{
					periods: []vestingtypes.Period{
						{Length: 5, Amount: cs(c("ujolt", 5))},
						{Length: 5, Amount: cs(c("ujolt", 5))},
						{Length: 5, Amount: cs(c("ujolt", 5))},
						{Length: 5, Amount: cs(c("ujolt", 5))},
					},
					origVestingCoins: cs(c("ujolt", 20)),
					startTime:        100,
					endTime:          120,
				},
				period:              vestingtypes.Period{Length: 2, Amount: cs(c("ujolt", 50000000000006))},
				ctxTime:             time.Unix(101, 0),
				mintModAccountCoins: true,
				expectedPeriods: []vestingtypes.Period{
					{Length: 3, Amount: cs(c("ujolt", 50000000000006))},
					{Length: 2, Amount: cs(c("ujolt", 5))},
					{Length: 5, Amount: cs(c("ujolt", 5))},
					{Length: 5, Amount: cs(c("ujolt", 5))},
					{Length: 5, Amount: cs(c("ujolt", 5))},
				},
				expectedStartTime: 100,
				expectedEndTime:   120,
			},
			errArgs: errArgs{
				expectErr: false,
				contains:  "",
			},
		},
		{
			name: "insert period at beginning with new start time",
			args: args{
				accArgs: accountArgs{
					periods: []vestingtypes.Period{
						{Length: 5, Amount: cs(c("ujolt", 5))},
						{Length: 5, Amount: cs(c("ujolt", 5))},
						{Length: 5, Amount: cs(c("ujolt", 5))},
						{Length: 5, Amount: cs(c("ujolt", 5))},
					},
					origVestingCoins: cs(c("ujolt", 20)),
					startTime:        100,
					endTime:          120,
				},
				period:              vestingtypes.Period{Length: 7, Amount: cs(c("ujolt", 50000000000006))},
				ctxTime:             time.Unix(80, 0),
				mintModAccountCoins: true,
				expectedPeriods: []vestingtypes.Period{
					{Length: 7, Amount: cs(c("ujolt", 50000000000006))},
					{Length: 18, Amount: cs(c("ujolt", 5))},
					{Length: 5, Amount: cs(c("ujolt", 5))},
					{Length: 5, Amount: cs(c("ujolt", 5))},
					{Length: 5, Amount: cs(c("ujolt", 5))},
				},
				expectedStartTime: 80,
				expectedEndTime:   120,
			},
			errArgs: errArgs{
				expectErr: false,
				contains:  "",
			},
		},
		{
			name: "insert period in middle of schedule",
			args: args{
				accArgs: accountArgs{
					periods: []vestingtypes.Period{
						{Length: 5, Amount: cs(c("ujolt", 5))},
						{Length: 5, Amount: cs(c("ujolt", 5))},
						{Length: 5, Amount: cs(c("ujolt", 5))},
						{Length: 5, Amount: cs(c("ujolt", 5))},
					},
					origVestingCoins: cs(c("ujolt", 20)),
					startTime:        100,
					endTime:          120,
				},
				period:              vestingtypes.Period{Length: 7, Amount: cs(c("ujolt", 50000000000006))},
				ctxTime:             time.Unix(101, 0),
				mintModAccountCoins: true,
				expectedPeriods: []vestingtypes.Period{
					{Length: 5, Amount: cs(c("ujolt", 5))},
					{Length: 3, Amount: cs(c("ujolt", 50000000000006))},
					{Length: 2, Amount: cs(c("ujolt", 5))},
					{Length: 5, Amount: cs(c("ujolt", 5))},
					{Length: 5, Amount: cs(c("ujolt", 5))},
				},
				expectedStartTime: 100,
				expectedEndTime:   120,
			},
			errArgs: errArgs{
				expectErr: false,
				contains:  "",
			},
		},
		{
			name: "append to end of schedule",
			args: args{
				accArgs: accountArgs{
					periods: []vestingtypes.Period{
						{Length: 5, Amount: cs(c("ujolt", 5))},
						{Length: 5, Amount: cs(c("ujolt", 5))},
						{Length: 5, Amount: cs(c("ujolt", 5))},
						{Length: 5, Amount: cs(c("ujolt", 5))},
					},
					origVestingCoins: cs(c("ujolt", 20)),
					startTime:        100,
					endTime:          120,
				},
				period:              vestingtypes.Period{Length: 7, Amount: cs(c("ujolt", 50000000000006))},
				ctxTime:             time.Unix(125, 0),
				mintModAccountCoins: true,
				expectedPeriods: []vestingtypes.Period{
					{Length: 5, Amount: cs(c("ujolt", 5))},
					{Length: 5, Amount: cs(c("ujolt", 5))},
					{Length: 5, Amount: cs(c("ujolt", 5))},
					{Length: 5, Amount: cs(c("ujolt", 5))},
					{Length: 12, Amount: cs(c("ujolt", 50000000000006))},
				},
				expectedStartTime: 100,
				expectedEndTime:   132,
			},
			errArgs: errArgs{
				expectErr: false,
				contains:  "",
			},
		},
		{
			name: "add coins to existing period",
			args: args{
				accArgs: accountArgs{
					periods: []vestingtypes.Period{
						{Length: 5, Amount: cs(c("ujolt", 5))},
						{Length: 5, Amount: cs(c("ujolt", 5))},
						{Length: 5, Amount: cs(c("ujolt", 5))},
						{Length: 5, Amount: cs(c("ujolt", 5))},
					},
					origVestingCoins: cs(c("ujolt", 20)),
					startTime:        100,
					endTime:          120,
				},
				period:              vestingtypes.Period{Length: 5, Amount: cs(c("ujolt", 50000000000006))},
				ctxTime:             time.Unix(110, 0),
				mintModAccountCoins: true,
				expectedPeriods: []vestingtypes.Period{
					{Length: 5, Amount: cs(c("ujolt", 5))},
					{Length: 5, Amount: cs(c("ujolt", 5))},
					{Length: 5, Amount: cs(c("ujolt", 50000000000011))},
					{Length: 5, Amount: cs(c("ujolt", 5))},
				},
				expectedStartTime: 100,
				expectedEndTime:   120,
			},
			errArgs: errArgs{
				expectErr: false,
				contains:  "",
			},
		},
		{
			name: "insufficient mod account balance",
			args: args{
				accArgs: accountArgs{
					periods: []vestingtypes.Period{
						{Length: 5, Amount: cs(c("ujolt", 5))},
						{Length: 5, Amount: cs(c("ujolt", 5))},
						{Length: 5, Amount: cs(c("ujolt", 5))},
						{Length: 5, Amount: cs(c("ujolt", 5))},
					},
					origVestingCoins: cs(c("ujolt", 20)),
					startTime:        100,
					endTime:          120,
				},
				period:              vestingtypes.Period{Length: 7, Amount: cs(c("ujolt", 50000000000006))},
				ctxTime:             time.Unix(125, 0),
				mintModAccountCoins: false,
				expectedPeriods: []vestingtypes.Period{
					{Length: 5, Amount: cs(c("ujolt", 5))},
					{Length: 5, Amount: cs(c("ujolt", 5))},
					{Length: 5, Amount: cs(c("ujolt", 5))},
					{Length: 5, Amount: cs(c("ujolt", 5))},
					{Length: 12, Amount: cs(c("ujolt", 50000000000006))},
				},
				expectedStartTime: 100,
				expectedEndTime:   132,
			},
			errArgs: errArgs{
				expectErr: true,
				contains:  "insufficient funds",
			},
		},
		{
			name: "add large period mid schedule",
			args: args{
				accArgs: accountArgs{
					periods: []vestingtypes.Period{
						{Length: 5, Amount: cs(c("ujolt", 5))},
						{Length: 5, Amount: cs(c("ujolt", 5))},
						{Length: 5, Amount: cs(c("ujolt", 5))},
						{Length: 5, Amount: cs(c("ujolt", 5))},
					},
					origVestingCoins: cs(c("ujolt", 20)),
					startTime:        100,
					endTime:          120,
				},
				period:              vestingtypes.Period{Length: 50, Amount: cs(c("ujolt", 50000000000006))},
				ctxTime:             time.Unix(110, 0),
				mintModAccountCoins: true,
				expectedPeriods: []vestingtypes.Period{
					{Length: 5, Amount: cs(c("ujolt", 5))},
					{Length: 5, Amount: cs(c("ujolt", 5))},
					{Length: 5, Amount: cs(c("ujolt", 5))},
					{Length: 5, Amount: cs(c("ujolt", 5))},
					{Length: 40, Amount: cs(c("ujolt", 50000000000006))},
				},
				expectedStartTime: 100,
				expectedEndTime:   160,
			},
			errArgs: errArgs{
				expectErr: false,
				contains:  "",
			},
		},
	}
	for _, tc := range tests {
		suite.Run(tc.name, func() {
			authBuilder := app.NewAuthBankGenesisBuilder().WithSimplePeriodicVestingAccount(
				suite.addrs[0],
				tc.args.accArgs.origVestingCoins,
				tc.args.accArgs.periods,
				tc.args.accArgs.startTime,
			)

			if tc.args.mintModAccountCoins {
				fmt.Printf(">>>>>%v\n", tc.args.period.Amount)
				authBuilder = authBuilder.WithSimpleModuleAccount(types2.ModuleName, tc.args.period.Amount)
			}

			suite.genesisTime = tc.args.ctxTime
			suite.SetupApp()
			suite.app.InitializeFromGenesisStates(nil, nil,
				authBuilder.BuildMarshalled(suite.app.AppCodec()),
			)

			if tc.args.mintModAccountCoins {
				err := fundModuleAccount(suite.app.GetBankKeeper(), suite.ctx, types2.ModuleName, tc.args.period.Amount)
				suite.Require().NoError(err)

			}

			err := suite.keeper.SendTimeLockedCoinsToPeriodicVestingAccount(suite.ctx, types2.ModuleName, suite.addrs[0], tc.args.period.Amount, tc.args.period.Length)

			if tc.errArgs.expectErr {
				suite.Require().True(strings.Contains(err.Error(), tc.errArgs.contains))
			} else {
				suite.Require().NoError(err)

				acc := suite.getAccount(suite.addrs[0])
				vacc, ok := acc.(*vestingtypes.PeriodicVestingAccount)
				suite.Require().True(ok)
				suite.Require().Equal(tc.args.expectedPeriods, vacc.VestingPeriods)
				suite.Require().Equal(tc.args.expectedStartTime, vacc.StartTime)
				suite.Require().Equal(tc.args.expectedEndTime, vacc.EndTime)
			}
		})
	}
}

func fundModuleAccount(bankKeeper bankkeeper.Keeper, ctx context.Context, recipientMod string, amounts sdk.Coins) error {
	if err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, amounts); err != nil {
		return err
	}
	return bankKeeper.SendCoinsFromModuleToModule(ctx, minttypes.ModuleName, recipientMod, amounts)
}

func fundAccount(bankKeeper bankkeeper.Keeper, ctx context.Context, addr sdk.AccAddress, amounts sdk.Coins) error {
	if err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, amounts); err != nil {
		return err
	}
	return bankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr, amounts)
}

func (suite *PayoutTestSuite) TestSendCoinsToBaseAccount() {
	authBuilder := app.NewAuthBankGenesisBuilder().
		WithSimpleAccount(suite.addrs[1], cs(c("ujolt", 400))).
		WithSimpleModuleAccount(types2.ModuleName, cs(c("ujolt", 600)))

	suite.genesisTime = time.Unix(100, 0)
	suite.SetupApp()

	var genAcc []authtypes.GenesisAccount
	b := authtypes.NewBaseAccount(suite.addrs[1], nil, 0, 0)
	genAcc = append(genAcc, b)

	suite.app.InitializeFromGenesisStates(genAcc, cs(c("ujolt", 400)),
		authBuilder.BuildMarshalled(suite.app.AppCodec()),
	)

	err := fundModuleAccount(suite.app.GetBankKeeper(), suite.app.Ctx, types2.ModuleName, cs(c("ujolt", 600)))
	suite.Require().NoError(err)

	// send coins to base account
	err = suite.keeper.SendTimeLockedCoinsToAccount(suite.ctx, types2.ModuleName, suite.addrs[1], cs(c("ujolt", 100)), 5)
	suite.Require().NoError(err)
	acc := suite.getAccount(suite.addrs[1])
	vacc, ok := acc.(*vestingtypes.PeriodicVestingAccount)
	suite.True(ok)
	expectedPeriods := []vestingtypes.Period{
		{Length: int64(5), Amount: cs(c("ujolt", 100))},
	}

	bk := suite.app.GetBankKeeper()

	suite.Equal(expectedPeriods, vacc.VestingPeriods)
	suite.Equal(cs(c("ujolt", 100)), vacc.OriginalVesting)
	suite.Equal(cs(c("ujolt", 500), c("stake", 100000000000000)), bk.GetAllBalances(suite.ctx, vacc.GetAddress()))
	suite.Equal(int64(105), vacc.EndTime)
	suite.Equal(int64(100), vacc.StartTime)
}

func (suite *PayoutTestSuite) TestSendCoinsToInvalidAccount() {
	authBuilder := app.NewAuthBankGenesisBuilder().
		WithSimpleModuleAccount(types2.ModuleName, cs(c("ujolt", 600)))

	suite.SetupApp()
	suite.app.InitializeFromGenesisStates(nil, nil,
		authBuilder.BuildMarshalled(suite.app.AppCodec()),
	)

	// No longer an empty validator vesting account, just a regular addr
	err := suite.keeper.SendTimeLockedCoinsToAccount(suite.ctx, types2.ModuleName, suite.addrs[2], cs(c("ujolt", 100)), 5)
	suite.Require().ErrorIs(err, types2.ErrAccountNotFound)

	macc := suite.getModuleAccount(jolttypes.ModuleName)
	err = suite.keeper.SendTimeLockedCoinsToAccount(suite.ctx, types2.ModuleName, macc.GetAddress(), cs(c("ujolt", 100)), 5)
	suite.Require().ErrorIs(err, types2.ErrInvalidAccountType)
}

func (suite *PayoutTestSuite) TestGetPeriodLength() {
	type args struct {
		blockTime time.Time
		lockup    int64
	}
	type periodTest struct {
		name           string
		args           args
		expectedLength int64
	}
	testCases := []periodTest{
		{
			name: "first half of month",
			args: args{
				blockTime: time.Date(2020, 11, 2, 15, 0, 0, 0, time.UTC),
				lockup:    6,
			},
			expectedLength: time.Date(2021, 5, 15, 14, 0, 0, 0, time.UTC).Unix() - time.Date(2020, 11, 2, 15, 0, 0, 0, time.UTC).Unix(),
		},
		{
			name: "first half of month long lockup",
			args: args{
				blockTime: time.Date(2020, 11, 2, 15, 0, 0, 0, time.UTC),
				lockup:    24,
			},
			expectedLength: time.Date(2022, 11, 15, 14, 0, 0, 0, time.UTC).Unix() - time.Date(2020, 11, 2, 15, 0, 0, 0, time.UTC).Unix(),
		},
		{
			name: "second half of month",
			args: args{
				blockTime: time.Date(2020, 12, 31, 15, 0, 0, 0, time.UTC),
				lockup:    6,
			},
			expectedLength: time.Date(2021, 7, 1, 14, 0, 0, 0, time.UTC).Unix() - time.Date(2020, 12, 31, 15, 0, 0, 0, time.UTC).Unix(),
		},
		{
			name: "second half of month long lockup",
			args: args{
				blockTime: time.Date(2020, 12, 31, 15, 0, 0, 0, time.UTC),
				lockup:    24,
			},
			expectedLength: time.Date(2023, 1, 1, 14, 0, 0, 0, time.UTC).Unix() - time.Date(2020, 12, 31, 15, 0, 0, 0, time.UTC).Unix(),
		},
		{
			name: "end of feb",
			args: args{
				blockTime: time.Date(2021, 2, 28, 15, 0, 0, 0, time.UTC),
				lockup:    6,
			},
			expectedLength: time.Date(2021, 9, 1, 14, 0, 0, 0, time.UTC).Unix() - time.Date(2021, 2, 28, 15, 0, 0, 0, time.UTC).Unix(),
		},
		{
			name: "leap year",
			args: args{
				blockTime: time.Date(2020, 2, 29, 15, 0, 0, 0, time.UTC),
				lockup:    6,
			},
			expectedLength: time.Date(2020, 9, 1, 14, 0, 0, 0, time.UTC).Unix() - time.Date(2020, 2, 29, 15, 0, 0, 0, time.UTC).Unix(),
		},
		{
			name: "leap year long lockup",
			args: args{
				blockTime: time.Date(2020, 2, 29, 15, 0, 0, 0, time.UTC),
				lockup:    24,
			},
			expectedLength: time.Date(2022, 3, 1, 14, 0, 0, 0, time.UTC).Unix() - time.Date(2020, 2, 29, 15, 0, 0, 0, time.UTC).Unix(),
		},
		{
			name: "exactly half of month, is pushed to start of month + lockup",
			args: args{
				blockTime: time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC),
				lockup:    6,
			},
			expectedLength: time.Date(2021, 7, 1, 14, 0, 0, 0, time.UTC).Unix() - time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC).Unix(),
		},
		{
			name: "just before half of month",
			args: args{
				blockTime: time.Date(2020, 12, 15, 13, 59, 59, 0, time.UTC),
				lockup:    6,
			},
			expectedLength: time.Date(2021, 6, 15, 14, 0, 0, 0, time.UTC).Unix() - time.Date(2020, 12, 15, 13, 59, 59, 0, time.UTC).Unix(),
		},
		{
			name: "just after start of month payout time, is pushed to mid month + lockup",
			args: args{
				blockTime: time.Date(2020, 12, 1, 14, 0, 1, 0, time.UTC),
				lockup:    1,
			},
			expectedLength: time.Date(2021, 1, 15, 14, 0, 0, 0, time.UTC).Unix() - time.Date(2020, 12, 1, 14, 0, 1, 0, time.UTC).Unix(),
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			length := suite.keeper.GetPeriodLength(tc.args.blockTime, tc.args.lockup)
			suite.Require().Equal(tc.expectedLength, length)
		})
	}
}

func TestPayoutTestSuite(t *testing.T) {
	suite.Run(t, new(PayoutTestSuite))
}
