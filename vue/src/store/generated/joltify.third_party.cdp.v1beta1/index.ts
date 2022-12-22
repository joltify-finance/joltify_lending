import { Client, registry, MissingWalletError } from 'joltify-finance-joltify_lending-client-ts'

import { CDP } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.cdp.v1beta1/types"
import { Deposit } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.cdp.v1beta1/types"
import { TotalPrincipal } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.cdp.v1beta1/types"
import { TotalCollateral } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.cdp.v1beta1/types"
import { OwnerCDPIndex } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.cdp.v1beta1/types"
import { Params } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.cdp.v1beta1/types"
import { DebtParam } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.cdp.v1beta1/types"
import { CollateralParam } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.cdp.v1beta1/types"
import { GenesisAccumulationTime } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.cdp.v1beta1/types"
import { GenesisTotalPrincipal } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.cdp.v1beta1/types"
import { CDPResponse } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.cdp.v1beta1/types"


export { CDP, Deposit, TotalPrincipal, TotalCollateral, OwnerCDPIndex, Params, DebtParam, CollateralParam, GenesisAccumulationTime, GenesisTotalPrincipal, CDPResponse };

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
				Accounts: {},
				TotalPrincipal: {},
				TotalCollateral: {},
				Cdps: {},
				Cdp: {},
				Deposits: {},
				
				_Structure: {
						CDP: getStructure(CDP.fromPartial({})),
						Deposit: getStructure(Deposit.fromPartial({})),
						TotalPrincipal: getStructure(TotalPrincipal.fromPartial({})),
						TotalCollateral: getStructure(TotalCollateral.fromPartial({})),
						OwnerCDPIndex: getStructure(OwnerCDPIndex.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						DebtParam: getStructure(DebtParam.fromPartial({})),
						CollateralParam: getStructure(CollateralParam.fromPartial({})),
						GenesisAccumulationTime: getStructure(GenesisAccumulationTime.fromPartial({})),
						GenesisTotalPrincipal: getStructure(GenesisTotalPrincipal.fromPartial({})),
						CDPResponse: getStructure(CDPResponse.fromPartial({})),
						
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
				getAccounts: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Accounts[JSON.stringify(params)] ?? {}
		},
				getTotalPrincipal: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TotalPrincipal[JSON.stringify(params)] ?? {}
		},
				getTotalCollateral: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TotalCollateral[JSON.stringify(params)] ?? {}
		},
				getCdps: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Cdps[JSON.stringify(params)] ?? {}
		},
				getCdp: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Cdp[JSON.stringify(params)] ?? {}
		},
				getDeposits: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Deposits[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: joltify.third_party.cdp.v1beta1 initialized!')
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
				let value= (await client.JoltifyThirdPartyCdpV1Beta1.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryAccounts({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyThirdPartyCdpV1Beta1.query.queryAccounts()).data
				
					
				commit('QUERY', { query: 'Accounts', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAccounts', payload: { options: { all }, params: {...key},query }})
				return getters['getAccounts']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAccounts API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTotalPrincipal({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyThirdPartyCdpV1Beta1.query.queryTotalPrincipal(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JoltifyThirdPartyCdpV1Beta1.query.queryTotalPrincipal({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'TotalPrincipal', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTotalPrincipal', payload: { options: { all }, params: {...key},query }})
				return getters['getTotalPrincipal']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTotalPrincipal API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTotalCollateral({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyThirdPartyCdpV1Beta1.query.queryTotalCollateral(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JoltifyThirdPartyCdpV1Beta1.query.queryTotalCollateral({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'TotalCollateral', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTotalCollateral', payload: { options: { all }, params: {...key},query }})
				return getters['getTotalCollateral']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTotalCollateral API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCdps({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyThirdPartyCdpV1Beta1.query.queryCdps(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JoltifyThirdPartyCdpV1Beta1.query.queryCdps({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'Cdps', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCdps', payload: { options: { all }, params: {...key},query }})
				return getters['getCdps']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCdps API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCdp({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyThirdPartyCdpV1Beta1.query.queryCdp( key.owner,  key.collateral_type)).data
				
					
				commit('QUERY', { query: 'Cdp', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCdp', payload: { options: { all }, params: {...key},query }})
				return getters['getCdp']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCdp API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDeposits({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyThirdPartyCdpV1Beta1.query.queryDeposits( key.owner,  key.collateral_type)).data
				
					
				commit('QUERY', { query: 'Deposits', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDeposits', payload: { options: { all }, params: {...key},query }})
				return getters['getDeposits']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDeposits API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgWithdraw({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyThirdPartyCdpV1Beta1.tx.sendMsgWithdraw({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgWithdraw:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgWithdraw:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDrawDebt({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyThirdPartyCdpV1Beta1.tx.sendMsgDrawDebt({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDrawDebt:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDrawDebt:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRepayDebt({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyThirdPartyCdpV1Beta1.tx.sendMsgRepayDebt({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRepayDebt:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRepayDebt:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgLiquidate({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyThirdPartyCdpV1Beta1.tx.sendMsgLiquidate({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgLiquidate:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgLiquidate:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateCDP({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyThirdPartyCdpV1Beta1.tx.sendMsgCreateCDP({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateCDP:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateCDP:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeposit({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyThirdPartyCdpV1Beta1.tx.sendMsgDeposit({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeposit:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeposit:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgWithdraw({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyThirdPartyCdpV1Beta1.tx.msgWithdraw({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgWithdraw:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgWithdraw:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDrawDebt({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyThirdPartyCdpV1Beta1.tx.msgDrawDebt({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDrawDebt:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDrawDebt:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRepayDebt({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyThirdPartyCdpV1Beta1.tx.msgRepayDebt({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRepayDebt:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRepayDebt:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgLiquidate({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyThirdPartyCdpV1Beta1.tx.msgLiquidate({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgLiquidate:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgLiquidate:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateCDP({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyThirdPartyCdpV1Beta1.tx.msgCreateCDP({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateCDP:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateCDP:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeposit({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyThirdPartyCdpV1Beta1.tx.msgDeposit({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeposit:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeposit:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
