package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetPendingFee(c context.Context, req *types.QueryPendingFeeRequest) (*types.QueryPendingFeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	fees := k.GetAllFeeAmount(ctx)
	return &types.QueryPendingFeeResponse{Feecoin: fees}, nil
}
