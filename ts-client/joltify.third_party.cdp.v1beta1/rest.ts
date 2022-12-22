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
 * Params defines the parameters for the cdp module.
 */
export interface ThirdPartycdpv1Beta1Params {
  collateral_params?: V1Beta1CollateralParam[];
  debt_param?: V1Beta1DebtParam;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  global_debt_limit?: V1Beta1Coin;
  surplus_auction_threshold?: string;
  surplus_auction_lot?: string;
  debt_auction_threshold?: string;
  debt_auction_lot?: string;
  circuit_breaker?: boolean;
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
 * CDPResponse defines the state of a single collateralized debt position.
 */
export interface V1Beta1CDPResponse {
  /** @format uint64 */
  id?: string;
  owner?: string;
  type?: string;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  collateral?: V1Beta1Coin;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  principal?: V1Beta1Coin;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  accumulated_fees?: V1Beta1Coin;

  /** @format date-time */
  fees_updated?: string;
  interest_factor?: string;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  collateral_value?: V1Beta1Coin;
  collateralization_ratio?: string;
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

export interface V1Beta1CollateralParam {
  denom?: string;
  type?: string;
  liquidation_ratio?: string;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  debt_limit?: V1Beta1Coin;
  stability_fee?: string;
  auction_size?: string;
  liquidation_penalty?: string;
  spot_market_id?: string;
  liquidation_market_id?: string;
  keeper_reward_percentage?: string;
  check_collateralization_index_count?: string;
  conversion_factor?: string;
}

export interface V1Beta1DebtParam {
  denom?: string;
  reference_asset?: string;
  conversion_factor?: string;
  debt_floor?: string;
}

export interface V1Beta1Deposit {
  /** @format uint64 */
  cdp_id?: string;
  depositor?: string;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  amount?: V1Beta1Coin;
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
 * MsgCreateCDPResponse defines the Msg/CreateCDP response type.
 */
export interface V1Beta1MsgCreateCDPResponse {
  /** @format uint64 */
  cdp_id?: string;
}

/**
 * MsgDepositResponse defines the Msg/Deposit response type.
 */
export type V1Beta1MsgDepositResponse = object;

/**
 * MsgDrawDebtResponse defines the Msg/DrawDebt response type.
 */
export type V1Beta1MsgDrawDebtResponse = object;

/**
 * MsgLiquidateResponse defines the Msg/Liquidate response type.
 */
export type V1Beta1MsgLiquidateResponse = object;

/**
 * MsgRepayDebtResponse defines the Msg/RepayDebt response type.
 */
export type V1Beta1MsgRepayDebtResponse = object;

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
 * QueryAccountsResponse defines the response type for the Query/Accounts RPC method.
 */
export interface V1Beta1QueryAccountsResponse {
  accounts?: V1Beta1ModuleAccount[];
}

/**
 * QueryCdpResponse defines the response type for the Query/Cdp RPC method.
 */
export interface V1Beta1QueryCdpResponse {
  /** CDPResponse defines the state of a single collateralized debt position. */
  cdp?: V1Beta1CDPResponse;
}

/**
 * QueryCdpsResponse defines the response type for the Query/Cdps RPC method.
 */
export interface V1Beta1QueryCdpsResponse {
  cdps?: V1Beta1CDPResponse[];

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
 * QueryDepositsResponse defines the response type for the Query/Deposits RPC method.
 */
export interface V1Beta1QueryDepositsResponse {
  deposits?: V1Beta1Deposit[];
}

/**
 * QueryParamsResponse defines the response type for the Query/Params RPC method.
 */
export interface V1Beta1QueryParamsResponse {
  /** Params defines the parameters for the cdp module. */
  params?: ThirdPartycdpv1Beta1Params;
}

/**
 * QueryTotalCollateralResponse defines the response type for the Query/TotalCollateral RPC method.
 */
export interface V1Beta1QueryTotalCollateralResponse {
  total_collateral?: V1Beta1TotalCollateral[];
}

/**
 * QueryTotalPrincipalResponse defines the response type for the Query/TotalPrincipal RPC method.
 */
export interface V1Beta1QueryTotalPrincipalResponse {
  total_principal?: V1Beta1TotalPrincipal[];
}

export interface V1Beta1TotalCollateral {
  collateral_type?: string;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  amount?: V1Beta1Coin;
}

export interface V1Beta1TotalPrincipal {
  collateral_type?: string;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  amount?: V1Beta1Coin;
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
 * @title joltify/third_party/cdp/v1beta1/cdp.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryAccounts
   * @summary Accounts queries the CDP module accounts.
   * @request GET:/joltify/cdp/v1beta1/accounts
   */
  queryAccounts = (params: RequestParams = {}) =>
    this.request<V1Beta1QueryAccountsResponse, RpcStatus>({
      path: `/joltify/cdp/v1beta1/accounts`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryCdps
   * @summary Cdps queries all active CDPs.
   * @request GET:/joltify/cdp/v1beta1/cdps
   */
  queryCdps = (
    query?: {
      collateral_type?: string;
      owner?: string;
      id?: string;
      ratio?: string;
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<V1Beta1QueryCdpsResponse, RpcStatus>({
      path: `/joltify/cdp/v1beta1/cdps`,
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
   * @summary Deposits queries deposits associated with the CDP owned by an address for a collateral type.
   * @request GET:/joltify/cdp/v1beta1/cdps/deposits/{owner}/{collateral_type}
   */
  queryDeposits = (owner: string, collateralType: string, params: RequestParams = {}) =>
    this.request<V1Beta1QueryDepositsResponse, RpcStatus>({
      path: `/joltify/cdp/v1beta1/cdps/deposits/${owner}/${collateralType}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryCdp
   * @summary Cdp queries a CDP with the input owner address and collateral type.
   * @request GET:/joltify/cdp/v1beta1/cdps/{owner}/{collateral_type}
   */
  queryCdp = (owner: string, collateralType: string, params: RequestParams = {}) =>
    this.request<V1Beta1QueryCdpResponse, RpcStatus>({
      path: `/joltify/cdp/v1beta1/cdps/${owner}/${collateralType}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Params queries all parameters of the cdp module.
   * @request GET:/joltify/cdp/v1beta1/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<V1Beta1QueryParamsResponse, RpcStatus>({
      path: `/joltify/cdp/v1beta1/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTotalCollateral
   * @summary TotalCollateral queries the total collateral of a given collateral type.
   * @request GET:/joltify/cdp/v1beta1/totalCollateral
   */
  queryTotalCollateral = (query?: { collateral_type?: string }, params: RequestParams = {}) =>
    this.request<V1Beta1QueryTotalCollateralResponse, RpcStatus>({
      path: `/joltify/cdp/v1beta1/totalCollateral`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTotalPrincipal
   * @summary TotalPrincipal queries the total principal of a given collateral type.
   * @request GET:/joltify/cdp/v1beta1/totalPrincipal
   */
  queryTotalPrincipal = (query?: { collateral_type?: string }, params: RequestParams = {}) =>
    this.request<V1Beta1QueryTotalPrincipalResponse, RpcStatus>({
      path: `/joltify/cdp/v1beta1/totalPrincipal`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });
}
