package keeper_test

import (
	"math/rand"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/suite"
)

// Test suite used for all keeper tests
type liquidateTestSuite struct {
	suite.Suite
	keeper       *spvkeeper.Keeper
	nftKeeper    types.NFTKeeper
	app          types.MsgServer
	ctx          sdk.Context
	investorPool string
	investors    []string
}

func TestLiquidateSuite(t *testing.T) {
	suite.Run(t, new(liquidateTestSuite))
}

// The default state used by each test
func (suite *liquidateTestSuite) SetupTest() {
	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)

	app, k, nftKeeper, _, _, wctx := setupMsgServer(suite.T())
	ctx := sdk.UnwrapSDKContext(wctx)

	suite.ctx = ctx
	suite.keeper = k
	suite.nftKeeper = nftKeeper
	suite.app = app
}

func setupLiquidateEnv(suite *liquidateTestSuite) {
	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 3, PoolName: "hello", Apy: []string{"0.15", "0.12"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdk.NewInt(4*1e5)), sdk.NewCoin("ausdc", sdk.NewInt(4*1e5))}}
	resp, err := suite.app.CreatePool(suite.ctx, &req)
	suite.Require().NoError(err)

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

	creator1 := "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl"
	creator2 := "jolt1kkujrm0lqeu0e5va5f6mmwk87wva0k8cmam8jq"

	depositAmount := sdk.NewCoin("ausdc", sdk.NewInt(4e5))
	// suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{
		Creator:   creator1,
		PoolIndex: suite.investorPool,
		Token:     depositAmount,
	}

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	poolInfo.PoolTotalBorrowLimit = 100
	poolInfo.TargetAmount = sdk.NewCoin("ausdc", sdk.NewInt(4e5))
	suite.keeper.SetPool(suite.ctx, poolInfo)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: depositorPool, BorrowAmount: sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1.34e5))}

	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	poolInfo.PoolStatus = types.PoolInfo_Liquidation
	poolInfo.PoolTotalBorrowLimit = 100
	suite.keeper.SetPool(suite.ctx, poolInfo)

	suite.investors = []string{creator1, creator2}
}

func (suite *liquidateTestSuite) TestLiquidate() {
	setupLiquidateEnv(suite)
	type args struct {
		msgLiquidate *types.MsgLiquidate
		expectedErr  string
	}

	type test struct {
		name string
		args args
	}

	testCases := []test{
		{
			name: "invalid address",
			args: args{msgLiquidate: &types.MsgLiquidate{Creator: "invalid address"}, expectedErr: "invalid address invalid address: invalid address"},
		},

		{
			name: "amount cannot be zero",
			args: args{msgLiquidate: &types.MsgLiquidate{Creator: suite.investors[0], PoolIndex: suite.investorPool, Amount: sdk.NewCoin("abc", sdk.ZeroInt())}, expectedErr: "the amount cannot be zero"},
		},
		{
			name: "pool cannot be found",
			args: args{msgLiquidate: &types.MsgLiquidate{Creator: suite.investors[1], PoolIndex: "invalid", Amount: sdk.NewCoin("ausdc", sdk.OneInt())}, expectedErr: "pool cannot be found invalid"},
		},

		{
			name: "inconsistent demon",
			args: args{msgLiquidate: &types.MsgLiquidate{Creator: suite.investors[1], PoolIndex: suite.investorPool, Amount: sdk.NewCoin("abc", sdk.OneInt())}, expectedErr: "the token is not the same as the borrowed token"},
		},
		{
			name: "success",
			args: args{msgLiquidate: &types.MsgLiquidate{Creator: suite.investors[1], PoolIndex: suite.investorPool, Amount: sdk.NewCoin("ausdc", sdk.OneInt())}, expectedErr: ""},
		},

		{
			name: "pool is not in liquidation",
			args: args{msgLiquidate: &types.MsgLiquidate{Creator: suite.investors[1], PoolIndex: suite.investorPool, Amount: sdk.NewCoin("ausdc", sdk.OneInt())}, expectedErr: "pool is not in liquidation"},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			if tc.name == "pool is not in liquidation" {
				poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
				suite.Require().True(found)
				poolInfo.PoolStatus = types.PoolInfo_FREEZING
				poolInfo.PoolTotalBorrowLimit = 100
				suite.keeper.SetPool(suite.ctx, poolInfo)
				_, err := suite.app.Liquidate(suite.ctx, tc.args.msgLiquidate)
				if tc.args.expectedErr != "" {
					suite.Require().ErrorContains(err, tc.args.expectedErr)
				} else {
					suite.Require().NoError(err)
				}
			}

			_, err := suite.app.Liquidate(suite.ctx, tc.args.msgLiquidate)
			if tc.args.expectedErr != "" {
				suite.Require().ErrorContains(err, tc.args.expectedErr)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *liquidateTestSuite) TestLiquidateWithPaymentCheckSignleBorrow() {
	setupLiquidateEnv(suite)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	poolInfo.PoolTotalBorrowLimit = 100
	poolInfo.TargetAmount = sdk.NewCoin("ausdc", sdk.NewInt(600000))
	suite.keeper.SetPool(suite.ctx, poolInfo)

	samples := make([]int, 20)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		// Generate a random integer between 0 and remainingSum
		randomInt := rand.Intn(9) + 1
		samples[i] = randomInt
	}

	for i := 0; i < 20; i++ {
		amount := sdk.NewIntFromUint64(uint64(samples[i])).Mul(sdk.NewIntFromUint64(1e2))
		_, err := suite.app.Liquidate(suite.ctx, &types.MsgLiquidate{Creator: suite.investors[1], PoolIndex: suite.investorPool, Amount: sdk.NewCoin("ausdc", amount)})
		suite.Require().NoError(err)
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Hour))
	}

	histories := make([][]*types.LiquidationItem, 1)
	for i, el := range poolInfo.PoolNFTIds {
		class, found := suite.nftKeeper.GetClass(suite.ctx, el)
		if !found {
			panic(found)
		}
		var borrowInterest types.BorrowInterest
		var err error
		err = proto.Unmarshal(class.Data.Value, &borrowInterest)
		if err != nil {
			panic(err)
		}
		history := borrowInterest.LiquidationItems
		histories[i] = history
	}

	for i := 0; i < 20; i++ {
		total := sdk.NewIntFromUint64(uint64(samples[i])).Mul(sdk.NewIntFromUint64(1e2))
		suite.Require().True(total.Equal(histories[0][i].Amount.Amount))
	}
}

