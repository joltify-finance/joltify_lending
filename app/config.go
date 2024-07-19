package app

import (
	sdkmath "cosmossdk.io/math"

	cmtcfg "github.com/cometbft/cometbft/config"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// Bech32MainPrefix defines the Bech32 prefix for account addresses
	Bech32MainPrefix = "jolt"
	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key
	Bech32PrefixAccPub = Bech32MainPrefix + "pub"
	// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
	Bech32PrefixValAddr = Bech32MainPrefix + "val" + "oper"
	// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
	Bech32PrefixValPub = Bech32MainPrefix + "val" + "oper" + "pub"
	// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
	Bech32PrefixConsAddr = Bech32MainPrefix + "val" + "cons"
	// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
	Bech32PrefixConsPub = Bech32MainPrefix + "val" + "cons" + "pub"

	Bip44CoinType = 118 // we apply the atom coin type

	HumanCoinUnit = "jolt"
	BaseCoinUnit  = "ujolt"

	JoltExponent = 6
)

// SetSDKConfig configures the global config with joltify app specific parameters.
// It does not seal the config to allow modification in tests.
func SetSDKConfig() *sdk.Config {
	config := sdk.GetConfig()
	SetBech32AddressPrefixes(config)
	SetBip44CoinType(config)
	return config
}

// SetBech32AddressPrefixes sets the global prefix to be used when serializing addresses to bech32 strings.
func SetBech32AddressPrefixes(config *sdk.Config) {
	config.SetBech32PrefixForAccount(Bech32MainPrefix, Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)
}

func RegisterDenoms() {
	err := sdk.RegisterDenom(HumanCoinUnit, sdkmath.LegacyOneDec())
	if err != nil {
		panic(err)
	}
	err = sdk.RegisterDenom(BaseCoinUnit, sdkmath.LegacyNewDecWithPrec(1, JoltExponent))
	if err != nil {
		panic(err)
	}
}

// SetBip44CoinType sets the global coin type to be used in hierarchical deterministic wallets.
func SetBip44CoinType(config *sdk.Config) {
	config.SetCoinType(Bip44CoinType)
}

// initCometBFTConfig helps to override default CometBFT Config values.
// return cmtcfg.DefaultConfig if no custom configuration is required for the application.
func InitCometBFTConfig() *cmtcfg.Config {
	cfg := cmtcfg.DefaultConfig()

	// these values put a higher strain on node memory
	// cfg.P2P.MaxNumInboundPeers = 100
	// cfg.P2P.MaxNumOutboundPeers = 40

	return cfg
}

// initAppConfig helps to override default appConfig template and configs.
// return "", nil if no custom configuration is required for the application.
func InitAppConfig() (string, interface{}) {
	// The following code snippet is just for reference.
	type CustomAppConfig struct {
		serverconfig.Config `mapstructure:",squash"`
	}

	// Optionally allow the chain developer to overwrite the SDK's default
	// server config.
	srvCfg := serverconfig.DefaultConfig()
	// The SDK's default minimum gas price is set to "" (empty value) inside
	// app.toml. If left empty by validators, the node will halt on startup.
	// However, the chain developer can set a default app.toml value for their
	// validators here.
	//
	// In summary:
	// - if you leave srvCfg.MinGasPrices = "", all validators MUST tweak their
	//   own app.toml config,
	// - if you set srvCfg.MinGasPrices non-empty, validators CAN tweak their
	//   own app.toml to override, or use this default value.
	//
	// In tests, we set the min gas prices to 0.
	// srvCfg.MinGasPrices = "0stake"
	// srvCfg.BaseConfig.IAVLDisableFastNode = true // disable fastnode by default

	customAppConfig := CustomAppConfig{
		Config: *srvCfg,
	}

	customAppTemplate := serverconfig.DefaultConfigTemplate
	// Edit the default template file
	//
	// customAppTemplate := serverconfig.DefaultConfigTemplate + `
	// [wasm]
	// # This is the maximum sdk gas (wasm and storage) that we allow for any x/wasm "smart" queries
	// query_gas_limit = 300000
	// # This is the number of wasm vm instances we keep cached in memory for speed-up
	// # Warning: this is currently unstable and may lead to crashes, best to keep for 0 unless testing locally
	// lru_size = 0`

	return customAppTemplate, customAppConfig
}
