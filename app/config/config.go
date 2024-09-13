package config

import (
	"fmt"

	"github.com/joltify-finance/joltify_lending/dydx_helper/module"

	"cosmossdk.io/x/tx/signing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	cosproto "github.com/cosmos/gogoproto/proto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	AccountAddressPrefix = "jolt"
	// Bech32MainPrefix defines the Bech32 prefix for an account's address
	Bech32MainPrefix = AccountAddressPrefix
	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
	Bech32PrefixAccAddr = Bech32MainPrefix
	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key
	Bech32PrefixAccPub = Bech32MainPrefix + sdk.PrefixPublic
	// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
	Bech32PrefixValAddr = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixOperator
	// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
	Bech32PrefixValPub = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixOperator + sdk.PrefixPublic
	// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
	Bech32PrefixConsAddr = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixConsensus
	// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
	Bech32PrefixConsPub = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixConsensus + sdk.PrefixPublic
)

// SetupConfig sets up and seals the config.
// Note that importing and invoking this function also calls the `init` function in this package,
// which sets the address prefixes.
func SetupConfig() {
	config := sdk.GetConfig()
	config.Seal()
}

func init() {
	// This package does not contain the `app/config` package in its import chain, and therefore needs to call
	// SetAddressPrefixes() explicitly in order to set the `dydx` address prefixes.
	SetAddressPrefixes()
}

// SetAddressPrefixes sets the global prefixes to be used when serializing addresses and public keys to Bech32 strings.
func SetAddressPrefixes() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)
}

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Codec             codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

// TODO(CORE-846): Consider having app injected messages return an error instead of empty signers list.
func noSigners(_ proto.Message) ([][]byte, error) {
	return [][]byte{}, nil
}

func getLegacyMsgSignerFn(path []string) func(msg proto.Message) ([][]byte, error) {
	if len(path) == 0 {
		panic("path is expected to contain at least one value.")
	}

	return func(msg proto.Message) ([][]byte, error) {
		m := msg.ProtoReflect()
		for _, p := range path[:len(path)-1] {
			fieldDesc := m.Descriptor().Fields().ByName(protoreflect.Name(p))
			if fieldDesc.Kind() != protoreflect.MessageKind {
				return nil, fmt.Errorf("Expected for field %s to be Message type in path %+v for msg %+v.", p, path, msg)
			}
			v := m.Get(fieldDesc)
			if !v.IsValid() {
				return nil, fmt.Errorf("Expected for field %s to be populated in path %+v for msg %+v.", p, path, msg)
			}
			m = v.Message()
		}

		fieldDesc := m.Descriptor().Fields().ByName(protoreflect.Name(path[len(path)-1]))
		if fieldDesc.Kind() != protoreflect.StringKind {
			return nil, fmt.Errorf(
				"Expected for final field %s to be String type in path %+v for msg %+v.",
				path[len(path)-1],
				path,
				msg,
			)
		}
		signer, err := sdk.AccAddressFromBech32(m.Get(fieldDesc).String())
		if err != nil {
			return nil, err
		}
		return [][]byte{signer}, nil
	}
}

// MakeEncodingConfig creates a new EncodingConfig.
func MakeEncodingConfig() EncodingConfig {
	// TODO(CORE-840): cosmos.msg.v1.signer annotation doesn't supported nested messages beyond a depth of 1
	// which requires any message where the authority is nested further to implement its own accessor. Once
	// https://github.com/cosmos/cosmos-sdk/issues/18722 is fixed, replace this with the cosmos.msg.v1.signing
	// annotation on the protos.

	customerSigner := make(map[protoreflect.FullName]signing.GetSignersFunc)

	customerSigner["dydxprotocol.bridge.MsgAcknowledgeBridges"] = noSigners
	customerSigner["dydxprotocol.clob.MsgProposedOperations"] = noSigners
	customerSigner["dydxprotocol.perpetuals.MsgAddPremiumVotes"] = noSigners
	customerSigner["dydxprotocol.prices.MsgUpdateMarketPrices"] = noSigners

	customerSigner["dydxprotocol.clob.MsgBatchCancel"] = getLegacyMsgSignerFn(
		[]string{"subaccount_id", "owner"})
	customerSigner["dydxprotocol.clob.MsgCancelOrder"] = getLegacyMsgSignerFn(
		[]string{"order_id", "subaccount_id", "owner"})
	customerSigner["dydxprotocol.clob.MsgPlaceOrder"] = getLegacyMsgSignerFn([]string{"order", "order_id", "subaccount_id", "owner"})
	customerSigner["dydxprotocol.sending.MsgCreateTransfer"] = getLegacyMsgSignerFn([]string{"transfer", "sender", "owner"})
	customerSigner["dydxprotocol.sending.MsgWithdrawFromSubaccount"] = getLegacyMsgSignerFn([]string{"sender", "owner"})
	customerSigner["dydxprotocol.vault.MsgDepositToVault"] = getLegacyMsgSignerFn([]string{"subaccount_id", "owner"})

	interfaceRegistryold, _ := types.NewInterfaceRegistryWithOptions(types.InterfaceRegistryOptions{
		ProtoFiles: cosproto.HybridResolver,
		SigningOptions: signing.Options{
			AddressCodec: address.Bech32Codec{
				Bech32Prefix: sdk.GetConfig().GetBech32AccountAddrPrefix(),
			},
			ValidatorAddressCodec: address.Bech32Codec{
				Bech32Prefix: sdk.GetConfig().GetBech32ValidatorAddrPrefix(),
			},
			CustomGetSigners: customerSigner,
		},
	})
	_ = interfaceRegistryold

	interfaceRegistry, err := module.NewInterfaceRegistry(Bech32PrefixAccAddr, Bech32PrefixValAddr)
	if err != nil {
		panic(err)
	}
	legacyAmino := codec.NewLegacyAmino()
	// interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(marshaler, tx.DefaultSignModes)

	std.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterInterfaces(interfaceRegistry)

	return EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             marshaler,
		TxConfig:          txCfg,
		Amino:             legacyAmino,
	}
}
