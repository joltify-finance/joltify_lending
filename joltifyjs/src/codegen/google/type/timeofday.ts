import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
/**
 * Represents a time of day. The date and time zone are either not significant
 * or are specified elsewhere. An API may choose to allow leap seconds. Related
 * types are [google.type.Date][google.type.Date] and
 * `google.protobuf.Timestamp`.
 */

export interface TimeOfDay {
  /**
   * Hours of day in 24 hour format. Should be from 0 to 23. An API may choose
   * to allow the value "24:00:00" for scenarios like business closing time.
   */
  hours: number;
  /** Minutes of hour of day. Must be from 0 to 59. */

  minutes: number;
  /**
   * Seconds of minutes of the time. Must normally be from 0 to 59. An API may
   * allow the value 60 if it allows leap-seconds.
   */

  seconds: number;
  /** Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999. */

  nanos: number;
}
/**
 * Represents a time of day. The date and time zone are either not significant
 * or are specified elsewhere. An API may choose to allow leap seconds. Related
 * types are [google.type.Date][google.type.Date] and
 * `google.protobuf.Timestamp`.
 */

export interface TimeOfDaySDKType {
  hours: number;
  minutes: number;
  seconds: number;
  nanos: number;
}

function createBaseTimeOfDay(): TimeOfDay {
  return {
    hours: 0,
    minutes: 0,
    seconds: 0,
    nanos: 0
  };
}

export const TimeOfDay = {
  encode(message: TimeOfDay, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.hours !== 0) {
      writer.uint32(8).int32(message.hours);
    }

    if (message.minutes !== 0) {
      writer.uint32(16).int32(message.minutes);
    }

    if (message.seconds !== 0) {
      writer.uint32(24).int32(message.seconds);
    }

    if (message.nanos !== 0) {
      writer.uint32(32).int32(message.nanos);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TimeOfDay {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTimeOfDay();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.hours = reader.int32();
          break;

        case 2:
          message.minutes = reader.int32();
          break;

        case 3:
          message.seconds = reader.int32();
          break;

        case 4:
          message.nanos = reader.int32();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<TimeOfDay>): TimeOfDay {
    const message = createBaseTimeOfDay();
    message.hours = object.hours ?? 0;
    message.minutes = object.minutes ?? 0;
    message.seconds = object.seconds ?? 0;
    message.nanos = object.nanos ?? 0;
    return message;
  }

};