package main

import (
	sdkmath "cosmossdk.io/math"
	"fmt"
	"os"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/joltify-finance/joltify_lending/x/spv/keeper"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("use : cmd apy payfreq(seconds) userAmount(optional)\n")
		return
	}

	apy, err := sdk.NewDecFromStr(os.Args[1])
	if err != nil {
		// print the log of the err
		fmt.Printf("use : cmd apy payfreq(seconds) ")
		return
	}
	payFreq, err := strconv.Atoi(os.Args[2])
	if err != nil {
		// print the log of the err
		fmt.Printf("use : cmd apy payfreq(seconds) ")
		return
	}

	apyToPayFreq, err := keeper.CalculateInterestAmount(apy, payFreq)
	if err != nil {
		// print the log of the err
		log.Error("fail to calculate interest amount", "err", err)
		return
	}

	var a, interestToUser, interestToReserve sdkmath.Int
	if len(os.Args) == 4 {
		userAmount := os.Args[3]
		amount, ok := sdk.NewIntFromString(userAmount)
		if !ok {
			// print the log of the err
			fmt.Printf("use : cmd apy payfreq(seconds) userAmount(optional)")
			return
		}
		// 85% of the interest is paid to the user, 15% is to the pool
		a = apyToPayFreq.MulInt(amount).TruncateInt()
		interestToReserve = sdk.NewDecFromInt(a).Mul(sdk.MustNewDecFromStr("0.15")).TruncateInt()
		interestToUser = a.Sub(interestToReserve)
	}

	fmt.Printf("apy: %s, payFreq: %d, apyToPayFreq: %s, interest_to_user: %s interest_to_reserve %s, total %s\n", apy, payFreq, apyToPayFreq, interestToUser, interestToReserve, a)
}
