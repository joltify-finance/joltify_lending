package msgs

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govbeta "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	icacontrollertypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/types"
)

// UnsupportedMsgSamples are msgs that are registered with the app, but are not supported.
var UnsupportedMsgSamples = map[string]sdk.Msg{
	// gov
	// MsgCancelProposal is not allowed by protocol, due to it's potential for abuse.
	"/cosmos.gov.v1.MsgCancelProposal":         &gov.MsgCancelProposal{},
	"/cosmos.gov.v1.MsgCancelProposalResponse": nil,
	// These are deprecated/legacy msgs that we should not support.
	"/cosmos.gov.v1beta1.MsgSubmitProposal":         &govbeta.MsgSubmitProposal{},
	"/cosmos.gov.v1beta1.MsgSubmitProposalResponse": nil,

	// ICA Controller messages - these are not used since ICA Controller is disabled.
	"/ibc.applications.interchain_accounts.controller.v1.MsgRegisterInterchainAccount": &icacontrollertypes.
		MsgRegisterInterchainAccount{},
	"/ibc.applications.interchain_accounts.controller.v1.MsgRegisterInterchainAccountResponse": nil,
	"/ibc.applications.interchain_accounts.controller.v1.MsgSendTx": &icacontrollertypes.
		MsgSendTx{},
	"/ibc.applications.interchain_accounts.controller.v1.MsgSendTxResponse": nil,
	"/ibc.applications.interchain_accounts.controller.v1.MsgUpdateParams": &icacontrollertypes.
		MsgUpdateParams{},
	"/ibc.applications.interchain_accounts.controller.v1.MsgUpdateParamsResponse": nil,
	"/ibc.core.channel.v1.MsgUpdateParams":                                        nil,
	"/ibc.core.channel.v1.MsgUpdateParamsResponse":                                nil,
}

var OtherMsg = map[string]sdk.Msg{
	"/ibc.lightclients.solomachine.v3.ClientState":    nil,
	"/ibc.lightclients.solomachine.v3.ConsensusState": nil,
	"/ibc.lightclients.solomachine.v3.Header":         nil,
	"/ibc.lightclients.solomachine.v3.Misbehaviour":   nil,

	"/ibc.core.channel.v1.MsgChannelUpgradeAck":             nil,
	"/ibc.core.channel.v1.MsgChannelUpgradeAckResponse":     nil,
	"/ibc.core.channel.v1.MsgChannelUpgradeCancel":          nil,
	"/ibc.core.channel.v1.MsgChannelUpgradeCancelResponse":  nil,
	"/ibc.core.channel.v1.MsgChannelUpgradeConfirm":         nil,
	"/ibc.core.channel.v1.MsgChannelUpgradeConfirmResponse": nil,
	"/ibc.core.channel.v1.MsgChannelUpgradeInit	":           nil,
	"/ibc.core.channel.v1.MsgChannelUpgradeInitResponse":    nil,
	"/ibc.core.channel.v1.MsgChannelUpgradeOpen":            nil,
	"/ibc.core.channel.v1.MsgChannelUpgradeOpenResponse":    nil,
	"/ibc.core.channel.v1.MsgChannelUpgradeTimeout":         nil,
	"/ibc.core.channel.v1.MsgChannelUpgradeTimeoutResponse": nil,
	"/ibc.core.channel.v1.MsgChannelUpgradeTry":             nil,
	"/ibc.core.channel.v1.MsgChannelUpgradeTryResponse":     nil,
	"/ibc.core.channel.v1.MsgPruneAcknowledgements":         nil,
	"/ibc.core.channel.v1.MsgPruneAcknowledgementsResponse": nil,
}

