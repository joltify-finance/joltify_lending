/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "joltifyfinance.joltify_lending.spv";

export interface MsgTransferOwnership {
  creator: string;
  poolIndex: string;
}

export interface MsgTransferOwnershipResponse {
  operationResult: boolean;
}

function createBaseMsgTransferOwnership(): MsgTransferOwnership {
  return { creator: "", poolIndex: "" };
}

export const MsgTransferOwnership = {
  encode(message: MsgTransferOwnership, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgTransferOwnership {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgTransferOwnership();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.poolIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgTransferOwnership {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      poolIndex: isSet(object.poolIndex) ? String(object.poolIndex) : "",
    };
  },

  toJSON(message: MsgTransferOwnership): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.poolIndex !== undefined && (obj.poolIndex = message.poolIndex);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgTransferOwnership>, I>>(object: I): MsgTransferOwnership {
    const message = createBaseMsgTransferOwnership();
    message.creator = object.creator ?? "";
    message.poolIndex = object.poolIndex ?? "";
    return message;
  },
};

function createBaseMsgTransferOwnershipResponse(): MsgTransferOwnershipResponse {
  return { operationResult: false };
}

export const MsgTransferOwnershipResponse = {
  encode(message: MsgTransferOwnershipResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.operationResult === true) {
      writer.uint32(8).bool(message.operationResult);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgTransferOwnershipResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgTransferOwnershipResponse();
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

  fromJSON(object: any): MsgTransferOwnershipResponse {
    return { operationResult: isSet(object.operationResult) ? Boolean(object.operationResult) : false };
  },

  toJSON(message: MsgTransferOwnershipResponse): unknown {
    const obj: any = {};
    message.operationResult !== undefined && (obj.operationResult = message.operationResult);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgTransferOwnershipResponse>, I>>(object: I): MsgTransferOwnershipResponse {
    const message = createBaseMsgTransferOwnershipResponse();
    message.operationResult = object.operationResult ?? false;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  TransferOwnership(request: MsgTransferOwnership): Promise<MsgTransferOwnershipResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.TransferOwnership = this.TransferOwnership.bind(this);
  }
  TransferOwnership(request: MsgTransferOwnership): Promise<MsgTransferOwnershipResponse> {
    const data = MsgTransferOwnership.encode(request).finish();
    const promise = this.rpc.request("joltifyfinance.joltify_lending.spv.Msg", "TransferOwnership", data);
    return promise.then((data) => MsgTransferOwnershipResponse.decode(new _m0.Reader(data)));
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
