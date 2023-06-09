package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	zlog "github.com/rs/zerolog"

	"github.com/joho/godotenv"

	"github.com/joltify-finance/joltify_lending/contrib/devnet/integrationtest/common"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/rs/zerolog/log"
)

const (
	poolIndex = "0x43ce7e072884180e125328e727911ad83fcaba1cc487ece1ccc3e19376f51118"
	RED       = "red"
	GREEN     = "green"
	BLUE      = "blue"
	YELLOW    = "yellow"
	WHITE     = "white"
	denom     = "ausdc"
)

var (
	gbase             = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	needWrite         = false
	logger            = log.With().Logger()
	transferAmount    = sdk.NewInt(0)
	guiWithdrawAmount = 0
)

type outputData struct {
	msg       []string
	counter   int
	locker    *sync.Mutex
	outPannel *widgets.List
}

func newPaymentGauge() *widgets.Gauge {
	g := widgets.NewGauge()
	g.Title = "      Time for the next Payment"
	g.SetRect(45, 0, 145, 5)
	g.Percent = 50
	g.Label = "Time for the next payment"
	g.BarColor = ui.ColorGreen
	g.LabelStyle = ui.NewStyle(ui.ColorYellow)
	return g
}

func newSubmitWithdrawalGauge() *widgets.Gauge {
	g := widgets.NewGauge()
	g.Title = "      Time for submit withdrawal request"
	g.SetRect(45, 6, 145, 11)
	g.Percent = 50
	g.Label = "Gauge with custom highlighted label"
	g.BarColor = ui.ColorBlue
	g.LabelStyle = ui.NewStyle(ui.ColorYellow)
	return g
}

func newPartialPaymentGauge() *widgets.Gauge {
	g := widgets.NewGauge()
	g.Title = "      Time for partial payment"
	g.SetRect(45, 12, 145, 17)
	g.Percent = 50
	g.Label = "Gauge with custom highlighted label"
	g.BarColor = ui.ColorMagenta
	g.LabelStyle = ui.NewStyle(ui.ColorYellow)
	return g
}

func newOutput() *widgets.List {
	output := widgets.NewList()
	output.Title = "Output of the execution"
	output.TextStyle = ui.NewStyle(ui.ColorYellow)
	output.WrapText = false
	output.SetRect(2, 42, 145, 17)
	return output
}

func newList() *widgets.List {
	l := widgets.NewList()
	l.Title = "Enter number to select"
	l.Rows = []string{
		"[0] Exit",
		"[1] Deposit",
		"[2] Withdraw",
		"[3] Claim interest",
		"[4] Transfer ownership",
		"[5] Submit withdrawal",
		"[6] Dump pool",
		"[7] All user withdraw",
	}
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(2, 0, 40, 17)
	return l
}

func updateGauge(g1, g2, g3 *widgets.Gauge) error {
	w, poolInfo, err := common.GetWindow(poolIndex)
	if err != nil {
		logger.Error().Err(err).Msgf("fail to update gauge")
		return err
	}
	paymentLength := poolInfo.PoolInfo.PayFreq
	projectlength, err := strconv.Atoi(poolInfo.PoolInfo.ProjectLength)
	if err != nil {
		panic(err)
	}

	g1.Title = fmt.Sprintf("The payment is due in %v seconds (total %v)", w.PaymentDue, paymentLength)
	g1.Label = g1.Title
	if w.PaymentDue < 0 {
		g1.Percent = 0
	} else {
		g1.Percent = int(float32(w.PaymentDue) / float32(paymentLength) * 100)
	}

	g2.Title = fmt.Sprintf("The submit request in %v seconds (total %v)", w.WithdrawStartTime, projectlength)
	g2.Label = g2.Title
	if w.WithdrawStartTime < 0 {
		g2.Percent = 0
	} else {
		g2.Percent = int(float32(w.WithdrawStartTime) / float32(projectlength) * 100)
	}

	g3.Title = fmt.Sprintf("The partial payment in %v seconds (total %v)", w.PayPartialStartTime, projectlength)
	g3.Label = g3.Title
	if w.PayPartialStartTime < 0 {
		g3.Percent = 0
	} else {
		g3.Percent = int(float32(w.PayPartialStartTime) / float32(projectlength) * 100)
	}
	return nil
}

