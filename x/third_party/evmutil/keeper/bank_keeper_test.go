package keeper_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	tmtime "github.com/cometbft/cometbft/types/time"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vesting "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	"github.com/joltify-finance/joltify_lending/x/third_party/evmutil/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party/evmutil/testutil"
	"github.com/joltify-finance/joltify_lending/x/third_party/evmutil/types"
)

type evmBankKeeperTestSuite struct {
	testutil.Suite
}

func (suite *evmBankKeeperTestSuite) SetupTest() {
	suite.Suite.SetupTest()
}

func (suite *evmBankKeeperTestSuite) TestGetBalance_ReturnsSpendable() {
	startingCoins := sdk.NewCoins(sdk.NewInt64Coin("ujolt", 10))
	startingAjolt := sdkmath.NewInt(100)

	now := tmtime.Now()
	endTime := now.Add(24 * time.Hour)
	bacc := authtypes.NewBaseAccountWithAddress(suite.Addrs[0])
	vacc := vesting.NewContinuousVestingAccount(bacc, startingCoins, now.Unix(), endTime.Unix())
	suite.AccountKeeper.SetAccount(suite.Ctx, vacc)

	err := suite.App.FundAccount(suite.Ctx, suite.Addrs[0], startingCoins)
	suite.Require().NoError(err)
	err = suite.Keeper.SetBalance(suite.Ctx, suite.Addrs[0], startingAjolt)
	suite.Require().NoError(err)

	coin := suite.EvmBankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "ajolt")
	suite.Require().Equal(startingAjolt, coin.Amount)

	ctx := suite.Ctx.WithBlockTime(now.Add(12 * time.Hour))
	coin = suite.EvmBankKeeper.GetBalance(ctx, suite.Addrs[0], "ajolt")
	suite.Require().Equal(sdkmath.NewIntFromUint64(5_000_000_000_100), coin.Amount)
}

func (suite *evmBankKeeperTestSuite) TestGetBalance_NotEvmDenom() {
	suite.Require().Panics(func() {
		suite.EvmBankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "ujolt")
	})
	suite.Require().Panics(func() {
		suite.EvmBankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "busd")
	})
}

