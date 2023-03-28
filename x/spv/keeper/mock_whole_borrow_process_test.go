package keeper_test

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
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
func generateRandomIntegersWithSum(n int, targetSum int) ([]int, []sdk.Dec) {
	result := make([]int, n)
	ratio := make([]sdk.Dec, n)
	remainingSum := targetSum

	rand.Seed(100)
	for i := 0; i < n-1; i++ {
		// Generate a random integer between 0 and remainingSum
		randomInt := rand.Intn(remainingSum)
		result[i] = randomInt
		ratio[i] = sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(randomInt))).QuoTruncate(sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(targetSum))))
		remainingSum -= randomInt
	}

	// Set the last integer to the remaining sum
	result[n-1] = remainingSum
	ratio[n-1] = sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(result[n-1]))).QuoTruncate(sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(targetSum))))
	return result, ratio
}

func (suite *mockWholeProcessSuite) TestMockSystemOneYearSimple() {
	setupMockPool(suite)

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
	// Generate  random integer for junior pool

	juniorAmounts, ratioJunior := generateRandomIntegersWithSum(6, 300000)
	seniorAmounts, ratioSenior := generateRandomIntegersWithSum(6, 1000000)

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

	//  borrow another one
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Minute))
	msgJunior := types.MsgBorrow{Creator: suite.creator, PoolIndex: juniorPool, BorrowAmount: sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base)))}
	suite.app.Borrow(suite.ctx, &msgJunior)

	startTime := suite.ctx.BlockTime()

	// we pay the interest  and principle in the escrow account
	_, err := suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{
		Creator:   suite.creator,
		PoolIndex: juniorPool,
		Token:     sdk.Coin{Denom: "ausdc", Amount: sdk.NewInt(50000).Mul(sdk.NewIntFromBigInt(base))},
	})
	suite.Require().NoError(err)

	_, err = suite.app.PayPrincipalForWithdrawalRequests(suite.ctx, &types.MsgPayPrincipal{
		Creator:   suite.creator,
		PoolIndex: juniorPool,
		Token: sdk.Coin{
			Denom:  "ausdc",
			Amount: sdk.NewInt(200000).Mul(sdk.NewIntFromBigInt(base)),
		},
	})
	suite.Require().NoError(err)

	_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{
		Creator:   suite.creator,
		PoolIndex: seniorPool,
		Token:     sdk.Coin{Denom: "ausdc", Amount: sdk.NewInt(80000).Mul(sdk.NewIntFromBigInt(base))},
	})
	suite.Require().NoError(err)

	_, err = suite.app.PayPrincipalForWithdrawalRequests(suite.ctx, &types.MsgPayPrincipal{
		Creator:   suite.creator,
		PoolIndex: seniorPool,
		Token: sdk.Coin{
			Denom:  "ausdc",
			Amount: sdk.NewInt(800000).Mul(sdk.NewIntFromBigInt(base)),
		},
	})
	suite.Require().NoError(err)

	poolInfo, ok = suite.keeper.GetPools(suite.ctx, seniorPool)
	suite.Require().True(ok)
	spew.Dump(poolInfo)
	suite.Require().True(poolInfo.BorrowedAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(800000), base))))
	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base))))
	poolInfo, ok = suite.keeper.GetPools(suite.ctx, juniorPool)
	suite.Require().True(ok)
	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(100000), base))))
	suite.Require().True(poolInfo.BorrowedAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base))))

	// we simulate 6 seconds for each block
	deltaTime := time.Second * time.Duration(60)
	checkPoints := []time.Duration{
		time.Duration(oneMonth) * time.Second,
		time.Duration(oneYear) * time.Second,
	}

	checkPointCounter := 0

	allCoinsSenior := make(map[int]sdk.Coins)
	allCoinsJunior := make(map[int]sdk.Coins)
	for {
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(deltaTime))
		spv.EndBlock(suite.ctx, *suite.keeper)

		//if suite.ctx.BlockTime().After(startTime.Add(time.Minute)) && firstTime {
		//	firstTime = false
		//	// 1 minutes, borrow another one
		//	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Minute))
		//	msgJunior := types.MsgBorrow{Creator: suite.creator, PoolIndex: juniorPool, BorrowAmount: sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base)))}
		//	suite.app.Borrow(suite.ctx, &msgJunior)
		//
		//	poolInfo, ok = suite.keeper.GetPools(suite.ctx, seniorPool)
		//	suite.Require().True(ok)
		//	spew.Dump(poolInfo)
		//	suite.Require().True(poolInfo.BorrowedAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(800000), base))))
		//	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base))))
		//	poolInfo, ok = suite.keeper.GetPools(suite.ctx, juniorPool)
		//	suite.Require().True(ok)
		//	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(100000), base))))
		//	suite.Require().True(poolInfo.BorrowedAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base))))
		//}
		// now we submit the withdrawal request

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

	sumJuniorAllYear := sdk.NewCoin("ausdc", sdk.ZeroInt())
	sumSeniorAllYear := sdk.NewCoin("ausdc", sdk.ZeroInt())

	// we check the first month
	// total junior is interest is
	firstMonthJuniorCoins := allCoinsJunior[0]
	for _, el := range firstMonthJuniorCoins {
		sumJunior = sumJunior.Add(el)
	}

	firstMonthSeniorCoins := allCoinsSenior[0]
	for _, el := range firstMonthSeniorCoins {
		sumSenior = sumSenior.Add(el)
	}

	twelveMonthJuniorCoins := allCoinsJunior[1]
	for _, el := range twelveMonthJuniorCoins {
		sumJuniorAllYear = sumJuniorAllYear.Add(el)
	}

	twelveMonthSeniorCoins := allCoinsSenior[1]
	for _, el := range twelveMonthSeniorCoins {
		sumSeniorAllYear = sumSeniorAllYear.Add(el)
	}

	fmt.Printf(">sssssssss>>>>>%v\n", sumJunior)
	expected := sumJunior.Amount.MulRaw(13)
	expectedS := sumSenior.Amount.MulRaw(13)
	suite.Require().True(expected.Equal(sumJuniorAllYear.Amount))
	suite.Require().True(expectedS.Equal(sumSeniorAllYear.Amount))

	// now we borrow 200,000 and 800,000
	totalToInvestor := sumSenior.Add(sumJunior)
	reserve, ok := suite.keeper.GetReserve(suite.ctx, "ausdc")
	suite.Require().True(ok)
	totalInterest := totalToInvestor.Add(reserve)

	for i, el := range firstMonthJuniorCoins {
		ratioGet := sdk.NewDecFromInt(el.Amount).QuoTruncate(sdk.NewDecFromInt(sumJunior.Amount))
		suite.Require().True(ratioGet.Sub(ratioJunior[i]).Abs().LT(sdk.NewDecFromBigInt(big.NewInt(100))))
	}

	for i, el := range firstMonthSeniorCoins {
		ratioGet := sdk.NewDecFromInt(el.Amount).QuoTruncate(sdk.NewDecFromInt(sumSenior.Amount))
		suite.Require().True(ratioGet.Sub(ratioSenior[i]).Abs().LT(sdk.NewDecFromBigInt(big.NewInt(100))))
	}

	//so the total interest should be 100000, we calcualte the 4 weeks' interest, and one year we have 52/4=13 payments
	suite.Require().True(sdk.NewIntFromUint64(100000).Sub(totalInterest.Amount).LTE(sdk.NewIntFromUint64(1000000)))

}

