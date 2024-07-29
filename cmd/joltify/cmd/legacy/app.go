package cmd

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/types/module"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"

	"cosmossdk.io/log"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/app/params"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

const (
	flagMempoolEnableAuth    = "mempool.enable-authentication"
	flagMempoolAuthAddresses = "mempool.authorized-addresses"
	ChainID                  = "joltify_1729-1"
)

// appCreator holds functions used by the sdk server to control the joltify app.
// The methods implement types in cosmos-sdk/server/types
type appCreator struct {
	encodingConfig params.EncodingConfig
}

//func initRootCmd(
//	rootCmd *cobra.Command,
//	txConfig client.TxConfig,
//	basicManager module.BasicManager,
//) {
//	rootCmd.AddCommand(
//		genutilcli.InitCmd(basicManager, app.DefaultNodeHome),
//		debug.Cmd(),
//		confixcmd.ConfigCommand(),
//		pruning.Cmd(newApp, app.DefaultNodeHome),
//		snapshot.Cmd(newApp),
//	)
//
//	server.AddCommands(rootCmd, app.DefaultNodeHome, newApp, appExport, addModuleInitFlags)
//
//	// add keybase, auxiliary RPC, query, genesis, and tx child commands
//	rootCmd.AddCommand(
//		server.StatusCommand(),
//		genesisCommand(txConfig, basicManager),
//		queryCommand(),
//		txCommand(),
//		keys.Commands(),
//	)
//}

func addModuleInitFlags(startCmd *cobra.Command) {
	crisis.AddModuleInitFlags(startCmd)
}

// genesisCommand builds genesis-related `exampled genesis` command. Users may provide application specific commands as a parameter
func genesisCommand(txConfig client.TxConfig, basicManager module.BasicManager, cmds ...*cobra.Command) *cobra.Command {
	cmd := genutilcli.Commands(txConfig, basicManager, app.DefaultNodeHome)

	for _, subCmd := range cmds {
		cmd.AddCommand(subCmd)
	}
	return cmd
}

func queryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "query",
		Aliases:                    []string{"q"},
		Short:                      "Querying subcommands",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		rpc.QueryEventForTxCmd(),
		rpc.ValidatorCommand(),
		server.QueryBlockCmd(),
		authcmd.QueryTxsByEventsCmd(),
		server.QueryBlocksCmd(),
		authcmd.QueryTxCmd(),
		server.QueryBlockResultsCmd(),
	)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

func txCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "tx",
		Short:                      "Transactions subcommands",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetSignCommand(),
		authcmd.GetSignBatchCommand(),
		authcmd.GetMultiSignCommand(),
		authcmd.GetMultiSignBatchCmd(),
		authcmd.GetValidateSignaturesCommand(),
		flags.LineBreak,
		authcmd.GetBroadcastCommand(),
		authcmd.GetEncodeCommand(),
		authcmd.GetDecodeCommand(),
		authcmd.GetSimulateCmd(),
	)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

