package msgs

import (
	"github.com/joltify-finance/joltify_lending/lib"
)

var (
	// AllTypeMessages is a list of all messages and types that are used in the app.
	// This list comes from the app's `InterfaceRegistry`.
	//AllTypeMessages = map[string]struct{}{
	//	// auth
	//	"/cosmos.auth.v1beta1.BaseAccount":      {},
	//	"/cosmos.auth.v1beta1.ModuleAccount":    {},
	//	"/cosmos.auth.v1beta1.ModuleCredential": {},
	//	"/cosmos.auth.v1beta1.MsgUpdateParams":  {},
	//
	//	// authz
	//	"/cosmos.authz.v1beta1.GenericAuthorization": {},
	//	"/cosmos.authz.v1beta1.MsgExec":              {},
	//	"/cosmos.authz.v1beta1.MsgExecResponse":      {},
	//	"/cosmos.authz.v1beta1.MsgGrant":             {},
	//	"/cosmos.authz.v1beta1.MsgGrantResponse":     {},
	//	"/cosmos.authz.v1beta1.MsgRevoke":            {},
	//	"/cosmos.authz.v1beta1.MsgRevokeResponse":    {},
	//
	//	// bank
	//	"/cosmos.bank.v1beta1.MsgMultiSend":              {},
	//	"/cosmos.bank.v1beta1.MsgMultiSendResponse":      {},
	//	"/cosmos.bank.v1beta1.MsgSend":                   {},
	//	"/cosmos.bank.v1beta1.MsgSendResponse":           {},
	//	"/cosmos.bank.v1beta1.MsgSetSendEnabled":         {},
	//	"/cosmos.bank.v1beta1.MsgSetSendEnabledResponse": {},
	//	"/cosmos.bank.v1beta1.MsgUpdateParams":           {},
	//	"/cosmos.bank.v1beta1.MsgUpdateParamsResponse":   {},
	//	"/cosmos.bank.v1beta1.SendAuthorization":         {},
	//	"/cosmos.bank.v1beta1.Supply":                    {},
	//
	//	// george we disable it
	//	// consensus
	//	//"/cosmos.consensus.v1.MsgUpdateParams":         {},
	//	//"/cosmos.consensus.v1.MsgUpdateParamsResponse": {},
	//
	//	// crisis
	//	"/cosmos.crisis.v1beta1.MsgUpdateParams":            {},
	//	"/cosmos.crisis.v1beta1.MsgUpdateParamsResponse":    {},
	//	"/cosmos.crisis.v1beta1.MsgVerifyInvariant":         {},
	//	"/cosmos.crisis.v1beta1.MsgVerifyInvariantResponse": {},
	//
	//	// crypto
	//	"/cosmos.crypto.ed25519.PrivKey":            {},
	//	"/cosmos.crypto.ed25519.PubKey":             {},
	//	"/cosmos.crypto.multisig.LegacyAminoPubKey": {},
	//	"/cosmos.crypto.secp256k1.PrivKey":          {},
	//	"/cosmos.crypto.secp256k1.PubKey":           {},
	//	"/cosmos.crypto.secp256r1.PubKey":           {},
	//
	//	// distribution
	//	"/cosmos.distribution.v1beta1.CommunityPoolSpendProposal":             {},
	//	"/cosmos.distribution.v1beta1.MsgCommunityPoolSpend":                  {},
	//	"/cosmos.distribution.v1beta1.MsgCommunityPoolSpendResponse":          {},
	//	"/cosmos.distribution.v1beta1.MsgDepositValidatorRewardsPool":         {},
	//	"/cosmos.distribution.v1beta1.MsgDepositValidatorRewardsPoolResponse": {},
	//	"/cosmos.distribution.v1beta1.MsgFundCommunityPool":                   {},
	//	"/cosmos.distribution.v1beta1.MsgFundCommunityPoolResponse":           {},
	//	"/cosmos.distribution.v1beta1.MsgSetWithdrawAddress":                  {},
	//	"/cosmos.distribution.v1beta1.MsgSetWithdrawAddressResponse":          {},
	//	"/cosmos.distribution.v1beta1.MsgUpdateParams":                        {},
	//	"/cosmos.distribution.v1beta1.MsgUpdateParamsResponse":                {},
	//	"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward":             {},
	//	"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorRewardResponse":     {},
	//	"/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission":         {},
	//	"/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommissionResponse": {},
	//
	//	// evidence
	//	"/cosmos.evidence.v1beta1.Equivocation":              {},
	//	"/cosmos.evidence.v1beta1.MsgSubmitEvidence":         {},
	//	"/cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse": {},
	//
	//	// feegrant
	//	"/cosmos.feegrant.v1beta1.AllowedMsgAllowance":        {},
	//	"/cosmos.feegrant.v1beta1.BasicAllowance":             {},
	//	"/cosmos.feegrant.v1beta1.MsgGrantAllowance":          {},
	//	"/cosmos.feegrant.v1beta1.MsgGrantAllowanceResponse":  {},
	//	"/cosmos.feegrant.v1beta1.MsgPruneAllowances":         {},
	//	"/cosmos.feegrant.v1beta1.MsgPruneAllowancesResponse": {},
	//	"/cosmos.feegrant.v1beta1.MsgRevokeAllowance":         {},
	//	"/cosmos.feegrant.v1beta1.MsgRevokeAllowanceResponse": {},
	//	"/cosmos.feegrant.v1beta1.PeriodicAllowance":          {},
	//
	//	// gov
	//	"/cosmos.gov.v1.MsgCancelProposal":              {},
	//	"/cosmos.gov.v1.MsgCancelProposalResponse":      {},
	//	"/cosmos.gov.v1.MsgDeposit":                     {},
	//	"/cosmos.gov.v1.MsgDepositResponse":             {},
	//	"/cosmos.gov.v1.MsgExecLegacyContent":           {},
	//	"/cosmos.gov.v1.MsgExecLegacyContentResponse":   {},
	//	"/cosmos.gov.v1.MsgSubmitProposal":              {},
	//	"/cosmos.gov.v1.MsgSubmitProposalResponse":      {},
	//	"/cosmos.gov.v1.MsgUpdateParams":                {},
	//	"/cosmos.gov.v1.MsgUpdateParamsResponse":        {},
	//	"/cosmos.gov.v1.MsgVote":                        {},
	//	"/cosmos.gov.v1.MsgVoteResponse":                {},
	//	"/cosmos.gov.v1.MsgVoteWeighted":                {},
	//	"/cosmos.gov.v1.MsgVoteWeightedResponse":        {},
	//	"/cosmos.gov.v1beta1.MsgDeposit":                {},
	//	"/cosmos.gov.v1beta1.MsgDepositResponse":        {},
	//	"/cosmos.gov.v1beta1.MsgSubmitProposal":         {},
	//	"/cosmos.gov.v1beta1.MsgSubmitProposalResponse": {},
	//	"/cosmos.gov.v1beta1.MsgVote":                   {},
	//	"/cosmos.gov.v1beta1.MsgVoteResponse":           {},
	//	"/cosmos.gov.v1beta1.MsgVoteWeighted":           {},
	//	"/cosmos.gov.v1beta1.MsgVoteWeightedResponse":   {},
	//	"/cosmos.gov.v1beta1.TextProposal":              {},
	//
	//	// params
	//	"/cosmos.params.v1beta1.ParameterChangeProposal": {},
	//
	//	// slashing
	//	"/cosmos.slashing.v1beta1.MsgUnjail":               {},
	//	"/cosmos.slashing.v1beta1.MsgUnjailResponse":       {},
	//	"/cosmos.slashing.v1beta1.MsgUpdateParams":         {},
	//	"/cosmos.slashing.v1beta1.MsgUpdateParamsResponse": {},
	//
	//	// staking
	//	"/cosmos.staking.v1beta1.MsgBeginRedelegate":                   {},
	//	"/cosmos.staking.v1beta1.MsgBeginRedelegateResponse":           {},
	//	"/cosmos.staking.v1beta1.MsgCancelUnbondingDelegation":         {},
	//	"/cosmos.staking.v1beta1.MsgCancelUnbondingDelegationResponse": {},
	//	"/cosmos.staking.v1beta1.MsgCreateValidator":                   {},
	//	"/cosmos.staking.v1beta1.MsgCreateValidatorResponse":           {},
	//	"/cosmos.staking.v1beta1.MsgDelegate":                          {},
	//	"/cosmos.staking.v1beta1.MsgDelegateResponse":                  {},
	//	"/cosmos.staking.v1beta1.MsgEditValidator":                     {},
	//	"/cosmos.staking.v1beta1.MsgEditValidatorResponse":             {},
	//	"/cosmos.staking.v1beta1.MsgUndelegate":                        {},
	//	"/cosmos.staking.v1beta1.MsgUndelegateResponse":                {},
	//	"/cosmos.staking.v1beta1.MsgUpdateParams":                      {},
	//	"/cosmos.staking.v1beta1.MsgUpdateParamsResponse":              {},
	//	"/cosmos.staking.v1beta1.StakeAuthorization":                   {},
	//
	//	// tx
	//	"/cosmos.tx.v1beta1.Tx": {},
	//
	//	// upgrade
	//	"/cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal": {},
	//	"/cosmos.upgrade.v1beta1.MsgCancelUpgrade":              {},
	//	"/cosmos.upgrade.v1beta1.MsgCancelUpgradeResponse":      {},
	//	"/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade":            {},
	//	"/cosmos.upgrade.v1beta1.MsgSoftwareUpgradeResponse":    {},
	//	"/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal":       {},
	//
	//	// vesting
	//	"/cosmos.vesting.v1beta1.BaseVestingAccount":                      {},
	//	"/cosmos.vesting.v1beta1.ContinuousVestingAccount":                {},
	//	"/cosmos.vesting.v1beta1.DelayedVestingAccount":                   {},
	//	"/cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount":         {},
	//	"/cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse": {},
	//	"/cosmos.vesting.v1beta1.MsgCreatePermanentLockedAccount":         {},
	//	"/cosmos.vesting.v1beta1.MsgCreatePermanentLockedAccountResponse": {},
	//	"/cosmos.vesting.v1beta1.MsgCreateVestingAccount":                 {},
	//	"/cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse":         {},
	//	"/cosmos.vesting.v1beta1.PeriodicVestingAccount":                  {},
	//	"/cosmos.vesting.v1beta1.PermanentLockedAccount":                  {},
	//
	//	// blocktime
	//	"/joltify.third_party.dydxprotocol.blocktime.MsgUpdateDowntimeParams":         {},
	//	"/joltify.third_party.dydxprotocol.blocktime.MsgUpdateDowntimeParamsResponse": {},
	//
	//	// bridge
	//	"/joltify.third_party.dydxprotocol.bridge.MsgAcknowledgeBridges":          {},
	//	"/joltify.third_party.dydxprotocol.bridge.MsgAcknowledgeBridgesResponse":  {},
	//	"/joltify.third_party.dydxprotocol.bridge.MsgCompleteBridge":              {},
	//	"/joltify.third_party.dydxprotocol.bridge.MsgCompleteBridgeResponse":      {},
	//	"/joltify.third_party.dydxprotocol.bridge.MsgUpdateEventParams":           {},
	//	"/joltify.third_party.dydxprotocol.bridge.MsgUpdateEventParamsResponse":   {},
	//	"/joltify.third_party.dydxprotocol.bridge.MsgUpdateProposeParams":         {},
	//	"/joltify.third_party.dydxprotocol.bridge.MsgUpdateProposeParamsResponse": {},
	//	"/joltify.third_party.dydxprotocol.bridge.MsgUpdateSafetyParams":          {},
	//	"/joltify.third_party.dydxprotocol.bridge.MsgUpdateSafetyParamsResponse":  {},
	//
	//	// clob
	//	"/joltify.third_party.dydxprotocol.clob.MsgBatchCancel":                                {},
	//	"/joltify.third_party.dydxprotocol.clob.MsgBatchCancelResponse":                        {},
	//	"/joltify.third_party.dydxprotocol.clob.MsgCancelOrder":                                {},
	//	"/joltify.third_party.dydxprotocol.clob.MsgCancelOrderResponse":                        {},
	//	"/joltify.third_party.dydxprotocol.clob.MsgCreateClobPair":                             {},
	//	"/joltify.third_party.dydxprotocol.clob.MsgCreateClobPairResponse":                     {},
	//	"/joltify.third_party.dydxprotocol.clob.MsgPlaceOrder":                                 {},
	//	"/joltify.third_party.dydxprotocol.clob.MsgPlaceOrderResponse":                         {},
	//	"/joltify.third_party.dydxprotocol.clob.MsgProposedOperations":                         {},
	//	"/joltify.third_party.dydxprotocol.clob.MsgProposedOperationsResponse":                 {},
	//	"/joltify.third_party.dydxprotocol.clob.MsgUpdateBlockRateLimitConfiguration":          {},
	//	"/joltify.third_party.dydxprotocol.clob.MsgUpdateBlockRateLimitConfigurationResponse":  {},
	//	"/joltify.third_party.dydxprotocol.clob.MsgUpdateClobPair":                             {},
	//	"/joltify.third_party.dydxprotocol.clob.MsgUpdateClobPairResponse":                     {},
	//	"/joltify.third_party.dydxprotocol.clob.MsgUpdateEquityTierLimitConfiguration":         {},
	//	"/joltify.third_party.dydxprotocol.clob.MsgUpdateEquityTierLimitConfigurationResponse": {},
	//	"/joltify.third_party.dydxprotocol.clob.MsgUpdateLiquidationsConfig":                   {},
	//	"/joltify.third_party.dydxprotocol.clob.MsgUpdateLiquidationsConfigResponse":           {},
	//
	//	// delaymsg
	//	"/joltify.third_party.dydxprotocol.delaymsg.MsgDelayMessage":         {},
	//	"/joltify.third_party.dydxprotocol.delaymsg.MsgDelayMessageResponse": {},
	//
	//	// feetiers
	//	"/joltify.third_party.dydxprotocol.feetiers.MsgUpdatePerpetualFeeParams":         {},
	//	"/joltify.third_party.dydxprotocol.feetiers.MsgUpdatePerpetualFeeParamsResponse": {},
	//
	//	// govplus
	//	"/joltify.third_party.dydxprotocol.govplus.MsgSlashValidator":         {},
	//	"/joltify.third_party.dydxprotocol.govplus.MsgSlashValidatorResponse": {},
	//
	//	// perpetuals
	//	"/joltify.third_party.dydxprotocol.perpetuals.MsgAddPremiumVotes":               {},
	//	"/joltify.third_party.dydxprotocol.perpetuals.MsgAddPremiumVotesResponse":       {},
	//	"/joltify.third_party.dydxprotocol.perpetuals.MsgCreatePerpetual":               {},
	//	"/joltify.third_party.dydxprotocol.perpetuals.MsgCreatePerpetualResponse":       {},
	//	"/joltify.third_party.dydxprotocol.perpetuals.MsgSetLiquidityTier":              {},
	//	"/joltify.third_party.dydxprotocol.perpetuals.MsgSetLiquidityTierResponse":      {},
	//	"/joltify.third_party.dydxprotocol.perpetuals.MsgUpdateParams":                  {},
	//	"/joltify.third_party.dydxprotocol.perpetuals.MsgUpdateParamsResponse":          {},
	//	"/joltify.third_party.dydxprotocol.perpetuals.MsgUpdatePerpetualParams":         {},
	//	"/joltify.third_party.dydxprotocol.perpetuals.MsgUpdatePerpetualParamsResponse": {},
	//
	//	// prices
	//	"/joltify.third_party.dydxprotocol.prices.MsgCreateOracleMarket":         {},
	//	"/joltify.third_party.dydxprotocol.prices.MsgCreateOracleMarketResponse": {},
	//	"/joltify.third_party.dydxprotocol.prices.MsgUpdateMarketPrices":         {},
	//	"/joltify.third_party.dydxprotocol.prices.MsgUpdateMarketPricesResponse": {},
	//	"/joltify.third_party.dydxprotocol.prices.MsgUpdateMarketParam":          {},
	//	"/joltify.third_party.dydxprotocol.prices.MsgUpdateMarketParamResponse":  {},
	//
	//	// ratelimit
	//	"/joltify.third_party.dydxprotocol.ratelimit.MsgSetLimitParams":         {},
	//	"/joltify.third_party.dydxprotocol.ratelimit.MsgSetLimitParamsResponse": {},
	//
	//	// sending
	//	"/joltify.third_party.dydxprotocol.sending.MsgCreateTransfer":                  {},
	//	"/joltify.third_party.dydxprotocol.sending.MsgCreateTransferResponse":          {},
	//	"/joltify.third_party.dydxprotocol.sending.MsgDepositToSubaccount":             {},
	//	"/joltify.third_party.dydxprotocol.sending.MsgDepositToSubaccountResponse":     {},
	//	"/joltify.third_party.dydxprotocol.sending.MsgWithdrawFromSubaccount":          {},
	//	"/joltify.third_party.dydxprotocol.sending.MsgWithdrawFromSubaccountResponse":  {},
	//	"/joltify.third_party.dydxprotocol.sending.MsgSendFromModuleToAccount":         {},
	//	"/joltify.third_party.dydxprotocol.sending.MsgSendFromModuleToAccountResponse": {},
	//
	//	// stats
	//	"/joltify.third_party.dydxprotocol.stats.MsgUpdateParams":         {},
	//	"/joltify.third_party.dydxprotocol.stats.MsgUpdateParamsResponse": {},
	//
	//	// vault
	//	"/joltify.third_party.dydxprotocol.vault.MsgDepositToVault":         {},
	//	"/joltify.third_party.dydxprotocol.vault.MsgDepositToVaultResponse": {},
	//	"/joltify.third_party.dydxprotocol.vault.MsgUpdateParams":           {},
	//	"/joltify.third_party.dydxprotocol.vault.MsgUpdateParamsResponse":   {},
	//
	//	// vest
	//	"/joltify.third_party.dydxprotocol.vest.MsgSetVestEntry":            {},
	//	"/joltify.third_party.dydxprotocol.vest.MsgSetVestEntryResponse":    {},
	//	"/joltify.third_party.dydxprotocol.vest.MsgDeleteVestEntry":         {},
	//	"/joltify.third_party.dydxprotocol.vest.MsgDeleteVestEntryResponse": {},
	//
	//	// rewards
	//	"/joltify.third_party.dydxprotocol.rewards.MsgUpdateParams":         {},
	//	"/joltify.third_party.dydxprotocol.rewards.MsgUpdateParamsResponse": {},
	//
	//	// ibc.applications
	//	"/ibc.applications.transfer.v1.MsgTransfer":             {},
	//	"/ibc.applications.transfer.v1.MsgTransferResponse":     {},
	//	"/ibc.applications.transfer.v1.MsgUpdateParams":         {},
	//	"/ibc.applications.transfer.v1.MsgUpdateParamsResponse": {},
	//	"/ibc.applications.transfer.v1.TransferAuthorization":   {},
	//
	//	// ibc.core.channel
	//	"/ibc.core.channel.v1.Channel":                        {},
	//	"/ibc.core.channel.v1.Counterparty":                   {},
	//	"/ibc.core.channel.v1.MsgAcknowledgement":             {},
	//	"/ibc.core.channel.v1.MsgAcknowledgementResponse":     {},
	//	"/ibc.core.channel.v1.MsgChannelCloseConfirm":         {},
	//	"/ibc.core.channel.v1.MsgChannelCloseConfirmResponse": {},
	//	"/ibc.core.channel.v1.MsgChannelCloseInit":            {},
	//	"/ibc.core.channel.v1.MsgChannelCloseInitResponse":    {},
	//	"/ibc.core.channel.v1.MsgChannelOpenAck":              {},
	//	"/ibc.core.channel.v1.MsgChannelOpenAckResponse":      {},
	//	"/ibc.core.channel.v1.MsgChannelOpenConfirm":          {},
	//	"/ibc.core.channel.v1.MsgChannelOpenConfirmResponse":  {},
	//	"/ibc.core.channel.v1.MsgChannelOpenInit":             {},
	//	"/ibc.core.channel.v1.MsgChannelOpenInitResponse":     {},
	//	"/ibc.core.channel.v1.MsgChannelOpenTry":              {},
	//	"/ibc.core.channel.v1.MsgChannelOpenTryResponse":      {},
	//	"/ibc.core.channel.v1.MsgRecvPacket":                  {},
	//	"/ibc.core.channel.v1.MsgRecvPacketResponse":          {},
	//	"/ibc.core.channel.v1.MsgTimeout":                     {},
	//	"/ibc.core.channel.v1.MsgTimeoutOnClose":              {},
	//	"/ibc.core.channel.v1.MsgTimeoutOnCloseResponse":      {},
	//	"/ibc.core.channel.v1.MsgTimeoutResponse":             {},
	//	"/ibc.core.channel.v1.Packet":                         {},
	//
	//	"/ibc.core.channel.v1.MsgChannelUpgradeAck":             {},
	//	"/ibc.core.channel.v1.MsgChannelUpgradeAckResponse":     {},
	//	"/ibc.core.channel.v1.MsgChannelUpgradeCancel":          {},
	//	"/ibc.core.channel.v1.MsgChannelUpgradeCancelResponse":  {},
	//	"/ibc.core.channel.v1.MsgChannelUpgradeConfirm":         {},
	//	"/ibc.core.channel.v1.MsgChannelUpgradeConfirmResponse": {},
	//	"/ibc.core.channel.v1.MsgChannelUpgradeInit":            {},
	//	"/ibc.core.channel.v1.MsgChannelUpgradeInitResponse":    {},
	//	"/ibc.core.channel.v1.MsgChannelUpgradeOpen":            {},
	//	"/ibc.core.channel.v1.MsgChannelUpgradeOpenResponse":    {},
	//	"/ibc.core.channel.v1.MsgChannelUpgradeTimeout":         {},
	//	"/ibc.core.channel.v1.MsgChannelUpgradeTimeoutResponse": {},
	//	"/ibc.core.channel.v1.MsgChannelUpgradeTry":             {},
	//	"/ibc.core.channel.v1.MsgChannelUpgradeTryResponse":     {},
	//	"/ibc.core.channel.v1.MsgPruneAcknowledgements":         {},
	//	"/ibc.core.channel.v1.MsgPruneAcknowledgementsResponse": {},
	//	"/ibc.core.channel.v1.MsgUpdateParams":                  {},
	//	"/ibc.core.channel.v1.MsgUpdateParamsResponse":          {},
	//
	//	"/ibc.lightclients.solomachine.v3.ClientState":    {},
	//	"/ibc.lightclients.solomachine.v3.ConsensusState": {},
	//	"/ibc.lightclients.solomachine.v3.Header":         {},
	//	"/ibc.lightclients.solomachine.v3.Misbehaviour":   {},
	//
	//	// ibc.core.client
	//	"/ibc.core.client.v1.ClientUpdateProposal":          {},
	//	"/ibc.core.client.v1.Height":                        {},
	//	"/ibc.core.client.v1.MsgCreateClient":               {},
	//	"/ibc.core.client.v1.MsgCreateClientResponse":       {},
	//	"/ibc.core.client.v1.MsgIBCSoftwareUpgrade":         {},
	//	"/ibc.core.client.v1.MsgIBCSoftwareUpgradeResponse": {},
	//	"/ibc.core.client.v1.MsgRecoverClient":              {},
	//	"/ibc.core.client.v1.MsgRecoverClientResponse":      {},
	//	"/ibc.core.client.v1.MsgSubmitMisbehaviour":         {},
	//	"/ibc.core.client.v1.MsgSubmitMisbehaviourResponse": {},
	//	"/ibc.core.client.v1.MsgUpdateClient":               {},
	//	"/ibc.core.client.v1.MsgUpdateClientResponse":       {},
	//	"/ibc.core.client.v1.MsgUpgradeClient":              {},
	//	"/ibc.core.client.v1.MsgUpgradeClientResponse":      {},
	//	"/ibc.core.client.v1.MsgUpdateParams":               {},
	//	"/ibc.core.client.v1.MsgUpdateParamsResponse":       {},
	//	"/ibc.core.client.v1.UpgradeProposal":               {},
	//
	//	// ibc.core.commitment
	//	"/ibc.core.commitment.v1.MerklePath":   {},
	//	"/ibc.core.commitment.v1.MerklePrefix": {},
	//	"/ibc.core.commitment.v1.MerkleProof":  {},
	//	"/ibc.core.commitment.v1.MerkleRoot":   {},
	//
	//	// ibc.core.connection
	//	"/ibc.core.connection.v1.ConnectionEnd":                    {},
	//	"/ibc.core.connection.v1.Counterparty":                     {},
	//	"/ibc.core.connection.v1.MsgConnectionOpenAck":             {},
	//	"/ibc.core.connection.v1.MsgConnectionOpenAckResponse":     {},
	//	"/ibc.core.connection.v1.MsgConnectionOpenConfirm":         {},
	//	"/ibc.core.connection.v1.MsgConnectionOpenConfirmResponse": {},
	//	"/ibc.core.connection.v1.MsgConnectionOpenInit":            {},
	//	"/ibc.core.connection.v1.MsgConnectionOpenInitResponse":    {},
	//	"/ibc.core.connection.v1.MsgConnectionOpenTry":             {},
	//	"/ibc.core.connection.v1.MsgConnectionOpenTryResponse":     {},
	//	"/ibc.core.connection.v1.MsgUpdateParams":                  {},
	//	"/ibc.core.connection.v1.MsgUpdateParamsResponse":          {},
	//
	//	// ibc.lightclients
	//	"/ibc.lightclients.localhost.v2.ClientState":     {},
	//	"/ibc.lightclients.tendermint.v1.ClientState":    {},
	//	"/ibc.lightclients.tendermint.v1.ConsensusState": {},
	//	"/ibc.lightclients.tendermint.v1.Header":         {},
	//	"/ibc.lightclients.tendermint.v1.Misbehaviour":   {},
	//
	//	// originial joltify messages
	//
	//	"/joltify.burnauction.MsgSubmitrequest":                    {},
	//	"/joltify.burnauction.MsgSubmitrequestResponse":            {},
	//	"/joltify.kyc.MSgCreateProjectResponse":                    {},
	//	"/joltify.kyc.MsgCreateProject":                            {},
	//	"/joltify.kyc.MsgUploadInvestor":                           {},
	//	"/joltify.kyc.MsgUploadInvestorResponse":                   {},
	//	"/joltify.spv.BorrowInterest":                              {},
	//	"/joltify.spv.MsgActivePool":                               {},
	//	"/joltify.spv.MsgActivePoolResponse":                       {},
	//	"/joltify.spv.MsgAddInvestors":                             {},
	//	"/joltify.spv.MsgAddInvestorsResponse":                     {},
	//	"/joltify.spv.MsgBorrow":                                   {},
	//	"/joltify.spv.MsgBorrowResponse":                           {},
	//	"/joltify.spv.MsgClaimInterest":                            {},
	//	"/joltify.spv.MsgClaimInterestResponse":                    {},
	//	"/joltify.spv.MsgCreatePool":                               {},
	//	"/joltify.spv.MsgCreatePoolResponse":                       {},
	//	"/joltify.spv.MsgDeposit":                                  {},
	//	"/joltify.spv.MsgDepositResponse":                          {},
	//	"/joltify.spv.MsgLiquidate":                                {},
	//	"/joltify.spv.MsgLiquidateResponse":                        {},
	//	"/joltify.spv.MsgPayPrincipal":                             {},
	//	"/joltify.spv.MsgPayPrincipalPartial":                      {},
	//	"/joltify.spv.MsgPayPrincipalPartialResponse":              {},
	//	"/joltify.spv.MsgPayPrincipalResponse":                     {},
	//	"/joltify.spv.MsgRepayInterest":                            {},
	//	"/joltify.spv.MsgRepayInterestResponse":                    {},
	//	"/joltify.spv.MsgSubmitWithdrawProposal":                   {},
	//	"/joltify.spv.MsgSubmitWithdrawProposalResponse":           {},
	//	"/joltify.spv.MsgTransferOwnership":                        {},
	//	"/joltify.spv.MsgTransferOwnershipResponse":                {},
	//	"/joltify.spv.MsgUpdatePool":                               {},
	//	"/joltify.spv.MsgUpdatePoolResponse":                       {},
	//	"/joltify.spv.MsgWithdrawPrincipal":                        {},
	//	"/joltify.spv.MsgWithdrawPrincipalResponse":                {},
	//	"/joltify.spv.NftInfo":                                     {},
	//	"/joltify.third_party.auction.v1beta1.CollateralAuction":   {},
	//	"/joltify.third_party.auction.v1beta1.DebtAuction":         {},
	//	"/joltify.third_party.auction.v1beta1.MsgPlaceBid":         {},
	//	"/joltify.third_party.auction.v1beta1.MsgPlaceBidResponse": {},
	//	"/joltify.third_party.auction.v1beta1.SurplusAuction":      {},
	//
	//	"/joltify.third_party.incentive.v1beta1.MsgClaimJoltReward":            {},
	//	"/joltify.third_party.incentive.v1beta1.MsgClaimJoltRewardResponse":    {},
	//	"/joltify.third_party.incentive.v1beta1.MsgClaimSPVReward":             {},
	//	"/joltify.third_party.incentive.v1beta1.MsgClaimSPVRewardResponse":     {},
	//	"/joltify.third_party.incentive.v1beta1.MsgClaimSwapReward":            {},
	//	"/joltify.third_party.incentive.v1beta1.MsgClaimSwapRewardResponse":    {},
	//	"/joltify.third_party.jolt.v1beta1.MsgBorrow":                          {},
	//	"/joltify.third_party.jolt.v1beta1.MsgBorrowResponse":                  {},
	//	"/joltify.third_party.jolt.v1beta1.MsgDeposit":                         {},
	//	"/joltify.third_party.jolt.v1beta1.MsgDepositResponse":                 {},
	//	"/joltify.third_party.jolt.v1beta1.MsgLiquidate":                       {},
	//	"/joltify.third_party.jolt.v1beta1.MsgLiquidateResponse":               {},
	//	"/joltify.third_party.jolt.v1beta1.MsgRepay":                           {},
	//	"/joltify.third_party.jolt.v1beta1.MsgRepayResponse":                   {},
	//	"/joltify.third_party.jolt.v1beta1.MsgWithdraw":                        {},
	//	"/joltify.third_party.jolt.v1beta1.MsgWithdrawResponse":                {},
	//	"/joltify.third_party.pricefeed.v1beta1.MsgPostPrice":                  {},
	//	"/joltify.third_party.pricefeed.v1beta1.MsgPostPriceResponse":          {},
	//	"/joltify.third_party.swap.v1beta1.MsgDeposit":                         {},
	//	"/joltify.third_party.swap.v1beta1.MsgDepositResponse":                 {},
	//	"/joltify.third_party.swap.v1beta1.MsgSwapExactForBatchTokens":         {},
	//	"/joltify.third_party.swap.v1beta1.MsgSwapExactForBatchTokensResponse": {},
	//	"/joltify.third_party.swap.v1beta1.MsgSwapExactForTokens":              {},
	//	"/joltify.third_party.swap.v1beta1.MsgSwapExactForTokensResponse":      {},
	//	"/joltify.third_party.swap.v1beta1.MsgSwapForExactTokens":              {},
	//	"/joltify.third_party.swap.v1beta1.MsgSwapForExactTokensResponse":      {},
	//	"/joltify.third_party.swap.v1beta1.MsgWithdraw":                        {},
	//	"/joltify.third_party.swap.v1beta1.MsgWithdrawResponse":                {},
	//
	//	// nft
	//	"/cosmos.nft.v1beta1.MsgSend":         {},
	//	"/cosmos.nft.v1beta1.MsgSendResponse": {},
	//
	//	// ica messages
	//	// Note: the `interchain_accounts.controller` messages are not actually used by the app,
	//	// since ICA Controller Keeper is initialized as nil.
	//	// However, since the ica.AppModuleBasic{} needs to be passed to basic_mananger as a whole, these messages
	//	// registered in the interface registry.
	//	//"/ibc.applications.interchain_accounts.v1.InterchainAccount":                               {},
	//	//"/ibc.applications.interchain_accounts.controller.v1.MsgSendTx":                            {},
	//	//"/ibc.applications.interchain_accounts.controller.v1.MsgSendTxResponse":                    {},
	//	//"/ibc.applications.interchain_accounts.controller.v1.MsgRegisterInterchainAccount":         {},
	//	//"/ibc.applications.interchain_accounts.controller.v1.MsgRegisterInterchainAccountResponse": {},
	//	//"/ibc.applications.interchain_accounts.controller.v1.MsgUpdateParams":                      {},
	//	//"/ibc.applications.interchain_accounts.controller.v1.MsgUpdateParamsResponse":              {},
	//	//"/ibc.applications.interchain_accounts.host.v1.MsgUpdateParams":                            {},
	//	//"/ibc.applications.interchain_accounts.host.v1.MsgUpdateParamsResponse":                    {},
	//}

	// DisallowMsgs are messages that cannot be externally submitted.
	DisallowMsgs = lib.MergeAllMapsMustHaveDistinctKeys(
		AppInjectedMsgSamples,
		InternalMsgSamplesAll,
		NestedMsgSamples,
		UnsupportedMsgSamples,
	)

	// AllowMsgs are messages that can be externally submitted.
	AllowMsgs = NormalMsgs
)
