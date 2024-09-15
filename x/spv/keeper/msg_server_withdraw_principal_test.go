package keeper_test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/joltify-finance/joltify_lending/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/joltify-finance/joltify_lending/utils"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/suite"
)

// Test suite used for all keeper tests
type withDrawPrincipalSuite struct {
	suite.Suite
	keeper       *spvkeeper.Keeper
	nftKeeper    types.NFTKeeper
	app          types.MsgServer
	ctx          context.Context
	investors    []string
	investorPool string
}

func setupPool(suite *withDrawPrincipalSuite) {
	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 3, PoolName: "hello", Apy: []string{"0.15", "0.12"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdkmath.NewInt(3*1e9)), sdk.NewCoin("ausdc", sdkmath.NewInt(3*1e9))}}
	resp, err := suite.app.CreatePool(suite.ctx, &req)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, resp.PoolIndex[0])
	suite.Require().True(found)
	poolInfo.PoolTotalBorrowLimit = 100
	poolInfo.TargetAmount = sdk.NewCoin("ausdc", sdkmath.NewInt(3*1e5))
	suite.keeper.SetPool(suite.ctx, poolInfo)

	suite.keeper.SetReserve(suite.ctx, sdk.NewCoin("ausdc", sdkmath.ZeroInt()))

	depositorPool := resp.PoolIndex[0]
	suite.investorPool = depositorPool

	_, err = suite.app.ActivePool(suite.ctx, types.NewMsgActivePool("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", resp.PoolIndex[0]))
	suite.Require().NoError(err)

	_, err = suite.app.ActivePool(suite.ctx, types.NewMsgActivePool("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", resp.PoolIndex[1]))
	suite.Require().NoError(err)

	req2 := types.MsgAddInvestors{
		Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: resp.PoolIndex[0],
		InvestorID: []string{"2"},
	}
	_, err = suite.app.AddInvestors(suite.ctx, &req2)
	suite.Require().NoError(err)

	// creator1 := "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl"
	// creator2 := "jolt1kkujrm0lqeu0e5va5f6mmwk87wva0k8cmam8jq"
	// creator3 := "jolt1z0y0zl0trsnuqmqf5v034pyv9sp39jg3rv6lsm"
	// creator4 := "jolt1fcaa73cc9c2l3l2u57skddgd0zm749ncukx90g"

	suite.keeper.GetPools(suite.ctx, resp.PoolIndex[0])

	suite.investors = keeper.Wallets
}

// The default state used by each test
func (suite *withDrawPrincipalSuite) SetupTest() {
	config := sdk.NewConfig()
	utils.SetBech32AddressPrefixes(config)
	lapp, k, nftKeeper, _, _, wctx := setupMsgServer(suite.T())
	ctx := sdk.UnwrapSDKContext(wctx)
	// create the first pool apy 7.8%

	k.SetParams(ctx, types.NewTestParams())

	suite.ctx = ctx
	suite.keeper = k
	suite.app = lapp
	suite.nftKeeper = nftKeeper
}

func convertBorrowToUsd(in sdkmath.Int) sdkmath.Int {
	ratio := sdkmath.LegacyMustNewDecFromStr("0.7")
	return ratio.MulInt(in).TruncateInt()
}

func convertBorrowToLocal(in sdkmath.Int) sdkmath.Int {
	ratio := sdkmath.LegacyMustNewDecFromStr("0.7")
	return (sdkmath.LegacyNewDecFromInt(in)).Quo(ratio).TruncateInt()
}

func TestWithdrawPrincipalInterest(t *testing.T) {
	suite.Run(t, new(withDrawPrincipalSuite))
}

