import { PageRequest, PageRequestSDKType, PageResponse, PageResponseSDKType } from "../../../../cosmos/base/query/v1beta1/pagination";
import { Params, ParamsSDKType } from "./genesis";
import { ModuleAccount, ModuleAccountSDKType } from "../../../../cosmos/auth/v1beta1/auth";
import { Deposit, DepositSDKType, TotalPrincipal, TotalPrincipalSDKType, TotalCollateral, TotalCollateralSDKType } from "./cdp";
import { Coin, CoinSDKType } from "../../../../cosmos/base/v1beta1/coin";
import { Timestamp } from "../../../../google/protobuf/timestamp";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial, Long, toTimestamp, fromTimestamp } from "../../../../helpers";
/** QueryParamsRequest defines the request type for the Query/Params RPC method. */

export interface QueryParamsRequest {}
/** QueryParamsRequest defines the request type for the Query/Params RPC method. */

export interface QueryParamsRequestSDKType {}
/** QueryParamsResponse defines the response type for the Query/Params RPC method. */

export interface QueryParamsResponse {
  params?: Params;
}
/** QueryParamsResponse defines the response type for the Query/Params RPC method. */

export interface QueryParamsResponseSDKType {
  params?: ParamsSDKType;
}
/** QueryAccountsRequest defines the request type for the Query/Accounts RPC method. */

export interface QueryAccountsRequest {}
/** QueryAccountsRequest defines the request type for the Query/Accounts RPC method. */

export interface QueryAccountsRequestSDKType {}
/** QueryAccountsResponse defines the response type for the Query/Accounts RPC method. */

export interface QueryAccountsResponse {
  accounts: ModuleAccount[];
}
/** QueryAccountsResponse defines the response type for the Query/Accounts RPC method. */

export interface QueryAccountsResponseSDKType {
  accounts: ModuleAccountSDKType[];
}
/** QueryCdpRequest defines the request type for the Query/Cdp RPC method. */

export interface QueryCdpRequest {
  collateralType: string;
  owner: string;
}
/** QueryCdpRequest defines the request type for the Query/Cdp RPC method. */

export interface QueryCdpRequestSDKType {
  collateral_type: string;
  owner: string;
}
/** QueryCdpResponse defines the response type for the Query/Cdp RPC method. */

export interface QueryCdpResponse {
  cdp?: CDPResponse;
}
/** QueryCdpResponse defines the response type for the Query/Cdp RPC method. */

export interface QueryCdpResponseSDKType {
  cdp?: CDPResponseSDKType;
}
/** QueryCdpsRequest is the params for a filtered CDP query, the request type for the Query/Cdps RPC method. */

export interface QueryCdpsRequest {
  collateralType: string;
  owner: string;
  id: Long;
  /** sdk.Dec as a string */

  ratio: string;
  pagination?: PageRequest;
}
/** QueryCdpsRequest is the params for a filtered CDP query, the request type for the Query/Cdps RPC method. */

export interface QueryCdpsRequestSDKType {
  collateral_type: string;
  owner: string;
  id: Long;
  ratio: string;
  pagination?: PageRequestSDKType;
}
/** QueryCdpsResponse defines the response type for the Query/Cdps RPC method. */

export interface QueryCdpsResponse {
  cdps: CDPResponse[];
  pagination?: PageResponse;
}
/** QueryCdpsResponse defines the response type for the Query/Cdps RPC method. */

export interface QueryCdpsResponseSDKType {
  cdps: CDPResponseSDKType[];
  pagination?: PageResponseSDKType;
}
/** QueryDepositsRequest defines the request type for the Query/Deposits RPC method. */

export interface QueryDepositsRequest {
  collateralType: string;
  owner: string;
}
/** QueryDepositsRequest defines the request type for the Query/Deposits RPC method. */

export interface QueryDepositsRequestSDKType {
  collateral_type: string;
  owner: string;
}
/** QueryDepositsResponse defines the response type for the Query/Deposits RPC method. */

export interface QueryDepositsResponse {
  deposits: Deposit[];
}
/** QueryDepositsResponse defines the response type for the Query/Deposits RPC method. */

export interface QueryDepositsResponseSDKType {
  deposits: DepositSDKType[];
}
/** QueryTotalPrincipalRequest defines the request type for the Query/TotalPrincipal RPC method. */

