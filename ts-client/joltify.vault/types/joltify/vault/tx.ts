/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "joltify.vault";

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

export interface MsgCreateOutboundTxResponse {
  successful: boolean;
}

/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCreateIssueToken {
  creator: Uint8Array;
  index: string;
  coin: Uint8Array;
  receiver: Uint8Array;
}

export interface MsgCreateIssueTokenResponse {
  successful: boolean;
}

export interface MsgCreateCreatePool {
  creator: Uint8Array;
  poolPubKey: string;
  blockHeight: string;
}

export interface MsgCreateCreatePoolResponse {
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
    receiverAddress: new Uint8Array(),
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

  fromJSON(object: any): MsgCreateOutboundTx {
    return {
      creator: isSet(object.creator) ? bytesFromBase64(object.creator) : new Uint8Array(),
      requestID: isSet(object.requestID) ? String(object.requestID) : "",
      outboundTx: isSet(object.outboundTx) ? String(object.outboundTx) : "",
      blockHeight: isSet(object.blockHeight) ? String(object.blockHeight) : "",
      feecoin: Array.isArray(object?.feecoin) ? object.feecoin.map((e: any) => Coin.fromJSON(e)) : [],
      chainType: isSet(object.chainType) ? String(object.chainType) : "",
      needMint: isSet(object.needMint) ? Boolean(object.needMint) : false,
      inTxHash: isSet(object.inTxHash) ? String(object.inTxHash) : "",
      receiverAddress: isSet(object.receiverAddress) ? bytesFromBase64(object.receiverAddress) : new Uint8Array(),
    };
  },

