/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "joltify.vault";

export interface Entity {
  address: Uint8Array;
  feecoin: Coin[];
}

export interface Proposals {
  entry: Entity[];
}

export interface OutboundTx {
  index: string;
  processed: boolean;
  items: { [key: string]: Proposals };
  chainType: string;
  inTxHash: string;
  receiverAddress: Uint8Array;
  needMint: boolean;
  feecoin: Coin[];
}

export interface OutboundTx_ItemsEntry {
  key: string;
  value: Proposals | undefined;
}

function createBaseEntity(): Entity {
  return { address: new Uint8Array(), feecoin: [] };
}

export const Entity = {
  encode(message: Entity, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address.length !== 0) {
      writer.uint32(10).bytes(message.address);
    }
    for (const v of message.feecoin) {
      Coin.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Entity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.bytes();
          break;
        case 2:
          message.feecoin.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Entity {
    return {
      address: isSet(object.address) ? bytesFromBase64(object.address) : new Uint8Array(),
      feecoin: Array.isArray(object?.feecoin) ? object.feecoin.map((e: any) => Coin.fromJSON(e)) : [],
    };
  },

  toJSON(message: Entity): unknown {
    const obj: any = {};
    message.address !== undefined
      && (obj.address = base64FromBytes(message.address !== undefined ? message.address : new Uint8Array()));
    if (message.feecoin) {
      obj.feecoin = message.feecoin.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.feecoin = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Entity>, I>>(object: I): Entity {
    const message = createBaseEntity();
    message.address = object.address ?? new Uint8Array();
    message.feecoin = object.feecoin?.map((e) => Coin.fromPartial(e)) || [];
    return message;
  },
};

function createBaseProposals(): Proposals {
  return { entry: [] };
}

export const Proposals = {
  encode(message: Proposals, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.entry) {
      Entity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Proposals {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProposals();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.entry.push(Entity.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Proposals {
    return { entry: Array.isArray(object?.entry) ? object.entry.map((e: any) => Entity.fromJSON(e)) : [] };
  },

  toJSON(message: Proposals): unknown {
    const obj: any = {};
    if (message.entry) {
      obj.entry = message.entry.map((e) => e ? Entity.toJSON(e) : undefined);
    } else {
      obj.entry = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Proposals>, I>>(object: I): Proposals {
    const message = createBaseProposals();
    message.entry = object.entry?.map((e) => Entity.fromPartial(e)) || [];
    return message;
  },
};

function createBaseOutboundTx(): OutboundTx {
  return {
    index: "",
    processed: false,
    items: {},
    chainType: "",
    inTxHash: "",
    receiverAddress: new Uint8Array(),
    needMint: false,
    feecoin: [],
  };
}

export const OutboundTx = {
  encode(message: OutboundTx, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.processed === true) {
      writer.uint32(16).bool(message.processed);
    }
    Object.entries(message.items).forEach(([key, value]) => {
      OutboundTx_ItemsEntry.encode({ key: key as any, value }, writer.uint32(26).fork()).ldelim();
    });
    if (message.chainType !== "") {
      writer.uint32(34).string(message.chainType);
    }
    if (message.inTxHash !== "") {
      writer.uint32(42).string(message.inTxHash);
    }
    if (message.receiverAddress.length !== 0) {
      writer.uint32(50).bytes(message.receiverAddress);
    }
    if (message.needMint === true) {
      writer.uint32(56).bool(message.needMint);
    }
    for (const v of message.feecoin) {
      Coin.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OutboundTx {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOutboundTx();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.processed = reader.bool();
          break;
        case 3:
          const entry3 = OutboundTx_ItemsEntry.decode(reader, reader.uint32());
          if (entry3.value !== undefined) {
            message.items[entry3.key] = entry3.value;
          }
          break;
        case 4:
          message.chainType = reader.string();
          break;
        case 5:
          message.inTxHash = reader.string();
          break;
        case 6:
          message.receiverAddress = reader.bytes();
          break;
        case 7:
          message.needMint = reader.bool();
          break;
        case 8:
          message.feecoin.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OutboundTx {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      processed: isSet(object.processed) ? Boolean(object.processed) : false,
      items: isObject(object.items)
        ? Object.entries(object.items).reduce<{ [key: string]: Proposals }>((acc, [key, value]) => {
          acc[key] = Proposals.fromJSON(value);
          return acc;
        }, {})
        : {},
      chainType: isSet(object.chainType) ? String(object.chainType) : "",
      inTxHash: isSet(object.inTxHash) ? String(object.inTxHash) : "",
      receiverAddress: isSet(object.receiverAddress) ? bytesFromBase64(object.receiverAddress) : new Uint8Array(),
      needMint: isSet(object.needMint) ? Boolean(object.needMint) : false,
      feecoin: Array.isArray(object?.feecoin) ? object.feecoin.map((e: any) => Coin.fromJSON(e)) : [],
    };
  },

  toJSON(message: OutboundTx): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.processed !== undefined && (obj.processed = message.processed);
    obj.items = {};
    if (message.items) {
      Object.entries(message.items).forEach(([k, v]) => {
        obj.items[k] = Proposals.toJSON(v);
      });
    }
    message.chainType !== undefined && (obj.chainType = message.chainType);
    message.inTxHash !== undefined && (obj.inTxHash = message.inTxHash);
    message.receiverAddress !== undefined
      && (obj.receiverAddress = base64FromBytes(
        message.receiverAddress !== undefined ? message.receiverAddress : new Uint8Array(),
      ));
    message.needMint !== undefined && (obj.needMint = message.needMint);
    if (message.feecoin) {
      obj.feecoin = message.feecoin.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.feecoin = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OutboundTx>, I>>(object: I): OutboundTx {
    const message = createBaseOutboundTx();
    message.index = object.index ?? "";
    message.processed = object.processed ?? false;
    message.items = Object.entries(object.items ?? {}).reduce<{ [key: string]: Proposals }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = Proposals.fromPartial(value);
      }
      return acc;
    }, {});
    message.chainType = object.chainType ?? "";
    message.inTxHash = object.inTxHash ?? "";
    message.receiverAddress = object.receiverAddress ?? new Uint8Array();
    message.needMint = object.needMint ?? false;
    message.feecoin = object.feecoin?.map((e) => Coin.fromPartial(e)) || [];
    return message;
  },
};

function createBaseOutboundTx_ItemsEntry(): OutboundTx_ItemsEntry {
  return { key: "", value: undefined };
}

export const OutboundTx_ItemsEntry = {
  encode(message: OutboundTx_ItemsEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      Proposals.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OutboundTx_ItemsEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOutboundTx_ItemsEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = Proposals.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OutboundTx_ItemsEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? Proposals.fromJSON(object.value) : undefined,
    };
  },

  toJSON(message: OutboundTx_ItemsEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value ? Proposals.toJSON(message.value) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OutboundTx_ItemsEntry>, I>>(object: I): OutboundTx_ItemsEntry {
    const message = createBaseOutboundTx_ItemsEntry();
    message.key = object.key ?? "";
    message.value = (object.value !== undefined && object.value !== null)
      ? Proposals.fromPartial(object.value)
      : undefined;
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

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
