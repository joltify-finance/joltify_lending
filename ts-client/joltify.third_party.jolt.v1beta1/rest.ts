/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

/**
* `Any` contains an arbitrary serialized protocol buffer message along with a
URL that describes the type of the serialized message.

Protobuf library provides support to pack/unpack Any values in the form
of utility functions or additional generated methods of the Any type.

Example 1: Pack and unpack a message in C++.

    Foo foo = ...;
    Any any;
    any.PackFrom(foo);
    ...
    if (any.UnpackTo(&foo)) {
      ...
    }

Example 2: Pack and unpack a message in Java.

    Foo foo = ...;
    Any any = Any.pack(foo);
    ...
    if (any.is(Foo.class)) {
      foo = any.unpack(Foo.class);
    }

 Example 3: Pack and unpack a message in Python.

    foo = Foo(...)
    any = Any()
    any.Pack(foo)
    ...
    if any.Is(Foo.DESCRIPTOR):
      any.Unpack(foo)
      ...

 Example 4: Pack and unpack a message in Go

     foo := &pb.Foo{...}
     any, err := anypb.New(foo)
     if err != nil {
       ...
     }
     ...
     foo := &pb.Foo{}
     if err := any.UnmarshalTo(foo); err != nil {
       ...
     }

The pack methods provided by protobuf library will by default use
'type.googleapis.com/full.type.name' as the type URL and the unpack
methods only use the fully qualified type name after the last '/'
in the type URL, for example "foo.bar.com/x/y.z" will yield type
name "y.z".


JSON
====
The JSON representation of an `Any` value uses the regular
representation of the deserialized, embedded message, with an
additional field `@type` which contains the type URL. Example:

    package google.profile;
    message Person {
      string first_name = 1;
      string last_name = 2;
    }

    {
      "@type": "type.googleapis.com/google.profile.Person",
      "firstName": <string>,
      "lastName": <string>
    }

If the embedded message type is well-known and has a custom JSON
representation, that representation will be embedded adding a field
`value` which holds the custom JSON in addition to the `@type`
field. Example (for message [google.protobuf.Duration][]):

    {
      "@type": "type.googleapis.com/google.protobuf.Duration",
      "value": "1.212s"
    }
*/
export interface ProtobufAny {
  /**
   * A URL/resource name that uniquely identifies the type of the serialized
   * protocol buffer message. This string must contain at least
   * one "/" character. The last segment of the URL's path must represent
   * the fully qualified name of the type (as in
   * `path/google.protobuf.Duration`). The name should be in a canonical form
   * (e.g., leading "." is not accepted).
   *
   * In practice, teams usually precompile into the binary all types that they
   * expect it to use in the context of Any. However, for URLs which use the
   * scheme `http`, `https`, or no scheme, one can optionally set up a type
   * server that maps type URLs to message definitions as follows:
   * * If no scheme is provided, `https` is assumed.
   * * An HTTP GET on the URL must yield a [google.protobuf.Type][]
   *   value in binary format, or produce an error.
   * * Applications are allowed to cache lookup results based on the
   *   URL, or have them precompiled into a binary to avoid any
   *   lookup. Therefore, binary compatibility needs to be preserved
   *   on changes to types. (Use versioned type names to manage
   *   breaking changes.)
   * Note: this functionality is not currently available in the official
   * protobuf release, and it is not used for type URLs beginning with
   * type.googleapis.com.
   * Schemes other than `http`, `https` (or the empty scheme) might be
   * used with implementation specific semantics.
   */
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

/**
 * Params defines the parameters for the jolt module.
 */
export interface ThirdPartyjoltv1Beta1Params {
  money_markets?: V1Beta1MoneyMarket[];
  minimum_borrow_usd_value?: string;
  surplus_auction_threshold?: string;
}

/**
* BaseAccount defines a base account type. It contains all the necessary fields
for basic account functionality. Any custom account type should extend this
type for additional functionality (e.g. vesting).
*/
export interface V1Beta1BaseAccount {
  address?: string;