func (suite *withDrawPrincipalSuite) TestMsgWithdrawPrincipalTest() {
	setupPool(suite)
	// now we deposit some token and it should be enough to borrow
	creator1 := suite.investors[0]
	creator2 := suite.investors[1]
	creatorAddr1, err := sdk.AccAddressFromBech32(creator1)
	suite.Require().NoError(err)
	creatorAddr2, err := sdk.AccAddressFromBech32(creator2)
	suite.Require().NoError(err)
	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(4e5))
	// suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{
		Creator:   creator1,
		PoolIndex: suite.investorPool,
		Token:     depositAmount,
	}

	// user two deposit half of the amount of the user 1
	msgDepositUser2 := &types.MsgDeposit{
		Creator:   creator2,
		PoolIndex: suite.investorPool,
		Token:     depositAmount.SubAmount(sdkmath.NewInt(2e5)),
	}

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	// now we borrow 1.34e5
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)
	totalBorrowed := borrow.BorrowAmount

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.Require().True(borrow.BorrowAmount.Amount.Sub(convertBorrowToUsd(poolInfo.BorrowedAmount.Amount)).Abs().LT(sdkmath.NewInt(2)))

	borrowable := sdkmath.NewIntFromUint64(6e5).Sub(sdkmath.NewIntFromUint64(1.34e5))
	suite.Require().EqualValues(borrowable, poolInfo.UsableAmount.Amount)
	suite.Require().True(convertBorrowToUsd(poolInfo.BorrowedAmount.Amount).Sub(sdkmath.NewIntFromUint64(1.34e5)).Abs().LT(sdkmath.NewIntFromUint64(2)))

	depositor1, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	depositor2, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)

	totalWithdrawbleFromInvestor := depositor1.WithdrawalAmount.Add(depositor2.WithdrawalAmount)
	suite.Require().EqualValues(totalWithdrawbleFromInvestor.Amount, poolInfo.UsableAmount.Amount)

	suite.Require().EqualValues(sdkmath.NewInt(4e5), depositor1.WithdrawalAmount.AddAmount(convertBorrowToUsd(depositor1.LockedAmount.Amount)).Amount)
	suite.Require().EqualValues(sdkmath.NewInt(1.99999e5), depositor2.WithdrawalAmount.AddAmount(convertBorrowToUsd(depositor2.LockedAmount.Amount)).Amount)

	getRatio := sdkmath.LegacyNewDecFromInt(depositor1.LockedAmount.Amount).Quo(sdkmath.LegacyNewDecFromInt(depositor2.LockedAmount.Amount))
	// the ratio should close to 2
	suite.Require().True(getRatio.Sub(sdkmath.LegacyMustNewDecFromStr("2")).LTE(sdkmath.LegacyNewDecWithPrec(1, 4)))

	// now we borrow more money
	borrow.BorrowAmount = sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.1e5))
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)
	totalBorrowed = totalBorrowed.Add(borrow.BorrowAmount)

	borrowable = borrowable.Sub(borrow.BorrowAmount.Amount)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.Require().EqualValues(borrowable, poolInfo.UsableAmount.Amount)

	depositor1, found = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	depositor2, found = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)
	totalWithdrawbleFromInvestor = depositor1.WithdrawalAmount.Add(depositor2.WithdrawalAmount)
	suite.Require().EqualValues(totalWithdrawbleFromInvestor.Amount, poolInfo.UsableAmount.Amount)
	suite.Require().True(sdkmath.NewInt(4e5).Sub(depositor1.WithdrawalAmount.AddAmount(convertBorrowToUsd(depositor1.LockedAmount.Amount)).Amount).Abs().LTE(sdkmath.NewIntFromUint64(1)))
	suite.Require().True(sdkmath.NewInt(2e5).Sub(depositor2.WithdrawalAmount.AddAmount(convertBorrowToUsd(depositor2.LockedAmount.Amount)).Amount).Abs().LTE(sdkmath.NewIntFromUint64(1)))

	getRatio = sdkmath.LegacyNewDecFromInt(depositor1.LockedAmount.Amount).Quo(sdkmath.LegacyNewDecFromInt(depositor2.LockedAmount.Amount))
	// the ratio should close to 2
	suite.Require().True(getRatio.Sub(sdkmath.LegacyMustNewDecFromStr("2")).LTE(sdkmath.LegacyNewDecWithPrec(1, 4)))

	// now we run the withdrawal
	withdrawReq := types.MsgWithdrawPrincipal{Creator: "invalid"}
	_, err = suite.app.WithdrawPrincipal(suite.ctx, &withdrawReq)
	suite.Require().ErrorContains(err, "invalid address")

	withdrawReq = types.MsgWithdrawPrincipal{Creator: suite.investors[0], PoolIndex: "invalid"}
	_, err = suite.app.WithdrawPrincipal(suite.ctx, &withdrawReq)
	suite.Require().ErrorContains(err, "depositor not found for pool")

	withdrawReq.PoolIndex = suite.investorPool
	withdrawReq.Token = sdk.NewCoin("invalid", sdkmath.NewIntFromUint64(22))
	_, err = suite.app.WithdrawPrincipal(suite.ctx, &withdrawReq)
	suite.Require().ErrorContains(err, "you can only withdraw ausdc")

	getRatio = sdkmath.LegacyNewDecFromInt(depositor1.LockedAmount.Amount).Quo(sdkmath.LegacyNewDecFromInt(depositor2.LockedAmount.Amount))
	// the ratio should close to 2

	depositor1, found = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	failedAmount := depositor1.WithdrawalAmount.SubAmount(sdkmath.NewIntFromUint64(3))
	withdrawReq.Token = sdk.NewCoin(depositor1.WithdrawalAmount.Denom, failedAmount.Amount)
	_, err = suite.app.WithdrawPrincipal(suite.ctx, &withdrawReq)
	suite.Require().ErrorContains(err, "deposit amount 3 is less than the minimum deposit amount 5")

	suite.Require().True(getRatio.Sub(sdkmath.LegacyMustNewDecFromStr("2")).LTE(sdkmath.LegacyNewDecWithPrec(1, 4)))

	withdrawReq.Token = sdk.NewCoin(depositor1.WithdrawalAmount.Denom, sdkmath.NewIntFromUint64(100))
	_, err = suite.app.WithdrawPrincipal(suite.ctx, &withdrawReq)
	suite.Require().NoError(err)
	before := depositor1.WithdrawalAmount.Amount
	beforePool := poolInfo.UsableAmount.Amount
	depositor1, found = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)
	suite.Require().True(before.Sub(depositor1.WithdrawalAmount.Amount).Equal(sdkmath.NewIntFromUint64(100)))
	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.Require().True(beforePool.Sub(poolInfo.UsableAmount.Amount).Equal(sdkmath.NewIntFromUint64(100)))
}

