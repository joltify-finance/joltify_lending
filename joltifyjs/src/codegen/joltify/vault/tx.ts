import { Coin, CoinSDKType } from "../../cosmos/base/v1beta1/coin";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
export interface MsgCreateOutboundTx {
  creator: Uint8Array;
  requestID: string;
  outboundTx: string;
  blockHeight: string;
  feecoin: Coin[];
  chainType: string;
  needMint: boolean;
  inTxHash: string;
  receiverAddress: Uint8Array;
}
export interface MsgCreateOutboundTxSDKType {
  creator: Uint8Array;
  request_iD: string;
  outbound_tx: string;
  block_height: string;
  feecoin: CoinSDKType[];
  chain_type: string;
  need_mint: boolean;
  in_tx_hash: string;
  receiver_address: Uint8Array;
}
export interface MsgCreateOutboundTxResponse {
  successful: boolean;
}
export interface MsgCreateOutboundTxResponseSDKType {
  successful: boolean;
}
/** this line is used by starport scaffolding # proto/tx/message */

export interface MsgCreateIssueToken {
  creator: Uint8Array;
  index: string;
  coin: Uint8Array;
  receiver: Uint8Array;
}
/** this line is used by starport scaffolding # proto/tx/message */

export interface MsgCreateIssueTokenSDKType {
  creator: Uint8Array;
  index: string;
  coin: Uint8Array;
  receiver: Uint8Array;
}
export interface MsgCreateIssueTokenResponse {
  successful: boolean;
}
export interface MsgCreateIssueTokenResponseSDKType {
  successful: boolean;
}
export interface MsgCreateCreatePool {
  creator: Uint8Array;
  poolPubKey: string;
  blockHeight: string;
}
export interface MsgCreateCreatePoolSDKType {
  creator: Uint8Array;
  pool_pub_key: string;
  block_height: string;
}
export interface MsgCreateCreatePoolResponse {
  successful: boolean;
}
export interface MsgCreateCreatePoolResponseSDKType {
  successful: boolean;
}

function createBaseMsgCreateOutboundTx(): MsgCreateOutboundTx {
  return {
    creator: new Uint8Array(),
    requestID: "",
    outboundTx: "",
    blockHeight: "",
    feecoin: [],
    chainType: "",
    needMint: false,
    inTxHash: "",
    receiverAddress: new Uint8Array()
  };
}

