package ante

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
	ibcclient "github.com/cosmos/ibc-go/v8/modules/core/02-client/types" //nolint:staticcheck
	ibcconn "github.com/cosmos/ibc-go/v8/modules/core/03-connection/types"
	blocktime "github.com/joltify-finance/joltify_lending/dydx_helper/x/blocktime/types"
	delaymsg "github.com/joltify-finance/joltify_lending/dydx_helper/x/delaymsg/types"
	feetiers "github.com/joltify-finance/joltify_lending/dydx_helper/x/feetiers/types"
	govplus "github.com/joltify-finance/joltify_lending/dydx_helper/x/govplus/types"
	perpetuals "github.com/joltify-finance/joltify_lending/dydx_helper/x/perpetuals/types"
	prices "github.com/joltify-finance/joltify_lending/dydx_helper/x/prices/types"
	ratelimit "github.com/joltify-finance/joltify_lending/dydx_helper/x/ratelimit/types"
	rewards "github.com/joltify-finance/joltify_lending/dydx_helper/x/rewards/types"
	sending "github.com/joltify-finance/joltify_lending/dydx_helper/x/sending/types"
	stats "github.com/joltify-finance/joltify_lending/dydx_helper/x/stats/types"
	vault "github.com/joltify-finance/joltify_lending/dydx_helper/x/vault/types"
	vest "github.com/joltify-finance/joltify_lending/dydx_helper/x/vest/types"
	bridge "github.com/joltify-finance/joltify_lending/x/third_party_dydx/bridge/types"
	clob "github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/types"
)

// IsInternalMsg returns true if the given msg is an internal message.
func IsInternalMsg(msg sdk.Msg) bool {
	switch msg.(type) {
	case
		// ------- CosmosSDK default modules
		// auth
		*auth.MsgUpdateParams,

		// bank
		*bank.MsgSetSendEnabled,
		*bank.MsgUpdateParams,

		// consensus
		*consensus.MsgUpdateParams,

		// crisis
		*crisis.MsgUpdateParams,

		// distribution
		*distribution.MsgCommunityPoolSpend,
		*distribution.MsgUpdateParams,

		// gov
		*gov.MsgExecLegacyContent,
		*gov.MsgUpdateParams,

		// slashing
		*slashing.MsgUpdateParams,

		// staking
		*staking.MsgUpdateParams,

		// upgrade
		*upgrade.MsgCancelUpgrade,
		*upgrade.MsgSoftwareUpgrade,

		// ------- Custom modules
		// blocktime
		*blocktime.MsgUpdateDowntimeParams,

		// bridge
		*bridge.MsgCompleteBridge,
		*bridge.MsgUpdateEventParams,
		*bridge.MsgUpdateProposeParams,
		*bridge.MsgUpdateSafetyParams,

		// clob
		*clob.MsgCreateClobPair,
		*clob.MsgUpdateBlockRateLimitConfiguration,
		*clob.MsgUpdateClobPair,
		*clob.MsgUpdateEquityTierLimitConfiguration,
		*clob.MsgUpdateLiquidationsConfig,

		// delaymsg
		*delaymsg.MsgDelayMessage,

		// feetiers
		*feetiers.MsgUpdatePerpetualFeeParams,

		// govplus
		*govplus.MsgSlashValidator,

		// perpetuals
		*perpetuals.MsgCreatePerpetual,
		*perpetuals.MsgSetLiquidityTier,
		*perpetuals.MsgUpdateParams,
		*perpetuals.MsgUpdatePerpetualParams,

		// prices
		*prices.MsgCreateOracleMarket,
		*prices.MsgUpdateMarketParam,

		// ratelimit
		*ratelimit.MsgSetLimitParams,
		*ratelimit.MsgSetLimitParamsResponse,

		// rewards
		*rewards.MsgUpdateParams,

		// sending
		*sending.MsgSendFromModuleToAccount,

		// stats
		*stats.MsgUpdateParams,

		// vault
		*vault.MsgUpdateParams,

		// vest
		*vest.MsgDeleteVestEntry,
		*vest.MsgSetVestEntry,

		// ibc
		*icahosttypes.MsgUpdateParams,
		*ibctransfer.MsgUpdateParams,
		*ibcclient.MsgUpdateParams,
		*ibcconn.MsgUpdateParams:

		return true

	default:
		return false
	}
}
