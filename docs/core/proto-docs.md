 <!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [joltify/auction/v1beta1/auction.proto](#joltify/auction/v1beta1/auction.proto)
    - [BaseAuction](#joltify.auction.v1beta1.BaseAuction)
    - [CollateralAuction](#joltify.auction.v1beta1.CollateralAuction)
    - [DebtAuction](#joltify.auction.v1beta1.DebtAuction)
    - [SurplusAuction](#joltify.auction.v1beta1.SurplusAuction)
    - [WeightedAddresses](#joltify.auction.v1beta1.WeightedAddresses)
  
- [joltify/auction/v1beta1/genesis.proto](#joltify/auction/v1beta1/genesis.proto)
    - [GenesisState](#joltify.auction.v1beta1.GenesisState)
    - [Params](#joltify.auction.v1beta1.Params)
  
- [joltify/auction/v1beta1/query.proto](#joltify/auction/v1beta1/query.proto)
    - [QueryAuctionRequest](#joltify.auction.v1beta1.QueryAuctionRequest)
    - [QueryAuctionResponse](#joltify.auction.v1beta1.QueryAuctionResponse)
    - [QueryAuctionsRequest](#joltify.auction.v1beta1.QueryAuctionsRequest)
    - [QueryAuctionsResponse](#joltify.auction.v1beta1.QueryAuctionsResponse)
    - [QueryNextAuctionIDRequest](#joltify.auction.v1beta1.QueryNextAuctionIDRequest)
    - [QueryNextAuctionIDResponse](#joltify.auction.v1beta1.QueryNextAuctionIDResponse)
    - [QueryParamsRequest](#joltify.auction.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#joltify.auction.v1beta1.QueryParamsResponse)
  
    - [Query](#joltify.auction.v1beta1.Query)
  
- [joltify/auction/v1beta1/tx.proto](#joltify/auction/v1beta1/tx.proto)
    - [MsgPlaceBid](#joltify.auction.v1beta1.MsgPlaceBid)
    - [MsgPlaceBidResponse](#joltify.auction.v1beta1.MsgPlaceBidResponse)
  
    - [Msg](#joltify.auction.v1beta1.Msg)
  
- [joltify/cdp/v1beta1/cdp.proto](#joltify/cdp/v1beta1/cdp.proto)
    - [CDP](#joltify.cdp.v1beta1.CDP)
    - [Deposit](#joltify.cdp.v1beta1.Deposit)
    - [OwnerCDPIndex](#joltify.cdp.v1beta1.OwnerCDPIndex)
    - [TotalCollateral](#joltify.cdp.v1beta1.TotalCollateral)
    - [TotalPrincipal](#joltify.cdp.v1beta1.TotalPrincipal)
  
- [joltify/cdp/v1beta1/genesis.proto](#joltify/cdp/v1beta1/genesis.proto)
    - [CollateralParam](#joltify.cdp.v1beta1.CollateralParam)
    - [DebtParam](#joltify.cdp.v1beta1.DebtParam)
    - [GenesisAccumulationTime](#joltify.cdp.v1beta1.GenesisAccumulationTime)
    - [GenesisState](#joltify.cdp.v1beta1.GenesisState)
    - [GenesisTotalPrincipal](#joltify.cdp.v1beta1.GenesisTotalPrincipal)
    - [Params](#joltify.cdp.v1beta1.Params)
  
- [joltify/cdp/v1beta1/query.proto](#joltify/cdp/v1beta1/query.proto)
    - [CDPResponse](#joltify.cdp.v1beta1.CDPResponse)
    - [QueryAccountsRequest](#joltify.cdp.v1beta1.QueryAccountsRequest)
    - [QueryAccountsResponse](#joltify.cdp.v1beta1.QueryAccountsResponse)
    - [QueryCdpRequest](#joltify.cdp.v1beta1.QueryCdpRequest)
    - [QueryCdpResponse](#joltify.cdp.v1beta1.QueryCdpResponse)
    - [QueryCdpsRequest](#joltify.cdp.v1beta1.QueryCdpsRequest)
    - [QueryCdpsResponse](#joltify.cdp.v1beta1.QueryCdpsResponse)
    - [QueryDepositsRequest](#joltify.cdp.v1beta1.QueryDepositsRequest)
    - [QueryDepositsResponse](#joltify.cdp.v1beta1.QueryDepositsResponse)
    - [QueryParamsRequest](#joltify.cdp.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#joltify.cdp.v1beta1.QueryParamsResponse)
    - [QueryTotalCollateralRequest](#joltify.cdp.v1beta1.QueryTotalCollateralRequest)
    - [QueryTotalCollateralResponse](#joltify.cdp.v1beta1.QueryTotalCollateralResponse)
    - [QueryTotalPrincipalRequest](#joltify.cdp.v1beta1.QueryTotalPrincipalRequest)
    - [QueryTotalPrincipalResponse](#joltify.cdp.v1beta1.QueryTotalPrincipalResponse)
  
    - [Query](#joltify.cdp.v1beta1.Query)
  
- [joltify/cdp/v1beta1/tx.proto](#joltify/cdp/v1beta1/tx.proto)
    - [MsgCreateCDP](#joltify.cdp.v1beta1.MsgCreateCDP)
    - [MsgCreateCDPResponse](#joltify.cdp.v1beta1.MsgCreateCDPResponse)
    - [MsgDeposit](#joltify.cdp.v1beta1.MsgDeposit)
    - [MsgDepositResponse](#joltify.cdp.v1beta1.MsgDepositResponse)
    - [MsgDrawDebt](#joltify.cdp.v1beta1.MsgDrawDebt)
    - [MsgDrawDebtResponse](#joltify.cdp.v1beta1.MsgDrawDebtResponse)
    - [MsgLiquidate](#joltify.cdp.v1beta1.MsgLiquidate)
    - [MsgLiquidateResponse](#joltify.cdp.v1beta1.MsgLiquidateResponse)
    - [MsgRepayDebt](#joltify.cdp.v1beta1.MsgRepayDebt)
    - [MsgRepayDebtResponse](#joltify.cdp.v1beta1.MsgRepayDebtResponse)
    - [MsgWithdraw](#joltify.cdp.v1beta1.MsgWithdraw)
    - [MsgWithdrawResponse](#joltify.cdp.v1beta1.MsgWithdrawResponse)
  
    - [Msg](#joltify.cdp.v1beta1.Msg)
  
- [joltify/incentive/v1beta1/claims.proto](#joltify/incentive/v1beta1/claims.proto)
    - [BaseClaim](#joltify.incentive.v1beta1.BaseClaim)
    - [BaseMultiClaim](#joltify.incentive.v1beta1.BaseMultiClaim)
    - [DelegatorClaim](#joltify.incentive.v1beta1.DelegatorClaim)
    - [JoltLiquidityProviderClaim](#joltify.incentive.v1beta1.JoltLiquidityProviderClaim)
    - [MultiRewardIndex](#joltify.incentive.v1beta1.MultiRewardIndex)
    - [MultiRewardIndexesProto](#joltify.incentive.v1beta1.MultiRewardIndexesProto)
    - [RewardIndex](#joltify.incentive.v1beta1.RewardIndex)
    - [RewardIndexesProto](#joltify.incentive.v1beta1.RewardIndexesProto)
    - [SavingsClaim](#joltify.incentive.v1beta1.SavingsClaim)
    - [SwapClaim](#joltify.incentive.v1beta1.SwapClaim)
    - [USDXMintingClaim](#joltify.incentive.v1beta1.USDXMintingClaim)
  
- [joltify/incentive/v1beta1/params.proto](#joltify/incentive/v1beta1/params.proto)
    - [MultiRewardPeriod](#joltify.incentive.v1beta1.MultiRewardPeriod)
    - [Multiplier](#joltify.incentive.v1beta1.Multiplier)
    - [MultipliersPerDenom](#joltify.incentive.v1beta1.MultipliersPerDenom)
    - [Params](#joltify.incentive.v1beta1.Params)
    - [RewardPeriod](#joltify.incentive.v1beta1.RewardPeriod)
  
- [joltify/incentive/v1beta1/genesis.proto](#joltify/incentive/v1beta1/genesis.proto)
    - [AccumulationTime](#joltify.incentive.v1beta1.AccumulationTime)
    - [GenesisRewardState](#joltify.incentive.v1beta1.GenesisRewardState)
    - [GenesisState](#joltify.incentive.v1beta1.GenesisState)
  
- [joltify/incentive/v1beta1/tx.proto](#joltify/incentive/v1beta1/tx.proto)
    - [MsgClaimDelegatorReward](#joltify.incentive.v1beta1.MsgClaimDelegatorReward)
    - [MsgClaimDelegatorRewardResponse](#joltify.incentive.v1beta1.MsgClaimDelegatorRewardResponse)
    - [MsgClaimJoltReward](#joltify.incentive.v1beta1.MsgClaimJoltReward)
    - [MsgClaimJoltRewardResponse](#joltify.incentive.v1beta1.MsgClaimJoltRewardResponse)
    - [MsgClaimSavingsReward](#joltify.incentive.v1beta1.MsgClaimSavingsReward)
    - [MsgClaimSavingsRewardResponse](#joltify.incentive.v1beta1.MsgClaimSavingsRewardResponse)
    - [MsgClaimSwapReward](#joltify.incentive.v1beta1.MsgClaimSwapReward)
    - [MsgClaimSwapRewardResponse](#joltify.incentive.v1beta1.MsgClaimSwapRewardResponse)
    - [MsgClaimUSDXMintingReward](#joltify.incentive.v1beta1.MsgClaimUSDXMintingReward)
    - [MsgClaimUSDXMintingRewardResponse](#joltify.incentive.v1beta1.MsgClaimUSDXMintingRewardResponse)
    - [Selection](#joltify.incentive.v1beta1.Selection)
  
    - [Msg](#joltify.incentive.v1beta1.Msg)
  
- [joltify/issuance/v1beta1/genesis.proto](#joltify/issuance/v1beta1/genesis.proto)
    - [Asset](#joltify.issuance.v1beta1.Asset)
    - [AssetSupply](#joltify.issuance.v1beta1.AssetSupply)
    - [GenesisState](#joltify.issuance.v1beta1.GenesisState)
    - [Params](#joltify.issuance.v1beta1.Params)
    - [RateLimit](#joltify.issuance.v1beta1.RateLimit)
  
- [joltify/issuance/v1beta1/query.proto](#joltify/issuance/v1beta1/query.proto)
    - [QueryParamsRequest](#joltify.issuance.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#joltify.issuance.v1beta1.QueryParamsResponse)
  
    - [Query](#joltify.issuance.v1beta1.Query)
  
- [joltify/issuance/v1beta1/tx.proto](#joltify/issuance/v1beta1/tx.proto)
    - [MsgBlockAddress](#joltify.issuance.v1beta1.MsgBlockAddress)
    - [MsgBlockAddressResponse](#joltify.issuance.v1beta1.MsgBlockAddressResponse)
    - [MsgIssueTokens](#joltify.issuance.v1beta1.MsgIssueTokens)
    - [MsgIssueTokensResponse](#joltify.issuance.v1beta1.MsgIssueTokensResponse)
    - [MsgRedeemTokens](#joltify.issuance.v1beta1.MsgRedeemTokens)
    - [MsgRedeemTokensResponse](#joltify.issuance.v1beta1.MsgRedeemTokensResponse)
    - [MsgSetPauseStatus](#joltify.issuance.v1beta1.MsgSetPauseStatus)
    - [MsgSetPauseStatusResponse](#joltify.issuance.v1beta1.MsgSetPauseStatusResponse)
    - [MsgUnblockAddress](#joltify.issuance.v1beta1.MsgUnblockAddress)
    - [MsgUnblockAddressResponse](#joltify.issuance.v1beta1.MsgUnblockAddressResponse)
  
    - [Msg](#joltify.issuance.v1beta1.Msg)
  
- [joltify/jolt/v1beta1/jolt.proto](#joltify/jolt/v1beta1/jolt.proto)
    - [Borrow](#joltify.jolt.v1beta1.Borrow)
    - [BorrowInterestFactor](#joltify.jolt.v1beta1.BorrowInterestFactor)
    - [BorrowLimit](#joltify.jolt.v1beta1.BorrowLimit)
    - [CoinsProto](#joltify.jolt.v1beta1.CoinsProto)
    - [Deposit](#joltify.jolt.v1beta1.Deposit)
    - [InterestRateModel](#joltify.jolt.v1beta1.InterestRateModel)
    - [MoneyMarket](#joltify.jolt.v1beta1.MoneyMarket)
    - [Params](#joltify.jolt.v1beta1.Params)
    - [SupplyInterestFactor](#joltify.jolt.v1beta1.SupplyInterestFactor)
  
- [joltify/jolt/v1beta1/genesis.proto](#joltify/jolt/v1beta1/genesis.proto)
    - [GenesisAccumulationTime](#joltify.jolt.v1beta1.GenesisAccumulationTime)
    - [GenesisState](#joltify.jolt.v1beta1.GenesisState)
  
- [joltify/jolt/v1beta1/query.proto](#joltify/jolt/v1beta1/query.proto)
    - [BorrowInterestFactorResponse](#joltify.jolt.v1beta1.BorrowInterestFactorResponse)
    - [BorrowResponse](#joltify.jolt.v1beta1.BorrowResponse)
    - [DepositResponse](#joltify.jolt.v1beta1.DepositResponse)
    - [InterestFactor](#joltify.jolt.v1beta1.InterestFactor)
    - [LiquidateItem](#joltify.jolt.v1beta1.LiquidateItem)
    - [MoneyMarketInterestRate](#joltify.jolt.v1beta1.MoneyMarketInterestRate)
    - [QueryAccountsRequest](#joltify.jolt.v1beta1.QueryAccountsRequest)
    - [QueryAccountsResponse](#joltify.jolt.v1beta1.QueryAccountsResponse)
    - [QueryBorrowsRequest](#joltify.jolt.v1beta1.QueryBorrowsRequest)
    - [QueryBorrowsResponse](#joltify.jolt.v1beta1.QueryBorrowsResponse)
    - [QueryDepositsRequest](#joltify.jolt.v1beta1.QueryDepositsRequest)
    - [QueryDepositsResponse](#joltify.jolt.v1beta1.QueryDepositsResponse)
    - [QueryInterestFactorsRequest](#joltify.jolt.v1beta1.QueryInterestFactorsRequest)
    - [QueryInterestFactorsResponse](#joltify.jolt.v1beta1.QueryInterestFactorsResponse)
    - [QueryInterestRateRequest](#joltify.jolt.v1beta1.QueryInterestRateRequest)
    - [QueryInterestRateResponse](#joltify.jolt.v1beta1.QueryInterestRateResponse)
    - [QueryLiquidateRequest](#joltify.jolt.v1beta1.QueryLiquidateRequest)
    - [QueryLiquidateResp](#joltify.jolt.v1beta1.QueryLiquidateResp)
    - [QueryParamsRequest](#joltify.jolt.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#joltify.jolt.v1beta1.QueryParamsResponse)
    - [QueryReservesRequest](#joltify.jolt.v1beta1.QueryReservesRequest)
    - [QueryReservesResponse](#joltify.jolt.v1beta1.QueryReservesResponse)
    - [QueryTotalBorrowedRequest](#joltify.jolt.v1beta1.QueryTotalBorrowedRequest)
    - [QueryTotalBorrowedResponse](#joltify.jolt.v1beta1.QueryTotalBorrowedResponse)
    - [QueryTotalDepositedRequest](#joltify.jolt.v1beta1.QueryTotalDepositedRequest)
    - [QueryTotalDepositedResponse](#joltify.jolt.v1beta1.QueryTotalDepositedResponse)
    - [QueryUnsyncedBorrowsRequest](#joltify.jolt.v1beta1.QueryUnsyncedBorrowsRequest)
    - [QueryUnsyncedBorrowsResponse](#joltify.jolt.v1beta1.QueryUnsyncedBorrowsResponse)
    - [QueryUnsyncedDepositsRequest](#joltify.jolt.v1beta1.QueryUnsyncedDepositsRequest)
    - [QueryUnsyncedDepositsResponse](#joltify.jolt.v1beta1.QueryUnsyncedDepositsResponse)
    - [SupplyInterestFactorResponse](#joltify.jolt.v1beta1.SupplyInterestFactorResponse)
  
    - [Query](#joltify.jolt.v1beta1.Query)
  
- [joltify/jolt/v1beta1/tx.proto](#joltify/jolt/v1beta1/tx.proto)
    - [MsgBorrow](#joltify.jolt.v1beta1.MsgBorrow)
    - [MsgBorrowResponse](#joltify.jolt.v1beta1.MsgBorrowResponse)
    - [MsgDeposit](#joltify.jolt.v1beta1.MsgDeposit)
    - [MsgDepositResponse](#joltify.jolt.v1beta1.MsgDepositResponse)
    - [MsgLiquidate](#joltify.jolt.v1beta1.MsgLiquidate)
    - [MsgLiquidateResponse](#joltify.jolt.v1beta1.MsgLiquidateResponse)
    - [MsgRepay](#joltify.jolt.v1beta1.MsgRepay)
    - [MsgRepayResponse](#joltify.jolt.v1beta1.MsgRepayResponse)
    - [MsgWithdraw](#joltify.jolt.v1beta1.MsgWithdraw)
    - [MsgWithdrawResponse](#joltify.jolt.v1beta1.MsgWithdrawResponse)
  
    - [Msg](#joltify.jolt.v1beta1.Msg)
  
- [joltify/mint/dist.proto](#joltify/mint/dist.proto)
    - [HistoricalDistInfo](#joltify.mint.HistoricalDistInfo)
  
- [joltify/mint/params.proto](#joltify/mint/params.proto)
    - [Params](#joltify.mint.Params)
  
- [joltify/mint/genesis.proto](#joltify/mint/genesis.proto)
    - [GenesisState](#joltify.mint.GenesisState)
  
- [joltify/mint/query.proto](#joltify/mint/query.proto)
    - [QueryHistoryDistRequest](#joltify.mint.QueryHistoryDistRequest)
    - [QueryHistoryDistResponse](#joltify.mint.QueryHistoryDistResponse)
    - [QueryParamsRequest](#joltify.mint.QueryParamsRequest)
    - [QueryParamsResponse](#joltify.mint.QueryParamsResponse)
  
    - [Query](#joltify.mint.Query)
  
- [joltify/mint/tx.proto](#joltify/mint/tx.proto)
    - [Msg](#joltify.mint.Msg)
  
- [joltify/pricefeed/v1beta1/store.proto](#joltify/pricefeed/v1beta1/store.proto)
    - [CurrentPrice](#joltify.pricefeed.v1beta1.CurrentPrice)
    - [Market](#joltify.pricefeed.v1beta1.Market)
    - [Params](#joltify.pricefeed.v1beta1.Params)
    - [PostedPrice](#joltify.pricefeed.v1beta1.PostedPrice)
  
- [joltify/pricefeed/v1beta1/genesis.proto](#joltify/pricefeed/v1beta1/genesis.proto)
    - [GenesisState](#joltify.pricefeed.v1beta1.GenesisState)
  
- [joltify/pricefeed/v1beta1/query.proto](#joltify/pricefeed/v1beta1/query.proto)
    - [CurrentPriceResponse](#joltify.pricefeed.v1beta1.CurrentPriceResponse)
    - [MarketResponse](#joltify.pricefeed.v1beta1.MarketResponse)
    - [PostedPriceResponse](#joltify.pricefeed.v1beta1.PostedPriceResponse)
    - [QueryMarketsRequest](#joltify.pricefeed.v1beta1.QueryMarketsRequest)
    - [QueryMarketsResponse](#joltify.pricefeed.v1beta1.QueryMarketsResponse)
    - [QueryOraclesRequest](#joltify.pricefeed.v1beta1.QueryOraclesRequest)
    - [QueryOraclesResponse](#joltify.pricefeed.v1beta1.QueryOraclesResponse)
    - [QueryParamsRequest](#joltify.pricefeed.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#joltify.pricefeed.v1beta1.QueryParamsResponse)
    - [QueryPriceRequest](#joltify.pricefeed.v1beta1.QueryPriceRequest)
    - [QueryPriceResponse](#joltify.pricefeed.v1beta1.QueryPriceResponse)
    - [QueryPricesRequest](#joltify.pricefeed.v1beta1.QueryPricesRequest)
    - [QueryPricesResponse](#joltify.pricefeed.v1beta1.QueryPricesResponse)
    - [QueryRawPricesRequest](#joltify.pricefeed.v1beta1.QueryRawPricesRequest)
    - [QueryRawPricesResponse](#joltify.pricefeed.v1beta1.QueryRawPricesResponse)
  
    - [Query](#joltify.pricefeed.v1beta1.Query)
  
- [joltify/pricefeed/v1beta1/tx.proto](#joltify/pricefeed/v1beta1/tx.proto)
    - [MsgPostPrice](#joltify.pricefeed.v1beta1.MsgPostPrice)
    - [MsgPostPriceResponse](#joltify.pricefeed.v1beta1.MsgPostPriceResponse)
  
    - [Msg](#joltify.pricefeed.v1beta1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="joltify/auction/v1beta1/auction.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/auction/v1beta1/auction.proto



<a name="joltify.auction.v1beta1.BaseAuction"></a>

### BaseAuction
BaseAuction defines common attributes of all auctions


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `initiator` | [string](#string) |  |  |
| `lot` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `bidder` | [bytes](#bytes) |  |  |
| `bid` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `has_received_bids` | [bool](#bool) |  |  |
| `end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `max_end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="joltify.auction.v1beta1.CollateralAuction"></a>

### CollateralAuction
CollateralAuction is a two phase auction.
Initially, in forward auction phase, bids can be placed up to a max bid.
Then it switches to a reverse auction phase, where the initial amount up for auction is bid down.
Unsold Lot is sent to LotReturns, being divided among the addresses by weight.
Collateral auctions are normally used to sell off collateral seized from CDPs.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_auction` | [BaseAuction](#joltify.auction.v1beta1.BaseAuction) |  |  |
| `corresponding_debt` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `max_bid` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `lot_returns` | [WeightedAddresses](#joltify.auction.v1beta1.WeightedAddresses) |  |  |






<a name="joltify.auction.v1beta1.DebtAuction"></a>

### DebtAuction
DebtAuction is a reverse auction that mints what it pays out.
It is normally used to acquire pegged asset to cover the CDP system's debts that were not covered by selling
collateral.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_auction` | [BaseAuction](#joltify.auction.v1beta1.BaseAuction) |  |  |
| `corresponding_debt` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="joltify.auction.v1beta1.SurplusAuction"></a>

### SurplusAuction
SurplusAuction is a forward auction that burns what it receives from bids.
It is normally used to sell off excess pegged asset acquired by the CDP system.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_auction` | [BaseAuction](#joltify.auction.v1beta1.BaseAuction) |  |  |






<a name="joltify.auction.v1beta1.WeightedAddresses"></a>

### WeightedAddresses
WeightedAddresses is a type for storing some addresses and associated weights.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `addresses` | [bytes](#bytes) | repeated |  |
| `weights` | [bytes](#bytes) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="joltify/auction/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/auction/v1beta1/genesis.proto



<a name="joltify.auction.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the auction module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `next_auction_id` | [uint64](#uint64) |  |  |
| `params` | [Params](#joltify.auction.v1beta1.Params) |  |  |
| `auctions` | [google.protobuf.Any](#google.protobuf.Any) | repeated | Genesis auctions |






<a name="joltify.auction.v1beta1.Params"></a>

### Params
Params defines the parameters for the issuance module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `max_auction_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `forward_bid_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `reverse_bid_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `increment_surplus` | [bytes](#bytes) |  |  |
| `increment_debt` | [bytes](#bytes) |  |  |
| `increment_collateral` | [bytes](#bytes) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="joltify/auction/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/auction/v1beta1/query.proto



<a name="joltify.auction.v1beta1.QueryAuctionRequest"></a>

### QueryAuctionRequest
QueryAuctionRequest is the request type for the Query/Auction RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auction_id` | [uint64](#uint64) |  |  |






<a name="joltify.auction.v1beta1.QueryAuctionResponse"></a>

### QueryAuctionResponse
QueryAuctionResponse is the response type for the Query/Auction RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auction` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="joltify.auction.v1beta1.QueryAuctionsRequest"></a>

### QueryAuctionsRequest
QueryAuctionsRequest is the request type for the Query/Auctions RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `type` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `denom` | [string](#string) |  |  |
| `phase` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="joltify.auction.v1beta1.QueryAuctionsResponse"></a>

### QueryAuctionsResponse
QueryAuctionsResponse is the response type for the Query/Auctions RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auctions` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="joltify.auction.v1beta1.QueryNextAuctionIDRequest"></a>

### QueryNextAuctionIDRequest
QueryNextAuctionIDRequest defines the request type for querying x/auction next auction ID.






<a name="joltify.auction.v1beta1.QueryNextAuctionIDResponse"></a>

### QueryNextAuctionIDResponse
QueryNextAuctionIDResponse defines the response type for querying x/auction next auction ID.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |






<a name="joltify.auction.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for querying x/auction parameters.






<a name="joltify.auction.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for querying x/auction parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#joltify.auction.v1beta1.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="joltify.auction.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for auction module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#joltify.auction.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#joltify.auction.v1beta1.QueryParamsResponse) | Params queries all parameters of the auction module. | GET|/joltify/auction/v1beta1/params|
| `Auction` | [QueryAuctionRequest](#joltify.auction.v1beta1.QueryAuctionRequest) | [QueryAuctionResponse](#joltify.auction.v1beta1.QueryAuctionResponse) | Auction queries an individual Auction by auction ID | GET|/joltify/auction/v1beta1/auctions/{auction_id}|
| `Auctions` | [QueryAuctionsRequest](#joltify.auction.v1beta1.QueryAuctionsRequest) | [QueryAuctionsResponse](#joltify.auction.v1beta1.QueryAuctionsResponse) | Auctions queries auctions filtered by asset denom, owner address, phase, and auction type | GET|/joltify/auction/v1beta1/auctions|
| `NextAuctionID` | [QueryNextAuctionIDRequest](#joltify.auction.v1beta1.QueryNextAuctionIDRequest) | [QueryNextAuctionIDResponse](#joltify.auction.v1beta1.QueryNextAuctionIDResponse) | NextAuctionID queries the next auction ID | GET|/joltify/auction/v1beta1/next-auction-id|

 <!-- end services -->



<a name="joltify/auction/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/auction/v1beta1/tx.proto



<a name="joltify.auction.v1beta1.MsgPlaceBid"></a>

### MsgPlaceBid
MsgPlaceBid represents a message used by bidders to place bids on auctions


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auction_id` | [uint64](#uint64) |  |  |
| `bidder` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="joltify.auction.v1beta1.MsgPlaceBidResponse"></a>

### MsgPlaceBidResponse
MsgPlaceBidResponse defines the Msg/PlaceBid response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="joltify.auction.v1beta1.Msg"></a>

### Msg
Msg defines the auction Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `PlaceBid` | [MsgPlaceBid](#joltify.auction.v1beta1.MsgPlaceBid) | [MsgPlaceBidResponse](#joltify.auction.v1beta1.MsgPlaceBidResponse) | PlaceBid message type used by bidders to place bids on auctions | |

 <!-- end services -->



<a name="joltify/cdp/v1beta1/cdp.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/cdp/v1beta1/cdp.proto



<a name="joltify.cdp.v1beta1.CDP"></a>

### CDP
CDP defines the state of a single collateralized debt position.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `owner` | [bytes](#bytes) |  |  |
| `type` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `principal` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `accumulated_fees` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fees_updated` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `interest_factor` | [string](#string) |  |  |






<a name="joltify.cdp.v1beta1.Deposit"></a>

### Deposit
Deposit defines an amount of coins deposited by an account to a cdp


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp_id` | [uint64](#uint64) |  |  |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="joltify.cdp.v1beta1.OwnerCDPIndex"></a>

### OwnerCDPIndex
OwnerCDPIndex defines the cdp ids for a single cdp owner


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp_ids` | [uint64](#uint64) | repeated |  |






<a name="joltify.cdp.v1beta1.TotalCollateral"></a>

### TotalCollateral
TotalCollateral defines the total collateral of a given collateral type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="joltify.cdp.v1beta1.TotalPrincipal"></a>

### TotalPrincipal
TotalPrincipal defines the total principal of a given collateral type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="joltify/cdp/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/cdp/v1beta1/genesis.proto



<a name="joltify.cdp.v1beta1.CollateralParam"></a>

### CollateralParam
CollateralParam defines governance parameters for each collateral type within the cdp module


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |
| `liquidation_ratio` | [string](#string) |  |  |
| `debt_limit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `stability_fee` | [string](#string) |  |  |
| `auction_size` | [string](#string) |  |  |
| `liquidation_penalty` | [string](#string) |  |  |
| `spot_market_id` | [string](#string) |  |  |
| `liquidation_market_id` | [string](#string) |  |  |
| `keeper_reward_percentage` | [string](#string) |  |  |
| `check_collateralization_index_count` | [string](#string) |  |  |
| `conversion_factor` | [string](#string) |  |  |






<a name="joltify.cdp.v1beta1.DebtParam"></a>

### DebtParam
DebtParam defines governance params for debt assets


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `reference_asset` | [string](#string) |  |  |
| `conversion_factor` | [string](#string) |  |  |
| `debt_floor` | [string](#string) |  |  |






<a name="joltify.cdp.v1beta1.GenesisAccumulationTime"></a>

### GenesisAccumulationTime
GenesisAccumulationTime defines the previous distribution time and its corresponding denom


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `previous_accumulation_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `interest_factor` | [string](#string) |  |  |






<a name="joltify.cdp.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the cdp module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#joltify.cdp.v1beta1.Params) |  | params defines all the paramaters of the module. |
| `cdps` | [CDP](#joltify.cdp.v1beta1.CDP) | repeated |  |
| `deposits` | [Deposit](#joltify.cdp.v1beta1.Deposit) | repeated |  |
| `starting_cdp_id` | [uint64](#uint64) |  |  |
| `debt_denom` | [string](#string) |  |  |
| `gov_denom` | [string](#string) |  |  |
| `previous_accumulation_times` | [GenesisAccumulationTime](#joltify.cdp.v1beta1.GenesisAccumulationTime) | repeated |  |
| `total_principals` | [GenesisTotalPrincipal](#joltify.cdp.v1beta1.GenesisTotalPrincipal) | repeated |  |






<a name="joltify.cdp.v1beta1.GenesisTotalPrincipal"></a>

### GenesisTotalPrincipal
GenesisTotalPrincipal defines the total principal and its corresponding collateral type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `total_principal` | [string](#string) |  |  |






<a name="joltify.cdp.v1beta1.Params"></a>

### Params
Params defines the parameters for the cdp module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_params` | [CollateralParam](#joltify.cdp.v1beta1.CollateralParam) | repeated |  |
| `debt_param` | [DebtParam](#joltify.cdp.v1beta1.DebtParam) |  |  |
| `global_debt_limit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `surplus_auction_threshold` | [string](#string) |  |  |
| `surplus_auction_lot` | [string](#string) |  |  |
| `debt_auction_threshold` | [string](#string) |  |  |
| `debt_auction_lot` | [string](#string) |  |  |
| `circuit_breaker` | [bool](#bool) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="joltify/cdp/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/cdp/v1beta1/query.proto



<a name="joltify.cdp.v1beta1.CDPResponse"></a>

### CDPResponse
CDPResponse defines the state of a single collateralized debt position.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `owner` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `principal` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `accumulated_fees` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fees_updated` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `interest_factor` | [string](#string) |  |  |
| `collateral_value` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateralization_ratio` | [string](#string) |  |  |






<a name="joltify.cdp.v1beta1.QueryAccountsRequest"></a>

### QueryAccountsRequest
QueryAccountsRequest defines the request type for the Query/Accounts RPC method.






<a name="joltify.cdp.v1beta1.QueryAccountsResponse"></a>

### QueryAccountsResponse
QueryAccountsResponse defines the response type for the Query/Accounts RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `accounts` | [cosmos.auth.v1beta1.ModuleAccount](#cosmos.auth.v1beta1.ModuleAccount) | repeated |  |






<a name="joltify.cdp.v1beta1.QueryCdpRequest"></a>

### QueryCdpRequest
QueryCdpRequest defines the request type for the Query/Cdp RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |






<a name="joltify.cdp.v1beta1.QueryCdpResponse"></a>

### QueryCdpResponse
QueryCdpResponse defines the response type for the Query/Cdp RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp` | [CDPResponse](#joltify.cdp.v1beta1.CDPResponse) |  |  |






<a name="joltify.cdp.v1beta1.QueryCdpsRequest"></a>

### QueryCdpsRequest
QueryCdpsRequest is the params for a filtered CDP query, the request type for the Query/Cdps RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `id` | [uint64](#uint64) |  |  |
| `ratio` | [string](#string) |  | sdk.Dec as a string |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="joltify.cdp.v1beta1.QueryCdpsResponse"></a>

### QueryCdpsResponse
QueryCdpsResponse defines the response type for the Query/Cdps RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdps` | [CDPResponse](#joltify.cdp.v1beta1.CDPResponse) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="joltify.cdp.v1beta1.QueryDepositsRequest"></a>

### QueryDepositsRequest
QueryDepositsRequest defines the request type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |






<a name="joltify.cdp.v1beta1.QueryDepositsResponse"></a>

### QueryDepositsResponse
QueryDepositsResponse defines the response type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposits` | [Deposit](#joltify.cdp.v1beta1.Deposit) | repeated |  |






<a name="joltify.cdp.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for the Query/Params RPC method.






<a name="joltify.cdp.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#joltify.cdp.v1beta1.Params) |  |  |






<a name="joltify.cdp.v1beta1.QueryTotalCollateralRequest"></a>

### QueryTotalCollateralRequest
QueryTotalCollateralRequest defines the request type for the Query/TotalCollateral RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |






<a name="joltify.cdp.v1beta1.QueryTotalCollateralResponse"></a>

### QueryTotalCollateralResponse
QueryTotalCollateralResponse defines the response type for the Query/TotalCollateral RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `total_collateral` | [TotalCollateral](#joltify.cdp.v1beta1.TotalCollateral) | repeated |  |






<a name="joltify.cdp.v1beta1.QueryTotalPrincipalRequest"></a>

### QueryTotalPrincipalRequest
QueryTotalPrincipalRequest defines the request type for the Query/TotalPrincipal RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |






<a name="joltify.cdp.v1beta1.QueryTotalPrincipalResponse"></a>

### QueryTotalPrincipalResponse
QueryTotalPrincipalResponse defines the response type for the Query/TotalPrincipal RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `total_principal` | [TotalPrincipal](#joltify.cdp.v1beta1.TotalPrincipal) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="joltify.cdp.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for cdp module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#joltify.cdp.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#joltify.cdp.v1beta1.QueryParamsResponse) | Params queries all parameters of the cdp module. | GET|/joltify/cdp/v1beta1/params|
| `Accounts` | [QueryAccountsRequest](#joltify.cdp.v1beta1.QueryAccountsRequest) | [QueryAccountsResponse](#joltify.cdp.v1beta1.QueryAccountsResponse) | Accounts queries the CDP module accounts. | GET|/joltify/cdp/v1beta1/accounts|
| `TotalPrincipal` | [QueryTotalPrincipalRequest](#joltify.cdp.v1beta1.QueryTotalPrincipalRequest) | [QueryTotalPrincipalResponse](#joltify.cdp.v1beta1.QueryTotalPrincipalResponse) | TotalPrincipal queries the total principal of a given collateral type. | GET|/joltify/cdp/v1beta1/totalPrincipal|
| `TotalCollateral` | [QueryTotalCollateralRequest](#joltify.cdp.v1beta1.QueryTotalCollateralRequest) | [QueryTotalCollateralResponse](#joltify.cdp.v1beta1.QueryTotalCollateralResponse) | TotalCollateral queries the total collateral of a given collateral type. | GET|/joltify/cdp/v1beta1/totalCollateral|
| `Cdps` | [QueryCdpsRequest](#joltify.cdp.v1beta1.QueryCdpsRequest) | [QueryCdpsResponse](#joltify.cdp.v1beta1.QueryCdpsResponse) | Cdps queries all active CDPs. | GET|/joltify/cdp/v1beta1/cdps|
| `Cdp` | [QueryCdpRequest](#joltify.cdp.v1beta1.QueryCdpRequest) | [QueryCdpResponse](#joltify.cdp.v1beta1.QueryCdpResponse) | Cdp queries a CDP with the input owner address and collateral type. | GET|/joltify/cdp/v1beta1/cdps/{owner}/{collateral_type}|
| `Deposits` | [QueryDepositsRequest](#joltify.cdp.v1beta1.QueryDepositsRequest) | [QueryDepositsResponse](#joltify.cdp.v1beta1.QueryDepositsResponse) | Deposits queries deposits associated with the CDP owned by an address for a collateral type. | GET|/joltify/cdp/v1beta1/cdps/deposits/{owner}/{collateral_type}|

 <!-- end services -->



<a name="joltify/cdp/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/cdp/v1beta1/tx.proto



<a name="joltify.cdp.v1beta1.MsgCreateCDP"></a>

### MsgCreateCDP
MsgCreateCDP defines a message to create a new CDP.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `principal` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="joltify.cdp.v1beta1.MsgCreateCDPResponse"></a>

### MsgCreateCDPResponse
MsgCreateCDPResponse defines the Msg/CreateCDP response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp_id` | [uint64](#uint64) |  |  |






<a name="joltify.cdp.v1beta1.MsgDeposit"></a>

### MsgDeposit
MsgDeposit defines a message to deposit to a CDP.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="joltify.cdp.v1beta1.MsgDepositResponse"></a>

### MsgDepositResponse
MsgDepositResponse defines the Msg/Deposit response type.






<a name="joltify.cdp.v1beta1.MsgDrawDebt"></a>

### MsgDrawDebt
MsgDrawDebt defines a message to draw debt from a CDP.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `principal` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="joltify.cdp.v1beta1.MsgDrawDebtResponse"></a>

### MsgDrawDebtResponse
MsgDrawDebtResponse defines the Msg/DrawDebt response type.






<a name="joltify.cdp.v1beta1.MsgLiquidate"></a>

### MsgLiquidate
MsgLiquidate defines a message to attempt to liquidate a CDP whos
collateralization ratio is under its liquidation ratio.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `keeper` | [string](#string) |  |  |
| `borrower` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="joltify.cdp.v1beta1.MsgLiquidateResponse"></a>

### MsgLiquidateResponse
MsgLiquidateResponse defines the Msg/Liquidate response type.






<a name="joltify.cdp.v1beta1.MsgRepayDebt"></a>

### MsgRepayDebt
MsgRepayDebt defines a message to repay debt from a CDP.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `payment` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="joltify.cdp.v1beta1.MsgRepayDebtResponse"></a>

### MsgRepayDebtResponse
MsgRepayDebtResponse defines the Msg/RepayDebt response type.






<a name="joltify.cdp.v1beta1.MsgWithdraw"></a>

### MsgWithdraw
MsgWithdraw defines a message to withdraw collateral from a CDP.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="joltify.cdp.v1beta1.MsgWithdrawResponse"></a>

### MsgWithdrawResponse
MsgWithdrawResponse defines the Msg/Withdraw response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="joltify.cdp.v1beta1.Msg"></a>

### Msg
Msg defines the cdp Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateCDP` | [MsgCreateCDP](#joltify.cdp.v1beta1.MsgCreateCDP) | [MsgCreateCDPResponse](#joltify.cdp.v1beta1.MsgCreateCDPResponse) | CreateCDP defines a method to create a new CDP. | |
| `Deposit` | [MsgDeposit](#joltify.cdp.v1beta1.MsgDeposit) | [MsgDepositResponse](#joltify.cdp.v1beta1.MsgDepositResponse) | Deposit defines a method to deposit to a CDP. | |
| `Withdraw` | [MsgWithdraw](#joltify.cdp.v1beta1.MsgWithdraw) | [MsgWithdrawResponse](#joltify.cdp.v1beta1.MsgWithdrawResponse) | Withdraw defines a method to withdraw collateral from a CDP. | |
| `DrawDebt` | [MsgDrawDebt](#joltify.cdp.v1beta1.MsgDrawDebt) | [MsgDrawDebtResponse](#joltify.cdp.v1beta1.MsgDrawDebtResponse) | DrawDebt defines a method to draw debt from a CDP. | |
| `RepayDebt` | [MsgRepayDebt](#joltify.cdp.v1beta1.MsgRepayDebt) | [MsgRepayDebtResponse](#joltify.cdp.v1beta1.MsgRepayDebtResponse) | RepayDebt defines a method to repay debt from a CDP. | |
| `Liquidate` | [MsgLiquidate](#joltify.cdp.v1beta1.MsgLiquidate) | [MsgLiquidateResponse](#joltify.cdp.v1beta1.MsgLiquidateResponse) | Liquidate defines a method to attempt to liquidate a CDP whos collateralization ratio is under its liquidation ratio. | |

 <!-- end services -->



<a name="joltify/incentive/v1beta1/claims.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/incentive/v1beta1/claims.proto



<a name="joltify.incentive.v1beta1.BaseClaim"></a>

### BaseClaim
BaseClaim is a claim with a single reward coin types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [bytes](#bytes) |  |  |
| `reward` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="joltify.incentive.v1beta1.BaseMultiClaim"></a>

### BaseMultiClaim
BaseMultiClaim is a claim with multiple reward coin types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [bytes](#bytes) |  |  |
| `reward` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="joltify.incentive.v1beta1.DelegatorClaim"></a>

### DelegatorClaim
DelegatorClaim stores delegation rewards that can be claimed by owner


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_claim` | [BaseMultiClaim](#joltify.incentive.v1beta1.BaseMultiClaim) |  |  |
| `reward_indexes` | [MultiRewardIndex](#joltify.incentive.v1beta1.MultiRewardIndex) | repeated |  |






<a name="joltify.incentive.v1beta1.JoltLiquidityProviderClaim"></a>

### JoltLiquidityProviderClaim
JoltLiquidityProviderClaim stores the jolt liquidity provider rewards that can be claimed by owner


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_claim` | [BaseMultiClaim](#joltify.incentive.v1beta1.BaseMultiClaim) |  |  |
| `supply_reward_indexes` | [MultiRewardIndex](#joltify.incentive.v1beta1.MultiRewardIndex) | repeated |  |
| `borrow_reward_indexes` | [MultiRewardIndex](#joltify.incentive.v1beta1.MultiRewardIndex) | repeated |  |






<a name="joltify.incentive.v1beta1.MultiRewardIndex"></a>

### MultiRewardIndex
MultiRewardIndex stores reward accumulation information on multiple reward types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `reward_indexes` | [RewardIndex](#joltify.incentive.v1beta1.RewardIndex) | repeated |  |






<a name="joltify.incentive.v1beta1.MultiRewardIndexesProto"></a>

### MultiRewardIndexesProto
MultiRewardIndexesProto defines a Protobuf wrapper around a MultiRewardIndexes slice


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `multi_reward_indexes` | [MultiRewardIndex](#joltify.incentive.v1beta1.MultiRewardIndex) | repeated |  |






<a name="joltify.incentive.v1beta1.RewardIndex"></a>

### RewardIndex
RewardIndex stores reward accumulation information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `reward_factor` | [bytes](#bytes) |  |  |






<a name="joltify.incentive.v1beta1.RewardIndexesProto"></a>

### RewardIndexesProto
RewardIndexesProto defines a Protobuf wrapper around a RewardIndexes slice


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `reward_indexes` | [RewardIndex](#joltify.incentive.v1beta1.RewardIndex) | repeated |  |






<a name="joltify.incentive.v1beta1.SavingsClaim"></a>

### SavingsClaim
SavingsClaim stores the savings rewards that can be claimed by owner


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_claim` | [BaseMultiClaim](#joltify.incentive.v1beta1.BaseMultiClaim) |  |  |
| `reward_indexes` | [MultiRewardIndex](#joltify.incentive.v1beta1.MultiRewardIndex) | repeated |  |






<a name="joltify.incentive.v1beta1.SwapClaim"></a>

### SwapClaim
SwapClaim stores the swap rewards that can be claimed by owner


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_claim` | [BaseMultiClaim](#joltify.incentive.v1beta1.BaseMultiClaim) |  |  |
| `reward_indexes` | [MultiRewardIndex](#joltify.incentive.v1beta1.MultiRewardIndex) | repeated |  |






<a name="joltify.incentive.v1beta1.USDXMintingClaim"></a>

### USDXMintingClaim
USDXMintingClaim is for USDX minting rewards


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_claim` | [BaseClaim](#joltify.incentive.v1beta1.BaseClaim) |  |  |
| `reward_indexes` | [RewardIndex](#joltify.incentive.v1beta1.RewardIndex) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="joltify/incentive/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/incentive/v1beta1/params.proto



<a name="joltify.incentive.v1beta1.MultiRewardPeriod"></a>

### MultiRewardPeriod
MultiRewardPeriod supports multiple reward types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `active` | [bool](#bool) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `start` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `end` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `rewards_per_second` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="joltify.incentive.v1beta1.Multiplier"></a>

### Multiplier
Multiplier amount the claim rewards get increased by, along with how long the claim rewards are locked


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `months_lockup` | [int64](#int64) |  |  |
| `factor` | [bytes](#bytes) |  |  |






<a name="joltify.incentive.v1beta1.MultipliersPerDenom"></a>

### MultipliersPerDenom
MultipliersPerDenom is a map of denoms to a set of multipliers


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `multipliers` | [Multiplier](#joltify.incentive.v1beta1.Multiplier) | repeated |  |






<a name="joltify.incentive.v1beta1.Params"></a>

### Params
Params


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `usdx_minting_reward_periods` | [RewardPeriod](#joltify.incentive.v1beta1.RewardPeriod) | repeated |  |
| `jolt_supply_reward_periods` | [MultiRewardPeriod](#joltify.incentive.v1beta1.MultiRewardPeriod) | repeated |  |
| `jolt_borrow_reward_periods` | [MultiRewardPeriod](#joltify.incentive.v1beta1.MultiRewardPeriod) | repeated |  |
| `delegator_reward_periods` | [MultiRewardPeriod](#joltify.incentive.v1beta1.MultiRewardPeriod) | repeated |  |
| `swap_reward_periods` | [MultiRewardPeriod](#joltify.incentive.v1beta1.MultiRewardPeriod) | repeated |  |
| `claim_multipliers` | [MultipliersPerDenom](#joltify.incentive.v1beta1.MultipliersPerDenom) | repeated |  |
| `claim_end` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `savings_reward_periods` | [MultiRewardPeriod](#joltify.incentive.v1beta1.MultiRewardPeriod) | repeated |  |






<a name="joltify.incentive.v1beta1.RewardPeriod"></a>

### RewardPeriod
RewardPeriod stores the state of an ongoing reward


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `active` | [bool](#bool) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `start` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `end` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `rewards_per_second` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="joltify/incentive/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/incentive/v1beta1/genesis.proto



<a name="joltify.incentive.v1beta1.AccumulationTime"></a>

### AccumulationTime
AccumulationTime stores the previous reward distribution time and its corresponding collateral type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `previous_accumulation_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="joltify.incentive.v1beta1.GenesisRewardState"></a>

### GenesisRewardState
GenesisRewardState groups together the global state for a particular reward so it can be exported in genesis.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `accumulation_times` | [AccumulationTime](#joltify.incentive.v1beta1.AccumulationTime) | repeated |  |
| `multi_reward_indexes` | [MultiRewardIndex](#joltify.incentive.v1beta1.MultiRewardIndex) | repeated |  |






<a name="joltify.incentive.v1beta1.GenesisState"></a>

### GenesisState
GenesisState is the state that must be provided at genesis.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#joltify.incentive.v1beta1.Params) |  |  |
| `usdx_reward_state` | [GenesisRewardState](#joltify.incentive.v1beta1.GenesisRewardState) |  |  |
| `jolt_supply_reward_state` | [GenesisRewardState](#joltify.incentive.v1beta1.GenesisRewardState) |  |  |
| `jolt_borrow_reward_state` | [GenesisRewardState](#joltify.incentive.v1beta1.GenesisRewardState) |  |  |
| `delegator_reward_state` | [GenesisRewardState](#joltify.incentive.v1beta1.GenesisRewardState) |  |  |
| `swap_reward_state` | [GenesisRewardState](#joltify.incentive.v1beta1.GenesisRewardState) |  |  |
| `usdx_minting_claims` | [USDXMintingClaim](#joltify.incentive.v1beta1.USDXMintingClaim) | repeated |  |
| `jolt_liquidity_provider_claims` | [JoltLiquidityProviderClaim](#joltify.incentive.v1beta1.JoltLiquidityProviderClaim) | repeated |  |
| `delegator_claims` | [DelegatorClaim](#joltify.incentive.v1beta1.DelegatorClaim) | repeated |  |
| `swap_claims` | [SwapClaim](#joltify.incentive.v1beta1.SwapClaim) | repeated |  |
| `savings_reward_state` | [GenesisRewardState](#joltify.incentive.v1beta1.GenesisRewardState) |  |  |
| `savings_claims` | [SavingsClaim](#joltify.incentive.v1beta1.SavingsClaim) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="joltify/incentive/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/incentive/v1beta1/tx.proto



<a name="joltify.incentive.v1beta1.MsgClaimDelegatorReward"></a>

### MsgClaimDelegatorReward
MsgClaimDelegatorReward message type used to claim delegator rewards


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denoms_to_claim` | [Selection](#joltify.incentive.v1beta1.Selection) | repeated |  |






<a name="joltify.incentive.v1beta1.MsgClaimDelegatorRewardResponse"></a>

### MsgClaimDelegatorRewardResponse
MsgClaimDelegatorRewardResponse defines the Msg/ClaimDelegatorReward response type.






<a name="joltify.incentive.v1beta1.MsgClaimJoltReward"></a>

### MsgClaimJoltReward
MsgClaimHardReward message type used to claim Hard liquidity provider rewards


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denoms_to_claim` | [Selection](#joltify.incentive.v1beta1.Selection) | repeated |  |






<a name="joltify.incentive.v1beta1.MsgClaimJoltRewardResponse"></a>

### MsgClaimJoltRewardResponse
MsgClaimJoltRewardResponse defines the Msg/ClaimHardReward response type.






<a name="joltify.incentive.v1beta1.MsgClaimSavingsReward"></a>

### MsgClaimSavingsReward
MsgClaimSavingsReward message type used to claim savings rewards


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denoms_to_claim` | [Selection](#joltify.incentive.v1beta1.Selection) | repeated |  |






<a name="joltify.incentive.v1beta1.MsgClaimSavingsRewardResponse"></a>

### MsgClaimSavingsRewardResponse
MsgClaimSavingsRewardResponse defines the Msg/ClaimSavingsReward response type.






<a name="joltify.incentive.v1beta1.MsgClaimSwapReward"></a>

### MsgClaimSwapReward
MsgClaimSwapReward message type used to claim delegator rewards


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denoms_to_claim` | [Selection](#joltify.incentive.v1beta1.Selection) | repeated |  |






<a name="joltify.incentive.v1beta1.MsgClaimSwapRewardResponse"></a>

### MsgClaimSwapRewardResponse
MsgClaimSwapRewardResponse defines the Msg/ClaimSwapReward response type.






<a name="joltify.incentive.v1beta1.MsgClaimUSDXMintingReward"></a>

### MsgClaimUSDXMintingReward
MsgClaimUSDXMintingReward message type used to claim USDX minting rewards


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `multiplier_name` | [string](#string) |  |  |






<a name="joltify.incentive.v1beta1.MsgClaimUSDXMintingRewardResponse"></a>

### MsgClaimUSDXMintingRewardResponse
MsgClaimUSDXMintingRewardResponse defines the Msg/ClaimUSDXMintingReward response type.






<a name="joltify.incentive.v1beta1.Selection"></a>

### Selection
Selection is a pair of denom and multiplier name. It holds the choice of multiplier a user makes when they claim a
denom.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `multiplier_name` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="joltify.incentive.v1beta1.Msg"></a>

### Msg
Msg defines the incentive Msg service.

ClaimUSDXMintingReward is a message type used to claim USDX minting rewards
 rpc ClaimUSDXMintingReward(MsgClaimUSDXMintingReward) returns (MsgClaimUSDXMintingRewardResponse);

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ClaimJoltReward` | [MsgClaimJoltReward](#joltify.incentive.v1beta1.MsgClaimJoltReward) | [MsgClaimJoltRewardResponse](#joltify.incentive.v1beta1.MsgClaimJoltRewardResponse) | ClaimJoltReward is a message type used to claim Hard liquidity provider rewards | |

 <!-- end services -->



<a name="joltify/issuance/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/issuance/v1beta1/genesis.proto



<a name="joltify.issuance.v1beta1.Asset"></a>

### Asset
Asset type for assets in the issuance module


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `denom` | [string](#string) |  |  |
| `blocked_addresses` | [string](#string) | repeated |  |
| `paused` | [bool](#bool) |  |  |
| `blockable` | [bool](#bool) |  |  |
| `rate_limit` | [RateLimit](#joltify.issuance.v1beta1.RateLimit) |  |  |






<a name="joltify.issuance.v1beta1.AssetSupply"></a>

### AssetSupply
AssetSupply contains information about an asset's rate-limited supply (the
total supply of the asset is tracked in the top-level supply module)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `current_supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `time_elapsed` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="joltify.issuance.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the issuance module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#joltify.issuance.v1beta1.Params) |  | params defines all the paramaters of the module. |
| `supplies` | [AssetSupply](#joltify.issuance.v1beta1.AssetSupply) | repeated |  |






<a name="joltify.issuance.v1beta1.Params"></a>

### Params
Params defines the parameters for the issuance module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `assets` | [Asset](#joltify.issuance.v1beta1.Asset) | repeated |  |






<a name="joltify.issuance.v1beta1.RateLimit"></a>

### RateLimit
RateLimit parameters for rate-limiting the supply of an issued asset


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `active` | [bool](#bool) |  |  |
| `limit` | [bytes](#bytes) |  |  |
| `time_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="joltify/issuance/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/issuance/v1beta1/query.proto



<a name="joltify.issuance.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for querying x/issuance parameters.






<a name="joltify.issuance.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for querying x/issuance parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#joltify.issuance.v1beta1.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="joltify.issuance.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for issuance module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#joltify.issuance.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#joltify.issuance.v1beta1.QueryParamsResponse) | Params queries all parameters of the issuance module. | GET|/joltify/issuance/v1beta1/params|

 <!-- end services -->



<a name="joltify/issuance/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/issuance/v1beta1/tx.proto



<a name="joltify.issuance.v1beta1.MsgBlockAddress"></a>

### MsgBlockAddress
MsgBlockAddress represents a message used by the issuer to block an address from holding or transferring tokens


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denom` | [string](#string) |  |  |
| `blocked_address` | [string](#string) |  |  |






<a name="joltify.issuance.v1beta1.MsgBlockAddressResponse"></a>

### MsgBlockAddressResponse
MsgBlockAddressResponse defines the Msg/BlockAddress response type.






<a name="joltify.issuance.v1beta1.MsgIssueTokens"></a>

### MsgIssueTokens
MsgIssueTokens represents a message used by the issuer to issue new tokens


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `tokens` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `receiver` | [string](#string) |  |  |






<a name="joltify.issuance.v1beta1.MsgIssueTokensResponse"></a>

### MsgIssueTokensResponse
MsgIssueTokensResponse defines the Msg/IssueTokens response type.






<a name="joltify.issuance.v1beta1.MsgRedeemTokens"></a>

### MsgRedeemTokens
MsgRedeemTokens represents a message used by the issuer to redeem (burn) tokens


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `tokens` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="joltify.issuance.v1beta1.MsgRedeemTokensResponse"></a>

### MsgRedeemTokensResponse
MsgRedeemTokensResponse defines the Msg/RedeemTokens response type.






<a name="joltify.issuance.v1beta1.MsgSetPauseStatus"></a>

### MsgSetPauseStatus
MsgSetPauseStatus message type used by the issuer to pause or unpause status


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denom` | [string](#string) |  |  |
| `status` | [bool](#bool) |  |  |






<a name="joltify.issuance.v1beta1.MsgSetPauseStatusResponse"></a>

### MsgSetPauseStatusResponse
MsgSetPauseStatusResponse defines the Msg/SetPauseStatus response type.






<a name="joltify.issuance.v1beta1.MsgUnblockAddress"></a>

### MsgUnblockAddress
MsgUnblockAddress message type used by the issuer to unblock an address from holding or transferring tokens


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denom` | [string](#string) |  |  |
| `blocked_address` | [string](#string) |  |  |






<a name="joltify.issuance.v1beta1.MsgUnblockAddressResponse"></a>

### MsgUnblockAddressResponse
MsgUnblockAddressResponse defines the Msg/UnblockAddress response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="joltify.issuance.v1beta1.Msg"></a>

### Msg
Msg defines the issuance Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `IssueTokens` | [MsgIssueTokens](#joltify.issuance.v1beta1.MsgIssueTokens) | [MsgIssueTokensResponse](#joltify.issuance.v1beta1.MsgIssueTokensResponse) | IssueTokens message type used by the issuer to issue new tokens | |
| `RedeemTokens` | [MsgRedeemTokens](#joltify.issuance.v1beta1.MsgRedeemTokens) | [MsgRedeemTokensResponse](#joltify.issuance.v1beta1.MsgRedeemTokensResponse) | RedeemTokens message type used by the issuer to redeem (burn) tokens | |
| `BlockAddress` | [MsgBlockAddress](#joltify.issuance.v1beta1.MsgBlockAddress) | [MsgBlockAddressResponse](#joltify.issuance.v1beta1.MsgBlockAddressResponse) | BlockAddress message type used by the issuer to block an address from holding or transferring tokens | |
| `UnblockAddress` | [MsgUnblockAddress](#joltify.issuance.v1beta1.MsgUnblockAddress) | [MsgUnblockAddressResponse](#joltify.issuance.v1beta1.MsgUnblockAddressResponse) | UnblockAddress message type used by the issuer to unblock an address from holding or transferring tokens | |
| `SetPauseStatus` | [MsgSetPauseStatus](#joltify.issuance.v1beta1.MsgSetPauseStatus) | [MsgSetPauseStatusResponse](#joltify.issuance.v1beta1.MsgSetPauseStatusResponse) | SetPauseStatus message type used to pause or unpause status | |

 <!-- end services -->



<a name="joltify/jolt/v1beta1/jolt.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/jolt/v1beta1/jolt.proto



<a name="joltify.jolt.v1beta1.Borrow"></a>

### Borrow
Borrow defines an amount of coins borrowed from a jolt module account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `borrower` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `index` | [BorrowInterestFactor](#joltify.jolt.v1beta1.BorrowInterestFactor) | repeated |  |






<a name="joltify.jolt.v1beta1.BorrowInterestFactor"></a>

### BorrowInterestFactor
BorrowInterestFactor defines an individual borrow interest factor.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `value` | [string](#string) |  |  |






<a name="joltify.jolt.v1beta1.BorrowLimit"></a>

### BorrowLimit
BorrowLimit enforces restrictions on a money market.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `has_max_limit` | [bool](#bool) |  |  |
| `maximum_limit` | [string](#string) |  |  |
| `loan_to_value` | [string](#string) |  |  |






<a name="joltify.jolt.v1beta1.CoinsProto"></a>

### CoinsProto
CoinsProto defines a Protobuf wrapper around a Coins slice


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="joltify.jolt.v1beta1.Deposit"></a>

### Deposit
Deposit defines an amount of coins deposited into a jolt module account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `index` | [SupplyInterestFactor](#joltify.jolt.v1beta1.SupplyInterestFactor) | repeated |  |






<a name="joltify.jolt.v1beta1.InterestRateModel"></a>

### InterestRateModel
InterestRateModel contains information about an asset's interest rate.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_rate_apy` | [string](#string) |  |  |
| `base_multiplier` | [string](#string) |  |  |
| `kink` | [string](#string) |  |  |
| `jump_multiplier` | [string](#string) |  |  |






<a name="joltify.jolt.v1beta1.MoneyMarket"></a>

### MoneyMarket
MoneyMarket is a money market for an individual asset.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `borrow_limit` | [BorrowLimit](#joltify.jolt.v1beta1.BorrowLimit) |  |  |
| `spot_market_id` | [string](#string) |  |  |
| `conversion_factor` | [string](#string) |  |  |
| `interest_rate_model` | [InterestRateModel](#joltify.jolt.v1beta1.InterestRateModel) |  |  |
| `reserve_factor` | [string](#string) |  |  |
| `keeper_reward_percentage` | [string](#string) |  |  |






<a name="joltify.jolt.v1beta1.Params"></a>

### Params
Params defines the parameters for the jolt module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `money_markets` | [MoneyMarket](#joltify.jolt.v1beta1.MoneyMarket) | repeated |  |
| `minimum_borrow_usd_value` | [string](#string) |  |  |
| `surplus_auction_threshold` | [string](#string) |  |  |






<a name="joltify.jolt.v1beta1.SupplyInterestFactor"></a>

### SupplyInterestFactor
SupplyInterestFactor defines an individual borrow interest factor.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `value` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="joltify/jolt/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/jolt/v1beta1/genesis.proto



<a name="joltify.jolt.v1beta1.GenesisAccumulationTime"></a>

### GenesisAccumulationTime
GenesisAccumulationTime stores the previous distribution time and its corresponding denom.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `previous_accumulation_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `supply_interest_factor` | [string](#string) |  |  |
| `borrow_interest_factor` | [string](#string) |  |  |






<a name="joltify.jolt.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the jolt module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#joltify.jolt.v1beta1.Params) |  |  |
| `previous_accumulation_times` | [GenesisAccumulationTime](#joltify.jolt.v1beta1.GenesisAccumulationTime) | repeated |  |
| `deposits` | [Deposit](#joltify.jolt.v1beta1.Deposit) | repeated |  |
| `borrows` | [Borrow](#joltify.jolt.v1beta1.Borrow) | repeated |  |
| `total_supplied` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `total_borrowed` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `total_reserves` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="joltify/jolt/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/jolt/v1beta1/query.proto



<a name="joltify.jolt.v1beta1.BorrowInterestFactorResponse"></a>

### BorrowInterestFactorResponse
BorrowInterestFactorResponse defines an individual borrow interest factor.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `value` | [string](#string) |  | sdk.Dec as string |






<a name="joltify.jolt.v1beta1.BorrowResponse"></a>

### BorrowResponse
BorrowResponse defines an amount of coins borrowed from a jolt module account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `borrower` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `index` | [BorrowInterestFactorResponse](#joltify.jolt.v1beta1.BorrowInterestFactorResponse) | repeated |  |






<a name="joltify.jolt.v1beta1.DepositResponse"></a>

### DepositResponse
DepositResponse defines an amount of coins deposited into a jolt module account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `index` | [SupplyInterestFactorResponse](#joltify.jolt.v1beta1.SupplyInterestFactorResponse) | repeated |  |






<a name="joltify.jolt.v1beta1.InterestFactor"></a>

### InterestFactor
InterestFactor is a unique type returned by interest factor queries


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `borrow_interest_factor` | [string](#string) |  | sdk.Dec as String |
| `supply_interest_factor` | [string](#string) |  | sdk.Dec as String |






<a name="joltify.jolt.v1beta1.LiquidateItem"></a>

### LiquidateItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `ltv` | [string](#string) |  |  |






<a name="joltify.jolt.v1beta1.MoneyMarketInterestRate"></a>

### MoneyMarketInterestRate
MoneyMarketInterestRate is a unique type returned by interest rate queries


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `supply_interest_rate` | [string](#string) |  | sdk.Dec as String |
| `borrow_interest_rate` | [string](#string) |  | sdk.Dec as String |






<a name="joltify.jolt.v1beta1.QueryAccountsRequest"></a>

### QueryAccountsRequest
QueryAccountsRequest is the request type for the Query/Accounts RPC method.






<a name="joltify.jolt.v1beta1.QueryAccountsResponse"></a>

### QueryAccountsResponse
QueryAccountsResponse is the response type for the Query/Accounts RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `accounts` | [cosmos.auth.v1beta1.ModuleAccount](#cosmos.auth.v1beta1.ModuleAccount) | repeated |  |






<a name="joltify.jolt.v1beta1.QueryBorrowsRequest"></a>

### QueryBorrowsRequest
QueryBorrowsRequest is the request type for the Query/Borrows RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="joltify.jolt.v1beta1.QueryBorrowsResponse"></a>

### QueryBorrowsResponse
QueryBorrowsResponse is the response type for the Query/Borrows RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `borrows` | [BorrowResponse](#joltify.jolt.v1beta1.BorrowResponse) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="joltify.jolt.v1beta1.QueryDepositsRequest"></a>

### QueryDepositsRequest
QueryDepositsRequest is the request type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="joltify.jolt.v1beta1.QueryDepositsResponse"></a>

### QueryDepositsResponse
QueryDepositsResponse is the response type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposits` | [DepositResponse](#joltify.jolt.v1beta1.DepositResponse) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="joltify.jolt.v1beta1.QueryInterestFactorsRequest"></a>

### QueryInterestFactorsRequest
QueryInterestFactorsRequest is the request type for the Query/InterestFactors RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |






<a name="joltify.jolt.v1beta1.QueryInterestFactorsResponse"></a>

### QueryInterestFactorsResponse
QueryInterestFactorsResponse is the response type for the Query/InterestFactors RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `interest_factors` | [InterestFactor](#joltify.jolt.v1beta1.InterestFactor) | repeated |  |






<a name="joltify.jolt.v1beta1.QueryInterestRateRequest"></a>

### QueryInterestRateRequest
QueryInterestRateRequest is the request type for the Query/InterestRate RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |






<a name="joltify.jolt.v1beta1.QueryInterestRateResponse"></a>

### QueryInterestRateResponse
QueryInterestRateResponse is the response type for the Query/InterestRate RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `interest_rates` | [MoneyMarketInterestRate](#joltify.jolt.v1beta1.MoneyMarketInterestRate) | repeated |  |






<a name="joltify.jolt.v1beta1.QueryLiquidateRequest"></a>

### QueryLiquidateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `borrower` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="joltify.jolt.v1beta1.QueryLiquidateResp"></a>

### QueryLiquidateResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `liquidateItems` | [LiquidateItem](#joltify.jolt.v1beta1.LiquidateItem) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="joltify.jolt.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="joltify.jolt.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#joltify.jolt.v1beta1.Params) |  |  |






<a name="joltify.jolt.v1beta1.QueryReservesRequest"></a>

### QueryReservesRequest
QueryReservesRequest is the request type for the Query/Reserves RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |






<a name="joltify.jolt.v1beta1.QueryReservesResponse"></a>

### QueryReservesResponse
QueryReservesResponse is the response type for the Query/Reserves RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="joltify.jolt.v1beta1.QueryTotalBorrowedRequest"></a>

### QueryTotalBorrowedRequest
QueryTotalBorrowedRequest is the request type for the Query/TotalBorrowed RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |






<a name="joltify.jolt.v1beta1.QueryTotalBorrowedResponse"></a>

### QueryTotalBorrowedResponse
QueryTotalBorrowedResponse is the response type for the Query/TotalBorrowed RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `borrowed_coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="joltify.jolt.v1beta1.QueryTotalDepositedRequest"></a>

### QueryTotalDepositedRequest
QueryTotalDepositedRequest is the request type for the Query/TotalDeposited RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |






<a name="joltify.jolt.v1beta1.QueryTotalDepositedResponse"></a>

### QueryTotalDepositedResponse
QueryTotalDepositedResponse is the response type for the Query/TotalDeposited RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `supplied_coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="joltify.jolt.v1beta1.QueryUnsyncedBorrowsRequest"></a>

### QueryUnsyncedBorrowsRequest
QueryUnsyncedBorrowsRequest is the request type for the Query/UnsyncedBorrows RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="joltify.jolt.v1beta1.QueryUnsyncedBorrowsResponse"></a>

### QueryUnsyncedBorrowsResponse
QueryUnsyncedBorrowsResponse is the response type for the Query/UnsyncedBorrows RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `borrows` | [BorrowResponse](#joltify.jolt.v1beta1.BorrowResponse) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="joltify.jolt.v1beta1.QueryUnsyncedDepositsRequest"></a>

### QueryUnsyncedDepositsRequest
QueryUnsyncedDepositsRequest is the request type for the Query/UnsyncedDeposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="joltify.jolt.v1beta1.QueryUnsyncedDepositsResponse"></a>

### QueryUnsyncedDepositsResponse
QueryUnsyncedDepositsResponse is the response type for the Query/UnsyncedDeposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposits` | [DepositResponse](#joltify.jolt.v1beta1.DepositResponse) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="joltify.jolt.v1beta1.SupplyInterestFactorResponse"></a>

### SupplyInterestFactorResponse
SupplyInterestFactorResponse defines an individual borrow interest factor.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `value` | [string](#string) |  | sdk.Dec as string |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="joltify.jolt.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for bep3 module.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#joltify.jolt.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#joltify.jolt.v1beta1.QueryParamsResponse) | Params queries module params. | GET|/joltify/jolt/v1beta1/params|
| `Accounts` | [QueryAccountsRequest](#joltify.jolt.v1beta1.QueryAccountsRequest) | [QueryAccountsResponse](#joltify.jolt.v1beta1.QueryAccountsResponse) | Accounts queries module accounts. | GET|/joltify/jolt/v1beta1/accounts|
| `Deposits` | [QueryDepositsRequest](#joltify.jolt.v1beta1.QueryDepositsRequest) | [QueryDepositsResponse](#joltify.jolt.v1beta1.QueryDepositsResponse) | Deposits queries jolt deposits. | GET|/joltify/jolt/v1beta1/deposits|
| `UnsyncedDeposits` | [QueryUnsyncedDepositsRequest](#joltify.jolt.v1beta1.QueryUnsyncedDepositsRequest) | [QueryUnsyncedDepositsResponse](#joltify.jolt.v1beta1.QueryUnsyncedDepositsResponse) | UnsyncedDeposits queries unsynced deposits. | GET|/joltify/jolt/v1beta1/unsynced-deposits|
| `TotalDeposited` | [QueryTotalDepositedRequest](#joltify.jolt.v1beta1.QueryTotalDepositedRequest) | [QueryTotalDepositedResponse](#joltify.jolt.v1beta1.QueryTotalDepositedResponse) | TotalDeposited queries total coins deposited to jolt liquidity pools. | GET|/joltify/jolt/v1beta1/total-deposited/{denom}|
| `Borrows` | [QueryBorrowsRequest](#joltify.jolt.v1beta1.QueryBorrowsRequest) | [QueryBorrowsResponse](#joltify.jolt.v1beta1.QueryBorrowsResponse) | Borrows queries jolt borrows. | GET|/joltify/jolt/v1beta1/borrows|
| `UnsyncedBorrows` | [QueryUnsyncedBorrowsRequest](#joltify.jolt.v1beta1.QueryUnsyncedBorrowsRequest) | [QueryUnsyncedBorrowsResponse](#joltify.jolt.v1beta1.QueryUnsyncedBorrowsResponse) | UnsyncedBorrows queries unsynced borrows. | GET|/joltify/jolt/v1beta1/unsynced-borrows|
| `TotalBorrowed` | [QueryTotalBorrowedRequest](#joltify.jolt.v1beta1.QueryTotalBorrowedRequest) | [QueryTotalBorrowedResponse](#joltify.jolt.v1beta1.QueryTotalBorrowedResponse) | TotalBorrowed queries total coins borrowed from jolt liquidity pools. | GET|/joltify/jolt/v1beta1/total-borrowed/{denom}|
| `InterestRate` | [QueryInterestRateRequest](#joltify.jolt.v1beta1.QueryInterestRateRequest) | [QueryInterestRateResponse](#joltify.jolt.v1beta1.QueryInterestRateResponse) | InterestRate queries the jolt module interest rates. | GET|/joltify/jolt/v1beta1/interest-rate/{denom}|
| `Reserves` | [QueryReservesRequest](#joltify.jolt.v1beta1.QueryReservesRequest) | [QueryReservesResponse](#joltify.jolt.v1beta1.QueryReservesResponse) | Reserves queries total jolt reserve coins. | GET|/joltify/jolt/v1beta1/reserves/{denom}|
| `InterestFactors` | [QueryInterestFactorsRequest](#joltify.jolt.v1beta1.QueryInterestFactorsRequest) | [QueryInterestFactorsResponse](#joltify.jolt.v1beta1.QueryInterestFactorsResponse) | InterestFactors queries jolt module interest factors. | GET|/joltify/jolt/v1beta1/interest-factors/{denom}|
| `liquidate` | [QueryLiquidateRequest](#joltify.jolt.v1beta1.QueryLiquidateRequest) | [QueryLiquidateResp](#joltify.jolt.v1beta1.QueryLiquidateResp) | queries jolt module interest factors. | GET|/joltify/jolt/v1beta1/liquidate|

 <!-- end services -->



<a name="joltify/jolt/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/jolt/v1beta1/tx.proto



<a name="joltify.jolt.v1beta1.MsgBorrow"></a>

### MsgBorrow
MsgBorrow defines the Msg/Borrow request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `borrower` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="joltify.jolt.v1beta1.MsgBorrowResponse"></a>

### MsgBorrowResponse
MsgBorrowResponse defines the Msg/Borrow response type.






<a name="joltify.jolt.v1beta1.MsgDeposit"></a>

### MsgDeposit
MsgDeposit defines the Msg/Deposit request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="joltify.jolt.v1beta1.MsgDepositResponse"></a>

### MsgDepositResponse
MsgDepositResponse defines the Msg/Deposit response type.






<a name="joltify.jolt.v1beta1.MsgLiquidate"></a>

### MsgLiquidate
MsgLiquidate defines the Msg/Liquidate request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `keeper` | [string](#string) |  |  |
| `borrower` | [string](#string) |  |  |






<a name="joltify.jolt.v1beta1.MsgLiquidateResponse"></a>

### MsgLiquidateResponse
MsgLiquidateResponse defines the Msg/Liquidate response type.






<a name="joltify.jolt.v1beta1.MsgRepay"></a>

### MsgRepay
MsgRepay defines the Msg/Repay request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="joltify.jolt.v1beta1.MsgRepayResponse"></a>

### MsgRepayResponse
MsgRepayResponse defines the Msg/Repay response type.






<a name="joltify.jolt.v1beta1.MsgWithdraw"></a>

### MsgWithdraw
MsgWithdraw defines the Msg/Withdraw request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="joltify.jolt.v1beta1.MsgWithdrawResponse"></a>

### MsgWithdrawResponse
MsgWithdrawResponse defines the Msg/Withdraw response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="joltify.jolt.v1beta1.Msg"></a>

### Msg
Msg defines the jolt Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Deposit` | [MsgDeposit](#joltify.jolt.v1beta1.MsgDeposit) | [MsgDepositResponse](#joltify.jolt.v1beta1.MsgDepositResponse) | Deposit defines a method for depositing funds to jolt liquidity pool. | |
| `Withdraw` | [MsgWithdraw](#joltify.jolt.v1beta1.MsgWithdraw) | [MsgWithdrawResponse](#joltify.jolt.v1beta1.MsgWithdrawResponse) | Withdraw defines a method for withdrawing funds from jolt liquidity pool. | |
| `Borrow` | [MsgBorrow](#joltify.jolt.v1beta1.MsgBorrow) | [MsgBorrowResponse](#joltify.jolt.v1beta1.MsgBorrowResponse) | Borrow defines a method for borrowing funds from jolt liquidity pool. | |
| `Repay` | [MsgRepay](#joltify.jolt.v1beta1.MsgRepay) | [MsgRepayResponse](#joltify.jolt.v1beta1.MsgRepayResponse) | Repay defines a method for repaying funds borrowed from jolt liquidity pool. | |
| `Liquidate` | [MsgLiquidate](#joltify.jolt.v1beta1.MsgLiquidate) | [MsgLiquidateResponse](#joltify.jolt.v1beta1.MsgLiquidateResponse) | Liquidate defines a method for attempting to liquidate a borrower that is over their loan-to-value. | |

 <!-- end services -->



<a name="joltify/mint/dist.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/mint/dist.proto



<a name="joltify.mint.HistoricalDistInfo"></a>

### HistoricalDistInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `payout_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `distributed_round` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="joltify/mint/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/mint/params.proto



<a name="joltify.mint.Params"></a>

### Params
Params defines the parameters for the module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `totalCount` | [uint64](#uint64) |  |  |
| `each_provisions` | [string](#string) |  |  |
| `unit` | [string](#string) |  |  |
| `community_provisions` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="joltify/mint/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/mint/genesis.proto



<a name="joltify.mint.GenesisState"></a>

### GenesisState
GenesisState defines the joltmint module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#joltify.mint.Params) |  |  |
| `historical_dist_info` | [HistoricalDistInfo](#joltify.mint.HistoricalDistInfo) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="joltify/mint/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/mint/query.proto



<a name="joltify.mint.QueryHistoryDistRequest"></a>

### QueryHistoryDistRequest
QueryParamsRequest is request type for the Query/Params RPC method.






<a name="joltify.mint.QueryHistoryDistResponse"></a>

### QueryHistoryDistResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `history` | [HistoricalDistInfo](#joltify.mint.HistoricalDistInfo) |  | params holds all the parameters of this module. |






<a name="joltify.mint.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC method.






<a name="joltify.mint.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#joltify.mint.Params) |  | params holds all the parameters of this module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="joltify.mint.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#joltify.mint.QueryParamsRequest) | [QueryParamsResponse](#joltify.mint.QueryParamsResponse) | Parameters queries the parameters of the module. | GET|/joltify-finance/joltify_lending/mint/params|
| `Distribution` | [QueryHistoryDistRequest](#joltify.mint.QueryHistoryDistRequest) | [QueryHistoryDistResponse](#joltify.mint.QueryHistoryDistResponse) | Distribution queries the parameters of the module. | GET|/joltify-finance/joltify_lending/mint/dist|

 <!-- end services -->



<a name="joltify/mint/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/mint/tx.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="joltify.mint.Msg"></a>

### Msg
Msg defines the Msg service.

this line is used by starport scaffolding # proto/tx/rpc

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |

 <!-- end services -->



<a name="joltify/pricefeed/v1beta1/store.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/pricefeed/v1beta1/store.proto



<a name="joltify.pricefeed.v1beta1.CurrentPrice"></a>

### CurrentPrice
CurrentPrice defines a current price for a particular market in the pricefeed
module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |






<a name="joltify.pricefeed.v1beta1.Market"></a>

### Market
Market defines an asset in the pricefeed.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `base_asset` | [string](#string) |  |  |
| `quote_asset` | [string](#string) |  |  |
| `oracles` | [bytes](#bytes) | repeated |  |
| `active` | [bool](#bool) |  |  |






<a name="joltify.pricefeed.v1beta1.Params"></a>

### Params
Params defines the parameters for the pricefeed module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `markets` | [Market](#joltify.pricefeed.v1beta1.Market) | repeated |  |






<a name="joltify.pricefeed.v1beta1.PostedPrice"></a>

### PostedPrice
PostedPrice defines a price for market posted by a specific oracle.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `oracle_address` | [bytes](#bytes) |  |  |
| `price` | [string](#string) |  |  |
| `expiry` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="joltify/pricefeed/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/pricefeed/v1beta1/genesis.proto



<a name="joltify.pricefeed.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the pricefeed module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#joltify.pricefeed.v1beta1.Params) |  | params defines all the paramaters of the module. |
| `posted_prices` | [PostedPrice](#joltify.pricefeed.v1beta1.PostedPrice) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="joltify/pricefeed/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/pricefeed/v1beta1/query.proto



<a name="joltify.pricefeed.v1beta1.CurrentPriceResponse"></a>

### CurrentPriceResponse
CurrentPriceResponse defines a current price for a particular market in the pricefeed
module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |






<a name="joltify.pricefeed.v1beta1.MarketResponse"></a>

### MarketResponse
MarketResponse defines an asset in the pricefeed.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `base_asset` | [string](#string) |  |  |
| `quote_asset` | [string](#string) |  |  |
| `oracles` | [string](#string) | repeated |  |
| `active` | [bool](#bool) |  |  |






<a name="joltify.pricefeed.v1beta1.PostedPriceResponse"></a>

### PostedPriceResponse
PostedPriceResponse defines a price for market posted by a specific oracle.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `oracle_address` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |
| `expiry` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="joltify.pricefeed.v1beta1.QueryMarketsRequest"></a>

### QueryMarketsRequest
QueryMarketsRequest is the request type for the Query/Markets RPC method.






<a name="joltify.pricefeed.v1beta1.QueryMarketsResponse"></a>

### QueryMarketsResponse
QueryMarketsResponse is the response type for the Query/Markets RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `markets` | [MarketResponse](#joltify.pricefeed.v1beta1.MarketResponse) | repeated | List of markets |






<a name="joltify.pricefeed.v1beta1.QueryOraclesRequest"></a>

### QueryOraclesRequest
QueryOraclesRequest is the request type for the Query/Oracles RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |






<a name="joltify.pricefeed.v1beta1.QueryOraclesResponse"></a>

### QueryOraclesResponse
QueryOraclesResponse is the response type for the Query/Oracles RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `oracles` | [string](#string) | repeated | List of oracle addresses |






<a name="joltify.pricefeed.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for querying x/pricefeed
parameters.






<a name="joltify.pricefeed.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for querying x/pricefeed
parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#joltify.pricefeed.v1beta1.Params) |  |  |






<a name="joltify.pricefeed.v1beta1.QueryPriceRequest"></a>

### QueryPriceRequest
QueryPriceRequest is the request type for the Query/PriceRequest RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |






<a name="joltify.pricefeed.v1beta1.QueryPriceResponse"></a>

### QueryPriceResponse
QueryPriceResponse is the response type for the Query/Prices RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `price` | [CurrentPriceResponse](#joltify.pricefeed.v1beta1.CurrentPriceResponse) |  |  |






<a name="joltify.pricefeed.v1beta1.QueryPricesRequest"></a>

### QueryPricesRequest
QueryPricesRequest is the request type for the Query/Prices RPC method.






<a name="joltify.pricefeed.v1beta1.QueryPricesResponse"></a>

### QueryPricesResponse
QueryPricesResponse is the response type for the Query/Prices RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `prices` | [CurrentPriceResponse](#joltify.pricefeed.v1beta1.CurrentPriceResponse) | repeated |  |






<a name="joltify.pricefeed.v1beta1.QueryRawPricesRequest"></a>

### QueryRawPricesRequest
QueryRawPricesRequest is the request type for the Query/RawPrices RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |






<a name="joltify.pricefeed.v1beta1.QueryRawPricesResponse"></a>

### QueryRawPricesResponse
QueryRawPricesResponse is the response type for the Query/RawPrices RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `raw_prices` | [PostedPriceResponse](#joltify.pricefeed.v1beta1.PostedPriceResponse) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="joltify.pricefeed.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for pricefeed module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#joltify.pricefeed.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#joltify.pricefeed.v1beta1.QueryParamsResponse) | Params queries all parameters of the pricefeed module. | GET|/joltify/pricefeed/v1beta1/params|
| `Price` | [QueryPriceRequest](#joltify.pricefeed.v1beta1.QueryPriceRequest) | [QueryPriceResponse](#joltify.pricefeed.v1beta1.QueryPriceResponse) | Price queries price details based on a market | GET|/joltify/pricefeed/v1beta1/prices/{market_id}|
| `Prices` | [QueryPricesRequest](#joltify.pricefeed.v1beta1.QueryPricesRequest) | [QueryPricesResponse](#joltify.pricefeed.v1beta1.QueryPricesResponse) | Prices queries all prices | GET|/joltify/pricefeed/v1beta1/prices|
| `RawPrices` | [QueryRawPricesRequest](#joltify.pricefeed.v1beta1.QueryRawPricesRequest) | [QueryRawPricesResponse](#joltify.pricefeed.v1beta1.QueryRawPricesResponse) | RawPrices queries all raw prices based on a market | GET|/joltify/pricefeed/v1beta1/rawprices/{market_id}|
| `Oracles` | [QueryOraclesRequest](#joltify.pricefeed.v1beta1.QueryOraclesRequest) | [QueryOraclesResponse](#joltify.pricefeed.v1beta1.QueryOraclesResponse) | Oracles queries all oracles based on a market | GET|/joltify/pricefeed/v1beta1/oracles/{market_id}|
| `Markets` | [QueryMarketsRequest](#joltify.pricefeed.v1beta1.QueryMarketsRequest) | [QueryMarketsResponse](#joltify.pricefeed.v1beta1.QueryMarketsResponse) | Markets queries all markets | GET|/joltify/pricefeed/v1beta1/markets|

 <!-- end services -->



<a name="joltify/pricefeed/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## joltify/pricefeed/v1beta1/tx.proto



<a name="joltify.pricefeed.v1beta1.MsgPostPrice"></a>

### MsgPostPrice
MsgPostPrice represents a method for creating a new post price


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from` | [string](#string) |  | address of client |
| `market_id` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |
| `expiry` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="joltify.pricefeed.v1beta1.MsgPostPriceResponse"></a>

### MsgPostPriceResponse
MsgPostPriceResponse defines the Msg/PostPrice response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="joltify.pricefeed.v1beta1.Msg"></a>

### Msg
Msg defines the pricefeed Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `PostPrice` | [MsgPostPrice](#joltify.pricefeed.v1beta1.MsgPostPrice) | [MsgPostPriceResponse](#joltify.pricefeed.v1beta1.MsgPostPriceResponse) | PostPrice defines a method for creating a new post price | |

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

