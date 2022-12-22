/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "joltify.vault";

export interface historicalAmount {
  blockHeight: number;
  amount: Coin[];
}

export interface coinsQuota {
  history: historicalAmount[];
  CoinsSum: Coin[];
}

function createBasehistoricalAmount(): historicalAmount {
  return { blockHeight: 0, amount: [] };
}

export const historicalAmount = {
  encode(message: historicalAmount, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.blockHeight !== 0) {
      writer.uint32(8).int64(message.blockHeight);
    }
    for (const v of message.amount) {
      Coin.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): historicalAmount {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasehistoricalAmount();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.blockHeight = longToNumber(reader.int64() as Long);
          break;
        case 2:
          message.amount.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): historicalAmount {
    return {
      blockHeight: isSet(object.blockHeight) ? Number(object.blockHeight) : 0,
      amount: Array.isArray(object?.amount) ? object.amount.map((e: any) => Coin.fromJSON(e)) : [],
    };
  },

  toJSON(message: historicalAmount): unknown {
    const obj: any = {};
    message.blockHeight !== undefined && (obj.blockHeight = Math.round(message.blockHeight));
    if (message.amount) {
      obj.amount = message.amount.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.amount = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<historicalAmount>, I>>(object: I): historicalAmount {
    const message = createBasehistoricalAmount();
    message.blockHeight = object.blockHeight ?? 0;
    message.amount = object.amount?.map((e) => Coin.fromPartial(e)) || [];
    return message;
  },
};

function createBasecoinsQuota(): coinsQuota {
  return { history: [], CoinsSum: [] };
}

export const coinsQuota = {
  encode(message: coinsQuota, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.history) {
      historicalAmount.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.CoinsSum) {
      Coin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): coinsQuota {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasecoinsQuota();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          message.history.push(historicalAmount.decode(reader, reader.uint32()));
          break;
        case 3:
          message.CoinsSum.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): coinsQuota {
    return {
      history: Array.isArray(object?.history) ? object.history.map((e: any) => historicalAmount.fromJSON(e)) : [],
      CoinsSum: Array.isArray(object?.CoinsSum) ? object.CoinsSum.map((e: any) => Coin.fromJSON(e)) : [],
    };
  },

  toJSON(message: coinsQuota): unknown {
    const obj: any = {};
    if (message.history) {
      obj.history = message.history.map((e) => e ? historicalAmount.toJSON(e) : undefined);
    } else {
      obj.history = [];
    }
    if (message.CoinsSum) {
      obj.CoinsSum = message.CoinsSum.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.CoinsSum = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<coinsQuota>, I>>(object: I): coinsQuota {
    const message = createBasecoinsQuota();
    message.history = object.history?.map((e) => historicalAmount.fromPartial(e)) || [];
    message.CoinsSum = object.CoinsSum?.map((e) => Coin.fromPartial(e)) || [];
    return message;
  },
};

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

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
