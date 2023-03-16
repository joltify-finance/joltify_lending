import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
export interface Investor {
  investorId: string;
  walletAddress: string[];
}
export interface InvestorSDKType {
  investorId: string;
  walletAddress: string[];
}

function createBaseInvestor(): Investor {
  return {
    investorId: "",
    walletAddress: []
  };
}

export const Investor = {
  encode(message: Investor, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.investorId !== "") {
      writer.uint32(10).string(message.investorId);
    }

    for (const v of message.walletAddress) {
      writer.uint32(18).string(v!);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Investor {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInvestor();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.investorId = reader.string();
          break;

        case 2:
          message.walletAddress.push(reader.string());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<Investor>): Investor {
    const message = createBaseInvestor();
    message.investorId = object.investorId ?? "";
    message.walletAddress = object.walletAddress?.map(e => e) || [];
    return message;
  }

};