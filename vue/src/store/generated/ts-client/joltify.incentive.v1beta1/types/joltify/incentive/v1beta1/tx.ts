/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "joltify.incentive.v1beta1";

/**
 * Selection is a pair of denom and multiplier name. It holds the choice of multiplier a user makes when they claim a
 * denom.
 */
export interface Selection {
  denom: string;
  multiplierName: string;
}

/** MsgClaimUSDXMintingReward message type used to claim USDX minting rewards */
export interface MsgClaimUSDXMintingReward {
  sender: string;
  multiplierName: string;
}

/** MsgClaimUSDXMintingRewardResponse defines the Msg/ClaimUSDXMintingReward response type. */
export interface MsgClaimUSDXMintingRewardResponse {
}

/** MsgClaimHardReward message type used to claim Hard liquidity provider rewards */
export interface MsgClaimJoltReward {
  sender: string;
  denomsToClaim: Selection[];
}

/** MsgClaimJoltRewardResponse defines the Msg/ClaimHardReward response type. */
export interface MsgClaimJoltRewardResponse {
}

/** MsgClaimDelegatorReward message type used to claim delegator rewards */
export interface MsgClaimDelegatorReward {
  sender: string;
  denomsToClaim: Selection[];
}

/** MsgClaimDelegatorRewardResponse defines the Msg/ClaimDelegatorReward response type. */
export interface MsgClaimDelegatorRewardResponse {
}

/** MsgClaimSwapReward message type used to claim delegator rewards */
export interface MsgClaimSwapReward {
  sender: string;
  denomsToClaim: Selection[];
}

/** MsgClaimSwapRewardResponse defines the Msg/ClaimSwapReward response type. */
export interface MsgClaimSwapRewardResponse {
}

/** MsgClaimSavingsReward message type used to claim savings rewards */
export interface MsgClaimSavingsReward {
  sender: string;
  denomsToClaim: Selection[];
}

/** MsgClaimSavingsRewardResponse defines the Msg/ClaimSavingsReward response type. */
export interface MsgClaimSavingsRewardResponse {
}

