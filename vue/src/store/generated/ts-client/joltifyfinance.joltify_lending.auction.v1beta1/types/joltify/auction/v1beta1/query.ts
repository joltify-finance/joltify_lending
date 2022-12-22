/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../../cosmos/base/query/v1beta1/pagination";
import { Any } from "../../../google/protobuf/any";
import { Params } from "./genesis";

export const protobufPackage = "joltifyfinance.joltify_lending.auction.v1beta1";

/** QueryParamsRequest defines the request type for querying x/auction parameters. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse defines the response type for querying x/auction parameters. */
export interface QueryParamsResponse {
  params: Params | undefined;
}

/** QueryAuctionRequest is the request type for the Query/Auction RPC method. */
export interface QueryAuctionRequest {
  auctionId: number;
}

/** QueryAuctionResponse is the response type for the Query/Auction RPC method. */
export interface QueryAuctionResponse {
  auction: Any | undefined;
}

/** QueryAuctionsRequest is the request type for the Query/Auctions RPC method. */
export interface QueryAuctionsRequest {
  type: string;
  owner: string;
  denom: string;
  phase: string;
  /** pagination defines an optional pagination for the request. */
  pagination: PageRequest | undefined;
}

/** QueryAuctionsResponse is the response type for the Query/Auctions RPC method. */
export interface QueryAuctionsResponse {
  auctions: Any[];
  /** pagination defines the pagination in the response. */
  pagination: PageResponse | undefined;
}

/** QueryNextAuctionIDRequest defines the request type for querying x/auction next auction ID. */
export interface QueryNextAuctionIDRequest {
}

