package cmd

import (
	"os"
	"strings"

	tmcfg "github.com/cometbft/cometbft/config"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// EnvPrefix is the prefix environment variables must have to configure the app.
const EnvPrefix = "JOLT"

// NewRootCmd creates a new root command for the joltify blockchain.
func NewRootCmd() *cobra.Command {

	var (
		moduleBasicManager module.BasicManager
		clientCtx          client.Context
		err                error
	)
	app.SetSDKConfig().Seal()

	encodingConfig := app.MakeEncodingConfig()

	initClientCtx := client.Context{}.
		WithCodec(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(types.AccountRetriever{}).
		WithHomeDir(app.DefaultNodeHome).
		WithViper(EnvPrefix)

	rootCmd := &cobra.Command{
		Use:   "joltify",
		Short: "Daemon and CLI for the Joltify blockchain.",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			cmd.SetOut(cmd.OutOrStdout())
			cmd.SetErr(cmd.ErrOrStderr())

			initClientCtx, err = client.ReadPersistentCommandFlags(initClientCtx, cmd.Flags())
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

			customAppTemplate, customAppConfig := initAppConfig("ujolt")
			return server.InterceptConfigsPreRunHandler(cmd, customAppTemplate, customAppConfig, tmcfg.DefaultConfig())
		},
	}

	initRootCmd(rootCmd, clientCtx.TxConfig, moduleBasicManager)

	overwriteFlagDefaults(rootCmd, map[string]string{
		flags.FlagChainID:        strings.ReplaceAll(app.Name, "-", ""),
		flags.FlagKeyringBackend: "test",
	})

	return rootCmd
}

// addSubCmds registers all the sub commands used by joltify.
//func addSubCmds(rootCmd *cobra.Command, encodingConfig params.EncodingConfig, defaultNodeHome string) {
//	gentxModule := simapp.ModuleBasics[genutiltypes.ModuleName].(genutil.AppModuleBasic)
//	rootCmd.AddCommand(
//		ethermintclient.ValidateChainID(
//			genutilcli.InitCmd(app.ModuleBasics, defaultNodeHome),
//		),
//		genutilcli.CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, defaultNodeHome, gentxModule.GenTxValidator),
//		genutilcli.MigrateGenesisCmd(),
//		genutilcli.GenTxCmd(app.ModuleBasics, encodingConfig.TxConfig, banktypes.GenesisBalancesIterator{}, defaultNodeHome),
//		genutilcli.ValidateGenesisCmd(app.ModuleBasics),
//		AddGenesisAccountCmd(defaultNodeHome),
//		tmcli.NewCompletionCmd(rootCmd, true), // TODO add other shells, drop tmcli dependency, unhide?
//		// testnetCmd(app.ModuleBasics, banktypes.GenesisBalancesIterator{}), // TODO add
//		debug.Cmd(),
//		config.Cmd(),
//	)
//
//	ac := appCreator{
//		encodingConfig: encodingConfig,
//	}
//
//	// add keybase, auxiliary RPC, query, and tx child commands
//	rootCmd.AddCommand(
//		StatusCommand(),
//		newQueryCmd(),
//		newTxCmd(),
//		joltclient.KeyCommands(app.DefaultNodeHome),
//	)
//}

func overwriteFlagDefaults(c *cobra.Command, defaults map[string]string) {
	set := func(s *pflag.FlagSet, key, val string) {
		if f := s.Lookup(key); f != nil {
			f.DefValue = val
			_ = f.Value.Set(val)
		}
	}
	for key, val := range defaults {
		set(c.Flags(), key, val)
		set(c.PersistentFlags(), key, val)
	}
	for _, c := range c.Commands() {
		overwriteFlagDefaults(c, defaults)
	}
}
