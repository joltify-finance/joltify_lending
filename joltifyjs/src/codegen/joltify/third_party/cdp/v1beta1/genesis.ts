import { CDP, CDPSDKType, Deposit, DepositSDKType } from "./cdp";
import { Coin, CoinSDKType } from "../../../../cosmos/base/v1beta1/coin";
import { Timestamp } from "../../../../google/protobuf/timestamp";
import { Long, DeepPartial, toTimestamp, fromTimestamp } from "../../../../helpers";
import * as _m0 from "protobufjs/minimal";
/** GenesisState defines the cdp module's genesis state. */

export interface GenesisState {
  /** params defines all the paramaters of the module. */
  params?: Params;
  cdps: CDP[];
  deposits: Deposit[];
  startingCdpId: Long;
  debtDenom: string;
  govDenom: string;
  previousAccumulationTimes: GenesisAccumulationTime[];
  totalPrincipals: GenesisTotalPrincipal[];
}
/** GenesisState defines the cdp module's genesis state. */

export interface GenesisStateSDKType {
  params?: ParamsSDKType;
  cdps: CDPSDKType[];
  deposits: DepositSDKType[];
  starting_cdp_id: Long;
  debt_denom: string;
  gov_denom: string;
  previous_accumulation_times: GenesisAccumulationTimeSDKType[];
  total_principals: GenesisTotalPrincipalSDKType[];
}
/** Params defines the parameters for the cdp module. */

export interface Params {
  collateralParams: CollateralParam[];
  debtParam?: DebtParam;
  globalDebtLimit?: Coin;
  surplusAuctionThreshold: string;
  surplusAuctionLot: string;
  debtAuctionThreshold: string;
  debtAuctionLot: string;
  circuitBreaker: boolean;
}
/** Params defines the parameters for the cdp module. */

export interface ParamsSDKType {
  collateral_params: CollateralParamSDKType[];
  debt_param?: DebtParamSDKType;
  global_debt_limit?: CoinSDKType;
  surplus_auction_threshold: string;
  surplus_auction_lot: string;
  debt_auction_threshold: string;
  debt_auction_lot: string;
  circuit_breaker: boolean;
}
/** DebtParam defines governance params for debt assets */

export interface DebtParam {
  denom: string;
  referenceAsset: string;
  conversionFactor: string;
  debtFloor: string;
}
/** DebtParam defines governance params for debt assets */

export interface DebtParamSDKType {
  denom: string;
  reference_asset: string;
  conversion_factor: string;
  debt_floor: string;
}
/** CollateralParam defines governance parameters for each collateral type within the cdp module */

export interface CollateralParam {
  denom: string;
  type: string;
  liquidationRatio: string;
  debtLimit?: Coin;
  stabilityFee: string;
  auctionSize: string;
  liquidationPenalty: string;
  spotMarketId: string;
  liquidationMarketId: string;
  keeperRewardPercentage: string;
  checkCollateralizationIndexCount: string;
  conversionFactor: string;
}
/** CollateralParam defines governance parameters for each collateral type within the cdp module */

export interface CollateralParamSDKType {
  denom: string;
  type: string;
  liquidation_ratio: string;
  debt_limit?: CoinSDKType;
  stability_fee: string;
  auction_size: string;
  liquidation_penalty: string;
  spot_market_id: string;
  liquidation_market_id: string;
  keeper_reward_percentage: string;
  check_collateralization_index_count: string;
  conversion_factor: string;
}
/** GenesisAccumulationTime defines the previous distribution time and its corresponding denom */

export interface GenesisAccumulationTime {
  collateralType: string;
  previousAccumulationTime?: Date;
  interestFactor: string;
}
/** GenesisAccumulationTime defines the previous distribution time and its corresponding denom */

export interface GenesisAccumulationTimeSDKType {
  collateral_type: string;
  previous_accumulation_time?: Date;
  interest_factor: string;
}
/** GenesisTotalPrincipal defines the total principal and its corresponding collateral type */

