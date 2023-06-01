// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgRepayInterest } from "./types/joltify/spv/tx";
import { MsgActivePool } from "./types/joltify/spv/tx";
import { MsgCreatePool } from "./types/joltify/spv/tx";
import { MsgPayPrincipal } from "./types/joltify/spv/tx";
import { MsgTransferOwnership } from "./types/joltify/spv/tx";
import { MsgUpdatePool } from "./types/joltify/spv/tx";
import { MsgWithdrawPrincipal } from "./types/joltify/spv/tx";
import { MsgAddInvestors } from "./types/joltify/spv/tx";
import { MsgBorrow } from "./types/joltify/spv/tx";
import { MsgSubmitWithdrawProposal } from "./types/joltify/spv/tx";
import { MsgClaimInterest } from "./types/joltify/spv/tx";
import { MsgDeposit } from "./types/joltify/spv/tx";

import { DepositorInfo as typeDepositorInfo} from "./types"
import { BorrowDetail as typeBorrowDetail} from "./types"
import { NftInfo as typeNftInfo} from "./types"
import { PaymentItem as typePaymentItem} from "./types"
import { BorrowInterest as typeBorrowInterest} from "./types"
import { Params as typeParams} from "./types"
import { PoolInfo as typePoolInfo} from "./types"
import { PoolWithInvestors as typePoolWithInvestors} from "./types"
import { PoolDepositedInvestors as typePoolDepositedInvestors} from "./types"
import { WalletsLinkPool as typeWalletsLinkPool} from "./types"

export { MsgRepayInterest, MsgActivePool, MsgCreatePool, MsgPayPrincipal, MsgTransferOwnership, MsgUpdatePool, MsgWithdrawPrincipal, MsgAddInvestors, MsgBorrow, MsgSubmitWithdrawProposal, MsgClaimInterest, MsgDeposit };

type sendMsgRepayInterestParams = {
  value: MsgRepayInterest,
  fee?: StdFee,
  memo?: string
};

type sendMsgActivePoolParams = {
  value: MsgActivePool,
  fee?: StdFee,
  memo?: string
};

type sendMsgCreatePoolParams = {
  value: MsgCreatePool,
  fee?: StdFee,
  memo?: string
};

type sendMsgPayPrincipalParams = {
  value: MsgPayPrincipal,
  fee?: StdFee,
  memo?: string
};

type sendMsgTransferOwnershipParams = {
  value: MsgTransferOwnership,
  fee?: StdFee,
  memo?: string
};

type sendMsgUpdatePoolParams = {
  value: MsgUpdatePool,
  fee?: StdFee,
  memo?: string
};

type sendMsgWithdrawPrincipalParams = {
  value: MsgWithdrawPrincipal,
  fee?: StdFee,
  memo?: string
};

type sendMsgAddInvestorsParams = {
  value: MsgAddInvestors,
  fee?: StdFee,
  memo?: string
};

type sendMsgBorrowParams = {
  value: MsgBorrow,
  fee?: StdFee,
  memo?: string
};

type sendMsgSubmitWithdrawProposalParams = {
  value: MsgSubmitWithdrawProposal,
  fee?: StdFee,
  memo?: string
};

