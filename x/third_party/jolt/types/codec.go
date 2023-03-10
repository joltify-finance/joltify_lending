package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"

	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgDeposit{}, "jolt/MsgDeposit", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "jolt/MsgWithdraw", nil)
	cdc.RegisterConcrete(&MsgBorrow{}, "jolt/MsgBorrow", nil)
	cdc.RegisterConcrete(&MsgLiquidate{}, "jolt/MsgLiquidate", nil)
	cdc.RegisterConcrete(&MsgRepay{}, "jolt/MsgRepay", nil)
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
