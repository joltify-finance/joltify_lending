// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgClaimUSDXMintingReward } from "./types/joltify/third_party/incentive/v1beta1/tx";
import { MsgClaimDelegatorReward } from "./types/joltify/third_party/incentive/v1beta1/tx";
import { MsgClaimSwapReward } from "./types/joltify/third_party/incentive/v1beta1/tx";
import { MsgClaimSavingsReward } from "./types/joltify/third_party/incentive/v1beta1/tx";
import { MsgClaimJoltReward } from "./types/joltify/third_party/incentive/v1beta1/tx";


export { MsgClaimUSDXMintingReward, MsgClaimDelegatorReward, MsgClaimSwapReward, MsgClaimSavingsReward, MsgClaimJoltReward };

type sendMsgClaimUSDXMintingRewardParams = {
  value: MsgClaimUSDXMintingReward,
  fee?: StdFee,
  memo?: string
};

type sendMsgClaimDelegatorRewardParams = {
  value: MsgClaimDelegatorReward,
  fee?: StdFee,
  memo?: string
};

type sendMsgClaimSwapRewardParams = {
  value: MsgClaimSwapReward,
  fee?: StdFee,
  memo?: string
};

type sendMsgClaimSavingsRewardParams = {
  value: MsgClaimSavingsReward,
  fee?: StdFee,
  memo?: string
};

type sendMsgClaimJoltRewardParams = {
  value: MsgClaimJoltReward,
  fee?: StdFee,
  memo?: string
};


type msgClaimUSDXMintingRewardParams = {
  value: MsgClaimUSDXMintingReward,
};

type msgClaimDelegatorRewardParams = {
  value: MsgClaimDelegatorReward,
};

type msgClaimSwapRewardParams = {
  value: MsgClaimSwapReward,
};

type msgClaimSavingsRewardParams = {
  value: MsgClaimSavingsReward,
};

type msgClaimJoltRewardParams = {
  value: MsgClaimJoltReward,
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
		
		async sendMsgClaimUSDXMintingReward({ value, fee, memo }: sendMsgClaimUSDXMintingRewardParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgClaimUSDXMintingReward: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgClaimUSDXMintingReward({ value: MsgClaimUSDXMintingReward.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgClaimUSDXMintingReward: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgClaimDelegatorReward({ value, fee, memo }: sendMsgClaimDelegatorRewardParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgClaimDelegatorReward: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgClaimDelegatorReward({ value: MsgClaimDelegatorReward.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgClaimDelegatorReward: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgClaimSwapReward({ value, fee, memo }: sendMsgClaimSwapRewardParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgClaimSwapReward: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgClaimSwapReward({ value: MsgClaimSwapReward.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgClaimSwapReward: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgClaimSavingsReward({ value, fee, memo }: sendMsgClaimSavingsRewardParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgClaimSavingsReward: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgClaimSavingsReward({ value: MsgClaimSavingsReward.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgClaimSavingsReward: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgClaimJoltReward({ value, fee, memo }: sendMsgClaimJoltRewardParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgClaimJoltReward: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgClaimJoltReward({ value: MsgClaimJoltReward.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgClaimJoltReward: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgClaimUSDXMintingReward({ value }: msgClaimUSDXMintingRewardParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.third_party.incentive.v1beta1.MsgClaimUSDXMintingReward", value: MsgClaimUSDXMintingReward.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgClaimUSDXMintingReward: Could not create message: ' + e.message)
			}
		},
		
		msgClaimDelegatorReward({ value }: msgClaimDelegatorRewardParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.third_party.incentive.v1beta1.MsgClaimDelegatorReward", value: MsgClaimDelegatorReward.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgClaimDelegatorReward: Could not create message: ' + e.message)
			}
		},
		
		msgClaimSwapReward({ value }: msgClaimSwapRewardParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.third_party.incentive.v1beta1.MsgClaimSwapReward", value: MsgClaimSwapReward.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgClaimSwapReward: Could not create message: ' + e.message)
			}
		},
		
		msgClaimSavingsReward({ value }: msgClaimSavingsRewardParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.third_party.incentive.v1beta1.MsgClaimSavingsReward", value: MsgClaimSavingsReward.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgClaimSavingsReward: Could not create message: ' + e.message)
			}
		},
		
		msgClaimJoltReward({ value }: msgClaimJoltRewardParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.third_party.incentive.v1beta1.MsgClaimJoltReward", value: MsgClaimJoltReward.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgClaimJoltReward: Could not create message: ' + e.message)
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
			JoltifyThirdPartyIncentiveV1Beta1: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;