export interface QueryTotalPrincipalRequest {
  collateralType: string;
}
/** QueryTotalPrincipalRequest defines the request type for the Query/TotalPrincipal RPC method. */

export interface QueryTotalPrincipalRequestSDKType {
  collateral_type: string;
}
/** QueryTotalPrincipalResponse defines the response type for the Query/TotalPrincipal RPC method. */

export interface QueryTotalPrincipalResponse {
  totalPrincipal: TotalPrincipal[];
}
/** QueryTotalPrincipalResponse defines the response type for the Query/TotalPrincipal RPC method. */

export interface QueryTotalPrincipalResponseSDKType {
  total_principal: TotalPrincipalSDKType[];
}
/** QueryTotalCollateralRequest defines the request type for the Query/TotalCollateral RPC method. */

export interface QueryTotalCollateralRequest {
  collateralType: string;
}
/** QueryTotalCollateralRequest defines the request type for the Query/TotalCollateral RPC method. */

export interface QueryTotalCollateralRequestSDKType {
  collateral_type: string;
}
/** QueryTotalCollateralResponse defines the response type for the Query/TotalCollateral RPC method. */

export interface QueryTotalCollateralResponse {
  totalCollateral: TotalCollateral[];
}
/** QueryTotalCollateralResponse defines the response type for the Query/TotalCollateral RPC method. */

export interface QueryTotalCollateralResponseSDKType {
  total_collateral: TotalCollateralSDKType[];
}
/** CDPResponse defines the state of a single collateralized debt position. */

export interface CDPResponse {
  id: Long;
  owner: string;
  type: string;
  collateral?: Coin;
  principal?: Coin;
  accumulatedFees?: Coin;
  feesUpdated?: Date;
  interestFactor: string;
  collateralValue?: Coin;
  collateralizationRatio: string;
}
/** CDPResponse defines the state of a single collateralized debt position. */

