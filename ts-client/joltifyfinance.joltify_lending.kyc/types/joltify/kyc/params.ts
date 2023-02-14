/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "joltifyfinance.joltify_lending.kyc";

export interface BasicInfo {
  description: string;
  projectsUrl: string;
  projectCountry: string;
  businessNumber: string;
  reserved: Uint8Array;
  projectName: string;
}

/** Market defines an asset in the pricefeed. */
export interface ProjectInfo {
  index: number;
  SPVName: string;
  basicInfo: BasicInfo | undefined;
  projectOwner: Uint8Array;
  projectLength: number;
  projectTargetAmount: Coin | undefined;
  baseApy: string;
  payFreq: string;
}

/** Params defines the parameters for the module. */
export interface Params {
  projectsInfo: ProjectInfo[];
  submitter: Uint8Array[];
}

function createBaseBasicInfo(): BasicInfo {
  return {
    description: "",
    projectsUrl: "",
    projectCountry: "",
    businessNumber: "",
    reserved: new Uint8Array(),
    projectName: "",
  };
}

export const BasicInfo = {
  encode(message: BasicInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.description !== "") {
      writer.uint32(10).string(message.description);
    }
    if (message.projectsUrl !== "") {
      writer.uint32(18).string(message.projectsUrl);
    }
    if (message.projectCountry !== "") {
      writer.uint32(26).string(message.projectCountry);
    }
    if (message.businessNumber !== "") {
      writer.uint32(34).string(message.businessNumber);
    }
    if (message.reserved.length !== 0) {
      writer.uint32(42).bytes(message.reserved);
    }
    if (message.projectName !== "") {
      writer.uint32(50).string(message.projectName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BasicInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBasicInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.description = reader.string();
          break;
        case 2:
          message.projectsUrl = reader.string();
          break;
        case 3:
          message.projectCountry = reader.string();
          break;
        case 4:
          message.businessNumber = reader.string();
          break;
        case 5:
          message.reserved = reader.bytes();
          break;
        case 6:
          message.projectName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): BasicInfo {
    return {
      description: isSet(object.description) ? String(object.description) : "",
      projectsUrl: isSet(object.projectsUrl) ? String(object.projectsUrl) : "",
      projectCountry: isSet(object.projectCountry) ? String(object.projectCountry) : "",
      businessNumber: isSet(object.businessNumber) ? String(object.businessNumber) : "",
      reserved: isSet(object.reserved) ? bytesFromBase64(object.reserved) : new Uint8Array(),
      projectName: isSet(object.projectName) ? String(object.projectName) : "",
    };
  },

  toJSON(message: BasicInfo): unknown {
    const obj: any = {};
    message.description !== undefined && (obj.description = message.description);
    message.projectsUrl !== undefined && (obj.projectsUrl = message.projectsUrl);
    message.projectCountry !== undefined && (obj.projectCountry = message.projectCountry);
    message.businessNumber !== undefined && (obj.businessNumber = message.businessNumber);
    message.reserved !== undefined
      && (obj.reserved = base64FromBytes(message.reserved !== undefined ? message.reserved : new Uint8Array()));
    message.projectName !== undefined && (obj.projectName = message.projectName);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<BasicInfo>, I>>(object: I): BasicInfo {
    const message = createBaseBasicInfo();
    message.description = object.description ?? "";
    message.projectsUrl = object.projectsUrl ?? "";
    message.projectCountry = object.projectCountry ?? "";
    message.businessNumber = object.businessNumber ?? "";
    message.reserved = object.reserved ?? new Uint8Array();
    message.projectName = object.projectName ?? "";
    return message;
  },
};

function createBaseProjectInfo(): ProjectInfo {
  return {
    index: 0,
    SPVName: "",
    basicInfo: undefined,
    projectOwner: new Uint8Array(),
    projectLength: 0,
    projectTargetAmount: undefined,
    baseApy: "",
    payFreq: "",
  };
}

export const ProjectInfo = {
  encode(message: ProjectInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== 0) {
      writer.uint32(8).int32(message.index);
    }
    if (message.SPVName !== "") {
      writer.uint32(18).string(message.SPVName);
    }
    if (message.basicInfo !== undefined) {
      BasicInfo.encode(message.basicInfo, writer.uint32(26).fork()).ldelim();
    }
    if (message.projectOwner.length !== 0) {
      writer.uint32(34).bytes(message.projectOwner);
    }
    if (message.projectLength !== 0) {
      writer.uint32(40).uint64(message.projectLength);
    }
    if (message.projectTargetAmount !== undefined) {
      Coin.encode(message.projectTargetAmount, writer.uint32(50).fork()).ldelim();
    }
    if (message.baseApy !== "") {
      writer.uint32(58).string(message.baseApy);
    }
    if (message.payFreq !== "") {
      writer.uint32(66).string(message.payFreq);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ProjectInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProjectInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.int32();
          break;
        case 2:
          message.SPVName = reader.string();
          break;
        case 3:
          message.basicInfo = BasicInfo.decode(reader, reader.uint32());
          break;
        case 4:
          message.projectOwner = reader.bytes();
          break;
        case 5:
          message.projectLength = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.projectTargetAmount = Coin.decode(reader, reader.uint32());
          break;
        case 7:
          message.baseApy = reader.string();
          break;
        case 8:
          message.payFreq = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ProjectInfo {
    return {
      index: isSet(object.index) ? Number(object.index) : 0,
      SPVName: isSet(object.SPVName) ? String(object.SPVName) : "",
      basicInfo: isSet(object.basicInfo) ? BasicInfo.fromJSON(object.basicInfo) : undefined,
      projectOwner: isSet(object.projectOwner) ? bytesFromBase64(object.projectOwner) : new Uint8Array(),
      projectLength: isSet(object.projectLength) ? Number(object.projectLength) : 0,
      projectTargetAmount: isSet(object.projectTargetAmount) ? Coin.fromJSON(object.projectTargetAmount) : undefined,
      baseApy: isSet(object.baseApy) ? String(object.baseApy) : "",
      payFreq: isSet(object.payFreq) ? String(object.payFreq) : "",
    };
  },

  toJSON(message: ProjectInfo): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = Math.round(message.index));
    message.SPVName !== undefined && (obj.SPVName = message.SPVName);
    message.basicInfo !== undefined
      && (obj.basicInfo = message.basicInfo ? BasicInfo.toJSON(message.basicInfo) : undefined);
    message.projectOwner !== undefined
      && (obj.projectOwner = base64FromBytes(
        message.projectOwner !== undefined ? message.projectOwner : new Uint8Array(),
      ));
    message.projectLength !== undefined && (obj.projectLength = Math.round(message.projectLength));
    message.projectTargetAmount !== undefined
      && (obj.projectTargetAmount = message.projectTargetAmount ? Coin.toJSON(message.projectTargetAmount) : undefined);
    message.baseApy !== undefined && (obj.baseApy = message.baseApy);
    message.payFreq !== undefined && (obj.payFreq = message.payFreq);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ProjectInfo>, I>>(object: I): ProjectInfo {
    const message = createBaseProjectInfo();
    message.index = object.index ?? 0;
    message.SPVName = object.SPVName ?? "";
    message.basicInfo = (object.basicInfo !== undefined && object.basicInfo !== null)
      ? BasicInfo.fromPartial(object.basicInfo)
      : undefined;
    message.projectOwner = object.projectOwner ?? new Uint8Array();
    message.projectLength = object.projectLength ?? 0;
    message.projectTargetAmount = (object.projectTargetAmount !== undefined && object.projectTargetAmount !== null)
      ? Coin.fromPartial(object.projectTargetAmount)
      : undefined;
    message.baseApy = object.baseApy ?? "";
    message.payFreq = object.payFreq ?? "";
    return message;
  },
};

function createBaseParams(): Params {
  return { projectsInfo: [], submitter: [] };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.projectsInfo) {
      ProjectInfo.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.submitter) {
      writer.uint32(34).bytes(v!);
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
          message.projectsInfo.push(ProjectInfo.decode(reader, reader.uint32()));
          break;
        case 4:
          message.submitter.push(reader.bytes());
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
      projectsInfo: Array.isArray(object?.projectsInfo)
        ? object.projectsInfo.map((e: any) => ProjectInfo.fromJSON(e))
        : [],
      submitter: Array.isArray(object?.submitter) ? object.submitter.map((e: any) => bytesFromBase64(e)) : [],
    };
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    if (message.projectsInfo) {
      obj.projectsInfo = message.projectsInfo.map((e) => e ? ProjectInfo.toJSON(e) : undefined);
    } else {
      obj.projectsInfo = [];
    }
    if (message.submitter) {
      obj.submitter = message.submitter.map((e) => base64FromBytes(e !== undefined ? e : new Uint8Array()));
    } else {
      obj.submitter = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Params>, I>>(object: I): Params {
    const message = createBaseParams();
    message.projectsInfo = object.projectsInfo?.map((e) => ProjectInfo.fromPartial(e)) || [];
    message.submitter = object.submitter?.map((e) => e) || [];
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
