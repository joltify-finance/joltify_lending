import { Long, DeepPartial } from "../../helpers";
import * as _m0 from "protobufjs/minimal";
/** Params defines the parameters for the module. */

export interface Params {
  halfCount: Long;
  firstProvisions: string;
  currentProvisions: string;
  unit: string;
  communityProvisions: string;
}
/** Params defines the parameters for the module. */

export interface ParamsSDKType {
  halfCount: Long;
  first_provisions: string;
  current_provisions: string;
  unit: string;
  community_provisions: string;
}

function createBaseParams(): Params {
  return {
    halfCount: Long.UZERO,
    firstProvisions: "",
    currentProvisions: "",
    unit: "",
    communityProvisions: ""
  };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (!message.halfCount.isZero()) {
      writer.uint32(8).uint64(message.halfCount);
    }

    if (message.firstProvisions !== "") {
      writer.uint32(18).string(message.firstProvisions);
    }

    if (message.currentProvisions !== "") {
      writer.uint32(26).string(message.currentProvisions);
    }

    if (message.unit !== "") {
      writer.uint32(34).string(message.unit);
    }

    if (message.communityProvisions !== "") {
      writer.uint32(42).string(message.communityProvisions);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseParams();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.halfCount = (reader.uint64() as Long);
          break;

        case 2:
          message.firstProvisions = reader.string();
          break;

        case 3:
          message.currentProvisions = reader.string();
          break;

        case 4:
          message.unit = reader.string();
          break;

        case 5:
          message.communityProvisions = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = createBaseParams();
    message.halfCount = object.halfCount !== undefined && object.halfCount !== null ? Long.fromValue(object.halfCount) : Long.UZERO;
    message.firstProvisions = object.firstProvisions ?? "";
    message.currentProvisions = object.currentProvisions ?? "";
    message.unit = object.unit ?? "";
    message.communityProvisions = object.communityProvisions ?? "";
    return message;
  }

};