func processWindow(ctx context.Context, g1, g2, g3 *widgets.Gauge, gwNotify chan int, display *outputData) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second * 8):
			err := updateGauge(g1, g2, g3)
			if err != nil {
				fmt.Printf("err %v", err)
				continue
			}
			ui.Render(g1, g2, g3)
			err = triggerEvent(poolIndex, gwNotify, display)
			if err != nil {
				display.showOutput(err.Error(), RED)
			}

		}
	}
}

func (o *outputData) showOutput(msg string, color string) {
	o.locker.Lock()
	msg = fmt.Sprintf("[%v](fg:%v)", msg, color)
	if len(o.msg) >= o.counter {
		o.msg = o.msg[1:]
	}
	o.msg = append(o.msg, msg)
	o.outPannel.Lock()
	o.outPannel.Rows = o.msg
	o.outPannel.Unlock()

	ui.Render(o.outPannel)
	o.locker.Unlock()
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

func processEvent(cancel context.CancelFunc, wg *sync.WaitGroup, inputChain chan string, display *outputData, notify chan int) {
	allInvestors := os.Getenv("ALL_INVESTORS")
	totalInvestors, err := strconv.Atoi(allInvestors)
	if err != nil {
		panic(err)
	}

	defer func() {
		cancel()
		wg.Done()
	}()

	choice := -1
	for {
		select {
		case input := <-inputChain:
			input = strings.Trim(input, "\n")
			switch input {
			case "0", "q":
				return
			case "1", "2", "3", "4", "5":
				display.showOutput("input the number of accounts want to run", RED)
				ui.Render(display.outPannel)
				choice, err = strconv.Atoi(input)
				if err != nil {
					panic(err)
				}
				continue

			case "6":
				baseName := time.Now().Format("2006-01-02 15-04")
				fileName := fmt.Sprintf("%s-%s.xlsx", baseName, "before")
				common.DumpAll(poolIndex, fileName, true, logger)
				display.showOutput("finish dumping the status", GREEN)

			case "7":
				display.showOutput("start withdraw all", GREEN)
				actuallyDone := common.WithdrawOrDeposit(poolIndex, totalInvestors, totalInvestors, true, true, logger)
				var missed []int
				for k := 0; k < totalInvestors; k++ {
					_, ok := actuallyDone[k+1]
					if ok {
						continue
					}
					missed = append(missed, k+1)
				}

				msg := fmt.Sprintf("withdraw complete with failed users %v", missed)
				display.showOutput(msg, GREEN)

			default:
				if len(input) == 0 {
					continue
				}
				identifier := input[0]
				if identifier != 'c' {
					msg := fmt.Sprintf("invalid input %v", input)
					display.showOutput(msg, RED)

					ui.Render(display.outPannel)
					continue
				}
				numAccounts, err := strconv.Atoi(input[1:])
				if err != nil {
					display.showOutput("invalid input, please input the numnber", RED)
					ui.Render(display.outPannel)
					continue
				}
				switch choice {
				case 1, 2:
					var actuallyDone map[int]int
					display.showOutput("we start process deposit/withdraw", YELLOW)

					baseName := time.Now().Format("2006-01-02 15-04")
					fileName := fmt.Sprintf("%s-%s.xlsx", baseName, "before")
					poolb, depositorsb, _, err := common.DumpAll(poolIndex, fileName, needWrite, logger)
					if err != nil {
						emsg := fmt.Errorf("error dumnp all %w", err)
						display.showOutput(emsg.Error(), RED)
					}
					if choice == 1 {
						actuallyDone = common.WithdrawOrDeposit(poolIndex, numAccounts, totalInvestors, false, false, logger)
					} else {
						actuallyDone = common.WithdrawOrDeposit(poolIndex, numAccounts, totalInvestors, false, true, logger)
					}
					choice = -1
					msg := fmt.Sprintf("we actually done the withdraw with users %v", actuallyDone)
					display.showOutput(msg, GREEN)
					poola, depositorsa, _, err := common.DumpAll(poolIndex, fileName, needWrite, logger)
					if err != nil {
						msg := fmt.Errorf("error dumnp all after %w", err)
						display.showOutput(msg.Error(), RED)

					}

					pb, ok := new(big.Int).SetString(poolb.PoolInfo.UsableAmount.Amount, 10)
					if !ok {
						panic("should never fail")
					}

					pa, ok := new(big.Int).SetString(poola.PoolInfo.UsableAmount.Amount, 10)
					if !ok {
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
							if el.Denom == denom {
								bb, ok = new(big.Int).SetString(el.Amount, 10)
								if !ok {
									panic("should never fail")
								}
							}
						}

						for _, el := range depositorsb[i].Balances {
							if el.Denom == denom {
								ba, ok = new(big.Int).SetString(el.Amount, 10)
								if !ok {
									panic("should never fail")
								}
							}
						}
						balanceChange := new(big.Int).Sub(bb, ba)
						msg := fmt.Sprintf("%v pool change: %v, withdraw change: %v, balance change: %v\n", i, poolChainge.String(), withdrawChange.String(), balanceChange.String())

						display.showOutput(msg, GREEN)
						if new(big.Int).Abs(withdrawChange).Cmp(new(big.Int).Mul(gbase, big.NewInt(int64(actuallyDone[i+1])))) != 0 {
							display.showOutput("withdraw change not patch", RED)
							incorrect = true
						}

						if new(big.Int).Abs(poolChainge).Cmp(new(big.Int).Mul(gbase, total)) != 0 {
							display.showOutput("pool change not match", RED)
							incorrect = true
						}
					}
					if !incorrect {
						tick := html.UnescapeString("&#" + "9989" + ";")
						msg := fmt.Sprintf("%v all checked correct\n", tick)
						display.showOutput(msg, GREEN)
					} else {
						tick := html.UnescapeString("&#" + "10060" + ";")
						msg := fmt.Sprintf("%v some checked incorrect\n", tick)
						display.showOutput(msg, RED)
					}
				case 3:
					choice = -1
					var actuallyDone []int
					display.showOutput("we start process claim interest", YELLOW)

					baseName := time.Now().Format("2006-01-02 15-04")
					fileName := fmt.Sprintf("%s-%s.xlsx", baseName, "before")
					_, depositorsb, _, err := common.DumpAll(poolIndex, fileName, needWrite, logger)
					if err != nil {
						logger.Error().Err(err).Msgf("error dumnp all")
					}

					actuallyDone = claimInterest(poolIndex, numAccounts, totalInvestors)
					msg := fmt.Sprintf("we actually done the withdraw interest with users %v", actuallyDone)
					display.showOutput(msg, GREEN)
					fileName = fmt.Sprintf("%s-%s.xlsx", baseName, "after")
					_, depositorsa, _, err := common.DumpAll(poolIndex, fileName, needWrite, logger)
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
						msg := fmt.Sprintf("%v %v interest change: %v, balance change: %v\n", i, tick, delta.String(), balanceChange.String())
						display.showOutput(msg, GREEN)
					}

				case 4:
					choice = -1
					var actuallyDone []int
					display.showOutput("we start process transfer ownership", YELLOW)

					baseName := time.Now().Format("2006-01-02 15-04")
					fileName := fmt.Sprintf("%s-%s.xlsx", baseName, "before")
					_, depositorb, _, err := common.DumpAll(poolIndex, fileName, needWrite, logger)
					if err != nil {
						logger.Error().Err(err).Msgf("error dumnp all")
					}
					actuallyDone = transferOwnership(poolIndex, numAccounts, totalInvestors)

					msg := fmt.Sprintf("we actually done the transfer with users %v", actuallyDone)
					display.showOutput(msg, GREEN)

					fileName = fmt.Sprintf("%s-%s.xlsx", baseName, "after")
					_, depositora, _, err := common.DumpAll(poolIndex, fileName, needWrite, logger)
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
							msg := fmt.Sprintf("%v %v depositor status change: %v -> %v lockedUsd: %v\n", i, html.UnescapeString("&#"+"9989"+";"), before, after, lockedUsd)
							display.showOutput(msg, YELLOW)
						}
					}
					msg = fmt.Sprintf("total transfer: %v\n", totalTransfer.String())
					display.showOutput(msg, GREEN)
					transferAmount = totalTransfer

				case 5:
					choice = -1
					guiWithdrawAmount = numAccounts
					msg := fmt.Sprintf("we have submitted the withdraw request for %v users", guiWithdrawAmount)
					display.showOutput(msg, YELLOW)

				}

			}
			ui.Render(display.outPannel)

		case option := <-notify:
			switch option {
			case common.WITHDRAW:
				if guiWithdrawAmount == -1 {
					display.showOutput("no users withdraw request", WHITE)
					continue
				}
				display.showOutput("we are about to submit the withdraw request", YELLOW)
				baseName := time.Now().Format("2006-01-02 15-04")
				fileName := fmt.Sprintf("%s-%s.xlsx", baseName, "before")
				_, depositorb, _, err := common.DumpAll(poolIndex, fileName, needWrite, logger)
				if err != nil {
					logger.Error().Err(err).Msgf("error dumnp all")
				}
				actualDone := submitWithdraw(poolIndex, guiWithdrawAmount, totalInvestors)
				msg := fmt.Sprintf("we actually done the withdraw with users %v", actualDone)
				display.showOutput(msg, GREEN)
				fileName = fmt.Sprintf("%s-%s.xlsx", baseName, "after")
				_, depositora, _, err := common.DumpAll(poolIndex, fileName, needWrite, logger)
				if err != nil {
					logger.Error().Err(err).Msgf("error dumnp all")
				}

				for i, el := range depositorb {
					before := el.Depositor.DepositType
					after := depositora[i].Depositor.DepositType
					if before != after {
						display.showOutput(fmt.Sprintf("%v %v depositor status change: %v -> %v\n", i, html.UnescapeString("&#"+"9989"+";"), before, after), RED)
					}
				}

			case common.PAYPRINCIPAL:
				baseName := time.Now().Format("2006-01-02 15-04")
				fileName := fmt.Sprintf("%s-%s.xlsx", baseName, "before")
				common.DumpAll(poolIndex, fileName, needWrite, logger)
				payPrincipalPartial(poolIndex)
				fileName = fmt.Sprintf("%s-%s.xlsx", baseName, "after")
				common.DumpAll(poolIndex, fileName, needWrite, logger)
				guiWithdrawAmount = -1
				display.showOutput("we have processed the withdraw request and pay the principal", BLUE)

			}
		}
	}
}

