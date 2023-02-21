// Generated by Ignite ignite.com/cli
import { Registry } from '@cosmjs/proto-signing'
import { IgniteClient } from "./client";
import { MissingWalletError } from "./helpers";
import { Module as CosmosAuthV1Beta1, msgTypes as CosmosAuthV1Beta1MsgTypes } from './cosmos.auth.v1beta1'
import { Module as CosmosAuthzV1Beta1, msgTypes as CosmosAuthzV1Beta1MsgTypes } from './cosmos.authz.v1beta1'
import { Module as CosmosBankV1Beta1, msgTypes as CosmosBankV1Beta1MsgTypes } from './cosmos.bank.v1beta1'
import { Module as CosmosBaseTendermintV1Beta1, msgTypes as CosmosBaseTendermintV1Beta1MsgTypes } from './cosmos.base.tendermint.v1beta1'
import { Module as CosmosCrisisV1Beta1, msgTypes as CosmosCrisisV1Beta1MsgTypes } from './cosmos.crisis.v1beta1'
import { Module as CosmosDistributionV1Beta1, msgTypes as CosmosDistributionV1Beta1MsgTypes } from './cosmos.distribution.v1beta1'
import { Module as CosmosEvidenceV1Beta1, msgTypes as CosmosEvidenceV1Beta1MsgTypes } from './cosmos.evidence.v1beta1'
import { Module as CosmosFeegrantV1Beta1, msgTypes as CosmosFeegrantV1Beta1MsgTypes } from './cosmos.feegrant.v1beta1'
import { Module as CosmosGovV1Beta1, msgTypes as CosmosGovV1Beta1MsgTypes } from './cosmos.gov.v1beta1'
import { Module as CosmosParamsV1Beta1, msgTypes as CosmosParamsV1Beta1MsgTypes } from './cosmos.params.v1beta1'
import { Module as CosmosSlashingV1Beta1, msgTypes as CosmosSlashingV1Beta1MsgTypes } from './cosmos.slashing.v1beta1'
import { Module as CosmosStakingV1Beta1, msgTypes as CosmosStakingV1Beta1MsgTypes } from './cosmos.staking.v1beta1'
import { Module as CosmosTxV1Beta1, msgTypes as CosmosTxV1Beta1MsgTypes } from './cosmos.tx.v1beta1'
import { Module as CosmosUpgradeV1Beta1, msgTypes as CosmosUpgradeV1Beta1MsgTypes } from './cosmos.upgrade.v1beta1'
import { Module as CosmosVestingV1Beta1, msgTypes as CosmosVestingV1Beta1MsgTypes } from './cosmos.vesting.v1beta1'
import { Module as JoltifyMint, msgTypes as JoltifyMintMsgTypes } from './joltify.mint'
import { Module as JoltifySpv, msgTypes as JoltifySpvMsgTypes } from './joltify.spv'
import { Module as JoltifyThirdPartyAuctionV1Beta1, msgTypes as JoltifyThirdPartyAuctionV1Beta1MsgTypes } from './joltify.third_party.auction.v1beta1'
import { Module as JoltifyThirdPartyCdpV1Beta1, msgTypes as JoltifyThirdPartyCdpV1Beta1MsgTypes } from './joltify.third_party.cdp.v1beta1'
import { Module as JoltifyThirdPartyIncentiveV1Beta1, msgTypes as JoltifyThirdPartyIncentiveV1Beta1MsgTypes } from './joltify.third_party.incentive.v1beta1'
import { Module as JoltifyThirdPartyIssuanceV1Beta1, msgTypes as JoltifyThirdPartyIssuanceV1Beta1MsgTypes } from './joltify.third_party.issuance.v1beta1'
import { Module as JoltifyThirdPartyJoltV1Beta1, msgTypes as JoltifyThirdPartyJoltV1Beta1MsgTypes } from './joltify.third_party.jolt.v1beta1'
import { Module as JoltifyThirdPartyPricefeedV1Beta1, msgTypes as JoltifyThirdPartyPricefeedV1Beta1MsgTypes } from './joltify.third_party.pricefeed.v1beta1'
import { Module as JoltifyVault, msgTypes as JoltifyVaultMsgTypes } from './joltify.vault'
import { Module as JoltifyfinanceJoltifyLendingKyc, msgTypes as JoltifyfinanceJoltifyLendingKycMsgTypes } from './joltifyfinance.joltify_lending.kyc'
import { Module as JoltifyfinanceJoltifyLendingSpv, msgTypes as JoltifyfinanceJoltifyLendingSpvMsgTypes } from './joltifyfinance.joltify_lending.spv'


const Client = IgniteClient.plugin([
    CosmosAuthV1Beta1, CosmosAuthzV1Beta1, CosmosBankV1Beta1, CosmosBaseTendermintV1Beta1, CosmosCrisisV1Beta1, CosmosDistributionV1Beta1, CosmosEvidenceV1Beta1, CosmosFeegrantV1Beta1, CosmosGovV1Beta1, CosmosParamsV1Beta1, CosmosSlashingV1Beta1, CosmosStakingV1Beta1, CosmosTxV1Beta1, CosmosUpgradeV1Beta1, CosmosVestingV1Beta1, JoltifyMint, JoltifySpv, JoltifyThirdPartyAuctionV1Beta1, JoltifyThirdPartyCdpV1Beta1, JoltifyThirdPartyIncentiveV1Beta1, JoltifyThirdPartyIssuanceV1Beta1, JoltifyThirdPartyJoltV1Beta1, JoltifyThirdPartyPricefeedV1Beta1, JoltifyVault, JoltifyfinanceJoltifyLendingKyc, JoltifyfinanceJoltifyLendingSpv
]);

const registry = new Registry([
  ...CosmosAuthV1Beta1MsgTypes,
  ...CosmosAuthzV1Beta1MsgTypes,
  ...CosmosBankV1Beta1MsgTypes,
  ...CosmosBaseTendermintV1Beta1MsgTypes,
  ...CosmosCrisisV1Beta1MsgTypes,
  ...CosmosDistributionV1Beta1MsgTypes,
  ...CosmosEvidenceV1Beta1MsgTypes,
  ...CosmosFeegrantV1Beta1MsgTypes,
  ...CosmosGovV1Beta1MsgTypes,
  ...CosmosParamsV1Beta1MsgTypes,
  ...CosmosSlashingV1Beta1MsgTypes,
  ...CosmosStakingV1Beta1MsgTypes,
  ...CosmosTxV1Beta1MsgTypes,
  ...CosmosUpgradeV1Beta1MsgTypes,
  ...CosmosVestingV1Beta1MsgTypes,
  ...JoltifyMintMsgTypes,
  ...JoltifySpvMsgTypes,
  ...JoltifyThirdPartyAuctionV1Beta1MsgTypes,
  ...JoltifyThirdPartyCdpV1Beta1MsgTypes,
  ...JoltifyThirdPartyIncentiveV1Beta1MsgTypes,
  ...JoltifyThirdPartyIssuanceV1Beta1MsgTypes,
  ...JoltifyThirdPartyJoltV1Beta1MsgTypes,
  ...JoltifyThirdPartyPricefeedV1Beta1MsgTypes,
  ...JoltifyVaultMsgTypes,
  ...JoltifyfinanceJoltifyLendingKycMsgTypes,
  ...JoltifyfinanceJoltifyLendingSpvMsgTypes,
  
])

export {
    Client,
    registry,
    MissingWalletError
}
