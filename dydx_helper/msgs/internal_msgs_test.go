package msgs_test

import (
	"sort"
	"testing"

	"github.com/joltify-finance/joltify_lending/dydx_helper/msgs"
	"github.com/joltify-finance/joltify_lending/lib"
	"github.com/stretchr/testify/require"
)

func TestInternalMsgSamples_All_Key(t *testing.T) {
	expectedAllInternalMsgs := lib.MergeAllMapsMustHaveDistinctKeys(msgs.InternalMsgSamplesGovAuth)
	require.Equal(t, expectedAllInternalMsgs, msgs.InternalMsgSamplesAll)
}

func TestInternalMsgSamples_All_Value(t *testing.T) {
	validateMsgValue(t, msgs.InternalMsgSamplesAll)
}

func TestInternalMsgSamples_Gov_Key(t *testing.T) {
	expectedMsgs := []string{
		// auth
		"/cosmos.auth.v1beta1.MsgUpdateParams",

		// bank
		"/cosmos.bank.v1beta1.MsgSetSendEnabled",
		"/cosmos.bank.v1beta1.MsgSetSendEnabledResponse",
		"/cosmos.bank.v1beta1.MsgUpdateParams",
		"/cosmos.bank.v1beta1.MsgUpdateParamsResponse",

		// consensus
		"/cosmos.consensus.v1.MsgUpdateParams",
		"/cosmos.consensus.v1.MsgUpdateParamsResponse",

		// crisis
		"/cosmos.crisis.v1beta1.MsgUpdateParams",
		"/cosmos.crisis.v1beta1.MsgUpdateParamsResponse",

		// distribution
		"/cosmos.distribution.v1beta1.MsgCommunityPoolSpend",
		"/cosmos.distribution.v1beta1.MsgCommunityPoolSpendResponse",
		"/cosmos.distribution.v1beta1.MsgUpdateParams",
		"/cosmos.distribution.v1beta1.MsgUpdateParamsResponse",

		// gov
		"/cosmos.gov.v1.MsgExecLegacyContent",
		"/cosmos.gov.v1.MsgExecLegacyContentResponse",
		"/cosmos.gov.v1.MsgUpdateParams",
		"/cosmos.gov.v1.MsgUpdateParamsResponse",

		// slashing
		"/cosmos.slashing.v1beta1.MsgUpdateParams",
		"/cosmos.slashing.v1beta1.MsgUpdateParamsResponse",

		// staking
		"/cosmos.staking.v1beta1.MsgUpdateParams",
		"/cosmos.staking.v1beta1.MsgUpdateParamsResponse",

		// upgrade
		"/cosmos.upgrade.v1beta1.MsgCancelUpgrade",
		"/cosmos.upgrade.v1beta1.MsgCancelUpgradeResponse",
		"/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade",
		"/cosmos.upgrade.v1beta1.MsgSoftwareUpgradeResponse",

		// blocktime
		"/joltify.third_party.dydxprotocol.blocktime.MsgUpdateDowntimeParams",
		"/joltify.third_party.dydxprotocol.blocktime.MsgUpdateDowntimeParamsResponse",

		// bridge
		"/joltify.third_party.dydxprotocol.bridge.MsgCompleteBridge",
		"/joltify.third_party.dydxprotocol.bridge.MsgCompleteBridgeResponse",
		"/joltify.third_party.dydxprotocol.bridge.MsgUpdateEventParams",
		"/joltify.third_party.dydxprotocol.bridge.MsgUpdateEventParamsResponse",
		"/joltify.third_party.dydxprotocol.bridge.MsgUpdateProposeParams",
		"/joltify.third_party.dydxprotocol.bridge.MsgUpdateProposeParamsResponse",
		"/joltify.third_party.dydxprotocol.bridge.MsgUpdateSafetyParams",
		"/joltify.third_party.dydxprotocol.bridge.MsgUpdateSafetyParamsResponse",

		// clob
		"/joltify.third_party.dydxprotocol.clob.MsgCreateClobPair",
		"/joltify.third_party.dydxprotocol.clob.MsgCreateClobPairResponse",
		"/joltify.third_party.dydxprotocol.clob.MsgUpdateBlockRateLimitConfiguration",
		"/joltify.third_party.dydxprotocol.clob.MsgUpdateBlockRateLimitConfigurationResponse",
		"/joltify.third_party.dydxprotocol.clob.MsgUpdateClobPair",
		"/joltify.third_party.dydxprotocol.clob.MsgUpdateClobPairResponse",
		"/joltify.third_party.dydxprotocol.clob.MsgUpdateEquityTierLimitConfiguration",
		"/joltify.third_party.dydxprotocol.clob.MsgUpdateEquityTierLimitConfigurationResponse",
		"/joltify.third_party.dydxprotocol.clob.MsgUpdateLiquidationsConfig",
		"/joltify.third_party.dydxprotocol.clob.MsgUpdateLiquidationsConfigResponse",

		// delaymsg
		"/joltify.third_party.dydxprotocol.delaymsg.MsgDelayMessage",
		"/joltify.third_party.dydxprotocol.delaymsg.MsgDelayMessageResponse",

		// feetiers
		"/joltify.third_party.dydxprotocol.feetiers.MsgUpdatePerpetualFeeParams",
		"/joltify.third_party.dydxprotocol.feetiers.MsgUpdatePerpetualFeeParamsResponse",

		// govplus
		"/joltify.third_party.dydxprotocol.govplus.MsgSlashValidator",
		"/joltify.third_party.dydxprotocol.govplus.MsgSlashValidatorResponse",

		// perpeutals
		"/joltify.third_party.dydxprotocol.perpetuals.MsgCreatePerpetual",
		"/joltify.third_party.dydxprotocol.perpetuals.MsgCreatePerpetualResponse",
		"/joltify.third_party.dydxprotocol.perpetuals.MsgSetLiquidityTier",
		"/joltify.third_party.dydxprotocol.perpetuals.MsgSetLiquidityTierResponse",
		"/joltify.third_party.dydxprotocol.perpetuals.MsgUpdateParams",
		"/joltify.third_party.dydxprotocol.perpetuals.MsgUpdateParamsResponse",
		"/joltify.third_party.dydxprotocol.perpetuals.MsgUpdatePerpetualParams",
		"/joltify.third_party.dydxprotocol.perpetuals.MsgUpdatePerpetualParamsResponse",

		// prices
		"/joltify.third_party.dydxprotocol.prices.MsgCreateOracleMarket",
		"/joltify.third_party.dydxprotocol.prices.MsgCreateOracleMarketResponse",
		"/joltify.third_party.dydxprotocol.prices.MsgUpdateMarketParam",
		"/joltify.third_party.dydxprotocol.prices.MsgUpdateMarketParamResponse",

		// ratelimit
		"/joltify.third_party.dydxprotocol.ratelimit.MsgSetLimitParams",
		"/joltify.third_party.dydxprotocol.ratelimit.MsgSetLimitParamsResponse",

		// rewards
		"/joltify.third_party.dydxprotocol.rewards.MsgUpdateParams",
		"/joltify.third_party.dydxprotocol.rewards.MsgUpdateParamsResponse",

		// sending
		"/joltify.third_party.dydxprotocol.sending.MsgSendFromModuleToAccount",
		"/joltify.third_party.dydxprotocol.sending.MsgSendFromModuleToAccountResponse",

		// stats
		"/joltify.third_party.dydxprotocol.stats.MsgUpdateParams",
		"/joltify.third_party.dydxprotocol.stats.MsgUpdateParamsResponse",

		// vault
		"/joltify.third_party.dydxprotocol.vault.MsgUpdateParams",
		"/joltify.third_party.dydxprotocol.vault.MsgUpdateParamsResponse",

		// vest
		"/joltify.third_party.dydxprotocol.vest.MsgDeleteVestEntry",
		"/joltify.third_party.dydxprotocol.vest.MsgDeleteVestEntryResponse",
		"/joltify.third_party.dydxprotocol.vest.MsgSetVestEntry",
		"/joltify.third_party.dydxprotocol.vest.MsgSetVestEntryResponse",

		// ibc
		"/ibc.applications.interchain_accounts.host.v1.MsgUpdateParams",
		"/ibc.applications.interchain_accounts.host.v1.MsgUpdateParamsResponse",
		"/ibc.applications.transfer.v1.MsgUpdateParams",
		"/ibc.applications.transfer.v1.MsgUpdateParamsResponse",
		"/ibc.core.client.v1.MsgUpdateParams",
		"/ibc.core.client.v1.MsgUpdateParamsResponse",
		"/ibc.core.connection.v1.MsgUpdateParams",
		"/ibc.core.connection.v1.MsgUpdateParamsResponse",
	}

	require.Equal(t, expectedMsgs, lib.GetSortedKeys[sort.StringSlice](msgs.InternalMsgSamplesGovAuth))
}

func TestInternalMsgSamples_Gov_Value(t *testing.T) {
	validateMsgValue(t, msgs.InternalMsgSamplesGovAuth)
}
