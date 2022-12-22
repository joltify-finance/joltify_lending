/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "joltify.vault";

export interface addressV16 {
  address: Uint8Array[];
}

export interface OutboundTxV16 {
  index: string;
  items: { [key: string]: addressV16 };
}

export interface OutboundTxV16_ItemsEntry {
  key: string;
  value: addressV16 | undefined;
}

function createBaseaddressV16(): addressV16 {
  return { address: [] };
}

export const addressV16 = {
  encode(message: addressV16, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.address) {
      writer.uint32(10).bytes(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): addressV16 {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseaddressV16();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address.push(reader.bytes());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): addressV16 {
    return { address: Array.isArray(object?.address) ? object.address.map((e: any) => bytesFromBase64(e)) : [] };
  },

  toJSON(message: addressV16): unknown {
    const obj: any = {};
    if (message.address) {
      obj.address = message.address.map((e) => base64FromBytes(e !== undefined ? e : new Uint8Array()));
    } else {
      obj.address = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<addressV16>, I>>(object: I): addressV16 {
    const message = createBaseaddressV16();
    message.address = object.address?.map((e) => e) || [];
    return message;
  },
};

function createBaseOutboundTxV16(): OutboundTxV16 {
  return { index: "", items: {} };
}

export const OutboundTxV16 = {
  encode(message: OutboundTxV16, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    Object.entries(message.items).forEach(([key, value]) => {
      OutboundTxV16_ItemsEntry.encode({ key: key as any, value }, writer.uint32(18).fork()).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OutboundTxV16 {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOutboundTxV16();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          const entry2 = OutboundTxV16_ItemsEntry.decode(reader, reader.uint32());
          if (entry2.value !== undefined) {
            message.items[entry2.key] = entry2.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OutboundTxV16 {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      items: isObject(object.items)
        ? Object.entries(object.items).reduce<{ [key: string]: addressV16 }>((acc, [key, value]) => {
          acc[key] = addressV16.fromJSON(value);
          return acc;
        }, {})
        : {},
    };
  },

  toJSON(message: OutboundTxV16): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    obj.items = {};
    if (message.items) {
      Object.entries(message.items).forEach(([k, v]) => {
        obj.items[k] = addressV16.toJSON(v);
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OutboundTxV16>, I>>(object: I): OutboundTxV16 {
    const message = createBaseOutboundTxV16();
    message.index = object.index ?? "";
    message.items = Object.entries(object.items ?? {}).reduce<{ [key: string]: addressV16 }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = addressV16.fromPartial(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseOutboundTxV16_ItemsEntry(): OutboundTxV16_ItemsEntry {
  return { key: "", value: undefined };
}

export const OutboundTxV16_ItemsEntry = {
  encode(message: OutboundTxV16_ItemsEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      addressV16.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OutboundTxV16_ItemsEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOutboundTxV16_ItemsEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = addressV16.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OutboundTxV16_ItemsEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? addressV16.fromJSON(object.value) : undefined,
    };
  },

  toJSON(message: OutboundTxV16_ItemsEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value ? addressV16.toJSON(message.value) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OutboundTxV16_ItemsEntry>, I>>(object: I): OutboundTxV16_ItemsEntry {
    const message = createBaseOutboundTxV16_ItemsEntry();
    message.key = object.key ?? "";
    message.value = (object.value !== undefined && object.value !== null)
      ? addressV16.fromPartial(object.value)
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
