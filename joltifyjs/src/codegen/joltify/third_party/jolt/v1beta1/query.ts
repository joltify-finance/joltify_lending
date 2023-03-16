import { PageRequest, PageRequestSDKType, PageResponse, PageResponseSDKType } from "../../../../cosmos/base/query/v1beta1/pagination";
import { Params, ParamsSDKType } from "./jolt";
import { ModuleAccount, ModuleAccountSDKType } from "../../../../cosmos/auth/v1beta1/auth";
import { Coin, CoinSDKType } from "../../../../cosmos/base/v1beta1/coin";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../../../helpers";
/** QueryParamsRequest is the request type for the Query/Params RPC method. */

export interface QueryParamsRequest {}
/** QueryParamsRequest is the request type for the Query/Params RPC method. */

export interface QueryParamsRequestSDKType {}
/** QueryParamsResponse is the response type for the Query/Params RPC method. */

export interface QueryParamsResponse {
  params?: Params;
}
/** QueryParamsResponse is the response type for the Query/Params RPC method. */

export interface QueryParamsResponseSDKType {
  params?: ParamsSDKType;
}
/** QueryAccountsRequest is the request type for the Query/Accounts RPC method. */

export interface QueryAccountsRequest {}
/** QueryAccountsRequest is the request type for the Query/Accounts RPC method. */

export interface QueryAccountsRequestSDKType {}
/** QueryAccountsResponse is the response type for the Query/Accounts RPC method. */

export interface QueryAccountsResponse {
  accounts: ModuleAccount[];
}
/** QueryAccountsResponse is the response type for the Query/Accounts RPC method. */

export interface QueryAccountsResponseSDKType {
  accounts: ModuleAccountSDKType[];
}
/** QueryDepositsRequest is the request type for the Query/Deposits RPC method. */

export interface QueryDepositsRequest {
  denom: string;
  owner: string;
  pagination?: PageRequest;
}
/** QueryDepositsRequest is the request type for the Query/Deposits RPC method. */

export interface QueryDepositsRequestSDKType {
  denom: string;
  owner: string;
  pagination?: PageRequestSDKType;
}
/** QueryDepositsResponse is the response type for the Query/Deposits RPC method. */

export interface QueryDepositsResponse {
  deposits: DepositResponse[];
  pagination?: PageResponse;
}
/** QueryDepositsResponse is the response type for the Query/Deposits RPC method. */

export interface QueryDepositsResponseSDKType {
  deposits: DepositResponseSDKType[];
  pagination?: PageResponseSDKType;
}
/** QueryUnsyncedDepositsRequest is the request type for the Query/UnsyncedDeposits RPC method. */

export interface QueryUnsyncedDepositsRequest {
  denom: string;
  owner: string;
  pagination?: PageRequest;
}
/** QueryUnsyncedDepositsRequest is the request type for the Query/UnsyncedDeposits RPC method. */

export interface QueryUnsyncedDepositsRequestSDKType {
  denom: string;
  owner: string;
  pagination?: PageRequestSDKType;
}
/** QueryUnsyncedDepositsResponse is the response type for the Query/UnsyncedDeposits RPC method. */

export interface QueryUnsyncedDepositsResponse {
  deposits: DepositResponse[];
  pagination?: PageResponse;
}
/** QueryUnsyncedDepositsResponse is the response type for the Query/UnsyncedDeposits RPC method. */

export interface QueryUnsyncedDepositsResponseSDKType {
  deposits: DepositResponseSDKType[];
  pagination?: PageResponseSDKType;
}
/** QueryTotalDepositedRequest is the request type for the Query/TotalDeposited RPC method. */

export interface QueryTotalDepositedRequest {
  denom: string;
}
/** QueryTotalDepositedRequest is the request type for the Query/TotalDeposited RPC method. */

export interface QueryTotalDepositedRequestSDKType {
  denom: string;
}
/** QueryTotalDepositedResponse is the response type for the Query/TotalDeposited RPC method. */

export interface QueryTotalDepositedResponse {
  suppliedCoins: Coin[];
}
/** QueryTotalDepositedResponse is the response type for the Query/TotalDeposited RPC method. */

export interface QueryTotalDepositedResponseSDKType {
  supplied_coins: CoinSDKType[];
}
/** QueryBorrowsRequest is the request type for the Query/Borrows RPC method. */

