package main

import (
	"fmt"
	"os"

	clienthelpers "cosmossdk.io/client/v2/helpers"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/cmd/joltify/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, clienthelpers.EnvPrefix, app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}

	// rootCmd := cmd.NewRootCmd()

	//if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
	//	switch e := err.(type) {
	//	case server.ErrorCode:
	//		os.Exit(e.Code)
	//
	//	default:
	//		os.Exit(1)
	//	}
	//}
}