func (suite *liquidateTestSuite) TestLiquidateWithPaymentCheckTwoBorrow() {
	setupLiquidateEnv(suite)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	poolInfo.PoolTotalBorrowLimit = 100
	poolInfo.TargetAmount = sdk.NewCoin("ausdc", sdk.NewInt(4e5))
	suite.keeper.SetPool(suite.ctx, poolInfo)

	depositorPool := suite.investorPool

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: depositorPool, BorrowAmount: sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1.34e5))}

	_, err := suite.app.Borrow(suite.ctx, borrow)
	suite.Require().ErrorContains(err, "pool is not in active status")

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	poolInfo.PoolStatus = types.PoolInfo_ACTIVE
	poolInfo.PoolTotalBorrowLimit = 100
	suite.keeper.SetPool(suite.ctx, poolInfo)

	borrow.BorrowAmount = sdk.NewCoin("ausdc", sdk.NewIntFromUint64(2e5))
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.Require().True(checkValueEqualWithExchange(poolInfo.BorrowedAmount.Amount, sdk.NewIntFromUint64(3.34e5)))
	poolInfo.PoolStatus = types.PoolInfo_Liquidation
	suite.keeper.SetPool(suite.ctx, poolInfo)

	// we mock the liquidation payments for 20 times

	samples := make([]int, 20)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		// Generate a random integer between 0 and remainingSum
		randomInt := rand.Intn(9) + 1
		samples[i] = randomInt
	}

	for i := 0; i < 20; i++ {
		amount := sdk.NewIntFromUint64(uint64(samples[i])).Mul(sdk.NewIntFromUint64(1e2))
		_, err := suite.app.Liquidate(suite.ctx, &types.MsgLiquidate{Creator: suite.investors[1], PoolIndex: suite.investorPool, Amount: sdk.NewCoin("ausdc", amount)})
		suite.Require().NoError(err)
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Hour))
	}

	histories := make([][]*types.LiquidationItem, 2)
	for i, el := range poolInfo.PoolNFTIds {
		class, found := suite.nftKeeper.GetClass(suite.ctx, el)
		if !found {
			panic(found)
		}
		var borrowInterest types.BorrowInterest
		var err error
		err = proto.Unmarshal(class.Data.Value, &borrowInterest)
		if err != nil {
			panic(err)
		}
		history := borrowInterest.LiquidationItems
		histories[i] = history
	}

	for i := 0; i < 20; i++ {
		total := sdk.NewIntFromUint64(uint64(samples[i])).Mul(sdk.NewIntFromUint64(1e2))
		v2 := sdk.NewDecFromInt(total.Mul(sdk.NewIntFromUint64(2e5))).Quo(sdk.NewDecFromInt(sdk.NewIntFromUint64(3.34e5))).TruncateInt()
		suite.Require().True(v2.Equal(histories[1][i].Amount.Amount))
		v1 := total.Sub(v2)
		suite.Require().True(v1.Equal(histories[0][i].Amount.Amount))
	}
}
