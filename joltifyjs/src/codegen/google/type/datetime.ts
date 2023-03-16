import { Duration, DurationSDKType } from "../protobuf/duration";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
/**
 * Represents civil time (or occasionally physical time).
 * 
 * This type can represent a civil time in one of a few possible ways:
 * 
 *  * When utc_offset is set and time_zone is unset: a civil time on a calendar
 *    day with a particular offset from UTC.
 *  * When time_zone is set and utc_offset is unset: a civil time on a calendar
 *    day in a particular time zone.
 *  * When neither time_zone nor utc_offset is set: a civil time on a calendar
 *    day in local time.
 * 
 * The date is relative to the Proleptic Gregorian Calendar.
 * 
 * If year is 0, the DateTime is considered not to have a specific year. month
 * and day must have valid, non-zero values.
 * 
 * This type may also be used to represent a physical time if all the date and
 * time fields are set and either case of the `time_offset` oneof is set.
 * Consider using `Timestamp` message for physical time instead. If your use
 * case also would like to store the user's timezone, that can be done in
 * another field.
 * 
 * This type is more flexible than some applications may want. Make sure to
 * document and validate your application's limitations.
 */

export interface DateTime {
  /**
   * Optional. Year of date. Must be from 1 to 9999, or 0 if specifying a
   * datetime without a year.
   */
  year: number;
  /** Required. Month of year. Must be from 1 to 12. */

  month: number;
  /**
   * Required. Day of month. Must be from 1 to 31 and valid for the year and
   * month.
   */

  day: number;
  /**
   * Required. Hours of day in 24 hour format. Should be from 0 to 23. An API
   * may choose to allow the value "24:00:00" for scenarios like business
   * closing time.
   */

  hours: number;
  /** Required. Minutes of hour of day. Must be from 0 to 59. */

  minutes: number;
  /**
   * Required. Seconds of minutes of the time. Must normally be from 0 to 59. An
   * API may allow the value 60 if it allows leap-seconds.
   */

  seconds: number;
  /**
   * Required. Fractions of seconds in nanoseconds. Must be from 0 to
   * 999,999,999.
   */

  nanos: number;
  /**
   * UTC offset. Must be whole seconds, between -18 hours and +18 hours.
   * For example, a UTC offset of -4:00 would be represented as
   * { seconds: -14400 }.
   */

  utcOffset?: Duration;
  /** Time zone. */

  timeZone?: TimeZone;
}
/**
 * Represents civil time (or occasionally physical time).
 * 
 * This type can represent a civil time in one of a few possible ways:
 * 
 *  * When utc_offset is set and time_zone is unset: a civil time on a calendar
 *    day with a particular offset from UTC.
 *  * When time_zone is set and utc_offset is unset: a civil time on a calendar
 *    day in a particular time zone.
 *  * When neither time_zone nor utc_offset is set: a civil time on a calendar
 *    day in local time.
 * 
 * The date is relative to the Proleptic Gregorian Calendar.
 * 
 * If year is 0, the DateTime is considered not to have a specific year. month
 * and day must have valid, non-zero values.
 * 
 * This type may also be used to represent a physical time if all the date and
 * time fields are set and either case of the `time_offset` oneof is set.
 * Consider using `Timestamp` message for physical time instead. If your use
 * case also would like to store the user's timezone, that can be done in
 * another field.
 * 
 * This type is more flexible than some applications may want. Make sure to
 * document and validate your application's limitations.
 */

export interface DateTimeSDKType {
  year: number;
  month: number;
  day: number;
  hours: number;
  minutes: number;
  seconds: number;
  nanos: number;
  utc_offset?: DurationSDKType;
  time_zone?: TimeZoneSDKType;
}
/**
 * Represents a time zone from the
 * [IANA Time Zone Database](https://www.iana.org/time-zones).
 */

export interface TimeZone {
  /** IANA Time Zone Database time zone, e.g. "America/New_York". */
  id: string;
  /** Optional. IANA Time Zone Database version number, e.g. "2019a". */

  version: string;
}
/**
 * Represents a time zone from the
 * [IANA Time Zone Database](https://www.iana.org/time-zones).
 */

export interface TimeZoneSDKType {
  id: string;
  version: string;
}

function createBaseDateTime(): DateTime {
  return {
    year: 0,
    month: 0,
    day: 0,
    hours: 0,
    minutes: 0,
    seconds: 0,
    nanos: 0,
    utcOffset: undefined,
    timeZone: undefined
  };
}

export const DateTime = {
  encode(message: DateTime, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.year !== 0) {
      writer.uint32(8).int32(message.year);
    }

    if (message.month !== 0) {
      writer.uint32(16).int32(message.month);
    }

    if (message.day !== 0) {
      writer.uint32(24).int32(message.day);
    }

    if (message.hours !== 0) {
      writer.uint32(32).int32(message.hours);
    }

    if (message.minutes !== 0) {
      writer.uint32(40).int32(message.minutes);
    }

    if (message.seconds !== 0) {
      writer.uint32(48).int32(message.seconds);
    }

    if (message.nanos !== 0) {
      writer.uint32(56).int32(message.nanos);
    }

    if (message.utcOffset !== undefined) {
      Duration.encode(message.utcOffset, writer.uint32(66).fork()).ldelim();
    }

    if (message.timeZone !== undefined) {
      TimeZone.encode(message.timeZone, writer.uint32(74).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DateTime {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDateTime();

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

        case 4:
          message.hours = reader.int32();
          break;

        case 5:
          message.minutes = reader.int32();
          break;

        case 6:
          message.seconds = reader.int32();
          break;

        case 7:
          message.nanos = reader.int32();
          break;

        case 8:
          message.utcOffset = Duration.decode(reader, reader.uint32());
          break;

        case 9:
          message.timeZone = TimeZone.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<DateTime>): DateTime {
    const message = createBaseDateTime();
    message.year = object.year ?? 0;
    message.month = object.month ?? 0;
    message.day = object.day ?? 0;
    message.hours = object.hours ?? 0;
    message.minutes = object.minutes ?? 0;
    message.seconds = object.seconds ?? 0;
    message.nanos = object.nanos ?? 0;
    message.utcOffset = object.utcOffset !== undefined && object.utcOffset !== null ? Duration.fromPartial(object.utcOffset) : undefined;
    message.timeZone = object.timeZone !== undefined && object.timeZone !== null ? TimeZone.fromPartial(object.timeZone) : undefined;
    return message;
  }

};

function createBaseTimeZone(): TimeZone {
  return {
    id: "",
    version: ""
  };
}

export const TimeZone = {
  encode(message: TimeZone, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }

    if (message.version !== "") {
      writer.uint32(18).string(message.version);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TimeZone {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTimeZone();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;

        case 2:
          message.version = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<TimeZone>): TimeZone {
    const message = createBaseTimeZone();
    message.id = object.id ?? "";
    message.version = object.version ?? "";
    return message;
  }

};