package keeper_test

import (
	"errors"
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/swap/types"
)

func (suite *keeperTestSuite) TestSwapExactForTokens() {
	suite.Keeper.SetParams(suite.Ctx, types.Params{
		SwapFee: sdkmath.LegacyMustNewDecFromStr("0.0025"),
	})
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	poolID := suite.setupPool(reserves, totalShares, owner.GetAddress())

	balance := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("uoppy", sdkmath.NewInt(1e6))
	coinB := sdk.NewCoin("usdc", sdkmath.NewInt(5e6))

	err := suite.Keeper.SwapExactForTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdkmath.LegacyMustNewDecFromStr("0.01"))
	suite.Require().NoError(err)

	expectedOutput := sdk.NewCoin("usdc", sdkmath.NewInt(4982529))

	suite.AccountBalanceEqual(requester.GetAddress(), balance.Sub(coinA).Add(expectedOutput))
	suite.ModuleAccountBalanceEqual(reserves.Add(coinA).Sub(expectedOutput))
	suite.PoolLiquidityEqual(reserves.Add(coinA).Sub(expectedOutput))

	suite.EventsContains(sdk.UnwrapSDKContext(suite.Ctx).EventManager().Events(), sdk.NewEvent(
		types.EventTypeSwapTrade,
		sdk.NewAttribute(types.AttributeKeyPoolID, poolID),
		sdk.NewAttribute(types.AttributeKeyRequester, requester.GetAddress().String()),
		sdk.NewAttribute(types.AttributeKeySwapInput, coinA.String()),
		sdk.NewAttribute(types.AttributeKeySwapOutput, expectedOutput.String()),
		sdk.NewAttribute(types.AttributeKeyFeePaid, "2500uoppy"),
		sdk.NewAttribute(types.AttributeKeyExactDirection, "input"),
	))
}

func (suite *keeperTestSuite) TestSwapExactForBatchTokens() {
	suite.Keeper.SetParams(suite.Ctx, types.Params{
		SwapFee: sdkmath.LegacyMustNewDecFromStr("0.0025"),
	})
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	poolID := suite.setupPool(reserves, totalShares, owner.GetAddress())

	reserves2 := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("usdt", sdkmath.NewInt(5000e6)),
	)

	poolID2 := suite.setupPool(reserves2, totalShares, owner.GetAddress())

	balance := sdk.NewCoins(
		sdk.NewCoin("usdc", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)

	coinA := sdk.NewCoin("usdc", sdkmath.NewInt(1e6))
	coinB := sdk.NewCoin("usdt", sdkmath.NewInt(1e6))

	err := suite.Keeper.SwapExactForBatchTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdkmath.LegacyMustNewDecFromStr("0.01"))
	suite.Require().NoError(err)

	expectedOutput := sdk.NewCoin("usdt", sdkmath.NewInt(994607))
	intermediateOutput := sdk.NewCoin("uoppy", sdkmath.NewInt(199460))
	//

	suite.AccountBalanceEqual(requester.GetAddress(), balance.Sub(coinA).Add(expectedOutput))
	suite.ModuleAccountBalanceEqual(reserves.Add(coinA).Add(reserves2...).Sub(expectedOutput))
	suite.PoolLiquidityEqual(reserves2.Add(intermediateOutput).Sub(expectedOutput))
	suite.PoolLiquidityEqual(reserves.Add(coinA).Sub(intermediateOutput))

	suite.EventsContains(sdk.UnwrapSDKContext(suite.Ctx).EventManager().Events(), sdk.NewEvent(
		types.EventTypeSwapTrade,
		sdk.NewAttribute(types.AttributeKeyPoolID, poolID),
		sdk.NewAttribute(types.AttributeKeyRequester, requester.GetAddress().String()),
		sdk.NewAttribute(types.AttributeKeySwapInput, coinA.String()),
		sdk.NewAttribute(types.AttributeKeySwapOutput, intermediateOutput.String()),
		sdk.NewAttribute(types.AttributeKeyFeePaid, "2500usdc"),
		sdk.NewAttribute(types.AttributeKeyExactDirection, "input"),
	))

	suite.EventsContains(sdk.UnwrapSDKContext(suite.Ctx).EventManager().Events(), sdk.NewEvent(
		types.EventTypeSwapTrade,
		sdk.NewAttribute(types.AttributeKeyPoolID, poolID2),
		sdk.NewAttribute(types.AttributeKeyRequester, requester.GetAddress().String()),
		sdk.NewAttribute(types.AttributeKeySwapInput, intermediateOutput.String()),
		sdk.NewAttribute(types.AttributeKeySwapOutput, expectedOutput.String()),
		sdk.NewAttribute(types.AttributeKeyFeePaid, "499uoppy"),
		sdk.NewAttribute(types.AttributeKeyExactDirection, "input"),
	))
}