// newApp loads config from AppOptions and returns a new app.
func (ac appCreator) newApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	appOpts servertypes.AppOptions,
) servertypes.Application {
	//var cache sdk.MultiStorePersistentCache
	//if cast.ToBool(appOpts.Get(server.FlagInterBlockCache)) {
	//	cache = store.NewCommitKVStoreCacheManager()
	//}

	skipUpgradeHeights := make(map[int64]bool)
	for _, h := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}

	pruningOpts, err := server.GetPruningOptionsFromFlags(appOpts)
	if err != nil {
		panic(err)
	}

	homeDir := cast.ToString(appOpts.Get(flags.FlagHome))
	// snapshotDir := filepath.Join(homeDir, "data", "snapshots") // TODO can these directory names be imported from somewhere?
	// snapshotDB, err := dbm.NewDB("metadata", dbm.GoLevelDBBackend, snapshotDir)
	if err != nil {
		panic(err)
	}
	//snapshotStore, err := snapshots.NewStore(snapshotDB, snapshotDir)
	//if err != nil {
	//	panic(err)
	//}

	mempoolEnableAuth := cast.ToBool(appOpts.Get(flagMempoolEnableAuth))
	mempoolAuthAddresses, err := accAddressesFromBech32(
		cast.ToStringSlice(appOpts.Get(flagMempoolAuthAddresses))...,
	)
	// snapOpts := types.NewSnapshotOptions(cast.ToUint64(appOpts.Get(server.FlagStateSyncSnapshotInterval)), cast.ToUint32(appOpts.Get(server.FlagStateSyncSnapshotKeepRecent)))
	if err != nil {
		panic(fmt.Sprintf("could not get authorized address from config: %v", err))
	}

	return app.NewApp(
		logger, db, homeDir, traceStore, ac.encodingConfig,
		app.Options{
			SkipLoadLatest:        false,
			SkipUpgradeHeights:    skipUpgradeHeights,
			SkipGenesisInvariants: cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants)),
			InvariantCheckPeriod:  cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)),
			MempoolEnableAuth:     mempoolEnableAuth,
			MempoolAuthAddresses:  mempoolAuthAddresses,
		},
		0,
		appOpts,
		baseapp.SetPruning(pruningOpts),
		baseapp.SetMinGasPrices(strings.ReplaceAll(cast.ToString(appOpts.Get(server.FlagMinGasPrices)), ";", ",")),
		baseapp.SetHaltHeight(cast.ToUint64(appOpts.Get(server.FlagHaltHeight))),
		baseapp.SetHaltTime(cast.ToUint64(appOpts.Get(server.FlagHaltTime))),
		baseapp.SetMinRetainBlocks(cast.ToUint64(appOpts.Get(server.FlagMinRetainBlocks))), // TODO what is this?
		baseapp.SetTrace(cast.ToBool(appOpts.Get(server.FlagTrace))),
		baseapp.SetIndexEvents(cast.ToStringSlice(appOpts.Get(server.FlagIndexEvents))),
		baseapp.SetChainID(ChainID),
	)
}

// appExport writes out an app's state to json.
func (ac appCreator) appExport(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	height int64,
	forZeroHeight bool,
	jailAllowedAddrs []string,
	appOpts servertypes.AppOptions,
	modulesToExport []string,
) (servertypes.ExportedApp, error) {
	homePath, ok := appOpts.Get(flags.FlagHome).(string)
	if !ok || homePath == "" {
		return servertypes.ExportedApp{}, errors.New("application home not set")
	}

	options := app.DefaultOptions
	options.SkipLoadLatest = true
	options.InvariantCheckPeriod = cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod))

	var tempApp *app.App
	if height != -1 {
		tempApp = app.NewApp(logger, db, homePath, traceStore, ac.encodingConfig, options, uint(1), appOpts, baseapp.SetChainID(ChainID))
		if err := tempApp.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	} else {
		tempApp = app.NewApp(logger, db, homePath, traceStore, ac.encodingConfig, options, uint(1), appOpts, baseapp.SetChainID(ChainID))
	}

	return tempApp.ExportAppStateAndValidators(forZeroHeight, jailAllowedAddrs, modulesToExport)
}

// addStartCmdFlags adds flags to the server start command.
func (ac appCreator) addStartCmdFlags(startCmd *cobra.Command) {
	crisis.AddModuleInitFlags(startCmd)
}

// accAddressesFromBech32 converts a slice of bech32 encoded addresses into a slice of address types.
func accAddressesFromBech32(addresses ...string) ([]sdk.AccAddress, error) {
	var decodedAddresses []sdk.AccAddress
	for _, s := range addresses {
		a, err := sdk.AccAddressFromBech32(s)
		if err != nil {
			return nil, err
		}
		decodedAddresses = append(decodedAddresses, a)
	}
	return decodedAddresses, nil
}
