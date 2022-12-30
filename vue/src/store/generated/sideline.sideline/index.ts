import { Client, registry, MissingWalletError } from 'sideline-client-ts'

import { Developer } from "sideline-client-ts/sideline.sideline/types"
import { Employer } from "sideline-client-ts/sideline.sideline/types"
import { Params } from "sideline-client-ts/sideline.sideline/types"
import { Task } from "sideline-client-ts/sideline.sideline/types"


export { Developer, Employer, Params, Task };

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
				Employer: {},
				EmployerAll: {},
				Developer: {},
				DeveloperAll: {},
				Task: {},
				TaskAll: {},
				
				_Structure: {
						Developer: getStructure(Developer.fromPartial({})),
						Employer: getStructure(Employer.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						Task: getStructure(Task.fromPartial({})),
						
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
				getEmployer: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Employer[JSON.stringify(params)] ?? {}
		},
				getEmployerAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.EmployerAll[JSON.stringify(params)] ?? {}
		},
				getDeveloper: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Developer[JSON.stringify(params)] ?? {}
		},
				getDeveloperAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.DeveloperAll[JSON.stringify(params)] ?? {}
		},
				getTask: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Task[JSON.stringify(params)] ?? {}
		},
				getTaskAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TaskAll[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: sideline.sideline initialized!')
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
				let value= (await client.SidelineSideline.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryEmployer({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SidelineSideline.query.queryEmployer( key.index)).data
				
					
				commit('QUERY', { query: 'Employer', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryEmployer', payload: { options: { all }, params: {...key},query }})
				return getters['getEmployer']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryEmployer API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryEmployerAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SidelineSideline.query.queryEmployerAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.SidelineSideline.query.queryEmployerAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'EmployerAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryEmployerAll', payload: { options: { all }, params: {...key},query }})
				return getters['getEmployerAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryEmployerAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDeveloper({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SidelineSideline.query.queryDeveloper( key.index)).data
				
					
				commit('QUERY', { query: 'Developer', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDeveloper', payload: { options: { all }, params: {...key},query }})
				return getters['getDeveloper']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDeveloper API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDeveloperAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SidelineSideline.query.queryDeveloperAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.SidelineSideline.query.queryDeveloperAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'DeveloperAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDeveloperAll', payload: { options: { all }, params: {...key},query }})
				return getters['getDeveloperAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDeveloperAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTask({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SidelineSideline.query.queryTask( key.id)).data
				
					
				commit('QUERY', { query: 'Task', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTask', payload: { options: { all }, params: {...key},query }})
				return getters['getTask']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTask API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTaskAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SidelineSideline.query.queryTaskAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.SidelineSideline.query.queryTaskAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'TaskAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTaskAll', payload: { options: { all }, params: {...key},query }})
				return getters['getTaskAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTaskAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgSuccessTask({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SidelineSideline.tx.sendMsgSuccessTask({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSuccessTask:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSuccessTask:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCancelTask({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SidelineSideline.tx.sendMsgCancelTask({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCancelTask:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCancelTask:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRegistEmployer({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SidelineSideline.tx.sendMsgRegistEmployer({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRegistEmployer:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRegistEmployer:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateTask({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SidelineSideline.tx.sendMsgCreateTask({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateTask:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateTask:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgStartJudgeTask({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SidelineSideline.tx.sendMsgStartJudgeTask({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgStartJudgeTask:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgStartJudgeTask:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUndoneTask({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SidelineSideline.tx.sendMsgUndoneTask({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUndoneTask:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUndoneTask:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgFailTask({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SidelineSideline.tx.sendMsgFailTask({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgFailTask:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgFailTask:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRegistDeveloper({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SidelineSideline.tx.sendMsgRegistDeveloper({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRegistDeveloper:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRegistDeveloper:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDoTask({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SidelineSideline.tx.sendMsgDoTask({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDoTask:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDoTask:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSubmitTask({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SidelineSideline.tx.sendMsgSubmitTask({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitTask:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSubmitTask:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgSuccessTask({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SidelineSideline.tx.msgSuccessTask({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSuccessTask:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSuccessTask:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCancelTask({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SidelineSideline.tx.msgCancelTask({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCancelTask:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCancelTask:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRegistEmployer({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SidelineSideline.tx.msgRegistEmployer({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRegistEmployer:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRegistEmployer:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateTask({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SidelineSideline.tx.msgCreateTask({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateTask:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateTask:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgStartJudgeTask({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SidelineSideline.tx.msgStartJudgeTask({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgStartJudgeTask:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgStartJudgeTask:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUndoneTask({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SidelineSideline.tx.msgUndoneTask({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUndoneTask:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUndoneTask:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgFailTask({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SidelineSideline.tx.msgFailTask({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgFailTask:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgFailTask:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRegistDeveloper({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SidelineSideline.tx.msgRegistDeveloper({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRegistDeveloper:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRegistDeveloper:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDoTask({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SidelineSideline.tx.msgDoTask({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDoTask:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDoTask:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSubmitTask({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SidelineSideline.tx.msgSubmitTask({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitTask:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSubmitTask:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
