import { PageRequest, PageRequestSDKType, PageResponse, PageResponseSDKType } from "../../cosmos/base/query/v1beta1/pagination";
import { Coin, CoinSDKType } from "../../cosmos/base/v1beta1/coin";
import { OutboundTx, OutboundTxSDKType } from "./outbound_tx";
import { Validators, ValidatorsSDKType } from "./staking";
import { CoinsQuota, CoinsQuotaSDKType } from "./quota";
import { IssueToken, IssueTokenSDKType } from "./issue_token";
import { PoolProposal, PoolProposalSDKType } from "./create_pool";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
export interface QueryPendingFeeRequest {
  pagination?: PageRequest;
}
export interface QueryPendingFeeRequestSDKType {
  pagination?: PageRequestSDKType;
}
export interface QueryPendingFeeResponse {
  feecoin: Coin[];
  pagination?: PageResponse;
}
export interface QueryPendingFeeResponseSDKType {
  feecoin: CoinSDKType[];
  pagination?: PageResponseSDKType;
}
export interface QueryGetOutboundTxRequest {
  requestID: string;
}
export interface QueryGetOutboundTxRequestSDKType {
  requestID: string;
}
export interface QueryGetOutboundTxResponse {
  outboundTx?: OutboundTx;
}
export interface QueryGetOutboundTxResponseSDKType {
  outboundTx?: OutboundTxSDKType;
}
export interface QueryAllOutboundTxRequest {
  pagination?: PageRequest;
}
export interface QueryAllOutboundTxRequestSDKType {
  pagination?: PageRequestSDKType;
}
export interface QueryAllOutboundTxResponse {
  outboundTx: OutboundTx[];
  pagination?: PageResponse;
}
export interface QueryAllOutboundTxResponseSDKType {
  outboundTx: OutboundTxSDKType[];
  pagination?: PageResponseSDKType;
}
export interface QueryGetValidatorsRequest {
  height: string;
}
export interface QueryGetValidatorsRequestSDKType {
  height: string;
}
export interface QueryGetValidatorsResponse {
  validators?: Validators;
}
export interface QueryGetValidatorsResponseSDKType {
  validators?: ValidatorsSDKType;
}
export interface QueryAllValidatorsRequest {
  pagination?: PageRequest;
}
export interface QueryAllValidatorsRequestSDKType {
  pagination?: PageRequestSDKType;
}
export interface QueryAllValidatorsResponse {
  allValidators: Validators[];
  pagination?: PageResponse;
}
export interface QueryAllValidatorsResponseSDKType {
  all_validators: ValidatorsSDKType[];
  pagination?: PageResponseSDKType;
}
export interface QueryGetQuotaRequest {
  queryLength: number;
  pagination?: PageRequest;
}
export interface QueryGetQuotaRequestSDKType {
  query_length: number;
  pagination?: PageRequestSDKType;
}
export interface QueryGetQuotaResponse {
  coinQuotaResponse?: CoinsQuota;
}
export interface QueryGetQuotaResponseSDKType {
  coinQuotaResponse?: CoinsQuotaSDKType;
}
/** this line is used by starport scaffolding # 3 */

export interface QueryGetIssueTokenRequest {
  index: string;
}
/** this line is used by starport scaffolding # 3 */

export interface QueryGetIssueTokenRequestSDKType {
  index: string;
}
export interface QueryGetIssueTokenResponse {
  IssueToken?: IssueToken;
}
export interface QueryGetIssueTokenResponseSDKType {
  IssueToken?: IssueTokenSDKType;
}
export interface QueryAllIssueTokenRequest {
  pagination?: PageRequest;
}
export interface QueryAllIssueTokenRequestSDKType {
  pagination?: PageRequestSDKType;
}
export interface QueryAllIssueTokenResponse {
  IssueToken: IssueToken[];
  pagination?: PageResponse;
}
export interface QueryAllIssueTokenResponseSDKType {
  IssueToken: IssueTokenSDKType[];
  pagination?: PageResponseSDKType;
}
export interface QueryGetCreatePoolRequest {
  index: string;
}
export interface QueryGetCreatePoolRequestSDKType {
  index: string;
}
export interface QueryGetCreatePoolResponse {
  CreatePool?: PoolProposal;
}
export interface QueryGetCreatePoolResponseSDKType {
  CreatePool?: PoolProposalSDKType;
}
export interface PoolInfo {
  BlockHeight: string;
  CreatePool?: PoolProposal;
}
export interface PoolInfoSDKType {
  BlockHeight: string;
  CreatePool?: PoolProposalSDKType;
}
export interface QueryLastPoolResponse {
  pools: PoolInfo[];
}
export interface QueryLastPoolResponseSDKType {
  pools: PoolInfoSDKType[];
}
export interface QueryAllCreatePoolRequest {
  pagination?: PageRequest;
}
export interface QueryAllCreatePoolRequestSDKType {
  pagination?: PageRequestSDKType;
}
export interface QueryLatestPoolRequest {
  pagination?: PageRequest;
}
export interface QueryLatestPoolRequestSDKType {
  pagination?: PageRequestSDKType;
}
export interface QueryAllCreatePoolResponse {
  CreatePool: PoolProposal[];
  pagination?: PageResponse;
}
export interface QueryAllCreatePoolResponseSDKType {
  CreatePool: PoolProposalSDKType[];
  pagination?: PageResponseSDKType;
}

