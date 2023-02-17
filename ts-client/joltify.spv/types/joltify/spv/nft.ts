/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";
import { Timestamp } from "../../google/protobuf/timestamp";

export const protobufPackage = "joltify.spv";

export interface NftInfo {
  issuer: string;
  receiver: string;
  ratio: string;
  issueTime: Date | undefined;
}

export interface BorrowInterest {
  poolIndex: string;
  apy: string;
  payFreq: number;
  issueTime: Date | undefined;
  borrowed: Coin | undefined;
}

function createBaseNftInfo(): NftInfo {
  return { issuer: "", receiver: "", ratio: "", issueTime: undefined };
}

export const NftInfo = {
  encode(message: NftInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.issuer !== "") {
      writer.uint32(10).string(message.issuer);
    }
    if (message.receiver !== "") {
      writer.uint32(18).string(message.receiver);
    }
    if (message.ratio !== "") {
      writer.uint32(26).string(message.ratio);
    }
    if (message.issueTime !== undefined) {
      Timestamp.encode(toTimestamp(message.issueTime), writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): NftInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseNftInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.issuer = reader.string();
          break;
        case 2:
          message.receiver = reader.string();
          break;
        case 3:
          message.ratio = reader.string();
          break;
        case 4:
          message.issueTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NftInfo {
    return {
      issuer: isSet(object.issuer) ? String(object.issuer) : "",
      receiver: isSet(object.receiver) ? String(object.receiver) : "",
      ratio: isSet(object.ratio) ? String(object.ratio) : "",
      issueTime: isSet(object.issueTime) ? fromJsonTimestamp(object.issueTime) : undefined,
    };
  },

  toJSON(message: NftInfo): unknown {
    const obj: any = {};
    message.issuer !== undefined && (obj.issuer = message.issuer);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.ratio !== undefined && (obj.ratio = message.ratio);
    message.issueTime !== undefined && (obj.issueTime = message.issueTime.toISOString());
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<NftInfo>, I>>(object: I): NftInfo {
    const message = createBaseNftInfo();
    message.issuer = object.issuer ?? "";
    message.receiver = object.receiver ?? "";
    message.ratio = object.ratio ?? "";
    message.issueTime = object.issueTime ?? undefined;
    return message;
  },
};

function createBaseBorrowInterest(): BorrowInterest {
  return { poolIndex: "", apy: "", payFreq: 0, issueTime: undefined, borrowed: undefined };
}

export const BorrowInterest = {
  encode(message: BorrowInterest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.poolIndex !== "") {
      writer.uint32(10).string(message.poolIndex);
    }
    if (message.apy !== "") {
      writer.uint32(18).string(message.apy);
    }
    if (message.payFreq !== 0) {
      writer.uint32(24).int32(message.payFreq);
    }
    if (message.issueTime !== undefined) {
      Timestamp.encode(toTimestamp(message.issueTime), writer.uint32(34).fork()).ldelim();
    }
    if (message.borrowed !== undefined) {
      Coin.encode(message.borrowed, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BorrowInterest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBorrowInterest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.poolIndex = reader.string();
          break;
        case 2:
          message.apy = reader.string();
          break;
        case 3:
          message.payFreq = reader.int32();
          break;
        case 4:
          message.issueTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 5:
          message.borrowed = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): BorrowInterest {
    return {
      poolIndex: isSet(object.poolIndex) ? String(object.poolIndex) : "",
      apy: isSet(object.apy) ? String(object.apy) : "",
      payFreq: isSet(object.payFreq) ? Number(object.payFreq) : 0,
      issueTime: isSet(object.issueTime) ? fromJsonTimestamp(object.issueTime) : undefined,
      borrowed: isSet(object.borrowed) ? Coin.fromJSON(object.borrowed) : undefined,
    };
  },

  toJSON(message: BorrowInterest): unknown {
    const obj: any = {};
    message.poolIndex !== undefined && (obj.poolIndex = message.poolIndex);
    message.apy !== undefined && (obj.apy = message.apy);
    message.payFreq !== undefined && (obj.payFreq = Math.round(message.payFreq));
    message.issueTime !== undefined && (obj.issueTime = message.issueTime.toISOString());
    message.borrowed !== undefined && (obj.borrowed = message.borrowed ? Coin.toJSON(message.borrowed) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<BorrowInterest>, I>>(object: I): BorrowInterest {
    const message = createBaseBorrowInterest();
    message.poolIndex = object.poolIndex ?? "";
    message.apy = object.apy ?? "";
    message.payFreq = object.payFreq ?? 0;
    message.issueTime = object.issueTime ?? undefined;
    message.borrowed = (object.borrowed !== undefined && object.borrowed !== null)
      ? Coin.fromPartial(object.borrowed)
      : undefined;
    return message;
  },
};

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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
