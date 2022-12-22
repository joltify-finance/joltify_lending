/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "joltify.vault";

export interface PoolProposal {
  poolPubKey: string;
  poolAddr: Uint8Array;
  nodes: Uint8Array[];
}

export interface CreatePool {
  blockHeight: string;
  validators: Uint8Array[];
  proposal: PoolProposal[];
}

function createBasePoolProposal(): PoolProposal {
  return { poolPubKey: "", poolAddr: new Uint8Array(), nodes: [] };
}

export const PoolProposal = {
  encode(message: PoolProposal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.poolPubKey !== "") {
      writer.uint32(10).string(message.poolPubKey);
    }
    if (message.poolAddr.length !== 0) {
      writer.uint32(18).bytes(message.poolAddr);
    }
    for (const v of message.nodes) {
      writer.uint32(26).bytes(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PoolProposal {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePoolProposal();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.poolPubKey = reader.string();
          break;
        case 2:
          message.poolAddr = reader.bytes();
          break;
        case 3:
          message.nodes.push(reader.bytes());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PoolProposal {
    return {
      poolPubKey: isSet(object.poolPubKey) ? String(object.poolPubKey) : "",
      poolAddr: isSet(object.poolAddr) ? bytesFromBase64(object.poolAddr) : new Uint8Array(),
      nodes: Array.isArray(object?.nodes) ? object.nodes.map((e: any) => bytesFromBase64(e)) : [],
    };
  },

  toJSON(message: PoolProposal): unknown {
    const obj: any = {};
    message.poolPubKey !== undefined && (obj.poolPubKey = message.poolPubKey);
    message.poolAddr !== undefined
      && (obj.poolAddr = base64FromBytes(message.poolAddr !== undefined ? message.poolAddr : new Uint8Array()));
    if (message.nodes) {
      obj.nodes = message.nodes.map((e) => base64FromBytes(e !== undefined ? e : new Uint8Array()));
    } else {
      obj.nodes = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PoolProposal>, I>>(object: I): PoolProposal {
    const message = createBasePoolProposal();
    message.poolPubKey = object.poolPubKey ?? "";
    message.poolAddr = object.poolAddr ?? new Uint8Array();
    message.nodes = object.nodes?.map((e) => e) || [];
    return message;
  },
};

function createBaseCreatePool(): CreatePool {
  return { blockHeight: "", validators: [], proposal: [] };
}

export const CreatePool = {
  encode(message: CreatePool, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.blockHeight !== "") {
      writer.uint32(10).string(message.blockHeight);
    }
    for (const v of message.validators) {
      writer.uint32(18).bytes(v!);
    }
    for (const v of message.proposal) {
      PoolProposal.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreatePool {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreatePool();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.blockHeight = reader.string();
          break;
        case 2:
          message.validators.push(reader.bytes());
          break;
        case 3:
          message.proposal.push(PoolProposal.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreatePool {
    return {
      blockHeight: isSet(object.blockHeight) ? String(object.blockHeight) : "",
      validators: Array.isArray(object?.validators) ? object.validators.map((e: any) => bytesFromBase64(e)) : [],
      proposal: Array.isArray(object?.proposal) ? object.proposal.map((e: any) => PoolProposal.fromJSON(e)) : [],
    };
  },

  toJSON(message: CreatePool): unknown {
    const obj: any = {};
    message.blockHeight !== undefined && (obj.blockHeight = message.blockHeight);
    if (message.validators) {
      obj.validators = message.validators.map((e) => base64FromBytes(e !== undefined ? e : new Uint8Array()));
    } else {
      obj.validators = [];
    }
    if (message.proposal) {
      obj.proposal = message.proposal.map((e) => e ? PoolProposal.toJSON(e) : undefined);
    } else {
      obj.proposal = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreatePool>, I>>(object: I): CreatePool {
    const message = createBaseCreatePool();
    message.blockHeight = object.blockHeight ?? "";
    message.validators = object.validators?.map((e) => e) || [];
    message.proposal = object.proposal?.map((e) => PoolProposal.fromPartial(e)) || [];
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