func (suite *withDrawPrincipalSuite) TestWithdrawPrincipalWithLiquidationMultipleBorrow() {
	setupPool(suite)
	creator1 := suite.investors[0]
	creator2 := suite.investors[1]
	creatorAddr1, err := sdk.AccAddressFromBech32(creator1)
	suite.Require().NoError(err)
	creatorAddr2, err := sdk.AccAddressFromBech32(creator2)
	suite.Require().NoError(err)
	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(4e5))
	// suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{
		Creator:   creator1,
		PoolIndex: suite.investorPool,
		Token:     depositAmount,
	}

	// user two deposit half of the amount of the user 1
	msgDepositUser2 := &types.MsgDeposit{
		Creator:   creator2,
		PoolIndex: suite.investorPool,
		Token:     depositAmount.SubAmount(sdkmath.NewInt(2e5)),
	}

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	// now we borrow 1.34e5
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	// we borrow again
	borrow.BorrowAmount = sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1e5))
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Second * 20))

	poolInfo.PoolStatus = types.PoolInfo_Liquidation
	suite.keeper.SetPool(suite.ctx, poolInfo)

	samples := make([]int, 20)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		// Generate a random integer between 0 and remainingSum
		randomInt := rand.Intn(9) + 1
		samples[i] = randomInt
	}

	amount := sdkmath.NewIntFromUint64(uint64(samples[0])).Mul(sdkmath.NewIntFromUint64(1e2))
	_, err = suite.app.Liquidate(suite.ctx, &types.MsgLiquidate{Creator: suite.investors[1], PoolIndex: suite.investorPool, Amount: sdk.NewCoin("ausdc", amount)})
	suite.Require().NoError(err)
	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Hour))

	depositorInfo1, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	depositorInfo2, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)

	expectedLiquidationToUser1 := sdkmath.LegacyNewDecFromInt(amount.Mul(depositorInfo1.LockedAmount.Amount)).Quo(sdkmath.LegacyNewDecFromInt(poolInfo.BorrowedAmount.Amount)).TruncateInt()
	expectedLiquidationToUser2 := sdkmath.LegacyNewDecFromInt(amount.Mul(depositorInfo2.LockedAmount.Amount)).Quo(sdkmath.LegacyNewDecFromInt(poolInfo.BorrowedAmount.Amount)).TruncateInt()

	totalExpected1 := depositorInfo1.WithdrawalAmount.AddAmount(expectedLiquidationToUser1)
	totalExpected2 := depositorInfo2.WithdrawalAmount.AddAmount(expectedLiquidationToUser2)

	resp1, err := suite.app.WithdrawPrincipal(suite.ctx, &types.MsgWithdrawPrincipal{Creator: suite.investors[0], PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1))})
	suite.Require().NoError(err)

	resp2, err := suite.app.WithdrawPrincipal(suite.ctx, &types.MsgWithdrawPrincipal{Creator: suite.investors[1], PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1e9))})
	suite.Require().NoError(err)

	// start from 1
	real1, err := sdk.ParseCoinNormalized(resp1.Amount)
	suite.Require().NoError(err)
	real2, err := sdk.ParseCoinNormalized(resp2.Amount)
	suite.Require().NoError(err)

	suite.Require().True(totalExpected1.Amount.Sub(real1.Amount).Abs().LT(sdkmath.NewIntFromUint64(10)))
	suite.Require().True(totalExpected2.Amount.Sub(real2.Amount).Abs().LT(sdkmath.NewIntFromUint64(10)))

	expectedLiquidationToUser1 = sdkmath.ZeroInt()
	expectedLiquidationToUser2 = sdkmath.ZeroInt()
	for i := 1; i < 5; i++ {
		amount := sdkmath.NewIntFromUint64(uint64(samples[i])).Mul(sdkmath.NewIntFromUint64(1e2))
		_, err := suite.app.Liquidate(suite.ctx, &types.MsgLiquidate{Creator: suite.investors[1], PoolIndex: suite.investorPool, Amount: sdk.NewCoin("ausdc", amount)})
		suite.Require().NoError(err)
		suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Hour))
		expectedLiquidationToUser1 = expectedLiquidationToUser1.Add(sdkmath.LegacyNewDecFromInt(amount.Mul(depositorInfo1.LockedAmount.Amount)).Quo(sdkmath.LegacyNewDecFromInt(poolInfo.BorrowedAmount.Amount)).TruncateInt())
		expectedLiquidationToUser2 = expectedLiquidationToUser2.Add(sdkmath.LegacyNewDecFromInt(amount.Mul(depositorInfo2.LockedAmount.Amount)).Quo(sdkmath.LegacyNewDecFromInt(poolInfo.BorrowedAmount.Amount)).TruncateInt())
	}
	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Hour))
	resp1, err = suite.app.WithdrawPrincipal(suite.ctx, &types.MsgWithdrawPrincipal{Creator: suite.investors[0], PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1))})
	suite.Require().NoError(err)

	resp2, err = suite.app.WithdrawPrincipal(suite.ctx, &types.MsgWithdrawPrincipal{Creator: suite.investors[1], PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1e9))})
	suite.Require().NoError(err)

	// start from 1
	real1, err = sdk.ParseCoinNormalized(resp1.Amount)
	suite.Require().NoError(err)
	real2, err = sdk.ParseCoinNormalized(resp2.Amount)
	suite.Require().NoError(err)

	suite.Require().True(expectedLiquidationToUser1.Sub(real1.Amount).Abs().LT(sdkmath.NewIntFromUint64(10)))
	suite.Require().True(expectedLiquidationToUser2.Sub(real2.Amount).Abs().LT(sdkmath.NewIntFromUint64(10)))

	expectedLiquidationToUser1 = sdkmath.ZeroInt()
	expectedLiquidationToUser2 = sdkmath.ZeroInt()
	for i := 5; i < 20; i++ {
		amount := sdkmath.NewIntFromUint64(uint64(samples[i])).Mul(sdkmath.NewIntFromUint64(1e2))
		_, err := suite.app.Liquidate(suite.ctx, &types.MsgLiquidate{Creator: suite.investors[1], PoolIndex: suite.investorPool, Amount: sdk.NewCoin("ausdc", amount)})
		suite.Require().NoError(err)
		suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Hour))
		expectedLiquidationToUser1 = expectedLiquidationToUser1.Add(sdkmath.LegacyNewDecFromInt(amount.Mul(depositorInfo1.LockedAmount.Amount)).Quo(sdkmath.LegacyNewDecFromInt(poolInfo.BorrowedAmount.Amount)).TruncateInt())
		expectedLiquidationToUser2 = expectedLiquidationToUser2.Add(sdkmath.LegacyNewDecFromInt(amount.Mul(depositorInfo2.LockedAmount.Amount)).Quo(sdkmath.LegacyNewDecFromInt(poolInfo.BorrowedAmount.Amount)).TruncateInt())
	}

	resp1, err = suite.app.WithdrawPrincipal(suite.ctx, &types.MsgWithdrawPrincipal{Creator: suite.investors[0], PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1))})
	suite.Require().NoError(err)

	resp2, err = suite.app.WithdrawPrincipal(suite.ctx, &types.MsgWithdrawPrincipal{Creator: suite.investors[1], PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1e9))})
	suite.Require().NoError(err)

	// start from 1
	real1, err = sdk.ParseCoinNormalized(resp1.Amount)
	suite.Require().NoError(err)
	real2, err = sdk.ParseCoinNormalized(resp2.Amount)
	suite.Require().NoError(err)

	suite.Require().True(expectedLiquidationToUser1.Sub(real1.Amount).Abs().LTE(sdkmath.NewIntFromUint64(12)))
	suite.Require().True(expectedLiquidationToUser2.Sub(real2.Amount).Abs().LTE(sdkmath.NewIntFromUint64(12)))
}

