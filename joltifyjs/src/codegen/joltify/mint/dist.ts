import { Timestamp } from "../../google/protobuf/timestamp";
import { Long, toTimestamp, fromTimestamp, DeepPartial } from "../../helpers";
import * as _m0 from "protobufjs/minimal";
export interface HistoricalDistInfo {
  payoutTime?: Date;
  distributedRound: Long;
}
export interface HistoricalDistInfoSDKType {
  payout_time?: Date;
  distributed_round: Long;
}

function createBaseHistoricalDistInfo(): HistoricalDistInfo {
  return {
    payoutTime: undefined,
    distributedRound: Long.UZERO
  };
}

export const HistoricalDistInfo = {
  encode(message: HistoricalDistInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.payoutTime !== undefined) {
      Timestamp.encode(toTimestamp(message.payoutTime), writer.uint32(10).fork()).ldelim();
    }

    if (!message.distributedRound.isZero()) {
      writer.uint32(16).uint64(message.distributedRound);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): HistoricalDistInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHistoricalDistInfo();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.payoutTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;

        case 2:
          message.distributedRound = (reader.uint64() as Long);
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<HistoricalDistInfo>): HistoricalDistInfo {
    const message = createBaseHistoricalDistInfo();
    message.payoutTime = object.payoutTime ?? undefined;
    message.distributedRound = object.distributedRound !== undefined && object.distributedRound !== null ? Long.fromValue(object.distributedRound) : Long.UZERO;
    return message;
  }

};