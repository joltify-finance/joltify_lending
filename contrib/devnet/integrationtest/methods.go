package main

import (
	"math/rand"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"

	"github.com/joltify-finance/joltify_lending/contrib/devnet/integrationtest/common"
)

// claimInterest randomly choose some users to claim interest, total investor is the total number of investors
// wants to claim interest (not all the users are capable of claiming interest)
func claimInterest(poolIndex string, claimUsersNum, totalInvestors int) []int {
	usersClaim := rand.Perm(totalInvestors)[:claimUsersNum]

	logger.Debug().Msgf("we do the interest claim with users %v", usersClaim)
	var actuallyDone []int
	locker := sync.RWMutex{}
	wg := sync.WaitGroup{}
	wg.Add(claimUsersNum)
	for _, user := range usersClaim {
		go func(userIndex int) {
			defer wg.Done()
			out, err := common.RunCommandWithOutput("./scripts/claim_interest.sh", poolIndex, strconv.Itoa(userIndex))
			if err != nil {
				logger.Error().Err(err).Msgf("we got error %v when claim interest for user %d with error output %v", err, userIndex, out)
				return
			}
			locker.Lock()
			actuallyDone = append(actuallyDone, userIndex)
			locker.Unlock()
		}(user + 1)
	}
	wg.Wait()
	logger.Debug().Msgf("we actually done the interest claim with users %v", actuallyDone)
	return actuallyDone
}

func transferOwnership(poolIndex string, claimUsersNum, totalInvestors int) []int {
	usersClaim := rand.Perm(totalInvestors)[:claimUsersNum]
	logger.Debug().Msgf("we do the transfer owner with users %v", usersClaim)
	var actuallydone []int
	locker := sync.RWMutex{}
	wg := sync.WaitGroup{}
	wg.Add(claimUsersNum)
	for _, user := range usersClaim {
		go func(userIndex int) {
			defer wg.Done()
			out, err := common.RunCommandWithOutput("./scripts/transfer.sh", poolIndex, strconv.Itoa(userIndex))
			if err != nil {
				logger.Error().Err(err).Msgf("we got error %v when transfer for user %d with %v", err, userIndex, out)
				return
			}
			locker.Lock()
			actuallydone = append(actuallydone, userIndex)
			locker.Unlock()
		}(user + 1)
	}
	wg.Wait()
	logger.Debug().Msgf("we actually done the transfer owner with users %v", actuallydone)
	return actuallydone
}

func submitWithdraw(poolIndex string, claimUsersNum, totalinvestors int) []int {
	usersClaim := rand.Perm(totalinvestors)[:claimUsersNum]

	var actuallydone []int
	locker := sync.RWMutex{}
	wg := sync.WaitGroup{}
	wg.Add(claimUsersNum)
	for _, user := range usersClaim {
		go func(userIndex int) {
			defer wg.Done()
			out, err := common.RunCommandWithOutput("./scripts/partial_withdraw.sh", poolIndex, strconv.Itoa(userIndex))
			if err != nil {
				logger.Error().Err(err).Msgf("we got error %v when call partial withdraw  for user %d %v", err, userIndex, out)
				return
			}
			locker.Lock()
			actuallydone = append(actuallydone, userIndex)
			locker.Unlock()
		}(user + 1)
	}
	wg.Wait()
	logger.Debug().Msgf("we actually done the withdraw request with users %v", actuallydone)
	return actuallydone
}

func queryUSDCBalance(nickname string) (string, error) {
	out, err := common.RunCommandWithOutput("./scripts/show_balance.sh", nickname)
	if err != nil {
		logger.Error().Err(err).Msgf("we get error when pay partial principal with error %v", out)
		return "", err
	}
	return out, nil
}

func payPrincipalPartial(poolIndex string) {
	out, err := common.RunCommandWithOutput("./scripts/pay_partial_principal.sh", poolIndex, "200000")
	if err != nil {
		logger.Error().Err(err).Msgf("we get error when pay partial principal with error %v", out)
		return
	}
}

func depositAndBorrow() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	ret := os.Getenv("INIT_INVESTORS")
	common.RunCommand("./scripts/run_normal_borrow_random.sh", ret)
	return nil
}

func payInterest(poolIndex string) error {
	common.RunCommand("./scripts/run_pay_interest.sh", poolIndex, "1000000")
	return nil
}

func startChain() {
	common.RunCommand("./scripts/start_chain.sh")
}