func main() {
	zlog.SetGlobalLevel(zlog.InfoLevel)

	file, err := os.OpenFile(
		"gui.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0o664,
	)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	logger = zlog.New(file).With().Timestamp().Logger()

	fmt.Printf("Do you want to start the chain? (y/n):")
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

	if err := ui.Init(); err != nil {
		logger.Error().Err(err).Msgf("failed to initialize termui: %v", err)
		return
	}
	defer ui.Close()
	l := newList()
	g1 := newPaymentGauge()
	g2 := newSubmitWithdrawalGauge()
	g3 := newPartialPaymentGauge()
	output := newOutput()

	display := &outputData{
		msg:       make([]string, 0, 15),
		counter:   15,
		locker:    &sync.Mutex{},
		outPannel: output,
	}

	err = updateGauge(g1, g2, g3)
	if err != nil {
		fmt.Printf(">>>%v\n", err)
		return
	}
	ui.Render(l, output, g1, g2, g3)

	uiEvents := ui.PollEvents()

	windowNotify := make(chan int, 1)
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		processWindow(ctx, g1, g2, g3, windowNotify, display)
	}()

	err = godotenv.Load(".env")
	if err != nil {
		fmt.Printf("fail to load .env file")
		return
	}

	inputChan := make(chan string, 100)
	var cache bytes.Buffer
	// we listen to the keyboard input
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case e := <-uiEvents:

				if e.Type != ui.KeyboardEvent {
					continue
				}

				if e.ID == "<Enter>" {
					cache.WriteString("\n")
				} else {
					cache.WriteString(e.ID)
				}
				if e.ID == "<Enter>" {
					o, err := cache.ReadString('\n')
					if err != nil {
						panic(err)
					}
					inputChan <- o
				}
			}
		}
	}()

	wg.Add(1)
	go processEvent(cancel, &wg, inputChan, display, windowNotify)
	wg.Wait()
}
