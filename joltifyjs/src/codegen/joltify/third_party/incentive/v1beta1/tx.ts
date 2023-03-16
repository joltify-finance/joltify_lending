import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../../../helpers";
/**
 * Selection is a pair of denom and multiplier name. It holds the choice of multiplier a user makes when they claim a
 * denom.
 */

export interface Selection {
  denom: string;
  multiplierName: string;
}
/**
 * Selection is a pair of denom and multiplier name. It holds the choice of multiplier a user makes when they claim a
 * denom.
 */

export interface SelectionSDKType {
  denom: string;
  multiplier_name: string;
}
/** MsgClaimUSDXMintingReward message type used to claim USDX minting rewards */

export interface MsgClaimUSDXMintingReward {
  sender: string;
  multiplierName: string;
}
/** MsgClaimUSDXMintingReward message type used to claim USDX minting rewards */

export interface MsgClaimUSDXMintingRewardSDKType {
  sender: string;
  multiplier_name: string;
}
/** MsgClaimUSDXMintingRewardResponse defines the Msg/ClaimUSDXMintingReward response type. */

export interface MsgClaimUSDXMintingRewardResponse {}
/** MsgClaimUSDXMintingRewardResponse defines the Msg/ClaimUSDXMintingReward response type. */

export interface MsgClaimUSDXMintingRewardResponseSDKType {}
/** MsgClaimHardReward message type used to claim Hard liquidity provider rewards */

export interface MsgClaimJoltReward {
  sender: string;
  denomsToClaim: Selection[];
}
/** MsgClaimHardReward message type used to claim Hard liquidity provider rewards */

export interface MsgClaimJoltRewardSDKType {
  sender: string;
  denoms_to_claim: SelectionSDKType[];
}
/** MsgClaimJoltRewardResponse defines the Msg/ClaimHardReward response type. */

export interface MsgClaimJoltRewardResponse {}
/** MsgClaimJoltRewardResponse defines the Msg/ClaimHardReward response type. */

export interface MsgClaimJoltRewardResponseSDKType {}
/** MsgClaimDelegatorReward message type used to claim delegator rewards */

export interface MsgClaimDelegatorReward {
  sender: string;
  denomsToClaim: Selection[];
}
/** MsgClaimDelegatorReward message type used to claim delegator rewards */

export interface MsgClaimDelegatorRewardSDKType {
  sender: string;
  denoms_to_claim: SelectionSDKType[];
}
/** MsgClaimDelegatorRewardResponse defines the Msg/ClaimDelegatorReward response type. */

export interface MsgClaimDelegatorRewardResponse {}
/** MsgClaimDelegatorRewardResponse defines the Msg/ClaimDelegatorReward response type. */

export interface MsgClaimDelegatorRewardResponseSDKType {}
/** MsgClaimSwapReward message type used to claim delegator rewards */

export interface MsgClaimSwapReward {
  sender: string;
  denomsToClaim: Selection[];
}
/** MsgClaimSwapReward message type used to claim delegator rewards */

export interface MsgClaimSwapRewardSDKType {
  sender: string;
  denoms_to_claim: SelectionSDKType[];
}
/** MsgClaimSwapRewardResponse defines the Msg/ClaimSwapReward response type. */

export interface MsgClaimSwapRewardResponse {}
/** MsgClaimSwapRewardResponse defines the Msg/ClaimSwapReward response type. */

export interface MsgClaimSwapRewardResponseSDKType {}
/** MsgClaimSavingsReward message type used to claim savings rewards */

export interface MsgClaimSavingsReward {
  sender: string;
  denomsToClaim: Selection[];
}
/** MsgClaimSavingsReward message type used to claim savings rewards */

export interface MsgClaimSavingsRewardSDKType {
  sender: string;
  denoms_to_claim: SelectionSDKType[];
}
/** MsgClaimSavingsRewardResponse defines the Msg/ClaimSavingsReward response type. */

export interface MsgClaimSavingsRewardResponse {}
/** MsgClaimSavingsRewardResponse defines the Msg/ClaimSavingsReward response type. */

export interface MsgClaimSavingsRewardResponseSDKType {}

function createBaseSelection(): Selection {
  return {
    denom: "",
    multiplierName: ""
  };
}