export const MsgCreateOutboundTx = {
  encode(message: MsgCreateOutboundTx, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator.length !== 0) {
      writer.uint32(10).bytes(message.creator);
    }

    if (message.requestID !== "") {
      writer.uint32(18).string(message.requestID);
    }

    if (message.outboundTx !== "") {
      writer.uint32(26).string(message.outboundTx);
    }

    if (message.blockHeight !== "") {
      writer.uint32(34).string(message.blockHeight);
    }

    for (const v of message.feecoin) {
      Coin.encode(v!, writer.uint32(42).fork()).ldelim();
    }

    if (message.chainType !== "") {
      writer.uint32(50).string(message.chainType);
    }

    if (message.needMint === true) {
      writer.uint32(56).bool(message.needMint);
    }

    if (message.inTxHash !== "") {
      writer.uint32(66).string(message.inTxHash);
    }

    if (message.receiverAddress.length !== 0) {
      writer.uint32(74).bytes(message.receiverAddress);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateOutboundTx {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateOutboundTx();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.creator = reader.bytes();
          break;

        case 2:
          message.requestID = reader.string();
          break;

        case 3:
          message.outboundTx = reader.string();
          break;

        case 4:
          message.blockHeight = reader.string();
          break;

        case 5:
          message.feecoin.push(Coin.decode(reader, reader.uint32()));
          break;

        case 6:
          message.chainType = reader.string();
          break;

        case 7:
          message.needMint = reader.bool();
          break;

        case 8:
          message.inTxHash = reader.string();
          break;

        case 9:
          message.receiverAddress = reader.bytes();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgCreateOutboundTx>): MsgCreateOutboundTx {
    const message = createBaseMsgCreateOutboundTx();
    message.creator = object.creator ?? new Uint8Array();
    message.requestID = object.requestID ?? "";
    message.outboundTx = object.outboundTx ?? "";
    message.blockHeight = object.blockHeight ?? "";
    message.feecoin = object.feecoin?.map(e => Coin.fromPartial(e)) || [];
    message.chainType = object.chainType ?? "";
    message.needMint = object.needMint ?? false;
    message.inTxHash = object.inTxHash ?? "";
    message.receiverAddress = object.receiverAddress ?? new Uint8Array();
    return message;
  }

};

function createBaseMsgCreateOutboundTxResponse(): MsgCreateOutboundTxResponse {
  return {
    successful: false
  };
}

export const MsgCreateOutboundTxResponse = {
  encode(message: MsgCreateOutboundTxResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.successful === true) {
      writer.uint32(8).bool(message.successful);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateOutboundTxResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateOutboundTxResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.successful = reader.bool();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgCreateOutboundTxResponse>): MsgCreateOutboundTxResponse {
    const message = createBaseMsgCreateOutboundTxResponse();
    message.successful = object.successful ?? false;
    return message;
  }

};

function createBaseMsgCreateIssueToken(): MsgCreateIssueToken {
  return {
    creator: new Uint8Array(),
    index: "",
    coin: new Uint8Array(),
    receiver: new Uint8Array()
  };
}

export const MsgCreateIssueToken = {
  encode(message: MsgCreateIssueToken, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator.length !== 0) {
      writer.uint32(10).bytes(message.creator);
    }

    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }

    if (message.coin.length !== 0) {
      writer.uint32(26).bytes(message.coin);
    }

    if (message.receiver.length !== 0) {
      writer.uint32(34).bytes(message.receiver);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateIssueToken {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateIssueToken();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.creator = reader.bytes();
          break;

        case 2:
          message.index = reader.string();
          break;

        case 3:
          message.coin = reader.bytes();
          break;

        case 4:
          message.receiver = reader.bytes();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgCreateIssueToken>): MsgCreateIssueToken {
    const message = createBaseMsgCreateIssueToken();
    message.creator = object.creator ?? new Uint8Array();
    message.index = object.index ?? "";
    message.coin = object.coin ?? new Uint8Array();
    message.receiver = object.receiver ?? new Uint8Array();
    return message;
  }

};

function createBaseMsgCreateIssueTokenResponse(): MsgCreateIssueTokenResponse {
  return {
    successful: false
  };
}

export const MsgCreateIssueTokenResponse = {
  encode(message: MsgCreateIssueTokenResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.successful === true) {
      writer.uint32(8).bool(message.successful);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateIssueTokenResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateIssueTokenResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.successful = reader.bool();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgCreateIssueTokenResponse>): MsgCreateIssueTokenResponse {
    const message = createBaseMsgCreateIssueTokenResponse();
    message.successful = object.successful ?? false;
    return message;
  }

};

function createBaseMsgCreateCreatePool(): MsgCreateCreatePool {
  return {
    creator: new Uint8Array(),
    poolPubKey: "",
    blockHeight: ""
  };
}

export const MsgCreateCreatePool = {
  encode(message: MsgCreateCreatePool, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator.length !== 0) {
      writer.uint32(10).bytes(message.creator);
    }

    if (message.poolPubKey !== "") {
      writer.uint32(18).string(message.poolPubKey);
    }

    if (message.blockHeight !== "") {
      writer.uint32(26).string(message.blockHeight);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateCreatePool {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateCreatePool();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.creator = reader.bytes();
          break;

        case 2:
          message.poolPubKey = reader.string();
          break;

        case 3:
          message.blockHeight = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgCreateCreatePool>): MsgCreateCreatePool {
    const message = createBaseMsgCreateCreatePool();
    message.creator = object.creator ?? new Uint8Array();
    message.poolPubKey = object.poolPubKey ?? "";
    message.blockHeight = object.blockHeight ?? "";
    return message;
  }

};

function createBaseMsgCreateCreatePoolResponse(): MsgCreateCreatePoolResponse {
  return {
    successful: false
  };
}

export const MsgCreateCreatePoolResponse = {
  encode(message: MsgCreateCreatePoolResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.successful === true) {
      writer.uint32(8).bool(message.successful);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateCreatePoolResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateCreatePoolResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.successful = reader.bool();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgCreateCreatePoolResponse>): MsgCreateCreatePoolResponse {
    const message = createBaseMsgCreateCreatePoolResponse();
    message.successful = object.successful ?? false;
    return message;
  }

};