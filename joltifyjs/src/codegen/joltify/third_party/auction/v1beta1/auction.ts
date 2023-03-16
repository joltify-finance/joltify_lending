import { Coin, CoinSDKType } from "../../../../cosmos/base/v1beta1/coin";
import { Timestamp } from "../../../../google/protobuf/timestamp";
import { Long, toTimestamp, fromTimestamp, DeepPartial } from "../../../../helpers";
import * as _m0 from "protobufjs/minimal";
/** BaseAuction defines common attributes of all auctions */

export interface BaseAuction {
  id: Long;
  initiator: string;
  lot?: Coin;
  bidder: Uint8Array;
  bid?: Coin;
  hasReceivedBids: boolean;
  endTime?: Date;
  maxEndTime?: Date;
}
/** BaseAuction defines common attributes of all auctions */

export interface BaseAuctionSDKType {
  id: Long;
  initiator: string;
  lot?: CoinSDKType;
  bidder: Uint8Array;
  bid?: CoinSDKType;
  has_received_bids: boolean;
  end_time?: Date;
  max_end_time?: Date;
}
/**
 * SurplusAuction is a forward auction that burns what it receives from bids.
 * It is normally used to sell off excess pegged asset acquired by the CDP system.
 */

export interface SurplusAuction {
  baseAuction?: BaseAuction;
}
/**
 * SurplusAuction is a forward auction that burns what it receives from bids.
 * It is normally used to sell off excess pegged asset acquired by the CDP system.
 */

export interface SurplusAuctionSDKType {
  base_auction?: BaseAuctionSDKType;
}
/**
 * DebtAuction is a reverse auction that mints what it pays out.
 * It is normally used to acquire pegged asset to cover the CDP system's debts that were not covered by selling
 * collateral.
 */

export interface DebtAuction {
  baseAuction?: BaseAuction;
  correspondingDebt?: Coin;
}
/**
 * DebtAuction is a reverse auction that mints what it pays out.
 * It is normally used to acquire pegged asset to cover the CDP system's debts that were not covered by selling
 * collateral.
 */

export interface DebtAuctionSDKType {
  base_auction?: BaseAuctionSDKType;
  corresponding_debt?: CoinSDKType;
}
/**
 * CollateralAuction is a two phase auction.
 * Initially, in forward auction phase, bids can be placed up to a max bid.
 * Then it switches to a reverse auction phase, where the initial amount up for auction is bid down.
 * Unsold Lot is sent to LotReturns, being divided among the addresses by weight.
 * Collateral auctions are normally used to sell off collateral seized from CDPs.
 */

export interface CollateralAuction {
  baseAuction?: BaseAuction;
  correspondingDebt?: Coin;
  maxBid?: Coin;
  lotReturns?: WeightedAddresses;
}
/**
 * CollateralAuction is a two phase auction.
 * Initially, in forward auction phase, bids can be placed up to a max bid.
 * Then it switches to a reverse auction phase, where the initial amount up for auction is bid down.
 * Unsold Lot is sent to LotReturns, being divided among the addresses by weight.
 * Collateral auctions are normally used to sell off collateral seized from CDPs.
 */

export interface CollateralAuctionSDKType {
  base_auction?: BaseAuctionSDKType;
  corresponding_debt?: CoinSDKType;
  max_bid?: CoinSDKType;
  lot_returns?: WeightedAddressesSDKType;
}
/** WeightedAddresses is a type for storing some addresses and associated weights. */

export interface WeightedAddresses {
  addresses: Uint8Array[];
  weights: Uint8Array[];
}
/** WeightedAddresses is a type for storing some addresses and associated weights. */

export interface WeightedAddressesSDKType {
  addresses: Uint8Array[];
  weights: Uint8Array[];
}

function createBaseBaseAuction(): BaseAuction {
  return {
    id: Long.UZERO,
    initiator: "",
    lot: undefined,
    bidder: new Uint8Array(),
    bid: undefined,
    hasReceivedBids: false,
    endTime: undefined,
    maxEndTime: undefined
  };
}

