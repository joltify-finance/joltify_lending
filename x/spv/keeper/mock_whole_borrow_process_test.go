package keeper_test

import (
	"fmt"
	"math/big"
	"math/rand"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/davecgh/go-spew/spew"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/utils"
	"github.com/joltify-finance/joltify_lending/x/spv"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/suite"
)

const (
	oneWeek  = 7 * 24 * 3600
	oneMonth = oneWeek * 4
	oneYear  = oneWeek * 52
)

// Test suite used for all keeper tests
type mockWholeProcessSuite struct {
	suite.Suite
	keeper       *spvkeeper.Keeper
	nftKeeper    types.NFTKeeper
	app          types.MsgServer
	ctx          sdk.Context
	investors    []string
	investorPool string
	creator      string
}

var base = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

func setupMockPool(suite *mockWholeProcessSuite) {
	// create the first pool apy 7.8%
	// senior pool is 300,000
	amount := new(big.Int).Mul(big.NewInt(200000), base)
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 4, PoolName: "hello", Apy: "0.15", TargetTokenAmount: sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(amount))}
	resp, err := suite.app.CreatePool(suite.ctx, &req)
	suite.Require().NoError(err)

	depositorPool := resp.PoolIndex[0]

	suite.investorPool = depositorPool

	_, err = suite.app.ActivePool(suite.ctx, types.NewMsgActivePool("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", resp.PoolIndex[0]))
	suite.Require().NoError(err)

	_, err = suite.app.ActivePool(suite.ctx, types.NewMsgActivePool("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", resp.PoolIndex[1]))
	suite.Require().NoError(err)

	req2 := types.MsgAddInvestors{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: resp.PoolIndex[0],
		InvestorID: []string{"2"}}
	_, err = suite.app.AddInvestors(suite.ctx, &req2)
	suite.Require().NoError(err)

	//creator1 := "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl"
	//creator2 := "jolt1kkujrm0lqeu0e5va5f6mmwk87wva0k8cmam8jq"
	//creator3 := "jolt1z0y0zl0trsnuqmqf5v034pyv9sp39jg3rv6lsm"
	//creator4 := "jolt1fcaa73cc9c2l3l2u57skddgd0zm749ncukx90g"

	suite.investors = keeper.Wallets
	suite.creator = "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0"
}

// The default state used by each test
func (suite *mockWholeProcessSuite) SetupTest() {

	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)
	app, k, nftKeeper, wctx := setupMsgServer(suite.T())
	ctx := sdk.UnwrapSDKContext(wctx)
	// create the first pool apy 7.8%

	suite.ctx = ctx
	suite.keeper = k
	suite.app = app
	suite.nftKeeper = nftKeeper
}
func TestMockWholeProcess(t *testing.T) {
	suite.Run(t, new(mockWholeProcessSuite))
}

// generateRandomIntegersWithSum generates n random integers with a given sum
func generateRandomIntegersWithSum(n int, targetSum int) []int {
	result := make([]int, n)
	remainingSum := targetSum

	for i := 0; i < n-1; i++ {
		// Generate a random integer between 0 and remainingSum
		randomInt := rand.Intn(remainingSum)
		result[i] = randomInt
		remainingSum -= randomInt
	}

	// Set the last integer to the remaining sum
	result[n-1] = remainingSum

	return result
}

