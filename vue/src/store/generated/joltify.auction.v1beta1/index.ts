import { Client, registry, MissingWalletError } from 'joltify-finance-joltify_lending-client-ts'

import { BaseAuction } from "joltify-finance-joltify_lending-client-ts/joltify.auction.v1beta1/types"
import { SurplusAuction } from "joltify-finance-joltify_lending-client-ts/joltify.auction.v1beta1/types"
import { DebtAuction } from "joltify-finance-joltify_lending-client-ts/joltify.auction.v1beta1/types"
import { CollateralAuction } from "joltify-finance-joltify_lending-client-ts/joltify.auction.v1beta1/types"
import { WeightedAddresses } from "joltify-finance-joltify_lending-client-ts/joltify.auction.v1beta1/types"
import { Params } from "joltify-finance-joltify_lending-client-ts/joltify.auction.v1beta1/types"


export { BaseAuction, SurplusAuction, DebtAuction, CollateralAuction, WeightedAddresses, Params };

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
				Auction: {},
				Auctions: {},
				NextAuctionID: {},
				
				_Structure: {
						BaseAuction: getStructure(BaseAuction.fromPartial({})),
						SurplusAuction: getStructure(SurplusAuction.fromPartial({})),
						DebtAuction: getStructure(DebtAuction.fromPartial({})),
						CollateralAuction: getStructure(CollateralAuction.fromPartial({})),
						WeightedAddresses: getStructure(WeightedAddresses.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						
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
				getAuction: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Auction[JSON.stringify(params)] ?? {}
		},
				getAuctions: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Auctions[JSON.stringify(params)] ?? {}
		},
				getNextAuctionID: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.NextAuctionID[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: joltify.auction.v1beta1 initialized!')
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
				let value= (await client.JoltifyAuctionV1Beta1.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryAuction({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyAuctionV1Beta1.query.queryAuction( key.auction_id)).data
				
					
				commit('QUERY', { query: 'Auction', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAuction', payload: { options: { all }, params: {...key},query }})
				return getters['getAuction']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAuction API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryAuctions({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyAuctionV1Beta1.query.queryAuctions(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JoltifyAuctionV1Beta1.query.queryAuctions({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'Auctions', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAuctions', payload: { options: { all }, params: {...key},query }})
				return getters['getAuctions']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAuctions API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryNextAuctionID({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JoltifyAuctionV1Beta1.query.queryNextAuctionId()).data
				
					
				commit('QUERY', { query: 'NextAuctionID', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryNextAuctionID', payload: { options: { all }, params: {...key},query }})
				return getters['getNextAuctionID']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryNextAuctionID API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgPlaceBid({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyAuctionV1Beta1.tx.sendMsgPlaceBid({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPlaceBid:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgPlaceBid:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgPlaceBid({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyAuctionV1Beta1.tx.msgPlaceBid({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPlaceBid:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgPlaceBid:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