func (suite *keeperTestSuite) TestSwapExactForTokens_OutputGreaterThanZero() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(10e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(50e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	suite.setupPool(reserves, totalShares, owner.GetAddress())

	balance := sdk.NewCoins(
		sdk.NewCoin("usdc", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("usdc", sdkmath.NewInt(5))
	coinB := sdk.NewCoin("uoppy", sdkmath.NewInt(1))

	err := suite.Keeper.SwapExactForTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdkmath.LegacyMustNewDecFromStr("1"))
	suite.EqualError(err, "swap output rounds to zero, increase input amount: insufficient liquidity")
}

func (suite *keeperTestSuite) TestSwapBatchExactForTokens_OutputGreaterThanZeroFirstSwap() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(10e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(50e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	suite.setupPool(reserves, totalShares, owner.GetAddress())

	balance := sdk.NewCoins(
		sdk.NewCoin("usdc", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("usdc", sdkmath.NewInt(5))
	coinB := sdk.NewCoin("uoppy", sdkmath.NewInt(1))

	err := suite.Keeper.SwapExactForBatchTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdkmath.LegacyMustNewDecFromStr("1"))
	suite.EqualError(err, "swap1 output rounds to zero, increase input amount: insufficient liquidity")
}

func (suite *keeperTestSuite) TestSwapBatchExactForTokens_OutputGreaterThanZeroSecondSwap() {
	owner := suite.CreateAccount(sdk.Coins{})

	reserves := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	suite.setupPool(reserves, totalShares, owner.GetAddress())

	reserves2 := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("usdt", sdkmath.NewInt(1e2)),
	)

	suite.setupPool(reserves2, totalShares, owner.GetAddress())

	balance := sdk.NewCoins(
		sdk.NewCoin("usdc", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("usdc", sdkmath.NewInt(1e6))
	coinB := sdk.NewCoin("usdt", sdkmath.NewInt(1))

	err := suite.Keeper.SwapExactForBatchTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdkmath.LegacyMustNewDecFromStr("1"))
	suite.EqualError(err, "swap2 output rounds to zero, increase input amount: insufficient liquidity")
}

func (suite *keeperTestSuite) TestSwapExactBatchForTokens_Slippage() {
	owner := suite.CreateAccount(sdk.Coins{})

	reserves := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	suite.setupPool(reserves, totalShares, owner.GetAddress())

	reserves2 := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("usdt", sdkmath.NewInt(10000e6)),
	)

	suite.setupPool(reserves2, totalShares, owner.GetAddress())

	testCases := []struct {
		coinA      sdk.Coin
		coinB      sdk.Coin
		slippage   sdkmath.LegacyDec
		fee        sdkmath.LegacyDec
		shouldFail bool
	}{
		// negtive slippage OK
		{sdk.NewCoin("usdt", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(1e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("usdt", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(1.004653e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("usdt", sdkmath.NewInt(50e6)), sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("usdt", sdkmath.NewInt(50e6)), sdk.NewCoin("usdc", sdkmath.NewInt(1e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		// positive slippage with zero slippage OK
		{sdk.NewCoin("usdt", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(0.1e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("usdt", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(0.2e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("usdt", sdkmath.NewInt(50e6)), sdk.NewCoin("usdc", sdkmath.NewInt(0.9e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("usdt", sdkmath.NewInt(50e6)), sdk.NewCoin("usdc", sdkmath.NewInt(1e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		// exact zero slippage OK
		{sdk.NewCoin("usdt", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(999600)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0"), false},
		{sdk.NewCoin("usdt", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(993607)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.003"), false},
		{sdk.NewCoin("usdt", sdkmath.NewInt(5e6)), sdk.NewCoin("usdc", sdkmath.NewInt(4990014)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0"), false},
		{sdk.NewCoin("usdt", sdkmath.NewInt(5e6)), sdk.NewCoin("usdc", sdkmath.NewInt(987158)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.003"), false},
		{sdk.NewCoin("usdt", sdkmath.NewInt(5e6)), sdk.NewCoin("usdc", sdkmath.NewInt(941059)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.05"), false},
		// slippage failure, zero slippage tolerance
		{sdk.NewCoin("usdt", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(999601)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0"), true},
		{sdk.NewCoin("usdt", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(993608)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.003"), true},
		{sdk.NewCoin("usdt", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(902158)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.05"), true},
		{sdk.NewCoin("usdt", sdkmath.NewInt(5e6)), sdk.NewCoin("usdc", sdkmath.NewInt(4990015)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0"), true},
		// slippage failure, 1 percent slippage
		{sdk.NewCoin("usdt", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(5000501)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0"), true},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("coinA=%s coinB=%s slippage=%s fee=%s", tc.coinA, tc.coinB, tc.slippage, tc.fee), func() {
			suite.SetupTest()
			suite.Keeper.SetParams(suite.Ctx, types.Params{
				SwapFee: tc.fee,
			})
			owner := suite.CreateAccount(sdk.Coins{})

			reserves := sdk.NewCoins(
				sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
				sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)),
			)
			totalShares := sdkmath.NewInt(30e6)
			suite.setupPool(reserves, totalShares, owner.GetAddress())

			reserves2 := sdk.NewCoins(
				sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
				sdk.NewCoin("usdt", sdkmath.NewInt(5000e6)),
			)

			suite.setupPool(reserves2, totalShares, owner.GetAddress())

			balance := sdk.NewCoins(
				sdk.NewCoin("usdt", sdkmath.NewInt(100e6)),
				sdk.NewCoin("uoppy", sdkmath.NewInt(100e6)),
			)
			requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)

			ctx := suite.App.NewContext(true)
			err := suite.Keeper.SwapExactForBatchTokens(ctx, requester.GetAddress(), tc.coinA, tc.coinB, tc.slippage)

			if tc.shouldFail {
				suite.Require().Error(err)
				suite.Contains(err.Error(), "slippage exceeded")
			} else {
				suite.NoError(err)
			}
		})
	}
}

func (suite *keeperTestSuite) TestSwapExactForTokens_Slippage() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(100e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(500e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	suite.setupPool(reserves, totalShares, owner.GetAddress())

	testCases := []struct {
		coinA      sdk.Coin
		coinB      sdk.Coin
		slippage   sdkmath.LegacyDec
		fee        sdkmath.LegacyDec
		shouldFail bool
	}{
		// positive slippage OK
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(2e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(4e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("usdc", sdkmath.NewInt(50e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(5e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("usdc", sdkmath.NewInt(50e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		// positive slippage with zero slippage OK
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(2e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(4e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("usdc", sdkmath.NewInt(50e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(5e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("usdc", sdkmath.NewInt(50e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		// exact zero slippage OK
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(4950495)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0"), false},
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(4935790)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.003"), false},
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(4705299)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.05"), false},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(990099)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0"), false},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(987158)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.003"), false},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(941059)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.05"), false},
		// slippage failure, zero slippage tolerance
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(4950496)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0"), true},
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(4935793)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.003"), true},
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(4705300)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.05"), true},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(990100)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0"), true},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(987159)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.003"), true},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(941060)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.05"), true},
		// slippage failure, 1 percent slippage
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(5000501)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0"), true},
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(4985647)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.003"), true},
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(4752828)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.05"), true},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(1000101)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0"), true},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(997130)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.003"), true},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(950565)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.05"), true},
		// slippage OK, 1 percent slippage
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(5000500)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0"), false},
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(4985646)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.003"), false},
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdk.NewCoin("usdc", sdkmath.NewInt(4752827)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.05"), false},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(1000100)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0"), false},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(997129)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.003"), false},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(950564)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.05"), false},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("coinA=%s coinB=%s slippage=%s fee=%s", tc.coinA, tc.coinB, tc.slippage, tc.fee), func() {
			suite.SetupTest()
			suite.Keeper.SetParams(suite.Ctx, types.Params{
				SwapFee: tc.fee,
			})
			owner := suite.CreateAccount(sdk.Coins{})
			reserves := sdk.NewCoins(
				sdk.NewCoin("uoppy", sdkmath.NewInt(100e6)),
				sdk.NewCoin("usdc", sdkmath.NewInt(500e6)),
			)
			totalShares := sdkmath.NewInt(30e6)
			suite.setupPool(reserves, totalShares, owner.GetAddress())
			balance := sdk.NewCoins(
				sdk.NewCoin("uoppy", sdkmath.NewInt(100e6)),
				sdk.NewCoin("usdc", sdkmath.NewInt(100e6)),
			)
			requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)

			ctx := suite.App.NewContext(true)
			err := suite.Keeper.SwapExactForTokens(ctx, requester.GetAddress(), tc.coinA, tc.coinB, tc.slippage)

			if tc.shouldFail {
				suite.Require().Error(err)
				suite.Contains(err.Error(), "slippage exceeded")
			} else {
				suite.NoError(err)
			}
		})
	}
}

func (suite *keeperTestSuite) TestSwapExactForTokens_InsufficientFunds() {
	testCases := []struct {
		name     string
		balanceA sdk.Coin
		coinA    sdk.Coin
		coinB    sdk.Coin
	}{
		{"no uoppy balance", sdk.NewCoin("uoppy", sdkmath.ZeroInt()), sdk.NewCoin("uoppy", sdkmath.NewInt(100)), sdk.NewCoin("usdc", sdkmath.NewInt(500))},
		{"low uoppy balance", sdk.NewCoin("uoppy", sdkmath.NewInt(1000000)), sdk.NewCoin("uoppy", sdkmath.NewInt(1000001)), sdk.NewCoin("usdc", sdkmath.NewInt(5000000))},
		{"low uoppy balance", sdk.NewCoin("usdc", sdkmath.NewInt(5000000)), sdk.NewCoin("usdc", sdkmath.NewInt(5000001)), sdk.NewCoin("uoppy", sdkmath.NewInt(1000000))},
		{"large uoppy balance difference", sdk.NewCoin("uoppy", sdkmath.NewInt(100e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)), sdk.NewCoin("usdc", sdkmath.NewInt(5000e6))},
		{"large usdc balance difference", sdk.NewCoin("usdc", sdkmath.NewInt(500e6)), sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6))},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			owner := suite.CreateAccount(sdk.Coins{})
			reserves := sdk.NewCoins(
				sdk.NewCoin("uoppy", sdkmath.NewInt(100000e6)),
				sdk.NewCoin("usdc", sdkmath.NewInt(500000e6)),
			)
			totalShares := sdkmath.NewInt(30000e6)
			suite.setupPool(reserves, totalShares, owner.GetAddress())
			balance := sdk.NewCoins(tc.balanceA)
			requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)

			ctx := suite.App.NewContext(true)
			err := suite.Keeper.SwapExactForTokens(ctx, requester.GetAddress(), tc.coinA, tc.coinB, sdkmath.LegacyMustNewDecFromStr("0.1"))
			suite.Require().True(errors.Is(err, sdkerrors.ErrInsufficientFunds), fmt.Sprintf("got err %s", err))
		})
	}
}

func (suite *keeperTestSuite) TestSwapBatchExactForTokens_InsufficientFunds() {
	testCases := []struct {
		name     string
		balanceA sdk.Coin
		coinA    sdk.Coin
		coinB    sdk.Coin
	}{
		{"no usdt balance", sdk.NewCoin("usdt", sdkmath.ZeroInt()), sdk.NewCoin("usdt", sdkmath.NewInt(100)), sdk.NewCoin("usdc", sdkmath.NewInt(90))},
		{"low usdt balance", sdk.NewCoin("usdt", sdkmath.NewInt(1000000)), sdk.NewCoin("usdt", sdkmath.NewInt(1000001)), sdk.NewCoin("usdc", sdkmath.NewInt(1000000))},
		{"low usdc balance", sdk.NewCoin("usdc", sdkmath.NewInt(5000000)), sdk.NewCoin("usdc", sdkmath.NewInt(5000001)), sdk.NewCoin("usdt", sdkmath.NewInt(1000000))},
		{"large usdt balance difference", sdk.NewCoin("usdt", sdkmath.NewInt(100e6)), sdk.NewCoin("usdt", sdkmath.NewInt(1000e6)), sdk.NewCoin("usdc", sdkmath.NewInt(100e6))},
		{"large usdc balance difference", sdk.NewCoin("usdc", sdkmath.NewInt(500e6)), sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)), sdk.NewCoin("usdt", sdkmath.NewInt(900e6))},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			owner := suite.CreateAccount(sdk.Coins{})

			reserves := sdk.NewCoins(
				sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
				sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)),
			)
			totalShares := sdkmath.NewInt(30e6)
			suite.setupPool(reserves, totalShares, owner.GetAddress())

			reserves2 := sdk.NewCoins(
				sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
				sdk.NewCoin("usdt", sdkmath.NewInt(5000e6)),
			)

			suite.setupPool(reserves2, totalShares, owner.GetAddress())

			balance := sdk.NewCoins(tc.balanceA)
			requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)

			ctx := suite.App.NewContext(true)
			err := suite.Keeper.SwapExactForBatchTokens(ctx, requester.GetAddress(), tc.coinA, tc.coinB, sdkmath.LegacyMustNewDecFromStr("0.1"))
			suite.Require().True(errors.Is(err, sdkerrors.ErrInsufficientFunds), fmt.Sprintf("got err %s", err))
		})
	}
}

func (suite *keeperTestSuite) TestSwapExactForTokens_InsufficientFunds_Vesting() {
	testCases := []struct {
		name     string
		balanceA sdk.Coin
		vestingA sdk.Coin
		coinA    sdk.Coin
		coinB    sdk.Coin
	}{
		{"no uoppy balance, vesting only", sdk.NewCoin("uoppy", sdkmath.ZeroInt()), sdk.NewCoin("uoppy", sdkmath.NewInt(100)), sdk.NewCoin("uoppy", sdkmath.NewInt(100)), sdk.NewCoin("usdc", sdkmath.NewInt(500))},
		{"no usdc balance, vesting only", sdk.NewCoin("usdc", sdkmath.ZeroInt()), sdk.NewCoin("usdc", sdkmath.NewInt(500)), sdk.NewCoin("usdc", sdkmath.NewInt(500)), sdk.NewCoin("uoppy", sdkmath.NewInt(100))},
		{"low uoppy balance, vesting matches exact", sdk.NewCoin("uoppy", sdkmath.NewInt(1000000)), sdk.NewCoin("uoppy", sdkmath.NewInt(1)), sdk.NewCoin("uoppy", sdkmath.NewInt(1000001)), sdk.NewCoin("usdc", sdkmath.NewInt(5000000))},
		{"low uoppy balance, vesting matches exact", sdk.NewCoin("usdc", sdkmath.NewInt(5000000)), sdk.NewCoin("usdc", sdkmath.NewInt(1)), sdk.NewCoin("usdc", sdkmath.NewInt(5000001)), sdk.NewCoin("uoppy", sdkmath.NewInt(1000000))},
		{"large uoppy balance difference, vesting covers difference", sdk.NewCoin("uoppy", sdkmath.NewInt(100e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)), sdk.NewCoin("usdc", sdkmath.NewInt(5000e6))},
		{"large usdc balance difference, vesting covers difference", sdk.NewCoin("usdc", sdkmath.NewInt(500e6)), sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)), sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6))},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			owner := suite.CreateAccount(sdk.Coins{})
			reserves := sdk.NewCoins(
				sdk.NewCoin("uoppy", sdkmath.NewInt(100000e6)),
				sdk.NewCoin("usdc", sdkmath.NewInt(500000e6)),
			)
			totalShares := sdkmath.NewInt(30000e6)
			suite.setupPool(reserves, totalShares, owner.GetAddress())
			balance := sdk.NewCoins(tc.balanceA)
			vesting := sdk.NewCoins(tc.vestingA)
			requester := suite.CreateVestingAccount(balance, vesting)

			ctx := suite.App.NewContext(true)
			err := suite.Keeper.SwapExactForTokens(ctx, requester.GetAddress(), tc.coinA, tc.coinB, sdkmath.LegacyMustNewDecFromStr("0.1"))
			suite.Require().True(errors.Is(err, sdkerrors.ErrInsufficientFunds), fmt.Sprintf("got err %s", err))
		})
	}
}

func (suite *keeperTestSuite) TestSwapExactForTokens_PoolNotFound() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)),
	)
	totalShares := sdkmath.NewInt(3000e6)
	poolID := suite.setupPool(reserves, totalShares, owner.GetAddress())
	suite.Keeper.DeletePool(suite.Ctx, poolID)

	balance := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(10e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("uoppy", sdkmath.NewInt(1e6))
	coinB := sdk.NewCoin("usdc", sdkmath.NewInt(5e6))

	err := suite.Keeper.SwapExactForTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdkmath.LegacyMustNewDecFromStr("0.01"))
	suite.EqualError(err, "pool uoppy:usdc not found: invalid pool")

	err = suite.Keeper.SwapExactForTokens(suite.Ctx, requester.GetAddress(), coinB, coinA, sdkmath.LegacyMustNewDecFromStr("0.01"))
	suite.EqualError(err, "pool uoppy:usdc not found: invalid pool")
}

func (suite *keeperTestSuite) TestSwapExactBatchForTokens_PoolNotFound() {
	owner := suite.CreateAccount(sdk.Coins{})

	reserves := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	suite.setupPool(reserves, totalShares, owner.GetAddress())

	reserves2 := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("usdt", sdkmath.NewInt(5000e6)),
	)

	suite.setupPool(reserves2, totalShares, owner.GetAddress())

	balance := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(10e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("usdt", sdkmath.NewInt(1e6))
	coinB := sdk.NewCoin("usdf", sdkmath.NewInt(5e6))

	err := suite.Keeper.SwapExactForBatchTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdkmath.LegacyMustNewDecFromStr("0.01"))
	suite.EqualError(err, "pool uoppy:usdf not found: invalid pool")
}

func (suite *keeperTestSuite) TestSwapExactForTokens_PanicOnInvalidPool() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)),
	)
	totalShares := sdkmath.NewInt(3000e6)
	poolID := suite.setupPool(reserves, totalShares, owner.GetAddress())

	poolRecord, found := suite.Keeper.GetPool(suite.Ctx, poolID)
	suite.Require().True(found, "expected pool record to exist")

	poolRecord.TotalShares = sdkmath.ZeroInt()
	suite.Keeper.SetPool_Raw(suite.Ctx, poolRecord)

	balance := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(10e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("uoppy", sdkmath.NewInt(1e6))
	coinB := sdk.NewCoin("usdc", sdkmath.NewInt(5e6))

	suite.PanicsWithValue("invalid pool uoppy:usdc: total shares must be greater than zero: invalid pool", func() {
		_ = suite.Keeper.SwapExactForTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdkmath.LegacyMustNewDecFromStr("0.01"))
	}, "expected invalid pool record to panic")

	suite.PanicsWithValue("invalid pool uoppy:usdc: total shares must be greater than zero: invalid pool", func() {
		_ = suite.Keeper.SwapExactForTokens(suite.Ctx, requester.GetAddress(), coinB, coinA, sdkmath.LegacyMustNewDecFromStr("0.01"))
	}, "expected invalid pool record to panic")
}

func (suite *keeperTestSuite) TestSwapBatchExactForTokens_PanicOnInvalidPool() {
	owner := suite.CreateAccount(sdk.Coins{})

	reserves := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	poolID := suite.setupPool(reserves, totalShares, owner.GetAddress())

	reserves2 := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("usdt", sdkmath.NewInt(5000e6)),
	)

	poolID2 := suite.setupPool(reserves2, totalShares, owner.GetAddress())

	poolRecord, found := suite.Keeper.GetPool(suite.Ctx, poolID)
	suite.Require().True(found, "expected pool record to exist")

	poolRecord2, found := suite.Keeper.GetPool(suite.Ctx, poolID2)
	suite.Require().True(found, "expected pool record to exist")

	poolRecord.TotalShares = sdkmath.ZeroInt()
	suite.Keeper.SetPool_Raw(suite.Ctx, poolRecord)

	poolRecord.TotalShares = sdkmath.ZeroInt()
	suite.Keeper.SetPool_Raw(suite.Ctx, poolRecord2)

	balance := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(10e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("usdt", sdkmath.NewInt(1e6))
	coinB := sdk.NewCoin("usdc", sdkmath.NewInt(5e6))

	suite.PanicsWithValue("invalid pool uoppy:usdc: total shares must be greater than zero: invalid pool", func() {
		_ = suite.Keeper.SwapExactForBatchTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdkmath.LegacyMustNewDecFromStr("0.01"))
	}, "expected invalid pool record to panic")

	suite.PanicsWithValue("invalid pool uoppy:usdc: total shares must be greater than zero: invalid pool", func() {
		_ = suite.Keeper.SwapExactForBatchTokens(suite.Ctx, requester.GetAddress(), coinB, coinA, sdkmath.LegacyMustNewDecFromStr("0.01"))
	}, "expected invalid pool record to panic")
}

func (suite *keeperTestSuite) TestSwapExactForTokens_PanicOnInsufficientModuleAccFunds() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)),
	)
	totalShares := sdkmath.NewInt(3000e6)
	suite.setupPool(reserves, totalShares, owner.GetAddress())

	suite.RemoveCoinsFromModule(sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)),
	))

	balance := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(10e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("uoppy", sdkmath.NewInt(1e6))
	coinB := sdk.NewCoin("usdc", sdkmath.NewInt(5e6))

	suite.Panics(func() {
		_ = suite.Keeper.SwapExactForTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdkmath.LegacyMustNewDecFromStr("0.1"))
	}, "expected panic when module account does not have enough funds")

	suite.Panics(func() {
		_ = suite.Keeper.SwapExactForTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdkmath.LegacyMustNewDecFromStr("0.5"))
	}, "expected panic when module account does not have enough funds")
}