type sendMsgClaimInterestParams = {
  value: MsgClaimInterest,
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

type msgActivePoolParams = {
  value: MsgActivePool,
};

type msgCreatePoolParams = {
  value: MsgCreatePool,
};

type msgPayPrincipalParams = {
  value: MsgPayPrincipal,
};

type msgTransferOwnershipParams = {
  value: MsgTransferOwnership,
};

type msgUpdatePoolParams = {
  value: MsgUpdatePool,
};

type msgWithdrawPrincipalParams = {
  value: MsgWithdrawPrincipal,
};

type msgAddInvestorsParams = {
  value: MsgAddInvestors,
};

type msgBorrowParams = {
  value: MsgBorrow,
};

type msgSubmitWithdrawProposalParams = {
  value: MsgSubmitWithdrawProposal,
};

type msgClaimInterestParams = {
  value: MsgClaimInterest,
};

type msgDepositParams = {
  value: MsgDeposit,
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
		
		async sendMsgActivePool({ value, fee, memo }: sendMsgActivePoolParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgActivePool: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgActivePool({ value: MsgActivePool.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgActivePool: Could not broadcast Tx: '+ e.message)
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
		
		async sendMsgPayPrincipal({ value, fee, memo }: sendMsgPayPrincipalParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgPayPrincipal: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgPayPrincipal({ value: MsgPayPrincipal.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgPayPrincipal: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgTransferOwnership({ value, fee, memo }: sendMsgTransferOwnershipParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgTransferOwnership: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgTransferOwnership({ value: MsgTransferOwnership.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgTransferOwnership: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgUpdatePool({ value, fee, memo }: sendMsgUpdatePoolParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgUpdatePool: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgUpdatePool({ value: MsgUpdatePool.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgUpdatePool: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgWithdrawPrincipal({ value, fee, memo }: sendMsgWithdrawPrincipalParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgWithdrawPrincipal: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgWithdrawPrincipal({ value: MsgWithdrawPrincipal.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgWithdrawPrincipal: Could not broadcast Tx: '+ e.message)
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
		
		async sendMsgSubmitWithdrawProposal({ value, fee, memo }: sendMsgSubmitWithdrawProposalParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgSubmitWithdrawProposal: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgSubmitWithdrawProposal({ value: MsgSubmitWithdrawProposal.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgSubmitWithdrawProposal: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgClaimInterest({ value, fee, memo }: sendMsgClaimInterestParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgClaimInterest: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgClaimInterest({ value: MsgClaimInterest.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgClaimInterest: Could not broadcast Tx: '+ e.message)
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
		
		msgActivePool({ value }: msgActivePoolParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.spv.MsgActivePool", value: MsgActivePool.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgActivePool: Could not create message: ' + e.message)
			}
		},
		
		msgCreatePool({ value }: msgCreatePoolParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.spv.MsgCreatePool", value: MsgCreatePool.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreatePool: Could not create message: ' + e.message)
			}
		},
		
		msgPayPrincipal({ value }: msgPayPrincipalParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.spv.MsgPayPrincipal", value: MsgPayPrincipal.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgPayPrincipal: Could not create message: ' + e.message)
			}
		},
		
		msgTransferOwnership({ value }: msgTransferOwnershipParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.spv.MsgTransferOwnership", value: MsgTransferOwnership.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgTransferOwnership: Could not create message: ' + e.message)
			}
		},
		
		msgUpdatePool({ value }: msgUpdatePoolParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.spv.MsgUpdatePool", value: MsgUpdatePool.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgUpdatePool: Could not create message: ' + e.message)
			}
		},
		
		msgWithdrawPrincipal({ value }: msgWithdrawPrincipalParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.spv.MsgWithdrawPrincipal", value: MsgWithdrawPrincipal.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgWithdrawPrincipal: Could not create message: ' + e.message)
			}
		},
		
		msgAddInvestors({ value }: msgAddInvestorsParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.spv.MsgAddInvestors", value: MsgAddInvestors.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAddInvestors: Could not create message: ' + e.message)
			}
		},
		
		msgBorrow({ value }: msgBorrowParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.spv.MsgBorrow", value: MsgBorrow.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgBorrow: Could not create message: ' + e.message)
			}
		},
		
		msgSubmitWithdrawProposal({ value }: msgSubmitWithdrawProposalParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.spv.MsgSubmitWithdrawProposal", value: MsgSubmitWithdrawProposal.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgSubmitWithdrawProposal: Could not create message: ' + e.message)
			}
		},
		
		msgClaimInterest({ value }: msgClaimInterestParams): EncodeObject {
			try {
				return { typeUrl: "/joltify.spv.MsgClaimInterest", value: MsgClaimInterest.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgClaimInterest: Could not create message: ' + e.message)
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
	public structure: Record<string,unknown>;
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		this.structure =  {
						DepositorInfo: getStructure(typeDepositorInfo.fromPartial({})),
						BorrowDetail: getStructure(typeBorrowDetail.fromPartial({})),
						NftInfo: getStructure(typeNftInfo.fromPartial({})),
						PaymentItem: getStructure(typePaymentItem.fromPartial({})),
						BorrowInterest: getStructure(typeBorrowInterest.fromPartial({})),
						Params: getStructure(typeParams.fromPartial({})),
						PoolInfo: getStructure(typePoolInfo.fromPartial({})),
						PoolWithInvestors: getStructure(typePoolWithInvestors.fromPartial({})),
						PoolDepositedInvestors: getStructure(typePoolDepositedInvestors.fromPartial({})),
						WalletsLinkPool: getStructure(typeWalletsLinkPool.fromPartial({})),
						
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
			JoltifySpv: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;