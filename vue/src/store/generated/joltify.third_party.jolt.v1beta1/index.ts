import { Client, registry, MissingWalletError } from 'joltify-finance-joltify_lending-client-ts'

import { GenesisAccumulationTime } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.jolt.v1beta1/types"
import { Params } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.jolt.v1beta1/types"
import { MoneyMarket } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.jolt.v1beta1/types"
import { BorrowLimit } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.jolt.v1beta1/types"
import { InterestRateModel } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.jolt.v1beta1/types"
import { Deposit } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.jolt.v1beta1/types"
import { Borrow } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.jolt.v1beta1/types"
import { SupplyInterestFactor } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.jolt.v1beta1/types"
import { BorrowInterestFactor } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.jolt.v1beta1/types"
import { CoinsProto } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.jolt.v1beta1/types"
import { LiquidateItem } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.jolt.v1beta1/types"
import { DepositResponse } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.jolt.v1beta1/types"
import { SupplyInterestFactorResponse } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.jolt.v1beta1/types"
import { BorrowResponse } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.jolt.v1beta1/types"
import { BorrowInterestFactorResponse } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.jolt.v1beta1/types"
import { MoneyMarketInterestRate } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.jolt.v1beta1/types"
import { InterestFactor } from "joltify-finance-joltify_lending-client-ts/joltify.third_party.jolt.v1beta1/types"