  toJSON(message: MsgCreateOutboundTx): unknown {
    const obj: any = {};
    message.creator !== undefined
      && (obj.creator = base64FromBytes(message.creator !== undefined ? message.creator : new Uint8Array()));
    message.requestID !== undefined && (obj.requestID = message.requestID);
    message.outboundTx !== undefined && (obj.outboundTx = message.outboundTx);
    message.blockHeight !== undefined && (obj.blockHeight = message.blockHeight);
    if (message.feecoin) {
      obj.feecoin = message.feecoin.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.feecoin = [];
    }
    message.chainType !== undefined && (obj.chainType = message.chainType);
    message.needMint !== undefined && (obj.needMint = message.needMint);
    message.inTxHash !== undefined && (obj.inTxHash = message.inTxHash);
    message.receiverAddress !== undefined
      && (obj.receiverAddress = base64FromBytes(
        message.receiverAddress !== undefined ? message.receiverAddress : new Uint8Array(),
      ));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateOutboundTx>, I>>(object: I): MsgCreateOutboundTx {
    const message = createBaseMsgCreateOutboundTx();
    message.creator = object.creator ?? new Uint8Array();
    message.requestID = object.requestID ?? "";
    message.outboundTx = object.outboundTx ?? "";
    message.blockHeight = object.blockHeight ?? "";
    message.feecoin = object.feecoin?.map((e) => Coin.fromPartial(e)) || [];
    message.chainType = object.chainType ?? "";
    message.needMint = object.needMint ?? false;
    message.inTxHash = object.inTxHash ?? "";
    message.receiverAddress = object.receiverAddress ?? new Uint8Array();
    return message;
  },
};

function createBaseMsgCreateOutboundTxResponse(): MsgCreateOutboundTxResponse {
  return { successful: false };
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

  fromJSON(object: any): MsgCreateOutboundTxResponse {
    return { successful: isSet(object.successful) ? Boolean(object.successful) : false };
  },

  toJSON(message: MsgCreateOutboundTxResponse): unknown {
    const obj: any = {};
    message.successful !== undefined && (obj.successful = message.successful);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateOutboundTxResponse>, I>>(object: I): MsgCreateOutboundTxResponse {
    const message = createBaseMsgCreateOutboundTxResponse();
    message.successful = object.successful ?? false;
    return message;
  },
};

function createBaseMsgCreateIssueToken(): MsgCreateIssueToken {
  return { creator: new Uint8Array(), index: "", coin: new Uint8Array(), receiver: new Uint8Array() };
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

  fromJSON(object: any): MsgCreateIssueToken {
    return {
      creator: isSet(object.creator) ? bytesFromBase64(object.creator) : new Uint8Array(),
      index: isSet(object.index) ? String(object.index) : "",
      coin: isSet(object.coin) ? bytesFromBase64(object.coin) : new Uint8Array(),
      receiver: isSet(object.receiver) ? bytesFromBase64(object.receiver) : new Uint8Array(),
    };
  },

  toJSON(message: MsgCreateIssueToken): unknown {
    const obj: any = {};
    message.creator !== undefined
      && (obj.creator = base64FromBytes(message.creator !== undefined ? message.creator : new Uint8Array()));
    message.index !== undefined && (obj.index = message.index);
    message.coin !== undefined
      && (obj.coin = base64FromBytes(message.coin !== undefined ? message.coin : new Uint8Array()));
    message.receiver !== undefined
      && (obj.receiver = base64FromBytes(message.receiver !== undefined ? message.receiver : new Uint8Array()));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateIssueToken>, I>>(object: I): MsgCreateIssueToken {
    const message = createBaseMsgCreateIssueToken();
    message.creator = object.creator ?? new Uint8Array();
    message.index = object.index ?? "";
    message.coin = object.coin ?? new Uint8Array();
    message.receiver = object.receiver ?? new Uint8Array();
    return message;
  },
};

function createBaseMsgCreateIssueTokenResponse(): MsgCreateIssueTokenResponse {
  return { successful: false };
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

  fromJSON(object: any): MsgCreateIssueTokenResponse {
    return { successful: isSet(object.successful) ? Boolean(object.successful) : false };
  },

  toJSON(message: MsgCreateIssueTokenResponse): unknown {
    const obj: any = {};
    message.successful !== undefined && (obj.successful = message.successful);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateIssueTokenResponse>, I>>(object: I): MsgCreateIssueTokenResponse {
    const message = createBaseMsgCreateIssueTokenResponse();
    message.successful = object.successful ?? false;
    return message;
  },
};

function createBaseMsgCreateCreatePool(): MsgCreateCreatePool {
  return { creator: new Uint8Array(), poolPubKey: "", blockHeight: "" };
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

  fromJSON(object: any): MsgCreateCreatePool {
    return {
      creator: isSet(object.creator) ? bytesFromBase64(object.creator) : new Uint8Array(),
      poolPubKey: isSet(object.poolPubKey) ? String(object.poolPubKey) : "",
      blockHeight: isSet(object.blockHeight) ? String(object.blockHeight) : "",
    };
  },

  toJSON(message: MsgCreateCreatePool): unknown {
    const obj: any = {};
    message.creator !== undefined
      && (obj.creator = base64FromBytes(message.creator !== undefined ? message.creator : new Uint8Array()));
    message.poolPubKey !== undefined && (obj.poolPubKey = message.poolPubKey);
    message.blockHeight !== undefined && (obj.blockHeight = message.blockHeight);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateCreatePool>, I>>(object: I): MsgCreateCreatePool {
    const message = createBaseMsgCreateCreatePool();
    message.creator = object.creator ?? new Uint8Array();
    message.poolPubKey = object.poolPubKey ?? "";
    message.blockHeight = object.blockHeight ?? "";
    return message;
  },
};

function createBaseMsgCreateCreatePoolResponse(): MsgCreateCreatePoolResponse {
  return { successful: false };
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

  fromJSON(object: any): MsgCreateCreatePoolResponse {
    return { successful: isSet(object.successful) ? Boolean(object.successful) : false };
  },

  toJSON(message: MsgCreateCreatePoolResponse): unknown {
    const obj: any = {};
    message.successful !== undefined && (obj.successful = message.successful);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateCreatePoolResponse>, I>>(object: I): MsgCreateCreatePoolResponse {
    const message = createBaseMsgCreateCreatePoolResponse();
    message.successful = object.successful ?? false;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateOutboundTx(request: MsgCreateOutboundTx): Promise<MsgCreateOutboundTxResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateIssueToken(request: MsgCreateIssueToken): Promise<MsgCreateIssueTokenResponse>;
  CreateCreatePool(request: MsgCreateCreatePool): Promise<MsgCreateCreatePoolResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateOutboundTx = this.CreateOutboundTx.bind(this);
    this.CreateIssueToken = this.CreateIssueToken.bind(this);
    this.CreateCreatePool = this.CreateCreatePool.bind(this);
  }
  CreateOutboundTx(request: MsgCreateOutboundTx): Promise<MsgCreateOutboundTxResponse> {
    const data = MsgCreateOutboundTx.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Msg", "CreateOutboundTx", data);
    return promise.then((data) => MsgCreateOutboundTxResponse.decode(new _m0.Reader(data)));
  }

  CreateIssueToken(request: MsgCreateIssueToken): Promise<MsgCreateIssueTokenResponse> {
    const data = MsgCreateIssueToken.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Msg", "CreateIssueToken", data);
    return promise.then((data) => MsgCreateIssueTokenResponse.decode(new _m0.Reader(data)));
  }

  CreateCreatePool(request: MsgCreateCreatePool): Promise<MsgCreateCreatePoolResponse> {
    const data = MsgCreateCreatePool.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Msg", "CreateCreatePool", data);
    return promise.then((data) => MsgCreateCreatePoolResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

function bytesFromBase64(b64: string): Uint8Array {
  if (globalThis.Buffer) {
    return Uint8Array.from(globalThis.Buffer.from(b64, "base64"));
  } else {
    const bin = globalThis.atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
      arr[i] = bin.charCodeAt(i);
    }
    return arr;
  }
}

function base64FromBytes(arr: Uint8Array): string {
  if (globalThis.Buffer) {
    return globalThis.Buffer.from(arr).toString("base64");
  } else {
    const bin: string[] = [];
    arr.forEach((byte) => {
      bin.push(String.fromCharCode(byte));
    });
    return globalThis.btoa(bin.join(""));
  }
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
