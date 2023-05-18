package main

import (
	"fmt"
	"html"
	"math/big"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/contrib/devnet/integrationtest/common"
)

var (
	wrong       = html.UnescapeString("&#" + "10060" + ";")
	correct     = html.UnescapeString("&#" + "9989" + ";")
	depositorsb []common.SPV
	depositorsa []common.SPV
)

func compareWithinError(a, b, e *big.Int) bool {
	delta := new(big.Int).Sub(new(big.Int).Abs(a), new(big.Int).Abs(b))
	return delta.CmpAbs(e) != 1
}

func triggerEvent(poolIndex string, wNotify chan int, display *outputData) error {
	w, poolInfo, err := common.GetWindow(poolIndex)
	if err != nil {
		return err
	}
	if w.WithdrawStartTime <= 0 && w.WithdrawStartTime > -10 {
		display.showOutput("send withdraw notify", BLUE)
		wNotify <- common.WITHDRAW
	}
	if w.PayPartialStartTime <= 0 && w.PayPartialStartTime > -10 {
		display.showOutput("send pay principal notify", BLUE)

		if len(poolInfo.PoolInfo.WithdrawAccounts) != 0 {
			wNotify <- common.PAYPRINCIPAL
		}
	}

	if w.PaymentDue <= 6 && len(depositorsb) == 0 {
		display.showOutput("we take dump before payment", YELLOW)
		_, depositorsb, _, err = common.DumpAll(poolIndex, "before.xlsx", false, logger)
		if err != nil {
			return fmt.Errorf("error dumnp all: %v", err)
		}
	}
	if w.PaymentDue > 100 && len(depositorsa) == 0 {
		display.showOutput("we take dump after payment", YELLOW)
		_, depositorsa, _, err = common.DumpAll(poolIndex, "after.xlsx", false, logger)
		if err != nil {
			return fmt.Errorf("error dumnp all: %v", err)
		}
	}

	withdrawChangeMap := make(map[int]*big.Int)
	lockedChangeMap := make(map[int]*big.Int)

	totalWithdrawChange := big.NewInt(0)
	totalLockedChange := big.NewInt(0)
	totalTransferedLocked := big.NewInt(0)
	totalTransferedWithdrawed := big.NewInt(0)
	if len(depositorsa) != 0 && len(depositorsb) != 0 {
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
	ratio := types.MustNewDecFromStr(price.Price.Price)
	lockedUsd := ratio.MulInt(types.NewIntFromBigInt(totalTransferedLocked)).TruncateInt()
	delta := lockedUsd.Sub(types.NewIntFromBigInt(totalTransferedWithdrawed)).Abs()

	if totalTransferedWithdrawed.Cmp(big.NewInt(0)) != 0 {
		mnsg := fmt.Sprintf("transfer amount is %v and difference between total locked change and withdrawalable change is %v\n", totalTransferedWithdrawed.String(), delta.String())
		display.showOutput(mnsg, RED)
		if !transferAmount.IsZero() {
			if transferAmount.Equal(types.NewIntFromBigInt(totalTransferedWithdrawed)) {
				msg := fmt.Sprintf("%v transfer amount is equal to total transfer request", correct)
				display.showOutput(msg, GREEN)
			} else {
				msg := fmt.Sprintf("%v transfer amount is NOT equal to total transfer request", wrong)
				display.showOutput(msg, RED)
			}
		}
	}

	if !compareWithinError(totalWithdrawChange, big.NewInt(0), big.NewInt(10)) || !compareWithinError(totalLockedChange, big.NewInt(0), big.NewInt(10)) {
		tick := html.UnescapeString("&#" + "10060" + ";")
		msg := fmt.Sprintf("%v total withdraw change %v and total locked %v\n", tick, totalWithdrawChange.String(), totalLockedChange.String())
		display.showOutput(msg, WHITE)
	}
	return nil
}
