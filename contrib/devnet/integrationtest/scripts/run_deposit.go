package main

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"sync"
	// rand2 "github.com/tendermint/tendermint/libs/rand"
)

// the random here is pre-determined order
func generateRandomIntegersWithSum(n int, sum int) []float64 {
	rand.New(rand.NewSource(100))
	mean := float64(sum) / float64(n) // mean of the normal distribution
	stdDev := math.Sqrt(mean)         // standard deviation of the normal distribution

	// generate random values with a normal distribution
	values := make([]float64, n)
	for i := range values {
		values[i] = rand.NormFloat64()*stdDev + mean
	}

	// adjust values to ensure their sum is equal to the given sum
	currentSum := 0.0
	for _, v := range values {
		currentSum += v
	}
	scaleFactor := float64(sum) / currentSum
	for i := range values {
		values[i] *= scaleFactor
	}
	return values
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

	values := generateRandomIntegersWithSum(investorsNum, totalAmount)
	fmt.Println("values: ", values)
	// return
	wg := sync.WaitGroup{}
	wg.Add(len(values))
	for i, v := range values {
		go func(index int, value float64) {
			defer wg.Done()
			var cmd *exec.Cmd
			if withdraw {
				cmd = exec.Command("./withdraw_principal.sh", poolIndex, "1000000000000000000000000ausdc", strconv.Itoa(offset+index+1))
			} else {
				valueStr := strconv.FormatFloat(value, 'f', 6, 64)
				// run the shell scripts
				cmd = exec.Command("./scripts/deposit.sh", poolIndex, valueStr, strconv.Itoa(offset+index+1))
			}

			// pipe the commands output to the applications
			// standard output
			var outb, errb bytes.Buffer
			cmd.Stdout = &outb
			cmd.Stderr = &errb

			// Run still runs the command and waits for completion
			// but the output is instantly piped to Stdout
			if err := cmd.Run(); err != nil {
				fmt.Printf("could not run command: %v", outb.String())
				return
			} else {
				fmt.Println(outb.String())
			}
		}(i, v)
	}
	wg.Wait()
	os.Exit(0)
}
