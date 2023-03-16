import { setPaginationParams } from "../../helpers";
import { LCDClient } from "@osmonauts/lcd";
import { QueryParamsRequest, QueryParamsResponseSDKType, QueryInvestorWalletsRequest, QueryInvestorWalletsResponseSDKType, QueryByWalletRequest, QueryByWalletResponseSDKType, ListInvestorsRequest, ListInvestorsResponseSDKType } from "./query";
export class LCDQueryClient {
  req: LCDClient;

  constructor({
    requestClient
  }: {
    requestClient: LCDClient;
  }) {
    this.req = requestClient;
    this.params = this.params.bind(this);
    this.queryInvestorWallets = this.queryInvestorWallets.bind(this);
    this.queryByWallet = this.queryByWallet.bind(this);
    this.listInvestors = this.listInvestors.bind(this);
  }
  /* Parameters queries the parameters of the module. */


  async params(_params: QueryParamsRequest = {}): Promise<QueryParamsResponseSDKType> {
    const endpoint = `joltify/kyc/params`;
    return await this.req.get<QueryParamsResponseSDKType>(endpoint);
  }
  /* Queries a list of QueryInvestorWallets items. */


  async queryInvestorWallets(params: QueryInvestorWalletsRequest): Promise<QueryInvestorWalletsResponseSDKType> {
    const endpoint = `joltify/kyc/query_investor_wallets/${params.investorId}`;
    return await this.req.get<QueryInvestorWalletsResponseSDKType>(endpoint);
  }
  /* Queries a list of QueryByWallet items. */


  async queryByWallet(params: QueryByWalletRequest): Promise<QueryByWalletResponseSDKType> {
    const endpoint = `joltify/kyc/query_by_wallet/${params.wallet}`;
    return await this.req.get<QueryByWalletResponseSDKType>(endpoint);
  }
  /* Queries a list of ListInvestors items. */


  async listInvestors(params: ListInvestorsRequest = {
    pagination: undefined
  }): Promise<ListInvestorsResponseSDKType> {
    const options: any = {
      params: {}
    };

    if (typeof params?.pagination !== "undefined") {
      setPaginationParams(options, params.pagination);
    }

    const endpoint = `joltify/kyc/list_investors`;
    return await this.req.get<ListInvestorsResponseSDKType>(endpoint, options);
  }

}