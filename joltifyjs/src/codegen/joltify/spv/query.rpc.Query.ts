import { Rpc } from "../../helpers";
import * as _m0 from "protobufjs/minimal";
import { QueryClient, createProtobufRpcClient } from "@cosmjs/stargate";
import { QueryParamsRequest, QueryParamsResponse, QueryListPoolsRequest, QueryListPoolsResponse, QueryQueryPoolRequest, QueryQueryPoolResponse, QueryDepositorRequest, QueryDepositorResponse, QueryAllowedPoolsRequest, QueryAllowedPoolsResponse, QueryOutstandingInterestRequest, QueryOutstandingInterestResponse, QueryClaimableInterestRequest, QueryClaimableInterestResponse, QuerywithdrawalPrincipalRequest, QuerywithdrawalPrincipalResponse } from "./query";
/** Query defines the gRPC querier service. */

export interface Query {
  /** Parameters queries the parameters of the module. */
  params(request?: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of Listpools items. */

  listPools(request?: QueryListPoolsRequest): Promise<QueryListPoolsResponse>;
  /** Queries a list of QueryPool items. */

  queryPool(request: QueryQueryPoolRequest): Promise<QueryQueryPoolResponse>;
  depositor(request: QueryDepositorRequest): Promise<QueryDepositorResponse>;
  /** Queries a list of AllowedPools items. */

  allowedPools(request: QueryAllowedPoolsRequest): Promise<QueryAllowedPoolsResponse>;
  outstandingInterest(request: QueryOutstandingInterestRequest): Promise<QueryOutstandingInterestResponse>;
  /** Queries a list of ClaimableInterest items. */

  claimableInterest(request: QueryClaimableInterestRequest): Promise<QueryClaimableInterestResponse>;
  /** Queries a list of withdrawalPrincipal items. */

  withdrawalPrincipal(request: QuerywithdrawalPrincipalRequest): Promise<QuerywithdrawalPrincipalResponse>;
}
export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.params = this.params.bind(this);
    this.listPools = this.listPools.bind(this);
    this.queryPool = this.queryPool.bind(this);
    this.depositor = this.depositor.bind(this);
    this.allowedPools = this.allowedPools.bind(this);
    this.outstandingInterest = this.outstandingInterest.bind(this);
    this.claimableInterest = this.claimableInterest.bind(this);
    this.withdrawalPrincipal = this.withdrawalPrincipal.bind(this);
  }

  params(request: QueryParamsRequest = {}): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Query", "Params", data);
    return promise.then(data => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  listPools(request: QueryListPoolsRequest = {
    pagination: undefined
  }): Promise<QueryListPoolsResponse> {
    const data = QueryListPoolsRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Query", "ListPools", data);
    return promise.then(data => QueryListPoolsResponse.decode(new _m0.Reader(data)));
  }

  queryPool(request: QueryQueryPoolRequest): Promise<QueryQueryPoolResponse> {
    const data = QueryQueryPoolRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Query", "QueryPool", data);
    return promise.then(data => QueryQueryPoolResponse.decode(new _m0.Reader(data)));
  }

  depositor(request: QueryDepositorRequest): Promise<QueryDepositorResponse> {
    const data = QueryDepositorRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Query", "Depositor", data);
    return promise.then(data => QueryDepositorResponse.decode(new _m0.Reader(data)));
  }

  allowedPools(request: QueryAllowedPoolsRequest): Promise<QueryAllowedPoolsResponse> {
    const data = QueryAllowedPoolsRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Query", "AllowedPools", data);
    return promise.then(data => QueryAllowedPoolsResponse.decode(new _m0.Reader(data)));
  }

  outstandingInterest(request: QueryOutstandingInterestRequest): Promise<QueryOutstandingInterestResponse> {
    const data = QueryOutstandingInterestRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Query", "OutstandingInterest", data);
    return promise.then(data => QueryOutstandingInterestResponse.decode(new _m0.Reader(data)));
  }

  claimableInterest(request: QueryClaimableInterestRequest): Promise<QueryClaimableInterestResponse> {
    const data = QueryClaimableInterestRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Query", "ClaimableInterest", data);
    return promise.then(data => QueryClaimableInterestResponse.decode(new _m0.Reader(data)));
  }

  withdrawalPrincipal(request: QuerywithdrawalPrincipalRequest): Promise<QuerywithdrawalPrincipalResponse> {
    const data = QuerywithdrawalPrincipalRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Query", "withdrawalPrincipal", data);
    return promise.then(data => QuerywithdrawalPrincipalResponse.decode(new _m0.Reader(data)));
  }

}
export const createRpcQueryExtension = (base: QueryClient) => {
  const rpc = createProtobufRpcClient(base);
  const queryService = new QueryClientImpl(rpc);
  return {
    params(request?: QueryParamsRequest): Promise<QueryParamsResponse> {
      return queryService.params(request);
    },

    listPools(request?: QueryListPoolsRequest): Promise<QueryListPoolsResponse> {
      return queryService.listPools(request);
    },

    queryPool(request: QueryQueryPoolRequest): Promise<QueryQueryPoolResponse> {
      return queryService.queryPool(request);
    },

    depositor(request: QueryDepositorRequest): Promise<QueryDepositorResponse> {
      return queryService.depositor(request);
    },

    allowedPools(request: QueryAllowedPoolsRequest): Promise<QueryAllowedPoolsResponse> {
      return queryService.allowedPools(request);
    },

    outstandingInterest(request: QueryOutstandingInterestRequest): Promise<QueryOutstandingInterestResponse> {
      return queryService.outstandingInterest(request);
    },

    claimableInterest(request: QueryClaimableInterestRequest): Promise<QueryClaimableInterestResponse> {
      return queryService.claimableInterest(request);
    },

    withdrawalPrincipal(request: QuerywithdrawalPrincipalRequest): Promise<QuerywithdrawalPrincipalResponse> {
      return queryService.withdrawalPrincipal(request);
    }

  };
};