func (suite *keeperTestSuite) TestSwapBatchExactForTokens_PanicOnInsufficientModuleAccFunds() {
	owner := suite.CreateAccount(sdk.Coins{})

	reserves := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	suite.setupPool(reserves, totalShares, owner.GetAddress())

	reserves2 := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("usdt", sdkmath.NewInt(5000e6)),
	)

	suite.setupPool(reserves2, totalShares, owner.GetAddress())

	suite.RemoveCoinsFromModule(sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)),
	))

	balance := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(10e6)),
		sdk.NewCoin("usdt", sdkmath.NewInt(10e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("usdt", sdkmath.NewInt(1e6))
	coinB := sdk.NewCoin("usdc", sdkmath.NewInt(0.9e6))

	suite.Panics(func() {
		suite.Keeper.SwapExactForBatchTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdkmath.LegacyMustNewDecFromStr("0.1"))
	}, "expected panic when module account does not have enough funds")

	suite.Panics(func() {
		_ = suite.Keeper.SwapExactForBatchTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdkmath.LegacyMustNewDecFromStr("0.5"))
	}, "expected panic when module account does not have enough funds")

	// usdc -> usdt will not panic as the module has the balance oppy as supplied by usdt,oppy pair
	//suite.Panics(func() {
	//	err := suite.Keeper.SwapExactForBatchTokens(suite.Ctx, requester.GetAddress(), coinB, coinA, sdkmath.LegacyMustNewDecFromStr("0.9"))
	//	fmt.Printf(">>>>> err: %v\n", err)
	//}, "expected panic when module account does not have enough funds")
}

