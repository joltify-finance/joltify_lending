package cmd

import (
	appflags "github.com/joltify-finance/joltify_lending/app/flags"
	daemonflags "github.com/joltify-finance/joltify_lending/daemons/flags"
	"github.com/joltify-finance/joltify_lending/dydx_helper/indexer"
	clobflags "github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/flags"
	"github.com/spf13/cobra"
)

// GetOptionWithCustomStartCmd returns a root command option with custom start commands.
func GetOptionWithCustomStartCmd() *RootCmdOption {
	option := newRootCmdOption()
	f := func(cmd *cobra.Command) {
		// Add app flags.
		appflags.AddFlagsToCmd(cmd)

		// Add daemon flags.
		daemonflags.AddDaemonFlagsToCmd(cmd)

		// Add indexer flags.
		indexer.AddIndexerFlagsToCmd(cmd)

		// Add clob flags.
		clobflags.AddClobFlagsToCmd(cmd)
	}
	option.setCustomizeStartCmd(f)
	return option
}
