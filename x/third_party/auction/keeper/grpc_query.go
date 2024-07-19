package keeper

import (
	"context"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/auction/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cosmossdk.io/store/prefix"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/gogo/protobuf/proto"
)

type queryServer struct {
	keeper Keeper
}

// NewQueryServerImpl creates a new server for handling gRPC queries.
func NewQueryServerImpl(k Keeper) types2.QueryServer {
	return &queryServer{keeper: k}
}

var _ types2.QueryServer = queryServer{}

// Params implements the gRPC service handler for querying x/auction parameters.
func (s queryServer) Params(ctx context.Context, req *types2.QueryParamsRequest) (*types2.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	params := s.keeper.GetParams(sdkCtx)

	return &types2.QueryParamsResponse{Params: params}, nil
}

// Auction implements the Query/Auction gRPC method
func (s queryServer) Auction(c context.Context, req *types2.QueryAuctionRequest) (*types2.QueryAuctionResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	auction, found := s.keeper.GetAuction(ctx, req.AuctionId)
	if !found {
		return &types2.QueryAuctionResponse{}, nil
	}

	msg, ok := auction.(proto.Message)
	if !ok {
		return nil, status.Errorf(codes.Internal, "can't protomarshal %T", msg)
	}

	auctionAny, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &types2.QueryAuctionResponse{
		Auction: auctionAny,
	}, nil
}

// Auctions implements the Query/Auctions gRPC method
func (s queryServer) Auctions(c context.Context, req *types2.QueryAuctionsRequest) (*types2.QueryAuctionsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	var auctions []*codectypes.Any
	auctionStore := prefix.NewStore(ctx.KVStore(s.keeper.storeKey), types2.AuctionKeyPrefix)

	pageRes, err := query.FilteredPaginate(auctionStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		result, err := s.keeper.UnmarshalAuction(value)
		if err != nil {
			return false, err
		}

		// True if empty owner, otherwise check if auction contains owner
		ownerIsMatch := req.Owner == ""
		if req.Owner != "" {
			if cAuc, ok := result.(*types2.CollateralAuction); ok {
				for _, addr := range cAuc.GetLotReturns().Addresses {
					if addr.String() == req.Owner {
						ownerIsMatch = true
						break
					}
				}
			}
		}

		phaseIsMatch := req.Phase == "" || req.Phase == result.GetPhase()
		typeIsMatch := req.Type == "" || req.Type == result.GetType()
		denomIsMatch := req.Denom == "" || req.Denom == result.GetBid().Denom || req.Denom == result.GetLot().Denom

		if ownerIsMatch && phaseIsMatch && typeIsMatch && denomIsMatch {
			if accumulate {
				msg, ok := result.(proto.Message)
				if !ok {
					return false, status.Errorf(codes.Internal, "can't protomarshal %T", msg)
				}

				auctionAny, err := codectypes.NewAnyWithValue(msg)
				if err != nil {
					return false, err
				}
				auctions = append(auctions, auctionAny)
			}

			return true, nil
		}

		return false, nil
	})
	if err != nil {
		return &types2.QueryAuctionsResponse{}, err
	}

	return &types2.QueryAuctionsResponse{
		Auctions:   auctions,
		Pagination: pageRes,
	}, nil
}

// NextAuctionID implements the gRPC service handler for querying x/auction next auction ID.
func (s queryServer) NextAuctionID(ctx context.Context, req *types2.QueryNextAuctionIDRequest) (*types2.QueryNextAuctionIDResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	nextAuctionID, err := s.keeper.GetNextAuctionID(sdkCtx)
	if err != nil {
		return &types2.QueryNextAuctionIDResponse{}, err
	}

	return &types2.QueryNextAuctionIDResponse{Id: nextAuctionID}, nil
}
