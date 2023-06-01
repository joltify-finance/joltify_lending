/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { Coin } from "../../cosmos/base/v1beta1/coin";
import { PoolProposal } from "./create_pool";
import { IssueToken } from "./issue_token";
import { OutboundTx } from "./outbound_tx";
import { CoinsQuota } from "./quota";
import { Validators } from "./staking";

export const protobufPackage = "joltify.vault";

export interface QueryPendingFeeRequest {
  pagination: PageRequest | undefined;
}

export interface QueryPendingFeeResponse {
  feecoin: Coin[];
  pagination: PageResponse | undefined;
}

export interface QueryGetOutboundTxRequest {
  requestID: string;
}

export interface QueryGetOutboundTxResponse {
  outboundTx: OutboundTx | undefined;
}

export interface QueryAllOutboundTxRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllOutboundTxResponse {
  outboundTx: OutboundTx[];
  pagination: PageResponse | undefined;
}

export interface QueryGetValidatorsRequest {
  height: string;
}

export interface QueryGetValidatorsResponse {
  validators: Validators | undefined;
}

export interface QueryAllValidatorsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllValidatorsResponse {
  allValidators: Validators[];
  pagination: PageResponse | undefined;
}

export interface QueryGetQuotaRequest {
  queryLength: number;
  pagination: PageRequest | undefined;
}

export interface QueryGetQuotaResponse {
  coinQuotaResponse: CoinsQuota | undefined;
}

/** this line is used by starport scaffolding # 3 */
export interface QueryGetIssueTokenRequest {
  index: string;
}

export interface QueryGetIssueTokenResponse {
  IssueToken: IssueToken | undefined;
}

export interface QueryAllIssueTokenRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllIssueTokenResponse {
  IssueToken: IssueToken[];
  pagination: PageResponse | undefined;
}

export interface QueryGetCreatePoolRequest {
  index: string;
}

export interface QueryGetCreatePoolResponse {
  CreatePool: PoolProposal | undefined;
}

export interface PoolInfo {
  BlockHeight: string;
  CreatePool: PoolProposal | undefined;
}

export interface QueryLastPoolResponse {
  pools: PoolInfo[];
}

export interface QueryAllCreatePoolRequest {
  pagination: PageRequest | undefined;
}

export interface QueryLatestPoolRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllCreatePoolResponse {
  CreatePool: PoolProposal[];
  pagination: PageResponse | undefined;
}

function createBaseQueryPendingFeeRequest(): QueryPendingFeeRequest {
  return { pagination: undefined };
}

