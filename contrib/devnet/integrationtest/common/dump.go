package common

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	zlog "github.com/rs/zerolog"

	"github.com/rs/zerolog/log"
)

func dumpPool(poolIndex, fileName string, needWrite bool) (SPV, error) {
	out, err := RunCommandWithOutput("joltify", "q", "spv", "query-pool", poolIndex, "--output", "json")
	if err != nil {
		fmt.Printf("error(%v) is %v\n", err, out)
		return SPV{}, err
	}

	var poolInfo SPV
	err = json.Unmarshal([]byte(out), &poolInfo)
	_ = err

	// var data [][]string
	data := make([][]string, 2)
	data[0] = []string{"pool name", "pool index", "usable amount", "borrowed amount", "pool status", "project length", "pay freq", "last payment time"}
	data[1] = []string{
		poolInfo.PoolInfo.PoolName, poolInfo.PoolInfo.Index, poolInfo.PoolInfo.UsableAmount.Amount, poolInfo.PoolInfo.BorrowedAmount.Amount, poolInfo.PoolInfo.PoolStatus, poolInfo.PoolInfo.ProjectLength, strconv.Itoa(poolInfo.PoolInfo.PayFreq), poolInfo.PoolInfo.LastPaymentTime.String(),
	}
	if needWrite {
		WriteDataToExcel("pool_info", data, fileName)
	}
	return poolInfo, nil
}

func dumpInvestorsAndInterest(poolIndex, fileName string, needWrite bool, logger zlog.Logger) ([]SPV, error) {
	ret := os.Getenv("ALL_INVESTORS")
	initInvestors, err := strconv.Atoi(ret)
	if err != nil {
		fmt.Printf("incorrect ret")
		return nil, err
	}
	wg := sync.WaitGroup{}
	wg.Add(initInvestors)
	var errG error
	data := make([][]string, initInvestors+1)
	data[0] = []string{"depositor address", "withdrawal amount", "locked amount", "deposit type", "pending interest", "claimable interest", "balance"}
	locker := sync.RWMutex{}

	depositorsInterest := make([]SPV, initInvestors)
	for i := 1; i <= initInvestors; i++ {
		go func(index int) {
			defer wg.Done()
			keyName := fmt.Sprintf("key_%d", index)
			address, err := RunCommandWithOutput("joltify", "keys", "show", keyName, "--address")
			if err != nil {
				errG = err
				return
			}
			address = strings.Trim(address, "\n")
			out, err := RunCommandWithOutput("joltify", "q", "spv", "depositor", poolIndex, address, "--output", "json")
			if err != nil {
				logger.Error().Err(err).Msgf(">>> no deposit found for key %v at pool %v", index, poolIndex)
			}

			out2, err := RunCommandWithOutput("joltify", "q", "spv", "claimable-interest", address, poolIndex, "--output", "json")
			if err != nil {
				// this means the depositor cannot be found
				logger.Debug().Msgf(">>> no interest found for key %v at pool %v", index, poolIndex)
			}

			out3, err := RunCommandWithOutput("joltify", "q", "bank", "balances", address, "--output", "json")
			if err != nil {
				fmt.Printf("error to get the balance")
			}
			var balances SPV
			err = json.Unmarshal([]byte(out3), &balances)
			if err != nil {
				panic(err)
			}
			amountAusdc := "0"
			for _, coin := range balances.Balances {
				if coin.Denom == "ausdc" {
					amountAusdc = coin.Amount
				}
			}

			var depositor SPV
			err = json.Unmarshal([]byte(out), &depositor)
			_ = err // we ignore the error here as it is expected

			var interest SPV
			err = json.Unmarshal([]byte(out2), &interest)
			_ = err // we ignore the error here as it is expected
			locker.Lock()
			data[index] = []string{address, depositor.Depositor.WithdrawalAmount.Amount, depositor.Depositor.LockedAmount.Amount, depositor.Depositor.DepositType, depositor.Depositor.PendingInterest.Amount, interest.ClaimableInterestAmount.Amount, amountAusdc}
			depositorsInterest[index-1].Depositor = depositor.Depositor
			depositorsInterest[index-1].Balances = balances.Balances
			depositorsInterest[index-1].ClaimableInterestAmount = interest.ClaimableInterestAmount
			locker.Unlock()
		}(i)
	}
	wg.Wait()

	if needWrite {
		WriteDataToExcel("depositor_info", data, fileName)
	}
	return depositorsInterest, errG
}

func dumpBorrowNFT(poolIndex, fileName string, needWrite bool, logger zlog.Logger) ([]SPV, error) {
	out, err := RunCommandWithOutput("joltify", "q", "spv", "query-pool", poolIndex, "--output", "json")
	if err != nil {
		fmt.Printf("error(%v) is %v\n", err, out)
		return nil, err
	}

	var poolInfo SPV
	err = json.Unmarshal([]byte(out), &poolInfo)
	_ = err

	nfts := poolInfo.PoolInfo.PoolNFTIds

	data := make([][]string, len(nfts)+1)
	data[0] = []string{"nft-id", "borrowed amount", "borrowed time", "exchange ratio", "total interest paid counter", "accumulate interest", "interest paid", "delta interest"}
	nftsResult := make([]SPV, len(nfts))
	for i, el := range nfts {
		out, err := RunCommandWithOutput("joltify", "q", "nft", "class", el, "--output", "json")
		if err != nil {
			logger.Error().Err(err).Msgf("error is %v\n", out)
			continue
		}
		var nft SPV
		err = json.Unmarshal([]byte(out), &nft)
		_ = err
		borrow := nft.Class.Data.BorrowDetails[len(nft.Class.Data.BorrowDetails)-1]
		paidCounter := len(nft.Class.Data.Payments)
		counter := strconv.Itoa(paidCounter)

		acc, ok := new(big.Int).SetString(nft.Class.Data.AccInterest.Amount, 10)
		if !ok {
			panic("accInterest is not a number")
		}
		paid, ok := new(big.Int).SetString(nft.Class.Data.InterestPaid.Amount, 10)
		if !ok {
			panic("paid is not a number")
		}

		data[i+1] = []string{nft.Class.ID, borrow.BorrowedAmount.Amount, borrow.TimeStamp.String(), borrow.ExchangeRatio, counter, nft.Class.Data.AccInterest.Amount, nft.Class.Data.InterestPaid.Amount, new(big.Int).Sub(acc, paid).String()}
		nftsResult[i] = nft
	}
	if needWrite {
		WriteDataToExcel("borrow_info", data, fileName)
	}
	return nftsResult, nil
}

func DumpAll(poolIndex, fileName string, needWrite bool, logger zlog.Logger) (SPV, []SPV, []SPV, error) {
	if fileName == "" {
		fileName = time.Now().Format("2006-01-02 15-04") + ".xlsx"
	}
	var err error
	poolSPV, err := dumpPool(poolIndex, fileName, needWrite)
	if err != nil {
		log.Error().Err(err).Msg(">>>>error in dump pool")
		return SPV{}, nil, nil, err
	}
	depositors, err := dumpInvestorsAndInterest(poolIndex, fileName, needWrite, logger)
	if err != nil {
		logger.Error().Err(err).Msg(">>>>error in dump depositor n")
		return SPV{}, nil, nil, err
	}
	nftsSPV, err := dumpBorrowNFT(poolIndex, fileName, needWrite, logger)
	if err != nil {
		logger.Error().Err(err).Msg(">>>>error in dump dump")
		return SPV{}, nil, nil, err
	}
	return poolSPV, depositors, nftsSPV, nil
}
