package common

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"

	"github.com/gookit/color"
)

func RunCommand(cmdStr string, parameters ...string) {
	cmd := exec.Command(cmdStr, parameters...)
	cmd.Env = os.Environ()
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Printf("%s\n", m)
	}
	cmd.Wait()
}

func RunCommandWithOutput(cmdStr string, parameters ...string) (string, error) {
	var outb, errb bytes.Buffer
	cmd := exec.Command(cmdStr, parameters...)
	cmd.Env = os.Environ()
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	if err := cmd.Run(); err != nil {
		return errb.String(), err
	}
	return outb.String(), nil
}

func GetWindow(poolIndex string) (Window, SPV, error) {
	var poolInfo SPV
	out, err := RunCommandWithOutput("joltify", "q", "spv", "query-pool", poolIndex, "--output", "json")
	if err != nil {
		err = fmt.Errorf("%v fail to query pool info: %v", err, out)
		return Window{}, SPV{}, err
	}

	err = json.Unmarshal([]byte(out), &poolInfo)
	if err != nil {
		color.Red.Printf("fail to unmarshal")
		panic(err)
	}

	w := GetTimeWindow(poolInfo)

	return w, poolInfo, nil
}

type Window struct {
	PaymentDue          int
	WithdrawStartTime   int
	PayPartialStartTime int
}

func GetTimeWindow(poolInfo SPV) Window {
	projectDueTime := poolInfo.PoolInfo.ProjectDueTime

	proposalDate := projectDueTime.Add(-time.Second * time.Duration(poolInfo.PoolInfo.WithdrawRequestWindowSeconds*3))

	payPrincipalDueDate := projectDueTime.Add(-time.Second * time.Duration(poolInfo.PoolInfo.WithdrawRequestWindowSeconds*2-10))
	currentTime := time.Now()
	withdrawStart := proposalDate.Sub(currentTime).Seconds()
	payPrincipalStart := payPrincipalDueDate.Sub(currentTime).Seconds()

	paymentDate := poolInfo.PoolInfo.LastPaymentTime.Add(time.Duration(poolInfo.PoolInfo.PayFreq) * time.Second)

	paymentDue := int(paymentDate.Sub(currentTime).Seconds())

	w := Window{
		PaymentDue:          paymentDue,
		WithdrawStartTime:   int(withdrawStart),
		PayPartialStartTime: int(payPrincipalStart),
	}

	return w
}

func WithdrawOrDeposit(poolIndex string, claimUsersNum, totalInvestors int, all, withdraw bool) map[int]int {
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
			// var out string
			var err error
			if withdraw {
				_, err = RunCommandWithOutput("./scripts/withdraw.sh", poolIndex, strconv.Itoa(amount), strconv.Itoa(userIndex))
			} else {
				_, err = RunCommandWithOutput("./scripts/deposit.sh", poolIndex, strconv.Itoa(amount), strconv.Itoa(userIndex))
			}
			if err != nil {
				// logger.Error().Err(err).Msgf("we got error %v when call withdraw for user %d, %v", err, userIndex, out)
				return
			}
			locker.Lock()
			actuallyDone[userIndex] = amount
			locker.Unlock()
		}(k, v)
	}
	wg.Wait()
	if withdraw {
		logger.Debug().Msgf("we actually done the withdraw with users %v", actuallyDone)
	} else {
		logger.Debug().Msgf("we actually done the deposit with users %v", actuallyDone)
	}
	return actuallyDone
}