func (suite *keeperTestSuite) TestSwapForExactTokens() {
	suite.Keeper.SetParams(suite.Ctx, types.Params{
		SwapFee: sdkmath.LegacyMustNewDecFromStr("0.0025"),
	})
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	poolID := suite.setupPool(reserves, totalShares, owner.GetAddress())

	balance := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("uoppy", sdkmath.NewInt(1e6))
	coinB := sdk.NewCoin("usdc", sdkmath.NewInt(5e6))

	err := suite.Keeper.SwapForExactTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdkmath.LegacyMustNewDecFromStr("0.01"))
	suite.Require().NoError(err)

	expectedInput := sdk.NewCoin("uoppy", sdkmath.NewInt(1003511))

	suite.AccountBalanceEqual(requester.GetAddress(), balance.Sub(expectedInput).Add(coinB))
	suite.ModuleAccountBalanceEqual(reserves.Add(expectedInput).Sub(coinB))
	suite.PoolLiquidityEqual(reserves.Add(expectedInput).Sub(coinB))

	suite.EventsContains(sdk.UnwrapSDKContext(suite.Ctx).EventManager().Events(), sdk.NewEvent(
		types.EventTypeSwapTrade,
		sdk.NewAttribute(types.AttributeKeyPoolID, poolID),
		sdk.NewAttribute(types.AttributeKeyRequester, requester.GetAddress().String()),
		sdk.NewAttribute(types.AttributeKeySwapInput, expectedInput.String()),
		sdk.NewAttribute(types.AttributeKeySwapOutput, coinB.String()),
		sdk.NewAttribute(types.AttributeKeyFeePaid, "2509uoppy"),
		sdk.NewAttribute(types.AttributeKeyExactDirection, "output"),
	))
}

