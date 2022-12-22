/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Any } from "../../../google/protobuf/any";
import { Duration } from "../../../google/protobuf/duration";

export const protobufPackage = "joltify.auction.v1beta1";

/** GenesisState defines the auction module's genesis state. */
export interface GenesisState {
  nextAuctionId: number;
  params:
    | Params
    | undefined;
  /** Genesis auctions */
  auctions: Any[];
}

/** Params defines the parameters for the issuance module. */
export interface Params {
  maxAuctionDuration: Duration | undefined;
  forwardBidDuration: Duration | undefined;
  reverseBidDuration: Duration | undefined;
  incrementSurplus: Uint8Array;
  incrementDebt: Uint8Array;
  incrementCollateral: Uint8Array;
}

function createBaseGenesisState(): GenesisState {
  return { nextAuctionId: 0, params: undefined, auctions: [] };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.nextAuctionId !== 0) {
      writer.uint32(8).uint64(message.nextAuctionId);
    }
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.auctions) {
      Any.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nextAuctionId = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 3:
          message.auctions.push(Any.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      nextAuctionId: isSet(object.nextAuctionId) ? Number(object.nextAuctionId) : 0,
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      auctions: Array.isArray(object?.auctions) ? object.auctions.map((e: any) => Any.fromJSON(e)) : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.nextAuctionId !== undefined && (obj.nextAuctionId = Math.round(message.nextAuctionId));
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.auctions) {
      obj.auctions = message.auctions.map((e) => e ? Any.toJSON(e) : undefined);
    } else {
      obj.auctions = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.nextAuctionId = object.nextAuctionId ?? 0;
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.auctions = object.auctions?.map((e) => Any.fromPartial(e)) || [];
    return message;
  },
};

function createBaseParams(): Params {
  return {
    maxAuctionDuration: undefined,
    forwardBidDuration: undefined,
    reverseBidDuration: undefined,
    incrementSurplus: new Uint8Array(),
    incrementDebt: new Uint8Array(),
    incrementCollateral: new Uint8Array(),
  };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.maxAuctionDuration !== undefined) {
      Duration.encode(message.maxAuctionDuration, writer.uint32(10).fork()).ldelim();
    }
    if (message.forwardBidDuration !== undefined) {
      Duration.encode(message.forwardBidDuration, writer.uint32(50).fork()).ldelim();
    }
    if (message.reverseBidDuration !== undefined) {
      Duration.encode(message.reverseBidDuration, writer.uint32(58).fork()).ldelim();
    }
    if (message.incrementSurplus.length !== 0) {
      writer.uint32(26).bytes(message.incrementSurplus);
    }
    if (message.incrementDebt.length !== 0) {
      writer.uint32(34).bytes(message.incrementDebt);
    }
    if (message.incrementCollateral.length !== 0) {
      writer.uint32(42).bytes(message.incrementCollateral);
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
          message.maxAuctionDuration = Duration.decode(reader, reader.uint32());
          break;
        case 6:
          message.forwardBidDuration = Duration.decode(reader, reader.uint32());
          break;
        case 7:
          message.reverseBidDuration = Duration.decode(reader, reader.uint32());
          break;
        case 3:
          message.incrementSurplus = reader.bytes();
          break;
        case 4:
          message.incrementDebt = reader.bytes();
          break;
        case 5:
          message.incrementCollateral = reader.bytes();
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
      maxAuctionDuration: isSet(object.maxAuctionDuration) ? Duration.fromJSON(object.maxAuctionDuration) : undefined,
      forwardBidDuration: isSet(object.forwardBidDuration) ? Duration.fromJSON(object.forwardBidDuration) : undefined,
      reverseBidDuration: isSet(object.reverseBidDuration) ? Duration.fromJSON(object.reverseBidDuration) : undefined,
      incrementSurplus: isSet(object.incrementSurplus) ? bytesFromBase64(object.incrementSurplus) : new Uint8Array(),
      incrementDebt: isSet(object.incrementDebt) ? bytesFromBase64(object.incrementDebt) : new Uint8Array(),
      incrementCollateral: isSet(object.incrementCollateral)
        ? bytesFromBase64(object.incrementCollateral)
        : new Uint8Array(),
    };
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.maxAuctionDuration !== undefined
      && (obj.maxAuctionDuration = message.maxAuctionDuration
        ? Duration.toJSON(message.maxAuctionDuration)
        : undefined);
    message.forwardBidDuration !== undefined
      && (obj.forwardBidDuration = message.forwardBidDuration
        ? Duration.toJSON(message.forwardBidDuration)
        : undefined);
    message.reverseBidDuration !== undefined
      && (obj.reverseBidDuration = message.reverseBidDuration
        ? Duration.toJSON(message.reverseBidDuration)
        : undefined);
    message.incrementSurplus !== undefined
      && (obj.incrementSurplus = base64FromBytes(
        message.incrementSurplus !== undefined ? message.incrementSurplus : new Uint8Array(),
      ));
    message.incrementDebt !== undefined
      && (obj.incrementDebt = base64FromBytes(
        message.incrementDebt !== undefined ? message.incrementDebt : new Uint8Array(),
      ));
    message.incrementCollateral !== undefined
      && (obj.incrementCollateral = base64FromBytes(
        message.incrementCollateral !== undefined ? message.incrementCollateral : new Uint8Array(),
      ));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Params>, I>>(object: I): Params {
    const message = createBaseParams();
    message.maxAuctionDuration = (object.maxAuctionDuration !== undefined && object.maxAuctionDuration !== null)
      ? Duration.fromPartial(object.maxAuctionDuration)
      : undefined;
    message.forwardBidDuration = (object.forwardBidDuration !== undefined && object.forwardBidDuration !== null)
      ? Duration.fromPartial(object.forwardBidDuration)
      : undefined;
    message.reverseBidDuration = (object.reverseBidDuration !== undefined && object.reverseBidDuration !== null)
      ? Duration.fromPartial(object.reverseBidDuration)
      : undefined;
    message.incrementSurplus = object.incrementSurplus ?? new Uint8Array();
    message.incrementDebt = object.incrementDebt ?? new Uint8Array();
    message.incrementCollateral = object.incrementCollateral ?? new Uint8Array();
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
