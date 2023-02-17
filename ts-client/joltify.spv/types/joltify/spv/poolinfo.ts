/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";
import { Timestamp } from "../../google/protobuf/timestamp";

export const protobufPackage = "joltify.spv";

export interface PoolInfo {
  index: string;
  poolName: string;
  linkedProject: number;
  ownerAddress: Uint8Array;
  apy: string;
  totalAmount: Coin | undefined;
  payFreq: number;
  reserveFactor: string;
  /**
   * string            pool_nFT_class      = 9 [
   *    (cosmos_proto.scalar)  = "cosmos.Class",
   *    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/x/nft.Class",
   *    (gogoproto.nullable)   = false
   *  ];
   */
  poolNFTIds: string[];
  poolStartTime: Date | undefined;
  poolStatus: PoolInfo_POOLSTATUS;
  borrowedAmount: Coin | undefined;
  poolInterest: string;
  projectLength: number;
  borrowableAmount: Coin | undefined;
}

export enum PoolInfo_POOLSTATUS {
  ACTIVE = 0,
  INACTIVE = 1,
  CLOSED = 2,
  OPEN = 3,
  UNRECOGNIZED = -1,
}

export function poolInfo_POOLSTATUSFromJSON(object: any): PoolInfo_POOLSTATUS {
  switch (object) {
    case 0:
    case "ACTIVE":
      return PoolInfo_POOLSTATUS.ACTIVE;
    case 1:
    case "INACTIVE":
      return PoolInfo_POOLSTATUS.INACTIVE;
    case 2:
    case "CLOSED":
      return PoolInfo_POOLSTATUS.CLOSED;
    case 3:
    case "OPEN":
      return PoolInfo_POOLSTATUS.OPEN;
    case -1:
    case "UNRECOGNIZED":
    default:
      return PoolInfo_POOLSTATUS.UNRECOGNIZED;
  }
}