func (suite *evmBankKeeperTestSuite) TestGetBalance() {
	tests := []struct {
		name           string
		startingAmount sdk.Coins
		expAmount      sdkmath.Int
	}{
		{
			"ujolt with ajolt",
			sdk.NewCoins(
				sdk.NewInt64Coin("ajolt", 100),
				sdk.NewInt64Coin("ujolt", 10),
			),
			sdkmath.NewInt(10_000_000_000_100),
		},
		{
			"just ajolt",
			sdk.NewCoins(
				sdk.NewInt64Coin("ajolt", 100),
				sdk.NewInt64Coin("busd", 100),
			),
			sdkmath.NewInt(100),
		},
		{
			"just ujolt",
			sdk.NewCoins(
				sdk.NewInt64Coin("ujolt", 10),
				sdk.NewInt64Coin("busd", 100),
			),
			sdkmath.NewInt(10_000_000_000_000),
		},
		{
			"no ujolt or ajolt",
			sdk.NewCoins(),
			sdkmath.ZeroInt(),
		},
		{
			"with avaka that is more than 1 ujolt",
			sdk.NewCoins(
				sdk.NewInt64Coin("ajolt", 20_000_000_000_220),
				sdk.NewInt64Coin("ujolt", 11),
			),
			sdkmath.NewInt(31_000_000_000_220),
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()

			suite.FundAccountWithJolt(suite.Addrs[0], tt.startingAmount)
			coin := suite.EvmBankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "ajolt")
			suite.Require().Equal(tt.expAmount, coin.Amount)
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestSendCoinsFromModuleToAccount() {
	startingModuleCoins := sdk.NewCoins(
		sdk.NewInt64Coin("ajolt", 200),
		sdk.NewInt64Coin("ujolt", 100),
	)
	tests := []struct {
		name           string
		sendCoins      sdk.Coins
		startingAccBal sdk.Coins
		expAccBal      sdk.Coins
		hasErr         bool
	}{
		{
			"send more than 1 ujolt",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 12_000_000_000_010)),
			sdk.Coins{},
			sdk.NewCoins(
				sdk.NewInt64Coin("ajolt", 10),
				sdk.NewInt64Coin("ujolt", 12),
			),
			false,
		},
		{
			"send less than 1 ujolt",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 122)),
			sdk.Coins{},
			sdk.NewCoins(
				sdk.NewInt64Coin("ajolt", 122),
				sdk.NewInt64Coin("ujolt", 0),
			),
			false,
		},
		{
			"send an exact amount of ujolt",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 98_000_000_000_000)),
			sdk.Coins{},
			sdk.NewCoins(
				sdk.NewInt64Coin("ajolt", 0o0),
				sdk.NewInt64Coin("ujolt", 98),
			),
			false,
		},
		{
			"send no ajolt",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 0)),
			sdk.Coins{},
			sdk.NewCoins(
				sdk.NewInt64Coin("ajolt", 0),
				sdk.NewInt64Coin("ujolt", 0),
			),
			false,
		},
		{
			"errors if sending other coins",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 500), sdk.NewInt64Coin("busd", 1000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"errors if not enough total ajolt to cover",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 100_000_000_001_000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"errors if not enough ujolt to cover",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 200_000_000_000_000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"converts receiver's ajolt to ujolt if there's enough ajolt after the transfer",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 99_000_000_000_200)),
			sdk.NewCoins(
				sdk.NewInt64Coin("ajolt", 999_999_999_900),
				sdk.NewInt64Coin("ujolt", 1),
			),
			sdk.NewCoins(
				sdk.NewInt64Coin("ajolt", 100),
				sdk.NewInt64Coin("ujolt", 101),
			),
			false,
		},
		{
			"converts all of receiver's ajolt to ujolt even if somehow receiver has more than 1ujolt of ajolt",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 12_000_000_000_100)),
			sdk.NewCoins(
				sdk.NewInt64Coin("ajolt", 5_999_999_999_990),
				sdk.NewInt64Coin("ujolt", 1),
			),
			sdk.NewCoins(
				sdk.NewInt64Coin("ajolt", 90),
				sdk.NewInt64Coin("ujolt", 19),
			),
			false,
		},
		{
			"swap 1 ujolt for ajolt if module account doesn't have enough ajolt",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 99_000_000_001_000)),
			sdk.NewCoins(
				sdk.NewInt64Coin("ajolt", 200),
				sdk.NewInt64Coin("ujolt", 1),
			),
			sdk.NewCoins(
				sdk.NewInt64Coin("ajolt", 1200),
				sdk.NewInt64Coin("ujolt", 100),
			),
			false,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()

			suite.FundAccountWithJolt(suite.Addrs[0], tt.startingAccBal)
			suite.FundModuleAccountWithJolt(evmtypes.ModuleName, startingModuleCoins)

			// fund our module with some ujolt to account for converting extra ajolt back to ujolt
			suite.FundModuleAccountWithJolt(types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin("ujolt", 10)))

			err := suite.EvmBankKeeper.SendCoinsFromModuleToAccount(suite.Ctx, evmtypes.ModuleName, suite.Addrs[0], tt.sendCoins)
			if tt.hasErr {
				suite.Require().Error(err)
				return
			} else {
				suite.Require().NoError(err)
			}

			// check ujolt
			ujoltSender := suite.BankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "ujolt")
			suite.Require().Equal(tt.expAccBal.AmountOf("ujolt").Int64(), ujoltSender.Amount.Int64())

			// check ajolt
			actualAjolt := suite.Keeper.GetBalance(suite.Ctx, suite.Addrs[0])
			suite.Require().Equal(tt.expAccBal.AmountOf("ajolt").Int64(), actualAjolt.Int64())
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestSendCoinsFromAccountToModule() {
	startingAccCoins := sdk.NewCoins(
		sdk.NewInt64Coin("ajolt", 200),
		sdk.NewInt64Coin("ujolt", 100),
	)
	startingModuleCoins := sdk.NewCoins(
		sdk.NewInt64Coin("ajolt", 100_000_000_000),
	)
	tests := []struct {
		name           string
		sendCoins      sdk.Coins
		expSenderCoins sdk.Coins
		expModuleCoins sdk.Coins
		hasErr         bool
	}{
		{
			"send more than 1 ujolt",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 12_000_000_000_010)),
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 190), sdk.NewInt64Coin("ujolt", 88)),
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 100_000_000_010), sdk.NewInt64Coin("ujolt", 12)),
			false,
		},
		{
			"send less than 1 ujolt",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 122)),
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 78), sdk.NewInt64Coin("ujolt", 100)),
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 100_000_000_122), sdk.NewInt64Coin("ujolt", 0)),
			false,
		},
		{
			"send an exact amount of ujolt",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 98_000_000_000_000)),
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 200), sdk.NewInt64Coin("ujolt", 2)),
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 100_000_000_000), sdk.NewInt64Coin("ujolt", 98)),
			false,
		},
		{
			"send no ajolt",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 0)),
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 200), sdk.NewInt64Coin("ujolt", 100)),
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 100_000_000_000), sdk.NewInt64Coin("ujolt", 0)),
			false,
		},
		{
			"errors if sending other coins",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 500), sdk.NewInt64Coin("busd", 1000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"errors if have dup coins",
			sdk.Coins{
				sdk.NewInt64Coin("ajolt", 12_000_000_000_000),
				sdk.NewInt64Coin("ajolt", 2_000_000_000_000),
			},
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"errors if not enough total ajolt to cover",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 100_000_000_001_000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"errors if not enough ujolt to cover",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 200_000_000_000_000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"converts 1 ujolt to ajolt if not enough ajolt to cover",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 99_001_000_000_000)),
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 999_000_000_200), sdk.NewInt64Coin("ujolt", 0)),
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 101_000_000_000), sdk.NewInt64Coin("ujolt", 99)),
			false,
		},
		{
			"converts receiver's ajolt to ujolt if there's enough ajolt after the transfer",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 5_900_000_000_200)),
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 100_000_000_000), sdk.NewInt64Coin("ujolt", 94)),
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 200), sdk.NewInt64Coin("ujolt", 6)),
			false,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()
			suite.FundAccountWithJolt(suite.Addrs[0], startingAccCoins)
			suite.FundModuleAccountWithJolt(evmtypes.ModuleName, startingModuleCoins)

			err := suite.EvmBankKeeper.SendCoinsFromAccountToModule(suite.Ctx, suite.Addrs[0], evmtypes.ModuleName, tt.sendCoins)
			if tt.hasErr {
				suite.Require().Error(err)
				return
			} else {
				suite.Require().NoError(err)
			}

			// check sender balance
			ujoltSender := suite.BankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "ujolt")
			suite.Require().Equal(tt.expSenderCoins.AmountOf("ujolt").Int64(), ujoltSender.Amount.Int64())
			actualAjolt := suite.Keeper.GetBalance(suite.Ctx, suite.Addrs[0])
			suite.Require().Equal(tt.expSenderCoins.AmountOf("ajolt").Int64(), actualAjolt.Int64())

			// check module balance
			moduleAddr := suite.AccountKeeper.GetModuleAddress(evmtypes.ModuleName)
			ujoltSender = suite.BankKeeper.GetBalance(suite.Ctx, moduleAddr, "ujolt")
			suite.Require().Equal(tt.expModuleCoins.AmountOf("ujolt").Int64(), ujoltSender.Amount.Int64())
			actualAjolt = suite.Keeper.GetBalance(suite.Ctx, moduleAddr)
			suite.Require().Equal(tt.expModuleCoins.AmountOf("ajolt").Int64(), actualAjolt.Int64())
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestBurnCoins() {
	startingUjolt := sdkmath.NewInt(100)
	tests := []struct {
		name       string
		burnCoins  sdk.Coins
		expUjolt   sdkmath.Int
		expAjolt   sdkmath.Int
		hasErr     bool
		ajoltStart sdkmath.Int
	}{
		{
			"burn more than 1 ujolt",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 12_021_000_000_002)),
			sdkmath.NewInt(88),
			sdkmath.NewInt(100_000_000_000),
			false,
			sdkmath.NewInt(121_000_000_002),
		},
		{
			"burn less than 1 ujolt",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 122)),
			sdkmath.NewInt(100),
			sdkmath.NewInt(878),
			false,
			sdkmath.NewInt(1000),
		},
		{
			"burn an exact amount of ujolt",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 98_000_000_000_000)),
			sdkmath.NewInt(2),
			sdkmath.NewInt(10),
			false,
			sdkmath.NewInt(10),
		},
		{
			"burn no ajolt",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 0)),
			startingUjolt,
			sdkmath.ZeroInt(),
			false,
			sdkmath.ZeroInt(),
		},
		{
			"errors if burning other coins",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 500), sdk.NewInt64Coin("busd", 1000)),
			startingUjolt,
			sdkmath.NewInt(100),
			true,
			sdkmath.NewInt(100),
		},
		{
			"errors if have dup coins",
			sdk.Coins{
				sdk.NewInt64Coin("ajolt", 12_000_000_000_000),
				sdk.NewInt64Coin("ajolt", 2_000_000_000_000),
			},
			startingUjolt,
			sdkmath.ZeroInt(),
			true,
			sdkmath.ZeroInt(),
		},
		{
			"errors if burn amount is negative",
			sdk.Coins{sdk.Coin{Denom: "ajolt", Amount: sdkmath.NewInt(-100)}},
			startingUjolt,
			sdkmath.NewInt(50),
			true,
			sdkmath.NewInt(50),
		},
		{
			"errors if not enough ajolt to cover burn",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 100_999_000_000_000)),
			sdkmath.NewInt(0),
			sdkmath.NewInt(99_000_000_000),
			true,
			sdkmath.NewInt(99_000_000_000),
		},
		{
			"errors if not enough ujolt to cover burn",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 200_000_000_000_000)),
			sdkmath.NewInt(100),
			sdkmath.ZeroInt(),
			true,
			sdkmath.ZeroInt(),
		},
		{
			"converts 1 ujolt to ajolt if not enough ajolt to cover",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 12_021_000_000_002)),
			sdkmath.NewInt(87),
			sdkmath.NewInt(980_000_000_000),
			false,
			sdkmath.NewInt(1_000_000_002),
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()
			startingCoins := sdk.NewCoins(
				sdk.NewCoin("ujolt", startingUjolt),
				sdk.NewCoin("ajolt", tt.ajoltStart),
			)
			suite.FundModuleAccountWithJolt(evmtypes.ModuleName, startingCoins)

			err := suite.EvmBankKeeper.BurnCoins(suite.Ctx, evmtypes.ModuleName, tt.burnCoins)
			if tt.hasErr {
				suite.Require().Error(err)
				return
			} else {
				suite.Require().NoError(err)
			}

			// check ujolt
			ujoltActual := suite.BankKeeper.GetBalance(suite.Ctx, suite.EvmModuleAddr, "ujolt")
			suite.Require().Equal(tt.expUjolt, ujoltActual.Amount)

			// check ajolt
			ajoltActual := suite.Keeper.GetBalance(suite.Ctx, suite.EvmModuleAddr)
			suite.Require().Equal(tt.expAjolt, ajoltActual)
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestMintCoins() {
	tests := []struct {
		name       string
		mintCoins  sdk.Coins
		ujolt      sdkmath.Int
		ajolt      sdkmath.Int
		hasErr     bool
		ajoltStart sdkmath.Int
	}{
		{
			"mint more than 1 ujolt",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 12_021_000_000_002)),
			sdkmath.NewInt(12),
			sdkmath.NewInt(21_000_000_002),
			false,
			sdkmath.ZeroInt(),
		},
		{
			"mint less than 1 ujolt",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 901_000_000_001)),
			sdkmath.ZeroInt(),
			sdkmath.NewInt(901_000_000_001),
			false,
			sdkmath.ZeroInt(),
		},
		{
			"mint an exact amount of ujolt",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 123_000_000_000_000_000)),
			sdkmath.NewInt(123_000),
			sdkmath.ZeroInt(),
			false,
			sdkmath.ZeroInt(),
		},
		{
			"mint no ajolt",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 0)),
			sdkmath.ZeroInt(),
			sdkmath.ZeroInt(),
			false,
			sdkmath.ZeroInt(),
		},
		{
			"errors if minting other coins",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 500), sdk.NewInt64Coin("busd", 1000)),
			sdkmath.ZeroInt(),
			sdkmath.NewInt(100),
			true,
			sdkmath.NewInt(100),
		},
		{
			"errors if have dup coins",
			sdk.Coins{
				sdk.NewInt64Coin("ajolt", 12_000_000_000_000),
				sdk.NewInt64Coin("ajolt", 2_000_000_000_000),
			},
			sdkmath.ZeroInt(),
			sdkmath.ZeroInt(),
			true,
			sdkmath.ZeroInt(),
		},
		{
			"errors if mint amount is negative",
			sdk.Coins{sdk.Coin{Denom: "ajolt", Amount: sdkmath.NewInt(-100)}},
			sdkmath.ZeroInt(),
			sdkmath.NewInt(50),
			true,
			sdkmath.NewInt(50),
		},
		{
			"adds to existing ajolt balance",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 12_021_000_000_002)),
			sdkmath.NewInt(12),
			sdkmath.NewInt(21_000_000_102),
			false,
			sdkmath.NewInt(100),
		},
		{
			"convert ajolt balance to ujolt if it exceeds 1 ujolt",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 10_999_000_000_000)),
			sdkmath.NewInt(12),
			sdkmath.NewInt(1_200_000_001),
			false,
			sdkmath.NewInt(1_002_200_000_001),
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()
			suite.FundModuleAccountWithJolt(types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin("ujolt", 10)))
			suite.FundModuleAccountWithJolt(evmtypes.ModuleName, sdk.NewCoins(sdk.NewCoin("ajolt", tt.ajoltStart)))

			err := suite.EvmBankKeeper.MintCoins(suite.Ctx, evmtypes.ModuleName, tt.mintCoins)
			if tt.hasErr {
				suite.Require().Error(err)
				return
			} else {
				suite.Require().NoError(err)
			}

			// check ujolt
			ujoltActual := suite.BankKeeper.GetBalance(suite.Ctx, suite.EvmModuleAddr, "ujolt")
			suite.Require().Equal(tt.ujolt, ujoltActual.Amount)

			// check ajolt
			ajoltActual := suite.Keeper.GetBalance(suite.Ctx, suite.EvmModuleAddr)
			suite.Require().Equal(tt.ajolt, ajoltActual)
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestValidateEvmCoins() {
	tests := []struct {
		name      string
		coins     sdk.Coins
		shouldErr bool
	}{
		{
			"valid coins",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 500)),
			false,
		},
		{
			"dup coins",
			sdk.Coins{sdk.NewInt64Coin("ajolt", 500), sdk.NewInt64Coin("ajolt", 500)},
			true,
		},
		{
			"not evm coins",
			sdk.NewCoins(sdk.NewInt64Coin("ujolt", 500)),
			true,
		},
		{
			"negative coins",
			sdk.Coins{sdk.Coin{Denom: "ajolt", Amount: sdkmath.NewInt(-500)}},
			true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			err := keeper.ValidateEvmCoins(tt.coins)
			if tt.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestConvertOneUjoltToAjoltIfNeeded() {
	ajoltNeeded := sdkmath.NewInt(200)
	tests := []struct {
		name          string
		startingCoins sdk.Coins
		expectedCoins sdk.Coins
		success       bool
	}{
		{
			"not enough ujolt for conversion",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 100)),
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 100)),
			false,
		},
		{
			"converts 1 ujolt to ajolt",
			sdk.NewCoins(sdk.NewInt64Coin("ujolt", 10), sdk.NewInt64Coin("ajolt", 100)),
			sdk.NewCoins(sdk.NewInt64Coin("ujolt", 9), sdk.NewInt64Coin("ajolt", 1_000_000_000_100)),
			true,
		},
		{
			"conversion not needed",
			sdk.NewCoins(sdk.NewInt64Coin("ujolt", 10), sdk.NewInt64Coin("ajolt", 200)),
			sdk.NewCoins(sdk.NewInt64Coin("ujolt", 10), sdk.NewInt64Coin("ajolt", 200)),
			true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()

			suite.FundAccountWithJolt(suite.Addrs[0], tt.startingCoins)
			err := suite.EvmBankKeeper.ConvertOneUjoltToAjoltIfNeeded(suite.Ctx, suite.Addrs[0], ajoltNeeded)
			moduleJolt := suite.BankKeeper.GetBalance(suite.Ctx, suite.AccountKeeper.GetModuleAddress(types.ModuleName), "ujolt")
			if tt.success {
				suite.Require().NoError(err)
				if tt.startingCoins.AmountOf("ajolt").LT(ajoltNeeded) {
					suite.Require().Equal(sdk.OneInt(), moduleJolt.Amount)
				}
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(sdkmath.ZeroInt(), moduleJolt.Amount)
			}

			ajolt := suite.Keeper.GetBalance(suite.Ctx, suite.Addrs[0])
			suite.Require().Equal(tt.expectedCoins.AmountOf("ajolt"), ajolt)
			ujolt := suite.BankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "ujolt")
			suite.Require().Equal(tt.expectedCoins.AmountOf("ujolt"), ujolt.Amount)
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestConvertAjoltToUjolt() {
	tests := []struct {
		name          string
		startingCoins sdk.Coins
		expectedCoins sdk.Coins
	}{
		{
			"not enough ujolt",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 100)),
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 100), sdk.NewInt64Coin("ujolt", 0)),
		},
		{
			"converts ajolt for 1 ujolt",
			sdk.NewCoins(sdk.NewInt64Coin("ujolt", 10), sdk.NewInt64Coin("ajolt", 1_000_000_000_003)),
			sdk.NewCoins(sdk.NewInt64Coin("ujolt", 11), sdk.NewInt64Coin("ajolt", 3)),
		},
		{
			"converts more than 1 ujolt of ajolt",
			sdk.NewCoins(sdk.NewInt64Coin("ujolt", 10), sdk.NewInt64Coin("ajolt", 8_000_000_000_123)),
			sdk.NewCoins(sdk.NewInt64Coin("ujolt", 18), sdk.NewInt64Coin("ajolt", 123)),
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()

			err := suite.App.FundModuleAccount(suite.Ctx, types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin("ujolt", 10)))
			suite.Require().NoError(err)
			suite.FundAccountWithJolt(suite.Addrs[0], tt.startingCoins)
			err = suite.EvmBankKeeper.ConvertAJoltToUJolt(suite.Ctx, suite.Addrs[0])
			suite.Require().NoError(err)
			ajolt := suite.Keeper.GetBalance(suite.Ctx, suite.Addrs[0])
			suite.Require().Equal(tt.expectedCoins.AmountOf("ajolt"), ajolt)
			ujolt := suite.BankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "ujolt")
			suite.Require().Equal(tt.expectedCoins.AmountOf("ujolt"), ujolt.Amount)
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestSplitAjoltCoins() {
	tests := []struct {
		name          string
		coins         sdk.Coins
		expectedCoins sdk.Coins
		shouldErr     bool
	}{
		{
			"invalid coins",
			sdk.NewCoins(sdk.NewInt64Coin("ujolt", 500)),
			nil,
			true,
		},
		{
			"empty coins",
			sdk.NewCoins(),
			sdk.NewCoins(),
			false,
		},
		{
			"ujolt & ajolt coins",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 8_000_000_000_123)),
			sdk.NewCoins(sdk.NewInt64Coin("ujolt", 8), sdk.NewInt64Coin("ajolt", 123)),
			false,
		},
		{
			"only ajolt",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 10_123)),
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 10_123)),
			false,
		},
		{
			"only ujolt",
			sdk.NewCoins(sdk.NewInt64Coin("ajolt", 5_000_000_000_000)),
			sdk.NewCoins(sdk.NewInt64Coin("ujolt", 5)),
			false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			ujolt, ajolt, err := keeper.SplitAJoltCoins(tt.coins)
			if tt.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tt.expectedCoins.AmountOf("ujolt"), ujolt.Amount)
				suite.Require().Equal(tt.expectedCoins.AmountOf("ajolt"), ajolt)
			}
		})
	}
}

func TestEvmBankKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(evmBankKeeperTestSuite))
}
