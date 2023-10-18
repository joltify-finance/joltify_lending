package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers all the necessary types and interfaces for the
// governance module.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgDeposit{}, "/joltify.third_party.swap.v1beta1.MsgDeposit", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "/joltify.third_party.swap.v1beta1.MsgWithdraw", nil)
	cdc.RegisterConcrete(&MsgSwapExactForTokens{}, "/joltify.third_party.swap.v1beta1.MsgSwapExactForTokens", nil)
	cdc.RegisterConcrete(&MsgSwapExactForBatchTokens{}, "/joltify.third_party.swap.v1beta1.MsgSwapForExactBatchTokens", nil)
	cdc.RegisterConcrete(&MsgSwapForExactTokens{}, "/joltify.third_party.swap.v1beta1.MsgSwapForExactTokens", nil)
}

// RegisterInterfaces registers proto messages under their interfaces for unmarshalling,
// in addition to registerting the msg service for handling tx msgs
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeposit{},
		&MsgWithdraw{},
		&MsgSwapExactForTokens{},
		&MsgSwapForExactTokens{},
		&MsgSwapExactForBatchTokens{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()
	// ModuleCdc represents the legacy amino codec for the module
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
}
