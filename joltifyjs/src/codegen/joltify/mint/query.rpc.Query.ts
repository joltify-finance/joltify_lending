import { Rpc } from "../../helpers";
import * as _m0 from "protobufjs/minimal";
import { QueryClient, createProtobufRpcClient } from "@cosmjs/stargate";
import { QueryParamsRequest, QueryParamsResponse, QueryHistoryDistRequest, QueryHistoryDistResponse } from "./query";
/** Query defines the gRPC querier service. */

export interface Query {
  /** Parameters queries the parameters of the module. */
  params(request?: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Distribution queries the parameters of the module. */

  distribution(request?: QueryHistoryDistRequest): Promise<QueryHistoryDistResponse>;
}
export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.params = this.params.bind(this);
    this.distribution = this.distribution.bind(this);
  }

  params(request: QueryParamsRequest = {}): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.mint.Query", "Params", data);
    return promise.then(data => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  distribution(request: QueryHistoryDistRequest = {}): Promise<QueryHistoryDistResponse> {
    const data = QueryHistoryDistRequest.encode(request).finish();
    const promise = this.rpc.request("joltify.mint.Query", "Distribution", data);
    return promise.then(data => QueryHistoryDistResponse.decode(new _m0.Reader(data)));
  }

}
export const createRpcQueryExtension = (base: QueryClient) => {
  const rpc = createProtobufRpcClient(base);
  const queryService = new QueryClientImpl(rpc);
  return {
    params(request?: QueryParamsRequest): Promise<QueryParamsResponse> {
      return queryService.params(request);
    },

    distribution(request?: QueryHistoryDistRequest): Promise<QueryHistoryDistResponse> {
      return queryService.distribution(request);
    }

  };
};