export const QueryPendingFeeRequest = {
  encode(message: QueryPendingFeeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryPendingFeeRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryPendingFeeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryPendingFeeRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryPendingFeeRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryPendingFeeRequest>, I>>(object: I): QueryPendingFeeRequest {
    const message = createBaseQueryPendingFeeRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryPendingFeeResponse(): QueryPendingFeeResponse {
  return { feecoin: [], pagination: undefined };
}

export const QueryPendingFeeResponse = {
  encode(message: QueryPendingFeeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.feecoin) {
      Coin.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryPendingFeeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryPendingFeeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.feecoin.push(Coin.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryPendingFeeResponse {
    return {
      feecoin: Array.isArray(object?.feecoin) ? object.feecoin.map((e: any) => Coin.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryPendingFeeResponse): unknown {
    const obj: any = {};
    if (message.feecoin) {
      obj.feecoin = message.feecoin.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.feecoin = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryPendingFeeResponse>, I>>(object: I): QueryPendingFeeResponse {
    const message = createBaseQueryPendingFeeResponse();
    message.feecoin = object.feecoin?.map((e) => Coin.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetOutboundTxRequest(): QueryGetOutboundTxRequest {
  return { requestID: "" };
}

export const QueryGetOutboundTxRequest = {
  encode(message: QueryGetOutboundTxRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.requestID !== "") {
      writer.uint32(10).string(message.requestID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetOutboundTxRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetOutboundTxRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.requestID = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOutboundTxRequest {
    return { requestID: isSet(object.requestID) ? String(object.requestID) : "" };
  },

  toJSON(message: QueryGetOutboundTxRequest): unknown {
    const obj: any = {};
    message.requestID !== undefined && (obj.requestID = message.requestID);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetOutboundTxRequest>, I>>(object: I): QueryGetOutboundTxRequest {
    const message = createBaseQueryGetOutboundTxRequest();
    message.requestID = object.requestID ?? "";
    return message;
  },
};

function createBaseQueryGetOutboundTxResponse(): QueryGetOutboundTxResponse {
  return { outboundTx: undefined };
}

export const QueryGetOutboundTxResponse = {
  encode(message: QueryGetOutboundTxResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.outboundTx !== undefined) {
      OutboundTx.encode(message.outboundTx, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetOutboundTxResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetOutboundTxResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.outboundTx = OutboundTx.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOutboundTxResponse {
    return { outboundTx: isSet(object.outboundTx) ? OutboundTx.fromJSON(object.outboundTx) : undefined };
  },

  toJSON(message: QueryGetOutboundTxResponse): unknown {
    const obj: any = {};
    message.outboundTx !== undefined
      && (obj.outboundTx = message.outboundTx ? OutboundTx.toJSON(message.outboundTx) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetOutboundTxResponse>, I>>(object: I): QueryGetOutboundTxResponse {
    const message = createBaseQueryGetOutboundTxResponse();
    message.outboundTx = (object.outboundTx !== undefined && object.outboundTx !== null)
      ? OutboundTx.fromPartial(object.outboundTx)
      : undefined;
    return message;
  },
};

function createBaseQueryAllOutboundTxRequest(): QueryAllOutboundTxRequest {
  return { pagination: undefined };
}

export const QueryAllOutboundTxRequest = {
  encode(message: QueryAllOutboundTxRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllOutboundTxRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllOutboundTxRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllOutboundTxRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllOutboundTxRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllOutboundTxRequest>, I>>(object: I): QueryAllOutboundTxRequest {
    const message = createBaseQueryAllOutboundTxRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllOutboundTxResponse(): QueryAllOutboundTxResponse {
  return { outboundTx: [], pagination: undefined };
}

export const QueryAllOutboundTxResponse = {
  encode(message: QueryAllOutboundTxResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.outboundTx) {
      OutboundTx.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllOutboundTxResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllOutboundTxResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.outboundTx.push(OutboundTx.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllOutboundTxResponse {
    return {
      outboundTx: Array.isArray(object?.outboundTx) ? object.outboundTx.map((e: any) => OutboundTx.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllOutboundTxResponse): unknown {
    const obj: any = {};
    if (message.outboundTx) {
      obj.outboundTx = message.outboundTx.map((e) => e ? OutboundTx.toJSON(e) : undefined);
    } else {
      obj.outboundTx = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllOutboundTxResponse>, I>>(object: I): QueryAllOutboundTxResponse {
    const message = createBaseQueryAllOutboundTxResponse();
    message.outboundTx = object.outboundTx?.map((e) => OutboundTx.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetValidatorsRequest(): QueryGetValidatorsRequest {
  return { height: "" };
}

export const QueryGetValidatorsRequest = {
  encode(message: QueryGetValidatorsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.height !== "") {
      writer.uint32(10).string(message.height);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetValidatorsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetValidatorsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.height = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetValidatorsRequest {
    return { height: isSet(object.height) ? String(object.height) : "" };
  },

  toJSON(message: QueryGetValidatorsRequest): unknown {
    const obj: any = {};
    message.height !== undefined && (obj.height = message.height);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetValidatorsRequest>, I>>(object: I): QueryGetValidatorsRequest {
    const message = createBaseQueryGetValidatorsRequest();
    message.height = object.height ?? "";
    return message;
  },
};

function createBaseQueryGetValidatorsResponse(): QueryGetValidatorsResponse {
  return { validators: undefined };
}

export const QueryGetValidatorsResponse = {
  encode(message: QueryGetValidatorsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.validators !== undefined) {
      Validators.encode(message.validators, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetValidatorsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetValidatorsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.validators = Validators.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetValidatorsResponse {
    return { validators: isSet(object.validators) ? Validators.fromJSON(object.validators) : undefined };
  },

  toJSON(message: QueryGetValidatorsResponse): unknown {
    const obj: any = {};
    message.validators !== undefined
      && (obj.validators = message.validators ? Validators.toJSON(message.validators) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetValidatorsResponse>, I>>(object: I): QueryGetValidatorsResponse {
    const message = createBaseQueryGetValidatorsResponse();
    message.validators = (object.validators !== undefined && object.validators !== null)
      ? Validators.fromPartial(object.validators)
      : undefined;
    return message;
  },
};

function createBaseQueryAllValidatorsRequest(): QueryAllValidatorsRequest {
  return { pagination: undefined };
}

export const QueryAllValidatorsRequest = {
  encode(message: QueryAllValidatorsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllValidatorsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllValidatorsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllValidatorsRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllValidatorsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllValidatorsRequest>, I>>(object: I): QueryAllValidatorsRequest {
    const message = createBaseQueryAllValidatorsRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllValidatorsResponse(): QueryAllValidatorsResponse {
  return { allValidators: [], pagination: undefined };
}

export const QueryAllValidatorsResponse = {
  encode(message: QueryAllValidatorsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.allValidators) {
      Validators.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllValidatorsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllValidatorsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.allValidators.push(Validators.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllValidatorsResponse {
    return {
      allValidators: Array.isArray(object?.allValidators)
        ? object.allValidators.map((e: any) => Validators.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllValidatorsResponse): unknown {
    const obj: any = {};
    if (message.allValidators) {
      obj.allValidators = message.allValidators.map((e) => e ? Validators.toJSON(e) : undefined);
    } else {
      obj.allValidators = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllValidatorsResponse>, I>>(object: I): QueryAllValidatorsResponse {
    const message = createBaseQueryAllValidatorsResponse();
    message.allValidators = object.allValidators?.map((e) => Validators.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetQuotaRequest(): QueryGetQuotaRequest {
  return { queryLength: 0, pagination: undefined };
}

export const QueryGetQuotaRequest = {
  encode(message: QueryGetQuotaRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.queryLength !== 0) {
      writer.uint32(8).int32(message.queryLength);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetQuotaRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetQuotaRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.queryLength = reader.int32();
          break;
        case 2:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetQuotaRequest {
    return {
      queryLength: isSet(object.queryLength) ? Number(object.queryLength) : 0,
      pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryGetQuotaRequest): unknown {
    const obj: any = {};
    message.queryLength !== undefined && (obj.queryLength = Math.round(message.queryLength));
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetQuotaRequest>, I>>(object: I): QueryGetQuotaRequest {
    const message = createBaseQueryGetQuotaRequest();
    message.queryLength = object.queryLength ?? 0;
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetQuotaResponse(): QueryGetQuotaResponse {
  return { coinQuotaResponse: undefined };
}

export const QueryGetQuotaResponse = {
  encode(message: QueryGetQuotaResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.coinQuotaResponse !== undefined) {
      CoinsQuota.encode(message.coinQuotaResponse, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetQuotaResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetQuotaResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.coinQuotaResponse = CoinsQuota.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetQuotaResponse {
    return {
      coinQuotaResponse: isSet(object.coinQuotaResponse) ? CoinsQuota.fromJSON(object.coinQuotaResponse) : undefined,
    };
  },

  toJSON(message: QueryGetQuotaResponse): unknown {
    const obj: any = {};
    message.coinQuotaResponse !== undefined
      && (obj.coinQuotaResponse = message.coinQuotaResponse ? CoinsQuota.toJSON(message.coinQuotaResponse) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetQuotaResponse>, I>>(object: I): QueryGetQuotaResponse {
    const message = createBaseQueryGetQuotaResponse();
    message.coinQuotaResponse = (object.coinQuotaResponse !== undefined && object.coinQuotaResponse !== null)
      ? CoinsQuota.fromPartial(object.coinQuotaResponse)
      : undefined;
    return message;
  },
};

function createBaseQueryGetIssueTokenRequest(): QueryGetIssueTokenRequest {
  return { index: "" };
}

export const QueryGetIssueTokenRequest = {
  encode(message: QueryGetIssueTokenRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetIssueTokenRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetIssueTokenRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetIssueTokenRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetIssueTokenRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetIssueTokenRequest>, I>>(object: I): QueryGetIssueTokenRequest {
    const message = createBaseQueryGetIssueTokenRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetIssueTokenResponse(): QueryGetIssueTokenResponse {
  return { IssueToken: undefined };
}

export const QueryGetIssueTokenResponse = {
  encode(message: QueryGetIssueTokenResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.IssueToken !== undefined) {
      IssueToken.encode(message.IssueToken, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetIssueTokenResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetIssueTokenResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.IssueToken = IssueToken.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetIssueTokenResponse {
    return { IssueToken: isSet(object.IssueToken) ? IssueToken.fromJSON(object.IssueToken) : undefined };
  },

  toJSON(message: QueryGetIssueTokenResponse): unknown {
    const obj: any = {};
    message.IssueToken !== undefined
      && (obj.IssueToken = message.IssueToken ? IssueToken.toJSON(message.IssueToken) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetIssueTokenResponse>, I>>(object: I): QueryGetIssueTokenResponse {
    const message = createBaseQueryGetIssueTokenResponse();
    message.IssueToken = (object.IssueToken !== undefined && object.IssueToken !== null)
      ? IssueToken.fromPartial(object.IssueToken)
      : undefined;
    return message;
  },
};

function createBaseQueryAllIssueTokenRequest(): QueryAllIssueTokenRequest {
  return { pagination: undefined };
}

export const QueryAllIssueTokenRequest = {
  encode(message: QueryAllIssueTokenRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllIssueTokenRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllIssueTokenRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllIssueTokenRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllIssueTokenRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllIssueTokenRequest>, I>>(object: I): QueryAllIssueTokenRequest {
    const message = createBaseQueryAllIssueTokenRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllIssueTokenResponse(): QueryAllIssueTokenResponse {
  return { IssueToken: [], pagination: undefined };
}

export const QueryAllIssueTokenResponse = {
  encode(message: QueryAllIssueTokenResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.IssueToken) {
      IssueToken.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllIssueTokenResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllIssueTokenResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.IssueToken.push(IssueToken.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllIssueTokenResponse {
    return {
      IssueToken: Array.isArray(object?.IssueToken) ? object.IssueToken.map((e: any) => IssueToken.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllIssueTokenResponse): unknown {
    const obj: any = {};
    if (message.IssueToken) {
      obj.IssueToken = message.IssueToken.map((e) => e ? IssueToken.toJSON(e) : undefined);
    } else {
      obj.IssueToken = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllIssueTokenResponse>, I>>(object: I): QueryAllIssueTokenResponse {
    const message = createBaseQueryAllIssueTokenResponse();
    message.IssueToken = object.IssueToken?.map((e) => IssueToken.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetCreatePoolRequest(): QueryGetCreatePoolRequest {
  return { index: "" };
}

export const QueryGetCreatePoolRequest = {
  encode(message: QueryGetCreatePoolRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetCreatePoolRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetCreatePoolRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetCreatePoolRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetCreatePoolRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetCreatePoolRequest>, I>>(object: I): QueryGetCreatePoolRequest {
    const message = createBaseQueryGetCreatePoolRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetCreatePoolResponse(): QueryGetCreatePoolResponse {
  return { CreatePool: undefined };
}

export const QueryGetCreatePoolResponse = {
  encode(message: QueryGetCreatePoolResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.CreatePool !== undefined) {
      PoolProposal.encode(message.CreatePool, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetCreatePoolResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetCreatePoolResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.CreatePool = PoolProposal.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetCreatePoolResponse {
    return { CreatePool: isSet(object.CreatePool) ? PoolProposal.fromJSON(object.CreatePool) : undefined };
  },

  toJSON(message: QueryGetCreatePoolResponse): unknown {
    const obj: any = {};
    message.CreatePool !== undefined
      && (obj.CreatePool = message.CreatePool ? PoolProposal.toJSON(message.CreatePool) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetCreatePoolResponse>, I>>(object: I): QueryGetCreatePoolResponse {
    const message = createBaseQueryGetCreatePoolResponse();
    message.CreatePool = (object.CreatePool !== undefined && object.CreatePool !== null)
      ? PoolProposal.fromPartial(object.CreatePool)
      : undefined;
    return message;
  },
};

function createBasePoolInfo(): PoolInfo {
  return { BlockHeight: "", CreatePool: undefined };
}

export const PoolInfo = {
  encode(message: PoolInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.BlockHeight !== "") {
      writer.uint32(10).string(message.BlockHeight);
    }
    if (message.CreatePool !== undefined) {
      PoolProposal.encode(message.CreatePool, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PoolInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePoolInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.BlockHeight = reader.string();
          break;
        case 2:
          message.CreatePool = PoolProposal.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PoolInfo {
    return {
      BlockHeight: isSet(object.BlockHeight) ? String(object.BlockHeight) : "",
      CreatePool: isSet(object.CreatePool) ? PoolProposal.fromJSON(object.CreatePool) : undefined,
    };
  },

  toJSON(message: PoolInfo): unknown {
    const obj: any = {};
    message.BlockHeight !== undefined && (obj.BlockHeight = message.BlockHeight);
    message.CreatePool !== undefined
      && (obj.CreatePool = message.CreatePool ? PoolProposal.toJSON(message.CreatePool) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PoolInfo>, I>>(object: I): PoolInfo {
    const message = createBasePoolInfo();
    message.BlockHeight = object.BlockHeight ?? "";
    message.CreatePool = (object.CreatePool !== undefined && object.CreatePool !== null)
      ? PoolProposal.fromPartial(object.CreatePool)
      : undefined;
    return message;
  },
};

function createBaseQueryLastPoolResponse(): QueryLastPoolResponse {
  return { pools: [] };
}

export const QueryLastPoolResponse = {
  encode(message: QueryLastPoolResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.pools) {
      PoolInfo.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryLastPoolResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryLastPoolResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pools.push(PoolInfo.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryLastPoolResponse {
    return { pools: Array.isArray(object?.pools) ? object.pools.map((e: any) => PoolInfo.fromJSON(e)) : [] };
  },

  toJSON(message: QueryLastPoolResponse): unknown {
    const obj: any = {};
    if (message.pools) {
      obj.pools = message.pools.map((e) => e ? PoolInfo.toJSON(e) : undefined);
    } else {
      obj.pools = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryLastPoolResponse>, I>>(object: I): QueryLastPoolResponse {
    const message = createBaseQueryLastPoolResponse();
    message.pools = object.pools?.map((e) => PoolInfo.fromPartial(e)) || [];
    return message;
  },
};

function createBaseQueryAllCreatePoolRequest(): QueryAllCreatePoolRequest {
  return { pagination: undefined };
}

export const QueryAllCreatePoolRequest = {
  encode(message: QueryAllCreatePoolRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllCreatePoolRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllCreatePoolRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllCreatePoolRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllCreatePoolRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllCreatePoolRequest>, I>>(object: I): QueryAllCreatePoolRequest {
    const message = createBaseQueryAllCreatePoolRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryLatestPoolRequest(): QueryLatestPoolRequest {
  return { pagination: undefined };
}

export const QueryLatestPoolRequest = {
  encode(message: QueryLatestPoolRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryLatestPoolRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryLatestPoolRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryLatestPoolRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryLatestPoolRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryLatestPoolRequest>, I>>(object: I): QueryLatestPoolRequest {
    const message = createBaseQueryLatestPoolRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllCreatePoolResponse(): QueryAllCreatePoolResponse {
  return { CreatePool: [], pagination: undefined };
}

export const QueryAllCreatePoolResponse = {
  encode(message: QueryAllCreatePoolResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.CreatePool) {
      PoolProposal.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllCreatePoolResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllCreatePoolResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.CreatePool.push(PoolProposal.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllCreatePoolResponse {
    return {
      CreatePool: Array.isArray(object?.CreatePool) ? object.CreatePool.map((e: any) => PoolProposal.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllCreatePoolResponse): unknown {
    const obj: any = {};
    if (message.CreatePool) {
      obj.CreatePool = message.CreatePool.map((e) => e ? PoolProposal.toJSON(e) : undefined);
    } else {
      obj.CreatePool = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllCreatePoolResponse>, I>>(object: I): QueryAllCreatePoolResponse {
    const message = createBaseQueryAllCreatePoolResponse();
    message.CreatePool = object.CreatePool?.map((e) => PoolProposal.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a OutboundTx by index. */
  OutboundTx(request: QueryGetOutboundTxRequest): Promise<QueryGetOutboundTxResponse>;
  /** Queries a list of OutboundTx items. */
  OutboundTxAll(request: QueryAllOutboundTxRequest): Promise<QueryAllOutboundTxResponse>;
  /** Queries a list of GetValidators items. */
  GetValidators(request: QueryGetValidatorsRequest): Promise<QueryGetValidatorsResponse>;
  /** Queries a list of GetValidators items. */
  GetAllValidators(request: QueryAllValidatorsRequest): Promise<QueryAllValidatorsResponse>;
  /** Queries a list of GetQuota items. */
  GetQuota(request: QueryGetQuotaRequest): Promise<QueryGetQuotaResponse>;
  /** Queries a issueToken by index. */
  IssueToken(request: QueryGetIssueTokenRequest): Promise<QueryGetIssueTokenResponse>;
  /** Queries a list of issueToken items. */
  IssueTokenAll(request: QueryAllIssueTokenRequest): Promise<QueryAllIssueTokenResponse>;
  /** Queries a createPool by index. */
  CreatePool(request: QueryGetCreatePoolRequest): Promise<QueryGetCreatePoolResponse>;
  /** Queries a list of createPool items. */
  CreatePoolAll(request: QueryAllCreatePoolRequest): Promise<QueryAllCreatePoolResponse>;
  /** Queries a createPool by index. */
  GetLastPool(request: QueryLatestPoolRequest): Promise<QueryLastPoolResponse>;
  /** Queries the pending fee */
  GetPendingFee(request: QueryPendingFeeRequest): Promise<QueryPendingFeeResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.OutboundTx = this.OutboundTx.bind(this);
    this.OutboundTxAll = this.OutboundTxAll.bind(this);
    this.GetValidators = this.GetValidators.bind(this);
    this.GetAllValidators = this.GetAllValidators.bind(this);
    this.GetQuota = this.GetQuota.bind(this);
    this.IssueToken = this.IssueToken.bind(this);
    this.IssueTokenAll = this.IssueTokenAll.bind(this);
    this.CreatePool = this.CreatePool.bind(this);
    this.CreatePoolAll = this.CreatePoolAll.bind(this);
    this.GetLastPool = this.GetLastPool.bind(this);
    this.GetPendingFee = this.GetPendingFee.bind(this);
  }
  OutboundTx(request: QueryGetOutboundTxRequest): Promise<QueryGetOutboundTxResponse> {
    const data = QueryGetOutboundTxRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Query", "OutboundTx", data);
    return promise.then((data) => QueryGetOutboundTxResponse.decode(new _m0.Reader(data)));
  }

  OutboundTxAll(request: QueryAllOutboundTxRequest): Promise<QueryAllOutboundTxResponse> {
    const data = QueryAllOutboundTxRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Query", "OutboundTxAll", data);
    return promise.then((data) => QueryAllOutboundTxResponse.decode(new _m0.Reader(data)));
  }

  GetValidators(request: QueryGetValidatorsRequest): Promise<QueryGetValidatorsResponse> {
    const data = QueryGetValidatorsRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Query", "GetValidators", data);
    return promise.then((data) => QueryGetValidatorsResponse.decode(new _m0.Reader(data)));
  }

  GetAllValidators(request: QueryAllValidatorsRequest): Promise<QueryAllValidatorsResponse> {
    const data = QueryAllValidatorsRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Query", "GetAllValidators", data);
    return promise.then((data) => QueryAllValidatorsResponse.decode(new _m0.Reader(data)));
  }

  GetQuota(request: QueryGetQuotaRequest): Promise<QueryGetQuotaResponse> {
    const data = QueryGetQuotaRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Query", "GetQuota", data);
    return promise.then((data) => QueryGetQuotaResponse.decode(new _m0.Reader(data)));
  }

  IssueToken(request: QueryGetIssueTokenRequest): Promise<QueryGetIssueTokenResponse> {
    const data = QueryGetIssueTokenRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Query", "IssueToken", data);
    return promise.then((data) => QueryGetIssueTokenResponse.decode(new _m0.Reader(data)));
  }

  IssueTokenAll(request: QueryAllIssueTokenRequest): Promise<QueryAllIssueTokenResponse> {
    const data = QueryAllIssueTokenRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Query", "IssueTokenAll", data);
    return promise.then((data) => QueryAllIssueTokenResponse.decode(new _m0.Reader(data)));
  }

  CreatePool(request: QueryGetCreatePoolRequest): Promise<QueryGetCreatePoolResponse> {
    const data = QueryGetCreatePoolRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Query", "CreatePool", data);
    return promise.then((data) => QueryGetCreatePoolResponse.decode(new _m0.Reader(data)));
  }

  CreatePoolAll(request: QueryAllCreatePoolRequest): Promise<QueryAllCreatePoolResponse> {
    const data = QueryAllCreatePoolRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Query", "CreatePoolAll", data);
    return promise.then((data) => QueryAllCreatePoolResponse.decode(new _m0.Reader(data)));
  }

  GetLastPool(request: QueryLatestPoolRequest): Promise<QueryLastPoolResponse> {
    const data = QueryLatestPoolRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Query", "GetLastPool", data);
    return promise.then((data) => QueryLastPoolResponse.decode(new _m0.Reader(data)));
  }

  GetPendingFee(request: QueryPendingFeeRequest): Promise<QueryPendingFeeResponse> {
    const data = QueryPendingFeeRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Query", "GetPendingFee", data);
    return promise.then((data) => QueryPendingFeeResponse.decode(new _m0.Reader(data)));
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
