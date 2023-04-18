package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"

	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgDeposit{}, "/joltify.third_party.jolt.v1beta1.MsgDeposit", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "/joltify.third_party.jolt.v1beta1.MsgWithdraw", nil)
	cdc.RegisterConcrete(&MsgBorrow{}, "/joltify.third_party.jolt.v1beta1.MsgBorrow", nil)
	cdc.RegisterConcrete(&MsgLiquidate{}, "/joltify.third_party.jolt.v1beta1.MsgLiquidate", nil)
	cdc.RegisterConcrete(&MsgRepay{}, "/joltify.third_party.jolt.v1beta1.MsgRepay", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeposit{},
		&MsgWithdraw{},
		&MsgBorrow{},
		&MsgLiquidate{},
		&MsgRepay{},
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
