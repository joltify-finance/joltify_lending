package ante_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	appmsgs "github.com/joltify-finance/joltify_lending/dydx_helper/app/msgs"
	"github.com/joltify-finance/joltify_lending/dydx_helper/lib"
	"github.com/joltify-finance/joltify_lending/dydx_helper/lib/ante"
	testmsgs "github.com/joltify-finance/joltify_lending/dydx_helper/testutil/msgs"
	"github.com/stretchr/testify/require"
)

func TestIsUnsupportedMsg_Empty(t *testing.T) {
	tests := map[string]struct {
		msg sdk.Msg
	}{
		"empty msg": {
			msg: nil,
		},
		"not unsupported msg": {
			msg: testMsg,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			require.False(t, ante.IsUnsupportedMsg(tc.msg))
		})
	}
}

func TestIsUnsupportedMsg_Invalid(t *testing.T) {
	allMsgsMinusUnsupported := lib.MergeAllMapsMustHaveDistinctKeys(appmsgs.AllowMsgs, appmsgs.DisallowMsgs)
	for key := range appmsgs.UnsupportedMsgSamples {
		delete(allMsgsMinusUnsupported, key)
	}
	allNonNilSampleMsgs := testmsgs.GetNonNilSampleMsgs(allMsgsMinusUnsupported)

	for _, sampleMsg := range allNonNilSampleMsgs {
		t.Run(sampleMsg.Name, func(t *testing.T) {
			require.False(t, ante.IsUnsupportedMsg(sampleMsg.Msg))
		})
	}
}

func TestIsUnsupportedMsg_Valid(t *testing.T) {
	sampleMsgs := testmsgs.GetNonNilSampleMsgs(appmsgs.UnsupportedMsgSamples)

	for _, sampleMsg := range sampleMsgs {
		t.Run(sampleMsg.Name, func(t *testing.T) {
			require.True(t, ante.IsUnsupportedMsg(sampleMsg.Msg))
		})
	}
}