func (suite *keeperTestSuite) TestSwapForExactTokens_OutputLessThanPoolReserves() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(100e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(500e6)),
	)
	totalShares := sdkmath.NewInt(300e6)
	suite.setupPool(reserves, totalShares, owner.GetAddress())

	balance := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("uoppy", sdkmath.NewInt(1e6))

	coinB := sdk.NewCoin("usdc", sdkmath.NewInt(500e6).Add(sdkmath.OneInt()))
	err := suite.Keeper.SwapForExactTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdkmath.LegacyMustNewDecFromStr("0.01"))
	suite.EqualError(err, "output 500000001 >= pool reserves 500000000: insufficient liquidity")

	coinB = sdk.NewCoin("usdc", sdkmath.NewInt(500e6))
	err = suite.Keeper.SwapForExactTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdkmath.LegacyMustNewDecFromStr("0.01"))
	suite.EqualError(err, "output 500000000 >= pool reserves 500000000: insufficient liquidity")
}

func (suite *keeperTestSuite) TestSwapForExactTokens_Slippage() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(100e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(500e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	suite.setupPool(reserves, totalShares, owner.GetAddress())

	testCases := []struct {
		coinA      sdk.Coin
		coinB      sdk.Coin
		slippage   sdkmath.LegacyDec
		fee        sdkmath.LegacyDec
		shouldFail bool
	}{
		// positive slippage OK
		{sdk.NewCoin("uoppy", sdkmath.NewInt(5e6)), sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("uoppy", sdkmath.NewInt(5e6)), sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("usdc", sdkmath.NewInt(100e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(10e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("usdc", sdkmath.NewInt(100e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(10e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		// positive slippage with zero slippage OK
		{sdk.NewCoin("uoppy", sdkmath.NewInt(5e6)), sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("uoppy", sdkmath.NewInt(5e6)), sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("usdc", sdkmath.NewInt(100e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(10e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("usdc", sdkmath.NewInt(100e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(10e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.0025"), false},
		// exact zero slippage OK
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1010102)), sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0"), false},
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1010102)), sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.003"), false},
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1010102)), sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.05"), false},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5050506)), sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0"), false},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5050506)), sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.003"), false},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5050506)), sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.05"), false},
		// slippage failure, zero slippage tolerance
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1010101)), sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0"), true},
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1010101)), sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.003"), true},
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1010101)), sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.05"), true},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5050505)), sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0"), true},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5050505)), sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.003"), true},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5050505)), sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdkmath.LegacyZeroDec(), sdkmath.LegacyMustNewDecFromStr("0.05"), true},
		// slippage failure, 1 percent slippage
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1000000)), sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0"), true},
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1000000)), sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.003"), true},
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1000000)), sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.05"), true},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5000000)), sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0"), true},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5000000)), sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.003"), true},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5000000)), sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.05"), true},
		// slippage OK, 1 percent slippage
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1000001)), sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0"), false},
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1000001)), sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.003"), false},
		{sdk.NewCoin("uoppy", sdkmath.NewInt(1000001)), sdk.NewCoin("usdc", sdkmath.NewInt(5e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.05"), false},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5000001)), sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0"), false},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5000001)), sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.003"), false},
		{sdk.NewCoin("usdc", sdkmath.NewInt(5000001)), sdk.NewCoin("uoppy", sdkmath.NewInt(1e6)), sdkmath.LegacyMustNewDecFromStr("0.01"), sdkmath.LegacyMustNewDecFromStr("0.05"), false},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("coinA=%s coinB=%s slippage=%s fee=%s", tc.coinA, tc.coinB, tc.slippage, tc.fee), func() {
			suite.SetupTest()
			suite.Keeper.SetParams(suite.Ctx, types.Params{
				SwapFee: tc.fee,
			})
			owner := suite.CreateAccount(sdk.Coins{})
			reserves := sdk.NewCoins(
				sdk.NewCoin("uoppy", sdkmath.NewInt(100e6)),
				sdk.NewCoin("usdc", sdkmath.NewInt(500e6)),
			)
			totalShares := sdkmath.NewInt(30e6)
			suite.setupPool(reserves, totalShares, owner.GetAddress())
			balance := sdk.NewCoins(
				sdk.NewCoin("uoppy", sdkmath.NewInt(100e6)),
				sdk.NewCoin("usdc", sdkmath.NewInt(100e6)),
			)
			requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)

			ctx := suite.App.NewContext(true)
			err := suite.Keeper.SwapForExactTokens(ctx, requester.GetAddress(), tc.coinA, tc.coinB, tc.slippage)

			if tc.shouldFail {
				suite.Require().Error(err)
				suite.Contains(err.Error(), "slippage exceeded")
			} else {
				suite.NoError(err)
			}
		})
	}
}

