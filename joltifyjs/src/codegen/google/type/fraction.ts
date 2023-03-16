import { Long, DeepPartial } from "../../helpers";
import * as _m0 from "protobufjs/minimal";
/** Represents a fraction in terms of a numerator divided by a denominator. */

export interface Fraction {
  /** The numerator in the fraction, e.g. 2 in 2/3. */
  numerator: Long;
  /**
   * The value by which the numerator is divided, e.g. 3 in 2/3. Must be
   * positive.
   */

  denominator: Long;
}
/** Represents a fraction in terms of a numerator divided by a denominator. */

export interface FractionSDKType {
  numerator: Long;
  denominator: Long;
}

function createBaseFraction(): Fraction {
  return {
    numerator: Long.ZERO,
    denominator: Long.ZERO
  };
}

export const Fraction = {
  encode(message: Fraction, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (!message.numerator.isZero()) {
      writer.uint32(8).int64(message.numerator);
    }

    if (!message.denominator.isZero()) {
      writer.uint32(16).int64(message.denominator);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Fraction {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFraction();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.numerator = (reader.int64() as Long);
          break;

        case 2:
          message.denominator = (reader.int64() as Long);
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<Fraction>): Fraction {
    const message = createBaseFraction();
    message.numerator = object.numerator !== undefined && object.numerator !== null ? Long.fromValue(object.numerator) : Long.ZERO;
    message.denominator = object.denominator !== undefined && object.denominator !== null ? Long.fromValue(object.denominator) : Long.ZERO;
    return message;
  }

};