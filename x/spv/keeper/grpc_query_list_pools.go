package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListPools(goCtx context.Context, req *types.QueryListPoolsRequest) (*types.QueryListPoolsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	investorStores := prefix.NewStore(store, types.KeyPrefix(types.Pool))

	var poolsInfos []*types.PoolInfo

	pageRes, err := query.Paginate(investorStores, req.Pagination, func(key []byte, value []byte) error {
		var investor types.PoolInfo
		if err := k.cdc.Unmarshal(value, &investor); err != nil {
			return err
		}
		poolsInfos = append(poolsInfos, &investor)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListPoolsResponse{PoolsInfo: poolsInfos, Pagination: pageRes}, nil
}

func (k Keeper) ListHistoryPools(goCtx context.Context, req *types.QueryListPoolsRequest) (*types.QueryListPoolsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	investorStores := prefix.NewStore(store, types.KeyPrefix(types.HistoryPool))

	var poolsInfos []*types.PoolInfo

	pageRes, err := query.Paginate(investorStores, req.Pagination, func(key []byte, value []byte) error {
		var poolInfo types.PoolInfo
		if err := k.cdc.Unmarshal(value, &poolInfo); err != nil {
			return err
		}

		poolsInfos = append(poolsInfos, &poolInfo)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListPoolsResponse{PoolsInfo: poolsInfos, Pagination: pageRes}, nil
}
