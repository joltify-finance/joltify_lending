import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
/**
 * An object representing a phone number, suitable as an API wire format.
 * 
 * This representation:
 * 
 *  - should not be used for locale-specific formatting of a phone number, such
 *    as "+1 (650) 253-0000 ext. 123"
 * 
 *  - is not designed for efficient storage
 *  - may not be suitable for dialing - specialized libraries (see references)
 *    should be used to parse the number for that purpose
 * 
 * To do something meaningful with this number, such as format it for various
 * use-cases, convert it to an `i18n.phonenumbers.PhoneNumber` object first.
 * 
 * For instance, in Java this would be:
 * 
 *    com.google.type.PhoneNumber wireProto =
 *        com.google.type.PhoneNumber.newBuilder().build();
 *    com.google.i18n.phonenumbers.Phonenumber.PhoneNumber phoneNumber =
 *        PhoneNumberUtil.getInstance().parse(wireProto.getE164Number(), "ZZ");
 *    if (!wireProto.getExtension().isEmpty()) {
 *      phoneNumber.setExtension(wireProto.getExtension());
 *    }
 * 
 *  Reference(s):
 *   - https://github.com/google/libphonenumber
 */

export interface PhoneNumber {
  /**
   * The phone number, represented as a leading plus sign ('+'), followed by a
   * phone number that uses a relaxed ITU E.164 format consisting of the
   * country calling code (1 to 3 digits) and the subscriber number, with no
   * additional spaces or formatting, e.g.:
   *  - correct: "+15552220123"
   *  - incorrect: "+1 (555) 222-01234 x123".
   * 
   * The ITU E.164 format limits the latter to 12 digits, but in practice not
   * all countries respect that, so we relax that restriction here.
   * National-only numbers are not allowed.
   * 
   * References:
   *  - https://www.itu.int/rec/T-REC-E.164-201011-I
   *  - https://en.wikipedia.org/wiki/E.164.
   *  - https://en.wikipedia.org/wiki/List_of_country_calling_codes
   */
  e164Number?: string;
  /**
   * A short code.
   * 
   * Reference(s):
   *  - https://en.wikipedia.org/wiki/Short_code
   */

  shortCode?: PhoneNumber_ShortCode;
  /**
   * The phone number's extension. The extension is not standardized in ITU
   * recommendations, except for being defined as a series of numbers with a
   * maximum length of 40 digits. Other than digits, some other dialing
   * characters such as ',' (indicating a wait) or '#' may be stored here.
   * 
   * Note that no regions currently use extensions with short codes, so this
   * field is normally only set in conjunction with an E.164 number. It is held
   * separately from the E.164 number to allow for short code extensions in the
   * future.
   */

  extension: string;
}
/**
 * An object representing a phone number, suitable as an API wire format.
 * 
 * This representation:
 * 
 *  - should not be used for locale-specific formatting of a phone number, such
 *    as "+1 (650) 253-0000 ext. 123"
 * 
 *  - is not designed for efficient storage
 *  - may not be suitable for dialing - specialized libraries (see references)
 *    should be used to parse the number for that purpose
 * 
 * To do something meaningful with this number, such as format it for various
 * use-cases, convert it to an `i18n.phonenumbers.PhoneNumber` object first.
 * 
 * For instance, in Java this would be:
 * 
 *    com.google.type.PhoneNumber wireProto =
 *        com.google.type.PhoneNumber.newBuilder().build();
 *    com.google.i18n.phonenumbers.Phonenumber.PhoneNumber phoneNumber =
 *        PhoneNumberUtil.getInstance().parse(wireProto.getE164Number(), "ZZ");
 *    if (!wireProto.getExtension().isEmpty()) {
 *      phoneNumber.setExtension(wireProto.getExtension());
 *    }
 * 
 *  Reference(s):
 *   - https://github.com/google/libphonenumber
 */

export interface PhoneNumberSDKType {
  e164_number?: string;
  short_code?: PhoneNumber_ShortCodeSDKType;
  extension: string;
}
/**
 * An object representing a short code, which is a phone number that is
 * typically much shorter than regular phone numbers and can be used to
 * address messages in MMS and SMS systems, as well as for abbreviated dialing
 * (e.g. "Text 611 to see how many minutes you have remaining on your plan.").
 * 
 * Short codes are restricted to a region and are not internationally
 * dialable, which means the same short code can exist in different regions,
 * with different usage and pricing, even if those regions share the same
 * country calling code (e.g. US and CA).
 */

export interface PhoneNumber_ShortCode {
  /**
   * Required. The BCP-47 region code of the location where calls to this
   * short code can be made, such as "US" and "BB".
   * 
   * Reference(s):
   *  - http://www.unicode.org/reports/tr35/#unicode_region_subtag
   */
  regionCode: string;
  /**
   * Required. The short code digits, without a leading plus ('+') or country
   * calling code, e.g. "611".
   */

  number: string;
}
/**
 * An object representing a short code, which is a phone number that is
 * typically much shorter than regular phone numbers and can be used to
 * address messages in MMS and SMS systems, as well as for abbreviated dialing
 * (e.g. "Text 611 to see how many minutes you have remaining on your plan.").
 * 
 * Short codes are restricted to a region and are not internationally
 * dialable, which means the same short code can exist in different regions,
 * with different usage and pricing, even if those regions share the same
 * country calling code (e.g. US and CA).
 */

export interface PhoneNumber_ShortCodeSDKType {
  region_code: string;
  number: string;
}

function createBasePhoneNumber(): PhoneNumber {
  return {
    e164Number: undefined,
    shortCode: undefined,
    extension: ""
  };
}

export const PhoneNumber = {
  encode(message: PhoneNumber, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.e164Number !== undefined) {
      writer.uint32(10).string(message.e164Number);
    }

    if (message.shortCode !== undefined) {
      PhoneNumber_ShortCode.encode(message.shortCode, writer.uint32(18).fork()).ldelim();
    }

    if (message.extension !== "") {
      writer.uint32(26).string(message.extension);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PhoneNumber {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePhoneNumber();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.e164Number = reader.string();
          break;

        case 2:
          message.shortCode = PhoneNumber_ShortCode.decode(reader, reader.uint32());
          break;

        case 3:
          message.extension = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<PhoneNumber>): PhoneNumber {
    const message = createBasePhoneNumber();
    message.e164Number = object.e164Number ?? undefined;
    message.shortCode = object.shortCode !== undefined && object.shortCode !== null ? PhoneNumber_ShortCode.fromPartial(object.shortCode) : undefined;
    message.extension = object.extension ?? "";
    return message;
  }

};

function createBasePhoneNumber_ShortCode(): PhoneNumber_ShortCode {
  return {
    regionCode: "",
    number: ""
  };
}

export const PhoneNumber_ShortCode = {
  encode(message: PhoneNumber_ShortCode, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.regionCode !== "") {
      writer.uint32(10).string(message.regionCode);
    }

    if (message.number !== "") {
      writer.uint32(18).string(message.number);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PhoneNumber_ShortCode {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePhoneNumber_ShortCode();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.regionCode = reader.string();
          break;

        case 2:
          message.number = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<PhoneNumber_ShortCode>): PhoneNumber_ShortCode {
    const message = createBasePhoneNumber_ShortCode();
    message.regionCode = object.regionCode ?? "";
    message.number = object.number ?? "";
    return message;
  }

};