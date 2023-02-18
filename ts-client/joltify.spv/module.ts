// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgRepayInterest } from "./types/joltify/spv/tx";
import { MsgCreatePool } from "./types/joltify/spv/tx";
import { MsgBorrow } from "./types/joltify/spv/tx";
import { MsgAddInvestors } from "./types/joltify/spv/tx";
import { MsgDeposit } from "./types/joltify/spv/tx";


export { MsgRepayInterest, MsgCreatePool, MsgBorrow, MsgAddInvestors, MsgDeposit };

type sendMsgRepayInterestParams = {
  value: MsgRepayInterest,
  fee?: StdFee,
  memo?: string
};

type sendMsgCreatePoolParams = {
  value: MsgCreatePool,
  fee?: StdFee,
  memo?: string
};

type sendMsgBorrowParams = {
  value: MsgBorrow,
  fee?: StdFee,
  memo?: string
};

type sendMsgAddInvestorsParams = {
  value: MsgAddInvestors,
  fee?: StdFee,
  memo?: string
};

type sendMsgDepositParams = {
  value: MsgDeposit,
  fee?: StdFee,
  memo?: string
};


type msgRepayInterestParams = {
  value: MsgRepayInterest,
};

type msgCreatePoolParams = {
  value: MsgCreatePool,
};

type msgBorrowParams = {
  value: MsgBorrow,
};

type msgAddInvestorsParams = {
  value: MsgAddInvestors,
};

type msgDepositParams = {
  value: MsgDeposit,
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
		
		async sendMsgRepayInterest({ value, fee, memo }: sendMsgRepayInterestParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRepayInterest: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRepayInterest({ value: MsgRepayInterest.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRepayInterest: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCreatePool({ value, fee, memo }: sendMsgCreatePoolParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreatePool: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreatePool({ value: MsgCreatePool.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreatePool: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgBorrow({ value, fee, memo }: sendMsgBorrowParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgBorrow: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgBorrow({ value: MsgBorrow.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgBorrow: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgAddInvestors({ value, fee, memo }: sendMsgAddInvestorsParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgAddInvestors: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgAddInvestors({ value: MsgAddInvestors.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgAddInvestors: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgDeposit({ value, fee, memo }: sendMsgDepositParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgDeposit: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgDeposit({ value: MsgDeposit.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgDeposit: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgRepayInterest({ value }: msgRepayInterestParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.spv.MsgRepayInterest", value: MsgRepayInterest.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRepayInterest: Could not create message: ' + e.message)
			}
		},
		
		msgCreatePool({ value }: msgCreatePoolParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.spv.MsgCreatePool", value: MsgCreatePool.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreatePool: Could not create message: ' + e.message)
			}
		},
		
		msgBorrow({ value }: msgBorrowParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.spv.MsgBorrow", value: MsgBorrow.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgBorrow: Could not create message: ' + e.message)
			}
		},
		
		msgAddInvestors({ value }: msgAddInvestorsParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.spv.MsgAddInvestors", value: MsgAddInvestors.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAddInvestors: Could not create message: ' + e.message)
			}
		},
		
		msgDeposit({ value }: msgDepositParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.spv.MsgDeposit", value: MsgDeposit.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgDeposit: Could not create message: ' + e.message)
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
			JoltifySpv: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;