function createBaseQueryPendingFeeRequest(): QueryPendingFeeRequest {
  return {
    pagination: undefined
  };
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

  fromPartial(object: DeepPartial<QueryPendingFeeRequest>): QueryPendingFeeRequest {
    const message = createBaseQueryPendingFeeRequest();
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageRequest.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryPendingFeeResponse(): QueryPendingFeeResponse {
  return {
    feecoin: [],
    pagination: undefined
  };
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

  fromPartial(object: DeepPartial<QueryPendingFeeResponse>): QueryPendingFeeResponse {
    const message = createBaseQueryPendingFeeResponse();
    message.feecoin = object.feecoin?.map(e => Coin.fromPartial(e)) || [];
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageResponse.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryGetOutboundTxRequest(): QueryGetOutboundTxRequest {
  return {
    requestID: ""
  };
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

  fromPartial(object: DeepPartial<QueryGetOutboundTxRequest>): QueryGetOutboundTxRequest {
    const message = createBaseQueryGetOutboundTxRequest();
    message.requestID = object.requestID ?? "";
    return message;
  }

};

function createBaseQueryGetOutboundTxResponse(): QueryGetOutboundTxResponse {
  return {
    outboundTx: undefined
  };
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

  fromPartial(object: DeepPartial<QueryGetOutboundTxResponse>): QueryGetOutboundTxResponse {
    const message = createBaseQueryGetOutboundTxResponse();
    message.outboundTx = object.outboundTx !== undefined && object.outboundTx !== null ? OutboundTx.fromPartial(object.outboundTx) : undefined;
    return message;
  }

};

function createBaseQueryAllOutboundTxRequest(): QueryAllOutboundTxRequest {
  return {
    pagination: undefined
  };
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

  fromPartial(object: DeepPartial<QueryAllOutboundTxRequest>): QueryAllOutboundTxRequest {
    const message = createBaseQueryAllOutboundTxRequest();
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageRequest.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryAllOutboundTxResponse(): QueryAllOutboundTxResponse {
  return {
    outboundTx: [],
    pagination: undefined
  };
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

  fromPartial(object: DeepPartial<QueryAllOutboundTxResponse>): QueryAllOutboundTxResponse {
    const message = createBaseQueryAllOutboundTxResponse();
    message.outboundTx = object.outboundTx?.map(e => OutboundTx.fromPartial(e)) || [];
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageResponse.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryGetValidatorsRequest(): QueryGetValidatorsRequest {
  return {
    height: ""
  };
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

  fromPartial(object: DeepPartial<QueryGetValidatorsRequest>): QueryGetValidatorsRequest {
    const message = createBaseQueryGetValidatorsRequest();
    message.height = object.height ?? "";
    return message;
  }

};

function createBaseQueryGetValidatorsResponse(): QueryGetValidatorsResponse {
  return {
    validators: undefined
  };
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

  fromPartial(object: DeepPartial<QueryGetValidatorsResponse>): QueryGetValidatorsResponse {
    const message = createBaseQueryGetValidatorsResponse();
    message.validators = object.validators !== undefined && object.validators !== null ? Validators.fromPartial(object.validators) : undefined;
    return message;
  }

};

function createBaseQueryAllValidatorsRequest(): QueryAllValidatorsRequest {
  return {
    pagination: undefined
  };
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

  fromPartial(object: DeepPartial<QueryAllValidatorsRequest>): QueryAllValidatorsRequest {
    const message = createBaseQueryAllValidatorsRequest();
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageRequest.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryAllValidatorsResponse(): QueryAllValidatorsResponse {
  return {
    allValidators: [],
    pagination: undefined
  };
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

  fromPartial(object: DeepPartial<QueryAllValidatorsResponse>): QueryAllValidatorsResponse {
    const message = createBaseQueryAllValidatorsResponse();
    message.allValidators = object.allValidators?.map(e => Validators.fromPartial(e)) || [];
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageResponse.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryGetQuotaRequest(): QueryGetQuotaRequest {
  return {
    queryLength: 0,
    pagination: undefined
  };
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

  fromPartial(object: DeepPartial<QueryGetQuotaRequest>): QueryGetQuotaRequest {
    const message = createBaseQueryGetQuotaRequest();
    message.queryLength = object.queryLength ?? 0;
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageRequest.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryGetQuotaResponse(): QueryGetQuotaResponse {
  return {
    coinQuotaResponse: undefined
  };
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

  fromPartial(object: DeepPartial<QueryGetQuotaResponse>): QueryGetQuotaResponse {
    const message = createBaseQueryGetQuotaResponse();
    message.coinQuotaResponse = object.coinQuotaResponse !== undefined && object.coinQuotaResponse !== null ? CoinsQuota.fromPartial(object.coinQuotaResponse) : undefined;
    return message;
  }

};

function createBaseQueryGetIssueTokenRequest(): QueryGetIssueTokenRequest {
  return {
    index: ""
  };
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

  fromPartial(object: DeepPartial<QueryGetIssueTokenRequest>): QueryGetIssueTokenRequest {
    const message = createBaseQueryGetIssueTokenRequest();
    message.index = object.index ?? "";
    return message;
  }

};

function createBaseQueryGetIssueTokenResponse(): QueryGetIssueTokenResponse {
  return {
    IssueToken: undefined
  };
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

  fromPartial(object: DeepPartial<QueryGetIssueTokenResponse>): QueryGetIssueTokenResponse {
    const message = createBaseQueryGetIssueTokenResponse();
    message.IssueToken = object.IssueToken !== undefined && object.IssueToken !== null ? IssueToken.fromPartial(object.IssueToken) : undefined;
    return message;
  }

};

function createBaseQueryAllIssueTokenRequest(): QueryAllIssueTokenRequest {
  return {
    pagination: undefined
  };
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

  fromPartial(object: DeepPartial<QueryAllIssueTokenRequest>): QueryAllIssueTokenRequest {
    const message = createBaseQueryAllIssueTokenRequest();
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageRequest.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryAllIssueTokenResponse(): QueryAllIssueTokenResponse {
  return {
    IssueToken: [],
    pagination: undefined
  };
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

  fromPartial(object: DeepPartial<QueryAllIssueTokenResponse>): QueryAllIssueTokenResponse {
    const message = createBaseQueryAllIssueTokenResponse();
    message.IssueToken = object.IssueToken?.map(e => IssueToken.fromPartial(e)) || [];
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageResponse.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryGetCreatePoolRequest(): QueryGetCreatePoolRequest {
  return {
    index: ""
  };
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

  fromPartial(object: DeepPartial<QueryGetCreatePoolRequest>): QueryGetCreatePoolRequest {
    const message = createBaseQueryGetCreatePoolRequest();
    message.index = object.index ?? "";
    return message;
  }

};

function createBaseQueryGetCreatePoolResponse(): QueryGetCreatePoolResponse {
  return {
    CreatePool: undefined
  };
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

  fromPartial(object: DeepPartial<QueryGetCreatePoolResponse>): QueryGetCreatePoolResponse {
    const message = createBaseQueryGetCreatePoolResponse();
    message.CreatePool = object.CreatePool !== undefined && object.CreatePool !== null ? PoolProposal.fromPartial(object.CreatePool) : undefined;
    return message;
  }

};

function createBasePoolInfo(): PoolInfo {
  return {
    BlockHeight: "",
    CreatePool: undefined
  };
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

  fromPartial(object: DeepPartial<PoolInfo>): PoolInfo {
    const message = createBasePoolInfo();
    message.BlockHeight = object.BlockHeight ?? "";
    message.CreatePool = object.CreatePool !== undefined && object.CreatePool !== null ? PoolProposal.fromPartial(object.CreatePool) : undefined;
    return message;
  }

};

function createBaseQueryLastPoolResponse(): QueryLastPoolResponse {
  return {
    pools: []
  };
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

  fromPartial(object: DeepPartial<QueryLastPoolResponse>): QueryLastPoolResponse {
    const message = createBaseQueryLastPoolResponse();
    message.pools = object.pools?.map(e => PoolInfo.fromPartial(e)) || [];
    return message;
  }

};

function createBaseQueryAllCreatePoolRequest(): QueryAllCreatePoolRequest {
  return {
    pagination: undefined
  };
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

  fromPartial(object: DeepPartial<QueryAllCreatePoolRequest>): QueryAllCreatePoolRequest {
    const message = createBaseQueryAllCreatePoolRequest();
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageRequest.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryLatestPoolRequest(): QueryLatestPoolRequest {
  return {
    pagination: undefined
  };
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

  fromPartial(object: DeepPartial<QueryLatestPoolRequest>): QueryLatestPoolRequest {
    const message = createBaseQueryLatestPoolRequest();
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageRequest.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryAllCreatePoolResponse(): QueryAllCreatePoolResponse {
  return {
    CreatePool: [],
    pagination: undefined
  };
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

  fromPartial(object: DeepPartial<QueryAllCreatePoolResponse>): QueryAllCreatePoolResponse {
    const message = createBaseQueryAllCreatePoolResponse();
    message.CreatePool = object.CreatePool?.map(e => PoolProposal.fromPartial(e)) || [];
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageResponse.fromPartial(object.pagination) : undefined;
    return message;
  }

};