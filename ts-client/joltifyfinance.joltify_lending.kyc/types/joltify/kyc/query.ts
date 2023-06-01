/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { Investor } from "./investor";
import { Params } from "./params";

export const protobufPackage = "joltifyfinance.joltify_lending.kyc";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryInvestorWalletsRequest {
  investorId: string;
}

export interface QueryInvestorWalletsResponse {
  wallets: string[];
}

export interface QueryByWalletRequest {
  wallet: string;
}

export interface QueryByWalletResponse {
  investor: Investor | undefined;
}

export interface ListInvestorsRequest {
  pagination: PageRequest | undefined;
}

export interface ListInvestorsResponse {
  investor: Investor[];
  pagination: PageResponse | undefined;
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

function createBaseQueryInvestorWalletsRequest(): QueryInvestorWalletsRequest {
  return { investorId: "" };
}

export const QueryInvestorWalletsRequest = {
  encode(message: QueryInvestorWalletsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.investorId !== "") {
      writer.uint32(10).string(message.investorId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryInvestorWalletsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryInvestorWalletsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.investorId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryInvestorWalletsRequest {
    return { investorId: isSet(object.investorId) ? String(object.investorId) : "" };
  },

  toJSON(message: QueryInvestorWalletsRequest): unknown {
    const obj: any = {};
    message.investorId !== undefined && (obj.investorId = message.investorId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryInvestorWalletsRequest>, I>>(object: I): QueryInvestorWalletsRequest {
    const message = createBaseQueryInvestorWalletsRequest();
    message.investorId = object.investorId ?? "";
    return message;
  },
};

function createBaseQueryInvestorWalletsResponse(): QueryInvestorWalletsResponse {
  return { wallets: [] };
}

export const QueryInvestorWalletsResponse = {
  encode(message: QueryInvestorWalletsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.wallets) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryInvestorWalletsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryInvestorWalletsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.wallets.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryInvestorWalletsResponse {
    return { wallets: Array.isArray(object?.wallets) ? object.wallets.map((e: any) => String(e)) : [] };
  },

  toJSON(message: QueryInvestorWalletsResponse): unknown {
    const obj: any = {};
    if (message.wallets) {
      obj.wallets = message.wallets.map((e) => e);
    } else {
      obj.wallets = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryInvestorWalletsResponse>, I>>(object: I): QueryInvestorWalletsResponse {
    const message = createBaseQueryInvestorWalletsResponse();
    message.wallets = object.wallets?.map((e) => e) || [];
    return message;
  },
};

function createBaseQueryByWalletRequest(): QueryByWalletRequest {
  return { wallet: "" };
}

export const QueryByWalletRequest = {
  encode(message: QueryByWalletRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.wallet !== "") {
      writer.uint32(10).string(message.wallet);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryByWalletRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryByWalletRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.wallet = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryByWalletRequest {
    return { wallet: isSet(object.wallet) ? String(object.wallet) : "" };
  },

  toJSON(message: QueryByWalletRequest): unknown {
    const obj: any = {};
    message.wallet !== undefined && (obj.wallet = message.wallet);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryByWalletRequest>, I>>(object: I): QueryByWalletRequest {
    const message = createBaseQueryByWalletRequest();
    message.wallet = object.wallet ?? "";
    return message;
  },
};

function createBaseQueryByWalletResponse(): QueryByWalletResponse {
  return { investor: undefined };
}

export const QueryByWalletResponse = {
  encode(message: QueryByWalletResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.investor !== undefined) {
      Investor.encode(message.investor, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryByWalletResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryByWalletResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.investor = Investor.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryByWalletResponse {
    return { investor: isSet(object.investor) ? Investor.fromJSON(object.investor) : undefined };
  },

  toJSON(message: QueryByWalletResponse): unknown {
    const obj: any = {};
    message.investor !== undefined && (obj.investor = message.investor ? Investor.toJSON(message.investor) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryByWalletResponse>, I>>(object: I): QueryByWalletResponse {
    const message = createBaseQueryByWalletResponse();
    message.investor = (object.investor !== undefined && object.investor !== null)
      ? Investor.fromPartial(object.investor)
      : undefined;
    return message;
  },
};

function createBaseListInvestorsRequest(): ListInvestorsRequest {
  return { pagination: undefined };
}

export const ListInvestorsRequest = {
  encode(message: ListInvestorsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListInvestorsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListInvestorsRequest();
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

  fromJSON(object: any): ListInvestorsRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: ListInvestorsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListInvestorsRequest>, I>>(object: I): ListInvestorsRequest {
    const message = createBaseListInvestorsRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseListInvestorsResponse(): ListInvestorsResponse {
  return { investor: [], pagination: undefined };
}

export const ListInvestorsResponse = {
  encode(message: ListInvestorsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.investor) {
      Investor.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListInvestorsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListInvestorsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.investor.push(Investor.decode(reader, reader.uint32()));
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

  fromJSON(object: any): ListInvestorsResponse {
    return {
      investor: Array.isArray(object?.investor) ? object.investor.map((e: any) => Investor.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: ListInvestorsResponse): unknown {
    const obj: any = {};
    if (message.investor) {
      obj.investor = message.investor.map((e) => e ? Investor.toJSON(e) : undefined);
    } else {
      obj.investor = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListInvestorsResponse>, I>>(object: I): ListInvestorsResponse {
    const message = createBaseListInvestorsResponse();
    message.investor = object.investor?.map((e) => Investor.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of QueryInvestorWallets items. */
  QueryInvestorWallets(request: QueryInvestorWalletsRequest): Promise<QueryInvestorWalletsResponse>;
  /** Queries a list of QueryByWallet items. */
  QueryByWallet(request: QueryByWalletRequest): Promise<QueryByWalletResponse>;
  /** Queries a list of ListInvestors items. */
  ListInvestors(request: ListInvestorsRequest): Promise<ListInvestorsResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.QueryInvestorWallets = this.QueryInvestorWallets.bind(this);
    this.QueryByWallet = this.QueryByWallet.bind(this);
    this.ListInvestors = this.ListInvestors.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("joltifyfinance.joltify_lending.kyc.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  QueryInvestorWallets(request: QueryInvestorWalletsRequest): Promise<QueryInvestorWalletsResponse> {
    const data = QueryInvestorWalletsRequest.encode(request).finish();
    const promise = this.rpc.request("joltifyfinance.joltify_lending.kyc.Query", "QueryInvestorWallets", data);
    return promise.then((data) => QueryInvestorWalletsResponse.decode(new _m0.Reader(data)));
  }

  QueryByWallet(request: QueryByWalletRequest): Promise<QueryByWalletResponse> {
    const data = QueryByWalletRequest.encode(request).finish();
    const promise = this.rpc.request("joltifyfinance.joltify_lending.kyc.Query", "QueryByWallet", data);
    return promise.then((data) => QueryByWalletResponse.decode(new _m0.Reader(data)));
  }

  ListInvestors(request: ListInvestorsRequest): Promise<ListInvestorsResponse> {
    const data = ListInvestorsRequest.encode(request).finish();
    const promise = this.rpc.request("joltifyfinance.joltify_lending.kyc.Query", "ListInvestors", data);
    return promise.then((data) => ListInvestorsResponse.decode(new _m0.Reader(data)));
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
