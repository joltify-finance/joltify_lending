package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"html"
	"math/big"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/gookit/color"
	"github.com/joho/godotenv"
	"github.com/joltify-finance/joltify_lending/contrib/devnet/integrationtest/common"
	zlog "github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	logger         = log.With().Logger()
	needWrite      = false
	inputTimeout   = time.Second * 5
	base           = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	transferAmount = sdk.NewInt(0)
	wrong          = html.UnescapeString("&#" + "10060" + ";")
	correct        = html.UnescapeString("&#" + "9989" + ";")
)

type Window struct {
	paymentDue          int
	withdrawStartTime   int
	payPartialStartTime int
}

func getTimeWindow(poolInfo common.SPV) Window {
	projectDueTime := poolInfo.PoolInfo.ProjectDueTime

	proposalDate := projectDueTime.Add(-time.Second * time.Duration(poolInfo.PoolInfo.WithdrawRequestWindowSeconds*3))

	payPrincipalDueDate := projectDueTime.Add(-time.Second * time.Duration(poolInfo.PoolInfo.WithdrawRequestWindowSeconds*2-10))
	currentTime := time.Now()
	withdrawStart := proposalDate.Sub(currentTime).Seconds()
	payPrincipalStart := payPrincipalDueDate.Sub(currentTime).Seconds()

	paymentDate := poolInfo.PoolInfo.LastPaymentTime.Add(time.Duration(poolInfo.PoolInfo.PayFreq) * time.Second)

	paymentDue := int(paymentDate.Sub(currentTime).Seconds())

	w := Window{
		paymentDue:          paymentDue,
		withdrawStartTime:   int(withdrawStart),
		payPartialStartTime: int(payPrincipalStart),
	}

	return w
}

func printMenu(poolInfo common.SPV) {
	window := getTimeWindow(poolInfo)

	color.Green.Printf("\n#####payment: %v(s)  submit proposal in %v(s) pay partial in %v(s) #########\n", window.paymentDue, window.withdrawStartTime, window.payPartialStartTime)
	color.Green.Println("1. deposit")
	color.Green.Println("2. withdraw")
	color.Green.Println("3. claim interest")
	color.Green.Println("4. transfer ownership")
	color.Green.Println("5. submit withdraw request")
	color.Green.Println("8. pay partial")
	color.Green.Println("6. dump the pool")
	color.Green.Println("7. all users withdraw all")
	color.Green.Println("0. exit")
	color.Green.Println("#####################################################################\n")
}

