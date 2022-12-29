// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgCreateTask } from "./types/sideline/sideline/tx";
import { MsgRegistEmployer } from "./types/sideline/sideline/tx";
import { MsgRegistDeveloper } from "./types/sideline/sideline/tx";
import { MsgDoTask } from "./types/sideline/sideline/tx";


export { MsgCreateTask, MsgRegistEmployer, MsgRegistDeveloper, MsgDoTask };

type sendMsgCreateTaskParams = {
  value: MsgCreateTask,
  fee?: StdFee,
  memo?: string
};

type sendMsgRegistEmployerParams = {
  value: MsgRegistEmployer,
  fee?: StdFee,
  memo?: string
};

type sendMsgRegistDeveloperParams = {
  value: MsgRegistDeveloper,
  fee?: StdFee,
  memo?: string
};

type sendMsgDoTaskParams = {
  value: MsgDoTask,
  fee?: StdFee,
  memo?: string
};


type msgCreateTaskParams = {
  value: MsgCreateTask,
};

type msgRegistEmployerParams = {
  value: MsgRegistEmployer,
};

type msgRegistDeveloperParams = {
  value: MsgRegistDeveloper,
};

type msgDoTaskParams = {
  value: MsgDoTask,
};


export const registry = new Registry(msgTypes);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgCreateTask({ value, fee, memo }: sendMsgCreateTaskParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateTask: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateTask({ value: MsgCreateTask.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateTask: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgRegistEmployer({ value, fee, memo }: sendMsgRegistEmployerParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRegistEmployer: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRegistEmployer({ value: MsgRegistEmployer.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRegistEmployer: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgRegistDeveloper({ value, fee, memo }: sendMsgRegistDeveloperParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRegistDeveloper: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRegistDeveloper({ value: MsgRegistDeveloper.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRegistDeveloper: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgDoTask({ value, fee, memo }: sendMsgDoTaskParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgDoTask: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgDoTask({ value: MsgDoTask.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgDoTask: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgCreateTask({ value }: msgCreateTaskParams): EncodeObject {
			try {
				return { typeUrl: "/sideline.sideline.MsgCreateTask", value: MsgCreateTask.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateTask: Could not create message: ' + e.message)
			}
		},
		
		msgRegistEmployer({ value }: msgRegistEmployerParams): EncodeObject {
			try {
				return { typeUrl: "/sideline.sideline.MsgRegistEmployer", value: MsgRegistEmployer.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRegistEmployer: Could not create message: ' + e.message)
			}
		},
		
		msgRegistDeveloper({ value }: msgRegistDeveloperParams): EncodeObject {
			try {
				return { typeUrl: "/sideline.sideline.MsgRegistDeveloper", value: MsgRegistDeveloper.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRegistDeveloper: Could not create message: ' + e.message)
			}
		},
		
		msgDoTask({ value }: msgDoTaskParams): EncodeObject {
			try {
				return { typeUrl: "/sideline.sideline.MsgDoTask", value: MsgDoTask.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgDoTask: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseURL: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		client.on('signer-changed',(signer) => {			
		 this.updateTX(client);
		})
	}
	updateTX(client: IgniteClient) {
    const methods = txClient({
        signer: client.signer,
        addr: client.env.rpcURL,
        prefix: client.env.prefix ?? "cosmos",
    })
	
    this.tx = methods;
    for (let m in methods) {
        this.tx[m] = methods[m].bind(this.tx);
    }
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			SidelineSideline: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;