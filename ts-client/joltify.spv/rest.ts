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

export enum DepositorInfoDEPOSITTYPE {
  WithdrawProposal = "withdraw_proposal",
  TransferRequest = "transfer_request",
  DepositClose = "deposit_close",
  Unset = "unset",
  Processed = "processed",
  Deactive = "deactive",
}

export enum PoolInfoPOOLSTATUS {
  ACTIVE = "ACTIVE",
  INACTIVE = "INACTIVE",
  CLOSED = "CLOSED",
  PREPARE = "PREPARE",
  CLOSING = "CLOSING",
}

export enum PoolInfoPOOLTYPE {
  JUNIOR = "JUNIOR",
  SENIOR = "SENIOR",
}

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

export interface SpvDepositorInfo {
  investor_id?: string;

  /** @format byte */
  depositor_address?: string;
  pool_index?: string;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  locked_amount?: V1Beta1Coin;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  withdrawal_amount?: V1Beta1Coin;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  incentive_amount?: V1Beta1Coin;
  linkedNFT?: string[];
  deposit_type?: DepositorInfoDEPOSITTYPE;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  pending_interest?: V1Beta1Coin;
}

export type SpvMsgActivePoolResponse = object;

export interface SpvMsgAddInvestorsResponse {
  operation_result?: boolean;
}

export interface SpvMsgBorrowResponse {
  borrow_amount?: string;
}

export interface SpvMsgClaimInterestResponse {
  amount?: string;
}

export interface SpvMsgCreatePoolResponse {
  pool_index?: string[];
}

export type SpvMsgDepositResponse = object;

export type SpvMsgPayPrincipalResponse = object;

export type SpvMsgRepayInterestResponse = object;

export interface SpvMsgSubmitWithdrawProposalResponse {
  operation_result?: boolean;
}

export interface SpvMsgTransferOwnershipResponse {
  operation_result?: boolean;
}

export type SpvMsgUpdatePoolResponse = object;

export interface SpvMsgWithdrawPrincipalResponse {
  amount?: string;
}

/**
 * Params defines the parameters for the module.
 */
export type SpvParams = object;

export interface SpvPoolInfo {
  index?: string;
  pool_name?: string;

  /** @format int32 */
  linked_project?: number;

  /** @format byte */
  owner_address?: string;
  apy?: string;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  total_amount?: V1Beta1Coin;

  /** @format int32 */
  pay_freq?: number;
  reserve_factor?: string;

  /**
   * string            pool_nFT_class      = 9 [
   *    (cosmos_proto.scalar)  = "cosmos.Class",
   *    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/x/nft.Class",
   *    (gogoproto.nullable)   = false
   *  ];
   */
  pool_nFT_ids?: string[];

  /** @format date-time */
  last_payment_time?: string;
  pool_status?: PoolInfoPOOLSTATUS;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  borrowed_amount?: V1Beta1Coin;
  pool_interest?: string;

  /** @format uint64 */
  project_length?: string;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  usable_amount?: V1Beta1Coin;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  target_amount?: V1Beta1Coin;
  pool_type?: PoolInfoPOOLTYPE;
  escrow_interest_amount?: string;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  escrow_principal_amount?: V1Beta1Coin;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  withdraw_proposal_amount?: V1Beta1Coin;

  /** @format date-time */
  project_due_time?: string;
  withdraw_accounts?: string[];
  transfer_accounts?: string[];
}

export interface SpvQueryAllowedPoolsResponse {
  pools_index?: string[];
}

export interface SpvQueryClaimableInterestResponse {
  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  claimable_interest_amount?: V1Beta1Coin;
}

export interface SpvQueryDepositorResponse {
  depositor?: SpvDepositorInfo;
}

export interface SpvQueryListPoolsResponse {
  pools_info?: SpvPoolInfo[];

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

export interface SpvQueryOutstandingInterestResponse {
  amount?: string;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface SpvQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: SpvParams;
}

export interface SpvQueryQueryPoolResponse {
  pool_info?: SpvPoolInfo;
}

export interface SpvQuerywithdrawalPrincipalResponse {
  amount?: string;
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
 * @title joltify/spv/deposit.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QuerywithdrawalPrincipal
   * @summary Queries a list of withdrawalPrincipal items.
   * @request GET:/joltify-finance/joltify_lending/spv/withdrawal_principal/{pool_index}
   */
  querywithdrawalPrincipal = (poolIndex: string, query?: { walletAddress?: string }, params: RequestParams = {}) =>
    this.request<SpvQuerywithdrawalPrincipalResponse, RpcStatus>({
      path: `/joltify-finance/joltify_lending/spv/withdrawal_principal/${poolIndex}`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryClaimableInterest
   * @summary Queries a list of ClaimableInterest items.
   * @request GET:/joltify/spv/claimable_interest/{wallet}
   */
  queryClaimableInterest = (wallet: string, query?: { pool_index?: string }, params: RequestParams = {}) =>
    this.request<SpvQueryClaimableInterestResponse, RpcStatus>({
      path: `/joltify/spv/claimable_interest/${wallet}`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryDepositor
   * @request GET:/joltify/spv/depositor/{walletAddress}
   */
  queryDepositor = (walletAddress: string, query?: { deposit_pool_index?: string }, params: RequestParams = {}) =>
    this.request<SpvQueryDepositorResponse, RpcStatus>({
      path: `/joltify/spv/depositor/${walletAddress}`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryListPools
   * @summary Queries a list of Listpools items.
   * @request GET:/joltify/spv/list_pools
   */
  queryListPools = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<SpvQueryListPoolsResponse, RpcStatus>({
      path: `/joltify/spv/list_pools`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryOutstandingInterest
   * @request GET:/joltify/spv/outstanding_interest/{wallet}/{pool_index}
   */
  queryOutstandingInterest = (wallet: string, poolIndex: string, params: RequestParams = {}) =>
    this.request<SpvQueryOutstandingInterestResponse, RpcStatus>({
      path: `/joltify/spv/outstanding_interest/${wallet}/${poolIndex}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQueryPool
   * @summary Queries a list of QueryPool items.
   * @request GET:/joltify/spv/query_pool/{pool_index}
   */
  queryQueryPool = (poolIndex: string, params: RequestParams = {}) =>
    this.request<SpvQueryQueryPoolResponse, RpcStatus>({
      path: `/joltify/spv/query_pool/${poolIndex}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryAllowedPools
   * @summary Queries a list of AllowedPools items.
   * @request GET:/joltify/spve/allowed_pools/{wallet_address}
   */
  queryAllowedPools = (walletAddress: string, params: RequestParams = {}) =>
    this.request<SpvQueryAllowedPoolsResponse, RpcStatus>({
      path: `/joltify/spve/allowed_pools/${walletAddress}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