func runService(ctx context.Context, poolIndex string, totalInvestors int, wg *sync.WaitGroup, wNotify chan int, cancel context.CancelFunc, ch chan string) {
	defer wg.Done()

	requestWithdrawnums := -1
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(2 * time.Second):

			out, err := common.RunCommandWithOutput("joltify", "q", "spv", "query-pool", poolIndex, "--output", "json")
			if err != nil {
				color.Red.Printf("error in run command")
			}

			var poolInfo common.SPV
			err = json.Unmarshal([]byte(out), &poolInfo)
			if err != nil {
				color.Red.Printf("fail to unmarshal")
			}

			if requestWithdrawnums > 0 {
				printMenu(poolInfo)
				color.Gray.Println("we are in submit withdraw request phase, not allow to do other operations")
				continue
			}

			printMenu(poolInfo)
			fmt.Printf("input the choice:     ")
			var input string
			select {
			case input = <-ch:
			case <-time.After(inputTimeout):
			}
			if input == "" {
				continue
			}
			input = strings.Trim(input, "\n")
			choice, err := strconv.Atoi(input)
			if err != nil {
				logger.Error().Err(err).Msgf("error input")
				continue
			}
			switch choice {
			case 1, 2:
				color.Cyan.Printf("input the number of accounts wants to run ")

				var input string
				select {
				case input = <-ch:
				case <-time.After(inputTimeout):
				}

				if input == "" {
					continue
				}

				input = strings.Trim(input, "\n")

				numAccounts, err := strconv.Atoi(input)
				if err != nil {
					logger.Error().Err(err).Msgf("error input")
					continue
				}
				baseName := time.Now().Format("2006-01-02 15-04")
				fileName := fmt.Sprintf("%s-%s.xlsx", baseName, "before")
				poolb, depositorsb, _, err := common.DumpAll(poolIndex, fileName, needWrite)
				if err != nil {
					logger.Error().Err(err).Msgf("error dumnp all")
				}
				var actuallyDone map[int]int
				if choice == 1 {
					actuallyDone = withdrawOrDeposit(poolIndex, numAccounts, totalInvestors, false, false)
				} else {
					actuallyDone = withdrawOrDeposit(poolIndex, numAccounts, totalInvestors, false, true)
				}
				fileName = fmt.Sprintf("%s-%s.xlsx", baseName, "after")
				poola, depositorsa, _, err := common.DumpAll(poolIndex, fileName, needWrite)
				if err != nil {
					logger.Error().Err(err).Msgf("error dumnp all after")
				}

				pb, bool := new(big.Int).SetString(poolb.PoolInfo.UsableAmount.Amount, 10)
				if !bool {
					panic("should never fail")
				}

				pa, bool := new(big.Int).SetString(poola.PoolInfo.UsableAmount.Amount, 10)
				if !bool {
					panic("should never fail")
				}
				poolChainge := new(big.Int).Sub(pa, pb)
				if poolChainge.Cmp(big.NewInt(0)) == 0 {
					continue
				}

				total := big.NewInt(0)
				for _, el := range actuallyDone {
					total = total.Add(total, big.NewInt(int64(el)))
				}

				incorrect := false

				for i, el := range depositorsb {

					before, ok := new(big.Int).SetString(el.Depositor.WithdrawalAmount.Amount, 10)
					if !ok {
						continue
					}
					after, ok := new(big.Int).SetString(depositorsa[i].Depositor.WithdrawalAmount.Amount, 10)
					if !ok {
						continue
					}

					withdrawChange := new(big.Int).Sub(after, before)

					if withdrawChange.Cmp(big.NewInt(0)) == 0 {
						continue
					}

					var bb, ba *big.Int
					for _, el := range depositorsa[i].Balances {
						if el.Denom == "ausdc" {
							bb, ok = new(big.Int).SetString(el.Amount, 10)
							if !ok {
								panic("should never fail")
							}
						}
					}

					for _, el := range depositorsb[i].Balances {
						if el.Denom == "ausdc" {
							ba, ok = new(big.Int).SetString(el.Amount, 10)
							if !ok {
								panic("should never fail")
							}
						}
					}
					balanceChange := new(big.Int).Sub(bb, ba)
					color.Red.Printf("%v pool change: %v, withdraw change: %v, balance change: %v\n", i, poolChainge.String(), withdrawChange.String(), balanceChange.String())

					if new(big.Int).Abs(withdrawChange).Cmp(new(big.Int).Mul(base, big.NewInt(int64(actuallyDone[i+1])))) != 0 {
						color.Red.Printf("withdraw change not match\n")
						incorrect = true
					}

					if new(big.Int).Abs(poolChainge).Cmp(new(big.Int).Mul(base, total)) != 0 {
						color.Red.Printf("pool change not match\n")
						incorrect = true
					}
				}
				if !incorrect {
					tick := html.UnescapeString("&#" + "9989" + ";")
					color.Green.Printf("%v all checked correct\n", tick)
				} else {
					tick := html.UnescapeString("&#" + "10060" + ";")
					color.Red.Printf("%v some checked incorrect\n", tick)
				}

			case 3:
				color.Cyan.Printf("input the number of accounts wants to run claim interest ")

				var input string
				select {
				case input = <-ch:
				case <-time.After(inputTimeout):
				}

				if input == "" {
					continue
				}

				input = strings.Trim(input, "\n")

				numAccounts, err := strconv.Atoi(input)
				if err != nil {
					logger.Error().Err(err).Msgf("error input")
					continue
				}
				baseName := time.Now().Format("2006-01-02 15-04")
				fileName := fmt.Sprintf("%s-%s.xlsx", baseName, "before")
				_, depositorsb, _, err := common.DumpAll(poolIndex, fileName, needWrite)
				if err != nil {
					logger.Error().Err(err).Msgf("error dumnp all")
				}

				claimInterest(poolIndex, numAccounts, totalInvestors)
				fileName = fmt.Sprintf("%s-%s.xlsx", baseName, "after")
				_, depositorsa, _, err := common.DumpAll(poolIndex, fileName, needWrite)
				if err != nil {
					logger.Error().Err(err).Msgf("error dumnp all")
				}

				for i, el := range depositorsb {
					beforep, ok := new(big.Int).SetString(el.Depositor.PendingInterest.Amount, 10)
					if !ok {
						continue
					}
					beforeC, ok := new(big.Int).SetString(el.ClaimableInterestAmount.Amount, 10)
					if !ok {
						continue
					}
					totalBefore := new(big.Int).Add(beforep, beforeC)

					afterp, ok := new(big.Int).SetString(depositorsa[i].Depositor.PendingInterest.Amount, 10)
					if !ok {
						continue
					}
					afterC, ok := new(big.Int).SetString(depositorsa[i].ClaimableInterestAmount.Amount, 10)
					if !ok {
						continue
					}
					totalAfter := new(big.Int).Add(afterp, afterC)

					delta := new(big.Int).Sub(totalAfter, totalBefore)
					if delta.Cmp(big.NewInt(0)) == 0 {
						continue
					}

					var bb, ba *big.Int
					for _, el := range depositorsa[i].Balances {
						if el.Denom == "ausdc" {
							bb, ok = new(big.Int).SetString(el.Amount, 10)
							if !ok {
								panic("should never fail")
							}
						}
					}

					for _, el := range depositorsb[i].Balances {
						if el.Denom == "ausdc" {
							ba, ok = new(big.Int).SetString(el.Amount, 10)
							if !ok {
								panic("should never fail")
							}
						}
					}
					balanceChange := new(big.Int).Sub(bb, ba)
					var tick string
					if balanceChange.Cmp(new(big.Int).Mul(delta, big.NewInt(-1))) == 0 {
						tick = html.UnescapeString("&#" + "9989" + ";")
					} else {
						tick = html.UnescapeString("&#" + "10062" + ";")
					}
					color.Magenta.Printf("%v %v interest change: %v, balance change: %v\n", i, tick, delta.String(), balanceChange.String())
				}

			case 4:
				color.Cyan.Printf("input the number of accounts wants to run transfer ownership ")

				var input string
				select {
				case input = <-ch:
				case <-time.After(inputTimeout):
				}

				if input == "" {
					continue
				}

				input = strings.Trim(input, "\n")
				numAccounts, err := strconv.Atoi(input)
				if err != nil {
					logger.Error().Err(err).Msgf("error input")
					continue
				}
				baseName := time.Now().Format("2006-01-02 15-04")
				fileName := fmt.Sprintf("%s-%s.xlsx", baseName, "before")
				_, depositorb, _, err := common.DumpAll(poolIndex, fileName, needWrite)
				if err != nil {
					logger.Error().Err(err).Msgf("error dumnp all")
				}
				transferOwnership(poolIndex, numAccounts, totalInvestors)
				fileName = fmt.Sprintf("%s-%s.xlsx", baseName, "after")
				_, depositora, _, err := common.DumpAll(poolIndex, fileName, needWrite)
				if err != nil {
					logger.Error().Err(err).Msgf("error dumnp all")
				}
				totalTransfer := sdk.NewIntFromUint64(0)
				price := getprice()
				ratio := sdk.MustNewDecFromStr(price.Price.Price)
				for i, el := range depositorb {
					before := el.Depositor.DepositType
					after := depositora[i].Depositor.DepositType

					if before != after {
						locked := depositora[i].Depositor.LockedAmount.Amount
						lockedd, ok := sdk.NewIntFromString(locked)
						if !ok {
							panic("should not fail in convert string to digit")
						}
						lockedUsd := ratio.MulInt(lockedd).TruncateInt()
						totalTransfer = totalTransfer.Add(lockedUsd)
						color.Red.Printf("%v %v depositor status change: %v -> %v lockedUsd: %v\n", i, html.UnescapeString("&#"+"9989"+";"), before, after, lockedUsd)
					}
				}
				color.Cyan.Printf("total transfer: %v\n", totalTransfer.String())
				transferAmount = totalTransfer

			case 5:
				if requestWithdrawnums > 0 {
					logger.Info().Msgf("already have submitted the request to be processed")
					continue
				}
				color.Cyan.Printf("input the number of accounts wants to run submit withdraw request ")

				var input string
				select {
				case input = <-ch:
				case <-time.After(inputTimeout):
				}

				if input == "" {
					continue
				}

				input = strings.Trim(input, "\n")
				numAccounts, err := strconv.Atoi(input)
				if err != nil {
					logger.Error().Err(err).Msgf("error input")
					continue
				}
				requestWithdrawnums = numAccounts

			case 6:
				baseName := time.Now().Format("2006-01-02 15-04")
				fileName := fmt.Sprintf("%s-%s.xlsx", baseName, "before")
				common.DumpAll(poolIndex, fileName, true)

			case 7:
				withdrawOrDeposit(poolIndex, totalInvestors, totalInvestors, true, true)
			case 0:
				cancel()
				return

			}

		case option := <-wNotify:
			switch option {
			case common.WITHDRAW:
				if requestWithdrawnums == -1 {
					logger.Info().Msgf("no users withdraw request")
					continue
				}
				color.Yellow.Println("we are about to submit the withdraw request")
				baseName := time.Now().Format("2006-01-02 15-04")
				fileName := fmt.Sprintf("%s-%s.xlsx", baseName, "before")
				_, depositorb, _, err := common.DumpAll(poolIndex, fileName, needWrite)
				if err != nil {
					logger.Error().Err(err).Msgf("error dumnp all")
				}
				submitWithdraw(poolIndex, requestWithdrawnums, totalInvestors)
				fileName = fmt.Sprintf("%s-%s.xlsx", baseName, "after")
				_, depositora, _, err := common.DumpAll(poolIndex, fileName, needWrite)
				if err != nil {
					logger.Error().Err(err).Msgf("error dumnp all")
				}

				for i, el := range depositorb {
					before := el.Depositor.DepositType
					after := depositora[i].Depositor.DepositType
					if before != after {
						color.Red.Printf("%v %v depositor status change: %v -> %v\n", i, html.UnescapeString("&#"+"9989"+";"), before, after)
					}
				}

			case common.PAYPRINCIPAL:
				baseName := time.Now().Format("2006-01-02 15-04")
				fileName := fmt.Sprintf("%s-%s.xlsx", baseName, "before")
				common.DumpAll(poolIndex, fileName, needWrite)
				payPrincipalPartial(poolIndex)
				fileName = fmt.Sprintf("%s-%s.xlsx", baseName, "after")
				common.DumpAll(poolIndex, fileName, needWrite)
				requestWithdrawnums = -1
				color.Blue.Println("we have processed the withdraw request and pay the principal")

			}
		}
	}
}

