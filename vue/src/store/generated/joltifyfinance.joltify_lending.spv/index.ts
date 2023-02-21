import { Client, registry, MissingWalletError } from 'joltify-finance-joltify_lending-client-ts'



export {  };

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
			console.log('Vuex module: joltifyfinance.joltify_lending.spv initialized!')
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
		
		async sendMsgActivePool({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyfinanceJoltifyLendingSpv.tx.sendMsgActivePool({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgActivePool:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgActivePool:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdatePool({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyfinanceJoltifyLendingSpv.tx.sendMsgUpdatePool({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdatePool:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdatePool:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgClaimInterest({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyfinanceJoltifyLendingSpv.tx.sendMsgClaimInterest({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgClaimInterest:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgClaimInterest:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRepayInterest({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JoltifyfinanceJoltifyLendingSpv.tx.sendMsgRepayInterest({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRepayInterest:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRepayInterest:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgActivePool({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyfinanceJoltifyLendingSpv.tx.msgActivePool({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgActivePool:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgActivePool:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdatePool({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyfinanceJoltifyLendingSpv.tx.msgUpdatePool({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdatePool:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdatePool:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgClaimInterest({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyfinanceJoltifyLendingSpv.tx.msgClaimInterest({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgClaimInterest:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgClaimInterest:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRepayInterest({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JoltifyfinanceJoltifyLendingSpv.tx.msgRepayInterest({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRepayInterest:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRepayInterest:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
