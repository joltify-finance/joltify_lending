/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { DepositorInfo } from "./deposit";
import { Params } from "./params";
import { PoolInfo } from "./poolinfo";

export const protobufPackage = "joltify.spv";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryListPoolsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryListPoolsResponse {
  poolsInfo: PoolInfo[];
  pagination: PageResponse | undefined;
}

export interface QueryQueryPoolRequest {
  poolIndex: string;
}

export interface QueryQueryPoolResponse {
  poolInfo: PoolInfo | undefined;
}

export interface QueryDepositorRequest {
  walletAddress: string;
  depositPoolIndex: string;
}

export interface QueryDepositorResponse {
  depositor: DepositorInfo | undefined;
}

export interface QueryAllowedPoolsRequest {
  walletAddress: string;
}

export interface QueryAllowedPoolsResponse {
  poolsIndex: string[];
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

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
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

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryListPoolsRequest(): QueryListPoolsRequest {
  return { pagination: undefined };
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

  fromJSON(object: any): QueryListPoolsRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryListPoolsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryListPoolsRequest>, I>>(object: I): QueryListPoolsRequest {
    const message = createBaseQueryListPoolsRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryListPoolsResponse(): QueryListPoolsResponse {
  return { poolsInfo: [], pagination: undefined };
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

  fromJSON(object: any): QueryListPoolsResponse {
    return {
      poolsInfo: Array.isArray(object?.poolsInfo) ? object.poolsInfo.map((e: any) => PoolInfo.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryListPoolsResponse): unknown {
    const obj: any = {};
    if (message.poolsInfo) {
      obj.poolsInfo = message.poolsInfo.map((e) => e ? PoolInfo.toJSON(e) : undefined);
    } else {
      obj.poolsInfo = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryListPoolsResponse>, I>>(object: I): QueryListPoolsResponse {
    const message = createBaseQueryListPoolsResponse();
    message.poolsInfo = object.poolsInfo?.map((e) => PoolInfo.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryQueryPoolRequest(): QueryQueryPoolRequest {
  return { poolIndex: "" };
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

  fromJSON(object: any): QueryQueryPoolRequest {
    return { poolIndex: isSet(object.poolIndex) ? String(object.poolIndex) : "" };
  },

  toJSON(message: QueryQueryPoolRequest): unknown {
    const obj: any = {};
    message.poolIndex !== undefined && (obj.poolIndex = message.poolIndex);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQueryPoolRequest>, I>>(object: I): QueryQueryPoolRequest {
    const message = createBaseQueryQueryPoolRequest();
    message.poolIndex = object.poolIndex ?? "";
    return message;
  },
};

function createBaseQueryQueryPoolResponse(): QueryQueryPoolResponse {
  return { poolInfo: undefined };
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

  fromJSON(object: any): QueryQueryPoolResponse {
    return { poolInfo: isSet(object.poolInfo) ? PoolInfo.fromJSON(object.poolInfo) : undefined };
  },

  toJSON(message: QueryQueryPoolResponse): unknown {
    const obj: any = {};
    message.poolInfo !== undefined && (obj.poolInfo = message.poolInfo ? PoolInfo.toJSON(message.poolInfo) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQueryPoolResponse>, I>>(object: I): QueryQueryPoolResponse {
    const message = createBaseQueryQueryPoolResponse();
    message.poolInfo = (object.poolInfo !== undefined && object.poolInfo !== null)
      ? PoolInfo.fromPartial(object.poolInfo)
      : undefined;
    return message;
  },
};

function createBaseQueryDepositorRequest(): QueryDepositorRequest {
  return { walletAddress: "", depositPoolIndex: "" };
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

  fromJSON(object: any): QueryDepositorRequest {
    return {
      walletAddress: isSet(object.walletAddress) ? String(object.walletAddress) : "",
      depositPoolIndex: isSet(object.depositPoolIndex) ? String(object.depositPoolIndex) : "",
    };
  },

  toJSON(message: QueryDepositorRequest): unknown {
    const obj: any = {};
    message.walletAddress !== undefined && (obj.walletAddress = message.walletAddress);
    message.depositPoolIndex !== undefined && (obj.depositPoolIndex = message.depositPoolIndex);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryDepositorRequest>, I>>(object: I): QueryDepositorRequest {
    const message = createBaseQueryDepositorRequest();
    message.walletAddress = object.walletAddress ?? "";
    message.depositPoolIndex = object.depositPoolIndex ?? "";
    return message;
  },
};

function createBaseQueryDepositorResponse(): QueryDepositorResponse {
  return { depositor: undefined };
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

  fromJSON(object: any): QueryDepositorResponse {
    return { depositor: isSet(object.depositor) ? DepositorInfo.fromJSON(object.depositor) : undefined };
  },

  toJSON(message: QueryDepositorResponse): unknown {
    const obj: any = {};
    message.depositor !== undefined
      && (obj.depositor = message.depositor ? DepositorInfo.toJSON(message.depositor) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryDepositorResponse>, I>>(object: I): QueryDepositorResponse {
    const message = createBaseQueryDepositorResponse();
    message.depositor = (object.depositor !== undefined && object.depositor !== null)
      ? DepositorInfo.fromPartial(object.depositor)
      : undefined;
    return message;
  },
};

function createBaseQueryAllowedPoolsRequest(): QueryAllowedPoolsRequest {
  return { walletAddress: "" };
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

  fromJSON(object: any): QueryAllowedPoolsRequest {
    return { walletAddress: isSet(object.walletAddress) ? String(object.walletAddress) : "" };
  },

  toJSON(message: QueryAllowedPoolsRequest): unknown {
    const obj: any = {};
    message.walletAddress !== undefined && (obj.walletAddress = message.walletAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllowedPoolsRequest>, I>>(object: I): QueryAllowedPoolsRequest {
    const message = createBaseQueryAllowedPoolsRequest();
    message.walletAddress = object.walletAddress ?? "";
    return message;
  },
};

function createBaseQueryAllowedPoolsResponse(): QueryAllowedPoolsResponse {
  return { poolsIndex: [] };
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

  fromJSON(object: any): QueryAllowedPoolsResponse {
    return { poolsIndex: Array.isArray(object?.poolsIndex) ? object.poolsIndex.map((e: any) => String(e)) : [] };
  },

  toJSON(message: QueryAllowedPoolsResponse): unknown {
    const obj: any = {};
    if (message.poolsIndex) {
      obj.poolsIndex = message.poolsIndex.map((e) => e);
    } else {
      obj.poolsIndex = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllowedPoolsResponse>, I>>(object: I): QueryAllowedPoolsResponse {
    const message = createBaseQueryAllowedPoolsResponse();
    message.poolsIndex = object.poolsIndex?.map((e) => e) || [];
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of Listpools items. */
  ListPools(request: QueryListPoolsRequest): Promise<QueryListPoolsResponse>;
  /** Queries a list of QueryPool items. */
  QueryPool(request: QueryQueryPoolRequest): Promise<QueryQueryPoolResponse>;
  Depositor(request: QueryDepositorRequest): Promise<QueryDepositorResponse>;
  /** Queries a list of AllowedPools items. */
  AllowedPools(request: QueryAllowedPoolsRequest): Promise<QueryAllowedPoolsResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.ListPools = this.ListPools.bind(this);
    this.QueryPool = this.QueryPool.bind(this);
    this.Depositor = this.Depositor.bind(this);
    this.AllowedPools = this.AllowedPools.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  ListPools(request: QueryListPoolsRequest): Promise<QueryListPoolsResponse> {
    const data = QueryListPoolsRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Query", "ListPools", data);
    return promise.then((data) => QueryListPoolsResponse.decode(new _m0.Reader(data)));
  }

  QueryPool(request: QueryQueryPoolRequest): Promise<QueryQueryPoolResponse> {
    const data = QueryQueryPoolRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Query", "QueryPool", data);
    return promise.then((data) => QueryQueryPoolResponse.decode(new _m0.Reader(data)));
  }

  Depositor(request: QueryDepositorRequest): Promise<QueryDepositorResponse> {
    const data = QueryDepositorRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Query", "Depositor", data);
    return promise.then((data) => QueryDepositorResponse.decode(new _m0.Reader(data)));
  }

  AllowedPools(request: QueryAllowedPoolsRequest): Promise<QueryAllowedPoolsResponse> {
    const data = QueryAllowedPoolsRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Query", "AllowedPools", data);
    return promise.then((data) => QueryAllowedPoolsResponse.decode(new _m0.Reader(data)));
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
