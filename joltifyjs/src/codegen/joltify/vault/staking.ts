import { Coin, CoinSDKType } from "../../cosmos/base/v1beta1/coin";
import { Long, DeepPartial } from "../../helpers";
import * as _m0 from "protobufjs/minimal";
export interface Params {
  blockChurnInterval: Long;
  power: Long;
  step: Long;
  candidateRatio: string;
  targetQuota: Coin[];
  historyLength: number;
}
export interface ParamsSDKType {
  block_churn_interval: Long;
  power: Long;
  step: Long;
  candidate_ratio: string;
  target_quota: CoinSDKType[];
  history_length: number;
}
export interface Validator {
  pubkey: Uint8Array;
  power: Long;
}
export interface ValidatorSDKType {
  pubkey: Uint8Array;
  power: Long;
}
export interface StandbyPower {
  addr: string;
  power: Long;
}
export interface StandbyPowerSDKType {
  addr: string;
  power: Long;
}
export interface Validators {
  allValidators: Validator[];
  height: Long;
}
export interface ValidatorsSDKType {
  all_validators: ValidatorSDKType[];
  height: Long;
}

function createBaseParams(): Params {
  return {
    blockChurnInterval: Long.ZERO,
    power: Long.ZERO,
    step: Long.ZERO,
    candidateRatio: "",
    targetQuota: [],
    historyLength: 0
  };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (!message.blockChurnInterval.isZero()) {
      writer.uint32(8).int64(message.blockChurnInterval);
    }

    if (!message.power.isZero()) {
      writer.uint32(16).int64(message.power);
    }

    if (!message.step.isZero()) {
      writer.uint32(24).int64(message.step);
    }

    if (message.candidateRatio !== "") {
      writer.uint32(34).string(message.candidateRatio);
    }

    for (const v of message.targetQuota) {
      Coin.encode(v!, writer.uint32(42).fork()).ldelim();
    }

    if (message.historyLength !== 0) {
      writer.uint32(48).int32(message.historyLength);
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
          message.blockChurnInterval = (reader.int64() as Long);
          break;

        case 2:
          message.power = (reader.int64() as Long);
          break;

        case 3:
          message.step = (reader.int64() as Long);
          break;

        case 4:
          message.candidateRatio = reader.string();
          break;

        case 5:
          message.targetQuota.push(Coin.decode(reader, reader.uint32()));
          break;

        case 6:
          message.historyLength = reader.int32();
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
    message.blockChurnInterval = object.blockChurnInterval !== undefined && object.blockChurnInterval !== null ? Long.fromValue(object.blockChurnInterval) : Long.ZERO;
    message.power = object.power !== undefined && object.power !== null ? Long.fromValue(object.power) : Long.ZERO;
    message.step = object.step !== undefined && object.step !== null ? Long.fromValue(object.step) : Long.ZERO;
    message.candidateRatio = object.candidateRatio ?? "";
    message.targetQuota = object.targetQuota?.map(e => Coin.fromPartial(e)) || [];
    message.historyLength = object.historyLength ?? 0;
    return message;
  }

};

function createBaseValidator(): Validator {
  return {
    pubkey: new Uint8Array(),
    power: Long.ZERO
  };
}

export const Validator = {
  encode(message: Validator, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pubkey.length !== 0) {
      writer.uint32(10).bytes(message.pubkey);
    }

    if (!message.power.isZero()) {
      writer.uint32(16).int64(message.power);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Validator {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseValidator();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.pubkey = reader.bytes();
          break;

        case 2:
          message.power = (reader.int64() as Long);
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<Validator>): Validator {
    const message = createBaseValidator();
    message.pubkey = object.pubkey ?? new Uint8Array();
    message.power = object.power !== undefined && object.power !== null ? Long.fromValue(object.power) : Long.ZERO;
    return message;
  }

};

function createBaseStandbyPower(): StandbyPower {
  return {
    addr: "",
    power: Long.ZERO
  };
}

export const StandbyPower = {
  encode(message: StandbyPower, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.addr !== "") {
      writer.uint32(10).string(message.addr);
    }

    if (!message.power.isZero()) {
      writer.uint32(16).int64(message.power);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): StandbyPower {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStandbyPower();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.addr = reader.string();
          break;

        case 2:
          message.power = (reader.int64() as Long);
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<StandbyPower>): StandbyPower {
    const message = createBaseStandbyPower();
    message.addr = object.addr ?? "";
    message.power = object.power !== undefined && object.power !== null ? Long.fromValue(object.power) : Long.ZERO;
    return message;
  }

};

function createBaseValidators(): Validators {
  return {
    allValidators: [],
    height: Long.ZERO
  };
}

export const Validators = {
  encode(message: Validators, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.allValidators) {
      Validator.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    if (!message.height.isZero()) {
      writer.uint32(16).int64(message.height);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Validators {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseValidators();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.allValidators.push(Validator.decode(reader, reader.uint32()));
          break;

        case 2:
          message.height = (reader.int64() as Long);
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<Validators>): Validators {
    const message = createBaseValidators();
    message.allValidators = object.allValidators?.map(e => Validator.fromPartial(e)) || [];
    message.height = object.height !== undefined && object.height !== null ? Long.fromValue(object.height) : Long.ZERO;
    return message;
  }

};