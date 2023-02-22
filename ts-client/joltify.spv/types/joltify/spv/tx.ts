/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "joltify.spv";

export interface MsgCreatePool {
  creator: string;
  projectIndex: number;
  poolName: string;
  apy: string;
  targetTokenAmount: Coin | undefined;
}

export interface MsgCreatePoolResponse {
  poolIndex: string;
}

export interface MsgAddInvestors {
  creator: string;
  poolIndex: string;
  investorID: string[];
}

export interface MsgAddInvestorsResponse {
  operationResult: boolean;
}

export interface MsgDeposit {
  creator: string;
  poolIndex: string;
  token: Coin | undefined;
}

export interface MsgDepositResponse {
}

export interface MsgBorrow {
  creator: string;
  poolIndex: string;
  borrowAmount: Coin | undefined;
}

export interface MsgBorrowResponse {
}

export interface MsgRepayInterest {
  creator: string;
  poolIndex: string;
  token: Coin | undefined;
}

export interface MsgRepayInterestResponse {
}

export interface MsgClaimInterest {
  creator: string;
  poolIndex: string;
}

export interface MsgClaimInterestResponse {
}

export interface MsgUpdatePool {
  creator: string;
  poolIndex: string;
  poolName: string;
  poolApy: string;
  targetTokenAmount: Coin | undefined;
}

export interface MsgUpdatePoolResponse {
}

export interface MsgActivePool {
  creator: string;
  poolIndex: string;
}

export interface MsgActivePoolResponse {
}

export interface MsgPayPrincipal {
  creator: string;
  poolIndex: string;
  token: Coin | undefined;
}

export interface MsgPayPrincipalResponse {
}

export interface MsgWithdrawPrincipal {
  creator: string;
  poolIndex: string;
  token: Coin | undefined;
}

export interface MsgWithdrawPrincipalResponse {
}

function createBaseMsgCreatePool(): MsgCreatePool {
  return { creator: "", projectIndex: 0, poolName: "", apy: "", targetTokenAmount: undefined };
}