export interface QueryBorrowsRequest {
  denom: string;
  owner: string;
  pagination?: PageRequest;
}
/** QueryBorrowsRequest is the request type for the Query/Borrows RPC method. */

export interface QueryBorrowsRequestSDKType {
  denom: string;
  owner: string;
  pagination?: PageRequestSDKType;
}
export interface QueryLiquidateRequest {
  borrower: string;
  pagination?: PageRequest;
}
export interface QueryLiquidateRequestSDKType {
  borrower: string;
  pagination?: PageRequestSDKType;
}
export interface LiquidateItem {
  owner: string;
  ltv: string;
}
export interface LiquidateItemSDKType {
  owner: string;
  ltv: string;
}
export interface QueryLiquidateResp {
  liquidateItems: LiquidateItem[];
  pagination?: PageResponse;
}
export interface QueryLiquidateRespSDKType {
  liquidateItems: LiquidateItemSDKType[];
  pagination?: PageResponseSDKType;
}
/** QueryBorrowsResponse is the response type for the Query/Borrows RPC method. */

export interface QueryBorrowsResponse {
  borrows: BorrowResponse[];
  pagination?: PageResponse;
}
/** QueryBorrowsResponse is the response type for the Query/Borrows RPC method. */

export interface QueryBorrowsResponseSDKType {
  borrows: BorrowResponseSDKType[];
  pagination?: PageResponseSDKType;
}
/** QueryUnsyncedBorrowsRequest is the request type for the Query/UnsyncedBorrows RPC method. */

export interface QueryUnsyncedBorrowsRequest {
  denom: string;
  owner: string;
  pagination?: PageRequest;
}
/** QueryUnsyncedBorrowsRequest is the request type for the Query/UnsyncedBorrows RPC method. */

export interface QueryUnsyncedBorrowsRequestSDKType {
  denom: string;
  owner: string;
  pagination?: PageRequestSDKType;
}
/** QueryUnsyncedBorrowsResponse is the response type for the Query/UnsyncedBorrows RPC method. */

export interface QueryUnsyncedBorrowsResponse {
  borrows: BorrowResponse[];
  pagination?: PageResponse;
}
/** QueryUnsyncedBorrowsResponse is the response type for the Query/UnsyncedBorrows RPC method. */

export interface QueryUnsyncedBorrowsResponseSDKType {
  borrows: BorrowResponseSDKType[];
  pagination?: PageResponseSDKType;
}
/** QueryTotalBorrowedRequest is the request type for the Query/TotalBorrowed RPC method. */

export interface QueryTotalBorrowedRequest {
  denom: string;
}
/** QueryTotalBorrowedRequest is the request type for the Query/TotalBorrowed RPC method. */

export interface QueryTotalBorrowedRequestSDKType {
  denom: string;
}
/** QueryTotalBorrowedResponse is the response type for the Query/TotalBorrowed RPC method. */

export interface QueryTotalBorrowedResponse {
  borrowedCoins: Coin[];
}
/** QueryTotalBorrowedResponse is the response type for the Query/TotalBorrowed RPC method. */

export interface QueryTotalBorrowedResponseSDKType {
  borrowed_coins: CoinSDKType[];
}
/** QueryInterestRateRequest is the request type for the Query/InterestRate RPC method. */

export interface QueryInterestRateRequest {
  denom: string;
}
/** QueryInterestRateRequest is the request type for the Query/InterestRate RPC method. */

export interface QueryInterestRateRequestSDKType {
  denom: string;
}
/** QueryInterestRateResponse is the response type for the Query/InterestRate RPC method. */

export interface QueryInterestRateResponse {
  interestRates: MoneyMarketInterestRate[];
}
/** QueryInterestRateResponse is the response type for the Query/InterestRate RPC method. */

export interface QueryInterestRateResponseSDKType {
  interest_rates: MoneyMarketInterestRateSDKType[];
}
/** QueryReservesRequest is the request type for the Query/Reserves RPC method. */

export interface QueryReservesRequest {
  denom: string;
}
/** QueryReservesRequest is the request type for the Query/Reserves RPC method. */

export interface QueryReservesRequestSDKType {
  denom: string;
}
/** QueryReservesResponse is the response type for the Query/Reserves RPC method. */

