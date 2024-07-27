package keeper_test

import (
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/joltify-finance/joltify_lending/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	"github.com/joltify-finance/joltify_lending/x/spv"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/suite"
)

// Test suite used for all keeper tests
type mockBurnSuite struct {
	suite.Suite
	keeper        *spvkeeper.Keeper
	nftKeeper     types.NFTKeeper
	auctionKeeper keeper.MockAuctionKeeper
	app           types.MsgServer
	ctx           context.Context
	investors     []string
	investorPool  string
	creator       string
}

func lsetupMockPool(suite *mockBurnSuite) {
	// create the first pool apy 7.8%
	// senior pool is 300,000
	amount := new(big.Int).Mul(big.NewInt(200000), base)
	amountSenior := new(big.Int).Mul(big.NewInt(800000), base)
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 4, PoolName: "hello", Apy: []string{"0.15", "0.0875"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdkmath.NewIntFromBigInt(amount)), sdk.NewCoin("ausdc", sdkmath.NewIntFromBigInt(amountSenior))}}
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

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)

	poolInfo.PoolTotalBorrowLimit = 100
	poolInfo.PoolLockedSeconds = 360000
	// we separate the pool to make sure we can borrow from the senior pool independently
	poolInfo.SeparatePool = true
	suite.keeper.SetPool(suite.ctx, poolInfo)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, resp.PoolIndex[1])
	suite.Require().True(found)

	poolInfo.PoolTotalBorrowLimit = 100
	poolInfo.PoolLockedSeconds = 360000
	// we separate the pool to make sure we can borrow from the senior pool independently
	poolInfo.SeparatePool = true
	suite.keeper.SetPool(suite.ctx, poolInfo)

	suite.investors = keeper.Wallets
	suite.creator = "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0"
}

// The default state used by each test
func (suite *mockBurnSuite) SetupTest() {
	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)
	lapp, k, nftKeeper, _, auctionKeeper, wctx := setupMsgServer(suite.T())
	ctx := sdk.UnwrapSDKContext(wctx)
	// create the first pool apy 7.8%

	suite.ctx = ctx
	suite.keeper = k
	suite.app = lapp
	suite.nftKeeper = nftKeeper
	suite.auctionKeeper = auctionKeeper
}

func TestMockBurn(t *testing.T) {
	suite.Run(t, new(mockBurnSuite))
}

