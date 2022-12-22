/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Timestamp } from "../../google/protobuf/timestamp";

export const protobufPackage = "joltify.mint";

export interface HistoricalDistInfo {
  payoutTime: Date | undefined;
  distributedRound: number;
}

function createBaseHistoricalDistInfo(): HistoricalDistInfo {
  return { payoutTime: undefined, distributedRound: 0 };
}

export const HistoricalDistInfo = {
  encode(message: HistoricalDistInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.payoutTime !== undefined) {
      Timestamp.encode(toTimestamp(message.payoutTime), writer.uint32(10).fork()).ldelim();
    }
    if (message.distributedRound !== 0) {
      writer.uint32(16).uint64(message.distributedRound);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): HistoricalDistInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHistoricalDistInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.payoutTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 2:
          message.distributedRound = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): HistoricalDistInfo {
    return {
      payoutTime: isSet(object.payoutTime) ? fromJsonTimestamp(object.payoutTime) : undefined,
      distributedRound: isSet(object.distributedRound) ? Number(object.distributedRound) : 0,
    };
  },

  toJSON(message: HistoricalDistInfo): unknown {
    const obj: any = {};
    message.payoutTime !== undefined && (obj.payoutTime = message.payoutTime.toISOString());
    message.distributedRound !== undefined && (obj.distributedRound = Math.round(message.distributedRound));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<HistoricalDistInfo>, I>>(object: I): HistoricalDistInfo {
    const message = createBaseHistoricalDistInfo();
    message.payoutTime = object.payoutTime ?? undefined;
    message.distributedRound = object.distributedRound ?? 0;
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

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

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