export const MsgCreatePool = {
  encode(message: MsgCreatePool, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.projectIndex !== 0) {
      writer.uint32(16).int32(message.projectIndex);
    }
    if (message.poolName !== "") {
      writer.uint32(26).string(message.poolName);
    }
    if (message.apy !== "") {
      writer.uint32(34).string(message.apy);
    }
    if (message.targetTokenAmount !== undefined) {
      Coin.encode(message.targetTokenAmount, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreatePool {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreatePool();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.projectIndex = reader.int32();
          break;
        case 3:
          message.poolName = reader.string();
          break;
        case 4:
          message.apy = reader.string();
          break;
        case 5:
          message.targetTokenAmount = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreatePool {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      projectIndex: isSet(object.projectIndex) ? Number(object.projectIndex) : 0,
      poolName: isSet(object.poolName) ? String(object.poolName) : "",
      apy: isSet(object.apy) ? String(object.apy) : "",
      targetTokenAmount: isSet(object.targetTokenAmount) ? Coin.fromJSON(object.targetTokenAmount) : undefined,
    };
  },

  toJSON(message: MsgCreatePool): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.projectIndex !== undefined && (obj.projectIndex = Math.round(message.projectIndex));
    message.poolName !== undefined && (obj.poolName = message.poolName);
    message.apy !== undefined && (obj.apy = message.apy);
    message.targetTokenAmount !== undefined
      && (obj.targetTokenAmount = message.targetTokenAmount ? Coin.toJSON(message.targetTokenAmount) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreatePool>, I>>(object: I): MsgCreatePool {
    const message = createBaseMsgCreatePool();
    message.creator = object.creator ?? "";
    message.projectIndex = object.projectIndex ?? 0;
    message.poolName = object.poolName ?? "";
    message.apy = object.apy ?? "";
    message.targetTokenAmount = (object.targetTokenAmount !== undefined && object.targetTokenAmount !== null)
      ? Coin.fromPartial(object.targetTokenAmount)
      : undefined;
    return message;
  },
};

function createBaseMsgCreatePoolResponse(): MsgCreatePoolResponse {
  return { poolIndex: "" };
}

export const MsgCreatePoolResponse = {
  encode(message: MsgCreatePoolResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.poolIndex !== "") {
      writer.uint32(10).string(message.poolIndex);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreatePoolResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreatePoolResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.poolIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreatePoolResponse {
    return { poolIndex: isSet(object.poolIndex) ? String(object.poolIndex) : "" };
  },

  toJSON(message: MsgCreatePoolResponse): unknown {
    const obj: any = {};
    message.poolIndex !== undefined && (obj.poolIndex = message.poolIndex);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreatePoolResponse>, I>>(object: I): MsgCreatePoolResponse {
    const message = createBaseMsgCreatePoolResponse();
    message.poolIndex = object.poolIndex ?? "";
    return message;
  },
};

function createBaseMsgAddInvestors(): MsgAddInvestors {
  return { creator: "", poolIndex: "", investorID: [] };
}

export const MsgAddInvestors = {
  encode(message: MsgAddInvestors, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }
    for (const v of message.investorID) {
      writer.uint32(26).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddInvestors {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddInvestors();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.poolIndex = reader.string();
          break;
        case 3:
          message.investorID.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddInvestors {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      poolIndex: isSet(object.poolIndex) ? String(object.poolIndex) : "",
      investorID: Array.isArray(object?.investorID) ? object.investorID.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: MsgAddInvestors): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.poolIndex !== undefined && (obj.poolIndex = message.poolIndex);
    if (message.investorID) {
      obj.investorID = message.investorID.map((e) => e);
    } else {
      obj.investorID = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddInvestors>, I>>(object: I): MsgAddInvestors {
    const message = createBaseMsgAddInvestors();
    message.creator = object.creator ?? "";
    message.poolIndex = object.poolIndex ?? "";
    message.investorID = object.investorID?.map((e) => e) || [];
    return message;
  },
};

function createBaseMsgAddInvestorsResponse(): MsgAddInvestorsResponse {
  return { operationResult: false };
}

export const MsgAddInvestorsResponse = {
  encode(message: MsgAddInvestorsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.operationResult === true) {
      writer.uint32(8).bool(message.operationResult);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddInvestorsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddInvestorsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.operationResult = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddInvestorsResponse {
    return { operationResult: isSet(object.operationResult) ? Boolean(object.operationResult) : false };
  },

  toJSON(message: MsgAddInvestorsResponse): unknown {
    const obj: any = {};
    message.operationResult !== undefined && (obj.operationResult = message.operationResult);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddInvestorsResponse>, I>>(object: I): MsgAddInvestorsResponse {
    const message = createBaseMsgAddInvestorsResponse();
    message.operationResult = object.operationResult ?? false;
    return message;
  },
};

function createBaseMsgDeposit(): MsgDeposit {
  return { creator: "", poolIndex: "", token: undefined };
}

export const MsgDeposit = {
  encode(message: MsgDeposit, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }
    if (message.token !== undefined) {
      Coin.encode(message.token, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeposit {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeposit();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.poolIndex = reader.string();
          break;
        case 3:
          message.token = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeposit {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      poolIndex: isSet(object.poolIndex) ? String(object.poolIndex) : "",
      token: isSet(object.token) ? Coin.fromJSON(object.token) : undefined,
    };
  },

  toJSON(message: MsgDeposit): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.poolIndex !== undefined && (obj.poolIndex = message.poolIndex);
    message.token !== undefined && (obj.token = message.token ? Coin.toJSON(message.token) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeposit>, I>>(object: I): MsgDeposit {
    const message = createBaseMsgDeposit();
    message.creator = object.creator ?? "";
    message.poolIndex = object.poolIndex ?? "";
    message.token = (object.token !== undefined && object.token !== null) ? Coin.fromPartial(object.token) : undefined;
    return message;
  },
};

function createBaseMsgDepositResponse(): MsgDepositResponse {
  return {};
}

export const MsgDepositResponse = {
  encode(_: MsgDepositResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDepositResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDepositResponse();
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

  fromJSON(_: any): MsgDepositResponse {
    return {};
  },

  toJSON(_: MsgDepositResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDepositResponse>, I>>(_: I): MsgDepositResponse {
    const message = createBaseMsgDepositResponse();
    return message;
  },
};

function createBaseMsgBorrow(): MsgBorrow {
  return { creator: "", poolIndex: "", borrowAmount: undefined };
}

export const MsgBorrow = {
  encode(message: MsgBorrow, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }
    if (message.borrowAmount !== undefined) {
      Coin.encode(message.borrowAmount, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBorrow {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBorrow();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.poolIndex = reader.string();
          break;
        case 3:
          message.borrowAmount = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgBorrow {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      poolIndex: isSet(object.poolIndex) ? String(object.poolIndex) : "",
      borrowAmount: isSet(object.borrowAmount) ? Coin.fromJSON(object.borrowAmount) : undefined,
    };
  },

  toJSON(message: MsgBorrow): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.poolIndex !== undefined && (obj.poolIndex = message.poolIndex);
    message.borrowAmount !== undefined
      && (obj.borrowAmount = message.borrowAmount ? Coin.toJSON(message.borrowAmount) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgBorrow>, I>>(object: I): MsgBorrow {
    const message = createBaseMsgBorrow();
    message.creator = object.creator ?? "";
    message.poolIndex = object.poolIndex ?? "";
    message.borrowAmount = (object.borrowAmount !== undefined && object.borrowAmount !== null)
      ? Coin.fromPartial(object.borrowAmount)
      : undefined;
    return message;
  },
};

function createBaseMsgBorrowResponse(): MsgBorrowResponse {
  return {};
}

export const MsgBorrowResponse = {
  encode(_: MsgBorrowResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBorrowResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBorrowResponse();
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

  fromJSON(_: any): MsgBorrowResponse {
    return {};
  },

  toJSON(_: MsgBorrowResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgBorrowResponse>, I>>(_: I): MsgBorrowResponse {
    const message = createBaseMsgBorrowResponse();
    return message;
  },
};

function createBaseMsgRepayInterest(): MsgRepayInterest {
  return { creator: "", poolIndex: "", token: undefined };
}

export const MsgRepayInterest = {
  encode(message: MsgRepayInterest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }
    if (message.token !== undefined) {
      Coin.encode(message.token, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRepayInterest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRepayInterest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.poolIndex = reader.string();
          break;
        case 4:
          message.token = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRepayInterest {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      poolIndex: isSet(object.poolIndex) ? String(object.poolIndex) : "",
      token: isSet(object.token) ? Coin.fromJSON(object.token) : undefined,
    };
  },

  toJSON(message: MsgRepayInterest): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.poolIndex !== undefined && (obj.poolIndex = message.poolIndex);
    message.token !== undefined && (obj.token = message.token ? Coin.toJSON(message.token) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRepayInterest>, I>>(object: I): MsgRepayInterest {
    const message = createBaseMsgRepayInterest();
    message.creator = object.creator ?? "";
    message.poolIndex = object.poolIndex ?? "";
    message.token = (object.token !== undefined && object.token !== null) ? Coin.fromPartial(object.token) : undefined;
    return message;
  },
};

function createBaseMsgRepayInterestResponse(): MsgRepayInterestResponse {
  return {};
}

export const MsgRepayInterestResponse = {
  encode(_: MsgRepayInterestResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRepayInterestResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRepayInterestResponse();
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

  fromJSON(_: any): MsgRepayInterestResponse {
    return {};
  },

  toJSON(_: MsgRepayInterestResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRepayInterestResponse>, I>>(_: I): MsgRepayInterestResponse {
    const message = createBaseMsgRepayInterestResponse();
    return message;
  },
};

function createBaseMsgClaimInterest(): MsgClaimInterest {
  return { creator: "", poolIndex: "" };
}

export const MsgClaimInterest = {
  encode(message: MsgClaimInterest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimInterest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimInterest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.poolIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgClaimInterest {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      poolIndex: isSet(object.poolIndex) ? String(object.poolIndex) : "",
    };
  },

  toJSON(message: MsgClaimInterest): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.poolIndex !== undefined && (obj.poolIndex = message.poolIndex);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimInterest>, I>>(object: I): MsgClaimInterest {
    const message = createBaseMsgClaimInterest();
    message.creator = object.creator ?? "";
    message.poolIndex = object.poolIndex ?? "";
    return message;
  },
};

function createBaseMsgClaimInterestResponse(): MsgClaimInterestResponse {
  return {};
}

export const MsgClaimInterestResponse = {
  encode(_: MsgClaimInterestResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimInterestResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimInterestResponse();
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

  fromJSON(_: any): MsgClaimInterestResponse {
    return {};
  },

  toJSON(_: MsgClaimInterestResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimInterestResponse>, I>>(_: I): MsgClaimInterestResponse {
    const message = createBaseMsgClaimInterestResponse();
    return message;
  },
};

function createBaseMsgUpdatePool(): MsgUpdatePool {
  return { creator: "", poolIndex: "", poolName: "", poolApy: "", targetTokenAmount: undefined };
}

export const MsgUpdatePool = {
  encode(message: MsgUpdatePool, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }
    if (message.poolName !== "") {
      writer.uint32(26).string(message.poolName);
    }
    if (message.poolApy !== "") {
      writer.uint32(34).string(message.poolApy);
    }
    if (message.targetTokenAmount !== undefined) {
      Coin.encode(message.targetTokenAmount, writer.uint32(50).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdatePool {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdatePool();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.poolIndex = reader.string();
          break;
        case 3:
          message.poolName = reader.string();
          break;
        case 4:
          message.poolApy = reader.string();
          break;
        case 6:
          message.targetTokenAmount = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdatePool {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      poolIndex: isSet(object.poolIndex) ? String(object.poolIndex) : "",
      poolName: isSet(object.poolName) ? String(object.poolName) : "",
      poolApy: isSet(object.poolApy) ? String(object.poolApy) : "",
      targetTokenAmount: isSet(object.targetTokenAmount) ? Coin.fromJSON(object.targetTokenAmount) : undefined,
    };
  },

  toJSON(message: MsgUpdatePool): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.poolIndex !== undefined && (obj.poolIndex = message.poolIndex);
    message.poolName !== undefined && (obj.poolName = message.poolName);
    message.poolApy !== undefined && (obj.poolApy = message.poolApy);
    message.targetTokenAmount !== undefined
      && (obj.targetTokenAmount = message.targetTokenAmount ? Coin.toJSON(message.targetTokenAmount) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdatePool>, I>>(object: I): MsgUpdatePool {
    const message = createBaseMsgUpdatePool();
    message.creator = object.creator ?? "";
    message.poolIndex = object.poolIndex ?? "";
    message.poolName = object.poolName ?? "";
    message.poolApy = object.poolApy ?? "";
    message.targetTokenAmount = (object.targetTokenAmount !== undefined && object.targetTokenAmount !== null)
      ? Coin.fromPartial(object.targetTokenAmount)
      : undefined;
    return message;
  },
};

function createBaseMsgUpdatePoolResponse(): MsgUpdatePoolResponse {
  return {};
}

export const MsgUpdatePoolResponse = {
  encode(_: MsgUpdatePoolResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdatePoolResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdatePoolResponse();
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

  fromJSON(_: any): MsgUpdatePoolResponse {
    return {};
  },

  toJSON(_: MsgUpdatePoolResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdatePoolResponse>, I>>(_: I): MsgUpdatePoolResponse {
    const message = createBaseMsgUpdatePoolResponse();
    return message;
  },
};

function createBaseMsgActivePool(): MsgActivePool {
  return { creator: "", poolIndex: "" };
}

export const MsgActivePool = {
  encode(message: MsgActivePool, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgActivePool {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgActivePool();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.poolIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgActivePool {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      poolIndex: isSet(object.poolIndex) ? String(object.poolIndex) : "",
    };
  },

  toJSON(message: MsgActivePool): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.poolIndex !== undefined && (obj.poolIndex = message.poolIndex);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgActivePool>, I>>(object: I): MsgActivePool {
    const message = createBaseMsgActivePool();
    message.creator = object.creator ?? "";
    message.poolIndex = object.poolIndex ?? "";
    return message;
  },
};

function createBaseMsgActivePoolResponse(): MsgActivePoolResponse {
  return {};
}

export const MsgActivePoolResponse = {
  encode(_: MsgActivePoolResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgActivePoolResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgActivePoolResponse();
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

  fromJSON(_: any): MsgActivePoolResponse {
    return {};
  },

  toJSON(_: MsgActivePoolResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgActivePoolResponse>, I>>(_: I): MsgActivePoolResponse {
    const message = createBaseMsgActivePoolResponse();
    return message;
  },
};

function createBaseMsgPayPrincipal(): MsgPayPrincipal {
  return { creator: "", poolIndex: "", token: undefined };
}

export const MsgPayPrincipal = {
  encode(message: MsgPayPrincipal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }
    if (message.token !== undefined) {
      Coin.encode(message.token, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgPayPrincipal {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgPayPrincipal();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.poolIndex = reader.string();
          break;
        case 3:
          message.token = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPayPrincipal {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      poolIndex: isSet(object.poolIndex) ? String(object.poolIndex) : "",
      token: isSet(object.token) ? Coin.fromJSON(object.token) : undefined,
    };
  },

  toJSON(message: MsgPayPrincipal): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.poolIndex !== undefined && (obj.poolIndex = message.poolIndex);
    message.token !== undefined && (obj.token = message.token ? Coin.toJSON(message.token) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgPayPrincipal>, I>>(object: I): MsgPayPrincipal {
    const message = createBaseMsgPayPrincipal();
    message.creator = object.creator ?? "";
    message.poolIndex = object.poolIndex ?? "";
    message.token = (object.token !== undefined && object.token !== null) ? Coin.fromPartial(object.token) : undefined;
    return message;
  },
};

function createBaseMsgPayPrincipalResponse(): MsgPayPrincipalResponse {
  return {};
}

export const MsgPayPrincipalResponse = {
  encode(_: MsgPayPrincipalResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgPayPrincipalResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgPayPrincipalResponse();
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

  fromJSON(_: any): MsgPayPrincipalResponse {
    return {};
  },

  toJSON(_: MsgPayPrincipalResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgPayPrincipalResponse>, I>>(_: I): MsgPayPrincipalResponse {
    const message = createBaseMsgPayPrincipalResponse();
    return message;
  },
};

function createBaseMsgWithdrawPrincipal(): MsgWithdrawPrincipal {
  return { creator: "", poolIndex: "", token: undefined };
}

export const MsgWithdrawPrincipal = {
  encode(message: MsgWithdrawPrincipal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }
    if (message.token !== undefined) {
      Coin.encode(message.token, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgWithdrawPrincipal {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgWithdrawPrincipal();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.poolIndex = reader.string();
          break;
        case 3:
          message.token = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgWithdrawPrincipal {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      poolIndex: isSet(object.poolIndex) ? String(object.poolIndex) : "",
      token: isSet(object.token) ? Coin.fromJSON(object.token) : undefined,
    };
  },

  toJSON(message: MsgWithdrawPrincipal): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.poolIndex !== undefined && (obj.poolIndex = message.poolIndex);
    message.token !== undefined && (obj.token = message.token ? Coin.toJSON(message.token) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgWithdrawPrincipal>, I>>(object: I): MsgWithdrawPrincipal {
    const message = createBaseMsgWithdrawPrincipal();
    message.creator = object.creator ?? "";
    message.poolIndex = object.poolIndex ?? "";
    message.token = (object.token !== undefined && object.token !== null) ? Coin.fromPartial(object.token) : undefined;
    return message;
  },
};

function createBaseMsgWithdrawPrincipalResponse(): MsgWithdrawPrincipalResponse {
  return {};
}

export const MsgWithdrawPrincipalResponse = {
  encode(_: MsgWithdrawPrincipalResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgWithdrawPrincipalResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgWithdrawPrincipalResponse();
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

  fromJSON(_: any): MsgWithdrawPrincipalResponse {
    return {};
  },

  toJSON(_: MsgWithdrawPrincipalResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgWithdrawPrincipalResponse>, I>>(_: I): MsgWithdrawPrincipalResponse {
    const message = createBaseMsgWithdrawPrincipalResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreatePool(request: MsgCreatePool): Promise<MsgCreatePoolResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  AddInvestors(request: MsgAddInvestors): Promise<MsgAddInvestorsResponse>;
  Deposit(request: MsgDeposit): Promise<MsgDepositResponse>;
  Borrow(request: MsgBorrow): Promise<MsgBorrowResponse>;
  RepayInterest(request: MsgRepayInterest): Promise<MsgRepayInterestResponse>;
  ClaimInterest(request: MsgClaimInterest): Promise<MsgClaimInterestResponse>;
  UpdatePool(request: MsgUpdatePool): Promise<MsgUpdatePoolResponse>;
  ActivePool(request: MsgActivePool): Promise<MsgActivePoolResponse>;
  PayPrincipal(request: MsgPayPrincipal): Promise<MsgPayPrincipalResponse>;
  WithdrawPrincipal(request: MsgWithdrawPrincipal): Promise<MsgWithdrawPrincipalResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreatePool = this.CreatePool.bind(this);
    this.AddInvestors = this.AddInvestors.bind(this);
    this.Deposit = this.Deposit.bind(this);
    this.Borrow = this.Borrow.bind(this);
    this.RepayInterest = this.RepayInterest.bind(this);
    this.ClaimInterest = this.ClaimInterest.bind(this);
    this.UpdatePool = this.UpdatePool.bind(this);
    this.ActivePool = this.ActivePool.bind(this);
    this.PayPrincipal = this.PayPrincipal.bind(this);
    this.WithdrawPrincipal = this.WithdrawPrincipal.bind(this);
  }
  CreatePool(request: MsgCreatePool): Promise<MsgCreatePoolResponse> {
    const data = MsgCreatePool.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "CreatePool", data);
    return promise.then((data) => MsgCreatePoolResponse.decode(new _m0.Reader(data)));
  }

  AddInvestors(request: MsgAddInvestors): Promise<MsgAddInvestorsResponse> {
    const data = MsgAddInvestors.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "AddInvestors", data);
    return promise.then((data) => MsgAddInvestorsResponse.decode(new _m0.Reader(data)));
  }

  Deposit(request: MsgDeposit): Promise<MsgDepositResponse> {
    const data = MsgDeposit.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "Deposit", data);
    return promise.then((data) => MsgDepositResponse.decode(new _m0.Reader(data)));
  }

  Borrow(request: MsgBorrow): Promise<MsgBorrowResponse> {
    const data = MsgBorrow.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "Borrow", data);
    return promise.then((data) => MsgBorrowResponse.decode(new _m0.Reader(data)));
  }

  RepayInterest(request: MsgRepayInterest): Promise<MsgRepayInterestResponse> {
    const data = MsgRepayInterest.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "RepayInterest", data);
    return promise.then((data) => MsgRepayInterestResponse.decode(new _m0.Reader(data)));
  }

  ClaimInterest(request: MsgClaimInterest): Promise<MsgClaimInterestResponse> {
    const data = MsgClaimInterest.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "ClaimInterest", data);
    return promise.then((data) => MsgClaimInterestResponse.decode(new _m0.Reader(data)));
  }

  UpdatePool(request: MsgUpdatePool): Promise<MsgUpdatePoolResponse> {
    const data = MsgUpdatePool.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "UpdatePool", data);
    return promise.then((data) => MsgUpdatePoolResponse.decode(new _m0.Reader(data)));
  }

  ActivePool(request: MsgActivePool): Promise<MsgActivePoolResponse> {
    const data = MsgActivePool.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "ActivePool", data);
    return promise.then((data) => MsgActivePoolResponse.decode(new _m0.Reader(data)));
  }

  PayPrincipal(request: MsgPayPrincipal): Promise<MsgPayPrincipalResponse> {
    const data = MsgPayPrincipal.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "PayPrincipal", data);
    return promise.then((data) => MsgPayPrincipalResponse.decode(new _m0.Reader(data)));
  }

  WithdrawPrincipal(request: MsgWithdrawPrincipal): Promise<MsgWithdrawPrincipalResponse> {
    const data = MsgWithdrawPrincipal.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "WithdrawPrincipal", data);
    return promise.then((data) => MsgWithdrawPrincipalResponse.decode(new _m0.Reader(data)));
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
