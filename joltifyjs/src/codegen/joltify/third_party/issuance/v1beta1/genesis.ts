import { Duration, DurationSDKType } from "../../../../google/protobuf/duration";
import { Coin, CoinSDKType } from "../../../../cosmos/base/v1beta1/coin";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../../../helpers";
/** GenesisState defines the issuance module's genesis state. */

export interface GenesisState {
  /** params defines all the paramaters of the module. */
  params?: Params;
  supplies: AssetSupply[];
}
/** GenesisState defines the issuance module's genesis state. */

export interface GenesisStateSDKType {
  params?: ParamsSDKType;
  supplies: AssetSupplySDKType[];
}
/** Params defines the parameters for the issuance module. */

export interface Params {
  assets: Asset[];
}
/** Params defines the parameters for the issuance module. */

export interface ParamsSDKType {
  assets: AssetSDKType[];
}
/** Asset type for assets in the issuance module */

export interface Asset {
  owner: string;
  denom: string;
  blockedAddresses: string[];
  paused: boolean;
  blockable: boolean;
  rateLimit?: RateLimit;
}
/** Asset type for assets in the issuance module */

export interface AssetSDKType {
  owner: string;
  denom: string;
  blocked_addresses: string[];
  paused: boolean;
  blockable: boolean;
  rate_limit?: RateLimitSDKType;
}
/** RateLimit parameters for rate-limiting the supply of an issued asset */

export interface RateLimit {
  active: boolean;
  limit: Uint8Array;
  timePeriod?: Duration;
}
/** RateLimit parameters for rate-limiting the supply of an issued asset */

export interface RateLimitSDKType {
  active: boolean;
  limit: Uint8Array;
  time_period?: DurationSDKType;
}
/**
 * AssetSupply contains information about an asset's rate-limited supply (the
 * total supply of the asset is tracked in the top-level supply module)
 */

export interface AssetSupply {
  currentSupply?: Coin;
  timeElapsed?: Duration;
}
/**
 * AssetSupply contains information about an asset's rate-limited supply (the
 * total supply of the asset is tracked in the top-level supply module)
 */

export interface AssetSupplySDKType {
  current_supply?: CoinSDKType;
  time_elapsed?: DurationSDKType;
}

function createBaseGenesisState(): GenesisState {
  return {
    params: undefined,
    supplies: []
  };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }

    for (const v of message.supplies) {
      AssetSupply.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;

        case 2:
          message.supplies.push(AssetSupply.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = createBaseGenesisState();
    message.params = object.params !== undefined && object.params !== null ? Params.fromPartial(object.params) : undefined;
    message.supplies = object.supplies?.map(e => AssetSupply.fromPartial(e)) || [];
    return message;
  }

};

function createBaseParams(): Params {
  return {
    assets: []
  };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.assets) {
      Asset.encode(v!, writer.uint32(10).fork()).ldelim();
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
          message.assets.push(Asset.decode(reader, reader.uint32()));
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
    message.assets = object.assets?.map(e => Asset.fromPartial(e)) || [];
    return message;
  }

};

function createBaseAsset(): Asset {
  return {
    owner: "",
    denom: "",
    blockedAddresses: [],
    paused: false,
    blockable: false,
    rateLimit: undefined
  };
}

export const Asset = {
  encode(message: Asset, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }

    if (message.denom !== "") {
      writer.uint32(18).string(message.denom);
    }

    for (const v of message.blockedAddresses) {
      writer.uint32(26).string(v!);
    }

    if (message.paused === true) {
      writer.uint32(32).bool(message.paused);
    }

    if (message.blockable === true) {
      writer.uint32(40).bool(message.blockable);
    }

    if (message.rateLimit !== undefined) {
      RateLimit.encode(message.rateLimit, writer.uint32(50).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Asset {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAsset();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;

        case 2:
          message.denom = reader.string();
          break;

        case 3:
          message.blockedAddresses.push(reader.string());
          break;

        case 4:
          message.paused = reader.bool();
          break;

        case 5:
          message.blockable = reader.bool();
          break;

        case 6:
          message.rateLimit = RateLimit.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<Asset>): Asset {
    const message = createBaseAsset();
    message.owner = object.owner ?? "";
    message.denom = object.denom ?? "";
    message.blockedAddresses = object.blockedAddresses?.map(e => e) || [];
    message.paused = object.paused ?? false;
    message.blockable = object.blockable ?? false;
    message.rateLimit = object.rateLimit !== undefined && object.rateLimit !== null ? RateLimit.fromPartial(object.rateLimit) : undefined;
    return message;
  }

};

function createBaseRateLimit(): RateLimit {
  return {
    active: false,
    limit: new Uint8Array(),
    timePeriod: undefined
  };
}

export const RateLimit = {
  encode(message: RateLimit, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.active === true) {
      writer.uint32(8).bool(message.active);
    }

    if (message.limit.length !== 0) {
      writer.uint32(18).bytes(message.limit);
    }

    if (message.timePeriod !== undefined) {
      Duration.encode(message.timePeriod, writer.uint32(26).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RateLimit {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRateLimit();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.active = reader.bool();
          break;

        case 2:
          message.limit = reader.bytes();
          break;

        case 3:
          message.timePeriod = Duration.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<RateLimit>): RateLimit {
    const message = createBaseRateLimit();
    message.active = object.active ?? false;
    message.limit = object.limit ?? new Uint8Array();
    message.timePeriod = object.timePeriod !== undefined && object.timePeriod !== null ? Duration.fromPartial(object.timePeriod) : undefined;
    return message;
  }

};

function createBaseAssetSupply(): AssetSupply {
  return {
    currentSupply: undefined,
    timeElapsed: undefined
  };
}

export const AssetSupply = {
  encode(message: AssetSupply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.currentSupply !== undefined) {
      Coin.encode(message.currentSupply, writer.uint32(10).fork()).ldelim();
    }

    if (message.timeElapsed !== undefined) {
      Duration.encode(message.timeElapsed, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AssetSupply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAssetSupply();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.currentSupply = Coin.decode(reader, reader.uint32());
          break;

        case 2:
          message.timeElapsed = Duration.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<AssetSupply>): AssetSupply {
    const message = createBaseAssetSupply();
    message.currentSupply = object.currentSupply !== undefined && object.currentSupply !== null ? Coin.fromPartial(object.currentSupply) : undefined;
    message.timeElapsed = object.timeElapsed !== undefined && object.timeElapsed !== null ? Duration.fromPartial(object.timeElapsed) : undefined;
    return message;
  }

};