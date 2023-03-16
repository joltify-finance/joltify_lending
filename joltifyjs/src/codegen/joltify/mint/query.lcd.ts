import { LCDClient } from "@osmonauts/lcd";
import { QueryParamsRequest, QueryParamsResponseSDKType, QueryHistoryDistRequest, QueryHistoryDistResponseSDKType } from "./query";
export class LCDQueryClient {
  req: LCDClient;

  constructor({
    requestClient
  }: {
    requestClient: LCDClient;
  }) {
    this.req = requestClient;
    this.params = this.params.bind(this);
    this.distribution = this.distribution.bind(this);
  }
  /* Parameters queries the parameters of the module. */


  async params(_params: QueryParamsRequest = {}): Promise<QueryParamsResponseSDKType> {
    const endpoint = `joltify-finance/joltify_lending/mint/params`;
    return await this.req.get<QueryParamsResponseSDKType>(endpoint);
  }
  /* Distribution queries the parameters of the module. */


  async distribution(_params: QueryHistoryDistRequest = {}): Promise<QueryHistoryDistResponseSDKType> {
    const endpoint = `joltify-finance/joltify_lending/mint/dist`;
    return await this.req.get<QueryHistoryDistResponseSDKType>(endpoint);
  }

}