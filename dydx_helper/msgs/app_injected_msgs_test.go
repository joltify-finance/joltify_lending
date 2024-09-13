package msgs_test

import (
	"sort"
	"strings"
	"testing"

	"github.com/cosmos/gogoproto/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	"github.com/joltify-finance/joltify_lending/dydx_helper/msgs"
	"github.com/joltify-finance/joltify_lending/lib"
	"github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/encoding"
	testmsgs "github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/msgs"
	"github.com/stretchr/testify/require"
)

func TestAppInjectedMsgSamples_Key(t *testing.T) {
	expectedMsgs := []string{
		// bridge
		"/joltify.third_party.dydxprotocol.bridge.MsgAcknowledgeBridges",
		"/joltify.third_party.dydxprotocol.bridge.MsgAcknowledgeBridgesResponse",

		// clob
		"/joltify.third_party.dydxprotocol.clob.MsgProposedOperations",
		"/joltify.third_party.dydxprotocol.clob.MsgProposedOperationsResponse",

		// perpetuals
		"/joltify.third_party.dydxprotocol.perpetuals.MsgAddPremiumVotes",
		"/joltify.third_party.dydxprotocol.perpetuals.MsgAddPremiumVotesResponse",

		// prices
		"/joltify.third_party.dydxprotocol.prices.MsgUpdateMarketPrices",
		"/joltify.third_party.dydxprotocol.prices.MsgUpdateMarketPricesResponse",
	}

	require.Equal(t, expectedMsgs, lib.GetSortedKeys[sort.StringSlice](msgs.AppInjectedMsgSamples))
}

func TestAppInjectedMsgSamples_Value(t *testing.T) {
	validateMsgValue(t, msgs.AppInjectedMsgSamples)
}

func TestAppInjectedMsgSamples_GetSigners(t *testing.T) {
	testEncodingCfg := encoding.GetTestEncodingCfg()
	testTxBuilder := testEncodingCfg.TxConfig.NewTxBuilder()

	for _, sample := range testmsgs.GetNonNilSampleMsgs(msgs.AppInjectedMsgSamples) {
		_ = testTxBuilder.SetMsgs(sample.Msg)
		sigTx, ok := testTxBuilder.GetTx().(authsigning.SigVerifiableTx)
		require.True(t, ok)
		signers, err := sigTx.GetSigners()
		require.NoError(t, err)
		require.Empty(t, signers)
	}
}

// validateMsgValue ensures that the message is
//  1. not nil for "<module>.<version>.Msg<Name>"
//  2. sample msg's proto msg name matches the key it's registered under
//  3. nil sample message for others
func validateMsgValue(
	t *testing.T,
	sampleMsgs map[string]sdk.Msg,
) {
	for key, sample := range sampleMsgs {
		keyTokens := strings.Split(key, ".")
		if testmsgs.IsValidMsgFormat(keyTokens) && !strings.HasSuffix(key, "Response") {
			// Sample msg cannot be nil.
			require.NotNil(t, sample, "key: %s", key)

			// Sample msg type must match the key it's registered under.
			expectedTypeUrl := "/joltify.third_party." + proto.MessageName(sample)
			require.Equal(t, expectedTypeUrl, key)
		} else {
			// "Response" messages are msgs that cannot be submitted, so no sample is provided.
			// Additionally, all other intermediary msgs should not be submitted as a top-level msg.
			require.Nil(t, sample)
		}
	}
}