func (suite *withDrawPrincipalSuite) TestWithdrawPrincipalWithLiquidation() {
	setupPool(suite)
	creator1 := suite.investors[0]
	creator2 := suite.investors[1]
	creatorAddr1, err := sdk.AccAddressFromBech32(creator1)
	suite.Require().NoError(err)
	creatorAddr2, err := sdk.AccAddressFromBech32(creator2)
	suite.Require().NoError(err)
	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(4e5))
	// suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{
		Creator:   creator1,
		PoolIndex: suite.investorPool,
		Token:     depositAmount,
	}

	// user two deposit half of the amount of the user 1
	msgDepositUser2 := &types.MsgDeposit{
		Creator:   creator2,
		PoolIndex: suite.investorPool,
		Token:     depositAmount.SubAmount(sdkmath.NewInt(2e5)),
	}

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	// now we borrow 1.34e5
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Second * 20))

	poolInfo.PoolStatus = types.PoolInfo_Liquidation
	suite.keeper.SetPool(suite.ctx, poolInfo)

	realPaid := sdkmath.ZeroInt()

	samples := make([]int, 20)
	rand.Seed(time.Now().UnixNano())
	a := 0
	for i := 0; i < 20; i++ {
		// Generate a random integer between 0 and remainingSum
		randomInt := rand.Intn(9) + 1
		samples[i] = randomInt
		a += randomInt
	}

	amount := sdkmath.NewIntFromUint64(uint64(samples[0])).Mul(sdkmath.NewIntFromUint64(1e2))
	_, err = suite.app.Liquidate(suite.ctx, &types.MsgLiquidate{Creator: suite.investors[1], PoolIndex: suite.investorPool, Amount: sdk.NewCoin("ausdc", amount)})
	suite.Require().NoError(err)
	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Hour))

	depositorInfo1, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	depositorInfo2, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)

	expectedLiquidationToUser1 := sdkmath.LegacyNewDecFromInt(amount.Mul(depositorInfo1.LockedAmount.Amount)).Quo(sdkmath.LegacyNewDecFromInt(poolInfo.BorrowedAmount.Amount)).TruncateInt()
	expectedLiquidationToUser2 := sdkmath.LegacyNewDecFromInt(amount.Mul(depositorInfo2.LockedAmount.Amount)).Quo(sdkmath.LegacyNewDecFromInt(poolInfo.BorrowedAmount.Amount)).TruncateInt()

	totalExpected1 := depositorInfo1.WithdrawalAmount.AddAmount(expectedLiquidationToUser1)
	totalExpected2 := depositorInfo2.WithdrawalAmount.AddAmount(expectedLiquidationToUser2)

	resp1, err := suite.app.WithdrawPrincipal(suite.ctx, &types.MsgWithdrawPrincipal{Creator: suite.investors[0], PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1))})
	suite.Require().NoError(err)

	resp2, err := suite.app.WithdrawPrincipal(suite.ctx, &types.MsgWithdrawPrincipal{Creator: suite.investors[1], PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1e9))})
	suite.Require().NoError(err)

	// start from 1
	real1, err := sdk.ParseCoinNormalized(resp1.Amount)
	suite.Require().NoError(err)

	// we cannot use the resp1 as it includes the withdrawal amount
	realPaid = realPaid.Add(expectedLiquidationToUser1)
	real2, err := sdk.ParseCoinNormalized(resp2.Amount)
	suite.Require().NoError(err)
	realPaid = realPaid.Add(expectedLiquidationToUser2)
	suite.Require().True(totalExpected1.Amount.Equal(real1.Amount))
	suite.Require().True(totalExpected2.Amount.Equal(real2.Amount))

	expectedLiquidationToUser1 = sdkmath.ZeroInt()
	expectedLiquidationToUser2 = sdkmath.ZeroInt()
	for i := 1; i < 5; i++ {
		amount := sdkmath.NewIntFromUint64(uint64(samples[i])).Mul(sdkmath.NewIntFromUint64(1e2))
		_, err := suite.app.Liquidate(suite.ctx, &types.MsgLiquidate{Creator: suite.investors[1], PoolIndex: suite.investorPool, Amount: sdk.NewCoin("ausdc", amount)})
		suite.Require().NoError(err)
		suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Hour))
		expectedLiquidationToUser1 = expectedLiquidationToUser1.Add(sdkmath.LegacyNewDecFromInt(amount.Mul(depositorInfo1.LockedAmount.Amount)).Quo(sdkmath.LegacyNewDecFromInt(poolInfo.BorrowedAmount.Amount)).TruncateInt())
		expectedLiquidationToUser2 = expectedLiquidationToUser2.Add(sdkmath.LegacyNewDecFromInt(amount.Mul(depositorInfo2.LockedAmount.Amount)).Quo(sdkmath.LegacyNewDecFromInt(poolInfo.BorrowedAmount.Amount)).TruncateInt())
	}
	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Hour))
	resp1, err = suite.app.WithdrawPrincipal(suite.ctx, &types.MsgWithdrawPrincipal{Creator: suite.investors[0], PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1))})
	suite.Require().NoError(err)

	resp2, err = suite.app.WithdrawPrincipal(suite.ctx, &types.MsgWithdrawPrincipal{Creator: suite.investors[1], PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1e9))})
	suite.Require().NoError(err)

	// start from 1
	real1, err = sdk.ParseCoinNormalized(resp1.Amount)
	suite.Require().NoError(err)
	real2, err = sdk.ParseCoinNormalized(resp2.Amount)
	suite.Require().NoError(err)
	realPaid = realPaid.Add(real1.Amount)
	realPaid = realPaid.Add(real2.Amount)

	suite.Require().True(expectedLiquidationToUser1.Equal(real1.Amount))
	suite.Require().True(expectedLiquidationToUser2.Equal(real2.Amount))

	expectedLiquidationToUser1 = sdkmath.ZeroInt()
	expectedLiquidationToUser2 = sdkmath.ZeroInt()
	for i := 5; i < 20; i++ {
		amount := sdkmath.NewIntFromUint64(uint64(samples[i])).Mul(sdkmath.NewIntFromUint64(1e2))
		_, err := suite.app.Liquidate(suite.ctx, &types.MsgLiquidate{Creator: suite.investors[1], PoolIndex: suite.investorPool, Amount: sdk.NewCoin("ausdc", amount)})
		suite.Require().NoError(err)
		suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Hour))
		expectedLiquidationToUser1 = expectedLiquidationToUser1.Add(sdkmath.LegacyNewDecFromInt(amount.Mul(depositorInfo1.LockedAmount.Amount)).Quo(sdkmath.LegacyNewDecFromInt(poolInfo.BorrowedAmount.Amount)).TruncateInt())
		expectedLiquidationToUser2 = expectedLiquidationToUser2.Add(sdkmath.LegacyNewDecFromInt(amount.Mul(depositorInfo2.LockedAmount.Amount)).Quo(sdkmath.LegacyNewDecFromInt(poolInfo.BorrowedAmount.Amount)).TruncateInt())
	}

	resp1, err = suite.app.WithdrawPrincipal(suite.ctx, &types.MsgWithdrawPrincipal{Creator: suite.investors[0], PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1))})
	suite.Require().NoError(err)

	resp2, err = suite.app.WithdrawPrincipal(suite.ctx, &types.MsgWithdrawPrincipal{Creator: suite.investors[1], PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1e9))})
	suite.Require().NoError(err)

	// start from 1
	real1, err = sdk.ParseCoinNormalized(resp1.Amount)
	suite.Require().NoError(err)
	real2, err = sdk.ParseCoinNormalized(resp2.Amount)
	suite.Require().NoError(err)

	realPaid = realPaid.Add(real1.Amount)
	realPaid = realPaid.Add(real2.Amount)

	suite.Require().True(expectedLiquidationToUser1.Equal(real1.Amount))
	suite.Require().True(expectedLiquidationToUser2.Equal(real2.Amount))

	nfts := poolInfo.PoolNFTIds[0]
	class, found := suite.nftKeeper.GetClass(suite.ctx, nfts)
	if !found {
		panic(found)
	}
	var borrowInterest types.BorrowInterest
	err = proto.Unmarshal(class.Data.Value, &borrowInterest)
	if err != nil {
		panic(err)
	}

	suite.Require().True(borrowInterest.TotalPaidLiquidationAmount.Equal(realPaid))
}

