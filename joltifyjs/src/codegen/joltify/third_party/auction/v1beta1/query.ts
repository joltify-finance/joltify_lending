import { PageRequest, PageRequestSDKType, PageResponse, PageResponseSDKType } from "../../../../cosmos/base/query/v1beta1/pagination";
import { Params, ParamsSDKType } from "./genesis";
import { Any, AnySDKType } from "../../../../google/protobuf/any";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial, Long } from "../../../../helpers";
/** QueryParamsRequest defines the request type for querying x/auction parameters. */

export interface QueryParamsRequest {}
/** QueryParamsRequest defines the request type for querying x/auction parameters. */

export interface QueryParamsRequestSDKType {}
/** QueryParamsResponse defines the response type for querying x/auction parameters. */

export interface QueryParamsResponse {
  params?: Params;
}
/** QueryParamsResponse defines the response type for querying x/auction parameters. */

export interface QueryParamsResponseSDKType {
  params?: ParamsSDKType;
}
/** QueryAuctionRequest is the request type for the Query/Auction RPC method. */

export interface QueryAuctionRequest {
  auctionId: Long;
}
/** QueryAuctionRequest is the request type for the Query/Auction RPC method. */

export interface QueryAuctionRequestSDKType {
  auction_id: Long;
}
/** QueryAuctionResponse is the response type for the Query/Auction RPC method. */

export interface QueryAuctionResponse {
  auction?: Any;
}
/** QueryAuctionResponse is the response type for the Query/Auction RPC method. */

export interface QueryAuctionResponseSDKType {
  auction?: AnySDKType;
}
/** QueryAuctionsRequest is the request type for the Query/Auctions RPC method. */

export interface QueryAuctionsRequest {
  type: string;
  owner: string;
  denom: string;
  phase: string;
  /** pagination defines an optional pagination for the request. */

  pagination?: PageRequest;
}
/** QueryAuctionsRequest is the request type for the Query/Auctions RPC method. */

export interface QueryAuctionsRequestSDKType {
  type: string;
  owner: string;
  denom: string;
  phase: string;
  pagination?: PageRequestSDKType;
}
/** QueryAuctionsResponse is the response type for the Query/Auctions RPC method. */

export interface QueryAuctionsResponse {
  auctions: Any[];
  /** pagination defines the pagination in the response. */

  pagination?: PageResponse;
}
/** QueryAuctionsResponse is the response type for the Query/Auctions RPC method. */

export interface QueryAuctionsResponseSDKType {
  auctions: AnySDKType[];
  pagination?: PageResponseSDKType;
}
/** QueryNextAuctionIDRequest defines the request type for querying x/auction next auction ID. */

export interface QueryNextAuctionIDRequest {}
/** QueryNextAuctionIDRequest defines the request type for querying x/auction next auction ID. */

export interface QueryNextAuctionIDRequestSDKType {}
/** QueryNextAuctionIDResponse defines the response type for querying x/auction next auction ID. */

export interface QueryNextAuctionIDResponse {
  id: Long;
}
/** QueryNextAuctionIDResponse defines the response type for querying x/auction next auction ID. */

export interface QueryNextAuctionIDResponseSDKType {
  id: Long;
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

function createBaseQueryAuctionRequest(): QueryAuctionRequest {
  return {
    auctionId: Long.UZERO
  };
}

export const QueryAuctionRequest = {
  encode(message: QueryAuctionRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (!message.auctionId.isZero()) {
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
          message.auctionId = (reader.uint64() as Long);
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryAuctionRequest>): QueryAuctionRequest {
    const message = createBaseQueryAuctionRequest();
    message.auctionId = object.auctionId !== undefined && object.auctionId !== null ? Long.fromValue(object.auctionId) : Long.UZERO;
    return message;
  }

};

function createBaseQueryAuctionResponse(): QueryAuctionResponse {
  return {
    auction: undefined
  };
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

  fromPartial(object: DeepPartial<QueryAuctionResponse>): QueryAuctionResponse {
    const message = createBaseQueryAuctionResponse();
    message.auction = object.auction !== undefined && object.auction !== null ? Any.fromPartial(object.auction) : undefined;
    return message;
  }

};

function createBaseQueryAuctionsRequest(): QueryAuctionsRequest {
  return {
    type: "",
    owner: "",
    denom: "",
    phase: "",
    pagination: undefined
  };
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

  fromPartial(object: DeepPartial<QueryAuctionsRequest>): QueryAuctionsRequest {
    const message = createBaseQueryAuctionsRequest();
    message.type = object.type ?? "";
    message.owner = object.owner ?? "";
    message.denom = object.denom ?? "";
    message.phase = object.phase ?? "";
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageRequest.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryAuctionsResponse(): QueryAuctionsResponse {
  return {
    auctions: [],
    pagination: undefined
  };
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

  fromPartial(object: DeepPartial<QueryAuctionsResponse>): QueryAuctionsResponse {
    const message = createBaseQueryAuctionsResponse();
    message.auctions = object.auctions?.map(e => Any.fromPartial(e)) || [];
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageResponse.fromPartial(object.pagination) : undefined;
    return message;
  }

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

  fromPartial(_: DeepPartial<QueryNextAuctionIDRequest>): QueryNextAuctionIDRequest {
    const message = createBaseQueryNextAuctionIDRequest();
    return message;
  }

};

function createBaseQueryNextAuctionIDResponse(): QueryNextAuctionIDResponse {
  return {
    id: Long.UZERO
  };
}

export const QueryNextAuctionIDResponse = {
  encode(message: QueryNextAuctionIDResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (!message.id.isZero()) {
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
          message.id = (reader.uint64() as Long);
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryNextAuctionIDResponse>): QueryNextAuctionIDResponse {
    const message = createBaseQueryNextAuctionIDResponse();
    message.id = object.id !== undefined && object.id !== null ? Long.fromValue(object.id) : Long.UZERO;
    return message;
  }

};