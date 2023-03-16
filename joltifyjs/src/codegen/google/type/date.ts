import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
/**
 * Represents a whole or partial calendar date, such as a birthday. The time of
 * day and time zone are either specified elsewhere or are insignificant. The
 * date is relative to the Gregorian Calendar. This can represent one of the
 * following:
 * 
 * * A full date, with non-zero year, month, and day values
 * * A month and day value, with a zero year, such as an anniversary
 * * A year on its own, with zero month and day values
 * * A year and month value, with a zero day, such as a credit card expiration
 * date
 * 
 * Related types are [google.type.TimeOfDay][google.type.TimeOfDay] and
 * `google.protobuf.Timestamp`.
 */

export interface Date {
  /**
   * Year of the date. Must be from 1 to 9999, or 0 to specify a date without
   * a year.
   */
  year: number;
  /**
   * Month of a year. Must be from 1 to 12, or 0 to specify a year without a
   * month and day.
   */

  month: number;
  /**
   * Day of a month. Must be from 1 to 31 and valid for the year and month, or 0
   * to specify a year by itself or a year and month where the day isn't
   * significant.
   */

  day: number;
}
/**
 * Represents a whole or partial calendar date, such as a birthday. The time of
 * day and time zone are either specified elsewhere or are insignificant. The
 * date is relative to the Gregorian Calendar. This can represent one of the
 * following:
 * 
 * * A full date, with non-zero year, month, and day values
 * * A month and day value, with a zero year, such as an anniversary
 * * A year on its own, with zero month and day values
 * * A year and month value, with a zero day, such as a credit card expiration
 * date
 * 
 * Related types are [google.type.TimeOfDay][google.type.TimeOfDay] and
 * `google.protobuf.Timestamp`.
 */

export interface DateSDKType {
  year: number;
  month: number;
  day: number;
}

function createBaseDate(): Date {
  return {
    year: 0,
    month: 0,
    day: 0
  };
}

export const Date = {
  encode(message: Date, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.year !== 0) {
      writer.uint32(8).int32(message.year);
    }

    if (message.month !== 0) {
      writer.uint32(16).int32(message.month);
    }

    if (message.day !== 0) {
      writer.uint32(24).int32(message.day);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Date {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDate();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.year = reader.int32();
          break;

        case 2:
          message.month = reader.int32();
          break;

        case 3:
          message.day = reader.int32();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<Date>): Date {
    const message = createBaseDate();
    message.year = object.year ?? 0;
    message.month = object.month ?? 0;
    message.day = object.day ?? 0;
    return message;
  }

};