export function poolInfo_POOLSTATUSToJSON(object: PoolInfo_POOLSTATUS): string {
  switch (object) {
    case PoolInfo_POOLSTATUS.ACTIVE:
      return "ACTIVE";
    case PoolInfo_POOLSTATUS.INACTIVE:
      return "INACTIVE";
    case PoolInfo_POOLSTATUS.CLOSED:
      return "CLOSED";
    case PoolInfo_POOLSTATUS.OPEN:
      return "OPEN";
    case PoolInfo_POOLSTATUS.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface PoolWithInvestors {
  poolIndex: string;
  investors: string[];
}

export interface PoolDepositedInvestors {
  poolIndex: string;
  walletAddress: Uint8Array[];
}

function createBasePoolInfo(): PoolInfo {
  return {
    index: "",
    poolName: "",
    linkedProject: 0,
    ownerAddress: new Uint8Array(),
    apy: "",
    totalAmount: undefined,
    payFreq: 0,
    reserveFactor: "",
    poolNFTIds: [],
    poolStartTime: undefined,
    poolStatus: 0,
    borrowedAmount: undefined,
    poolInterest: "",
    projectLength: 0,
    borrowableAmount: undefined,
  };
}

export const PoolInfo = {
  encode(message: PoolInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.poolName !== "") {
      writer.uint32(18).string(message.poolName);
    }
    if (message.linkedProject !== 0) {
      writer.uint32(24).int32(message.linkedProject);
    }
    if (message.ownerAddress.length !== 0) {
      writer.uint32(34).bytes(message.ownerAddress);
    }
    if (message.apy !== "") {
      writer.uint32(42).string(message.apy);
    }
    if (message.totalAmount !== undefined) {
      Coin.encode(message.totalAmount, writer.uint32(50).fork()).ldelim();
    }
    if (message.payFreq !== 0) {
      writer.uint32(56).int32(message.payFreq);
    }
    if (message.reserveFactor !== "") {
      writer.uint32(66).string(message.reserveFactor);
    }
    for (const v of message.poolNFTIds) {
      writer.uint32(74).string(v!);
    }
    if (message.poolStartTime !== undefined) {
      Timestamp.encode(toTimestamp(message.poolStartTime), writer.uint32(82).fork()).ldelim();
    }
    if (message.poolStatus !== 0) {
      writer.uint32(88).int32(message.poolStatus);
    }
    if (message.borrowedAmount !== undefined) {
      Coin.encode(message.borrowedAmount, writer.uint32(98).fork()).ldelim();
    }
    if (message.poolInterest !== "") {
      writer.uint32(106).string(message.poolInterest);
    }
    if (message.projectLength !== 0) {
      writer.uint32(112).uint64(message.projectLength);
    }
    if (message.borrowableAmount !== undefined) {
      Coin.encode(message.borrowableAmount, writer.uint32(122).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PoolInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePoolInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.poolName = reader.string();
          break;
        case 3:
          message.linkedProject = reader.int32();
          break;
        case 4:
          message.ownerAddress = reader.bytes();
          break;
        case 5:
          message.apy = reader.string();
          break;
        case 6:
          message.totalAmount = Coin.decode(reader, reader.uint32());
          break;
        case 7:
          message.payFreq = reader.int32();
          break;
        case 8:
          message.reserveFactor = reader.string();
          break;
        case 9:
          message.poolNFTIds.push(reader.string());
          break;
        case 10:
          message.poolStartTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 11:
          message.poolStatus = reader.int32() as any;
          break;
        case 12:
          message.borrowedAmount = Coin.decode(reader, reader.uint32());
          break;
        case 13:
          message.poolInterest = reader.string();
          break;
        case 14:
          message.projectLength = longToNumber(reader.uint64() as Long);
          break;
        case 15:
          message.borrowableAmount = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PoolInfo {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      poolName: isSet(object.poolName) ? String(object.poolName) : "",
      linkedProject: isSet(object.linkedProject) ? Number(object.linkedProject) : 0,
      ownerAddress: isSet(object.ownerAddress) ? bytesFromBase64(object.ownerAddress) : new Uint8Array(),
      apy: isSet(object.apy) ? String(object.apy) : "",
      totalAmount: isSet(object.totalAmount) ? Coin.fromJSON(object.totalAmount) : undefined,
      payFreq: isSet(object.payFreq) ? Number(object.payFreq) : 0,
      reserveFactor: isSet(object.reserveFactor) ? String(object.reserveFactor) : "",
      poolNFTIds: Array.isArray(object?.poolNFTIds) ? object.poolNFTIds.map((e: any) => String(e)) : [],
      poolStartTime: isSet(object.poolStartTime) ? fromJsonTimestamp(object.poolStartTime) : undefined,
      poolStatus: isSet(object.poolStatus) ? poolInfo_POOLSTATUSFromJSON(object.poolStatus) : 0,
      borrowedAmount: isSet(object.borrowedAmount) ? Coin.fromJSON(object.borrowedAmount) : undefined,
      poolInterest: isSet(object.poolInterest) ? String(object.poolInterest) : "",
      projectLength: isSet(object.projectLength) ? Number(object.projectLength) : 0,
      borrowableAmount: isSet(object.borrowableAmount) ? Coin.fromJSON(object.borrowableAmount) : undefined,
    };
  },

  toJSON(message: PoolInfo): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.poolName !== undefined && (obj.poolName = message.poolName);
    message.linkedProject !== undefined && (obj.linkedProject = Math.round(message.linkedProject));
    message.ownerAddress !== undefined
      && (obj.ownerAddress = base64FromBytes(
        message.ownerAddress !== undefined ? message.ownerAddress : new Uint8Array(),
      ));
    message.apy !== undefined && (obj.apy = message.apy);
    message.totalAmount !== undefined
      && (obj.totalAmount = message.totalAmount ? Coin.toJSON(message.totalAmount) : undefined);
    message.payFreq !== undefined && (obj.payFreq = Math.round(message.payFreq));
    message.reserveFactor !== undefined && (obj.reserveFactor = message.reserveFactor);
    if (message.poolNFTIds) {
      obj.poolNFTIds = message.poolNFTIds.map((e) => e);
    } else {
      obj.poolNFTIds = [];
    }
    message.poolStartTime !== undefined && (obj.poolStartTime = message.poolStartTime.toISOString());
    message.poolStatus !== undefined && (obj.poolStatus = poolInfo_POOLSTATUSToJSON(message.poolStatus));
    message.borrowedAmount !== undefined
      && (obj.borrowedAmount = message.borrowedAmount ? Coin.toJSON(message.borrowedAmount) : undefined);
    message.poolInterest !== undefined && (obj.poolInterest = message.poolInterest);
    message.projectLength !== undefined && (obj.projectLength = Math.round(message.projectLength));
    message.borrowableAmount !== undefined
      && (obj.borrowableAmount = message.borrowableAmount ? Coin.toJSON(message.borrowableAmount) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PoolInfo>, I>>(object: I): PoolInfo {
    const message = createBasePoolInfo();
    message.index = object.index ?? "";
    message.poolName = object.poolName ?? "";
    message.linkedProject = object.linkedProject ?? 0;
    message.ownerAddress = object.ownerAddress ?? new Uint8Array();
    message.apy = object.apy ?? "";
    message.totalAmount = (object.totalAmount !== undefined && object.totalAmount !== null)
      ? Coin.fromPartial(object.totalAmount)
      : undefined;
    message.payFreq = object.payFreq ?? 0;
    message.reserveFactor = object.reserveFactor ?? "";
    message.poolNFTIds = object.poolNFTIds?.map((e) => e) || [];
    message.poolStartTime = object.poolStartTime ?? undefined;
    message.poolStatus = object.poolStatus ?? 0;
    message.borrowedAmount = (object.borrowedAmount !== undefined && object.borrowedAmount !== null)
      ? Coin.fromPartial(object.borrowedAmount)
      : undefined;
    message.poolInterest = object.poolInterest ?? "";
    message.projectLength = object.projectLength ?? 0;
    message.borrowableAmount = (object.borrowableAmount !== undefined && object.borrowableAmount !== null)
      ? Coin.fromPartial(object.borrowableAmount)
      : undefined;
    return message;
  },
};

function createBasePoolWithInvestors(): PoolWithInvestors {
  return { poolIndex: "", investors: [] };
}

export const PoolWithInvestors = {
  encode(message: PoolWithInvestors, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.poolIndex !== "") {
      writer.uint32(10).string(message.poolIndex);
    }
    for (const v of message.investors) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PoolWithInvestors {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePoolWithInvestors();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.poolIndex = reader.string();
          break;
        case 2:
          message.investors.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PoolWithInvestors {
    return {
      poolIndex: isSet(object.poolIndex) ? String(object.poolIndex) : "",
      investors: Array.isArray(object?.investors) ? object.investors.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: PoolWithInvestors): unknown {
    const obj: any = {};
    message.poolIndex !== undefined && (obj.poolIndex = message.poolIndex);
    if (message.investors) {
      obj.investors = message.investors.map((e) => e);
    } else {
      obj.investors = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PoolWithInvestors>, I>>(object: I): PoolWithInvestors {
    const message = createBasePoolWithInvestors();
    message.poolIndex = object.poolIndex ?? "";
    message.investors = object.investors?.map((e) => e) || [];
    return message;
  },
};

function createBasePoolDepositedInvestors(): PoolDepositedInvestors {
  return { poolIndex: "", walletAddress: [] };
}

export const PoolDepositedInvestors = {
  encode(message: PoolDepositedInvestors, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.poolIndex !== "") {
      writer.uint32(10).string(message.poolIndex);
    }
    for (const v of message.walletAddress) {
      writer.uint32(34).bytes(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PoolDepositedInvestors {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePoolDepositedInvestors();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.poolIndex = reader.string();
          break;
        case 4:
          message.walletAddress.push(reader.bytes());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PoolDepositedInvestors {
    return {
      poolIndex: isSet(object.poolIndex) ? String(object.poolIndex) : "",
      walletAddress: Array.isArray(object?.walletAddress)
        ? object.walletAddress.map((e: any) => bytesFromBase64(e))
        : [],
    };
  },

  toJSON(message: PoolDepositedInvestors): unknown {
    const obj: any = {};
    message.poolIndex !== undefined && (obj.poolIndex = message.poolIndex);
    if (message.walletAddress) {
      obj.walletAddress = message.walletAddress.map((e) => base64FromBytes(e !== undefined ? e : new Uint8Array()));
    } else {
      obj.walletAddress = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PoolDepositedInvestors>, I>>(object: I): PoolDepositedInvestors {
    const message = createBasePoolDepositedInvestors();
    message.poolIndex = object.poolIndex ?? "";
    message.walletAddress = object.walletAddress?.map((e) => e) || [];
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
