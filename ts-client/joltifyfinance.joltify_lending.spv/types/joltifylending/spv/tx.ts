/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "joltifyfinance.joltify_lending.spv";

export interface MsgAddInvestors {
  creator: string;
  walletAddress: string;
}

export interface MsgAddInvestorsResponse {
  operationResult: boolean;
}

function createBaseMsgAddInvestors(): MsgAddInvestors {
  return { creator: "", walletAddress: "" };
}

export const MsgAddInvestors = {
  encode(message: MsgAddInvestors, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.walletAddress !== "") {
      writer.uint32(18).string(message.walletAddress);
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
          message.walletAddress = reader.string();
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
      walletAddress: isSet(object.walletAddress) ? String(object.walletAddress) : "",
    };
  },

  toJSON(message: MsgAddInvestors): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.walletAddress !== undefined && (obj.walletAddress = message.walletAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddInvestors>, I>>(object: I): MsgAddInvestors {
    const message = createBaseMsgAddInvestors();
    message.creator = object.creator ?? "";
    message.walletAddress = object.walletAddress ?? "";
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
  /** this line is used by starport scaffolding # proto/tx/rpc */
  AddInvestors(request: MsgAddInvestors): Promise<MsgAddInvestorsResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.AddInvestors = this.AddInvestors.bind(this);
  }
  AddInvestors(request: MsgAddInvestors): Promise<MsgAddInvestorsResponse> {
    const data = MsgAddInvestors.encode(request).finish();
    const promise = this.rpc.request("joltifyfinance.joltify_lending.spv.Msg", "AddInvestors", data);
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