export const Selection = {
  encode(message: Selection, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }

    if (message.multiplierName !== "") {
      writer.uint32(18).string(message.multiplierName);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Selection {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSelection();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;

        case 2:
          message.multiplierName = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<Selection>): Selection {
    const message = createBaseSelection();
    message.denom = object.denom ?? "";
    message.multiplierName = object.multiplierName ?? "";
    return message;
  }

};

function createBaseMsgClaimUSDXMintingReward(): MsgClaimUSDXMintingReward {
  return {
    sender: "",
    multiplierName: ""
  };
}

export const MsgClaimUSDXMintingReward = {
  encode(message: MsgClaimUSDXMintingReward, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sender !== "") {
      writer.uint32(10).string(message.sender);
    }

    if (message.multiplierName !== "") {
      writer.uint32(18).string(message.multiplierName);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimUSDXMintingReward {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimUSDXMintingReward();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.sender = reader.string();
          break;

        case 2:
          message.multiplierName = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgClaimUSDXMintingReward>): MsgClaimUSDXMintingReward {
    const message = createBaseMsgClaimUSDXMintingReward();
    message.sender = object.sender ?? "";
    message.multiplierName = object.multiplierName ?? "";
    return message;
  }

};

function createBaseMsgClaimUSDXMintingRewardResponse(): MsgClaimUSDXMintingRewardResponse {
  return {};
}

export const MsgClaimUSDXMintingRewardResponse = {
  encode(_: MsgClaimUSDXMintingRewardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimUSDXMintingRewardResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimUSDXMintingRewardResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(_: DeepPartial<MsgClaimUSDXMintingRewardResponse>): MsgClaimUSDXMintingRewardResponse {
    const message = createBaseMsgClaimUSDXMintingRewardResponse();
    return message;
  }

};

function createBaseMsgClaimJoltReward(): MsgClaimJoltReward {
  return {
    sender: "",
    denomsToClaim: []
  };
}

export const MsgClaimJoltReward = {
  encode(message: MsgClaimJoltReward, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sender !== "") {
      writer.uint32(10).string(message.sender);
    }

    for (const v of message.denomsToClaim) {
      Selection.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimJoltReward {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimJoltReward();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.sender = reader.string();
          break;

        case 2:
          message.denomsToClaim.push(Selection.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgClaimJoltReward>): MsgClaimJoltReward {
    const message = createBaseMsgClaimJoltReward();
    message.sender = object.sender ?? "";
    message.denomsToClaim = object.denomsToClaim?.map(e => Selection.fromPartial(e)) || [];
    return message;
  }

};

function createBaseMsgClaimJoltRewardResponse(): MsgClaimJoltRewardResponse {
  return {};
}

export const MsgClaimJoltRewardResponse = {
  encode(_: MsgClaimJoltRewardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimJoltRewardResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimJoltRewardResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(_: DeepPartial<MsgClaimJoltRewardResponse>): MsgClaimJoltRewardResponse {
    const message = createBaseMsgClaimJoltRewardResponse();
    return message;
  }

};

function createBaseMsgClaimDelegatorReward(): MsgClaimDelegatorReward {
  return {
    sender: "",
    denomsToClaim: []
  };
}

export const MsgClaimDelegatorReward = {
  encode(message: MsgClaimDelegatorReward, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sender !== "") {
      writer.uint32(10).string(message.sender);
    }

    for (const v of message.denomsToClaim) {
      Selection.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimDelegatorReward {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimDelegatorReward();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.sender = reader.string();
          break;

        case 2:
          message.denomsToClaim.push(Selection.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgClaimDelegatorReward>): MsgClaimDelegatorReward {
    const message = createBaseMsgClaimDelegatorReward();
    message.sender = object.sender ?? "";
    message.denomsToClaim = object.denomsToClaim?.map(e => Selection.fromPartial(e)) || [];
    return message;
  }

};

function createBaseMsgClaimDelegatorRewardResponse(): MsgClaimDelegatorRewardResponse {
  return {};
}

export const MsgClaimDelegatorRewardResponse = {
  encode(_: MsgClaimDelegatorRewardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimDelegatorRewardResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimDelegatorRewardResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(_: DeepPartial<MsgClaimDelegatorRewardResponse>): MsgClaimDelegatorRewardResponse {
    const message = createBaseMsgClaimDelegatorRewardResponse();
    return message;
  }

};

function createBaseMsgClaimSwapReward(): MsgClaimSwapReward {
  return {
    sender: "",
    denomsToClaim: []
  };
}

export const MsgClaimSwapReward = {
  encode(message: MsgClaimSwapReward, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sender !== "") {
      writer.uint32(10).string(message.sender);
    }

    for (const v of message.denomsToClaim) {
      Selection.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimSwapReward {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimSwapReward();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.sender = reader.string();
          break;

        case 2:
          message.denomsToClaim.push(Selection.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgClaimSwapReward>): MsgClaimSwapReward {
    const message = createBaseMsgClaimSwapReward();
    message.sender = object.sender ?? "";
    message.denomsToClaim = object.denomsToClaim?.map(e => Selection.fromPartial(e)) || [];
    return message;
  }

};

function createBaseMsgClaimSwapRewardResponse(): MsgClaimSwapRewardResponse {
  return {};
}

export const MsgClaimSwapRewardResponse = {
  encode(_: MsgClaimSwapRewardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimSwapRewardResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimSwapRewardResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(_: DeepPartial<MsgClaimSwapRewardResponse>): MsgClaimSwapRewardResponse {
    const message = createBaseMsgClaimSwapRewardResponse();
    return message;
  }

};

function createBaseMsgClaimSavingsReward(): MsgClaimSavingsReward {
  return {
    sender: "",
    denomsToClaim: []
  };
}

export const MsgClaimSavingsReward = {
  encode(message: MsgClaimSavingsReward, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sender !== "") {
      writer.uint32(10).string(message.sender);
    }

    for (const v of message.denomsToClaim) {
      Selection.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimSavingsReward {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimSavingsReward();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.sender = reader.string();
          break;

        case 2:
          message.denomsToClaim.push(Selection.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgClaimSavingsReward>): MsgClaimSavingsReward {
    const message = createBaseMsgClaimSavingsReward();
    message.sender = object.sender ?? "";
    message.denomsToClaim = object.denomsToClaim?.map(e => Selection.fromPartial(e)) || [];
    return message;
  }

};

function createBaseMsgClaimSavingsRewardResponse(): MsgClaimSavingsRewardResponse {
  return {};
}

export const MsgClaimSavingsRewardResponse = {
  encode(_: MsgClaimSavingsRewardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimSavingsRewardResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimSavingsRewardResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(_: DeepPartial<MsgClaimSavingsRewardResponse>): MsgClaimSavingsRewardResponse {
    const message = createBaseMsgClaimSavingsRewardResponse();
    return message;
  }

};