export interface QueryReservesResponse {
  amount: Coin[];
}
/** QueryReservesResponse is the response type for the Query/Reserves RPC method. */

export interface QueryReservesResponseSDKType {
  amount: CoinSDKType[];
}
/** QueryInterestFactorsRequest is the request type for the Query/InterestFactors RPC method. */

export interface QueryInterestFactorsRequest {
  denom: string;
}
/** QueryInterestFactorsRequest is the request type for the Query/InterestFactors RPC method. */

export interface QueryInterestFactorsRequestSDKType {
  denom: string;
}
/** QueryInterestFactorsResponse is the response type for the Query/InterestFactors RPC method. */

export interface QueryInterestFactorsResponse {
  interestFactors: InterestFactor[];
}
/** QueryInterestFactorsResponse is the response type for the Query/InterestFactors RPC method. */

export interface QueryInterestFactorsResponseSDKType {
  interest_factors: InterestFactorSDKType[];
}
/** DepositResponse defines an amount of coins deposited into a jolt module account. */

export interface DepositResponse {
  depositor: string;
  amount: Coin[];
  index: SupplyInterestFactorResponse[];
}
/** DepositResponse defines an amount of coins deposited into a jolt module account. */

export interface DepositResponseSDKType {
  depositor: string;
  amount: CoinSDKType[];
  index: SupplyInterestFactorResponseSDKType[];
}
/** SupplyInterestFactorResponse defines an individual borrow interest factor. */

export interface SupplyInterestFactorResponse {
  denom: string;
  /** sdk.Dec as string */

  value: string;
}
/** SupplyInterestFactorResponse defines an individual borrow interest factor. */

export interface SupplyInterestFactorResponseSDKType {
  denom: string;
  value: string;
}
/** BorrowResponse defines an amount of coins borrowed from a jolt module account. */

export interface BorrowResponse {
  borrower: string;
  amount: Coin[];
  index: BorrowInterestFactorResponse[];
}
/** BorrowResponse defines an amount of coins borrowed from a jolt module account. */

export interface BorrowResponseSDKType {
  borrower: string;
  amount: CoinSDKType[];
  index: BorrowInterestFactorResponseSDKType[];
}
/** BorrowInterestFactorResponse defines an individual borrow interest factor. */

export interface BorrowInterestFactorResponse {
  denom: string;
  /** sdk.Dec as string */

  value: string;
}
/** BorrowInterestFactorResponse defines an individual borrow interest factor. */

export interface BorrowInterestFactorResponseSDKType {
  denom: string;
  value: string;
}
/** MoneyMarketInterestRate is a unique type returned by interest rate queries */

export interface MoneyMarketInterestRate {
  denom: string;
  /** sdk.Dec as String */

  supplyInterestRate: string;
  /** sdk.Dec as String */

  borrowInterestRate: string;
}
/** MoneyMarketInterestRate is a unique type returned by interest rate queries */

export interface MoneyMarketInterestRateSDKType {
  denom: string;
  supply_interest_rate: string;
  borrow_interest_rate: string;
}
/** InterestFactor is a unique type returned by interest factor queries */

export interface InterestFactor {
  denom: string;
  /** sdk.Dec as String */

  borrowInterestFactor: string;
  /** sdk.Dec as String */

  supplyInterestFactor: string;
}
/** InterestFactor is a unique type returned by interest factor queries */

export interface InterestFactorSDKType {
  denom: string;
  borrow_interest_factor: string;
  supply_interest_factor: string;
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

function createBaseQueryDepositsRequest(): QueryDepositsRequest {
  return {
    denom: "",
    owner: "",
    pagination: undefined
  };
}

export const QueryDepositsRequest = {
  encode(message: QueryDepositsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }

    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }

    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(26).fork()).ldelim();
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
          message.denom = reader.string();
          break;

        case 2:
          message.owner = reader.string();
          break;

        case 3:
          message.pagination = PageRequest.decode(reader, reader.uint32());
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
    message.denom = object.denom ?? "";
    message.owner = object.owner ?? "";
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageRequest.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryDepositsResponse(): QueryDepositsResponse {
  return {
    deposits: [],
    pagination: undefined
  };
}

