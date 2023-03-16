import { Timestamp } from "../../../../google/protobuf/timestamp";
import { Coin, CoinSDKType } from "../../../../cosmos/base/v1beta1/coin";
import * as _m0 from "protobufjs/minimal";
import { toTimestamp, fromTimestamp, DeepPartial, Long } from "../../../../helpers";
/** RewardPeriod stores the state of an ongoing reward */

export interface RewardPeriod {
  active: boolean;
  collateralType: string;
  start?: Date;
  end?: Date;
  rewardsPerSecond?: Coin;
}
/** RewardPeriod stores the state of an ongoing reward */

export interface RewardPeriodSDKType {
  active: boolean;
  collateral_type: string;
  start?: Date;
  end?: Date;
  rewards_per_second?: CoinSDKType;
}
/** MultiRewardPeriod supports multiple reward types */

export interface MultiRewardPeriod {
  active: boolean;
  collateralType: string;
  start?: Date;
  end?: Date;
  rewardsPerSecond: Coin[];
}
/** MultiRewardPeriod supports multiple reward types */

export interface MultiRewardPeriodSDKType {
  active: boolean;
  collateral_type: string;
  start?: Date;
  end?: Date;
  rewards_per_second: CoinSDKType[];
}
/** Multiplier amount the claim rewards get increased by, along with how long the claim rewards are locked */

export interface Multiplier {
  name: string;
  monthsLockup: Long;
  factor: Uint8Array;
}
/** Multiplier amount the claim rewards get increased by, along with how long the claim rewards are locked */

export interface MultiplierSDKType {
  name: string;
  months_lockup: Long;
  factor: Uint8Array;
}
/** MultipliersPerDenom is a map of denoms to a set of multipliers */

export interface MultipliersPerDenom {
  denom: string;
  multipliers: Multiplier[];
}
/** MultipliersPerDenom is a map of denoms to a set of multipliers */

export interface MultipliersPerDenomSDKType {
  denom: string;
  multipliers: MultiplierSDKType[];
}
/** Params */

export interface Params {
  usdxMintingRewardPeriods: RewardPeriod[];
  joltSupplyRewardPeriods: MultiRewardPeriod[];
  joltBorrowRewardPeriods: MultiRewardPeriod[];
  delegatorRewardPeriods: MultiRewardPeriod[];
  swapRewardPeriods: MultiRewardPeriod[];
  claimMultipliers: MultipliersPerDenom[];
  claimEnd?: Date;
  savingsRewardPeriods: MultiRewardPeriod[];
}
/** Params */

export interface ParamsSDKType {
  usdx_minting_reward_periods: RewardPeriodSDKType[];
  jolt_supply_reward_periods: MultiRewardPeriodSDKType[];
  jolt_borrow_reward_periods: MultiRewardPeriodSDKType[];
  delegator_reward_periods: MultiRewardPeriodSDKType[];
  swap_reward_periods: MultiRewardPeriodSDKType[];
  claim_multipliers: MultipliersPerDenomSDKType[];
  claim_end?: Date;
  savings_reward_periods: MultiRewardPeriodSDKType[];
}

function createBaseRewardPeriod(): RewardPeriod {
  return {
    active: false,
    collateralType: "",
    start: undefined,
    end: undefined,
    rewardsPerSecond: undefined
  };
}

