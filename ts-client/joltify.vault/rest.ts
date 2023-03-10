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

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
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
   * query the next page most efficiently. It will be empty if
   * there are no more results.
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

export interface VaultIssueToken {
  /** @format byte */
  creator?: string;
  index?: string;

  /** @format byte */
  coin?: string;

  /** @format byte */
  receiver?: string;
}

export interface VaultMsgCreateCreatePoolResponse {
  successful?: boolean;
}

export interface VaultMsgCreateIssueTokenResponse {
  successful?: boolean;
}

export interface VaultMsgCreateOutboundTxResponse {
  successful?: boolean;
}

export interface VaultOutboundTx {
  index?: string;
  processed?: boolean;
  items?: Record<string, Vaultproposals>;
  chainType?: string;
  inTxHash?: string;

  /** @format byte */
  receiverAddress?: string;
  needMint?: boolean;
  feecoin?: V1Beta1Coin[];
}

export interface VaultPoolProposal {
  poolPubKey?: string;

  /** @format byte */
  poolAddr?: string;
  nodes?: string[];
}

export interface VaultQueryAllCreatePoolResponse {
  CreatePool?: VaultPoolProposal[];

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

export interface VaultQueryAllIssueTokenResponse {
  IssueToken?: VaultIssueToken[];

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

export interface VaultQueryAllOutboundTxResponse {
  outboundTx?: VaultOutboundTx[];

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

export interface VaultQueryAllValidatorsResponse {
  all_validators?: VaultValidators[];

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

export interface VaultQueryGetCreatePoolResponse {
  CreatePool?: VaultPoolProposal;
}

export interface VaultQueryGetIssueTokenResponse {
  IssueToken?: VaultIssueToken;
}

export interface VaultQueryGetOutboundTxResponse {
  outboundTx?: VaultOutboundTx;
}

export interface VaultQueryGetQuotaResponse {
  coinQuotaResponse?: VaultcoinsQuota;
}

export interface VaultQueryGetValidatorsResponse {
  validators?: VaultValidators;
}

export interface VaultQueryLastPoolResponse {
  pools?: VaultpoolInfo[];
}

export interface VaultQueryPendingFeeResponse {
  feecoin?: V1Beta1Coin[];

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

export interface VaultValidator {
  /** @format byte */
  pubkey?: string;

  /** @format int64 */
  power?: string;
}

export interface VaultValidators {
  all_validators?: VaultValidator[];

  /** @format int64 */
  height?: string;
}

export interface VaultcoinsQuota {
  history?: VaulthistoricalAmount[];
  CoinsSum?: V1Beta1Coin[];
}

export interface Vaultentity {
  /** @format byte */
  address?: string;
  feecoin?: V1Beta1Coin[];
}

export interface VaulthistoricalAmount {
  /** @format int64 */
  blockHeight?: string;
  amount?: V1Beta1Coin[];
}

export interface VaultpoolInfo {
  BlockHeight?: string;
  CreatePool?: VaultPoolProposal;
}

export interface Vaultproposals {
  entry?: Vaultentity[];
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
 * @title joltify/vault/create_pool.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryCreatePoolAll
   * @summary Queries a list of createPool items.
   * @request GET:/joltify/vault/createPool
   */
  queryCreatePoolAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<VaultQueryAllCreatePoolResponse, RpcStatus>({
      path: `/joltify/vault/createPool`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryCreatePool
   * @summary Queries a createPool by index.
   * @request GET:/joltify/vault/createPool/{index}
   */
  queryCreatePool = (index: string, params: RequestParams = {}) =>
    this.request<VaultQueryGetCreatePoolResponse, RpcStatus>({
      path: `/joltify/vault/createPool/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetLastPool
   * @summary Queries a createPool by index.
   * @request GET:/joltify/vault/getLastPool
   */
  queryGetLastPool = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<VaultQueryLastPoolResponse, RpcStatus>({
      path: `/joltify/vault/getLastPool`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetPendingFee
   * @summary Queries the pending fee
   * @request GET:/joltify/vault/get_pending_fee
   */
  queryGetPendingFee = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<VaultQueryPendingFeeResponse, RpcStatus>({
      path: `/joltify/vault/get_pending_fee`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetQuota
   * @summary Queries a list of GetQuota items.
   * @request GET:/joltify/vault/get_quota/{query_length}
   */
  queryGetQuota = (
    queryLength: number,
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<VaultQueryGetQuotaResponse, RpcStatus>({
      path: `/joltify/vault/get_quota/${queryLength}`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetValidators
   * @summary Queries a list of GetValidators items.
   * @request GET:/joltify/vault/get_validators/{height}
   */
  queryGetValidators = (height: string, params: RequestParams = {}) =>
    this.request<VaultQueryGetValidatorsResponse, RpcStatus>({
      path: `/joltify/vault/get_validators/${height}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryIssueTokenAll
   * @summary Queries a list of issueToken items.
   * @request GET:/joltify/vault/issueToken
   */
  queryIssueTokenAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<VaultQueryAllIssueTokenResponse, RpcStatus>({
      path: `/joltify/vault/issueToken`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryIssueToken
   * @summary Queries a issueToken by index.
   * @request GET:/joltify/vault/issueToken/{index}
   */
  queryIssueToken = (index: string, params: RequestParams = {}) =>
    this.request<VaultQueryGetIssueTokenResponse, RpcStatus>({
      path: `/joltify/vault/issueToken/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryOutboundTxAll
   * @summary Queries a list of OutboundTx items.
   * @request GET:/joltify/vault/outbound_tx
   */
  queryOutboundTxAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<VaultQueryAllOutboundTxResponse, RpcStatus>({
      path: `/joltify/vault/outbound_tx`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryOutboundTx
   * @summary Queries a OutboundTx by index.
   * @request GET:/joltify/vault/outbound_tx/{requestID}
   */
  queryOutboundTx = (requestId: string, params: RequestParams = {}) =>
    this.request<VaultQueryGetOutboundTxResponse, RpcStatus>({
      path: `/joltify/vault/outbound_tx/${requestId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetAllValidators
   * @summary Queries a list of GetValidators items.
   * @request GET:/joltify/vault/validators
   */
  queryGetAllValidators = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<VaultQueryAllValidatorsResponse, RpcStatus>({
      path: `/joltify/vault/validators`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });
}
