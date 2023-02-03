import { Client, registry, MissingWalletError } from 'joltify-finance-joltify_lending-client-ts'

import { Params } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.issuance.v1beta1/types"
import { Asset } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.issuance.v1beta1/types"
import { RateLimit } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.issuance.v1beta1/types"
import { AssetSupply } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.issuance.v1beta1/types"


export { Params, Asset, RateLimit, AssetSupply };

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
				Params: {},
				
				_Structure: {
						Params: getStructure(Params.fromPartial({})),
						Asset: getStructure(Asset.fromPartial({})),
						RateLimit: getStructure(RateLimit.fromPartial({})),
						AssetSupply: getStructure(AssetSupply.fromPartial({})),
						
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
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: joltify.third_party.issuance.v1beta1 initialized!')
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
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyThirdPartyIssuanceV1Beta1.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgBlockAddress({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyThirdPartyIssuanceV1Beta1.tx.sendMsgBlockAddress({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBlockAddress:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgBlockAddress:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSetPauseStatus({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyThirdPartyIssuanceV1Beta1.tx.sendMsgSetPauseStatus({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetPauseStatus:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSetPauseStatus:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRedeemTokens({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyThirdPartyIssuanceV1Beta1.tx.sendMsgRedeemTokens({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRedeemTokens:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRedeemTokens:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUnblockAddress({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyThirdPartyIssuanceV1Beta1.tx.sendMsgUnblockAddress({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUnblockAddress:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUnblockAddress:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgIssueTokens({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyThirdPartyIssuanceV1Beta1.tx.sendMsgIssueTokens({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgIssueTokens:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgIssueTokens:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgBlockAddress({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyThirdPartyIssuanceV1Beta1.tx.msgBlockAddress({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBlockAddress:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgBlockAddress:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSetPauseStatus({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyThirdPartyIssuanceV1Beta1.tx.msgSetPauseStatus({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetPauseStatus:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSetPauseStatus:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRedeemTokens({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyThirdPartyIssuanceV1Beta1.tx.msgRedeemTokens({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRedeemTokens:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRedeemTokens:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUnblockAddress({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyThirdPartyIssuanceV1Beta1.tx.msgUnblockAddress({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUnblockAddress:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUnblockAddress:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgIssueTokens({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyThirdPartyIssuanceV1Beta1.tx.msgIssueTokens({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgIssueTokens:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgIssueTokens:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
