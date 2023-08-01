package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers the necessary evmutil interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgConvertCoinToERC20{}, "/joltify.third_party.evmutil.v1beta1.MsgConvertCoinToERC20", nil)
	cdc.RegisterConcrete(&MsgConvertERC20ToCoin{}, "/joltify.third_party.evmutil.v1beta1.MsgConvertERC20ToCoin", nil)
	cdc.RegisterConcrete(&MsgConvertCosmosCoinToERC20{}, "/joltify.third_party.evmutil.v1beta1.MsgConvertCosmosCoinToERC20", nil)
	cdc.RegisterConcrete(&MsgConvertCosmosCoinFromERC20{}, "/joltify.third_party.evmutil.v1beta1.MsgConvertCosmosCoinFromERC20", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgConvertCoinToERC20{},
		&MsgConvertERC20ToCoin{},
		&MsgConvertCosmosCoinToERC20{},
		&MsgConvertCosmosCoinFromERC20{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
}
