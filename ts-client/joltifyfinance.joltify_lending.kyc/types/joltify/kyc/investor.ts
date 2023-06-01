/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "joltifyfinance.joltify_lending.kyc";

export interface Investor {
  investorId: string;
  walletAddress: string[];
}

function createBaseInvestor(): Investor {
  return { investorId: "", walletAddress: [] };
}

export const Investor = {
  encode(message: Investor, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.investorId !== "") {
      writer.uint32(10).string(message.investorId);
    }
    for (const v of message.walletAddress) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Investor {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInvestor();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.investorId = reader.string();
          break;
        case 2:
          message.walletAddress.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Investor {
    return {
      investorId: isSet(object.investorId) ? String(object.investorId) : "",
      walletAddress: Array.isArray(object?.walletAddress) ? object.walletAddress.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: Investor): unknown {
    const obj: any = {};
    message.investorId !== undefined && (obj.investorId = message.investorId);
    if (message.walletAddress) {
      obj.walletAddress = message.walletAddress.map((e) => e);
    } else {
      obj.walletAddress = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Investor>, I>>(object: I): Investor {
    const message = createBaseInvestor();
    message.investorId = object.investorId ?? "";
    message.walletAddress = object.walletAddress?.map((e) => e) || [];
    return message;
  },
};

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