func (suite *keeperTestSuite) TestSwapForExactTokens_InsufficientFunds() {
	testCases := []struct {
		name     string
		balanceA sdk.Coin
		coinA    sdk.Coin
		coinB    sdk.Coin
	}{
		{"no uoppy balance", sdk.NewCoin("uoppy", sdkmath.ZeroInt()), sdk.NewCoin("uoppy", sdkmath.NewInt(100)), sdk.NewCoin("usdc", sdkmath.NewInt(500))},
		{"no usdc balance", sdk.NewCoin("usdc", sdkmath.ZeroInt()), sdk.NewCoin("usdc", sdkmath.NewInt(500)), sdk.NewCoin("uoppy", sdkmath.NewInt(100))},
		{"low uoppy balance", sdk.NewCoin("uoppy", sdkmath.NewInt(1000000)), sdk.NewCoin("uoppy", sdkmath.NewInt(1000000)), sdk.NewCoin("usdc", sdkmath.NewInt(5000000))},
		{"low uoppy balance", sdk.NewCoin("usdc", sdkmath.NewInt(5000000)), sdk.NewCoin("usdc", sdkmath.NewInt(5000000)), sdk.NewCoin("uoppy", sdkmath.NewInt(1000000))},
		{"large uoppy balance difference", sdk.NewCoin("uoppy", sdkmath.NewInt(100e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)), sdk.NewCoin("usdc", sdkmath.NewInt(5000e6))},
		{"large usdc balance difference", sdk.NewCoin("usdc", sdkmath.NewInt(500e6)), sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6))},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			owner := suite.CreateAccount(sdk.Coins{})
			reserves := sdk.NewCoins(
				sdk.NewCoin("uoppy", sdkmath.NewInt(100000e6)),
				sdk.NewCoin("usdc", sdkmath.NewInt(500000e6)),
			)
			totalShares := sdkmath.NewInt(30000e6)
			suite.setupPool(reserves, totalShares, owner.GetAddress())
			balance := sdk.NewCoins(tc.balanceA)
			requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)

			ctx := suite.App.NewContext(true)
			err := suite.Keeper.SwapForExactTokens(ctx, requester.GetAddress(), tc.coinA, tc.coinB, sdkmath.LegacyMustNewDecFromStr("0.1"))
			suite.Require().True(errors.Is(err, sdkerrors.ErrInsufficientFunds), fmt.Sprintf("got err %s", err))
		})
	}
}

