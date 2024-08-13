package constants

import (
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/dydx_helper/x/delaymsg/types"
	bridgetypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/bridge/types"
)

var (
	// MsgCompleteBridge is an example of an expected Msg type in the delaymsg module.
	TestMsg1 = &bridgetypes.MsgCompleteBridge{
		Authority: types.ModuleAddress.String(),
		Event: bridgetypes.BridgeEvent{
			Id: 1,
		},
	}
	TestMsg2 = &bridgetypes.MsgCompleteBridge{
		Authority: types.ModuleAddress.String(),
		Event: bridgetypes.BridgeEvent{
			Id: 2,
		},
	}
	TestMsg3 = &bridgetypes.MsgCompleteBridge{
		Authority: types.ModuleAddress.String(),
		Event: bridgetypes.BridgeEvent{
			Id: 3,
		},
	}
	NoHandlerMsg = &testdata.TestMsg{Signers: []string{types.ModuleAddress.String()}}

	AllMsgs = []sdk.Msg{TestMsg1, TestMsg2, TestMsg3}
)
