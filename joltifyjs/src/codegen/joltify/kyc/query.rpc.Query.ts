import { Rpc } from "../../helpers";
import * as _m0 from "protobufjs/minimal";
import { QueryClient, createProtobufRpcClient } from "@cosmjs/stargate";
import { QueryParamsRequest, QueryParamsResponse, QueryInvestorWalletsRequest, QueryInvestorWalletsResponse, QueryByWalletRequest, QueryByWalletResponse, ListInvestorsRequest, ListInvestorsResponse } from "./query";
/** Query defines the gRPC querier service. */

export interface Query {
  /** Parameters queries the parameters of the module. */
  params(request?: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of QueryInvestorWallets items. */

  queryInvestorWallets(request: QueryInvestorWalletsRequest): Promise<QueryInvestorWalletsResponse>;
  /** Queries a list of QueryByWallet items. */

  queryByWallet(request: QueryByWalletRequest): Promise<QueryByWalletResponse>;
  /** Queries a list of ListInvestors items. */

  listInvestors(request?: ListInvestorsRequest): Promise<ListInvestorsResponse>;
}
export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.params = this.params.bind(this);
    this.queryInvestorWallets = this.queryInvestorWallets.bind(this);
    this.queryByWallet = this.queryByWallet.bind(this);
    this.listInvestors = this.listInvestors.bind(this);
  }

  params(request: QueryParamsRequest = {}): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("joltifyfinance.joltify_lending.kyc.Query", "Params", data);
    return promise.then(data => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  queryInvestorWallets(request: QueryInvestorWalletsRequest): Promise<QueryInvestorWalletsResponse> {
    const data = QueryInvestorWalletsRequest.encode(request).finish();
    const promise = this.rpc.request("joltifyfinance.joltify_lending.kyc.Query", "QueryInvestorWallets", data);
    return promise.then(data => QueryInvestorWalletsResponse.decode(new _m0.Reader(data)));
  }

  queryByWallet(request: QueryByWalletRequest): Promise<QueryByWalletResponse> {
    const data = QueryByWalletRequest.encode(request).finish();
    const promise = this.rpc.request("joltifyfinance.joltify_lending.kyc.Query", "QueryByWallet", data);
    return promise.then(data => QueryByWalletResponse.decode(new _m0.Reader(data)));
  }

  listInvestors(request: ListInvestorsRequest = {
    pagination: undefined
  }): Promise<ListInvestorsResponse> {
    const data = ListInvestorsRequest.encode(request).finish();
    const promise = this.rpc.request("joltifyfinance.joltify_lending.kyc.Query", "ListInvestors", data);
    return promise.then(data => ListInvestorsResponse.decode(new _m0.Reader(data)));
  }

}
export const createRpcQueryExtension = (base: QueryClient) => {
  const rpc = createProtobufRpcClient(base);
  const queryService = new QueryClientImpl(rpc);
  return {
    params(request?: QueryParamsRequest): Promise<QueryParamsResponse> {
      return queryService.params(request);
    },

    queryInvestorWallets(request: QueryInvestorWalletsRequest): Promise<QueryInvestorWalletsResponse> {
      return queryService.queryInvestorWallets(request);
    },

    queryByWallet(request: QueryByWalletRequest): Promise<QueryByWalletResponse> {
      return queryService.queryByWallet(request);
    },

    listInvestors(request?: ListInvestorsRequest): Promise<ListInvestorsResponse> {
      return queryService.listInvestors(request);
    }

  };
};