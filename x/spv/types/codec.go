package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreatePool{}, "spv/CreatePool", nil)
	cdc.RegisterConcrete(&MsgAddInvestors{}, "spv/AddInvestors", nil)
	cdc.RegisterConcrete(&MsgDeposit{}, "spv/Deposit", nil)
	cdc.RegisterConcrete(&MsgBorrow{}, "spv/Borrow", nil)
	cdc.RegisterConcrete(&MsgRepayInterest{}, "spv/RepayInterest", nil)
	cdc.RegisterConcrete(&MsgClaimInterest{}, "spv/ClaimInterest", nil)
	cdc.RegisterConcrete(&MsgUpdatePool{}, "spv/UpdatePool", nil)
	cdc.RegisterConcrete(&MsgActivePool{}, "spv/ActivePool", nil)
	cdc.RegisterConcrete(&MsgPayPrincipal{}, "spv/PayPrincipal", nil)
	cdc.RegisterConcrete(&MsgWithdrawPrincipal{}, "spv/WithdrawPrincipal", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddInvestors{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeposit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBorrow{},
	)

	registry.RegisterInterface(
		"joltify.spv.BorrowInterest",
		(*NFTBorrowInterest)(nil),
		(*NftInfo)(nil),
	)

	registry.RegisterImplementations((*NFTBorrowInterest)(nil), &BorrowInterest{})

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRepayInterest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgClaimInterest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdatePool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgActivePool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPayPrincipal{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdrawPrincipal{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
