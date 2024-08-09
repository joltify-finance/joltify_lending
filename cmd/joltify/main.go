package main

import (
	"fmt"
	"os"

	"github.com/joltify-finance/joltify_lending/cmd/joltify/cmd"

	clienthelpers "cosmossdk.io/client/v2/helpers"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/joltify-finance/joltify_lending/app"
)

func main() {
	app.RegisterDenoms()
	app.SetSDKConfig()

	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, clienthelpers.EnvPrefix, app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
