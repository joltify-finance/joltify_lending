package keeper

import (
	"context"
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListAllProjects(goCtx context.Context, req *types.ListAllProjectsRequest) (*types.ListAllProjectsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	projectNum := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectInfoNum))
	data := projectNum.Get(types.KeyPrefix(types.ProjectInfoNum))
	currentNum := binary.BigEndian.Uint32(data)

	store := ctx.KVStore(k.storeKey)
	projectStore := prefix.NewStore(store, types.KeyPrefix(types.ProjectInfoPrefix))

	var allProjects []*types.ProjectInfo
	pageRes, err := query.Paginate(projectStore, req.Pagination, func(key []byte, value []byte) error {
		var investor types.ProjectInfo
		if err := k.cdc.Unmarshal(value, &investor); err != nil {
			return err
		}
		allProjects = append(allProjects, &investor)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.ListAllProjectsResponse{Project: allProjects, TotalNumber: int32(currentNum), Pagination: pageRes}, nil
}

func (k Keeper) QueryProject(goCtx context.Context, req *types.QueryProjectRequest) (*types.QueryProjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	project, ok := k.GetProject(ctx, req.ProjectId)
	if !ok {
		return nil, status.Error(codes.NotFound, "project not found")
	}

	return &types.QueryProjectResponse{Project: &project}, nil
}