function createBaseSelection(): Selection {
  return { denom: "", multiplierName: "" };
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

  fromJSON(object: any): Selection {
    return {
      denom: isSet(object.denom) ? String(object.denom) : "",
      multiplierName: isSet(object.multiplierName) ? String(object.multiplierName) : "",
    };
  },

  toJSON(message: Selection): unknown {
    const obj: any = {};
    message.denom !== undefined && (obj.denom = message.denom);
    message.multiplierName !== undefined && (obj.multiplierName = message.multiplierName);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Selection>, I>>(object: I): Selection {
    const message = createBaseSelection();
    message.denom = object.denom ?? "";
    message.multiplierName = object.multiplierName ?? "";
    return message;
  },
};

function createBaseMsgClaimUSDXMintingReward(): MsgClaimUSDXMintingReward {
  return { sender: "", multiplierName: "" };
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

  fromJSON(object: any): MsgClaimUSDXMintingReward {
    return {
      sender: isSet(object.sender) ? String(object.sender) : "",
      multiplierName: isSet(object.multiplierName) ? String(object.multiplierName) : "",
    };
  },

  toJSON(message: MsgClaimUSDXMintingReward): unknown {
    const obj: any = {};
    message.sender !== undefined && (obj.sender = message.sender);
    message.multiplierName !== undefined && (obj.multiplierName = message.multiplierName);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimUSDXMintingReward>, I>>(object: I): MsgClaimUSDXMintingReward {
    const message = createBaseMsgClaimUSDXMintingReward();
    message.sender = object.sender ?? "";
    message.multiplierName = object.multiplierName ?? "";
    return message;
  },
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

  fromJSON(_: any): MsgClaimUSDXMintingRewardResponse {
    return {};
  },

  toJSON(_: MsgClaimUSDXMintingRewardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimUSDXMintingRewardResponse>, I>>(
    _: I,
  ): MsgClaimUSDXMintingRewardResponse {
    const message = createBaseMsgClaimUSDXMintingRewardResponse();
    return message;
  },
};

function createBaseMsgClaimJoltReward(): MsgClaimJoltReward {
  return { sender: "", denomsToClaim: [] };
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

  fromJSON(object: any): MsgClaimJoltReward {
    return {
      sender: isSet(object.sender) ? String(object.sender) : "",
      denomsToClaim: Array.isArray(object?.denomsToClaim)
        ? object.denomsToClaim.map((e: any) => Selection.fromJSON(e))
        : [],
    };
  },

  toJSON(message: MsgClaimJoltReward): unknown {
    const obj: any = {};
    message.sender !== undefined && (obj.sender = message.sender);
    if (message.denomsToClaim) {
      obj.denomsToClaim = message.denomsToClaim.map((e) => e ? Selection.toJSON(e) : undefined);
    } else {
      obj.denomsToClaim = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimJoltReward>, I>>(object: I): MsgClaimJoltReward {
    const message = createBaseMsgClaimJoltReward();
    message.sender = object.sender ?? "";
    message.denomsToClaim = object.denomsToClaim?.map((e) => Selection.fromPartial(e)) || [];
    return message;
  },
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

  fromJSON(_: any): MsgClaimJoltRewardResponse {
    return {};
  },

  toJSON(_: MsgClaimJoltRewardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimJoltRewardResponse>, I>>(_: I): MsgClaimJoltRewardResponse {
    const message = createBaseMsgClaimJoltRewardResponse();
    return message;
  },
};

function createBaseMsgClaimDelegatorReward(): MsgClaimDelegatorReward {
  return { sender: "", denomsToClaim: [] };
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

  fromJSON(object: any): MsgClaimDelegatorReward {
    return {
      sender: isSet(object.sender) ? String(object.sender) : "",
      denomsToClaim: Array.isArray(object?.denomsToClaim)
        ? object.denomsToClaim.map((e: any) => Selection.fromJSON(e))
        : [],
    };
  },

  toJSON(message: MsgClaimDelegatorReward): unknown {
    const obj: any = {};
    message.sender !== undefined && (obj.sender = message.sender);
    if (message.denomsToClaim) {
      obj.denomsToClaim = message.denomsToClaim.map((e) => e ? Selection.toJSON(e) : undefined);
    } else {
      obj.denomsToClaim = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimDelegatorReward>, I>>(object: I): MsgClaimDelegatorReward {
    const message = createBaseMsgClaimDelegatorReward();
    message.sender = object.sender ?? "";
    message.denomsToClaim = object.denomsToClaim?.map((e) => Selection.fromPartial(e)) || [];
    return message;
  },
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

  fromJSON(_: any): MsgClaimDelegatorRewardResponse {
    return {};
  },

  toJSON(_: MsgClaimDelegatorRewardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimDelegatorRewardResponse>, I>>(_: I): MsgClaimDelegatorRewardResponse {
    const message = createBaseMsgClaimDelegatorRewardResponse();
    return message;
  },
};

function createBaseMsgClaimSwapReward(): MsgClaimSwapReward {
  return { sender: "", denomsToClaim: [] };
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

  fromJSON(object: any): MsgClaimSwapReward {
    return {
      sender: isSet(object.sender) ? String(object.sender) : "",
      denomsToClaim: Array.isArray(object?.denomsToClaim)
        ? object.denomsToClaim.map((e: any) => Selection.fromJSON(e))
        : [],
    };
  },

  toJSON(message: MsgClaimSwapReward): unknown {
    const obj: any = {};
    message.sender !== undefined && (obj.sender = message.sender);
    if (message.denomsToClaim) {
      obj.denomsToClaim = message.denomsToClaim.map((e) => e ? Selection.toJSON(e) : undefined);
    } else {
      obj.denomsToClaim = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimSwapReward>, I>>(object: I): MsgClaimSwapReward {
    const message = createBaseMsgClaimSwapReward();
    message.sender = object.sender ?? "";
    message.denomsToClaim = object.denomsToClaim?.map((e) => Selection.fromPartial(e)) || [];
    return message;
  },
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

  fromJSON(_: any): MsgClaimSwapRewardResponse {
    return {};
  },

  toJSON(_: MsgClaimSwapRewardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimSwapRewardResponse>, I>>(_: I): MsgClaimSwapRewardResponse {
    const message = createBaseMsgClaimSwapRewardResponse();
    return message;
  },
};

function createBaseMsgClaimSavingsReward(): MsgClaimSavingsReward {
  return { sender: "", denomsToClaim: [] };
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

  fromJSON(object: any): MsgClaimSavingsReward {
    return {
      sender: isSet(object.sender) ? String(object.sender) : "",
      denomsToClaim: Array.isArray(object?.denomsToClaim)
        ? object.denomsToClaim.map((e: any) => Selection.fromJSON(e))
        : [],
    };
  },

  toJSON(message: MsgClaimSavingsReward): unknown {
    const obj: any = {};
    message.sender !== undefined && (obj.sender = message.sender);
    if (message.denomsToClaim) {
      obj.denomsToClaim = message.denomsToClaim.map((e) => e ? Selection.toJSON(e) : undefined);
    } else {
      obj.denomsToClaim = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimSavingsReward>, I>>(object: I): MsgClaimSavingsReward {
    const message = createBaseMsgClaimSavingsReward();
    message.sender = object.sender ?? "";
    message.denomsToClaim = object.denomsToClaim?.map((e) => Selection.fromPartial(e)) || [];
    return message;
  },
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

  fromJSON(_: any): MsgClaimSavingsRewardResponse {
    return {};
  },

  toJSON(_: MsgClaimSavingsRewardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimSavingsRewardResponse>, I>>(_: I): MsgClaimSavingsRewardResponse {
    const message = createBaseMsgClaimSavingsRewardResponse();
    return message;
  },
};

/** Msg defines the incentive Msg service. */
export interface Msg {
  /** ClaimJoltReward is a message type used to claim Hard liquidity provider rewards */
  ClaimJoltReward(request: MsgClaimJoltReward): Promise<MsgClaimJoltRewardResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.ClaimJoltReward = this.ClaimJoltReward.bind(this);
  }
  ClaimJoltReward(request: MsgClaimJoltReward): Promise<MsgClaimJoltRewardResponse> {
    const data = MsgClaimJoltReward.encode(request).finish();
    const promise = this.rpc.request("joltify.incentive.v1beta1.Msg", "ClaimJoltReward", data);
    return promise.then((data) => MsgClaimJoltRewardResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
