/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "joltify.vault";

export interface Params {
  blockChurnInterval: number;
  power: number;
  step: number;
  candidateRatio: string;
  targetQuota: Coin[];
  historyLength: number;
}

export interface Validator {
  pubkey: Uint8Array;
  power: number;
}

export interface StandbyPower {
  addr: string;
  power: number;
}

export interface Validators {
  allValidators: Validator[];
  height: number;
}

function createBaseParams(): Params {
  return { blockChurnInterval: 0, power: 0, step: 0, candidateRatio: "", targetQuota: [], historyLength: 0 };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.blockChurnInterval !== 0) {
      writer.uint32(8).int64(message.blockChurnInterval);
    }
    if (message.power !== 0) {
      writer.uint32(16).int64(message.power);
    }
    if (message.step !== 0) {
      writer.uint32(24).int64(message.step);
    }
    if (message.candidateRatio !== "") {
      writer.uint32(34).string(message.candidateRatio);
    }
    for (const v of message.targetQuota) {
      Coin.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    if (message.historyLength !== 0) {
      writer.uint32(48).int32(message.historyLength);
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
          message.blockChurnInterval = longToNumber(reader.int64() as Long);
          break;
        case 2:
          message.power = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.step = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.candidateRatio = reader.string();
          break;
        case 5:
          message.targetQuota.push(Coin.decode(reader, reader.uint32()));
          break;
        case 6:
          message.historyLength = reader.int32();
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
      blockChurnInterval: isSet(object.blockChurnInterval) ? Number(object.blockChurnInterval) : 0,
      power: isSet(object.power) ? Number(object.power) : 0,
      step: isSet(object.step) ? Number(object.step) : 0,
      candidateRatio: isSet(object.candidateRatio) ? String(object.candidateRatio) : "",
      targetQuota: Array.isArray(object?.targetQuota) ? object.targetQuota.map((e: any) => Coin.fromJSON(e)) : [],
      historyLength: isSet(object.historyLength) ? Number(object.historyLength) : 0,
    };
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.blockChurnInterval !== undefined && (obj.blockChurnInterval = Math.round(message.blockChurnInterval));
    message.power !== undefined && (obj.power = Math.round(message.power));
    message.step !== undefined && (obj.step = Math.round(message.step));
    message.candidateRatio !== undefined && (obj.candidateRatio = message.candidateRatio);
    if (message.targetQuota) {
      obj.targetQuota = message.targetQuota.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.targetQuota = [];
    }
    message.historyLength !== undefined && (obj.historyLength = Math.round(message.historyLength));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Params>, I>>(object: I): Params {
    const message = createBaseParams();
    message.blockChurnInterval = object.blockChurnInterval ?? 0;
    message.power = object.power ?? 0;
    message.step = object.step ?? 0;
    message.candidateRatio = object.candidateRatio ?? "";
    message.targetQuota = object.targetQuota?.map((e) => Coin.fromPartial(e)) || [];
    message.historyLength = object.historyLength ?? 0;
    return message;
  },
};

function createBaseValidator(): Validator {
  return { pubkey: new Uint8Array(), power: 0 };
}

export const Validator = {
  encode(message: Validator, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pubkey.length !== 0) {
      writer.uint32(10).bytes(message.pubkey);
    }
    if (message.power !== 0) {
      writer.uint32(16).int64(message.power);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Validator {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseValidator();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pubkey = reader.bytes();
          break;
        case 2:
          message.power = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Validator {
    return {
      pubkey: isSet(object.pubkey) ? bytesFromBase64(object.pubkey) : new Uint8Array(),
      power: isSet(object.power) ? Number(object.power) : 0,
    };
  },

  toJSON(message: Validator): unknown {
    const obj: any = {};
    message.pubkey !== undefined
      && (obj.pubkey = base64FromBytes(message.pubkey !== undefined ? message.pubkey : new Uint8Array()));
    message.power !== undefined && (obj.power = Math.round(message.power));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Validator>, I>>(object: I): Validator {
    const message = createBaseValidator();
    message.pubkey = object.pubkey ?? new Uint8Array();
    message.power = object.power ?? 0;
    return message;
  },
};

function createBaseStandbyPower(): StandbyPower {
  return { addr: "", power: 0 };
}

export const StandbyPower = {
  encode(message: StandbyPower, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.addr !== "") {
      writer.uint32(10).string(message.addr);
    }
    if (message.power !== 0) {
      writer.uint32(16).int64(message.power);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): StandbyPower {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStandbyPower();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.addr = reader.string();
          break;
        case 2:
          message.power = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StandbyPower {
    return {
      addr: isSet(object.addr) ? String(object.addr) : "",
      power: isSet(object.power) ? Number(object.power) : 0,
    };
  },

  toJSON(message: StandbyPower): unknown {
    const obj: any = {};
    message.addr !== undefined && (obj.addr = message.addr);
    message.power !== undefined && (obj.power = Math.round(message.power));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<StandbyPower>, I>>(object: I): StandbyPower {
    const message = createBaseStandbyPower();
    message.addr = object.addr ?? "";
    message.power = object.power ?? 0;
    return message;
  },
};

function createBaseValidators(): Validators {
  return { allValidators: [], height: 0 };
}

export const Validators = {
  encode(message: Validators, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.allValidators) {
      Validator.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.height !== 0) {
      writer.uint32(16).int64(message.height);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Validators {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseValidators();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.allValidators.push(Validator.decode(reader, reader.uint32()));
          break;
        case 2:
          message.height = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Validators {
    return {
      allValidators: Array.isArray(object?.allValidators)
        ? object.allValidators.map((e: any) => Validator.fromJSON(e))
        : [],
      height: isSet(object.height) ? Number(object.height) : 0,
    };
  },

  toJSON(message: Validators): unknown {
    const obj: any = {};
    if (message.allValidators) {
      obj.allValidators = message.allValidators.map((e) => e ? Validator.toJSON(e) : undefined);
    } else {
      obj.allValidators = [];
    }
    message.height !== undefined && (obj.height = Math.round(message.height));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Validators>, I>>(object: I): Validators {
    const message = createBaseValidators();
    message.allValidators = object.allValidators?.map((e) => Validator.fromPartial(e)) || [];
    message.height = object.height ?? 0;
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

function bytesFromBase64(b64: string): Uint8Array {
  if (globalThis.Buffer) {
    return Uint8Array.from(globalThis.Buffer.from(b64, "base64"));
  } else {
    const bin = globalThis.atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
      arr[i] = bin.charCodeAt(i);
    }
    return arr;
  }
}

function base64FromBytes(arr: Uint8Array): string {
  if (globalThis.Buffer) {
    return globalThis.Buffer.from(arr).toString("base64");
  } else {
    const bin: string[] = [];
    arr.forEach((byte) => {
      bin.push(String.fromCharCode(byte));
    });
    return globalThis.btoa(bin.join(""));
  }
}

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