export const BaseAuction = {
  encode(message: BaseAuction, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (!message.id.isZero()) {
      writer.uint32(8).uint64(message.id);
    }

    if (message.initiator !== "") {
      writer.uint32(18).string(message.initiator);
    }

    if (message.lot !== undefined) {
      Coin.encode(message.lot, writer.uint32(26).fork()).ldelim();
    }

    if (message.bidder.length !== 0) {
      writer.uint32(34).bytes(message.bidder);
    }

    if (message.bid !== undefined) {
      Coin.encode(message.bid, writer.uint32(42).fork()).ldelim();
    }

    if (message.hasReceivedBids === true) {
      writer.uint32(48).bool(message.hasReceivedBids);
    }

    if (message.endTime !== undefined) {
      Timestamp.encode(toTimestamp(message.endTime), writer.uint32(58).fork()).ldelim();
    }

    if (message.maxEndTime !== undefined) {
      Timestamp.encode(toTimestamp(message.maxEndTime), writer.uint32(66).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BaseAuction {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBaseAuction();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.id = (reader.uint64() as Long);
          break;

        case 2:
          message.initiator = reader.string();
          break;

        case 3:
          message.lot = Coin.decode(reader, reader.uint32());
          break;

        case 4:
          message.bidder = reader.bytes();
          break;

        case 5:
          message.bid = Coin.decode(reader, reader.uint32());
          break;

        case 6:
          message.hasReceivedBids = reader.bool();
          break;

        case 7:
          message.endTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;

        case 8:
          message.maxEndTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<BaseAuction>): BaseAuction {
    const message = createBaseBaseAuction();
    message.id = object.id !== undefined && object.id !== null ? Long.fromValue(object.id) : Long.UZERO;
    message.initiator = object.initiator ?? "";
    message.lot = object.lot !== undefined && object.lot !== null ? Coin.fromPartial(object.lot) : undefined;
    message.bidder = object.bidder ?? new Uint8Array();
    message.bid = object.bid !== undefined && object.bid !== null ? Coin.fromPartial(object.bid) : undefined;
    message.hasReceivedBids = object.hasReceivedBids ?? false;
    message.endTime = object.endTime ?? undefined;
    message.maxEndTime = object.maxEndTime ?? undefined;
    return message;
  }

};

function createBaseSurplusAuction(): SurplusAuction {
  return {
    baseAuction: undefined
  };
}

export const SurplusAuction = {
  encode(message: SurplusAuction, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.baseAuction !== undefined) {
      BaseAuction.encode(message.baseAuction, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SurplusAuction {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSurplusAuction();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.baseAuction = BaseAuction.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<SurplusAuction>): SurplusAuction {
    const message = createBaseSurplusAuction();
    message.baseAuction = object.baseAuction !== undefined && object.baseAuction !== null ? BaseAuction.fromPartial(object.baseAuction) : undefined;
    return message;
  }

};

function createBaseDebtAuction(): DebtAuction {
  return {
    baseAuction: undefined,
    correspondingDebt: undefined
  };
}

export const DebtAuction = {
  encode(message: DebtAuction, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.baseAuction !== undefined) {
      BaseAuction.encode(message.baseAuction, writer.uint32(10).fork()).ldelim();
    }

    if (message.correspondingDebt !== undefined) {
      Coin.encode(message.correspondingDebt, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DebtAuction {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDebtAuction();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.baseAuction = BaseAuction.decode(reader, reader.uint32());
          break;

        case 2:
          message.correspondingDebt = Coin.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<DebtAuction>): DebtAuction {
    const message = createBaseDebtAuction();
    message.baseAuction = object.baseAuction !== undefined && object.baseAuction !== null ? BaseAuction.fromPartial(object.baseAuction) : undefined;
    message.correspondingDebt = object.correspondingDebt !== undefined && object.correspondingDebt !== null ? Coin.fromPartial(object.correspondingDebt) : undefined;
    return message;
  }

};

function createBaseCollateralAuction(): CollateralAuction {
  return {
    baseAuction: undefined,
    correspondingDebt: undefined,
    maxBid: undefined,
    lotReturns: undefined
  };
}

export const CollateralAuction = {
  encode(message: CollateralAuction, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.baseAuction !== undefined) {
      BaseAuction.encode(message.baseAuction, writer.uint32(10).fork()).ldelim();
    }

    if (message.correspondingDebt !== undefined) {
      Coin.encode(message.correspondingDebt, writer.uint32(18).fork()).ldelim();
    }

    if (message.maxBid !== undefined) {
      Coin.encode(message.maxBid, writer.uint32(26).fork()).ldelim();
    }

    if (message.lotReturns !== undefined) {
      WeightedAddresses.encode(message.lotReturns, writer.uint32(34).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CollateralAuction {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCollateralAuction();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.baseAuction = BaseAuction.decode(reader, reader.uint32());
          break;

        case 2:
          message.correspondingDebt = Coin.decode(reader, reader.uint32());
          break;

        case 3:
          message.maxBid = Coin.decode(reader, reader.uint32());
          break;

        case 4:
          message.lotReturns = WeightedAddresses.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<CollateralAuction>): CollateralAuction {
    const message = createBaseCollateralAuction();
    message.baseAuction = object.baseAuction !== undefined && object.baseAuction !== null ? BaseAuction.fromPartial(object.baseAuction) : undefined;
    message.correspondingDebt = object.correspondingDebt !== undefined && object.correspondingDebt !== null ? Coin.fromPartial(object.correspondingDebt) : undefined;
    message.maxBid = object.maxBid !== undefined && object.maxBid !== null ? Coin.fromPartial(object.maxBid) : undefined;
    message.lotReturns = object.lotReturns !== undefined && object.lotReturns !== null ? WeightedAddresses.fromPartial(object.lotReturns) : undefined;
    return message;
  }

};

function createBaseWeightedAddresses(): WeightedAddresses {
  return {
    addresses: [],
    weights: []
  };
}

export const WeightedAddresses = {
  encode(message: WeightedAddresses, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.addresses) {
      writer.uint32(10).bytes(v!);
    }

    for (const v of message.weights) {
      writer.uint32(18).bytes(v!);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WeightedAddresses {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWeightedAddresses();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.addresses.push(reader.bytes());
          break;

        case 2:
          message.weights.push(reader.bytes());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<WeightedAddresses>): WeightedAddresses {
    const message = createBaseWeightedAddresses();
    message.addresses = object.addresses?.map(e => e) || [];
    message.weights = object.weights?.map(e => e) || [];
    return message;
  }

};