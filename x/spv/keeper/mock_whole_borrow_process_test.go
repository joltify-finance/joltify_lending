package keeper_test

import (
	"math/big"
	"math/rand"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
	amountSenior := new(big.Int).Mul(big.NewInt(800000), base)
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 4, PoolName: "hello", Apy: []string{"0.15", "0.0875"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(amount)), sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(amountSenior))}}
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
func (suite *mockWholeProcessSuite) SetupTest() {
	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)
	lapp, k, nftKeeper, _, _, wctx := setupMsgServer(suite.T())
	ctx := sdk.UnwrapSDKContext(wctx)
	// create the first pool apy 7.8%

	suite.ctx = ctx
	suite.keeper = k
	suite.app = lapp
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
		msg := &types.MsgDeposit{
			Creator:   suite.investors[i],
			PoolIndex: juniorPool,
			Token:     token,
		}
		msgDepositUsersJunior[i] = msg

		token = sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(int64(seniorAmounts[i])), base)))
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
	msgSenior := types.MsgBorrow{Creator: suite.creator, PoolIndex: seniorPool, BorrowAmount: sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(800000), base)))}
	_, err := suite.app.Borrow(suite.ctx, &msgSenior)
	suite.Require().NoError(err)

	//  borrow another one
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Minute))
	msgJunior := types.MsgBorrow{Creator: suite.creator, PoolIndex: juniorPool, BorrowAmount: sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base)))}
	_, err = suite.app.Borrow(suite.ctx, &msgJunior)
	suite.Require().NoError(err)

	startTime := suite.ctx.BlockTime()

	// we pay the interest  and principle in the escrow account
	_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{
		Creator:   suite.creator,
		PoolIndex: juniorPool,
		Token:     sdk.Coin{Denom: "ausdc", Amount: sdk.NewInt(50000).Mul(sdk.NewIntFromBigInt(base))},
	})
	suite.Require().NoError(err)

	_, err = suite.app.PayPrincipalForWithdrawalRequests(suite.ctx, &types.MsgPayPrincipalPartial{
		Creator:   suite.creator,
		PoolIndex: juniorPool,
		Token: sdk.Coin{
			Denom:  "ausdc",
			Amount: sdk.NewInt(200000).Mul(sdk.NewIntFromBigInt(base)),
		},
	})
	suite.Require().ErrorContains(err, "no withdraw proposal to be paid: invalid request")

	poolInfo, ok = suite.keeper.GetPools(suite.ctx, seniorPool)
	suite.Require().True(ok)

	_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{
		Creator:   suite.creator,
		PoolIndex: seniorPool,
		Token:     sdk.Coin{Denom: "ausdc", Amount: sdk.NewInt(80000).Mul(sdk.NewIntFromBigInt(base))},
	})
	suite.Require().NoError(err)

	_, err = suite.app.PayPrincipalForWithdrawalRequests(suite.ctx, &types.MsgPayPrincipalPartial{
		Creator:   suite.creator,
		PoolIndex: seniorPool,
		Token: sdk.Coin{
			Denom:  "ausdc",
			Amount: sdk.NewInt(800000).Mul(sdk.NewIntFromBigInt(base)),
		},
	})
	suite.Require().ErrorContains(err, "no withdraw proposal to be paid: invalid request")

	poolInfo, ok = suite.keeper.GetPools(suite.ctx, seniorPool)
	suite.Require().True(ok)
	suite.Require().True(checkValueEqualWithExchange(poolInfo.BorrowedAmount.Amount, sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(800000), base))))
	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base))))
	poolInfo, ok = suite.keeper.GetPools(suite.ctx, juniorPool)
	suite.Require().True(ok)
	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(100000), base))))
	suite.Require().True(checkValueEqualWithExchange(poolInfo.BorrowedAmount.Amount, sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base))))

	// we simulate 6 seconds for each block
	deltaTime := time.Second * time.Duration(60)
	checkPoints := []time.Duration{
		time.Duration(oneMonth) * time.Second,
		time.Duration(oneMonth) * time.Second * 2,
		time.Duration(oneYear) * time.Second,
	}

	checkPointCounter := 0

	allCoinsSenior := make(map[int]sdk.Coins)
	allCoinsJunior := make(map[int]sdk.Coins)
	for {
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(deltaTime))
		spv.EndBlock(suite.ctx, *suite.keeper)

		if startTime.Add(checkPoints[checkPointCounter]).Before(suite.ctx.BlockTime()) {
			coins := suite.getAllInvestorInterest(seniorPool, suite.investors[:6], -1)
			allCoinsSenior[checkPointCounter] = coins
			coins2 := suite.getAllInvestorInterest(juniorPool, suite.investors[:6], -1)
			allCoinsJunior[checkPointCounter] = coins2
			checkPointCounter++
			if checkPointCounter == len(checkPoints) {
				break
			}
		}
	}

	sumJuniorF := sdk.NewCoin("ausdc", sdk.ZeroInt())
	sumSeniorF := sdk.NewCoin("ausdc", sdk.ZeroInt())
	sumJunior := sdk.NewCoin("ausdc", sdk.ZeroInt())
	sumSenior := sdk.NewCoin("ausdc", sdk.ZeroInt())

	sumJuniorAllYear := sdk.NewCoin("ausdc", sdk.ZeroInt())
	sumSeniorAllYear := sdk.NewCoin("ausdc", sdk.ZeroInt())

	// we check the first month
	// total junior is interest is

	firstMonthJuniorCoins := allCoinsJunior[0]
	for _, el := range firstMonthJuniorCoins {
		sumJuniorF = sumJuniorF.Add(el)
	}

	firstMonthSeniorCoins := allCoinsSenior[0]
	for _, el := range firstMonthSeniorCoins {
		sumSeniorF = sumSeniorF.Add(el)
	}

	secondMonthJuniorCoins := allCoinsJunior[1]
	for _, el := range secondMonthJuniorCoins {
		sumJunior = sumJunior.Add(el)
	}

	secondMonthSeniorCoins := allCoinsSenior[1]
	for _, el := range secondMonthSeniorCoins {
		sumSenior = sumSenior.Add(el)
	}

	// as the first month has not pay the whole pay freq, so we add it seperately
	fullPayFreqJunior := sumJunior.Sub(sumJuniorF)
	fullPayFreqSenior := sumSenior.Sub(sumSeniorF)

	twelveMonthJuniorCoins := allCoinsJunior[2]
	for _, el := range twelveMonthJuniorCoins {
		sumJuniorAllYear = sumJuniorAllYear.Add(el)
	}

	twelveMonthSeniorCoins := allCoinsSenior[2]
	for _, el := range twelveMonthSeniorCoins {
		sumSeniorAllYear = sumSeniorAllYear.Add(el)
	}

	expected := fullPayFreqJunior.Amount.MulRaw(12).Add(sumJuniorF.Amount)
	expectedS := fullPayFreqSenior.Amount.MulRaw(12).Add(sumSeniorF.Amount)
	suite.Require().True(expected.Equal(sumJuniorAllYear.Amount))
	suite.Require().True(expectedS.Equal(sumSeniorAllYear.Amount))

	// now we borrow 200,000 and 800,000
	totalToInvestor := sumSenior.Add(sumJunior)
	totalToInvestorUSD := convertBorrowToUsd(totalToInvestor.Amount)
	reserve, ok := suite.keeper.GetReserve(suite.ctx, "ausdc")
	suite.Require().True(ok)
	totalInterest := reserve.AddAmount(totalToInvestorUSD)

	for i, el := range secondMonthJuniorCoins {
		ratioGet := sdk.NewDecFromInt(el.Amount).QuoTruncate(sdk.NewDecFromInt(sumJunior.Amount))
		suite.Require().True(ratioGet.Sub(ratioJunior[i]).Abs().LT(sdk.NewDecWithPrec(1, 8)))
	}

	for i, el := range secondMonthSeniorCoins {
		ratioGet := sdk.NewDecFromInt(el.Amount).QuoTruncate(sdk.NewDecFromInt(sumSenior.Amount))
		suite.Require().True(ratioGet.Sub(ratioSenior[i]).Abs().LT(sdk.NewDecWithPrec(1, 8)))
	}

	// so the total interest should be 100000, we calcualte the 4 weeks' interest, and one year we have 52/4=13 payments
	suite.Require().True(sdk.NewIntFromUint64(100000).Sub(totalInterest.Amount).LTE(sdk.NewIntFromUint64(1000000)))
}