export const RewardPeriod = {
  encode(message: RewardPeriod, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.active === true) {
      writer.uint32(8).bool(message.active);
    }

    if (message.collateralType !== "") {
      writer.uint32(18).string(message.collateralType);
    }

    if (message.start !== undefined) {
      Timestamp.encode(toTimestamp(message.start), writer.uint32(26).fork()).ldelim();
    }

    if (message.end !== undefined) {
      Timestamp.encode(toTimestamp(message.end), writer.uint32(34).fork()).ldelim();
    }

    if (message.rewardsPerSecond !== undefined) {
      Coin.encode(message.rewardsPerSecond, writer.uint32(42).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RewardPeriod {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRewardPeriod();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.active = reader.bool();
          break;

        case 2:
          message.collateralType = reader.string();
          break;

        case 3:
          message.start = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;

        case 4:
          message.end = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;

        case 5:
          message.rewardsPerSecond = Coin.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<RewardPeriod>): RewardPeriod {
    const message = createBaseRewardPeriod();
    message.active = object.active ?? false;
    message.collateralType = object.collateralType ?? "";
    message.start = object.start ?? undefined;
    message.end = object.end ?? undefined;
    message.rewardsPerSecond = object.rewardsPerSecond !== undefined && object.rewardsPerSecond !== null ? Coin.fromPartial(object.rewardsPerSecond) : undefined;
    return message;
  }

};

function createBaseMultiRewardPeriod(): MultiRewardPeriod {
  return {
    active: false,
    collateralType: "",
    start: undefined,
    end: undefined,
    rewardsPerSecond: []
  };
}

export const MultiRewardPeriod = {
  encode(message: MultiRewardPeriod, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.active === true) {
      writer.uint32(8).bool(message.active);
    }

    if (message.collateralType !== "") {
      writer.uint32(18).string(message.collateralType);
    }

    if (message.start !== undefined) {
      Timestamp.encode(toTimestamp(message.start), writer.uint32(26).fork()).ldelim();
    }

    if (message.end !== undefined) {
      Timestamp.encode(toTimestamp(message.end), writer.uint32(34).fork()).ldelim();
    }

    for (const v of message.rewardsPerSecond) {
      Coin.encode(v!, writer.uint32(42).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MultiRewardPeriod {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMultiRewardPeriod();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.active = reader.bool();
          break;

        case 2:
          message.collateralType = reader.string();
          break;

        case 3:
          message.start = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;

        case 4:
          message.end = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;

        case 5:
          message.rewardsPerSecond.push(Coin.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MultiRewardPeriod>): MultiRewardPeriod {
    const message = createBaseMultiRewardPeriod();
    message.active = object.active ?? false;
    message.collateralType = object.collateralType ?? "";
    message.start = object.start ?? undefined;
    message.end = object.end ?? undefined;
    message.rewardsPerSecond = object.rewardsPerSecond?.map(e => Coin.fromPartial(e)) || [];
    return message;
  }

};

function createBaseMultiplier(): Multiplier {
  return {
    name: "",
    monthsLockup: Long.ZERO,
    factor: new Uint8Array()
  };
}

export const Multiplier = {
  encode(message: Multiplier, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }

    if (!message.monthsLockup.isZero()) {
      writer.uint32(16).int64(message.monthsLockup);
    }

    if (message.factor.length !== 0) {
      writer.uint32(26).bytes(message.factor);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Multiplier {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMultiplier();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;

        case 2:
          message.monthsLockup = (reader.int64() as Long);
          break;

        case 3:
          message.factor = reader.bytes();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<Multiplier>): Multiplier {
    const message = createBaseMultiplier();
    message.name = object.name ?? "";
    message.monthsLockup = object.monthsLockup !== undefined && object.monthsLockup !== null ? Long.fromValue(object.monthsLockup) : Long.ZERO;
    message.factor = object.factor ?? new Uint8Array();
    return message;
  }

};

function createBaseMultipliersPerDenom(): MultipliersPerDenom {
  return {
    denom: "",
    multipliers: []
  };
}

export const MultipliersPerDenom = {
  encode(message: MultipliersPerDenom, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }

    for (const v of message.multipliers) {
      Multiplier.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MultipliersPerDenom {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMultipliersPerDenom();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;

        case 2:
          message.multipliers.push(Multiplier.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MultipliersPerDenom>): MultipliersPerDenom {
    const message = createBaseMultipliersPerDenom();
    message.denom = object.denom ?? "";
    message.multipliers = object.multipliers?.map(e => Multiplier.fromPartial(e)) || [];
    return message;
  }

};

function createBaseParams(): Params {
  return {
    usdxMintingRewardPeriods: [],
    joltSupplyRewardPeriods: [],
    joltBorrowRewardPeriods: [],
    delegatorRewardPeriods: [],
    swapRewardPeriods: [],
    claimMultipliers: [],
    claimEnd: undefined,
    savingsRewardPeriods: []
  };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.usdxMintingRewardPeriods) {
      RewardPeriod.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    for (const v of message.joltSupplyRewardPeriods) {
      MultiRewardPeriod.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    for (const v of message.joltBorrowRewardPeriods) {
      MultiRewardPeriod.encode(v!, writer.uint32(26).fork()).ldelim();
    }

    for (const v of message.delegatorRewardPeriods) {
      MultiRewardPeriod.encode(v!, writer.uint32(34).fork()).ldelim();
    }

    for (const v of message.swapRewardPeriods) {
      MultiRewardPeriod.encode(v!, writer.uint32(42).fork()).ldelim();
    }

    for (const v of message.claimMultipliers) {
      MultipliersPerDenom.encode(v!, writer.uint32(50).fork()).ldelim();
    }

    if (message.claimEnd !== undefined) {
      Timestamp.encode(toTimestamp(message.claimEnd), writer.uint32(58).fork()).ldelim();
    }

    for (const v of message.savingsRewardPeriods) {
      MultiRewardPeriod.encode(v!, writer.uint32(66).fork()).ldelim();
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
          message.usdxMintingRewardPeriods.push(RewardPeriod.decode(reader, reader.uint32()));
          break;

        case 2:
          message.joltSupplyRewardPeriods.push(MultiRewardPeriod.decode(reader, reader.uint32()));
          break;

        case 3:
          message.joltBorrowRewardPeriods.push(MultiRewardPeriod.decode(reader, reader.uint32()));
          break;

        case 4:
          message.delegatorRewardPeriods.push(MultiRewardPeriod.decode(reader, reader.uint32()));
          break;

        case 5:
          message.swapRewardPeriods.push(MultiRewardPeriod.decode(reader, reader.uint32()));
          break;

        case 6:
          message.claimMultipliers.push(MultipliersPerDenom.decode(reader, reader.uint32()));
          break;

        case 7:
          message.claimEnd = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;

        case 8:
          message.savingsRewardPeriods.push(MultiRewardPeriod.decode(reader, reader.uint32()));
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
    message.usdxMintingRewardPeriods = object.usdxMintingRewardPeriods?.map(e => RewardPeriod.fromPartial(e)) || [];
    message.joltSupplyRewardPeriods = object.joltSupplyRewardPeriods?.map(e => MultiRewardPeriod.fromPartial(e)) || [];
    message.joltBorrowRewardPeriods = object.joltBorrowRewardPeriods?.map(e => MultiRewardPeriod.fromPartial(e)) || [];
    message.delegatorRewardPeriods = object.delegatorRewardPeriods?.map(e => MultiRewardPeriod.fromPartial(e)) || [];
    message.swapRewardPeriods = object.swapRewardPeriods?.map(e => MultiRewardPeriod.fromPartial(e)) || [];
    message.claimMultipliers = object.claimMultipliers?.map(e => MultipliersPerDenom.fromPartial(e)) || [];
    message.claimEnd = object.claimEnd ?? undefined;
    message.savingsRewardPeriods = object.savingsRewardPeriods?.map(e => MultiRewardPeriod.fromPartial(e)) || [];
    return message;
  }

};