export { GenesisAccumulationTime, Params, MoneyMarket, BorrowLimit, InterestRateModel, Deposit, Borrow, SupplyInterestFactor, BorrowInterestFactor, CoinsProto, LiquidateItem, DepositResponse, SupplyInterestFactorResponse, BorrowResponse, BorrowInterestFactorResponse, MoneyMarketInterestRate, InterestFactor };

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
				Deposits: {},
				UnsyncedDeposits: {},
				TotalDeposited: {},
				Borrows: {},
				UnsyncedBorrows: {},
				TotalBorrowed: {},
				InterestRate: {},
				Reserves: {},
				InterestFactors: {},
				liquidate: {},
				
				_Structure: {
						GenesisAccumulationTime: getStructure(GenesisAccumulationTime.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						MoneyMarket: getStructure(MoneyMarket.fromPartial({})),
						BorrowLimit: getStructure(BorrowLimit.fromPartial({})),
						InterestRateModel: getStructure(InterestRateModel.fromPartial({})),
						Deposit: getStructure(Deposit.fromPartial({})),
						Borrow: getStructure(Borrow.fromPartial({})),
						SupplyInterestFactor: getStructure(SupplyInterestFactor.fromPartial({})),
						BorrowInterestFactor: getStructure(BorrowInterestFactor.fromPartial({})),
						CoinsProto: getStructure(CoinsProto.fromPartial({})),
						LiquidateItem: getStructure(LiquidateItem.fromPartial({})),
						DepositResponse: getStructure(DepositResponse.fromPartial({})),
						SupplyInterestFactorResponse: getStructure(SupplyInterestFactorResponse.fromPartial({})),
						BorrowResponse: getStructure(BorrowResponse.fromPartial({})),
						BorrowInterestFactorResponse: getStructure(BorrowInterestFactorResponse.fromPartial({})),
						MoneyMarketInterestRate: getStructure(MoneyMarketInterestRate.fromPartial({})),
						InterestFactor: getStructure(InterestFactor.fromPartial({})),
						
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
				getDeposits: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Deposits[JSON.stringify(params)] ?? {}
		},
				getUnsyncedDeposits: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.UnsyncedDeposits[JSON.stringify(params)] ?? {}
		},
				getTotalDeposited: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TotalDeposited[JSON.stringify(params)] ?? {}
		},
				getBorrows: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Borrows[JSON.stringify(params)] ?? {}
		},
				getUnsyncedBorrows: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.UnsyncedBorrows[JSON.stringify(params)] ?? {}
		},
				getTotalBorrowed: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TotalBorrowed[JSON.stringify(params)] ?? {}
		},
				getInterestRate: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.InterestRate[JSON.stringify(params)] ?? {}
		},
				getReserves: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Reserves[JSON.stringify(params)] ?? {}
		},
				getInterestFactors: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.InterestFactors[JSON.stringify(params)] ?? {}
		},
				getliquidate: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.liquidate[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: joltify.third_party.jolt.v1beta1 initialized!')
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
				let value= (await client.JoltifyThirdPartyJoltV1Beta1.query.queryParams()).data
				
					
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
				let value= (await client.JoltifyThirdPartyJoltV1Beta1.query.queryAccounts()).data
				
					
				commit('QUERY', { query: 'Accounts', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAccounts', payload: { options: { all }, params: {...key},query }})
				return getters['getAccounts']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAccounts API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDeposits({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyThirdPartyJoltV1Beta1.query.queryDeposits(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JoltifyThirdPartyJoltV1Beta1.query.queryDeposits({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'Deposits', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDeposits', payload: { options: { all }, params: {...key},query }})
				return getters['getDeposits']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDeposits API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryUnsyncedDeposits({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyThirdPartyJoltV1Beta1.query.queryUnsyncedDeposits(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JoltifyThirdPartyJoltV1Beta1.query.queryUnsyncedDeposits({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'UnsyncedDeposits', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryUnsyncedDeposits', payload: { options: { all }, params: {...key},query }})
				return getters['getUnsyncedDeposits']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryUnsyncedDeposits API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTotalDeposited({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyThirdPartyJoltV1Beta1.query.queryTotalDeposited( key.denom)).data
				
					
				commit('QUERY', { query: 'TotalDeposited', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTotalDeposited', payload: { options: { all }, params: {...key},query }})
				return getters['getTotalDeposited']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTotalDeposited API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryBorrows({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyThirdPartyJoltV1Beta1.query.queryBorrows(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JoltifyThirdPartyJoltV1Beta1.query.queryBorrows({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'Borrows', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryBorrows', payload: { options: { all }, params: {...key},query }})
				return getters['getBorrows']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryBorrows API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryUnsyncedBorrows({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyThirdPartyJoltV1Beta1.query.queryUnsyncedBorrows(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JoltifyThirdPartyJoltV1Beta1.query.queryUnsyncedBorrows({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'UnsyncedBorrows', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryUnsyncedBorrows', payload: { options: { all }, params: {...key},query }})
				return getters['getUnsyncedBorrows']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryUnsyncedBorrows API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTotalBorrowed({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyThirdPartyJoltV1Beta1.query.queryTotalBorrowed( key.denom)).data
				
					
				commit('QUERY', { query: 'TotalBorrowed', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTotalBorrowed', payload: { options: { all }, params: {...key},query }})
				return getters['getTotalBorrowed']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTotalBorrowed API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryInterestRate({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyThirdPartyJoltV1Beta1.query.queryInterestRate( key.denom)).data
				
					
				commit('QUERY', { query: 'InterestRate', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryInterestRate', payload: { options: { all }, params: {...key},query }})
				return getters['getInterestRate']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryInterestRate API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryReserves({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyThirdPartyJoltV1Beta1.query.queryReserves( key.denom)).data
				
					
				commit('QUERY', { query: 'Reserves', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryReserves', payload: { options: { all }, params: {...key},query }})
				return getters['getReserves']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryReserves API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryInterestFactors({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyThirdPartyJoltV1Beta1.query.queryInterestFactors( key.denom)).data
				
					
				commit('QUERY', { query: 'InterestFactors', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryInterestFactors', payload: { options: { all }, params: {...key},query }})
				return getters['getInterestFactors']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryInterestFactors API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async Queryliquidate({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyThirdPartyJoltV1Beta1.query.queryliquidate(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JoltifyThirdPartyJoltV1Beta1.query.queryliquidate({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'liquidate', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'Queryliquidate', payload: { options: { all }, params: {...key},query }})
				return getters['getliquidate']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:Queryliquidate API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgWithdraw({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyThirdPartyJoltV1Beta1.tx.sendMsgWithdraw({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgWithdraw:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgWithdraw:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeposit({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyThirdPartyJoltV1Beta1.tx.sendMsgDeposit({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeposit:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeposit:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgBorrow({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyThirdPartyJoltV1Beta1.tx.sendMsgBorrow({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBorrow:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgBorrow:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgLiquidate({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyThirdPartyJoltV1Beta1.tx.sendMsgLiquidate({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgLiquidate:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgLiquidate:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRepay({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyThirdPartyJoltV1Beta1.tx.sendMsgRepay({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRepay:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRepay:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgWithdraw({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyThirdPartyJoltV1Beta1.tx.msgWithdraw({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgWithdraw:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgWithdraw:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeposit({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyThirdPartyJoltV1Beta1.tx.msgDeposit({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeposit:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeposit:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgBorrow({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyThirdPartyJoltV1Beta1.tx.msgBorrow({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBorrow:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgBorrow:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgLiquidate({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyThirdPartyJoltV1Beta1.tx.msgLiquidate({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgLiquidate:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgLiquidate:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRepay({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyThirdPartyJoltV1Beta1.tx.msgRepay({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRepay:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRepay:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
