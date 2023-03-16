import { setPaginationParams } from "../../helpers";
import { LCDClient } from "@osmonauts/lcd";
import { QueryGetOutboundTxRequest, QueryGetOutboundTxResponseSDKType, QueryAllOutboundTxRequest, QueryAllOutboundTxResponseSDKType, QueryGetValidatorsRequest, QueryGetValidatorsResponseSDKType, QueryAllValidatorsRequest, QueryAllValidatorsResponseSDKType, QueryGetQuotaRequest, QueryGetQuotaResponseSDKType, QueryGetIssueTokenRequest, QueryGetIssueTokenResponseSDKType, QueryAllIssueTokenRequest, QueryAllIssueTokenResponseSDKType, QueryGetCreatePoolRequest, QueryGetCreatePoolResponseSDKType, QueryAllCreatePoolRequest, QueryAllCreatePoolResponseSDKType, QueryLatestPoolRequest, QueryLastPoolResponseSDKType, QueryPendingFeeRequest, QueryPendingFeeResponseSDKType } from "./query";
export class LCDQueryClient {
  req: LCDClient;

  constructor({
    requestClient
  }: {
    requestClient: LCDClient;
  }) {
    this.req = requestClient;
    this.outboundTx = this.outboundTx.bind(this);
    this.outboundTxAll = this.outboundTxAll.bind(this);
    this.getValidators = this.getValidators.bind(this);
    this.getAllValidators = this.getAllValidators.bind(this);
    this.getQuota = this.getQuota.bind(this);
    this.issueToken = this.issueToken.bind(this);
    this.issueTokenAll = this.issueTokenAll.bind(this);
    this.createPool = this.createPool.bind(this);
    this.createPoolAll = this.createPoolAll.bind(this);
    this.getLastPool = this.getLastPool.bind(this);
    this.getPendingFee = this.getPendingFee.bind(this);
  }
  /* Queries a OutboundTx by index. */


  async outboundTx(params: QueryGetOutboundTxRequest): Promise<QueryGetOutboundTxResponseSDKType> {
    const endpoint = `joltify/vault/outbound_tx/${params.requestID}`;
    return await this.req.get<QueryGetOutboundTxResponseSDKType>(endpoint);
  }
  /* Queries a list of OutboundTx items. */


  async outboundTxAll(params: QueryAllOutboundTxRequest = {
    pagination: undefined
  }): Promise<QueryAllOutboundTxResponseSDKType> {
    const options: any = {
      params: {}
    };

    if (typeof params?.pagination !== "undefined") {
      setPaginationParams(options, params.pagination);
    }

    const endpoint = `joltify/vault/outbound_tx`;
    return await this.req.get<QueryAllOutboundTxResponseSDKType>(endpoint, options);
  }
  /* Queries a list of GetValidators items. */


  async getValidators(params: QueryGetValidatorsRequest): Promise<QueryGetValidatorsResponseSDKType> {
    const endpoint = `joltify/vault/get_validators/${params.height}`;
    return await this.req.get<QueryGetValidatorsResponseSDKType>(endpoint);
  }
  /* Queries a list of GetValidators items. */


  async getAllValidators(params: QueryAllValidatorsRequest = {
    pagination: undefined
  }): Promise<QueryAllValidatorsResponseSDKType> {
    const options: any = {
      params: {}
    };

    if (typeof params?.pagination !== "undefined") {
      setPaginationParams(options, params.pagination);
    }

    const endpoint = `joltify/vault/validators`;
    return await this.req.get<QueryAllValidatorsResponseSDKType>(endpoint, options);
  }
  /* Queries a list of GetQuota items. */


  async getQuota(params: QueryGetQuotaRequest): Promise<QueryGetQuotaResponseSDKType> {
    const options: any = {
      params: {}
    };

    if (typeof params?.pagination !== "undefined") {
      setPaginationParams(options, params.pagination);
    }

    const endpoint = `joltify/vault/get_quota/${params.queryLength}`;
    return await this.req.get<QueryGetQuotaResponseSDKType>(endpoint, options);
  }
  /* Queries a issueToken by index. */


  async issueToken(params: QueryGetIssueTokenRequest): Promise<QueryGetIssueTokenResponseSDKType> {
    const endpoint = `joltify/vault/issueToken/${params.index}`;
    return await this.req.get<QueryGetIssueTokenResponseSDKType>(endpoint);
  }
  /* Queries a list of issueToken items. */


  async issueTokenAll(params: QueryAllIssueTokenRequest = {
    pagination: undefined
  }): Promise<QueryAllIssueTokenResponseSDKType> {
    const options: any = {
      params: {}
    };

    if (typeof params?.pagination !== "undefined") {
      setPaginationParams(options, params.pagination);
    }

    const endpoint = `joltify/vault/issueToken`;
    return await this.req.get<QueryAllIssueTokenResponseSDKType>(endpoint, options);
  }
  /* Queries a createPool by index. */


  async createPool(params: QueryGetCreatePoolRequest): Promise<QueryGetCreatePoolResponseSDKType> {
    const endpoint = `joltify/vault/createPool/${params.index}`;
    return await this.req.get<QueryGetCreatePoolResponseSDKType>(endpoint);
  }
  /* Queries a list of createPool items. */


  async createPoolAll(params: QueryAllCreatePoolRequest = {
    pagination: undefined
  }): Promise<QueryAllCreatePoolResponseSDKType> {
    const options: any = {
      params: {}
    };

    if (typeof params?.pagination !== "undefined") {
      setPaginationParams(options, params.pagination);
    }

    const endpoint = `joltify/vault/createPool`;
    return await this.req.get<QueryAllCreatePoolResponseSDKType>(endpoint, options);
  }
  /* Queries a createPool by index. */


  async getLastPool(params: QueryLatestPoolRequest = {
    pagination: undefined
  }): Promise<QueryLastPoolResponseSDKType> {
    const options: any = {
      params: {}
    };

    if (typeof params?.pagination !== "undefined") {
      setPaginationParams(options, params.pagination);
    }

    const endpoint = `joltify/vault/getLastPool`;
    return await this.req.get<QueryLastPoolResponseSDKType>(endpoint, options);
  }
  /* Queries the pending fee */


  async getPendingFee(params: QueryPendingFeeRequest = {
    pagination: undefined
  }): Promise<QueryPendingFeeResponseSDKType> {
    const options: any = {
      params: {}
    };

    if (typeof params?.pagination !== "undefined") {
      setPaginationParams(options, params.pagination);
    }

    const endpoint = `joltify/vault/get_pending_fee`;
    return await this.req.get<QueryPendingFeeResponseSDKType>(endpoint, options);
  }

}