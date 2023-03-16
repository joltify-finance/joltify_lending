import { Coin, CoinSDKType } from "../../cosmos/base/v1beta1/coin";
import { Long, DeepPartial } from "../../helpers";
import * as _m0 from "protobufjs/minimal";
export interface HistoricalAmount {
  blockHeight: Long;
  amount: Coin[];
}
export interface HistoricalAmountSDKType {
  blockHeight: Long;
  amount: CoinSDKType[];
}
export interface CoinsQuota {
  history: HistoricalAmount[];
  CoinsSum: Coin[];
}
export interface CoinsQuotaSDKType {
  history: HistoricalAmountSDKType[];
  Coins_sum: CoinSDKType[];
}

function createBaseHistoricalAmount(): HistoricalAmount {
  return {
    blockHeight: Long.ZERO,
    amount: []
  };
}

export const HistoricalAmount = {
  encode(message: HistoricalAmount, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (!message.blockHeight.isZero()) {
      writer.uint32(8).int64(message.blockHeight);
    }

    for (const v of message.amount) {
      Coin.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): HistoricalAmount {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHistoricalAmount();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.blockHeight = (reader.int64() as Long);
          break;

        case 2:
          message.amount.push(Coin.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<HistoricalAmount>): HistoricalAmount {
    const message = createBaseHistoricalAmount();
    message.blockHeight = object.blockHeight !== undefined && object.blockHeight !== null ? Long.fromValue(object.blockHeight) : Long.ZERO;
    message.amount = object.amount?.map(e => Coin.fromPartial(e)) || [];
    return message;
  }

};

function createBaseCoinsQuota(): CoinsQuota {
  return {
    history: [],
    CoinsSum: []
  };
}

export const CoinsQuota = {
  encode(message: CoinsQuota, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.history) {
      HistoricalAmount.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    for (const v of message.CoinsSum) {
      Coin.encode(v!, writer.uint32(26).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CoinsQuota {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCoinsQuota();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 2:
          message.history.push(HistoricalAmount.decode(reader, reader.uint32()));
          break;

        case 3:
          message.CoinsSum.push(Coin.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<CoinsQuota>): CoinsQuota {
    const message = createBaseCoinsQuota();
    message.history = object.history?.map(e => HistoricalAmount.fromPartial(e)) || [];
    message.CoinsSum = object.CoinsSum?.map(e => Coin.fromPartial(e)) || [];
    return message;
  }

};