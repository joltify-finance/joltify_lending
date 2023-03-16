import { Timestamp } from "../../../../google/protobuf/timestamp";
import { MultiRewardIndex, MultiRewardIndexSDKType, USDXMintingClaim, USDXMintingClaimSDKType, JoltLiquidityProviderClaim, JoltLiquidityProviderClaimSDKType, DelegatorClaim, DelegatorClaimSDKType, SwapClaim, SwapClaimSDKType, SavingsClaim, SavingsClaimSDKType } from "./claims";
import { Params, ParamsSDKType } from "./params";
import * as _m0 from "protobufjs/minimal";
import { toTimestamp, fromTimestamp, DeepPartial } from "../../../../helpers";
/** AccumulationTime stores the previous reward distribution time and its corresponding collateral type */

export interface AccumulationTime {
  collateralType: string;
  previousAccumulationTime?: Date;
}
/** AccumulationTime stores the previous reward distribution time and its corresponding collateral type */

export interface AccumulationTimeSDKType {
  collateral_type: string;
  previous_accumulation_time?: Date;
}
/** GenesisRewardState groups together the global state for a particular reward so it can be exported in genesis. */

export interface GenesisRewardState {
  accumulationTimes: AccumulationTime[];
  multiRewardIndexes: MultiRewardIndex[];
}
/** GenesisRewardState groups together the global state for a particular reward so it can be exported in genesis. */

export interface GenesisRewardStateSDKType {
  accumulation_times: AccumulationTimeSDKType[];
  multi_reward_indexes: MultiRewardIndexSDKType[];
}
/** GenesisState is the state that must be provided at genesis. */

export interface GenesisState {
  params?: Params;
  usdxRewardState?: GenesisRewardState;
  joltSupplyRewardState?: GenesisRewardState;
  joltBorrowRewardState?: GenesisRewardState;
  delegatorRewardState?: GenesisRewardState;
  swapRewardState?: GenesisRewardState;
  usdxMintingClaims: USDXMintingClaim[];
  joltLiquidityProviderClaims: JoltLiquidityProviderClaim[];
  delegatorClaims: DelegatorClaim[];
  swapClaims: SwapClaim[];
  savingsRewardState?: GenesisRewardState;
  savingsClaims: SavingsClaim[];
}
/** GenesisState is the state that must be provided at genesis. */

export interface GenesisStateSDKType {
  params?: ParamsSDKType;
  usdx_reward_state?: GenesisRewardStateSDKType;
  jolt_supply_reward_state?: GenesisRewardStateSDKType;
  jolt_borrow_reward_state?: GenesisRewardStateSDKType;
  delegator_reward_state?: GenesisRewardStateSDKType;
  swap_reward_state?: GenesisRewardStateSDKType;
  usdx_minting_claims: USDXMintingClaimSDKType[];
  jolt_liquidity_provider_claims: JoltLiquidityProviderClaimSDKType[];
  delegator_claims: DelegatorClaimSDKType[];
  swap_claims: SwapClaimSDKType[];
  savings_reward_state?: GenesisRewardStateSDKType;
  savings_claims: SavingsClaimSDKType[];
}

function createBaseAccumulationTime(): AccumulationTime {
  return {
    collateralType: "",
    previousAccumulationTime: undefined
  };
}

