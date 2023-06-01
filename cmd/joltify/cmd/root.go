package cmd

import (
	"os"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/app/params"
	"github.com/spf13/cobra"
	tmcfg "github.com/tendermint/tendermint/config"
	tmcli "github.com/tendermint/tendermint/libs/cli"
)

// EnvPrefix is the prefix environment variables must have to configure the app.
const EnvPrefix = "JOLT"

// NewRootCmd creates a new root command for the joltify blockchain.
func NewRootCmd() *cobra.Command {
	app.SetSDKConfig().Seal()

	encodingConfig := app.MakeEncodingConfig()

	initClientCtx := client.Context{}.
		WithCodec(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(types.AccountRetriever{}).
		WithBroadcastMode(flags.BroadcastBlock).
		WithHomeDir(app.DefaultNodeHome).
		WithViper(EnvPrefix)

	rootCmd := &cobra.Command{
		Use:   "joltify",
		Short: "Daemon and CLI for the Joltify blockchain.",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			cmd.SetOut(cmd.OutOrStdout())
			cmd.SetErr(cmd.ErrOrStderr())

			initClientCtx, err := client.ReadPersistentCommandFlags(initClientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			initClientCtx, err = config.ReadFromClientConfig(initClientCtx)
			if err != nil {
				return err
			}

			if err = client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
				return err
			}
			srvCfg := serverconfig.DefaultConfig()
			srvCfg.MinGasPrices = "0ujolt"

			type CustomAppConfig struct {
				serverconfig.Config
			}
			customAppConfig := CustomAppConfig{
				Config: *srvCfg,
			}
			customAppTemplate := serverconfig.DefaultConfigTemplate
			return server.InterceptConfigsPreRunHandler(cmd, customAppTemplate, customAppConfig, tmcfg.DefaultConfig())
		},
	}

	addSubCmds(rootCmd, encodingConfig, app.DefaultNodeHome)

	return rootCmd
}

// addSubCmds registers all the sub commands used by joltify.
func addSubCmds(rootCmd *cobra.Command, encodingConfig params.EncodingConfig, defaultNodeHome string) {
	rootCmd.AddCommand(
		genutilcli.InitCmd(app.ModuleBasics, defaultNodeHome),
		genutilcli.CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, defaultNodeHome),
		genutilcli.MigrateGenesisCmd(),
		genutilcli.GenTxCmd(app.ModuleBasics, encodingConfig.TxConfig, banktypes.GenesisBalancesIterator{}, defaultNodeHome),
		genutilcli.ValidateGenesisCmd(app.ModuleBasics),
		AddGenesisAccountCmd(defaultNodeHome),
		tmcli.NewCompletionCmd(rootCmd, true), // TODO add other shells, drop tmcli dependency, unhide?
		// testnetCmd(app.ModuleBasics, banktypes.GenesisBalancesIterator{}), // TODO add
		debug.Cmd(),
		config.Cmd(),
	)

	ac := appCreator{
		encodingConfig: encodingConfig,
	}

	server.AddCommands(
		rootCmd,
		defaultNodeHome,
		ac.newApp,
		ac.appExport,
		func(cmd *cobra.Command) {
			ac.addStartCmdFlags(cmd)
		},
	)

	// add keybase, auxiliary RPC, query, and tx child commands
	rootCmd.AddCommand(
		StatusCommand(),
		newQueryCmd(),
		newTxCmd(),
		keys.Commands(app.DefaultNodeHome),
	)
}
