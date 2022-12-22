/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "joltify.mint";

/** Params defines the parameters for the module. */
export interface Params {
  halfCount: number;
  firstProvisions: string;
  currentProvisions: string;
  unit: string;
  communityProvisions: string;
}

function createBaseParams(): Params {
  return { halfCount: 0, firstProvisions: "", currentProvisions: "", unit: "", communityProvisions: "" };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.halfCount !== 0) {
      writer.uint32(8).uint64(message.halfCount);
    }
    if (message.firstProvisions !== "") {
      writer.uint32(18).string(message.firstProvisions);
    }
    if (message.currentProvisions !== "") {
      writer.uint32(26).string(message.currentProvisions);
    }
    if (message.unit !== "") {
      writer.uint32(34).string(message.unit);
    }
    if (message.communityProvisions !== "") {
      writer.uint32(42).string(message.communityProvisions);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseParams();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.halfCount = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.firstProvisions = reader.string();
          break;
        case 3:
          message.currentProvisions = reader.string();
          break;
        case 4:
          message.unit = reader.string();
          break;
        case 5:
          message.communityProvisions = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    return {
      halfCount: isSet(object.halfCount) ? Number(object.halfCount) : 0,
      firstProvisions: isSet(object.firstProvisions) ? String(object.firstProvisions) : "",
      currentProvisions: isSet(object.currentProvisions) ? String(object.currentProvisions) : "",
      unit: isSet(object.unit) ? String(object.unit) : "",
      communityProvisions: isSet(object.communityProvisions) ? String(object.communityProvisions) : "",
    };
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.halfCount !== undefined && (obj.halfCount = Math.round(message.halfCount));
    message.firstProvisions !== undefined && (obj.firstProvisions = message.firstProvisions);
    message.currentProvisions !== undefined && (obj.currentProvisions = message.currentProvisions);
    message.unit !== undefined && (obj.unit = message.unit);
    message.communityProvisions !== undefined && (obj.communityProvisions = message.communityProvisions);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Params>, I>>(object: I): Params {
    const message = createBaseParams();
    message.halfCount = object.halfCount ?? 0;
    message.firstProvisions = object.firstProvisions ?? "";
    message.currentProvisions = object.currentProvisions ?? "";
    message.unit = object.unit ?? "";
    message.communityProvisions = object.communityProvisions ?? "";
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