export interface GenesisTotalPrincipal {
  collateralType: string;
  totalPrincipal: string;
}
/** GenesisTotalPrincipal defines the total principal and its corresponding collateral type */

export interface GenesisTotalPrincipalSDKType {
  collateral_type: string;
  total_principal: string;
}

function createBaseGenesisState(): GenesisState {
  return {
    params: undefined,
    cdps: [],
    deposits: [],
    startingCdpId: Long.UZERO,
    debtDenom: "",
    govDenom: "",
    previousAccumulationTimes: [],
    totalPrincipals: []
  };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }

    for (const v of message.cdps) {
      CDP.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    for (const v of message.deposits) {
      Deposit.encode(v!, writer.uint32(26).fork()).ldelim();
    }

    if (!message.startingCdpId.isZero()) {
      writer.uint32(32).uint64(message.startingCdpId);
    }

    if (message.debtDenom !== "") {
      writer.uint32(42).string(message.debtDenom);
    }

    if (message.govDenom !== "") {
      writer.uint32(50).string(message.govDenom);
    }

    for (const v of message.previousAccumulationTimes) {
      GenesisAccumulationTime.encode(v!, writer.uint32(58).fork()).ldelim();
    }

    for (const v of message.totalPrincipals) {
      GenesisTotalPrincipal.encode(v!, writer.uint32(66).fork()).ldelim();
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
          message.cdps.push(CDP.decode(reader, reader.uint32()));
          break;

        case 3:
          message.deposits.push(Deposit.decode(reader, reader.uint32()));
          break;

        case 4:
          message.startingCdpId = (reader.uint64() as Long);
          break;

        case 5:
          message.debtDenom = reader.string();
          break;

        case 6:
          message.govDenom = reader.string();
          break;

        case 7:
          message.previousAccumulationTimes.push(GenesisAccumulationTime.decode(reader, reader.uint32()));
          break;

        case 8:
          message.totalPrincipals.push(GenesisTotalPrincipal.decode(reader, reader.uint32()));
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
    message.cdps = object.cdps?.map(e => CDP.fromPartial(e)) || [];
    message.deposits = object.deposits?.map(e => Deposit.fromPartial(e)) || [];
    message.startingCdpId = object.startingCdpId !== undefined && object.startingCdpId !== null ? Long.fromValue(object.startingCdpId) : Long.UZERO;
    message.debtDenom = object.debtDenom ?? "";
    message.govDenom = object.govDenom ?? "";
    message.previousAccumulationTimes = object.previousAccumulationTimes?.map(e => GenesisAccumulationTime.fromPartial(e)) || [];
    message.totalPrincipals = object.totalPrincipals?.map(e => GenesisTotalPrincipal.fromPartial(e)) || [];
    return message;
  }

};