func (suite *withDrawPrincipalSuite) TestWithdrawPrincipalWithClosePool() {
	setupPool(suite)
	// now we deposit some token and it should be enough to borrow
	creator1 := suite.investors[0]
	creator2 := suite.investors[1]
	creatorAddr1, err := sdk.AccAddressFromBech32(creator1)
	suite.Require().NoError(err)
	// creatorAddr2, err := sdk.AccAddressFromBech32(creator2)
	// suite.Require().NoError(err)
	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(4e5))
	// suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{
		Creator:   creator1,
		PoolIndex: suite.investorPool,
		Token:     depositAmount,
	}

	// user two deposit half of the amount of the user 1
	msgDepositUser2 := &types.MsgDeposit{
		Creator:   creator2,
		PoolIndex: suite.investorPool,
		Token:     depositAmount.SubAmount(sdkmath.NewInt(2e5)),
	}

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	// now we borrow 1.34e5
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Second * 20))

	_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))})
	suite.Require().NoError(err)

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(oneYear * time.Second))
	req := types.MsgPayPrincipal{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(5e6))}
	_, err = suite.app.PayPrincipal(suite.ctx, &req)
	suite.Require().NoError(err)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Duration(poolInfo.PayFreq) * time.Second))
	suite.Require().EqualValues(poolInfo.PoolStatus, types.PoolInfo_FREEZING)
	suite.keeper.HandlePrincipalPayment(suite.ctx, &poolInfo)

	suite.ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Second * 200))

	withdrawreq1 := types.MsgWithdrawPrincipal{Creator: suite.investors[0], PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(100))}
	resp, err := suite.app.WithdrawPrincipal(suite.ctx, &withdrawreq1)
	suite.Require().NoError(err)

	_, found = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().NoError(err)
	suite.Require().False(found)
	suite.Require().EqualValues(resp.Amount, "400000ausdc")

	withdrawreq1 = types.MsgWithdrawPrincipal{Creator: suite.investors[1], PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(100))}
	resp, err = suite.app.WithdrawPrincipal(suite.ctx, &withdrawreq1)
	suite.Require().NoError(err)
	respAmount, err := sdk.ParseCoinNormalized(resp.Amount)
	suite.Require().NoError(err)

	suite.Require().True(checkValueWithRangeTwo(respAmount.Amount, sdkmath.NewIntFromUint64(2e5)))

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().False(found)
}