var JoltMsg = map[string]sdk.Msg{
	"/joltify.third_party.incentive.v1beta1.MsgClaimJoltReward":            nil,
	"/joltify.third_party.incentive.v1beta1.MsgClaimJoltRewardResponse":    nil,
	"/joltify.third_party.incentive.v1beta1.MsgClaimSPVReward":             nil,
	"/joltify.third_party.incentive.v1beta1.MsgClaimSPVRewardResponse":     nil,
	"/joltify.third_party.incentive.v1beta1.MsgClaimSwapReward":            nil,
	"/joltify.third_party.incentive.v1beta1.MsgClaimSwapRewardResponse":    nil,
	"/joltify.third_party.jolt.v1beta1.MsgBorrow":                          nil,
	"/joltify.third_party.jolt.v1beta1.MsgBorrowResponse":                  nil,
	"/joltify.third_party.jolt.v1beta1.MsgDeposit":                         nil,
	"/joltify.third_party.jolt.v1beta1.MsgDepositResponse":                 nil,
	"/joltify.third_party.jolt.v1beta1.MsgLiquidate":                       nil,
	"/joltify.third_party.jolt.v1beta1.MsgLiquidateResponse":               nil,
	"/joltify.third_party.jolt.v1beta1.MsgRepay":                           nil,
	"/joltify.third_party.jolt.v1beta1.MsgRepayResponse":                   nil,
	"/joltify.third_party.jolt.v1beta1.MsgWithdraw":                        nil,
	"/joltify.third_party.jolt.v1beta1.MsgWithdrawResponse":                nil,
	"/joltify.third_party.pricefeed.v1beta1.MsgPostPrice":                  nil,
	"/joltify.third_party.pricefeed.v1beta1.MsgPostPriceResponse":          nil,
	"/joltify.third_party.swap.v1beta1.MsgDeposit":                         nil,
	"/joltify.third_party.swap.v1beta1.MsgDepositResponse":                 nil,
	"/joltify.third_party.swap.v1beta1.MsgSwapExactForBatchTokens":         nil,
	"/joltify.third_party.swap.v1beta1.MsgSwapExactForBatchTokensResponse": nil,
	"/joltify.third_party.swap.v1beta1.MsgSwapExactForTokens":              nil,
	"/joltify.third_party.swap.v1beta1.MsgSwapExactForTokensResponse":      nil,
	"/joltify.third_party.swap.v1beta1.MsgSwapForExactTokens":              nil,
	"/joltify.third_party.swap.v1beta1.MsgSwapForExactTokensResponse":      nil,
	"/joltify.third_party.swap.v1beta1.MsgWithdraw":                        nil,
	"/joltify.third_party.swap.v1beta1.MsgWithdrawResponse":                nil,

	"/joltify.burnauction.MsgSubmitrequest":                    nil,
	"/joltify.burnauction.MsgSubmitrequestResponse":            nil,
	"/joltify.kyc.MSgCreateProjectResponse":                    nil,
	"/joltify.kyc.MsgCreateProject":                            nil,
	"/joltify.kyc.MsgUploadInvestor":                           nil,
	"/joltify.kyc.MsgUploadInvestorResponse":                   nil,
	"/joltify.spv.BorrowInterest":                              nil,
	"/joltify.spv.MsgActivePool":                               nil,
	"/joltify.spv.MsgActivePoolResponse":                       nil,
	"/joltify.spv.MsgAddInvestors":                             nil,
	"/joltify.spv.MsgAddInvestorsResponse":                     nil,
	"/joltify.spv.MsgBorrow":                                   nil,
	"/joltify.spv.MsgBorrowResponse":                           nil,
	"/joltify.spv.MsgClaimInterest":                            nil,
	"/joltify.spv.MsgClaimInterestResponse":                    nil,
	"/joltify.spv.MsgCreatePool":                               nil,
	"/joltify.spv.MsgCreatePoolResponse":                       nil,
	"/joltify.spv.MsgDeposit":                                  nil,
	"/joltify.spv.MsgDepositResponse":                          nil,
	"/joltify.spv.MsgLiquidate":                                nil,
	"/joltify.spv.MsgLiquidateResponse":                        nil,
	"/joltify.spv.MsgPayPrincipal":                             nil,
	"/joltify.spv.MsgPayPrincipalPartial":                      nil,
	"/joltify.spv.MsgPayPrincipalPartialResponse":              nil,
	"/joltify.spv.MsgPayPrincipalResponse":                     nil,
	"/joltify.spv.MsgRepayInterest":                            nil,
	"/joltify.spv.MsgRepayInterestResponse":                    nil,
	"/joltify.spv.MsgSubmitWithdrawProposal":                   nil,
	"/joltify.spv.MsgSubmitWithdrawProposalResponse":           nil,
	"/joltify.spv.MsgTransferOwnership":                        nil,
	"/joltify.spv.MsgTransferOwnershipResponse":                nil,
	"/joltify.spv.MsgUpdatePool":                               nil,
	"/joltify.spv.MsgUpdatePoolResponse":                       nil,
	"/joltify.spv.MsgWithdrawPrincipal":                        nil,
	"/joltify.spv.MsgWithdrawPrincipalResponse":                nil,
	"/joltify.spv.NftInfo":                                     nil,
	"/joltify.third_party.auction.v1beta1.CollateralAuction":   nil,
	"/joltify.third_party.auction.v1beta1.DebtAuction":         nil,
	"/joltify.third_party.auction.v1beta1.MsgPlaceBid":         nil,
	"/joltify.third_party.auction.v1beta1.MsgPlaceBidResponse": nil,
	"/joltify.third_party.auction.v1beta1.SurplusAuction":      nil,
}
