/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "joltify.spv";

export interface DepositorInfo {
  investorId: string;
  depositorAddress: Uint8Array;
  poolIndex: string;
  lockedAmount: Coin | undefined;
  withdrawalAmount: Coin | undefined;
  incentiveAmount: Coin | undefined;
  linkedNFT: string[];
}

function createBaseDepositorInfo(): DepositorInfo {
  return {
    investorId: "",
    depositorAddress: new Uint8Array(),
    poolIndex: "",
    lockedAmount: undefined,
    withdrawalAmount: undefined,
    incentiveAmount: undefined,
    linkedNFT: [],
  };
}

export const DepositorInfo = {
  encode(message: DepositorInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.investorId !== "") {
      writer.uint32(10).string(message.investorId);
    }
    if (message.depositorAddress.length !== 0) {
      writer.uint32(18).bytes(message.depositorAddress);
    }
    if (message.poolIndex !== "") {
      writer.uint32(26).string(message.poolIndex);
    }
    if (message.lockedAmount !== undefined) {
      Coin.encode(message.lockedAmount, writer.uint32(34).fork()).ldelim();
    }
    if (message.withdrawalAmount !== undefined) {
      Coin.encode(message.withdrawalAmount, writer.uint32(42).fork()).ldelim();
    }
    if (message.incentiveAmount !== undefined) {
      Coin.encode(message.incentiveAmount, writer.uint32(50).fork()).ldelim();
    }
    for (const v of message.linkedNFT) {
      writer.uint32(58).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DepositorInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDepositorInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.investorId = reader.string();
          break;
        case 2:
          message.depositorAddress = reader.bytes();
          break;
        case 3:
          message.poolIndex = reader.string();
          break;
        case 4:
          message.lockedAmount = Coin.decode(reader, reader.uint32());
          break;
        case 5:
          message.withdrawalAmount = Coin.decode(reader, reader.uint32());
          break;
        case 6:
          message.incentiveAmount = Coin.decode(reader, reader.uint32());
          break;
        case 7:
          message.linkedNFT.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DepositorInfo {
    return {
      investorId: isSet(object.investorId) ? String(object.investorId) : "",
      depositorAddress: isSet(object.depositorAddress) ? bytesFromBase64(object.depositorAddress) : new Uint8Array(),
      poolIndex: isSet(object.poolIndex) ? String(object.poolIndex) : "",
      lockedAmount: isSet(object.lockedAmount) ? Coin.fromJSON(object.lockedAmount) : undefined,
      withdrawalAmount: isSet(object.withdrawalAmount) ? Coin.fromJSON(object.withdrawalAmount) : undefined,
      incentiveAmount: isSet(object.incentiveAmount) ? Coin.fromJSON(object.incentiveAmount) : undefined,
      linkedNFT: Array.isArray(object?.linkedNFT) ? object.linkedNFT.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: DepositorInfo): unknown {
    const obj: any = {};
    message.investorId !== undefined && (obj.investorId = message.investorId);
    message.depositorAddress !== undefined
      && (obj.depositorAddress = base64FromBytes(
        message.depositorAddress !== undefined ? message.depositorAddress : new Uint8Array(),
      ));
    message.poolIndex !== undefined && (obj.poolIndex = message.poolIndex);
    message.lockedAmount !== undefined
      && (obj.lockedAmount = message.lockedAmount ? Coin.toJSON(message.lockedAmount) : undefined);
    message.withdrawalAmount !== undefined
      && (obj.withdrawalAmount = message.withdrawalAmount ? Coin.toJSON(message.withdrawalAmount) : undefined);
    message.incentiveAmount !== undefined
      && (obj.incentiveAmount = message.incentiveAmount ? Coin.toJSON(message.incentiveAmount) : undefined);
    if (message.linkedNFT) {
      obj.linkedNFT = message.linkedNFT.map((e) => e);
    } else {
      obj.linkedNFT = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DepositorInfo>, I>>(object: I): DepositorInfo {
    const message = createBaseDepositorInfo();
    message.investorId = object.investorId ?? "";
    message.depositorAddress = object.depositorAddress ?? new Uint8Array();
    message.poolIndex = object.poolIndex ?? "";
    message.lockedAmount = (object.lockedAmount !== undefined && object.lockedAmount !== null)
      ? Coin.fromPartial(object.lockedAmount)
      : undefined;
    message.withdrawalAmount = (object.withdrawalAmount !== undefined && object.withdrawalAmount !== null)
      ? Coin.fromPartial(object.withdrawalAmount)
      : undefined;
    message.incentiveAmount = (object.incentiveAmount !== undefined && object.incentiveAmount !== null)
      ? Coin.fromPartial(object.incentiveAmount)
      : undefined;
    message.linkedNFT = object.linkedNFT?.map((e) => e) || [];
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
