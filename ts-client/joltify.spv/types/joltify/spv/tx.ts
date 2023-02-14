/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "joltify.spv";

export interface MsgCreatePool {
  creator: string;
  projectIndex: number;
  poolName: string;
  apy: string;
  payFreq: string;
  targetAmount: string;
}

export interface MsgCreatePoolResponse {
  poolIndex: string;
}

export interface MsgAddInvestors {
  creator: string;
  poolIndex: string;
  invetorID: string[];
}

export interface MsgAddInvestorsResponse {
  operationResult: boolean;
}

function createBaseMsgCreatePool(): MsgCreatePool {
  return { creator: "", projectIndex: 0, poolName: "", apy: "", payFreq: "", targetAmount: "" };
}

export const MsgCreatePool = {
  encode(message: MsgCreatePool, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.projectIndex !== 0) {
      writer.uint32(16).int32(message.projectIndex);
    }
    if (message.poolName !== "") {
      writer.uint32(26).string(message.poolName);
    }
    if (message.apy !== "") {
      writer.uint32(34).string(message.apy);
    }
    if (message.payFreq !== "") {
      writer.uint32(42).string(message.payFreq);
    }
    if (message.targetAmount !== "") {
      writer.uint32(50).string(message.targetAmount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreatePool {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreatePool();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.projectIndex = reader.int32();
          break;
        case 3:
          message.poolName = reader.string();
          break;
        case 4:
          message.apy = reader.string();
          break;
        case 5:
          message.payFreq = reader.string();
          break;
        case 6:
          message.targetAmount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreatePool {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      projectIndex: isSet(object.projectIndex) ? Number(object.projectIndex) : 0,
      poolName: isSet(object.poolName) ? String(object.poolName) : "",
      apy: isSet(object.apy) ? String(object.apy) : "",
      payFreq: isSet(object.payFreq) ? String(object.payFreq) : "",
      targetAmount: isSet(object.targetAmount) ? String(object.targetAmount) : "",
    };
  },

  toJSON(message: MsgCreatePool): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.projectIndex !== undefined && (obj.projectIndex = Math.round(message.projectIndex));
    message.poolName !== undefined && (obj.poolName = message.poolName);
    message.apy !== undefined && (obj.apy = message.apy);
    message.payFreq !== undefined && (obj.payFreq = message.payFreq);
    message.targetAmount !== undefined && (obj.targetAmount = message.targetAmount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreatePool>, I>>(object: I): MsgCreatePool {
    const message = createBaseMsgCreatePool();
    message.creator = object.creator ?? "";
    message.projectIndex = object.projectIndex ?? 0;
    message.poolName = object.poolName ?? "";
    message.apy = object.apy ?? "";
    message.payFreq = object.payFreq ?? "";
    message.targetAmount = object.targetAmount ?? "";
    return message;
  },
};

function createBaseMsgCreatePoolResponse(): MsgCreatePoolResponse {
  return { poolIndex: "" };
}

export const MsgCreatePoolResponse = {
  encode(message: MsgCreatePoolResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.poolIndex !== "") {
      writer.uint32(10).string(message.poolIndex);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreatePoolResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreatePoolResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.poolIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreatePoolResponse {
    return { poolIndex: isSet(object.poolIndex) ? String(object.poolIndex) : "" };
  },

  toJSON(message: MsgCreatePoolResponse): unknown {
    const obj: any = {};
    message.poolIndex !== undefined && (obj.poolIndex = message.poolIndex);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreatePoolResponse>, I>>(object: I): MsgCreatePoolResponse {
    const message = createBaseMsgCreatePoolResponse();
    message.poolIndex = object.poolIndex ?? "";
    return message;
  },
};

function createBaseMsgAddInvestors(): MsgAddInvestors {
  return { creator: "", poolIndex: "", invetorID: [] };
}

export const MsgAddInvestors = {
  encode(message: MsgAddInvestors, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }
    for (const v of message.invetorID) {
      writer.uint32(26).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddInvestors {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddInvestors();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.poolIndex = reader.string();
          break;
        case 3:
          message.invetorID.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddInvestors {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      poolIndex: isSet(object.poolIndex) ? String(object.poolIndex) : "",
      invetorID: Array.isArray(object?.invetorID) ? object.invetorID.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: MsgAddInvestors): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.poolIndex !== undefined && (obj.poolIndex = message.poolIndex);
    if (message.invetorID) {
      obj.invetorID = message.invetorID.map((e) => e);
    } else {
      obj.invetorID = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddInvestors>, I>>(object: I): MsgAddInvestors {
    const message = createBaseMsgAddInvestors();
    message.creator = object.creator ?? "";
    message.poolIndex = object.poolIndex ?? "";
    message.invetorID = object.invetorID?.map((e) => e) || [];
    return message;
  },
};

function createBaseMsgAddInvestorsResponse(): MsgAddInvestorsResponse {
  return { operationResult: false };
}

export const MsgAddInvestorsResponse = {
  encode(message: MsgAddInvestorsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.operationResult === true) {
      writer.uint32(8).bool(message.operationResult);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddInvestorsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddInvestorsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.operationResult = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddInvestorsResponse {
    return { operationResult: isSet(object.operationResult) ? Boolean(object.operationResult) : false };
  },

  toJSON(message: MsgAddInvestorsResponse): unknown {
    const obj: any = {};
    message.operationResult !== undefined && (obj.operationResult = message.operationResult);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddInvestorsResponse>, I>>(object: I): MsgAddInvestorsResponse {
    const message = createBaseMsgAddInvestorsResponse();
    message.operationResult = object.operationResult ?? false;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreatePool(request: MsgCreatePool): Promise<MsgCreatePoolResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  AddInvestors(request: MsgAddInvestors): Promise<MsgAddInvestorsResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreatePool = this.CreatePool.bind(this);
    this.AddInvestors = this.AddInvestors.bind(this);
  }
  CreatePool(request: MsgCreatePool): Promise<MsgCreatePoolResponse> {
    const data = MsgCreatePool.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "CreatePool", data);
    return promise.then((data) => MsgCreatePoolResponse.decode(new _m0.Reader(data)));
  }

  AddInvestors(request: MsgAddInvestors): Promise<MsgAddInvestorsResponse> {
    const data = MsgAddInvestors.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "AddInvestors", data);
    return promise.then((data) => MsgAddInvestorsResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
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
