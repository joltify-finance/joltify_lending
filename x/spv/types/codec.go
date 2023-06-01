package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreatePool{}, "/joltify.spv.MsgCreatePool", nil)
	cdc.RegisterConcrete(&MsgAddInvestors{}, "/joltify.spv.MsgAddInvestors", nil)
	cdc.RegisterConcrete(&MsgDeposit{}, "/joltify.spv.MsgDeposit", nil)
	cdc.RegisterConcrete(&MsgBorrow{}, "/joltify.spv.MsgBorrow", nil)
	cdc.RegisterConcrete(&MsgRepayInterest{}, "/joltify.spv.MsgRepayInterest", nil)
	cdc.RegisterConcrete(&MsgClaimInterest{}, "/joltify.spv.MsgClaimInterest", nil)
	cdc.RegisterConcrete(&MsgUpdatePool{}, "/joltify.spv.MsgUpdatePool", nil)
	cdc.RegisterConcrete(&MsgActivePool{}, "/joltify.spv.MsgActivePool", nil)
	cdc.RegisterConcrete(&MsgPayPrincipal{}, "/joltify.spv.MsgPayPrincipal", nil)
	cdc.RegisterConcrete(&MsgPayPrincipalPartial{}, "/joltify.spv.MsgPayPrincipalPartial", nil)
	cdc.RegisterConcrete(&MsgWithdrawPrincipal{}, "/joltify.spv.MsgWithdrawPrincipal", nil)
	cdc.RegisterConcrete(&MsgSubmitWithdrawProposal{}, "/joltify.spv.MsgSubmitWithdrawProposal", nil)
	cdc.RegisterConcrete(&MsgTransferOwnership{}, "/joltify.spv.MsgTransferOwnership", nil)
	cdc.RegisterConcrete(&MsgLiquidate{}, "/joltify.spv.MsgLiquidate", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterInterface(
		"joltify.spv.BorrowInterest",
		(*NFTBorrowInterest)(nil),
		(*NftInfo)(nil),
	)

	registry.RegisterImplementations((*NFTBorrowInterest)(nil), &BorrowInterest{})

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePool{},
		&MsgAddInvestors{},
		&MsgDeposit{},
		&MsgBorrow{},
		&MsgRepayInterest{},
		&MsgClaimInterest{},
		&MsgUpdatePool{},
		&MsgActivePool{},
		&MsgPayPrincipal{},
		&MsgPayPrincipalPartial{},
		&MsgWithdrawPrincipal{},
		&MsgSubmitWithdrawProposal{},
		&MsgTransferOwnership{},
		&MsgLiquidate{},
	)
	// this line is used by starport scaffolding # 3
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino = codec.NewLegacyAmino()
	// ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func init() {
	RegisterCodec(Amino)
	cryptocodec.RegisterCrypto(Amino)
}
