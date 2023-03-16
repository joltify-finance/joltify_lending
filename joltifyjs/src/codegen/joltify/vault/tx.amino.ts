import { AminoMsg } from "@cosmjs/amino";
import { MsgCreateOutboundTx, MsgCreateIssueToken, MsgCreateCreatePool } from "./tx";
export interface MsgCreateOutboundTxAminoType extends AminoMsg {
  type: "/joltify.vault.MsgCreateOutboundTx";
  value: {
    creator: Uint8Array;
    request_iD: string;
    outbound_tx: string;
    block_height: string;
    feecoin: {
      denom: string;
      amount: string;
    }[];
    chain_type: string;
    need_mint: boolean;
    in_tx_hash: string;
    receiver_address: Uint8Array;
  };
}
export interface MsgCreateIssueTokenAminoType extends AminoMsg {
  type: "/joltify.vault.MsgCreateIssueToken";
  value: {
    creator: Uint8Array;
    index: string;
    coin: Uint8Array;
    receiver: Uint8Array;
  };
}
export interface MsgCreateCreatePoolAminoType extends AminoMsg {
  type: "/joltify.vault.MsgCreateCreatePool";
  value: {
    creator: Uint8Array;
    pool_pub_key: string;
    block_height: string;
  };
}
export const AminoConverter = {
  "/joltify.vault.MsgCreateOutboundTx": {
    aminoType: "/joltify.vault.MsgCreateOutboundTx",
    toAmino: ({
      creator,
      requestID,
      outboundTx,
      blockHeight,
      feecoin,
      chainType,
      needMint,
      inTxHash,
      receiverAddress
    }: MsgCreateOutboundTx): MsgCreateOutboundTxAminoType["value"] => {
      return {
        creator,
        request_iD: requestID,
        outbound_tx: outboundTx,
        block_height: blockHeight,
        feecoin: feecoin.map(el0 => ({
          denom: el0.denom,
          amount: el0.amount
        })),
        chain_type: chainType,
        need_mint: needMint,
        in_tx_hash: inTxHash,
        receiver_address: receiverAddress
      };
    },
    fromAmino: ({
      creator,
      request_iD,
      outbound_tx,
      block_height,
      feecoin,
      chain_type,
      need_mint,
      in_tx_hash,
      receiver_address
    }: MsgCreateOutboundTxAminoType["value"]): MsgCreateOutboundTx => {
      return {
        creator,
        requestID: request_iD,
        outboundTx: outbound_tx,
        blockHeight: block_height,
        feecoin: feecoin.map(el0 => ({
          denom: el0.denom,
          amount: el0.amount
        })),
        chainType: chain_type,
        needMint: need_mint,
        inTxHash: in_tx_hash,
        receiverAddress: receiver_address
      };
    }
  },
  "/joltify.vault.MsgCreateIssueToken": {
    aminoType: "/joltify.vault.MsgCreateIssueToken",
    toAmino: ({
      creator,
      index,
      coin,
      receiver
    }: MsgCreateIssueToken): MsgCreateIssueTokenAminoType["value"] => {
      return {
        creator,
        index,
        coin,
        receiver
      };
    },
    fromAmino: ({
      creator,
      index,
      coin,
      receiver
    }: MsgCreateIssueTokenAminoType["value"]): MsgCreateIssueToken => {
      return {
        creator,
        index,
        coin,
        receiver
      };
    }
  },
  "/joltify.vault.MsgCreateCreatePool": {
    aminoType: "/joltify.vault.MsgCreateCreatePool",
    toAmino: ({
      creator,
      poolPubKey,
      blockHeight
    }: MsgCreateCreatePool): MsgCreateCreatePoolAminoType["value"] => {
      return {
        creator,
        pool_pub_key: poolPubKey,
        block_height: blockHeight
      };
    },
    fromAmino: ({
      creator,
      pool_pub_key,
      block_height
    }: MsgCreateCreatePoolAminoType["value"]): MsgCreateCreatePool => {
      return {
        creator,
        poolPubKey: pool_pub_key,
        blockHeight: block_height
      };
    }
  }
};