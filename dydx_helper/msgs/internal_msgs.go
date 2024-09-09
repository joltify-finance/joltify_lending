package msgs

import (
	upgrade "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensus "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisis "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distribution "github.com/cosmos/cosmos-sdk/x/distribution/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	slashing "github.com/cosmos/cosmos-sdk/x/slashing/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	icahosttypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/types"
	ibctransfer "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibcclient "github.com/cosmos/ibc-go/v8/modules/core/02-client/types" //nolint:static-check
	ibcconn "github.com/cosmos/ibc-go/v8/modules/core/03-connection/types"
	"github.com/joltify-finance/joltify_lending/lib"
	blocktime "github.com/joltify-finance/joltify_lending/x/third_party_dydx/blocktime/types"
	bridge "github.com/joltify-finance/joltify_lending/x/third_party_dydx/bridge/types"
	clob "github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/types"
	delaymsg "github.com/joltify-finance/joltify_lending/x/third_party_dydx/delaymsg/types"
	feetiers "github.com/joltify-finance/joltify_lending/x/third_party_dydx/feetiers/types"
	govplus "github.com/joltify-finance/joltify_lending/x/third_party_dydx/govplus/types"
	perpetuals "github.com/joltify-finance/joltify_lending/x/third_party_dydx/perpetuals/types"
	prices "github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/types"
	ratelimit "github.com/joltify-finance/joltify_lending/x/third_party_dydx/ratelimit/types"
	rewards "github.com/joltify-finance/joltify_lending/x/third_party_dydx/rewards/types"
	sending "github.com/joltify-finance/joltify_lending/x/third_party_dydx/sending/types"
	stats "github.com/joltify-finance/joltify_lending/x/third_party_dydx/stats/types"
	vault "github.com/joltify-finance/joltify_lending/x/third_party_dydx/vault/types"
	vest "github.com/joltify-finance/joltify_lending/x/third_party_dydx/vest/types"
)

