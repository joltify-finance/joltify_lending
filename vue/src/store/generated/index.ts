// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import JoltifyfinanceJoltifyLendingKyc from './joltifyfinance.joltify_lending.kyc'
import JoltifyMint from './joltify.mint'
import JoltifySpv from './joltify.spv'
import JoltifyThirdPartyAuctionV1Beta1 from './joltify.third_party.auction.v1beta1'
import JoltifyThirdPartyCdpV1Beta1 from './joltify.third_party.cdp.v1beta1'
import JoltifyThirdPartyIncentiveV1Beta1 from './joltify.third_party.incentive.v1beta1'
import JoltifyThirdPartyIssuanceV1Beta1 from './joltify.third_party.issuance.v1beta1'
import JoltifyThirdPartyJoltV1Beta1 from './joltify.third_party.jolt.v1beta1'
import JoltifyThirdPartyPricefeedV1Beta1 from './joltify.third_party.pricefeed.v1beta1'
import JoltifyVault from './joltify.vault'
import JoltifyfinanceJoltifyLendingSpv from './joltifyfinance.joltify_lending.spv'
import CosmosAuthV1Beta1 from './cosmos.auth.v1beta1'
import CosmosAuthzV1Beta1 from './cosmos.authz.v1beta1'
import CosmosBankV1Beta1 from './cosmos.bank.v1beta1'
import CosmosBaseTendermintV1Beta1 from './cosmos.base.tendermint.v1beta1'
import CosmosCrisisV1Beta1 from './cosmos.crisis.v1beta1'
import CosmosDistributionV1Beta1 from './cosmos.distribution.v1beta1'
import CosmosEvidenceV1Beta1 from './cosmos.evidence.v1beta1'
import CosmosFeegrantV1Beta1 from './cosmos.feegrant.v1beta1'
import CosmosGovV1Beta1 from './cosmos.gov.v1beta1'
import CosmosParamsV1Beta1 from './cosmos.params.v1beta1'
import CosmosSlashingV1Beta1 from './cosmos.slashing.v1beta1'
import CosmosStakingV1Beta1 from './cosmos.staking.v1beta1'
import CosmosTxV1Beta1 from './cosmos.tx.v1beta1'
import CosmosUpgradeV1Beta1 from './cosmos.upgrade.v1beta1'
import CosmosVestingV1Beta1 from './cosmos.vesting.v1beta1'


export default { 
  JoltifyfinanceJoltifyLendingKyc: load(JoltifyfinanceJoltifyLendingKyc, 'joltifyfinance.joltify_lending.kyc'),
  JoltifyMint: load(JoltifyMint, 'joltify.mint'),
  JoltifySpv: load(JoltifySpv, 'joltify.spv'),
  JoltifyThirdPartyAuctionV1Beta1: load(JoltifyThirdPartyAuctionV1Beta1, 'joltify.third_party.auction.v1beta1'),
  JoltifyThirdPartyCdpV1Beta1: load(JoltifyThirdPartyCdpV1Beta1, 'joltify.third_party.cdp.v1beta1'),
  JoltifyThirdPartyIncentiveV1Beta1: load(JoltifyThirdPartyIncentiveV1Beta1, 'joltify.third_party.incentive.v1beta1'),
  JoltifyThirdPartyIssuanceV1Beta1: load(JoltifyThirdPartyIssuanceV1Beta1, 'joltify.third_party.issuance.v1beta1'),
  JoltifyThirdPartyJoltV1Beta1: load(JoltifyThirdPartyJoltV1Beta1, 'joltify.third_party.jolt.v1beta1'),
  JoltifyThirdPartyPricefeedV1Beta1: load(JoltifyThirdPartyPricefeedV1Beta1, 'joltify.third_party.pricefeed.v1beta1'),
  JoltifyVault: load(JoltifyVault, 'joltify.vault'),
  JoltifyfinanceJoltifyLendingSpv: load(JoltifyfinanceJoltifyLendingSpv, 'joltifyfinance.joltify_lending.spv'),
  CosmosAuthV1Beta1: load(CosmosAuthV1Beta1, 'cosmos.auth.v1beta1'),
  CosmosAuthzV1Beta1: load(CosmosAuthzV1Beta1, 'cosmos.authz.v1beta1'),
  CosmosBankV1Beta1: load(CosmosBankV1Beta1, 'cosmos.bank.v1beta1'),
  CosmosBaseTendermintV1Beta1: load(CosmosBaseTendermintV1Beta1, 'cosmos.base.tendermint.v1beta1'),
  CosmosCrisisV1Beta1: load(CosmosCrisisV1Beta1, 'cosmos.crisis.v1beta1'),
  CosmosDistributionV1Beta1: load(CosmosDistributionV1Beta1, 'cosmos.distribution.v1beta1'),
  CosmosEvidenceV1Beta1: load(CosmosEvidenceV1Beta1, 'cosmos.evidence.v1beta1'),
  CosmosFeegrantV1Beta1: load(CosmosFeegrantV1Beta1, 'cosmos.feegrant.v1beta1'),
  CosmosGovV1Beta1: load(CosmosGovV1Beta1, 'cosmos.gov.v1beta1'),
  CosmosParamsV1Beta1: load(CosmosParamsV1Beta1, 'cosmos.params.v1beta1'),
  CosmosSlashingV1Beta1: load(CosmosSlashingV1Beta1, 'cosmos.slashing.v1beta1'),
  CosmosStakingV1Beta1: load(CosmosStakingV1Beta1, 'cosmos.staking.v1beta1'),
  CosmosTxV1Beta1: load(CosmosTxV1Beta1, 'cosmos.tx.v1beta1'),
  CosmosUpgradeV1Beta1: load(CosmosUpgradeV1Beta1, 'cosmos.upgrade.v1beta1'),
  CosmosVestingV1Beta1: load(CosmosVestingV1Beta1, 'cosmos.vesting.v1beta1'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}