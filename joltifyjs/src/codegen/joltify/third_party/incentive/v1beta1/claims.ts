import { Coin, CoinSDKType } from "../../../../cosmos/base/v1beta1/coin";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../../../helpers";
/** BaseClaim is a claim with a single reward coin types */

export interface BaseClaim {
  owner: Uint8Array;
  reward?: Coin;
}
/** BaseClaim is a claim with a single reward coin types */

export interface BaseClaimSDKType {
  owner: Uint8Array;
  reward?: CoinSDKType;
}
/** BaseMultiClaim is a claim with multiple reward coin types */

export interface BaseMultiClaim {
  owner: Uint8Array;
  reward: Coin[];
}
/** BaseMultiClaim is a claim with multiple reward coin types */

export interface BaseMultiClaimSDKType {
  owner: Uint8Array;
  reward: CoinSDKType[];
}
/** RewardIndex stores reward accumulation information */

export interface RewardIndex {
  collateralType: string;
  rewardFactor: Uint8Array;
}
/** RewardIndex stores reward accumulation information */

export interface RewardIndexSDKType {
  collateral_type: string;
  reward_factor: Uint8Array;
}
/** RewardIndexesProto defines a Protobuf wrapper around a RewardIndexes slice */

export interface RewardIndexesProto {
  rewardIndexes: RewardIndex[];
}
/** RewardIndexesProto defines a Protobuf wrapper around a RewardIndexes slice */

export interface RewardIndexesProtoSDKType {
  reward_indexes: RewardIndexSDKType[];
}
/** MultiRewardIndex stores reward accumulation information on multiple reward types */

export interface MultiRewardIndex {
  collateralType: string;
  rewardIndexes: RewardIndex[];
}
/** MultiRewardIndex stores reward accumulation information on multiple reward types */

export interface MultiRewardIndexSDKType {
  collateral_type: string;
  reward_indexes: RewardIndexSDKType[];
}
/** MultiRewardIndexesProto defines a Protobuf wrapper around a MultiRewardIndexes slice */

export interface MultiRewardIndexesProto {
  multiRewardIndexes: MultiRewardIndex[];
}
/** MultiRewardIndexesProto defines a Protobuf wrapper around a MultiRewardIndexes slice */

export interface MultiRewardIndexesProtoSDKType {
  multi_reward_indexes: MultiRewardIndexSDKType[];
}
/** USDXMintingClaim is for USDX minting rewards */

export interface USDXMintingClaim {
  baseClaim?: BaseClaim;
  rewardIndexes: RewardIndex[];
}
/** USDXMintingClaim is for USDX minting rewards */

export interface USDXMintingClaimSDKType {
  base_claim?: BaseClaimSDKType;
  reward_indexes: RewardIndexSDKType[];
}
/** JoltLiquidityProviderClaim stores the jolt liquidity provider rewards that can be claimed by owner */

export interface JoltLiquidityProviderClaim {
  baseClaim?: BaseMultiClaim;
  supplyRewardIndexes: MultiRewardIndex[];
  borrowRewardIndexes: MultiRewardIndex[];
}
/** JoltLiquidityProviderClaim stores the jolt liquidity provider rewards that can be claimed by owner */

export interface JoltLiquidityProviderClaimSDKType {
  base_claim?: BaseMultiClaimSDKType;
  supply_reward_indexes: MultiRewardIndexSDKType[];
  borrow_reward_indexes: MultiRewardIndexSDKType[];
}
/** DelegatorClaim stores delegation rewards that can be claimed by owner */

export interface DelegatorClaim {
  baseClaim?: BaseMultiClaim;
  rewardIndexes: MultiRewardIndex[];
}
/** DelegatorClaim stores delegation rewards that can be claimed by owner */

export interface DelegatorClaimSDKType {
  base_claim?: BaseMultiClaimSDKType;
  reward_indexes: MultiRewardIndexSDKType[];
}
/** SwapClaim stores the swap rewards that can be claimed by owner */

export interface SwapClaim {
  baseClaim?: BaseMultiClaim;
  rewardIndexes: MultiRewardIndex[];
}
/** SwapClaim stores the swap rewards that can be claimed by owner */

export interface SwapClaimSDKType {
  base_claim?: BaseMultiClaimSDKType;
  reward_indexes: MultiRewardIndexSDKType[];
}
/** SavingsClaim stores the savings rewards that can be claimed by owner */

export interface SavingsClaim {
  baseClaim?: BaseMultiClaim;
  rewardIndexes: MultiRewardIndex[];
}
/** SavingsClaim stores the savings rewards that can be claimed by owner */

export interface SavingsClaimSDKType {
  base_claim?: BaseMultiClaimSDKType;
  reward_indexes: MultiRewardIndexSDKType[];
}

function createBaseBaseClaim(): BaseClaim {
  return {
    owner: new Uint8Array(),
    reward: undefined
  };
}