func (suite *mockWholeProcessSuite) TestMockSystemOneYearWithWithdrawal() {
	setupMockPool(suite)

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
	// Generate  random integer for junior pool

	juniorAmounts, ratioJunior := generateRandomIntegersWithSum(6, 300000)
	seniorAmounts, ratioSenior := generateRandomIntegersWithSum(6, 1000000)

	_ = ratioSenior
	_ = ratioJunior

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

	//  borrow another one
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Minute))
	msgJunior := types.MsgBorrow{Creator: suite.creator, PoolIndex: juniorPool, BorrowAmount: sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base)))}
	suite.app.Borrow(suite.ctx, &msgJunior)

	startTime := suite.ctx.BlockTime()

	// we pay the interest  and principle in the escrow account
	_, err := suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{
		Creator:   suite.creator,
		PoolIndex: juniorPool,
		Token:     sdk.Coin{Denom: "ausdc", Amount: sdk.NewInt(30000).Mul(sdk.NewIntFromBigInt(base))},
	})
	suite.Require().NoError(err)

	_, err = suite.app.PayPrincipalForWithdrawalRequests(suite.ctx, &types.MsgPayPrincipal{
		Creator:   suite.creator,
		PoolIndex: juniorPool,
		Token: sdk.Coin{
			Denom:  "ausdc",
			Amount: sdk.NewInt(200000).Mul(sdk.NewIntFromBigInt(base)),
		},
	})
	suite.Require().NoError(err)

	_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{
		Creator:   suite.creator,
		PoolIndex: seniorPool,
		Token:     sdk.Coin{Denom: "ausdc", Amount: sdk.NewInt(70000).Mul(sdk.NewIntFromBigInt(base))},
	})
	suite.Require().NoError(err)

	_, err = suite.app.PayPrincipalForWithdrawalRequests(suite.ctx, &types.MsgPayPrincipal{
		Creator:   suite.creator,
		PoolIndex: seniorPool,
		Token: sdk.Coin{
			Denom:  "ausdc",
			Amount: sdk.NewInt(800000).Mul(sdk.NewIntFromBigInt(base)),
		},
	})
	suite.Require().NoError(err)

	poolInfo, _ = suite.keeper.GetPools(suite.ctx, seniorPool)
	principalEscrow := poolInfo.EscrowPrincipalAmount.Amount

	addr, err := sdk.AccAddressFromBech32(suite.investors[0])
	suite.Require().NoError(err)
	d0, _ := suite.keeper.GetDepositor(suite.ctx, seniorPool, addr)
	//suite.Require().True(newinvestor1.WithdrawalAmount.Amount.Equal(sdk.NewIntFromUint64(uint64(seniorAmounts[0])).Mul(sdk.NewIntFromBigInt(base))))
	addr, err = sdk.AccAddressFromBech32(suite.investors[1])
	suite.Require().NoError(err)
	d1, _ := suite.keeper.GetDepositor(suite.ctx, seniorPool, addr)

	totalLocked := d0.LockedAmount.AddAmount(d1.LockedAmount.Amount)

	poolInfo, ok = suite.keeper.GetPools(suite.ctx, seniorPool)
	suite.Require().True(ok)
	spew.Dump(poolInfo)
	suite.Require().True(poolInfo.BorrowedAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(800000), base))))
	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base))))
	poolInfo, ok = suite.keeper.GetPools(suite.ctx, juniorPool)
	suite.Require().True(ok)
	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(100000), base))))
	suite.Require().True(poolInfo.BorrowedAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base))))

	// we simulate 6 seconds for each block
	deltaTime := time.Second * time.Duration(60)
	checkPoints := []time.Duration{
		time.Duration(oneMonth) * time.Second,
		time.Duration(oneYear) * time.Second,
	}
	withdrawRequest := []time.Duration{
		time.Duration(oneMonth*11) * time.Second,
	}

	checkPointCounter := 0

	allCoinsSenior := make(map[int]sdk.Coins)
	allCoinsJunior := make(map[int]sdk.Coins)
	processed := false
	for {
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(deltaTime))
		spv.EndBlock(suite.ctx, *suite.keeper)

		//if suite.ctx.BlockTime().After(startTime.Add(time.Minute)) && firstTime {
		//	firstTime = false
		//	// 1 minutes, borrow another one
		//	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Minute))
		//	msgJunior := types.MsgBorrow{Creator: suite.creator, PoolIndex: juniorPool, BorrowAmount: sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base)))}
		//	suite.app.Borrow(suite.ctx, &msgJunior)
		//
		//	poolInfo, ok = suite.keeper.GetPools(suite.ctx, seniorPool)
		//	suite.Require().True(ok)
		//	spew.Dump(poolInfo)
		//	suite.Require().True(poolInfo.BorrowedAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(800000), base))))
		//	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base))))
		//	poolInfo, ok = suite.keeper.GetPools(suite.ctx, juniorPool)
		//	suite.Require().True(ok)
		//	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(100000), base))))
		//	suite.Require().True(poolInfo.BorrowedAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base))))
		//}
		// now we submit the withdrawal request
		if startTime.Add(withdrawRequest[0]).Before(suite.ctx.BlockTime()) && !processed {
			for i := 0; i < 2; i++ {
				_, err := suite.app.SubmitWithdrawProposal(suite.ctx, &types.MsgSubmitWithdrawProposal{
					Creator:   suite.investors[i],
					PoolIndex: seniorPool,
				})
				suite.Require().NoError(err)
			}
			processed = true
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

	//
	//spv.EndBlock(suite.ctx, *suite.keeper)

	addr, err = sdk.AccAddressFromBech32(suite.investors[0])
	suite.Require().NoError(err)
	newinvestor1, _ := suite.keeper.GetDepositor(suite.ctx, seniorPool, addr)
	//suite.Require().True(newinvestor1.WithdrawalAmount.Amount.Equal(sdk.NewIntFromUint64(uint64(seniorAmounts[0])).Mul(sdk.NewIntFromBigInt(base))))
	_ = newinvestor1
	addr, err = sdk.AccAddressFromBech32(suite.investors[1])
	suite.Require().NoError(err)
	newinvestor2, _ := suite.keeper.GetDepositor(suite.ctx, seniorPool, addr)
	_ = newinvestor2
	//suite.Require().True(newinvestor2.WithdrawalAmount.Amount.Equal(sdk.NewIntFromUint64(uint64(seniorAmounts[1])).Mul(sdk.NewIntFromBigInt(base))))

	// total 59499999999999975519844
	coins := suite.getAllInvestorInterest(seniorPool, suite.investors[:6])
	allCoinsSenior[checkPointCounter] = coins

	sumSenior := sdk.NewCoin("ausdc", sdk.ZeroInt())
	sumJunior := sdk.NewCoin("ausdc", sdk.ZeroInt())
	allMonthSeniorCoins := allCoinsSenior[1]
	for _, el := range allMonthSeniorCoins {
		sumSenior = sumSenior.Add(el)
	}

	allMonthJuniorCoins := allCoinsJunior[1]
	for _, el := range allMonthJuniorCoins {
		sumJunior = sumJunior.Add(el)
	}

	totalToInvestor := sumSenior.Add(sumJunior)
	reserve, ok := suite.keeper.GetReserve(suite.ctx, "ausdc")
	suite.Require().True(ok)
	totalInterest := totalToInvestor.Add(reserve)
	suite.Require().True(sdk.NewIntFromUint64(100000).Sub(totalInterest.Amount).LTE(sdk.NewIntFromUint64(1000000)))

	seniorPoolInfo, _ := suite.keeper.GetPools(suite.ctx, seniorPool)
	juniorPoolInfo, _ := suite.keeper.GetPools(suite.ctx, juniorPool)
	totalLeftInterest := seniorPoolInfo.EscrowInterestAmount.Add(juniorPoolInfo.EscrowInterestAmount)
	suite.Require().True(totalLeftInterest.LT(sdk.NewIntFromUint64(1000000000)))
	suite.Require().True(totalLocked.Add(seniorPoolInfo.EscrowPrincipalAmount).Amount.Equal(principalEscrow))

	// now we withdraw
	totalWithdrawal := sdk.ZeroInt()
	resp, err := suite.app.WithdrawPrincipal(suite.ctx, &types.MsgWithdrawPrincipal{Creator: suite.investors[0], PoolIndex: seniorPool, Token: sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1))})
	suite.Require().NoError(err)
	token, _ := sdk.ParseCoinNormalized(resp.Amount)
	suite.Require().True(sdk.NewIntFromUint64(uint64(seniorAmounts[0])).Mul(sdk.NewIntFromBigInt(base)).Equal(token.Amount))
	totalWithdrawal = totalWithdrawal.Add(token.Amount)
	resp, err = suite.app.WithdrawPrincipal(suite.ctx, &types.MsgWithdrawPrincipal{Creator: suite.investors[1], PoolIndex: seniorPool, Token: sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1))})
	suite.Require().NoError(err)
	token, _ = sdk.ParseCoinNormalized(resp.Amount)
	suite.Require().True(sdk.NewIntFromUint64(uint64(seniorAmounts[1])).Mul(sdk.NewIntFromBigInt(base)).Equal(token.Amount))
	totalWithdrawal = totalWithdrawal.Add(token.Amount)

	addr, err = sdk.AccAddressFromBech32(suite.investors[0])
	suite.Require().NoError(err)
	_, ok = suite.keeper.GetDepositor(suite.ctx, seniorPool, addr)
	suite.Require().False(ok)
	addr, err = sdk.AccAddressFromBech32(suite.investors[1])
	suite.Require().NoError(err)
	_, ok = suite.keeper.GetDepositor(suite.ctx, seniorPool, addr)
	suite.Require().False(ok)

	for _, el := range suite.investors[2:6] {
		_, err := suite.app.ClaimInterest(suite.ctx, &types.MsgClaimInterest{
			el,
			seniorPool,
		})
		suite.Require().NoError(err)
		_, err = suite.app.ClaimInterest(suite.ctx, &types.MsgClaimInterest{
			el,
			juniorPool,
		})
		suite.Require().NoError(err)
	}

	for _, el := range suite.investors[:2] {
		_, err := suite.app.ClaimInterest(suite.ctx, &types.MsgClaimInterest{
			el,
			juniorPool,
		})
		suite.Require().NoError(err)
	}

	_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{
		Creator:   suite.creator,
		PoolIndex: juniorPool,
		Token:     sdk.Coin{Denom: "ausdc", Amount: sdk.NewInt(30000).Mul(sdk.NewIntFromBigInt(base))},
	})
	suite.Require().NoError(err)

	suite.Require().NoError(err)

	_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{
		Creator:   suite.creator,
		PoolIndex: seniorPool,
		Token:     sdk.Coin{Denom: "ausdc", Amount: sdk.NewInt(70000).Mul(sdk.NewIntFromBigInt(base))},
	})
	suite.Require().NoError(err)

	// now we can test another 2 months, spv pay the interest, then the spv close the pool

	p1, _ := suite.keeper.GetPools(suite.ctx, seniorPool)
	p2, _ := suite.keeper.GetPools(suite.ctx, juniorPool)

	seniorLocked := p1.BorrowedAmount
	juniorLocked := p2.BorrowedAmount
	fmt.Printf(">>>>>>>>>>>%v=====%v\n", seniorLocked, juniorLocked)
	// senior apy is
	seniorInterest := sdk.NewDecFromInt(seniorLocked.Amount).Mul(sdk.NewDecWithPrec(875, 4))
	juniorInterest := sdk.NewDecFromInt(juniorLocked.Amount).Mul(sdk.NewDecWithPrec(15, 2))

	_ = seniorInterest
	_ = juniorInterest

	currentTime := suite.ctx.BlockTime()

	totalSeniorInterest := sdk.ZeroInt()
	totalJuniorInterest := sdk.ZeroInt()
	eightWeeks := oneWeek * 8
	for {
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(deltaTime))
		spv.EndBlock(suite.ctx, *suite.keeper)

		if currentTime.Add(time.Second * time.Duration(eightWeeks)).Before(suite.ctx.BlockTime()) {
			_, err := suite.app.PayPrincipal(suite.ctx, &types.MsgPayPrincipal{
				Creator:   suite.creator,
				PoolIndex: juniorPool,
				Token:     sdk.Coin{Denom: "ausdc", Amount: sdk.NewIntFromUint64(1)},
			})
			suite.Require().NoError(err)

			_, err = suite.app.PayPrincipal(suite.ctx, &types.MsgPayPrincipal{
				Creator:   suite.creator,
				PoolIndex: seniorPool,
				Token:     sdk.Coin{Denom: "ausdc", Amount: sdk.NewIntFromUint64(1)},
			})
			suite.Require().NoError(err)
			break
		}
	}

	p2, _ = suite.keeper.GetPools(suite.ctx, juniorPool)
	class, found := suite.nftKeeper.GetClass(suite.ctx, p2.PoolNFTIds[0])
	suite.Require().True(found)
	var borrowInterest types.BorrowInterest
	err = proto.Unmarshal(class.Data.Value, &borrowInterest)
	if err != nil {
		panic(err)
	}

	fmt.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>\n")
	item := borrowInterest.Payments[len(borrowInterest.Payments)-1]
	spew.Dump(item)
	item = borrowInterest.Payments[1]
	spew.Dump(item)

	fmt.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>\n")

	var tokens sdk.Coins
	for _, el := range suite.investors[2:6] {
		resp, err := suite.app.ClaimInterest(suite.ctx, &types.MsgClaimInterest{
			Creator:   el,
			PoolIndex: seniorPool,
		})
		tokens, err = sdk.ParseCoinsNormalized(resp.Amount)
		suite.Require().NoError(err)
		totalSeniorInterest = totalSeniorInterest.Add(tokens[0].Amount)

		resp, err = suite.app.ClaimInterest(suite.ctx, &types.MsgClaimInterest{
			Creator:   el,
			PoolIndex: juniorPool,
		})
		tokens, err = sdk.ParseCoinsNormalized(resp.Amount)
		suite.Require().NoError(err)
		totalJuniorInterest = totalJuniorInterest.Add(tokens[0].Amount)
	}

	for _, el := range suite.investors[:2] {
		resp, err := suite.app.ClaimInterest(suite.ctx, &types.MsgClaimInterest{
			Creator:   el,
			PoolIndex: juniorPool,
		})
		tokens, err = sdk.ParseCoinsNormalized(resp.Amount)
		suite.Require().NoError(err)
		totalJuniorInterest = totalJuniorInterest.Add(tokens[0].Amount)
	}

	expectedYearSenior := sdk.NewDecFromInt(totalSeniorInterest).Mul(sdk.NewDec(1)).QuoTruncate(sdk.NewDec(8))
	expectedYearJunior := sdk.NewDecFromInt(totalJuniorInterest).Mul(sdk.NewDec(1)).QuoTruncate(sdk.NewDec(8))

	fmt.Printf(">>>>>%v===aaa==%v\n", expectedYearSenior, expectedYearJunior)

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(oneWeek)))
	spv.EndBlock(suite.ctx, *suite.keeper)
	p, _ := suite.keeper.GetPools(suite.ctx, seniorPool)
	suite.Require().Equal(p.PoolStatus, types.PoolInfo_CLOSED)
	p, _ = suite.keeper.GetPools(suite.ctx, juniorPool)
	suite.Require().Equal(p.PoolStatus, types.PoolInfo_CLOSED)

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second))

	for i, el := range suite.investors[2:6] {
		resp, err := suite.app.WithdrawPrincipal(suite.ctx, &types.MsgWithdrawPrincipal{
			Creator:   el,
			PoolIndex: seniorPool,
			Token:     sdk.NewCoin("ausdc", sdk.NewInt(1)),
		})
		tokens, err = sdk.ParseCoinsNormalized(resp.Amount)
		suite.Require().NoError(err)
		totalWithdrawal = totalWithdrawal.Add(tokens[0].Amount)
		fmt.Printf(">>>>senior>>total amount %v===%v\n", tokens[0].Amount, seniorAmounts[i+2])
		addr, _ := sdk.AccAddressFromBech32(el)
		_, found := suite.keeper.GetDepositor(suite.ctx, seniorPool, addr)
		suite.Require().False(found)
		resp, err = suite.app.WithdrawPrincipal(suite.ctx, &types.MsgWithdrawPrincipal{
			Creator:   el,
			PoolIndex: juniorPool,
			Token:     sdk.NewCoin("ausdc", sdk.NewInt(1)),
		})
		tokens, err = sdk.ParseCoinsNormalized(resp.Amount)
		suite.Require().NoError(err)
		fmt.Printf(">>>>junior>>total amount %v===%v\n", tokens[0].Amount, juniorAmounts[i+2])
		totalWithdrawal = totalWithdrawal.Add(tokens[0].Amount)
		_, found = suite.keeper.GetDepositor(suite.ctx, juniorPool, addr)
		suite.Require().False(found)
		suite.Require().NoError(err)
	}

	// we need to add the principal for the first two investors in the junior pool
	resp, err = suite.app.WithdrawPrincipal(suite.ctx, &types.MsgWithdrawPrincipal{
		Creator:   suite.investors[0],
		PoolIndex: juniorPool,
		Token:     sdk.NewCoin("ausdc", sdk.NewInt(1)),
	})
	tokens, err = sdk.ParseCoinsNormalized(resp.Amount)
	suite.Require().NoError(err)

	totalWithdrawal = totalWithdrawal.Add(tokens[0].Amount)

	resp, err = suite.app.WithdrawPrincipal(suite.ctx, &types.MsgWithdrawPrincipal{
		Creator:   suite.investors[1],
		PoolIndex: juniorPool,
		Token:     sdk.NewCoin("ausdc", sdk.NewInt(1)),
	})
	tokens, err = sdk.ParseCoinsNormalized(resp.Amount)
	suite.Require().NoError(err)
	totalWithdrawal = totalWithdrawal.Add(tokens[0].Amount)
	suite.Require().True(totalWithdrawal.Equal(sdk.NewIntFromUint64(1300000).Mul(sdk.NewIntFromBigInt(base))))

	_, found = suite.keeper.GetPools(suite.ctx, seniorPool)
	suite.Require().False(found)
	_, found = suite.keeper.GetPools(suite.ctx, juniorPool)
	suite.Require().False(found)

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
