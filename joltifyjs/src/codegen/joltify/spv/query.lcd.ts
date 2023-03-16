import { setPaginationParams } from "../../helpers";
import { LCDClient } from "@osmonauts/lcd";
import { QueryListPoolsRequest, QueryListPoolsResponseSDKType, QueryQueryPoolRequest, QueryQueryPoolResponseSDKType, QueryDepositorRequest, QueryDepositorResponseSDKType, QueryAllowedPoolsRequest, QueryAllowedPoolsResponseSDKType, QueryOutstandingInterestRequest, QueryOutstandingInterestResponseSDKType, QueryClaimableInterestRequest, QueryClaimableInterestResponseSDKType, QuerywithdrawalPrincipalRequest, QuerywithdrawalPrincipalResponseSDKType } from "./query";
export class LCDQueryClient {
  req: LCDClient;

  constructor({
    requestClient
  }: {
    requestClient: LCDClient;
  }) {
    this.req = requestClient;
    this.listPools = this.listPools.bind(this);
    this.queryPool = this.queryPool.bind(this);
    this.depositor = this.depositor.bind(this);
    this.allowedPools = this.allowedPools.bind(this);
    this.outstandingInterest = this.outstandingInterest.bind(this);
    this.claimableInterest = this.claimableInterest.bind(this);
    this.withdrawalPrincipal = this.withdrawalPrincipal.bind(this);
  }
  /* Queries a list of Listpools items. */


  async listPools(params: QueryListPoolsRequest = {
    pagination: undefined
  }): Promise<QueryListPoolsResponseSDKType> {
    const options: any = {
      params: {}
    };

    if (typeof params?.pagination !== "undefined") {
      setPaginationParams(options, params.pagination);
    }

    const endpoint = `joltify/spv/list_pools`;
    return await this.req.get<QueryListPoolsResponseSDKType>(endpoint, options);
  }
  /* Queries a list of QueryPool items. */


  async queryPool(params: QueryQueryPoolRequest): Promise<QueryQueryPoolResponseSDKType> {
    const endpoint = `joltify/spv/query_pool/${params.poolIndex}`;
    return await this.req.get<QueryQueryPoolResponseSDKType>(endpoint);
  }
  /* Depositor */


  async depositor(params: QueryDepositorRequest): Promise<QueryDepositorResponseSDKType> {
    const options: any = {
      params: {}
    };

    if (typeof params?.depositPoolIndex !== "undefined") {
      options.params.deposit_pool_index = params.depositPoolIndex;
    }

    const endpoint = `joltify/spv/depositor/${params.walletAddress}`;
    return await this.req.get<QueryDepositorResponseSDKType>(endpoint, options);
  }
  /* Queries a list of AllowedPools items. */


  async allowedPools(params: QueryAllowedPoolsRequest): Promise<QueryAllowedPoolsResponseSDKType> {
    const endpoint = `joltify/spve/allowed_pools/${params.walletAddress}`;
    return await this.req.get<QueryAllowedPoolsResponseSDKType>(endpoint);
  }
  /* OutstandingInterest */


  async outstandingInterest(params: QueryOutstandingInterestRequest): Promise<QueryOutstandingInterestResponseSDKType> {
    const endpoint = `joltify/spv/outstanding_interest/${params.wallet}/${params.poolIndex}`;
    return await this.req.get<QueryOutstandingInterestResponseSDKType>(endpoint);
  }
  /* Queries a list of ClaimableInterest items. */


  async claimableInterest(params: QueryClaimableInterestRequest): Promise<QueryClaimableInterestResponseSDKType> {
    const options: any = {
      params: {}
    };

    if (typeof params?.poolIndex !== "undefined") {
      options.params.pool_index = params.poolIndex;
    }

    const endpoint = `joltify/spv/claimable_interest/${params.wallet}`;
    return await this.req.get<QueryClaimableInterestResponseSDKType>(endpoint, options);
  }
  /* Queries a list of withdrawalPrincipal items. */


  async withdrawalPrincipal(params: QuerywithdrawalPrincipalRequest): Promise<QuerywithdrawalPrincipalResponseSDKType> {
    const options: any = {
      params: {}
    };

    if (typeof params?.walletAddress !== "undefined") {
      options.params.walletAddress = params.walletAddress;
    }

    const endpoint = `joltify-finance/joltify_lending/spv/withdrawal_principal/${params.poolIndex}`;
    return await this.req.get<QuerywithdrawalPrincipalResponseSDKType>(endpoint, options);
  }

}