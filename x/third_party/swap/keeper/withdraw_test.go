package keeper_test

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/swap/types"
)

func (suite *keeperTestSuite) TestWithdraw_AllShares() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("ukava", sdkmath.NewInt(10e6)),
		sdk.NewCoin("usdx", sdkmath.NewInt(50e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	poolID := suite.setupPool(reserves, totalShares, owner.GetAddress())

	err := suite.Keeper.Withdraw(suite.Ctx, owner.GetAddress(), totalShares, reserves[0], reserves[1])
	suite.Require().NoError(err)

	suite.PoolDeleted(reserves[0].Denom, reserves[1].Denom)
	suite.PoolSharesDeleted(owner.GetAddress(), reserves[0].Denom, reserves[1].Denom)
	suite.AccountBalanceEqual(owner.GetAddress(), reserves)
	suite.ModuleAccountBalanceEqual(sdk.Coins{})

	suite.EventsContains(suite.Ctx.EventManager().Events(), sdk.NewEvent(
		types.EventTypeSwapWithdraw,
		sdk.NewAttribute(types.AttributeKeyPoolID, poolID),
		sdk.NewAttribute(types.AttributeKeyOwner, owner.GetAddress().String()),
		sdk.NewAttribute(sdk.AttributeKeyAmount, reserves.String()),
		sdk.NewAttribute(types.AttributeKeyShares, totalShares.String()),
	))
}

func (suite *keeperTestSuite) TestWithdraw_PartialShares() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("ukava", sdkmath.NewInt(10e6)),
		sdk.NewCoin("usdx", sdkmath.NewInt(50e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	poolID := suite.setupPool(reserves, totalShares, owner.GetAddress())

	sharesToWithdraw := sdkmath.NewInt(15e6)
	minCoinA := sdk.NewCoin("usdx", sdkmath.NewInt(25e6))
	minCoinB := sdk.NewCoin("ukava", sdkmath.NewInt(5e6))

	err := suite.Keeper.Withdraw(suite.Ctx, owner.GetAddress(), sharesToWithdraw, minCoinA, minCoinB)
	suite.Require().NoError(err)

	sharesLeft := totalShares.Sub(sharesToWithdraw)
	reservesLeft := sdk.NewCoins(reserves[0].Sub(minCoinB), reserves[1].Sub(minCoinA))

	suite.PoolShareTotalEqual(poolID, sharesLeft)
	suite.PoolDepositorSharesEqual(owner.GetAddress(), poolID, sharesLeft)
	suite.PoolReservesEqual(poolID, reservesLeft)
	suite.AccountBalanceEqual(owner.GetAddress(), sdk.NewCoins(minCoinA, minCoinB))
	suite.ModuleAccountBalanceEqual(reservesLeft)

	suite.EventsContains(suite.Ctx.EventManager().Events(), sdk.NewEvent(
		types.EventTypeSwapWithdraw,
		sdk.NewAttribute(types.AttributeKeyPoolID, poolID),
		sdk.NewAttribute(types.AttributeKeyOwner, owner.GetAddress().String()),
		sdk.NewAttribute(sdk.AttributeKeyAmount, sdk.NewCoins(minCoinA, minCoinB).String()),
		sdk.NewAttribute(types.AttributeKeyShares, sharesToWithdraw.String()),
	))
}

func (suite *keeperTestSuite) TestWithdraw_NoSharesOwned() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("ukava", sdkmath.NewInt(10e6)),
		sdk.NewCoin("usdx", sdkmath.NewInt(50e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	poolID := suite.setupPool(reserves, totalShares, owner.GetAddress())

	accWithNoDeposit := sdk.AccAddress("some account")

	err := suite.Keeper.Withdraw(suite.Ctx, accWithNoDeposit, totalShares, reserves[0], reserves[1])
	suite.EqualError(err, fmt.Sprintf("no deposit for account %s and pool %s: deposit not found", accWithNoDeposit.String(), poolID))
}

func (suite *keeperTestSuite) TestWithdraw_GreaterThanSharesOwned() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("ukava", sdkmath.NewInt(10e6)),
		sdk.NewCoin("usdx", sdkmath.NewInt(50e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	suite.setupPool(reserves, totalShares, owner.GetAddress())

	sharesToWithdraw := totalShares.Add(sdk.OneInt())
	err := suite.Keeper.Withdraw(suite.Ctx, owner.GetAddress(), sharesToWithdraw, reserves[0], reserves[1])
	suite.EqualError(err, fmt.Sprintf("withdraw of %s shares greater than %s shares owned: invalid shares", sharesToWithdraw, totalShares))
}

func (suite *keeperTestSuite) TestWithdraw_MinWithdraw() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("ukava", sdkmath.NewInt(10e6)),
		sdk.NewCoin("usdx", sdkmath.NewInt(50e6)),
	)
	totalShares := sdkmath.NewInt(30e6)

	testCases := []struct {
		shares     sdkmath.Int
		minCoinA   sdk.Coin
		minCoinB   sdk.Coin
		shouldFail bool
	}{
		{sdkmath.NewInt(1), sdk.NewCoin("ukava", sdkmath.NewInt(1)), sdk.NewCoin("usdx", sdkmath.NewInt(1)), true},
		{sdkmath.NewInt(1), sdk.NewCoin("usdx", sdkmath.NewInt(5)), sdk.NewCoin("ukava", sdkmath.NewInt(1)), true},

		{sdkmath.NewInt(2), sdk.NewCoin("ukava", sdkmath.NewInt(1)), sdk.NewCoin("usdx", sdkmath.NewInt(1)), true},
		{sdkmath.NewInt(2), sdk.NewCoin("usdx", sdkmath.NewInt(5)), sdk.NewCoin("ukava", sdkmath.NewInt(1)), true},

		{sdkmath.NewInt(3), sdk.NewCoin("ukava", sdkmath.NewInt(1)), sdk.NewCoin("usdx", sdkmath.NewInt(5)), false},
		{sdkmath.NewInt(3), sdk.NewCoin("usdx", sdkmath.NewInt(5)), sdk.NewCoin("ukava", sdkmath.NewInt(1)), false},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("shares=%s minCoinA=%s minCoinB=%s", tc.shares, tc.minCoinA, tc.minCoinB), func() {
			suite.SetupTest()
			suite.setupPool(reserves, totalShares, owner.GetAddress())

			err := suite.Keeper.Withdraw(suite.Ctx, owner.GetAddress(), tc.shares, tc.minCoinA, tc.minCoinB)
			if tc.shouldFail {
				suite.EqualError(err, "shares must be increased: insufficient liquidity")
			} else {
				suite.NoError(err, "expected no liquidity error")
			}
		})
	}
}

func (suite *keeperTestSuite) TestWithdraw_BelowMinimum() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("ukava", sdkmath.NewInt(10e6)),
		sdk.NewCoin("usdx", sdkmath.NewInt(50e6)),
	)
	totalShares := sdkmath.NewInt(30e6)

	testCases := []struct {
		shares     sdkmath.Int
		minCoinA   sdk.Coin
		minCoinB   sdk.Coin
		shouldFail bool
	}{
		{sdkmath.NewInt(15e6), sdk.NewCoin("ukava", sdkmath.NewInt(5000001)), sdk.NewCoin("usdx", sdkmath.NewInt(25e6)), true},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("shares=%s minCoinA=%s minCoinB=%s", tc.shares, tc.minCoinA, tc.minCoinB), func() {
			suite.SetupTest()
			suite.setupPool(reserves, totalShares, owner.GetAddress())

			err := suite.Keeper.Withdraw(suite.Ctx, owner.GetAddress(), tc.shares, tc.minCoinA, tc.minCoinB)
			if tc.shouldFail {
				suite.EqualError(err, "minimum withdraw not met: slippage exceeded")
			} else {
				suite.NoError(err, "expected no slippage error")
			}
		})
	}
}

