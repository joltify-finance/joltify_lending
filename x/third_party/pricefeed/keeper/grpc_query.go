package keeper

import (
	"context"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type queryServer struct {
	keeper Keeper
}

// NewQueryServerImpl creates a new server for handling gRPC queries.
func NewQueryServerImpl(k Keeper) types2.QueryServer {
	return &queryServer{keeper: k}
}

var _ types2.QueryServer = queryServer{}

// Params implements the gRPC service handler for querying x/pricefeed parameters.
func (s queryServer) Params(c context.Context, req *types2.QueryParamsRequest) (*types2.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	sdkCtx := sdk.UnwrapSDKContext(c)
	params := s.keeper.GetParams(sdkCtx)

	return &types2.QueryParamsResponse{Params: params}, nil
}

// Price implements the gRPC service handler for querying x/pricefeed price.
func (s queryServer) Price(c context.Context, req *types2.QueryPriceRequest) (*types2.QueryPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	_, found := s.keeper.GetMarket(ctx, req.MarketId)
	if !found {
		return nil, status.Error(codes.NotFound, "invalid market ID")
	}
	currentPrice, sdkErr := s.keeper.GetCurrentPrice(ctx, req.MarketId)
	if sdkErr != nil {
		return nil, sdkErr
	}

	return &types2.QueryPriceResponse{
		Price: types2.CurrentPriceResponse(currentPrice),
	}, nil
}

func (s queryServer) Prices(c context.Context, req *types2.QueryPricesRequest) (*types2.QueryPricesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	var currentPrices types2.CurrentPriceResponses
	for _, cp := range s.keeper.GetCurrentPrices(ctx) {
		if cp.MarketID != "" {
			currentPrices = append(currentPrices, types2.CurrentPriceResponse(cp))
		}
	}

	return &types2.QueryPricesResponse{
		Prices: currentPrices,
	}, nil
}

func (s queryServer) RawPrices(c context.Context, req *types2.QueryRawPricesRequest) (*types2.QueryRawPricesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	_, found := s.keeper.GetMarket(ctx, req.MarketId)
	if !found {
		return nil, status.Error(codes.NotFound, "invalid market ID")
	}

	var prices types2.PostedPriceResponses
	for _, rp := range s.keeper.GetRawPrices(ctx, req.MarketId) {
		prices = append(prices, types2.PostedPriceResponse{
			MarketID:      rp.MarketID,
			OracleAddress: rp.OracleAddress.String(),
			Price:         rp.Price,
			Expiry:        rp.Expiry,
		})
	}

	return &types2.QueryRawPricesResponse{
		RawPrices: prices,
	}, nil
}

func (s queryServer) Oracles(c context.Context, req *types2.QueryOraclesRequest) (*types2.QueryOraclesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	oracles, err := s.keeper.GetOracles(ctx, req.MarketId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "invalid market ID")
	}

	var strOracles []string
	for _, oracle := range oracles {
		strOracles = append(strOracles, oracle.String())
	}

	return &types2.QueryOraclesResponse{
		Oracles: strOracles,
	}, nil
}

func (s queryServer) Markets(c context.Context, req *types2.QueryMarketsRequest) (*types2.QueryMarketsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	var markets types2.MarketResponses
	for _, market := range s.keeper.GetMarkets(ctx) {
		markets = append(markets, market.ToMarketResponse())
	}

	return &types2.QueryMarketsResponse{
		Markets: markets,
	}, nil
}
