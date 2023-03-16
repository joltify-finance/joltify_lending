import { Coin, CoinSDKType } from "../../cosmos/base/v1beta1/coin";
import { Timestamp } from "../../google/protobuf/timestamp";
import * as _m0 from "protobufjs/minimal";
import { toTimestamp, fromTimestamp, DeepPartial } from "../../helpers";
export interface BorrowDetail {
  borrowedAmount?: Coin;
  timeStamp?: Date;
}
export interface BorrowDetailSDKType {
  borrowed_amount?: CoinSDKType;
  time_stamp?: Date;
}
export interface NftInfo {
  issuer: string;
  receiver: string;
  borrowed?: Coin;
  lastPayment?: Date;
}
export interface NftInfoSDKType {
  issuer: string;
  receiver: string;
  borrowed?: CoinSDKType;
  last_payment?: Date;
}
export interface PaymentItem {
  paymentTime?: Date;
  paymentAmount?: Coin;
}
export interface PaymentItemSDKType {
  payment_time?: Date;
  payment_amount?: CoinSDKType;
}
export interface BorrowInterest {
  poolIndex: string;
  apy: string;
  payFreq: number;
  issueTime?: Date;
  borrowDetails: BorrowDetail[];
  monthlyRatio: string;
  interestSPY: string;
  payments: PaymentItem[];
  interestPaid?: Coin;
}
export interface BorrowInterestSDKType {
  pool_index: string;
  apy: string;
  pay_freq: number;
  issue_time?: Date;
  borrow_details: BorrowDetailSDKType[];
  monthly_ratio: string;
  interest_sPY: string;
  payments: PaymentItemSDKType[];
  interestPaid?: CoinSDKType;
}

function createBaseBorrowDetail(): BorrowDetail {
  return {
    borrowedAmount: undefined,
    timeStamp: undefined
  };
}

export const BorrowDetail = {
  encode(message: BorrowDetail, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.borrowedAmount !== undefined) {
      Coin.encode(message.borrowedAmount, writer.uint32(10).fork()).ldelim();
    }

    if (message.timeStamp !== undefined) {
      Timestamp.encode(toTimestamp(message.timeStamp), writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BorrowDetail {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBorrowDetail();

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

  fromPartial(object: DeepPartial<BorrowDetail>): BorrowDetail {
    const message = createBaseBorrowDetail();
    message.borrowedAmount = object.borrowedAmount !== undefined && object.borrowedAmount !== null ? Coin.fromPartial(object.borrowedAmount) : undefined;
    message.timeStamp = object.timeStamp ?? undefined;
    return message;
  }

};

function createBaseNftInfo(): NftInfo {
  return {
    issuer: "",
    receiver: "",
    borrowed: undefined,
    lastPayment: undefined
  };
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

  fromPartial(object: DeepPartial<NftInfo>): NftInfo {
    const message = createBaseNftInfo();
    message.issuer = object.issuer ?? "";
    message.receiver = object.receiver ?? "";
    message.borrowed = object.borrowed !== undefined && object.borrowed !== null ? Coin.fromPartial(object.borrowed) : undefined;
    message.lastPayment = object.lastPayment ?? undefined;
    return message;
  }

};

function createBasePaymentItem(): PaymentItem {
  return {
    paymentTime: undefined,
    paymentAmount: undefined
  };
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

  fromPartial(object: DeepPartial<PaymentItem>): PaymentItem {
    const message = createBasePaymentItem();
    message.paymentTime = object.paymentTime ?? undefined;
    message.paymentAmount = object.paymentAmount !== undefined && object.paymentAmount !== null ? Coin.fromPartial(object.paymentAmount) : undefined;
    return message;
  }

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
    interestPaid: undefined
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
      BorrowDetail.encode(v!, writer.uint32(42).fork()).ldelim();
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
          message.borrowDetails.push(BorrowDetail.decode(reader, reader.uint32()));
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

  fromPartial(object: DeepPartial<BorrowInterest>): BorrowInterest {
    const message = createBaseBorrowInterest();
    message.poolIndex = object.poolIndex ?? "";
    message.apy = object.apy ?? "";
    message.payFreq = object.payFreq ?? 0;
    message.issueTime = object.issueTime ?? undefined;
    message.borrowDetails = object.borrowDetails?.map(e => BorrowDetail.fromPartial(e)) || [];
    message.monthlyRatio = object.monthlyRatio ?? "";
    message.interestSPY = object.interestSPY ?? "";
    message.payments = object.payments?.map(e => PaymentItem.fromPartial(e)) || [];
    message.interestPaid = object.interestPaid !== undefined && object.interestPaid !== null ? Coin.fromPartial(object.interestPaid) : undefined;
    return message;
  }

};