func (suite *keeperTestSuite) TestWithdraw_PanicOnMissingPool() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("ukava", sdkmath.NewInt(10e6)),
		sdk.NewCoin("usdx", sdkmath.NewInt(50e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	poolID := suite.setupPool(reserves, totalShares, owner.GetAddress())

	suite.Keeper.DeletePool(suite.Ctx, poolID)

	suite.PanicsWithValue("pool ukava:usdx not found", func() {
		_ = suite.Keeper.Withdraw(suite.Ctx, owner.GetAddress(), totalShares, reserves[0], reserves[1])
	}, "expected missing pool record to panic")
}

func (suite *keeperTestSuite) TestWithdraw_PanicOnInvalidPool() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("ukava", sdkmath.NewInt(10e6)),
		sdk.NewCoin("usdx", sdkmath.NewInt(50e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	poolID := suite.setupPool(reserves, totalShares, owner.GetAddress())

	poolRecord, found := suite.Keeper.GetPool(suite.Ctx, poolID)
	suite.Require().True(found, "expected pool record to exist")

	poolRecord.TotalShares = sdk.ZeroInt()
	suite.Keeper.SetPool_Raw(suite.Ctx, poolRecord)

	suite.PanicsWithValue("invalid pool ukava:usdx: total shares must be greater than zero: invalid pool", func() {
		_ = suite.Keeper.Withdraw(suite.Ctx, owner.GetAddress(), totalShares, reserves[0], reserves[1])
	}, "expected invalid pool record to panic")
}

func (suite *keeperTestSuite) TestWithdraw_PanicOnModuleInsufficientFunds() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("ukava", sdkmath.NewInt(10e6)),
		sdk.NewCoin("usdx", sdkmath.NewInt(50e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	suite.setupPool(reserves, totalShares, owner.GetAddress())

	suite.RemoveCoinsFromModule(sdk.NewCoins(
		sdk.NewCoin("ukava", sdkmath.NewInt(1e6)),
		sdk.NewCoin("usdx", sdkmath.NewInt(5e6)),
	))

	suite.Panics(func() {
		_ = suite.Keeper.Withdraw(suite.Ctx, owner.GetAddress(), totalShares, reserves[0], reserves[1])
	}, "expected panic when module account does not have enough funds")
}