// false means it exceeds the tolerance
func compareWithinError(a, b, e *big.Int) bool {
	delta := new(big.Int).Sub(new(big.Int).Abs(a), new(big.Int).Abs(b))
	if delta.CmpAbs(e) == 1 {
		return false
	}
	return true
}

func runWindowMonitor(ctx context.Context, poolIndex string, wg *sync.WaitGroup, wNotify chan int) {
	done1 := false

	var depositorsb []common.SPV
	var depositorsa []common.SPV
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(5 * time.Second):

			out, err := common.RunCommandWithOutput("joltify", "q", "spv", "query-pool", poolIndex, "--output", "json")
			if err != nil {
				color.Red.Printf("error in run command")
			}

			var poolInfo common.SPV
			err = json.Unmarshal([]byte(out), &poolInfo)
			if err != nil {
				color.Red.Printf("fail to unmarshal")
			}

			w := getTimeWindow(poolInfo)
			if w.withdrawStartTime <= 0 && w.withdrawStartTime > -10 && !done1 {
				color.Green.Printf("send withdraw notify\n")
				wNotify <- common.WITHDRAW
				color.Green.Println("done withdrawal notification")
				done1 = true

			}
			if w.payPartialStartTime <= 0 && w.payPartialStartTime > -10 {
				color.Green.Printf("send pay principal notify\n")

				if len(poolInfo.PoolInfo.WithdrawAccounts) != 0 {
					wNotify <- common.PAYPRINCIPAL
					done1 = false
				}
			}

			if w.paymentDue <= 6 {
				color.Yellow.Println("\nwe take dump before payment")
				done1 = true
				_, depositorsb, _, err = common.DumpAll(poolIndex, "before.xlsx", false)
				if err != nil {
					logger.Error().Err(err).Msgf("error dumnp all")
				}
			}
			if w.paymentDue > 100 {
				done1 = false
				color.Yellow.Println("we take dump after payment")
				_, depositorsa, _, err = common.DumpAll(poolIndex, "after.xlsx", false)
				if err != nil {
					logger.Error().Err(err).Msgf("error dumnp all")
				}
			}
			// if set1 && set2 {

			withdrawChangeMap := make(map[int]*big.Int)
			lockedChangeMap := make(map[int]*big.Int)

			totalWithdrawChange := big.NewInt(0)
			totalLockedChange := big.NewInt(0)
			totalTransferedLocked := big.NewInt(0)
			totalTransferedWithdrawed := big.NewInt(0)
			if depositorsa != nil && depositorsb != nil {
				for i, el := range depositorsb {
					before, ok := new(big.Int).SetString(el.Depositor.WithdrawalAmount.Amount, 10)
					if !ok {
						continue
					}
					after, ok := new(big.Int).SetString(depositorsa[i].Depositor.WithdrawalAmount.Amount, 10)
					if !ok {
						continue
					}

					withdrawChange := new(big.Int).Sub(after, before)

					if withdrawChange.Cmp(big.NewInt(0)) == 0 {
						continue
					}

					lb, ok := new(big.Int).SetString(el.Depositor.LockedAmount.Amount, 10)
					if !ok {
						continue
					}
					la, ok := new(big.Int).SetString(depositorsa[i].Depositor.LockedAmount.Amount, 10)
					if !ok {
						continue
					}

					lockChange := new(big.Int).Sub(la, lb)

					if withdrawChange.Cmp(big.NewInt(0)) != 0 {
						withdrawChangeMap[i] = withdrawChange
						totalWithdrawChange = totalWithdrawChange.Add(totalWithdrawChange, withdrawChange)
					}

					if lockChange.Cmp(big.NewInt(0)) != 0 {
						lockedChangeMap[i] = lockChange
						totalLockedChange = totalLockedChange.Add(totalLockedChange, lockChange)
					}
				}

				if len(withdrawChangeMap) != 0 {
					for _, v := range withdrawChangeMap {
						if v.Cmp(big.NewInt(0)) == 1 {
							totalTransferedWithdrawed = totalTransferedWithdrawed.Add(totalTransferedWithdrawed, v)
						}
					}
				}

				if len(lockedChangeMap) != 0 {
					for _, v := range lockedChangeMap {
						if v.Cmp(big.NewInt(0)) == 1 {
							totalTransferedLocked = totalTransferedLocked.Add(totalTransferedLocked, v)
						}
					}
				}

				depositorsb = nil
				depositorsa = nil
			}

			price := getprice()
			ratio := sdk.MustNewDecFromStr(price.Price.Price)
			lockedUsd := ratio.MulInt(sdk.NewIntFromBigInt(totalTransferedLocked)).TruncateInt()
			delta := lockedUsd.Sub(sdk.NewIntFromBigInt(totalTransferedWithdrawed)).Abs()

			if totalTransferedWithdrawed.Cmp(big.NewInt(0)) != 0 {
				color.Yellow.Printf(">>>>>transfer amount is %v and difference between total locked change and withdrawalable change is %v\n", totalTransferedWithdrawed.String(), delta.String())
				if !transferAmount.IsZero() {
					if transferAmount.Equal(sdk.NewIntFromBigInt(totalTransferedWithdrawed)) {
						color.Yellow.Printf("%v transfer amount is equal to total transfer request", correct)
					} else {
						color.Red.Printf("%v transfer amount is NOT equal to total transfer request", wrong)
					}
				}
			}

			if !compareWithinError(totalWithdrawChange, big.NewInt(0), big.NewInt(10)) || !compareWithinError(totalLockedChange, big.NewInt(0), big.NewInt(10)) {
				tick := html.UnescapeString("&#" + "10060" + ";")
				color.HiBlue.Printf("%v total withdraw change %v and total locked %v\n", tick, totalWithdrawChange.String(), totalLockedChange.String())
			}
		}
	}
}

