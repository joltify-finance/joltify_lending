import { Client, registry, MissingWalletError } from 'joltify-finance-joltify_lending-client-ts'

import { PoolProposal } from "joltify-finance-joltify_lending-client-ts/joltify.vault/types"
import { CreatePool } from "joltify-finance-joltify_lending-client-ts/joltify.vault/types"
import { IssueToken } from "joltify-finance-joltify_lending-client-ts/joltify.vault/types"
import { entity } from "joltify-finance-joltify_lending-client-ts/joltify.vault/types"
import { proposals } from "joltify-finance-joltify_lending-client-ts/joltify.vault/types"
import { OutboundTx } from "joltify-finance-joltify_lending-client-ts/joltify.vault/types"
import { addressV16 } from "joltify-finance-joltify_lending-client-ts/joltify.vault/types"
import { OutboundTxV16 } from "joltify-finance-joltify_lending-client-ts/joltify.vault/types"
import { poolInfo } from "joltify-finance-joltify_lending-client-ts/joltify.vault/types"
import { historicalAmount } from "joltify-finance-joltify_lending-client-ts/joltify.vault/types"
import { coinsQuota } from "joltify-finance-joltify_lending-client-ts/joltify.vault/types"
import { Params } from "joltify-finance-joltify_lending-client-ts/joltify.vault/types"
import { Validator } from "joltify-finance-joltify_lending-client-ts/joltify.vault/types"
import { StandbyPower } from "joltify-finance-joltify_lending-client-ts/joltify.vault/types"
import { Validators } from "joltify-finance-joltify_lending-client-ts/joltify.vault/types"


