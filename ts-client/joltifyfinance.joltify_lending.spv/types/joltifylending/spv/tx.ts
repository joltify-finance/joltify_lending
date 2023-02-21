/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "joltifyfinance.joltify_lending.spv";

export interface MsgRepayInterest {
  creator: string;
  borrowAmount: string;
}

export interface MsgRepayInterestResponse {
}

export interface MsgClaimInterest {
  creator: string;
  borrowAmount: string;
}

export interface MsgClaimInterestResponse {
}

function createBaseMsgRepayInterest(): MsgRepayInterest {
  return { creator: "", borrowAmount: "" };
}

export const MsgRepayInterest = {
  encode(message: MsgRepayInterest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.borrowAmount !== "") {
      writer.uint32(18).string(message.borrowAmount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRepayInterest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRepayInterest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.borrowAmount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRepayInterest {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      borrowAmount: isSet(object.borrowAmount) ? String(object.borrowAmount) : "",
    };
  },

  toJSON(message: MsgRepayInterest): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.borrowAmount !== undefined && (obj.borrowAmount = message.borrowAmount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRepayInterest>, I>>(object: I): MsgRepayInterest {
    const message = createBaseMsgRepayInterest();
    message.creator = object.creator ?? "";
    message.borrowAmount = object.borrowAmount ?? "";
    return message;
  },
};

function createBaseMsgRepayInterestResponse(): MsgRepayInterestResponse {
  return {};
}

export const MsgRepayInterestResponse = {
  encode(_: MsgRepayInterestResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRepayInterestResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRepayInterestResponse();
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

  fromJSON(_: any): MsgRepayInterestResponse {
    return {};
  },

  toJSON(_: MsgRepayInterestResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRepayInterestResponse>, I>>(_: I): MsgRepayInterestResponse {
    const message = createBaseMsgRepayInterestResponse();
    return message;
  },
};

function createBaseMsgClaimInterest(): MsgClaimInterest {
  return { creator: "", borrowAmount: "" };
}

export const MsgClaimInterest = {
  encode(message: MsgClaimInterest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.borrowAmount !== "") {
      writer.uint32(18).string(message.borrowAmount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimInterest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimInterest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.borrowAmount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgClaimInterest {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      borrowAmount: isSet(object.borrowAmount) ? String(object.borrowAmount) : "",
    };
  },

  toJSON(message: MsgClaimInterest): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.borrowAmount !== undefined && (obj.borrowAmount = message.borrowAmount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimInterest>, I>>(object: I): MsgClaimInterest {
    const message = createBaseMsgClaimInterest();
    message.creator = object.creator ?? "";
    message.borrowAmount = object.borrowAmount ?? "";
    return message;
  },
};

function createBaseMsgClaimInterestResponse(): MsgClaimInterestResponse {
  return {};
}

export const MsgClaimInterestResponse = {
  encode(_: MsgClaimInterestResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimInterestResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimInterestResponse();
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

  fromJSON(_: any): MsgClaimInterestResponse {
    return {};
  },

  toJSON(_: MsgClaimInterestResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimInterestResponse>, I>>(_: I): MsgClaimInterestResponse {
    const message = createBaseMsgClaimInterestResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  RepayInterest(request: MsgRepayInterest): Promise<MsgRepayInterestResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  ClaimInterest(request: MsgClaimInterest): Promise<MsgClaimInterestResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.RepayInterest = this.RepayInterest.bind(this);
    this.ClaimInterest = this.ClaimInterest.bind(this);
  }
  RepayInterest(request: MsgRepayInterest): Promise<MsgRepayInterestResponse> {
    const data = MsgRepayInterest.encode(request).finish();
    const promise = this.rpc.request("joltifyfinance.joltify_lending.spv.Msg", "RepayInterest", data);
    return promise.then((data) => MsgRepayInterestResponse.decode(new _m0.Reader(data)));
  }

  ClaimInterest(request: MsgClaimInterest): Promise<MsgClaimInterestResponse> {
    const data = MsgClaimInterest.encode(request).finish();
    const promise = this.rpc.request("joltifyfinance.joltify_lending.spv.Msg", "ClaimInterest", data);
    return promise.then((data) => MsgClaimInterestResponse.decode(new _m0.Reader(data)));
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
