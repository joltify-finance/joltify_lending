package msgs_test

import (
	"sort"
	"testing"

	"github.com/joltify-finance/joltify_lending/dydx_helper/msgs"
	"github.com/joltify-finance/joltify_lending/lib"
	"github.com/stretchr/testify/require"
)

func TestNestedMsgs_Key(t *testing.T) {
	expectedMsgs := []string{
		// authz
		"/cosmos.authz.v1beta1.MsgExec",
		"/cosmos.authz.v1beta1.MsgExecResponse",

		// gov
		"/cosmos.gov.v1.MsgSubmitProposal",
		"/cosmos.gov.v1.MsgSubmitProposalResponse",
	}
	require.Equal(t, expectedMsgs, lib.GetSortedKeys[sort.StringSlice](msgs.NestedMsgSamples))
}

func TestNestedMsgs_Value(t *testing.T) {
	validateMsgValue(t, msgs.NestedMsgSamples)
}
