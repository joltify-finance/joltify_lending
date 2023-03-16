import { Rpc } from "../../helpers";
import * as _m0 from "protobufjs/minimal";
import { QueryClient, createProtobufRpcClient } from "@cosmjs/stargate";
import { QueryGetOutboundTxRequest, QueryGetOutboundTxResponse, QueryAllOutboundTxRequest, QueryAllOutboundTxResponse, QueryGetValidatorsRequest, QueryGetValidatorsResponse, QueryAllValidatorsRequest, QueryAllValidatorsResponse, QueryGetQuotaRequest, QueryGetQuotaResponse, QueryGetIssueTokenRequest, QueryGetIssueTokenResponse, QueryAllIssueTokenRequest, QueryAllIssueTokenResponse, QueryGetCreatePoolRequest, QueryGetCreatePoolResponse, QueryAllCreatePoolRequest, QueryAllCreatePoolResponse, QueryLatestPoolRequest, QueryLastPoolResponse, QueryPendingFeeRequest, QueryPendingFeeResponse } from "./query";
/** Query defines the gRPC querier service. */

export interface Query {
  /** Queries a OutboundTx by index. */
  outboundTx(request: QueryGetOutboundTxRequest): Promise<QueryGetOutboundTxResponse>;
  /** Queries a list of OutboundTx items. */

  outboundTxAll(request?: QueryAllOutboundTxRequest): Promise<QueryAllOutboundTxResponse>;
  /** Queries a list of GetValidators items. */

  getValidators(request: QueryGetValidatorsRequest): Promise<QueryGetValidatorsResponse>;
  /** Queries a list of GetValidators items. */

  getAllValidators(request?: QueryAllValidatorsRequest): Promise<QueryAllValidatorsResponse>;
  /** Queries a list of GetQuota items. */

  getQuota(request: QueryGetQuotaRequest): Promise<QueryGetQuotaResponse>;
  /** Queries a issueToken by index. */

  issueToken(request: QueryGetIssueTokenRequest): Promise<QueryGetIssueTokenResponse>;
  /** Queries a list of issueToken items. */

  issueTokenAll(request?: QueryAllIssueTokenRequest): Promise<QueryAllIssueTokenResponse>;
  /** Queries a createPool by index. */

  createPool(request: QueryGetCreatePoolRequest): Promise<QueryGetCreatePoolResponse>;
  /** Queries a list of createPool items. */

  createPoolAll(request?: QueryAllCreatePoolRequest): Promise<QueryAllCreatePoolResponse>;
  /** Queries a createPool by index. */

  getLastPool(request?: QueryLatestPoolRequest): Promise<QueryLastPoolResponse>;
  /** Queries the pending fee */