export const QueryDepositsResponse = {
  encode(message: QueryDepositsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.deposits) {
      DepositResponse.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
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
          message.deposits.push(DepositResponse.decode(reader, reader.uint32()));
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

  fromPartial(object: DeepPartial<QueryDepositsResponse>): QueryDepositsResponse {
    const message = createBaseQueryDepositsResponse();
    message.deposits = object.deposits?.map(e => DepositResponse.fromPartial(e)) || [];
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageResponse.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryUnsyncedDepositsRequest(): QueryUnsyncedDepositsRequest {
  return {
    denom: "",
    owner: "",
    pagination: undefined
  };
}

export const QueryUnsyncedDepositsRequest = {
  encode(message: QueryUnsyncedDepositsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }

    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }

    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(26).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryUnsyncedDepositsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryUnsyncedDepositsRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;

        case 2:
          message.owner = reader.string();
          break;

        case 3:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryUnsyncedDepositsRequest>): QueryUnsyncedDepositsRequest {
    const message = createBaseQueryUnsyncedDepositsRequest();
    message.denom = object.denom ?? "";
    message.owner = object.owner ?? "";
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageRequest.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryUnsyncedDepositsResponse(): QueryUnsyncedDepositsResponse {
  return {
    deposits: [],
    pagination: undefined
  };
}

export const QueryUnsyncedDepositsResponse = {
  encode(message: QueryUnsyncedDepositsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.deposits) {
      DepositResponse.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryUnsyncedDepositsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryUnsyncedDepositsResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.deposits.push(DepositResponse.decode(reader, reader.uint32()));
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

  fromPartial(object: DeepPartial<QueryUnsyncedDepositsResponse>): QueryUnsyncedDepositsResponse {
    const message = createBaseQueryUnsyncedDepositsResponse();
    message.deposits = object.deposits?.map(e => DepositResponse.fromPartial(e)) || [];
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageResponse.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryTotalDepositedRequest(): QueryTotalDepositedRequest {
  return {
    denom: ""
  };
}

export const QueryTotalDepositedRequest = {
  encode(message: QueryTotalDepositedRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryTotalDepositedRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryTotalDepositedRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryTotalDepositedRequest>): QueryTotalDepositedRequest {
    const message = createBaseQueryTotalDepositedRequest();
    message.denom = object.denom ?? "";
    return message;
  }

};

function createBaseQueryTotalDepositedResponse(): QueryTotalDepositedResponse {
  return {
    suppliedCoins: []
  };
}

export const QueryTotalDepositedResponse = {
  encode(message: QueryTotalDepositedResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.suppliedCoins) {
      Coin.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryTotalDepositedResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryTotalDepositedResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 2:
          message.suppliedCoins.push(Coin.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryTotalDepositedResponse>): QueryTotalDepositedResponse {
    const message = createBaseQueryTotalDepositedResponse();
    message.suppliedCoins = object.suppliedCoins?.map(e => Coin.fromPartial(e)) || [];
    return message;
  }

};

function createBaseQueryBorrowsRequest(): QueryBorrowsRequest {
  return {
    denom: "",
    owner: "",
    pagination: undefined
  };
}

export const QueryBorrowsRequest = {
  encode(message: QueryBorrowsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }

    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }

    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(26).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryBorrowsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryBorrowsRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;

        case 2:
          message.owner = reader.string();
          break;

        case 3:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryBorrowsRequest>): QueryBorrowsRequest {
    const message = createBaseQueryBorrowsRequest();
    message.denom = object.denom ?? "";
    message.owner = object.owner ?? "";
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageRequest.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryLiquidateRequest(): QueryLiquidateRequest {
  return {
    borrower: "",
    pagination: undefined
  };
}

export const QueryLiquidateRequest = {
  encode(message: QueryLiquidateRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.borrower !== "") {
      writer.uint32(10).string(message.borrower);
    }

    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryLiquidateRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryLiquidateRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.borrower = reader.string();
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

  fromPartial(object: DeepPartial<QueryLiquidateRequest>): QueryLiquidateRequest {
    const message = createBaseQueryLiquidateRequest();
    message.borrower = object.borrower ?? "";
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageRequest.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseLiquidateItem(): LiquidateItem {
  return {
    owner: "",
    ltv: ""
  };
}

export const LiquidateItem = {
  encode(message: LiquidateItem, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }

    if (message.ltv !== "") {
      writer.uint32(18).string(message.ltv);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LiquidateItem {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLiquidateItem();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;

        case 2:
          message.ltv = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<LiquidateItem>): LiquidateItem {
    const message = createBaseLiquidateItem();
    message.owner = object.owner ?? "";
    message.ltv = object.ltv ?? "";
    return message;
  }

};

function createBaseQueryLiquidateResp(): QueryLiquidateResp {
  return {
    liquidateItems: [],
    pagination: undefined
  };
}

export const QueryLiquidateResp = {
  encode(message: QueryLiquidateResp, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.liquidateItems) {
      LiquidateItem.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryLiquidateResp {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryLiquidateResp();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.liquidateItems.push(LiquidateItem.decode(reader, reader.uint32()));
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

  fromPartial(object: DeepPartial<QueryLiquidateResp>): QueryLiquidateResp {
    const message = createBaseQueryLiquidateResp();
    message.liquidateItems = object.liquidateItems?.map(e => LiquidateItem.fromPartial(e)) || [];
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageResponse.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryBorrowsResponse(): QueryBorrowsResponse {
  return {
    borrows: [],
    pagination: undefined
  };
}

export const QueryBorrowsResponse = {
  encode(message: QueryBorrowsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.borrows) {
      BorrowResponse.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryBorrowsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryBorrowsResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.borrows.push(BorrowResponse.decode(reader, reader.uint32()));
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

  fromPartial(object: DeepPartial<QueryBorrowsResponse>): QueryBorrowsResponse {
    const message = createBaseQueryBorrowsResponse();
    message.borrows = object.borrows?.map(e => BorrowResponse.fromPartial(e)) || [];
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageResponse.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryUnsyncedBorrowsRequest(): QueryUnsyncedBorrowsRequest {
  return {
    denom: "",
    owner: "",
    pagination: undefined
  };
}

export const QueryUnsyncedBorrowsRequest = {
  encode(message: QueryUnsyncedBorrowsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }

    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }

    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(26).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryUnsyncedBorrowsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryUnsyncedBorrowsRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;

        case 2:
          message.owner = reader.string();
          break;

        case 3:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryUnsyncedBorrowsRequest>): QueryUnsyncedBorrowsRequest {
    const message = createBaseQueryUnsyncedBorrowsRequest();
    message.denom = object.denom ?? "";
    message.owner = object.owner ?? "";
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageRequest.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryUnsyncedBorrowsResponse(): QueryUnsyncedBorrowsResponse {
  return {
    borrows: [],
    pagination: undefined
  };
}

export const QueryUnsyncedBorrowsResponse = {
  encode(message: QueryUnsyncedBorrowsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.borrows) {
      BorrowResponse.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryUnsyncedBorrowsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryUnsyncedBorrowsResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.borrows.push(BorrowResponse.decode(reader, reader.uint32()));
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

  fromPartial(object: DeepPartial<QueryUnsyncedBorrowsResponse>): QueryUnsyncedBorrowsResponse {
    const message = createBaseQueryUnsyncedBorrowsResponse();
    message.borrows = object.borrows?.map(e => BorrowResponse.fromPartial(e)) || [];
    message.pagination = object.pagination !== undefined && object.pagination !== null ? PageResponse.fromPartial(object.pagination) : undefined;
    return message;
  }

};

function createBaseQueryTotalBorrowedRequest(): QueryTotalBorrowedRequest {
  return {
    denom: ""
  };
}

export const QueryTotalBorrowedRequest = {
  encode(message: QueryTotalBorrowedRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryTotalBorrowedRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryTotalBorrowedRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryTotalBorrowedRequest>): QueryTotalBorrowedRequest {
    const message = createBaseQueryTotalBorrowedRequest();
    message.denom = object.denom ?? "";
    return message;
  }

};

function createBaseQueryTotalBorrowedResponse(): QueryTotalBorrowedResponse {
  return {
    borrowedCoins: []
  };
}

export const QueryTotalBorrowedResponse = {
  encode(message: QueryTotalBorrowedResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.borrowedCoins) {
      Coin.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryTotalBorrowedResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryTotalBorrowedResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 2:
          message.borrowedCoins.push(Coin.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryTotalBorrowedResponse>): QueryTotalBorrowedResponse {
    const message = createBaseQueryTotalBorrowedResponse();
    message.borrowedCoins = object.borrowedCoins?.map(e => Coin.fromPartial(e)) || [];
    return message;
  }

};

function createBaseQueryInterestRateRequest(): QueryInterestRateRequest {
  return {
    denom: ""
  };
}

export const QueryInterestRateRequest = {
  encode(message: QueryInterestRateRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryInterestRateRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryInterestRateRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryInterestRateRequest>): QueryInterestRateRequest {
    const message = createBaseQueryInterestRateRequest();
    message.denom = object.denom ?? "";
    return message;
  }

};

function createBaseQueryInterestRateResponse(): QueryInterestRateResponse {
  return {
    interestRates: []
  };
}

export const QueryInterestRateResponse = {
  encode(message: QueryInterestRateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.interestRates) {
      MoneyMarketInterestRate.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryInterestRateResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryInterestRateResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.interestRates.push(MoneyMarketInterestRate.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryInterestRateResponse>): QueryInterestRateResponse {
    const message = createBaseQueryInterestRateResponse();
    message.interestRates = object.interestRates?.map(e => MoneyMarketInterestRate.fromPartial(e)) || [];
    return message;
  }

};

function createBaseQueryReservesRequest(): QueryReservesRequest {
  return {
    denom: ""
  };
}

export const QueryReservesRequest = {
  encode(message: QueryReservesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryReservesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryReservesRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryReservesRequest>): QueryReservesRequest {
    const message = createBaseQueryReservesRequest();
    message.denom = object.denom ?? "";
    return message;
  }

};

function createBaseQueryReservesResponse(): QueryReservesResponse {
  return {
    amount: []
  };
}

export const QueryReservesResponse = {
  encode(message: QueryReservesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.amount) {
      Coin.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryReservesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryReservesResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 2:
          message.amount.push(Coin.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryReservesResponse>): QueryReservesResponse {
    const message = createBaseQueryReservesResponse();
    message.amount = object.amount?.map(e => Coin.fromPartial(e)) || [];
    return message;
  }

};

function createBaseQueryInterestFactorsRequest(): QueryInterestFactorsRequest {
  return {
    denom: ""
  };
}

export const QueryInterestFactorsRequest = {
  encode(message: QueryInterestFactorsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryInterestFactorsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryInterestFactorsRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryInterestFactorsRequest>): QueryInterestFactorsRequest {
    const message = createBaseQueryInterestFactorsRequest();
    message.denom = object.denom ?? "";
    return message;
  }

};

function createBaseQueryInterestFactorsResponse(): QueryInterestFactorsResponse {
  return {
    interestFactors: []
  };
}

export const QueryInterestFactorsResponse = {
  encode(message: QueryInterestFactorsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.interestFactors) {
      InterestFactor.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryInterestFactorsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryInterestFactorsResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.interestFactors.push(InterestFactor.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryInterestFactorsResponse>): QueryInterestFactorsResponse {
    const message = createBaseQueryInterestFactorsResponse();
    message.interestFactors = object.interestFactors?.map(e => InterestFactor.fromPartial(e)) || [];
    return message;
  }

};

function createBaseDepositResponse(): DepositResponse {
  return {
    depositor: "",
    amount: [],
    index: []
  };
}

export const DepositResponse = {
  encode(message: DepositResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.depositor !== "") {
      writer.uint32(10).string(message.depositor);
    }

    for (const v of message.amount) {
      Coin.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    for (const v of message.index) {
      SupplyInterestFactorResponse.encode(v!, writer.uint32(26).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DepositResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDepositResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.depositor = reader.string();
          break;

        case 2:
          message.amount.push(Coin.decode(reader, reader.uint32()));
          break;

        case 3:
          message.index.push(SupplyInterestFactorResponse.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<DepositResponse>): DepositResponse {
    const message = createBaseDepositResponse();
    message.depositor = object.depositor ?? "";
    message.amount = object.amount?.map(e => Coin.fromPartial(e)) || [];
    message.index = object.index?.map(e => SupplyInterestFactorResponse.fromPartial(e)) || [];
    return message;
  }

};

function createBaseSupplyInterestFactorResponse(): SupplyInterestFactorResponse {
  return {
    denom: "",
    value: ""
  };
}

export const SupplyInterestFactorResponse = {
  encode(message: SupplyInterestFactorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }

    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SupplyInterestFactorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSupplyInterestFactorResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;

        case 2:
          message.value = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<SupplyInterestFactorResponse>): SupplyInterestFactorResponse {
    const message = createBaseSupplyInterestFactorResponse();
    message.denom = object.denom ?? "";
    message.value = object.value ?? "";
    return message;
  }

};

function createBaseBorrowResponse(): BorrowResponse {
  return {
    borrower: "",
    amount: [],
    index: []
  };
}

export const BorrowResponse = {
  encode(message: BorrowResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.borrower !== "") {
      writer.uint32(10).string(message.borrower);
    }

    for (const v of message.amount) {
      Coin.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    for (const v of message.index) {
      BorrowInterestFactorResponse.encode(v!, writer.uint32(26).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BorrowResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBorrowResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.borrower = reader.string();
          break;

        case 2:
          message.amount.push(Coin.decode(reader, reader.uint32()));
          break;

        case 3:
          message.index.push(BorrowInterestFactorResponse.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<BorrowResponse>): BorrowResponse {
    const message = createBaseBorrowResponse();
    message.borrower = object.borrower ?? "";
    message.amount = object.amount?.map(e => Coin.fromPartial(e)) || [];
    message.index = object.index?.map(e => BorrowInterestFactorResponse.fromPartial(e)) || [];
    return message;
  }

};

function createBaseBorrowInterestFactorResponse(): BorrowInterestFactorResponse {
  return {
    denom: "",
    value: ""
  };
}

export const BorrowInterestFactorResponse = {
  encode(message: BorrowInterestFactorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }

    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BorrowInterestFactorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBorrowInterestFactorResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;

        case 2:
          message.value = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<BorrowInterestFactorResponse>): BorrowInterestFactorResponse {
    const message = createBaseBorrowInterestFactorResponse();
    message.denom = object.denom ?? "";
    message.value = object.value ?? "";
    return message;
  }

};

function createBaseMoneyMarketInterestRate(): MoneyMarketInterestRate {
  return {
    denom: "",
    supplyInterestRate: "",
    borrowInterestRate: ""
  };
}

export const MoneyMarketInterestRate = {
  encode(message: MoneyMarketInterestRate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }

    if (message.supplyInterestRate !== "") {
      writer.uint32(18).string(message.supplyInterestRate);
    }

    if (message.borrowInterestRate !== "") {
      writer.uint32(26).string(message.borrowInterestRate);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MoneyMarketInterestRate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMoneyMarketInterestRate();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;

        case 2:
          message.supplyInterestRate = reader.string();
          break;

        case 3:
          message.borrowInterestRate = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MoneyMarketInterestRate>): MoneyMarketInterestRate {
    const message = createBaseMoneyMarketInterestRate();
    message.denom = object.denom ?? "";
    message.supplyInterestRate = object.supplyInterestRate ?? "";
    message.borrowInterestRate = object.borrowInterestRate ?? "";
    return message;
  }

};

function createBaseInterestFactor(): InterestFactor {
  return {
    denom: "",
    borrowInterestFactor: "",
    supplyInterestFactor: ""
  };
}

export const InterestFactor = {
  encode(message: InterestFactor, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }

    if (message.borrowInterestFactor !== "") {
      writer.uint32(18).string(message.borrowInterestFactor);
    }

    if (message.supplyInterestFactor !== "") {
      writer.uint32(26).string(message.supplyInterestFactor);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): InterestFactor {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInterestFactor();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;

        case 2:
          message.borrowInterestFactor = reader.string();
          break;

        case 3:
          message.supplyInterestFactor = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<InterestFactor>): InterestFactor {
    const message = createBaseInterestFactor();
    message.denom = object.denom ?? "";
    message.borrowInterestFactor = object.borrowInterestFactor ?? "";
    message.supplyInterestFactor = object.supplyInterestFactor ?? "";
    return message;
  }

};