func (suite *keeperTestSuite) TestSwapForExactTokens_InsufficientFunds_Vesting() {
	testCases := []struct {
		name     string
		balanceA sdk.Coin
		vestingA sdk.Coin
		coinA    sdk.Coin
		coinB    sdk.Coin
	}{
		{"no uoppy balance, vesting only", sdk.NewCoin("uoppy", sdkmath.ZeroInt()), sdk.NewCoin("uoppy", sdkmath.NewInt(100)), sdk.NewCoin("uoppy", sdkmath.NewInt(1000)), sdk.NewCoin("usdc", sdkmath.NewInt(500))},
		{"no usdc balance, vesting only", sdk.NewCoin("usdc", sdkmath.ZeroInt()), sdk.NewCoin("usdc", sdkmath.NewInt(500)), sdk.NewCoin("usdc", sdkmath.NewInt(5000)), sdk.NewCoin("uoppy", sdkmath.NewInt(100))},
		{"low uoppy balance, vesting matches exact", sdk.NewCoin("uoppy", sdkmath.NewInt(1000000)), sdk.NewCoin("uoppy", sdkmath.NewInt(100000)), sdk.NewCoin("uoppy", sdkmath.NewInt(1000000)), sdk.NewCoin("usdc", sdkmath.NewInt(5000000))},
		{"low uoppy balance, vesting matches exact", sdk.NewCoin("usdc", sdkmath.NewInt(5000000)), sdk.NewCoin("usdc", sdkmath.NewInt(500000)), sdk.NewCoin("usdc", sdkmath.NewInt(5000000)), sdk.NewCoin("uoppy", sdkmath.NewInt(1000000))},
		{"large uoppy balance difference, vesting covers difference", sdk.NewCoin("uoppy", sdkmath.NewInt(100e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(10000e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)), sdk.NewCoin("usdc", sdkmath.NewInt(5000e6))},
		{"large usdc balance difference, vesting covers difference", sdk.NewCoin("usdc", sdkmath.NewInt(500e6)), sdk.NewCoin("usdc", sdkmath.NewInt(500000e6)), sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)), sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6))},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			owner := suite.CreateAccount(sdk.Coins{})
			reserves := sdk.NewCoins(
				sdk.NewCoin("uoppy", sdkmath.NewInt(100000e6)),
				sdk.NewCoin("usdc", sdkmath.NewInt(500000e6)),
			)
			totalShares := sdkmath.NewInt(30000e6)
			suite.setupPool(reserves, totalShares, owner.GetAddress())
			balance := sdk.NewCoins(tc.balanceA)
			vesting := sdk.NewCoins(tc.vestingA)
			requester := suite.CreateVestingAccount(balance, vesting)

			ctx := suite.App.NewContext(true)
			err := suite.Keeper.SwapForExactTokens(ctx, requester.GetAddress(), tc.coinA, tc.coinB, sdkmath.LegacyMustNewDecFromStr("0.1"))
			suite.Require().True(errors.Is(err, sdkerrors.ErrInsufficientFunds), fmt.Sprintf("got err %s", err))
		})
	}
}