  /**
   * `Any` contains an arbitrary serialized protocol buffer message along with a
   * URL that describes the type of the serialized message.
   *
   * Protobuf library provides support to pack/unpack Any values in the form
   * of utility functions or additional generated methods of the Any type.
   * Example 1: Pack and unpack a message in C++.
   *     Foo foo = ...;
   *     Any any;
   *     any.PackFrom(foo);
   *     ...
   *     if (any.UnpackTo(&foo)) {
   *       ...
   *     }
   * Example 2: Pack and unpack a message in Java.
   *     Any any = Any.pack(foo);
   *     if (any.is(Foo.class)) {
   *       foo = any.unpack(Foo.class);
   *  Example 3: Pack and unpack a message in Python.
   *     foo = Foo(...)
   *     any = Any()
   *     any.Pack(foo)
   *     if any.Is(Foo.DESCRIPTOR):
   *       any.Unpack(foo)
   *  Example 4: Pack and unpack a message in Go
   *      foo := &pb.Foo{...}
   *      any, err := anypb.New(foo)
   *      if err != nil {
   *        ...
   *      }
   *      ...
   *      foo := &pb.Foo{}
   *      if err := any.UnmarshalTo(foo); err != nil {
   * The pack methods provided by protobuf library will by default use
   * 'type.googleapis.com/full.type.name' as the type URL and the unpack
   * methods only use the fully qualified type name after the last '/'
   * in the type URL, for example "foo.bar.com/x/y.z" will yield type
   * name "y.z".
   * JSON
   * ====
   * The JSON representation of an `Any` value uses the regular
   * representation of the deserialized, embedded message, with an
   * additional field `@type` which contains the type URL. Example:
   *     package google.profile;
   *     message Person {
   *       string first_name = 1;
   *       string last_name = 2;
   *     {
   *       "@type": "type.googleapis.com/google.profile.Person",
   *       "firstName": <string>,
   *       "lastName": <string>
   * If the embedded message type is well-known and has a custom JSON
   * representation, that representation will be embedded adding a field
   * `value` which holds the custom JSON in addition to the `@type`
   * field. Example (for message [google.protobuf.Duration][]):
   *       "@type": "type.googleapis.com/google.protobuf.Duration",
   *       "value": "1.212s"
   */
  pub_key?: ProtobufAny;

  /** @format uint64 */
  account_number?: string;

  /** @format uint64 */
  sequence?: string;
}

/**
 * BorrowInterestFactorResponse defines an individual borrow interest factor.
 */
export interface V1Beta1BorrowInterestFactorResponse {
  denom?: string;

  /** sdk.Dec as string */
  value?: string;
}

/**
 * BorrowLimit enforces restrictions on a money market.
 */
export interface V1Beta1BorrowLimit {
  has_max_limit?: boolean;
  maximum_limit?: string;
  loan_to_value?: string;
}

/**
 * BorrowResponse defines an amount of coins borrowed from a jolt module account.
 */
export interface V1Beta1BorrowResponse {
  borrower?: string;
  amount?: V1Beta1Coin[];
  index?: V1Beta1BorrowInterestFactorResponse[];
}

/**
* Coin defines a token with a denomination and an amount.

NOTE: The amount field is an Int which implements the custom method
signatures required by gogoproto.
*/
export interface V1Beta1Coin {
  denom?: string;
  amount?: string;
}

/**
 * DepositResponse defines an amount of coins deposited into a jolt module account.
 */
export interface V1Beta1DepositResponse {
  depositor?: string;
  amount?: V1Beta1Coin[];
  index?: V1Beta1SupplyInterestFactorResponse[];
}

export interface V1Beta1InterestFactor {
  denom?: string;

  /** sdk.Dec as String */
  borrow_interest_factor?: string;

  /** sdk.Dec as String */
  supply_interest_factor?: string;
}

/**
 * InterestRateModel contains information about an asset's interest rate.
 */
export interface V1Beta1InterestRateModel {
  base_rate_apy?: string;
  base_multiplier?: string;
  kink?: string;
  jump_multiplier?: string;
}

export interface V1Beta1LiquidateItem {
  owner?: string;
  ltv?: string;
}

/**
 * ModuleAccount defines an account for modules that holds coins on a pool.
 */
export interface V1Beta1ModuleAccount {
  /**
   * BaseAccount defines a base account type. It contains all the necessary fields
   * for basic account functionality. Any custom account type should extend this
   * type for additional functionality (e.g. vesting).
   */
  base_account?: V1Beta1BaseAccount;
  name?: string;
  permissions?: string[];
}

/**
 * MoneyMarket is a money market for an individual asset.
 */
export interface V1Beta1MoneyMarket {
  denom?: string;