function createBaseParams(): Params {
  return {
    collateralParams: [],
    debtParam: undefined,
    globalDebtLimit: undefined,
    surplusAuctionThreshold: "",
    surplusAuctionLot: "",
    debtAuctionThreshold: "",
    debtAuctionLot: "",
    circuitBreaker: false
  };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.collateralParams) {
      CollateralParam.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    if (message.debtParam !== undefined) {
      DebtParam.encode(message.debtParam, writer.uint32(18).fork()).ldelim();
    }

    if (message.globalDebtLimit !== undefined) {
      Coin.encode(message.globalDebtLimit, writer.uint32(26).fork()).ldelim();
    }

    if (message.surplusAuctionThreshold !== "") {
      writer.uint32(34).string(message.surplusAuctionThreshold);
    }

    if (message.surplusAuctionLot !== "") {
      writer.uint32(42).string(message.surplusAuctionLot);
    }

    if (message.debtAuctionThreshold !== "") {
      writer.uint32(50).string(message.debtAuctionThreshold);
    }

    if (message.debtAuctionLot !== "") {
      writer.uint32(58).string(message.debtAuctionLot);
    }

    if (message.circuitBreaker === true) {
      writer.uint32(64).bool(message.circuitBreaker);
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
          message.collateralParams.push(CollateralParam.decode(reader, reader.uint32()));
          break;

        case 2:
          message.debtParam = DebtParam.decode(reader, reader.uint32());
          break;

        case 3:
          message.globalDebtLimit = Coin.decode(reader, reader.uint32());
          break;

        case 4:
          message.surplusAuctionThreshold = reader.string();
          break;

        case 5:
          message.surplusAuctionLot = reader.string();
          break;

        case 6:
          message.debtAuctionThreshold = reader.string();
          break;

        case 7:
          message.debtAuctionLot = reader.string();
          break;

        case 8:
          message.circuitBreaker = reader.bool();
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
    message.collateralParams = object.collateralParams?.map(e => CollateralParam.fromPartial(e)) || [];
    message.debtParam = object.debtParam !== undefined && object.debtParam !== null ? DebtParam.fromPartial(object.debtParam) : undefined;
    message.globalDebtLimit = object.globalDebtLimit !== undefined && object.globalDebtLimit !== null ? Coin.fromPartial(object.globalDebtLimit) : undefined;
    message.surplusAuctionThreshold = object.surplusAuctionThreshold ?? "";
    message.surplusAuctionLot = object.surplusAuctionLot ?? "";
    message.debtAuctionThreshold = object.debtAuctionThreshold ?? "";
    message.debtAuctionLot = object.debtAuctionLot ?? "";
    message.circuitBreaker = object.circuitBreaker ?? false;
    return message;
  }

};

function createBaseDebtParam(): DebtParam {
  return {
    denom: "",
    referenceAsset: "",
    conversionFactor: "",
    debtFloor: ""
  };
}

