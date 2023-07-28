package cmd

import (
	"fmt"
	"os"

	"github.com/tendermint/tendermint/store"

	cmtjson "github.com/tendermint/tendermint/libs/json"

	"github.com/tendermint/tendermint/node"
	sm "github.com/tendermint/tendermint/state"

	"github.com/evmos/ethermint/crypto/hd"

	"github.com/cosmos/cosmos-sdk/client/flags"
	ethermintclient "github.com/evmos/ethermint/client"
	ethermintserver "github.com/evmos/ethermint/server"
	servercfg "github.com/evmos/ethermint/server/config"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/app/params"
	joltclient "github.com/joltify-finance/joltify_lending/client"
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
		WithKeyringOptions(hd.EthSecp256k1Option()).
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

			customAppTemplate, customAppConfig := servercfg.AppConfig("ujolt")
			return server.InterceptConfigsPreRunHandler(cmd, customAppTemplate, customAppConfig, tmcfg.DefaultConfig())
		},
	}

	addSubCmds(rootCmd, encodingConfig, app.DefaultNodeHome)

	return rootCmd
}

// addSubCmds registers all the sub commands used by joltify.
func addSubCmds(rootCmd *cobra.Command, encodingConfig params.EncodingConfig, defaultNodeHome string) {
	rootCmd.AddCommand(
		ethermintclient.ValidateChainID(
			genutilcli.InitCmd(app.ModuleBasics, defaultNodeHome),
		),
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

	//config, err := servercfg.GetConfig(serverCtx.Viper)
	//if err != nil {
	//	panic(err)
	//}
	//

	dbProvider := node.DefaultDBProvider
	cfg := tmcfg.DefaultConfig()
	cfg.DBBackend = "goleveldb"
	cfg.DBPath = "/home/yb/.joltify/data"
	cfg.RootDir = "data"
	stateDB, err := dbProvider(&node.DBContext{ID: "state", Config: cfg})
	if err != nil {
		panic(err)
	}

	blockStoreDB, err := dbProvider(&node.DBContext{"blockstore", cfg})
	if err != nil {
		return
	}

	blockStore := store.NewBlockStore(blockStoreDB)
	block := blockStore.LoadBlock(50)
	// blockStore.SaveBlock(b, b.Evidence)
	_ = block.ChainID

	////
	genesisDocKey := []byte("genesisDoc")
	state, genDoc, err := node.LoadStateFromDBOrGenesisDocProvider(stateDB, nil)
	if err != nil {
		panic(err)
	}

	stateStore := sm.NewStore(stateDB, sm.StoreOptions{
		DiscardABCIResponses: false,
	})

	fmt.Printf(">>>>>>>genIDDIIIIII>>>>>>>>%v\n", genDoc.ChainID)
	fmt.Printf(">>>>>>>state ID>>>>>>>>%v\n", state.ChainID)
	genDoc.ChainID = "joltify-133"

	// panics if failed to marshal the given genesis document
	b, err := cmtjson.Marshal(genDoc)
	if err != nil {
		panic(err)
	}
	if err := stateDB.SetSync(genesisDocKey, b); err != nil {
		panic(err)
	}

	state.ChainID = genDoc.ChainID
	fmt.Printf(">>>>>>>>>>>>>%v\n", state.ChainID)
	fmt.Printf(">>>>>>>>>last height >>%v\n", state.LastBlockHeight)
	state.
		// state.LastBlockHeight = state.LastBlockHeight + 1
		fmt.Printf(">>>>>>>>>>%v\n", state)
	err = stateStore.Save(state)
	if err != nil {
		panic(err)
	}

	err = stateDB.Close()
	if err != nil {
		panic(err)
	}

	err = blockStoreDB.Close()
	if err != nil {
		panic(err)
	}

	//
	//
	//
	//
	//
	//stateDB, err = dbProvider(&DBContext{"state", config})
	//
	//
	//
	//state, genDoc, err := LoadStateFromDBOrGenesisDocProvider(stateDB, genesisDocProvider)
	//if err != nil {
	//	return nil, err
	//}
	//
	//

	ethermintserver.AddCommands(
		rootCmd,
		ethermintserver.NewDefaultStartOptions(
			ac.newApp,
			app.DefaultNodeHome,
		),
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
		joltclient.KeyCommands(app.DefaultNodeHome),
	)
}
