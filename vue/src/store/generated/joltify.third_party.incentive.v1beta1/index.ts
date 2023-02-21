import { Client, registry, MissingWalletError } from 'joltify-finance-joltify_lending-client-ts'

import { BaseClaim } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.incentive.v1beta1/types"
import { BaseMultiClaim } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.incentive.v1beta1/types"
import { RewardIndex } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.incentive.v1beta1/types"
import { RewardIndexesProto } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.incentive.v1beta1/types"
import { MultiRewardIndex } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.incentive.v1beta1/types"
import { MultiRewardIndexesProto } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.incentive.v1beta1/types"
import { USDXMintingClaim } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.incentive.v1beta1/types"
import { JoltLiquidityProviderClaim } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.incentive.v1beta1/types"
import { DelegatorClaim } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.incentive.v1beta1/types"
import { SwapClaim } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.incentive.v1beta1/types"
import { SavingsClaim } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.incentive.v1beta1/types"
import { AccumulationTime } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.incentive.v1beta1/types"
import { GenesisRewardState } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.incentive.v1beta1/types"
import { RewardPeriod } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.incentive.v1beta1/types"
import { MultiRewardPeriod } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.incentive.v1beta1/types"
import { Multiplier } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.incentive.v1beta1/types"
import { MultipliersPerDenom } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.incentive.v1beta1/types"
import { Params } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.incentive.v1beta1/types"
import { Selection } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.incentive.v1beta1/types"
import { MsgClaimUSDXMintingRewardResponse } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.incentive.v1beta1/types"
import { MsgClaimDelegatorRewardResponse } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.incentive.v1beta1/types"
import { MsgClaimSwapRewardResponse } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.incentive.v1beta1/types"
import { MsgClaimSavingsRewardResponse } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.incentive.v1beta1/types"


export { BaseClaim, BaseMultiClaim, RewardIndex, RewardIndexesProto, MultiRewardIndex, MultiRewardIndexesProto, USDXMintingClaim, JoltLiquidityProviderClaim, DelegatorClaim, SwapClaim, SavingsClaim, AccumulationTime, GenesisRewardState, RewardPeriod, MultiRewardPeriod, Multiplier, MultipliersPerDenom, Params, Selection, MsgClaimUSDXMintingRewardResponse, MsgClaimDelegatorRewardResponse, MsgClaimSwapRewardResponse, MsgClaimSavingsRewardResponse };

function initClient(vuexGetters) {
	return new Client(vuexGetters['common/env/getEnv'], vuexGetters['common/wallet/signer'])
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	let structure: {fields: Field[]} = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const getDefaultState = () => {
	return {
				
				_Structure: {
						BaseClaim: getStructure(BaseClaim.fromPartial({})),
						BaseMultiClaim: getStructure(BaseMultiClaim.fromPartial({})),
						RewardIndex: getStructure(RewardIndex.fromPartial({})),
						RewardIndexesProto: getStructure(RewardIndexesProto.fromPartial({})),
						MultiRewardIndex: getStructure(MultiRewardIndex.fromPartial({})),
						MultiRewardIndexesProto: getStructure(MultiRewardIndexesProto.fromPartial({})),
						USDXMintingClaim: getStructure(USDXMintingClaim.fromPartial({})),
						JoltLiquidityProviderClaim: getStructure(JoltLiquidityProviderClaim.fromPartial({})),
						DelegatorClaim: getStructure(DelegatorClaim.fromPartial({})),
						SwapClaim: getStructure(SwapClaim.fromPartial({})),
						SavingsClaim: getStructure(SavingsClaim.fromPartial({})),
						AccumulationTime: getStructure(AccumulationTime.fromPartial({})),
						GenesisRewardState: getStructure(GenesisRewardState.fromPartial({})),
						RewardPeriod: getStructure(RewardPeriod.fromPartial({})),
						MultiRewardPeriod: getStructure(MultiRewardPeriod.fromPartial({})),
						Multiplier: getStructure(Multiplier.fromPartial({})),
						MultipliersPerDenom: getStructure(MultipliersPerDenom.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						Selection: getStructure(Selection.fromPartial({})),
						MsgClaimUSDXMintingRewardResponse: getStructure(MsgClaimUSDXMintingRewardResponse.fromPartial({})),
						MsgClaimDelegatorRewardResponse: getStructure(MsgClaimDelegatorRewardResponse.fromPartial({})),
						MsgClaimSwapRewardResponse: getStructure(MsgClaimSwapRewardResponse.fromPartial({})),
						MsgClaimSavingsRewardResponse: getStructure(MsgClaimSavingsRewardResponse.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: joltify.third_party.incentive.v1beta1 initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		async sendMsgClaimUSDXMintingReward({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyThirdPartyIncentiveV1Beta1.tx.sendMsgClaimUSDXMintingReward({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgClaimUSDXMintingReward:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgClaimUSDXMintingReward:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgClaimSavingsReward({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyThirdPartyIncentiveV1Beta1.tx.sendMsgClaimSavingsReward({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgClaimSavingsReward:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgClaimSavingsReward:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgClaimJoltReward({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyThirdPartyIncentiveV1Beta1.tx.sendMsgClaimJoltReward({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgClaimJoltReward:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgClaimJoltReward:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgClaimSwapReward({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyThirdPartyIncentiveV1Beta1.tx.sendMsgClaimSwapReward({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgClaimSwapReward:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgClaimSwapReward:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgClaimDelegatorReward({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyThirdPartyIncentiveV1Beta1.tx.sendMsgClaimDelegatorReward({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgClaimDelegatorReward:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgClaimDelegatorReward:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgClaimUSDXMintingReward({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyThirdPartyIncentiveV1Beta1.tx.msgClaimUSDXMintingReward({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgClaimUSDXMintingReward:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgClaimUSDXMintingReward:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgClaimSavingsReward({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyThirdPartyIncentiveV1Beta1.tx.msgClaimSavingsReward({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgClaimSavingsReward:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgClaimSavingsReward:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgClaimJoltReward({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyThirdPartyIncentiveV1Beta1.tx.msgClaimJoltReward({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgClaimJoltReward:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgClaimJoltReward:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgClaimSwapReward({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyThirdPartyIncentiveV1Beta1.tx.msgClaimSwapReward({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgClaimSwapReward:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgClaimSwapReward:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgClaimDelegatorReward({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyThirdPartyIncentiveV1Beta1.tx.msgClaimDelegatorReward({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgClaimDelegatorReward:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgClaimDelegatorReward:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
