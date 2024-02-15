package main

import (
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/fatih/color"

	"github.com/xuri/excelize/v2"
)

const (
	WITHDRAW = 1
	LOCKED   = 2
	PENDING  = 4
	INTEREST = 5
	BALANCE  = 6
)

func readExpectedBorrowFromNFTs() (*big.Int, *big.Int, error) {
	f, err := excelize.OpenFile(os.Args[1])
	if err != nil {
		fmt.Printf(">>>>open file error %v\n", err)
		return nil, nil, err
	}
	data, err := f.GetRows("borrow_info")
	if err != nil {
		fmt.Printf(">>>>get col error %v\n", err)
		return nil, nil, err
	}

	totalBorrowed := big.NewInt(0)
	accInterest := big.NewInt(0)

	for row := 1; row < len(data); row++ {
		borrowed, ok := new(big.Int).SetString(data[row][1], 10)
		if !ok {
			return nil, nil, errors.New("fail to convert borrowed amount")
		}

		interest, ok := new(big.Int).SetString(data[row][5], 10)
		if !ok {
			return nil, nil, errors.New("fail to convert interest amount")
		}

		totalBorrowed.Add(totalBorrowed, borrowed)
		accInterest.Add(accInterest, interest)
	}

	return totalBorrowed, accInterest, nil
}

func readExpectedBorrowFromPoolInfo() (*big.Int, *big.Int, error) {
	f, err := excelize.OpenFile(os.Args[1])
	if err != nil {
		fmt.Printf(">>>>open file error %v\n", err)
		return nil, nil, err
	}
	data, err := f.GetRows("pool_info")
	if err != nil {
		fmt.Printf(">>>>get col error %v\n", err)
		return nil, nil, err
	}

	usableAmount, ok := new(big.Int).SetString(data[1][2], 10)
	if !ok {
		return nil, nil, errors.New("fail to convert usable amount")
	}

	borrowedAmount, ok := new(big.Int).SetString(data[1][3], 10)
	if !ok {
		return nil, nil, errors.New("fail to convert borrowed amount")
	}

	return usableAmount, borrowedAmount, nil
}

func main() {
	filename := os.Args[1]
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Printf(">>>>open file error %v\n", err)
		return
	}
	data, err := f.GetRows("depositor_info")
	if err != nil {
		fmt.Printf(">>>>get col error %v\n", err)
		return
	}

	totalw := big.NewInt(0)
	totall := big.NewInt(0)
	totalp := big.NewInt(0)
	totali := big.NewInt(0)
	totalb := big.NewInt(0)

	// fmt.Printf(">>>>>data is %v\n\n", data[1][WITHDRAW])
	// return

	// a := data[1:]
	// fmt.Printf(">>>>%v----\n", a[1][WITHDRAW])
	// return
	numData := data[1:]
	// fmt.Printf(">>>>>numData is %v\n", numData[0][1])
	// return
	for row := 0; row < len(numData); row++ {
		w, ok := new(big.Int).SetString(numData[row][WITHDRAW], 10)
		if !ok {
			continue
		}

		l, ok := new(big.Int).SetString(numData[row][LOCKED], 10)
		if !ok {
			continue
		}

		p, ok := new(big.Int).SetString(numData[row][PENDING], 10)
		if !ok {
			continue
		}
		i, ok := new(big.Int).SetString(numData[row][INTEREST], 10)
		if !ok {
			continue
		}
		b, ok := new(big.Int).SetString(numData[row][BALANCE], 10)
		if !ok {
			continue
		}

		totalw.Add(totalw, w)
		totall.Add(totall, l)
		totalp.Add(totalp, p)
		totali.Add(totali, i)
		totalb.Add(totalb, b)
	}

	usable, borrowed, err := readExpectedBorrowFromPoolInfo()
	if err != nil {
		fmt.Printf(">>>>>read expected borrow error %v\n", err)
		return
	}

	_ = usable
	_ = borrowed

	totalBorrowedNft, accInterest, err := readExpectedBorrowFromNFTs()
	if err != nil {
		fmt.Printf(">>>>>read expected borrow error %v\n", err)
		return
	}

	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	gree := color.New(color.FgGreen).SprintFunc()

	borrowedEqual := borrowed.Cmp(totalBorrowedNft)
	if borrowedEqual != 0 {
		fmt.Printf("borrowed Amount is qual to pool: %v\n", red("FALSE"))
	} else {
		fmt.Printf("borrowed Amount is qual to pool: %v\n", gree("TRUE"))
	}

	fmt.Printf("withdrawal total is %v\n", yellow(totalw.String()))
	fmt.Printf("locked total is %v\n", yellow(totall.String()))
	fmt.Printf("pending total is %v\n", yellow(totalp.String()))
	fmt.Printf("interest total is %v\n", yellow(totali.String()))
	fmt.Printf("balance total is %v\n", yellow(totalb.String()))

	if totalw.Cmp(usable) != 0 {
		fmt.Printf("withdrawal total is equal to usable: %v\n", red("FALSE"))
	} else {
		fmt.Printf("withdrawal total is equal to usable: %v\n", gree("TRUE"))
	}

	if totall.Cmp(borrowed) != 0 {
		fmt.Printf("locked total is equal to borrowed: %v\n", red("FALSE"))
	} else {
		fmt.Printf("locked total is equal to borrowed: %v\n", gree("TRUE"))
	}

	if totali.Cmp(accInterest) != 0 {
		fmt.Printf("interest total is equal to accInterest: %v; DIFF:%v\n", red("FALSE"), totali.Sub(totali, accInterest))
	} else {
		fmt.Printf("interest total is equal to accInterest: %v\n", gree("TRUE"))
	}
}