export const AccumulationTime = {
  encode(message: AccumulationTime, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.collateralType !== "") {
      writer.uint32(10).string(message.collateralType);
    }

    if (message.previousAccumulationTime !== undefined) {
      Timestamp.encode(toTimestamp(message.previousAccumulationTime), writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AccumulationTime {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAccumulationTime();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.collateralType = reader.string();
          break;

        case 2:
          message.previousAccumulationTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<AccumulationTime>): AccumulationTime {
    const message = createBaseAccumulationTime();
    message.collateralType = object.collateralType ?? "";
    message.previousAccumulationTime = object.previousAccumulationTime ?? undefined;
    return message;
  }

};

function createBaseGenesisRewardState(): GenesisRewardState {
  return {
    accumulationTimes: [],
    multiRewardIndexes: []
  };
}

export const GenesisRewardState = {
  encode(message: GenesisRewardState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.accumulationTimes) {
      AccumulationTime.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    for (const v of message.multiRewardIndexes) {
      MultiRewardIndex.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisRewardState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisRewardState();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.accumulationTimes.push(AccumulationTime.decode(reader, reader.uint32()));
          break;

        case 2:
          message.multiRewardIndexes.push(MultiRewardIndex.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<GenesisRewardState>): GenesisRewardState {
    const message = createBaseGenesisRewardState();
    message.accumulationTimes = object.accumulationTimes?.map(e => AccumulationTime.fromPartial(e)) || [];
    message.multiRewardIndexes = object.multiRewardIndexes?.map(e => MultiRewardIndex.fromPartial(e)) || [];
    return message;
  }

};

function createBaseGenesisState(): GenesisState {
  return {
    params: undefined,
    usdxRewardState: undefined,
    joltSupplyRewardState: undefined,
    joltBorrowRewardState: undefined,
    delegatorRewardState: undefined,
    swapRewardState: undefined,
    usdxMintingClaims: [],
    joltLiquidityProviderClaims: [],
    delegatorClaims: [],
    swapClaims: [],
    savingsRewardState: undefined,
    savingsClaims: []
  };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }

    if (message.usdxRewardState !== undefined) {
      GenesisRewardState.encode(message.usdxRewardState, writer.uint32(18).fork()).ldelim();
    }

    if (message.joltSupplyRewardState !== undefined) {
      GenesisRewardState.encode(message.joltSupplyRewardState, writer.uint32(26).fork()).ldelim();
    }

    if (message.joltBorrowRewardState !== undefined) {
      GenesisRewardState.encode(message.joltBorrowRewardState, writer.uint32(34).fork()).ldelim();
    }

    if (message.delegatorRewardState !== undefined) {
      GenesisRewardState.encode(message.delegatorRewardState, writer.uint32(42).fork()).ldelim();
    }

    if (message.swapRewardState !== undefined) {
      GenesisRewardState.encode(message.swapRewardState, writer.uint32(50).fork()).ldelim();
    }

    for (const v of message.usdxMintingClaims) {
      USDXMintingClaim.encode(v!, writer.uint32(58).fork()).ldelim();
    }

    for (const v of message.joltLiquidityProviderClaims) {
      JoltLiquidityProviderClaim.encode(v!, writer.uint32(66).fork()).ldelim();
    }

    for (const v of message.delegatorClaims) {
      DelegatorClaim.encode(v!, writer.uint32(74).fork()).ldelim();
    }

    for (const v of message.swapClaims) {
      SwapClaim.encode(v!, writer.uint32(82).fork()).ldelim();
    }

    if (message.savingsRewardState !== undefined) {
      GenesisRewardState.encode(message.savingsRewardState, writer.uint32(90).fork()).ldelim();
    }

    for (const v of message.savingsClaims) {
      SavingsClaim.encode(v!, writer.uint32(98).fork()).ldelim();
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
          message.usdxRewardState = GenesisRewardState.decode(reader, reader.uint32());
          break;

        case 3:
          message.joltSupplyRewardState = GenesisRewardState.decode(reader, reader.uint32());
          break;

        case 4:
          message.joltBorrowRewardState = GenesisRewardState.decode(reader, reader.uint32());
          break;

        case 5:
          message.delegatorRewardState = GenesisRewardState.decode(reader, reader.uint32());
          break;

        case 6:
          message.swapRewardState = GenesisRewardState.decode(reader, reader.uint32());
          break;

        case 7:
          message.usdxMintingClaims.push(USDXMintingClaim.decode(reader, reader.uint32()));
          break;

        case 8:
          message.joltLiquidityProviderClaims.push(JoltLiquidityProviderClaim.decode(reader, reader.uint32()));
          break;

        case 9:
          message.delegatorClaims.push(DelegatorClaim.decode(reader, reader.uint32()));
          break;

        case 10:
          message.swapClaims.push(SwapClaim.decode(reader, reader.uint32()));
          break;

        case 11:
          message.savingsRewardState = GenesisRewardState.decode(reader, reader.uint32());
          break;

        case 12:
          message.savingsClaims.push(SavingsClaim.decode(reader, reader.uint32()));
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
    message.usdxRewardState = object.usdxRewardState !== undefined && object.usdxRewardState !== null ? GenesisRewardState.fromPartial(object.usdxRewardState) : undefined;
    message.joltSupplyRewardState = object.joltSupplyRewardState !== undefined && object.joltSupplyRewardState !== null ? GenesisRewardState.fromPartial(object.joltSupplyRewardState) : undefined;
    message.joltBorrowRewardState = object.joltBorrowRewardState !== undefined && object.joltBorrowRewardState !== null ? GenesisRewardState.fromPartial(object.joltBorrowRewardState) : undefined;
    message.delegatorRewardState = object.delegatorRewardState !== undefined && object.delegatorRewardState !== null ? GenesisRewardState.fromPartial(object.delegatorRewardState) : undefined;
    message.swapRewardState = object.swapRewardState !== undefined && object.swapRewardState !== null ? GenesisRewardState.fromPartial(object.swapRewardState) : undefined;
    message.usdxMintingClaims = object.usdxMintingClaims?.map(e => USDXMintingClaim.fromPartial(e)) || [];
    message.joltLiquidityProviderClaims = object.joltLiquidityProviderClaims?.map(e => JoltLiquidityProviderClaim.fromPartial(e)) || [];
    message.delegatorClaims = object.delegatorClaims?.map(e => DelegatorClaim.fromPartial(e)) || [];
    message.swapClaims = object.swapClaims?.map(e => SwapClaim.fromPartial(e)) || [];
    message.savingsRewardState = object.savingsRewardState !== undefined && object.savingsRewardState !== null ? GenesisRewardState.fromPartial(object.savingsRewardState) : undefined;
    message.savingsClaims = object.savingsClaims?.map(e => SavingsClaim.fromPartial(e)) || [];
    return message;
  }

};