func (suite *mockBurnSuite) TestBurn() {
	lsetupMockPool(suite)

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
	// Generate  random integer for junior pool

	juniorAmounts, _ := generateRandomIntegersWithSum(6, 300000)
	seniorAmounts, _ := generateRandomIntegersWithSum(6, 1000000)

	// we have 8 users, and the last two will be used as the transfer one

	// senior 0x8083eaa3584b60c163fe63e5ab6937526022cd47c35f9cb1e0790005c3ae9d00
	// junior 0xe0d49c3eed41e408b493a14042a8aa31375d64e3e357f911afbb085e02bde083

	seniorPool := "0x8083eaa3584b60c163fe63e5ab6937526022cd47c35f9cb1e0790005c3ae9d00"
	juniorPool := "0xe0d49c3eed41e408b493a14042a8aa31375d64e3e357f911afbb085e02bde083"
	msgDepositUsersJunior := make([]*types.MsgDeposit, 6)
	msgDepositUsersSenior := make([]*types.MsgDeposit, 6)
	for i := 0; i < 6; i++ {
		token := sdk.NewCoin("ausdc", sdkmath.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(int64(juniorAmounts[i])), base)))
		msg := &types.MsgDeposit{
			Creator:   suite.investors[i],
			PoolIndex: juniorPool,
			Token:     token,
		}
		msgDepositUsersJunior[i] = msg

		token = sdk.NewCoin("ausdc", sdkmath.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(int64(seniorAmounts[i])), base)))
		msg = &types.MsgDeposit{
			Creator:   suite.investors[i],
			PoolIndex: seniorPool,
			Token:     token,
		}
		msgDepositUsersSenior[i] = msg
	}

	poolInfo, ok := suite.keeper.GetPools(suite.ctx, seniorPool)
	suite.Require().True(ok)
	suite.Require().Equal(poolInfo.Apy.String(), "0.087500000000000000")
	poolInfo.PoolLockedSeconds = 360000
	poolInfo.PoolTotalBorrowLimit = 10000
	poolInfo.SeparatePool = true
	suite.keeper.SetPool(suite.ctx, poolInfo)
	poolInfo, ok = suite.keeper.GetPools(suite.ctx, juniorPool)
	suite.Require().True(ok)
	suite.Require().Equal(poolInfo.Apy.String(), "0.150000000000000000")
	poolInfo.PoolLockedSeconds = 360000
	poolInfo.PoolTotalBorrowLimit = 100000
	poolInfo.SeparatePool = true
	suite.keeper.SetPool(suite.ctx, poolInfo)

	// now we deposit
	for i := 0; i < 6; i++ {
		_, err := suite.app.Deposit(suite.ctx, msgDepositUsersSenior[i])
		suite.Require().NoError(err)
		_, err = suite.app.Deposit(suite.ctx, msgDepositUsersJunior[i])
		suite.Require().NoError(err)
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * 10))
	}

	// now we check the total amount of the pool is 300,000 and 1000,0000, and the borrowed is 0
	poolInfo, ok = suite.keeper.GetPools(suite.ctx, seniorPool)
	suite.Require().True(ok)
	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(1000000), base))))
	suite.Require().True(poolInfo.BorrowedAmount.IsZero())
	poolInfo, ok = suite.keeper.GetPools(suite.ctx, juniorPool)
	suite.Require().True(ok)
	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(300000), base))))
	suite.Require().True(poolInfo.BorrowedAmount.IsZero())

	// now we borrow 200,000 and 800,000
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(oneWeek))
	msgSenior := types.MsgBorrow{Creator: suite.creator, PoolIndex: seniorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(800000), base)))}
	_, err := suite.app.Borrow(suite.ctx, &msgSenior)
	suite.Require().NoError(err)

	//  borrow another one
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Minute))
	msgJunior := types.MsgBorrow{Creator: suite.creator, PoolIndex: juniorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base)))}
	_, err = suite.app.Borrow(suite.ctx, &msgJunior)
	suite.Require().NoError(err)

	startTime := suite.ctx.BlockTime()

	// we pay the interest  and principle in the escrow account
	_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{
		Creator:   suite.creator,
		PoolIndex: juniorPool,
		Token:     sdk.Coin{Denom: "ausdc", Amount: sdkmath.NewInt(50000).Mul(sdk.NewIntFromBigInt(base))},
	})
	suite.Require().NoError(err)

	_, err = suite.app.PayPrincipalForWithdrawalRequests(suite.ctx, &types.MsgPayPrincipalPartial{
		Creator:   suite.creator,
		PoolIndex: juniorPool,
		Token: sdk.Coin{
			Denom:  "ausdc",
			Amount: sdkmath.NewInt(200000).Mul(sdk.NewIntFromBigInt(base)),
		},
	})
	suite.Require().ErrorContains(err, "no withdraw proposal to be paid: invalid request")

	poolInfo, ok = suite.keeper.GetPools(suite.ctx, seniorPool)
	suite.Require().True(ok)

	_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{
		Creator:   suite.creator,
		PoolIndex: seniorPool,
		Token:     sdk.Coin{Denom: "ausdc", Amount: sdkmath.NewInt(80000).Mul(sdk.NewIntFromBigInt(base))},
	})
	suite.Require().NoError(err)

	_, err = suite.app.PayPrincipalForWithdrawalRequests(suite.ctx, &types.MsgPayPrincipalPartial{
		Creator:   suite.creator,
		PoolIndex: seniorPool,
		Token: sdk.Coin{
			Denom:  "ausdc",
			Amount: sdkmath.NewInt(800000).Mul(sdk.NewIntFromBigInt(base)),
		},
	})
	suite.Require().ErrorContains(err, "no withdraw proposal to be paid: invalid request")

	poolInfo, ok = suite.keeper.GetPools(suite.ctx, seniorPool)
	suite.Require().True(ok)
	suite.Require().True(checkValueEqualWithExchange(poolInfo.BorrowedAmount.Amount, sdkmath.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(800000), base))))
	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base))))
	poolInfo, ok = suite.keeper.GetPools(suite.ctx, juniorPool)
	suite.Require().True(ok)
	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(100000), base))))
	suite.Require().True(checkValueEqualWithExchange(poolInfo.BorrowedAmount.Amount, sdkmath.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base))))

	// we simulate 6 seconds for each block
	deltaTime := time.Second * time.Duration(60)
	checkPoints := []time.Duration{
		time.Duration(oneMonth) * time.Second,
		time.Duration(oneMonth) * time.Second * 2,
		time.Duration(oneYear) * time.Second,
	}

	checkPointCounter := 0

	for {
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(deltaTime))
		spv.EndBlock(suite.ctx, *suite.keeper)
		if startTime.Add(checkPoints[checkPointCounter]).Before(suite.ctx.BlockTime()) {
			checkPointCounter++
			suite.Require().True(suite.auctionKeeper.AuctionAmount[0].Amount.String() == "288461538461538359998")
			if checkPointCounter == len(checkPoints) {
				break
			}
		}
	}
}
