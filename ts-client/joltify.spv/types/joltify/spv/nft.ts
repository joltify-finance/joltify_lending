/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";
import { Timestamp } from "../../google/protobuf/timestamp";

export const protobufPackage = "joltify.spv";

export interface borrowDetail {
  borrowedAmount: Coin | undefined;
  timeStamp: Date | undefined;
}

export interface NftInfo {
  issuer: string;
  receiver: string;
  borrowed: Coin | undefined;
  lastPayment: Date | undefined;
}

export interface PaymentItem {
  paymentTime: Date | undefined;
  paymentAmount: Coin | undefined;
}

export interface BorrowInterest {
  poolIndex: string;
  apy: string;
  payFreq: number;
  issueTime:
    | Date
    | undefined;
  /**
   * cosmos.base.v1beta1.Coin borrowed_last = 6[
   * 		(gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins",
   * 		(gogoproto.nullable) = false];
   */
  borrowDetails: borrowDetail[];
  monthlyRatio: string;
  interestSPY: string;
  payments: PaymentItem[];
  interestPaid: Coin | undefined;
}

function createBaseborrowDetail(): borrowDetail {
  return { borrowedAmount: undefined, timeStamp: undefined };
}

export const borrowDetail = {
  encode(message: borrowDetail, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.borrowedAmount !== undefined) {
      Coin.encode(message.borrowedAmount, writer.uint32(10).fork()).ldelim();
    }
    if (message.timeStamp !== undefined) {
      Timestamp.encode(toTimestamp(message.timeStamp), writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): borrowDetail {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseborrowDetail();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.borrowedAmount = Coin.decode(reader, reader.uint32());
          break;
        case 2:
          message.timeStamp = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): borrowDetail {
    return {
      borrowedAmount: isSet(object.borrowedAmount) ? Coin.fromJSON(object.borrowedAmount) : undefined,
      timeStamp: isSet(object.timeStamp) ? fromJsonTimestamp(object.timeStamp) : undefined,
    };
  },

  toJSON(message: borrowDetail): unknown {
    const obj: any = {};
    message.borrowedAmount !== undefined
      && (obj.borrowedAmount = message.borrowedAmount ? Coin.toJSON(message.borrowedAmount) : undefined);
    message.timeStamp !== undefined && (obj.timeStamp = message.timeStamp.toISOString());
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<borrowDetail>, I>>(object: I): borrowDetail {
    const message = createBaseborrowDetail();
    message.borrowedAmount = (object.borrowedAmount !== undefined && object.borrowedAmount !== null)
      ? Coin.fromPartial(object.borrowedAmount)
      : undefined;
    message.timeStamp = object.timeStamp ?? undefined;
    return message;
  },
};

function createBaseNftInfo(): NftInfo {
  return { issuer: "", receiver: "", borrowed: undefined, lastPayment: undefined };
}

export const NftInfo = {
  encode(message: NftInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.issuer !== "") {
      writer.uint32(10).string(message.issuer);
    }
    if (message.receiver !== "") {
      writer.uint32(18).string(message.receiver);
    }
    if (message.borrowed !== undefined) {
      Coin.encode(message.borrowed, writer.uint32(26).fork()).ldelim();
    }
    if (message.lastPayment !== undefined) {
      Timestamp.encode(toTimestamp(message.lastPayment), writer.uint32(34).fork()).ldelim();
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
          message.borrowed = Coin.decode(reader, reader.uint32());
          break;
        case 4:
          message.lastPayment = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
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
      borrowed: isSet(object.borrowed) ? Coin.fromJSON(object.borrowed) : undefined,
      lastPayment: isSet(object.lastPayment) ? fromJsonTimestamp(object.lastPayment) : undefined,
    };
  },

  toJSON(message: NftInfo): unknown {
    const obj: any = {};
    message.issuer !== undefined && (obj.issuer = message.issuer);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.borrowed !== undefined && (obj.borrowed = message.borrowed ? Coin.toJSON(message.borrowed) : undefined);
    message.lastPayment !== undefined && (obj.lastPayment = message.lastPayment.toISOString());
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<NftInfo>, I>>(object: I): NftInfo {
    const message = createBaseNftInfo();
    message.issuer = object.issuer ?? "";
    message.receiver = object.receiver ?? "";
    message.borrowed = (object.borrowed !== undefined && object.borrowed !== null)
      ? Coin.fromPartial(object.borrowed)
      : undefined;
    message.lastPayment = object.lastPayment ?? undefined;
    return message;
  },
};

function createBasePaymentItem(): PaymentItem {
  return { paymentTime: undefined, paymentAmount: undefined };
}

export const PaymentItem = {
  encode(message: PaymentItem, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.paymentTime !== undefined) {
      Timestamp.encode(toTimestamp(message.paymentTime), writer.uint32(10).fork()).ldelim();
    }
    if (message.paymentAmount !== undefined) {
      Coin.encode(message.paymentAmount, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PaymentItem {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePaymentItem();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.paymentTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 2:
          message.paymentAmount = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PaymentItem {
    return {
      paymentTime: isSet(object.paymentTime) ? fromJsonTimestamp(object.paymentTime) : undefined,
      paymentAmount: isSet(object.paymentAmount) ? Coin.fromJSON(object.paymentAmount) : undefined,
    };
  },

  toJSON(message: PaymentItem): unknown {
    const obj: any = {};
    message.paymentTime !== undefined && (obj.paymentTime = message.paymentTime.toISOString());
    message.paymentAmount !== undefined
      && (obj.paymentAmount = message.paymentAmount ? Coin.toJSON(message.paymentAmount) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PaymentItem>, I>>(object: I): PaymentItem {
    const message = createBasePaymentItem();
    message.paymentTime = object.paymentTime ?? undefined;
    message.paymentAmount = (object.paymentAmount !== undefined && object.paymentAmount !== null)
      ? Coin.fromPartial(object.paymentAmount)
      : undefined;
    return message;
  },
};

function createBaseBorrowInterest(): BorrowInterest {
  return {
    poolIndex: "",
    apy: "",
    payFreq: 0,
    issueTime: undefined,
    borrowDetails: [],
    monthlyRatio: "",
    interestSPY: "",
    payments: [],
    interestPaid: undefined,
  };
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
    for (const v of message.borrowDetails) {
      borrowDetail.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    if (message.monthlyRatio !== "") {
      writer.uint32(50).string(message.monthlyRatio);
    }
    if (message.interestSPY !== "") {
      writer.uint32(58).string(message.interestSPY);
    }
    for (const v of message.payments) {
      PaymentItem.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    if (message.interestPaid !== undefined) {
      Coin.encode(message.interestPaid, writer.uint32(74).fork()).ldelim();
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
          message.borrowDetails.push(borrowDetail.decode(reader, reader.uint32()));
          break;
        case 6:
          message.monthlyRatio = reader.string();
          break;
        case 7:
          message.interestSPY = reader.string();
          break;
        case 8:
          message.payments.push(PaymentItem.decode(reader, reader.uint32()));
          break;
        case 9:
          message.interestPaid = Coin.decode(reader, reader.uint32());
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
      borrowDetails: Array.isArray(object?.borrowDetails)
        ? object.borrowDetails.map((e: any) => borrowDetail.fromJSON(e))
        : [],
      monthlyRatio: isSet(object.monthlyRatio) ? String(object.monthlyRatio) : "",
      interestSPY: isSet(object.interestSPY) ? String(object.interestSPY) : "",
      payments: Array.isArray(object?.payments) ? object.payments.map((e: any) => PaymentItem.fromJSON(e)) : [],
      interestPaid: isSet(object.interestPaid) ? Coin.fromJSON(object.interestPaid) : undefined,
    };
  },

  toJSON(message: BorrowInterest): unknown {
    const obj: any = {};
    message.poolIndex !== undefined && (obj.poolIndex = message.poolIndex);
    message.apy !== undefined && (obj.apy = message.apy);
    message.payFreq !== undefined && (obj.payFreq = Math.round(message.payFreq));
    message.issueTime !== undefined && (obj.issueTime = message.issueTime.toISOString());
    if (message.borrowDetails) {
      obj.borrowDetails = message.borrowDetails.map((e) => e ? borrowDetail.toJSON(e) : undefined);
    } else {
      obj.borrowDetails = [];
    }
    message.monthlyRatio !== undefined && (obj.monthlyRatio = message.monthlyRatio);
    message.interestSPY !== undefined && (obj.interestSPY = message.interestSPY);
    if (message.payments) {
      obj.payments = message.payments.map((e) => e ? PaymentItem.toJSON(e) : undefined);
    } else {
      obj.payments = [];
    }
    message.interestPaid !== undefined
      && (obj.interestPaid = message.interestPaid ? Coin.toJSON(message.interestPaid) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<BorrowInterest>, I>>(object: I): BorrowInterest {
    const message = createBaseBorrowInterest();
    message.poolIndex = object.poolIndex ?? "";
    message.apy = object.apy ?? "";
    message.payFreq = object.payFreq ?? 0;
    message.issueTime = object.issueTime ?? undefined;
    message.borrowDetails = object.borrowDetails?.map((e) => borrowDetail.fromPartial(e)) || [];
    message.monthlyRatio = object.monthlyRatio ?? "";
    message.interestSPY = object.interestSPY ?? "";
    message.payments = object.payments?.map((e) => PaymentItem.fromPartial(e)) || [];
    message.interestPaid = (object.interestPaid !== undefined && object.interestPaid !== null)
      ? Coin.fromPartial(object.interestPaid)
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
