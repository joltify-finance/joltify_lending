package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/joltify-finance/joltify_lending/x/quota/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QuotaAll(goCtx context.Context, req *types.QueryAllQuotaRequest) (*types.QueryAllQuotaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var quotas []types.CoinsQuota
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	quotaStore := prefix.NewStore(store, types.KeyPrefix(types.QuotaKey))

	pageRes, err := query.Paginate(quotaStore, req.Pagination, func(key []byte, value []byte) error {
		var quota types.CoinsQuota
		if err := k.cdc.Unmarshal(value, &quota); err != nil {
			return err
		}

		quotas = append(quotas, quota)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllQuotaResponse{Quota: quotas, Pagination: pageRes}, nil
}

func (k Keeper) Quota(goCtx context.Context, req *types.QueryGetQuotaRequest) (*types.QueryGetQuotaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	quota, found := k.GetQuotaData(ctx, req.QuotaModuleName)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetQuotaResponse{Quota: quota}, nil
}

func (k Keeper) AccountQuota(goCtx context.Context, req *types.QueryGetAccountQuotaRequest) (*types.QueryGetAccountQuotaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	quota, found := k.getAccountQuotaData(ctx, req.QuotaModuleName, req.GetAccountAddress())
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}
	return &types.QueryGetAccountQuotaResponse{Quota: quota}, nil
}
