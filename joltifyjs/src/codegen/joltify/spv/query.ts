import { PageRequest, PageRequestSDKType, PageResponse, PageResponseSDKType } from "../../cosmos/base/query/v1beta1/pagination";
import { Params, ParamsSDKType } from "./params";
import { PoolInfo, PoolInfoSDKType } from "./poolinfo";
import { DepositorInfo, DepositorInfoSDKType } from "./deposit";
import { Coin, CoinSDKType } from "../../cosmos/base/v1beta1/coin";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
/** QueryParamsRequest is request type for the Query/Params RPC method. */

export interface QueryParamsRequest {}
/** QueryParamsRequest is request type for the Query/Params RPC method. */

export interface QueryParamsRequestSDKType {}
/** QueryParamsResponse is response type for the Query/Params RPC method. */

export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: Params;
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */

export interface QueryParamsResponseSDKType {
  params?: ParamsSDKType;
}
export interface QueryListPoolsRequest {
  pagination?: PageRequest;
}
export interface QueryListPoolsRequestSDKType {
  pagination?: PageRequestSDKType;
}
export interface QueryListPoolsResponse {
  poolsInfo: PoolInfo[];
  pagination?: PageResponse;
}
export interface QueryListPoolsResponseSDKType {
  pools_info: PoolInfoSDKType[];
  pagination?: PageResponseSDKType;
}
export interface QueryQueryPoolRequest {
  poolIndex: string;
}
export interface QueryQueryPoolRequestSDKType {
  pool_index: string;
}
export interface QueryQueryPoolResponse {
  poolInfo?: PoolInfo;
}
export interface QueryQueryPoolResponseSDKType {
  pool_info?: PoolInfoSDKType;
}
export interface QueryDepositorRequest {
  walletAddress: string;
  depositPoolIndex: string;
}
export interface QueryDepositorRequestSDKType {
  walletAddress: string;
  deposit_pool_index: string;
}
export interface QueryDepositorResponse {
  depositor?: DepositorInfo;
}
export interface QueryDepositorResponseSDKType {
  depositor?: DepositorInfoSDKType;
}
export interface QueryAllowedPoolsRequest {
  walletAddress: string;
}
export interface QueryAllowedPoolsRequestSDKType {
  wallet_address: string;
}
export interface QueryAllowedPoolsResponse {
  poolsIndex: string[];
}
export interface QueryAllowedPoolsResponseSDKType {
  pools_index: string[];
}
export interface QueryClaimableInterestRequest {
  wallet: string;
  poolIndex: string;
}
export interface QueryClaimableInterestRequestSDKType {
  wallet: string;
  pool_index: string;
}
export interface QueryClaimableInterestResponse {
  claimableInterestAmount?: Coin;
}
export interface QueryClaimableInterestResponseSDKType {
  claimable_interest_amount?: CoinSDKType;
}
export interface QueryOutstandingInterestRequest {
  wallet: string;
  poolIndex: string;
}
export interface QueryOutstandingInterestRequestSDKType {
  wallet: string;
  pool_index: string;
}
export interface QueryOutstandingInterestResponse {
  amount: string;
}
export interface QueryOutstandingInterestResponseSDKType {
  amount: string;
}
export interface QuerywithdrawalPrincipalRequest {
  poolIndex: string;
  walletAddress: string;
}
export interface QuerywithdrawalPrincipalRequestSDKType {
  pool_index: string;
  walletAddress: string;
}
export interface QuerywithdrawalPrincipalResponse {
  amount: string;
}
export interface QuerywithdrawalPrincipalResponseSDKType {
  amount: string;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();

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

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  }

};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return {
    params: undefined
  };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = object.params !== undefined && object.params !== null ? Params.fromPartial(object.params) : undefined;
    return message;
  }

};

function createBaseQueryListPoolsRequest(): QueryListPoolsRequest {
  return {
    pagination: undefined
  };
}