export { PoolProposal, CreatePool, IssueToken, entity, proposals, OutboundTx, addressV16, OutboundTxV16, poolInfo, historicalAmount, coinsQuota, Params, Validator, StandbyPower, Validators };

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
				OutboundTx: {},
				OutboundTxAll: {},
				GetValidators: {},
				GetAllValidators: {},
				GetQuota: {},
				IssueToken: {},
				IssueTokenAll: {},
				CreatePool: {},
				CreatePoolAll: {},
				GetLastPool: {},
				GetPendingFee: {},
				
				_Structure: {
						PoolProposal: getStructure(PoolProposal.fromPartial({})),
						CreatePool: getStructure(CreatePool.fromPartial({})),
						IssueToken: getStructure(IssueToken.fromPartial({})),
						entity: getStructure(entity.fromPartial({})),
						proposals: getStructure(proposals.fromPartial({})),
						OutboundTx: getStructure(OutboundTx.fromPartial({})),
						addressV16: getStructure(addressV16.fromPartial({})),
						OutboundTxV16: getStructure(OutboundTxV16.fromPartial({})),
						poolInfo: getStructure(poolInfo.fromPartial({})),
						historicalAmount: getStructure(historicalAmount.fromPartial({})),
						coinsQuota: getStructure(coinsQuota.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						Validator: getStructure(Validator.fromPartial({})),
						StandbyPower: getStructure(StandbyPower.fromPartial({})),
						Validators: getStructure(Validators.fromPartial({})),
						
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
				getOutboundTx: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.OutboundTx[JSON.stringify(params)] ?? {}
		},
				getOutboundTxAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.OutboundTxAll[JSON.stringify(params)] ?? {}
		},
				getGetValidators: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetValidators[JSON.stringify(params)] ?? {}
		},
				getGetAllValidators: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetAllValidators[JSON.stringify(params)] ?? {}
		},
				getGetQuota: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetQuota[JSON.stringify(params)] ?? {}
		},
				getIssueToken: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.IssueToken[JSON.stringify(params)] ?? {}
		},
				getIssueTokenAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.IssueTokenAll[JSON.stringify(params)] ?? {}
		},
				getCreatePool: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.CreatePool[JSON.stringify(params)] ?? {}
		},
				getCreatePoolAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.CreatePoolAll[JSON.stringify(params)] ?? {}
		},
				getGetLastPool: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetLastPool[JSON.stringify(params)] ?? {}
		},
				getGetPendingFee: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetPendingFee[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: joltify.vault initialized!')
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
		
		
		
		 		
		
		
		async QueryOutboundTx({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyVault.query.queryOutboundTx( key.requestID)).data
				
					
				commit('QUERY', { query: 'OutboundTx', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryOutboundTx', payload: { options: { all }, params: {...key},query }})
				return getters['getOutboundTx']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryOutboundTx API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryOutboundTxAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyVault.query.queryOutboundTxAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JoltifyVault.query.queryOutboundTxAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'OutboundTxAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryOutboundTxAll', payload: { options: { all }, params: {...key},query }})
				return getters['getOutboundTxAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryOutboundTxAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetValidators({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyVault.query.queryGetValidators( key.height)).data
				
					
				commit('QUERY', { query: 'GetValidators', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetValidators', payload: { options: { all }, params: {...key},query }})
				return getters['getGetValidators']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryGetValidators API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetAllValidators({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyVault.query.queryGetAllValidators(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JoltifyVault.query.queryGetAllValidators({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'GetAllValidators', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetAllValidators', payload: { options: { all }, params: {...key},query }})
				return getters['getGetAllValidators']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryGetAllValidators API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetQuota({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyVault.query.queryGetQuota( key.query_length, query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JoltifyVault.query.queryGetQuota( key.query_length, {...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'GetQuota', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetQuota', payload: { options: { all }, params: {...key},query }})
				return getters['getGetQuota']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryGetQuota API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryIssueToken({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyVault.query.queryIssueToken( key.index)).data
				
					
				commit('QUERY', { query: 'IssueToken', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryIssueToken', payload: { options: { all }, params: {...key},query }})
				return getters['getIssueToken']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryIssueToken API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryIssueTokenAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyVault.query.queryIssueTokenAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JoltifyVault.query.queryIssueTokenAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'IssueTokenAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryIssueTokenAll', payload: { options: { all }, params: {...key},query }})
				return getters['getIssueTokenAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryIssueTokenAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCreatePool({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyVault.query.queryCreatePool( key.index)).data
				
					
				commit('QUERY', { query: 'CreatePool', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCreatePool', payload: { options: { all }, params: {...key},query }})
				return getters['getCreatePool']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCreatePool API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCreatePoolAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyVault.query.queryCreatePoolAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JoltifyVault.query.queryCreatePoolAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'CreatePoolAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCreatePoolAll', payload: { options: { all }, params: {...key},query }})
				return getters['getCreatePoolAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCreatePoolAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetLastPool({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyVault.query.queryGetLastPool(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JoltifyVault.query.queryGetLastPool({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'GetLastPool', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetLastPool', payload: { options: { all }, params: {...key},query }})
				return getters['getGetLastPool']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryGetLastPool API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetPendingFee({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyVault.query.queryGetPendingFee(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JoltifyVault.query.queryGetPendingFee({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'GetPendingFee', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetPendingFee', payload: { options: { all }, params: {...key},query }})
				return getters['getGetPendingFee']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryGetPendingFee API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgCreateCreatePool({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyVault.tx.sendMsgCreateCreatePool({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateCreatePool:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateCreatePool:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateOutboundTx({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyVault.tx.sendMsgCreateOutboundTx({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateOutboundTx:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateOutboundTx:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateIssueToken({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyVault.tx.sendMsgCreateIssueToken({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateIssueToken:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateIssueToken:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgCreateCreatePool({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyVault.tx.msgCreateCreatePool({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateCreatePool:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateCreatePool:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateOutboundTx({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyVault.tx.msgCreateOutboundTx({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateOutboundTx:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateOutboundTx:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateIssueToken({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyVault.tx.msgCreateIssueToken({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateIssueToken:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateIssueToken:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
