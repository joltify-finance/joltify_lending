package main

import (
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"

	"github.com/joltify-finance/joltify_lending/contrib/devnet/integrationtest/common"
)

// claimInterest randomly choose some users to claim interest, total investor is the total number of investors
// wants to claim interest (not all the users are capable of claiming interest)
func claimInterest(poolIndex string, claimUsersNum, totalInvestors int) {
	rand.Seed(time.Now().UnixNano()) // seed the random number generator

	usersClaim := make([]int, claimUsersNum)
	for i := 0; i < claimUsersNum; i++ {
		num := rand.Intn(totalInvestors) + 1 // generate a random number between 1 and y
		usersClaim[i] = num
	}
	logger.Info().Msgf("we do the interest claim with users %v", usersClaim)
	var actuallyDone []int
	locker := sync.RWMutex{}
	wg := sync.WaitGroup{}
	wg.Add(claimUsersNum)
	for _, user := range usersClaim {
		go func(userIndex int) {
			defer wg.Done()
			_, err := common.RunCommandWithOutput("./claim_interest.sh", poolIndex, strconv.Itoa(userIndex))
			_ = err // we ignore the err here as some account may not have deposit
			if err != nil {
				logger.Debug().Msgf("we got error %v when claim interest for user %d", err, userIndex)
				return
			}
			locker.Lock()
			actuallyDone = append(actuallyDone, userIndex)
			locker.Unlock()
		}(user)
	}
	wg.Wait()
	logger.Info().Msgf("we actually done the interest claim with users %v", actuallyDone)
}

func withdrawOrDeposit(poolIndex string, claimUsersNum, totalInvestors int, all, withdraw bool) map[int]int {
	rand.Seed(time.Now().UnixNano()) // seed the random number generator
	usersClaimData := make(map[int]int)
	actuallyDone := make(map[int]int)
	for i := 0; i < claimUsersNum; i++ {
		num := rand.Intn(totalInvestors) + 1 // generate a random number between 1 and y
		amount := rand.Intn(100) + 1         // generate a random number between 1 and y
		if all {
			amount = 100000 // generate a random number between 1 and y
		}
		if all {
			usersClaimData[i] = amount
		} else {
			usersClaimData[num] = amount
		}
	}
	locker := sync.RWMutex{}
	wg := sync.WaitGroup{}
	wg.Add(claimUsersNum)
	for k, v := range usersClaimData {
		go func(userIndex, amount int) {
			defer wg.Done()
			var out string
			var err error
			if withdraw {
				out, err = common.RunCommandWithOutput("./withdraw.sh", poolIndex, strconv.Itoa(amount), strconv.Itoa(userIndex))
			} else {
				out, err = common.RunCommandWithOutput("./deposit.sh", poolIndex, strconv.Itoa(amount), strconv.Itoa(userIndex))
			}
			if err != nil {
				logger.Error().Err(err).Msgf("we got error %v when call withdraw for user %d, %v", err, userIndex, out)
				return
			}
			locker.Lock()
			actuallyDone[userIndex] = amount
			locker.Unlock()
		}(k, v)
	}
	wg.Wait()
	if withdraw {
		logger.Info().Msgf("we actually done the withdraw with users %v", actuallyDone)
	} else {
		logger.Info().Msgf("we actually done the deposit with users %v", actuallyDone)
	}
	return actuallyDone
}

func transferOwnership(poolIndex string, claimUsersNum, totalInvestors int) {
	rand.Seed(time.Now().UnixNano()) // seed the random number generator
	usersClaim := make([]int, claimUsersNum)
	for i := 0; i < claimUsersNum; i++ {
		num := rand.Intn(totalInvestors) + 1 // generate a random number between 1 and y
		usersClaim[i] = num
	}
	logger.Info().Msgf("we do the transfer owner with users %v", usersClaim)
	var actuallydone []int
	locker := sync.RWMutex{}
	wg := sync.WaitGroup{}
	wg.Add(claimUsersNum)
	for _, user := range usersClaim {
		go func(userIndex int) {
			defer wg.Done()
			_, err := common.RunCommandWithOutput("./transfer.sh", poolIndex, strconv.Itoa(userIndex))
			_ = err // we ignore the err here as some account may not have deposit
			if err != nil {
				logger.Debug().Msgf("we got error %v when transfer for user %d", err, userIndex)
				return
			}
			locker.Lock()
			actuallydone = append(actuallydone, userIndex)
			locker.Unlock()
		}(user)
	}
	wg.Wait()
	logger.Info().Msgf("we actually done the transfer owner with users %v", actuallydone)
}

func submitWithdraw(poolIndex string, claimUsersNum, totalInvestors int) {
	rand.Seed(time.Now().UnixNano()) // seed the random number generator
	usersClaim := make([]int, claimUsersNum)
	for i := 0; i < claimUsersNum; i++ {
		num := rand.Intn(totalInvestors) + 1 // generate a random number between 1 and y
		usersClaim[i] = num
	}
	logger.Info().Msgf("we do the submit withdraw with users %v", usersClaim)

	var actuallydone []int
	locker := sync.RWMutex{}
	wg := sync.WaitGroup{}
	wg.Add(claimUsersNum)
	for _, user := range usersClaim {
		go func(userIndex int) {
			defer wg.Done()
			out, err := common.RunCommandWithOutput("./partial_withdraw.sh", poolIndex, strconv.Itoa(userIndex))
			_ = err // we ignore the err here as some account may not have deposit
			if err != nil {
				// logger.Debug().Msgf("we got error %v when claim interest for user %d", err, userIndex)
				logger.Error().Err(err).Msgf("we got error %v when call partial withdraw  for user %d %v", err, userIndex, out)
				return
			}
			locker.Lock()
			actuallydone = append(actuallydone, userIndex)
			locker.Unlock()
		}(user)
	}
	wg.Wait()
	logger.Info().Msgf("we actually done the withdraw request with users %v", actuallydone)
}

func payPrincipalPartial(poolIndex string) {
	out, err := common.RunCommandWithOutput("./pay_partial_principal.sh", poolIndex, "200000")
	if err != nil {
		logger.Error().Err(err).Msgf("we get error when pay partial principal with error %v", out)
		return
	}
	logger.Info().Msgf("pay partial principal successfully")
}

func depositAndBorrow() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	ret := os.Getenv("INIT_INVESTORS")
	common.RunCommand("./run_normal_borrow_random.sh", ret)
	return nil
}

func payInterest(poolIndex string) error {
	common.RunCommand("./run_pay_interest.sh", poolIndex, "1000000")
	return nil
}

func startChain() {
	common.RunCommand("./start_chain.sh")
}