  /** BorrowLimit enforces restrictions on a money market. */
  borrow_limit?: V1Beta1BorrowLimit;
  spot_market_id?: string;
  conversion_factor?: string;

  /** InterestRateModel contains information about an asset's interest rate. */
  interest_rate_model?: V1Beta1InterestRateModel;
  reserve_factor?: string;
  keeper_reward_percentage?: string;
}

export interface V1Beta1MoneyMarketInterestRate {
  denom?: string;

  /** sdk.Dec as String */
  supply_interest_rate?: string;

  /** sdk.Dec as String */
  borrow_interest_rate?: string;
}

/**
 * MsgBorrowResponse defines the Msg/Borrow response type.
 */
export type V1Beta1MsgBorrowResponse = object;

/**
 * MsgDepositResponse defines the Msg/Deposit response type.
 */
export type V1Beta1MsgDepositResponse = object;

/**
 * MsgLiquidateResponse defines the Msg/Liquidate response type.
 */
export type V1Beta1MsgLiquidateResponse = object;

/**
 * MsgRepayResponse defines the Msg/Repay response type.
 */
export type V1Beta1MsgRepayResponse = object;

/**
 * MsgWithdrawResponse defines the Msg/Withdraw response type.
 */
export type V1Beta1MsgWithdrawResponse = object;

/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
  /**
   * key is a value returned in PageResponse.next_key to begin
   * querying the next page most efficiently. Only one of offset or key
   * should be set.
   * @format byte
   */
  key?: string;

  /**
   * offset is a numeric offset that can be used when key is unavailable.
   * It is less efficient than using key. Only one of offset or key should
   * be set.
   * @format uint64
   */
  offset?: string;

  /**
   * limit is the total number of results to be returned in the result page.
   * If left empty it will default to a value to be set by each app.
   * @format uint64
   */
  limit?: string;

  /**
   * count_total is set to true  to indicate that the result set should include
   * a count of the total number of items available for pagination in UIs.
   * count_total is only respected when offset is used. It is ignored when key
   * is set.
   */
  count_total?: boolean;

  /**
   * reverse is set to true if results are to be returned in the descending order.
   *
   * Since: cosmos-sdk 0.43
   */
  reverse?: boolean;
}

/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
  /**
   * next_key is the key to be passed to PageRequest.key to
   * query the next page most efficiently
   * @format byte
   */
  next_key?: string;

  /**
   * total is total number of results available if PageRequest.count_total
   * was set, its value is undefined otherwise
   * @format uint64
   */
  total?: string;
}

/**
 * QueryAccountsResponse is the response type for the Query/Accounts RPC method.
 */
export interface V1Beta1QueryAccountsResponse {
  accounts?: V1Beta1ModuleAccount[];
}

/**
 * QueryBorrowsResponse is the response type for the Query/Borrows RPC method.
 */
