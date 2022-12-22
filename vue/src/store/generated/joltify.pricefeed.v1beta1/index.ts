import { Client, registry, MissingWalletError } from 'joltify-finance-joltify_lending-client-ts'

import { PostedPriceResponse } from "joltify-finance-joltify_lending-client-ts/joltify.pricefeed.v1beta1/types"
import { CurrentPriceResponse } from "joltify-finance-joltify_lending-client-ts/joltify.pricefeed.v1beta1/types"
import { MarketResponse } from "joltify-finance-joltify_lending-client-ts/joltify.pricefeed.v1beta1/types"
import { Params } from "joltify-finance-joltify_lending-client-ts/joltify.pricefeed.v1beta1/types"
import { Market } from "joltify-finance-joltify_lending-client-ts/joltify.pricefeed.v1beta1/types"
import { PostedPrice } from "joltify-finance-joltify_lending-client-ts/joltify.pricefeed.v1beta1/types"
import { CurrentPrice } from "joltify-finance-joltify_lending-client-ts/joltify.pricefeed.v1beta1/types"


export { PostedPriceResponse, CurrentPriceResponse, MarketResponse, Params, Market, PostedPrice, CurrentPrice };

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
				Price: {},
				Prices: {},
				RawPrices: {},
				Oracles: {},
				Markets: {},
				
				_Structure: {
						PostedPriceResponse: getStructure(PostedPriceResponse.fromPartial({})),
						CurrentPriceResponse: getStructure(CurrentPriceResponse.fromPartial({})),
						MarketResponse: getStructure(MarketResponse.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						Market: getStructure(Market.fromPartial({})),
						PostedPrice: getStructure(PostedPrice.fromPartial({})),
						CurrentPrice: getStructure(CurrentPrice.fromPartial({})),
						
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
				getPrice: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Price[JSON.stringify(params)] ?? {}
		},
				getPrices: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Prices[JSON.stringify(params)] ?? {}
		},
				getRawPrices: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RawPrices[JSON.stringify(params)] ?? {}
		},
				getOracles: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Oracles[JSON.stringify(params)] ?? {}
		},
				getMarkets: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Markets[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: joltify.pricefeed.v1beta1 initialized!')
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
				let value= (await client.JoltifyPricefeedV1Beta1.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPrice({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyPricefeedV1Beta1.query.queryPrice( key.market_id)).data
				
					
				commit('QUERY', { query: 'Price', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPrice', payload: { options: { all }, params: {...key},query }})
				return getters['getPrice']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPrice API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPrices({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyPricefeedV1Beta1.query.queryPrices()).data
				
					
				commit('QUERY', { query: 'Prices', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPrices', payload: { options: { all }, params: {...key},query }})
				return getters['getPrices']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPrices API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRawPrices({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyPricefeedV1Beta1.query.queryRawPrices( key.market_id)).data
				
					
				commit('QUERY', { query: 'RawPrices', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRawPrices', payload: { options: { all }, params: {...key},query }})
				return getters['getRawPrices']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRawPrices API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryOracles({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyPricefeedV1Beta1.query.queryOracles( key.market_id)).data
				
					
				commit('QUERY', { query: 'Oracles', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryOracles', payload: { options: { all }, params: {...key},query }})
				return getters['getOracles']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryOracles API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryMarkets({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyPricefeedV1Beta1.query.queryMarkets()).data
				
					
				commit('QUERY', { query: 'Markets', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryMarkets', payload: { options: { all }, params: {...key},query }})
				return getters['getMarkets']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryMarkets API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgPostPrice({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyPricefeedV1Beta1.tx.sendMsgPostPrice({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPostPrice:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgPostPrice:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgPostPrice({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyPricefeedV1Beta1.tx.msgPostPrice({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPostPrice:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgPostPrice:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