func (suite *keeperTestSuite) TestSwapForExactTokens_PoolNotFound() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)),
	)
	totalShares := sdkmath.NewInt(3000e6)
	poolID := suite.setupPool(reserves, totalShares, owner.GetAddress())
	suite.Keeper.DeletePool(suite.Ctx, poolID)

	balance := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(10e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("uoppy", sdkmath.NewInt(1e6))
	coinB := sdk.NewCoin("usdc", sdkmath.NewInt(5e6))

	err := suite.Keeper.SwapForExactTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdkmath.LegacyMustNewDecFromStr("0.01"))
	suite.EqualError(err, "pool uoppy:usdc not found: invalid pool")

	err = suite.Keeper.SwapForExactTokens(suite.Ctx, requester.GetAddress(), coinB, coinA, sdkmath.LegacyMustNewDecFromStr("0.01"))
	suite.EqualError(err, "pool uoppy:usdc not found: invalid pool")
}

func (suite *keeperTestSuite) TestSwapForExactTokens_PanicOnInvalidPool() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)),
	)
	totalShares := sdkmath.NewInt(3000e6)
	poolID := suite.setupPool(reserves, totalShares, owner.GetAddress())

	poolRecord, found := suite.Keeper.GetPool(suite.Ctx, poolID)
	suite.Require().True(found, "expected pool record to exist")

	poolRecord.TotalShares = sdkmath.ZeroInt()
	suite.Keeper.SetPool_Raw(suite.Ctx, poolRecord)

	balance := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(10e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("uoppy", sdkmath.NewInt(1e6))
	coinB := sdk.NewCoin("usdc", sdkmath.NewInt(5e6))

	suite.PanicsWithValue("invalid pool uoppy:usdc: total shares must be greater than zero: invalid pool", func() {
		_ = suite.Keeper.SwapForExactTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdkmath.LegacyMustNewDecFromStr("0.01"))
	}, "expected invalid pool record to panic")

	suite.PanicsWithValue("invalid pool uoppy:usdc: total shares must be greater than zero: invalid pool", func() {
		_ = suite.Keeper.SwapForExactTokens(suite.Ctx, requester.GetAddress(), coinB, coinA, sdkmath.LegacyMustNewDecFromStr("0.01"))
	}, "expected invalid pool record to panic")
}

func (suite *keeperTestSuite) TestSwapForExactTokens_PanicOnInsufficientModuleAccFunds() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)),
	)
	totalShares := sdkmath.NewInt(3000e6)
	suite.setupPool(reserves, totalShares, owner.GetAddress())

	suite.RemoveCoinsFromModule(sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(5000e6)),
	))

	balance := sdk.NewCoins(
		sdk.NewCoin("uoppy", sdkmath.NewInt(10e6)),
		sdk.NewCoin("usdc", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("uoppy", sdkmath.NewInt(1e6))
	coinB := sdk.NewCoin("usdc", sdkmath.NewInt(5e6))

	suite.Panics(func() {
		_ = suite.Keeper.SwapForExactTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdkmath.LegacyMustNewDecFromStr("0.01"))
	}, "expected panic when module account does not have enough funds")

	suite.Panics(func() {
		_ = suite.Keeper.SwapForExactTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdkmath.LegacyMustNewDecFromStr("0.01"))
	}, "expected panic when module account does not have enough funds")
}