export interface V1Beta1QueryBorrowsResponse {
  borrows?: V1Beta1BorrowResponse[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

/**
 * QueryDepositsResponse is the response type for the Query/Deposits RPC method.
 */
export interface V1Beta1QueryDepositsResponse {
  deposits?: V1Beta1DepositResponse[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

/**
 * QueryInterestFactorsResponse is the response type for the Query/InterestFactors RPC method.
 */
export interface V1Beta1QueryInterestFactorsResponse {
  interest_factors?: V1Beta1InterestFactor[];
}

/**
 * QueryInterestRateResponse is the response type for the Query/InterestRate RPC method.
 */
export interface V1Beta1QueryInterestRateResponse {
  interest_rates?: V1Beta1MoneyMarketInterestRate[];
}

export interface V1Beta1QueryLiquidateResp {
  liquidateItems?: V1Beta1LiquidateItem[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

/**
 * QueryParamsResponse is the response type for the Query/Params RPC method.
 */
export interface V1Beta1QueryParamsResponse {
  /** Params defines the parameters for the jolt module. */
  params?: ThirdPartyjoltv1Beta1Params;
}

/**
 * QueryReservesResponse is the response type for the Query/Reserves RPC method.
 */
export interface V1Beta1QueryReservesResponse {
  amount?: V1Beta1Coin[];
}

/**
 * QueryTotalBorrowedResponse is the response type for the Query/TotalBorrowed RPC method.
 */
export interface V1Beta1QueryTotalBorrowedResponse {
  borrowed_coins?: V1Beta1Coin[];
}

/**
 * QueryTotalDepositedResponse is the response type for the Query/TotalDeposited RPC method.
 */
export interface V1Beta1QueryTotalDepositedResponse {
  supplied_coins?: V1Beta1Coin[];
}

/**
 * QueryUnsyncedBorrowsResponse is the response type for the Query/UnsyncedBorrows RPC method.
 */
export interface V1Beta1QueryUnsyncedBorrowsResponse {
  borrows?: V1Beta1BorrowResponse[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

/**
 * QueryUnsyncedDepositsResponse is the response type for the Query/UnsyncedDeposits RPC method.
 */
export interface V1Beta1QueryUnsyncedDepositsResponse {
  deposits?: V1Beta1DepositResponse[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

/**
 * SupplyInterestFactorResponse defines an individual borrow interest factor.
 */
export interface V1Beta1SupplyInterestFactorResponse {
  denom?: string;

  /** sdk.Dec as string */
  value?: string;
}

import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, ResponseType } from "axios";

export type QueryParamsType = Record<string | number, any>;

export interface FullRequestParams extends Omit<AxiosRequestConfig, "data" | "params" | "url" | "responseType"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseType;
  /** request body */
  body?: unknown;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> extends Omit<AxiosRequestConfig, "data" | "cancelToken"> {
  securityWorker?: (
    securityData: SecurityDataType | null,
  ) => Promise<AxiosRequestConfig | void> | AxiosRequestConfig | void;
  secure?: boolean;
  format?: ResponseType;
}

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public instance: AxiosInstance;
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private secure?: boolean;
  private format?: ResponseType;

  constructor({ securityWorker, secure, format, ...axiosConfig }: ApiConfig<SecurityDataType> = {}) {
    this.instance = axios.create({ ...axiosConfig, baseURL: axiosConfig.baseURL || "" });
    this.secure = secure;
    this.format = format;
    this.securityWorker = securityWorker;
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  private mergeRequestParams(params1: AxiosRequestConfig, params2?: AxiosRequestConfig): AxiosRequestConfig {
    return {
      ...this.instance.defaults,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.instance.defaults.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createFormData(input: Record<string, unknown>): FormData {
    return Object.keys(input || {}).reduce((formData, key) => {
      const property = input[key];
      formData.append(
        key,
        property instanceof Blob
          ? property
          : typeof property === "object" && property !== null
          ? JSON.stringify(property)
          : `${property}`,
      );
      return formData;
    }, new FormData());
  }

  public request = async <T = any, _E = any>({
    secure,
    path,
    type,
    query,
    format,
    body,
    ...params
  }: FullRequestParams): Promise<AxiosResponse<T>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const responseFormat = (format && this.format) || void 0;

    if (type === ContentType.FormData && body && body !== null && typeof body === "object") {
      requestParams.headers.common = { Accept: "*/*" };
      requestParams.headers.post = {};
      requestParams.headers.put = {};

      body = this.createFormData(body as Record<string, unknown>);
    }

    return this.instance.request({
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      params: query,
      responseType: responseFormat,
      data: body,
      url: path,
    });
  };
}

/**
 * @title joltify/third_party/jolt/v1beta1/genesis.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryAccounts
   * @summary Accounts queries module accounts.
   * @request GET:/joltify/jolt/v1beta1/accounts
   */
  queryAccounts = (params: RequestParams = {}) =>
    this.request<V1Beta1QueryAccountsResponse, RpcStatus>({
      path: `/joltify/jolt/v1beta1/accounts`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryBorrows
   * @summary Borrows queries jolt borrows.
   * @request GET:/joltify/jolt/v1beta1/borrows
   */
  queryBorrows = (
    query?: {
      denom?: string;
      owner?: string;
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<V1Beta1QueryBorrowsResponse, RpcStatus>({
      path: `/joltify/jolt/v1beta1/borrows`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryDeposits
   * @summary Deposits queries jolt deposits.
   * @request GET:/joltify/jolt/v1beta1/deposits
   */
  queryDeposits = (
    query?: {
      denom?: string;
      owner?: string;
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<V1Beta1QueryDepositsResponse, RpcStatus>({
      path: `/joltify/jolt/v1beta1/deposits`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryInterestFactors
   * @summary InterestFactors queries jolt module interest factors.
   * @request GET:/joltify/jolt/v1beta1/interest-factors/{denom}
   */
  queryInterestFactors = (denom: string, params: RequestParams = {}) =>
    this.request<V1Beta1QueryInterestFactorsResponse, RpcStatus>({
      path: `/joltify/jolt/v1beta1/interest-factors/${denom}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryInterestRate
   * @summary InterestRate queries the jolt module interest rates.
   * @request GET:/joltify/jolt/v1beta1/interest-rate/{denom}
   */
  queryInterestRate = (denom: string, params: RequestParams = {}) =>
    this.request<V1Beta1QueryInterestRateResponse, RpcStatus>({
      path: `/joltify/jolt/v1beta1/interest-rate/${denom}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLiquidate
   * @summary queries jolt module interest factors.
   * @request GET:/joltify/jolt/v1beta1/liquidate
   */
  queryLiquidate = (
    query?: {
      borrower?: string;
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<V1Beta1QueryLiquidateResp, RpcStatus>({
      path: `/joltify/jolt/v1beta1/liquidate`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Params queries module params.
   * @request GET:/joltify/jolt/v1beta1/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<V1Beta1QueryParamsResponse, RpcStatus>({
      path: `/joltify/jolt/v1beta1/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryReserves
   * @summary Reserves queries total jolt reserve coins.
   * @request GET:/joltify/jolt/v1beta1/reserves/{denom}
   */
  queryReserves = (denom: string, params: RequestParams = {}) =>
    this.request<V1Beta1QueryReservesResponse, RpcStatus>({
      path: `/joltify/jolt/v1beta1/reserves/${denom}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTotalBorrowed
   * @summary TotalBorrowed queries total coins borrowed from jolt liquidity pools.
   * @request GET:/joltify/jolt/v1beta1/total-borrowed/{denom}
   */
  queryTotalBorrowed = (denom: string, params: RequestParams = {}) =>
    this.request<V1Beta1QueryTotalBorrowedResponse, RpcStatus>({
      path: `/joltify/jolt/v1beta1/total-borrowed/${denom}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTotalDeposited
   * @summary TotalDeposited queries total coins deposited to jolt liquidity pools.
   * @request GET:/joltify/jolt/v1beta1/total-deposited/{denom}
   */
  queryTotalDeposited = (denom: string, params: RequestParams = {}) =>
    this.request<V1Beta1QueryTotalDepositedResponse, RpcStatus>({
      path: `/joltify/jolt/v1beta1/total-deposited/${denom}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUnsyncedBorrows
   * @summary UnsyncedBorrows queries unsynced borrows.
   * @request GET:/joltify/jolt/v1beta1/unsynced-borrows
   */
  queryUnsyncedBorrows = (
    query?: {
      denom?: string;
      owner?: string;
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<V1Beta1QueryUnsyncedBorrowsResponse, RpcStatus>({
      path: `/joltify/jolt/v1beta1/unsynced-borrows`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUnsyncedDeposits
   * @summary UnsyncedDeposits queries unsynced deposits.
   * @request GET:/joltify/jolt/v1beta1/unsynced-deposits
   */
  queryUnsyncedDeposits = (
    query?: {
      denom?: string;
      owner?: string;
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<V1Beta1QueryUnsyncedDepositsResponse, RpcStatus>({
      path: `/joltify/jolt/v1beta1/unsynced-deposits`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });
}
