import { PageRequest, PageRequestSDKType, PageResponse, PageResponseSDKType } from "../../cosmos/base/query/v1beta1/pagination";
import { Params, ParamsSDKType } from "./params";
import { Investor, InvestorSDKType } from "./investor";
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
export interface QueryInvestorWalletsRequest {
  investorId: string;
}
export interface QueryInvestorWalletsRequestSDKType {
  investorId: string;
}
export interface QueryInvestorWalletsResponse {
  wallets: string[];
}
export interface QueryInvestorWalletsResponseSDKType {
  wallets: string[];
}
export interface QueryByWalletRequest {
  wallet: string;
}
export interface QueryByWalletRequestSDKType {
  wallet: string;
}
export interface QueryByWalletResponse {
  investor?: Investor;
}
export interface QueryByWalletResponseSDKType {
  investor?: InvestorSDKType;
}
export interface ListInvestorsRequest {
  pagination?: PageRequest;
}
export interface ListInvestorsRequestSDKType {
  pagination?: PageRequestSDKType;
}
export interface ListInvestorsResponse {
  investor: Investor[];
  pagination?: PageResponse;
}
export interface ListInvestorsResponseSDKType {
  investor: InvestorSDKType[];
  pagination?: PageResponseSDKType;
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

function createBaseQueryInvestorWalletsRequest(): QueryInvestorWalletsRequest {
  return {
    investorId: ""
  };
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

  fromPartial(object: DeepPartial<QueryInvestorWalletsRequest>): QueryInvestorWalletsRequest {
    const message = createBaseQueryInvestorWalletsRequest();
    message.investorId = object.investorId ?? "";
    return message;
  }

};

function createBaseQueryInvestorWalletsResponse(): QueryInvestorWalletsResponse {
  return {
    wallets: []
  };
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

  fromPartial(object: DeepPartial<QueryInvestorWalletsResponse>): QueryInvestorWalletsResponse {
    const message = createBaseQueryInvestorWalletsResponse();
    message.wallets = object.wallets?.map(e => e) || [];
    return message;
  }

};

function createBaseQueryByWalletRequest(): QueryByWalletRequest {
  return {
    wallet: ""
  };
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

  fromPartial(object: DeepPartial<QueryByWalletRequest>): QueryByWalletRequest {
    const message = createBaseQueryByWalletRequest();
    message.wallet = object.wallet ?? "";
    return message;
  }

};

function createBaseQueryByWalletResponse(): QueryByWalletResponse {
  return {
    investor: undefined
  };
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

  fromPartial(object: DeepPartial<QueryByWalletResponse>): QueryByWalletResponse {
    const message = createBaseQueryByWalletResponse();
    message.investor = object.investor !== undefined && object.investor !== null ? Investor.fromPartial(object.investor) : undefined;
    return message;
  }

};

function createBaseListInvestorsRequest(): ListInvestorsRequest {
  return {
    pagination: undefined
  };
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

  fromPartial(object: DeepPartial<ListInvestorsRequest>): ListInvestorsRequest {
    const message = createBaseListInvestorsRequest();
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageRequest.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseListInvestorsResponse(): ListInvestorsResponse {
  return {
    investor: [],
    pagination: undefined
  };
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

  fromPartial(object: DeepPartial<ListInvestorsResponse>): ListInvestorsResponse {
    const message = createBaseListInvestorsResponse();
    message.investor = object.investor?.map(e => Investor.fromPartial(e)) || [];
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageResponse.fromPartial(object.pagination) : undefined;
    return message;
  }

};