func (suite *mockWholeProcessSuite) TestMockSystemOneYearSimple() {
	setupMockPool(suite)

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
	// Generate  random integer for junior pool

	juniorAmounts := generateRandomIntegersWithSum(6, 300000)
	seniorAmounts := generateRandomIntegersWithSum(6, 1000000)

	_ = seniorAmounts
	// we have 8 users, and the last two will be used as the transfer one

	// senior 0x8083eaa3584b60c163fe63e5ab6937526022cd47c35f9cb1e0790005c3ae9d00
	// junior 0xe0d49c3eed41e408b493a14042a8aa31375d64e3e357f911afbb085e02bde083

	seniorPool := "0x8083eaa3584b60c163fe63e5ab6937526022cd47c35f9cb1e0790005c3ae9d00"
	juniorPool := "0xe0d49c3eed41e408b493a14042a8aa31375d64e3e357f911afbb085e02bde083"
	msgDepositUsersJunior := make([]*types.MsgDeposit, 6)
	msgDepositUsersSenior := make([]*types.MsgDeposit, 6)
	for i := 0; i < 6; i++ {
		token := sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(int64(juniorAmounts[i])), base)))
		msg := &types.MsgDeposit{Creator: suite.investors[i],
			PoolIndex: juniorPool,
			Token:     token}
		msgDepositUsersJunior[i] = msg

		token = sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(int64(seniorAmounts[i])), base)))
		msg = &types.MsgDeposit{Creator: suite.investors[i],
			PoolIndex: seniorPool,
			Token:     token}
		msgDepositUsersSenior[i] = msg
	}

	poolInfo, ok := suite.keeper.GetPools(suite.ctx, seniorPool)
	suite.Require().True(ok)
	suite.Require().Equal(poolInfo.Apy.String(), "0.087500000000000000")
	spew.Dump(poolInfo)
	poolInfo, ok = suite.keeper.GetPools(suite.ctx, juniorPool)
	suite.Require().True(ok)
	suite.Require().Equal(poolInfo.Apy.String(), "0.150000000000000000")

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
	spew.Dump(poolInfo)
	poolInfo, ok = suite.keeper.GetPools(suite.ctx, juniorPool)
	suite.Require().True(ok)
	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(300000), base))))
	suite.Require().True(poolInfo.BorrowedAmount.IsZero())

	// now we borrow 200,000 and 800,000
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(oneWeek))
	msgSenior := types.MsgBorrow{Creator: suite.creator, PoolIndex: seniorPool, BorrowAmount: sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(800000), base)))}
	suite.app.Borrow(suite.ctx, &msgSenior)

	// we simulate 6 seconds for each block
	startTime := suite.ctx.BlockTime()
	firstTime := true
	deltaTime := time.Second * time.Duration(60)
	checkPoints := []time.Duration{
		time.Duration(oneMonth) * time.Second,
		//time.Duration(oneMonth*2) * time.Second,
		//time.Duration(oneMonth*3) * time.Second,
		//time.Duration(oneMonth*4) * time.Second,
		//time.Duration(oneMonth*5) * time.Second,
		time.Duration(oneMonth*6) * time.Second,
	}

	checkPointCounter := 0

	allCoinsSenior := make(map[int]sdk.Coins)
	allCoinsJunior := make(map[int]sdk.Coins)
	for {
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(deltaTime))

		spv.EndBlock(suite.ctx, *suite.keeper)

		if suite.ctx.BlockTime().After(startTime.Add(time.Minute)) && firstTime {
			firstTime = false
			// 1 minutes, borrow another one
			suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Minute))
			msgJunior := types.MsgBorrow{Creator: suite.creator, PoolIndex: juniorPool, BorrowAmount: sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base)))}
			suite.app.Borrow(suite.ctx, &msgJunior)

			poolInfo, ok = suite.keeper.GetPools(suite.ctx, seniorPool)
			suite.Require().True(ok)
			spew.Dump(poolInfo)
			suite.Require().True(poolInfo.BorrowedAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(800000), base))))
			suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base))))
			poolInfo, ok = suite.keeper.GetPools(suite.ctx, juniorPool)
			suite.Require().True(ok)
			suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(100000), base))))
			suite.Require().True(poolInfo.BorrowedAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base))))
		}

		if startTime.Add(checkPoints[checkPointCounter]).Before(suite.ctx.BlockTime()) {
			coins := suite.getAllInvestorInterest(seniorPool, suite.investors[:6])
			allCoinsSenior[checkPointCounter] = coins
			coins2 := suite.getAllInvestorInterest(juniorPool, suite.investors[:6])
			allCoinsJunior[checkPointCounter] = coins2
			checkPointCounter++
			if checkPointCounter == len(checkPoints) {
				break
			}
		}
	}

	sumJunior := sdk.NewCoin("ausdc", sdk.ZeroInt())
	sumSenior := sdk.NewCoin("ausdc", sdk.ZeroInt())

	// we check the first month
	firstMonthJuniorCoins := allCoinsJunior[0]
	for _, el := range firstMonthJuniorCoins {
		sumJunior = sumJunior.Add(el)
	}

	firstMonthSeniorCoins := allCoinsSenior[0]
	for _, el := range firstMonthSeniorCoins {
		sumSenior = sumSenior.Add(el)
	}

	// now we borrow 200,000 and 800,000

	fmt.Printf(">>>>>>%v\n", sumSenior.Amount.MulRaw(12))
	fmt.Printf(">>>>>>%v\n", sumJunior.Amount.MulRaw(12))

}

func (suite *mockWholeProcessSuite) getAllInvestorInterest(poolType string, investors []string) sdk.Coins {

	interests := make(sdk.Coins, len(investors))
	for i, el := range investors {
		resp, err := suite.keeper.ClaimableInterest(suite.ctx, &types.QueryClaimableInterestRequest{el, poolType})
		suite.Require().NoError(err)
		interests[i] = resp.ClaimableInterestAmount
	}
	return interests

}