func (suite *mockWholeProcessSuite) TestMockSystemOneYearWithWithdrawal() {
	setupMockPool(suite)

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
		token := sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(int64(juniorAmounts[i])), base)))
		msg := &types.MsgDeposit{
			Creator:   suite.investors[i],
			PoolIndex: juniorPool,
			Token:     token,
		}
		msgDepositUsersJunior[i] = msg

		token = sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(int64(seniorAmounts[i])), base)))
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
	poolInfo.PoolTotalBorrowLimit = 100000
	poolInfo.WithdrawRequestWindowSeconds = oneMonth * 3
	suite.keeper.SetPool(suite.ctx, poolInfo)
	poolInfo, ok = suite.keeper.GetPools(suite.ctx, juniorPool)
	suite.Require().True(ok)
	suite.Require().Equal(poolInfo.Apy.String(), "0.150000000000000000")

	poolInfo.PoolLockedSeconds = 360000
	poolInfo.PoolTotalBorrowLimit = 100000
	poolInfo.WithdrawRequestWindowSeconds = oneMonth * 3
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
	msgSenior := types.MsgBorrow{Creator: suite.creator, PoolIndex: seniorPool, BorrowAmount: sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(800000), base)))}
	_, err := suite.app.Borrow(suite.ctx, &msgSenior)
	suite.Require().NoError(err)

	//  borrow another one
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Minute))
	msgJunior := types.MsgBorrow{Creator: suite.creator, PoolIndex: juniorPool, BorrowAmount: sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base)))}
	_, err = suite.app.Borrow(suite.ctx, &msgJunior)
	suite.Require().NoError(err)

	startTime := suite.ctx.BlockTime()

	// we pay the interest  and principle in the escrow account
	_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{
		Creator:   suite.creator,
		PoolIndex: juniorPool,
		Token:     sdk.Coin{Denom: "ausdc", Amount: sdk.NewInt(30000).Mul(sdk.NewIntFromBigInt(base))},
	})
	suite.Require().NoError(err)

	_, err = suite.app.PayPrincipalForWithdrawalRequests(suite.ctx, &types.MsgPayPrincipalPartial{
		Creator:   suite.creator,
		PoolIndex: juniorPool,
		Token: sdk.Coin{
			Denom:  "ausdc",
			Amount: sdk.NewInt(200000).Mul(sdk.NewIntFromBigInt(base)),
		},
	})
	suite.Require().ErrorContains(err, "no withdraw proposal to be paid: invalid request")

	_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{
		Creator:   suite.creator,
		PoolIndex: seniorPool,
		Token:     sdk.Coin{Denom: "ausdc", Amount: sdk.NewInt(70000).Mul(sdk.NewIntFromBigInt(base))},
	})
	suite.Require().NoError(err)

	_, err = suite.app.PayPrincipalForWithdrawalRequests(suite.ctx, &types.MsgPayPrincipalPartial{
		Creator:   suite.creator,
		PoolIndex: seniorPool,
		Token: sdk.Coin{
			Denom:  "ausdc",
			Amount: sdk.NewInt(800000).Mul(sdk.NewIntFromBigInt(base)),
		},
	})
	suite.Require().ErrorContains(err, "no withdraw proposal to be paid: invalid request")

	poolInfo, _ = suite.keeper.GetPools(suite.ctx, seniorPool)
	// principalEscrow := poolInfo.EscrowPrincipalAmount.Amount

	addr, err := sdk.AccAddressFromBech32(suite.investors[0])
	suite.Require().NoError(err)
	// d0, _ := suite.keeper.GetDepositor(suite.ctx, seniorPool, addr)
	// suite.Require().True(newinvestor1.WithdrawalAmount.Amount.Equal(sdk.NewIntFromUint64(uint64(seniorAmounts[0])).Mul(sdk.NewIntFromBigInt(base))))
	addr, err = sdk.AccAddressFromBech32(suite.investors[1])
	suite.Require().NoError(err)
	// d1, _ := suite.keeper.GetDepositor(suite.ctx, seniorPool, addr)

	// totalLocked := d0.LockedAmount.AddAmount(d1.LockedAmount.Amount)

	poolInfo, ok = suite.keeper.GetPools(suite.ctx, seniorPool)
	suite.Require().True(ok)
	suite.Require().True(checkValueEqualWithExchange(poolInfo.BorrowedAmount.Amount, sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(800000), base))))
	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base))))
	poolInfo, ok = suite.keeper.GetPools(suite.ctx, juniorPool)
	suite.Require().True(ok)
	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(100000), base))))
	suite.Require().True(checkValueEqualWithExchange(poolInfo.BorrowedAmount.Amount, sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base))))

	// we simulate 6 seconds for each block
	deltaTime := time.Second * time.Duration(60)
	checkPoints := []time.Duration{
		time.Duration(oneMonth) * time.Second,
		time.Duration(oneYear)*time.Second - time.Hour*2,
		time.Duration(oneYear) * time.Second,
	}
	withdrawRequest := []time.Duration{
		time.Duration(oneMonth*6) * time.Second,
	}

	checkPointCounter := 0

	allCoinsSenior := make(map[int]sdk.Coins)
	allCoinsJunior := make(map[int]sdk.Coins)
	processed := false

	for {
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(deltaTime))
		spv.EndBlock(suite.ctx, *suite.keeper)

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
			coins := suite.getAllInvestorInterest(seniorPool, suite.investors[:6], -1)
			allCoinsSenior[checkPointCounter] = coins
			coins2 := suite.getAllInvestorInterest(juniorPool, suite.investors[:6], -1)
			allCoinsJunior[checkPointCounter] = coins2

			if checkPointCounter == 1 {

				principalPayment := new(big.Int).Mul(big.NewInt(1000000), base)

				// here we pay the partial principal
				_, err = suite.app.PayPrincipalForWithdrawalRequests(suite.ctx, &types.MsgPayPrincipalPartial{
					Creator:   suite.creator,
					PoolIndex: seniorPool,
					Token:     sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(principalPayment)),
				})
				suite.Require().NoError(err)

			}

			checkPointCounter++
			if checkPointCounter == len(checkPoints) {
				break
			}
		}
	}

	addr, err = sdk.AccAddressFromBech32(suite.investors[0])
	suite.Require().NoError(err)
	newinvestor1, _ := suite.keeper.GetDepositor(suite.ctx, seniorPool, addr)
	// suite.Require().True(newinvestor1.WithdrawalAmount.Amount.Equal(sdk.NewIntFromUint64(uint64(seniorAmounts[0])).Mul(sdk.NewIntFromBigInt(base))))
	_ = newinvestor1
	addr, err = sdk.AccAddressFromBech32(suite.investors[1])
	suite.Require().NoError(err)
	newinvestor2, _ := suite.keeper.GetDepositor(suite.ctx, seniorPool, addr)
	_ = newinvestor2
	// suite.Require().True(newinvestor2.WithdrawalAmount.Amount.Equal(sdk.NewIntFromUint64(uint64(seniorAmounts[1])).Mul(sdk.NewIntFromBigInt(base))))

	// total 59499999999999975519844
	coins := suite.getAllInvestorInterest(seniorPool, suite.investors[:6], -1)
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
	totalToInvestorUSD := convertBorrowToUsd(totalToInvestor.Amount)
	reserve, ok := suite.keeper.GetReserve(suite.ctx, "ausdc")
	suite.Require().True(ok)
	totalInterest := reserve.AddAmount(totalToInvestorUSD)
	suite.Require().True(sdk.NewIntFromUint64(100000).Sub(totalInterest.Amount).LTE(sdk.NewIntFromUint64(1000000)))

	// seniorPoolInfo, _ := suite.keeper.GetPools(suite.ctx, seniorPool)
	// juniorPoolInfo, _ := suite.keeper.GetPools(suite.ctx, juniorPool)
	// totalLeftInterest := seniorPoolInfo.EscrowInterestAmount.Add(juniorPoolInfo.EscrowInterestAmount)
	// suite.Require().True(totalLeftInterest.LT(sdk.NewIntFromUint64(100)))

	addr, err = sdk.AccAddressFromBech32(suite.investors[0])
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, seniorPool)
	suite.Require().True(found)

	//amount := new(big.Int).Mul(big.NewInt(2000), base)
	//_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{
	//	Creator:   suite.creator,
	//	PoolIndex: seniorPool,
	//	Token:     sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(amount)),
	//})
	//suite.Require().NoError(err)
	//

	//
	//principalPayment := new(big.Int).Mul(big.NewInt(1000000), base)
	//
	//// here we pay the partial principal
	//_, err = suite.app.PayPrincipalForWithdrawalRequests(suite.ctx, &types.MsgPayPrincipalPartial{
	//	Creator:   suite.creator,
	//	PoolIndex: seniorPool,
	//	Token:     sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(principalPayment)),
	//})
	//suite.Require().NoError(err)

	newinvestor1, _ = suite.keeper.GetDepositor(suite.ctx, seniorPool, addr)
	// suite.Require().True(newinvestor1.WithdrawalAmount.Amount.Equal(sdk.NewIntFromUint64(uint64(seniorAmounts[0])).Mul(sdk.NewIntFromBigInt(base))))

	poolInfo, found = suite.keeper.GetPools(suite.ctx, seniorPool)
	suite.Require().True(found)
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Hour * 2))
	ret := suite.keeper.HandlePartialPrincipalPayment(suite.ctx, &poolInfo, poolInfo.WithdrawAccounts)
	suite.Require().True(ret)

	// manually update the pool
	suite.keeper.SetPool(suite.ctx, poolInfo)

	newinvestor1, _ = suite.keeper.GetDepositor(suite.ctx, seniorPool, addr)
	// suite.Require().True(newinvestor1.WithdrawalAmount.Amount.Equal(sdk.NewIntFromUint64(uint64(seniorAmounts[0])).Mul(sdk.NewIntFromBigInt(base))))

	// now we withdraw
	// ####################withdraw principal for investor who submit the withdrawal proposal##################################
	totalWithdrawal := sdk.ZeroInt()
	resp, err := suite.app.WithdrawPrincipal(suite.ctx, &types.MsgWithdrawPrincipal{Creator: suite.investors[0], PoolIndex: seniorPool, Token: sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1000e10))})
	suite.Require().NoError(err)
	token, _ := sdk.ParseCoinNormalized(resp.Amount)
	suite.Require().True(checkValueWithRangeTwo(sdk.NewIntFromUint64(uint64(seniorAmounts[0])).Mul(sdk.NewIntFromBigInt(base)), token.Amount.Sub(allCoinsSenior[2][0].Amount)))

	totalWithdrawal = totalWithdrawal.Add(token.Amount)
	resp, err = suite.app.WithdrawPrincipal(suite.ctx, &types.MsgWithdrawPrincipal{Creator: suite.investors[1], PoolIndex: seniorPool, Token: sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1))})
	suite.Require().NoError(err)
	// ######################################################

	token, _ = sdk.ParseCoinNormalized(resp.Amount)
	suite.Require().True(checkValueWithRangeTwo(sdk.NewIntFromUint64(uint64(seniorAmounts[1])).Mul(sdk.NewIntFromBigInt(base)), token.Amount.Sub(allCoinsSenior[2][1].Amount)))
	totalWithdrawal = totalWithdrawal.Add(token.Amount)

	// total withdrawal should exclude the interest
	totalWithdrawal = totalWithdrawal.Sub(allCoinsSenior[1][0].Amount)
	totalWithdrawal = totalWithdrawal.Sub(allCoinsSenior[1][1].Amount)

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
			Creator:   el,
			PoolIndex: seniorPool,
		})
		suite.Require().NoError(err)
		_, err = suite.app.ClaimInterest(suite.ctx, &types.MsgClaimInterest{
			Creator:   el,
			PoolIndex: juniorPool,
		})
		suite.Require().NoError(err)
	}

	for _, el := range suite.investors[:2] {
		_, err := suite.app.ClaimInterest(suite.ctx, &types.MsgClaimInterest{
			Creator:   el,
			PoolIndex: juniorPool,
		})
		suite.Require().NoError(err)
	}

	//_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{
	//	Creator:   suite.creator,
	//	PoolIndex: juniorPool,
	//	Token:     sdk.Coin{Denom: "ausdc", Amount: sdk.NewInt(30000).Mul(sdk.NewIntFromBigInt(base))},
	//})
	//suite.Require().NoError(err)
	//
	//_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{
	//	Creator:   suite.creator,
	//	PoolIndex: seniorPool,
	//	Token:     sdk.Coin{Denom: "ausdc", Amount: sdk.NewInt(70000).Mul(sdk.NewIntFromBigInt(base))},
	//})
	//suite.Require().NoError(err)
	//suite.Require().ErrorContains(err, "we have the prepayment interest, not accepting new interest payment")

	// now we can test another 2 months, spv pay the interest, then the spv close the pool

	p1, _ := suite.keeper.GetPools(suite.ctx, seniorPool)
	p2, _ := suite.keeper.GetPools(suite.ctx, juniorPool)

	seniorLocked := p1.BorrowedAmount
	juniorLocked := p2.BorrowedAmount
	// senior apy is
	seniorInterest := sdk.NewDecFromInt(seniorLocked.Amount).Mul(sdk.NewDecWithPrec(875, 4))
	juniorInterest := sdk.NewDecFromInt(juniorLocked.Amount).Mul(sdk.NewDecWithPrec(15, 2))

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
			suite.Require().ErrorContains(err, "principal is not fully paid. you have paid 1 and borrowed")

			amount := new(big.Int).Mul(big.NewInt(2000), base)
			_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{
				Creator:   suite.creator,
				PoolIndex: seniorPool,
				Token:     sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(amount)),
			})
			suite.Require().NoError(err)

			amount = new(big.Int).Mul(big.NewInt(200000), base)
			_, err = suite.app.PayPrincipal(suite.ctx, &types.MsgPayPrincipal{
				Creator:   suite.creator,
				PoolIndex: seniorPool,
				Token:     sdk.Coin{Denom: "ausdc", Amount: sdk.NewIntFromBigInt(amount)},
			})
			suite.Require().NoError(err)

			amount = new(big.Int).Mul(big.NewInt(200000), base)
			_, err = suite.app.PayPrincipal(suite.ctx, &types.MsgPayPrincipal{
				Creator:   suite.creator,
				PoolIndex: juniorPool,
				Token:     sdk.Coin{Denom: "ausdc", Amount: sdk.NewIntFromBigInt(amount)},
			})
			suite.Require().NoError(err)

			break
		}
	}

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

	expectedYearSenior := sdk.NewDecFromInt(totalSeniorInterest).Mul(sdk.NewDec(52)).Quo(sdk.NewDec(8)).QuoTruncate(sdk.NewDecWithPrec(85, 2))
	expectedYearJunior := sdk.NewDecFromInt(totalJuniorInterest).Mul(sdk.NewDec(52)).Quo(sdk.NewDec(8)).QuoTruncate(sdk.NewDecWithPrec(85, 2))

	oneWeekTotal := sdk.NewDecFromInt(totalSeniorInterest).Quo(sdk.NewDec(8)).Add(sdk.NewDecFromInt(totalJuniorInterest).Quo(sdk.NewDec(8)))

	suite.Require().True(sdk.NewDecFromInt(convertBorrowToUsd(seniorInterest.TruncateInt())).Sub(expectedYearSenior).LT(sdk.NewDecWithPrec(1e8, 0)))
	suite.Require().True(sdk.NewDecFromInt(convertBorrowToUsd(juniorInterest.TruncateInt())).Sub(expectedYearJunior).LT(sdk.NewDecWithPrec(1e8, 0)))

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * time.Duration(oneWeek)))
	spv.EndBlock(suite.ctx, *suite.keeper)
	p, _ := suite.keeper.GetPools(suite.ctx, seniorPool)
	suite.Require().Equal(p.PoolStatus, types.PoolInfo_FROZEN)
	p, _ = suite.keeper.GetPools(suite.ctx, juniorPool)
	suite.Require().Equal(p.PoolStatus, types.PoolInfo_FROZEN)

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second))

	for _, el := range suite.investors[2:6] {
		resp, err := suite.app.WithdrawPrincipal(suite.ctx, &types.MsgWithdrawPrincipal{
			Creator:   el,
			PoolIndex: seniorPool,
			Token:     sdk.NewCoin("ausdc", sdk.NewInt(1)),
		})
		tokens, err = sdk.ParseCoinsNormalized(resp.Amount)
		suite.Require().NoError(err)
		totalWithdrawal = totalWithdrawal.Add(tokens[0].Amount)
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
		totalWithdrawal = totalWithdrawal.Add(tokens[0].Amount)
		_, found = suite.keeper.GetDepositor(suite.ctx, juniorPool, addr)
		suite.Require().False(found)
	}

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
	// deltaPrincipal := totalWithdrawal.Sub(sdk.NewIntFromUint64(1300000).Mul(sdk.NewIntFromBigInt(base)).Add(oneWeekTotal.TruncateInt()))
	suite.Require().True(totalWithdrawal.Sub(sdk.NewIntFromUint64(1300000).Mul(sdk.NewIntFromBigInt(base)).Add(oneWeekTotal.TruncateInt())).Abs().LTE(sdk.NewIntFromUint64(10)))

	_, found = suite.keeper.GetPools(suite.ctx, seniorPool)
	suite.Require().False(found)
	_, found = suite.keeper.GetPools(suite.ctx, juniorPool)
	suite.Require().False(found)
}

