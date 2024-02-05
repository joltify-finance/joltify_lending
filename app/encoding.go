package app

import (
	"github.com/cosmos/cosmos-sdk/types/tx"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	enccodec "github.com/evmos/ethermint/encoding/codec"
	"github.com/joltify-finance/joltify_lending/app/params"
)

// MakeEncodingConfig creates an EncodingConfig and registers the app's types on it.
func MakeEncodingConfig() params.EncodingConfig {
	encodingConfig := params.MakeEncodingConfig()
	enccodec.RegisterLegacyAminoCodec(encodingConfig.Amino)
	enccodec.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	encodingConfig.InterfaceRegistry.RegisterImplementations(
		(*tx.TxExtensionOptionI)(nil),
		&authztypes.MsgGrant{},
		&authztypes.MsgRevoke{},
		&authztypes.MsgExec{},
	)
	return encodingConfig
}
