package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

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

func main() {
	poolIndex := os.Args[1]
	investorsNum, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(-1)
		return
	}

	offset, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(-1)
		return
	}

	totalAmount, err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(-1)
		return
	}

	withdraw, err := strconv.ParseBool(os.Args[5])
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(-1)
		return
	}

	values, _ := generateRandomIntegersWithSum(investorsNum, totalAmount)
	wg := sync.WaitGroup{}
	wg.Add(len(values))
	for i, v := range values {
		go func(index int, value int) {
			defer wg.Done()
			var cmd *exec.Cmd
			if withdraw {
				cmd = exec.Command("./withdraw_after_close.sh", poolIndex, "10ausdc", strconv.Itoa(offset+index+1))
			} else {
				// run the shell scripts
				cmd = exec.Command("./deposit.sh", poolIndex, strconv.Itoa(value), strconv.Itoa(offset+index+1))
			}

			// pipe the commands output to the applications
			// standard output
			var outb, errb bytes.Buffer
			cmd.Stdout = &outb
			cmd.Stderr = &errb

			// Run still runs the command and waits for completion
			// but the output is instantly piped to Stdout
			if err := cmd.Run(); err != nil {
				fmt.Println("could not run command: ", err)
				fmt.Println(errb.String())
				return
			} else {
				fmt.Println(outb.String())
			}
		}(i, v)
	}
	wg.Wait()
	os.Exit(0)
}
