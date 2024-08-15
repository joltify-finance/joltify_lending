package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/joltify-finance/joltify_lending/lib"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ClobPairAll(
	c context.Context,
	req *types.QueryAllClobPairRequest,
) (*types.QueryClobPairAllResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var clobPairs []types.ClobPair
	ctx := lib.UnwrapSDKContext(c, types.ModuleName)

	store := ctx.KVStore(k.storeKey)
	clobPairStore := prefix.NewStore(store, []byte(types.ClobPairKeyPrefix))

	pageRes, err := query.Paginate(clobPairStore, req.Pagination, func(key []byte, value []byte) error {
		var clobPair types.ClobPair
		if err := k.cdc.Unmarshal(value, &clobPair); err != nil {
			return err
		}

		clobPairs = append(clobPairs, clobPair)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryClobPairAllResponse{ClobPair: clobPairs, Pagination: pageRes}, nil
}

func (k Keeper) ClobPair(c context.Context, req *types.QueryGetClobPairRequest) (*types.QueryClobPairResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := lib.UnwrapSDKContext(c, types.ModuleName)

	val, found := k.GetClobPair(
		ctx,
		types.ClobPairId(req.Id),
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryClobPairResponse{ClobPair: val}, nil
}