  getPendingFee(request?: QueryPendingFeeRequest): Promise<QueryPendingFeeResponse>;
}
export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
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

  outboundTx(request: QueryGetOutboundTxRequest): Promise<QueryGetOutboundTxResponse> {
    const data = QueryGetOutboundTxRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Query", "OutboundTx", data);
    return promise.then(data => QueryGetOutboundTxResponse.decode(new _m0.Reader(data)));
  }

  outboundTxAll(request: QueryAllOutboundTxRequest = {
    pagination: undefined
  }): Promise<QueryAllOutboundTxResponse> {
    const data = QueryAllOutboundTxRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Query", "OutboundTxAll", data);
    return promise.then(data => QueryAllOutboundTxResponse.decode(new _m0.Reader(data)));
  }

  getValidators(request: QueryGetValidatorsRequest): Promise<QueryGetValidatorsResponse> {
    const data = QueryGetValidatorsRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Query", "GetValidators", data);
    return promise.then(data => QueryGetValidatorsResponse.decode(new _m0.Reader(data)));
  }

  getAllValidators(request: QueryAllValidatorsRequest = {
    pagination: undefined
  }): Promise<QueryAllValidatorsResponse> {
    const data = QueryAllValidatorsRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Query", "GetAllValidators", data);
    return promise.then(data => QueryAllValidatorsResponse.decode(new _m0.Reader(data)));
  }

  getQuota(request: QueryGetQuotaRequest): Promise<QueryGetQuotaResponse> {
    const data = QueryGetQuotaRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Query", "GetQuota", data);
    return promise.then(data => QueryGetQuotaResponse.decode(new _m0.Reader(data)));
  }

  issueToken(request: QueryGetIssueTokenRequest): Promise<QueryGetIssueTokenResponse> {
    const data = QueryGetIssueTokenRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Query", "IssueToken", data);
    return promise.then(data => QueryGetIssueTokenResponse.decode(new _m0.Reader(data)));
  }

  issueTokenAll(request: QueryAllIssueTokenRequest = {
    pagination: undefined
  }): Promise<QueryAllIssueTokenResponse> {
    const data = QueryAllIssueTokenRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Query", "IssueTokenAll", data);
    return promise.then(data => QueryAllIssueTokenResponse.decode(new _m0.Reader(data)));
  }

  createPool(request: QueryGetCreatePoolRequest): Promise<QueryGetCreatePoolResponse> {
    const data = QueryGetCreatePoolRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Query", "CreatePool", data);
    return promise.then(data => QueryGetCreatePoolResponse.decode(new _m0.Reader(data)));
  }

  createPoolAll(request: QueryAllCreatePoolRequest = {
    pagination: undefined
  }): Promise<QueryAllCreatePoolResponse> {
    const data = QueryAllCreatePoolRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Query", "CreatePoolAll", data);
    return promise.then(data => QueryAllCreatePoolResponse.decode(new _m0.Reader(data)));
  }

  getLastPool(request: QueryLatestPoolRequest = {
    pagination: undefined
  }): Promise<QueryLastPoolResponse> {
    const data = QueryLatestPoolRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Query", "GetLastPool", data);
    return promise.then(data => QueryLastPoolResponse.decode(new _m0.Reader(data)));
  }

  getPendingFee(request: QueryPendingFeeRequest = {
    pagination: undefined
  }): Promise<QueryPendingFeeResponse> {
    const data = QueryPendingFeeRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Query", "GetPendingFee", data);
    return promise.then(data => QueryPendingFeeResponse.decode(new _m0.Reader(data)));
  }

}
export const createRpcQueryExtension = (base: QueryClient) => {
  const rpc = createProtobufRpcClient(base);
  const queryService = new QueryClientImpl(rpc);
  return {
    outboundTx(request: QueryGetOutboundTxRequest): Promise<QueryGetOutboundTxResponse> {
      return queryService.outboundTx(request);
    },

    outboundTxAll(request?: QueryAllOutboundTxRequest): Promise<QueryAllOutboundTxResponse> {
      return queryService.outboundTxAll(request);
    },

    getValidators(request: QueryGetValidatorsRequest): Promise<QueryGetValidatorsResponse> {
      return queryService.getValidators(request);
    },

    getAllValidators(request?: QueryAllValidatorsRequest): Promise<QueryAllValidatorsResponse> {
      return queryService.getAllValidators(request);
    },

    getQuota(request: QueryGetQuotaRequest): Promise<QueryGetQuotaResponse> {
      return queryService.getQuota(request);
    },

    issueToken(request: QueryGetIssueTokenRequest): Promise<QueryGetIssueTokenResponse> {
      return queryService.issueToken(request);
    },

    issueTokenAll(request?: QueryAllIssueTokenRequest): Promise<QueryAllIssueTokenResponse> {
      return queryService.issueTokenAll(request);
    },

    createPool(request: QueryGetCreatePoolRequest): Promise<QueryGetCreatePoolResponse> {
      return queryService.createPool(request);
    },

    createPoolAll(request?: QueryAllCreatePoolRequest): Promise<QueryAllCreatePoolResponse> {
      return queryService.createPoolAll(request);
    },

    getLastPool(request?: QueryLatestPoolRequest): Promise<QueryLastPoolResponse> {
      return queryService.getLastPool(request);
    },

    getPendingFee(request?: QueryPendingFeeRequest): Promise<QueryPendingFeeResponse> {
      return queryService.getPendingFee(request);
    }

  };
};