export const DebtParam = {
  encode(message: DebtParam, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }

    if (message.referenceAsset !== "") {
      writer.uint32(18).string(message.referenceAsset);
    }

    if (message.conversionFactor !== "") {
      writer.uint32(26).string(message.conversionFactor);
    }

    if (message.debtFloor !== "") {
      writer.uint32(34).string(message.debtFloor);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DebtParam {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDebtParam();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;

        case 2:
          message.referenceAsset = reader.string();
          break;

        case 3:
          message.conversionFactor = reader.string();
          break;

        case 4:
          message.debtFloor = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<DebtParam>): DebtParam {
    const message = createBaseDebtParam();
    message.denom = object.denom ?? "";
    message.referenceAsset = object.referenceAsset ?? "";
    message.conversionFactor = object.conversionFactor ?? "";
    message.debtFloor = object.debtFloor ?? "";
    return message;
  }

};

function createBaseCollateralParam(): CollateralParam {
  return {
    denom: "",
    type: "",
    liquidationRatio: "",
    debtLimit: undefined,
    stabilityFee: "",
    auctionSize: "",
    liquidationPenalty: "",
    spotMarketId: "",
    liquidationMarketId: "",
    keeperRewardPercentage: "",
    checkCollateralizationIndexCount: "",
    conversionFactor: ""
  };
}

export const CollateralParam = {
  encode(message: CollateralParam, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }

    if (message.type !== "") {
      writer.uint32(18).string(message.type);
    }

    if (message.liquidationRatio !== "") {
      writer.uint32(26).string(message.liquidationRatio);
    }

    if (message.debtLimit !== undefined) {
      Coin.encode(message.debtLimit, writer.uint32(34).fork()).ldelim();
    }

    if (message.stabilityFee !== "") {
      writer.uint32(42).string(message.stabilityFee);
    }

    if (message.auctionSize !== "") {
      writer.uint32(50).string(message.auctionSize);
    }

    if (message.liquidationPenalty !== "") {
      writer.uint32(58).string(message.liquidationPenalty);
    }

    if (message.spotMarketId !== "") {
      writer.uint32(66).string(message.spotMarketId);
    }

    if (message.liquidationMarketId !== "") {
      writer.uint32(74).string(message.liquidationMarketId);
    }

    if (message.keeperRewardPercentage !== "") {
      writer.uint32(82).string(message.keeperRewardPercentage);
    }

    if (message.checkCollateralizationIndexCount !== "") {
      writer.uint32(90).string(message.checkCollateralizationIndexCount);
    }

    if (message.conversionFactor !== "") {
      writer.uint32(98).string(message.conversionFactor);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CollateralParam {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCollateralParam();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;

        case 2:
          message.type = reader.string();
          break;

        case 3:
          message.liquidationRatio = reader.string();
          break;

        case 4:
          message.debtLimit = Coin.decode(reader, reader.uint32());
          break;

        case 5:
          message.stabilityFee = reader.string();
          break;

        case 6:
          message.auctionSize = reader.string();
          break;

        case 7:
          message.liquidationPenalty = reader.string();
          break;

        case 8:
          message.spotMarketId = reader.string();
          break;

        case 9:
          message.liquidationMarketId = reader.string();
          break;

        case 10:
          message.keeperRewardPercentage = reader.string();
          break;

        case 11:
          message.checkCollateralizationIndexCount = reader.string();
          break;

        case 12:
          message.conversionFactor = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<CollateralParam>): CollateralParam {
    const message = createBaseCollateralParam();
    message.denom = object.denom ?? "";
    message.type = object.type ?? "";
    message.liquidationRatio = object.liquidationRatio ?? "";
    message.debtLimit = object.debtLimit !== undefined && object.debtLimit !== null ? Coin.fromPartial(object.debtLimit) : undefined;
    message.stabilityFee = object.stabilityFee ?? "";
    message.auctionSize = object.auctionSize ?? "";
    message.liquidationPenalty = object.liquidationPenalty ?? "";
    message.spotMarketId = object.spotMarketId ?? "";
    message.liquidationMarketId = object.liquidationMarketId ?? "";
    message.keeperRewardPercentage = object.keeperRewardPercentage ?? "";
    message.checkCollateralizationIndexCount = object.checkCollateralizationIndexCount ?? "";
    message.conversionFactor = object.conversionFactor ?? "";
    return message;
  }

};

function createBaseGenesisAccumulationTime(): GenesisAccumulationTime {
  return {
    collateralType: "",
    previousAccumulationTime: undefined,
    interestFactor: ""
  };
}

export const GenesisAccumulationTime = {
  encode(message: GenesisAccumulationTime, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.collateralType !== "") {
      writer.uint32(10).string(message.collateralType);
    }

    if (message.previousAccumulationTime !== undefined) {
      Timestamp.encode(toTimestamp(message.previousAccumulationTime), writer.uint32(18).fork()).ldelim();
    }

    if (message.interestFactor !== "") {
      writer.uint32(26).string(message.interestFactor);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisAccumulationTime {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisAccumulationTime();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.collateralType = reader.string();
          break;

        case 2:
          message.previousAccumulationTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;

        case 3:
          message.interestFactor = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<GenesisAccumulationTime>): GenesisAccumulationTime {
    const message = createBaseGenesisAccumulationTime();
    message.collateralType = object.collateralType ?? "";
    message.previousAccumulationTime = object.previousAccumulationTime ?? undefined;
    message.interestFactor = object.interestFactor ?? "";
    return message;
  }

};

function createBaseGenesisTotalPrincipal(): GenesisTotalPrincipal {
  return {
    collateralType: "",
    totalPrincipal: ""
  };
}

export const GenesisTotalPrincipal = {
  encode(message: GenesisTotalPrincipal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.collateralType !== "") {
      writer.uint32(10).string(message.collateralType);
    }

    if (message.totalPrincipal !== "") {
      writer.uint32(18).string(message.totalPrincipal);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisTotalPrincipal {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisTotalPrincipal();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.collateralType = reader.string();
          break;

        case 2:
          message.totalPrincipal = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<GenesisTotalPrincipal>): GenesisTotalPrincipal {
    const message = createBaseGenesisTotalPrincipal();
    message.collateralType = object.collateralType ?? "";
    message.totalPrincipal = object.totalPrincipal ?? "";
    return message;
  }

};