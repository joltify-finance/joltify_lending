/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "joltifyfinance.joltify_lending.spv";

export interface MsgBorrow {
  creator: string;
  poolIndex: string;
  borrowAmount: string;
}

export interface MsgBorrowResponse {
}

function createBaseMsgBorrow(): MsgBorrow {
  return { creator: "", poolIndex: "", borrowAmount: "" };
}

export const MsgBorrow = {
  encode(message: MsgBorrow, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }
    if (message.borrowAmount !== "") {
      writer.uint32(26).string(message.borrowAmount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBorrow {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBorrow();
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
          message.borrowAmount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgBorrow {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      poolIndex: isSet(object.poolIndex) ? String(object.poolIndex) : "",
      borrowAmount: isSet(object.borrowAmount) ? String(object.borrowAmount) : "",
    };
  },

  toJSON(message: MsgBorrow): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.poolIndex !== undefined && (obj.poolIndex = message.poolIndex);
    message.borrowAmount !== undefined && (obj.borrowAmount = message.borrowAmount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgBorrow>, I>>(object: I): MsgBorrow {
    const message = createBaseMsgBorrow();
    message.creator = object.creator ?? "";
    message.poolIndex = object.poolIndex ?? "";
    message.borrowAmount = object.borrowAmount ?? "";
    return message;
  },
};

function createBaseMsgBorrowResponse(): MsgBorrowResponse {
  return {};
}

export const MsgBorrowResponse = {
  encode(_: MsgBorrowResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBorrowResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBorrowResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgBorrowResponse {
    return {};
  },

  toJSON(_: MsgBorrowResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgBorrowResponse>, I>>(_: I): MsgBorrowResponse {
    const message = createBaseMsgBorrowResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Borrow(request: MsgBorrow): Promise<MsgBorrowResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Borrow = this.Borrow.bind(this);
  }
  Borrow(request: MsgBorrow): Promise<MsgBorrowResponse> {
    const data = MsgBorrow.encode(request).finish();
    const promise = this.rpc.request("joltifyfinance.joltify_lending.spv.Msg", "Borrow", data);
    return promise.then((data) => MsgBorrowResponse.decode(new _m0.Reader(data)));
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
