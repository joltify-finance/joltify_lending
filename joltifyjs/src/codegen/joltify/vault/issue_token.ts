import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
export interface IssueToken {
  creator: Uint8Array;
  index: string;
  coin: Uint8Array;
  receiver: Uint8Array;
}
export interface IssueTokenSDKType {
  creator: Uint8Array;
  index: string;
  coin: Uint8Array;
  receiver: Uint8Array;
}

function createBaseIssueToken(): IssueToken {
  return {
    creator: new Uint8Array(),
    index: "",
    coin: undefined,
    receiver: new Uint8Array()
  };
}

export const IssueToken = {
  encode(message: IssueToken, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator.length !== 0) {
      writer.uint32(10).bytes(message.creator);
    }

    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }

    if (message.coin !== undefined) {
      writer.uint32(26).bytes(message.coin);
    }

    if (message.receiver.length !== 0) {
      writer.uint32(34).bytes(message.receiver);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IssueToken {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIssueToken();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.creator = reader.bytes();
          break;

        case 2:
          message.index = reader.string();
          break;

        case 3:
          message.coin = reader.bytes();
          break;

        case 4:
          message.receiver = reader.bytes();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<IssueToken>): IssueToken {
    const message = createBaseIssueToken();
    message.creator = object.creator ?? new Uint8Array();
    message.index = object.index ?? "";
    message.coin = object.coin ?? undefined;
    message.receiver = object.receiver ?? new Uint8Array();
    return message;
  }

};