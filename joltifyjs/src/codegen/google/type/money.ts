import { Long, DeepPartial } from "../../helpers";
import * as _m0 from "protobufjs/minimal";
/** Represents an amount of money with its currency type. */

export interface Money {
  /** The three-letter currency code defined in ISO 4217. */
  currencyCode: string;
  /**
   * The whole units of the amount.
   * For example if `currencyCode` is `"USD"`, then 1 unit is one US dollar.
   */

  units: Long;
  /**
   * Number of nano (10^-9) units of the amount.
   * The value must be between -999,999,999 and +999,999,999 inclusive.
   * If `units` is positive, `nanos` must be positive or zero.
   * If `units` is zero, `nanos` can be positive, zero, or negative.
   * If `units` is negative, `nanos` must be negative or zero.
   * For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.
   */

  nanos: number;
}
/** Represents an amount of money with its currency type. */

export interface MoneySDKType {
  currency_code: string;
  units: Long;
  nanos: number;
}

function createBaseMoney(): Money {
  return {
    currencyCode: "",
    units: Long.ZERO,
    nanos: 0
  };
}

export const Money = {
  encode(message: Money, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.currencyCode !== "") {
      writer.uint32(10).string(message.currencyCode);
    }

    if (!message.units.isZero()) {
      writer.uint32(16).int64(message.units);
    }

    if (message.nanos !== 0) {
      writer.uint32(24).int32(message.nanos);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Money {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMoney();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.currencyCode = reader.string();
          break;

        case 2:
          message.units = (reader.int64() as Long);
          break;

        case 3:
          message.nanos = reader.int32();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<Money>): Money {
    const message = createBaseMoney();
    message.currencyCode = object.currencyCode ?? "";
    message.units = object.units !== undefined && object.units !== null ? Long.fromValue(object.units) : Long.ZERO;
    message.nanos = object.nanos ?? 0;
    return message;
  }

};