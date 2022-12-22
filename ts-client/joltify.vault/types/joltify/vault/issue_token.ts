/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "joltify.vault";

export interface IssueToken {
  creator: Uint8Array;
  index: string;
  coin: Uint8Array;
  receiver: Uint8Array;
}

function createBaseIssueToken(): IssueToken {
  return { creator: new Uint8Array(), index: "", coin: new Uint8Array(), receiver: new Uint8Array() };
}

export const IssueToken = {
  encode(message: IssueToken, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator.length !== 0) {
      writer.uint32(10).bytes(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.coin.length !== 0) {
      writer.uint32(26).bytes(message.coin);
    }
    if (message.receiver.length !== 0) {
      writer.uint32(34).bytes(message.receiver);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IssueToken {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIssueToken();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.bytes();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.coin = reader.bytes();
          break;
        case 4:
          message.receiver = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): IssueToken {
    return {
      creator: isSet(object.creator) ? bytesFromBase64(object.creator) : new Uint8Array(),
      index: isSet(object.index) ? String(object.index) : "",
      coin: isSet(object.coin) ? bytesFromBase64(object.coin) : new Uint8Array(),
      receiver: isSet(object.receiver) ? bytesFromBase64(object.receiver) : new Uint8Array(),
    };
  },

  toJSON(message: IssueToken): unknown {
    const obj: any = {};
    message.creator !== undefined
      && (obj.creator = base64FromBytes(message.creator !== undefined ? message.creator : new Uint8Array()));
    message.index !== undefined && (obj.index = message.index);
    message.coin !== undefined
      && (obj.coin = base64FromBytes(message.coin !== undefined ? message.coin : new Uint8Array()));
    message.receiver !== undefined
      && (obj.receiver = base64FromBytes(message.receiver !== undefined ? message.receiver : new Uint8Array()));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<IssueToken>, I>>(object: I): IssueToken {
    const message = createBaseIssueToken();
    message.creator = object.creator ?? new Uint8Array();
    message.index = object.index ?? "";
    message.coin = object.coin ?? new Uint8Array();
    message.receiver = object.receiver ?? new Uint8Array();
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