func (suite *withDrawPrincipalSuite) TestWithdrawWithSPVBorrowAndRepay() {
	// skip the test now
	setupPool(suite)
	// suite.keeper.SetHooks(&fakeSPVFunctions{})
	// now we deposit some token and it should be enough to borrow
	creator1 := suite.investors[0]
	creator2 := suite.investors[1]
	creatorAddr1, err := sdk.AccAddressFromBech32(creator1)
	suite.Require().NoError(err)
	creatorAddr2, err := sdk.AccAddressFromBech32(creator2)
	suite.Require().NoError(err)
	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(4e5))
	// suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{
		Creator:   creator1,
		PoolIndex: suite.investorPool,
		Token:     depositAmount,
	}

	// user two deposit half of the amount of the user 1
	msgDepositUser2 := &types.MsgDeposit{
		Creator:   creator2,
		PoolIndex: suite.investorPool,
		Token:     depositAmount.SubAmount(sdkmath.NewInt(2e5)),
	}

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	// now we borrow 1.34e5
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.Require().True(checkValueEqualWithExchange(poolInfo.BorrowedAmount.Amount, borrow.BorrowAmount.Amount))

	borrowable := sdkmath.NewIntFromUint64(6e5).Sub(sdkmath.NewIntFromUint64(1.34e5))
	suite.Require().EqualValues(borrowable, poolInfo.UsableAmount.Amount)
	suite.Require().True(checkValueEqualWithExchange(poolInfo.BorrowedAmount.Amount, sdkmath.NewIntFromUint64(1.34e5)))

	depositor1, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	depositor2, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)

	totalWithdrawbleFromInvestor := depositor1.WithdrawalAmount.Add(depositor2.WithdrawalAmount)
	suite.Require().EqualValues(totalWithdrawbleFromInvestor.Amount, poolInfo.UsableAmount.Amount)

	suite.Require().EqualValues(sdkmath.NewInt(4e5), convertBorrowToUsd(depositor1.LockedAmount.Amount).Add(depositor1.WithdrawalAmount.Amount))
	suite.Require().True(checkValueWithRangeTwo(sdkmath.NewInt(2e5), convertBorrowToUsd(depositor2.LockedAmount.Amount).Add(depositor2.WithdrawalAmount.Amount)))

	firstWithdrawable := depositor2.WithdrawalAmount
	firstWithdrawable1 := depositor1.WithdrawalAmount

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	// suite.keeper.HandlePrincipalPayment(suite.ctx, &poolInfo)
	poolInfoNew, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	withdraw := types.MsgWithdrawPrincipal{
		Creator:   suite.investors[1],
		PoolIndex: suite.investorPool,
		Token:     sdk.NewCoin("ausdc", sdkmath.ZeroInt()),
	}
	//_, err = suite.app.WithdrawPrincipal(suite.ctx, &withdraw)
	//suite.Require().NoError(err)
	depositor2, found = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)
	delta := depositor2.WithdrawalAmount.Sub(firstWithdrawable)

	totalMoney2 := depositor2.WithdrawalAmount.AddAmount(convertBorrowToUsd(depositor2.LockedAmount.Amount))
	suite.Require().True(checkValueWithRangeTwo(sdkmath.NewIntFromUint64(2e5), totalMoney2.Amount))

	depositor1, found = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)
	delta2 := depositor1.WithdrawalAmount.Sub(firstWithdrawable1)

	totalMoney1 := depositor1.WithdrawalAmount.AddAmount(convertBorrowToUsd(depositor1.LockedAmount.Amount))
	suite.Require().True(checkValueWithRangeTwo(sdkmath.NewIntFromUint64(4e5), totalMoney1.Amount))

	// we have another two borrows

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	borrow.BorrowAmount = sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(4e3))
	// now we borrow 4e3
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	borrow.BorrowAmount = sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(2e3))
	// now we borrow 2e3
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	poolInfoNew, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	nfts := poolInfoNew.PoolNFTIds[1]
	class, found := suite.nftKeeper.GetClass(suite.ctx, nfts)
	if !found {
		panic(found)
	}
	var borrowInterest types.BorrowInterest
	err = proto.Unmarshal(class.Data.Value, &borrowInterest)
	if err != nil {
		panic(err)
	}

	// now we verify the depositor 1 and 2
	depositorBeforeWithdraw2, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)
	depositorBeforeWithdraw1, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	depositor2, found = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)

	delta = depositor2.WithdrawalAmount.Sub(depositorBeforeWithdraw2.WithdrawalAmount)

	totalMoney2 = depositor2.WithdrawalAmount.AddAmount(convertBorrowToUsd(depositor2.LockedAmount.Amount))
	suite.Require().True(checkValueWithRangeTwo(sdkmath.NewIntFromUint64(2e5), totalMoney2.Amount))

	depositor1, found = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)
	delta2 = depositor1.WithdrawalAmount.Sub(depositorBeforeWithdraw1.WithdrawalAmount)

	totalMoney1 = depositor1.WithdrawalAmount.AddAmount(convertBorrowToUsd(depositor1.LockedAmount.Amount))
	suite.Require().True(checkValueWithRangeTwo(sdkmath.NewIntFromUint64(4e5), totalMoney1.Amount))

	allWithdrawAbs := delta.Add(delta2).Amount.Abs()
	// as we have two investors, so max of each extra pay is 1, in total is 2 at most
	suite.Require().True(allWithdrawAbs.Sub(sdkmath.NewIntFromUint64(1e2)).LT(sdkmath.NewIntFromUint64(3)))

	// poolInfoBeforePayAll, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	// suite.Require().True(found)

	//
	//
	//
	_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1e5))})
	suite.Require().NoError(err)
	// now we pay all the money
	req := types.MsgPayPrincipal{
		Creator:   "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0",
		PoolIndex: suite.investorPool,
		Token:     sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1e10)),
	}
	_, err = suite.app.PayPrincipal(suite.ctx, &req)
	suite.Require().NoError(err)

	poolInfoNew, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	suite.keeper.HandlePrincipalPayment(suite.ctx, &poolInfoNew)
	poolInfoNew, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	// now we verify the depositor 1 and 2

	depositorBeforeWithdraw2, found = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr2)
	suite.Require().True(found)

	depositorBeforeWithdraw1, found = suite.keeper.GetDepositor(suite.ctx, suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	withdraw.Creator = suite.investors[1]
	withdraw.Token = withdraw.Token.AddAmount(sdkmath.OneInt())

	resp, err := suite.app.WithdrawPrincipal(suite.ctx, &withdraw)
	suite.Require().NoError(err)

	depositor2, found = suite.keeper.GetDepositorHistory(suite.ctx, sdk.UnwrapSDKContext(suite.ctx).BlockTime(), suite.investorPool, creatorAddr2)
	suite.Require().True(found)

	suite.Require().True(depositor2.LockedAmount.Amount.Equal(sdkmath.ZeroInt()))
	// the original v2 deposit is 2e5

	coin, err := sdk.ParseCoinNormalized(resp.Amount)
	suite.Require().NoError(err)

	suite.Require().True(checkValueWithRangeTwo(coin.Amount, sdkmath.NewIntFromUint64(2e5)))

	withdraw.Creator = suite.investors[0]
	resp, err = suite.app.WithdrawPrincipal(suite.ctx, &withdraw)
	suite.Require().NoError(err)

	depositor1, found = suite.keeper.GetDepositorHistory(suite.ctx, sdk.UnwrapSDKContext(suite.ctx).BlockTime(), suite.investorPool, creatorAddr1)
	suite.Require().True(found)

	suite.Require().True(depositor1.LockedAmount.Amount.Equal(sdkmath.ZeroInt()))
	// the original v2 deposit is 2e5

	coin, err = sdk.ParseCoinNormalized(resp.Amount)
	suite.Require().NoError(err)
	suite.Require().True(checkValueWithRangeTwo(coin.Amount, sdkmath.NewIntFromUint64(4e5)))
}
