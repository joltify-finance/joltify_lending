/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "joltifyfinance.joltify_lending.kyc";

export interface MsgUploadInvestor {
  creator: string;
  investorId: string;
  walletAddress: string[];
}

export interface MsgUploadInvestorResponse {
  wallets: string[];
}

function createBaseMsgUploadInvestor(): MsgUploadInvestor {
  return { creator: "", investorId: "", walletAddress: [] };
}

export const MsgUploadInvestor = {
  encode(message: MsgUploadInvestor, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.investorId !== "") {
      writer.uint32(18).string(message.investorId);
    }
    for (const v of message.walletAddress) {
      writer.uint32(26).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUploadInvestor {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUploadInvestor();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.investorId = reader.string();
          break;
        case 3:
          message.walletAddress.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUploadInvestor {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      investorId: isSet(object.investorId) ? String(object.investorId) : "",
      walletAddress: Array.isArray(object?.walletAddress) ? object.walletAddress.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: MsgUploadInvestor): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.investorId !== undefined && (obj.investorId = message.investorId);
    if (message.walletAddress) {
      obj.walletAddress = message.walletAddress.map((e) => e);
    } else {
      obj.walletAddress = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUploadInvestor>, I>>(object: I): MsgUploadInvestor {
    const message = createBaseMsgUploadInvestor();
    message.creator = object.creator ?? "";
    message.investorId = object.investorId ?? "";
    message.walletAddress = object.walletAddress?.map((e) => e) || [];
    return message;
  },
};

function createBaseMsgUploadInvestorResponse(): MsgUploadInvestorResponse {
  return { wallets: [] };
}

export const MsgUploadInvestorResponse = {
  encode(message: MsgUploadInvestorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.wallets) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUploadInvestorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUploadInvestorResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.wallets.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUploadInvestorResponse {
    return { wallets: Array.isArray(object?.wallets) ? object.wallets.map((e: any) => String(e)) : [] };
  },

  toJSON(message: MsgUploadInvestorResponse): unknown {
    const obj: any = {};
    if (message.wallets) {
      obj.wallets = message.wallets.map((e) => e);
    } else {
      obj.wallets = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUploadInvestorResponse>, I>>(object: I): MsgUploadInvestorResponse {
    const message = createBaseMsgUploadInvestorResponse();
    message.wallets = object.wallets?.map((e) => e) || [];
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  UploadInvestor(request: MsgUploadInvestor): Promise<MsgUploadInvestorResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.UploadInvestor = this.UploadInvestor.bind(this);
  }
  UploadInvestor(request: MsgUploadInvestor): Promise<MsgUploadInvestorResponse> {
    const data = MsgUploadInvestor.encode(request).finish();
    const promise = this.rpc.request("joltifyfinance.joltify_lending.kyc.Msg", "UploadInvestor", data);
    return promise.then((data) => MsgUploadInvestorResponse.decode(new _m0.Reader(data)));
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