var (
	// InternalMsgSamplesAll are msgs that are used only used internally.
	InternalMsgSamplesAll = lib.MergeAllMapsMustHaveDistinctKeys(InternalMsgSamplesGovAuth)

	// InternalMsgSamplesGovAuth are msgs that are used only used internally.
	// GovAuth means that these messages must originate from the gov module and
	// signed by gov module account.
	// InternalMsgSamplesAll are msgs that are used only used internally.
	InternalMsgSamplesGovAuth = lib.MergeAllMapsMustHaveDistinctKeys(
		InternalMsgSamplesDefault,
		InternalMsgSamplesDydxCustom,
	)

	// CosmosSDK default modules
	InternalMsgSamplesDefault = map[string]sdk.Msg{
		// auth
		"/cosmos.auth.v1beta1.MsgUpdateParams": &auth.MsgUpdateParams{},

		// bank
		"/cosmos.bank.v1beta1.MsgSetSendEnabled":         &bank.MsgSetSendEnabled{},
		"/cosmos.bank.v1beta1.MsgSetSendEnabledResponse": nil,
		"/cosmos.bank.v1beta1.MsgUpdateParams":           &bank.MsgUpdateParams{},
		"/cosmos.bank.v1beta1.MsgUpdateParamsResponse":   nil,

		// consensus
		"/cosmos.consensus.v1.MsgUpdateParams":         &consensus.MsgUpdateParams{},
		"/cosmos.consensus.v1.MsgUpdateParamsResponse": nil,

		// crisis
		"/cosmos.crisis.v1beta1.MsgUpdateParams":         &crisis.MsgUpdateParams{},
		"/cosmos.crisis.v1beta1.MsgUpdateParamsResponse": nil,

		// distribution
		"/cosmos.distribution.v1beta1.MsgCommunityPoolSpend":         &distribution.MsgCommunityPoolSpend{},
		"/cosmos.distribution.v1beta1.MsgCommunityPoolSpendResponse": nil,
		"/cosmos.distribution.v1beta1.MsgUpdateParams":               &distribution.MsgUpdateParams{},
		"/cosmos.distribution.v1beta1.MsgUpdateParamsResponse":       nil,

		// gov
		"/cosmos.gov.v1.MsgExecLegacyContent":         &gov.MsgExecLegacyContent{},
		"/cosmos.gov.v1.MsgExecLegacyContentResponse": nil,
		"/cosmos.gov.v1.MsgUpdateParams":              &gov.MsgUpdateParams{},
		"/cosmos.gov.v1.MsgUpdateParamsResponse":      nil,

		// slashing
		"/cosmos.slashing.v1beta1.MsgUpdateParams":         &slashing.MsgUpdateParams{},
		"/cosmos.slashing.v1beta1.MsgUpdateParamsResponse": nil,

		// staking
		"/cosmos.staking.v1beta1.MsgUpdateParams":         &staking.MsgUpdateParams{},
		"/cosmos.staking.v1beta1.MsgUpdateParamsResponse": nil,

		// upgrade
		"/cosmos.upgrade.v1beta1.MsgCancelUpgrade":           &upgrade.MsgCancelUpgrade{},
		"/cosmos.upgrade.v1beta1.MsgCancelUpgradeResponse":   nil,
		"/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade":         &upgrade.MsgSoftwareUpgrade{},
		"/cosmos.upgrade.v1beta1.MsgSoftwareUpgradeResponse": nil,

		// ibc
		"/ibc.applications.interchain_accounts.host.v1.MsgUpdateParams":         &icahosttypes.MsgUpdateParams{},
		"/ibc.applications.interchain_accounts.host.v1.MsgUpdateParamsResponse": nil,
		"/ibc.applications.transfer.v1.MsgUpdateParams":                         &ibctransfer.MsgUpdateParams{},
		"/ibc.applications.transfer.v1.MsgUpdateParamsResponse":                 nil,
		"/ibc.core.client.v1.MsgUpdateParams":                                   &ibcclient.MsgUpdateParams{},
		"/ibc.core.client.v1.MsgUpdateParamsResponse":                           nil,
		"/ibc.core.connection.v1.MsgUpdateParams":                               &ibcconn.MsgUpdateParams{},
		"/ibc.core.connection.v1.MsgUpdateParamsResponse":                       nil,
	}

	// Custom modules
	InternalMsgSamplesDydxCustom = map[string]sdk.Msg{
		// blocktime
		"/joltify.third_party.dydxprotocol.blocktime.MsgUpdateDowntimeParams":         &blocktime.MsgUpdateDowntimeParams{},
		"/joltify.third_party.dydxprotocol.blocktime.MsgUpdateDowntimeParamsResponse": nil,

		// bridge
		"/joltify.third_party.dydxprotocol.bridge.MsgCompleteBridge":              &bridge.MsgCompleteBridge{},
		"/joltify.third_party.dydxprotocol.bridge.MsgCompleteBridgeResponse":      nil,
		"/joltify.third_party.dydxprotocol.bridge.MsgUpdateEventParams":           &bridge.MsgUpdateEventParams{},
		"/joltify.third_party.dydxprotocol.bridge.MsgUpdateEventParamsResponse":   nil,
		"/joltify.third_party.dydxprotocol.bridge.MsgUpdateProposeParams":         &bridge.MsgUpdateProposeParams{},
		"/joltify.third_party.dydxprotocol.bridge.MsgUpdateProposeParamsResponse": nil,
		"/joltify.third_party.dydxprotocol.bridge.MsgUpdateSafetyParams":          &bridge.MsgUpdateSafetyParams{},
		"/joltify.third_party.dydxprotocol.bridge.MsgUpdateSafetyParamsResponse":  nil,

		// clob
		"/joltify.third_party.dydxprotocol.clob.MsgCreateClobPair":                             &clob.MsgCreateClobPair{},
		"/joltify.third_party.dydxprotocol.clob.MsgCreateClobPairResponse":                     nil,
		"/joltify.third_party.dydxprotocol.clob.MsgUpdateBlockRateLimitConfiguration":          &clob.MsgUpdateBlockRateLimitConfiguration{},
		"/joltify.third_party.dydxprotocol.clob.MsgUpdateBlockRateLimitConfigurationResponse":  nil,
		"/joltify.third_party.dydxprotocol.clob.MsgUpdateClobPair":                             &clob.MsgUpdateClobPair{},
		"/joltify.third_party.dydxprotocol.clob.MsgUpdateClobPairResponse":                     nil,
		"/joltify.third_party.dydxprotocol.clob.MsgUpdateEquityTierLimitConfiguration":         &clob.MsgUpdateEquityTierLimitConfiguration{},
		"/joltify.third_party.dydxprotocol.clob.MsgUpdateEquityTierLimitConfigurationResponse": nil,
		"/joltify.third_party.dydxprotocol.clob.MsgUpdateLiquidationsConfig":                   &clob.MsgUpdateLiquidationsConfig{},
		"/joltify.third_party.dydxprotocol.clob.MsgUpdateLiquidationsConfigResponse":           nil,

		// delaymsg
		"/joltify.third_party.dydxprotocol.delaymsg.MsgDelayMessage":         &delaymsg.MsgDelayMessage{},
		"/joltify.third_party.dydxprotocol.delaymsg.MsgDelayMessageResponse": nil,

		// feetiers
		"/joltify.third_party.dydxprotocol.feetiers.MsgUpdatePerpetualFeeParams":         &feetiers.MsgUpdatePerpetualFeeParams{},
		"/joltify.third_party.dydxprotocol.feetiers.MsgUpdatePerpetualFeeParamsResponse": nil,

		// govplus
		"/joltify.third_party.dydxprotocol.govplus.MsgSlashValidator":         &govplus.MsgSlashValidator{},
		"/joltify.third_party.dydxprotocol.govplus.MsgSlashValidatorResponse": nil,

		// perpetuals
		"/joltify.third_party.dydxprotocol.perpetuals.MsgCreatePerpetual":               &perpetuals.MsgCreatePerpetual{},
		"/joltify.third_party.dydxprotocol.perpetuals.MsgCreatePerpetualResponse":       nil,
		"/joltify.third_party.dydxprotocol.perpetuals.MsgSetLiquidityTier":              &perpetuals.MsgSetLiquidityTier{},
		"/joltify.third_party.dydxprotocol.perpetuals.MsgSetLiquidityTierResponse":      nil,
		"/joltify.third_party.dydxprotocol.perpetuals.MsgUpdateParams":                  &perpetuals.MsgUpdateParams{},
		"/joltify.third_party.dydxprotocol.perpetuals.MsgUpdateParamsResponse":          nil,
		"/joltify.third_party.dydxprotocol.perpetuals.MsgUpdatePerpetualParams":         &perpetuals.MsgUpdatePerpetualParams{},
		"/joltify.third_party.dydxprotocol.perpetuals.MsgUpdatePerpetualParamsResponse": nil,

		// prices
		"/joltify.third_party.dydxprotocol.prices.MsgCreateOracleMarket":         &prices.MsgCreateOracleMarket{},
		"/joltify.third_party.dydxprotocol.prices.MsgCreateOracleMarketResponse": nil,
		"/joltify.third_party.dydxprotocol.prices.MsgUpdateMarketParam":          &prices.MsgUpdateMarketParam{},
		"/joltify.third_party.dydxprotocol.prices.MsgUpdateMarketParamResponse":  nil,

		// ratelimit
		"/joltify.third_party.dydxprotocol.ratelimit.MsgSetLimitParams":         &ratelimit.MsgSetLimitParams{},
		"/joltify.third_party.dydxprotocol.ratelimit.MsgSetLimitParamsResponse": nil,

		// rewards
		"/joltify.third_party.dydxprotocol.rewards.MsgUpdateParams":         &rewards.MsgUpdateParams{},
		"/joltify.third_party.dydxprotocol.rewards.MsgUpdateParamsResponse": nil,

		// sending
		"/joltify.third_party.dydxprotocol.sending.MsgSendFromModuleToAccount":         &sending.MsgSendFromModuleToAccount{},
		"/joltify.third_party.dydxprotocol.sending.MsgSendFromModuleToAccountResponse": nil,

		// stats
		"/joltify.third_party.dydxprotocol.stats.MsgUpdateParams":         &stats.MsgUpdateParams{},
		"/joltify.third_party.dydxprotocol.stats.MsgUpdateParamsResponse": nil,

		// vault
		"/joltify.third_party.dydxprotocol.vault.MsgUpdateParams":         &vault.MsgUpdateParams{},
		"/joltify.third_party.dydxprotocol.vault.MsgUpdateParamsResponse": nil,

		// vest
		"/joltify.third_party.dydxprotocol.vest.MsgSetVestEntry":            &vest.MsgSetVestEntry{},
		"/joltify.third_party.dydxprotocol.vest.MsgSetVestEntryResponse":    nil,
		"/joltify.third_party.dydxprotocol.vest.MsgDeleteVestEntry":         &vest.MsgDeleteVestEntry{},
		"/joltify.third_party.dydxprotocol.vest.MsgDeleteVestEntryResponse": nil,
	}
)
