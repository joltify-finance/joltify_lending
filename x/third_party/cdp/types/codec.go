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
	cdc.RegisterConcrete(&MsgCreateCDP{}, "/joltify.third_party.cdp.v1beta1.MsgCreateCDP", nil)
	cdc.RegisterConcrete(&MsgDeposit{}, "/joltify.third_party.cdp.v1beta1.MsgDeposit", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "/joltify.third_party.cdp.v1beta1.MsgWithdraw", nil)
	cdc.RegisterConcrete(&MsgDrawDebt{}, "/joltify.third_party.cdp.v1beta1.MsgDrawDebt", nil)
	cdc.RegisterConcrete(&MsgRepayDebt{}, "/joltify.third_party.cdp.v1beta1.MsgRepayDebt", nil)
	cdc.RegisterConcrete(&MsgLiquidate{}, "/joltify.third_party.cdp.v1beta1.MsgLiquidate", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCDP{},
		&MsgDeposit{},
		&MsgWithdraw{},
		&MsgDrawDebt{},
		&MsgRepayDebt{},
		&MsgLiquidate{},
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