export const QueryListPoolsRequest = {
  encode(message: QueryListPoolsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryListPoolsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryListPoolsRequest();

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

  fromPartial(object: DeepPartial<QueryListPoolsRequest>): QueryListPoolsRequest {
    const message = createBaseQueryListPoolsRequest();
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageRequest.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryListPoolsResponse(): QueryListPoolsResponse {
  return {
    poolsInfo: [],
    pagination: undefined
  };
}

export const QueryListPoolsResponse = {
  encode(message: QueryListPoolsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.poolsInfo) {
      PoolInfo.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryListPoolsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryListPoolsResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.poolsInfo.push(PoolInfo.decode(reader, reader.uint32()));
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

  fromPartial(object: DeepPartial<QueryListPoolsResponse>): QueryListPoolsResponse {
    const message = createBaseQueryListPoolsResponse();
    message.poolsInfo = object.poolsInfo?.map(e => PoolInfo.fromPartial(e)) || [];
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageResponse.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryQueryPoolRequest(): QueryQueryPoolRequest {
  return {
    poolIndex: ""
  };
}

export const QueryQueryPoolRequest = {
  encode(message: QueryQueryPoolRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.poolIndex !== "") {
      writer.uint32(10).string(message.poolIndex);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQueryPoolRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQueryPoolRequest();

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

  fromPartial(object: DeepPartial<QueryQueryPoolRequest>): QueryQueryPoolRequest {
    const message = createBaseQueryQueryPoolRequest();
    message.poolIndex = object.poolIndex ?? "";
    return message;
  }

};

function createBaseQueryQueryPoolResponse(): QueryQueryPoolResponse {
  return {
    poolInfo: undefined
  };
}

export const QueryQueryPoolResponse = {
  encode(message: QueryQueryPoolResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.poolInfo !== undefined) {
      PoolInfo.encode(message.poolInfo, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQueryPoolResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQueryPoolResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.poolInfo = PoolInfo.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryQueryPoolResponse>): QueryQueryPoolResponse {
    const message = createBaseQueryQueryPoolResponse();
    message.poolInfo = object.poolInfo !== undefined && object.poolInfo !== null ? PoolInfo.fromPartial(object.poolInfo) : undefined;
    return message;
  }

};

function createBaseQueryDepositorRequest(): QueryDepositorRequest {
  return {
    walletAddress: "",
    depositPoolIndex: ""
  };
}

export const QueryDepositorRequest = {
  encode(message: QueryDepositorRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.walletAddress !== "") {
      writer.uint32(10).string(message.walletAddress);
    }

    if (message.depositPoolIndex !== "") {
      writer.uint32(18).string(message.depositPoolIndex);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryDepositorRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryDepositorRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.walletAddress = reader.string();
          break;

        case 2:
          message.depositPoolIndex = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryDepositorRequest>): QueryDepositorRequest {
    const message = createBaseQueryDepositorRequest();
    message.walletAddress = object.walletAddress ?? "";
    message.depositPoolIndex = object.depositPoolIndex ?? "";
    return message;
  }

};

function createBaseQueryDepositorResponse(): QueryDepositorResponse {
  return {
    depositor: undefined
  };
}

export const QueryDepositorResponse = {
  encode(message: QueryDepositorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.depositor !== undefined) {
      DepositorInfo.encode(message.depositor, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryDepositorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryDepositorResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.depositor = DepositorInfo.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryDepositorResponse>): QueryDepositorResponse {
    const message = createBaseQueryDepositorResponse();
    message.depositor = object.depositor !== undefined && object.depositor !== null ? DepositorInfo.fromPartial(object.depositor) : undefined;
    return message;
  }

};

function createBaseQueryAllowedPoolsRequest(): QueryAllowedPoolsRequest {
  return {
    walletAddress: ""
  };
}

export const QueryAllowedPoolsRequest = {
  encode(message: QueryAllowedPoolsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.walletAddress !== "") {
      writer.uint32(10).string(message.walletAddress);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllowedPoolsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllowedPoolsRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.walletAddress = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryAllowedPoolsRequest>): QueryAllowedPoolsRequest {
    const message = createBaseQueryAllowedPoolsRequest();
    message.walletAddress = object.walletAddress ?? "";
    return message;
  }

};

function createBaseQueryAllowedPoolsResponse(): QueryAllowedPoolsResponse {
  return {
    poolsIndex: []
  };
}

export const QueryAllowedPoolsResponse = {
  encode(message: QueryAllowedPoolsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.poolsIndex) {
      writer.uint32(10).string(v!);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllowedPoolsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllowedPoolsResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.poolsIndex.push(reader.string());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryAllowedPoolsResponse>): QueryAllowedPoolsResponse {
    const message = createBaseQueryAllowedPoolsResponse();
    message.poolsIndex = object.poolsIndex?.map(e => e) || [];
    return message;
  }

};

function createBaseQueryClaimableInterestRequest(): QueryClaimableInterestRequest {
  return {
    wallet: "",
    poolIndex: ""
  };
}

export const QueryClaimableInterestRequest = {
  encode(message: QueryClaimableInterestRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.wallet !== "") {
      writer.uint32(10).string(message.wallet);
    }

    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryClaimableInterestRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryClaimableInterestRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.wallet = reader.string();
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

  fromPartial(object: DeepPartial<QueryClaimableInterestRequest>): QueryClaimableInterestRequest {
    const message = createBaseQueryClaimableInterestRequest();
    message.wallet = object.wallet ?? "";
    message.poolIndex = object.poolIndex ?? "";
    return message;
  }

};

function createBaseQueryClaimableInterestResponse(): QueryClaimableInterestResponse {
  return {
    claimableInterestAmount: undefined
  };
}

export const QueryClaimableInterestResponse = {
  encode(message: QueryClaimableInterestResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.claimableInterestAmount !== undefined) {
      Coin.encode(message.claimableInterestAmount, writer.uint32(50).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryClaimableInterestResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryClaimableInterestResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 6:
          message.claimableInterestAmount = Coin.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryClaimableInterestResponse>): QueryClaimableInterestResponse {
    const message = createBaseQueryClaimableInterestResponse();
    message.claimableInterestAmount = object.claimableInterestAmount !== undefined && object.claimableInterestAmount !== null ? Coin.fromPartial(object.claimableInterestAmount) : undefined;
    return message;
  }

};

function createBaseQueryOutstandingInterestRequest(): QueryOutstandingInterestRequest {
  return {
    wallet: "",
    poolIndex: ""
  };
}

export const QueryOutstandingInterestRequest = {
  encode(message: QueryOutstandingInterestRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.wallet !== "") {
      writer.uint32(10).string(message.wallet);
    }

    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryOutstandingInterestRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryOutstandingInterestRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.wallet = reader.string();
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

  fromPartial(object: DeepPartial<QueryOutstandingInterestRequest>): QueryOutstandingInterestRequest {
    const message = createBaseQueryOutstandingInterestRequest();
    message.wallet = object.wallet ?? "";
    message.poolIndex = object.poolIndex ?? "";
    return message;
  }

};

function createBaseQueryOutstandingInterestResponse(): QueryOutstandingInterestResponse {
  return {
    amount: ""
  };
}

export const QueryOutstandingInterestResponse = {
  encode(message: QueryOutstandingInterestResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.amount !== "") {
      writer.uint32(10).string(message.amount);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryOutstandingInterestResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryOutstandingInterestResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.amount = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryOutstandingInterestResponse>): QueryOutstandingInterestResponse {
    const message = createBaseQueryOutstandingInterestResponse();
    message.amount = object.amount ?? "";
    return message;
  }

};

function createBaseQuerywithdrawalPrincipalRequest(): QuerywithdrawalPrincipalRequest {
  return {
    poolIndex: "",
    walletAddress: ""
  };
}

export const QuerywithdrawalPrincipalRequest = {
  encode(message: QuerywithdrawalPrincipalRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.poolIndex !== "") {
      writer.uint32(10).string(message.poolIndex);
    }

    if (message.walletAddress !== "") {
      writer.uint32(18).string(message.walletAddress);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QuerywithdrawalPrincipalRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQuerywithdrawalPrincipalRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.poolIndex = reader.string();
          break;

        case 2:
          message.walletAddress = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QuerywithdrawalPrincipalRequest>): QuerywithdrawalPrincipalRequest {
    const message = createBaseQuerywithdrawalPrincipalRequest();
    message.poolIndex = object.poolIndex ?? "";
    message.walletAddress = object.walletAddress ?? "";
    return message;
  }

};

function createBaseQuerywithdrawalPrincipalResponse(): QuerywithdrawalPrincipalResponse {
  return {
    amount: ""
  };
}

export const QuerywithdrawalPrincipalResponse = {
  encode(message: QuerywithdrawalPrincipalResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.amount !== "") {
      writer.uint32(10).string(message.amount);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QuerywithdrawalPrincipalResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQuerywithdrawalPrincipalResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.amount = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QuerywithdrawalPrincipalResponse>): QuerywithdrawalPrincipalResponse {
    const message = createBaseQuerywithdrawalPrincipalResponse();
    message.amount = object.amount ?? "";
    return message;
  }

};