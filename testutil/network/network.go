package network

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	appconfig "github.com/joltify-finance/joltify_lending/app/config"
	configs2 "github.com/joltify-finance/joltify_lending/daemons/configs"
	"github.com/labstack/gommon/random"

	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/stretchr/testify/assert"

	pruningtypes "cosmossdk.io/store/pruning/types"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/joltify-finance/joltify_lending/app"
)

type (
	Network = network.Network
)

// New creates instance with fully configured cosmos network.
// Accepts optional config, that will be used in place of the DefaultConfig() if provided.
func New(t *testing.T, configs ...network.Config) *network.Network {
	if len(configs) > 1 {
		panic("at most one config should be provided")
	}
	var cfg network.Config
	if len(configs) == 0 {
		cfg = DefaultConfig()
	} else {
		cfg = configs[0]
	}

	p := t.TempDir()
	homeDir := path.Join(p, "node0/simd")
	fmt.Printf("we write to %v\n", homeDir)
	err := os.MkdirAll(path.Join(homeDir, "config"), 0755)
	assert.NoError(t, err)
	configs2.WriteDefaultPricefeedExchangeToml(homeDir) // must manually create config file.
	net, err := network.New(t, p, cfg)
	assert.NoError(t, err)
	t.Cleanup(net.Cleanup)
	return net
}

// DefaultConfig will initialize config for the network with custom application,
// genesis and single validator. All other parameters are inherited from cosmos-sdk/testutil/network.DefaultConfig
func DefaultConfig() network.Config {
	rd := random.New()
	randomChainName := rd.String(6, random.Alphabetic)
	randomChainID := strings.ToLower(randomChainName) + "localnet" + "_888-1"
	// randomChainID := "joltifydev_1729-1"

	encoding := appconfig.MakeEncodingConfig()
	app.ModuleBasics.RegisterInterfaces(encoding.InterfaceRegistry)

	net := network.Config{
		Codec:             encoding.Codec,
		TxConfig:          encoding.TxConfig,
		LegacyAmino:       encoding.Amino,
		InterfaceRegistry: encoding.InterfaceRegistry,
		AccountRetriever:  authtypes.AccountRetriever{},
		AppConstructor: func(val network.ValidatorI) servertypes.Application {
			localApp := app.NewApp(
				val.GetCtx().Logger, dbm.NewMemDB(), nil,
				true,
				simtestutil.NewAppOptionsWithFlagHome(val.GetCtx().Config.RootDir),
				baseapp.SetPruning(pruningtypes.NewPruningOptionsFromString(pruningtypes.PruningOptionDefault)),
				baseapp.SetMinGasPrices("0stake"),
				baseapp.SetChainID("joltifytest_888-1"),
			)
			encoding = localApp.EncodingConfig()
			return localApp
		},

		GenesisState:    app.ModuleBasics.DefaultGenesis(encoding.Codec),
		TimeoutCommit:   2 * time.Second,
		ChainID:         randomChainID,
		NumValidators:   1,
		BondDenom:       sdk.DefaultBondDenom,
		MinGasPrices:    fmt.Sprintf("0.000006%s", "ujolt"),
		AccountTokens:   sdk.TokensFromConsensusPower(1000, sdk.DefaultPowerReduction),
		StakingTokens:   sdk.TokensFromConsensusPower(500, sdk.DefaultPowerReduction),
		BondedTokens:    sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction),
		PruningStrategy: pruningtypes.PruningOptionNothing,
		CleanupDir:      true,
		SigningAlgo:     string(hd.Secp256k1Type),
		KeyringOptions:  []keyring.Option{},
	}
	return net
}
