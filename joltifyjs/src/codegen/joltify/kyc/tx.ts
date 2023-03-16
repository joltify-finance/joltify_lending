import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
export interface MsgUploadInvestor {
  creator: string;
  investorId: string;
  walletAddress: string[];
}
export interface MsgUploadInvestorSDKType {
  creator: string;
  investorId: string;
  walletAddress: string[];
}
export interface MsgUploadInvestorResponse {
  wallets: string[];
}
export interface MsgUploadInvestorResponseSDKType {
  wallets: string[];
}

function createBaseMsgUploadInvestor(): MsgUploadInvestor {
  return {
    creator: "",
    investorId: "",
    walletAddress: []
  };
}

export const MsgUploadInvestor = {
  encode(message: MsgUploadInvestor, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }

    if (message.investorId !== "") {
      writer.uint32(18).string(message.investorId);
    }

    for (const v of message.walletAddress) {
      writer.uint32(26).string(v!);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUploadInvestor {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUploadInvestor();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;

        case 2:
          message.investorId = reader.string();
          break;

        case 3:
          message.walletAddress.push(reader.string());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgUploadInvestor>): MsgUploadInvestor {
    const message = createBaseMsgUploadInvestor();
    message.creator = object.creator ?? "";
    message.investorId = object.investorId ?? "";
    message.walletAddress = object.walletAddress?.map(e => e) || [];
    return message;
  }

};

function createBaseMsgUploadInvestorResponse(): MsgUploadInvestorResponse {
  return {
    wallets: []
  };
}

export const MsgUploadInvestorResponse = {
  encode(message: MsgUploadInvestorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.wallets) {
      writer.uint32(10).string(v!);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUploadInvestorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUploadInvestorResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.wallets.push(reader.string());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgUploadInvestorResponse>): MsgUploadInvestorResponse {
    const message = createBaseMsgUploadInvestorResponse();
    message.wallets = object.wallets?.map(e => e) || [];
    return message;
  }

};