func getprice() common.SPV {
	result, err := common.RunCommandWithOutput("joltify", "q", "pricefeed", "price", "aud:usd", "--output", "json")
	if err != nil {
		panic(err)
	}

	var price common.SPV
	json.Unmarshal([]byte(result), &price)
	return price
}

func main() {
	poolIndex := "0x43ce7e072884180e125328e727911ad83fcaba1cc487ece1ccc3e19376f51118"
	zlog.SetGlobalLevel(zlog.InfoLevel)

	err := godotenv.Load(".env")
	if err != nil {
		logger.Error().Err(err).Msgf("fail to load .env file")
		return
	}

	color.Cyan.Printf("need to run brand new chain? (y/n): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Trim(input, "\n")
	switch input {
	case "y":
		startChain()
		depositAndBorrow()
		payInterest(poolIndex)
	default:
	}

	allInvestors := os.Getenv("ALL_INVESTORS")
	n, err := strconv.Atoi(allInvestors)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	windowNotify := make(chan int, 1)
	wg.Add(3)

	ch := make(chan string)
	go func() {
		for {
			input, _ := reader.ReadString('\n')
			ch <- input
		}
	}()

	go runService(ctx, poolIndex, n, &wg, windowNotify, cancel, ch)
	go runWindowMonitor(ctx, poolIndex, &wg, windowNotify)
	go func() {
		defer wg.Done()
		select {
		case <-c:
			cancel()
			return
		case <-ctx.Done():
			return
		}
	}()
	wg.Wait()

	defer cancel()
}
