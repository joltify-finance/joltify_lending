// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgDeposit } from "./types/joltify/third_party/cdp/v1beta1/tx";
import { MsgCreateCDP } from "./types/joltify/third_party/cdp/v1beta1/tx";
import { MsgWithdraw } from "./types/joltify/third_party/cdp/v1beta1/tx";
import { MsgDrawDebt } from "./types/joltify/third_party/cdp/v1beta1/tx";
import { MsgLiquidate } from "./types/joltify/third_party/cdp/v1beta1/tx";
import { MsgRepayDebt } from "./types/joltify/third_party/cdp/v1beta1/tx";

import { CDP as typeCDP} from "./types"
import { Deposit as typeDeposit} from "./types"
import { TotalPrincipal as typeTotalPrincipal} from "./types"
import { TotalCollateral as typeTotalCollateral} from "./types"
import { OwnerCDPIndex as typeOwnerCDPIndex} from "./types"
import { Params as typeParams} from "./types"
import { DebtParam as typeDebtParam} from "./types"
import { CollateralParam as typeCollateralParam} from "./types"
import { GenesisAccumulationTime as typeGenesisAccumulationTime} from "./types"
import { GenesisTotalPrincipal as typeGenesisTotalPrincipal} from "./types"
import { CDPResponse as typeCDPResponse} from "./types"

export { MsgDeposit, MsgCreateCDP, MsgWithdraw, MsgDrawDebt, MsgLiquidate, MsgRepayDebt };

type sendMsgDepositParams = {
  value: MsgDeposit,
  fee?: StdFee,
  memo?: string
};

type sendMsgCreateCDPParams = {
  value: MsgCreateCDP,
  fee?: StdFee,
  memo?: string
};

type sendMsgWithdrawParams = {
  value: MsgWithdraw,
  fee?: StdFee,
  memo?: string
};

type sendMsgDrawDebtParams = {
  value: MsgDrawDebt,
  fee?: StdFee,
  memo?: string
};

type sendMsgLiquidateParams = {
  value: MsgLiquidate,
  fee?: StdFee,
  memo?: string
};

type sendMsgRepayDebtParams = {
  value: MsgRepayDebt,
  fee?: StdFee,
  memo?: string
};


type msgDepositParams = {
  value: MsgDeposit,
};

type msgCreateCDPParams = {
  value: MsgCreateCDP,
};

type msgWithdrawParams = {
  value: MsgWithdraw,
};

type msgDrawDebtParams = {
  value: MsgDrawDebt,
};

type msgLiquidateParams = {
  value: MsgLiquidate,
};

type msgRepayDebtParams = {
  value: MsgRepayDebt,
};


export const registry = new Registry(msgTypes);

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	const structure: {fields: Field[]} = { fields: [] }
	for (let [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
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
		
		async sendMsgCreateCDP({ value, fee, memo }: sendMsgCreateCDPParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateCDP: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateCDP({ value: MsgCreateCDP.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateCDP: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgWithdraw({ value, fee, memo }: sendMsgWithdrawParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgWithdraw: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgWithdraw({ value: MsgWithdraw.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgWithdraw: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgDrawDebt({ value, fee, memo }: sendMsgDrawDebtParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgDrawDebt: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgDrawDebt({ value: MsgDrawDebt.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgDrawDebt: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgLiquidate({ value, fee, memo }: sendMsgLiquidateParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgLiquidate: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgLiquidate({ value: MsgLiquidate.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgLiquidate: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgRepayDebt({ value, fee, memo }: sendMsgRepayDebtParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRepayDebt: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRepayDebt({ value: MsgRepayDebt.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRepayDebt: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgDeposit({ value }: msgDepositParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.third_party.cdp.v1beta1.MsgDeposit", value: MsgDeposit.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgDeposit: Could not create message: ' + e.message)
			}
		},
		
		msgCreateCDP({ value }: msgCreateCDPParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.third_party.cdp.v1beta1.MsgCreateCDP", value: MsgCreateCDP.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateCDP: Could not create message: ' + e.message)
			}
		},
		
		msgWithdraw({ value }: msgWithdrawParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.third_party.cdp.v1beta1.MsgWithdraw", value: MsgWithdraw.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgWithdraw: Could not create message: ' + e.message)
			}
		},
		
		msgDrawDebt({ value }: msgDrawDebtParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.third_party.cdp.v1beta1.MsgDrawDebt", value: MsgDrawDebt.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgDrawDebt: Could not create message: ' + e.message)
			}
		},
		
		msgLiquidate({ value }: msgLiquidateParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.third_party.cdp.v1beta1.MsgLiquidate", value: MsgLiquidate.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgLiquidate: Could not create message: ' + e.message)
			}
		},
		
		msgRepayDebt({ value }: msgRepayDebtParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.third_party.cdp.v1beta1.MsgRepayDebt", value: MsgRepayDebt.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRepayDebt: Could not create message: ' + e.message)
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
	public structure: Record<string,unknown>;
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		this.structure =  {
						CDP: getStructure(typeCDP.fromPartial({})),
						Deposit: getStructure(typeDeposit.fromPartial({})),
						TotalPrincipal: getStructure(typeTotalPrincipal.fromPartial({})),
						TotalCollateral: getStructure(typeTotalCollateral.fromPartial({})),
						OwnerCDPIndex: getStructure(typeOwnerCDPIndex.fromPartial({})),
						Params: getStructure(typeParams.fromPartial({})),
						DebtParam: getStructure(typeDebtParam.fromPartial({})),
						CollateralParam: getStructure(typeCollateralParam.fromPartial({})),
						GenesisAccumulationTime: getStructure(typeGenesisAccumulationTime.fromPartial({})),
						GenesisTotalPrincipal: getStructure(typeGenesisTotalPrincipal.fromPartial({})),
						CDPResponse: getStructure(typeCDPResponse.fromPartial({})),
						
		};
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
			JoltifyThirdPartyCdpV1Beta1: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;