export interface CDPResponseSDKType {
  id: Long;
  owner: string;
  type: string;
  collateral?: CoinSDKType;
  principal?: CoinSDKType;
  accumulated_fees?: CoinSDKType;
  fees_updated?: Date;
  interest_factor: string;
  collateral_value?: CoinSDKType;
  collateralization_ratio: string;
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

function createBaseQueryAccountsRequest(): QueryAccountsRequest {
  return {};
}

export const QueryAccountsRequest = {
  encode(_: QueryAccountsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAccountsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAccountsRequest();

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

  fromPartial(_: DeepPartial<QueryAccountsRequest>): QueryAccountsRequest {
    const message = createBaseQueryAccountsRequest();
    return message;
  }

};

function createBaseQueryAccountsResponse(): QueryAccountsResponse {
  return {
    accounts: []
  };
}

export const QueryAccountsResponse = {
  encode(message: QueryAccountsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.accounts) {
      ModuleAccount.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAccountsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAccountsResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.accounts.push(ModuleAccount.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryAccountsResponse>): QueryAccountsResponse {
    const message = createBaseQueryAccountsResponse();
    message.accounts = object.accounts?.map(e => ModuleAccount.fromPartial(e)) || [];
    return message;
  }

};

function createBaseQueryCdpRequest(): QueryCdpRequest {
  return {
    collateralType: "",
    owner: ""
  };
}

export const QueryCdpRequest = {
  encode(message: QueryCdpRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.collateralType !== "") {
      writer.uint32(10).string(message.collateralType);
    }

    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCdpRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCdpRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.collateralType = reader.string();
          break;

        case 2:
          message.owner = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryCdpRequest>): QueryCdpRequest {
    const message = createBaseQueryCdpRequest();
    message.collateralType = object.collateralType ?? "";
    message.owner = object.owner ?? "";
    return message;
  }

};

function createBaseQueryCdpResponse(): QueryCdpResponse {
  return {
    cdp: undefined
  };
}

export const QueryCdpResponse = {
  encode(message: QueryCdpResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.cdp !== undefined) {
      CDPResponse.encode(message.cdp, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCdpResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCdpResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.cdp = CDPResponse.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryCdpResponse>): QueryCdpResponse {
    const message = createBaseQueryCdpResponse();
    message.cdp = object.cdp !== undefined && object.cdp !== null ? CDPResponse.fromPartial(object.cdp) : undefined;
    return message;
  }

};

function createBaseQueryCdpsRequest(): QueryCdpsRequest {
  return {
    collateralType: "",
    owner: "",
    id: Long.UZERO,
    ratio: "",
    pagination: undefined
  };
}

export const QueryCdpsRequest = {
  encode(message: QueryCdpsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.collateralType !== "") {
      writer.uint32(10).string(message.collateralType);
    }

    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }

    if (!message.id.isZero()) {
      writer.uint32(24).uint64(message.id);
    }

    if (message.ratio !== "") {
      writer.uint32(34).string(message.ratio);
    }

    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(42).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCdpsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCdpsRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.collateralType = reader.string();
          break;

        case 2:
          message.owner = reader.string();
          break;

        case 3:
          message.id = (reader.uint64() as Long);
          break;

        case 4:
          message.ratio = reader.string();
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

  fromPartial(object: DeepPartial<QueryCdpsRequest>): QueryCdpsRequest {
    const message = createBaseQueryCdpsRequest();
    message.collateralType = object.collateralType ?? "";
    message.owner = object.owner ?? "";
    message.id = object.id !== undefined && object.id !== null ? Long.fromValue(object.id) : Long.UZERO;
    message.ratio = object.ratio ?? "";
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageRequest.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryCdpsResponse(): QueryCdpsResponse {
  return {
    cdps: [],
    pagination: undefined
  };
}

export const QueryCdpsResponse = {
  encode(message: QueryCdpsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.cdps) {
      CDPResponse.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCdpsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCdpsResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.cdps.push(CDPResponse.decode(reader, reader.uint32()));
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

  fromPartial(object: DeepPartial<QueryCdpsResponse>): QueryCdpsResponse {
    const message = createBaseQueryCdpsResponse();
    message.cdps = object.cdps?.map(e => CDPResponse.fromPartial(e)) || [];
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageResponse.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryDepositsRequest(): QueryDepositsRequest {
  return {
    collateralType: "",
    owner: ""
  };
}

export const QueryDepositsRequest = {
  encode(message: QueryDepositsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.collateralType !== "") {
      writer.uint32(10).string(message.collateralType);
    }

    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryDepositsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryDepositsRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.collateralType = reader.string();
          break;

        case 2:
          message.owner = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryDepositsRequest>): QueryDepositsRequest {
    const message = createBaseQueryDepositsRequest();
    message.collateralType = object.collateralType ?? "";
    message.owner = object.owner ?? "";
    return message;
  }

};

function createBaseQueryDepositsResponse(): QueryDepositsResponse {
  return {
    deposits: []
  };
}

export const QueryDepositsResponse = {
  encode(message: QueryDepositsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.deposits) {
      Deposit.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryDepositsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryDepositsResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.deposits.push(Deposit.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryDepositsResponse>): QueryDepositsResponse {
    const message = createBaseQueryDepositsResponse();
    message.deposits = object.deposits?.map(e => Deposit.fromPartial(e)) || [];
    return message;
  }

};

function createBaseQueryTotalPrincipalRequest(): QueryTotalPrincipalRequest {
  return {
    collateralType: ""
  };
}

export const QueryTotalPrincipalRequest = {
  encode(message: QueryTotalPrincipalRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.collateralType !== "") {
      writer.uint32(10).string(message.collateralType);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryTotalPrincipalRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryTotalPrincipalRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.collateralType = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryTotalPrincipalRequest>): QueryTotalPrincipalRequest {
    const message = createBaseQueryTotalPrincipalRequest();
    message.collateralType = object.collateralType ?? "";
    return message;
  }

};

function createBaseQueryTotalPrincipalResponse(): QueryTotalPrincipalResponse {
  return {
    totalPrincipal: []
  };
}

export const QueryTotalPrincipalResponse = {
  encode(message: QueryTotalPrincipalResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.totalPrincipal) {
      TotalPrincipal.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryTotalPrincipalResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryTotalPrincipalResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.totalPrincipal.push(TotalPrincipal.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryTotalPrincipalResponse>): QueryTotalPrincipalResponse {
    const message = createBaseQueryTotalPrincipalResponse();
    message.totalPrincipal = object.totalPrincipal?.map(e => TotalPrincipal.fromPartial(e)) || [];
    return message;
  }

};

function createBaseQueryTotalCollateralRequest(): QueryTotalCollateralRequest {
  return {
    collateralType: ""
  };
}

export const QueryTotalCollateralRequest = {
  encode(message: QueryTotalCollateralRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.collateralType !== "") {
      writer.uint32(10).string(message.collateralType);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryTotalCollateralRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryTotalCollateralRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.collateralType = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryTotalCollateralRequest>): QueryTotalCollateralRequest {
    const message = createBaseQueryTotalCollateralRequest();
    message.collateralType = object.collateralType ?? "";
    return message;
  }

};

function createBaseQueryTotalCollateralResponse(): QueryTotalCollateralResponse {
  return {
    totalCollateral: []
  };
}

export const QueryTotalCollateralResponse = {
  encode(message: QueryTotalCollateralResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.totalCollateral) {
      TotalCollateral.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryTotalCollateralResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryTotalCollateralResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.totalCollateral.push(TotalCollateral.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryTotalCollateralResponse>): QueryTotalCollateralResponse {
    const message = createBaseQueryTotalCollateralResponse();
    message.totalCollateral = object.totalCollateral?.map(e => TotalCollateral.fromPartial(e)) || [];
    return message;
  }

};

function createBaseCDPResponse(): CDPResponse {
  return {
    id: Long.UZERO,
    owner: "",
    type: "",
    collateral: undefined,
    principal: undefined,
    accumulatedFees: undefined,
    feesUpdated: undefined,
    interestFactor: "",
    collateralValue: undefined,
    collateralizationRatio: ""
  };
}

export const CDPResponse = {
  encode(message: CDPResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (!message.id.isZero()) {
      writer.uint32(8).uint64(message.id);
    }

    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }

    if (message.type !== "") {
      writer.uint32(26).string(message.type);
    }

    if (message.collateral !== undefined) {
      Coin.encode(message.collateral, writer.uint32(34).fork()).ldelim();
    }

    if (message.principal !== undefined) {
      Coin.encode(message.principal, writer.uint32(42).fork()).ldelim();
    }

    if (message.accumulatedFees !== undefined) {
      Coin.encode(message.accumulatedFees, writer.uint32(50).fork()).ldelim();
    }

    if (message.feesUpdated !== undefined) {
      Timestamp.encode(toTimestamp(message.feesUpdated), writer.uint32(58).fork()).ldelim();
    }

    if (message.interestFactor !== "") {
      writer.uint32(66).string(message.interestFactor);
    }

    if (message.collateralValue !== undefined) {
      Coin.encode(message.collateralValue, writer.uint32(74).fork()).ldelim();
    }

    if (message.collateralizationRatio !== "") {
      writer.uint32(82).string(message.collateralizationRatio);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CDPResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCDPResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.id = (reader.uint64() as Long);
          break;

        case 2:
          message.owner = reader.string();
          break;

        case 3:
          message.type = reader.string();
          break;

        case 4:
          message.collateral = Coin.decode(reader, reader.uint32());
          break;

        case 5:
          message.principal = Coin.decode(reader, reader.uint32());
          break;

        case 6:
          message.accumulatedFees = Coin.decode(reader, reader.uint32());
          break;

        case 7:
          message.feesUpdated = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;

        case 8:
          message.interestFactor = reader.string();
          break;

        case 9:
          message.collateralValue = Coin.decode(reader, reader.uint32());
          break;

        case 10:
          message.collateralizationRatio = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<CDPResponse>): CDPResponse {
    const message = createBaseCDPResponse();
    message.id = object.id !== undefined && object.id !== null ? Long.fromValue(object.id) : Long.UZERO;
    message.owner = object.owner ?? "";
    message.type = object.type ?? "";
    message.collateral = object.collateral !== undefined && object.collateral !== null ? Coin.fromPartial(object.collateral) : undefined;
    message.principal = object.principal !== undefined && object.principal !== null ? Coin.fromPartial(object.principal) : undefined;
    message.accumulatedFees = object.accumulatedFees !== undefined && object.accumulatedFees !== null ? Coin.fromPartial(object.accumulatedFees) : undefined;
    message.feesUpdated = object.feesUpdated ?? undefined;
    message.interestFactor = object.interestFactor ?? "";
    message.collateralValue = object.collateralValue !== undefined && object.collateralValue !== null ? Coin.fromPartial(object.collateralValue) : undefined;
    message.collateralizationRatio = object.collateralizationRatio ?? "";
    return message;
  }

};