export const BaseClaim = {
  encode(message: BaseClaim, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner.length !== 0) {
      writer.uint32(10).bytes(message.owner);
    }

    if (message.reward !== undefined) {
      Coin.encode(message.reward, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BaseClaim {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBaseClaim();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.owner = reader.bytes();
          break;

        case 2:
          message.reward = Coin.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<BaseClaim>): BaseClaim {
    const message = createBaseBaseClaim();
    message.owner = object.owner ?? new Uint8Array();
    message.reward = object.reward !== undefined && object.reward !== null ? Coin.fromPartial(object.reward) : undefined;
    return message;
  }

};

function createBaseBaseMultiClaim(): BaseMultiClaim {
  return {
    owner: new Uint8Array(),
    reward: []
  };
}

export const BaseMultiClaim = {
  encode(message: BaseMultiClaim, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner.length !== 0) {
      writer.uint32(10).bytes(message.owner);
    }

    for (const v of message.reward) {
      Coin.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BaseMultiClaim {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBaseMultiClaim();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.owner = reader.bytes();
          break;

        case 2:
          message.reward.push(Coin.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<BaseMultiClaim>): BaseMultiClaim {
    const message = createBaseBaseMultiClaim();
    message.owner = object.owner ?? new Uint8Array();
    message.reward = object.reward?.map(e => Coin.fromPartial(e)) || [];
    return message;
  }

};

function createBaseRewardIndex(): RewardIndex {
  return {
    collateralType: "",
    rewardFactor: new Uint8Array()
  };
}

export const RewardIndex = {
  encode(message: RewardIndex, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.collateralType !== "") {
      writer.uint32(10).string(message.collateralType);
    }

    if (message.rewardFactor.length !== 0) {
      writer.uint32(18).bytes(message.rewardFactor);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RewardIndex {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRewardIndex();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.collateralType = reader.string();
          break;

        case 2:
          message.rewardFactor = reader.bytes();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<RewardIndex>): RewardIndex {
    const message = createBaseRewardIndex();
    message.collateralType = object.collateralType ?? "";
    message.rewardFactor = object.rewardFactor ?? new Uint8Array();
    return message;
  }

};

function createBaseRewardIndexesProto(): RewardIndexesProto {
  return {
    rewardIndexes: []
  };
}

export const RewardIndexesProto = {
  encode(message: RewardIndexesProto, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.rewardIndexes) {
      RewardIndex.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RewardIndexesProto {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRewardIndexesProto();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.rewardIndexes.push(RewardIndex.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<RewardIndexesProto>): RewardIndexesProto {
    const message = createBaseRewardIndexesProto();
    message.rewardIndexes = object.rewardIndexes?.map(e => RewardIndex.fromPartial(e)) || [];
    return message;
  }

};

function createBaseMultiRewardIndex(): MultiRewardIndex {
  return {
    collateralType: "",
    rewardIndexes: []
  };
}

export const MultiRewardIndex = {
  encode(message: MultiRewardIndex, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.collateralType !== "") {
      writer.uint32(10).string(message.collateralType);
    }

    for (const v of message.rewardIndexes) {
      RewardIndex.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MultiRewardIndex {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMultiRewardIndex();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.collateralType = reader.string();
          break;

        case 2:
          message.rewardIndexes.push(RewardIndex.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MultiRewardIndex>): MultiRewardIndex {
    const message = createBaseMultiRewardIndex();
    message.collateralType = object.collateralType ?? "";
    message.rewardIndexes = object.rewardIndexes?.map(e => RewardIndex.fromPartial(e)) || [];
    return message;
  }

};

function createBaseMultiRewardIndexesProto(): MultiRewardIndexesProto {
  return {
    multiRewardIndexes: []
  };
}

export const MultiRewardIndexesProto = {
  encode(message: MultiRewardIndexesProto, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.multiRewardIndexes) {
      MultiRewardIndex.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MultiRewardIndexesProto {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMultiRewardIndexesProto();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.multiRewardIndexes.push(MultiRewardIndex.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MultiRewardIndexesProto>): MultiRewardIndexesProto {
    const message = createBaseMultiRewardIndexesProto();
    message.multiRewardIndexes = object.multiRewardIndexes?.map(e => MultiRewardIndex.fromPartial(e)) || [];
    return message;
  }

};

function createBaseUSDXMintingClaim(): USDXMintingClaim {
  return {
    baseClaim: undefined,
    rewardIndexes: []
  };
}

export const USDXMintingClaim = {
  encode(message: USDXMintingClaim, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.baseClaim !== undefined) {
      BaseClaim.encode(message.baseClaim, writer.uint32(10).fork()).ldelim();
    }

    for (const v of message.rewardIndexes) {
      RewardIndex.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): USDXMintingClaim {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUSDXMintingClaim();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.baseClaim = BaseClaim.decode(reader, reader.uint32());
          break;

        case 2:
          message.rewardIndexes.push(RewardIndex.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<USDXMintingClaim>): USDXMintingClaim {
    const message = createBaseUSDXMintingClaim();
    message.baseClaim = object.baseClaim !== undefined && object.baseClaim !== null ? BaseClaim.fromPartial(object.baseClaim) : undefined;
    message.rewardIndexes = object.rewardIndexes?.map(e => RewardIndex.fromPartial(e)) || [];
    return message;
  }

};

function createBaseJoltLiquidityProviderClaim(): JoltLiquidityProviderClaim {
  return {
    baseClaim: undefined,
    supplyRewardIndexes: [],
    borrowRewardIndexes: []
  };
}

export const JoltLiquidityProviderClaim = {
  encode(message: JoltLiquidityProviderClaim, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.baseClaim !== undefined) {
      BaseMultiClaim.encode(message.baseClaim, writer.uint32(10).fork()).ldelim();
    }

    for (const v of message.supplyRewardIndexes) {
      MultiRewardIndex.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    for (const v of message.borrowRewardIndexes) {
      MultiRewardIndex.encode(v!, writer.uint32(26).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): JoltLiquidityProviderClaim {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseJoltLiquidityProviderClaim();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.baseClaim = BaseMultiClaim.decode(reader, reader.uint32());
          break;

        case 2:
          message.supplyRewardIndexes.push(MultiRewardIndex.decode(reader, reader.uint32()));
          break;

        case 3:
          message.borrowRewardIndexes.push(MultiRewardIndex.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<JoltLiquidityProviderClaim>): JoltLiquidityProviderClaim {
    const message = createBaseJoltLiquidityProviderClaim();
    message.baseClaim = object.baseClaim !== undefined && object.baseClaim !== null ? BaseMultiClaim.fromPartial(object.baseClaim) : undefined;
    message.supplyRewardIndexes = object.supplyRewardIndexes?.map(e => MultiRewardIndex.fromPartial(e)) || [];
    message.borrowRewardIndexes = object.borrowRewardIndexes?.map(e => MultiRewardIndex.fromPartial(e)) || [];
    return message;
  }

};

function createBaseDelegatorClaim(): DelegatorClaim {
  return {
    baseClaim: undefined,
    rewardIndexes: []
  };
}

export const DelegatorClaim = {
  encode(message: DelegatorClaim, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.baseClaim !== undefined) {
      BaseMultiClaim.encode(message.baseClaim, writer.uint32(10).fork()).ldelim();
    }

    for (const v of message.rewardIndexes) {
      MultiRewardIndex.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DelegatorClaim {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDelegatorClaim();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.baseClaim = BaseMultiClaim.decode(reader, reader.uint32());
          break;

        case 2:
          message.rewardIndexes.push(MultiRewardIndex.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<DelegatorClaim>): DelegatorClaim {
    const message = createBaseDelegatorClaim();
    message.baseClaim = object.baseClaim !== undefined && object.baseClaim !== null ? BaseMultiClaim.fromPartial(object.baseClaim) : undefined;
    message.rewardIndexes = object.rewardIndexes?.map(e => MultiRewardIndex.fromPartial(e)) || [];
    return message;
  }

};

function createBaseSwapClaim(): SwapClaim {
  return {
    baseClaim: undefined,
    rewardIndexes: []
  };
}

export const SwapClaim = {
  encode(message: SwapClaim, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.baseClaim !== undefined) {
      BaseMultiClaim.encode(message.baseClaim, writer.uint32(10).fork()).ldelim();
    }

    for (const v of message.rewardIndexes) {
      MultiRewardIndex.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SwapClaim {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSwapClaim();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.baseClaim = BaseMultiClaim.decode(reader, reader.uint32());
          break;

        case 2:
          message.rewardIndexes.push(MultiRewardIndex.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<SwapClaim>): SwapClaim {
    const message = createBaseSwapClaim();
    message.baseClaim = object.baseClaim !== undefined && object.baseClaim !== null ? BaseMultiClaim.fromPartial(object.baseClaim) : undefined;
    message.rewardIndexes = object.rewardIndexes?.map(e => MultiRewardIndex.fromPartial(e)) || [];
    return message;
  }

};

function createBaseSavingsClaim(): SavingsClaim {
  return {
    baseClaim: undefined,
    rewardIndexes: []
  };
}

export const SavingsClaim = {
  encode(message: SavingsClaim, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.baseClaim !== undefined) {
      BaseMultiClaim.encode(message.baseClaim, writer.uint32(10).fork()).ldelim();
    }

    for (const v of message.rewardIndexes) {
      MultiRewardIndex.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SavingsClaim {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSavingsClaim();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.baseClaim = BaseMultiClaim.decode(reader, reader.uint32());
          break;

        case 2:
          message.rewardIndexes.push(MultiRewardIndex.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<SavingsClaim>): SavingsClaim {
    const message = createBaseSavingsClaim();
    message.baseClaim = object.baseClaim !== undefined && object.baseClaim !== null ? BaseMultiClaim.fromPartial(object.baseClaim) : undefined;
    message.rewardIndexes = object.rewardIndexes?.map(e => MultiRewardIndex.fromPartial(e)) || [];
    return message;
  }

};