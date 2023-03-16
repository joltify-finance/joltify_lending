import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
/** Localized variant of a text in a particular language. */

export interface LocalizedText {
  /** Localized string in the language corresponding to `language_code' below. */
  text: string;
  /**
   * The text's BCP-47 language code, such as "en-US" or "sr-Latn".
   * 
   * For more information, see
   * http://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
   */

  languageCode: string;
}
/** Localized variant of a text in a particular language. */

export interface LocalizedTextSDKType {
  text: string;
  language_code: string;
}

function createBaseLocalizedText(): LocalizedText {
  return {
    text: "",
    languageCode: ""
  };
}

export const LocalizedText = {
  encode(message: LocalizedText, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.text !== "") {
      writer.uint32(10).string(message.text);
    }

    if (message.languageCode !== "") {
      writer.uint32(18).string(message.languageCode);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LocalizedText {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLocalizedText();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.text = reader.string();
          break;

        case 2:
          message.languageCode = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<LocalizedText>): LocalizedText {
    const message = createBaseLocalizedText();
    message.text = object.text ?? "";
    message.languageCode = object.languageCode ?? "";
    return message;
  }

};