func (suite *mockWholeProcessSuite) TestMockSystemOneYearWithWithdrawalTransferNFTPartially() {
	setupMockPool(suite)

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
	// Generate  random integer for junior pool

	juniorAmounts, _ := generateRandomIntegersWithSum(6, 300000)
	seniorAmounts, _ := generateRandomIntegersWithSum(6, 1000000)

	juniorAmountsLastTwo, _ := generateRandomIntegersWithSum(2, 2000)
	seniorAmountsLastTwo, _ := generateRandomIntegersWithSum(2, 3000)

	seniorPool := "0x8083eaa3584b60c163fe63e5ab6937526022cd47c35f9cb1e0790005c3ae9d00"
	juniorPool := "0xe0d49c3eed41e408b493a14042a8aa31375d64e3e357f911afbb085e02bde083"
	msgDepositUsersJunior := make([]*types.MsgDeposit, 6)
	msgDepositUsersSenior := make([]*types.MsgDeposit, 6)
	for i := 0; i < 6; i++ {
		token := sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(int64(juniorAmounts[i])), base)))
		msg := &types.MsgDeposit{
			Creator:   suite.investors[i],
			PoolIndex: juniorPool,
			Token:     token,
		}
		msgDepositUsersJunior[i] = msg

		token = sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(int64(seniorAmounts[i])), base)))
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
	suite.keeper.SetPool(suite.ctx, poolInfo)

	poolInfo, ok = suite.keeper.GetPools(suite.ctx, juniorPool)
	suite.Require().True(ok)
	suite.Require().Equal(poolInfo.Apy.String(), "0.150000000000000000")

	poolInfo.PoolLockedSeconds = 360000
	poolInfo.PoolTotalBorrowLimit = 100000
	poolInfo.WithdrawRequestWindowSeconds = oneMonth * 3
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
	token1 := sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(800000), base))
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(oneWeek))
	msgSenior := types.MsgBorrow{Creator: suite.creator, PoolIndex: seniorPool, BorrowAmount: sdk.NewCoin("ausdc", token1)}
	_, err := suite.app.Borrow(suite.ctx, &msgSenior)
	suite.Require().NoError(err)

	token2 := sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base))
	//  borrow another one
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Minute))
	msgJunior := types.MsgBorrow{Creator: suite.creator, PoolIndex: juniorPool, BorrowAmount: sdk.NewCoin("ausdc", token2)}
	_, err = suite.app.Borrow(suite.ctx, &msgJunior)
	suite.Require().NoError(err)

	seniorInterest := sdk.NewDecFromInt(token1).Mul(sdk.NewDecWithPrec(875, 4)).Mul(sdk.NewDecWithPrec(85, 2))
	juniorInterest := sdk.NewDecFromInt(token2).Mul(sdk.NewDecWithPrec(15, 2)).Mul(sdk.NewDecWithPrec(85, 2))

	startTime := suite.ctx.BlockTime()

	poolInfo, _ = suite.keeper.GetPools(suite.ctx, seniorPool)

	poolInfo, ok = suite.keeper.GetPools(suite.ctx, seniorPool)
	suite.Require().True(ok)
	suite.Require().True(checkValueEqualWithExchange(poolInfo.BorrowedAmount.Amount, sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(800000), base))))
	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base))))
	poolInfo, ok = suite.keeper.GetPools(suite.ctx, juniorPool)
	suite.Require().True(ok)
	suite.Require().True(poolInfo.UsableAmount.Amount.Equal(sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(100000), base))))
	suite.Require().True(checkValueEqualWithExchange(poolInfo.BorrowedAmount.Amount, sdk.NewIntFromBigInt(new(big.Int).Mul(big.NewInt(200000), base))))

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

		if startTime.Add(checkPoints[checkPointCounter]).Before(suite.ctx.BlockTime()) {
			coins := suite.getAllInvestorInterest(seniorPool, suite.investors[:6], -1)
			allCoinsSenior[checkPointCounter] = coins
			coins2 := suite.getAllInvestorInterest(juniorPool, suite.investors[:6], -1)
			allCoinsJunior[checkPointCounter] = coins2
			checkPointCounter++
			if checkPointCounter == 1 {
				break
			}
		}
	}

	seniorTotalInterest := sdk.ZeroInt()
	juniorTotalInterest := sdk.ZeroInt()

	amount := new(big.Int).Mul(big.NewInt(20000), base)
	_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{
		Creator:   suite.creator,
		PoolIndex: seniorPool,
		Token:     sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(amount)),
	})
	suite.Require().NoError(err)

	_, err = suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{
		Creator:   suite.creator,
		PoolIndex: juniorPool,
		Token:     sdk.NewCoin("ausdc", sdk.NewIntFromBigInt(amount)),
	})
	suite.Require().NoError(err)

	// get all interests
	for _, el := range suite.investors[:6] {

		r, err := suite.app.ClaimInterest(suite.ctx, &types.MsgClaimInterest{
			Creator:   el,
			PoolIndex: seniorPool,
		})
		suite.Require().NoError(err)

		tokens, err := sdk.ParseCoinsNormalized(r.Amount)
		suite.Require().NoError(err)

		seniorTotalInterest = seniorTotalInterest.Add(tokens[0].Amount)

		r, err = suite.app.ClaimInterest(suite.ctx, &types.MsgClaimInterest{
			Creator:   el,
			PoolIndex: juniorPool,
		})
		suite.Require().NoError(err)

		tokens, err = sdk.ParseCoinsNormalized(r.Amount)
		suite.Require().NoError(err)
		juniorTotalInterest = juniorTotalInterest.Add(tokens[0].Amount)

	}

	suite.Require().True(juniorTotalInterest.MulRaw(13).Sub(juniorInterest.TruncateInt()).LT(sdk.NewIntFromUint64(1e8)))
	suite.Require().True(seniorTotalInterest.MulRaw(13).Sub(seniorInterest.TruncateInt()).LT(sdk.NewIntFromUint64(1e8)))

	addr11, _ := sdk.AccAddressFromBech32(suite.investors[1])
	dbeforeTransfer, found := suite.keeper.GetDepositor(suite.ctx, juniorPool, addr11)
	suite.Require().True(found)

	addr10, _ := sdk.AccAddressFromBech32(suite.investors[0])
	dbeforeTransferU1, found := suite.keeper.GetDepositor(suite.ctx, seniorPool, addr10)
	suite.Require().True(found)

	// now we transfer owner
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Hour))
	for i, el := range suite.investors[6:] {

		_, err := suite.app.Deposit(suite.ctx, &types.MsgDeposit{
			Creator:   el,
			PoolIndex: seniorPool,
			Token:     sdk.NewCoin("ausdc", sdk.NewIntFromUint64(uint64(seniorAmountsLastTwo[i])).Mul(sdk.NewIntFromBigInt(base))),
		})
		suite.Require().NoError(err)

		_, err = suite.app.Deposit(suite.ctx, &types.MsgDeposit{
			Creator:   el,
			PoolIndex: juniorPool,
			Token:     sdk.NewCoin("ausdc", sdk.NewIntFromUint64(uint64(juniorAmountsLastTwo[i])).Mul(sdk.NewIntFromBigInt(base))),
		})
		suite.Require().NoError(err)
	}

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Hour))

	_, err = suite.app.TransferOwnership(suite.ctx, &types.MsgTransferOwnership{
		Creator:   suite.investors[0],
		PoolIndex: seniorPool,
	})
	suite.Require().NoError(err)

	_, err = suite.app.TransferOwnership(suite.ctx, &types.MsgTransferOwnership{
		Creator:   suite.investors[1],
		PoolIndex: juniorPool,
	})

	suite.Require().NoError(err)

	// first month after transfer ownership request sent
	checkPointCounter = 0
	newStartTime := suite.ctx.BlockTime()
	for {
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(deltaTime))
		spv.EndBlock(suite.ctx, *suite.keeper)

		if newStartTime.Add(checkPoints[checkPointCounter]).Before(suite.ctx.BlockTime()) {
			checkPointCounter++
			if checkPointCounter == 1 {
				break
			}
		}
	}

	// investor 1 should can be fully withdrawed
	d, found := suite.keeper.GetDepositor(suite.ctx, juniorPool, addr11)
	suite.Require().True(found)
	suite.Require().Equal(d.DepositType, types.DepositorInfo_deposit_close)

	// total 59499999999999975519844
	allSeniorsInterest := suite.getAllInvestorInterest(seniorPool, suite.investors[:8], -1)
	allJuniorInterest := suite.getAllInvestorInterest(juniorPool, suite.investors[:8], -1)

	sumSenior := sdk.NewCoin("ausdc", sdk.ZeroInt())
	sumJunior := sdk.NewCoin("ausdc", sdk.ZeroInt())
	for _, el := range allSeniorsInterest {
		sumSenior = sumSenior.Add(el)
	}

	for _, el := range allJuniorInterest {
		sumJunior = sumJunior.Add(el)
	}

	suite.Require().True(sumJunior.Amount.MulRaw(13).Sub(juniorInterest.TruncateInt()).LT(sdk.NewIntFromUint64(1e8)))
	suite.Require().True(sumSenior.Amount.MulRaw(13).Sub(seniorInterest.TruncateInt()).LT(sdk.NewIntFromUint64(1e8)))

	// check investor 1 interest

	// investor 1 can withdraw all the amount
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Hour))
	resp, err := suite.app.WithdrawPrincipal(suite.ctx, &types.MsgWithdrawPrincipal{
		Creator:   suite.investors[1],
		PoolIndex: juniorPool,
		Token:     sdk.NewCoin("ausdc", sdk.NewInt(1)),
	})
	suite.Require().NoError(err)

	principal := sdk.NewIntFromUint64(uint64(juniorAmounts[1])).Mul(sdk.NewIntFromBigInt(base))

	// when we call the withdraw, we return the withdrawal amount
	newprincipal := principal.Sub(dbeforeTransfer.WithdrawalAmount.Amount)

	expected := newprincipal.Add(allJuniorInterest[1].Amount)
	tokens, _ := sdk.ParseCoinsNormalized(resp.Amount)
	suite.Require().True(checkValueWithRangeTwo(expected, tokens[0].Amount))

	// we check whether the total borrow is equal to the estimated
	totalLockedSenior := sdk.NewCoin("aud-ausdc", sdk.ZeroInt())
	totalWithdrawalbleSenior := sdk.NewCoin("ausdc", sdk.ZeroInt())
	totalLockedJunior := sdk.NewCoin("aud-ausdc", sdk.ZeroInt())
	totalWithdrawalbleJunior := sdk.NewCoin("ausdc", sdk.ZeroInt())
	for i, el := range suite.investors {
		addr, err := sdk.AccAddressFromBech32(el)
		suite.Require().NoError(err)
		d, found := suite.keeper.GetDepositor(suite.ctx, seniorPool, addr)
		suite.Require().True(found)
		if i != 1 {
			d1, found := suite.keeper.GetDepositor(suite.ctx, juniorPool, addr)
			suite.Require().True(found)
			totalLockedJunior = totalLockedJunior.Add(d1.LockedAmount)
			totalWithdrawalbleJunior = totalWithdrawalbleJunior.Add(d1.WithdrawalAmount)
		}
		totalLockedSenior = totalLockedSenior.Add(d.LockedAmount)
		totalWithdrawalbleSenior = totalWithdrawalbleSenior.Add(d.WithdrawalAmount)
	}

	totalLockedJuniorUsd := sdk.NewCoin(poolInfo.TargetAmount.Denom, convertBorrowToUsd(totalLockedJunior.Amount))
	totalLockedSeniorUsd := sdk.NewCoin(poolInfo.TargetAmount.Denom, convertBorrowToUsd(totalLockedSenior.Amount))
	suite.Require().True(checkValueWithRangeTwo(totalLockedJuniorUsd.Add(totalWithdrawalbleJunior).Amount, sdk.NewIntFromUint64(uint64(302000-juniorAmounts[1])).Mul(sdk.NewIntFromBigInt(base))))

	suite.Require().True(checkValueWithRangeTwo(totalLockedSeniorUsd.Add(totalWithdrawalbleSenior).Amount.Add(dbeforeTransferU1.WithdrawalAmount.Amount), sdk.NewIntFromUint64(uint64(1003000)).Mul(sdk.NewIntFromBigInt(base))))

	seniorRatios := make([]sdk.Dec, len(suite.investors))
	juniorRatios := make([]sdk.Dec, len(suite.investors))

	for i, el := range suite.investors {
		addr, err := sdk.AccAddressFromBech32(el)
		suite.Require().NoError(err)
		d, found := suite.keeper.GetDepositor(suite.ctx, seniorPool, addr)
		suite.Require().True(found)
		if i != 1 {
			d1, found := suite.keeper.GetDepositor(suite.ctx, juniorPool, addr)
			suite.Require().True(found)
			ratio1 := sdk.NewDecFromInt(d1.LockedAmount.Amount).Quo(sdk.NewDecFromInt(totalLockedJunior.Amount))
			juniorRatios[i] = ratio1
		}
		ratio := sdk.NewDecFromInt(d.LockedAmount.Amount).Quo(sdk.NewDecFromInt(totalLockedSenior.Amount))
		seniorRatios[i] = ratio
	}

	// clean up the interest
	for i, el := range suite.investors {
		resp, err := suite.app.ClaimInterest(suite.ctx, &types.MsgClaimInterest{
			Creator:   el,
			PoolIndex: seniorPool,
		})

		tokens, _ := sdk.ParseCoinsNormalized(resp.Amount)
		suite.Require().True(tokens[0].Amount.Equal(allSeniorsInterest[i].Amount))
		suite.Require().NoError(err)

		if i != 1 {
			resp, err := suite.app.ClaimInterest(suite.ctx, &types.MsgClaimInterest{
				Creator:   el,
				PoolIndex: juniorPool,
			})
			tokens, _ := sdk.ParseCoinsNormalized(resp.Amount)
			suite.Require().True(tokens[0].Amount.Equal(allJuniorInterest[i].Amount))
			suite.Require().NoError(err)
		}
	}

	// now we calculate get the interest again, it should be distributed according to the distribution ratio

	// we need to increase the blocktime to avoid pay the nodes again for the past time
	checkPointCounter = 0
	newStartTime = suite.ctx.BlockTime()
	for {
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(deltaTime))
		spv.EndBlock(suite.ctx, *suite.keeper)
		if newStartTime.Add(checkPoints[checkPointCounter]).Before(suite.ctx.BlockTime()) {
			checkPointCounter++
			if checkPointCounter == 1 {
				break
			}
		}
	}

	allSeniorsInterest = suite.getAllInvestorInterest(seniorPool, suite.investors, -1)
	allJuniorInterest = suite.getAllInvestorInterest(juniorPool, suite.investors, 1)

	sumSenior = sdk.NewCoin("ausdc", sdk.ZeroInt())
	for _, el := range allSeniorsInterest {
		sumSenior = sumSenior.Add(el)
	}

	sumJunior = sdk.NewCoin("ausdc", sdk.ZeroInt())
	for _, el := range allJuniorInterest {
		sumJunior = sumJunior.Add(el)
	}

	for i, el := range allSeniorsInterest {
		ratioGet := sdk.NewDecFromInt(el.Amount).QuoTruncate(sdk.NewDecFromInt(sumSenior.Amount))
		suite.Require().True(ratioGet.Sub(seniorRatios[i]).Abs().LT(sdk.NewDecWithPrec(1, 8)))
	}

	for i, el := range allJuniorInterest {
		ratioGet := sdk.NewDecFromInt(el.Amount).QuoTruncate(sdk.NewDecFromInt(sumJunior.Amount))
		if i != 1 {
			suite.Require().True(ratioGet.Sub(juniorRatios[i]).Abs().LT(sdk.NewDecWithPrec(1, 8)))
		}
	}

	return
}

func (suite *mockWholeProcessSuite) getAllInvestorInterest(poolType string, investors []string, skip int) sdk.Coins {
	interests := make(sdk.Coins, len(investors))
	for i, el := range investors {
		if i != skip {
			resp, err := suite.keeper.ClaimableInterest(suite.ctx, &types.QueryClaimableInterestRequest{Wallet: el, PoolIndex: poolType})
			suite.Require().NoError(err)
			interests[i] = resp.ClaimableInterestAmount
		} else {
			interests[i] = sdk.NewCoin("ausdc", sdk.ZeroInt())
		}
	}
	return interests
}