/** QueryNextAuctionIDResponse defines the response type for querying x/auction next auction ID. */
export interface QueryNextAuctionIDResponse {
  id: number;
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

function createBaseQueryAuctionRequest(): QueryAuctionRequest {
  return { auctionId: 0 };
}

export const QueryAuctionRequest = {
  encode(message: QueryAuctionRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.auctionId !== 0) {
      writer.uint32(8).uint64(message.auctionId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAuctionRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAuctionRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.auctionId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAuctionRequest {
    return { auctionId: isSet(object.auctionId) ? Number(object.auctionId) : 0 };
  },

  toJSON(message: QueryAuctionRequest): unknown {
    const obj: any = {};
    message.auctionId !== undefined && (obj.auctionId = Math.round(message.auctionId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAuctionRequest>, I>>(object: I): QueryAuctionRequest {
    const message = createBaseQueryAuctionRequest();
    message.auctionId = object.auctionId ?? 0;
    return message;
  },
};

function createBaseQueryAuctionResponse(): QueryAuctionResponse {
  return { auction: undefined };
}

export const QueryAuctionResponse = {
  encode(message: QueryAuctionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.auction !== undefined) {
      Any.encode(message.auction, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAuctionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAuctionResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.auction = Any.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAuctionResponse {
    return { auction: isSet(object.auction) ? Any.fromJSON(object.auction) : undefined };
  },

  toJSON(message: QueryAuctionResponse): unknown {
    const obj: any = {};
    message.auction !== undefined && (obj.auction = message.auction ? Any.toJSON(message.auction) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAuctionResponse>, I>>(object: I): QueryAuctionResponse {
    const message = createBaseQueryAuctionResponse();
    message.auction = (object.auction !== undefined && object.auction !== null)
      ? Any.fromPartial(object.auction)
      : undefined;
    return message;
  },
};

function createBaseQueryAuctionsRequest(): QueryAuctionsRequest {
  return { type: "", owner: "", denom: "", phase: "", pagination: undefined };
}

export const QueryAuctionsRequest = {
  encode(message: QueryAuctionsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.type !== "") {
      writer.uint32(10).string(message.type);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    if (message.denom !== "") {
      writer.uint32(26).string(message.denom);
    }
    if (message.phase !== "") {
      writer.uint32(34).string(message.phase);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAuctionsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAuctionsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.type = reader.string();
          break;
        case 2:
          message.owner = reader.string();
          break;
        case 3:
          message.denom = reader.string();
          break;
        case 4:
          message.phase = reader.string();
          break;
        case 5:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAuctionsRequest {
    return {
      type: isSet(object.type) ? String(object.type) : "",
      owner: isSet(object.owner) ? String(object.owner) : "",
      denom: isSet(object.denom) ? String(object.denom) : "",
      phase: isSet(object.phase) ? String(object.phase) : "",
      pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAuctionsRequest): unknown {
    const obj: any = {};
    message.type !== undefined && (obj.type = message.type);
    message.owner !== undefined && (obj.owner = message.owner);
    message.denom !== undefined && (obj.denom = message.denom);
    message.phase !== undefined && (obj.phase = message.phase);
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAuctionsRequest>, I>>(object: I): QueryAuctionsRequest {
    const message = createBaseQueryAuctionsRequest();
    message.type = object.type ?? "";
    message.owner = object.owner ?? "";
    message.denom = object.denom ?? "";
    message.phase = object.phase ?? "";
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAuctionsResponse(): QueryAuctionsResponse {
  return { auctions: [], pagination: undefined };
}

export const QueryAuctionsResponse = {
  encode(message: QueryAuctionsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.auctions) {
      Any.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAuctionsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAuctionsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.auctions.push(Any.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAuctionsResponse {
    return {
      auctions: Array.isArray(object?.auctions) ? object.auctions.map((e: any) => Any.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAuctionsResponse): unknown {
    const obj: any = {};
    if (message.auctions) {
      obj.auctions = message.auctions.map((e) => e ? Any.toJSON(e) : undefined);
    } else {
      obj.auctions = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAuctionsResponse>, I>>(object: I): QueryAuctionsResponse {
    const message = createBaseQueryAuctionsResponse();
    message.auctions = object.auctions?.map((e) => Any.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryNextAuctionIDRequest(): QueryNextAuctionIDRequest {
  return {};
}

export const QueryNextAuctionIDRequest = {
  encode(_: QueryNextAuctionIDRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryNextAuctionIDRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryNextAuctionIDRequest();
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

  fromJSON(_: any): QueryNextAuctionIDRequest {
    return {};
  },

  toJSON(_: QueryNextAuctionIDRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryNextAuctionIDRequest>, I>>(_: I): QueryNextAuctionIDRequest {
    const message = createBaseQueryNextAuctionIDRequest();
    return message;
  },
};

function createBaseQueryNextAuctionIDResponse(): QueryNextAuctionIDResponse {
  return { id: 0 };
}

export const QueryNextAuctionIDResponse = {
  encode(message: QueryNextAuctionIDResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryNextAuctionIDResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryNextAuctionIDResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryNextAuctionIDResponse {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: QueryNextAuctionIDResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryNextAuctionIDResponse>, I>>(object: I): QueryNextAuctionIDResponse {
    const message = createBaseQueryNextAuctionIDResponse();
    message.id = object.id ?? 0;
    return message;
  },
};

/** Query defines the gRPC querier service for auction module */
export interface Query {
  /** Params queries all parameters of the auction module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Auction queries an individual Auction by auction ID */
  Auction(request: QueryAuctionRequest): Promise<QueryAuctionResponse>;
  /** Auctions queries auctions filtered by asset denom, owner address, phase, and auction type */
  Auctions(request: QueryAuctionsRequest): Promise<QueryAuctionsResponse>;
  /** NextAuctionID queries the next auction ID */
  NextAuctionID(request: QueryNextAuctionIDRequest): Promise<QueryNextAuctionIDResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Auction = this.Auction.bind(this);
    this.Auctions = this.Auctions.bind(this);
    this.NextAuctionID = this.NextAuctionID.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("joltifyfinance.joltify_lending.auction.v1beta1.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  Auction(request: QueryAuctionRequest): Promise<QueryAuctionResponse> {
    const data = QueryAuctionRequest.encode(request).finish();
    const promise = this.rpc.request("joltifyfinance.joltify_lending.auction.v1beta1.Query", "Auction", data);
    return promise.then((data) => QueryAuctionResponse.decode(new _m0.Reader(data)));
  }

  Auctions(request: QueryAuctionsRequest): Promise<QueryAuctionsResponse> {
    const data = QueryAuctionsRequest.encode(request).finish();
    const promise = this.rpc.request("joltifyfinance.joltify_lending.auction.v1beta1.Query", "Auctions", data);
    return promise.then((data) => QueryAuctionsResponse.decode(new _m0.Reader(data)));
  }

  NextAuctionID(request: QueryNextAuctionIDRequest): Promise<QueryNextAuctionIDResponse> {
    const data = QueryNextAuctionIDRequest.encode(request).finish();
    const promise = this.rpc.request("joltifyfinance.joltify_lending.auction.v1beta1.Query", "NextAuctionID", data);
    return promise.then((data) => QueryNextAuctionIDResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
